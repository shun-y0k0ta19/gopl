// Copyright © 2016 "Shun Yokota" All rights reserved

package sexpr

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
)

// Marshal encodes a Go value in S-expression form.
func Marshal(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	if err := encode(&buf, reflect.ValueOf(v)); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func encode(buf *bytes.Buffer, v reflect.Value) error {
	switch v.Kind() {
	case reflect.Invalid:
		buf.WriteString("nil")

	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		fmt.Fprintf(buf, "%d", v.Int())

	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		fmt.Fprintf(buf, "%d", v.Uint())

	case reflect.Float32, reflect.Float64:
		fmt.Fprintf(buf, "%v", v.Float())

	case reflect.Complex64, reflect.Complex128:
		fmt.Fprintf(buf, "#C(%v %v)", real(v.Complex()), imag(v.Complex()))

	case reflect.Bool:
		if v.Bool() {
			fmt.Fprintf(buf, "t")
		} else {
			fmt.Fprintf(buf, "nil")
		}

	case reflect.String:
		fmt.Fprintf(buf, "%q", v.String())

	case reflect.Ptr:
		return encode(buf, v.Elem())

	case reflect.Array, reflect.Slice: // (value ...)
		buf.WriteByte('(')
		lines := bytes.Split(buf.Bytes(), []byte("\n"))
		bcounts := len(lines[len(lines)-1]) - 1
		arrayBuf := bytes.NewBuffer([]byte{})
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				arrayBuf.WriteByte(' ')
			}
			if err := encode(arrayBuf, v.Index(i)); err != nil {
				return err
			}
			if v.Type() == reflect.SliceOf(reflect.TypeOf("")) || v.Type() == reflect.ArrayOf(v.Len(), reflect.TypeOf("")) {
				fmt.Fprintf(arrayBuf, "\n%s", strings.Repeat(" ", bcounts))
			}
		}
		trim := bytes.TrimSpace(arrayBuf.Bytes())
		arrayBuf = bytes.NewBuffer(trim)
		fmt.Fprintf(buf, "%s)", arrayBuf)

	case reflect.Struct: // ((name value) ...)
		buf.WriteByte('(')
		for i := 0; i < v.NumField(); i++ {
			if i > 0 {
				buf.WriteByte(' ')
			}
			fmt.Fprintf(buf, "(%s ", v.Type().Field(i).Name)
			if err := encode(buf, v.Field(i)); err != nil {
				return err
			}
			fmt.Fprintf(buf, ")\n")
		}
		buf = bytes.NewBuffer(bytes.TrimSpace(buf.Bytes()))
		fmt.Fprintf(buf, ")")

	case reflect.Map: // ((key value) ...)
		buf.WriteByte('(')
		lines := bytes.Split(buf.Bytes(), []byte("\n"))
		bcounts := len(lines[len(lines)-1]) - 1
		mapBuf := bytes.NewBuffer([]byte{})
		for i, key := range v.MapKeys() {
			if i > 0 {
				mapBuf.WriteByte(' ')
			}
			mapBuf.WriteByte('(')
			if err := encode(mapBuf, key); err != nil {
				return err
			}
			mapBuf.WriteByte(' ')
			if err := encode(mapBuf, v.MapIndex(key)); err != nil {
				return err
			}
			fmt.Fprintf(mapBuf, ")\n%s", strings.Repeat(" ", bcounts))
		}
		trim := bytes.TrimSpace(mapBuf.Bytes())
		fmt.Fprintf(buf, "%s)", trim)

	case reflect.Interface:
		if v.IsNil() {
			fmt.Fprintf(buf, "nil")
		} else {
			fmt.Fprintf(buf, "(\"%v\" ", v.Elem().Type())
			if err := encode(buf, v.Elem()); err != nil {
				return err
			}
			fmt.Fprint(buf, ")")
		}
	default: //  func, interface
		return fmt.Errorf("unsupported type: %s", v.Type())
	}
	return nil
}
