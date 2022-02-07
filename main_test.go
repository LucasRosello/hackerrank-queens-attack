package main

import (
	"testing"

	"github.com/LucasRosello/queens-attack/cached"
	"github.com/LucasRosello/queens-attack/calculed"
	"github.com/LucasRosello/queens-attack/simple"
)

func BenchmarkSimpleObstacles(b *testing.B) {
	for i := 0; i < b.N; i++ {
		simple.QueensAttack(20, 12, 9, 9, [][]int32{{5, 5}, {4, 2}, {2, 3}, {12, 2}, {7, 1}, {3, 5}, {4, 5}, {1, 1}, {7, 17}, {9, 10}, {3, 2}, {11, 10}, {8, 11}, {11, 8}, {1, 2}, {1, 3}, {1, 4}})
	}
}

func BenchmarkCachedObstacles(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cached.QueensAttack(20, 12, 9, 9, [][]int32{{5, 5}, {4, 2}, {2, 3}, {12, 2}, {7, 1}, {3, 5}, {4, 5}, {1, 1}, {7, 17}, {9, 10}, {3, 2}, {11, 10}, {8, 11}, {11, 8}, {1, 2}, {1, 3}, {1, 4}})
	}
}

func BenchmarkCalculedObstacles(b *testing.B) {
	for i := 0; i < b.N; i++ {
		calculed.QueensAttack(20, 12, 9, 9, [][]int32{{5, 5}, {4, 2}, {2, 3}, {12, 2}, {7, 1}, {3, 5}, {4, 5}, {1, 1}, {7, 17}, {9, 10}, {3, 2}, {11, 10}, {8, 11}, {11, 8}, {1, 2}, {1, 3}, {1, 4}})
	}
}
