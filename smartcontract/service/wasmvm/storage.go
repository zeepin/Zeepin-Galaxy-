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
package wasmvm

import (
	"bytes"

	"github.com/zeepin/ZeepinChain/common"
	"github.com/zeepin/ZeepinChain/core/states"
	scommon "github.com/zeepin/ZeepinChain/core/store/common"
	"github.com/zeepin/ZeepinChain/errors"
	"github.com/zeepin/ZeepinChain/vm/wasmvm/exec"
	"github.com/zeepin/ZeepinChain/vm/wasmvm/memory"
	"github.com/zeepin/ZeepinChain/vm/wasmvm/util"
)

//======================store apis here============================================
func (this *WasmVmService) putstore(engine *exec.ExecutionEngine) (bool, error) {
	vm := engine.GetVM()
	envCall := vm.GetEnvCall()
	params := envCall.GetParams()
	if len(params) != 2 {
		return false, errors.NewErr("[putstore] parameter count error")
	}

	key, err := vm.GetPointerMemory(params[0])
	if err != nil {
		return false, err
	}
	if len(key) > 1024 {
		return false, errors.NewErr("[putstore] Get Storage key to long")
	}

	value, err := vm.GetPointerMemory(params[1])
	if err != nil {
		return false, err
	}
	k, err := serializeStorageKey(vm.ContractAddress, []byte(util.TrimBuffToString(key)))
	if err != nil {
		return false, err
	}
	this.CloneCache.Add(scommon.ST_STORAGE, k, &states.StorageItem{Value: value})

	vm.RestoreCtx()

	return true, nil
}

func (this *WasmVmService) getstore(engine *exec.ExecutionEngine) (bool, error) {
	vm := engine.GetVM()
	envCall := vm.GetEnvCall()
	params := envCall.GetParams()

	if len(params) != 1 {
		return false, errors.NewErr("[getstore] parameter count error ")
	}

	key, err := vm.GetPointerMemory(params[0])
	if err != nil {
		return false, err
	}
	k, err := serializeStorageKey(vm.ContractAddress, []byte(util.TrimBuffToString(key)))
	if err != nil {
		return false, err
	}
	item, err := this.CloneCache.Get(scommon.ST_STORAGE, k)
	if err != nil {
		return false, err
	}

	if item == nil {
		vm.RestoreCtx()
		if envCall.GetReturns() {
			vm.PushResult(uint64(memory.VM_NIL_POINTER))
		}
		return true, nil
	}
	idx, err := vm.SetPointerMemory(item.(*states.StorageItem).Value)
	if err != nil {
		return false, err
	}

	vm.RestoreCtx()
	if envCall.GetReturns() {
		vm.PushResult(uint64(idx))
	}
	return true, nil
}

func (this *WasmVmService) deletestore(engine *exec.ExecutionEngine) (bool, error) {

	vm := engine.GetVM()
	envCall := vm.GetEnvCall()
	params := envCall.GetParams()

	if len(params) != 1 {
		return false, errors.NewErr("[deletestore] parameter count error")
	}

	key, err := vm.GetPointerMemory(params[0])
	if err != nil {
		return false, err
	}

	k, err := serializeStorageKey(vm.ContractAddress, []byte(util.TrimBuffToString(key)))
	if err != nil {
		return false, err
	}

	this.CloneCache.Delete(scommon.ST_STORAGE, k)
	vm.RestoreCtx()

	return true, nil
}

func serializeStorageKey(contractAddress common.Address, key []byte) ([]byte, error) {
	bf := new(bytes.Buffer)
	storageKey := &states.StorageKey{ContractAddress: contractAddress, Key: key}
	if _, err := storageKey.Serialize(bf); err != nil {
		return []byte{}, errors.NewErr("[serializeStorageKey] StorageKey serialize error!")
	}
	return bf.Bytes(), nil
}
