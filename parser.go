package zombie

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
	"sync"
)

//Parse format file to generates the candidates
func Parse(file string) (<-chan string, *sync.WaitGroup) {
	out := make(chan string, runtime.NumCPU()*10)
	var wg sync.WaitGroup
	//https://golang.org/pkg/sync/#WaitGroup.Add
	// This ensure that the first call to wait happens with the waitGroup greater than 0
	wg.Add(1)
	format, err := os.Open(file)
	if err != nil {
		log.Fatalf("Failed to: %v", err)
	}
	defer format.Close()
	scanner := bufio.NewScanner(format)
	formatArray := [][]string{}

	var lineNumber int
	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()
		fmt.Println(line)
		formatArray = append(formatArray, readFormat(line, lineNumber))
	}
	go buildCandidates(&wg, formatArray, out, "", true)
	return out, &wg
}

// read the format line and put it in an array of candidates
// one time symbol between parts candidates
// first symbol is a delimiter follow by a part
// repeat the delimiter before each part
// ex: aEa3  	=> [E 3]
// ex: !g!d!e!p => [g d e p]
// ex: %OUI%NON%YES%NO => [OUI NON YES NO]
func readFormat(line string, lineNumber int) []string {
	length := len(line)
	if length == 0 || length == 1 {
		log.Fatalf("Format of file is wrong on line: %d", lineNumber)
	}
	return strings.Split(line, string(line[0]))[1:]
}

//buildCandidates actually builds the candidates from the formatArray and send them to
//the out channel
func buildCandidates(wg *sync.WaitGroup, candidates [][]string, out chan<- string, part string, firstCall bool) {
	if firstCall {
		defer close(out)
		defer wg.Done()
	}
	if len(candidates) == 0 {
		wg.Add(1)
		out <- part
		//fmt.Println(part)
	} else if len(candidates) == 1 {
		for _, element := range candidates[0] {
			buildCandidates(wg, [][]string{}, out, part+element, false)
		}
	} else {
		for _, element := range candidates[0] {
			buildCandidates(wg, candidates[1:], out, part+element, false)
		}
	}
}

//Deprecated. The new one use Dispatcher func
func oldPrint(in <-chan string) {
	for elem := range in {

		fmt.Println(elem)
	}
}
