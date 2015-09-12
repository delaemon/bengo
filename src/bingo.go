package main
import (
	"fmt"
	"sort"
	"math/rand"
	"time"
	"strconv"
)

const (
	Red    = "0;31m"
	Green  = "0;32m"
	Yellow = "0;33m"
	Blue   = "0;34m"
	PurPle = "0;35m"
	Cyan   = "0;36m"
)

var (
	debug bool
)

func printColor(out, color string) {
	fmt.Printf("\x1b[%s%s\x1b[0m", color, out)
}

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

func gridHit(m map[int][]int, x ,y int) {
	keys := keys(m)
	sort.Sort(sort.IntSlice(keys))

	length := len(m[0]) * 3 + 1
	line(length)
	for row := 0; row < len(m[0]); row++ {
		fmt.Print("|")
		for _, col := range keys {
			if x == col && y == row {
				printColor(fmt.Sprintf("%2d", m[col][row]), Green)
			} else {
				fmt.Printf("%2d", m[col][row])
			}
			fmt.Print("|")
		}
		fmt.Println("")
		line(length)
	}
}

func gridReach(m map[int][]int, hits [][2]int) {
	keys := keys(m)
	sort.Sort(sort.IntSlice(keys))

	length := len(m[0]) * 3 + 1
	line(length)
	h := 0
	for row := 0; row < len(m[0]); row++ {
		fmt.Print("|")
		for _, col := range keys {
			if m[col][row] == 0 {
				if len(hits[h]) > 0 && col == hits[h][0] && row == hits[h][1] {
					if h < (len(hits) - 1) {
						h++
					}
					printColor(fmt.Sprintf("%2d", m[col][row]), Yellow)
				} else {
					fmt.Printf("%2d", m[col][row])
				}
			} else {
				fmt.Printf("%2d", m[col][row])
			}
			fmt.Print("|")
		}
		fmt.Println("")
		line(length)
	}

}

func gridGoal(m map[int][]int, hits [][2]int) {
	keys := keys(m)
	sort.Sort(sort.IntSlice(keys))

	length := len(m[0]) * 3 + 1
	line(length)
	h := 0
	for row := 0; row < len(m[0]); row++ {
		fmt.Print("|")
		for _, col := range keys {
			if m[col][row] == 0 {
				if len(hits[h]) > 0 && col == hits[h][0] && row == hits[h][1] {
					if h < (len(hits) - 1) {
						h++
					}
					printColor(fmt.Sprintf("%2d", m[col][row]), Red)
				} else {
					printColor(fmt.Sprintf("%2d", m[col][row]), Cyan)
				}
			} else {
				fmt.Printf("%2d", m[col][row])
			}
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
	col := (target - 1) / 15
	line := card[col]
	for row, v := range line {
		if v == target {
			card[col][row] = 0
			return true, col, row
		}
	}
	return false, 0, 0
}

func goalLeftOblique(card map[int][]int) [][2]int {
	line := make([][2]int, 0, 5)
	row := 0
	col := 0
	for i := 0; i < 5; i++ {
		row, col = i, i
		if debug {
			fmt.Printf("[\\] col: %d, row: %d, val: %d\n", col, row, card[col][row])
		}
		if card[col][row] == 0 {
			line = append(line, [2]int{col, row})
		}
	}
	return line
}

func goalRightOblique(card map[int][]int) [][2]int {
	line := make([][2]int, 0, 5)
	row := 0
	col := 0
	for i := 4; i >= 0; i-- {
		row, col = 4 - i, i
		if debug {
			fmt.Printf("[/] col: %d, row: %d, val: %d\n", col, row, card[col][row])
		}
		if card[col][row] == 0 {
			line = append(line, [2]int{col, row})
		}
	}
	return line
}

func goalCol(card map[int][]int, col int) [][2]int {
	line := make([][2]int, 0, 5)
	for row := 0; row < 5; row++ {
		if debug {
			fmt.Printf("[|] col: %d, row: %d, val: %d\n", col, row, card[col][row])
		}
		if card[col][row] == 0 {
			line = append(line, [2]int{col, row})
		}
	}
	return line
}

func goalRow(card map[int][]int, row int) [][2]int {
	line := make([][2]int, 0, 5)
	for col := 0; col < 5; col++ {
		if debug {
			fmt.Printf("[-] col: %d, row: %d, val: %d\n", col, row, card[col][row])
		}
		if card[col][row] == 0 {
			line = append(line, [2]int{col, row})
		}
	}
	return line
}

func result(player int, card map[int][]int, line[][2]int) bool {
	switch len(line) {
	case 5:
		printColor("player: " + strconv.Itoa(player + 1) + ", Goal!!!\n", Red)
		gridGoal(card, line)
		return true
	case 4:
		printColor("player: " + strconv.Itoa(player + 1) + ", Reach!!\n", Yellow)
		gridReach(card, line)
	}
	return false
}

func main() {
	cards := make([]map[int][]int, 0, 2)
	p1_card := getCard()
	p2_card := getCard()
	printColor("player: 1 Start\n", Blue)
	grid(p1_card)
	printColor("player: 2 Start\n", Blue)
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
			hit, col, row := hit(target, card)
			if hit {
				printColor("player: " + strconv.Itoa(player + 1) + ", Hit!\n", Green)
				gridHit(card, col, row)
				if debug {
					fmt.Printf("player: %d, target: %d, hit: %t, row:%d, col:%d\n",
						player + 1, target, hit, col, row)
				}
			}
			if hit {
				hits := goalCol(card, col)
				if result(player, card, hits) { return }
				hits = goalRow(card, row)
				if result(player, card, hits) { return }
				hits = goalRightOblique(card)
				if result(player, card, hits) { return }
				hits = goalLeftOblique(card)
				if result(player, card, hits) { return }
			}
		}
	}
}

