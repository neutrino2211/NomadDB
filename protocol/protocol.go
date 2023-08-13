package protocol

import (
	"errors"
	"net"
	"reflect"

	"github.com/neutrino2211/go-option"
)

var errPacketDefinitionInstantiationError = errors.New("Error instantiating a packet")

type PacketDefinition[T any] struct {
	Fields []string
	Types  []reflect.Type
	Sizes  []uint

	ValidationFunc func(p *T) bool
}

func getSize(typ reflect.Type) int {
	size := 0
	switch typ.Kind() {
	case reflect.String:
		size = typ.Len() * 8
	case reflect.Array:
		size = typ.Len() * 8
	default:
		size = typ.Bits()
	}
	return size
}

func (def *PacketDefinition[T]) Confirm(data []byte) bool {
	l := len(data)
	s := 0
	for _, t := range def.Types {
		s += getSize(t)
	}

	return l == s/8
}

func (def *PacketDefinition[T]) GetSize() uint64 {
	var r uint64 = 0

	for _, s := range def.Sizes {
		r += uint64(s)
	}

	return r
}

func (def *PacketDefinition[T]) Validate(t *T) bool {
	return def.ValidationFunc(t)
}

func (def *PacketDefinition[T]) Instance(data []byte) *option.Optional[T] {
	var r T
	pkt := reflect.New(reflect.TypeOf(r))

	for i, f := range def.Fields {
		size := def.Sizes[i]

		var start uint = 0

		n := 0
		for n < i {
			start += def.Sizes[n]
			n++
		}
		var val = reflect.ValueOf(data[start : start+size])
		if size == 1 {
			pkt.Elem().FieldByName(f).Set(reflect.ValueOf(data[start]))
		} else {
			reflect.Copy(pkt.Elem().FieldByName(f), val)
		}
	}

	r, ok := pkt.Elem().Interface().(T)

	if !ok {
		return option.Err[T](errPacketDefinitionInstantiationError)
	}

	return option.Some(r)
}

func (def *PacketDefinition[T]) ReadFromConn(conn net.Conn) *option.Optional[T] {
	buf := make([]byte, int(def.GetSize()))

	conn.Read(buf)

	if def.Confirm(buf) {
		return def.Instance(buf)
	}

	return option.None[T]()
}

func GeneratePacketDefinition[T any](v func(*T) bool) *PacketDefinition[T] {
	fields := []string{}
	types := []reflect.Type{}
	sizes := []uint{}

	var t T

	val := reflect.Indirect(reflect.ValueOf(t))

	for i := 0; i < val.NumField(); i++ {
		fields = append(fields, val.Type().Field(i).Name)
		types = append(types, val.Type().Field(i).Type)
		sizes = append(sizes, uint(getSize(val.Type().Field(i).Type)/8))
	}

	return &PacketDefinition[T]{
		Fields:         fields,
		Types:          types,
		Sizes:          sizes,
		ValidationFunc: v,
	}
}
