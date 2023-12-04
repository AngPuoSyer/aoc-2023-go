package file

import (
	"reflect"
	"testing"
)

func TestHello(t *testing.T) {
	ch := NewFileIOChannels()

	var strArray []string

	go ScanFile(ch, "../test_input.txt")

loop:
	for {
		select {
		case line := <-ch.data:
			strArray = append(strArray, line)

		case err := <-ch.err:
			t.Fatal(err)

		case <-ch.done:
			break loop
		}
	}

	if !reflect.DeepEqual(strArray, []string{"1", "2", "3"}) {
		t.Fatal("Wrong output")
	}
}
