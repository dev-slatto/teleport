/*
Copyright 2021 Gravitational, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package srv

import (
	"io"
	"sync"

	log "github.com/sirupsen/logrus"
)

// MultiReader is a wrapper around io.MultiReader which allows to add and remove readers
type MultiReader struct {
	sync.RWMutex
	readers map[string]io.Reader
	multi   io.Reader
}

// NewMultiReader creates a new multitreader.
func NewMultiReader() *MultiReader {
	return &MultiReader{
		readers: make(map[string]io.Reader),
		multi:   io.MultiReader(),
	}
}

// Read reads from any of the underlying readers.
func (r *MultiReader) Read(p []byte) (int, error) {
	r.RLock()
	multi := r.multi
	r.RUnlock()

	read, err := multi.Read(p)
	if err != nil {
		log.Warnf("failed to read from multireader: %v", err)
	}

	return read, nil
}

// AddReader adds the reader.
func (r *MultiReader) AddReader(name string, reader io.Reader) {
	r.Lock()
	defer r.Unlock()
	r.readers[name] = reader

	var readers []io.Reader
	for _, reader := range r.readers {
		readers = append(readers, reader)
	}

	r.multi = io.MultiReader(readers...)
}

// RemoveReader removes a reader.
func (r *MultiReader) RemoveReader(name string) {
	r.Lock()
	defer r.Unlock()
	delete(r.readers, name)

	var readers []io.Reader
	for _, reader := range r.readers {
		readers = append(readers, reader)
	}

	r.multi = io.MultiReader(readers...)
}