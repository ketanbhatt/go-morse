package main

import (
	"strings"
	"os"
	"regexp"
	"github.com/atotto/clipboard"
	"github.com/fatih/color"
)

var textToMorseMap = map[string]string{
  "0": "-----",
  "1": ".----",
  "2": "..---",
  "3": "...--",
  "4": "....-",
  "5": ".....",
  "6": "-....",
  "7": "--...",
  "8": "---..",
  "9": "----.",
  "a": ".-",
  "b": "-...",
  "c": "-.-.",
  "d": "-..",
  "e": ".",
  "f": "..-.",
  "g": "--.",
  "h": "....",
  "i": "..",
  "j": ".---",
  "k": "-.-",
  "l": ".-..",
  "m": "--",
  "n": "-.",
  "o": "---",
  "p": ".--.",
  "q": "--.-",
  "r": ".-.",
  "s": "...",
  "t": "-",
  "u": "..-",
  "v": "...-",
  "w": ".--",
  "x": "-..-",
  "y": "-.--",
  "z": "--..",
  ".": ".-.-.-",
  ",": "--..--",
  "?": "..--..",
  "!": "-.-.--",
  "-": "-....-",
  "/": "-..-.",
  "@": ".--.-.",
  "(": "-.--.",
  ")": "-.--.-",
  " ": "/",
}

var morseToTextMap = reverseMap(textToMorseMap)
var isMorse = regexp.MustCompile("^[.\\- \\/]*$").MatchString

func reverseMap(originalMap map[string]string) map[string]string {
    revMap := make(map[string]string)

    for k, v := range originalMap {
        revMap[v] = k
    }
    return revMap
}

func textToMorse(inputString string) string {
	var tString []string

	for _, val := range inputString {
		tString = append(tString, textToMorseMap[string(val)])
	}

	return strings.Join(tString, " ")
}

func morseToText(inputString string) string {
	var tString []string

	inputStringList := strings.Split(inputString, " ")

	for _, val := range inputStringList {
		tString = append(tString, morseToTextMap[string(val)])
	}

	return strings.Join(tString, "")

}

func main() {
	inputString := strings.ToLower(os.Args[1])

	var tString string

	if isMorse(inputString) {
		tString = morseToText(inputString)
	} else {
		tString = textToMorse(inputString)
	}

	color.Green(tString)
	color.White("")

	clipboard.WriteAll(tString)
	color.Yellow("Translated string copied to clipboard!")
}
