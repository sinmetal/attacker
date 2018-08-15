package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"

	"cloud.google.com/go/compute/metadata"
	"github.com/google/uuid"
)

func main() {
	var projectID string
	if metadata.OnGCE() {
		p, err := metadata.ProjectID()
		if err != nil {
			panic(err)
		}
		projectID = p
	}

	fmt.Printf("ProjectID is %s\n", projectID)

	run(sample)
}

func run(f func() error) {
	var wg sync.WaitGroup
	// 5min * i の時間回る
	for i := 0; i < 10; i++ {
		km := float64(500) * math.Pow(1.5, float64(i)) // 並列実行数

		// 秒間1回程度実行されるとして、300秒で5分ぐらいになる
		for j := 0; j < 300; j++ {
			for k := 0.0; k < km; k++ {
				wg.Add(1)
				go func(i int, j int, k float64) {
					ms := rand.Intn(1000)
					ms++
					t := time.NewTicker(time.Duration(ms) * time.Millisecond)
					select {
					case <-t.C:
						defer wg.Done()
						fmt.Printf("Start %d,%d,%f\n", i, j, k)
						if err := f(); err != nil {
							fmt.Printf("Faile %d,%d,%f\n", i, j, k)
						}
					}

				}(i, j, k)

			}
			time.Sleep(1000 * time.Millisecond)
		}
	}

	wg.Wait()
}

func sample() error {
	id := uuid.New().String()
	fmt.Printf("UUID is %s\n", id)

	return nil
}
