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

package utils

import (
	"testing"

	"github.com/zeepin/ZeepinChain/common/log"
	msgCommon "github.com/zeepin/ZeepinChain/p2pserver/common"
	"github.com/zeepin/ZeepinChain/p2pserver/net/netserver"
	"github.com/zeepin/ZeepinChain/p2pserver/net/protocol"
	"github.com/zeepin/ZeepinChain-Crypto/keypair"
	"github.com/zeepin/ZeepinChain-Eventbus/actor"
	"github.com/stretchr/testify/assert"
)

func testHandler(data *msgCommon.MsgPayload, p2p p2p.P2P, pid *actor.PID, args ...interface{}) error {
	log.Info("Test handler")
	return nil
}

// TestMsgRouter tests a basic function of a message router
func TestMsgRouter(t *testing.T) {
	_, pub, _ := keypair.GenerateKeyPair(keypair.PK_ECDSA, keypair.P256)
	network := netserver.NewNetServer(pub)
	msgRouter := NewMsgRouter(network)
	assert.NotNil(t, msgRouter)

	msgRouter.RegisterMsgHandler("test", testHandler)
	msgRouter.UnRegisterMsgHandler("test")
	msgRouter.Start()
	msgRouter.Stop()
}
