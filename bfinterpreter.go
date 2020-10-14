package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

var scanner = bufio.NewScanner(os.Stdin)

// Menu, choices to add Brainfuckcode
func code() []byte {
	returns := make([]byte, 1000000)

	fmt.Println("Do You want to write your own Brainfuck code or read a file? ")
	fmt.Println("write Text or File")
	scanner.Scan()

	if scanner.Text() == "Text" {
		fmt.Println("Enter your Brainfuck code and let the magic begin (:")
		scanner.Scan()
		returns = scanner.Bytes()

	} else if scanner.Text() == "File" {

		code, err := ioutil.ReadFile("fileName.bf")

		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s\n", err)
			os.Exit(-1)

		}
		returns = code
	}
	return returns
}

func main() {

	bfCode := code()

	programm := make([]byte, 30000) // my array, the user code will in there
	prgPointer := 0

	fmt.Println()

	for codePointer := 0; codePointer < len(bfCode); codePointer++ {

		switch bfCode[codePointer] {

		case 60: // Instruction "<"
			if prgPointer != 29999 {
				prgPointer--
			} else if prgPointer == 29999 {
				prgPointer = 0
			}

		case 62: // Instruction ">"
			if prgPointer != 29999 {
				prgPointer++
			} else if prgPointer == 0 {
				prgPointer = 0
			}

		case 43: // Instruction "+"
			if programm[prgPointer] != 255 {
				programm[prgPointer]++
			} else if programm[prgPointer] == 255 {
				programm[prgPointer] = 0
			}

		case 45: // Instruction "-"
			if programm[prgPointer] != 0 {
				programm[prgPointer]--
			} else if programm[prgPointer] == 0 {
				programm[prgPointer] = 255
			}

		case 91: // Instruction "["
			if programm[prgPointer] == 0 {

				loopCount := 0
				codePointer++

				for codePointer < len(bfCode) {
					if bfCode[codePointer] == 93 && loopCount == 0 {
						break
					} else if bfCode[codePointer] == 91 {
						loopCount++
					} else if bfCode[codePointer] == 93 {
						loopCount--
					} else {
						codePointer++
					}

				}
			}

		case 93: // Instruction "]"
			if programm[prgPointer] != 0 {

				loopCount := 0
				codePointer--

				for codePointer < len(bfCode) {
					if bfCode[codePointer] == 91 && loopCount == 0 {
						codePointer--
						break
					} else if bfCode[codePointer] == 91 {
						loopCount--
					} else if bfCode[codePointer] == 93 {
						loopCount++
					} else {
						codePointer--
					}
				}
			}

		case 46: // Instruction "."
			fmt.Print(string(programm[prgPointer]))

		case 44: // Instruction ","
			fmt.Println("Please enter one character..")
			scanner.Scan()
			programm[prgPointer] = scanner.Bytes()[0]
		}

	}
	fmt.Println()
	fmt.Println("Tadaaa!")

}
