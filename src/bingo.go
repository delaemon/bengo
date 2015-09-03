package main
import (
	"fmt"
	"sort"
	"math/rand"
	"time"
)

func getCard() map[int][]int {
	card := make(map[int][]int)
	row := 0
	rand.Seed(time.Now().UnixNano())
	for row < 5 {
		line := make([]int, 0, 5)
		for len(line) < 5 {
			if row == 2 && len(line) == 2 {
				line = append(line, 0)
			} else {
				val := rand.Intn(15) + (15 * row) + 1
				uniq := true
				for _, l := range line {
					if l == val {
						uniq = false
					}
				}
				if uniq {
					line = append(line, val)
				}
			}
		}
		card[row] = line
		row++
	}
	return card
}

func keys(m map[int][]int) []int {
	keys := make([]int, 0, len(m))
	for key, _ := range m {
		keys = append(keys, key)
	}
	return keys
}

func line(length int) {
	for i := 0; i < length; i++ {
		fmt.Print("-")
	}
	fmt.Println("")
}

func grid(m map[int][]int) {
	keys := keys(m)
	sort.Sort(sort.IntSlice(keys))

	length := len(m[0]) * 3 + 1
	line(length)
	for row := 0; row < len(m[0]); row++ {
		fmt.Print("|")
		for _, col := range keys {
			fmt.Printf("%2d", m[col][row])
			fmt.Print("|")
		}
		fmt.Println("")
		line(length)
	}
}

func shuffle(list []int) {
	rand.Seed(time.Now().UnixNano())
	for i := range list {
		j := rand.Intn(i + 1)
		list[i], list[j] = list[j], list[i]
	}
}

func drawing(numbers []int, index int) int {
	return numbers[index]
}

func hit(target int, card map[int][]int) (bool, int, int) {
	row := (target - 1) / 15
	line := card[row]
	for col, v := range line {
		if v == target {
			card[row][col] = 0
			grid(card)
			return true, row, col
		}
	}
	return false, 0, 0
}

func goal(card map[int][]int, row int, col int) bool {

	if goalCol(card, col) {
		return true
	}

	if goalRow(card, row) {
		return true
	}

	if goalRightOblique(card) {
		return true
	}

	if goalLeftOblique(card) {
		return true
	}

	return false
}

func goalRightOblique(card map[int][]int) bool {
	goal := true
	row := 0
	col := 0
	for i := 0; i < 5; i++ {
		row, col = i, i
		fmt.Printf("row: %d, col: %d, val: %d\n", row, col, card[row][col])
		if card[col][row] != 0 {
			goal = false
			break
		}
	}
	return goal
}

func goalLeftOblique(card map[int][]int) bool {
	goal := true
	row := 0
	col := 0
	for i := 4; i >= 0; i-- {
		row, col = 4 - i, i
		fmt.Printf("row: %d, col: %d, val: %d\n", row, col, card[row][col])
		if card[col][row] != 0 {
			goal = false
			break
		}
	}
	return goal
}

func goalCol(card map[int][]int, col int) bool {
	goal := true
	for row := 0; row < 5; row++ {
		fmt.Printf("row: %d, col: %d, val: %d\n", row, col, card[row][col])
		if card[col][row] != 0 {
			goal = false
			break
		}
	}
	return goal
}

func goalRow(card map[int][]int, row int) bool {
	goal := true
	for col := 0; col < 5; col++ {
		fmt.Printf("row: %d, col: %d, val: %d\n", row, col, card[row][col])
		if card[col][row] != 0 {
			goal = false
			break
		}
	}
	return goal
}

func main() {
	cards := make([]map[int][]int, 0, 2)
	p1_card := getCard()
	p2_card := getCard()
	fmt.Println("player1")
	grid(p1_card)
	fmt.Println("player2")
	grid(p2_card)
	cards = append(cards, p1_card)
	cards = append(cards, p2_card)

	numbers := make([]int, 0, 100)
	for i := 1; i <= 75; i++ {
		numbers = append(numbers, i)
	}
	shuffle(numbers)

	for i := 1; i <= 75; i++ {
		for player, card := range cards {
			target := drawing(numbers, i)
			hit, row, col := hit(target, card)
			fmt.Printf("player: %d, target: %d, hit: %t \n", player+1, target, hit)
			if hit {
				goal := goal(card, row, col)
				if goal {
					fmt.Printf("player: %d goal!!!\n", player+1)
					return
				}
			}
		}
	}
}
