package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"reflect"
	"time"

	"connectrpc.com/connect"
	"github.com/LayerTwo-Labs/sidesail/drivechain-server/bdk"
	rpc "github.com/LayerTwo-Labs/sidesail/drivechain-server/gen/drivechain/v1/drivechainv1connect"
	"github.com/LayerTwo-Labs/sidesail/drivechain-server/server"
	pb "github.com/barebitcoin/btc-buf/gen/bitcoin/bitcoind/v1alpha"
	coreproxy "github.com/barebitcoin/btc-buf/server"
	"github.com/jessevdk/go-flags"
	"github.com/rs/zerolog"
	"github.com/samber/lo"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	logger := zerolog.
		New(zerolog.NewConsoleWriter()).
		Level(zerolog.TraceLevel)

	zerolog.DefaultContextLogger = &logger

	if err := realMain(ctx); err != nil {
		cancel()

		// Error has been printed to the console!
		if _, ok := lo.ErrorsAs[*flags.Error](err); ok {
			os.Exit(1)
		}

		zerolog.Ctx(ctx).
			Fatal().
			Err(err).
			Msgf("main: got error: %T", err)
	}
}

func realMain(ctx context.Context) error {
	conf, err := readConfig()
	if err != nil {
		return err
	}

	proxy, err := startCoreProxy(ctx, conf)
	if err != nil {
		return err
	}

	info, err := proxy.GetBlockchainInfo(ctx, connect.NewRequest(&pb.GetBlockchainInfoRequest{}))
	if err != nil {
		return err
	}

	zerolog.Ctx(ctx).Info().Msgf("blockchain info: %s", info.Msg.String())

	// Can be generated by:
	// $ bdk-cli key generate | jq -r .xprv
	priv := "xprv9s21ZrQH143K4KPGEHUNPYep8dztoaJm1MUhMAx5mjh4Y8iGVNe32VwReJZ3rR1btgB82rbrLyRnpmeVWqshbNVumSdJY9FBBXeu3wN8z6T"

	pwd, err := os.Getwd()
	if err != nil {
		return err
	}

	electrumProtocol := "ssl"
	if conf.ElectrumNoSSL {
		electrumProtocol = "tcp"
	}

	wallet := bdk.Wallet{
		Datadir:    filepath.Join(pwd, ".data"),
		Descriptor: fmt.Sprintf("wpkh(%s/84h/1h/0h/0/*)", priv),

		// This is all wonky stuff. We're on some kind of botched regtest...
		// However, the address format is mainnet - and that's the only thing
		// that matters.
		Network:  "bitcoin", // "bitcoin" == mainnet
		Electrum: fmt.Sprintf("%s://%s", electrumProtocol, conf.ElectrumHost),
	}

	zerolog.Ctx(ctx).Debug().
		Msgf("initiating electrum connection at %s", wallet.Electrum)

	initialBalance, err := wallet.GetBalance(ctx)
	if err != nil {
		return err
	}

	ticker := time.NewTicker(time.Second * 5)
	defer ticker.Stop()
	go func() {
		for range ticker.C {
			if err := wallet.Sync(ctx); err != nil {
				zerolog.Ctx(ctx).Err(err).Msgf("unable to sync")
				continue
			}

			balance, err := wallet.GetBalance(ctx)
			if err != nil {
				zerolog.Ctx(ctx).Err(err).Msgf("unable to get balance")
				continue
			}

			if reflect.DeepEqual(balance, initialBalance) {
				continue
			}

			zerolog.Ctx(ctx).Info().
				Msgf("balance changed: %+v -> %+v", initialBalance, balance)

			initialBalance = balance
		}
	}()

	mux := http.NewServeMux()
	path, handler := rpc.NewDrivechainServiceHandler(server.New(&wallet, proxy))

	mux.Handle(path, handler)

	const address = "localhost:8080"
	zerolog.Ctx(ctx).Info().Msgf("server: listening on %s", address)

	httpServer := http.Server{
		Addr: address,
		// Use h2c so we can serve HTTP/2 without TLS.
		Handler: h2c.NewHandler(mux, &http2.Server{}),
	}

	errs := make(chan error)
	go func() {
		errs <- httpServer.ListenAndServe()
	}()
	go func() {
		<-ctx.Done()

		ctx, cancel := context.WithTimeout(context.WithoutCancel(ctx), time.Second*1)
		defer cancel()

		if err := httpServer.Shutdown(ctx); err != nil {
			errs <- fmt.Errorf("shutdown HTTP server: %w", err)
			return
		}

		errs <- nil
	}()

	return <-errs
}

func startCoreProxy(ctx context.Context, conf Config) (*coreproxy.Bitcoind, error) {
	core, err := coreproxy.NewBitcoind(
		ctx, conf.BitcoinCoreHost,
		conf.BitcoinCoreRpcUser, conf.BitcoinCoreRpcPassword,
	)
	if err != nil {
		return nil, err
	}

	return core, nil
}
