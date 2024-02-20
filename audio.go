package simpleui

import (
	"github.com/gordonklaus/portaudio"
	"log"
	"time"
)

type Audio struct {
	stream         *portaudio.Stream
	sampleRate     float64
	outputChannels int
	channel        chan float32
}

func NewAudio() *Audio {
	a := Audio{}
	return &a
}

func (a *Audio) Start() error {
	host, err := portaudio.DefaultHostApi()
	if err != nil {
		return err
	}
	// HighLatencyParameters are mono in, stereo out (if supported),
	// high latency, the smaller of the default sample rates of the two devices, and FramesPerBufferUnspecified. One of the devices may be nil.
	parameters := portaudio.HighLatencyParameters(nil, host.DefaultOutputDevice)

	delay := time.Second / 5
	bufLen := int(parameters.SampleRate * delay.Seconds())

	a.channel = make(chan float32, bufLen)

	stream, err := portaudio.OpenStream(parameters, a.Callback)
	if err != nil {
		return err
	}
	if err := stream.Start(); err != nil {
		return err
	}
	a.stream = stream
	a.sampleRate = parameters.SampleRate
	a.outputChannels = parameters.Output.Channels

	log.Println("Audio started, sampleRate:", a.sampleRate, "outputChannels:", a.outputChannels)
	return nil
}

func (a *Audio) Stop() error {
	return a.stream.Close()
}

// `Callback` invoked in a separate fixed goroutine
func (a *Audio) Callback(out []float32) {
	var output float32
	for i := range out {
		if i%a.outputChannels == 0 {
			select {
			case sample := <-a.channel:
				output = sample
			default:
				output = 0
			}
		}
		out[i] = output
	}
}

func (a *Audio) GetSampleRate() float64 {
	return a.sampleRate
}
func (a *Audio) GetAudioChannel() chan float32 {
	return a.channel
}
