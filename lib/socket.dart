import 'dart:convert';
import 'dart:developer';
import 'package:slotman/messages/message.dart';
import 'package:slotman/messages/tracks.dart';
import 'package:slotman/status.dart';
import 'package:web_socket_channel/web_socket_channel.dart';

class Socket {
  static WebSocketChannel? channel;

  static initialize() async {
    log('########### channel connect ${Uri.base}');

    channel = WebSocketChannel.connect(
      //Uri.parse('wss://echo.websocket.events'),
      Uri.parse('ws://localhost:8888/ws'),
    );

    log('########### channel $channel');

    await channel!.ready;

    log('########### channel ready $channel');

    channel!.stream.listen((json) {

      log('########### channel rcv $json');
      var message = Message.fromJson(jsonDecode(json));

      switch (message.tag()) {
        case 'tracks|set':
        case 'tracks|get':
          var tracks = Tracks.fromJson(jsonDecode(json));
          Status.rcvNumberOfTracks(tracks);
          break;
      }
    });
  }

  static void transmit(String json) {
    log('########### channel snd $json');
    channel!.sink.add(json);
  }
}
