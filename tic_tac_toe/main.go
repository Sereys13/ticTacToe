package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

/*
Функция конкатенации строк

@param str1, str2 строки для конкатенации
@return результат конкатенации
*/
func ConcatStrings(str1, str2 string) string {
	return str1 + str2
}

/*
Функция рандомного хода компьютера

@param data[][] игровое поле
@param name имя за кого играет компьютер (крестик или нолик)
*/
func RandomMove(data [][]string, name *string) {
	for {
		rand.Seed(time.Now().UnixNano())
		rowV := rand.Intn(3)
		colV := rand.Intn(3)
		if data[rowV][colV] == "" {
			data[rowV][colV] = *name
			break
		}
	}
}

/*
Функция победного хода компьютера

@param data[][] игровое поле
@param name имя за кого играет компьютер (крестик или нолик)
@return true если есть возможность сделать победный ход, иначе false
*/
func MoveWinComputer(data [][]string, name *string) bool {
	var count0_0, count0_2 int
	for i := 0; i < 3; i++ {
		var countRow, countCol int
		for j := 0; j < 3; j++ {
			if data[i][j] == *name {
				countRow++
			} else if data[i][j] != "" {
				countRow--
			}
			if data[j][i] == *name {
				countCol++
			} else if data[j][i] != "" {
				countCol--
			}
			if i == j && data[i][j] == *name {
				count0_0++
			} else if i == j && data[j][i] != "" {
				count0_0--
			}
			if (i == 0 && j == 2) || (i == 2 && j == 0) || (i == 1 && j == 1) {
				if data[i][j] == *name {
					count0_2++
				} else if data[i][j] != "" {
					count0_2--
				}
			}
		}
		if countRow == 2 {
			for j := 0; j < 3; j++ {
				if data[i][j] == "" {
					data[i][j] = *name
					return true
				}
			}
		}
		if countCol == 2 {
			for j := 0; j < 3; j++ {
				if data[j][i] == "" {
					data[j][i] = *name
					return true
				}
			}
		}
	}
	if count0_0 == 2 {
		for i := 0; i < 3; i++ {
			if data[i][i] == "" {
				data[i][i] = *name
				return true
			}
		}
	}
	if count0_2 == 2 {
		j := 2
		for i := 0; i < 3; i++ {
			if data[i][j] == "" {
				data[i][j] = *name
				return true
			}
			j--
		}
	}
	return false
}

/*
Функция подводящего к победе хода компьютера

@param data[][] игровое поле
@param name имя за кого играет компьютер (крестик или нолик)
@param namePlayer имя за кого играет игрок (крестик или нолик)
*/
func MoveAttackComputer(data [][]string, name, namePlayer *string) {
	found := false
	for i := 0; i < 3 && !found; i++ {
		for j := 0; j < 3 && !found; j++ {
			var ind1, ind2 int
			switch {
			case j == 0:
				ind1, ind2 = 1, 2
			case j == 1:
				ind1, ind2 = 0, 2
			case j == 2:
				ind1, ind2 = 0, 1
			}
			if data[i][j] == *name {
				if data[i][ind1] == "" && data[i][ind2] != *namePlayer {
					data[i][ind1] = *name
					found = true
					break
				} else if data[i][ind2] == "" && data[i][ind1] != *namePlayer {
					data[i][ind2] = *name
					found = true
					break
				}
			}
			if data[j][i] == *name {
				if data[ind1][i] == "" && data[ind2][i] != *namePlayer {
					data[ind1][i] = *name
					found = true
					break
				} else if data[ind2][i] == "" && data[ind1][i] != *namePlayer {
					data[ind2][i] = *name
					found = true
					break
				}
			}
			if i == j && data[i][j] == *name {
				if data[ind1][ind1] == "" && data[ind2][ind2] != *namePlayer {
					data[ind1][ind1] = *name
					found = true
				} else if data[ind2][ind2] == "" && data[ind1][ind1] != *namePlayer {
					data[ind2][ind2] = *name
					found = true
				}
			}
		}
	}
	if !found {
		RandomMove(data, name)
	}
}

/*
Функция хода компьютера не допускающего победы игрока

@param data[][] игровое поле
@param x,y координаты последнего хода игрока
@param name1 имя за кого играет игрок (крестик или нолик)
@param name2 имя за кого играет компьютер (крестик или нолик)
*/
func MoveProtectiveComputer(data [][]string, x, y *int, name1, name2 *string) {
	var ind1, ind2, ind3, ind4 int
	switch {
	case *y == 0:
		ind1, ind2 = 1, 2
	case *y == 1:
		ind1, ind2 = 0, 2
	case *y == 2:
		ind1, ind2 = 0, 1
	}
	switch {
	case *x == 0:
		ind3, ind4 = 1, 2
	case *x == 1:
		ind3, ind4 = 0, 2
	case *x == 2:
		ind3, ind4 = 0, 1
	}
	if data[*x][ind1] == *name1 && data[*x][ind2] == "" {
		data[*x][ind2] = *name2
	} else if data[*x][ind2] == *name1 && data[*x][ind1] == "" {
		data[*x][ind1] = *name2
	} else if data[ind3][*y] == *name1 && data[ind4][*y] == "" {
		data[ind4][*y] = *name2
	} else if data[ind4][*y] == *name1 && data[ind3][*y] == "" {
		data[ind3][*y] = *name2
	} else if (*x == 0 && *y == 0) || (*x == 1 && *y == 1) || (*x == 2 && *y == 2) {
		switch {
		case data[ind1][ind3] == *name1 && data[ind2][ind4] == "":
			data[ind2][ind4] = *name2
		case data[ind2][ind4] == *name1 && data[ind1][ind3] == "":
			data[ind1][ind3] = *name2
		case data[ind1][ind4] == *name1 && data[ind2][ind3] == "":
			data[ind2][ind3] = *name2
		case data[ind2][ind3] == *name1 && data[ind1][ind4] == "":
			data[ind1][ind4] = *name2
		default:
			MoveAttackComputer(data, name2, name1)
		}
	} else if (*x == 0 && *y == 2) || (*x == 2 && *y == 0) {
		if *y == 2 {
			ind1, ind2 = ind2, ind1
		} else {
			ind3, ind4 = ind4, ind3
		}
		switch {
		case data[ind3][ind1] == *name1 && data[ind4][ind2] == "":
			data[ind4][ind2] = *name2
		case data[ind4][ind2] == *name1 && data[ind3][ind1] == "":
			data[ind3][ind1] = *name2
		default:
			MoveAttackComputer(data, name2, name1)
		}
	} else {
		MoveAttackComputer(data, name2, name1)
	}
}

/*
Функция проверяющая победу игрока

@param data[][] игровое поле
@param name имя за кого играет игрок (крестик или нолик)
@return true если игрок победил, иначе false
*/
func CheckVictory(data [][]string, name *string) bool {
	switch {
	case data[0][0] == *name && data[0][1] == *name && data[0][2] == *name:
		return true
	case data[0][0] == *name && data[1][1] == *name && data[2][2] == *name:
		return true
	case data[0][0] == *name && data[1][0] == *name && data[2][0] == *name:
		return true
	case data[2][0] == *name && data[1][1] == *name && data[0][2] == *name:
		return true
	case data[0][2] == *name && data[1][2] == *name && data[2][2] == *name:
		return true
	case data[2][0] == *name && data[2][1] == *name && data[2][2] == *name:
		return true
	case data[1][0] == *name && data[1][1] == *name && data[1][2] == *name:
		return true
	case data[0][1] == *name && data[1][1] == *name && data[2][1] == *name:
		return true
	default:
		return false
	}
}

func main() {
	data := [][]string{make([]string, 3), make([]string, 3), make([]string, 3)}
	var namePlayer, nameRival string
	var count int
	a := app.New()
	w := a.NewWindow("Tic-Tac-Toe")
	b := a.NewWindow("GameField")
	w.Resize(fyne.NewSize(400, 400))
	b.Resize(fyne.NewSize(300, 130))
	var contentA, contentB, contentC, contentD *fyne.Container
	informationLine := widget.NewLabel("")
	input := widget.NewEntry()
	input.SetPlaceHolder("Добро пожаловать!Введите ваше имя")

	gameField := widget.NewTable(
		func() (int, int) {
			return len(data), len(data[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("Default test")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			if data[i.Row][i.Col] == "" {
				o.(*widget.Label).SetText(fmt.Sprintf("Rof: %d, Col: %d", i.Row, i.Col))
			} else if data[i.Row][i.Col] == "X" {
				o.(*widget.Label).SetText("         X")
			} else {
				o.(*widget.Label).SetText("         O")
			}
		})

	buttonName := widget.NewButton("Enter", func() {
		if input.Text != "" {
			informationLine.SetText(ConcatStrings(input.Text, " выберите за кого будете играть"))
			w.SetContent(contentB)
			w.Show()
		} else {
			input.SetPlaceHolder("Чтобы продолжить введите имя")
		}
	})

	buttonX := widget.NewButton("X", func() {
		namePlayer = "X"
		nameRival = "O"
		informationLine.SetText(ConcatStrings(input.Text, " вы ходите первыми"))
		w.SetContent(contentC)
		b.SetContent(gameField)
		b.Show()
		w.Show()
	})

	buttonY := widget.NewButton("O", func() {
		namePlayer = "O"
		nameRival = "X"
		informationLine.SetText("Вы играете за нолики. Компьютер ходит первым")
		w.SetContent(contentC)
		b.SetContent(gameField)
		b.Show()
		w.Show()
		time.Sleep(1 * time.Second)
		RandomMove(data, &nameRival)
		count = 1
		informationLine.SetText("Компьютер сделал свой ход. Теперь ваш ход ")
		b.SetContent(gameField)
		b.Show()
	})

	row := widget.NewEntry()
	row.SetPlaceHolder("Введите номер ряда(row)от 0 до 2...")
	col := widget.NewEntry()
	col.SetPlaceHolder("Введите номер колонки(col)от 0 до 2...")

	sel := widget.NewButton("Подсветить будущий ход", func() {
		rowV, _ := strconv.Atoi(row.Text)
		colV, _ := strconv.Atoi(col.Text)
		gameField.Select(widget.TableCellID{
			Row: rowV,
			Col: colV,
		})
	})

	play := widget.NewButton("Потвердить выбор", func() {
		moveWin := false
		rowV, _ := strconv.Atoi(row.Text)
		colV, _ := strconv.Atoi(col.Text)
		if rowV > 2 || rowV < 0 || colV > 2 || colV < 0 {
			informationLine.SetText("Некоректный ввод, повторите попытку")
			row.SetText("")
			col.SetText("")
		} else if data[rowV][colV] == "" {
			count++
			data[rowV][colV] = namePlayer
			b.SetContent(gameField)
			b.Show()
			if count >= 3 && CheckVictory(data, &namePlayer) {
				moveWin = true
				informationLine.SetText("Вы победили!!!")
				w.SetContent(contentD)
				w.Show()
			} else if (count == 5 && namePlayer == "X") || (count == 5 && namePlayer == "O" && !MoveWinComputer(data, &nameRival)) {
				moveWin = true
				informationLine.SetText("Победила дружба")
				w.SetContent(contentD)
				w.Show()
			} else {
				row.SetText("")
				col.SetText("")
				switch {
				case count == 1:
					RandomMove(data, &nameRival)
				case count == 2 || ((count > 2 && count != 5) && !MoveWinComputer(data, &nameRival)):
					MoveProtectiveComputer(data, &rowV, &colV, &namePlayer, &nameRival)
				default:
					moveWin = true
					informationLine.SetText("Вы проиграли компьютеру, в следующий раз обязательно повезет!!!")
					w.SetContent(contentD)
					w.Show()
				}
			}
			if !moveWin {
				b.SetContent(gameField)
				b.Show()
				informationLine.SetText(ConcatStrings("Компьютер сделал свой ход. Теперь ваш ход ", input.Text))
			}
		} else {
			informationLine.SetText("Такой ход уже сделан, повторите попытку")
		}
	})

	buttonClose := widget.NewButton("Выйти из игры", func() {
		informationLine.SetText("До встречи!!!")
		time.Sleep(1 * time.Second)
		b.Close()
		w.Close()
	})

	buttonRestart := widget.NewButton("Попробовать снова", func() {
		data = [][]string{make([]string, 3), make([]string, 3), make([]string, 3)}
		row.SetText("")
		col.SetText("")
		if namePlayer == "X" {
			informationLine.SetText(ConcatStrings(input.Text, " вы ходите первыми"))
			count = 0
		} else {
			RandomMove(data, &nameRival)
			informationLine.SetText(ConcatStrings("Компьютер сделал свой ход. Теперь ваш ход ", input.Text))
			count = 1
		}
		w.SetContent(contentC)
		b.SetContent(gameField)
		w.Show()
		b.Show()
	})

	contentA = fyne.NewContainerWithLayout(layout.NewVBoxLayout(), input, buttonName)
	contentB = fyne.NewContainerWithLayout(layout.NewVBoxLayout(), informationLine, buttonX, buttonY)
	contentC = fyne.NewContainerWithLayout(layout.NewVBoxLayout(), informationLine, row, col, sel, play)
	contentD = fyne.NewContainerWithLayout(layout.NewVBoxLayout(), informationLine, buttonRestart, buttonClose)

	w.SetContent(contentA)
	w.ShowAndRun()
}
