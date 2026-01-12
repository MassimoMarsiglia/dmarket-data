package pool

import (
	"bytes"
	"encoding/json"
	"io"
	"sync"
)

func (p *Pool[T, Z]) Transform(input T) ([]Z, error) {
	data, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	// First reader
	var r io.Reader = bytes.NewReader(data)

	var wg sync.WaitGroup
	errChan := make(chan error, 1)

	// Build the chain
	for _, tr := range p.TransformFuncs {
		pr, pw := io.Pipe()

		wg.Add(1)
		go func(in io.Reader, out *io.PipeWriter, tf TransformFunc) {
			defer wg.Done()
			defer out.Close()

			if err := tf(in, out); err != nil {
				out.CloseWithError(err)
				select {
				case errChan <- err:
				default:
				}
			}
		}(r, pw, tr)

		r = pr // next stage reads from previous
	}

	// Final output collector
	var results []Z
	dec := json.NewDecoder(r)

	for dec.More() {
		var z Z
		if err := dec.Decode(&z); err != nil {
			return nil, err
		}
		results = append(results, z)
	}

	wg.Wait()

	select {
	case err := <-errChan:
		return results, err
	default:
	}

	return results, nil
}
