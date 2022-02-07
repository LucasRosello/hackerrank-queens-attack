package simple

func QueensAttack(n int32, k int32, r_q int32, c_q int32, obstacles [][]int32) int32 {
	var sum int32

	sum = sum + GetDownPath(r_q, obstacles, c_q)
	sum = sum + GetUpPath(r_q, obstacles, c_q, n)
	sum = sum + GetLeftPath(r_q, obstacles, c_q)
	sum = sum + GetRightPath(r_q, obstacles, c_q, n)
	sum = sum + GetUpRightDiagonalPath(n, c_q, r_q, obstacles)
	sum = sum + GetUpLeftDiagonalPath(n, c_q, r_q, obstacles)
	sum = sum + GetDownLeftDiagonalPath(n, c_q, r_q, obstacles)
	sum = sum + GetDownRightDiagonalPath(n, c_q, r_q, obstacles)

	return sum
}

func GetDownPath(rowQueen int32, obstacles [][]int32, columnQueen int32) int32 {
	var sum int32

	for i := rowQueen - 1; i > 0; i-- {
		if obstacle(obstacles, i, columnQueen) {
			break
		}
		sum++
	}

	return sum
}

func GetUpPath(rowQueen int32, obstacles [][]int32, columnQueen int32, boardLength int32) int32 {
	var sum int32

	for i := rowQueen + 1; i <= boardLength; i++ {
		if obstacle(obstacles, i, columnQueen) {
			break
		}
		sum++
	}

	return sum
}

func GetLeftPath(rowQueen int32, obstacles [][]int32, columnQueen int32) int32 {
	var sum int32

	for i := columnQueen - 1; i > 0; i-- {
		if obstacle(obstacles, rowQueen, i) {
			break
		}
		sum++
	}

	return sum
}

func GetRightPath(rowQueen int32, obstacles [][]int32, columnQueen int32, boardLength int32) int32 {
	var sum int32

	for i := columnQueen + 1; i <= boardLength; i++ {
		if obstacle(obstacles, rowQueen, i) {
			break
		}
		sum++
	}

	return sum
}

func GetUpRightDiagonalPath(boardLength int32, columnQueen int32, rowQueen int32, obstacles [][]int32) int32 {
	var sum int32

	actualColumn := columnQueen
	actualRow := rowQueen

	for {
		actualColumn++
		actualRow++

		if actualColumn > boardLength || actualRow > boardLength {
			break
		}

		if obstacle(obstacles, actualRow, actualColumn) {
			break
		}
		sum++
	}

	return sum
}

func GetUpLeftDiagonalPath(boardLength int32, columnQueen int32, rowQueen int32, obstacles [][]int32) int32 {
	var sum int32

	actualColumn := columnQueen
	actualRow := rowQueen

	for {
		actualColumn--
		actualRow++

		if actualColumn <= 0 || actualRow > boardLength {
			break
		}

		if obstacle(obstacles, actualRow, actualColumn) {
			break
		}
		sum++
	}

	return sum
}

func GetDownLeftDiagonalPath(boardLength int32, columnQueen int32, rowQueen int32, obstacles [][]int32) int32 {
	var sum int32

	actualColumn := columnQueen
	actualRow := rowQueen

	for {
		actualColumn--
		actualRow--

		if actualColumn <= 0 || actualRow <= 0 {
			break
		}

		if obstacle(obstacles, actualRow, actualColumn) {
			break
		}
		sum++
	}

	return sum
}

func GetDownRightDiagonalPath(boardLength int32, columnQueen int32, rowQueen int32, obstacles [][]int32) int32 {
	var sum int32

	actualColumn := columnQueen
	actualRow := rowQueen

	for {
		actualColumn++
		actualRow--

		if actualColumn > boardLength || actualRow <= 0 {
			break
		}

		if obstacle(obstacles, actualRow, actualColumn) {
			break
		}
		sum++
	}

	return sum
}

func obstacle(obstacles [][]int32, row int32, column int32) bool {
	for i := 0; i < len(obstacles); i++ {
		if obstacles[i][0] == row && obstacles[i][1] == column {
			return true
		}

	}

	return false
}
