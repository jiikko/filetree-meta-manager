package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/jiikko/filetree-meta-manager/client/internal"
	"gopkg.in/yaml.v2"
)

type Config struct {
	url    string `yaml:"url"`
	apiKey string `yaml:"api_key"`
}

func main() {
	outputJSON := flag.Bool("json", true, "JSON形式で出力する")

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
	config, err := loadConfig(directoryPath)
	if err != nil {
		fmt.Println("設定ファイルの読み込みに失敗しました:", err)
		return
	}
	fmt.Println("URL:", config.url)

	fileTree, err := internal.RetrieveFileTree(directoryPath)
	if err != nil {
		fmt.Println("Failed to retrieve file tree from " + directoryPath)
		return
	} else {
		if *outputJSON {
			modifyFileTree(fileTree)
			printFileTreeJSON(fileTree)
		} else {
			modifyFileTree(fileTree)
			printFileTree(fileTree, "")
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

func printFileTreeJSON(node *internal.FileInfo) {
	jsonData, err := json.Marshal(node)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(jsonData))
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

	return &config, nil
}
