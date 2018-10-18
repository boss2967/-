package main

import (
	"BlockChain/src/bolt/bolt-master/bolt-master"

	"fmt"
	"log"
)

//引入区块链
//type BlockChain struct {
//	blocks []*Block
//}

//v3, 使用数据库代替数据
type BlockChain struct {
	//blocks []*Block
	db   *bolt.DB
	tail []byte //存储最后一个区块的哈希

}

const blockChainDb = "blockChain.db"
const blockBucket = "blockBucket"

/****
* 	NewBlockChain    定义一个区块链
* 	parmas(nil)
*	return([*BlockChain)
*
 */
// 05.	定义一个区块链
func NewBlockChain(address string) *BlockChain {
	//01.	创建一个创世块，并且添加到区块链中
	var lastHash []byte //从数据库中读取出来的变量
	//fmt.Println("开始2455")
	//02.	打开数据库
	db, err := bolt.Open(blockChainDb, 0600, nil)
	//nil
	if err != nil {
		log.Panic("打开数据库失败：", err)
	}
	//03.	将要操作数据库
	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil {
			//fmt.Println("开始4")
			//如果没有抽屉，我们就创建
			bucket, err = tx.CreateBucket([]byte(blockBucket))
			if err != nil {
				//fmt.Println("开始5")
				//如果没有抽屉，创建
				log.Panic("创建bucket（b1）失败")
			}
			//创建一个创世块，并且作为第一个区块添加到区块中
			//fmt.Println("开始6")
			gensisblock := GenesisBlock(address)
			//写数据,强转写入数据
			fmt.Printf("genesisBlock :%s\n", gensisblock)
			bucket.Put(gensisblock.Hash, gensisblock.Serialize())
			//写最后一个hash
			bucket.Put([]byte("LastHashKey"), gensisblock.Hash)
			//重写最后一个区块 hash
			lastHash = gensisblock.Hash

			//fmt.Println("开始111111111111111111111111111111111111111111111111111111111111111111111111111")
		} else {
			//重写最后一个区块hash
			lastHash = bucket.Get([]byte("LastHashKey"))
		}
		//返回
		return nil
	})
	//返回一个结构体，最后一个hash和 db
	return &BlockChain{db, lastHash}
}

/****
* 	GenesisBlock    生成创世区块
* 	parmas(nil)
*	return(*Block)
*
 */
//创世区块
func GenesisBlock(address string) *Block {
	coinbase := NewCoinbaseTX(address, "大家好")
	return NewBlock([]*Transaction{coinbase}, []byte{})
	//coinbase:= NewCoinbaseTX(address,gensisInfo)
}

/****
* 	BlockChain   方法 AddBlock  添加一个区块
* 	parmas(nil)
*	return(string)
*
 */
//添加区块
func (bc *BlockChain) AddBlock(txs []*Transaction) {
	//获取哈希
	db := bc.db         //变量承接收
	lastHash := bc.tail //最后一个区块的hash
	//操作数据库
	db.Update(func(tx *bolt.Tx) error {
		//完成数据的添加
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil {
			log.Panic("bucket不应该为空，请检查")
		}
		//生成最新的一个区块，要有内容， 和上一个区块的hash
		block := NewBlock(txs, lastHash)
		//赋值,存到数据库中， hash> >>>>>block块序列化的数据
		bucket.Put(block.Hash, block.Serialize())
		//修改数据库中最新的指向上一个hash
		bucket.Put([]byte("lastHashKey"), block.Hash)
		//修改内存中的最新区块，这个干什么用呢？他是为了最新区块追加上一个hash而准备中间量
		lastHash = block.Hash
		//更新bc 结构体
		bc.tail = lastHash
		//返回
		return nil
	})
	//创建新的区块
	//添加到db中
	// /lastBlock := bc.blocks[len(bc.blocks)-1]
	//prevHash := lastBlock.Hash
	////创建新的区块
	//block := NewBlock(data, prevHash)
	//bc.blocks = append(bc.blocks, block)

}

// v4 找到指定地址的所有的UTXO
func (bc *BlockChain) FindUTXOs(address string) []TXOutput {
	var UTXO []TXOutput
	//1.遍历区块
	//2.遍历交易
	//3.遍历output,找到属于自己的 utxo
	//先检查自己是否消耗过
	//4.遍历input 找到自己花费过的 utxo
	spentOutputs := make(map[string][]int64)
	//创建迭代器
	it := bc.NewIterator()
	for {
		//遍历区块
		block := it.Next()
		//遍历交易
		for _, tx := range block.Transactions {
			fmt.Printf("current tix:%x\n", tx.TXID)
			//
			//3.遍历output,找到属于自己的 utxo
		OUTPUT:
			//遍历output 找到和自己相关的utxo
			for i, output := range tx.TXOutputs {
				fmt.Printf("current index:%d\n", i)
				if spentOutputs[string(tx.TXID)] != nil {
					for _, j := range spentOutputs[string(tx.TXID)] {
						if int64(i) == j {
							continue OUTPUT
						}
					}
				}
				if output.PubKeyHash == address {
					UTXO = append(UTXO, output)
					fmt.Printf("%f\n", UTXO[0].Value)
				} else {
					fmt.Printf("-------------------------------3462")
				}
			}
			//如果当前交易是挖矿交易的话，那么不做遍历，直接跳过
			if !tx.IsCoinbase() {
				//遍历input，找到自己花费过的utxo的集合()
				for _, input := range tx.TXInputs {
					//
					if input.Sig == address {
						spentOutputs[string(input.TXID)] = append(spentOutputs[string(input.TXID)], input.Index)
					}
				}
			} else {
				fmt.Printf("这是coinbase,不做input遍历")
			}
		}
		if len(block.PrevHash) == 0 {
			fmt.Printf("区块遍历完成退出!")
			break
		}

		//这个output和我们目标的地址相通，蛮子条件，加到返回UTXO数组中

		//在这做一个过滤， 就是output中已经失效的，就是已经处理过的
		//如果相同，则跳过

		//如果当前交易id存在input map中，表示消耗过
		//

		//判断是否是挖矿交易 1.交易input只有一个，交易id为空， 交易索引为-1
		//4.遍历input， 把自己花费过的utxo查询出来

	}
	return UTXO
}

//
func (bc *BlockChain) FindNeedUTXOs(from string, amout float64) (map[string][]uint64, float64) {
	//找到合理的utxos集合
	var utxos map[string][]uint64
	var calc float64
	//TODO
	return utxos, calc
}
