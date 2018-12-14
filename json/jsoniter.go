// +build jsoniter

package json

import (
	"fmt"

	"github.com/json-iterator/go"
)

var (
	json          = jsoniter.ConfigCompatibleWithStandardLibrary
	Marshal       = json.Marshal
	MarshalIndent = json.MarshalIndent
	Unmarshal     = json.Unmarshal
	NewEncoder    = json.NewEncoder
	NewDecoder    = json.NewDecoder
)

func init() {
	fmt.Println("use jsoniter")
}
