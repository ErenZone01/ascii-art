package main

import (
	"os"
)

func main() {
	if len(os.Args) >= 2 && len(os.Args) <= 3 {
		arg1 := os.Args[1]
		arg2 := DefineArg()
		if CaseResolved(arg1) {
			return
		}
		texte := foundFile(arg2)
		asciiArt(arg1, texte)
	}
}
