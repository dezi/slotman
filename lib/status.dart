import 'dart:convert';
import 'package:slotman/messages/tracks.dart';
import 'package:slotman/socket.dart';

class Status {

  static int numberOfTracks = 0;

  static void initialize() async {
    var tracks = Tracks(mode:'get');
    Socket.transmit(jsonEncode(tracks));
  }

  static void sndNumberOfTracks(int val) {
    numberOfTracks = val;
    var tracks = Tracks(mode:'set', numberOfTracks: numberOfTracks);
    Socket.transmit(jsonEncode(tracks));
  }

  static void rcvNumberOfTracks(Tracks tracks) {
    numberOfTracks = tracks.numberOfTracks;
  }
}