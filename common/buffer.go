/*
 * Copyright © 2021 ZkBAS Protocol
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package common

import (
	"bytes"
	"encoding/hex"
	"errors"
	"math/big"

	curve "github.com/bnb-chain/zkbas-crypto/ecc/ztwistededwards/tebn254"
	"github.com/bnb-chain/zkbas/types"
)

func PaddingStringBigIntIntoBuf(buf *bytes.Buffer, aStr string) error {
	a, isValid := new(big.Int).SetString(aStr, 10)
	if !isValid {
		return errors.New("[PaddingStringBigIntIntoBuf] invalid string")
	}
	buf.Write(a.FillBytes(make([]byte, curve.PointSize)))
	return nil
}

func PaddingAddressIntoBuf(buf *bytes.Buffer, address string) (err error) {
	if address == types.NilL1Address {
		buf.Write(new(big.Int).FillBytes(make([]byte, 32)))
		return nil
	}
	addrBytes, err := DecodeAddress(address)
	if err != nil {
		return err
	}
	buf.Write(new(big.Int).SetBytes(addrBytes).FillBytes(make([]byte, curve.PointSize)))
	return nil
}

func DecodeAddress(addr string) ([]byte, error) {
	if len(addr) != 42 {
		return nil, errors.New("[DecodeAddress] invalid address")
	}
	addrBytes, err := hex.DecodeString(addr[2:])
	if err != nil {
		return nil, err
	}
	if len(addrBytes) != types.AddressSize {
		return nil, errors.New("[DecodeAddress] invalid address")
	}
	return addrBytes, nil
}

func PaddingInt64IntoBuf(buf *bytes.Buffer, a int64) {
	buf.Write(new(big.Int).SetInt64(a).FillBytes(make([]byte, curve.PointSize)))
}

func PaddingPkIntoBuf(buf *bytes.Buffer, pkStr string) (err error) {
	pk, err := ParsePubKey(pkStr)
	if err != nil {
		return err
	}
	writePointIntoBuf(buf, &pk.A)
	return nil
}

func writePointIntoBuf(buf *bytes.Buffer, p *curve.Point) {
	buf.Write(p.X.Marshal())
	buf.Write(p.Y.Marshal())
}
