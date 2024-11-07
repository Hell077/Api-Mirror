package generator

import (
	"fmt"
	"os"

	"github.com/Hell077/Api-Mirror/internal/parser"
)

// Generator генерирует HTML файл на основе конфигурации API
func Generator(config *parser.APIConfig, outputFileName string) error {
	file, err := os.Create(outputFileName)
	if err != nil {
		return fmt.Errorf("failed to create HTML file: %v", err)
	}
	defer file.Close()

	baseURL := fmt.Sprintf("http://%s:%s", config.APIMirror.SERVER, config.APIMirror.PORT)

	html := htmlDocStart()
	html += fmt.Sprintf(`<div>Url: %s </div>`, baseURL)

	for index, api := range config.APIMirror.APIList {
		uniqueID := fmt.Sprintf("%s-%d", api.Title, index)
		html += `<div class="api-container">`

		html += fmt.Sprintf(`<form id="%s-form" method="%s">`, uniqueID, api.Method)
		html += `<div class="api-details">`
		html += fmt.Sprintf(`<div class="api-title">%s</div>`, api.Title)
		html += fmt.Sprintf(`<div class="api-address">Address: %s</div>`, api.Address)
		html += fmt.Sprintf(`<div class="api-method">Method: %s</div>`, api.Method)

		// Рендеринг параметров URL как input-поля
		if len(api.Parameters) > 0 {
			html += generateParametersHTML(api.Parameters, uniqueID)
		}

		// Рендеринг полей тела запроса
		if len(api.Fields) > 0 {
			html += generateFieldsHTML(api.Fields, uniqueID)
		}

		// JavaScript вызов sendRequest с URL и параметрами формы
		html += fmt.Sprintf(
			`<button class="api-button" type="button" onclick="sendRequest('%s', '%s', '%s')">Send Request</button>`,
			baseURL+api.Address, api.Method, uniqueID,
		)
		html += `<ul class="response-status-list">` + GetSortStatus(api.Responses) + `</ul>`
		html += fmt.Sprintf(`<div id="%s-console-output" class="console-output"></div>`, uniqueID)
		html += `</div>`
		html += `</form>`
		html += `</div>`
	}

	html += SendScript()
	html += `</body></html>`

	_, err = file.WriteString(html)
	if err != nil {
		return fmt.Errorf("failed to write HTML content to file: %v", err)
	}

	return nil
}
