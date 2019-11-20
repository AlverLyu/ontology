/*
 * Copyright (C) 2018 The ontology Authors
 * This file is part of The ontology library.
 *
 * The ontology is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The ontology is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public License
 * along with The ontology.  If not, see <http://www.gnu.org/licenses/>.
 */
package ontid

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/ontio/ontology-crypto/keypair"
	"github.com/ontio/ontology/account"
	"github.com/ontio/ontology/common"
	"github.com/ontio/ontology/smartcontract/service/native"
	"github.com/ontio/ontology/smartcontract/service/native/utils"
)

func TestAttribute(t *testing.T) {
	testcase(t, CaseAttribute)
}

func CaseAttribute(t *testing.T, n *native.NativeService) {
	// 1. register id
	a := account.NewAccount("")
	id, err := account.GenerateID()
	if err != nil {
		t.Fatal("generate id error")
	}
	if err := regID(n, id, a); err != nil {
		t.Fatal(err)
	}

	// 2. add attribute
	attr := attribute{
		key:       []byte("test key"),
		valueType: []byte("test type"),
		value:     []byte("test value"),
	}
	if err := addAttr(n, id, attr, a); err != nil {
		t.Fatal(err)
	}

	// 3. check attribute
	if err := checkAttribute(n, id, []attribute{attr}); err != nil {
		t.Fatal(err)
	}
}

func addAttr(n *native.NativeService, id string, attr attribute, a *account.Account) error {
	sink := common.NewZeroCopySink(nil)
	sink.WriteString(id)
	utils.EncodeVarUint(sink, 1)
	attr.Serialization(sink)
	sink.WriteVarBytes(keypair.SerializePublicKey(a.PubKey()))
	n.Input = sink.Bytes()
	n.Tx.SignedAddr = []common.Address{a.Address}
	_, err := addAttributes(n)
	return err
}

func checkAttribute(n *native.NativeService, id string, attributes []attribute) error {
	sink := common.NewZeroCopySink(nil)
	sink.WriteString(id)
	n.Input = sink.Bytes()
	res, err := GetAttributes(n)
	if err != nil {
		return err
	}

	total := 0
	for _, a := range attributes {
		sink.Reset()
		a.Serialization(sink)
		b := sink.Bytes()
		if bytes.Index(res, b) == -1 {
			return fmt.Errorf("attribute %s not found", string(a.key))
		}
		total += len(b)
	}

	if len(res) != total {
		return fmt.Errorf("unmatched attribute number")
	}

	return nil
}
