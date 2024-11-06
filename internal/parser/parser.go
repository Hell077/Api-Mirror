package parser

import (
	"errors"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type APIConfig struct {
	APIMirror APIMirror `yaml:"API_MIRROR"`
}

type APIMirror struct {
	APIList map[string]API `yaml:"API_LIST"`
}

type API struct {
	Address   string         `yaml:"address"`
	Method    string         `yaml:"method"`
	Responses map[int]string `yaml:"responses"` // Убираем `,inline`
	Title     string         `yaml:"title"`
}

func ParseYAML(filepath string) (*APIConfig, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	fmt.Println("YAML file content:\n", string(data))

	var config APIConfig
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("error parsing YAML: %v", err)
	}

	// Дополнительная проверка структуры
	fmt.Printf("Parsed YAML content: %+v\n", config)

	if err := validateConfig(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

func validateConfig(config *APIConfig) error {
	if config.APIMirror.APIList == nil || len(config.APIMirror.APIList) == 0 {
		return errors.New("mandatory field API_MIRROR is missing or empty")
	}

	for name, api := range config.APIMirror.APIList {
		if api.Address == "" {
			return fmt.Errorf("API '%s' is missing the mandatory field 'address'", name)
		}
		if api.Method == "" {
			return fmt.Errorf("API '%s' is missing the mandatory field 'method'", name)
		}
		if len(api.Responses) == 0 {
			return fmt.Errorf("API '%s' has no response statuses", name)
		}
		if api.Title == "" {
			return fmt.Errorf("API '%s' is missing the mandatory field 'title'", name)
		}
	}
	return nil
}
