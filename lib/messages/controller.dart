class Controller {
  final String what = 'controller';
  String mode;

  int controller;
  bool isCalibrating;

  int minValue;
  int maxValue;

  Controller({
    this.mode = 'set',
    this.controller = 0,
    this.isCalibrating = false,
    this.minValue = 0,
    this.maxValue = 0,
  });

  Controller.fromJson(Map<String, dynamic> json)
      : mode = json['mode'] as String,
        controller = json['controller'] as int,
        isCalibrating = json['isCalibrating'] as bool,
        minValue = json['minValue'] as int,
        maxValue = json['maxValue'] as int;

  Map<String, dynamic> toJson() => {
        'what': what,
        'mode': mode,
        'controller': controller,
        'isCalibrating': isCalibrating,
        'minValue': minValue,
        'maxValue': maxValue,
      };

  Controller clone() {
    return Controller.fromJson(toJson());
  }
}
