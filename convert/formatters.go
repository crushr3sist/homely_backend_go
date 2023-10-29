package convert

import (
	"path/filepath"
	"strings"
	"unicode"
)

func ExtractShowName(pathString string) string {
	pathString = filepath.ToSlash(pathString)
	pathString = strings.ReplaceAll(pathString, " ", "")

	parts := strings.Split(pathString, "/")

	showKeywords := []string{"season", "series", "show", "s"}

	// Iterate over path parts from the end to the beginning
	for i := len(parts) - 1; i >= 0; i-- {
		// Check if the part contains a show keyword
		for _, keyword := range showKeywords {
			if strings.Contains(strings.ToLower(parts[i]), keyword) && i > 0 {
				// Find the next non-empty part as the show name
				for j := i - 1; j >= 0; j-- {
					if parts[j] != "" {
						return strings.TrimSpace(parts[j])
					}
				}
			}
		}
	}
	return ""
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

func DeFormatter(mediaName string) string {
	mediaName = strings.ReplaceAll(mediaName, "%20", "")
	return mediaName
}
