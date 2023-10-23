package directory

import (
	"path"
	"strings"
	"unicode"
)

type DirectoryInstance struct {
	Uuid            string
	Path            string
	FolderType      string
	FileCount       int8
	StreamableFiles []int
}

type Directories struct {
	Directories []DirectoryInstance
}

func FormatPath(pathString string) string {
	formattedString := ""

	for i := range pathString {
		if unicode.IsSpace(rune(pathString[i])) || string(pathString[i]) == " " {
			formattedString = strings.Replace(pathString, " ", "%20", -1)
		}
		if string(pathString[i]) == "/" {
			formattedString = strings.Replace(pathString, "/", "\\", -1)
		}
	}

	if len(formattedString) != 0 {
		return formattedString

	} else {
		return pathString
	}
}

func parseDirs(pathToDigest string) {
	returningDataStructure := &Directories{}
	pathBeingTargetted := path.Base(pathToDigest)

}
