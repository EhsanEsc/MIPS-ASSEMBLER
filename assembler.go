package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const instructionSignaturesFileAddress = "MIPSInstructions.txt"

var instructionSignatures = make(map[string][]string) // "Name" -> ["Inputs", "Binary Format"]

func readInstructionSignatures() {
	data, err := ioutil.ReadFile(instructionSignaturesFileAddress)
	if err != nil {
		panic(err)
	}
	signatures := strings.Split(string(data), "\n")
	for _, instSignature := range signatures {
		sigParts := strings.Split(instSignature, " ")
		if len(sigParts) != 2 {
			fmt.Println("Reading Instruction Signature Failed")
			os.Exit(1)
		}
		sig1, sig2 := sigParts[0], sigParts[1]
		sig2 = strings.Replace(sig2, "\r", "", -1)
		instName := strings.Split(sig1, ",")[0]
		instName = strings.ToLower(instName)
		instructionSignatures[instName] = []string{strings.Join(strings.Split(sig1, ",")[1:], ","), sig2}
	}
}

func convertToBinaryFormat(fileAddress string) {
	data, err := ioutil.ReadFile(fileAddress)
	if err != nil {
		panic(err)
	}
	instructionList := strings.Split(string(data), "\n")
	for _, instruction := range instructionList {
		instParts := strings.Split(instruction, " ")
		if len(instParts) == 1 {
			instParts = []string{"nop", ""}
		}
		instName := instParts[0]
		instName = strings.ToLower(instName)
		instInputs := instParts[1]
		instInputs = strings.Replace(instInputs, ")", "", -1)
		instInputs = strings.Replace(instInputs, "(", ",", -1)
		instInputs = strings.Replace(instInputs, "\r", "", -1)
		instInputsSplited := strings.Split(instInputs, ",")

		var result string
		var numbers = make([]int, 322)
		var registers = make([]int, 322)
		if instSignature, ok := instructionSignatures[instName]; ok {
			signatureInputs := strings.Split(instSignature[0], ",")
			for i, sigInput := range signatureInputs {
				if len(sigInput) < 1 {
					continue
				} else if sigInput[0:1] == "N" {
					index, _ := strconv.Atoi(sigInput[1:])
					numbers[index], err = strconv.Atoi(instInputsSplited[i])
				} else if sigInput[0:1] == "R" {
					index, _ := strconv.Atoi(sigInput[1:])
					registers[index], _ = strconv.Atoi(instInputsSplited[i][1:])
				} else {
					fmt.Printf("Undefined token in instruction signutares. Skiping")
				}
			}
			for _, sigPart := range strings.Split(instSignature[1], ",") {
				if sigPart[0:1] == "N" {
					index, _ := strconv.Atoi(strings.Split(sigPart, ":")[0][1:])
					numberOfBits, _ := strconv.Atoi(strings.Split(sigPart, ":")[1])
					ss := fmt.Sprintf("%b", numbers[index])
					for len(ss) < numberOfBits {
						ss = "0" + ss
					}
					result += ss
				} else if sigPart[0:1] == "R" {
					index, _ := strconv.Atoi(sigPart[1:])
					ss := fmt.Sprintf("%.5b", registers[index])
					result += ss
				} else {
					result += sigPart
				}
			}
		} else {
			fmt.Println("Instruction Not Found. Skiping...")
			result = instruction
		}
		fmt.Println(result)
	}
}

func main() {
	flag.Parse()
	var fileAddress string
	if flag.NArg() > 0 {
		fileAddress = filepath.FromSlash(flag.Arg(0))
	} else {
		fmt.Print("Enter Instruction File Address: ")
		fmt.Scanln(&fileAddress)
	}

	readInstructionSignatures()
	convertToBinaryFormat(fileAddress)
}
