package folderPath

import (
	"os"
	"path/filepath"
	"regexp"
)

func newRegexEntry(baseDir string, path string) CompositeEntry {
	pathReg := path
	compositeEntry := []Entry{}

	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if matched, _ := regexp.MatchString(pathReg, path); matched {
			entry := newDirEntry(path)
			compositeEntry = append(compositeEntry, entry)
			return filepath.SkipDir
		}

		return nil
	}
	_ = filepath.Walk(baseDir, walkFn)
	return compositeEntry
}

func exits(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}

	return true
}

