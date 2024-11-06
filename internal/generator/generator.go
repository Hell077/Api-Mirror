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

	// Формируем базовый URL на основе SERVER и PORT из YAML
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
            background-color: #f7f9fc;
            color: #333;
            line-height: 1.6;
            margin: 0;
            padding: 20px;
        }
        h1 {
            color: #2c3e50;
            font-size: 2em;
            margin-bottom: 1em;
            text-align: center;
        }
        .api-container {
            background-color: #ffffff;
            border-radius: 8px;
            box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
            padding: 20px;
            margin-bottom: 20px;
            transition: box-shadow 0.3s ease;
        }
        .api-container:hover {
            box-shadow: 0 6px 12px rgba(0, 0, 0, 0.15);
        }
        .api-title {
            font-size: 1.75em;
            color: #34495e;
            margin-bottom: 0.5em;
        }
        .api-address, .api-method {
            font-weight: bold;
            color: #16a085;
            margin-top: 0.5em;
        }
        .response-statuses {
            margin-top: 1em;
        }
        .response-statuses strong {
            font-size: 1.1em;
            color: #34495e;
        }
        .response-status {
            font-size: 0.95em;
            margin-left: 10px;
            color: #2980b9;
        }
        ul {
            list-style-type: none;
            padding: 0;
        }
        li {
            margin: 5px 0;
        }
        .api-form input {
            padding: 5px;
            margin: 5px;
            border-radius: 5px;
            border: 1px solid #ccc;
        }
        .api-form button {
            background-color: #2980b9;
            color: white;
            border: none;
            padding: 10px 20px;
            border-radius: 5px;
            cursor: pointer;
        }
        .api-form button:hover {
            background-color: #3498db;
        }
        .console-output {
            background-color: #2d3436;
            color: #ecf0f1;
            font-family: 'Courier New', Courier, monospace;
            padding: 10px;
            border-radius: 5px;
            margin-top: 10px;
            height: 150px;
            overflow-y: auto;
            white-space: pre-wrap;
        }
        .status-output {
            font-size: 1.2em;
            margin-top: 10px;
            font-weight: bold;
        }
    </style>
</head>
<body>
<h1>API Documentation</h1>
`

	// Перебираем список API из конфигурации
	for _, api := range config.APIMirror.APIList {
		html += `<div class="api-container">`
		html += fmt.Sprintf(`<div class="api-title">%s</div>`, api.Title)
		html += fmt.Sprintf(`<div class="api-address">Address: %s</div>`, api.Address)
		html += fmt.Sprintf(`<div class="api-method">Method: %s</div>`, api.Method)

		// Для параметров пути, таких как {user_id}, добавляем поле ввода
		if len(api.Parameters) > 0 {
			html += `<div class="api-form"><strong>Parameters:</strong><form id="params-` + api.Title + `">`
			for param, details := range api.Parameters {
				html += fmt.Sprintf(` 
					<li><label for="%s">%s:</label>
					<input type="text" id="%s" name="%s" value="%s" placeholder="%s" /></li>`,
					param, param, param, param, details.Placeholder, details.Placeholder)
			}
			html += fmt.Sprintf(`</form>`)
		}

		html += fmt.Sprintf(`<button type="button" onclick="sendRequest('%s', '%s', '%s')">Send Request</button>`, baseURL+api.Address, api.Method, api.Title)

		html += `<div class="response-statuses"><strong>Response Statuses:</strong><ul>`
		for status, response := range api.Responses {
			html += fmt.Sprintf(`<li><span class="response-status">%d: %s</span></li>`, status, response)
		}
		html += `</ul></div>`

		html += fmt.Sprintf(`<div class="status-output" id="status-%s">Status: Awaiting Response</div>`, api.Title)

		html += fmt.Sprintf(`<div class="console-output" id="console-%s"></div>`, api.Title)

		html += `	
		<script>
			function sendRequest(address, method, title) {
				// Получаем форму по уникальному id
				const form = document.getElementById('params-' + title);
				const formData = new FormData(form);
				let data = {};
				formData.forEach((value, key) => {
					data[key] = value;
				});
				
				// Очистка консоли перед новым запросом
				const consoleElement = document.getElementById('console-' + title);
				const statusElement = document.getElementById('status-' + title);
				consoleElement.textContent = 'Sending request...';
				statusElement.textContent = 'Status: Sending...';
				
				// Отправка запроса
				fetch(address, {
					method: method,
					headers: {
						'Content-Type': 'application/json'
					},
					body: method === 'GET' ? null : JSON.stringify(data) // Для GET-запросов тело не отправляется
				})
				.then(response => {
					// Проверка на ошибку CORS
					if (response.status === 0) {
						statusElement.textContent = 'Status: CORS Error - Forbidden';
						consoleElement.textContent = 'Error: CORS policy blocked the request';
					} else {
						statusElement.textContent = 'Status: ' + response.status;
						return response.json();
					}
				})
				.then(data => {
					consoleElement.textContent = 'Response: ' + JSON.stringify(data, null, 2);
				})
				.catch(error => {
					consoleElement.textContent = 'Error: ' + error;
					statusElement.textContent = 'Status: Error';
				});
			}
		</script>
		</div>`
	}

	html += `</body></html>`

	// Записываем HTML в файл
	_, err = file.WriteString(html)
	if err != nil {
		return fmt.Errorf("failed to write HTML content to file: %v", err)
	}

	return nil
}
