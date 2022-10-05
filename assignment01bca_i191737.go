package main

import (
	"crypto/sha256"
	"fmt"
	"strconv"
)



func CalculateHash(stringToHash string) string {
	sum := sha256.Sum256([]byte(stringToHash))

	// fmt.Printf("%x\n", sum) //hexadecimal
	return fmt.Sprintf("%x\n", sum)
}

type Block struct {
	Transaction string
	Nonce       int

	PrevHash string
	Hash     string
}

type Blockchain struct {
	blocks []*Block
}

func (bc *Blockchain) NewBlock(Transaction string, Nonce int, previousHash string) Block {
	nonce := strconv.Itoa(Nonce)
	block := Block{Transaction, Nonce, previousHash, CalculateHash(Transaction + nonce + previousHash)}
	bc.blocks = append(bc.blocks, &block)

	return block
}

func (block *Block) ChangeBlock(Mtransaction string) {
	block.Transaction = Mtransaction
	block.Hash = CalculateHash(Mtransaction + strconv.Itoa(block.Nonce) + block.PrevHash)
}

func (bc *Blockchain) VerifyChain() {
	for i := 0; i < len(bc.blocks); i++ {
		if i == 0 {
			continue
		}
		if bc.blocks[i].PrevHash != bc.blocks[i-1].Hash {
			fmt.Println("Blockchain is not valid")
			fmt.Println("Issue Occured at blocks: ", i, "and", i+1)
			return
		}
	}

	fmt.Println("Blockchain is valid")

}

func (bc *Blockchain) ListBlocks() {
	for i := 0; i < len(bc.blocks); i++ {

		fmt.Printf("*******     Block %d     *******\n, Transaction: %s\n, Nonce: %d\n, PrevHash: %s, Hash: %s\n\n", i+1, bc.blocks[i].Transaction, bc.blocks[i].Nonce, bc.blocks[i].PrevHash, bc.blocks[i].Hash)
	}

}

func main() {
	Blockchain := Blockchain{}

	//Block 1
	Blockchain.NewBlock("Transaction 1", 1, "0")
	//Block 2
	Blockchain.NewBlock("Transaction 2", 2, Blockchain.blocks[0].Hash)
	//Block 3
	Blockchain.NewBlock("Transaction 3", 3, Blockchain.blocks[1].Hash)
	//Block 4
	Blockchain.NewBlock("Transaction 4", 4, Blockchain.blocks[2].Hash)
	//Block 5
	Blockchain.NewBlock("Transaction 5", 5, Blockchain.blocks[3].Hash)

	Blockchain.ListBlocks()
	Blockchain.VerifyChain()

	Blockchain.blocks[2].ChangeBlock("Transaction 3.1")
	Blockchain.ListBlocks()

	Blockchain.VerifyChain()
}
