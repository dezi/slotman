import 'package:flutter/material.dart';
import 'package:slotman/drawer.dart';
import 'package:slotman/status.dart';

class SetupPilotPage extends StatefulWidget {
  const SetupPilotPage({super.key});

  final String title = 'Pilot Setup';

  @override
  State<SetupPilotPage> createState() => SetupPilotPageState();
}

class SetupPilotPageState extends State<SetupPilotPage> {
  var pilot = Status.pilot.clone();

  var appUuidCtl = TextEditingController(text: Status.pilot.appUuid);
  var firstNameCtl = TextEditingController(text: Status.pilot.firstName);
  var lastNameCtl = TextEditingController(text: Status.pilot.lastName);
  var carModelCtl = TextEditingController(text: Status.pilot.car);
  var minSpeedCtl = TextEditingController(text: '${Status.pilot.minSpeed}');
  var maxSpeedCtl = TextEditingController(text: '${Status.pilot.maxSpeed}');

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
          width: 360,
          child: Column(children: [
            SizedBox(height: 16),
            TextField(
              controller: appUuidCtl,
              decoration: InputDecoration(
                labelText: 'App Uuid',
                border: OutlineInputBorder(),
              ),
              enabled: false,
            ),
            SizedBox(height: 16),
            TextField(
              controller: firstNameCtl,
              decoration: InputDecoration(
                labelText: 'First Name',
                hintText: 'Please enter your first name',
                border: OutlineInputBorder(),
              ),
              maxLength: 20,
              onChanged: (text) {
                setState(() {
                  pilot.firstName = text;
                });
              },
            ),
            TextField(
              controller: lastNameCtl,
              decoration: InputDecoration(
                labelText: 'Last Name',
                hintText: 'Please enter your last name',
                border: OutlineInputBorder(),
              ),
              maxLength: 20,
              onChanged: (text) {
                setState(() {
                  pilot.lastName = text;
                });
              },
            ),
            TextField(
              controller: carModelCtl,
              decoration: InputDecoration(
                labelText: 'Car Model',
                hintText: 'Please enter your car model',
                border: OutlineInputBorder(),
              ),
              maxLength: 20,
              onChanged: (text) {
                setState(() {
                  pilot.car = text;
                });
              },
            ),
            TextField(
              controller: minSpeedCtl,
              decoration: InputDecoration(
                labelText: 'Controller Min Speed %',
                hintText: 'Please enter controller min speed %',
                border: OutlineInputBorder(),
              ),
              maxLength: 3,
              onChanged: (text) {
                setState(() {
                  pilot.minSpeed = int.parse(text);
                });
              },
            ),
            TextField(
              controller: maxSpeedCtl,
              decoration: InputDecoration(
                labelText: 'Controller Max Speed %',
                hintText: 'Please enter controller max speed %',
                border: OutlineInputBorder(),
              ),
              maxLength: 3,
              onChanged: (text) {
                setState(() {
                  pilot.maxSpeed = int.parse(text);
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
                Status.sndPilot(pilot);
              },
              child: Text('Update'),
            )
          ]),
        ),
      ),
    );
  }
}
