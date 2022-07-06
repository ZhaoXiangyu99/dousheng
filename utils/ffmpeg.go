package utils

import (
	"bytes"
	"os/exec"
)

// GetCoverForVideo 使用ffmpeg工具为视频截取封面
func GetCoverForVideo(videoPath, coverPath string) error {
	in := bytes.Buffer{}
	cmd := exec.Command("zsh")
	cmd.Stdin = &in

	//str := "ffmpeg -i /Users/zhaoxiangyu/Code/resources/video/123.mp4 -ss 1 -t 0.1 -frames:v 1 /Users/zhaoxiangyu/Code/resources/cover/678.jpg"
	str := "ffmpeg -i " + videoPath + " -ss 1 -t 0.1 -frames:v 1 " + coverPath
	in.WriteString(str)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
