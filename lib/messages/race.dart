class Race {
  final String what = 'race';
  final String mode;

  String title;
  int tracks;
  int rounds;

  Race({
    required this.mode,
    this.title = '',
    this.tracks = 0,
    this.rounds = 0,
  });

  Race.fromJson(Map<String, dynamic> json)
      : mode = json['mode'] as String,
        title = json['title'] as String,
        tracks = json['tracks'] as int,
        rounds = json['rounds'] as int;

  Map<String, dynamic> toJson() => {
        'what': what,
        'mode': mode,
        'title': title,
        'tracks': tracks,
        'rounds': rounds,
      };
}
