package main

import (
	"fmt"
	"log"
	"time"

	"gopkg.in/vmihailenco/msgpack.v2"

	"github.com/tinylib/msgp/msgp"
)

type Data struct {
	Persons []Person
}

type Person struct {
	Name       string `msg:"name"`
	Address    string `msg:"address"`
	Age        int    `msg:"age"`
	Hidden     string `msg:"-"` // this field is ignored
	unexported bool   // this field is also ignored
}

func main() {
	const N = 10000000
	data := new(Data)
	data.Persons = make([]Person, N)
	for i := 0; i < N; i++ {
		data.Persons[i].Name = fmt.Sprint(i)
		data.Persons[i].Age = i
	}
	start := time.Now()
	bts, err := data.MarshalMsg(nil)
	if err != nil {
		log.Println(err)
	}
	log.Println("size:", len(bts))
	log.Println("msgp time:", time.Now().Sub(start))

	start = time.Now()
	b, err := msgpack.Marshal(data)
	if err != nil {
		panic(err)
	}
	log.Println("size:", len(b))
	log.Println("msgpack time:", time.Now().Sub(start))
}

// DecodeMsg implements msgp.Decodable
func (z *Data) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var isz uint32
	isz, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for isz > 0 {
		isz--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Persons":
			var xsz uint32
			xsz, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Persons) >= int(xsz) {
				z.Persons = z.Persons[:xsz]
			} else {
				z.Persons = make([]Person, xsz)
			}
			for xvk := range z.Persons {
				var isz uint32
				isz, err = dc.ReadMapHeader()
				if err != nil {
					return
				}
				for isz > 0 {
					isz--
					field, err = dc.ReadMapKeyPtr()
					if err != nil {
						return
					}
					switch msgp.UnsafeString(field) {
					case "name":
						z.Persons[xvk].Name, err = dc.ReadString()
						if err != nil {
							return
						}
					case "address":
						z.Persons[xvk].Address, err = dc.ReadString()
						if err != nil {
							return
						}
					case "age":
						z.Persons[xvk].Age, err = dc.ReadInt()
						if err != nil {
							return
						}
					default:
						err = dc.Skip()
						if err != nil {
							return
						}
					}
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Data) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 1
	// write "Persons"
	err = en.Append(0x81, 0xa7, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Persons)))
	if err != nil {
		return
	}
	for xvk := range z.Persons {
		// map header, size 3
		// write "name"
		err = en.Append(0x83, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
		if err != nil {
			return err
		}
		err = en.WriteString(z.Persons[xvk].Name)
		if err != nil {
			return
		}
		// write "address"
		err = en.Append(0xa7, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73)
		if err != nil {
			return err
		}
		err = en.WriteString(z.Persons[xvk].Address)
		if err != nil {
			return
		}
		// write "age"
		err = en.Append(0xa3, 0x61, 0x67, 0x65)
		if err != nil {
			return err
		}
		err = en.WriteInt(z.Persons[xvk].Age)
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Data) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 1
	// string "Persons"
	o = append(o, 0x81, 0xa7, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Persons)))
	for xvk := range z.Persons {
		// map header, size 3
		// string "name"
		o = append(o, 0x83, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
		o = msgp.AppendString(o, z.Persons[xvk].Name)
		// string "address"
		o = append(o, 0xa7, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73)
		o = msgp.AppendString(o, z.Persons[xvk].Address)
		// string "age"
		o = append(o, 0xa3, 0x61, 0x67, 0x65)
		o = msgp.AppendInt(o, z.Persons[xvk].Age)
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Data) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var isz uint32
	isz, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for isz > 0 {
		isz--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Persons":
			var xsz uint32
			xsz, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Persons) >= int(xsz) {
				z.Persons = z.Persons[:xsz]
			} else {
				z.Persons = make([]Person, xsz)
			}
			for xvk := range z.Persons {
				var isz uint32
				isz, bts, err = msgp.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				for isz > 0 {
					isz--
					field, bts, err = msgp.ReadMapKeyZC(bts)
					if err != nil {
						return
					}
					switch msgp.UnsafeString(field) {
					case "name":
						z.Persons[xvk].Name, bts, err = msgp.ReadStringBytes(bts)
						if err != nil {
							return
						}
					case "address":
						z.Persons[xvk].Address, bts, err = msgp.ReadStringBytes(bts)
						if err != nil {
							return
						}
					case "age":
						z.Persons[xvk].Age, bts, err = msgp.ReadIntBytes(bts)
						if err != nil {
							return
						}
					default:
						bts, err = msgp.Skip(bts)
						if err != nil {
							return
						}
					}
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

func (z *Data) Msgsize() (s int) {
	s = 1 + 8 + msgp.ArrayHeaderSize
	for xvk := range z.Persons {
		s += 1 + 5 + msgp.StringPrefixSize + len(z.Persons[xvk].Name) + 8 + msgp.StringPrefixSize + len(z.Persons[xvk].Address) + 4 + msgp.IntSize
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Person) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var isz uint32
	isz, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for isz > 0 {
		isz--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "name":
			z.Name, err = dc.ReadString()
			if err != nil {
				return
			}
		case "address":
			z.Address, err = dc.ReadString()
			if err != nil {
				return
			}
		case "age":
			z.Age, err = dc.ReadInt()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z Person) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 3
	// write "name"
	err = en.Append(0x83, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Name)
	if err != nil {
		return
	}
	// write "address"
	err = en.Append(0xa7, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Address)
	if err != nil {
		return
	}
	// write "age"
	err = en.Append(0xa3, 0x61, 0x67, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteInt(z.Age)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Person) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "name"
	o = append(o, 0x83, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Name)
	// string "address"
	o = append(o, 0xa7, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73)
	o = msgp.AppendString(o, z.Address)
	// string "age"
	o = append(o, 0xa3, 0x61, 0x67, 0x65)
	o = msgp.AppendInt(o, z.Age)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Person) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var isz uint32
	isz, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for isz > 0 {
		isz--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "name":
			z.Name, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "address":
			z.Address, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "age":
			z.Age, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

func (z Person) Msgsize() (s int) {
	s = 1 + 5 + msgp.StringPrefixSize + len(z.Name) + 8 + msgp.StringPrefixSize + len(z.Address) + 4 + msgp.IntSize
	return
}
