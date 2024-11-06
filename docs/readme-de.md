
# API-Dokumentation für Api-Mirror

## 📋 Übersicht

**`Api-Mirror`** ist ein Tool, das zur Generierung und Anzeige interaktiver API-Dokumentation für jede API dient. Es unterstützt das Testen von API-Endpunkten direkt aus der generierten Dokumentation, sodass HTTP-Anfragen (GET, POST usw.) gesendet und die Antworten in Echtzeit angezeigt werden können. Das Tool generiert dynamisch API-Formulare basierend auf einer Konfigurationsdatei und zeigt die Abfrageergebnisse in der Konsole an.

---

## 🛠️ Funktionen

- **Dynamische HTML-Generierung**: Generiert automatisch eine interaktive HTML-Oberfläche basierend auf einer Konfigurationsdatei (YAML).
    - **Interaktives Testen von Anfragen**: Ermöglicht das Senden von HTTP-Anfragen direkt aus der Dokumentationsoberfläche.
    - **Echtzeit-Antworten**: Zeigt Antwortdaten und Status im Konsolenfenster an.
    - **Feldmaskierung**: Unterstützt die Maskierung bestimmter Eingabefelder.
    - **CORS-Handling**: Zeigt detaillierte Informationen zu CORS-Problemen an.

---

## ⚙️ Installation

### Schritt 1: Die neueste Version herunterladen

Um **Api-Mirror** zu installieren, laden Sie die neueste Version von der [GitHub-Veröffentlichung](https://github.com/Hell077/Api-Mirror-/releases) herunter. Gehen Sie zum Link und wählen Sie die entsprechende Datei für Ihr Betriebssystem (Windows, Linux, macOS) aus.

### Schritt 2: Hinzufügen zum PATH (Windows)

Um das Programm aus jeder Eingabeaufforderung auszuführen, fügen Sie die ausführbare Datei zum System-`PATH` hinzu.

1. **Finden Sie Ihr Go-Bin-Verzeichnis** (wo `api-mirror.exe` gespeichert ist).
2. Fügen Sie dieses Verzeichnis zu Ihrem `PATH` hinzu:
    - Klicken Sie mit der rechten Maustaste auf "Dieser PC" oder "Mein Computer" und wählen Sie **Eigenschaften**.
    - Wählen Sie **Erweiterte Systemeinstellungen**.
    - Klicken Sie auf **Umgebungsvariablen**.
    - Suchen Sie im Abschnitt **Systemvariablen** nach der Variablen `Path`, wählen Sie sie aus und klicken Sie auf **Bearbeiten**.
    - Fügen Sie den Pfad zum Verzeichnis hinzu, das `api-mirror.exe` enthält (z. B. `C:\path\to\Api-Mirror`).
    - Klicken Sie auf **OK**, um die Änderungen zu speichern.

Danach können Sie `api-mirror` aus jeder Eingabeaufforderung heraus ausführen.

---

## 📝 Konfigurationsdatei

`Api-Mirror` verwendet eine Konfigurationsdatei im YAML-Format, die Details zur API (Endpunkte, Methoden, Felder und Antworten) definiert.

### Beispielkonfiguration:

```yaml
API_MIRROR:
  SERVER: "localhost"
  PORT: "5000"
  API_LIST:
    API_Name:
      address: "/api/example"
      method: "POST"
      fields:
        name:
          type: "string"
          mask: "Anna"
        age:
          type: "int"
          mask: "20"
      responses:
        200: "OK"
        201: "Created"
        400: "Bad Request"
        500: "Internal Server Error"
      title: "Some Title"
    API_Name2:
      address: "/api/example2"
      method: "GET"
      responses:
        200: "OK"
        202: "Accepted"
        403: "Forbidden"
      title: "Another Title"
    API_Name3:
      address: "/api/example3"
      method: "POST"
      fields:
        name:
          type: "string"
          mask: "Anna"
        age:
          type: "int"
          mask: "20"
      responses:
        200: "OK"
        201: "Created"
        400: "Bad Request"
        500: "Internal Server Error"
      title: "Some Title"
```

- `SERVER`: Der Server, auf dem Ihre API gehostet wird (z. B. localhost).
    - `PORT`: Der Port, auf dem Ihre API zugänglich ist (z. B. 8080).
    - `APIList`: Eine Liste von API-Endpunkten mit ihren Details.

Jeder Endpunkt enthält:
- **Title**: Der Name der API.
    - **Address**: Die URL des API-Endpunkts.
    - **Method**: Die HTTP-Methode (GET, POST usw.).
    - **Fields**: Eingabefelder für den Endpunkt.
    - **Responses**: Mögliche HTTP-Antworten mit Codes und Beschreibungen

.

---

## 💻 Verwendung

### Anwendung starten

Sobald Sie die ausführbare Datei `Api-Mirror` zu Ihrem `PATH` hinzugefügt haben, können Sie sie über die Konsole ausführen. Um die API-Dokumentation zu generieren, verwenden Sie den folgenden Befehl:

```bash
Mirror --path "/path/to/config.yaml" --port [optional, ein freier Port wird gewählt]
```

- `--config` (erforderlich): Pfad zur YAML-Konfigurationsdatei.
    - `--output` (erforderlich): Pfad, in dem die generierte HTML-Dokumentation gespeichert wird.

### Beispiel:

```bash
Mirror --config api_config.yaml --output api_documentation.html
```

Dieser Befehl generiert eine HTML-Dokumentation basierend auf der Konfigurationsdatei `api_config.yaml` und speichert sie in der Datei `api_documentation.html`.

---

## 🔧 Flaggen-Unterstützung

Für detailliertere Anpassungen können zusätzliche Flags verwendet werden. Hier sind einige nützliche Optionen:

- `--config <Pfad>`: Pfad zur YAML-Konfigurationsdatei.
    - `--output <Pfad>`: Pfad, in dem die HTML-Dokumentation gespeichert wird.
    - `--help`: Zeigt die Liste aller verfügbaren Flags an.

---

## 📌 Hinweise

- Stellen Sie sicher, dass die Konfigurationsdatei ordnungsgemäß eingerichtet ist, bevor Sie den Server ausführen.
    - Sie können verschiedene Flags und Befehlszeilenparameter verwenden, um das Verhalten des Programms anzupassen.
    - Wenn Sie auf Fehler oder Probleme stoßen, wenden Sie sich an den Support auf der [Api-Mirror Issues](https://github.com/Hell077/Api-Mirror/issues) Repository.

---

## 🔗 Links

- **Releases**: [Laden Sie die neueste Version herunter](https://github.com/Hell077/Api-Mirror-/releases)
- **GitHub-Repository**: [https://github.com/Hell077/Api-Mirror](https://github.com/Hell077/Api-Mirror-)


