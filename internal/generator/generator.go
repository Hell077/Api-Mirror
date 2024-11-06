package generator

import (
	"fmt"
	"github.com/Hell077/Api-Mirror/internal/parser"
	"os"
)

func Generator(config *parser.APIConfig, outputFileName string) error {
	file, err := os.Create(outputFileName)
	if err != nil {
		return fmt.Errorf("failed to create HTML file: %v", err)
	}
	defer file.Close()

	baseURL := fmt.Sprintf("http://%s:%s", config.APIMirror.SERVER, config.APIMirror.PORT)

	html := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>API Documentation</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            margin: 20px;
            background-color: #f9f9f9;
            color: #333;
        }
        h1 {
            color: #444;
            text-align: center;
            margin-bottom: 40px;
            font-size: 36px;
            font-weight: 700;
        }
        .api-container {
            background-color: #ffffff;
            padding: 20px;
            margin-bottom: 30px;
            border-radius: 8px;
            box-shadow: 0 6px 20px rgba(0, 0, 0, 0.1);
            display: flex;
            flex-direction: row;
            justify-content: flex-start;
            align-items: flex-start;
            position: relative;
        }
        .api-details {
            width: 65%;
            padding-right: 20px;
        }
        .api-title {
            font-size: 24px;
            font-weight: bold;
            color: #2c3e50;
            margin-bottom: 20px;
        }
        .api-address, .api-method {
            margin-bottom: 15px;
            font-size: 16px;
            color: #7f8c8d;
        }
        .api-form li {
            list-style: none;
            margin-bottom: 15px;
        }
        .api-form label {
            display: block;
            font-weight: bold;
            color: #34495e;
            font-size: 14px;
        }
        .api-form input[type="text"] {
            width: 100%;
            padding: 10px;
            margin-top: 5px;
            border: 1px solid #ddd;
            border-radius: 5px;
            box-sizing: border-box;
            background-color: #fafafa;
            font-size: 14px;
            transition: border-color 0.3s ease;
        }
        .api-form input[type="text"]:focus {
            border-color: #3498db;
        }
        .api-button {
            padding: 12px 20px;
            color: #fff;
            background-color: #3498db;
            border: none;
            border-radius: 5px;
            cursor: pointer;
            font-size: 16px;
            transition: background-color 0.3s ease;
            margin-top: 20px;
        }
        .api-button:hover {
            background-color: #2980b9;
        }
        .response-statuses {
            margin-top: 20px;
        }
        .response-statuses ul {
            padding-left: 20px;
        }
        .response-status {
            font-size: 14px;
            display: inline-block;
            margin: 5px 0;
            padding: 5px 10px;
            background-color: #ecf0f1;
            border-radius: 3px;
            color: #34495e;
        }
        .console-output {
            background-color: #1e1e1e;
            color: #d4d4d4;
            padding: 15px;
            border-radius: 8px;
            width: 30%;
            font-family: Consolas, monospace;
            white-space: pre-wrap;
            font-size: 14px;
            min-height: 180px;
            box-shadow: 0 4px 10px rgba(0, 0, 0, 0.15);
            margin-left: 20px;
        }
        .status-output {
            color: #ffcc00;
            font-weight: bold;
            margin-top: 10px;
            font-size: 14px;
        }
    </style>
</head>
<body>
<h1>API Documentation</h1>
`

	for index, api := range config.APIMirror.APIList {
		uniqueID := fmt.Sprintf("%s-%d", api.Title, index)

		html += `<div class="api-container">`
		html += `<div class="api-details">`
		html += fmt.Sprintf(`<div class="api-title">%s</div>`, api.Title)
		html += fmt.Sprintf(`<div class="api-address">Address: %s</div>`, api.Address)
		html += fmt.Sprintf(`<div class="api-method">Method: %s</div>`, api.Method)

		if len(api.Parameters) > 0 {
			html += `<div class="api-form"><strong>Parameters (in URL):</strong><form id="params-` + uniqueID + `">`
			for paramName, param := range api.Parameters {
				html += fmt.Sprintf(`
        <li><label for="%s">%s:</label></li>
        <li><label for="%s">Type: %s</label></li>
        <input type="text" id="%s" name="%s" placeholder="Enter %s" /></li>
        `, paramName, paramName, paramName, param.Type, paramName, paramName, paramName)
			}
			html += `</form></div>`
		}

		if len(api.Fields) > 0 {
			html += `<div class="api-form"><strong>Fields (in body):</strong><form id="fields-` + uniqueID + `">`
			for fieldName, field := range api.Fields {
				html += fmt.Sprintf(`
        <li><label for="%s">%s:</label></li>
        <li><label for="%s">Type: %s</label></li>
        <input type="text" id="%s" name="%s" placeholder="Enter %s" /></li>
        `, fieldName, fieldName, fieldName, field.Type, fieldName, fieldName, fieldName)
			}
			html += `</form></div>`
		}

		html += fmt.Sprintf(`<button class="api-button" type="button" onclick="sendRequest('%s', '%s', '%s')">Send Request</button>`, baseURL+api.Address, api.Method, uniqueID)

		html += `<div class="response-statuses"><strong>Response Statuses:</strong><ul>`
		for status, response := range api.Responses {
			html += fmt.Sprintf(`<li><span class="response-status">%d: %s</span></li>`, status, response)
		}
		html += `</ul></div>`
		html += `</div>`

		html += fmt.Sprintf(`<div class="console-output" id="console-%s">Console Output</div>`, uniqueID)

		html += `
		<script>
			function sendRequest(address, method, uniqueID) {
    const paramsForm = document.getElementById('params-' + uniqueID);
    const fieldsForm = document.getElementById('fields-' + uniqueID);

    const paramData = new FormData(paramsForm);
    let paramObject = {};
    paramData.forEach((value, key) => {
        paramObject[key] = value;
    });

    let url = address;
    for (let key in paramObject) {
        url = url.replace("{" + key + "}", paramObject[key]);
    }

    const fieldsData = new FormData(fieldsForm);
    let bodyObject = {};
    fieldsData.forEach((value, key) => {
        bodyObject[key] = value;
    });

    const consoleElement = document.getElementById('console-' + uniqueID);
    consoleElement.innerHTML = ''; // Очистка консоли перед новым выводом

    const statusElement = document.createElement('div');
    statusElement.classList.add('status-output');
    statusElement.textContent = 'Sending ' + method + ' request to ' + url;
    consoleElement.appendChild(statusElement);

    fetch(url, {
        method: method,
        headers: {
            "Content-Type": "application/json"
        },
        body: method !== "GET" ? JSON.stringify(bodyObject) : null
    })
    .then(response => response.json())
    .then(data => {
        const successElement = document.createElement('div');
        successElement.classList.add('status-output');
        successElement.textContent = 'Response: ' + JSON.stringify(data);
        consoleElement.appendChild(successElement);
    })
    .catch(error => {
        const errorElement = document.createElement('div');
        errorElement.classList.add('status-output');
        errorElement.textContent = 'Error: ' + error.message;
        consoleElement.appendChild(errorElement);
    });
}

		</script>
		`

		html += `</div>`
	}

	html += `
</body>
</html>`

	_, err = file.WriteString(html)
	if err != nil {
		return fmt.Errorf("failed to write HTML content: %v", err)
	}

	return nil
}
