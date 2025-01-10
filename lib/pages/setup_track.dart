import 'package:flutter/material.dart';
import 'package:slotman/drawer.dart';
import 'package:slotman/status.dart';

import '../locals.dart';

class SetupTrackPage extends StatefulWidget {
  const SetupTrackPage({super.key});

  final String title = 'Track Setup';

  @override
  State<SetupTrackPage> createState() => _SetupTrackPageState();
}

class _SetupTrackPageState extends State<SetupTrackPage> {
  int numberOfTracks = Status.numberOfTracks;

  @override
  Widget build(BuildContext context) {
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
              groupValue: numberOfTracks,
              onChanged: (int? value) {
                setState(() {
                  numberOfTracks = value!;
                });
              },
            ),
            RadioListTile<int>(
              title: Text('4 Tracks'),
              value: 4,
              groupValue: numberOfTracks,
              onChanged: (int? value) {
                setState(() {
                  numberOfTracks = value!;
                });
              },
            ),
            RadioListTile<int>(
              title: Text('6 Tracks'),
              value: 6,
              groupValue: numberOfTracks,
              onChanged: (int? value) {
                setState(() {
                  numberOfTracks = value!;
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
                Status.sndNumberOfTracks(numberOfTracks);
              },
              child: Text('Update'),
            )
          ]),
        ),
      ),
    );
  }
}
