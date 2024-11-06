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

	html := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>API Documentation</title>
    <style>
        body { font-family: Arial, sans-serif; }
        .api-container { margin-bottom: 20px; padding: 15px; border: 1px solid #ddd; }
        .api-title { font-size: 1.5em; color: #333; }
        .api-address, .api-method { font-weight: bold; }
        .response-status { margin-left: 20px; }
    </style>
</head>
<body>
<h1>API Documentation</h1>
`
	for _, api := range config.APIMirror.APIList {
		html += `<div class="api-container">`
		html += fmt.Sprintf(`<div class="api-title">%s</div>`, api.Title)
		html += fmt.Sprintf(`<div class="api-address">Address: %s</div>`, api.Address)
		html += fmt.Sprintf(`<div class="api-method">Method: %s</div>`, api.Method)
		html += `<div class="response-statuses"><strong>Response Statuses:</strong><ul>`

		for status, response := range api.Responses {
			html += fmt.Sprintf(`<li><span class="response-status">%d: %s</span></li>`, status, response)
		}

		html += `</ul></div></div>`
	}

	html += `
</body>
</html>`

	if _, err := file.WriteString(html); err != nil {
		return fmt.Errorf("failed to write to HTML file: %v", err)
	}

	return nil
}
