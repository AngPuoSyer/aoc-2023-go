package file

import (
	"bufio"
	"os"
)

type FileIOChan struct {
	data chan string
	done chan bool
	err  chan error
}

func NewFileIOChannels() FileIOChan {
	return FileIOChan{
		data: make(chan string),
		done: make(chan bool),
		err:  make(chan error),
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ScanFile(ioChan FileIOChan, path string) {
	file, err := os.Open(path)
	check(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		ioChan.data <- scanner.Text()
	}

	if scanner.Err() != nil {
		ioChan.err <- scanner.Err()
	}

	ioChan.done <- true
}
