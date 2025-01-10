import 'package:flutter/material.dart';
import 'package:slotman/drawer.dart';
import 'package:slotman/status.dart';

class SetupRacePage extends StatefulWidget {
  const SetupRacePage({super.key});

  final String title = 'Race Setup';

  @override
  State<SetupRacePage> createState() => SetupRacePageState();
}

class SetupRacePageState extends State<SetupRacePage> {
  static SetupRacePageState? injector;

  void setContent() {
    setState(() {});
  }

  @override
  Widget build(BuildContext context) {
    injector = this;

    var tracks = Status.race.tracks != 0 ? '${Status.race.tracks}' : '';
    var rounds = Status.race.rounds != 0 ? '${Status.race.rounds}' : '';

    TextEditingController titleController = TextEditingController(text: Status.race.title);
    TextEditingController tracksController = TextEditingController(text: tracks);
    TextEditingController roundsController = TextEditingController(text: rounds);

    return Scaffold(
      appBar: AppBar(
        backgroundColor: Theme.of(context).colorScheme.inversePrimary,
        title: Text(widget.title),
      ),
      drawer: MainDrawer(),
      body: Center(
        child: SizedBox(
          width: 360,
          child: Column(children: [
            SizedBox(height: 16),
            TextField(
              controller: titleController,
              decoration: InputDecoration(
                labelText: 'Race Title',
                hintText: 'Please enter the race title',
                border: OutlineInputBorder(),
              ),
              maxLength: 48,
              onChanged: (text) {
                Status.race.title = text;
              },
            ),
            TextField(
              controller: tracksController,
              decoration: InputDecoration(
                labelText: 'Number of tracks',
                hintText: 'Please enter the number of tracks',
                border: OutlineInputBorder(),
              ),
              maxLength: 3,
              onChanged: (text) {
                setState(() {
                  Status.race.tracks = int.parse(text);
                });
              },
            ),
            TextField(
              controller: roundsController,
              decoration: InputDecoration(
                labelText: 'Number of rounds',
                hintText: 'Please enter the number of rounds',
                border: OutlineInputBorder(),
              ),
              maxLength: 3,
              onChanged: (text) {
                setState(() {
                  Status.race.rounds = int.parse(text);
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
                Status.sndRace();
              },
              child: Text('Update'),
            )
          ]),
        ),
      ),
    );
  }
}
