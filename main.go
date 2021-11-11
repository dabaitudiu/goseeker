package main

import (
	"fmt"
	"goseeker/handlers"
	"goseeker/tool"
	"unsafe"
)

func main() {
	fmt.Println("start loading...")
	content, err := tool.LoadArticle("corpus/wiki_00")
	if err != nil {
		panic(err)
	}
	//fmt.Println(content)

	tokenDict := make(map[string]string)

	handlers.ArticleToTokens(tokenDict, content, "0")
	fmt.Println(unsafe.Sizeof(tokenDict))

}
