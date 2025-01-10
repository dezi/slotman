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

  bool isCalibrating = false;

  int selected = 1;
  int numControllers = Status.numberOfTracks;

  int minValue = 0;
  int maxValue = 0;

  void setMinMaxValue(int selected, bool isCalibrating, int minValue, int maxValue) {
    setState(() {
      this.selected = selected;
      this.isCalibrating = isCalibrating;
      this.minValue = minValue;
      this.maxValue = maxValue;
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
              for (int controller = 1; controller <= numControllers; controller++)
                RadioListTile<int>(
                  title: Text('Controller $controller'),
                  value: controller,
                  groupValue: selected,
                  onChanged: (int? value) {
                    setState(() {
                      selected = value!;
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
                    isCalibrating = !isCalibrating;
                    Status.sndCalibrateController(selected, isCalibrating);
                  });
                },
                child: Text(isCalibrating ? 'Calibrating...' : 'Calibrate'),
              ),
              SizedBox(height: 16),
              if (isCalibrating)
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
                              '$minValue',
                              textAlign: TextAlign.center,
                              style: TextStyle(fontWeight: FontWeight.bold),
                            ))),
                  ],
                ),
              SizedBox(height: 16),
              if (isCalibrating)
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
                              '$maxValue',
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
