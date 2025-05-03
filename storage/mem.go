// package storage

// import (
// 	"fmt"
// 	"sync"
// )

// // mem implements the Storage interface for in-memory storage.
// type mem struct {
// 	data []*Entry
// 	mu   *sync.Mutex
// }

// // NewMemStorage creates a new in-memory storage instance.
// // func NewMemStorage() Storage {
// // 	return &mem{
// // 		data: make([]*Entry, 0),
// // 		mu:   &sync.Mutex{},
// // 	}
// // }

// // Load is a no-op for in-memory storage as data is not persisted.
// func (st *mem) Load() error {
// 	// No operation needed for in-memory storage
// 	fmt.Println("Initializing in-memory storage.")
// 	return nil
// }

// // Save is a no-op for in-memory storage as data is not persisted.
// func (st *mem) Save() error {
// 	// No operation needed for in-memory storage
// 	fmt.Println("In-memory storage state (not saving to disk):")
// 	st.mu.Lock()
// 	defer st.mu.Unlock()
// 	for _, e := range st.data {
// 		fmt.Println(e)
// 	}
// 	return nil
// }

// // Insert adds a new entry to the in-memory storage.
// // It returns an error if an entry with the same ID already exists.
// func (st *mem) Insert(e *Entry) error {
// 	st.mu.Lock()
// 	defer st.mu.Unlock()

// 	// Check if entry already exists
// 	if existing := st.get(e.ID); existing != nil {
// 		return fmt.Errorf("entry med id:%v findes allerede i storage", e.ID)
// 	}

// 	st.data = append(st.data, e)
// 	return nil
// }

// // Get retrieves an entry by its ID from the in-memory storage.
// // It returns nil if the entry is not found.
// func (st *mem) Get(id string) *Entry {
// 	st.mu.Lock()
// 	defer st.mu.Unlock()
// 	return st.get(id)
// }

// // get is an internal helper function to retrieve an entry without locking.
// // The caller must handle locking.
// func (st *mem) get(id string) *Entry {
// 	for _, e := range st.data {
// 		if e.ID == id {
// 			return e
// 		}
// 	}
// 	return nil
// }

// // Delete removes an entry by its ID from the in-memory storage.
// // It returns an error if the entry is not found.
// func (st *mem) Delete(id string) error {
// 	st.mu.Lock()
// 	defer st.mu.Unlock()

//		fmt.Println("sletter", id, "fra memory")
//		for i, e := range st.data {
//			if e.ID == id {
//				fmt.Println("Fundet entry i memory db", e)
//				st.data = removeFromSlice(st.data, i)
//				return nil
//			}
//		}
//		return fmt.Errorf("entry med id:%v findes ikke i memory storage", id)
//	}
package storage
