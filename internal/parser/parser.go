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
	Parameters map[string]Param `yaml:"parameters"`
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

	// Отладочная печать содержимого файла YAML
	fmt.Println("YAML file content:\n", string(data))

	var config APIConfig
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("error parsing YAML: %v", err)
	}

	// Печать разобранного содержимого YAML для отладки
	fmt.Printf("Parsed YAML content: %+v\n", config)

	// Валидация конфигурации
	if err := validateConfig(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

func validateConfig(config *APIConfig) error {
	// Проверка наличия обязательных данных
	if config.APIMirror.APIList == nil || len(config.APIMirror.APIList) == 0 {
		return errors.New("mandatory field API_MIRROR is missing or empty")
	}

	// Валидация каждого API
	for name, api := range config.APIMirror.APIList {
		// Проверка обязательных полей API
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

		// Валидация параметров API
		if len(api.Parameters) > 0 {
			for paramName, param := range api.Parameters {
				if param.Type == "" {
					return fmt.Errorf("API '%s' parameter '%s' is missing 'type'", name, paramName)
				}
				// Проверка на наличие плейсхолдера в параметре, если он предусмотрен
				if param.Placeholder == "" {
					return fmt.Errorf("API '%s' parameter '%s' is missing 'placeholder'", name, paramName)
				}
			}
		}

		// Валидация полей API
		if len(api.Fields) > 0 {
			for fieldName, field := range api.Fields {
				if field.Type == "" {
					return fmt.Errorf("API '%s' field '%s' is missing 'type'", name, fieldName)
				}
			}
		}
	}
	return nil
}
