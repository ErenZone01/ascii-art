package main

import (
	"fmt"
	"os"
	"strings"
)

func foundLine(texte string) []string {
	var compteur int = 0
	var phrase string
	var tabPhrase []string

	for i := 0; i < len(texte); i++ {
		if texte[i] != '\n' {
			phrase += string(texte[i])
		} else {
			compteur++
			if compteur == 9 {
				tabPhrase = append(tabPhrase, phrase)
				phrase = ""
				compteur = 0
				continue
			}
			if compteur != 9 {
				phrase += "\n"
			}

		}

	}
	return tabPhrase
}

func foundposition(position []int, tabPhrase []string) []string {
	var textePhrase []string
	for i := 0; i < len(position); i++ {
		for j := 0; j < len(tabPhrase); j++ {
			if position[i] == j {
				textePhrase = append(textePhrase, tabPhrase[j])
			}
		}
	}
	return textePhrase
}

func compraison(tabASCII []rune, arg1 string, position []int, tab []int) []int {
	var actif bool = false

	var special bool = true

	for i := 0; i < len(arg1); i++ {
		for j := 0; j < len(tabASCII); j++ {
			if i < len(arg1)-1 && (string(arg1[i]) == "\\" && string(arg1[i+1]) == "n") {
				actif = true
				break
			} else {
				if !actif {
					if rune(arg1[i]) == tabASCII[j] {
						position = append(position, tab[j])
						special = true
						break
					} else {
						special = false
					}
				} else {
					actif = false
					break
				}

			}
		}
		if !special {
			fmt.Println("Ce programme ne prend que les carcateres de la table ASCII")
			os.Exit(1)
		}
	}
	return position
}

func HorzontalLine(x [9][]string) string {
	var texte2 string
	for i, ligne := range x {
		if i > 0 {
			for _, part := range ligne {
				if len(part) > 0 {
					part = part[:len(part)-1]
				}
				texte2 += part
				texte2 += " "

			}
			if len(ligne[0]) == 1 && ligne[0] != "\n" {
				continue
			} else {
				texte2 += "\n"
			}

		}
	}
	return texte2

}

func SplitTexte(textePhrase []string) [9][]string {
	x := [9][]string{}
	for _, text := range textePhrase {
		for i, y := range strings.Split(text, "\n") {
			x[i] = append(x[i], y)

		}
	}
	return x
}

func ajoutASCII() []rune {
	var tabASCII []rune
	for i := ' '; i <= '~'; i++ {
		tabASCII = append(tabASCII, i)
	}
	return tabASCII
}

func ajoutNmbr(arg1 string) []int {
	var tab []int
	for i := 0; i < 95; i++ {
		tab = append(tab, i)
	}
	return tab
}

func CaseResolved(arg1 string) bool {
	var actif bool = true
	for i := 0; i < len(arg1); i = i + 2 {
		if (i+1 != len(arg1)) && (arg1[i] == '\\' && arg1[i+1] == 'n') {
			actif = false
		} else {
			actif = true
			break
		}

	}
	if !actif {
		for i := 0; i < len(arg1); i = i + 2 {
			fmt.Println()
		}
		return true
	}

	if len(os.Args) >= 2 && arg1 == "" {
		return true
	}

	if len(os.Args) >= 2 && arg1 == "\\n" {
		fmt.Println()
		return true
	}
	return false
}

func DefineArg() string {
	var arg2 string
	if len(os.Args) == 2 {
		arg2 = "standard"
	} else {
		arg2 = os.Args[2]
	}
	return arg2
}

func foundFile(arg2 string) string {
	file, err := os.ReadFile(arg2 + ".txt")
	if err != nil {
		fmt.Println("le fichier est introuvable")
		os.Exit(1)
	}
	return string(file)
}

func asciiArt(arg1 string, texte string) {
	var arguments = strings.Split(arg1, "\\n")

	for i := 0; i < len(arguments); i++ {
		if len(arguments[i]) == 0 {
			fmt.Println()
			continue
		}
		var tab []int = ajoutNmbr(arg1)
		var tabPhrase []string
		var tabASCII []rune = ajoutASCII()
		var position []int
		var textePhrase []string

		var texte2 string
		var x = [9][]string{}

		if arguments[i] != "\n" {
			arg1 = arguments[i]
			tab = ajoutNmbr(arg1)
			position = compraison(tabASCII, arg1, position, tab)
			tabPhrase = foundLine(texte)
			textePhrase = foundposition(position, tabPhrase)
			x = SplitTexte(textePhrase)
			texte2 = HorzontalLine(x)
			fmt.Print(texte2)
		}
	}
}
