// +build !jsoniter

package json

import (
	"encoding/json"
	"fmt"
)

var (
	Marshal       = json.Marshal
	MarshalIndent = json.MarshalIndent
	Unmarshal     = json.Unmarshal
	NewEncoder    = json.NewEncoder
	NewDecoder    = json.NewDecoder
)

func init() {
	fmt.Println("use json")
}
