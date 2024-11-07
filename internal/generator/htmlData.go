package generator

import (
	"fmt"
	"github.com/Hell077/Api-Mirror/internal/parser"
	"sort"
)

func htmlDocStart() string {
	return `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>API Documentation</title>
<style>
:root {
    --primary-color: #3498db;
    --secondary-color: #2ecc71;
    --danger-color: #e74c3c;
    --warning-color: #f39c12;
    --default-bg-color: #f9f9f9;
    --default-text-color: #333;
    --input-bg-color: #fafafa;
    --input-border-color: #ddd;
    --focus-border-color: #3498db;
    --console-bg-color: #1e1e1e;
    --console-text-color: #d4d4d4;
    --header-bg-color: #ffffff;
    --font-family: Arial, sans-serif;
}

body {
    font-family: var(--font-family);
    margin: 20px;
    background-color: var(--default-bg-color);
    color: var(--default-text-color);
}

h1 {
    color: #444;
    text-align: center;
    margin-bottom: 40px;
    font-size: 36px;
    font-weight: 700;
}

.api-container {
    background-color: var(--header-bg-color);
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
    width: 100%;
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
    border: 1px solid var(--input-border-color);
    border-radius: 5px;
    box-sizing: border-box;
    background-color: var(--input-bg-color);
    font-size: 14px;
    transition: border-color 0.3s ease;
}

.api-form input[type="text"]:focus {
    border-color: var(--focus-border-color);
    outline: none;
}

.api-button {
    padding: 12px 20px;
    color: #fff;
    background-color: var(--primary-color);
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
    background-color: var(--console-bg-color);
    color: var(--console-text-color);
    padding: 15px;
    border-radius: 8px;
    width: 30%;
    font-family: Consolas, monospace;
    white-space: pre-wrap;
    font-size: 14px;
    min-height: 180px;
    box-shadow: 0 4px 10px rgba(0, 0, 0, 0.15);
    position: absolute; /* Позиционируем консоль относительно родительского блока */
    top: 10px; /* Отступ сверху */
    right: 10px; /* Отступ справа */
    margin-left: 0; /* Убираем левый отступ */
    word-wrap: break-word; /* исправление */
}

.status-output {
    color: #ffcc00;
    font-weight: bold;
    margin-top: 10px;
    font-size: 14px;
}

.status-200 {
    color: var(--secondary-color);
}

.status-400 {
    color: var(--danger-color);
}

.status-500 {
    color: var(--warning-color);
}

.status-other {
    color: #7f8c8d;
}

fieldset.api-parameters, fieldset.api-fields {
    border: 1px solid #ccc;
    padding: 20px;
    margin-bottom: 20px;
    border-radius: 8px;
    background-color: var(--default-bg-color);
}

fieldset legend {
    font-size: 18px;
    font-weight: bold;
    color: var(--default-text-color);
    padding: 0 10px;
}

ul.parameter-list, ul.field-list {
    list-style: none;
    padding: 0;
    margin: 0;
}

.parameter-item, .field-item {
    display: flex;
    flex-direction: column;
    margin-bottom: 15px;
}

.parameter-name, .field-name {
    font-size: 14px;
    font-weight: 600;
    color: #555;
    margin-bottom: 4px;
}

.parameter-type, .field-type {
    font-size: 12px;
    color: #888;
    margin-bottom: 6px;
}

.parameter-input, .field-input {
    padding: 8px 12px;
    border: 1px solid var(--input-border-color);
    border-radius: 4px;
    font-size: 14px;
    color: #333;
    background-color: #ffffff;
    transition: border-color 0.3s ease;
}

.parameter-input:focus, .field-input:focus {
    border-color: var(--focus-border-color);
    outline: none;
}

form {
    width: 65%;
    margin: 0;
}

pre {
    text-wrap: wrap;
}
</style>

</head>
<body>

<h1>API Documentation</h1>

`
}

func generateParametersHTML(parameters map[string]parser.Param, uniqueID string) string {
	html := `<fieldset class="api-parameters"><legend>Parameters (in URL):</legend><form id="params-` + uniqueID + `">`
	html += `<ul class="parameter-list">`
	for paramName, param := range parameters {
		html += fmt.Sprintf(`
			<li class="parameter-item">
				<label for="%s" class="parameter-name">%s:</label>
				<span class="parameter-type">Type: %s</span>
				<input type="text" id="%s" name="%s" placeholder="Enter %s" class="parameter-input"/>
			</li>
		`, paramName, paramName, param.Type, paramName, paramName, paramName)
	}
	html += `</ul></form></fieldset>`
	return html
}

func generateFieldsHTML(fields map[string]parser.Field, uniqueID string) string {
	html := `<fieldset class="api-fields"><legend>Fields (in body):</legend><form id="fields-` + uniqueID + `">`
	html += `<ul class="field-list">`
	for fieldName, field := range fields {
		html += fmt.Sprintf(`
			<li class="field-item">
				<label for="%s" class="field-name">%s:</label>
				<span class="field-type">Type: %s</span>
				<input type="text" id="%s" name="%s" placeholder="Enter %s" class="field-input"/>
			</li>
		`, fieldName, fieldName, field.Type, fieldName, fieldName, fieldName)
	}
	html += `</ul></form></fieldset>`
	return html
}

func GetSortStatus(responses map[int]string) string {
	var html string
	statuses := make([]int, 0, len(responses))
	for status := range responses {
		statuses = append(statuses, status)
	}

	sort.Slice(statuses, func(i, j int) bool {
		return statuses[i] < statuses[j]
	})

	for _, status := range statuses {
		html += fmt.Sprintf(`<li><span class="response-status">%d: %s</span></li>`, status, responses[status])
	}

	return html
}
func SendScript() string {
	return `
      <script>
async function sendRequest(url, method, uniqueID) {
    const consoleOutput = document.getElementById(uniqueID + "-console-output");
    consoleOutput.innerHTML = "<pre>Status: Pending...</pre>";
    updateConsoleOutput(uniqueID, 'pending'); // Ожидание запроса
    
    const formData = new FormData(document.getElementById(uniqueID + '-form'));
    let options = {
        method: method,
        headers: {}
    };

    if (method === "GET") {
        const userId = formData.get('user_id'); 
        if (userId) {
            url = url.replace("{user_id}", userId); 
        } else {
            consoleOutput.innerHTML = "<pre>Error: user_id is required for GET request.</pre>";
            updateConsoleOutput(uniqueID, 'error'); // Ошибка в GET запросе
            return;
        }
    } else {
        options.body = JSON.stringify(Object.fromEntries(formData.entries()));
        options.headers['Content-Type'] = 'application/json';
    }

    try {
        const response = await fetch(url, options);

        if (!response.ok) {
            throw new Error("Request failed with status " + response.status);
        }

        const result = await response.json();
        
        // Проверяем, что сервер вернул ответ в формате JSON
        if (result) {
            consoleOutput.innerHTML = "<pre>Response: " + JSON.stringify(result, null, 2) + "</pre>";
            updateConsoleOutput(uniqueID, 'success', result); // Выводим ответ от сервера
        } else {
            consoleOutput.innerHTML = "<pre>Response: No JSON response received.</pre>";
            updateConsoleOutput(uniqueID, 'error'); // Ошибка в формате ответа
        }

    } catch (error) {
        console.error("Error:", error);

        if (error.message.includes("Failed to fetch")) {
            consoleOutput.innerHTML = "<pre>Error: The server is unreachable or CORS policy blocked the request.</pre>";
            updateConsoleOutput(uniqueID, 'error'); // Ошибка при доступе
        } else if (error.message.includes("NetworkError") || error.message.includes("ERR_CONNECTION_REFUSED")) {
            consoleOutput.innerHTML = "<pre>Error: Failed to connect to the server. Please check the server status or your network connection.</pre>";
            updateConsoleOutput(uniqueID, 'error'); // Ошибка сети
        } else {
            consoleOutput.innerHTML = "<pre>Error: " + error.message + "</pre>";
            updateConsoleOutput(uniqueID, 'error'); // Обработка других ошибок
        }
    }
}

function updateConsoleOutput(uniqueID, status, response = null) {
    var consoleElement = document.getElementById(uniqueID + '-console-output');
    if (!consoleElement) return;

    if (status === 'pending') {
        consoleElement.style.color = 'yellow';
        consoleElement.innerText = 'Pending...';
    } else if (status === 'success') {
        consoleElement.style.color = 'green';
        if (response) {
            consoleElement.innerHTML = "<pre>" + JSON.stringify(response, null, 2) + "</pre>";
        }
    }else if (status === 'error') {
        consoleElement.style.color = 'red';

        if (response) {
            consoleElement.innerText =   response;
        }
    }
}
</script>
    `
}
