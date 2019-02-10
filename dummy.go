package dummy

import (
	"math/rand"
	"reflect"
	"unsafe"
)

// LimitDepth represents the limit depth of recursively setting dummy values
var LimitDepth = 30

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// Set a dummy value recursively
func Set(i interface{}) {
	v := reflect.ValueOf(i)
	if v.Kind() != reflect.Ptr {
		panic("the argument must be a pointer")
	}
	set(v, 0)
}

func set(v reflect.Value, depth int) {
	if depth == LimitDepth {
		return
	}

	switch v.Kind() {
	case reflect.Bool:
		v.SetBool(rand.Int()&1 == 0)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v.SetInt(int64(rand.Uint64()))
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		v.SetUint(rand.Uint64())
	case reflect.Float32, reflect.Float64:
		v.SetFloat(rand.Float64())
	case reflect.Complex64, reflect.Complex128:
		v.SetComplex(complex(rand.Float64(), rand.Float64()))
	case reflect.String:
		b := make([]byte, rand.Intn(20))
		for i := range b {
			b[i] = letters[rand.Intn(len(letters))]
		}
		v.SetString(string(b))
	case reflect.Ptr:
		if v.IsNil() {
			v.Set(reflect.New(v.Type().Elem()))
		}
		set(v.Elem(), depth+1)
	case reflect.Struct:
		for i, n := 0, v.NumField(); i < n; i++ {
			v2 := v.Field(i)
			if !v2.CanSet() && v2.CanAddr() {
				v2 = reflect.NewAt(v2.Type(), unsafe.Pointer(v2.UnsafeAddr())).Elem()
			}
			set(v2, depth+1)
		}
	case reflect.Array:
		for i := 0; i < v.Len(); i++ {
			set(v.Index(i), depth+1)
		}
	case reflect.Slice:
		l := rand.Intn(5)
		v.Set(reflect.MakeSlice(v.Type(), l, l))
		for i := 0; i < v.Len(); i++ {
			set(v.Index(i), depth+1)
		}
	case reflect.Map:
		t := v.Type()
		if v.IsNil() {
			v.Set(reflect.MakeMap(t))
		}
		l := rand.Intn(5)
		for i := 0; i < l; i++ {
			key := reflect.New(t.Key()).Elem()
			value := reflect.New(t.Elem()).Elem()
			set(key, depth+1)
			set(value, depth+1)
			v.SetMapIndex(key, value)
		}
	}
}
