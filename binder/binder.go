package binder

import (
	"errors"
	"net/textproto"
	"reflect"
	"strconv"
	"time"
)

func bind(ptr interface{}, values map[string][]string, tagKey string, canonical bool) error {
	rt := reflect.TypeOf(ptr).Elem()
	rv := reflect.ValueOf(ptr).Elem()
	for i := 0; i < rt.NumField(); i++ {
		rtf := rt.Field(i)
		rvf := rv.Field(i)
		if !rvf.CanSet() {
			continue
		}

		kind := rvf.Kind()
		tag := rtf.Tag.Get(tagKey)
		switch tag {
		case "-":
			continue
		case "":
			tag = rtf.Name

			if kind == reflect.Struct {
				if err := bind(rvf.Addr().Interface(), values, tagKey, canonical); err != nil {
					return err
				}
				continue
			}
		}

		val, exists := values[canonicalKey(tag, canonical)]
		if !exists {
			continue
		}

		numElems := len(val)
		if kind == reflect.Slice && numElems > 0 {
			elemKind := rvf.Type().Elem().Kind()
			slice := reflect.MakeSlice(rvf.Type(), numElems, numElems)
			for j := 0; j < numElems; j++ {
				if err := setField(elemKind, val[j], slice.Index(j)); err != nil {
					return err
				}
			}
			rvf.Set(slice)
		} else {
			if _, ok := rvf.Interface().(time.Time); ok {
				if err := setTimeField(val[0], rtf, rvf); err != nil {
					return err
				}
				continue
			}
			if err := setField(kind, val[0], rvf); err != nil {
				return err
			}
		}
	}
	return nil
}

func setField(kind reflect.Kind, val string, field reflect.Value) error {
	switch kind {
	case reflect.Int:
		return setIntField(val, 0, field)
	case reflect.Int8:
		return setIntField(val, 8, field)
	case reflect.Int16:
		return setIntField(val, 16, field)
	case reflect.Int32:
		return setIntField(val, 32, field)
	case reflect.Int64:
		return setIntField(val, 64, field)
	case reflect.Uint:
		return setUintField(val, 0, field)
	case reflect.Uint8:
		return setUintField(val, 8, field)
	case reflect.Uint16:
		return setUintField(val, 16, field)
	case reflect.Uint32:
		return setUintField(val, 32, field)
	case reflect.Uint64:
		return setUintField(val, 64, field)
	case reflect.Bool:
		return setBoolField(val, field)
	case reflect.Float32:
		return setFloatField(val, 32, field)
	case reflect.Float64:
		return setFloatField(val, 64, field)
	case reflect.String:
		field.SetString(val)
		return nil
	}
	return errors.New("unknown type")
}

func setIntField(val string, bitSize int, field reflect.Value) error {
	if val == "" {
		val = "0"
	}

	intVal, err := strconv.ParseInt(val, 10, bitSize)
	if err == nil {
		field.SetInt(intVal)
	}
	return err
}

func setUintField(val string, bitSize int, field reflect.Value) error {
	if val == "" {
		val = "0"
	}

	uintVal, err := strconv.ParseUint(val, 10, bitSize)
	if err == nil {
		field.SetUint(uintVal)
	}
	return err
}

func setBoolField(val string, field reflect.Value) error {
	if val == "" {
		val = "false"
	}

	boolVal, err := strconv.ParseBool(val)
	if err == nil {
		field.SetBool(boolVal)
	}
	return err
}

func setFloatField(val string, bitSize int, field reflect.Value) error {
	if val == "" {
		val = "0.0"
	}

	floatVal, err := strconv.ParseFloat(val, bitSize)
	if err == nil {
		field.SetFloat(floatVal)
	}
	return err
}

func setTimeField(val string, structField reflect.StructField, field reflect.Value) error {
	timeFormat := structField.Tag.Get("time_format")
	if timeFormat == "" {
		return errors.New("time_format not exists")
	}

	if val == "" {
		field.Set(reflect.ValueOf(time.Time{}))
		return nil
	}

	loc := time.Local
	if isUTC, _ := strconv.ParseBool(structField.Tag.Get("time_utc")); isUTC {
		loc = time.UTC
	}

	t, err := time.ParseInLocation(timeFormat, val, loc)
	if err != nil {
		return err
	}

	field.Set(reflect.ValueOf(t))
	return nil
}

func canonicalKey(key string, canonical bool) string {
	if canonical {
		key = textproto.CanonicalMIMEHeaderKey(key)
	}
	return key
}
