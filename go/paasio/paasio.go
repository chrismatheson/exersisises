package paasio

import (
	"io"
	"sync"
)

type IODirection string
const (
	IORead IODirection = "read"
	IOWrite IODirection = "write"
)

type MyRWCounter struct{
	mut sync.Mutex
	readstream io.Reader
	wrtiestream io.Writer
	byteCount map[IODirection]int64
	callsCount map[IODirection]int
}

func (super *MyRWCounter) Read(p []byte) (n int, err error) {
	n, err = super.readstream.Read(p)
	super.mut.Lock()
	super.byteCount[IORead]+= int64(n)
	super.callsCount[IORead]++
	super.mut.Unlock()
	return n, err
}

func (super *MyRWCounter) ReadCount() (n int64, nops int) {
	super.mut.Lock()
	n = super.byteCount[IORead]
	nops = super.callsCount[IORead]
	super.mut.Unlock()
	return n, nops
}

func (super *MyRWCounter) Write(p []byte) (n int, err error) {
	n, err = super.wrtiestream.Write(p)
	super.mut.Lock()
	super.byteCount[IOWrite]+= int64(n)
	super.callsCount[IOWrite]++
	super.mut.Unlock()
	return n, err
}

func (super *MyRWCounter) WriteCount() (n int64, nops int) {
	super.mut.Lock()
	n = super.byteCount[IOWrite]
	nops = super.callsCount[IOWrite]
	super.mut.Unlock()
	return n, nops
}

func NewReadCounter(stream io.Reader) ReadCounter {
	counter := MyRWCounter{
		byteCount: map[IODirection]int64{},
		callsCount: map[IODirection]int{},
		readstream: stream,
	}
	return &counter
}

func NewWriteCounter(stream io.Writer) WriteCounter {
	return &MyRWCounter{
		byteCount: map[IODirection]int64{},
		callsCount: map[IODirection]int{},
		wrtiestream: stream,
	}
}

func NewReadWriteCounter(stream io.ReadWriter) ReadWriteCounter {
	return &MyRWCounter{
		byteCount: map[IODirection]int64{},
		callsCount: map[IODirection]int{},
		wrtiestream: stream,
		readstream: stream,
	}
}