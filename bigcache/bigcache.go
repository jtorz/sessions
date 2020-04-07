package bigcache

import (
	"github.com/allegro/bigcache/v2"
	"github.com/gin-contrib/sessions"
	gsb "github.com/jtorz/gorilla-sessions-bigcache"
)

// Store sessions.Store wrapper
type Store interface {
	sessions.Store
}

// NewStore returns a new store.
//
//	* client: bigcache client (github.com/bradfitz/gobigcache/bigcache)
//	* keyPrefix: prefix for the stored keys.
//	* keypairs: used to encode and decode cookie values.
func NewStore(
	client *bigcache.BigCache, keyPrefix string, keyPairs ...[]byte,
) Store {
	bigcacherClient := gsb.NewGoBigcacher(client)
	return NewBigcacheStore(bigcacherClient, keyPrefix, keyPairs...)
}

// NewBigcacheStore returns a new store.
//
//	* client: bigcache client which implements the gsb.Bigcacher interface
//	* keyPrefix: prefix for the stored keys.
//	* keypairs: used to encode and decode cookie values.
func NewBigcacheStore(
	client gsb.BigCacher, keyPrefix string, keyPairs ...[]byte,
) Store {
	return &store{gsb.NewBigCacherStore(client, keyPrefix, keyPairs...)}
}

// store gorilla-sessions-bigcache/BigcacheStore wrapper
type store struct {
	*gsb.BigcacheStore
}

func (c *store) Options(options sessions.Options) {
	c.BigcacheStore.Options = options.ToGorillaOptions()
}
