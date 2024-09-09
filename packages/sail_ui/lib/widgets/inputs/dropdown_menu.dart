import 'package:flutter/material.dart';
import 'package:sail_ui/sail_ui.dart';
import 'package:sail_ui/widgets/inputs/menu.dart';
import 'package:sail_ui/widgets/inputs/menu_items.dart';

class SailDropdownButton<T> extends StatefulWidget {
  const SailDropdownButton({
    required this.items,
    required this.onChanged,
    required this.value,
    this.icon = const Icon(
      Icons.expand_more,
      size: 16,
    ),
    this.width,
    this.large = false,
    super.key,
  });

  final List<SailDropdownItem<T>> items;
  final ValueChanged<T> onChanged;
  final T value;
  final Widget? icon;
  final double? width;
  final bool large;

  @override
  State<StatefulWidget> createState() => _SailDropdownButtonState<T>();
}

class _SailDropdownButtonState<T> extends State<SailDropdownButton<T>> {
  @override
  Widget build(BuildContext context) {
    var isWindows = Theme.of(context).platform == TargetPlatform.windows;

    var items = widget.items
        .map(
          (e) => SailMenuItem(
            height: widget.large ? 24 : 18,
            onSelected: () {
              widget.onChanged(e.value);
            },
            child: e,
          ),
        )
        .toList();

    var currentIndex = widget.items.indexWhere((element) => element.value == widget.value);
    Widget currentItem = widget.items[currentIndex];

    var offsetY = 0.0;
    for (var i = 0; i < currentIndex; i += 1) {
      offsetY += items[i].height;
    }

    if (widget.width != null) {
      currentItem = Expanded(
        child: currentItem,
      );
    }

    return SizedBox(
      width: widget.width,
      height: widget.large ? 32 : 24,
      child: _Button(
        padding: EdgeInsets.only(
          left: 8,
          right: widget.icon == null ? 8 : 4,
        ),
        onPressed: () {
          var bounds = getGlobalBoundsForContext(context);
          showSailMenu(
            context: context,
            preferredAnchorPoint: Offset(
              bounds.left - (isWindows ? 1 : 9),
              bounds.top - offsetY - 3,
            ),
            menu: SailMenu(
              items: items,
              width: widget.width,
            ),
          );
        },
        child: Row(
          mainAxisAlignment: MainAxisAlignment.spaceBetween,
          mainAxisSize: MainAxisSize.max,
          children: [
            currentItem,
            if (widget.icon != null)
              Padding(
                padding: const EdgeInsets.only(left: 4),
                child: widget.icon,
              ),
          ],
        ),
      ),
    );
  }
}

class SailDropdownItem<T> extends StatelessWidget {
  const SailDropdownItem({
    required this.value,
    required this.child,
    super.key,
  });

  final T value;
  final Widget child;

  @override
  Widget build(BuildContext context) {
    return Align(
      alignment: Alignment.centerLeft,
      child: child,
    );
  }
}

class _Button extends StatelessWidget {
  final VoidCallback? onPressed;
  final Widget child;
  final EdgeInsets padding;
  final bool large;
  final bool important;
  final bool enabled;

  const _Button({
    required this.onPressed,
    required this.child,
    this.padding = const EdgeInsets.symmetric(
      horizontal: 12,
      vertical: 0,
    ),
    this.large = false,
    this.important = false,
    this.enabled = true,
  });

  @override
  Widget build(BuildContext context) {
    final theme = Theme.of(context);
    final isWindows = theme.platform == TargetPlatform.windows;

    Color textColor;
    if (enabled && onPressed != null) {
      if (important) {
        textColor = theme.buttonTheme.colorScheme!.onPrimary;
      } else {
        textColor = theme.buttonTheme.colorScheme!.onSurface;
      }
    } else {
      textColor = theme.disabledColor;
    }

    return SizedBox(
      height: large ? 32 : 24,
      child: MaterialButton(
        // textTheme: textTheme,
        color: important ? theme.colorScheme.primary : theme.buttonTheme.colorScheme!.surface,
        hoverColor: Theme.of(context).hoverColor,
        shape: RoundedRectangleBorder(
          borderRadius: BorderRadius.all(Radius.circular(isWindows ? 3 : 4)),
          side: const BorderSide(
            color: Colors.grey,
            width: 0.5,
          ),
        ),
        onPressed: enabled ? onPressed : null,
        elevation: 0,
        hoverElevation: 0,
        focusElevation: 0,
        minWidth: 32,
        padding: EdgeInsets.zero,
        child: Container(
          alignment: Alignment.center,
          height: large ? 32 : 24,
          padding: padding,
          child: DefaultTextStyle(
            style: SailStyleValues.twelve.copyWith(
              color: textColor,
            ),
            child: child,
          ),
        ),
      ),
    );
  }
}