package bigcache

import (
	"testing"
	"time"

	"github.com/allegro/bigcache/v2"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/tester"
)

var defaultConfig = bigcache.DefaultConfig(time.Minute)

var newStore = func(t *testing.T) sessions.Store {
	client, err := bigcache.NewBigCache(defaultConfig)
	if err != nil {
		t.Fatal(err)
	}

	store := NewStore(client, "", []byte("secret"))
	return store
}

func TestBigcache_SessionGetSet(t *testing.T) {
	tester.GetSet(t, newStore)
}

func TestBigcache_SessionDeleteKey(t *testing.T) {
	tester.DeleteKey(t, newStore)
}

func TestBigcache_SessionFlashes(t *testing.T) {
	tester.Flashes(t, newStore)
}

func TestBigcache_SessionClear(t *testing.T) {
	tester.Clear(t, newStore)
}

func TestBigcache_SessionOptions(t *testing.T) {
	tester.Options(t, newStore)
}
