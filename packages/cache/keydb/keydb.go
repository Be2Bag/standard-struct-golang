package keydb

import (
	"fmt"
)

type KeyDB struct {
	prefix string
}

func New(prefix string) *KeyDB {
	return &KeyDB{prefix: prefix}
}

func (k *KeyDB) LineKey(sessionID string) string {
	key := fmt.Sprintf("%s:line:%s", k.prefix, sessionID)
	return key
}

func (k *KeyDB) AuthSessionKey(sessionID string) string {
	key := fmt.Sprintf("%s:auth:sessions:%s", k.prefix, sessionID)
	return key
}

func (k *KeyDB) RegisterKey(sessionID string) string {
	key := fmt.Sprintf("%s:auth:registration:%s", k.prefix, sessionID)
	return key
}
