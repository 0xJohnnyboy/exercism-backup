package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File = []bool
// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard = map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	if _, ok := cb[file]; !ok {
        return 0
    }
    
    occupiedSqsQty := 0
	for _, sq := range cb[file] {
        if sq == true { 
            occupiedSqsQty++
        }
    }

    return occupiedSqsQty
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
    if rank < 1 || rank > 8 {
        return 0
    }

    occupiedSqsQty := 0
    for _, file := range cb {
        if file[rank - 1] == true {
            occupiedSqsQty++
        }
    }

    return occupiedSqsQty
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	sqsQty := 0
    for _, f := range cb {
        sqsQty += len(f)
    }
    return sqsQty
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	occupiedSqsQty := 0 
    for _, file := range cb {
        for _, sq := range file {
            if sq == true {
                occupiedSqsQty++
            }
        }
    }
    return occupiedSqsQty
}
