package main

import (
	"BlockChain/src/bolt/bolt-master/bolt-master"
	"log"
)

/****
* 	BlockChainIterator   结构体sql
* 	parmas()
*	return()
*
 */
type BlockChainIterator struct {
	db *bolt.DB
	//游标
	currentHashPointer []byte
}

/****
* 	*BlockChain   NewIterator 游标卡尺
* 	parmas()
*	return(*BlockChainIterator)
*
 */
//游标卡尺
func (bc *BlockChain) NewIterator() *BlockChainIterator {
	//返回一个 sql结构体 参数是 db, 和上一个hash
	return &BlockChainIterator{
		bc.db,
		bc.tail,
	}
}

/****
* 	*BlockChainIterator   Next 游标卡尺
* 	parmas()
*	return(*Block)
*
 */
//迭代器是区块链，Next属于迭代器
func (it *BlockChainIterator) Next() *Block {
	//1.返回当前的区块
	//2. 指针前移
	var block Block
	it.db.View(func(tx *bolt.Tx) error {
		//
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil {
			log.Panic("迭代器遍历bucket不应该为空，请检查")
		}
		//获取上一个hash
		blockTmp := bucket.Get(it.currentHashPointer)
		//解码动作
		block := Deserialize(blockTmp)
		//游标移动
		it.currentHashPointer = block.PrevHash //游标哈希左移动
		//返回
		return nil
	})
	//返回
	return &block
}
