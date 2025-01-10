import 'package:shared_preferences/shared_preferences.dart';
import 'package:slotman/messages/pilot.dart';
import 'package:slotman/status.dart';
import 'package:uuid/uuid.dart';

class Locals {
  static SharedPreferences? prefs;

  static String appUuid = '';

  static String pilotFirstName = '';
  static String pilotLastName = '';
  static String pilotCarModel = '';

  static int pilotMinSpeed = 0;
  static int pilotMaxSpeed = 100;

  static Future<void> initialize() async {
    prefs = await SharedPreferences.getInstance();

    if (prefs?.getString('appUuid') == null) {
      saveAppUuid(Uuid().v4());

      savePilotFirstName('Max');
      savePilotLastName('Versägen');
      savePilotCarModel('Holzmotor 2000');

      savePilotMinSpeed(0);
      savePilotMaxSpeed(100);
    }

    appUuid = prefs?.getString('appUuid') ?? '';

    pilotFirstName = prefs?.getString('pilotFirstName') ?? '';
    pilotLastName = prefs?.getString('pilotLastName') ?? '';
    pilotCarModel = prefs?.getString('pilotCarModel') ?? '';

    pilotMinSpeed = prefs?.getInt('pilotMinSpeed') ?? 0;
    pilotMaxSpeed = prefs?.getInt('pilotMaxSpeed') ?? 0;

    sndPilot();
  }

  static void saveAppUuid(String val) {
    prefs?.setString('appUuid', appUuid = val);
  }

  static void savePilotFirstName(String val) {
    prefs?.setString('pilotFirstName', pilotFirstName = val);
  }

  static void savePilotLastName(String val) {
    prefs?.setString('pilotLastName', pilotLastName = val);
  }

  static void savePilotCarModel(String val) {
    prefs?.setString('pilotCarModel', pilotCarModel = val);
  }

  static void savePilotMinSpeed(int percent) {
    if (percent < 0) percent = 0;
    if (percent > 100) percent = 100;
    prefs?.setInt('pilotMinSpeed', pilotMinSpeed = percent);
  }

  static void savePilotMaxSpeed(int percent) {
    if (percent < 0) percent = 0;
    if (percent > 100) percent = 100;
    prefs?.setInt('pilotMaxSpeed', pilotMaxSpeed = percent);
  }

  static void sndPilot() {
    var pilot = Pilot(
      appUuid: appUuid,
      firstName: pilotFirstName,
      lastName: pilotLastName,
      carModel: pilotCarModel,
      minSpeed: pilotMinSpeed,
      maxSpeed: pilotMaxSpeed,
    );

    Status.sndPilot(pilot);
  }
}
