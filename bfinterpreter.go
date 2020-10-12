package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

var scanner = bufio.NewScanner(os.Stdin)
var input uint8
var pointer int           // obviously haha
var programm [30000]uint8 // my array, the user code will in there
var brainfuckcode string  // user code

func main() {

	pointer = 0

	fmt.Println("Enter youre BrainFuck code and let the magic begin (:")
	// fmt.Println("Do You want to write your own code or read a file? (Text or File)")
	// scanner.Scan()
	// choice := string(scanner.Text())

	code, err := ioutil.ReadFile("fileName.bf")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(-1)
	}
	brainfuckcode := code

	//	scanner.Scan()
	//	brainfuckcode = string(scanner.Text())

	fmt.Println()

	for i := 0; i < len(brainfuckcode); i++ {

		switch string(brainfuckcode[i]) {
		case "<":
			// eine zeile zurück
			pointer--
			if pointer == 30000 {
				pointer = 0
			}
		case ">":
			// zeile vor
			pointer++
			if pointer == 0 {
				pointer = 30000
			}

		case "+":
			// zelle plus eins
			programm[pointer]++
			if programm[pointer] > 255 {
				programm[pointer] = 0
			}

		case "-":
			// Zelle minus eins
			programm[pointer]--
			if programm[pointer] < 0 {
				programm[pointer] = 255
			}
		case "[":
			// loop anfang, wenn 0 raus
			p := i

			if programm[pointer] == 0 {
				count := 0

				for p < len(brainfuckcode) {
					if string(brainfuckcode[p]) == "]" && count == 0 && programm[pointer] == 0 {
						i = p
						break
					} else if string(brainfuckcode[p]) == "[" {
						count++
					} else if string(brainfuckcode[p]) == "]" {
						count--
					} else {
						p++
					}

				}
			}

		case "]":
			// loop ende, wenn nicht null zum aufang
			if programm[pointer] != 0 {
				count := 0
				p := i

				p--

				for p > 0 {
					if string(brainfuckcode[p]) == "[" && count == 0 && programm[pointer] != 0 {
						p--
						i = p
						break
					} else if string(brainfuckcode[p]) == "]" {
						count--
					} else {
						p--
					}
				}
			}

		case ".":
			// output
			fmt.Print(string(programm[pointer]))
		case ",":
			// input

			scanner.Scan()
			input = scanner.Bytes()[0]
			programm[pointer] = input
		}

	}
	fmt.Println()
	fmt.Println("Tadaaa!")

}
