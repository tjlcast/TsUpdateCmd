package folderPath

import (
	"io/ioutil"
	"path/filepath"
)

type DirEntry struct {
	absDir	string
}

func newDirEntry(path string) *DirEntry {
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &DirEntry{absDir}
}

func (self *DirEntry) readFile(fileName string) ([]byte, Entry, error) {
	filename := filepath.Join(self.absDir, fileName)
	data, err := ioutil.ReadFile(filename)
	return data, self, err
}

func (self *DirEntry) addFile(dataBits []byte, fileName string) (Entry, error) {
	filePath := filepath.Join(self.absDir, fileName)
	e := ioutil.WriteFile(filePath, dataBits, 0644)
	if e != nil {
		panic(e)
	}
	return self, nil
}

func (self *DirEntry) String() string {
	return self.absDir
}