package commons

import "testing"

func TestDatabase_Upsert(t *testing.T) {
	Upsert("foo","bar")
}
