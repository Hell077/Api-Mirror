Here is the translated documentation in English and German.

---

# API Documentation for Api-Mirror

## 📋 Overview

**`Api-Mirror`** is a tool designed to generate and display interactive API documentation for any API. It supports testing API endpoints directly from the generated documentation, allowing you to send HTTP requests (GET, POST, etc.) and view responses in real time. The tool dynamically generates API forms based on a configuration file and displays query results in the console output.

---

## 🛠️ Features

- **Dynamic HTML Generation**: Automatically generates an interactive HTML interface based on a configuration file (YAML).
    - **Interactive Request Testing**: Allows sending HTTP requests directly from the documentation interface.
    - **Real-time Responses**: Displays response data and status in the console window.
    - **Field Masking**: Supports masking specific input fields.
    - **CORS Handling**: Displays detailed information about CORS issues.

---

## ⚙️ Installation

### Step 1: Download the Latest Version

To install **Api-Mirror**, download the latest version from the [GitHub Release](https://github.com/Hell077/Api-Mirror-/releases). Go to the link and select the appropriate file for your operating system (Windows, Linux, macOS).

### Step 2: Add to PATH (Windows)

To run the program from any command line, add the executable to your system's `PATH`.

1. **Find your Go bin directory** (where `api-mirror.exe` is located).
2. Add this directory to your `PATH`:
    - Right-click on "This PC" or "My Computer" and select **Properties**.
    - Choose **Advanced system settings**.
    - Click **Environment Variables**.
    - In the **System variables** section, find the `Path` variable, select it, and click **Edit**.
    - Add the path to the directory containing `api-mirror.exe` (e.g., `C:\path\to\Api-Mirror`).
    - Click **OK** to save the changes.

After this, you will be able to run `api-mirror` from any command line.

---

## 📝 Configuration File

`Api-Mirror` uses a configuration file in YAML format that defines details about the API (endpoints, methods, fields, and responses).

### Sample Configuration:

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

- `SERVER`: The server where your API is hosted (e.g., localhost).
    - `PORT`: The port where your API is accessible (e.g., 8080).
    - `APIList`: A list of API endpoints with their details.

Each endpoint contains:
- **Title**: The name of the API.
    - **Address**: The URL of the API endpoint.
    - **Method**: The HTTP method (GET, POST, etc.).
    - **Fields**: Input fields for the endpoint.
    - **Responses**: Possible HTTP responses with codes and descriptions.

---

## 💻 Usage

### Running the Application

Once you add the executable file `Api-Mirror` to your `PATH`, you can run it from the terminal. To generate API documentation, use the following command:

```bash
Mirror --path "/path/to/config.yaml" --port [optional, a free port will be chosen]
```

- `--config` (required): Path to the YAML configuration file.
    - `--output` (required): Path where the generated HTML documentation will be saved.

### Example:

```bash
Mirror --config api_config.yaml --output api_documentation.html
```

This command will generate HTML documentation based on the configuration file `api_config.yaml` and save it to the file `api_documentation.html`.

---

## 🔧 Flag Support

For more detailed customization, you can use additional flags. Here are some useful options:

- `--config <path>`: Path to the YAML configuration file.
    - `--output <path>`: Path to save the HTML documentation.
    - `--help`: Show the list of all available flags.

---

## 📌 Notes

- Ensure that the configuration file is properly set up before running.
    - You can use various flags and command-line parameters to customize the program's behavior.
    - If you encounter any errors or issues, contact support at the [Api-Mirror Issues](https://github.com/Hell077/Api-Mirror/issues) repository.

---

## 🔗 Links

- **Releases**: [Download the latest version](https://github.com/Hell077/Api-Mirror-/releases)
- **GitHub Repository**: [https://github.com/Hell077/Api-Mirror](https://github.com/Hell077/Api-Mirror-)


