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
	"github.com/imZhuFei/zeepin/common"
	"github.com/imZhuFei/zeepin/errors"
	"github.com/imZhuFei/zeepin/smartcontract/event"
	"github.com/imZhuFei/zeepin/smartcontract/service/native"
)

func AddCommonEvent(native *native.NativeService, contract common.Address, name string, params interface{}) {
	native.Notifications = append(native.Notifications,
		&event.NotifyEventInfo{
			ContractAddress: contract,
			States:          []interface{}{name, params},
		})
}

func ConcatKey(contract common.Address, args ...[]byte) []byte {
	temp := contract[:]
	for _, arg := range args {
		temp = append(temp, arg...)
	}
	return temp
}

func ValidateOwner(native *native.NativeService, address common.Address) error {
	if native.ContextRef.CheckWitness(address) == false {
		return errors.NewErr("validateOwner, authentication failed!")
	}
	return nil
}
