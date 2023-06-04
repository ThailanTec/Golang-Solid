package main

import "strings"

type Journal struct {
	News []string
}

func (j *Journal) String() string {
	return strings.Join(j.News, "\n")
}

func (j *Journal) RemoveNews(news string) {
	//...
}

func (j *Journal) UpdateNews(news string) {
	//...
}

func (j *Journal) AddNews(news string) {
	//...
}

func (j *Journal) DeleteNews(news string) {
	//...
}

// Nâo deveria está presente

func (j *Journal) PrintNews(news string) {
	//...
}

func (j *Journal) LoadNewsFromFiles(news string) {
	//...
}
func (j *Journal) LoadNewsFromWeb(news string) {
	//...
}

func main() {

}
