package main

import (
	"fmt"
	"strconv"

	"github.com/yeqown/go-wal"
)

func main() {
	//example()
	example2()
}

func example() {
	config := wal.DefaultConfig()
	w, err := wal.NewWAL(config, wal.WithRoot("./testdata/wal"))
	if err != nil {
		panic(err)
	}

	err = w.TruncateBefore(140)
	if err != nil {
		panic(err)
	}

	// now you can use wal to write and read data
	var (
		newest int64
	)
	for i := 0; i < 100; i++ {
		newest, err = w.Write([]byte("hello world" + strconv.Itoa(i)))
		if err != nil {
			panic(err)
		}
	}

	// read data from newest
	data, err := w.Read(newest)
	if err != nil {
		panic(err)
	}
	fmt.Printf("read data: %s, offset: %d\n", string(data), newest)

	// read
	data2, offset, err := w.ReadLatest()
	if err != nil {
		panic(err)
	}
	fmt.Printf("read data: %s, offset: %d\n", string(data2), offset)

	if err := w.Close(); err != nil {
		panic(err)
	}
}
