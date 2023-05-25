package main

/*
Project go-reloaded by Abdulrahman Khaled
package name = project codename = inkwell

* this code is one im proud of, my most complicated project
*/

import (
	"bufio"
	"fmt"
	"inkwell"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

//* start file after all is done
func OpenResultFile(filename string) error {
	cmd := exec.Command("code", filename)
	err := cmd.Start()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	// read from file
	terminalArguments := os.Args[1:]

	if len(terminalArguments) == 0 || terminalArguments[0] == "-h" || terminalArguments[0] == "--help" {
		fmt.Println("Project inkwell by A.rahman Khaled" + "\n" + "Usage: " + "inkwell <input file path> <output file path>")
		fmt.Println("-h or --help: displays help page")
		return
	}

	if len(terminalArguments) == 1 {
		fmt.Println("No output file detected")
		return
	}

	var finalStr string

	file, err := os.Open(terminalArguments[0])
	if err != nil {
		fmt.Println("Error opening file: " + terminalArguments[0] + " is not found")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		finalStr += scanner.Text() + " "
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file")
	}
	// file ops and mods
	stringInFileAsArray := strings.Fields(finalStr)

	for i := 0; i < len(stringInFileAsArray); i++ {
		var nextStr string
		//* assign a name to the current and next string in iteration
		currentStr := stringInFileAsArray[i]
		if i != len(stringInFileAsArray)-1 {
			nextStr = stringInFileAsArray[i+1]
		}

		//* handling spaced out tags

		incorrectTags := []string{"(hex,", "(up,", "(low,", "(cap,", "(bin,"}

		for _, incorrectTag := range incorrectTags {
			if currentStr == incorrectTag {
				pattern := regexp.MustCompile(`\d+\)`)
				if pattern.MatchString(nextStr) {
					currentStr += nextStr
					stringInFileAsArray = inkwell.RemoveElement(stringInFileAsArray, i+1)
					fmt.Println("spaced out tag handled")
				}
			}
		}

		//* check for tags, and apply function accordingly
		tag, iterator, tagCheck := inkwell.IsTag(currentStr)
		if iterator > i {
			iterator = i
		}
		if tagCheck {
			fmt.Println(tag + " was detected")
			if i == 0 {
				fmt.Println(tag + " was at the start of the text, hence invalid")
				continue
			}
			//* main loop to apply functions
			for j := i - iterator; j < i; j++ {
				//* handling hexadecimal
				if strings.Contains(tag, "hex") {
					//* hexadecimal to decimal conversion
					u, err := strconv.ParseInt(stringInFileAsArray[j], 16, 32)
					if err != nil {
						log.Fatal("Inkwell project error: Error in conversion from hexadecimal, may be invalid syntax" + "\n" +
							"more Info: " + err.Error())
					}
					stringInFileAsArray[j] = fmt.Sprint(u)
				}

				//* handling binary
				if strings.Contains(tag, "bin") {
					//* binary to decimal conversion
					u, err := strconv.ParseInt(stringInFileAsArray[j], 2, 32)
					if err != nil {
						log.Fatal("Inkwell project error: Error in conversion from binary, may be invalid syntax" + "\n" +
							"more Info: " + err.Error())
					}
					stringInFileAsArray[j] = fmt.Sprint(u)
				}

				//* handling uppercase
				if strings.Contains(tag, "up") {
					stringInFileAsArray[j] = strings.ToUpper(stringInFileAsArray[j])

				}

				//* handling lowercase
				if strings.Contains(tag, "low") {
					stringInFileAsArray[j] = strings.ToLower(stringInFileAsArray[j])
				}

				//* handling capitalization
				if strings.Contains(tag, "cap") {
					//! DEPRECATED METHOD, CHANGE LATER
					stringInFileAsArray[j] = strings.Title(stringInFileAsArray[j])
				}
				// TODO: make a new function to pull off the same thing
			}
			stringInFileAsArray = inkwell.RemoveElement(stringInFileAsArray, i)
		}

		//* vowel handling
		if inkwell.IsVowelBegin(currentStr) {
			if i == 0 {
				continue
			}
			if stringInFileAsArray[i-1] == "a" {
				stringInFileAsArray[i-1] = "an"
			} else if stringInFileAsArray[i-1] == "A" {
				stringInFileAsArray[i-1] = "An"
			}
		}
	}

	finalStr = ""

	for _, v := range stringInFileAsArray {
		finalStr += v + " "
	}

	finalStr = inkwell.PuncConv(finalStr)
	finalStr = inkwell.QouteConv(finalStr)

	//* output to file
	filename := terminalArguments[1]

	content := finalStr

	er := inkwell.WriteToFile(filename, content)
	if er != nil {
		fmt.Println("Error writing file")
		return
	}

	fmt.Println("File written successfully.")

	// open file after execution
	if err := OpenResultFile(terminalArguments[1]); err != nil {
		fmt.Printf("Error opening file: %v\n", err)
	}
}
