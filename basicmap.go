package hashing

import (
	"bytes"
	"container/list"
	"encoding/gob"
	"errors"
	"strconv"
)

type basicHashMap struct {
	table []*list.List
}

func NewBasicHashMap() HashMap {
	initSize := 64
	table := make([]*list.List, 64)
	for i := 0; i < initSize; i++ {
		table[i] = list.New()
	}
	return &basicHashMap{table: table}
}

func hash(bytes []byte) int {
	h := 0
	for b := range bytes {
		h = 31*h + b
	}
	return h
}

func (b *basicHashMap) Put(key string, value interface{}) error {
	var buffer bytes.Buffer
	err := gob.NewEncoder(&buffer).Encode(value)
	if err != nil {
		return errors.New("bad value")
	}
	h := hash([]byte(key))
	index := h % len(b.table)
	encodedValue := buffer.Bytes()
	b.table[index].PushBack(pair{key, encodedValue})
	return nil
}

func search(l *list.List, key string) (encodedValue []byte) {
	element := l.Front()
	for element != nil {
		p := element.Value.(pair)
		if p.key == key {
			return p.value
		}
	}
	return nil
}

func (b *basicHashMap) Get(key string, value interface{}) error {
	h := hash([]byte(key))
	index := h % len(b.table)
	foundEncodedValue := search(b.table[index], key)
	if foundEncodedValue == nil {
		return errors.New("key not present")
	}
	buffer := bytes.NewReader(foundEncodedValue)
	err := gob.NewDecoder(buffer).Decode(value)
	if err != nil {
		return errors.New("bad encoded value")
	}
	return nil
}

func (b *basicHashMap) ToString() string {
	s := ""
	for i, l := range b.table {
		s += strconv.Itoa(i)
		e := l.Front()
		for e != nil {
			s += e.Value.(pair).key
		}
	}
}
