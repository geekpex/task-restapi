package data

import (
	"sync"
)

// IDs concurrent safe int64 slice
type IDs struct {
	ids []int64
	m   sync.Mutex
}

// Add adds id to internal slice
func (i *IDs) Add(id int64) {
	i.m.Lock()
	defer i.m.Unlock()

	i.ids = append(i.ids, id)
}

// Delete deletes all occurances of id in internal slice
func (i *IDs) Delete(id int64) {
	i.m.Lock()
	defer i.m.Unlock()

	for j := len(i.ids); j >= 0; j-- {
		if id == i.ids[j] {
			i.ids = append(i.ids[:j], i.ids[j+1:]...)
		}
	}
}

// IDs returns copy of internal slice
func (i *IDs) IDs() []int64 {
	i.m.Lock()
	defer i.m.Unlock()

	ids := make([]int64, len(i.ids))
	copy(ids, i.ids)
	return ids
}
