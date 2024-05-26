package internal

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type FileInfo struct {
	Path        string
	CreateTime  time.Time
	MD5Checksum string
	IsDir       bool
	Children    []*FileInfo
}

func calculateMD5(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}

func RetrieveFileTree(directoryPath string) (*FileInfo, error) {
	directoryPath = strings.TrimSuffix(directoryPath, "/")

	root := &FileInfo{
		Path:       directoryPath,
		IsDir:      true,
		Children:   []*FileInfo{},
		CreateTime: time.Now(),
	}

	err := filepath.Walk(directoryPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if path == directoryPath {
			return nil
		}

		var md5Checksum string
		if !info.IsDir() {
			md5Checksum, err = calculateMD5(path)
			if err != nil {
				return err
			}
		}

		fileInfo := &FileInfo{
			Path:        path,
			CreateTime:  info.ModTime(),
			MD5Checksum: md5Checksum,
			IsDir:       info.IsDir(),
			Children:    []*FileInfo{},
		}

		parentDir := filepath.Dir(path)
		addToParent(root, parentDir, fileInfo)

		return nil
	})

	if err != nil {
		return nil, err
	}

	return root, nil
}

func addToParent(node *FileInfo, parentDir string, fileInfo *FileInfo) bool {
	if node.Path == parentDir {
		node.Children = append(node.Children, fileInfo)
		return true
	}

	for _, child := range node.Children {
		if child.IsDir {
			if addToParent(child, parentDir, fileInfo) {
				return true
			}
		}
	}

	return false
}
