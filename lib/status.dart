import 'dart:convert';
import 'package:slotman/messages/controller.dart';
import 'package:slotman/messages/race.dart';
import 'package:slotman/messages/tracks.dart';
import 'package:slotman/pages/setup_controller.dart';
import 'package:slotman/pages/setup_race.dart';
import 'package:slotman/pages/setup_tracks.dart';
import 'package:slotman/socket.dart';

class Status {
  static int selectedController = 0;
  static bool isCalibrating = false;

  static Race race = Race(mode: 'set');
  static Tracks tracks = Tracks(mode: 'set');
  static Controller controller = Controller(mode: 'set');

  static Future<void> initialize() async {
    var tracks = Tracks(mode: 'get');
    Socket.transmit(jsonEncode(tracks));
    var race = Race(mode: 'get');
    Socket.transmit(jsonEncode(race));
  }

  static void sndRace() {
    Socket.transmit(jsonEncode(race));
  }

  static void rcvRace(Race race) {
    Status.race = race;
    if (SetupRacePageState.injector != null) {
      SetupRacePageState.injector!.setContent();
    }
  }

  static void sndTracks(Tracks tracks) {
    Status.tracks = tracks;
    Socket.transmit(jsonEncode(tracks));
  }

  static void rcvTracks(Tracks tracks) {
    Status.tracks = tracks;
    if (SetupTracksPageState.injector != null) {
      SetupTracksPageState.injector!.setContent();
    }
  }

  static void sndController(Controller controller) {
    Status.controller = controller;
    Socket.transmit(jsonEncode(controller));
  }

  static void rcvController(Controller controller) {
    Status.controller = controller;
    if (SetupControllerPageState.injector != null) {
      SetupControllerPageState.injector!.setContent();
    }
  }
}
