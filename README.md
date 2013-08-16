id3.go
======

An id3 tag reader for Go

Example
=======

```go
package main

import(
    "github.com/levicole/id3-go"
    "fmt"
    "os"
)

func printTag(tag *id3.ID3Tag) {
    fmt.Printf("%s\n", tag.Title)
    fmt.Printf("%s\n", tag.Album)
    fmt.Printf("%s\n", tag.Artist)
    fmt.Printf("%s\n", tag.Year)
    fmt.Printf("%s\n", tag.Comment)
}

func main() {
    if len(os.Args) < 1 {
        fmt.Println("Please give me a file name")
        os.Exit(1)
    }

    tag, err := id3.ReadFromFile(os.Args[1])
    if err != nil {
        fmt.Println(err)
    } else {
        printTag(tag)
    }
}

```
