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

// ===================================================================================================================
// Constants used during the tests
//
var oldBlock uint64 = 1
var initialParityBlocks = []uint64{oldBlock, 2, 4}
var expectedInitialParity uint64 = 7
var remainingBlocks = []uint64{2, 4}

var newBlock uint64 = 55
var newParityBlocks = []uint64{newBlock, 2, 4}

// ===================================================================================================================
// Functions used during the tests
//
func simulateFailure(blocks []uint64, failedBlock uint64, parity uint64, t *testing.T) {
	toRecover := blocks[failedBlock]
	remainingBlocks := []uint64{}

	for i, b := range blocks {
		if uint64(i) != failedBlock {
			remainingBlocks = append(remainingBlocks, b)
		}
	}

	recovered := RecoverLostBlock(remainingBlocks, parity)

	if recovered != toRecover {
		t.Error("Recovered block", recovered, "expected block", toRecover)
	}
}

func bruteForce(blocks []uint64, t *testing.T) {
	parity := ComputeParity(blocks)

	for i := range blocks {
		simulateFailure(blocks, uint64(i), parity, t)
	}
}

// ===================================================================================================================
// Tests
//
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

func TestBruteForceCloseTo0(t *testing.T) {
	const maxValue uint64 = 100

	for x := uint64(0); x < maxValue; x++ {
		for y := uint64(0); y < maxValue; y++ {
			for z := uint64(0); z < maxValue; z++ {
				bruteForce([]uint64{x, y, z}, t)
			}
		}
	}
}

func TestBruteForceCloseToMax(t *testing.T) {
	const startValue uint64 = ^uint64(0) - 100
	const maxValue uint64 = ^uint64(0)

	for x := uint64(startValue); x < maxValue; x++ {
		for y := uint64(startValue); y < maxValue; y++ {
			for z := uint64(startValue); z < maxValue; z++ {
				bruteForce([]uint64{x, y, z}, t)
			}
		}
	}
}
