package chain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/timshannon/badgerhold"
)

type ChainStore struct {
	filename   string
	store      badgerhold.Store
	fd         *os.File
	last_block string
}

func NewStore(filename string) *ChainStore {

	store, err := badgerhold.Open(filename, 0666, nil)
	if err != nil {
		fmt.Println("Cannot access database file")
		pantic(err)
	}

	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}
	return &ChainStore{filename: filename, fd: file}
}

/*
	checksum compute
	prev block checksum
*/

type ChainBlock struct {
	checksum   string
	data       []byte
	prev_block string // hash / checksum
}

/*
	Write block returns position of the block in the file.
*/

func (s *ChainStore) WriteBlock(b ChainBlock) (int64, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(b)
	slices := sha256.Sum256(b.data)
	checksum := string(slices[:])
	cb := ChainBlock(b)
	err = store.Insert("key", &cb{
		checksum:   checksum,
		prev_block: s.last_block, //TODO prev block populate
	})

	if err != nil {
		log.Fatal(err)
		return 0, err
	} else {
		offset, err := s.fd.Seek(0, io.SeekCurrent)
		if err != nil {
			panic(err)
		} else {
			s.fd.Write(buf.Bytes())
			return offset, nil
		}
	}
}
