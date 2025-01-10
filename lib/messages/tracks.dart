class Tracks {
  final String what = 'tracks';
  String mode;

  int tracks;

  Tracks({required this.mode, this.tracks = 0});

  Tracks.clone(Tracks tracks)
      : mode = tracks.mode,
        tracks = tracks.tracks;

  Tracks.fromJson(Map<String, dynamic> json)
      : mode = json['mode'] as String,
        tracks = json['tracks'] as int;

  Map<String, dynamic> toJson() => {
        'what': what,
        'mode': mode,
        'tracks': tracks,
      };
}
