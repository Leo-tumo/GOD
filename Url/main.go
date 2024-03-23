package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

//type permissions map[string]bool
//
//type user struct {
//	Name        string      `json:"username"`
//	Password    string      `json:"-"`
//	Permissions permissions `json:"perms,omitempty"`
//}
//
//func main() {
//	users := []user{
//		{"Leo", "1234", nil},
//		{"god", "42", permissions{"admin": true}},
//		{"evil", "666", permissions{"write": true}},
//	}
//	out, err := json.MarshalIndent(users, "", "\t")
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	fmt.Println(string(out))
//}

type user struct {
	Name        string          `json:"username"`
	Permissions map[string]bool `json:"perms,omitempty"`
}

func main() {
	var input []byte
	for in := bufio.NewScanner(os.Stdin); in.Scan(); {
		input = append(input, in.Bytes()...)
	}

	var users []user
	err := json.Unmarshal(input, &users)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(users)
}
