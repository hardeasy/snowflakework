package snowflakework

import (
	"testing"
)

func TestMakeId(t *testing.T) {
	c := MakeId()
	t.Log("id:", c)
}

func TestMakeIdMulti(t *testing.T) {
	for i := 0; i < 1000000; i++ {
		_ = MakeId()
	}
}
