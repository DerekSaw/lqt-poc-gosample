
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"errors"
)


// FLAGS & PARAMS
var DataDirPtr   = flag.String("d", "", "GKR path in `Data Directory`")
var InputFilePtr = flag.String("i", "", "The `Input Filename`")
var ProcessPtr   = flag.String("p", "v+g", "`Process List` to run")
var VersionPtr   = flag.Bool("v", false, "Output version")
var OutputFilePtr = flag.String("o", "", "`Output Filename`")


// package auto INIT
func init() {
	flag.Parse()
}


func printArrayString(arr []string) {
	fmt.Printf("[%d]", len(arr))
	for index, elem := range arr {
		fmt.Printf("(%d:%s)", index, elem)
	}
	fmt.Printf("\n")
}


func readInput(parserFn func(strarr []string)) (err error) {
	fmt.Println("Processing", *InputFilePtr)

	file, err := os.Open(*InputFilePtr)
	if err != nil {
		return err
	}
	defer file.Close()

	// start read
	reader := bufio.NewReader(file)
	count := 0
	fieldCount := 0

	var line string
	var sep string
	for {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			break
		}

		// trim line end
		line = strings.TrimRight(line, "\r\n")
		count++

		// obtain info from header
		if count == 1 {
			pos := strings.IndexAny(line, "|\t")
			if pos == -1 {
				return errors.New("Unable to find separator")
			}
			sep = string(line[pos])
			fmt.Printf("Seperator: %s\n", sep)
			if sep == "" {
				return errors.New("Separator empty")
			}
		}

		arr := strings.Split(line, sep)

		// get field count via header
		if count == 1 {
			fieldCount = len(arr)
		}

		// process fields
		if len(arr) == fieldCount {
			// TODO: gorountine here
			parserFn(arr)
		}

		// break out if EOF
		if err != nil {
			break
		}
	}
	if err != io.EOF {
		fmt.Printf(" > Failed with error: %v\n", err)
		return err
	}
	fmt.Printf("Total lines read: %d\n", count)
	return
}


func parseInput(arr []string) {
	//arr := strings.Split(line, sep)
	printArrayString(arr)
}


// MAIN
func main() {
	fmt.Println("data dir: ", *DataDirPtr)
	fmt.Println("Input file: ", *InputFilePtr)
	fmt.Println("Processes: ", *ProcessPtr)
	fmt.Println("Version: ", *VersionPtr)

	if *InputFilePtr != "" {
		readInput(parseInput)
	}
}




