import 'package:flutter/material.dart';
import 'package:slotman/drawer.dart';
import 'package:slotman/locals.dart';

class SetupPilotPage extends StatefulWidget {
  const SetupPilotPage({super.key});

  final String title = 'Pilot Setup';

  @override
  State<SetupPilotPage> createState() => _SetupPilotPageState();
}

class _SetupPilotPageState extends State<SetupPilotPage> {
  TextEditingController appUuidController = TextEditingController(text: Locals.appUuid);
  TextEditingController firstNameController = TextEditingController(text: Locals.pilotFirstName);
  TextEditingController lastNameController = TextEditingController(text: Locals.pilotLastName);
  TextEditingController carModelController = TextEditingController(text: Locals.pilotCarModel);
  TextEditingController minSpeedController =
      TextEditingController(text: '${Locals.pilotMinSpeed}');
  TextEditingController maxSpeedController =
      TextEditingController(text: '${Locals.pilotMaxSpeed}');

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
        child: SizedBox(
          width: 360,
          child: Column(children: [
            SizedBox(height: 16),
            TextField(
              controller: appUuidController,
              decoration: InputDecoration(
                labelText: 'App Uuid',
                border: OutlineInputBorder(),
              ),
              enabled: false,
            ),
            SizedBox(height: 16),
            TextField(
              controller: firstNameController,
              decoration: InputDecoration(
                labelText: 'First Name',
                hintText: 'Please enter your first name',
                border: OutlineInputBorder(),
              ),
              maxLength: 20,
              onChanged: (text) {
                Locals.savePilotFirstName(text);
              },
            ),
            TextField(
              controller: lastNameController,
              decoration: InputDecoration(
                labelText: 'Last Name',
                hintText: 'Please enter your last name',
                border: OutlineInputBorder(),
              ),
              maxLength: 20,
              onChanged: (text) {
                Locals.savePilotLastName(text);
              },
            ),
            TextField(
              controller: carModelController,
              decoration: InputDecoration(
                labelText: 'Car Model',
                hintText: 'Please enter your car model',
                border: OutlineInputBorder(),
              ),
              maxLength: 20,
              onChanged: (text) {
                Locals.savePilotCarModel(text);
              },
            ),
            TextField(
              controller: minSpeedController,
              decoration: InputDecoration(
                labelText: 'Controller Min Speed %',
                hintText: 'Please enter controller min speed %',
                border: OutlineInputBorder(),
              ),
              maxLength: 3,
              onChanged: (text) {
                setState(() {
                  Locals.savePilotMinSpeed(int.parse(text));
                });
              },
            ),
            TextField(
              controller: maxSpeedController,
              decoration: InputDecoration(
                labelText: 'Controller Max Speed %',
                hintText: 'Please enter controller max speed %',
                border: OutlineInputBorder(),
              ),
              maxLength: 3,
              onChanged: (text) {
                setState(() {
                  Locals.savePilotMaxSpeed(int.parse(text));
                });
              },
            ),
          ]),
        ),
      ),
    );
  }
}
