package xslog

import "sync"

type AttrKey int

var (
	nAttrKey AttrKey = 0

	attrKeyMu = &sync.RWMutex{}
)

func NewAttrKey() AttrKey {
	attrKeyMu.Lock()
	defer attrKeyMu.Unlock()

	key := nAttrKey
	nAttrKey++
	return key
}
