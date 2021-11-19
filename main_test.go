package main

import (
	"encoding/hex"
	"fmt"
	"goseeker/handlers"
	"goseeker/tool"
	"testing"
)

func Test_Main(t *testing.T) {
	fmt.Println("start loading...")

	for i := 1; i <= 2; i++ {
		filename := fmt.Sprintf("/Users/lisbon/Desktop/seeker/docs/AA/doc_%02d", i)
		fmt.Printf("trying to open %s\n", filename)
		content, err := tool.LoadFile(filename)
		if err != nil {
			panic(err)
		}
		//fmt.Println(content)

		docID := fmt.Sprintf("doc_%03d", i)
		tokenDict := handlers.TextToTokens(content, docID)
		err = handlers.InsertOrUpdateTokensInDB(tokenDict)
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

func Test_LoadWiki(t *testing.T) {
	counter := 1
	for i := 0; i < 100; i++ {
		if i%10 == 0 {
			fmt.Printf("handling file No.%d\n", i)
		}
		filename := fmt.Sprintf("/Users/lisbon/Desktop/seeker/wiki_zh/AA/wiki_%02d", i)
		freshCounter, err := tool.LoadWiki(filename, counter)
		if err != nil {
			panic(err)
		}
		counter = freshCounter
	}
}

func Test_Small(t *testing.T) {
	a := "五年计划"
	s := fmt.Sprintf("%x", []byte(a))
	fmt.Println(s)
	ori, _ := hex.DecodeString(s)
	fmt.Printf("%s", ori)
}
