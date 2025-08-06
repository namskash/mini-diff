package mini_diff

import(
	"fmt"
	"os"
	"bufio"
)

const (
	RESET = "\033[0m"
	RED = "\033[31m"
	GREEN = "\033[32m"
	YELLOW = "\033[33m"
	BLUE = "\033[34m"
	MAGENTA = "\033[35m"
	CYAN = "\033[36m"
	GRAY = "\033[37m"
	WHITE = "\033[97m"
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

	defer file1.Close() // Will run when the function exits
	defer file2.Close()

	for (scanner1.Scan() && scanner2.Scan()) || scanner2.Scan() {
		line1 := scanner1.Text()
		line2 := scanner2.Text()

		if line1 == line2 {
			fmt.Println(line1)
		} else {
			fmt.Printf(RED + "- %s\n" + GREEN + "+ %s\n" + RESET, line1, line2)
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
	
	var path1 string = "./example_files/example3.txt"
	var path2 string = "./example_files/example4.txt"

	file1 := fileReadWrapper(path1)
	file2 := fileReadWrapper(path2)

	if(file1 == nil || file2 == nil) {
		return
	}


	doTheDiff(file1, file2)

	fmt.Println("\nThank you for running mini_diff!\nNo LLMs were harmed in the making of this command line tool!")
}
