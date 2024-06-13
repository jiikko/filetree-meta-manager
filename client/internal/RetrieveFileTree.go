package internal

import (
	"os"
	"path/filepath"
	"strings"
	"time"
)

type FileInfo struct {
	Path       string    `json:"path"`
	CreateTime time.Time `json:"created_at"`
	isDir      bool
	Children   []*FileInfo `json:"children"`
}

func (f *FileInfo) PathWithoutDir() string {
	return filepath.Base(f.Path)
}

func (f *FileInfo) IsDir() bool {
	return f.isDir
}

var ignorePaths = []string{
	".Spotlight-V100",
	".DS_Store",
	".fseventsd",
	".unwanted",
}

func RetrieveFileTree(pm PathManager) (*FileInfo, error) {
	directoryPath := pm.BaseDirPath()
	rootInfo, err := os.Stat(directoryPath)
	if err != nil {
		return nil, err
	}

	root := &FileInfo{
		Path:       directoryPath,
		Children:   []*FileInfo{},
		CreateTime: rootInfo.ModTime().Truncate(time.Second),
	}

	err = filepath.Walk(directoryPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}

		if path == directoryPath {
			return nil
		}

		for _, ignoredPath := range ignorePaths {
			if strings.Contains(path, ignoredPath) {
				return nil
			}
		}

		// NOTE: 設定ファイルは無視する
		if strings.HasSuffix(path, pm.ConfigPath()) {
			return nil
		}

		fileInfo := &FileInfo{
			Path:       path,
			CreateTime: info.ModTime().Truncate(time.Second),
			isDir:      info.IsDir(),
			Children:   []*FileInfo{},
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
		if child.IsDir() {
			if addToParent(child, parentDir, fileInfo) {
				return true
			}
		}
	}

	return false
}
