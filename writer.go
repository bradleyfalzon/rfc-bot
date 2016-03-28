package main

import (
	"io"
	"log"
)

type writer struct {
	writers []io.Writer
	Wchan   chan string
}

func NewWriter() *writer {
	writer := &writer{}
	writer.Wchan = make(chan string)

	go func() {
		for {
			select {
			case str := <-writer.Wchan:
				writer.write(str)
			}
		}
	}()
	return writer
}

func (w *writer) write(str string) {
	if str == "" {
		return
	}
	log.Println("Writing msg:", str)
	for _, writer := range w.writers {
		if _, err := writer.Write([]byte(str)); err != nil {
			log.Println("Could not write str to writer:", err)
		}
	}
}

func (w *writer) AddWriter(writer io.Writer) {
	w.writers = append(w.writers, writer)
}
