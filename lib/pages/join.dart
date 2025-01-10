import 'package:flutter/material.dart';
import 'package:slotman/drawer.dart';

class JoinPage extends StatefulWidget {
  const JoinPage({super.key});

  final String title = 'Join Race';

  @override
  State<JoinPage> createState() => _JoinPageState();
}

class _JoinPageState extends State<JoinPage> {

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        backgroundColor: Theme.of(context).colorScheme.inversePrimary,
        title: Text(widget.title),
      ),
      drawer: MainDrawer(),
      body: Center(
        child: const Text(
              'Join Race...',
            ),
        ),
      );
  }
}
