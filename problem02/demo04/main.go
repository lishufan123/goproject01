package main

import (
	"fmt"
)

const (
	boardSize = 15
	emptyCell = " "
	player1   = "X"
	player2   = "O"
)

var (
	board  [boardSize][boardSize]string
	player string
)

func main() {
	initializeBoard()
	player = player1

	fmt.Println("欢迎来玩五子棋！")
	printBoard()

	for {
		fmt.Printf("玩家 %s，请输入你的下棋坐标（例如：3 5）: ", player)
		var x, y int
		_, err := fmt.Scanf("%d %d", &x, &y)
		if err != nil || x < 1 || x > boardSize || y < 1 || y > boardSize || board[x-1][y-1] != emptyCell {
			fmt.Println("无效的坐标，请重新输入。")
			continue
		}

		board[x-1][y-1] = player
		printBoard()

		if checkWin(x-1, y-1) {
			fmt.Printf("玩家 %s 赢了！游戏结束。\n", player)
			break
		}

		if isBoardFull() {
			fmt.Println("平局！游戏结束。")
			break
		}

		switchPlayer()
	}
}

func initializeBoard() {
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			board[i][j] = emptyCell
		}
	}
}

func printBoard() {
	fmt.Println("  1 2 3 4 5 6 7 8 9 10 11 12 13 14 15")
	for i := 0; i < boardSize; i++ {
		fmt.Printf("%2d", i+1)
		for j := 0; j < boardSize; j++ {
			fmt.Printf(" %s", board[i][j])
		}
		fmt.Println()
	}
}

func switchPlayer() {
	if player == player1 {
		player = player2
	} else {
		player = player1
	}
}

func checkWin(x, y int) bool {
	directions := [][2]int{{1, 0}, {0, 1}, {1, 1}, {1, -1}}
	for _, dir := range directions {
		dx, dy := dir[0], dir[1]
		count := 1

		for i := -4; i <= 4; i++ {
			if i == 0 {
				continue
			}

			newX, newY := x+i*dx, y+i*dy
			if newX >= 0 && newX < boardSize && newY >= 0 && newY < boardSize && board[newX][newY] == player {
				count++
				if count == 5 {
					return true
				}
			} else {
				break
			}
		}
	}

	return false
}

func isBoardFull() bool {
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			if board[i][j] == emptyCell {
				return false
			}
		}
	}
	return true
}
