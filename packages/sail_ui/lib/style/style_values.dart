import 'package:flutter/material.dart';

abstract class SailStyleValues {
  // ---------- PADDINGS AND SPACINGS ----------
  static BorderRadius borderRadiusRegular = BorderRadius.circular(21);
  static BorderRadius borderRadiusButton = BorderRadius.circular(4);

  static const double padding05 = 3;
  static const double padding08 = 4;
  static const double padding10 = 5;
  static const double padding12 = 6;
  static const double padding15 = 8;
  static const double padding20 = 10;
  static const double padding25 = 13;
  static const double padding30 = 15;
  static const double padding40 = 20;
  static const double padding50 = 25;

  // ---------- ICON SIZES ----------
  static const iconSizePrimary = 14.0;

  // ---------- FONTS ----------
  static const lightWeight = FontWeight.w300;
  static const regularWeight = FontWeight.w400;
  static const mediumWeight = FontWeight.w500;

  static const TextStyle twentyFour = TextStyle(
    fontSize: 24,
    fontWeight: regularWeight,
  );
  static const TextStyle twentyTwo = TextStyle(
    fontSize: 22,
    fontWeight: regularWeight,
  );
  static const TextStyle twenty = TextStyle(
    fontSize: 20,
    fontWeight: regularWeight,
  );
  static const TextStyle fifteen = TextStyle(
    fontSize: 16,
    fontWeight: regularWeight,
  );
  static const TextStyle thirteen = TextStyle(
    fontSize: 14,
    fontWeight: regularWeight,
  );
  static const TextStyle twelve = TextStyle(
    fontSize: 12,
    fontWeight: regularWeight,
  );
  static const TextStyle eleven = TextStyle(
    fontSize: 12,
    fontWeight: regularWeight,
  );
}
