package calculed

var upRow, upColumn, downRow, downColumn, leftRow, leftColumn, rightRow, rightColumn, upLeftRow, upLeftColumn, upRightRow, upRightColumn, downLeftRow, downLeftColumn, downRightRow, downRightColumn int32

func QueensAttack(n int32, k int32, r_q int32, c_q int32, obstacles [][]int32) int32 {
	var sum int32

	upRow, upColumn, downRow, downColumn, leftRow, leftColumn, rightRow, rightColumn, upLeftRow, upLeftColumn, upRightRow, upRightColumn, downLeftRow, downLeftColumn, downRightRow, downRightColumn = calculateObstacles(obstacles, r_q, c_q)

	sum = sum + getCalculedDownPath(r_q)
	sum = sum + getCalculedUpPath(r_q, n)
	sum = sum + getCalculedLeftPath(c_q)
	sum = sum + getCalculedRightPath(c_q, n)
	//sum = sum + getUpRightDiagonalPath(c_q, n, r_q)
	sum = sum + GetUpRightDiagonalPath(n, c_q, r_q, obstacles)
	sum = sum + GetUpLeftDiagonalPath(n, c_q, r_q, obstacles)
	sum = sum + GetDownLeftDiagonalPath(n, c_q, r_q, obstacles)
	sum = sum + GetDownRightDiagonalPath(n, c_q, r_q, obstacles)

	return sum
}

func calculateObstacles(obstacles [][]int32, queenRow int32, queenColumn int32) (int32, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32) {

	upRow := int32(999999999)
	downRow := int32(0)
	leftColumn := int32(0)
	rightColumn := int32(999999999)

	// Diagonals
	upRightRow := int32(999999999)
	upRightColumn := int32(999999999)

	upLeftRow := int32(999999999)
	upLeftColumn := int32(0)

	for i := 0; i < len(obstacles); i++ {
		actualRow := obstacles[i][0]
		actualColumn := obstacles[i][1]

		// obstacles[number_of_the_obstacle][axis] // 0 = row; 1 = column
		if actualRow > queenRow && actualColumn == queenColumn {
			if actualRow < upRow {
				upRow = actualRow
			}
		} else if actualRow < queenRow && actualColumn == queenColumn {
			if actualRow > downRow {
				downRow = actualRow
			}
		} else if actualRow == queenRow && actualColumn < queenColumn {
			if actualColumn > leftColumn {
				leftColumn = actualColumn
			}
		} else if actualRow == queenRow && actualColumn > queenColumn {
			if actualColumn < rightColumn {
				rightColumn = actualColumn
			}
		} else /*diagonals*/ if actualRow >= queenRow && actualColumn >= queenColumn && (queenRow-actualRow) == (queenColumn-actualColumn) {
			if actualColumn < rightColumn {
				upRightRow = actualRow
				upRightColumn = actualColumn
			}
		}
		//  else if actualRow >= queenRow && actualColumn >= queenColumn && actualRow-queenRow == queenColumn-actualColumn {
		// 	if actualRow < upLeftRow {
		// 		upLeftRow = actualRow
		// 		upLeftColumn = actualColumn
		// 	}
		// }
	}

	if upRow == 999999999 {
		upRow = 0
	}
	if rightColumn == 999999999 {
		rightColumn = 0
	}
	if upRightRow == 999999999 {
		upRightRow = 0
		upRightColumn = 0
	}
	if upLeftRow == 999999999 {
		upLeftRow = 0
	}

	return upRow, queenColumn, downRow, queenColumn, queenRow, leftColumn, queenRow, rightColumn, upLeftRow, upLeftColumn, upRightRow, upRightColumn, 0, 0, 0, 0
}

func getCalculedUpPath(rowQueen int32, boardLength int32) int32 {
	if upRow == 0 {
		return int32(boardLength - rowQueen)
	}
	return int32(upRow - rowQueen - 1)
}

func getCalculedDownPath(rowQueen int32) int32 {
	return int32(rowQueen - downRow - 1)
}

func getCalculedLeftPath(columnQueen int32) int32 {
	return int32(columnQueen - leftColumn - 1)
}

func getCalculedRightPath(columnQueen int32, boardLength int32) int32 {
	if rightColumn == 0 {
		return int32(boardLength - columnQueen)
	}
	return int32(rightColumn - columnQueen - 1)
}

func getUpRightDiagonalPath(columnQueen int32, boardLength int32, rowQueen int32) int32 {
	if upRightRow == 0 {
		columnMinimun := int32(boardLength - columnQueen)
		rowMinimun := int32(boardLength - rowQueen)
		if rowMinimun < columnMinimun {
			return rowMinimun
		} else {
			return columnMinimun
		}
	}

	if upRightRow < upRightColumn {
		return int32(upRightRow - rowQueen - 1)
	} else {
		return int32(upRightColumn - columnQueen - 1)
	}

}

/* CACA */

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
