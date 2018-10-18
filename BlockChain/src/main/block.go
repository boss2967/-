package main

import (
	"bytes"
	//"crypto/sha256"	  //"crypto/sha256"
	"crypto/sha256"
	"encoding/binary"
	"encoding/gob"
	"fmt"
	"log"
	"time"
)

//定义结构
type Block struct {
	//Version    uint64 //版本号
	//MerkelRoot []byte //merkel根
	//PrevHash   []byte //前区块哈希
	//Hash       []byte //当前区块哈希，正常中是没有当前区块的
	//Data       []byte //数据
	//TimeStamp  uint64 //时间戳
	//Difficulty uint64 //难度值
	//Nonce      uint64 //随机数
	//1.版本号
	Version uint64 //版本号
	//2. 前区块哈希
	PrevHash []byte
	//3. Merkel根（梅克尔根，这就是一个哈希值，我们先不管，我们后面v4再介绍）
	MerkelRoot []byte
	//4. 时间戳
	TimeStamp uint64
	//5. 难度值
	Difficulty uint64
	//6. 随机数，也就是挖矿要找的数据
	Nonce uint64
	//a. 当前区块哈希,正常比特币区块中没有当前区块的哈希，我们为了是方便做了简化！
	Hash []byte
	//b. 数据
	Data []byte
	///正式的交易数据组 区块体
	Transactions []*Transaction
}

/****
* 	Uint64ToByte 实现一个辅助函数， 功能将是将uint64转换成[]byte
* 	parmas(uint64)
*r	eturn([]byte)
*
 */
func Uint64ToByte(num uint64) []byte {
	var buffer bytes.Buffer
	err := binary.Write(&buffer, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buffer.Bytes()
}

/****
* 	NewBlock 实现创建一个区块
* 	parmas(string,[]byte)
*	return( *Block)
*
 */
//02.	创建一个区块
func NewBlock(txs []*Transaction, prevBlockHash []byte) *Block {
	//一个区块信息，和参数构成一个块
	var block Block
	//封装对象
	block = Block{
		Version:      00,
		PrevHash:     prevBlockHash,
		MerkelRoot:   []byte{},
		TimeStamp:    uint64(time.Now().Unix()),
		Nonce:        0,
		Difficulty:   0,
		Hash:         []byte{},
		Transactions: txs,
	}
	//默克尔树
	block.MerkelRoot = block.MarkMerklBoot()
	//创建一个pow对象
	pow := NewProofOWork(&block) //工作量证明
	hash, nonce := pow.Run()     //查找随机数，不停的进行哈希运算
	//根据挖矿结果对区块数据进行补充
	block.Hash = hash   //设置本区块hash
	block.Nonce = nonce //设置随机数
	//返回区块，注意是取地址
	return &block
}

////生成哈希
//func (block *Block) SetHash() {
//	//拼装数据
//	tmp := [][]byte{
//		Uint64ToByte(block.Version),
//		block.PrevHash,
//		block.MerkelRoot,
//		Uint64ToByte(block.TimeStamp),
//		Uint64ToByte(block.Difficulty),
//		Uint64ToByte(block.Nonce),
//		block.Data,
//	}
//	//将二维的切片数组连接起来，返回一个一组的片片
//	blockInfo := bytes.Join(tmp, []byte{})
//	//blockInfo := append(block.PrevHash, block.Data...)
//	// sha256
//	hash := sha256.Sum256(blockInfo)
//	//
//	block.Hash = hash[:]
//}
/****
* 	*Block  方法 Serialize
* 	parmas(nil)
*	return([]byte)
*
 */
func (block *Block) Serialize() []byte {
	var buffer bytes.Buffer //bytes类型缓冲区
	//01.	生成编码器
	encoder := gob.NewEncoder(&buffer)
	//02.	编码
	err := encoder.Encode(&block)
	//	// nil
	if err != nil {
		log.Panic("编码出错")
	}
	//返回
	return buffer.Bytes()
}

//反序列化
func Deserialize(data []byte) Block {
	//定义 block变量
	fmt.Println("-------bbbbbb", []byte{})
	var block Block
	//01.	生成解码器
	decoder := gob.NewDecoder(bytes.NewReader(data))
	//02.	解码
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic("解码出错")
	}
	return block
}
func (block *Block) MarkMerklBoot() []byte {
	return []byte{}

}

//模拟没课耳根， 支队交易的数据做简单的拼接，而不做二叉树处理
func (block *Block) MarkerlRoot() []byte {
	var info []byte
	//将哈希值拼接起来，
	for _, tx := range block.Transactions {
		info = append(info, tx.TXID...)
		//finalinfo = [][]byte{tx.TXID}
	}
	hash := sha256.Sum256(info)
	return hash[:]
}
