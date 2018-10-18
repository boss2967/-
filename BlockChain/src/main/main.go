package main

import "fmt"

func main() {
	fmt.Println("开始")
	bc := NewBlockChain("打开哦啊家平")
	cli := CLI{bc}
	cli.Run()
	//bc := NewBlockChain()
	//	//bc.AddBlock("张三向李四转了50个比特币")
	//	//bc.AddBlock("张三向李四转了50234个比特币")
	//	//
	//	////创建迭代器
	//	//it := bc.NewIterator()
	//	////调用迭代器 ，返回区块数据
	//	//for {
	//	//	block := it.Next()
	//	//	fmt.Printf("%x\n", block.Hash)
	//	//	if len(block.PrevHash) == 0 {
	//	//		break
	//	//	}
	//	//}
	//block := NewBlock("老师给班长转了一枚比特币", []byte{})
}
func main2() {
	//fmt.Println("----------")
	//db, err := bolt.Open(":test.db", 0600, nil)
	//if err != nil {
	//	log.Panic("打开数据库失败")
	//}
	//db.Update(func(tx *bolt.Tx) error {
	//	bucket := tx.Bucket([]byte("b1"))
	//	if bucket == nil {
	//		bucket, err = tx.CreateBucket([]byte("b1"))
	//		if err != nil {
	//			log.Panic("创建bucket（b1）失败")
	//		}
	//	}
	//	bucket.Put([]byte("1111"), []byte("hello"))
	//	bucket.Put([]byte("22222"), []byte("word"))
	//	return nil
	//})
	////读取数据
	//db.View(func(tx *bolt.Tx) error {
	//	//01.	找到抽屉，没有的话 退出
	//	//02.	直接读取数据
	//	bucket := tx.Bucket([]byte("b1"))
	//	if bucket == nil {
	//
	//		log.Panic("bucket b1不能为空")
	//	}
	//
	//	v1 := bucket.Get([]byte("1111"))
	//	v2 := bucket.Get([]byte("22222"))
	//	fmt.Printf("---%s", v1)
	//	fmt.Printf("---%s", v2)
	//	return nil
	//})
}
