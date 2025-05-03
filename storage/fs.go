package storage

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sync"
)

type fs struct {
	fileName string
	data     []*Entry
	mu       *sync.Mutex
}

func NewFsStorage(indexNavn string) Storage {
	return &fs{
		fileName: indexNavn,
		data:     make([]*Entry, 0),
		mu:       &sync.Mutex{},
	}
}

func (st *fs) Load() error {
	st.mu.Lock()
	defer st.mu.Unlock()

	fmt.Println("Opening", st.fileName)
	f, err := os.OpenFile(st.fileName, os.O_RDONLY, 0644) // Changed flags
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("File not found, starting with empty storage.")
			return nil
		}
		return err
	}

	defer f.Close()

	return json.NewDecoder(f).Decode(&st.data)
}

func (st *fs) Save() error {
	st.mu.Lock()
	defer st.mu.Unlock()

	f, err := os.OpenFile(st.fileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}

	fmt.Println("Saving storage")
	for _, e := range st.data {
		fmt.Println(e)
	}

	defer f.Close()

	return json.NewEncoder(f).Encode(st.data)
}

func (st *fs) Insert(e *Entry) error {
	st.mu.Lock()
	defer st.mu.Unlock()

	if e := st.Get(e.ID); e != nil {
		return fmt.Errorf("entry med id:%v findes allerede i storage", e.ID)
	}

	st.data = append(st.data, e)
	return nil
}

func (st *fs) Get(id string) *Entry {
	for _, e := range st.data {
		if e.ID == id {
			return e
		}
	}

	return nil
}

func (st *fs) Delete(e *Entry) error {
	st.mu.Lock()
	defer st.mu.Unlock()

	fmt.Println("sletter", e.ID)
	foundEntry := func() *Entry {
		for _, re := range st.data {
			if e.ID == re.ID {
				return e
			}
		}
		return nil
	}()

	if foundEntry == nil {
		return fmt.Errorf("entry med id:%v findes ikke i storage", e.ID)
	}

	st.data = removeFromSlice(st.data, e)

	fp := fmt.Sprintf("data/%s.bin", e.ID)

	if err := os.Remove(fp); err != nil {
		return err
	}

	// st.Save()

	return nil
}

func (st *fs) FileDest(id string) (io.WriteCloser, error) {
	st.mu.Lock()
	defer st.mu.Unlock()

	fp := fmt.Sprintf("data/%s.bin", id)

	fileDst, err := os.Create(fp)
	return fileDst, err
}

func (st *fs) FileSource(e *Entry) (io.ReadCloser, error) {
	st.mu.Lock()
	defer st.mu.Unlock()

	fp := fmt.Sprintf("data/%s.bin", e.ID)

	fileSrc, err := os.Open(fp)
	return fileSrc, err
}
