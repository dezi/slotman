import 'dart:convert';
import 'package:slotman/messages/controller.dart';
import 'package:slotman/messages/tracks.dart';
import 'package:slotman/pages/setup_controller.dart';
import 'package:slotman/socket.dart';

class Status {
  static int numberOfTracks = 0;
  static int selectedController = 0;
  static bool isCalibrating = false;

  static Future<void> initialize() async {
    var tracks = Tracks(mode: 'get');
    Socket.transmit(jsonEncode(tracks));
  }

  static void sndNumberOfTracks(int numberOfTracks) {
    Status.numberOfTracks = numberOfTracks;
    var tracks = Tracks(mode: 'set', numberOfTracks: numberOfTracks);
    Socket.transmit(jsonEncode(tracks));
  }

  static void rcvNumberOfTracks(Tracks tracks) {
    Status.numberOfTracks = tracks.numberOfTracks;
  }

  static void sndCalibrateController(int selectedController, bool isCalibrating) {
    Status.selectedController = selectedController;
    Status.isCalibrating = isCalibrating;
    var controller = Controller(
      mode: 'set',
      controller: selectedController,
      isCalibrating: isCalibrating,
    );
    Socket.transmit(jsonEncode(controller));
  }

  static void rcvCalibrateController(Controller controller) {
    if (SetupControllerPageState.injector != null) {
      SetupControllerPageState.injector!.setMinMaxValue(
        controller.controller,
        controller.isCalibrating,
        controller.minValue,
        controller.maxValue,
      );
    }
  }
}
