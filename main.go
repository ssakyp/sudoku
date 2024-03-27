package main

import (
	"os"

	"github.com/01-edu/z01"
)

// Выводит на стандартный вывод игровую доску судоку
func printSudokuBoard(board [][]rune) {
	for _, row := range board {
		for i, e := range row {
			z01.PrintRune(e)
			if i != len(row)-1 {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}

// Возвращает true, если значение 'value' находится в строке 'y' доски 'board'
func isInRow(board [][]rune, value rune, x, y int) bool {
	for i, v := range board[y] {
		if i != x && v == value {
			return true
		}
	}
	return false
}

// Возвращает true, если значение 'value' находится в столбце 'x' доски 'board'
func isInColumn(board [][]rune, value rune, x, y int) bool {
	for i := 0; i < 9; i++ {
		if i != y && board[i][x] == value {
			return true
		}
	}
	return false
}

// Принимает целое число 'x' и возвращает начало и конец
// интервала последовательных пар кратных трём, в котором находится 'x'.
// Например, для x = 2, x находится в (0, 3)
// для x = 4, x находится в (3, 6)
func intervalThree(x int, max int) (int, int) {
	var i int
	for i = 0; i < max; i += 3 {
		if x >= i && x < i+3 {
			break
		}
	}
	endi := 3 + i
	return i, endi
}

// Возвращает true, если значение 'value' разрешено на позиции (x, y) доски 'board'
func isAllowedInBox(board [][]rune, value rune, x, y int) bool {
	n := len(board)
	begi, endi := intervalThree(x, n)
	begj, endj := intervalThree(y, n)
	for j := begj; j < endj; j++ {
		for i := begi; i < endi; i++ {
			if (j != y || i != i) && board[j][i] == value {
				return false
			}
		}
	}
	return true
}

// Возвращает true, если позиция не имеет определенного значения
// То есть, если символ - точка '.'
func isEmpty(board [][]rune, x, y int) bool {
	return board[y][x] == '.'
}

// Возвращает все пустые позиции на доске
func availablePos(board [][]rune) [][]int {
	var ava [][]int
	for y, row := range board {
		for x, e := range row {
			if e == '.' {
				ava = append(ava, []int{x, y})
			}
		}
	}
	return ava
}

// Проверяет, является ли доска валидной судоку
func validBoard(board [][]rune) bool {
	size := 9
	if len(board) != size {
		return false
	}
	for y, row := range board {
		if len(row) != size {
			return false
		}
		for x, e := range row {
			if (e < '1' || e > '9') && e != '.' {
				return false
			}
			if e != '.' && !isAllowedInBox(board, e, x, y) {
				return false
			}
		}
	}
	return true
}

// Рекурсивная функция для решения головоломки судоку
func sudokuH(board [][]rune, available [][]int, i int) bool {
	n := len(available)

	if i >= n {
		return true
	}

	x := available[i][0]
	y := available[i][1]

	for c := '1'; c <= '9'; c++ {
		if isAllowedInBox(board, c, x, y) {
			board[y][x] = c
			if sudokuH(board, available, i+1) {
				return true
			}
			board[y][x] = '.'
		}
	}
	return false
}

// Заполняет пустые позиции доски правильными значениями
func sudokuSolver(board [][]rune) {
	available := availablePos(board)
	if sudokuH(board, available, 0) {
		printSudokuBoard(board)
	} else {
		// Выводит "Error", если судоку не имеет решения
		z01.PrintRune('E')
		z01.PrintRune('r')
		z01.PrintRune('r')
		z01.PrintRune('o')
		z01.PrintRune('r')
		z01.PrintRune('\n')
	}
}

func main() {
	var board [][]rune

	// Разбирает аргументы командной строки для получения доски судоку
	for _, v := range os.Args[1:] {
		board = append(board, []rune(v))
	}

	// Проверяет, является ли доска валидной, и решает судоку
	if validBoard(board) {
		sudokuSolver(board)
	} else {
		// Выводит "Error", если доска невалидна
		z01.PrintRune('E')
		z01.PrintRune('r')
		z01.PrintRune('r')
		z01.PrintRune('o')
		z01.PrintRune('r')
		z01.PrintRune('\n')
	}
}
