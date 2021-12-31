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

func bruteForceRecovery(blocks []uint64, t *testing.T) {
	parity := ComputeParity(blocks)

	for i := range blocks {
		simulateFailure(blocks, uint64(i), parity, t)
	}
}

func BruteForceSingleBlockUpdate(blockIndex uint64, blocks []uint64, startValue uint64, maxValue uint64, t *testing.T) {
	incrementalParity := ComputeParity(blocks)

	for i := uint64(startValue); i < maxValue; i++ {
		oldBlock := blocks[blockIndex]
		blocks[blockIndex] = i

		incrementalParity = UpdateParity(oldBlock, blocks[blockIndex], incrementalParity)

		parityComputedFromScratch := ComputeParity(blocks)
		if incrementalParity != parityComputedFromScratch {
			t.Error("Fresh computed parity", parityComputedFromScratch, "rolling parity", incrementalParity, "blocks", blocks)
		}

		for i := range blocks {
			simulateFailure(blocks, uint64(i), incrementalParity, t)
		}
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

func TestBruteForceRecoveryCloseTo0(t *testing.T) {
	const maxValue uint64 = 100

	for x := uint64(0); x < maxValue; x++ {
		for y := uint64(0); y < maxValue; y++ {
			for z := uint64(0); z < maxValue; z++ {
				bruteForceRecovery([]uint64{x, y, z}, t)
			}
		}
	}
}

func TestBruteForceRecoveryCloseToMax(t *testing.T) {
	const startValue uint64 = ^uint64(0) - 100
	const maxValue uint64 = ^uint64(0)

	for x := uint64(startValue); x < maxValue; x++ {
		for y := uint64(startValue); y < maxValue; y++ {
			for z := uint64(startValue); z < maxValue; z++ {
				bruteForceRecovery([]uint64{x, y, z}, t)
			}
		}
	}
}

func TestBruteForceUpdates(t *testing.T) {
	// In this test given a set of blocks and their parity:
	// [Block0, Block1, Block2] Parity
	//
	// We try to change each block selectivelly, calculate the parity incrementally
	// and verify that the incremental pairty is the same that we compute if we compute it
	// from scratch

	const block0 uint64 = 0
	const block1 uint64 = 1
	const block2 uint64 = 2

	const startValue uint64 = 0
	const maxValue uint64 = 10000

	const startValueMax uint64 = ^uint64(0) - 10000
	const maxValueMax uint64 = ^uint64(0)

	// Bruteforce close to 0
	BruteForceSingleBlockUpdate(block0, initialParityBlocks, startValue, maxValue, t)
	BruteForceSingleBlockUpdate(block1, initialParityBlocks, startValue, maxValue, t)
	BruteForceSingleBlockUpdate(block2, initialParityBlocks, startValue, maxValue, t)

	// Bruteforce close to max uint64
	BruteForceSingleBlockUpdate(block0, initialParityBlocks, startValueMax, maxValueMax, t)
	BruteForceSingleBlockUpdate(block1, initialParityBlocks, startValueMax, maxValueMax, t)
	BruteForceSingleBlockUpdate(block2, initialParityBlocks, startValueMax, maxValueMax, t)
}
