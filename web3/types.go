// Copyright (c) 2016, Alan Chen
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice,
//    this list of conditions and the following disclaimer.
//
// 2. Redistributions in binary form must reproduce the above copyright notice,
//    this list of conditions and the following disclaimer in the documentation
//    and/or other materials provided with the distribution.
//
// 3. Neither the name of the copyright holder nor the names of its contributors
//    may be used to endorse or promote products derived from this software
//    without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
// ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE
// LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
// CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF
// SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
// INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN
// CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
// ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
// POSSIBILITY OF SUCH DAMAGE.

package web3

import (
	"math/big"
)

const (
	hashLength    = 32
	addressLength = 20
)

// Hash ...
type Hash [hashLength]byte

// Address ...
type Address [addressLength]byte

// BlockIndicator ...
type BlockIndicator struct {
	Block interface{}
}

// TransactionRequest ...
type TransactionRequest struct {
	From     Address  `json:"from"`
	To       Address  `json:"to"`
	Gas      *big.Int `json:"gas"`
	GasPrice *big.Int `json:"gasprice"`
	Value    *big.Int `json:"value"`
	Data     []byte   `json:"data"`
}

// Transaction ...
type Transaction struct {
	Hash             Hash     `json:"hash"`
	Nonce            Hash     `json:"nonce"`
	BlockHash        Hash     `json:"blockHash"`
	BlockNumber      *big.Int `json:"blockNumber"`
	TransactionIndex uint64   `json:"transactionIndex"`
	From             Address  `json:"from"`
	To               Address  `json:"to"`
	Gas              *big.Int `json:"gas"`
	GasPrice         *big.Int `json:"gasprice"`
	Value            *big.Int `json:"value"`
	Data             []byte   `json:"input"`
}

// TransactionReceipt ...
type TransactionReceipt struct {
	Hash              Hash     `json:"transactionHash"`
	TransactionIndex  uint64   `json:"transactionIndex"`
	BlockNumber       *big.Int `json:"blockNumber"`
	BlockHash         Hash     `json:"blockHash"`
	CumulativeGasUsed *big.Int `json:"cumulativeGasUsed"`
	GasUsed           *big.Int `json:"gasUsed"`
	ContractAddress   Address  `json:"contractAddress"`
	Logs              []Log    `json:"logs"`
}

// Log ...
type Log struct {
	LogIndex         uint64     `json:"logIndex"`
	BlockNumber      *big.Int   `json:"blockNumber"`
	BlockHash        Hash       `json:"blockHash"`
	TransactionHash  Hash       `json:"transactionHash"`
	TransactionIndex uint64     `json:"transactionIndex"`
	Address          Address    `json:"address"`
	Data             [][32]byte `json:"data"`
	Topics           [][32]byte `json:"topics"`
}

// Block ...
type Block struct {
	Number          *big.Int `json:"number"`
	Hash            Hash     `json:"hash"`
	ParentHash      Hash     `json:"parentHash"`
	Nonce           Hash     `json:"nonce"`
	Sha3Uncles      Hash     `json:"sha3Uncles"`
	Bloom           Hash     `json:"logsBloom"`
	TransactionRoot Hash     `json:"transactionsRoot"`
	StateRoot       Hash     `json:"stateRoot"`
	Miner           Address  `json:"miner"`
	Difficulty      *big.Int `json:"difficulty"`
	TotalDifficulty *big.Int `json:"totalDifficulty"`
	ExtraData       Hash     `json:"extraData"`
	Size            *big.Int `json:"size"`
	GasLimit        *big.Int `json:"gasLimit"`
	MinGasPrice     *big.Int `json:"minGasPrice"`
	GasUsed         *big.Int `json:"gasUsed"`
	Timestamp       *big.Int `json:"timestamp"`
	Transactions    []Hash   `json:"transactions"`
	Uncles          []Hash   `json:"uncles"`
}