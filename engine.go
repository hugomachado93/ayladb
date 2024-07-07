package main

import (
	"io"
	"os"
	"strconv"
)

type OffsetMap struct {
	offset uint
	free   bool
}

type Engine struct {
	storage   *os.File
	hashTable map[string]OffsetMap
	// freeOffsetHashTable map[string]
	lastOffset        uint
	lastHashMapOffset uint
}

func NewEngine() *Engine {
	return &Engine{hashTable: make(map[string]OffsetMap), lastOffset: 0, lastHashMapOffset: 0}
}

func (sf *Engine) recordToFile(key string, data string) error {

	fileHash, err := os.OpenFile("hashmap.db", os.O_CREATE|os.O_RDWR, 0666)

	if err != nil {
		return err
	}

	file, err := os.OpenFile("db.db", os.O_CREATE|os.O_RDWR, 0666)

	if err != nil {
		return err
	}

	offset, err := file.Seek(0, io.SeekEnd)

	if err != nil {
		return err
	}

	file.WriteString(data + breakline)

	sf.hashTable[key] = OffsetMap{offset: uint(offset), free: false}

	offset, err = fileHash.Seek(0, io.SeekEnd)

	if err != nil {
		return err
	}

	fileHash.WriteString(key + separator + strconv.Itoa(int(offset)) + breakline)

	return nil
}

func (sf *Engine) loadFileToHashMap() error {
	fileHash, err := os.OpenFile("hashmap.db", os.O_RDONLY, 0666)

	if err != nil {
		return err
	}

	fileHash.Seek(0, io.SeekStart)

	buffer := make([]byte, 1)

	var key []byte
	var data []byte

	var readKey bool = true
	// var readData bool = false

	for {
		n, err := fileHash.Read(buffer)

		if err != nil {
			return err
		}

		if n == 0 {
			break
		}

		if buffer[0] == '|' {
			readKey = false
		}

		if buffer[0] == '\n' {
			readKey = true
		}

		if readKey {
			key = append(key, buffer[0])
		} else {
			data = append(data, buffer[0])
		}

		// sf.hashTable[string(key)] = string(data)
	}

	return nil
}
