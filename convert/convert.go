package convert

import (
	"fmt"
	"os"
	"path/filepath"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func ConvertAll(paths []string) {
	fmt.Println(paths)

	for _, file := range paths {

		fileDir := filepath.Join("videos", ExtractShowName(DeFormatter(ExtractFolderPath(FormatPath(file)))), filepath.Base(file[:len(file)-4]))
		err := os.MkdirAll(fileDir, os.ModePerm)
		if err != nil {
			fmt.Println(err)
			continue
		}

		in := ffmpeg.Input(file)
		out := in.Output(filepath.Join(fileDir, "output.m3u8"), ffmpeg.KwArgs{
			"c:v":           "h264_nvenc",
			"b:v":           "5M",
			"c:a":           "aac",
			"strict":        "-2",
			"start_number":  "0",
			"hls_time":      "10",
			"hls_list_size": "0",
			"f":             "hls",
		})
		err = out.OverWriteOutput().WithOutput(os.Stdout).Run()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("end")
		}
	}
}
