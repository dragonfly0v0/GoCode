package store

import (
	"testing"
)

func TestStore(t *testing.T) {
	monster := Monster{
		Name:  "牛魔王",
		Age:   8000,
		Skill: "法天象地",
	}

	monster.Store(&monster)
	t.Logf("执行正确\n")
}

func TestReStore(t *testing.T) {
	var monster Monster
	res := monster.Restore()

	if !res {
		t.Fatalf("moster.Restore() failed, reson: %v", res)
	}
	t.Logf("执行正确\n")
}
