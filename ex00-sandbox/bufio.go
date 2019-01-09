package main

import (
	"bufio"
	"fmt"
	"strings"
)

func byLines() {
	s := "I felt so good like anything was possible\n I hit cruise control and rubbed my eyes\n The last three days\n the rain was un-stoppable\n It was always cold\n no sunshine\n Yeah runnin' down a dream\n That never would come to me\n Workin' on a mystery\n goin' wherever it leads\n Runnin' down a dream"

	scanner := bufio.NewScanner(strings.NewReader(s))

	// This one is default.
	//scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
		// Also we can get slice of bytes.
		//fmt.Println(scanner.Bytes())
	}
}

func byWords() {
	s := "I felt so good like anything was possible\n I hit cruise control and rubbed my eyes\n The last three days\n the rain was un-stoppable\n It was always cold\n no sunshine\n Yeah runnin' down a dream\n That never would come to me\n Workin' on a mystery\n goin' wherever it leads\n Runnin' down a dream"

	scanner := bufio.NewScanner(strings.NewReader(s))
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func byRunes() {
	s := "I felt so good like anything was possible\n I hit cruise control and rubbed my eyes\n The last three days\n the rain was un-stoppable\n It was always cold\n no sunshine\n Yeah runnin' down a dream\n That never would come to me\n Workin' on a mystery\n goin' wherever it leads\n Runnin' down a dream"

	scanner := bufio.NewScanner(strings.NewReader(s))
	scanner.Split(bufio.ScanRunes)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	byLines()
	byWords()
	byRunes()
}

