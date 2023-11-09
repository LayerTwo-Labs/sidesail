// GENERATED CODE - DO NOT MODIFY BY HAND

// **************************************************************************
// AutoRouterGenerator
// **************************************************************************

// ignore_for_file: type=lint
// coverage:ignore-file

part of 'router.dart';

abstract class _$AppRouter extends RootStackRouter {
  // ignore: unused_element
  _$AppRouter({super.navigatorKey});

  @override
  final Map<String, PageFactory> pagesMap = {
    BlindMergedMiningTabRoute.name: (routeData) {
      return AutoRoutePage<dynamic>(
        routeData: routeData,
        child: const BlindMergedMiningTabPage(),
      );
    },
    DashboardTabRoute.name: (routeData) {
      return AutoRoutePage<dynamic>(
        routeData: routeData,
        child: const DashboardTabPage(),
      );
    },
    EthereumRPCTabRoute.name: (routeData) {
      return AutoRoutePage<dynamic>(
        routeData: routeData,
        child: const EthereumRPCTabPage(),
      );
    },
    HomeRoute.name: (routeData) {
      return AutoRoutePage<dynamic>(
        routeData: routeData,
        child: const HomePage(),
      );
    },
    NodeSettingsTabRoute.name: (routeData) {
      return AutoRoutePage<dynamic>(
        routeData: routeData,
        child: const NodeSettingsTabPage(),
      );
    },
    SailTestRoute.name: (routeData) {
      final args = routeData.argsAs<SailTestRouteArgs>();
      return AutoRoutePage<dynamic>(
        routeData: routeData,
        child: SailTestPage(
          key: args.key,
          child: args.child,
        ),
      );
    },
    SidechainExplorerTabRoute.name: (routeData) {
      return AutoRoutePage<dynamic>(
        routeData: routeData,
        child: const SidechainExplorerTabPage(),
      );
    },
    ThemeSettingsTabRoute.name: (routeData) {
      return AutoRoutePage<dynamic>(
        routeData: routeData,
        child: const ThemeSettingsTabPage(),
      );
    },
    TransferMainchainTabRoute.name: (routeData) {
      return AutoRoutePage<dynamic>(
        routeData: routeData,
        child: const TransferMainchainTabPage(),
      );
    },
    WithdrawalBundleTabRoute.name: (routeData) {
      return AutoRoutePage<dynamic>(
        routeData: routeData,
        child: const WithdrawalBundleTabPage(),
      );
    },
  };
}

/// generated route for
/// [BlindMergedMiningTabPage]
class BlindMergedMiningTabRoute extends PageRouteInfo<void> {
  const BlindMergedMiningTabRoute({List<PageRouteInfo>? children})
      : super(
          BlindMergedMiningTabRoute.name,
          initialChildren: children,
        );

  static const String name = 'BlindMergedMiningTabRoute';

  static const PageInfo<void> page = PageInfo<void>(name);
}

/// generated route for
/// [DashboardTabPage]
class DashboardTabRoute extends PageRouteInfo<void> {
  const DashboardTabRoute({List<PageRouteInfo>? children})
      : super(
          DashboardTabRoute.name,
          initialChildren: children,
        );

  static const String name = 'DashboardTabRoute';

  static const PageInfo<void> page = PageInfo<void>(name);
}

/// generated route for
/// [EthereumRPCTabPage]
class EthereumRPCTabRoute extends PageRouteInfo<void> {
  const EthereumRPCTabRoute({List<PageRouteInfo>? children})
      : super(
          EthereumRPCTabRoute.name,
          initialChildren: children,
        );

  static const String name = 'EthereumRPCTabRoute';

  static const PageInfo<void> page = PageInfo<void>(name);
}

/// generated route for
/// [HomePage]
class HomeRoute extends PageRouteInfo<void> {
  const HomeRoute({List<PageRouteInfo>? children})
      : super(
          HomeRoute.name,
          initialChildren: children,
        );

  static const String name = 'HomeRoute';

  static const PageInfo<void> page = PageInfo<void>(name);
}

/// generated route for
/// [NodeSettingsTabPage]
class NodeSettingsTabRoute extends PageRouteInfo<void> {
  const NodeSettingsTabRoute({List<PageRouteInfo>? children})
      : super(
          NodeSettingsTabRoute.name,
          initialChildren: children,
        );

  static const String name = 'NodeSettingsTabRoute';

  static const PageInfo<void> page = PageInfo<void>(name);
}

/// generated route for
/// [SailTestPage]
class SailTestRoute extends PageRouteInfo<SailTestRouteArgs> {
  SailTestRoute({
    Key? key,
    required Widget child,
    List<PageRouteInfo>? children,
  }) : super(
          SailTestRoute.name,
          args: SailTestRouteArgs(
            key: key,
            child: child,
          ),
          initialChildren: children,
        );

  static const String name = 'SailTestRoute';

  static const PageInfo<SailTestRouteArgs> page = PageInfo<SailTestRouteArgs>(name);
}

class SailTestRouteArgs {
  const SailTestRouteArgs({
    this.key,
    required this.child,
  });

  final Key? key;

  final Widget child;

  @override
  String toString() {
    return 'SailTestRouteArgs{key: $key, child: $child}';
  }
}

/// generated route for
/// [SidechainExplorerTabPage]
class SidechainExplorerTabRoute extends PageRouteInfo<void> {
  const SidechainExplorerTabRoute({List<PageRouteInfo>? children})
      : super(
          SidechainExplorerTabRoute.name,
          initialChildren: children,
        );

  static const String name = 'SidechainExplorerTabRoute';

  static const PageInfo<void> page = PageInfo<void>(name);
}

/// generated route for
/// [ThemeSettingsTabPage]
class ThemeSettingsTabRoute extends PageRouteInfo<void> {
  const ThemeSettingsTabRoute({List<PageRouteInfo>? children})
      : super(
          ThemeSettingsTabRoute.name,
          initialChildren: children,
        );

  static const String name = 'ThemeSettingsTabRoute';

  static const PageInfo<void> page = PageInfo<void>(name);
}

/// generated route for
/// [TransferMainchainTabPage]
class TransferMainchainTabRoute extends PageRouteInfo<void> {
  const TransferMainchainTabRoute({List<PageRouteInfo>? children})
      : super(
          TransferMainchainTabRoute.name,
          initialChildren: children,
        );

  static const String name = 'TransferMainchainTabRoute';

  static const PageInfo<void> page = PageInfo<void>(name);
}

/// generated route for
/// [WithdrawalBundleTabPage]
class WithdrawalBundleTabRoute extends PageRouteInfo<void> {
  const WithdrawalBundleTabRoute({List<PageRouteInfo>? children})
      : super(
          WithdrawalBundleTabRoute.name,
          initialChildren: children,
        );

  static const String name = 'WithdrawalBundleTabRoute';

  static const PageInfo<void> page = PageInfo<void>(name);
}
