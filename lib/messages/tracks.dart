class Tracks {
  final String what = 'tracks';
  final String mode;

  final int numberOfTracks;

  Tracks({required this.mode, this.numberOfTracks = 0});

  Tracks.fromJson(Map<String, dynamic> json)
      : mode = json['mode'] as String,
        numberOfTracks = json['numberOfTracks'] as int;

  Map<String, dynamic> toJson() => {
        'what': what,
        'mode': mode,
        'numberOfTracks': numberOfTracks,
      };
}
