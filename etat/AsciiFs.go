package etat

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Ascifs(s string, banner string) (string, error) {
	var (
		id  rune = 32
		tab map[rune][]string
		t   []string
	)
	d, err := os.Open(banner + ".txt")
	if err != nil {
		if err != nil {
			fmt.Println("Erreur lors de l'ouverture du fichier:", err)
			return "", err
		}
	}
	f := bufio.NewScanner(d)
	f.Split(bufio.ScanLines)
	for f.Scan() {
		t = append(t, f.Text())
	}
	tab = make(map[rune][]string)
	for i := 1; i < len(t); i += 9 {
		tab[id] = t[i : i+8]
		id++
	}
	text := ""
	l := strings.ReplaceAll(s, "\r\n", "\n")
	p := strings.Split(l, "\n")
	for _, v := range p {
		if v == "" {
			text += "\n"
		} else {
			for i := 0; i < 8; i++ {
				for j := 0; j < len(v); j++ {
					text += tab[rune(v[j])][i]
				}
				text += "\n"
			}
		}
	}
	return text, err
}

func Special(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] >= ' ' && s[i] <= '~' {
			return true
		}
	}
	return false
}
