package main

import (
	"fmt"
	"net/url"
	"os"
	"strings"
)

var entryCount = 0

type Journal struct {
	entries []string
}

func (j *Journal) AddEntry(text string) int {
	entryCount++
	entry := fmt.Sprintf("%d. %s", entryCount, text)
	j.entries = append(j.entries, entry)
	return entryCount
}

func (j *Journal) Print() {
	str := strings.Join(j.entries, "\n")
	fmt.Println(str)
}

func (j *Journal) RemoveEntry(index int) {
	//removes an entry

}

//SOC(Separation of concerns)

// Save this will break the single responsibility principle
// Let's say we want to be persistent with the entries of the file,
// so we are going to save the entries and then be able to load them too
func (j *Journal) Save(filename string) {
	_ = os.WriteFile(filename, []byte(strings.Join(j.entries, "\n")), 0644)
}

func (j *Journal) Load(filename string) {
	data, _ := os.ReadFile(filename)
	j.entries = strings.Split(string(data), "\n")
}

func (j *Journal) LoadFromWeb(url *url.URL) {
	//loads from web
}

//Now attaching the save to the journal will break the srp. The responsibility of the journal
//is to manage entries and not to persist them on a drive or a web or something
//persistence can be handled with a different component

//and addition to this, if other types have to save a file to the system, then a new function
//will have to be created, but if it is in a different component or package, it can easily be
//reused

//????????????? Solution ???????????? //
/*
1. making persistence a separate function or moving it to a separate package or component
*/

// creating a function to save the journal and other types that are like the journal

var LineSeparator = "\n"

func SaveToFile(j *Journal, filename string) {
	_ = os.WriteFile(filename, []byte(strings.Join(j.entries, LineSeparator)), 0644)
}

// CREATING A STRUCT

type Persistence struct {
	lineSeparator string
}

func (p *Persistence) SaveToFile(j *Journal, filename string) {
	_ = os.WriteFile(filename, []byte(strings.Join(j.entries, p.lineSeparator)), 0644)
}

func main() {
	//Test cases
	j := Journal{}
	j.AddEntry("I ate banku today")
	j.AddEntry("I think the new keyword works too")
	j.Print()

	//saving to file
	//1. functional approach
	SaveToFile(&j, "functional_save.txt")

	//modular or struct save
	p := Persistence{"\n"}
	p.SaveToFile(&j, "package_save.txt")
}
