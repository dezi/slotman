class Info {
  final String what = 'info';
  String mode;

  int track;
  int rounds;
  int position;

  double actRound;
  double topRound;
  double actSpeed;
  double topSpeed;

  String pilotUuid;

  Info({
    this.mode = 'set',
    this.track = 0,
    this.rounds = 0,
    this.position = 0,
    this.actRound = 0,
    this.topRound = 0,
    this.actSpeed = 0,
    this.topSpeed = 0,
    this.pilotUuid = '',
  });

  Info.fromJson(Map<String, dynamic> json)
      : mode = json['mode'] as String,
        track = json['track'] as int,
        rounds = json['rounds'] as int,
        position = json['position'] as int,
        actRound = json['actRound'] as double,
        topRound = json['topRound'] as double,
        actSpeed = json['actSpeed'] as double,
        topSpeed = json['topSpeed'] as double,
        pilotUuid = json['pilotUuid'] as String;

  Map<String, dynamic> toJson() => {
        'what': what,
        'mode': mode,
        'track': track,
        'rounds': rounds,
        'position': position,
        'actRound': actRound,
        'topRound': topRound,
        'actSpeed': actSpeed,
        'topSpeed': topSpeed,
        'pilotUuid': pilotUuid,
      };

  Info clone() {
    return Info.fromJson(toJson());
  }
}
