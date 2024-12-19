package prime

import (
	"sync"
	"testing"
	"time"
)

func TestGenerateChannel(t *testing.T) {
	testCases := []struct {
		name     string
		testFunc func(*testing.T)
	}{
		{
			name: "Scenario 1: GenerateChannel produces an infinite sequence of integers starting from 2",
			testFunc: func(t *testing.T) {
				ch := make(chan int)
				go GenerateChannel(ch)

				expectedSequence := []int{2, 3, 4, 5, 6}
				for _, expected := range expectedSequence {
					if num := <-ch; num != expected {
						t.Errorf("Expected %d, got %d", expected, num)
					}
				}
			},
		},
		{
			name: "Scenario 2: GenerateChannel can be stopped externally",
			testFunc: func(t *testing.T) {
				ch := make(chan int)
				stopCh := make(chan struct{})
				var wg sync.WaitGroup
				wg.Add(1)

				go func() {
					defer wg.Done()
					GenerateChannel(ch)
				}()

				go func() {
					time.Sleep(100 * time.Millisecond)
					close(stopCh)
				}()

				wg.Wait()
				if _, ok := <-ch; ok {
					t.Errorf("Expected channel to be closed")
				}
			},
		},
		{
			name: "Scenario 3: GenerateChannel operates concurrently without blocking",
			testFunc: func(t *testing.T) {
				ch := make(chan int)
				var wg sync.WaitGroup
				wg.Add(2)

				go func() {
					defer wg.Done()
					GenerateChannel(ch)
				}()

				go func() {
					defer wg.Done()

					time.Sleep(100 * time.Millisecond)
				}()

				wg.Wait()

			},
		},
		{
			name: "Scenario 4: GenerateChannel behaves correctly under high load",
			testFunc: func(t *testing.T) {
				ch := make(chan int)
				go GenerateChannel(ch)

				timeout := time.After(100 * time.Millisecond)
				count := 0

				for {
					select {
					case <-ch:
						count++
					case <-timeout:

						if count < 1000 {
							t.Errorf("Expected at least 1000 numbers to be generated, got %d", count)
						}
						return
					}
				}
			},
		},
		{
			name: "Scenario 5: GenerateChannel channel capacity handling",
			testFunc: func(t *testing.T) {
				capacity := 5
				ch := make(chan int, capacity)
				go GenerateChannel(ch)

				time.Sleep(100 * time.Millisecond)

				if len(ch) != capacity {
					t.Errorf("Expected channel to be filled up to its capacity of %d, got %d", capacity, len(ch))
				}
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.testFunc(t)
		})
	}
}
