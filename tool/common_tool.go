package tool

import (
	"io/ioutil"

	"github.com/pkg/errors"
)

// LoadArticle
func LoadArticle(filename string) (string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", errors.Wrap(err, "error when loading article")
	}
	return string(data), nil
}
