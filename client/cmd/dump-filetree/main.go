package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/jiikko/filetree-meta-manager/client/internal"
)

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
	PathManager := internal.PathManager{BaseDir: directoryPath}
	err := PathManager.IsNotExist()
	if err != nil {
		fmt.Println("ディレクトリが存在しません:", PathManager.BaseDir)
		return
	}

	if *initConfig {
		// NOTE: すでにファイルが存在していたら、エラー上書きする
		err := internal.CreateConfigTemplate(PathManager.ConfigPath())
		if err != nil {
			fmt.Println("設定ファイルの作成に失敗しました:", err)
			return
		}
		fmt.Println("設定ファイルの雛形を作成しました:", PathManager.ConfigPath())
		return
	}

	config, err := internal.LoadConfig(PathManager.ConfigPath())
	if err != nil {
		fmt.Println("設定ファイルの読み込みに失敗しました:", err)
		return
	}

	fileTree, err := internal.RetrieveFileTree(PathManager)
	if err != nil {
		fmt.Println("Failed to retrieve file tree from " + PathManager.BaseDir)
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
		fmt.Printf("%s%s, %s\n", indent, node.Path, node.CreateTime.Format(time.RFC3339))
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

func postFiletree(config *internal.Config, json string) error {
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
