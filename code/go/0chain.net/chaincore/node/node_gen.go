package node

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// MarshalMsg implements msgp.Marshaler
func (z *Node) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 11
	// string "Client"
	o = append(o, 0x8b, 0xa6, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74)
	o, err = z.Client.MarshalMsg(o)
	if err != nil {
		err = msgp.WrapError(err, "Client")
		return
	}
	// string "N2NHost"
	o = append(o, 0xa7, 0x4e, 0x32, 0x4e, 0x48, 0x6f, 0x73, 0x74)
	o = msgp.AppendString(o, z.N2NHost)
	// string "Host"
	o = append(o, 0xa4, 0x48, 0x6f, 0x73, 0x74)
	o = msgp.AppendString(o, z.Host)
	// string "Port"
	o = append(o, 0xa4, 0x50, 0x6f, 0x72, 0x74)
	o = msgp.AppendInt(o, z.Port)
	// string "Path"
	o = append(o, 0xa4, 0x50, 0x61, 0x74, 0x68)
	o = msgp.AppendString(o, z.Path)
	// string "Type"
	o = append(o, 0xa4, 0x54, 0x79, 0x70, 0x65)
	o = msgp.AppendInt8(o, z.Type)
	// string "Description"
	o = append(o, 0xab, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e)
	o = msgp.AppendString(o, z.Description)
	// string "SetIndex"
	o = append(o, 0xa8, 0x53, 0x65, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78)
	o = msgp.AppendInt(o, z.SetIndex)
	// string "Status"
	o = append(o, 0xa6, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73)
	o = msgp.AppendInt(o, z.Status)
	// string "InPrevMB"
	o = append(o, 0xa8, 0x49, 0x6e, 0x50, 0x72, 0x65, 0x76, 0x4d, 0x42)
	o = msgp.AppendBool(o, z.InPrevMB)
	// string "Info"
	o = append(o, 0xa4, 0x49, 0x6e, 0x66, 0x6f)
	o, err = z.Info.MarshalMsg(o)
	if err != nil {
		err = msgp.WrapError(err, "Info")
		return
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Node) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Client":
			bts, err = z.Client.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "Client")
				return
			}
		case "N2NHost":
			z.N2NHost, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "N2NHost")
				return
			}
		case "Host":
			z.Host, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Host")
				return
			}
		case "Port":
			z.Port, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Port")
				return
			}
		case "Path":
			z.Path, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Path")
				return
			}
		case "Type":
			z.Type, bts, err = msgp.ReadInt8Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Type")
				return
			}
		case "Description":
			z.Description, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Description")
				return
			}
		case "SetIndex":
			z.SetIndex, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "SetIndex")
				return
			}
		case "Status":
			z.Status, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Status")
				return
			}
		case "InPrevMB":
			z.InPrevMB, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "InPrevMB")
				return
			}
		case "Info":
			bts, err = z.Info.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "Info")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Node) Msgsize() (s int) {
	s = 1 + 7 + z.Client.Msgsize() + 8 + msgp.StringPrefixSize + len(z.N2NHost) + 5 + msgp.StringPrefixSize + len(z.Host) + 5 + msgp.IntSize + 5 + msgp.StringPrefixSize + len(z.Path) + 5 + msgp.Int8Size + 12 + msgp.StringPrefixSize + len(z.Description) + 9 + msgp.IntSize + 7 + msgp.IntSize + 9 + msgp.BoolSize + 5 + z.Info.Msgsize()
	return
}
