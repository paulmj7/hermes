package hermes

import (
	"errors"
	"os"
)

func ItemInfo(path string) (bool, int64, string, error) {
	f, err := os.Stat(path)
	if err != nil {
		return false, int64(0), "", errors.New("Error in filehandler:itemInfo, error in path")
	}

	isFile := false

	if f.Mode().IsRegular() {
		isFile = true
	}

	itemSize := int64(-1)
	date := ""
	if isFile {
		itemSize = f.Size()
		date = f.ModTime().String()[0:10]
	}

	return isFile, itemSize, date, nil
}
