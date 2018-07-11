package dlis

import (
	"io"
	"log"
)

func dclose(c io.Closer) {
	if err := c.Close(); err != nil {
		log.Printf("error with deferred closing: %v", err)
	}
}
