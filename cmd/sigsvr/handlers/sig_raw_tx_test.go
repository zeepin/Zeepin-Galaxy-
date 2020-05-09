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

package handlers

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"os"
	"testing"

	"github.com/zeepin/ZeepinChain/account"
	clisvrcom "github.com/zeepin/ZeepinChain/cmd/sigsvr/common"
	"github.com/zeepin/ZeepinChain/cmd/utils"
	"github.com/zeepin/ZeepinChain/common/log"
)

var (
	wallet *account.ClientImpl
	passwd = []byte("123456")
)

func TestMain(m *testing.M) {
	log.InitLog(0, os.Stdout)
	clisvrcom.DefAccount = account.NewAccount("")
	m.Run()
	os.RemoveAll("./ActorLog")
	os.RemoveAll("./Log")
}

func TestSigRawTx(t *testing.T) {
	acc := account.NewAccount("")
	defAcc := clisvrcom.DefAccount
	tx, err := utils.TransferTx(0, 0, "zpt", defAcc.Address.ToBase58(), acc.Address.ToBase58(), 10)
	if err != nil {
		t.Errorf("TransferTx error:%s", err)
		return
	}
	buf := bytes.NewBuffer(nil)
	err = tx.Serialize(buf)
	if err != nil {
		t.Errorf("tx.Serialize error:%s", err)
		return
	}
	rawReq := &SigRawTransactionReq{
		RawTx: hex.EncodeToString(buf.Bytes()),
	}
	data, err := json.Marshal(rawReq)
	if err != nil {
		t.Errorf("json.Marshal SigRawTransactionReq error:%s", err)
		return
	}
	req := &clisvrcom.CliRpcRequest{
		Qid:    "t",
		Method: "sigrawtx",
		Params: data,
	}
	resp := &clisvrcom.CliRpcResponse{}
	SigRawTransaction(req, resp)
	if resp.ErrorCode != 0 {
		t.Errorf("SigRawTransaction failed. ErrorCode:%d", resp.ErrorCode)
		return
	}
}
