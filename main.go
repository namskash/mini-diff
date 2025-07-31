package main

import(
	"fmt"
	"os"
	"bufio"
)

func mapConsoleColors() map[string]string {
	return map[string]string {
		"reset": "\033[0m",
		"red": "\033[31m",
		"green": "\033[32m",
		"yellow": "\033[33m",
		"blue": "\033[34m",
		"magenta": "\033[35m",
		"cyan": "\033[36m",
		"gray": "\033[37m",
		"white": "\033[97m",
	}
}

func fileReadWrapper(path string) *os.File {
	file, err := os.Open(path)

	if(err != nil) {
		fmt.Println(err) // C like printf! Yay!!
		return nil
	}

	return file
}

func doTheDiff(file1 *os.File, file2 *os.File) {
	var consoleColors = mapConsoleColors()

	scanner1 := bufio.NewScanner(file1)
	scanner2 := bufio.NewScanner(file2)

	for (scanner1.Scan() && scanner2.Scan()) || scanner2.Scan() {
		line1 := scanner1.Text()
		line2 := scanner2.Text()

		if line1 == line2 {
			fmt.Printf(line1)
		} else {
			fmt.Printf(consoleColors["red"] + "- %s\n" + consoleColors["green"] + "+ %s\n" + consoleColors["reset"], line1, line2)
		}
  }

	err1 := scanner1.Err()
	err2 := scanner2.Err()

  if err1 != nil {
    fmt.Printf("Error scanning file: %v\n", err1)
  }
	if err2 != nil {
		fmt.Printf("Error scanning file: %v\n", err2)
	}
}

func main() {
	
	var path1 string = "./example_files/example1.txt"
	var path2 string = "./example_files/example2.txt"

	file1 := fileReadWrapper(path1)
	file2 := fileReadWrapper(path2)

	if(file1 == nil || file2 == nil) {
		return
	}

	defer file1.Close() // Will run when (here) `main` exits.

	doTheDiff(file1, file2)

	fmt.Println("\nThank you for running mini_diff!\nNo LLMs were harmed in the making of this command line tool!")
}
