package node

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// MarshalMsg implements msgp.Marshaler
func (z *Info) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 4
	// string "BuildTag"
	o = append(o, 0x84, 0xa8, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x54, 0x61, 0x67)
	o = msgp.AppendString(o, z.BuildTag)
	// string "StateMissingNodes"
	o = append(o, 0xb1, 0x53, 0x74, 0x61, 0x74, 0x65, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x4e, 0x6f, 0x64, 0x65, 0x73)
	o = msgp.AppendInt64(o, z.StateMissingNodes)
	// string "MinersMedianNetworkTime"
	o = append(o, 0xb7, 0x4d, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x6e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x54, 0x69, 0x6d, 0x65)
	o = msgp.AppendInt64(o, z.MinersMedianNetworkTime)
	// string "AvgBlockTxns"
	o = append(o, 0xac, 0x41, 0x76, 0x67, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x78, 0x6e, 0x73)
	o = msgp.AppendInt(o, z.AvgBlockTxns)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Info) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "BuildTag":
			z.BuildTag, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "BuildTag")
				return
			}
		case "StateMissingNodes":
			z.StateMissingNodes, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "StateMissingNodes")
				return
			}
		case "MinersMedianNetworkTime":
			z.MinersMedianNetworkTime, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "MinersMedianNetworkTime")
				return
			}
		case "AvgBlockTxns":
			z.AvgBlockTxns, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "AvgBlockTxns")
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
func (z *Info) Msgsize() (s int) {
	s = 1 + 9 + msgp.StringPrefixSize + len(z.BuildTag) + 18 + msgp.Int64Size + 24 + msgp.Int64Size + 13 + msgp.IntSize
	return
}
