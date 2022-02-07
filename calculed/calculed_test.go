package calculed

import "testing"

func TestCalculed1x1EmptyBoard(t *testing.T) {
	expected := int32(0)
	got := QueensAttack(1, 0, 1, 1, nil)

	if got != expected {
		t.Errorf("Esperado: %v, obtenido: %v", expected, got)
	}
}

func TestCalculed9x9EmptyBoard(t *testing.T) {
	expected := int32(32)
	got := QueensAttack(9, 0, 5, 5, nil)

	if got != expected {
		t.Errorf("Esperado: %v, obtenido: %v", expected, got)
	}
}

func TestCalculed3x3EmptyBoard(t *testing.T) {
	expected := int32(8)
	got := QueensAttack(3, 0, 2, 2, nil)

	if got != expected {
		t.Errorf("Esperado: %v, obtenido: %v", expected, got)
	}
}

func TestCalculed4x4EmptyBoard(t *testing.T) {
	expected := int32(9)
	got := QueensAttack(4, 0, 4, 4, nil)

	if got != expected {
		t.Errorf("Esperado: %v, obtenido: %v", expected, got)
	}
}
func TestCalculed5x5EmptyBoard(t *testing.T) {
	expected := int32(14)
	got := QueensAttack(5, 0, 4, 3, nil)

	if got != expected {
		t.Errorf("Esperado: %v, obtenido: %v", expected, got)
	}
}

func TestCalculed5x5Board(t *testing.T) {
	expected := int32(10)
	got := QueensAttack(5, 3, 4, 3, [][]int32{{5, 5}, {4, 2}, {2, 3}})

	if got != expected {
		t.Errorf("Esperado: %v, obtenido: %v", expected, got)
	}
}

func TestCalculed8x8Board(t *testing.T) {
	expected := int32(24)
	got := QueensAttack(8, 1, 4, 4, [][]int32{{3, 5}})

	if got != expected {
		t.Errorf("Esperado: %v, obtenido: %v", expected, got)
	}
}

func TestCalculed100x100Board(t *testing.T) {
	expected := int32(40)
	got := QueensAttack(100, 100, 48, 81, [][]int32{{64, 97}, {54, 87}, {42, 75}, {32, 65}, {42, 87}, {32, 97}, {54, 75}, {64, 65}, {48, 87}, {48, 75}, {54, 81}, {42, 81}, {45, 17}, {14, 24}, {35, 15}, {95, 64}, {63, 87}, {25, 72}, {71, 38}, {96, 97}, {16, 30}, {60, 34}, {31, 67}, {26, 82}, {20, 93}, {81, 38}, {51, 94}, {75, 41}, {79, 84}, {79, 65}, {76, 80}, {52, 87}, {81, 54}, {89, 52}, {20, 31}, {10, 41}, {32, 73}, {83, 98}, {87, 61}, {82, 52}, {80, 64}, {82, 46}, {49, 21}, {73, 86}, {37, 70}, {43, 12}, {94, 28}, {10, 93}, {52, 25}, {50, 61}, {52, 68}, {52, 23}, {60, 91}, {79, 17}, {93, 82}, {12, 18}, {75, 64}, {69, 69}, {94, 74}, {61, 61}, {46, 57}, {67, 45}, {96, 64}, {83, 89}, {58, 87}, {76, 53}, {79, 21}, {94, 70}, {16, 10}, {50, 82}, {92, 20}, {40, 51}, {49, 28}, {51, 82}, {35, 16}, {15, 86}, {78, 89}, {41, 98}, {70, 46}, {79, 79}, {24, 40}, {91, 13}, {59, 73}, {35, 32}, {40, 31}, {14, 31}, {71, 35}, {96, 18}, {27, 39}, {28, 38}, {41, 36}, {31, 63}, {52, 48}, {81, 25}, {49, 90}, {32, 65}, {25, 45}, {63, 94}, {89, 50}, {43, 41}})

	if got != expected {
		t.Errorf("Esperado: %v, obtenido: %v", expected, got)
	}
}

func TestCalculed100000x100000EmptyBoard(t *testing.T) {
	expected := int32(308369)
	got := QueensAttack(100000, 0, 4187, 5068, nil)

	if got != expected {
		t.Errorf("Esperado: %v, obtenido: %v", expected, got)
	}
}

func TestCalculateUpObstacles(t *testing.T) {
	expectedRow := int32(3)
	expectedColumn := int32(4)
	upRow, upColumn, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := calculateObstacles([][]int32{{3, 4}}, 2, 4)

	if expectedRow != upRow {
		t.Errorf("Row Esperado: %v, obtenido: %v", expectedRow, upRow)
	}

	if expectedColumn != upColumn {
		t.Errorf("Column Esperado: %v, obtenido: %v", expectedColumn, upColumn)
	}
}

func TestCalculateDownObstacles(t *testing.T) {
	expectedRow := int32(1)
	expectedColumn := int32(4)
	//upRow, upColumn, downRow, downColumn, leftRow, leftColumn, rightRow, rightColumn
	_, _, downRow, downColumn, _, _, _, _, _, _, _, _, _, _, _, _ := calculateObstacles([][]int32{{1, 4}}, 2, 4)

	if expectedRow != downRow {
		t.Errorf("Row Esperado: %v, obtenido: %v", expectedRow, downRow)
	}

	if expectedColumn != downColumn {
		t.Errorf("Column Esperado: %v, obtenido: %v", expectedColumn, downColumn)
	}
}

func TestCalculateLeftObstacles(t *testing.T) {
	expectedRow := int32(2)
	expectedColumn := int32(3)
	_, _, _, _, leftRow, leftColumn, _, _, _, _, _, _, _, _, _, _ := calculateObstacles([][]int32{{2, 3}}, 2, 4)

	if expectedRow != leftRow {
		t.Errorf("Row Esperado: %v, obtenido: %v", expectedRow, leftRow)
	}

	if expectedColumn != leftColumn {
		t.Errorf("Column Esperado: %v, obtenido: %v", expectedColumn, leftColumn)
	}
}

func TestCalculateRightObstacles(t *testing.T) {
	expectedRow := int32(2)
	expectedColumn := int32(5)
	_, _, _, _, _, _, rightRow, rightColumn, _, _, _, _, _, _, _, _ := calculateObstacles([][]int32{{2, 5}}, 2, 4)

	if expectedRow != rightRow {
		t.Errorf("Row Esperado: %v, obtenido: %v", expectedRow, rightRow)
	}

	if expectedColumn != rightColumn {
		t.Errorf("Column Esperado: %v, obtenido: %v", expectedColumn, rightColumn)
	}
}

func TestCalculateUpRightDiagonalObstacles(t *testing.T) {
	expectedRow := int32(3)
	expectedColumn := int32(5)
	_, _, _, _, _, _, _, _, _, _, upRightRow, upRightColumn, _, _, _, _ := calculateObstacles([][]int32{{3, 5}}, 2, 4)

	if expectedRow != upRightRow {
		t.Errorf("Row Esperado: %v, obtenido: %v", expectedRow, upRightRow)
	}

	if expectedColumn != upRightColumn {
		t.Errorf("Column Esperado: %v, obtenido: %v", expectedColumn, upRightColumn)
	}
}

func TestCalculateUpRightDiagonalDoubleObstacle(t *testing.T) {
	expectedRow := int32(3)
	expectedColumn := int32(3)
	_, _, _, _, _, _, _, _, _, _, upRightRow, upRightColumn, _, _, _, _ := calculateObstacles([][]int32{{5, 5}, {3, 3}}, 3, 3)

	if expectedRow != upRightRow {
		t.Errorf("Row Esperado: %v, obtenido: %v", expectedRow, upRightRow)
	}

	if expectedColumn != upRightColumn {
		t.Errorf("Column Esperado: %v, obtenido: %v", expectedColumn, upRightColumn)
	}
}

// func TestCalculateUpLeftDiagonalObstacles(t *testing.T) {
// 	expectedRow := int32(3)
// 	expectedColumn := int32(3)
// 	_, _, _, _, _, _, _, _, upLeftRow, upLeftColumn, _, _, _, _, _, _ := calculateObstacles([][]int32{{3, 3}}, 2, 4)

// 	if expectedRow != upLeftRow {
// 		t.Errorf("Row Esperado: %v, obtenido: %v", expectedRow, upLeftRow)
// 	}

// 	if expectedColumn != upLeftColumn {
// 		t.Errorf("Column Esperado: %v, obtenido: %v", expectedColumn, upLeftColumn)
// 	}
// }
