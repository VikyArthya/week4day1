package main

import (
	"fmt"
	"sync"
	"time"
)

type DataModel struct {
	ID    int
	Value string
}

func StoreData(data DataModel, wg *sync.WaitGroup, mu *sync.Mutex, dataSlice *[]DataModel) {
	defer wg.Done()

	mu.Lock()
	*dataSlice = append(*dataSlice, data)
	mu.Unlock()

	LogDataStored(data)
}

func LogDataStored(data DataModel) {
	time.Sleep(10 * time.Millisecond)
	fmt.Printf("Data berhasil disimpan: ID=%d, Value=%s\n", data.ID, data.Value)
}

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	dataSlice := make([]DataModel, 0)

	for i := 1; i <= 100; i++ {
		wg.Add(1)
		data := DataModel{
			ID:    i,
			Value: fmt.Sprintf("Value %d", i),
		}

		go StoreData(data, &wg, &mu, &dataSlice)
	}

	wg.Wait()

	fmt.Printf("Total data yang disimpan: %d\n", len(dataSlice))
}
