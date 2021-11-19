package tool

import (
	"io/ioutil"
	"strings"

	gset "github.com/deckarep/golang-set"
	"github.com/pkg/errors"
)

// LoadFile
func LoadFile(filename string) (string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", errors.Wrapf(err, "error when loading file %s", filename)
	}
	return string(data), nil
}

// WriteFile
func WriteFile(filename string, input string) error {
	data := []byte(input)
	err := ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return errors.Wrap(err, "fail to call func:ioUtil.WriteFile()")
	}
	return nil
}

// ConvStrToSet 将sql的string结果转化为set
func ConvStrToSet(s string) gset.Set {
	set := gset.NewSet()
	strList := strings.Split(s, ",")
	for _, e := range strList {
		set.Add(e)
	}
	return set
}
