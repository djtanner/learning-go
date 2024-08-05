package main

import (
	"fmt"
	"io"
	"os"
)

type Book struct {
    title string
    price float64
    quantity int
}

func main() {
	var books []Book

  file, err := os.Open("products.txt")
  if err != nil {
    panic(err)
  }
  defer file.Close()
  
  for {
    var v1 string
	var v2 float64
	var v3 int

	//scan a line of text like "The ABC of Go";25.5;1500
	//into v1, v2, v3
	_, err = fmt.Fscanf(file, "%q;%f;%d\n", &v1, &v2, &v3)

    if err != nil {
		if err == io.EOF {
            break
        }
	  fmt.Println(err)
      break
    }
    
    var theBook Book
    theBook.title = v1
    theBook.price = v2
    theBook.quantity = v3

    books = append(books,theBook)
  }

  for _, ix := range books {
    fmt.Println(ix)
  }


}