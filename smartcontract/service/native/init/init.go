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

package init

import (
	"bytes"
	"math"
	"math/big"

	"github.com/zeepin/ZeepinChain/common"
	invoke "github.com/zeepin/ZeepinChain/core/utils"
	vm "github.com/zeepin/ZeepinChain/embed/simulator"
	"github.com/zeepin/ZeepinChain/smartcontract/service/native/auth"
	"github.com/zeepin/ZeepinChain/smartcontract/service/native/embed"
	"github.com/zeepin/ZeepinChain/smartcontract/service/native/gala"
	"github.com/zeepin/ZeepinChain/smartcontract/service/native/gid"
	params "github.com/zeepin/ZeepinChain/smartcontract/service/native/global_params"
	"github.com/zeepin/ZeepinChain/smartcontract/service/native/governance"
	"github.com/zeepin/ZeepinChain/smartcontract/service/native/utils"
	"github.com/zeepin/ZeepinChain/smartcontract/service/native/zpt"
)

var (
	COMMIT_DPOS_BYTES = InitBytes(utils.GovernanceContractAddress, governance.COMMIT_DPOS)
)

func init() {
	gala.InitGala()
	zpt.InitZpt()
	params.InitGlobalParams()
	gid.Init()
	auth.Init()
	governance.InitGovernance()
}

func InitBytes(addr common.Address, method string) []byte {
	bf := new(bytes.Buffer)
	builder := vm.NewParamsBuilder(bf)
	builder.EmitPushByteArray([]byte{})
	builder.EmitPushByteArray([]byte(method))
	builder.EmitPushByteArray(addr[:])
	builder.EmitPushInteger(big.NewInt(0))
	builder.Emit(vm.SYSCALL)
	builder.EmitPushByteArray([]byte(embed.NATIVE_INVOKE_NAME))

	tx := invoke.NewInvokeTransaction(builder.ToArray())
	tx.GasLimit = math.MaxUint64
	return bf.Bytes()
}
