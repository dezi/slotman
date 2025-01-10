class Pilot {
  final String what = 'pilot';
  final String mode;

  final String appUuid;
  final String firstName;
  final String lastName;
  final String carModel;
  final int minSpeed;
  final int maxSpeed;

  Pilot({
    required this.mode,
    required this.appUuid,
    required this.firstName,
    required this.lastName,
    required this.carModel,
    required this.minSpeed,
    required this.maxSpeed,
  });

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
}
