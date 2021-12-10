package bingo

type BoardMeta struct {
	Position int
	Marked   bool
}

type BingoBoard struct {
	Board  map[int]*BoardMeta
	Marked []bool
	// 5x5 = 25 places.
}

func NewBingoBoard() *BingoBoard {
	return &BingoBoard{
		Board:  make(map[int]*BoardMeta),
		Marked: make([]bool, 25),
	}
}

func checkIfWon(board *BingoBoard) bool {
	consRow := 0
	consCol := 0
	col := 0
	for row := 0; row < 25; row++ {
		// Start over counting we moved to a new row/col
		if row%5 == 0 {
			consRow = 0
			consCol = 0
			col = row / 5
		}
		if board.Marked[row] {
			consRow++
		}
		if board.Marked[col] {
			consCol++
		}
		// We won!
		if consCol == 5 || consRow == 5 {
			return true
		}
		col += 5
	}
	return false
}

func checkBoards(boards []*BingoBoard, number int) int {
	// Fill number
	for i, board := range boards {
		if meta, ok := board.Board[number]; ok {
			board.Marked[meta.Position] = true
			board.Board[number].Marked = true
			if checkIfWon(board) {
				return i
			}
		}
	}
	// We never found a board :(
	return -1
}

func checkAllBoards(boards []*BingoBoard, wonBoards []bool, number int) int {
	for i, board := range boards {
		if meta, ok := board.Board[number]; ok {
			board.Marked[meta.Position] = true
			board.Board[number].Marked = true
			if checkIfWon(board) && !wonBoards[i] {
				wonBoards[i] = true
				if allWon(wonBoards) {
					return i
				}
			}
		}
	}
	// We never found a board :(
	return -1
}

func countScore(board *BingoBoard, lastNumber int) int {
	unmarkedSum := 0
	for k, v := range board.Board {
		if !v.Marked {
			unmarkedSum += k
		}
	}
	return unmarkedSum * lastNumber
}

func allWon(wonBoards []bool) bool {
	for _, val := range wonBoards {
		if !val {
			return false
		}
	}
	return true
}

func PlayBingo(boards []*BingoBoard, lottoNumbers []int) int {
	for _, number := range lottoNumbers {
		if nbr := checkBoards(boards, number); nbr != -1 {
			return countScore(boards[nbr], number)
		}
	}
	return -1
}

func CheckLastBingo(boards []*BingoBoard, lottoNumbers []int) int {
	// Assume that all boards can win
	nbrBoards := len(boards)
	wonBoards := make([]bool, nbrBoards)
	for _, number := range lottoNumbers {
		if nbr := checkAllBoards(boards, wonBoards, number); nbr != -1 {
			wonBoards[nbr] = true
			if allWon(wonBoards) {
				return countScore(boards[nbr], number)
			}
		}
	}
	return -1
}
