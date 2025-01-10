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
          width: 300,
          child: Column(children: [
            SizedBox(height: 16),
            RadioListTile<int>(
              title: Text('2 Tracks'),
              value: 2,
              groupValue: Status.numberOfTracks,
              onChanged: (int? value) {
                setState(() {
                  Status.sndNumberOfTracks(value!);
                });
              },
            ),
            RadioListTile<int>(
              title: Text('4 Tracks'),
              value: 4,
              groupValue: Status.numberOfTracks,
              onChanged: (int? value) {
                setState(() {
                  Status.sndNumberOfTracks(value!);
                });
              },
            ),
            RadioListTile<int>(
              title: Text('6 Tracks'),
              value: 6,
              groupValue: Status.numberOfTracks,
              onChanged: (int? value) {
                setState(() {
                  Status.sndNumberOfTracks(value!);
                });
              },
            )
          ]),
        ),
      ),
    );
  }
}
