package main

import "github.com/tyatyushkin/gadget"

type Player interface {
	Play(string)
	Stop()
}

func TryOut(player Player) {
	player.Play("Test track")
	player.Stop()
	recorder, ok := player.(gadget.TapeRecorder)
	if ok {
		recorder.Record()
	}
}

func main() {
	TryOut(gadget.TapePlayer{})
	TryOut(gadget.TapeRecorder{})
}
