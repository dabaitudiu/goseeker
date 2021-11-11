package handlers

import (
	"log"
	"regexp"

	"github.com/yanyiwu/gojieba"
)

// NormalizeArticle
func NormalizeArticle(articleStr string) string {
	chineseReg := regexp.MustCompile("[^\u4e00-\u9fa5]")
	return chineseReg.ReplaceAllString(articleStr, "")
}

// ArticleToTokens
func ArticleToTokens(tokenDict map[string]string, articleStr string, articleID string) []string {

	articleStr = NormalizeArticle(articleStr)

	x := gojieba.NewJieba()
	defer x.Free()

	words := x.CutAll(articleStr)
	for _, word := range words {
		tokenDict[word] = articleID
	}

	log.Printf("tokenize finished, total tokens: %d\n", len(words))

	return nil
}
