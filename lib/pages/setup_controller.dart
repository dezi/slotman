import 'package:flutter/material.dart';
import 'package:slotman/drawer.dart';
import 'package:slotman/locals.dart';
import 'package:slotman/status.dart';

class SetupControllerPage extends StatefulWidget {
  const SetupControllerPage({super.key});

  final String title = 'Controller Setup';

  @override
  State<SetupControllerPage> createState() => _SetupControllerPageState();
}

class _SetupControllerPageState extends State<SetupControllerPage> {
  int selectedController = 1;
  int numControllers = Status.numberOfTracks;

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
          child: DecoratedBox(
            decoration: const BoxDecoration(
                color: Colors.red
            ),
            child: Column(children: [
              SizedBox(height: 16),
              for (int controller = 1; controller <= numControllers; controller++)
                RadioListTile<int>(
                  title: Text('Controller $controller'),
                  value: controller,
                  groupValue: selectedController,
                  onChanged: (int? value) {
                    setState(() {
                      selectedController = value!;
                    });
                  },
                ),
              SizedBox(height: 16),
              Text('123'),
              Text('544545545454545454366'),
            ]),
          ),
        ),
      ),
    );
  }
}
