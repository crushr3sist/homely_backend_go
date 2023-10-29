package directory

import (
	"os"
	"path"
	"path/filepath"
	"strings"
	"unicode"

	"github.com/google/uuid"
)

func ScanRecursive(dir_path string, ignore []string) ([]string, []string) {

	folders := []string{}
	files := []string{}

	filepath.Walk(dir_path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		_continue := false
		for _, i := range ignore {
			if strings.Index(path, i) != -1 {
				_continue = true
			}
		}

		if _continue == false {
			f_mode := info.Mode()
			if f_mode.IsDir() {
				folders = append(folders, path)
			} else if f_mode.IsRegular() {
				files = append(files, path)
			}
		}
		return nil
	})

	return folders, files
}

func FormatPath(pathString string) string {
	formattedString := ""

	for i := range pathString {
		if unicode.IsSpace(rune(pathString[i])) || string(pathString[i]) == " " {
			formattedString = strings.Replace(pathString, " ", "%20", -1)
		}
	}

	if len(formattedString) != 0 {
		return formattedString

	} else {
		return pathString
	}
}

func ExtractVideoName(pathString string) string {
	// takes a path string and extracts the video name
	// algorithm works by splitting the string by the last \\
	// right side is the file
	backSlashCount := strings.Count(pathString, "\\")
	counter := 0
	retIndex := 0

	for i := 0; i < len(pathString); i++ {
		if string(pathString[i]) == "\\" {
			counter += 1
			if counter == backSlashCount {
				retIndex = i
			}
		}
	}
	retIndex += 1
	return pathString[retIndex:]
}

func ExtractFolderPath(pathString string) string {
	// takes a path string and extracts the video name
	// algorithm works by splitting the string by the last \\
	// left side is the path

	backSlashCount := strings.Count(pathString, "\\")
	counter := 0
	retIndex := 0

	for i := 0; i < len(pathString); i++ {
		if string(pathString[i]) == "\\" {
			counter += 1
			if counter == backSlashCount {
				retIndex = i
			}
		}
	}

	retIndex += 1
	return pathString[:retIndex]
}

type DirectoryInstance struct {
	Uuid         string
	Path         string
	FolderType   string
	FileCount    int
	contentNames []string
}

type Directories struct {
	Directories []DirectoryInstance
	RawPaths    []string
}

func videoFileCollective(path []string) []string {
	var returningSlice []string

	for _, i := range path {
		returningSlice = append(returningSlice, ExtractVideoName(i))
	}

	return returningSlice
}

func generateRecord(videoDirString []string, givenFolder string, contentType string, fileCount int) DirectoryInstance {
	retStruct := &DirectoryInstance{}
	_id := uuid.New()
	id := _id.String()
	retStruct.Uuid = id
	retStruct.Path = givenFolder
	retStruct.FolderType = contentType
	retStruct.FileCount = fileCount
	retStruct.contentNames = videoFileCollective(videoDirString)
	return *retStruct
}

func ParseDirs(pathToDigest string, contentType string) *Directories {
	returningDataStructure := &Directories{}

	pathBeingTargeted := path.Base(pathToDigest)

	_, files := ScanRecursive(pathBeingTargeted, []string{".txt"})

	structInst := generateRecord(files, pathToDigest, contentType, len(files))

	returningDataStructure.Directories = append(returningDataStructure.Directories, structInst)

	returningDataStructure.RawPaths = files

	return returningDataStructure
}
