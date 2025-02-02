package main

import (
	"bufio"
	"os"
)

func UserPromptLoop(prompt string) {
	reader := bufio.NewReader(os.Stdin)
	for {
		input, _ := reader.ReadString('\n')
	}
}

func DoRequestToRefiner(prompt string) {

}

func DoRequestSemanticParser(prompt string) []byte {

}
