package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"time"

	"github.com/jiikko/filetree-meta-manager/client/internal"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Url        string `yaml:"url"`
	ApiKey     string `yaml:"api_key"`
	DeviceName string `yaml:"device"`
}

func main() {
	displayOnly := flag.Bool("display-only", false, "サーバに送信せずにJSONを画面に表示する")
	displayOnlyAsPlain := flag.Bool("display-only-as-plain", false, "サーバに送信せずにJSONを画面に表示する")
	initConfig := flag.Bool("init-config", false, "設定ファイルの雛形を作成する")

	flag.Parse()

	if len(flag.Args()) < 1 {
		fmt.Println("ディレクトリパスを指定してください")
		return
	}

	directoryPath := flag.Arg(0)
	if _, err := os.Stat(directoryPath); os.IsNotExist(err) {
		fmt.Println("ディレクトリが存在しません:", directoryPath)
		return
	}

	if *initConfig {
		configPath, err := createConfigTemplate(directoryPath)
		if err != nil {
			fmt.Println("設定ファイルの作成に失敗しました:", err)
			return
		}
		fmt.Println("設定ファイルの雛形を作成しました:", configPath)
		return
	}

	config, err := loadConfig(directoryPath)
	if err != nil {
		fmt.Println("設定ファイルの読み込みに失敗しました:", err)
		return
	}

	fileTree, err := internal.RetrieveFileTree(directoryPath)
	if err != nil {
		fmt.Println("Failed to retrieve file tree from " + directoryPath)
		return
	} else {
		modifyFileTree(fileTree)

		if *displayOnlyAsPlain {
			printFileTree(fileTree, "")
		} else {
			json := serializeAsJSON(fileTree)
			if *displayOnly {
				fmt.Println(json)
				return
			}

			err := postFiletree(config, json)
			if err != nil {
				fmt.Println("Failed to post file tree:", err)
			}
			return
		}
	}
}

func printFileTree(node *internal.FileInfo, indent string) {
	if node.IsDir() {
		fmt.Printf("%s%s (Directory)\n", indent, node.Path)
	} else {
		fmt.Printf("%s%s, %s, %s\n", indent, node.Path, node.MD5Checksum, node.CreateTime.Format(time.RFC3339))
	}
	for _, child := range node.Children {
		printFileTree(child, indent+"  ")
	}
}
func modifyFileTree(node *internal.FileInfo) {
	node.Path = node.PathWithoutDir()
	for _, child := range node.Children {
		modifyFileTree(child)
	}
}

func serializeAsJSON(node *internal.FileInfo) string {
	jsonData, err := json.Marshal(node)
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}
	return string(jsonData)
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

func loadConfig(baseDir string) (*Config, error) {
	configPath := filepath.Join(baseDir, "filetree_manager_config.yaml")
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

func createConfigTemplate(baseDir string) (string, error) {
	config := Config{
		Url:        "http://localhost:3000",
		ApiKey:     "your-api-key",
		DeviceName: "your-device-name",
	}

	configPath := filepath.Join(baseDir, "filetree_manager_config.yaml")
	file, err := os.Create(configPath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	encoder := yaml.NewEncoder(file)
	defer encoder.Close()

	err = encoder.Encode(config)
	if err != nil {
		return "", err
	}

	return configPath, nil
}

func postFiletree(config *Config, json string) error {
	reqBody := bytes.NewBuffer([]byte(json))
	requestPath := fmt.Sprintf("%s/api/v1/filetrees?device=%s", config.Url, url.QueryEscape(config.DeviceName))
	req, err := http.NewRequest("POST", requestPath, reqBody)
	if err != nil {
		return fmt.Errorf("リクエストの作成に失敗しました: %v", err)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", config.ApiKey)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("リクエストの送信に失敗しました: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("サーバーからエラーが返されました: %s", resp.Status) // TODO: response bodyも出力する
	}

	return nil
}
