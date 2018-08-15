package main

import (
	"context"
	"fmt"
	"math"
	"math/rand"
	"os"
	"sync"
	"time"

	"cloud.google.com/go/compute/metadata"
	"contrib.go.opencensus.io/exporter/stackdriver"
	"github.com/google/uuid"
	"go.opencensus.io/trace"
)

var ds *Datastore

func main() {
	var projectID string
	if metadata.OnGCE() {
		p, err := metadata.ProjectID()
		if err != nil {
			panic(err)
		}
		projectID = p
	}
	p := os.Getenv("GOOGLE_CLOUD_PROJECT")
	fmt.Printf("Env GOOGLE_CLOUD_PROJECT:%s\n", p)
	if len(p) > 0 {
		projectID = p
	}

	fmt.Printf("ProjectID is %s\n", projectID)

	{
		exporter, err := stackdriver.NewExporter(stackdriver.Options{
			ProjectID: projectID,
		})
		if err != nil {
			panic(err)
		}
		trace.RegisterExporter(exporter)
	}

	{
		var err error
		ds, err = NewDatastore(projectID)
		if err != nil {
			panic(err)
		}
	}

	run(sample)
}

func run(f func(i int, j int, k float64) error) {
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
						if err := runDatastore(i, j, k); err != nil {
							fmt.Printf("Faile %d,%d,%f : %+v\n", i, j, k, err)
						}
					}

				}(i, j, k)

			}
			time.Sleep(1000 * time.Millisecond)
		}
	}

	wg.Wait()
}

func sample(i int, j int, k float64) error {
	id := uuid.New().String()
	fmt.Printf("UUID is %s\n", id)

	return nil
}

func runDatastore(i int, j int, k float64) error {
	ctx := context.Background()
	return ds.PutHeavyEntity(ctx, i, j, k)
}
