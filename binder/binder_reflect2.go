// +build reflect2

package binder

import (
	"reflect"
	"strconv"
	"unsafe"

	"github.com/modern-go/concurrent"
	"github.com/modern-go/reflect2"
)

func bind(obj interface{}, values map[string][]string, tagKey string, canonical bool) error {
	rt := getTypeFromCache(obj)
	ptr := reflect2.PtrOf(obj)
	var rtf reflect2.StructField
	var fptr unsafe.Pointer
	var typ reflect2.Type
	var kind reflect.Kind
	var tag string

	for i := 0; i < rt.NumField(); i++ {
		rtf, fptr = rt.Field(i), rtf.UnsafeGet(ptr)
		typ, kind = rtf.Type(), typ.Kind()
		tag = rtf.Tag().Get(tagKey)

		switch tag {
		case "-":
			continue
		case "":
			tag = rtf.Name()

			if kind == reflect.Struct {
				if err := bind(typ.PackEFace(fptr), values, tagKey, canonical); err != nil {
					return err
				}
				continue
			}
		}

		vals, exists := values[canonicalKey(tag, canonical)]
		if !exists {
			continue
		}

		size := len(vals)
		if kind == reflect.Slice && size > 0 {
			sliceType := typ.(*reflect2.UnsafeSliceType)
			elemKind := sliceType.Elem().Kind()
			sliceType.UnsafeSet(fptr, sliceType.UnsafeMakeSlice(size, size))
			for j := 0; j < size; j++ {
				if err := setField(elemKind, vals[j], sliceType.UnsafeGetIndex(fptr, j)); err != nil {
					return err
				}
			}
		} else if size > 0 {
			if err := setField(kind, vals[0], fptr); err != nil {
				return err
			}
		}
	}
	return nil
}

func setField(kind reflect.Kind, val string, ptr unsafe.Pointer) error {
	switch kind {
	case reflect.Bool:
		return setBoolField(val, ptr)
	case reflect.Int:
		return setIntField(val, ptr)
	case reflect.Int8:
		return setInt8Field(val, ptr)
	case reflect.Int16:
		return setInt16Field(val, ptr)
	case reflect.Int32:
		return setInt32Field(val, ptr)
	case reflect.Int64:
		return setInt64Field(val, ptr)
	case reflect.Uint:
		return setUintField(val, ptr)
	case reflect.Uint8:
		return setUint8Field(val, ptr)
	case reflect.Uint16:
		return setUint16Field(val, ptr)
	case reflect.Uint32:
		return setUint32Field(val, ptr)
	case reflect.Uint64:
		return setUint64Field(val, ptr)
	case reflect.Float32:
		return setFloat32Field(val, ptr)
	case reflect.Float64:
		return setFloat64Field(val, ptr)
	case reflect.String:
		setString(val, ptr)
		return nil
	}
	return ErrInvalidType
}

func setBoolField(val string, ptr unsafe.Pointer) error {
	v, err := strconv.ParseBool(convertValue(val))
	if err == nil {
		*(*bool)(ptr) = v
	}
	return err
}

func setIntField(val string, ptr unsafe.Pointer) error {
	v, err := strconv.ParseInt(convertValue(val), 10, 0)
	if err == nil {
		*(*int)(ptr) = int(v)
	}
	return err
}

func setInt8Field(val string, ptr unsafe.Pointer) error {
	v, err := strconv.ParseInt(convertValue(val), 10, 8)
	if err == nil {
		*(*int8)(ptr) = int8(v)
	}
	return err
}

func setInt16Field(val string, ptr unsafe.Pointer) error {
	v, err := strconv.ParseInt(convertValue(val), 10, 16)
	if err == nil {
		*(*int16)(ptr) = int16(v)
	}
	return err
}

func setInt32Field(val string, ptr unsafe.Pointer) error {
	v, err := strconv.ParseInt(convertValue(val), 10, 32)
	if err == nil {
		*(*int32)(ptr) = int32(v)
	}
	return err
}

func setInt64Field(val string, ptr unsafe.Pointer) error {
	v, err := strconv.ParseInt(convertValue(val), 10, 64)
	if err == nil {
		*(*int64)(ptr) = v
	}
	return err
}

func setUintField(val string, ptr unsafe.Pointer) error {
	v, err := strconv.ParseUint(convertValue(val), 10, 0)
	if err == nil {
		*(*uint)(ptr) = uint(v)
	}
	return err
}

func setUint8Field(val string, ptr unsafe.Pointer) error {
	v, err := strconv.ParseUint(convertValue(val), 10, 8)
	if err == nil {
		*(*uint8)(ptr) = uint8(v)
	}
	return err
}

func setUint16Field(val string, ptr unsafe.Pointer) error {
	v, err := strconv.ParseUint(convertValue(val), 10, 16)
	if err == nil {
		*(*uint16)(ptr) = uint16(v)
	}
	return err
}

func setUint32Field(val string, ptr unsafe.Pointer) error {
	v, err := strconv.ParseUint(convertValue(val), 10, 32)
	if err == nil {
		*(*uint32)(ptr) = uint32(v)
	}
	return err
}

func setUint64Field(val string, ptr unsafe.Pointer) error {
	v, err := strconv.ParseUint(convertValue(val), 10, 64)
	if err == nil {
		*(*uint64)(ptr) = v
	}
	return err
}

func setFloat32Field(val string, ptr unsafe.Pointer) error {
	v, err := strconv.ParseFloat(convertValue(val), 32)
	if err == nil {
		*(*float32)(ptr) = float32(v)
	}
	return err
}

func setFloat64Field(val string, ptr unsafe.Pointer) error {
	v, err := strconv.ParseFloat(convertValue(val), 64)
	if err == nil {
		*(*float64)(ptr) = v
	}
	return err
}

func setString(val string, ptr unsafe.Pointer) {
	*(*string)(ptr) = val
}

var cache = concurrent.NewMap()

func getTypeFromCache(obj interface{}) (rt *reflect2.UnsafeStructType) {
	key := reflect2.RTypeOf(obj)
	if val, exists := cache.Load(key); exists {
		rt = val.(*reflect2.UnsafeStructType)
	} else {
		rt = reflect2.Type2(reflect.TypeOf(obj).Elem()).(*reflect2.UnsafeStructType)
		cache.Store(key, rt)
	}
	return
}
