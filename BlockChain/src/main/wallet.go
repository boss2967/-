package main

import (
	"BlockChain/src/base58"
	"BlockChain/src/ripemd160"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"log"
)

//这里的钱包是一个结构，每个钱包保存了公钥，私钥对
type Wallet struct {
	//私钥
	Private *ecdsa.PrivateKey
	//这里不存公钥,存的是拼接字符串，不存原始的公钥，而是存xy拼接的字符串，在校验端重新拆分

	PubKey []byte
}

//创建钱包
func NewWallet() *Wallet {
	//创建曲线
	curve := elliptic.P256()
	//生成私钥
	privateKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	//nil
	if err != nil {
		log.Panic()
	}
	//生成公钥
	pubKeyOrig := privateKey.PublicKey
	//拼接 X,Y
	pubKey := append(pubKeyOrig.X.Bytes(), pubKeyOrig.Y.Bytes()...)

	//返回
	return &Wallet{Private: privateKey, PubKey: pubKey}
}

//生成地址
func (w *Wallet) NewAddress() string {
	//拿到公钥
	pubKey := w.PubKey
	//160
	rip160HashValue := HashPubKey(pubKey)
	//hash160
	version := byte(00)
	//拼接version
	payload := append([]byte{version}, rip160HashValue...)
	//checksum
	checkCode := CheckSum(payload)
	//25字节数据
	payload = append(payload, checkCode...)
	//go语言有一个库，叫做btcd,这个语言是go实现的比特币全节点源码
	address := base58.Encode(payload)
	//
	return address
}

//Hash256
func HashPubKey(data []byte) []byte {
	hash := sha256.Sum256(data)
	//理解为编码器
	rip160hasher := ripemd160.New()
	_, err := rip160hasher.Write(hash[:])
	if err != nil {
		log.Panic(err)
	}
	//返回rip160的哈希结果
	rip160HashValue := rip160hasher.Sum(nil)
	return rip160HashValue
}

//两次哈希值
func CheckSum(data []byte) []byte {
	hash1 := sha256.Sum256(data)
	hash2 := sha256.Sum256(hash1[:])

	//前四字节校验码
	checkCode := hash2[:4]
	//返回
	return checkCode

}
