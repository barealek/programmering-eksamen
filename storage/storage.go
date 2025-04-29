package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

type Storage struct {
	fileName string
	data     []*Entry
	mu       *sync.Mutex
}

func NewStorage(indexNavn string) *Storage {
	return &Storage{
		fileName: indexNavn,
		data:     make([]*Entry, 0),
		mu:       &sync.Mutex{},
	}
}

func (st *Storage) Save() error {
	st.mu.Lock()
	defer st.mu.Unlock()

	f, err := os.OpenFile(st.fileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}

	defer f.Close()

	return json.NewEncoder(f).Encode(st.data)
}

func (st *Storage) Insert(e *Entry) error {
	st.mu.Lock()
	defer st.mu.Unlock()

	if e := st.Get(e.ID); e != nil {
		return fmt.Errorf("entry med id:%v findes allerede i storage", e.ID)
	}

	st.data = append(st.data, e)
	return nil
}

func (st *Storage) Get(id string) *Entry {
	for _, e := range st.data {
		if e.ID == id {
			return e
		}
	}

	return nil
}
