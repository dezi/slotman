import 'package:flutter/material.dart';
import 'package:slotman/drawer.dart';
import 'package:slotman/messages/info.dart';
import 'package:slotman/status.dart';

class InfoPage extends StatefulWidget {
  const InfoPage({super.key});

  final String title = 'Race Info';

  @override
  State<InfoPage> createState() => InfoPageState();
}

class InfoPageState extends State<InfoPage> {
  static InfoPageState? injector;

  final width = 420.0;

  var race = Status.race;
  var infos = Status.infos;
  var pilots = Status.pilots;

  void setContent() {
    setState(() {
      race = Status.race;
      infos = Status.infos;
      pilots = Status.pilots;
    });
  }

  Widget stdBox(double width, String text) {
    return SizedBox(
      width: width,
      child: Text(
        text,
        style: const TextStyle(
          fontSize: 20,
          fontWeight: FontWeight.bold,
        ),
      ),
    );
  }

  Widget divider() {
    return Center(
      child: SizedBox(
        width: width,
        child: Divider(thickness: 2),
      ),
    );
  }

  List<Widget> buildPositions() {
    final List<Widget> positions = [];

    final List<Info> sorted = [];

    infos.forEach((track, info) {
      sorted.add(info);
    });

    sorted.sort((a, b) => a.position.compareTo(b.position));

    positions.add(SizedBox(height: 24));

    positions.add(Center(
      child: Text(
        "Aktueller Rennstand",
        style: TextStyle(
          fontSize: 22,
          fontWeight: FontWeight.bold,
        ),
      ),
    ));

    positions.add(SizedBox(height: 16));

    positions.add(
      Center(
        child: SizedBox(
          width: width,
          child: Row(
            children: [
              stdBox(50, "#"),
              stdBox(160, "Pilot"),
              stdBox(50, "Lap"),
              stdBox(80, "Round"),
              stdBox(80, "Speed"),
            ],
          ),
        ),
      ),
    );

    positions.add(divider());

    var position = 1;

    for (var info in sorted) {
      final pilot = pilots[info.pilotUuid];
      if (pilot == null) continue;

      final nameTag = "${pilot.firstName.substring(0, 1)}. ${pilot.lastName}";

      positions.add(
        Center(
          child: SizedBox(
            width: width,
            child: Row(
              children: [
                stdBox(50, position.toString()),
                stdBox(160, nameTag),
                stdBox(50, info.rounds.toString()),
                stdBox(80, info.actRound.toStringAsFixed(3)),
                stdBox(80, info.actSpeed.toStringAsFixed(0)),
              ],
            ),
          ),
        ),
      );

      position++;
    }

    positions.add(divider());

    return positions;
  }

  List<Widget> buildTopRounds() {
    final List<Widget> positions = [];

    final List<Info> sorted = [];

    infos.forEach((track, info) {
      sorted.add(info);
    });

    sorted.sort((a, b) => a.topRound.compareTo(b.topRound));

    positions.add(SizedBox(height: 24));

    positions.add(Center(
      child: Text(
        "Schnellste Runden",
        style: TextStyle(
          fontSize: 22,
          fontWeight: FontWeight.bold,
        ),
      ),
    ));

    positions.add(SizedBox(height: 16));

    positions.add(
      Center(
        child: SizedBox(
          width: width,
          child: Row(
            children: [
              stdBox(50, "#"),
              stdBox(160, "Pilot"),
              stdBox(80, "Round"),
            ],
          ),
        ),
      ),
    );

    positions.add(divider());

    var position = 1;

    for (var info in sorted) {
      final pilot = pilots[info.pilotUuid];
      if (pilot == null) continue;

      final nameTag = "${pilot.firstName.substring(0, 1)}. ${pilot.lastName}";

      positions.add(
        Center(
          child: SizedBox(
            width: width,
            child: Row(
              children: [
                stdBox(50, position.toString()),
                stdBox(160, nameTag),
                stdBox(80, info.topRound.toStringAsFixed(3)),
              ],
            ),
          ),
        ),
      );

      position++;
    }

    positions.add(divider());

    return positions;
  }

  List<Widget> buildTopSpeeds() {
    final List<Widget> positions = [];

    final List<Info> sorted = [];

    infos.forEach((track, info) {
      sorted.add(info);
    });

    sorted.sort((a, b) => b.topSpeed.compareTo(a.topSpeed));

    positions.add(SizedBox(height: 24));

    positions.add(Center(
      child: Text(
        "Top - Speed",
        style: TextStyle(
          fontSize: 22,
          fontWeight: FontWeight.bold,
        ),
      ),
    ));

    positions.add(SizedBox(height: 16));

    positions.add(
      Center(
        child: SizedBox(
          width: width,
          child: Row(
            children: [
              stdBox(50, "#"),
              stdBox(160, "Pilot"),
              stdBox(80, "Speed"),
            ],
          ),
        ),
      ),
    );

    positions.add(divider());

    var position = 1;

    for (var info in sorted) {
      final pilot = pilots[info.pilotUuid];
      if (pilot == null) continue;

      final nameTag = "${pilot.firstName.substring(0, 1)}. ${pilot.lastName}";

      positions.add(
        Center(
          child: SizedBox(
            width: width,
            child: Row(
              children: [
                stdBox(50, position.toString()),
                stdBox(160, nameTag),
                stdBox(80, info.topSpeed.toStringAsFixed(0)),
              ],
            ),
          ),
        ),
      );

      position++;
    }

    positions.add(divider());

    return positions;
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
      body: Column(
        children: [
          Center(
            child: Text(
              race.title,
              style: TextStyle(
                fontSize: 28,
                fontWeight: FontWeight.bold,
              ),
            ),
          ),
          Center(
            child: Text(
              "${race.rounds} Runden",
              style: TextStyle(
                fontSize: 22,
                fontWeight: FontWeight.bold,
              ),
            ),
          ),
          ...buildPositions(),
          ...buildTopRounds(),
          ...buildTopSpeeds(),
        ],
      ),
    );
  }
}
