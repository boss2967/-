package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"log"
)

/*    交易*/

//1. 	创建交易结构
type Transaction struct {
	TXID      []byte     //交易ID
	TXInputs  []TXInput  //交易输入数组,可能是多个
	TXOutputs []TXOutput //交易输出数组，可能是多个

}

//输入input
type TXInput struct {
	TXID  []byte //引用的交易id
	Index int64  // 引用的output的索引值
	Sig   string //签名,解锁脚本，我们用地址来模拟
}

//输出output
type TXOutput struct {
	Value      float64 //转账金额
	PubKeyHash string  // 锁定脚本。地址模拟
}

// 设置交易ID
func (tx *Transaction) SetHash() {
	var buffer bytes.Buffer //bytes类型缓冲区
	//01.	生成编码器
	encoder := gob.NewEncoder(&buffer)
	//02.	编码
	err := encoder.Encode(&tx)
	// nil
	if err != nil {
		log.Panic("设置id出错：", err)
	}
	//buffer.bytes
	data := buffer.Bytes()
	//hash 256
	hash := sha256.Sum256(data)
	//数组转切片
	tx.TXID = hash[:]
}

const reward = 12.5

//实现一个函数，判断当前的交易是否为挖矿交易,判断是否挖矿交易
func (tx *Transaction) IsCoinbase() bool {
	//1.	交易input只有一个
	if len(tx.TXInputs) == 1 {
		input := tx.TXInputs[0]
		//2.	交易id为空
		//3.	交易的index为空 -1
		if !bytes.Equal(input.TXID, []byte{}) || input.Index != -1 {
			return false
		}
	}
	return true
}

//1. 	创建交易结构 挖矿交易
//2.  	提供创建交易方法
func NewCoinbaseTX(address string, data string) *Transaction {
	// 1.生成  两个数组
	//挖矿加一的特点：1 只有一个input 2. 无需 引用交易ID ,3无需引用index
	//矿工因为挖矿是无需指定签名，所以siz字段会有矿工来自由填写
	input := TXInput{[]byte{}, -1, data}
	//
	output := TXOutput{reward, address}
	//对于挖矿交易来说，只有一个input和output
	tx := Transaction{[]byte{}, []TXInput{input}, []TXOutput{output}}
	//计算hash
	tx.SetHash()
	return &tx
}

/*创建普通交易
 	1.找到自己 最合理的utxo集合	  map[string][]uint64
	2.将这些UTXO转成input
	3.创建output,
	4.如果有零钱，找零。
*/
func NewTransaction(from, to string, amout float64, bc *BlockChain) *Transaction {
	// 1.找到最合理UTXO集合，map[stirng] uint64
	utxos, resValue := bc.FindNeedUTXOs(from, amout)
	//2.	判断
	if resValue < amout {
		fmt.Printf("余额不足！，交易失败！")
		return nil
	}
	//声明数组
	var inputs []TXInput
	var outputs []TXOutput
	//2.	创建交易的输入，将这些UTXO卓一转成inputs
	for id, indexArray := range utxos {
		for _, i := range indexArray {
			input := TXInput{[]byte(id), int64(i), from}
			inputs = append(inputs, input)
		}
	}
	//创建交易输出
	output := TXOutput{amout, to}
	outputs = append(outputs, output)
	//找零
	if resValue > amout {
		outputs = append(outputs, TXOutput{resValue - amout, from})
	}
	tx := Transaction{[]byte{}, inputs, outputs}
	tx.SetHash()
	return &tx
}
//4. 	根据交易调整程序
