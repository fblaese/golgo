package main

import (
	"os/exec"
	"fmt"
	"os"
	"strconv"
	"strings"
	"log"
	"math/rand"
	"time"
	"bytes"
)

var field [][]bool
var newfield [][]bool
var width, height int
var generation int = 0

func main() {
	var err error

	// Get Terminal Size
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, _ := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	termsize := strings.Split(strings.TrimSuffix(string(out), "\n"), " ")
	width, err = strconv.Atoi(termsize[0])
	if err != nil {
		log.Fatal(err)
	}
	height, err = strconv.Atoi(termsize[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("width: %v, height: %v", width, height)

	rand.Seed(time.Now().UnixNano())

	field = make([][]bool, width)
	for col := range field {
		field[col] = make([]bool, height)
	}

	fmt.Printf("type: %T", field[0][0])
	for x := range field {
		for y := range field[x] {
			field[x][y] = rand.Float32() < 0.5
		}
	}

	newfield = make([][]bool, width)
	for col := range newfield {
		newfield[col] = make([]bool, height)
	}

	for {
		printField()
		for i := 0; i < 15; i++ {
			nextGen()
		}

		// swap fields
		fieldtmp := field
		field = newfield
		newfield = fieldtmp
	}
}

func countNeigh(x, y int) int {
	var count int = 0

	for xo := -1; xo <= 1; xo++ {
		for yo := -1; yo <= 1; yo++ {
			if xo == 0 && yo == 0 {
				continue
			}

			if x + xo >= 0 && y + yo >= 0 && x + xo < width && y + yo < height {
				if field[x+xo][y+yo] == true {
					count += 1
				}
			}
		}
	}

	return count
}

func nextGenCell(x, y int) {
	alive := field[x][y]
	neigh := countNeigh(x, y)

	if alive == false {
		if neigh == 3 {
			newfield[x][y] = true
		} else {
			newfield[x][y] = false
		}
	} else {
		if neigh < 2 || neigh > 3 {
			newfield[x][y] = false
		} else {
			newfield[x][y] = true
		}
	}
}

func nextGen() {
	for x := range field {
		for y := range field[x] {
			nextGenCell(x, y)
		}
	}
	generation++
}

func printField() {
    var buffer bytes.Buffer

	for x := range field {
		for y := range field[x] {
			if field[x][y] == false {
        buffer.WriteString(" ")
			} else {
        buffer.WriteString("■")
				//fmt.Print(countNeigh(x,y))
			}
		}
	}
	fmt.Print("\033[H\033[2J")
    fmt.Println(buffer.String())
	fmt.Printf("\rGeneration: %v", generation)
}
