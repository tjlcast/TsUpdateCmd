package folderPath

import (
	"os"
	"strings"
)

const pathListSeparator = string(os.PathListSeparator)
const pathSeparator = string(os.PathSeparator)

type Entry interface {
	readFile(fileName string) ([]byte, Entry, error)
	addFile(dataBits []byte, fileName string) (Entry, error)
	String() string
}

func newEntry(path string) Entry {
	// multiple paths splited by `pathSeparator`
	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	}

	//if strings.Contains(path, "*") {
	//	index := strings.Index(path, "*")
	//	indexLast := strings.LastIndex(path[:index], pathSeparator)
	//	basePath := path[:indexLast]
	//	return nil
	//	//newWildcardEntry(basePath, path)
	//}

	return newDirEntry(path)
}
