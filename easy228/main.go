package easy228

import (
	"bufio"
	// "fmt"
	"os"
	"strings"
)

func checkLineInFile(path string) {
	c := make(chan string)
	go readLines("test_input", c)
	var ok bool
	var w string
	for {
		select {
		case w, ok = <-c:
			WriteOrder(w)
		}
		if !ok {
			break
		}
	}
}

func readLines(path string, c chan<- string) {
	inFile, _ := os.Open(path)
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		c <- scanner.Text()
	}
	close(c)
}

func WriteOrder(w string) {
	w = strings.ToLower(w)
	if inOrder(w) {
		// fmt.Printf("%s IN ORDER\n", w)
	} else if inReverseOrder(w) {
		// fmt.Printf("%s IN REVERSE ORDER\n", w)
	} else {
		// fmt.Printf("%s NOT IN ORDER\n", w)
	}

}

func inOrder(w string) bool {
	val := 0

	for _, l := range w {
		i := int(l)
		if val > i {
			return false
		}
		val = i
	}
	return true
}

func inReverseOrder(w string) bool {
	val := 256

	for _, l := range w {
		i := int(l)
		if val < i {
			return false
		}
		val = i
	}
	return true
}
