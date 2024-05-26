package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/jiikko/filetree-meta-manager/client/internal"
)

func main() {
	commandName := "filetree-meta-manager"

	if len(os.Args) < 2 {
		fmt.Println("Usage: " + commandName + " [directory path]")
		return
	}

	directoryPath := os.Args[1]
	if _, err := os.Stat(directoryPath); os.IsNotExist(err) {
		fmt.Println("ディレクトリが存在しません:", directoryPath)
		return
	}

	fileTree, err := internal.RetrieveFileTree(directoryPath)
	if err != nil {
		fmt.Println("Failed to retrieve file tree from " + directoryPath)
		return
	} else {
		printFileTree(fileTree, "")
	}

}

// printFileTree はファイルツリーを再帰的に表示します
func printFileTree(node *internal.FileInfo, indent string) {
	if node.IsDir {
		fmt.Printf("%s%s (Directory)\n", indent, filepath.Base(node.Path))
	} else {
		fmt.Printf("%s%s, %s, %s\n", indent, filepath.Base(node.Path), node.MD5Checksum, node.CreateTime.Format(time.RFC3339))
	}
	for _, child := range node.Children {
		printFileTree(child, indent+"  ")
	}
}
