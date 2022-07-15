package Judge

import (
	"os"
)

// judge this path whether exist
func PathJudgeIsExists(path string) bool {
	// get file info
	_, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

// judge this path is file or folder
// true folder
// false file
func PathJudgeIsFolder(path string) bool {
	// get file info
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}
	// return file info
	return fileInfo.IsDir()
}
