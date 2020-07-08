package chain

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"os"
)

/*

Blockchain storage --
 .. SQL
    .. SQL Lite
		.. PG
 .. File Block Storage
    .. Indexfile and storage
		.. Berkely DB ?

Blockhain prootocol --
 .. JSON
 .. Some sort of binary. Send structs.
*/

type Foo struct {
	code uint16
}

func RegisterMessage(message string) bool {
	var network bytes.Buffer
	var enc = gob.NewEncoder(&network)
	f := Foo{code: 12}
	err := enc.Encode(f)

	var b []byte
	n, _ := network.Read(b)
	fmt.Printf("data: %d %x\n", n, network)

	if err != nil {
		os.Exit(1)
	}

	store := NewStore("")

	store.WriteBlock(&store.ChainBlock{
		checksum:   "123",
		message:    message,
		prev_block: "",
	})

	return false
}
