import 'package:drivechain_client/gen/drivechain/v1/drivechain.pbgrpc.dart';
import 'package:drivechain_client/gen/google/protobuf/empty.pb.dart';
import 'package:fixnum/fixnum.dart';
import 'package:flutter/widgets.dart';
import 'package:grpc/grpc.dart';

class DrivechainService extends InheritedWidget {
  late final DrivechainServiceClient _client;

  DrivechainService({required super.child, super.key}) {
    _client = DrivechainServiceClient(DrivechainChannel());
  }

  Future<String> sendTransaction(Map<String, int> destinations, [double? satoshiPerVbyte]) async {
    final request = SendTransactionRequest(
      destinations: destinations.map((k, v) => MapEntry(k, Int64(v))),
      satoshiPerVbyte: satoshiPerVbyte,
    );
    
    final response = await _client.sendTransaction(request);
    return response.txid;
  }

  /// Returns a tuple of the confirmed and pending balance in satoshi
  Future<(Int64, Int64)> getBalance() async {
    final response = await _client.getBalance(Empty());
    return (response.confirmedSatoshi, response.pendingSatoshi);
  }

  Future<String> getNewAddress() async {
    final response = await _client.getNewAddress(Empty());
    return response.address;
  }

  Future<List<Transaction>> listTransactions() async {
    final response = await _client.listTransactions(Empty());
    return response.transactions;
  }

  Future<List<UnconfirmedTransaction>> listUnconfirmedTransactions() async {
    final response = await _client.listUnconfirmedTransactions(Empty());
    return response.unconfirmedTransactions;
  }

  Future<List<ListRecentBlocksResponse_RecentBlock>> listRecentBlocks() async {
    final response = await _client.listRecentBlocks(Empty());
    return response.recentBlocks;
  }

  

  @override
  bool updateShouldNotify(covariant InheritedWidget oldWidget) => false;

  static DrivechainService of(BuildContext context) {
    return context.dependOnInheritedWidgetOfExactType<DrivechainService>()!;
  }
}

class DrivechainChannel extends ClientChannel {
  DrivechainChannel()
      : super(
          const String.fromEnvironment('DRIVECHAIN_URL', defaultValue: "http://localhost:8080").split(':').first,
          port: int.parse(
              const String.fromEnvironment('DRIVECHAIN_URL', defaultValue: "http://localhost:8080").split(':').last),
          options: const ChannelOptions(
            credentials: ChannelCredentials.insecure(),
          ),
        );
}
