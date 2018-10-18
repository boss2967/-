package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

type Person struct {
	Name string
	age  uint
}

func main() {
	//1. 定义一个结构Per'son
	var xiaoming Person
	//2.	buffer,辅助变量
	var buffer bytes.Buffer
	//编码器， 解码器

	//编码器
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(&xiaoming)
	if err != nil {
		log.Panic("编码出错")
	}
	//
	fmt.Println("----------", buffer.Bytes())
	//01.	解码器
	decoder := gob.NewDecoder(bytes.NewReader(buffer.Bytes()))
	//02.	解码
	var daming Person
	err = decoder.Decode(&daming)
	if err != nil {
		log.Panic("解码出错")
	}
}
