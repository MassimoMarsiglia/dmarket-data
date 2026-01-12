package pool

import (
	"bytes"
	"context"
	"io"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/MassimoMarsiglia/dmarket-bot/pkg/utils/pool/worker"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func NewTestWorker(ctx context.Context) (*worker.Worker[*string], error) {
	id := uuid.New()
	w := worker.New[*string](worker.Config[*string]{
		ID: id,
		Work: func() func(ctx context.Context) (*string, error) {
			counter := 0
			return func(ctx context.Context) (*string, error) {
				counter++
				result := strconv.Itoa(counter)

				return &result, nil
			}
		}(),
		Delay: 100 * time.Millisecond,
	})
	return w, nil
}

func DefaultTransformFunc(r io.Reader, w io.Writer) error {
	var buf bytes.Buffer
	_, err := buf.ReadFrom(r)
	if err != nil {
		return err
	}
	result := buf.String()
	_, err = w.Write([]byte(result))
	return err
}

func Test_Pool(t *testing.T) {
	pool := New[*string, string](
		SpawnFunc[*string](NewTestWorker),
		[]TransformFunc{
			DefaultTransformFunc,
		},
	)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	workerCount := 100
	expectedResults := workerCount * 10 // 10 results per worker
	workers, err := pool.Spawn(ctx, workerCount)
	assert.NoError(t, err)
	assert.Len(t, workers, workerCount)

	assert.NotNil(t, pool)
	sub := pool.Publisher.Subscribe()
	assert.NotNil(t, sub)

	t.Log("Starting pool...")
	go func() {
		err := pool.Start(ctx, 10*time.Millisecond)
		if err != nil {
			cancel()
			t.Errorf("failed to start pool: %v", err)
		}
	}()
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	wg.Add(1)
	results := make([]string, 0, expectedResults)
	completed := false
	go func() {
		for {
			select {
			case msg := <-sub.C:
				mu.Lock()
				results = append(results, msg)
				if len(results)%10 == 0 {
					t.Logf("Collected %d/%d results", len(results), expectedResults)
				}
				if len(results) >= expectedResults && !completed {
					completed = true
					wg.Done()
				}
				mu.Unlock()
			case err := <-sub.E:
				t.Logf("Error received: %v", err)
			case <-ctx.Done():
				t.Log("Context cancelled, stopping result collector")
				return
			}
		}
	}()

	// Wait for results with timeout
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		mu.Lock()
		t.Logf("Successfully collected %d/%d results (%.1f%%)", len(results), expectedResults, float64(len(results))/float64(expectedResults)*100)
		mu.Unlock()
	case <-time.After(15 * time.Second):
		mu.Lock()
		t.Fatalf("Test timed out waiting for results, only got %d/%d", len(results), expectedResults)
		mu.Unlock()
	}

	cancel()
	time.Sleep(100 * time.Millisecond) // Give goroutines time to clean up
}

func Test_WorkerDelay(t *testing.T) {
	delay := 200 * time.Millisecond
	tolerance := 50 * time.Millisecond // Allow 50ms variance

	// Create a worker with a specific delay
	id := uuid.New()
	w := worker.New[*string](worker.Config[*string]{
		ID: id,
		Work: func() func(ctx context.Context) (*string, error) {
			return func(ctx context.Context) (*string, error) {
				result := time.Now().Format(time.RFC3339Nano)
				return &result, nil
			}
		}(),
		Delay: delay,
	})

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Subscribe to worker output
	sub := w.Publisher.Subscribe()
	defer sub.Close()

	// Start the worker
	go func() {
		_ = w.Start(ctx)
	}()

	// Collect timestamps
	timestamps := make([]time.Time, 0, 5)
	expectedCount := 5

	for i := 0; i < expectedCount; i++ {
		select {
		case <-sub.C:
			timestamps = append(timestamps, time.Now())
		case <-time.After(delay + 1*time.Second):
			t.Fatalf("Timeout waiting for result %d", i+1)
		}
	}

	cancel()

	// Verify we got the expected number of results
	assert.Len(t, timestamps, expectedCount)

	// Measure intervals between results
	for i := 1; i < len(timestamps); i++ {
		interval := timestamps[i].Sub(timestamps[i-1])

		t.Logf("Interval %d: %v (expected ~%v)", i, interval, delay)

		// Check that the interval is within tolerance
		minExpected := delay - tolerance
		maxExpected := delay + tolerance

		assert.GreaterOrEqual(t, interval, minExpected,
			"Interval %d (%v) is too short, expected at least %v", i, interval, minExpected)
		assert.LessOrEqual(t, interval, maxExpected,
			"Interval %d (%v) is too long, expected at most %v", i, interval, maxExpected)
	}

	// Calculate average interval
	if len(timestamps) > 1 {
		totalDuration := timestamps[len(timestamps)-1].Sub(timestamps[0])
		avgInterval := totalDuration / time.Duration(len(timestamps)-1)
		t.Logf("Average interval: %v (expected ~%v)", avgInterval, delay)

		assert.InDelta(t, float64(delay), float64(avgInterval), float64(tolerance),
			"Average interval should be close to configured delay")
	}
}
