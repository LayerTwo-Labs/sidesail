import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:sail_ui/sail_ui.dart';
import 'package:sail_ui/widgets/core/sail_snackbar.dart';
import 'package:sail_ui/widgets/core/sail_text.dart';

class SingleValueContainer extends StatelessWidget {
  final String? label;
  final dynamic value;
  final double? width;
  final String? trailingText;
  final Widget? icon;
  final Widget? prefixAction;
  final bool copyable;
  final bool italic;
  final Color? color;

  const SingleValueContainer({
    super.key,
    required this.value,
    this.width,
    this.label,
    this.trailingText,
    this.icon,
    this.prefixAction,
    this.italic = false,
    this.copyable = true,
    this.color,
  });

  @override
  Widget build(BuildContext context) {
    return SailRow(
      crossAxisAlignment: CrossAxisAlignment.center,
      mainAxisAlignment: MainAxisAlignment.center,
      spacing: SailStyleValues.padding08,
      children: [
        if (prefixAction != null) prefixAction!,
        if (prefixAction != null) const SailSpacing(SailStyleValues.padding05),
        if (icon != null) icon!,
        if (label != null)
          if (width != null)
            SizedBox(
              width: width,
              child: SailText.secondary12(label!, italic: italic, customColor: color),
            )
          else
            SailText.secondary12(label!, italic: italic, customColor: color),
        Expanded(
          child: SailScaleButton(
            onPressed: copyable
                ? () {
                    Clipboard.setData(ClipboardData(text: value.toString()));
                    showSnackBar(context, 'Copied ${copyLabel()}');
                  }
                : null,
            child: SailText.primary12(value.toString()),
          ),
        ),
        if (trailingText != null) SailText.secondary12(trailingText!),
      ],
    );
  }

  String copyLabel() {
    final textToShow = label ?? value.toString();
    if (textToShow.length > 50) {
      return '${textToShow.substring(0, 50)}...';
    }

    return textToShow;
  }
}
