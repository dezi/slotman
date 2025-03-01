import 'dart:convert';
import 'dart:developer';
import 'package:slotman/locals.dart';
import 'package:slotman/messages/controller.dart';
import 'package:slotman/messages/info.dart';
import 'package:slotman/messages/message.dart';
import 'package:slotman/messages/pilot.dart';
import 'package:slotman/messages/race.dart';
import 'package:slotman/messages/tracks.dart';
import 'package:slotman/status.dart';
import 'package:web_socket_channel/web_socket_channel.dart';

class Socket {
  static WebSocketChannel? channel;

  static Future<void> initialize() async {
    if (channel != null) return;

    var uri = Uri.parse('ws://localhost:8888/ws/${Locals.appUuid}');

    log('########### channel connect ${uri.path}');

    channel = WebSocketChannel.connect(uri);

    log('########### channel $channel');

    await channel!.ready;

    log('########### channel ready $channel');

    channel!.stream.listen((json) {
      log('########### channel rcv $json');

      var message = Message.fromJson(jsonDecode(json));
      log('########### channel rcv tag=${message.tag()}');

      switch (message.tag()) {
        case 'race|set':
          var race = Race.fromJson(jsonDecode(json));
          Status.rcvRace(race);
          break;
        case 'tracks|set':
          var tracks = Tracks.fromJson(jsonDecode(json));
          Status.rcvTracks(tracks);
          break;
        case 'controller|set':
          var controller = Controller.fromJson(jsonDecode(json));
          Status.rcvController(controller);
        case 'pilot|set':
          var pilot = Pilot.fromJson(jsonDecode(json));
          Status.rcvPilot(pilot);
        case 'info|set':
          var info = Info.fromJson(jsonDecode(json));
          Status.rcvInfo(info);
      }
    });
  }

  static void transmit(String json) {
    log('########### channel snd $json');
    channel!.sink.add(json);
  }
}
