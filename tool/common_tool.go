package tool

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
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
	err := ioutil.WriteFile(filename, data, os.ModeAppend)
	if err != nil {
		return errors.Wrap(err, "fail to call func:ioUtil.WriteFile()")
	}
	return nil
}

// AppendFile
func AppendFile(filename string, input string) error {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return errors.Wrap(err, "fail to open file")
	}
	defer f.Close()

	if _, err = f.WriteString("\n" + input); err != nil {
		return errors.Wrap(err, "fail to write string")
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

// ConvStrToHex 将token的不同语言形式统一转为Byte数组，再转成16进制(小写字母)
func ConvStrToHex(s string) string {
	return fmt.Sprintf("%x", []byte(s))
}

// ConvHexToStr 将16进制字符串转化为正常语言的字符串
func ConvHexToStr(s string) (string, error) {
	ori, err := hex.DecodeString(s)
	if err != nil {
		return "", errors.Wrap(err, "fail to decode hex string")
	}
	return fmt.Sprintf("%s", ori), nil
}
