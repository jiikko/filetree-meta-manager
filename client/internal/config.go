package internal

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

// TODO: 無視するファイルのリストもフィールドに書きたい
type Config struct {
	Url        string `yaml:"url"`
	ApiKey     string `yaml:"api_key"`
	DeviceName string `yaml:"device"`
}

func LoadConfig(configPath string) (*Config, error) {
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config Config
	err = yaml.NewDecoder(file).Decode(&config)
	if err != nil {
		fmt.Println("設定ファイルのパースに失敗しました:", err)
		return nil, err
	}
	err = config.validateConfig()
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func (config *Config) validateConfig() error {
	if config.Url == "" {
		return fmt.Errorf("Urlが設定されていません")
	}
	if config.ApiKey == "" {
		return fmt.Errorf("ApiKeyが設定されていません")
	}
	if config.DeviceName == "" {
		return fmt.Errorf("DeviceNameが設定されていません")
	}
	return nil
}

func CreateConfigTemplate(configPath string) error {
	config := Config{
		Url:        "http://localhost:3000",
		ApiKey:     "your-api-key",
		DeviceName: "your-device-name",
	}

	file, err := os.Create(configPath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := yaml.NewEncoder(file)
	defer encoder.Close()

	err = encoder.Encode(config)
	if err != nil {
		return err
	}

	return nil
}
