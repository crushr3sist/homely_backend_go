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
	Uuid       string
	Path       string
	FolderType string
	FileCount  int
}

type Directories struct {
	Directories []DirectoryInstance
}

func videoFileCollective(path string) []string {
	var returningSlice []string

	return returningSlice

}

func generateRecord(videoDirString string, givenFolder string, contentType string, fileCount int) DirectoryInstance {
	retStruct := &DirectoryInstance{}
	_id := uuid.New()
	id := _id.String()
	retStruct.Uuid = id
	retStruct.Path = givenFolder
	retStruct.FolderType = contentType
	retStruct.FileCount = fileCount
	return *retStruct

}

func ParseDirs(pathToDigest string, contentType string) *Directories {
	returningDataStructure := &Directories{}
	pathBeingTargeted := path.Base(pathToDigest)
	// formatted := FormatPath(pathBeingTargeted)
	_, files := ScanRecursive(pathBeingTargeted, []string{".txt"})

	for _, file := range files {
		structInst := generateRecord(file, pathToDigest, contentType, len(files))
		returningDataStructure.Directories = append(returningDataStructure.Directories, structInst)
	}

	return returningDataStructure

}
