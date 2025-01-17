import 'dart:convert';
import 'package:slotman/messages/controller.dart';
import 'package:slotman/messages/pilot.dart';
import 'package:slotman/messages/race.dart';
import 'package:slotman/messages/tracks.dart';
import 'package:slotman/pages/setup_controller.dart';
import 'package:slotman/pages/setup_race.dart';
import 'package:slotman/pages/setup_tracks.dart';
import 'package:slotman/socket.dart';

class Status {

  static Race race = Race();
  static Pilot pilot = Pilot();
  static Tracks tracks = Tracks();
  static Controller controller = Controller();
  static Map<String,Pilot> pilots = <String,Pilot>{};

  static Future<void> initialize() async {
    var tracks = Tracks(mode: 'get');
    Socket.transmit(jsonEncode(tracks));
    var race = Race(mode: 'get');
    Socket.transmit(jsonEncode(race));
    Socket.transmit(jsonEncode(pilot));
  }

  static void sndPilot(Pilot pilot) {
    Status.pilot = pilot;
    Socket.transmit(jsonEncode(pilot));
  }

  static void rcvPilot(Pilot pilot) {
    pilots[pilot.appUuid] = pilot;
  }

  static void sndRace(Race race) {
    Status.race = race;
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
