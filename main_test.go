package main

import (
	"fmt"
	"goseeker/handlers"
	"goseeker/tool"
	"testing"
)

func Test_Main(t *testing.T) {
	fmt.Println("start loading...")

	for i := 1; i < 2; i++ {
		filename := fmt.Sprintf("corpus/wiki_%02d", i)
		fmt.Printf("trying to open %s\n", filename)
		content, err := tool.LoadFile(filename)
		if err != nil {
			panic(err)
		}
		//fmt.Println(content)

		docID := fmt.Sprintf("doc_%03d", i)
		tokenDict := handlers.TextToTokens(content, docID)
		keys := []string{handlers.TokenColumn, handlers.DocColumn, handlers.PosColumn}
		err = tool.InsertOrUpdate(keys, tokenDict, handlers.InvertedIndexDB)
		if err != nil {
			panic(err)
		}
	}

	// 查询语句
	//_, err = tool.Query("*", InvertedIndexDB, "")
	//if err != nil {
	//	panic(err)
	//}

	// 插入语句
	//err := handlers.InsertTokenSentence("必胜", "doc_003", "1,2,3,4")
	//if err != nil {
	//	panic(err)
	//}
	//
	//_ = &map[string][]string{
	//	"必胜": {"doc_004,doc_005", "5,29,30"},
	//}
}

func Test_Func(t *testing.T) {
	filename := fmt.Sprintf("corpus/wiki_%02d", 1)
	fmt.Printf("trying to open %s\n", filename)
	content, err := tool.LoadFile(filename)
	if err != nil {
		panic(err)
	}
	//fmt.Println(content)

	docID := fmt.Sprintf("doc_%03d", 1)
	tokenDict := handlers.TextToTokens(content, docID)
	fmt.Printf("len: %d", len(tokenDict))
}
