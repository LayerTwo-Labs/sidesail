import 'package:auto_route/auto_route.dart';
import 'package:flutter/widgets.dart';
import 'package:sidesail/pages/withdrawal_bundle_page.dart';
import 'package:sidesail/pages/withdrawal_page.dart';
import 'package:sidesail/routing/routes_export.dart';

part 'router.gr.dart';

/* In this Router, all of the app routes are defined. Router code is auto-generated.
*  In replaceInRouteName we replace the Page suffix with Route in all generated
*  routes. e.g: The annotated HomePage class will be generated as HomeRoute. HomeRoute
*  is the component you route to.
*  
*  Use the [watch] flag to watch the files' system for edits and rebuild as necessary.
*  $ flutter packages pub run build_runner watch
*  if you want the generator to run one time and exit, use
*  $ flutter packages pub run build_runner build

*/
@AutoRouterConfig(replaceInRouteName: 'Page,Route')
class AppRouter extends _$AppRouter {
  @override
  RouteType get defaultRouteType => const RouteType.adaptive();

  @override
  List<AutoRoute> get routes => [
        AutoRoute(
          page: HomeRoute.page,
          initial: true,
        ),
        AutoRoute(
          page: WithdrawalBundleRoute.page,
        ),

        AutoRoute(
          page: WithdrawalRoute.page,
        ),

        /// This route is used in tests so that we can pump a widget into a route
        /// and use the real router for our test
        AutoRoute(page: SailTestRoute.page),
      ];
}
