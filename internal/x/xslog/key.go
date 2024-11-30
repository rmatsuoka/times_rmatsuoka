package xslog

import "sync"

type AttrKey struct {
	id int
}

var (
	nAttrKey = 0

	attrKeyMu = &sync.RWMutex{}
)

func NewAttrKey() AttrKey {
	attrKeyMu.Lock()
	defer attrKeyMu.Unlock()

	id := nAttrKey
	nAttrKey++
	return AttrKey{id: id}
}
