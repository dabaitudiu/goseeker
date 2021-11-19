package tool

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/pkg/errors"
)

type WikiDoc struct {
	ID    string `json:"id"`
	Url   string `json:"url"`
	Title string `json:"title"`
	Text  string `json:"text"`
}

// LoadWiki
func LoadWiki(filename string, index int) (int, error) {
	file, openErr := os.Open(filename)
	if openErr != nil {
		log.Fatal(openErr)
	}
	defer file.Close()

	counter := index

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == nil || err == io.EOF {
			if line != "" {
				var doc WikiDoc
				if jsonErr := json.Unmarshal([]byte(line), &doc); jsonErr == nil {
					writeFileName := fmt.Sprintf("/Users/lisbon/Desktop/seeker/docs/AA/doc_%02d", counter)
					writeErr := WriteFile(writeFileName, doc.Text)
					if writeErr != nil {
						return -1, errors.Wrap(writeErr, "fail to write json to file")
					}
				} else {
					fmt.Print(jsonErr)
					return -1, errors.Wrap(jsonErr, "failed in json unmarshal")
				}
				counter += 1
			}
			if err == io.EOF {
				break
			}
		} else {
			fmt.Println(err)
			break
		}
	}

	fmt.Printf("counter: %d\n", counter)
	return counter, nil
}
