package zombie

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
	"sync"

	"github.com/btcsuite/btcutil/base58"
)

type Cracker func(string, *sync.WaitGroup)

func Dispatcher(fn Cracker, workers int, in <-chan string, wg *sync.WaitGroup) {
	var i int
	for i = 0; i < workers; i++ {
		go func() {
			for {
				line, ok := <-in
				if !ok {
					break
				}
				fn(line, wg)
			}
		}()
	}
	wg.Wait()
}

func wifChecker(line string, wg *sync.WaitGroup) {
	defer wg.Done()
	result, _, err := base58.CheckDecode(line)
	if err == nil {
		fmt.Println("We found a valid Wif ", hex.EncodeToString(result))
		fmt.Println("Wif ", line)
		file, errFile := os.Create(hex.EncodeToString(result))
		if errFile != nil {
			panic(errFile.Error())
		}
		defer file.Close()

		// Create a buffered writer from the file
		bufferedWriter := bufio.NewWriter(file)
		_, errorBuf := bufferedWriter.WriteString("line: " + line + "\nresult " + hex.EncodeToString(result) + "\n")
		if errorBuf != nil {
			panic(errorBuf.Error())
		}
		bufferedWriter.Flush()
	}
}

func CrackWif(workers int, in <-chan string, wg *sync.WaitGroup) {
	fmt.Println("Starting CrackWif")
	Dispatcher(wifChecker, workers, in, wg)
	wg.Wait()
	fmt.Println("Done Crackwif")
}

func Print(in <-chan string, wg *sync.WaitGroup) {
	Dispatcher(print, 1, in, wg)
}

func print(line string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(line)
}
