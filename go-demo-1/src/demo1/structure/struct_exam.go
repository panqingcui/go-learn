package main

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {
	var Book1 Books
	var Book2 Books

	/* book1*/
	Book1.title = "go 语言"
	Book1.author = "mmss"
	Book1.subject = "go 语言教程"
	Book1.book_id = 1
	/*打印book1*/
	printfBooks(Book1)
	/*Book2*/
	Book2.title = "python教程"
	Book2.author = "mzss"
	Book2.subject = "python教程"
	Book2.book_id = 2
	printfBooks(Book2)
	fmt.Printf("Book1 title %s \n", Book1.title)
	fmt.Printf("Book1 author %s \n", Book1.author)
	fmt.Printf("Book1 subject %s \n", Book1.subject)
	fmt.Printf("Book1 book_id %d \n", Book1.book_id)
	fmt.Printf("Book2 title %s \n", Book2.title)
	fmt.Printf("Book2 author %s \n", Book2.author)
	fmt.Printf("Book2 subject %s \n", Book2.subject)
	fmt.Printf("Book2 book_id %d \n", Book2.book_id)

}

/*打印*/
func printfBooks(books Books) {

	fmt.Printf("Book title %s \n", books.title)
	fmt.Printf("Book author %s \n", books.author)
	fmt.Printf("Book subject %s \n", books.subject)
	fmt.Printf("Book book_id %d \n", books.book_id)
}
