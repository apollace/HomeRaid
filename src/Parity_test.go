// HomeRaid
// Copyright (C) 2021  Alessandro Pollace
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package parity

import "testing"

var oldBlock uint64 = 1
var initialParityBlocks = []uint64{oldBlock, 2, 4}
var expectedInitialParity uint64 = 7
var remainingBlocks = []uint64{2, 4}

var newBlock uint64 = 55
var newParityBlocks = []uint64{newBlock, 2, 4}

func TestComputeParity_simple(t *testing.T) {
	var parityBlock = ComputeParity(initialParityBlocks)

	if parityBlock != expectedInitialParity {
		t.Error("Calculaed parity", parityBlock, "expected parity", expectedInitialParity)
	}
}

func TestUpdateParity_simple(t *testing.T) {
	var initialParity = ComputeParity(initialParityBlocks)
	var expectedNewParity = ComputeParity(newParityBlocks)
	var newParity = UpdateParity(oldBlock, newBlock, initialParity)

	if expectedNewParity != newParity {
		t.Error("Calculaed parity", newParity, "expected parity", expectedNewParity)
	}
}

func TestRecoverLostBlock_simple(t *testing.T) {
	var recoveredBlock = RecoverLostBlock(remainingBlocks, expectedInitialParity)

	if recoveredBlock != oldBlock {
		t.Error("Recovered block", recoveredBlock, "expected block", oldBlock)
	}
}
