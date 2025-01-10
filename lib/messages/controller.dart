class Controller {
  final String what = 'controller';
  final String mode;

  final int controller;
  final bool isCalibrating;

  final int minValue;
  final int maxValue;

  Controller({
    required this.mode,
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
      };
}
