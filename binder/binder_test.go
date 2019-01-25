package binder

import (
	"fmt"
	"testing"
)

type Teacher struct {
	Name string `header:"name" param:"name"`
}

type Student struct {
	Bool    bool     `header:"bool" param:"bool"`
	Int     int      `header:"int" param:"int"`
	Int8    int8     `header:"int8" param:"int8"`
	Int16   int16    `header:"int16" param:"int16"`
	Int32   int32    `header:"int32" param:"int32"`
	Int64   int64    `header:"int64" param:"int64"`
	Uint    uint     `header:"uint" param:"uint"`
	Uint8   uint8    `header:"uint8" param:"uint8"`
	Uint16  uint16   `header:"uint16" param:"uint16"`
	Uint32  uint32   `header:"uint32" param:"uint32"`
	Uint64  uint64   `header:"uint64" param:"uint64"`
	Float32 float32  `header:"float32" param:"float32"`
	Float64 float64  `header:"float64" param:"float64"`
	String  string   `header:"string" param:"string"`
	Slice   []string `header:"slice" param:"slice"`
	Teacher
}

var m = map[string][]string{
	"bool":    []string{"true"},
	"int":     []string{"10"},
	"int8":    []string{"8"},
	"int16":   []string{"16"},
	"int32":   []string{"32"},
	"int64":   []string{"64"},
	"uint":    []string{"20"},
	"uint8":   []string{"18"},
	"uint16":  []string{"26"},
	"uint32":  []string{"42"},
	"uint64":  []string{"74"},
	"float32": []string{"1.1"},
	"float64": []string{"2.2"},
	"string":  []string{"a"},
	"slice":   []string{"hello", "world"},
	"name": []string{"jack"},
}

func TestHeaderBinder_Bind(t *testing.T) {
	s := Student{}
	if err := bind(s, m, "header", false); err != nil {
		fmt.Printf("error - %v\n", err)
		return
	}
	fmt.Printf("s:%+v\n", s)

}
