import 'dart:async';
import 'dart:convert';
import 'dart:io';

import 'package:flutter/material.dart';
import 'package:logger/logger.dart';
import 'package:path_provider/path_provider.dart';
import 'package:sidesail/logger.dart';

class ProcessProvider extends ChangeNotifier {
  ProcessProvider();

  // List of currently running processes.
  List<int> pids = [];

  final Map<int, int> _exitCodes = {};
  final List<int> _runningProcesses = [];

  final Map<int, Stream<String>> _stdoutStreams = {};
  final Map<int, Stream<String>> _stderrStreams = {};

  Stream<String> stdout(int pid) => _stdoutStreams[pid] ?? const Stream.empty();
  Stream<String> stderr(int pid) => _stderrStreams[pid] ?? const Stream.empty();
  bool running(int pid) => _runningProcesses.contains(pid);

  // Starts a binary located in the asset bundle included with the app.
  Future<int> start(BuildContext context, String binary, List<String> args) async {
    // We're NOT looking up here based on platform and architecture. That is instead
    // handled at app bundling time, where it's up the person/process cutting the
    // release to place the appropriate binaries in the correct place.
    final binResource = await DefaultAssetBundle.of(context).load('assets/bin/$binary');

    final temp = await getTemporaryDirectory();

    final file = File('${temp.path}/$binary');

    // Have to convert the ByteData -> List<int>. Side note: why tf does writeAsBytes
    // operate on a list of numbers? Jesus
    // https://stackoverflow.com/a/50121777
    final buffer = binResource.buffer;
    await file.writeAsBytes(
      buffer.asUint8List(binResource.offsetInBytes, binResource.lengthInBytes),
    );

    // Must be executable before we can start.
    // TODO: what to do on Windows here
    await Process.run('chmod', ['+x', file.path]);

    final process = await Process.start(
      file.path,
      args,
      mode: ProcessStartMode.normal, // when the flutter app quits, this process quit
    );
    log.d('started "$binary" with pid ${process.pid}');
    _runningProcesses.add(process.pid);

    // Let output streaming chug in the background

    // Proper logs are written here
    _stdoutStreams[process.pid] = process.stdout.transform(utf8.decoder);

    // Error messages while starting up are written here
    _stderrStreams[process.pid] = process.stderr.transform(utf8.decoder);

    // Register a handler for when the process stops.
    unawaited(
      process.exitCode.then((code) async {
        var level = Level.error;
        // Node software exited with a success code. Someone called
        // `drivechain-cli stop`?
        if (code == 0) {
          level = Level.info;
        }

        final errLogs = await (_stderrStreams[process.pid] ?? const Stream.empty()).toList();
        log.log(level, '"$binary" exited with code $code: $errLogs');

        // Forward to listeners that the process finished.
        _exitCodes[process.pid] = code;
        _runningProcesses.remove(process.pid);
        notifyListeners();
      }),
    );

    notifyListeners();
    return process.pid;
  }

  @override
  void dispose() {
    super.dispose();

    log.d('dispose process provider: killing processes $_runningProcesses');
    _runningProcesses.forEach(Process.killPid);
  }
}
