package path

import (
	"DSGO/Graph/graph"
	"fmt"
	"time"
)

type Path = graph.Path

func BenchMark() {
	start := time.Now()
	roads, total, err := readGraph() //IO就是慢！！！
	if err != nil {
		fmt.Println("Illegal Input")
		return
	}
	roads_u := uAdjList(roads)
	matrix_u := transform(roads_u)
	matrix := sAdjMatrix(matrix_u)
	size := len(roads)
	fmt.Printf("Prepare Graph [%d vertexes & %d edges] in %v\n",
		size, total, time.Since(start))

	start = time.Now()
	for i := 0; i < size; i++ {
		SPFA(roads, i)
	}
	fmt.Println("SPFA:           ", time.Since(start))

	start = time.Now()
	for i := 0; i < size; i++ {
		Dijkstra(roads_u, i)
	}
	fmt.Println("Dijkstra:       ", time.Since(start))

	start = time.Now()
	for i := 0; i < size; i++ {
		DijkstraX(roads_u, i)
	}
	fmt.Println("Simple Dijkstra:", time.Since(start))

	start = time.Now()
	for i := 0; i < size; i++ {
		PlainDijkstra(matrix_u, i)
	}
	fmt.Println("Plain Dijkstra: ", time.Since(start))

	start = time.Now()
	FloydWarshall(matrix)
	fmt.Println("Floyd-Warshall: ", time.Since(start))
}

func readGraph() (roads [][]PathS, total int, err error) {
	var size int
	_, err = fmt.Scan(&size, &total)
	if err != nil || size < 2 || size > total {
		return nil, 0, err
	}
	roads = make([][]PathS, size)
	var a, b, dist int
	for i := 0; i < total; i++ {
		_, err = fmt.Scan(&a, &b, &dist)
		if err != nil {
			return nil, 0, err
		}
		roads[a] = append(roads[a], PathS{Next: b, Dist: dist})
	}
	return roads, total, nil
}

func uAdjList(roads [][]PathS) [][]Path {
	out := make([][]Path, len(roads))
	for i, vec := range roads {
		line := make([]Path, len(vec))
		for j, path := range vec {
			line[j] = Path{Next: path.Next, Weight: uint(path.Dist)}
		}
		out[i] = line
	}
	return out
}

func transform(roads [][]Path) [][]uint {
	size := len(roads)
	matrix := make([][]uint, size)
	for i, vec := range roads {
		line := make([]uint, size) //全零
		for _, path := range vec {
			line[path.Next] = path.Weight
		}
		matrix[i] = line
	}
	return matrix
}

func sAdjMatrix(matrix [][]uint) [][]int {
	size := len(matrix)
	out := make([][]int, size)
	for i := 0; i < size; i++ {
		out[i] = make([]int, size)
		for j := 0; j < size; j++ {
			if matrix[i][j] == 0 && i != j {
				out[i][j] = MAX_DIST
			} else {
				out[i][j] = int(matrix[i][j])
			}
		}
	}
	return out
}
