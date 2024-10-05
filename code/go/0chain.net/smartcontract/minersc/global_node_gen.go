package minersc

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"0chain.net/chaincore/block"
	"github.com/tinylib/msgp/msgp"
)

// MarshalMsg implements msgp.Marshaler
func (z *globalNodeBase) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 28
	// string "ViewChange"
	o = append(o, 0xde, 0x0, 0x1c, 0xaa, 0x56, 0x69, 0x65, 0x77, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65)
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
func (z *globalNodeBase) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
func (z *globalNodeBase) Msgsize() (s int) {
	s = 3 + 11 + msgp.Int64Size + 5 + msgp.IntSize + 5 + msgp.IntSize + 5 + msgp.IntSize + 5 + msgp.IntSize + 13 + msgp.IntSize + 9 + msgp.Float64Size + 9 + msgp.Float64Size + 9 + msgp.Float64Size + 10 + msgp.Int64Size + 9 + z.MaxStake.Msgsize() + 9 + z.MinStake.Msgsize() + 20 + z.MinStakePerDelegate.Msgsize() + 18 + msgp.DurationSize + 11 + msgp.Float64Size + 11 + msgp.Float64Size + 12 + z.BlockReward.Msgsize() + 10 + msgp.Float64Size + 6 + msgp.Int64Size + 18 + msgp.Float64Size + 26 + msgp.IntSize + 20 + msgp.IntSize + 28 + msgp.IntSize + 15
	if z.PrevMagicBlock == nil {
		s += msgp.NilSize
	} else {
		s += z.PrevMagicBlock.Msgsize()
	}
	s += 21 + msgp.Int64Size + 8 + msgp.StringPrefixSize + len(z.OwnerId) + 15 + msgp.Int64Size + 5 + msgp.MapHeaderSize
	if z.Cost != nil {
		for za0001, za0002 := range z.Cost {
			_ = za0002
			s += msgp.StringPrefixSize + len(za0001) + msgp.IntSize
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *globalNodeV1) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 28
	// string "ViewChange"
	o = append(o, 0xde, 0x0, 0x1c, 0xaa, 0x56, 0x69, 0x65, 0x77, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65)
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
func (z *globalNodeV1) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
func (z *globalNodeV1) Msgsize() (s int) {
	s = 3 + 11 + msgp.Int64Size + 5 + msgp.IntSize + 5 + msgp.IntSize + 5 + msgp.IntSize + 5 + msgp.IntSize + 13 + msgp.IntSize + 9 + msgp.Float64Size + 9 + msgp.Float64Size + 9 + msgp.Float64Size + 10 + msgp.Int64Size + 9 + z.MaxStake.Msgsize() + 9 + z.MinStake.Msgsize() + 20 + z.MinStakePerDelegate.Msgsize() + 18 + msgp.DurationSize + 11 + msgp.Float64Size + 11 + msgp.Float64Size + 12 + z.BlockReward.Msgsize() + 10 + msgp.Float64Size + 6 + msgp.Int64Size + 18 + msgp.Float64Size + 26 + msgp.IntSize + 20 + msgp.IntSize + 28 + msgp.IntSize + 15
	if z.PrevMagicBlock == nil {
		s += msgp.NilSize
	} else {
		s += z.PrevMagicBlock.Msgsize()
	}
	s += 21 + msgp.Int64Size + 8 + msgp.StringPrefixSize + len(z.OwnerId) + 15 + msgp.Int64Size + 5 + msgp.MapHeaderSize
	if z.Cost != nil {
		for za0001, za0002 := range z.Cost {
			_ = za0002
			s += msgp.StringPrefixSize + len(za0001) + msgp.IntSize
		}
	}
	return
}