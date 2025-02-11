import 'package:flutter/material.dart';
import 'package:slotman/status.dart';

class MainDrawer extends StatelessWidget {
  const MainDrawer({super.key});

  @override
  Widget build(BuildContext context) {

    String initials = "";
    if (Status.pilot.firstName.isNotEmpty) initials += Status.pilot.firstName[0];
    if (Status.pilot.lastName.isNotEmpty) initials += Status.pilot.lastName[0];

    return Drawer(
      child: ListView(
        padding: EdgeInsets.zero,
        children: <Widget>[
          UserAccountsDrawerHeader(
            accountName: Text("${Status.pilot.firstName} ${Status.pilot.lastName}"),
            accountEmail: Text(Status.pilot.car),
            currentAccountPicture: CircleAvatar(
              backgroundColor: Colors.orange,
              child: Text(
                initials,
                style: TextStyle(fontSize: 40.0),
              ),
            ),
          ),
          ListTile(
            leading: Icon(Icons.account_circle),
            title: Text('Start Page'),
            onTap: () {
              // Update the state of the app

              // Then close the drawer

              Navigator.pop(context);
              Navigator.pushNamed(context, '/');
            },
          ),
          ListTile(
            leading: Icon(Icons.join_inner),
            title: Text('Join Race'),
            onTap: () {
              Navigator.pop(context);
              Navigator.pushNamed(context, '/join');
            },
          ),
          ListTile(
            leading: Icon(Icons.track_changes),
            title: Text('Race Setup'),
            onTap: () {
              Navigator.pop(context);
              Navigator.pushNamed(context, '/setup/race');
            },
          ),
          ListTile(
            leading: Icon(Icons.account_circle),
            title: Text('Pilot Setup'),
            onTap: () {
              Navigator.pop(context);
              Navigator.pushNamed(context, '/setup/pilot');
            },
          ),
          ListTile(
            leading: Icon(Icons.waves),
            title: Text('Track Setup'),
            onTap: () {
              Navigator.pop(context);
              Navigator.pushNamed(context, '/setup/track');
            },
          ),
          ListTile(
            leading: Icon(Icons.barcode_reader),
            title: Text('Controller Setup'),
            onTap: () {
              Navigator.pop(context);
              Navigator.pushNamed(context, '/setup/controller');
            },
          ),
        ],
      ),
    );
  }
}
