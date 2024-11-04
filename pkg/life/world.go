package life

import (
	"errors"
	"math/rand"
	"time"
)

type World struct {
	Height int // Высота сетки
	Width  int // Ширина сетки
	Cells  [][]bool
}

// Используйте код из предыдущего урока по игре «Жизнь»
func NewWorld(height, width int) (*World, error) {
	if height <= 0 || width <= 0 {
		return &World{}, errors.New("Zero")
	}
	cells := make([][]bool, height)

	for i := range cells {
		cells[i] = make([]bool, width) // создаём новый слайс в каждой строке
	}
	return &World{
		Height: height,
		Width:  width,
		Cells:  cells,
	}, nil
}
func (w *World) next(x, y int) bool {
	n := w.neighbors(x, y)       // получим количество живых соседей
	alive := w.Cells[y][x]       // текущее состояние клетки
	if n < 4 && n > 1 && alive { // если соседей двое или трое, а клетка жива
		return true // то следующее состояние — жива
	}
	if n == 3 && !alive { // если клетка мертва, но у неё трое соседей
		return true // клетка оживает
	}

	return false
}
func (w *World) zero_zero() int {
	var n, a, b int
	if len(w.Cells) < 3 {
		a = len(w.Cells)
	} else {
		a = 2
	}
	if len(w.Cells[0]) < 3 {
		b = len(w.Cells[0])
	} else {
		b = 2
	}
	for i := 0; i < a; i++ {
		for j := 0; j < b; j++ {
			if i == 0 && j == 0 {
				continue
			} else if w.Cells[i][j] {
				n += 1
			}
		}
	}
	return n
}
func (w *World) zero_end() int {
	var n, a, b int
	if len(w.Cells) < 3 {
		a = len(w.Cells)
	} else {
		a = 2
	}
	b = len(w.Cells[0]) - 2
	for i := 0; i < a; i++ {
		for j := b; j < b+2; j++ {
			if i == 0 && j == b+1 {
				continue
			} else if w.Cells[i][j] {
				n += 1
			}
		}
	}
	return n
}
func (w *World) zero_central(b int) int {
	var n, a int
	if len(w.Cells) < 3 {
		a = len(w.Cells)
	} else {
		a = 2
	}
	for i := 0; i < a; i++ {
		for j := b - 1; j < b+2; j++ {
			if i == 0 && j == b {
				continue
			} else if w.Cells[i][j] {
				n += 1
			}
		}
	}
	return n
}
func (w *World) end_zero() int {
	var n, a, b int
	a = len(w.Cells) - 2
	if len(w.Cells[0]) < 3 {
		b = len(w.Cells[0])
	} else {
		b = 2
	}
	for i := a; i < a+2; i++ {
		for j := 0; j < b; j++ {
			if i == a+1 && j == 0 {
				continue
			} else if w.Cells[i][j] {
				n += 1
			}
		}
	}
	return n
}
func (w *World) end_end() int {
	var n, a, b int
	a = len(w.Cells) - 2
	b = len(w.Cells[0]) - 2
	for i := a; i < a+2; i++ {
		for j := b; j < b+2; j++ {
			if i == a+1 && j == b+1 {
				continue
			} else if w.Cells[i][j] {
				n += 1
			}
		}
	}
	return n
}
func (w *World) end_central(b int) int {
	var n, a int
	a = len(w.Cells) - 2
	for i := a; i < a+2; i++ {
		for j := b - 1; j < b+2; j++ {
			if i == a+1 && j == b {
				continue
			} else if w.Cells[i][j] {
				n += 1
			}
		}
	}
	return n
}
func (w *World) central_zero(a int) int {
	var n, b int
	if len(w.Cells[0]) < 3 {
		b = len(w.Cells[0])
	} else {
		b = 2
	}
	for i := a - 1; i < a+2; i++ {
		for j := 0; j < b; j++ {
			if i == a && j == 0 {
				continue
			} else if w.Cells[i][j] {
				n += 1
			}
		}
	}
	return n
}
func (w *World) central_end(a int) int {
	var n, b int
	b = len(w.Cells[0]) - 2
	for i := a - 1; i < a+2; i++ {
		for j := b; j < b+2; j++ {
			if i == a && j == b+1 {
				continue
			} else if w.Cells[i][j] {
				n += 1
			}
		}
	}
	return n
}
func (w *World) central_central(a, b int) int {
	var n int
	for i := a - 1; i < a+2; i++ {
		for j := b - 1; j < b+2; j++ {
			if i == a && j == b {
				continue
			} else if w.Cells[i][j] {
				n += 1
			}
		}
	}
	return n
}
func (w *World) neighbors(x, y int) int {
	var n int
	if y == 0 {
		if x == 0 {
			n = w.zero_zero()
		} else if x == len(w.Cells[0])-1 {
			n = w.zero_end()
		} else {
			n = w.zero_central(x)
		}
	} else if y == len(w.Cells)-1 {
		if x == 0 {
			n = w.end_zero()
		} else if x == len(w.Cells[0])-1 {
			n = w.end_end()
		} else {
			n = w.end_central(x)
		}
	} else {
		if x == 0 {
			n = w.central_zero(y)
		} else if x == len(w.Cells[0])-1 {
			n = w.central_end(y)
		} else {
			n = w.central_central(y, x)
		}
	}

	return n
}

func NextState(oldWorld, newWorld World) {
	for i := 0; i < oldWorld.Height; i++ {
		for j := 0; j < oldWorld.Width; j++ {
			// для каждой клетки получим новое состояние
			newWorld.Cells[i][j] = oldWorld.next(j, i)
		}
	}
}

// RandInit заполняет поля на указанное число процентов
func (w *World) RandInit(percentage int) {
	// Количество живых клеток
	numAlive := percentage * w.Height * w.Width / 100
	// Заполним живыми первые клетки
	w.fillAlive(numAlive)
	// Получаем рандомные числа
	r := rand.New(rand.NewSource(time.
		Now().Unix()))

	// Рандомно меняем местами
	for i := 0; i < w.Height*w.Width; i++ {
		randRowLeft := r.Intn(w.Width)
		randColLeft := r.Intn(w.Height)
		randRowRight := r.Intn(w.Width)
		randColRight := r.Intn(w.Height)

		w.Cells[randRowLeft][randColLeft] = w.Cells[randRowRight][randColRight]
	}
}

func (w *World) fillAlive(num int) {
	aliveCount := 0
	for j, row := range w.Cells {
		for k := range row {
			w.Cells[j][k] = true
			aliveCount++
			if aliveCount == num {

				return
			}
		}
	}
}
