package convert

import (
	"fmt"
	"os"
	"path/filepath"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func ConvertAll(paths []string) {
	for _, file := range paths {
		fileDir := filepath.Join("videos", filepath.Base(file[:len(file)-4]))
		err := os.MkdirAll(fileDir, os.ModePerm)
		if err != nil {
			fmt.Println(err)
			return
		}
		in := ffmpeg.Input(file)
		out := in.Output(filepath.Join(fileDir, "output.m3u8"), ffmpeg.KwArgs{
			"c:v":           "h264_nvenc", // NVIDIA NVENC hardware acceleration
			"b:v":           "5M",         // Video bitrate
			"c:a":           "aac",        // AAC audio codec
			"strict":        "-2",
			"start_number":  "0",
			"hls_time":      "10",
			"hls_list_size": "0",
			"f":             "hls", // Output format
		})
		err = out.OverWriteOutput().WithOutput(os.Stdout).Run()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("end")
		}
	}
}
