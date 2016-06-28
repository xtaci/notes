package main

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import (
	"github.com/tinylib/msgp/msgp"
)

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
				err = z.Persons[xvk].DecodeMsg(dc)
				if err != nil {
					return
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
		err = z.Persons[xvk].EncodeMsg(en)
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
		o, err = z.Persons[xvk].MarshalMsg(o)
		if err != nil {
			return
		}
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
				bts, err = z.Persons[xvk].UnmarshalMsg(bts)
				if err != nil {
					return
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
		s += z.Persons[xvk].Msgsize()
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
		case "Name":
			z.Name, err = dc.ReadString()
			if err != nil {
				return
			}
		case "Address":
			z.Address, err = dc.ReadString()
			if err != nil {
				return
			}
		case "Age":
			z.Age, err = dc.ReadInt()
			if err != nil {
				return
			}
		case "Hidden":
			z.Hidden, err = dc.ReadString()
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
func (z *Person) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 4
	// write "Name"
	err = en.Append(0x84, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Name)
	if err != nil {
		return
	}
	// write "Address"
	err = en.Append(0xa7, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Address)
	if err != nil {
		return
	}
	// write "Age"
	err = en.Append(0xa3, 0x41, 0x67, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteInt(z.Age)
	if err != nil {
		return
	}
	// write "Hidden"
	err = en.Append(0xa6, 0x48, 0x69, 0x64, 0x64, 0x65, 0x6e)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Hidden)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Person) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 4
	// string "Name"
	o = append(o, 0x84, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Name)
	// string "Address"
	o = append(o, 0xa7, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73)
	o = msgp.AppendString(o, z.Address)
	// string "Age"
	o = append(o, 0xa3, 0x41, 0x67, 0x65)
	o = msgp.AppendInt(o, z.Age)
	// string "Hidden"
	o = append(o, 0xa6, 0x48, 0x69, 0x64, 0x64, 0x65, 0x6e)
	o = msgp.AppendString(o, z.Hidden)
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
		case "Name":
			z.Name, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Address":
			z.Address, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Age":
			z.Age, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "Hidden":
			z.Hidden, bts, err = msgp.ReadStringBytes(bts)
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

func (z *Person) Msgsize() (s int) {
	s = 1 + 5 + msgp.StringPrefixSize + len(z.Name) + 8 + msgp.StringPrefixSize + len(z.Address) + 4 + msgp.IntSize + 7 + msgp.StringPrefixSize + len(z.Hidden)
	return
}
