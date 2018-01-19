package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"os"
)

// var fname = "TestDataSet.dlis"
var fname = "n802b_SHELL14.dls"

// SUL = storage unit label 80 bytes
type SUL struct {
	SeqNum      [4]byte  // sequence number
	DLISVersion [5]byte  // "V1.00" - most likely
	Struct      [6]byte  // storage unit structure - most likely a "RECORD"
	MaxRecLen   [5]byte  // maximum record length
	SetID       [60]byte // storage set identifier
}

func (s SUL) String() string {
	return fmt.Sprintf("Sequence Number: %s; DLISVersion: %s; Structure: %s; MaxiRecLen: %s, SetID: %s",
		string(s.SeqNum[:]), string(s.DLISVersion[:]),
		string(s.Struct[:]), string(s.MaxRecLen[:]),
		string(s.SetID[:]))
}

type VisibleEnvelop struct {
	SUL SUL
}

func readManual() (env VisibleEnvelop) {
	f, err := os.Open(fname)
	if err != nil {
		log.Printf("error opening file %s : %v", fname, err)
		return
	}
	defer dclose(f)

	_, err = f.Read(env.SUL.SeqNum[:])
	_, err = f.Read(env.SUL.DLISVersion[:])
	_, err = f.Read(env.SUL.Struct[:])
	_, err = f.Read(env.SUL.MaxRecLen[:])
	_, err = f.Read(env.SUL.SetID[:])

	return
}
func readReflect() (env VisibleEnvelop) {
	f, err := os.Open(fname)
	if err != nil {
		log.Printf("error opening file %s : %v", fname, err)
		return
	}
	defer dclose(f)

	err = binary.Read(f, binary.BigEndian, &env)

	return
}

func main() {

}

func dclose(c io.Closer) {
	if err := c.Close(); err != nil {
		log.Printf("error with deferred closing: %v", err)
	}
}
