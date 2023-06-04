package main

import (
	"io/ioutil"
)

// Modo certo de separar
type JournalClassic struct{}

func printeNews(filename string, jounal Journal) {
	_ = ioutil.WriteFile(filename, []jounal.String(), 0064)
}

func PrintNews(news string) {
	//...
}

func LoadNewsFromFiles(news string) {
	//...
}
func LoadNewsFromWeb(news string) {
	//...
}
