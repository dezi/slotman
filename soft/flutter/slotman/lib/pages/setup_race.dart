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

  var race = Status.race.clone();

  var titleCtl = TextEditingController(text: Status.race.title);
  var tracksCtl = TextEditingController(text: Status.race.tracks != 0 ? '${Status.race.tracks}' : '');
  var roundsCtl = TextEditingController(text: Status.race.rounds != 0 ? '${Status.race.rounds}' : '');

  void setContent() {
    setState(() {
      race = Status.race.clone();
    });
  }

  @override
  Widget build(BuildContext context) {
    injector = this;

    List<Widget> pilots = [];

    Status.pilots.forEach((appUuid, pilot) {
      String initials = "";
      if (pilot.firstName.isNotEmpty) initials += pilot.firstName[0];
      if (pilot.lastName.isNotEmpty) initials += pilot.lastName[0];

      pilots.add(SizedBox(height: 16));

      pilots.add(
        UserAccountsDrawerHeader(
          accountName: Text("${pilot.firstName} ${pilot.lastName}"),
          accountEmail: Text(pilot.carModel),
          currentAccountPicture: CircleAvatar(
            backgroundColor: Colors.orange,
            child: Text(
              initials,
              style: TextStyle(fontSize: 40.0),
            ),
          ),
          otherAccountsPictures: [
            Text(
              "2",
              style: TextStyle(
                fontSize: 64,
                fontWeight: FontWeight.bold,
                color: Colors.white,
              ),
            ),
          ],
          otherAccountsPicturesSize: Size.square(100),
        ),
      );
    });

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
              controller: titleCtl,
              decoration: InputDecoration(
                labelText: 'Race Title',
                hintText: 'Please enter the race title',
                border: OutlineInputBorder(),
              ),
              maxLength: 48,
              onChanged: (text) {
                setState(() {
                  race.title = text;
                });
              },
            ),
            TextField(
              controller: tracksCtl,
              decoration: InputDecoration(
                labelText: 'Number of tracks',
                hintText: 'Please enter the number of tracks',
                border: OutlineInputBorder(),
              ),
              maxLength: 3,
              onChanged: (text) {
                setState(() {
                  race.tracks = int.parse(text);
                });
              },
            ),
            TextField(
              controller: roundsCtl,
              decoration: InputDecoration(
                labelText: 'Number of rounds',
                hintText: 'Please enter the number of rounds',
                border: OutlineInputBorder(),
              ),
              maxLength: 3,
              onChanged: (text) {
                setState(() {
                  race.rounds = int.parse(text);
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
                Status.sndRace(race);
              },
              child: Text('Update'),
            ),
            ...pilots,
          ]),
        ),
      ),
    );
  }
}
