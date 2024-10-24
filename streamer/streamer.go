package streamer

import (
	"fmt"
	"github.com/tsawler/toolbox"
	"path"
	"path/filepath"
	"strings"
)

type ProcessingMessage struct {
	ID         int
	Successful bool
	Message    string
	OutputFile string
}

type VideoProcessingJob struct {
	Video Video
}

type Processor struct {
	Engine Encoder
}

type Video struct {
	ID           int
	InputFile    string
	OutputDir    string
	EncodingType string
	NotifyChan   chan ProcessingMessage
	Encoder      Processor
	Options      *VideoOptions
}

type VideoOptions struct {
	RenameOutput    bool
	SegmentDuration int
	MaxRate1080p    string
	MaxRate720p     string
	MaxRate480p     string
}

func (vd *VideoDispatcher) NewVideo(id int, input, output, encType string, notifyChan chan ProcessingMessage, ops *VideoOptions) Video {

	if ops == nil {
		ops = &VideoOptions{}
	}

	return Video{
		ID: id, InputFile: input, OutputDir: output, EncodingType: encType, NotifyChan: notifyChan, Encoder: vd.Processor, Options: ops,
	}
}

func (v *Video) encode() {
	var fileName string

	switch v.EncodingType {
	case "mp4":
		name, err := v.encodeToMP4()
		if err != nil {
			v.sendToNotifyChan(false, "", fmt.Sprintf("encode failed for %d, %s", v.ID, err.Error()))
			return
		}
		fileName = fmt.Sprintf("%s.mp4", name)
	case "hls":
		name, err := v.encodeToHLS()
		if err != nil {
			v.sendToNotifyChan(false, "", fmt.Sprintf("encode failed for %d, %s", v.ID, err.Error()))
			return
		}
		fileName = fmt.Sprintf("%s.m3u8", name)
	default:
		v.sendToNotifyChan(false, "", fmt.Sprintf("error processing for %d: invalid encoding type", v.ID))
		return
	}

	v.sendToNotifyChan(true, fileName, fmt.Sprintf("Video %d was encoded and saved as %s", v.ID, fileName))
}

func (v *Video) encodeToMP4() (string, error) {

	var t toolbox.Tools
	baseFileName := ""

	if !v.Options.RenameOutput {
		b := path.Base(v.InputFile)
		baseFileName = strings.TrimSuffix(b, filepath.Ext(b))
	} else {
		baseFileName = t.RandomString(10)
	}

	err := v.Encoder.Engine.EncodeToMP4(v, baseFileName)

	if err != nil {
		return "", err
	}

	return baseFileName, nil
}

func (v *Video) encodeToHLS() (string, error) {

	var t toolbox.Tools
	baseFileName := ""

	if !v.Options.RenameOutput {
		b := path.Base(v.InputFile)
		baseFileName = strings.TrimSuffix(b, filepath.Ext(b))
	} else {
		baseFileName = t.RandomString(10)
	}

	err := v.Encoder.Engine.EncodeToHLS(v, baseFileName)

	if err != nil {
		return "", err
	}

	return baseFileName, nil
}

func (v *Video) sendToNotifyChan(successful bool, fileName, message string) {
	v.NotifyChan <- ProcessingMessage{
		ID:         v.ID,
		Successful: successful,
		Message:    message,
		OutputFile: fileName,
	}
}

func New(jobQueue chan VideoProcessingJob, maxWorkers int) *VideoDispatcher {

	workerPool := make(chan chan VideoProcessingJob, maxWorkers)

	var e VideoEncoder

	processor := Processor{
		Engine: &e,
	}

	return &VideoDispatcher{
		WorkerPool: workerPool,
		maxWorkers: maxWorkers,
		jobQueue:   jobQueue,
		Processor:  processor,
	}
}
