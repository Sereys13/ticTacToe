package main

import "testing"

//Тесты функции MoveWinComputer
func TestMoveWinComputer1(t *testing.T) {
	data := [][]string{{"O", "O", ""}, make([]string, 3), make([]string, 3)}
	expected := "O"
	MoveWinComputer(data, &expected)
	if expected != data[0][2] {
		t.Errorf("TestMoveWinComputer1 failed")
	}
}

func TestMoveWinComputer2(t *testing.T) {
	data := [][]string{{"O", "O", "X"}, make([]string, 3), make([]string, 3)}
	expected := "X"
	MoveWinComputer(data, &expected)
	if expected != data[0][2] {
		t.Errorf("TestMoveWinComputer2 failed")
	}
}

func TestMoveWinComputer3(t *testing.T) {
	data := [][]string{{"O", "X", "X"}, make([]string, 3), {"O", "X", ""}}
	expected := "O"
	MoveWinComputer(data, &expected)
	if expected != data[1][0] {
		t.Errorf("TestMoveWinComputer3 failed")
	}
}

func TestMoveWinComputer4(t *testing.T) {
	data := [][]string{{"O", "X", "X"}, make([]string, 3), {"", "X", "O"}}
	expected := "O"
	MoveWinComputer(data, &expected)
	if expected != data[1][1] {
		t.Errorf("TestMoveWinComputer4 failed")
	}
}

func TestMoveWinComputer5(t *testing.T) {
	data := [][]string{{"", "", "O"}, {"", "O", ""}, {"", "X", ""}}
	expected := "O"
	MoveWinComputer(data, &expected)
	if expected != data[2][0] {
		t.Errorf("TestMoveWinComputer5 failed")
	}
}

func TestMoveWinComputer6(t *testing.T) {
	data := [][]string{{"O", "", ""}, {"", "O", ""}, {"", "", ""}}
	expected := "O"
	MoveWinComputer(data, &expected)
	if expected != data[2][2] {
		t.Errorf("TestMoveWinComputer6 failed")
	}
}

func TestMoveWinComputer7(t *testing.T) {
	data := [][]string{{"", "O", ""}, {"", "", ""}, {"", "O", "O"}}
	expected := "O"
	MoveWinComputer(data, &expected)
	if expected != data[1][1] && expected != data[2][0] {
		t.Errorf("TestMoveWinComputer7 failed")
	}
}

//Тесты функции MoveProtectiveComputer
func TestMoveProtectiveComputer1(t *testing.T) {
	data := [][]string{{"X", "X", ""}, {"", "", ""}, {"", "", ""}}
	namePlayer := "X"
	expected := "O"
	var x, y int
	x, y = 0, 1
	MoveProtectiveComputer(data, &x, &y, &namePlayer, &expected)
	if expected != data[0][2] {
		t.Errorf("TestMoveProtectiveComputer1 failed")
	}
}

func TestMoveProtectiveComputer2(t *testing.T) {
	data := [][]string{{"X", "", "X"}, {"", "", ""}, {"", "", ""}}
	namePlayer := "X"
	expected := "O"
	var x, y int
	MoveProtectiveComputer(data, &x, &y, &namePlayer, &expected)
	if expected != data[0][1] {
		t.Errorf("TestMoveProtectiveComputer2 failed")
	}
}

func TestMoveProtectiveComputer3(t *testing.T) {
	data := [][]string{{"X", "", ""}, {"", "", ""}, {"", "", "X"}}
	namePlayer := "X"
	expected := "O"
	var x, y int
	MoveProtectiveComputer(data, &x, &y, &namePlayer, &expected)
	if expected != data[1][1] {
		t.Errorf("TestMoveProtectiveComputer3 failed")
	}
}

func TestMoveProtectiveComputer4(t *testing.T) {
	data := [][]string{{"", "", ""}, {"", "X", ""}, {"X", "", ""}}
	namePlayer := "X"
	expected := "O"
	var x, y int
	x, y = 2, 0
	MoveProtectiveComputer(data, &x, &y, &namePlayer, &expected)
	if expected != data[0][2] {
		t.Errorf("TestMoveProtectiveComputer4 failed")
	}
}

func TestMoveProtectiveComputer5(t *testing.T) {
	data := [][]string{{"", "", ""}, {"", "X", ""}, {"", "X", ""}}
	namePlayer := "X"
	expected := "O"
	var x, y int
	x, y = 2, 1
	MoveProtectiveComputer(data, &x, &y, &namePlayer, &expected)
	if expected != data[0][1] {
		t.Errorf("TestMoveProtectiveComputer5 failed")
	}
}

func TestMoveProtectiveComputer6(t *testing.T) {
	data := [][]string{{"", "", ""}, {"", "X", ""}, {"", "", "X"}}
	namePlayer := "X"
	expected := "O"
	var x, y int
	x, y = 2, 2
	MoveProtectiveComputer(data, &x, &y, &namePlayer, &expected)
	if expected != data[0][0] {
		t.Errorf("TestMoveProtectiveComputer6 failed")
	}
}

func TestMoveProtectiveComputer7(t *testing.T) {
	data := [][]string{{"", "X", ""}, {"", "", "O"}, {"", "", "X"}}
	namePlayer := "X"
	expected := "O"
	var x, y int
	x, y = 2, 2
	MoveProtectiveComputer(data, &x, &y, &namePlayer, &expected)
	if expected != data[1][1] && expected != data[1][0] {
		t.Errorf("TestMoveProtectiveComputer7 failed")
	}
}

//Тесты функции MoveAttackComputer
func TestMoveAttackComputer1(t *testing.T) {
	data := [][]string{{"O", "X", ""}, {"X", "", ""}, {"", "", ""}}
	namePlayer := "X"
	expected := "O"
	MoveAttackComputer(data, &expected, &namePlayer)
	if expected != data[1][1] && expected != data[2][2] {
		t.Errorf("TestMoveAttackComputer1 failed")
	}
}

func TestMoveAttackComputer2(t *testing.T) {
	data := [][]string{{"", "X", ""}, {"X", "O", ""}, {"", "", ""}}
	namePlayer := "X"
	expected := "O"
	MoveAttackComputer(data, &expected, &namePlayer)
	if expected != data[0][0] && expected != data[2][2] {
		t.Errorf("TestMoveAttackComputer2 failed")
	}
}

func TestMoveAttackComputer3(t *testing.T) {
	data := [][]string{{"X", "X", ""}, {"", "O", ""}, {"", "", ""}}
	namePlayer := "X"
	expected := "O"
	MoveAttackComputer(data, &expected, &namePlayer)
	if expected != data[1][0] && expected != data[1][2] {
		t.Errorf("TestMoveAttackComputer3 failed")
	}
}

func TestMoveAttackComputer4(t *testing.T) {
	data := [][]string{{"", "", ""}, {"", "O", "X"}, {"", "", "X"}}
	namePlayer := "X"
	expected := "O"
	MoveAttackComputer(data, &expected, &namePlayer)
	if expected != data[0][1] && expected != data[2][1] {
		t.Errorf("TestMoveAttackComputer4 failed")
	}
}

func TestMoveAttackComputer5(t *testing.T) {
	data := [][]string{{"X", "", "O"}, {"", "", ""}, {"", "", ""}}
	namePlayer := "X"
	expected := "O"
	MoveAttackComputer(data, &expected, &namePlayer)
	if expected != data[1][2] && expected != data[2][2] {
		t.Errorf("TestMoveAttackComputer5 failed")
	}
}

func TestMoveAttackComputer6(t *testing.T) {
	data := [][]string{{"", "X", ""}, {"", "", ""}, {"", "O", ""}}
	namePlayer := "X"
	expected := "O"
	MoveAttackComputer(data, &expected, &namePlayer)
	if expected != data[2][0] && expected != data[2][2] {
		t.Errorf("TestMoveAttackComputer6 failed")
	}
}

func TestMoveAttackComputer7(t *testing.T) {
	data := [][]string{{"O", "X", ""}, {"", "X", "X"}, {"X", "", "X"}}
	namePlayer := "X"
	expected := "O"
	MoveAttackComputer(data, &expected, &namePlayer)
	if expected != data[0][2] && expected != data[1][0] && expected != data[2][1] {
		t.Errorf("TestMoveAttackComputer7 failed")
	}
}
