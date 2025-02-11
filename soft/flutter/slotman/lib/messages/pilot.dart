class Pilot {
  final String what = 'pilot';
  String mode;

  String uuid;

  String appUuid;
  String firstName;
  String lastName;

  String team;
  String car;

  int minSpeed;
  int maxSpeed;

  Pilot({
     this.mode = 'set',
     this.uuid = '',
     this.appUuid = '',
     this.firstName = '',
     this.lastName = '',
     this.team = '',
     this.car = '',
     this.minSpeed = 0,
     this.maxSpeed = 0,
  });

  Pilot.fromJson(Map<String, dynamic> json)
      : mode = json['mode'] as String,
        uuid = json['uuid'] as String,
        appUuid = json['appUuid'] as String,
        firstName = json['firstName'] as String,
        lastName = json['lastName'] as String,
        team = json['team'] as String,
        car = json['car'] as String,
        minSpeed = json['minSpeed'] as int,
        maxSpeed = json['maxSpeed'] as int;

  Map<String, dynamic> toJson() => {
        'what': what,
        'mode': mode,
        'uuid': uuid,
        'appUuid': appUuid,
        'firstName': firstName,
        'lastName': lastName,
        'team': team,
        'car': car,
        'minSpeed': minSpeed,
        'maxSpeed': maxSpeed,
      };

  Pilot clone() {
    return Pilot.fromJson(toJson());
  }
}
