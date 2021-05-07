package remote_package

import "testing"
import . "github.com/easierway/concurrent_map"

func TestRemotePackage(t *testing.T) {
	m := CreateConcurrentMap(99)
	m.Set(StrKey("key"), 10)
	t.Log(m.Get(StrKey("key")))
}
