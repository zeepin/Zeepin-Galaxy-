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
	"github.com/imZhuFei/zeepin/core/types"
	vm "github.com/imZhuFei/zeepin/embed/simulator"
	vmtypes "github.com/imZhuFei/zeepin/embed/simulator/types"
)

// BlockGetTransactionCount put block's transactions count to vm stack
func BlockGetTransactionCount(service *EmbeddedService, engine *vm.ExecutionEngine) error {
	i, err := vm.PopInteropInterface(engine)
	if err != nil {
		return err
	}
	vm.PushData(engine, len(i.(*types.Block).Transactions))
	return nil
}

// BlockGetTransactions put block's transactions to vm stack
func BlockGetTransactions(service *EmbeddedService, engine *vm.ExecutionEngine) error {
	i, err := vm.PopInteropInterface(engine)
	if err != nil {
		return err
	}
	transactions := i.(*types.Block).Transactions
	transactionList := make([]vmtypes.StackItems, 0)
	for _, v := range transactions {
		transactionList = append(transactionList, vmtypes.NewInteropInterface(v))
	}
	vm.PushData(engine, transactionList)
	return nil
}

// BlockGetTransaction put block's transaction to vm stack
func BlockGetTransaction(service *EmbeddedService, engine *vm.ExecutionEngine) error {
	i, err := vm.PopInteropInterface(engine)
	if err != nil {
		return err
	}
	index, err := vm.PopInt(engine)
	if err != nil {
		return err
	}
	vm.PushData(engine, i.(*types.Block).Transactions[index])
	return nil
}
