package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	RESET   = "\033[0m"
	RED     = "\033[31m"
	GREEN   = "\033[32m"
	YELLOW  = "\033[33m"
	BLUE    = "\033[34m"
	MAGENTA = "\033[35m"
	CYAN    = "\033[36m"
	GRAY    = "\033[37m"
	WHITE   = "\033[97m"
)

func readArgAt(index int) string {
	if index < len(os.Args) {
		return os.Args[index]
	}
	return ""
}

func fileReadWrapper(path string) *os.File {
	file, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return file
}

func doTheDiff(file1 *os.File, file2 *os.File) {
	scanner1 := bufio.NewScanner(file1)
	scanner2 := bufio.NewScanner(file2)

	defer file1.Close() // Will run when the function exits
	defer file2.Close()

	for (scanner1.Scan() && scanner2.Scan()) || scanner2.Scan() {
		line1 := scanner1.Text()
		line2 := scanner2.Text()

		if line1 == line2 {
			fmt.Println(line1)
		} else {
			fmt.Printf(RED+"- %s\n"+GREEN+"+ %s\n"+RESET, line1, line2)
		}
	}

	err1 := scanner1.Err()
	err2 := scanner2.Err()

	if err1 != nil {
		fmt.Printf("%sError scanning file: %v%s\n", RED, err1, RESET)
	}
	if err2 != nil {
		fmt.Printf("%sError scanning file: %v%s\n", RED, err2, RESET)
	}
}

func main() {
	var path1 string = readArgAt(1)
	var path2 string = readArgAt(2)

	if path1 == "" || path2 == "" {
		fmt.Printf("%sTwo paths not found!%s Usage: ./main <path-to-file-1> <path-to-file-1>%s\n", RED, MAGENTA, RESET)
		return
	}

	var file1 *os.File = fileReadWrapper(path1)
	var file2 *os.File = fileReadWrapper(path2)

	if file1 == nil {
		return
	}
	if file2 == nil {
		return
	}

	doTheDiff(file1, file2)

	fmt.Println("\nThank you for running mini_diff!\nNo LLMs were harmed in the making of this command line tool!")
}
