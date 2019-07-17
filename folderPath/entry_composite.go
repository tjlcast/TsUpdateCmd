package folderPath

import (
	"errors"
	"strings"
)

type CompositeEntry []Entry

func newCompositeEntry(pathList string) CompositeEntry {
	compositeEntry := []Entry{}
	for _, path := range strings.Split(pathList, pathListSeparator) {
		entry := newEntry(path)
		compositeEntry = append(compositeEntry, entry)
	}
	return compositeEntry
}

func (self CompositeEntry) readFile(fileName string) ([]byte, Entry, error) {
	for _, entry := range self {
		data, from, err := entry.readFile(fileName)
		if err == nil {
			return data, from, nil
		}
	}
	return nil, nil, errors.New("Can't find " + fileName)
}

func (self CompositeEntry) addFile(dataBits []byte, fileName string) (Entry, error) {
	for _, entry := range self  {
		_, e := entry.addFile(dataBits, fileName)
		if e != nil {
			return nil, e
		}
	}
	return self, nil
}

func (self CompositeEntry) String() string {
	strs := make([]string, len(self))
	for i, entry := range self {
		strs[i] = entry.String()
	}
	return strings.Join(strs, pathListSeparator)
}
