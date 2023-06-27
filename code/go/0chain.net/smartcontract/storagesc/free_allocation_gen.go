package storagesc

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// MarshalMsg implements msgp.Marshaler
func (z *freeStorageAssigner) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 6
	// string "i"
	o = append(o, 0x86, 0xa1, 0x69)
	o = msgp.AppendUint64(o, z.IndividualLimit)
	// string "t"
	o = append(o, 0xa1, 0x74)
	o = msgp.AppendUint64(o, z.TotalLimit)
	// string "r"
	o = append(o, 0xa1, 0x72)
	o = msgp.AppendUint64(o, z.CurrentRedeemed)
	// string "rt"
	o = append(o, 0xa2, 0x72, 0x74)
	o = msgp.AppendArrayHeader(o, uint32(len(z.RedeemedTimestamps)))
	for za0001 := range z.RedeemedTimestamps {
		o = msgp.AppendInt64(o, z.RedeemedTimestamps[za0001])
	}
	// string "c"
	o = append(o, 0xa1, 0x63)
	o = msgp.AppendString(o, z.ClientId)
	// string "p"
	o = append(o, 0xa1, 0x70)
	o = msgp.AppendString(o, z.PublicKey)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *freeStorageAssigner) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "i":
			z.IndividualLimit, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "IndividualLimit")
				return
			}
		case "t":
			z.TotalLimit, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "TotalLimit")
				return
			}
		case "r":
			z.CurrentRedeemed, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "CurrentRedeemed")
				return
			}
		case "rt":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "RedeemedTimestamps")
				return
			}
			if cap(z.RedeemedTimestamps) >= int(zb0002) {
				z.RedeemedTimestamps = (z.RedeemedTimestamps)[:zb0002]
			} else {
				z.RedeemedTimestamps = make([]int64, zb0002)
			}
			for za0001 := range z.RedeemedTimestamps {
				z.RedeemedTimestamps[za0001], bts, err = msgp.ReadInt64Bytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "RedeemedTimestamps", za0001)
					return
				}
			}
		case "c":
			z.ClientId, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "ClientId")
				return
			}
		case "p":
			z.PublicKey, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "PublicKey")
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
func (z *freeStorageAssigner) Msgsize() (s int) {
	s = 1 + 2 + msgp.Uint64Size + 2 + msgp.Uint64Size + 2 + msgp.Uint64Size + 3 + msgp.ArrayHeaderSize + (len(z.RedeemedTimestamps) * (msgp.Int64Size)) + 2 + msgp.StringPrefixSize + len(z.ClientId) + 2 + msgp.StringPrefixSize + len(z.PublicKey)
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *freeStorageMarker) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 5
	// string "Assigner"
	o = append(o, 0x85, 0xa8, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72)
	o = msgp.AppendString(o, z.Assigner)
	// string "Recipient"
	o = append(o, 0xa9, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74)
	o = msgp.AppendString(o, z.Recipient)
	// string "FreeTokens"
	o = append(o, 0xaa, 0x46, 0x72, 0x65, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73)
	o = msgp.AppendFloat64(o, z.FreeTokens)
	// string "Timestamp"
	o = append(o, 0xa9, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70)
	o, err = z.Timestamp.MarshalMsg(o)
	if err != nil {
		err = msgp.WrapError(err, "Timestamp")
		return
	}
	// string "Signature"
	o = append(o, 0xa9, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65)
	o = msgp.AppendString(o, z.Signature)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *freeStorageMarker) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "Assigner":
			z.Assigner, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Assigner")
				return
			}
		case "Recipient":
			z.Recipient, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Recipient")
				return
			}
		case "FreeTokens":
			z.FreeTokens, bts, err = msgp.ReadFloat64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "FreeTokens")
				return
			}
		case "Timestamp":
			bts, err = z.Timestamp.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "Timestamp")
				return
			}
		case "Signature":
			z.Signature, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Signature")
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
func (z *freeStorageMarker) Msgsize() (s int) {
	s = 1 + 9 + msgp.StringPrefixSize + len(z.Assigner) + 10 + msgp.StringPrefixSize + len(z.Recipient) + 11 + msgp.Float64Size + 10 + z.Timestamp.Msgsize() + 10 + msgp.StringPrefixSize + len(z.Signature)
	return
}

// MarshalMsg implements msgp.Marshaler
func (z freeStorageUpgradeInput) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "AllocationId"
	o = append(o, 0x82, 0xac, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64)
	o = msgp.AppendString(o, z.AllocationId)
	// string "Marker"
	o = append(o, 0xa6, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x72)
	o = msgp.AppendString(o, z.Marker)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *freeStorageUpgradeInput) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "AllocationId":
			z.AllocationId, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "AllocationId")
				return
			}
		case "Marker":
			z.Marker, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Marker")
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
func (z freeStorageUpgradeInput) Msgsize() (s int) {
	s = 1 + 13 + msgp.StringPrefixSize + len(z.AllocationId) + 7 + msgp.StringPrefixSize + len(z.Marker)
	return
}
