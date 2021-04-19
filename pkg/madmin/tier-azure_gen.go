package madmin

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *TierAzure) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Endpoint":
			z.Endpoint, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Endpoint")
				return
			}
		case "AccountName":
			z.AccountName, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "AccountName")
				return
			}
		case "AccountKey":
			z.AccountKey, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "AccountKey")
				return
			}
		case "Bucket":
			z.Bucket, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Bucket")
				return
			}
		case "Prefix":
			z.Prefix, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Prefix")
				return
			}
		case "Region":
			z.Region, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Region")
				return
			}
		case "StorageClass":
			z.StorageClass, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "StorageClass")
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *TierAzure) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 7
	// write "Endpoint"
	err = en.Append(0x87, 0xa8, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74)
	if err != nil {
		return
	}
	err = en.WriteString(z.Endpoint)
	if err != nil {
		err = msgp.WrapError(err, "Endpoint")
		return
	}
	// write "AccountName"
	err = en.Append(0xab, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65)
	if err != nil {
		return
	}
	err = en.WriteString(z.AccountName)
	if err != nil {
		err = msgp.WrapError(err, "AccountName")
		return
	}
	// write "AccountKey"
	err = en.Append(0xaa, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4b, 0x65, 0x79)
	if err != nil {
		return
	}
	err = en.WriteString(z.AccountKey)
	if err != nil {
		err = msgp.WrapError(err, "AccountKey")
		return
	}
	// write "Bucket"
	err = en.Append(0xa6, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74)
	if err != nil {
		return
	}
	err = en.WriteString(z.Bucket)
	if err != nil {
		err = msgp.WrapError(err, "Bucket")
		return
	}
	// write "Prefix"
	err = en.Append(0xa6, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78)
	if err != nil {
		return
	}
	err = en.WriteString(z.Prefix)
	if err != nil {
		err = msgp.WrapError(err, "Prefix")
		return
	}
	// write "Region"
	err = en.Append(0xa6, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e)
	if err != nil {
		return
	}
	err = en.WriteString(z.Region)
	if err != nil {
		err = msgp.WrapError(err, "Region")
		return
	}
	// write "StorageClass"
	err = en.Append(0xac, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73)
	if err != nil {
		return
	}
	err = en.WriteString(z.StorageClass)
	if err != nil {
		err = msgp.WrapError(err, "StorageClass")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *TierAzure) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 7
	// string "Endpoint"
	o = append(o, 0x87, 0xa8, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74)
	o = msgp.AppendString(o, z.Endpoint)
	// string "AccountName"
	o = append(o, 0xab, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.AccountName)
	// string "AccountKey"
	o = append(o, 0xaa, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4b, 0x65, 0x79)
	o = msgp.AppendString(o, z.AccountKey)
	// string "Bucket"
	o = append(o, 0xa6, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74)
	o = msgp.AppendString(o, z.Bucket)
	// string "Prefix"
	o = append(o, 0xa6, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78)
	o = msgp.AppendString(o, z.Prefix)
	// string "Region"
	o = append(o, 0xa6, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e)
	o = msgp.AppendString(o, z.Region)
	// string "StorageClass"
	o = append(o, 0xac, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73)
	o = msgp.AppendString(o, z.StorageClass)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *TierAzure) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "Endpoint":
			z.Endpoint, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Endpoint")
				return
			}
		case "AccountName":
			z.AccountName, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "AccountName")
				return
			}
		case "AccountKey":
			z.AccountKey, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "AccountKey")
				return
			}
		case "Bucket":
			z.Bucket, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Bucket")
				return
			}
		case "Prefix":
			z.Prefix, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Prefix")
				return
			}
		case "Region":
			z.Region, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Region")
				return
			}
		case "StorageClass":
			z.StorageClass, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "StorageClass")
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
func (z *TierAzure) Msgsize() (s int) {
	s = 1 + 9 + msgp.StringPrefixSize + len(z.Endpoint) + 12 + msgp.StringPrefixSize + len(z.AccountName) + 11 + msgp.StringPrefixSize + len(z.AccountKey) + 7 + msgp.StringPrefixSize + len(z.Bucket) + 7 + msgp.StringPrefixSize + len(z.Prefix) + 7 + msgp.StringPrefixSize + len(z.Region) + 13 + msgp.StringPrefixSize + len(z.StorageClass)
	return
}
