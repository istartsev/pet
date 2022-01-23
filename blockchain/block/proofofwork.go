package block

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
	"strconv"
)

const targetBits = 24
const maxNonce = math.MaxInt64


type ProofOfWork struct {
	block  *Block
	target *big.Int
}

func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))

	pow := &ProofOfWork{b, target}

	return pow
}

func IntToHex(i int64) []byte {
	return []byte(strconv.FormatInt(i, 10))
}

func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Data,
			IntToHex(pow.block.Timestamp),
			IntToHex(int64(targetBits)),
			IntToHex(int64(nonce)),
		},
		[]byte{},
	)

	return data
}

type POWRes struct {
	Nonce int
	Hash []byte
}

func (pow *ProofOfWork) workThread(requests <-chan int, done chan<- POWRes) {
	for req := range requests {
		go func(nonce int) {
			var hashInt big.Int
			data := pow.prepareData(nonce)
			hash := sha256.Sum256(data)
			//fmt.Printf("\r%x", hash)
			hashInt.SetBytes(hash[:])

			if hashInt.Cmp(pow.target) == -1 {
				done <- POWRes{Nonce: nonce, Hash: hash[:]}
			}
		}(req)
	}
}

func (pow *ProofOfWork) Run() (int, []byte) {
	var hash [32]byte
	nonce := 0

	fmt.Printf("Mining the block containing \"%s\"\n", pow.block.Data)
	done := make(chan POWRes)
	requestChan := make(chan int, 10)
	go pow.workThread(requestChan, done)

	for nonce < maxNonce {
		select {
		case res :=<- done:
			close(requestChan)
			return res.Nonce, res.Hash
		default:
			requestChan <- nonce
		}
		nonce++
	}
	fmt.Print("\n\n")

	return nonce, hash[:]
}

func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int

	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	isValid := hashInt.Cmp(pow.target) == -1

	return isValid
}