package main

import (
	"fmt"
)

func main() {
	tokenList, scanner := Parse("int a = 123")
	fmt.Printf("%+v\n", scanner)
	if tokenList == nil {
		fmt.Println("解析失败")
	} else {
		for _, token := range tokenList {
			fmt.Printf("位置:%d\t类型:%s\t文本:%s\n", token.Offset, token.TokenType, token.TokenText)
		}
	}
	// 执行结果
	// 位置:0  类型:Keyword    文本:int
	// 位置:4  类型:ID 文本:a
	// 位置:7  类型:Punctuation        文本:=
	// 位置:8  类型:IntLiteral 文本:123
}

func Parse(source string) ([]Token, *Scanner) {
	s := NewScanner(source)
	s.StackPush()
	parseAssign := AndNoWhiteToken(
		OneOfToken(ParseKeyword, "int"),
		ParseID,
		OneOfToken(ParsePunctuation, "="),
		ParseIntLiteral,
	)
	parser := Try(parseAssign)
	tokenList := parser(s)
	if tokenList == nil {
		s.StackPop(true)
	} else {
		s.StackPop(false)
	}
	return tokenList, s
}
