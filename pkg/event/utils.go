package event

func Length(chunkLn, chunkLastLn, chunkSize int) int {
	ln := chunkLn - 1
	// First segment of length calculation.
	fSegment := ln * chunkSize // [0 -> 9][...] -> [0 - 8][...] * utils.DefaultChunkSize
	// Last segment of length calculation.
	lSegment := chunkLastLn

	return fSegment + lSegment
}
