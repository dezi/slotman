class Message {
  final String what;
  final String mode;

  Message({
    this.what = '',
    this.mode = '',
  });

  Map<String, dynamic> toJson() => {
    'what': what,
    'mode': mode,
  };

  Message.fromJson(Map<String, dynamic> json)
      : what = json['what'] as String,
        mode = json['mode'] as String;

  String tag() {
    return "$what|$mode";
  }
}
