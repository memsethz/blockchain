package main

import "blockchain/src"

func main() {
	bc := blockchain.NewBlockchain()
	//defer bc.db.Close()

	cli := blockchain.NewCli(bc)
	cli.Run()
}
