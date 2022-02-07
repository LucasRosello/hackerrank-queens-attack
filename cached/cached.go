package cached

import "github.com/LucasRosello/queens-attack/simple"

var upRow, upColumn, downRow, downColumn, leftRow, leftColumn, rightRow, rightColumn int32

func QueensAttack(n int32, k int32, r_q int32, c_q int32, obstacles [][]int32) int32 {
	var sum int32

	upRow, upColumn, downRow, downColumn, _ = generateCachedObstacles(obstacles, r_q, c_q)
	sum = sum + getCachedDownPath(r_q, obstacles, c_q)
	sum = sum + getCachedUpPath(r_q, obstacles, c_q, n)
	sum = sum + simple.GetLeftPath(r_q, obstacles, c_q)
	sum = sum + simple.GetRightPath(r_q, obstacles, c_q, n)
	sum = sum + simple.GetUpRightDiagonalPath(n, c_q, r_q, obstacles)
	sum = sum + simple.GetUpLeftDiagonalPath(n, c_q, r_q, obstacles)
	sum = sum + simple.GetDownLeftDiagonalPath(n, c_q, r_q, obstacles)
	sum = sum + simple.GetDownRightDiagonalPath(n, c_q, r_q, obstacles)

	return sum
}

func getCachedUpPath(rowQueen int32, obstacles [][]int32, columnQueen int32, boardLength int32) int32 {
	var sum int32

	for i := rowQueen + 1; i <= boardLength; i++ {
		if i == upRow && columnQueen == upColumn {
			break
		}
		sum++
	}

	return sum
}

func getCachedDownPath(rowQueen int32, obstacles [][]int32, columnQueen int32) int32 {
	var sum int32

	for i := rowQueen - 1; i > 0; i-- {
		if i == downRow && columnQueen == downColumn {
			break
		}
		sum++
	}

	return sum
}

func generateCachedObstacles(obstacles [][]int32, r_q int32, c_q int32) (int32, int32, int32, int32, int32) {

	tempUpRow := int32(999999999)
	tempDownRow := int32(0)
	//tempRightColumnt := int32(999999999)
	tempLeftColumn := int32(0)

	for i := 0; i < len(obstacles); i++ {
		// obstacles[number_of_the_obstacle][axis] // 0 = row; 1 = column
		if obstacles[i][0] > r_q && obstacles[i][1] == c_q {
			if obstacles[i][0] < tempUpRow {
				tempUpRow = obstacles[i][0]
			}
		} else if obstacles[i][0] < r_q && obstacles[i][1] == c_q {
			if obstacles[i][0] > tempDownRow {
				tempDownRow = obstacles[i][0]
			}
		} else if obstacles[i][0] == r_q && obstacles[i][1] < c_q {
			if obstacles[i][1] > tempLeftColumn {
				tempLeftColumn = obstacles[i][1]
			}
		}
	}
	if tempUpRow == 999999999 {
		tempUpRow = 0
	}

	return tempUpRow, c_q, tempDownRow, c_q, tempLeftColumn
}
