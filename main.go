package main

import(
	"fmt"
	"os"
	"bufio"
)

func fileReadWrapper(path string) *os.File {
	file, err := os.Open(path)

	if(err != nil) {
		fmt.Println(err) // C like printf! Yay!!
		return nil
	}

	return file
}

func doTheDiff(file1 *os.File, file2 *os.File) {
	scanner1 := bufio.NewScanner(file1)
	scanner2 := bufio.NewScanner(file2)

	for (scanner1.Scan() && scanner2.Scan()) || scanner2.Scan() {
		line1 := scanner1.Text()
		line2 := scanner2.Text()

		fmt.Printf("Line1: %s\nLine2: %s\n\n", line1, line2)
		if line1 == line2 {
			fmt.Printf(line1)
		} else {
			fmt.Printf("- %s\n+ %s\n", line1, line2)
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
