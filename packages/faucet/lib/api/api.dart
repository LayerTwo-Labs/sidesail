import 'dart:async';
import 'dart:convert';

import 'package:http/http.dart' as http;
import 'package:sail_ui/bitcoin.dart';
import 'package:sail_ui/sail_ui.dart';

/// RPC connection to the mainchain node.
abstract class API {
  final String apiURL;

  API({required this.apiURL});

  Future<List<UTXO>> listClaims();
  Future<String> claim(String address, double amount);
}

class APILive extends API {
  APILive({required super.apiURL});

  @override
  Future<List<UTXO>> listClaims() async {
    final url = Uri.parse("$apiURL/listclaims");
    final res = await http.get(url);

    if (res.statusCode == 200) {
      final List<dynamic> claimsJSON = jsonDecode(res.body);

      List<UTXO> transactions = claimsJSON.map((jsonItem) => UTXO.fromMap(jsonItem)).toList();
      return transactions;
    } else {
      throw Exception('could not list claims');
    }
  }

  @override
  Future<String> claim(String address, double amount) async {
    amount = cleanAmount(amount);

    final url = Uri.parse("$apiURL/claim");

    Map<String, dynamic> requestBody = {
      'destination': address.trim(),
      'amount': amount.toString().trim(),
    };

    final res = await http.post(
      url,
      body: json.encode(requestBody),
    );

    if (res.statusCode == 200) {
      final jsonResponse = json.decode(res.body);
      final txid = jsonResponse['txid'];
      return txid;
    } else {
      throw Exception(res.body);
    }
  }
}
