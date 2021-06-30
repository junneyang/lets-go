package main

import (
	"encoding/json"
	"testing"
)

func TestUser(t *testing.T) {
	str := `{"username": "zhangsan", "age": 18}`
	user := &User{}
	json.Unmarshal([]byte(str), user)
	t.Logf("user=>%v", *user)
	b, _ := json.Marshal(user)
	t.Log(string(b))
}
