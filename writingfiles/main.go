package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/*
Make a save method on this struct to write a text-file with Title
as its name and Body as its content. Make a load function which,
given the title string, reads the corresponding text-file. Use *Page as an
argument because the structs could be quite large and we don’t want to make
copies of them in memory. Use the ioutil functions to read and write from/to a file.
*/

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	fileName := p.Title
	
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("Failed to create file: %s\n", err)
		return err
	}
	defer file.Close() 

	writer := bufio.NewWriter(file)

	_, err = writer.Write(p.Body)
	if err != nil {
		fmt.Printf("Failed to write to file: %s\n", err)
		return err
	}

	err = writer.Flush()
	if err != nil {
		fmt.Printf("Failed to flush writer: %s\n", err)
		return err
	}

	return nil

}

func load(title string) (*Page, error) {

	file, err := os.Open(title)
	if err != nil {
	  panic(err)
	}
	defer file.Close()

	var thePage Page
	thePage.Title = title
	thePage.Body, _ = io.ReadAll(file)

	if (err != nil) {
		return nil, err
	}

	return &thePage, nil
}

func main() {
	

	p1 := &Page{Title: "test1.txt", Body: []byte("This is a test.")}
	p1.save()

	p2, err := load("test.txt")
	if err != nil {
		fmt.Printf("Failed to load file: %s\n", err)
		return
	}

	fmt.Println(string(p2.Body))
}