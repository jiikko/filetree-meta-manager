package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/jiikko/filetree-meta-manager/client/internal"
)

func main() {
	outputJSON := flag.Bool("json", false, "JSON形式で出力する")
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
	if node.IsDir {
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

	jsonData, err := json.MarshalIndent(node, "", "  ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(jsonData))
}
