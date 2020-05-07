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

package account

import (
	"encoding/hex"
	"testing"
)

var id = "Gid:zpt:QbxSdfbWYsdWgN4TLrnZwXgL6bTPz4QRR9"

func TestCreate(t *testing.T) {
	nonce, _ := hex.DecodeString("4c6b58adc6b8c6774eee0eb07dac4e198df87aae28f8932db3982edf3ff026e4")
	id1, err := CreateID(nonce)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("result ID:", id1)
	if id != id1 {
		t.Fatal("expected ID:", id)
	}
}

func TestVerify(t *testing.T) {
	t.Log("verify", id)
	if !VerifyID(id) {
		t.Error("error: failed")
	}

	invalid := []string{
		"Gid:zpt:",
		"Gid:zpt:QbxSdfbWYsdWgN4TLrnZwXgL6bTPz4QRR9",
		"TSS6S4Xhzt5wtvRBTm4y3QCTRqB4BnU7vT",
		"Gid:else:TSS6S4Xhzt5wtvRBTm4y3QCT",
		"Gid:zpt: QbxSdfbWYsdWgN4TLrnZwXgL6bTPz4QRR9",
	}

	for _, v := range invalid {
		t.Log("verify", v)
		if VerifyID(v) {
			t.Error("error: passed")
		}
	}
}
