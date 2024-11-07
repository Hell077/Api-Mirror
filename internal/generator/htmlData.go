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
		function sendRequest(url, method, formID) {
			const form = document.getElementById(formID + '-form');
			const inputs = form.getElementsByTagName('input');

			// Замена плейсхолдеров параметров в URL
			let requestUrl = url;
			for (let input of inputs) {
				const placeholder = '{' + input.name + '}';
				requestUrl = requestUrl.replace(placeholder, input.value || input.placeholder);
			}

			// Настройка и отправка запроса
			fetch(requestUrl, {
				method: method,
				headers: {
					'Content-Type': 'application/json'
				},
				body: method === 'GET' ? null : JSON.stringify(Object.fromEntries(new FormData(form)))
			})
			.then(response => response.text())
			.then(data => {
				document.getElementById(formID + '-console-output').innerText = 'Response: ' + data;
			})
			.catch(error => {
				document.getElementById(formID + '-console-output').innerText = 'Error: ' + error;
			});
		}
	</script>`
}

func generateParametersHTML(parameters map[string]parser.Param, uniqueID string) string {
	html := `<div class="parameters">`
	for name, param := range parameters {
		html += fmt.Sprintf(
			`<div class="parameter"><label for="%s-%s">%s (%s):</label><input type="text" id="%s-%s" name="%s" placeholder="%s"></div>`,
			uniqueID, name, name, param.Type, uniqueID, name, name, param.Placeholder,
		)
	}
	html += `</div>`
	return html
}

func generateFieldsHTML(fields map[string]parser.Field, uniqueID string) string {
	html := `<div class="fields">`
	for name, field := range fields {
		html += fmt.Sprintf(
			`<div class="field"><label for="%s-%s">%s (%s):</label><input type="text" id="%s-%s" name="%s" placeholder="%s"></div>`,
			uniqueID, name, name, field.Type, uniqueID, name, name, field.Mask,
		)
	}
	html += `</div>`
	return html
}
