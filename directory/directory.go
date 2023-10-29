package directory

import (
	"log"
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

	// Scan
	filepath.Walk(dir_path, func(path string, f os.FileInfo, err error) error {

		_continue := false

		// Loop : Ignore Files & Folders
		for _, i := range ignore {

			// If ignored path
			if strings.Index(path, i) != -1 {

				// Continue
				_continue = true
			}
		}

		if _continue == false {

			f, err = os.Stat(path)

			// If no error
			if err != nil {
				log.Fatal(err)
			}

			// File & Folder Mode
			f_mode := f.Mode()

			// Is folder
			if f_mode.IsDir() {

				// Append to Folders Array
				folders = append(folders, path)

				// Is file
			} else if f_mode.IsRegular() {

				// Append to Files Array
				files = append(files, path)
			}
		}

		return nil
	})

	return folders, files
}

func FormatPath(pathString string) string {
	// makes the path formatted correctly to query

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
	_, filename := filepath.Split(pathString)
	return filename
}

func ExtractFolderPath(pathString string) string {
	dir, _ := filepath.Split(pathString)
	return dir
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
