/*
 * Copyright (C) 2018 The ZeepinChain Authors
 * This file is part of The ZeepinChain library.
 *
 * The ZeepinChain is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The ZeepinChain is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public License
 * along with The ZeepinChain.  If not, see <http://www.gnu.org/licenses/>.
 */

package types

import (
	"bytes"
	"fmt"
	"io"

	"github.com/imZhuFei/zeepin/common"
	"github.com/imZhuFei/zeepin/common/serialization"
)

type Block struct {
	Header       *Header
	Transactions []*Transaction
}

func (b *Block) Serialize(w io.Writer) error {
	err := b.Header.Serialize(w)
	if err != nil {
		return err
	}

	err = serialization.WriteUint32(w, uint32(len(b.Transactions)))
	if err != nil {
		return fmt.Errorf("Block item Transactions length serialization failed: %s", err)
	}

	for _, transaction := range b.Transactions {
		err := transaction.Serialize(w)
		if err != nil {
			return err
		}
	}
	return nil
}

func (b *Block) Deserialize(r io.Reader) error {
	if b.Header == nil {
		b.Header = new(Header)
	}
	err := b.Header.Deserialize(r)
	if err != nil {
		return err
	}

	//Transactions
	length, err := serialization.ReadUint32(r)
	if err != nil {
		return err
	}

	var hashes = make([]common.Uint256, 0, length)
	for i := uint32(0); i < length; i++ {
		transaction := new(Transaction)
		err := transaction.Deserialize(r)
		if err != nil {
			return err
		}
		txhash := transaction.Hash()
		b.Transactions = append(b.Transactions, transaction)
		hashes = append(hashes, txhash)
	}

	b.Header.TransactionsRoot = common.ComputeMerkleRoot(hashes)

	return nil
}

func (b *Block) ToArray() []byte {
	bf := new(bytes.Buffer)
	b.Serialize(bf)
	return bf.Bytes()
}

func (b *Block) Hash() common.Uint256 {
	return b.Header.Hash()
}

func (b *Block) Type() common.InventoryType {
	return common.BLOCK
}

func (b *Block) RebuildMerkleRoot() {
	txs := b.Transactions
	hashes := []common.Uint256{}
	for _, tx := range txs {
		hashes = append(hashes, tx.Hash())
	}
	hash := common.ComputeMerkleRoot(hashes)
	b.Header.TransactionsRoot = hash
}
