package id3

import(
    "os"
    "log"
    "bufio"
    "errors"
    "fmt"
)

var (
    ErrID3TagNotFound = errors.New("ID3Tag not found")
)

type ID3Tag struct {
    Title   string
    Album   string
    Artist  string
    Year    string
    Comment string
    Genre   string
}

func readTag(bytes []byte) []rune {
    tag := make([]rune, len(bytes))
    for _, r := range bytes {
        tag = append(tag, rune(r))
    }
    return tag
}


func ReadFromFile(filePath string) (*ID3Tag, error) {
    inFile, err := os.Open(filePath)

    if err != nil {
        log.Fatal(err)
    }

    reader := bufio.NewReader(inFile)

    header := make([]byte, 3)

    if _, err = inFile.Seek(-128, 2); err != nil {
        log.Fatal(err)
    }

    _, err = reader.Read(header)
    if err != nil {
        log.Fatal(err)
    }


    if string(header) != "TAG" {
        fmt.Println(string(header))
        return nil, ErrID3TagNotFound
    }

    tag := make([]byte, 125)
    _, err = reader.Read(tag)
    if err != nil {
        log.Fatal(err)
    }

    title     := readTag(tag[0:30])
    album     := readTag(tag[30:60])
    artist    := readTag(tag[60:90])
    year      := readTag(tag[90:94])
    comment   := readTag(tag[94:124])

    id3tag := &ID3Tag{ string(title), string(album), string(artist), string(year), string(comment), "Genre"}
    return id3tag, nil
}
