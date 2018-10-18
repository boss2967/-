package main

import (
	"fmt"
	"os"
)

//用来接收文件
type CLI struct {
	bc *BlockChain
}

const Usage = `i am  code
 getBlance - address ADDESS
`

func (cli *CLI) Run() {
	//01.	得到所有的命令
	args := os.Args
	if len(args) < 2 {
		fmt.Printf(Usage)
		return
	}
	//02.	执行命令C
	cmd := args[1]
	switch cmd {
	case "addBlock":
		//3. 执行相应动作
		fmt.Printf("添加区块\n")

		//确保命令有效
		if len(args) == 4 && args[2] == "--data" {
			//获取命令的数据
			//a. 获取数据
			//data := args[4]
			//b. 使用bc添加区块AddBlock
			//cli.bc.AddBlock(data)
		} else {
			fmt.Printf("添加区块参数使用不当，请检查")
			fmt.Printf(Usage)
		}
	case "printChain":
		fmt.Printf("正向打印区块\n")
		cli.PrinBlockChain()
	case "printChainR":
		fmt.Printf("反向打印区块\n")
		cli.PrinBlockChainReverse()
	case "getBalance":
		fmt.Printf("获取余额\n")
		if len(args) == 4 && args[2] == "--address" {
			address := args[3]
			cli.GetBalance(address)
		}
	default:
		fmt.Printf("无效的命令，请检查!\n")
		fmt.Printf(Usage)
	}
}
