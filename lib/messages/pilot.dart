class Pilot {
  final String what = 'pilot';
  String mode;

  String appUuid;
  String firstName;
  String lastName;
  String carModel;
  int minSpeed;
  int maxSpeed;

  Pilot({
     this.mode = 'set',
     this.appUuid = '',
     this.firstName = '',
     this.lastName = '',
     this.carModel = '',
     this.minSpeed = 0,
     this.maxSpeed = 0,
  });

  Pilot.fromJson(Map<String, dynamic> json)
      : mode = json['mode'] as String,
        appUuid = json['appUuid'] as String,
        firstName = json['firstName'] as String,
        lastName = json['lastName'] as String,
        carModel = json['carModel'] as String,
        minSpeed = json['minSpeed'] as int,
        maxSpeed = json['maxSpeed'] as int;

  Map<String, dynamic> toJson() => {
        'what': what,
        'mode': mode,
        'appUuid': appUuid,
        'firstName': firstName,
        'lastName': lastName,
        'carModel': carModel,
        'minSpeed': minSpeed,
        'maxSpeed': maxSpeed,
      };

  Pilot clone() {
    return Pilot.fromJson(toJson());
  }
}
