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

// This package contains all the primitive utility to manage the parity in our HomeRaid
package parity

// ComputeParity is used when you create your parity block for the first time. E.g. you create a new HomeRaid volume
// starting from existing data and you need to fill your parity disk
// In this case this function has to be called once for each block. A block is considered to be 64 bits
func ComputeParity(blocksFromDisks []uint64) (parityBlock uint64) {
	for i := 0; i < len(blocksFromDisks); i++ {
		parityBlock = parityBlock ^ blocksFromDisks[i]
	}

	return parityBlock
}

// UpdateParity is used when you write new data in an esisting HomeRaid volume.
func UpdateParity(oldBlock, newBlock, oldParityBlock uint64) (newParityBlock uint64) {
	return oldParityBlock ^ oldBlock ^ newBlock
}

// RecoverLostBlock is used to recover lost data in case of a drive failure inside an existing HomeRaid volume
func RecoverLostBlock(remainingBlocks []uint64, parityBlock uint64) (recoveredBlock uint64) {
	for i := 0; i < len(remainingBlocks); i++ {
		recoveredBlock = recoveredBlock ^ remainingBlocks[i]
	}

	return recoveredBlock ^ parityBlock
}
