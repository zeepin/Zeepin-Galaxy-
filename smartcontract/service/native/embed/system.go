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

package embed

import (
	vm "github.com/imZhuFei/zeepin/embed/simulator"
	"github.com/imZhuFei/zeepin/errors"
)

// GetCodeContainer push current transaction to vm stack
func GetCodeContainer(service *EmbeddedService, engine *vm.ExecutionEngine) error {
	vm.PushData(engine, service.Tx)
	return nil
}

// GetExecutingAddress push current context to vm stack
func GetExecutingAddress(service *EmbeddedService, engine *vm.ExecutionEngine) error {
	context := service.ContextRef.CurrentContext()
	if context == nil {
		return errors.NewErr("Current context invalid")
	}
	vm.PushData(engine, context.ContractAddress[:])
	return nil
}

// GetExecutingAddress push previous context to vm stack
func GetCallingAddress(service *EmbeddedService, engine *vm.ExecutionEngine) error {
	context := service.ContextRef.CallingContext()
	if context == nil {
		return errors.NewErr("Calling context invalid")
	}
	vm.PushData(engine, context.ContractAddress[:])
	return nil
}

// GetExecutingAddress push entry call context to vm stack
func GetEntryAddress(service *EmbeddedService, engine *vm.ExecutionEngine) error {
	context := service.ContextRef.EntryContext()
	if context == nil {
		return errors.NewErr("Entry context invalid")
	}
	vm.PushData(engine, context.ContractAddress[:])
	return nil
}
