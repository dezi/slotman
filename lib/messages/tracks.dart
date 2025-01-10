class Tracks {
  final String what = 'tracks';
  String mode;

  int tracks;

  Tracks({
    this.mode = 'set',
    this.tracks = 0,
  });

  Tracks.fromJson(Map<String, dynamic> json)
      : mode = json['mode'] as String,
        tracks = json['tracks'] as int;

  Map<String, dynamic> toJson() => {
        'what': what,
        'mode': mode,
        'tracks': tracks,
      };

  Tracks clone() {
    return Tracks.fromJson(toJson());
  }
}
