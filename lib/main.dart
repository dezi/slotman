import 'dart:developer';
import 'package:flutter/material.dart';
import 'package:slotman/locals.dart';
import 'package:slotman/pages/join.dart';
import 'package:slotman/pages/start.dart';
import 'package:slotman/pages/setup_pilot.dart';
import 'package:slotman/pages/setup_race.dart';
import 'package:slotman/pages/setup_controller.dart';
import 'package:slotman/pages/setup_track.dart';
import 'package:slotman/socket.dart';
import 'package:slotman/status.dart';

void main() async {
  await Locals.initialize();
  await Socket.initialize();
  await Status.initialize();
  runApp(MainApp());
}

class MainApp extends StatefulWidget {
  const MainApp({super.key});

  @override
  State<MainApp> createState() => _MainAppState();
}

class _MainAppState extends State<MainApp> {

  @override
  void initState() {
    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Slot Race Manager',
      theme: ThemeData(
        colorScheme: ColorScheme.fromSeed(seedColor: Colors.deepPurple),
        useMaterial3: true,
      ),
      home: const StartPage(),
      routes: {
        '/join': (context) => JoinPage(),
        '/setup/race': (context) => SetupRacePage(),
        '/setup/pilot': (context) => SetupPilotPage(),
        '/setup/track': (context) => SetupTrackPage(),
        '/setup/controller': (context) => SetupControllerPage(),
      },
    );
  }
}
