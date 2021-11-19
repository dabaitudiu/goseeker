package handlers

import (
	"fmt"
	"goseeker/tool"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"github.com/yanyiwu/gojieba"
)

const (
	InvertedIndexDB = "t_inverted_index"
	TokenColumn     = "c_token"
	DocColumn       = "c_doc_id"
	PosColumn       = "c_index_file"

	docIndex = 0
	posIndex = 1
)

// NormalizeText 将文章归一化，去掉非中文字符
func NormalizeText(textStr string) string {
	chineseReg := regexp.MustCompile("[^\u4e00-\u9fa5]")
	return chineseReg.ReplaceAllString(textStr, "")
}

// TextToTokens 将文章转化成token, tokenDict: map["token"]["doc_001","doc_001:1,2,3,4"]
func TextToTokens(textStr string, textID string) map[string][]string {
	tokenDict := make(map[string][]string)
	textStr = NormalizeText(textStr)
	x := gojieba.NewJieba()
	defer x.Free()

	words := x.CutAll(textStr)
	log.Printf("tokenize finished, total tokens: %d\n", len(words))

	for i, word := range words {
		if v, ok := tokenDict[word]; !ok {
			pos := textID + ":" + strconv.Itoa(i)
			tokenDict[word] = []string{textID, pos}
		} else {
			if len(v) < 2 {
				panic("length of tokenDict string list smaller than 2")
			}
			tokenDict[word][1] += "," + strconv.Itoa(i)
		}
	}

	return tokenDict
}

func LoadAllTokensFromDB() (map[string][]string, error) {
	tokenDict, err := tool.Query("*", InvertedIndexDB, "")
	if err != nil {
		return nil, errors.Wrap(err, "fail to query from InvertedIndexDB")
	}
	return tokenDict, nil
}

// InsertTokenSentence
func InsertTokenSentence(token string, docs string, positions string) error {
	keys := []string{TokenColumn, DocColumn, PosColumn}
	values := []interface{}{token, docs, positions}
	err := tool.Insert(keys, values, InvertedIndexDB)
	if err != nil {
		return errors.Wrap(err, "fail to insert into DB")
	}
	fmt.Printf("insert %s, %s, %s successfully\n", token, docs, positions)
	return nil
}

// list []string 格式:
// list[0]: doc_001, doc_002, doc_003...
// list[1]: doc_001:1,2,3,4;doc_002:4,23,122;doc_003:...

// InsertOrUpdateTokensInDB 插入或更新DB中的token信息
func InsertOrUpdateTokensInDB(freshMap map[string][]string) error {
	if len(freshMap) < 1 {
		return errors.New("invalid map size")
	}
	tokens := make([]string, 0)
	for token, _ := range freshMap {
		s := "'" + token + "'"
		tokens = append(tokens, s)
	}
	extraInfo := " WHERE c_token IN (" + strings.Join(tokens, ",") + ")"
	oldMap, err := tool.Query("*", InvertedIndexDB, extraInfo)
	if err != nil {
		return errors.Wrap(err, "fail to query DB in Func:InsertOrUpdateTokensInDB()")
	}

	for token, values := range freshMap {
		if _, ok := oldMap[token]; ok {
			if len(oldMap[token]) < 2 || len(values) < 2 {
				return errors.Wrap(err, "invalid length of map elements")
			}
			freshMap[token][0] = oldMap[token][0] + "," + values[0]
		}
	}

	keys := []string{TokenColumn, DocColumn, PosColumn}
	err = tool.InsertOrUpdate(keys, freshMap, InvertedIndexDB)
	if err != nil {
		return errors.Wrap(err, "fail to call func:tool.InsertOrUpdate()")
	}
	fmt.Println("Successfully insert or update token info in db.")
	return nil
}
