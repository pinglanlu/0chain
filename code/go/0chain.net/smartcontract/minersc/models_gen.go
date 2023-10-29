package minersc

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"0chain.net/chaincore/block"
	"github.com/tinylib/msgp/msgp"
)

// MarshalMsg implements msgp.Marshaler
func (z *DKGMinerNodes) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 12
	// string "MinN"
	o = append(o, 0x8c, 0xa4, 0x4d, 0x69, 0x6e, 0x4e)
	o = msgp.AppendInt(o, z.MinN)
	// string "MaxN"
	o = append(o, 0xa4, 0x4d, 0x61, 0x78, 0x4e)
	o = msgp.AppendInt(o, z.MaxN)
	// string "TPercent"
	o = append(o, 0xa8, 0x54, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74)
	o = msgp.AppendFloat64(o, z.TPercent)
	// string "KPercent"
	o = append(o, 0xa8, 0x4b, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74)
	o = msgp.AppendFloat64(o, z.KPercent)
	// string "SimpleNodes"
	o = append(o, 0xab, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x73)
	o = msgp.AppendMapHeader(o, uint32(len(z.SimpleNodes)))
	keys_za0001 := make([]string, 0, len(z.SimpleNodes))
	for k := range z.SimpleNodes {
		keys_za0001 = append(keys_za0001, k)
	}
	msgp.Sort(keys_za0001)
	for _, k := range keys_za0001 {
		za0002 := z.SimpleNodes[k]
		o = msgp.AppendString(o, k)
		if za0002 == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = za0002.MarshalMsg(o)
			if err != nil {
				err = msgp.WrapError(err, "SimpleNodes", k)
				return
			}
		}
	}
	// string "T"
	o = append(o, 0xa1, 0x54)
	o = msgp.AppendInt(o, z.T)
	// string "K"
	o = append(o, 0xa1, 0x4b)
	o = msgp.AppendInt(o, z.K)
	// string "N"
	o = append(o, 0xa1, 0x4e)
	o = msgp.AppendInt(o, z.N)
	// string "XPercent"
	o = append(o, 0xa8, 0x58, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74)
	o = msgp.AppendFloat64(o, z.XPercent)
	// string "RevealedShares"
	o = append(o, 0xae, 0x52, 0x65, 0x76, 0x65, 0x61, 0x6c, 0x65, 0x64, 0x53, 0x68, 0x61, 0x72, 0x65, 0x73)
	o = msgp.AppendMapHeader(o, uint32(len(z.RevealedShares)))
	keys_za0003 := make([]string, 0, len(z.RevealedShares))
	for k := range z.RevealedShares {
		keys_za0003 = append(keys_za0003, k)
	}
	msgp.Sort(keys_za0003)
	for _, k := range keys_za0003 {
		za0004 := z.RevealedShares[k]
		o = msgp.AppendString(o, k)
		o = msgp.AppendInt(o, za0004)
	}
	// string "Waited"
	o = append(o, 0xa6, 0x57, 0x61, 0x69, 0x74, 0x65, 0x64)
	o = msgp.AppendMapHeader(o, uint32(len(z.Waited)))
	keys_za0005 := make([]string, 0, len(z.Waited))
	for k := range z.Waited {
		keys_za0005 = append(keys_za0005, k)
	}
	msgp.Sort(keys_za0005)
	for _, k := range keys_za0005 {
		za0006 := z.Waited[k]
		o = msgp.AppendString(o, k)
		o = msgp.AppendBool(o, za0006)
	}
	// string "StartRound"
	o = append(o, 0xaa, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x6f, 0x75, 0x6e, 0x64)
	o = msgp.AppendInt64(o, z.StartRound)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *DKGMinerNodes) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "MinN":
			z.MinN, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "MinN")
				return
			}
		case "MaxN":
			z.MaxN, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "MaxN")
				return
			}
		case "TPercent":
			z.TPercent, bts, err = msgp.ReadFloat64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "TPercent")
				return
			}
		case "KPercent":
			z.KPercent, bts, err = msgp.ReadFloat64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "KPercent")
				return
			}
		case "SimpleNodes":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "SimpleNodes")
				return
			}
			if z.SimpleNodes == nil {
				z.SimpleNodes = make(SimpleNodes, zb0002)
			} else if len(z.SimpleNodes) > 0 {
				for key := range z.SimpleNodes {
					delete(z.SimpleNodes, key)
				}
			}
			for zb0002 > 0 {
				var za0001 string
				var za0002 *SimpleNode
				zb0002--
				za0001, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "SimpleNodes")
					return
				}
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					za0002 = nil
				} else {
					if za0002 == nil {
						za0002 = new(SimpleNode)
					}
					bts, err = za0002.UnmarshalMsg(bts)
					if err != nil {
						err = msgp.WrapError(err, "SimpleNodes", za0001)
						return
					}
				}
				z.SimpleNodes[za0001] = za0002
			}
		case "T":
			z.T, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "T")
				return
			}
		case "K":
			z.K, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "K")
				return
			}
		case "N":
			z.N, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "N")
				return
			}
		case "XPercent":
			z.XPercent, bts, err = msgp.ReadFloat64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "XPercent")
				return
			}
		case "RevealedShares":
			var zb0003 uint32
			zb0003, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "RevealedShares")
				return
			}
			if z.RevealedShares == nil {
				z.RevealedShares = make(map[string]int, zb0003)
			} else if len(z.RevealedShares) > 0 {
				for key := range z.RevealedShares {
					delete(z.RevealedShares, key)
				}
			}
			for zb0003 > 0 {
				var za0003 string
				var za0004 int
				zb0003--
				za0003, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "RevealedShares")
					return
				}
				za0004, bts, err = msgp.ReadIntBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "RevealedShares", za0003)
					return
				}
				z.RevealedShares[za0003] = za0004
			}
		case "Waited":
			var zb0004 uint32
			zb0004, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Waited")
				return
			}
			if z.Waited == nil {
				z.Waited = make(map[string]bool, zb0004)
			} else if len(z.Waited) > 0 {
				for key := range z.Waited {
					delete(z.Waited, key)
				}
			}
			for zb0004 > 0 {
				var za0005 string
				var za0006 bool
				zb0004--
				za0005, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "Waited")
					return
				}
				za0006, bts, err = msgp.ReadBoolBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "Waited", za0005)
					return
				}
				z.Waited[za0005] = za0006
			}
		case "StartRound":
			z.StartRound, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "StartRound")
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
func (z *DKGMinerNodes) Msgsize() (s int) {
	s = 1 + 5 + msgp.IntSize + 5 + msgp.IntSize + 9 + msgp.Float64Size + 9 + msgp.Float64Size + 12 + msgp.MapHeaderSize
	if z.SimpleNodes != nil {
		for za0001, za0002 := range z.SimpleNodes {
			_ = za0002
			s += msgp.StringPrefixSize + len(za0001)
			if za0002 == nil {
				s += msgp.NilSize
			} else {
				s += za0002.Msgsize()
			}
		}
	}
	s += 2 + msgp.IntSize + 2 + msgp.IntSize + 2 + msgp.IntSize + 9 + msgp.Float64Size + 15 + msgp.MapHeaderSize
	if z.RevealedShares != nil {
		for za0003, za0004 := range z.RevealedShares {
			_ = za0004
			s += msgp.StringPrefixSize + len(za0003) + msgp.IntSize
		}
	}
	s += 7 + msgp.MapHeaderSize
	if z.Waited != nil {
		for za0005, za0006 := range z.Waited {
			_ = za0006
			s += msgp.StringPrefixSize + len(za0005) + msgp.BoolSize
		}
	}
	s += 11 + msgp.Int64Size
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *GlobalNode) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 30
	// string "ViewChange"
	o = append(o, 0xde, 0x0, 0x1e, 0xaa, 0x56, 0x69, 0x65, 0x77, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65)
	o = msgp.AppendInt64(o, z.ViewChange)
	// string "MaxN"
	o = append(o, 0xa4, 0x4d, 0x61, 0x78, 0x4e)
	o = msgp.AppendInt(o, z.MaxN)
	// string "MinN"
	o = append(o, 0xa4, 0x4d, 0x69, 0x6e, 0x4e)
	o = msgp.AppendInt(o, z.MinN)
	// string "MaxS"
	o = append(o, 0xa4, 0x4d, 0x61, 0x78, 0x53)
	o = msgp.AppendInt(o, z.MaxS)
	// string "MinS"
	o = append(o, 0xa4, 0x4d, 0x69, 0x6e, 0x53)
	o = msgp.AppendInt(o, z.MinS)
	// string "MaxDelegates"
	o = append(o, 0xac, 0x4d, 0x61, 0x78, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73)
	o = msgp.AppendInt(o, z.MaxDelegates)
	// string "TPercent"
	o = append(o, 0xa8, 0x54, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74)
	o = msgp.AppendFloat64(o, z.TPercent)
	// string "KPercent"
	o = append(o, 0xa8, 0x4b, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74)
	o = msgp.AppendFloat64(o, z.KPercent)
	// string "XPercent"
	o = append(o, 0xa8, 0x58, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74)
	o = msgp.AppendFloat64(o, z.XPercent)
	// string "LastRound"
	o = append(o, 0xa9, 0x4c, 0x61, 0x73, 0x74, 0x52, 0x6f, 0x75, 0x6e, 0x64)
	o = msgp.AppendInt64(o, z.LastRound)
	// string "MaxStake"
	o = append(o, 0xa8, 0x4d, 0x61, 0x78, 0x53, 0x74, 0x61, 0x6b, 0x65)
	o, err = z.MaxStake.MarshalMsg(o)
	if err != nil {
		err = msgp.WrapError(err, "MaxStake")
		return
	}
	// string "MinStake"
	o = append(o, 0xa8, 0x4d, 0x69, 0x6e, 0x53, 0x74, 0x61, 0x6b, 0x65)
	o, err = z.MinStake.MarshalMsg(o)
	if err != nil {
		err = msgp.WrapError(err, "MinStake")
		return
	}
	// string "MinStakePerDelegate"
	o = append(o, 0xb3, 0x4d, 0x69, 0x6e, 0x53, 0x74, 0x61, 0x6b, 0x65, 0x50, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65)
	o, err = z.MinStakePerDelegate.MarshalMsg(o)
	if err != nil {
		err = msgp.WrapError(err, "MinStakePerDelegate")
		return
	}
	// string "HealthCheckPeriod"
	o = append(o, 0xb1, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64)
	o = msgp.AppendDuration(o, z.HealthCheckPeriod)
	// string "RewardRate"
	o = append(o, 0xaa, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x52, 0x61, 0x74, 0x65)
	o = msgp.AppendFloat64(o, z.RewardRate)
	// string "ShareRatio"
	o = append(o, 0xaa, 0x53, 0x68, 0x61, 0x72, 0x65, 0x52, 0x61, 0x74, 0x69, 0x6f)
	o = msgp.AppendFloat64(o, z.ShareRatio)
	// string "BlockReward"
	o = append(o, 0xab, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64)
	o, err = z.BlockReward.MarshalMsg(o)
	if err != nil {
		err = msgp.WrapError(err, "BlockReward")
		return
	}
	// string "MaxCharge"
	o = append(o, 0xa9, 0x4d, 0x61, 0x78, 0x43, 0x68, 0x61, 0x72, 0x67, 0x65)
	o = msgp.AppendFloat64(o, z.MaxCharge)
	// string "Epoch"
	o = append(o, 0xa5, 0x45, 0x70, 0x6f, 0x63, 0x68)
	o = msgp.AppendInt64(o, z.Epoch)
	// string "RewardDeclineRate"
	o = append(o, 0xb1, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x44, 0x65, 0x63, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x61, 0x74, 0x65)
	o = msgp.AppendFloat64(o, z.RewardDeclineRate)
	// string "MaxMint"
	o = append(o, 0xa7, 0x4d, 0x61, 0x78, 0x4d, 0x69, 0x6e, 0x74)
	o, err = z.MaxMint.MarshalMsg(o)
	if err != nil {
		err = msgp.WrapError(err, "MaxMint")
		return
	}
	// string "NumMinerDelegatesRewarded"
	o = append(o, 0xb9, 0x4e, 0x75, 0x6d, 0x4d, 0x69, 0x6e, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x65, 0x64)
	o = msgp.AppendInt(o, z.NumMinerDelegatesRewarded)
	// string "NumShardersRewarded"
	o = append(o, 0xb3, 0x4e, 0x75, 0x6d, 0x53, 0x68, 0x61, 0x72, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x65, 0x64)
	o = msgp.AppendInt(o, z.NumShardersRewarded)
	// string "NumSharderDelegatesRewarded"
	o = append(o, 0xbb, 0x4e, 0x75, 0x6d, 0x53, 0x68, 0x61, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x65, 0x64)
	o = msgp.AppendInt(o, z.NumSharderDelegatesRewarded)
	// string "PrevMagicBlock"
	o = append(o, 0xae, 0x50, 0x72, 0x65, 0x76, 0x4d, 0x61, 0x67, 0x69, 0x63, 0x42, 0x6c, 0x6f, 0x63, 0x6b)
	if z.PrevMagicBlock == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.PrevMagicBlock.MarshalMsg(o)
		if err != nil {
			err = msgp.WrapError(err, "PrevMagicBlock")
			return
		}
	}
	// string "Minted"
	o = append(o, 0xa6, 0x4d, 0x69, 0x6e, 0x74, 0x65, 0x64)
	o, err = z.Minted.MarshalMsg(o)
	if err != nil {
		err = msgp.WrapError(err, "Minted")
		return
	}
	// string "RewardRoundFrequency"
	o = append(o, 0xb4, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x46, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79)
	o = msgp.AppendInt64(o, z.RewardRoundFrequency)
	// string "OwnerId"
	o = append(o, 0xa7, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64)
	o = msgp.AppendString(o, z.OwnerId)
	// string "CooldownPeriod"
	o = append(o, 0xae, 0x43, 0x6f, 0x6f, 0x6c, 0x64, 0x6f, 0x77, 0x6e, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64)
	o = msgp.AppendInt64(o, z.CooldownPeriod)
	// string "Cost"
	o = append(o, 0xa4, 0x43, 0x6f, 0x73, 0x74)
	o = msgp.AppendMapHeader(o, uint32(len(z.Cost)))
	keys_za0001 := make([]string, 0, len(z.Cost))
	for k := range z.Cost {
		keys_za0001 = append(keys_za0001, k)
	}
	msgp.Sort(keys_za0001)
	for _, k := range keys_za0001 {
		za0002 := z.Cost[k]
		o = msgp.AppendString(o, k)
		o = msgp.AppendInt(o, za0002)
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *GlobalNode) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "ViewChange":
			z.ViewChange, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "ViewChange")
				return
			}
		case "MaxN":
			z.MaxN, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "MaxN")
				return
			}
		case "MinN":
			z.MinN, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "MinN")
				return
			}
		case "MaxS":
			z.MaxS, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "MaxS")
				return
			}
		case "MinS":
			z.MinS, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "MinS")
				return
			}
		case "MaxDelegates":
			z.MaxDelegates, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "MaxDelegates")
				return
			}
		case "TPercent":
			z.TPercent, bts, err = msgp.ReadFloat64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "TPercent")
				return
			}
		case "KPercent":
			z.KPercent, bts, err = msgp.ReadFloat64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "KPercent")
				return
			}
		case "XPercent":
			z.XPercent, bts, err = msgp.ReadFloat64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "XPercent")
				return
			}
		case "LastRound":
			z.LastRound, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "LastRound")
				return
			}
		case "MaxStake":
			bts, err = z.MaxStake.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "MaxStake")
				return
			}
		case "MinStake":
			bts, err = z.MinStake.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "MinStake")
				return
			}
		case "MinStakePerDelegate":
			bts, err = z.MinStakePerDelegate.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "MinStakePerDelegate")
				return
			}
		case "HealthCheckPeriod":
			z.HealthCheckPeriod, bts, err = msgp.ReadDurationBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "HealthCheckPeriod")
				return
			}
		case "RewardRate":
			z.RewardRate, bts, err = msgp.ReadFloat64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "RewardRate")
				return
			}
		case "ShareRatio":
			z.ShareRatio, bts, err = msgp.ReadFloat64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "ShareRatio")
				return
			}
		case "BlockReward":
			bts, err = z.BlockReward.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "BlockReward")
				return
			}
		case "MaxCharge":
			z.MaxCharge, bts, err = msgp.ReadFloat64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "MaxCharge")
				return
			}
		case "Epoch":
			z.Epoch, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Epoch")
				return
			}
		case "RewardDeclineRate":
			z.RewardDeclineRate, bts, err = msgp.ReadFloat64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "RewardDeclineRate")
				return
			}
		case "MaxMint":
			bts, err = z.MaxMint.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "MaxMint")
				return
			}
		case "NumMinerDelegatesRewarded":
			z.NumMinerDelegatesRewarded, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "NumMinerDelegatesRewarded")
				return
			}
		case "NumShardersRewarded":
			z.NumShardersRewarded, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "NumShardersRewarded")
				return
			}
		case "NumSharderDelegatesRewarded":
			z.NumSharderDelegatesRewarded, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "NumSharderDelegatesRewarded")
				return
			}
		case "PrevMagicBlock":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.PrevMagicBlock = nil
			} else {
				if z.PrevMagicBlock == nil {
					z.PrevMagicBlock = new(block.MagicBlock)
				}
				bts, err = z.PrevMagicBlock.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "PrevMagicBlock")
					return
				}
			}
		case "Minted":
			bts, err = z.Minted.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "Minted")
				return
			}
		case "RewardRoundFrequency":
			z.RewardRoundFrequency, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "RewardRoundFrequency")
				return
			}
		case "OwnerId":
			z.OwnerId, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "OwnerId")
				return
			}
		case "CooldownPeriod":
			z.CooldownPeriod, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "CooldownPeriod")
				return
			}
		case "Cost":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Cost")
				return
			}
			if z.Cost == nil {
				z.Cost = make(map[string]int, zb0002)
			} else if len(z.Cost) > 0 {
				for key := range z.Cost {
					delete(z.Cost, key)
				}
			}
			for zb0002 > 0 {
				var za0001 string
				var za0002 int
				zb0002--
				za0001, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "Cost")
					return
				}
				za0002, bts, err = msgp.ReadIntBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "Cost", za0001)
					return
				}
				z.Cost[za0001] = za0002
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
func (z *GlobalNode) Msgsize() (s int) {
	s = 3 + 11 + msgp.Int64Size + 5 + msgp.IntSize + 5 + msgp.IntSize + 5 + msgp.IntSize + 5 + msgp.IntSize + 13 + msgp.IntSize + 9 + msgp.Float64Size + 9 + msgp.Float64Size + 9 + msgp.Float64Size + 10 + msgp.Int64Size + 9 + z.MaxStake.Msgsize() + 9 + z.MinStake.Msgsize() + 20 + z.MinStakePerDelegate.Msgsize() + 18 + msgp.DurationSize + 11 + msgp.Float64Size + 11 + msgp.Float64Size + 12 + z.BlockReward.Msgsize() + 10 + msgp.Float64Size + 6 + msgp.Int64Size + 18 + msgp.Float64Size + 8 + z.MaxMint.Msgsize() + 26 + msgp.IntSize + 20 + msgp.IntSize + 28 + msgp.IntSize + 15
	if z.PrevMagicBlock == nil {
		s += msgp.NilSize
	} else {
		s += z.PrevMagicBlock.Msgsize()
	}
	s += 7 + z.Minted.Msgsize() + 21 + msgp.Int64Size + 8 + msgp.StringPrefixSize + len(z.OwnerId) + 15 + msgp.Int64Size + 5 + msgp.MapHeaderSize
	if z.Cost != nil {
		for za0001, za0002 := range z.Cost {
			_ = za0002
			s += msgp.StringPrefixSize + len(za0001) + msgp.IntSize
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z NodeIDs) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendArrayHeader(o, uint32(len(z)))
	for za0001 := range z {
		o = msgp.AppendString(o, z[za0001])
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *NodeIDs) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0002 uint32
	zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if cap((*z)) >= int(zb0002) {
		(*z) = (*z)[:zb0002]
	} else {
		(*z) = make(NodeIDs, zb0002)
	}
	for zb0001 := range *z {
		(*z)[zb0001], bts, err = msgp.ReadStringBytes(bts)
		if err != nil {
			err = msgp.WrapError(err, zb0001)
			return
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z NodeIDs) Msgsize() (s int) {
	s = msgp.ArrayHeaderSize
	for zb0003 := range z {
		s += msgp.StringPrefixSize + len(z[zb0003])
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z NodeType) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendInt(o, int(z))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *NodeType) UnmarshalMsg(bts []byte) (o []byte, err error) {
	{
		var zb0001 int
		zb0001, bts, err = msgp.ReadIntBytes(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		(*z) = NodeType(zb0001)
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z NodeType) Msgsize() (s int) {
	s = msgp.IntSize
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Phase) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendInt(o, int(z))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Phase) UnmarshalMsg(bts []byte) (o []byte, err error) {
	{
		var zb0001 int
		zb0001, bts, err = msgp.ReadIntBytes(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		(*z) = Phase(zb0001)
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Phase) Msgsize() (s int) {
	s = msgp.IntSize
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *PhaseNode) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 4
	// string "Phase"
	o = append(o, 0x84, 0xa5, 0x50, 0x68, 0x61, 0x73, 0x65)
	o = msgp.AppendInt(o, int(z.Phase))
	// string "StartRound"
	o = append(o, 0xaa, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x6f, 0x75, 0x6e, 0x64)
	o = msgp.AppendInt64(o, z.StartRound)
	// string "CurrentRound"
	o = append(o, 0xac, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x52, 0x6f, 0x75, 0x6e, 0x64)
	o = msgp.AppendInt64(o, z.CurrentRound)
	// string "Restarts"
	o = append(o, 0xa8, 0x52, 0x65, 0x73, 0x74, 0x61, 0x72, 0x74, 0x73)
	o = msgp.AppendInt64(o, z.Restarts)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *PhaseNode) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "Phase":
			{
				var zb0002 int
				zb0002, bts, err = msgp.ReadIntBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "Phase")
					return
				}
				z.Phase = Phase(zb0002)
			}
		case "StartRound":
			z.StartRound, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "StartRound")
				return
			}
		case "CurrentRound":
			z.CurrentRound, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "CurrentRound")
				return
			}
		case "Restarts":
			z.Restarts, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Restarts")
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
func (z *PhaseNode) Msgsize() (s int) {
	s = 1 + 6 + msgp.IntSize + 11 + msgp.Int64Size + 13 + msgp.Int64Size + 9 + msgp.Int64Size
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *SimpleNode) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 13
	// string "Provider"
	o = append(o, 0x8d, 0xa8, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72)
	o, err = z.Provider.MarshalMsg(o)
	if err != nil {
		err = msgp.WrapError(err, "Provider")
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
	// string "PublicKey"
	o = append(o, 0xa9, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79)
	o = msgp.AppendString(o, z.PublicKey)
	// string "ShortName"
	o = append(o, 0xa9, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.ShortName)
	// string "BuildTag"
	o = append(o, 0xa8, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x54, 0x61, 0x67)
	o = msgp.AppendString(o, z.BuildTag)
	// string "TotalStaked"
	o = append(o, 0xab, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x6b, 0x65, 0x64)
	o, err = z.TotalStaked.MarshalMsg(o)
	if err != nil {
		err = msgp.WrapError(err, "TotalStaked")
		return
	}
	// string "Delete"
	o = append(o, 0xa6, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65)
	o = msgp.AppendBool(o, z.Delete)
	// string "NodeType"
	o = append(o, 0xa8, 0x4e, 0x6f, 0x64, 0x65, 0x54, 0x79, 0x70, 0x65)
	o = msgp.AppendInt(o, int(z.NodeType))
	// string "LastHealthCheck"
	o = append(o, 0xaf, 0x4c, 0x61, 0x73, 0x74, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b)
	o, err = z.LastHealthCheck.MarshalMsg(o)
	if err != nil {
		err = msgp.WrapError(err, "LastHealthCheck")
		return
	}
	// string "LastSettingUpdateRound"
	o = append(o, 0xb6, 0x4c, 0x61, 0x73, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x75, 0x6e, 0x64)
	o = msgp.AppendInt64(o, z.LastSettingUpdateRound)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *SimpleNode) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "Provider":
			bts, err = z.Provider.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "Provider")
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
		case "PublicKey":
			z.PublicKey, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "PublicKey")
				return
			}
		case "ShortName":
			z.ShortName, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "ShortName")
				return
			}
		case "BuildTag":
			z.BuildTag, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "BuildTag")
				return
			}
		case "TotalStaked":
			bts, err = z.TotalStaked.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "TotalStaked")
				return
			}
		case "Delete":
			z.Delete, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Delete")
				return
			}
		case "NodeType":
			{
				var zb0002 int
				zb0002, bts, err = msgp.ReadIntBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "NodeType")
					return
				}
				z.NodeType = NodeType(zb0002)
			}
		case "LastHealthCheck":
			bts, err = z.LastHealthCheck.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "LastHealthCheck")
				return
			}
		case "LastSettingUpdateRound":
			z.LastSettingUpdateRound, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "LastSettingUpdateRound")
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
func (z *SimpleNode) Msgsize() (s int) {
	s = 1 + 9 + z.Provider.Msgsize() + 8 + msgp.StringPrefixSize + len(z.N2NHost) + 5 + msgp.StringPrefixSize + len(z.Host) + 5 + msgp.IntSize + 5 + msgp.StringPrefixSize + len(z.Path) + 10 + msgp.StringPrefixSize + len(z.PublicKey) + 10 + msgp.StringPrefixSize + len(z.ShortName) + 9 + msgp.StringPrefixSize + len(z.BuildTag) + 12 + z.TotalStaked.Msgsize() + 7 + msgp.BoolSize + 9 + msgp.IntSize + 16 + z.LastHealthCheck.Msgsize() + 23 + msgp.Int64Size
	return
}

// MarshalMsg implements msgp.Marshaler
func (z SimpleNodes) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendMapHeader(o, uint32(len(z)))
	keys_za0001 := make([]string, 0, len(z))
	for k := range z {
		keys_za0001 = append(keys_za0001, k)
	}
	msgp.Sort(keys_za0001)
	for _, k := range keys_za0001 {
		za0002 := z[k]
		o = msgp.AppendString(o, k)
		if za0002 == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = za0002.MarshalMsg(o)
			if err != nil {
				err = msgp.WrapError(err, k)
				return
			}
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *SimpleNodes) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0003 uint32
	zb0003, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if (*z) == nil {
		(*z) = make(SimpleNodes, zb0003)
	} else if len((*z)) > 0 {
		for key := range *z {
			delete((*z), key)
		}
	}
	for zb0003 > 0 {
		var zb0001 string
		var zb0002 *SimpleNode
		zb0003--
		zb0001, bts, err = msgp.ReadStringBytes(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		if msgp.IsNil(bts) {
			bts, err = msgp.ReadNilBytes(bts)
			if err != nil {
				return
			}
			zb0002 = nil
		} else {
			if zb0002 == nil {
				zb0002 = new(SimpleNode)
			}
			bts, err = zb0002.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, zb0001)
				return
			}
		}
		(*z)[zb0001] = zb0002
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z SimpleNodes) Msgsize() (s int) {
	s = msgp.MapHeaderSize
	if z != nil {
		for zb0004, zb0005 := range z {
			_ = zb0005
			s += msgp.StringPrefixSize + len(zb0004)
			if zb0005 == nil {
				s += msgp.NilSize
			} else {
				s += zb0005.Msgsize()
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z ViewChangeLock) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "DeleteViewChangeSet"
	o = append(o, 0x83, 0xb3, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x56, 0x69, 0x65, 0x77, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x65, 0x74)
	o = msgp.AppendBool(o, z.DeleteViewChangeSet)
	// string "DeleteVC"
	o = append(o, 0xa8, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x56, 0x43)
	o = msgp.AppendInt64(o, z.DeleteVC)
	// string "Owner"
	o = append(o, 0xa5, 0x4f, 0x77, 0x6e, 0x65, 0x72)
	o = msgp.AppendString(o, z.Owner)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ViewChangeLock) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "DeleteViewChangeSet":
			z.DeleteViewChangeSet, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "DeleteViewChangeSet")
				return
			}
		case "DeleteVC":
			z.DeleteVC, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "DeleteVC")
				return
			}
		case "Owner":
			z.Owner, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Owner")
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
func (z ViewChangeLock) Msgsize() (s int) {
	s = 1 + 20 + msgp.BoolSize + 9 + msgp.Int64Size + 6 + msgp.StringPrefixSize + len(z.Owner)
	return
}
