package main

import (
	"math/rand"
	"time"

	"github.com/k0kubun/pp"
	"github.com/kasari/dummy"
)

type test struct {
	b       bool
	i       int
	i8      int8
	i16     int16
	i32     int32
	i64     int64
	ui      uint
	ui8     uint8
	ui16    uint16
	ui32    uint32
	ui64    uint64
	uiptr   uintptr
	f32     float32
	f64     float64
	c64     complex64
	c128    complex128
	pb      *bool
	pi      *int
	pi8     *int8
	pi16    *int16
	pi32    *int32
	pi64    *int64
	pui     *uint
	pui8    *uint8
	pui16   *uint16
	pui32   *uint32
	pui64   *uint64
	puiptr  *uintptr
	pf32    *float32
	pf64    *float64
	pc64    *complex64
	pc128   *complex128
	slice   []uint32
	mp      map[string]int
	array   [3]float32
	inferno map[complex128][2]map[string]map[uint64]bool
	node    *node
}

type node struct {
	child *node
}

func main() {
	rand.Seed(time.Now().Unix())

	t := &test{}
	dummy.Set(t)
	pp.Print(t)
}
