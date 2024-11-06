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
	SERVER  string         `yaml:"SERVER"`
	PORT    string         `yaml:"PORT"`
	APIList map[string]API `yaml:"API_LIST"`
}

type API struct {
	Address    string           `yaml:"address"`
	Method     string           `yaml:"method"`
	Fields     map[string]Field `yaml:"fields"`
	Responses  map[int]string   `yaml:"responses"`
	Title      string           `yaml:"title"`
	Parameters map[string]Param `yaml:"parameters"` // Новое поле для параметров
}

type Field struct {
	Type string `yaml:"type"`
	Mask string `yaml:"mask"`
}

type Param struct {
	Type        string `yaml:"type"`        // Тип параметра, например, "int", "string"
	Placeholder string `yaml:"placeholder"` // Значение по умолчанию или описание для заполнителя, как {user_id}
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
		if len(api.Parameters) > 0 {
			// Валидируем параметры, если они есть
			for paramName, param := range api.Parameters {
				if param.Type == "" {
					return fmt.Errorf("API '%s' parameter '%s' is missing 'type'", name, paramName)
				}
			}
		}
	}
	return nil
}
