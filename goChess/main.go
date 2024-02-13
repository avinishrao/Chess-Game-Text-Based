package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const boardSize = 8

type ChessBoard [boardSize][boardSize]string

var upperTurn = true // Variable to track whose turn it is (true for upper side, false for lower side)

func initializeBoard() ChessBoard {
	board := ChessBoard{}
	// Initialize the board with pieces
	// This is a simplified starting position
	board[0] = [boardSize]string{"r", "n", "b", "q", "k", "b", "n", "r"}
	board[1] = [boardSize]string{"p", "p", "p", "p", "p", "p", "p", "p"}
	for i := 2; i < 6; i++ {
		board[i] = [boardSize]string{" ", " ", " ", " ", " ", " ", " ", " "}
	}
	board[6] = [boardSize]string{"P", "P", "P", "P", "P", "P", "P", "P"}
	board[7] = [boardSize]string{"R", "N", "B", "Q", "K", "B", "N", "R"}
	return board
}

func printBoard(board ChessBoard) {
	fmt.Println("  a b c d e f g h")
	for row := 0; row < boardSize; row++ {
		fmt.Printf("%d ", row+1)
		for col := 0; col < boardSize; col++ {
			fmt.Printf("%s ", board[row][col])
		}
		fmt.Println()
	}
}

func isValidMove(fromCol, fromRow, toCol, toRow int, board ChessBoard) bool {
	if fromCol < 0 || fromCol >= boardSize || fromRow < 0 || fromRow >= boardSize || toCol < 0 || toCol >= boardSize || toRow < 0 || toRow >= boardSize {
		return false
	}

	piece := board[fromRow][fromCol]
	// Check if the current player is allowed to move this piece
	if (upperTurn && !isUpper(piece)) || (!upperTurn && isUpper(piece)) {
		return false
	}

	switch piece {
	case "p": // Black Pawn
		// Check for regular pawn move (one square forward)
		if fromCol == toCol && toRow == fromRow+1 && board[toRow][toCol] == " " {
			return true
		}
		// Check for initial double pawn move (two squares forward)
		if fromCol == toCol && toRow == fromRow+2 && fromRow == 1 && board[fromRow+1][toCol] == " " && board[toRow][toCol] == " " {
			return true
		}
		// Check for capturing an opponent's piece diagonally
		if abs(toCol-fromCol) == 1 && toRow == fromRow+1 && board[toRow][toCol] != " " && isUpper(board[toRow][toCol]) {
			return true
		}
	case "P": // White Pawn
		// Check for regular pawn move (one square forward)
		if fromCol == toCol && toRow == fromRow-1 && board[toRow][toCol] == " " {
			return true
		}
		// Check for initial double pawn move (two squares forward)
		if fromCol == toCol && toRow == fromRow-2 && fromRow == 6 && board[fromRow-1][toCol] == " " && board[toRow][toCol] == " " {
			return true
		}
		// Check for capturing an opponent's piece diagonally
		if abs(toCol-fromCol) == 1 && toRow == fromRow-1 && board[toRow][toCol] != " " && !isUpper(board[toRow][toCol]) {
			return true
		}
	case "r", "R": // Rook
		return isValidRookMove(fromCol, fromRow, toCol, toRow, board)
	case "b", "B": // Bishop
		return isValidBishopMove(fromCol, fromRow, toCol, toRow, board)
	case "n", "N": // Knight
		return isValidKnightMove(fromCol, fromRow, toCol, toRow, board)
	case "q", "Q": // Queen
		return isValidQueenMove(fromCol, fromRow, toCol, toRow, board)
	case "k", "K": // King
		return isValidKingMove(fromCol, fromRow, toCol, toRow, board)
	default:
		return false
	}

	return false
}

func isUpper(s string) bool {
	return strings.ToUpper(s) == s
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func isValidKingMove(fromCol, fromRow, toCol, toRow int, board ChessBoard) bool {
	// Check if the destination is within the chessboard bounds
	if toRow < 0 || toRow >= boardSize || toCol < 0 || toCol >= boardSize {
		return false
	}

	// Calculate the absolute difference in rows and columns
	rowDiff := abs(fromRow - toRow)
	colDiff := abs(fromCol - toCol)

	// Check if the move is within one square in any direction
	if rowDiff <= 1 && colDiff <= 1 {
		return true
	}
	return false
}

func isValidQueenMove(fromCol, fromRow, toCol, toRow int, board ChessBoard) bool {
	// Check if the destination is within the chessboard bounds
	if toRow < 0 || toRow >= boardSize || toCol < 0 || toCol >= boardSize {
		return false
	}

	// Calculate the absolute difference in rows and columns
	rowDiff := abs(fromRow - toRow)
	colDiff := abs(fromCol - toCol)

	// Check if the move is either vertical, horizontal, or diagonal
	if fromRow == toRow || fromCol == toCol || rowDiff == colDiff {
		// Check if the path is clear for a rook-like move
		if fromRow == toRow {
			minCol, maxCol := min(fromCol, toCol), max(fromCol, toCol)
			for col := minCol + 1; col < maxCol; col++ {
				if board[fromRow][col] != " " {
					return false
				}
			}
		} else if fromCol == toCol {
			minRow, maxRow := min(fromRow, toRow), max(fromRow, toRow)
			for row := minRow + 1; row < maxRow; row++ {
				if board[row][fromCol] != " " {
					return false
				}
			}
		} else { // Diagonal move
			rowDir := 1
			if fromRow > toRow {
				rowDir = -1
			}
			colDir := 1
			if fromCol > toCol {
				colDir = -1
			}
			for row, col := fromRow+rowDir, fromCol+colDir; row != toRow; row, col = row+rowDir, col+colDir {
				if board[row][col] != " " {
					return false
				}
			}
		}
		return true
	}
	return false
}

func isValidKnightMove(fromCol, fromRow, toCol, toRow int, board ChessBoard) bool {
	// Check if the destination is within the chessboard bounds
	if toRow < 0 || toRow >= boardSize || toCol < 0 || toCol >= boardSize {
		return false
	}

	// Calculate the absolute difference in rows and columns
	rowDiff := abs(fromRow - toRow)
	colDiff := abs(fromCol - toCol)

	// Check if the move is in an L-shape (2 squares horizontally and 1 square vertically, or vice versa)
	return (rowDiff == 1 && colDiff == 2) || (rowDiff == 2 && colDiff == 1)
}

func isValidRookMove(fromCol, fromRow, toCol, toRow int, board ChessBoard) bool {
	// Check if the destination is within the chessboard bounds
	if toRow < 0 || toRow >= boardSize || toCol < 0 || toCol >= boardSize {
		return false
	}

	// Check if the move is either vertical or horizontal
	if fromRow == toRow || fromCol == toCol {
		// Check if there are any pieces blocking the path
		if fromRow == toRow {
			minCol, maxCol := min(fromCol, toCol), max(fromCol, toCol)
			for col := minCol + 1; col < maxCol; col++ {
				if board[fromRow][col] != " " {
					return false
				}
			}
		} else { // fromCol == toCol
			minRow, maxRow := min(fromRow, toRow), max(fromRow, toRow)
			for row := minRow + 1; row < maxRow; row++ {
				if board[row][fromCol] != " " {
					return false
				}
			}
		}
		return true
	}
	return false
}

func isValidBishopMove(fromCol, fromRow, toCol, toRow int, board ChessBoard) bool {
	// Check if the destination is within the chessboard bounds
	if toRow < 0 || toRow >= boardSize || toCol < 0 || toCol >= boardSize {
		return false
	}

	// Check if the move is diagonal
	if abs(toRow-fromRow) == abs(toCol-fromCol) {
		// Check if there are any pieces blocking the path
		rowDir := 1
		if toRow < fromRow {
			rowDir = -1
		}
		colDir := 1
		if toCol < fromCol {
			colDir = -1
		}
		for row, col := fromRow+rowDir, fromCol+colDir; row != toRow; row, col = row+rowDir, col+colDir {
			if board[row][col] != " " {
				return false
			}
		}
		return true
	}
	return false
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	board := initializeBoard()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		printBoard(board)
		var playerColor string
		if upperTurn {
			playerColor = "White"
		} else {
			playerColor = "Black"
		}
		fmt.Printf("Enter move for %s (e.g., 'e2 to e4') or 'exit' to quit: \n", playerColor)
		scanner.Scan()
		move := scanner.Text()

		if move == "exit" {
			break
		}

		parts := strings.Split(move, " to ")
		if len(parts) != 2 {
			fmt.Println("Invalid move format. Please use 'e2 to e4' format.")
			continue
		}
		fromSquare, toSquare := parts[0], parts[1]

		fromCol := int(fromSquare[0] - 'a')
		fromRow := int(fromSquare[1] - '1')
		toCol := int(toSquare[0] - 'a')
		toRow := int(toSquare[1] - '1')

		if !isValidMove(fromCol, fromRow, toCol, toRow, board) {
			fmt.Println("Invalid move.")
			continue
		}

		board[toRow][toCol] = board[fromRow][fromCol]
		board[fromRow][fromCol] = " "
		upperTurn = !upperTurn // Switch turn
	}
}
