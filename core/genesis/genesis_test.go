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
package genesis

import (
	"os"
	"testing"

	"github.com/zeepin/ZeepinChain/common"
	"github.com/zeepin/ZeepinChain/common/config"
	"github.com/zeepin/ZeepinChain/common/log"
	"github.com/zeepin/ZeepinChain-Crypto/keypair"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	log.InitLog(0, log.Stdout)
	m.Run()
	os.RemoveAll("./ActorLog")
}

func TestGenesisBlockInit(t *testing.T) {
	_, pub, _ := keypair.GenerateKeyPair(keypair.PK_ECDSA, keypair.P256)
	conf := &config.GenesisConfig{}
	block, err := BuildGenesisBlock([]keypair.PublicKey{pub}, conf)
	assert.Nil(t, err)
	assert.NotNil(t, block)
	assert.NotEqual(t, block.Header.TransactionsRoot, common.UINT256_EMPTY)
}

func TestNewParamDeployAndInit(t *testing.T) {
	deployTx := newParamContract()
	initTx := newParamInit()
	assert.NotNil(t, deployTx)
	assert.NotNil(t, initTx)
}
