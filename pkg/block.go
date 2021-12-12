package pkg

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type Block struct {
	Data         string
	Hash         string
	PreviousHash string
	TimeStamp    time.Time
	ProofOfWork  int
}

func NewBlock(prevHash, data string) *Block {
	b := &Block{
		Data:         data,
		PreviousHash: prevHash,
		TimeStamp:    time.Now(),
		ProofOfWork:  0,
	}
	b.Hash = b.CalculateHash()
	return b
}

func (b *Block) CalculateHash() string {
	return fmt.Sprintf("%x",
		sha256.Sum256([]byte(
			fmt.Sprintf("%s%s%s%d", b.PreviousHash, b.Data, b.TimeStamp.String(), b.ProofOfWork))))
}

func (b *Block) Mine(difficulty int) {
	for !strings.HasPrefix(b.Hash, b.zeros(difficulty)) {
		b.ProofOfWork++
		b.Hash = b.CalculateHash()
	}
}

func (b *Block) zeros(qty int) string {
	return fmt.Sprintf("%0"+fmt.Sprintf("%d", qty)+"d", 0)
}

func (b *Block) String() string {
	s, err := json.Marshal(b)
	if err != nil {
		return ""
	}
	return string(s)
}
