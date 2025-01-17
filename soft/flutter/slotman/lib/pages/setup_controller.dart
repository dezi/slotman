import 'package:flutter/material.dart';
import 'package:slotman/drawer.dart';
import 'package:slotman/status.dart';

class SetupControllerPage extends StatefulWidget {
  const SetupControllerPage({super.key});

  final String title = 'Controller Setup';

  @override
  State<SetupControllerPage> createState() => SetupControllerPageState();
}

class SetupControllerPageState extends State<SetupControllerPage> {

  static SetupControllerPageState? injector;

  var controller = Status.controller.clone();

  void setContent() {
    setState(() {
      controller = Status.controller.clone();
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
          child: DecoratedBox(
            decoration: const BoxDecoration(color: Colors.transparent),
            child: Column(children: [
              SizedBox(height: 16),
              for (int count = 1; count <= Status.tracks.tracks; count++)
                RadioListTile<int>(
                  title: Text('Controller $count'),
                  value: count,
                  groupValue: controller.controller,
                  onChanged: (int? value) {
                    setState(() {
                      controller.controller = value!;
                    });
                  },
                ),
              SizedBox(height: 16),
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
                  setState(() {
                    controller.isCalibrating = !controller.isCalibrating;
                    Status.sndController(controller);
                  });
                },
                child: Text(controller.isCalibrating ? 'Calibrating...' : 'Calibrate'),
              ),
              SizedBox(height: 16),
              if (controller.isCalibrating)
                Row(
                  children: [
                    SizedBox(width: 40, child: Text('Min')),
                    Container(
                        padding: const EdgeInsets.symmetric(vertical: 4, horizontal: 4),
                        margin: const EdgeInsets.symmetric(vertical: 0, horizontal: 8),
                        decoration: BoxDecoration(border: Border.all(width: 2)),
                        child: SizedBox(
                            width: 120,
                            child: Text(
                              '${controller.minValue}',
                              textAlign: TextAlign.center,
                              style: TextStyle(fontWeight: FontWeight.bold),
                            ))),
                  ],
                ),
              SizedBox(height: 16),
              if (controller.isCalibrating)
                Row(
                  children: [
                    SizedBox(width: 40, child: Text('Max')),
                    Container(
                        padding: const EdgeInsets.symmetric(vertical: 4, horizontal: 4),
                        margin: const EdgeInsets.symmetric(vertical: 0, horizontal: 8),
                        decoration: BoxDecoration(border: Border.all(width: 2)),
                        child: SizedBox(
                            width: 120,
                            child: Text(
                              '${controller.maxValue}',
                              textAlign: TextAlign.center,
                              style: TextStyle(fontWeight: FontWeight.bold),
                            ))),
                  ],
                ),
            ]),
          ),
        ),
      ),
    );
  }
}
