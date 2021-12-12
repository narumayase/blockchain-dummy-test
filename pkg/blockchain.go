package pkg

const BlockchainDifficulty = 2

type Blockchain struct {
	Chain []*Block
}

func NewBlockchain() *Blockchain {
	genesisBlock := NewBlock("0", "{\"isGenesis\":true}")
	return &Blockchain{Chain: []*Block{genesisBlock}}
}

func (b *Blockchain) AddBlock(data string) {
	lastBlock := b.Chain[len(b.Chain) -1 ]
	newBlock := NewBlock(lastBlock.Hash, data)
	newBlock.Mine(BlockchainDifficulty)
	b.Chain = append(b.Chain, newBlock)
}

func (b *Blockchain) IsValid() bool {
	for i := 1; i < len(b.Chain); i++ {
		currentBlock := b.Chain[i]
		previousBlock := b.Chain[i-1]

		if currentBlock.Hash != currentBlock.CalculateHash() {
			return false
		}
		if currentBlock.PreviousHash != previousBlock.Hash {
			return false
		}
	}
	return true
}
