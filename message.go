package main

type Message struct {
	data []byte
}

type Header struct {
	Size       int32
	APIKey     int16
	APIVersion int16
}
