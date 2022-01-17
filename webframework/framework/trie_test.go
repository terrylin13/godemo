package framework

import "testing"

func TestInsert(t *testing.T) {
	root := new(node)
	parts := []string{"user", "login"}
	root.insert("/user/login", parts, 0)
	for i, node := range root.children {
		t.Log("root", i, node)
	}

	n1 := root.search(parts, 0)

	t.Log(n1)

}
