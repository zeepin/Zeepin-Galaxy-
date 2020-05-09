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
	"math/big"

	"github.com/zeepin/ZeepinChain/embed/simulator/interfaces"
)

type Interop struct {
	_object interfaces.Interop
}

func NewInteropInterface(value interfaces.Interop) *Interop {
	var ii Interop
	ii._object = value
	return &ii
}

func (this *Interop) Equals(other StackItems) bool {
	v, err := other.GetInterface()
	if err != nil {
		return false
	}
	if this._object == nil || v == nil {
		return false
	}
	if !bytes.Equal(this._object.ToArray(), v.ToArray()) {
		return false
	}
	return true
}

func (this *Interop) GetBigInteger() (*big.Int, error) {
	return nil, fmt.Errorf("%s", "Not support interface to biginteger")
}

func (this *Interop) GetBoolean() (bool, error) {
	if this._object == nil {
		return false, nil
	}
	return true, nil
}

func (this *Interop) GetByteArray() ([]byte, error) {
	return nil, fmt.Errorf("%s", "Not support interface to bytearray")
}

func (this *Interop) GetInterface() (interfaces.Interop, error) {
	return this._object, nil
}

func (this *Interop) GetArray() ([]StackItems, error) {
	return nil, fmt.Errorf("%s", "Not support interface to array")
}

func (this *Interop) GetStruct() ([]StackItems, error) {
	return nil, fmt.Errorf("%s", "Not support interface to struct")
}

func (this *Interop) GetMap() (map[StackItems]StackItems, error) {
	return nil, fmt.Errorf("%s", "Not support interface to map")
}
