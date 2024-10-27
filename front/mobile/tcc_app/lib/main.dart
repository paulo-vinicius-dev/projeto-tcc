import 'package:flutter/material.dart';
import 'package:google_fonts/google_fonts.dart';

final theme = ThemeData(
  useMaterial3: true,
  colorScheme: ColorScheme.fromSeed(
    seedColor: Colors.black12,
    brightness: Brightness.dark,
  ),
  textTheme: GoogleFonts.latoTextTheme(),
);

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      home: Scaffold(
        body: Column(
          crossAxisAlignment: CrossAxisAlignment.center,
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            Padding(
              padding: const EdgeInsets.fromLTRB(20, 20, 20, 20),
              child: TextFormField(
                  decoration: const InputDecoration(
                border: OutlineInputBorder(),
                labelText: 'Digite o seu E-mail',
                suffixIcon: Icon(Icons.person),
                labelStyle: TextStyle(color: Colors.black87),
                floatingLabelStyle: TextStyle(color: Colors.green),
                focusedBorder: OutlineInputBorder(
                    borderSide: BorderSide(color: Colors.green)),
              )),
            ),
            Padding(
              padding: const EdgeInsets.fromLTRB(20, 20, 20, 20),
              child: TextFormField(
                  obscureText: true,
                  decoration: const InputDecoration(
                    border: OutlineInputBorder(),
                    labelText: 'Digite a sua Senha',
                    suffixIcon: Icon(Icons.lock),
                    labelStyle: TextStyle(color: Colors.black87),
                    floatingLabelStyle: TextStyle(color: Colors.green),
                    focusedBorder: OutlineInputBorder(
                      borderSide: BorderSide(color: Colors.green),
                    ),
                  )),
            ),
            Padding(
              padding: const EdgeInsets.fromLTRB(20, 20, 20, 20),
              child: ElevatedButton.icon(
                iconAlignment: IconAlignment.end,
                icon: const Icon(
                  Icons.login,
                  color: Colors.white,
                  size: 24.0,
                ),
                label: const Text("Acessar"),
                onPressed: () {},
                
                style: ButtonStyle(
                  backgroundColor: const WidgetStatePropertyAll(Colors.green),
                  foregroundColor: const WidgetStatePropertyAll(Colors.white),
                  shape: WidgetStatePropertyAll(RoundedRectangleBorder(
                      borderRadius: BorderRadius.circular(2),
                      side: const BorderSide(color: Colors.black12))),
                  minimumSize: const WidgetStatePropertyAll(Size(500.0, 50.0)),
                ),
              ),
            ),
          ],
        ),
      ),
    );
  }
}
