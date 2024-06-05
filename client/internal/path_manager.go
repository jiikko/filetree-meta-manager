package internal

import (
	"fmt"
	"os"
	"strings"
)

type PathManager struct {
	BaseDir string
}

func (pm *PathManager) IsNotExist() error {
	if _, err := os.Stat(pm.BaseDir); os.IsNotExist(err) {
		fmt.Println("ディレクトリが存在しません:", pm.BaseDir)
		return err
	}

	return nil
}

func (pm *PathManager) BaseDirPath() string {
	return strings.TrimSuffix(pm.BaseDir, "/")
}

func (pm *PathManager) ConfigPath() string {
	return fmt.Sprintf("%s/%s", pm.BaseDir, pm.ConfigFileName())
}

func (pm *PathManager) ConfigFileName() string {
	return ".filetree_manager_config.yaml"
}
