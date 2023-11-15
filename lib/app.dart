import 'dart:async';

import 'package:flutter/material.dart';
import 'package:get_it/get_it.dart';
import 'package:sail_ui/theme/theme.dart';
import 'package:sail_ui/theme/theme_data.dart';
import 'package:sail_ui/widgets/core/scaffold.dart';
import 'package:sail_ui/widgets/loading_indicator.dart';
import 'package:sidesail/routing/router.dart';
import 'package:sidesail/rpc/rpc_config.dart';
import 'package:sidesail/rpc/rpc_mainchain.dart';
import 'package:sidesail/rpc/rpc_sidechain.dart';
import 'package:sidesail/storage/client_settings.dart';
import 'package:sidesail/storage/sail_settings/theme_settings.dart';

class SailApp extends StatefulWidget {
  // This key is used to programmatically trigger a rebuild of the MyApp widget.
  static GlobalKey<SailAppState> sailAppKey = GlobalKey();

  final Widget Function(BuildContext context, AppRouter router) builder;

  SailApp({
    required this.builder,
  }) : super(key: sailAppKey);

  @override
  State<SailApp> createState() => SailAppState();

  static SailAppState of(BuildContext context) {
    final SailAppState? result = context.findAncestorStateOfType<SailAppState>();
    if (result != null) return result;
    throw FlutterError.fromParts(<DiagnosticsNode>[
      ErrorSummary(
        'SailAppState.of() called with a context that does not contain a SailApp.',
      ),
    ]);
  }
}

class SailAppState extends State<SailApp> with WidgetsBindingObserver {
  AppRouter get router => GetIt.I.get<AppRouter>();
  SidechainContainer get _sidechain => GetIt.I.get<SidechainContainer>();
  MainchainRPC get mainchain => GetIt.I.get<MainchainRPC>();
  final settings = GetIt.I.get<ClientSettings>();

  SailThemeData? theme;

  @override
  void initState() {
    super.initState();
    WidgetsBinding.instance.addObserver(this);
    unawaited(loadTheme());

    // always attempt to start binaries. If we're already
    // connected (handled in dependencies), the start binary
    // function makes sure to not restart it
    initMainchainBinary();
    initSidechainBinary();
  }

  void rebuildUI() {
    setState(() {
      // This is a workaround to trigger the root widget to rebuild.
      SailApp.sailAppKey = GlobalKey();
    });
  }

  Future<void> loadTheme([SailThemeValues? themeToLoad]) async {
    themeToLoad ??= (await settings.getValue(ThemeSetting())).value;
    if (themeToLoad == SailThemeValues.platform) {
      // ignore: deprecated_member_use
      themeToLoad = WidgetsBinding.instance.window.platformBrightness == Brightness.light
          ? SailThemeValues.light
          : SailThemeValues.dark;
    }

    theme = _themeDataFromTheme(themeToLoad);

    setState(() {});
    await settings.setValue(ThemeSetting(newValue: themeToLoad));
  }

  SailThemeData _themeDataFromTheme(SailThemeValues theme) {
    switch (theme) {
      case SailThemeValues.light:
        return SailThemeData.lightTheme(_sidechain.rpc.chain.color);
      case SailThemeValues.dark:
        return SailThemeData.darkTheme(_sidechain.rpc.chain.color);
      default:
        throw Exception('Could not get theme');
    }
  }

  Future<void> initMainchainBinary() async {
    return mainchain.initBinary(
      context,
      mainchain.binary,
      bitcoinCoreBinaryArgs(mainchain.conf),
    );
  }

  Future<void> initSidechainBinary() async {
    return _sidechain.rpc.initBinary(
      context,
      _sidechain.rpc.chain.binary,
      _sidechain.rpc.binaryArgs(mainchain.conf),
    );
  }

  bool get inInitialSetup => theme == null;

  @override
  Widget build(BuildContext context) {
    if (inInitialSetup) {
      return MaterialApp(
        home: Material(
          color: _sidechain.rpc.chain.color,
          child: SailTheme(
            data: SailThemeData.darkTheme(_sidechain.rpc.chain.color),
            child: const SailScaffold(
              body: Center(
                child: SailCircularProgressIndicator(),
              ),
            ),
          ),
        ),
      );
    }

    return SailTheme(
      data: theme!,
      child: widget.builder(context, router),
    );
  }

  @override
  void dispose() {
    WidgetsBinding.instance.removeObserver(this);
    super.dispose();
  }
}

Future<void> waitForBoolToBeTrue(
  Future<bool> Function() boolGetter, {
  Duration pollInterval = const Duration(milliseconds: 100),
}) async {
  bool result = await boolGetter();
  if (!result) {
    await Future.delayed(pollInterval);
    await waitForBoolToBeTrue(boolGetter, pollInterval: pollInterval);
  }
}
