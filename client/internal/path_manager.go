package internal

import (
	"fmt"
	"os"
	"strings"
)

type PathManager struct {
	BaseDir string // NOTE: 呼び出し厳禁. 末尾に/があったりなかったりするので、BaseDirPath()を使うこと
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
	return fmt.Sprintf("%s/%s", pm.BaseDirPath(), pm.ConfigFileName())
}

func (pm *PathManager) ConfigFileName() string {
	return ".filetree_manager_config.yaml"
}
