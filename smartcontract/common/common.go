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

package common

import (
	"github.com/zeepin/ZeepinChain/common"
	"github.com/zeepin/ZeepinChain/common/log"
	"github.com/zeepin/ZeepinChain/embed/simulator/types"
)

// ConvertReturnTypes return embeded stack element value
// According item types convert to hex string value
// Now embeded support type contain: ByteArray/Integer/Boolean/Array/Struct/Interop/StackItems
func ConvertEmbededTypeHexString(item interface{}) interface{} {
	if item == nil {
		return nil
	}
	switch v := item.(type) {
	case *types.ByteArray:
		arr, _ := v.GetByteArray()
		return common.ToHexString(arr)
	case *types.Integer:
		i, _ := v.GetBigInteger()
		if i.Sign() == 0 {
			return common.ToHexString([]byte{0})
		} else {
			return common.ToHexString(common.BigIntToEmbededBytes(i))
		}
	case *types.Boolean:
		b, _ := v.GetBoolean()
		if b {
			return common.ToHexString([]byte{1})
		} else {
			return common.ToHexString([]byte{0})
		}
	case *types.Array:
		var arr []interface{}
		ar, _ := v.GetArray()
		for _, val := range ar {
			arr = append(arr, ConvertEmbededTypeHexString(val))
		}
		return arr
	case *types.Struct:
		var arr []interface{}
		ar, _ := v.GetStruct()
		for _, val := range ar {
			arr = append(arr, ConvertEmbededTypeHexString(val))
		}
		return arr
	case *types.Interop:
		it, _ := v.GetInterface()
		return common.ToHexString(it.ToArray())
	default:
		log.Error("[ConvertTypes] Invalid Types!")
		return nil
	}
}
