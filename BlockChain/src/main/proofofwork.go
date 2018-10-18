package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

//01.	定义一个工作量证明的结构
type ProofOfWork struct {
	block  *Block   //blcok
	target *big.Int //非常大的数据
}

/****
* 	**Block   NewProofOWork 创建一个工作量证明
* 	parmas()
*	return( *ProofOfWork)
*
 */
//02.	创建一个工作量证明
func NewProofOWork(block *Block) *ProofOfWork {
	//生成一个结构体
	pow := ProofOfWork{
		block: block,
	}
	//目标值
	//指定的难度值，现在是string类型需要转换
	targetStr := "0000100000000000000000000000000000000000000000000000000000000000"
	//引入的辅助变量，目的是将上面的难度值转换成 big.int
	tmpInt := big.Int{}
	//将难度值赋值给big.int 指定16进制的格式
	tmpInt.SetString(targetStr, 16) //转换成16进制
	//区块本体赋值
	pow.target = &tmpInt
	//返回
	return &pow
}

/****
* 	**ProofOfWork   Run 函数 ，提供不断计算哈希函数
* 	parmas()
*	return( []byte ，uint64)
*
 */
//03.	run 函数 ，提供不断计算哈希函数
func (pow *ProofOfWork) Run() ([]byte, uint64) {
	//01.	拼装数据： 区块数据，不但变化的随机数
	//02.	做哈希运算
	//03.	做数据比较  pow中target 如果找到退出，如果没有找到随机数加一
	var nonce uint64   //随机数
	block := pow.block // 区块
	var hash [32]byte  // hash
	//开始挖矿
	fmt.Println("开始挖矿...")
	for {
		//fmt.Println("挖矿次数", nonce)
		//01.拼装数据
		tmp := [][]byte{
			Uint64ToByte(block.Version),
			block.PrevHash,
			block.MerkelRoot,
			Uint64ToByte(block.TimeStamp),
			Uint64ToByte(block.Difficulty),
			Uint64ToByte(nonce), //修改的随机值
			//block.Data,
		}
		//将二维的切牌你数组连接起来，返回一个一维的切片
		blockInfo := bytes.Join(tmp, []byte{})
		//02.	哈希运算
		hash = sha256.Sum256(blockInfo)
		//03.	比较
		tmpInt := big.Int{}
		//将我们的到的hash数组转换成一个big.Int
		tmpInt.SetBytes(hash[:])
		//比较当前的哈希和目标的哈希值，如果当前的哈希值小于目标的哈希值，就说明刚找到了满足，否则就继续找

		if tmpInt.Cmp(pow.target) == -1 {
			fmt.Println("-----", nonce)
			fmt.Printf("---------挖矿成功 %x\n,%d ,\n", hash, nonce)
			return hash[:], nonce
		} else {
			nonce++
		}
	}
	//直接返回对比成功的哈希 04.提供一个校验函数
}
