package main
import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var inputReader1 *bufio.Reader
var input string
var err error


func test_read(){
	inputFile, inputError := os.Open("./read/test_read.txt")
	if inputError != nil {
		fmt.Printf("An error occurred on opening the inputfile\n" +
			"Does the file exist?\n" +
			"Have you got acces to it?\n")
		return // exit the function on error
	}
	defer func() {
		err := inputFile.Close()
		if err != nil{
			fmt.Println(err)
		}
	}()
	inputReader2 := bufio.NewReader(inputFile)
	for {
		inputString, readerError := inputReader2.ReadString('\n')
		if readerError == io.EOF {
			fmt.Println("EOF")
			return
		}
		fmt.Printf("file | The input was: %s", inputString)
	}
}
func main() {
	inputReader1 = bufio.NewReader(os.Stdin)
	fmt.Println("Please enter some input: ")
	input, err = inputReader1.ReadString('\n')
	if err == nil {
		fmt.Printf("stdin | The input was: %s\n", input)
	}

	test_read()




}
