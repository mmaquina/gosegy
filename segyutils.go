package main

import (
    "fmt"
    "os"
    "log"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"

)


// Function to convert EBCDIC-encoded text to UTF-8
func ebcdicToUtf8(ebcdicData []byte) (string, error) {
	// IBM037 is a common EBCDIC code page. We use the charmap package here.
	decoder := charmap.CodePage037.NewDecoder()

	// Perform the transformation from EBCDIC to UTF-8
	utf8Data, _, err := transform.String(decoder, string(ebcdicData))
	if err != nil {
		return "", err
	}

	return utf8Data, nil
}



func readEbcdicHeader(file string) (string, error) {
    log.SetPrefix("readEbcdicHeader: ")
    log.SetFlags(0)

    buffer := make([]byte, 3200)

    f, err := os.Open(file)
    if err != nil {
        log.Fatal(err)
    }
    
    nb, err := f.Read(buffer)

    if err != nil {
        log.Fatal(err)
    }
    f.Close()

    fmt.Printf("Just read %d bytes.\nThe EBCDIC header for %s is the following:\n", nb, file) 

    return ebcdicToUtf8(buffer)
}

func main(){
    file := "hdr1.sgy"
    header, err := readEbcdicHeader(file)
    if err != nil {
        return
    }
    
    for i:=0; i<3200; i+=80 {
        fmt.Println(header[i:i+80])
    }
    
}

