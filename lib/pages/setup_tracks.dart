import 'package:flutter/material.dart';
import 'package:slotman/drawer.dart';
import 'package:slotman/messages/tracks.dart';
import 'package:slotman/status.dart';

import '../locals.dart';

class SetupTracksPage extends StatefulWidget {
  const SetupTracksPage({super.key});

  final String title = 'Track Setup';

  @override
  State<SetupTracksPage> createState() => SetupTracksPageState();
}

class SetupTracksPageState extends State<SetupTracksPage> {

  static SetupTracksPageState? injector;

  var tracks = Tracks.clone(Status.tracks);

  void setContent() {
    setState(() {
      tracks = Tracks.clone(Status.tracks);
    });
  }

  @override
  Widget build(BuildContext context) {

    injector = this;

    return Scaffold(
      appBar: AppBar(
        backgroundColor: Theme.of(context).colorScheme.inversePrimary,
        title: Text(widget.title),
      ),
      drawer: MainDrawer(),
      body: Center(
        child: SizedBox(
          width: 200,
          child: ListView(children: [
            SizedBox(height: 16),
            RadioListTile<int>(
              title: Text('2 Tracks'),
              value: 2,
              groupValue: tracks.tracks,
              onChanged: (int? value) {
                setState(() {
                  tracks.tracks = value!;
                });
              },
            ),
            RadioListTile<int>(
              title: Text('4 Tracks'),
              value: 4,
              groupValue: tracks.tracks,
              onChanged: (int? value) {
                setState(() {
                  tracks.tracks = value!;
                });
              },
            ),
            RadioListTile<int>(
              title: Text('6 Tracks'),
              value: 6,
              groupValue: tracks.tracks,
              onChanged: (int? value) {
                setState(() {
                  tracks.tracks = value!;
                });
              },
            ),
            ElevatedButton(
              style: ElevatedButton.styleFrom(
                foregroundColor: Colors.blue,
                minimumSize: Size(200, 48),
                padding: EdgeInsets.symmetric(horizontal: 16),
                shape: const RoundedRectangleBorder(
                  borderRadius: BorderRadius.all(Radius.circular(2)),
                ),
              ),
              onPressed: () {
                Status.sndTracks(tracks);
              },
              child: Text('Update'),
            )
          ]),
        ),
      ),
    );
  }
}
