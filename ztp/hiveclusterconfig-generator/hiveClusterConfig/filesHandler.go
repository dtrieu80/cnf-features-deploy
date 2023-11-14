package hiveClusterConfig

import (
	"os"
)

func ReadFile(filePath string) ([]byte, error) {
	return os.ReadFile(filePath)
}
