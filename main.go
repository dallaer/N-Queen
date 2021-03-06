package main

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"strconv"
	"time"
)

func createMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	return matrix
}

func printInFile(m [][]int) {
	s, err := os.OpenFile("file.txt", os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Unexpected error while open file")
	}
	for i := 0; i < len(m); i++ {
		data, err1 := json.Marshal(m[i])
		if err1 != nil {
			fmt.Println("Unexpected error")
		}
		s.Write(data)
		s.WriteString("\n")
	}
	s.WriteString("\n")

}

func printM(pos [][]int, s []int) {
	matrix := createMatrix(len(pos) + 1)
	for i := 0; i < len(pos); i++ {
		matrix[pos[i][0]-1][pos[i][1]-1] = 1
	}
	matrix[s[0]-1][s[1]-1] = 1
	printInFile(matrix)
}

func get_que(n, x, combs int, pos [][]int) int {
	for y := 1; y <= n; y++ {
		can_put := true
		for i := range pos {
			X, Y := pos[i][0], pos[i][1]
			if X == x || Y == y || math.Abs(float64(X-x)) == math.Abs(float64(Y-y)) {
				can_put = false
				break
			}
		}
		if can_put {
			if x == n {
				printM(pos, []int{x, y})
				return (combs + 1)
			} else {
				pos_copy := pos
				pos_copy = append(pos_copy, []int{x, y})
				combs = get_que(n, x+1, combs, pos_copy)
			}
		}
	}
	return combs
}

func W8() {
	for true {
		fmt.Println("Waiting...")
		time.Sleep(5 * time.Second)
	}
}

func main() {
	var n string
	for true {
		fmt.Println("Enter number")
		_, err := os.Create("file.txt")
		if err != nil {
			fmt.Println("Unexpected error while creating file")
		}
		fmt.Scan(&n)
		if n == "quit" {
			break
		}
		x, _ := strconv.Atoi(n)
		if x > 0 {
			go W8()
			fmt.Println("For ", x, " Queen ", get_que(x, 1, 0, [][]int{}), " difference combination\nYou can check it in the file")
			break
		} else {
			fmt.Println("Incorrect input. Try again. Enter quit for break.")
		}
	}
}
