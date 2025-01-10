import 'package:flutter/material.dart';
import 'package:slotman/drawer.dart';

class SetupRacePage extends StatefulWidget {
  const SetupRacePage({super.key});

  final String title = 'Race Setup';

  @override
  State<SetupRacePage> createState() => _SetupRacePageState();
}

class _SetupRacePageState extends State<SetupRacePage> {

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        backgroundColor: Theme.of(context).colorScheme.inversePrimary,
        title: Text(widget.title),
      ),
      drawer: MainDrawer(),
      body: Center(
        // Center is a layout widget. It takes a single child and positions it
        // in the middle of the parent.
        child: const Text(
              'Setup Race...',
            ),
        ),
      );
  }
}
