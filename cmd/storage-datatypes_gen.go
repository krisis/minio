package cmd

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *DiskInfo) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0001 uint32
	zb0001, err = dc.ReadArrayHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 11 {
		err = msgp.ArrayError{Wanted: 11, Got: zb0001}
		return
	}
	z.Total, err = dc.ReadUint64()
	if err != nil {
		err = msgp.WrapError(err, "Total")
		return
	}
	z.Free, err = dc.ReadUint64()
	if err != nil {
		err = msgp.WrapError(err, "Free")
		return
	}
	z.Used, err = dc.ReadUint64()
	if err != nil {
		err = msgp.WrapError(err, "Used")
		return
	}
	z.UsedInodes, err = dc.ReadUint64()
	if err != nil {
		err = msgp.WrapError(err, "UsedInodes")
		return
	}
	z.FSType, err = dc.ReadString()
	if err != nil {
		err = msgp.WrapError(err, "FSType")
		return
	}
	z.RootDisk, err = dc.ReadBool()
	if err != nil {
		err = msgp.WrapError(err, "RootDisk")
		return
	}
	z.Healing, err = dc.ReadBool()
	if err != nil {
		err = msgp.WrapError(err, "Healing")
		return
	}
	z.Endpoint, err = dc.ReadString()
	if err != nil {
		err = msgp.WrapError(err, "Endpoint")
		return
	}
	z.MountPath, err = dc.ReadString()
	if err != nil {
		err = msgp.WrapError(err, "MountPath")
		return
	}
	z.ID, err = dc.ReadString()
	if err != nil {
		err = msgp.WrapError(err, "ID")
		return
	}
	z.Error, err = dc.ReadString()
	if err != nil {
		err = msgp.WrapError(err, "Error")
		return
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *DiskInfo) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 11
	err = en.Append(0x9b)
	if err != nil {
		return
	}
	err = en.WriteUint64(z.Total)
	if err != nil {
		err = msgp.WrapError(err, "Total")
		return
	}
	err = en.WriteUint64(z.Free)
	if err != nil {
		err = msgp.WrapError(err, "Free")
		return
	}
	err = en.WriteUint64(z.Used)
	if err != nil {
		err = msgp.WrapError(err, "Used")
		return
	}
	err = en.WriteUint64(z.UsedInodes)
	if err != nil {
		err = msgp.WrapError(err, "UsedInodes")
		return
	}
	err = en.WriteString(z.FSType)
	if err != nil {
		err = msgp.WrapError(err, "FSType")
		return
	}
	err = en.WriteBool(z.RootDisk)
	if err != nil {
		err = msgp.WrapError(err, "RootDisk")
		return
	}
	err = en.WriteBool(z.Healing)
	if err != nil {
		err = msgp.WrapError(err, "Healing")
		return
	}
	err = en.WriteString(z.Endpoint)
	if err != nil {
		err = msgp.WrapError(err, "Endpoint")
		return
	}
	err = en.WriteString(z.MountPath)
	if err != nil {
		err = msgp.WrapError(err, "MountPath")
		return
	}
	err = en.WriteString(z.ID)
	if err != nil {
		err = msgp.WrapError(err, "ID")
		return
	}
	err = en.WriteString(z.Error)
	if err != nil {
		err = msgp.WrapError(err, "Error")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *DiskInfo) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 11
	o = append(o, 0x9b)
	o = msgp.AppendUint64(o, z.Total)
	o = msgp.AppendUint64(o, z.Free)
	o = msgp.AppendUint64(o, z.Used)
	o = msgp.AppendUint64(o, z.UsedInodes)
	o = msgp.AppendString(o, z.FSType)
	o = msgp.AppendBool(o, z.RootDisk)
	o = msgp.AppendBool(o, z.Healing)
	o = msgp.AppendString(o, z.Endpoint)
	o = msgp.AppendString(o, z.MountPath)
	o = msgp.AppendString(o, z.ID)
	o = msgp.AppendString(o, z.Error)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *DiskInfo) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 11 {
		err = msgp.ArrayError{Wanted: 11, Got: zb0001}
		return
	}
	z.Total, bts, err = msgp.ReadUint64Bytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "Total")
		return
	}
	z.Free, bts, err = msgp.ReadUint64Bytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "Free")
		return
	}
	z.Used, bts, err = msgp.ReadUint64Bytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "Used")
		return
	}
	z.UsedInodes, bts, err = msgp.ReadUint64Bytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "UsedInodes")
		return
	}
	z.FSType, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "FSType")
		return
	}
	z.RootDisk, bts, err = msgp.ReadBoolBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "RootDisk")
		return
	}
	z.Healing, bts, err = msgp.ReadBoolBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "Healing")
		return
	}
	z.Endpoint, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "Endpoint")
		return
	}
	z.MountPath, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "MountPath")
		return
	}
	z.ID, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "ID")
		return
	}
	z.Error, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "Error")
		return
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *DiskInfo) Msgsize() (s int) {
	s = 1 + msgp.Uint64Size + msgp.Uint64Size + msgp.Uint64Size + msgp.Uint64Size + msgp.StringPrefixSize + len(z.FSType) + msgp.BoolSize + msgp.BoolSize + msgp.StringPrefixSize + len(z.Endpoint) + msgp.StringPrefixSize + len(z.MountPath) + msgp.StringPrefixSize + len(z.ID) + msgp.StringPrefixSize + len(z.Error)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *FileInfo) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0001 uint32
	zb0001, err = dc.ReadArrayHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 20 {
		err = msgp.ArrayError{Wanted: 20, Got: zb0001}
		return
	}
	z.Volume, err = dc.ReadString()
	if err != nil {
		err = msgp.WrapError(err, "Volume")
		return
	}
	z.Name, err = dc.ReadString()
	if err != nil {
		err = msgp.WrapError(err, "Name")
		return
	}
	z.VersionID, err = dc.ReadString()
	if err != nil {
		err = msgp.WrapError(err, "VersionID")
		return
	}
	z.IsLatest, err = dc.ReadBool()
	if err != nil {
		err = msgp.WrapError(err, "IsLatest")
		return
	}
	z.Deleted, err = dc.ReadBool()
	if err != nil {
		err = msgp.WrapError(err, "Deleted")
		return
	}
	z.TransitionStatus, err = dc.ReadString()
	if err != nil {
		err = msgp.WrapError(err, "TransitionStatus")
		return
	}
	z.TransitionedObjName, err = dc.ReadString()
	if err != nil {
		err = msgp.WrapError(err, "TransitionedObjName")
		return
	}
	z.TransitionTier, err = dc.ReadString()
	if err != nil {
		err = msgp.WrapError(err, "TransitionTier")
		return
	}
	z.DataDir, err = dc.ReadString()
	if err != nil {
		err = msgp.WrapError(err, "DataDir")
		return
	}
	z.XLV1, err = dc.ReadBool()
	if err != nil {
		err = msgp.WrapError(err, "XLV1")
		return
	}
	z.ModTime, err = dc.ReadTime()
	if err != nil {
		err = msgp.WrapError(err, "ModTime")
		return
	}
	z.Size, err = dc.ReadInt64()
	if err != nil {
		err = msgp.WrapError(err, "Size")
		return
	}
	z.Mode, err = dc.ReadUint32()
	if err != nil {
		err = msgp.WrapError(err, "Mode")
		return
	}
	var zb0002 uint32
	zb0002, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err, "Metadata")
		return
	}
	if z.Metadata == nil {
		z.Metadata = make(map[string]string, zb0002)
	} else if len(z.Metadata) > 0 {
		for key := range z.Metadata {
			delete(z.Metadata, key)
		}
	}
	for zb0002 > 0 {
		zb0002--
		var za0001 string
		var za0002 string
		za0001, err = dc.ReadString()
		if err != nil {
			err = msgp.WrapError(err, "Metadata")
			return
		}
		za0002, err = dc.ReadString()
		if err != nil {
			err = msgp.WrapError(err, "Metadata", za0001)
			return
		}
		z.Metadata[za0001] = za0002
	}
	var zb0003 uint32
	zb0003, err = dc.ReadArrayHeader()
	if err != nil {
		err = msgp.WrapError(err, "Parts")
		return
	}
	if cap(z.Parts) >= int(zb0003) {
		z.Parts = (z.Parts)[:zb0003]
	} else {
		z.Parts = make([]ObjectPartInfo, zb0003)
	}
	for za0003 := range z.Parts {
		err = z.Parts[za0003].DecodeMsg(dc)
		if err != nil {
			err = msgp.WrapError(err, "Parts", za0003)
			return
		}
	}
	err = z.Erasure.DecodeMsg(dc)
	if err != nil {
		err = msgp.WrapError(err, "Erasure")
		return
	}
	z.MarkDeleted, err = dc.ReadBool()
	if err != nil {
		err = msgp.WrapError(err, "MarkDeleted")
		return
	}
	z.DeleteMarkerReplicationStatus, err = dc.ReadString()
	if err != nil {
		err = msgp.WrapError(err, "DeleteMarkerReplicationStatus")
		return
	}
	{
		var zb0004 string
		zb0004, err = dc.ReadString()
		if err != nil {
			err = msgp.WrapError(err, "VersionPurgeStatus")
			return
		}
		z.VersionPurgeStatus = VersionPurgeStatusType(zb0004)
	}
	z.Data, err = dc.ReadBytes(z.Data)
	if err != nil {
		err = msgp.WrapError(err, "Data")
		return
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *FileInfo) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 20
	err = en.Append(0xdc, 0x0, 0x14)
	if err != nil {
		return
	}
	err = en.WriteString(z.Volume)
	if err != nil {
		err = msgp.WrapError(err, "Volume")
		return
	}
	err = en.WriteString(z.Name)
	if err != nil {
		err = msgp.WrapError(err, "Name")
		return
	}
	err = en.WriteString(z.VersionID)
	if err != nil {
		err = msgp.WrapError(err, "VersionID")
		return
	}
	err = en.WriteBool(z.IsLatest)
	if err != nil {
		err = msgp.WrapError(err, "IsLatest")
		return
	}
	err = en.WriteBool(z.Deleted)
	if err != nil {
		err = msgp.WrapError(err, "Deleted")
		return
	}
	err = en.WriteString(z.TransitionStatus)
	if err != nil {
		err = msgp.WrapError(err, "TransitionStatus")
		return
	}
	err = en.WriteString(z.TransitionedObjName)
	if err != nil {
		err = msgp.WrapError(err, "TransitionedObjName")
		return
	}
	err = en.WriteString(z.TransitionTier)
	if err != nil {
		err = msgp.WrapError(err, "TransitionTier")
		return
	}
	err = en.WriteString(z.DataDir)
	if err != nil {
		err = msgp.WrapError(err, "DataDir")
		return
	}
	err = en.WriteBool(z.XLV1)
	if err != nil {
		err = msgp.WrapError(err, "XLV1")
		return
	}
	err = en.WriteTime(z.ModTime)
	if err != nil {
		err = msgp.WrapError(err, "ModTime")
		return
	}
	err = en.WriteInt64(z.Size)
	if err != nil {
		err = msgp.WrapError(err, "Size")
		return
	}
	err = en.WriteUint32(z.Mode)
	if err != nil {
		err = msgp.WrapError(err, "Mode")
		return
	}
	err = en.WriteMapHeader(uint32(len(z.Metadata)))
	if err != nil {
		err = msgp.WrapError(err, "Metadata")
		return
	}
	for za0001, za0002 := range z.Metadata {
		err = en.WriteString(za0001)
		if err != nil {
			err = msgp.WrapError(err, "Metadata")
			return
		}
		err = en.WriteString(za0002)
		if err != nil {
			err = msgp.WrapError(err, "Metadata", za0001)
			return
		}
	}
	err = en.WriteArrayHeader(uint32(len(z.Parts)))
	if err != nil {
		err = msgp.WrapError(err, "Parts")
		return
	}
	for za0003 := range z.Parts {
		err = z.Parts[za0003].EncodeMsg(en)
		if err != nil {
			err = msgp.WrapError(err, "Parts", za0003)
			return
		}
	}
	err = z.Erasure.EncodeMsg(en)
	if err != nil {
		err = msgp.WrapError(err, "Erasure")
		return
	}
	err = en.WriteBool(z.MarkDeleted)
	if err != nil {
		err = msgp.WrapError(err, "MarkDeleted")
		return
	}
	err = en.WriteString(z.DeleteMarkerReplicationStatus)
	if err != nil {
		err = msgp.WrapError(err, "DeleteMarkerReplicationStatus")
		return
	}
	err = en.WriteString(string(z.VersionPurgeStatus))
	if err != nil {
		err = msgp.WrapError(err, "VersionPurgeStatus")
		return
	}
	err = en.WriteBytes(z.Data)
	if err != nil {
		err = msgp.WrapError(err, "Data")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *FileInfo) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 20
	o = append(o, 0xdc, 0x0, 0x14)
	o = msgp.AppendString(o, z.Volume)
	o = msgp.AppendString(o, z.Name)
	o = msgp.AppendString(o, z.VersionID)
	o = msgp.AppendBool(o, z.IsLatest)
	o = msgp.AppendBool(o, z.Deleted)
	o = msgp.AppendString(o, z.TransitionStatus)
	o = msgp.AppendString(o, z.TransitionedObjName)
	o = msgp.AppendString(o, z.TransitionTier)
	o = msgp.AppendString(o, z.DataDir)
	o = msgp.AppendBool(o, z.XLV1)
	o = msgp.AppendTime(o, z.ModTime)
	o = msgp.AppendInt64(o, z.Size)
	o = msgp.AppendUint32(o, z.Mode)
	o = msgp.AppendMapHeader(o, uint32(len(z.Metadata)))
	for za0001, za0002 := range z.Metadata {
		o = msgp.AppendString(o, za0001)
		o = msgp.AppendString(o, za0002)
	}
	o = msgp.AppendArrayHeader(o, uint32(len(z.Parts)))
	for za0003 := range z.Parts {
		o, err = z.Parts[za0003].MarshalMsg(o)
		if err != nil {
			err = msgp.WrapError(err, "Parts", za0003)
			return
		}
	}
	o, err = z.Erasure.MarshalMsg(o)
	if err != nil {
		err = msgp.WrapError(err, "Erasure")
		return
	}
	o = msgp.AppendBool(o, z.MarkDeleted)
	o = msgp.AppendString(o, z.DeleteMarkerReplicationStatus)
	o = msgp.AppendString(o, string(z.VersionPurgeStatus))
	o = msgp.AppendBytes(o, z.Data)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *FileInfo) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 20 {
		err = msgp.ArrayError{Wanted: 20, Got: zb0001}
		return
	}
	z.Volume, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "Volume")
		return
	}
	z.Name, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "Name")
		return
	}
	z.VersionID, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "VersionID")
		return
	}
	z.IsLatest, bts, err = msgp.ReadBoolBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "IsLatest")
		return
	}
	z.Deleted, bts, err = msgp.ReadBoolBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "Deleted")
		return
	}
	z.TransitionStatus, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "TransitionStatus")
		return
	}
	z.TransitionedObjName, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "TransitionedObjName")
		return
	}
	z.TransitionTier, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "TransitionTier")
		return
	}
	z.DataDir, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "DataDir")
		return
	}
	z.XLV1, bts, err = msgp.ReadBoolBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "XLV1")
		return
	}
	z.ModTime, bts, err = msgp.ReadTimeBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "ModTime")
		return
	}
	z.Size, bts, err = msgp.ReadInt64Bytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "Size")
		return
	}
	z.Mode, bts, err = msgp.ReadUint32Bytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "Mode")
		return
	}
	var zb0002 uint32
	zb0002, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "Metadata")
		return
	}
	if z.Metadata == nil {
		z.Metadata = make(map[string]string, zb0002)
	} else if len(z.Metadata) > 0 {
		for key := range z.Metadata {
			delete(z.Metadata, key)
		}
	}
	for zb0002 > 0 {
		var za0001 string
		var za0002 string
		zb0002--
		za0001, bts, err = msgp.ReadStringBytes(bts)
		if err != nil {
			err = msgp.WrapError(err, "Metadata")
			return
		}
		za0002, bts, err = msgp.ReadStringBytes(bts)
		if err != nil {
			err = msgp.WrapError(err, "Metadata", za0001)
			return
		}
		z.Metadata[za0001] = za0002
	}
	var zb0003 uint32
	zb0003, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "Parts")
		return
	}
	if cap(z.Parts) >= int(zb0003) {
		z.Parts = (z.Parts)[:zb0003]
	} else {
		z.Parts = make([]ObjectPartInfo, zb0003)
	}
	for za0003 := range z.Parts {
		bts, err = z.Parts[za0003].UnmarshalMsg(bts)
		if err != nil {
			err = msgp.WrapError(err, "Parts", za0003)
			return
		}
	}
	bts, err = z.Erasure.UnmarshalMsg(bts)
	if err != nil {
		err = msgp.WrapError(err, "Erasure")
		return
	}
	z.MarkDeleted, bts, err = msgp.ReadBoolBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "MarkDeleted")
		return
	}
	z.DeleteMarkerReplicationStatus, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "DeleteMarkerReplicationStatus")
		return
	}
	{
		var zb0004 string
		zb0004, bts, err = msgp.ReadStringBytes(bts)
		if err != nil {
			err = msgp.WrapError(err, "VersionPurgeStatus")
			return
		}
		z.VersionPurgeStatus = VersionPurgeStatusType(zb0004)
	}
	z.Data, bts, err = msgp.ReadBytesBytes(bts, z.Data)
	if err != nil {
		err = msgp.WrapError(err, "Data")
		return
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *FileInfo) Msgsize() (s int) {
	s = 3 + msgp.StringPrefixSize + len(z.Volume) + msgp.StringPrefixSize + len(z.Name) + msgp.StringPrefixSize + len(z.VersionID) + msgp.BoolSize + msgp.BoolSize + msgp.StringPrefixSize + len(z.TransitionStatus) + msgp.StringPrefixSize + len(z.TransitionedObjName) + msgp.StringPrefixSize + len(z.TransitionTier) + msgp.StringPrefixSize + len(z.DataDir) + msgp.BoolSize + msgp.TimeSize + msgp.Int64Size + msgp.Uint32Size + msgp.MapHeaderSize
	if z.Metadata != nil {
		for za0001, za0002 := range z.Metadata {
			_ = za0002
			s += msgp.StringPrefixSize + len(za0001) + msgp.StringPrefixSize + len(za0002)
		}
	}
	s += msgp.ArrayHeaderSize
	for za0003 := range z.Parts {
		s += z.Parts[za0003].Msgsize()
	}
	s += z.Erasure.Msgsize() + msgp.BoolSize + msgp.StringPrefixSize + len(z.DeleteMarkerReplicationStatus) + msgp.StringPrefixSize + len(string(z.VersionPurgeStatus)) + msgp.BytesPrefixSize + len(z.Data)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *FileInfoVersions) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "Volume":
			z.Volume, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Volume")
				return
			}
		case "Name":
			z.Name, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Name")
				return
			}
		case "IsEmptyDir":
			z.IsEmptyDir, err = dc.ReadBool()
			if err != nil {
				err = msgp.WrapError(err, "IsEmptyDir")
				return
			}
		case "LatestModTime":
			z.LatestModTime, err = dc.ReadTime()
			if err != nil {
				err = msgp.WrapError(err, "LatestModTime")
				return
			}
		case "Versions":
			var zb0002 uint32
			zb0002, err = dc.ReadArrayHeader()
			if err != nil {
				err = msgp.WrapError(err, "Versions")
				return
			}
			if cap(z.Versions) >= int(zb0002) {
				z.Versions = (z.Versions)[:zb0002]
			} else {
				z.Versions = make([]FileInfo, zb0002)
			}
			for za0001 := range z.Versions {
				err = z.Versions[za0001].DecodeMsg(dc)
				if err != nil {
					err = msgp.WrapError(err, "Versions", za0001)
					return
				}
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
func (z *FileInfoVersions) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 5
	// write "Volume"
	err = en.Append(0x85, 0xa6, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65)
	if err != nil {
		return
	}
	err = en.WriteString(z.Volume)
	if err != nil {
		err = msgp.WrapError(err, "Volume")
		return
	}
	// write "Name"
	err = en.Append(0xa4, 0x4e, 0x61, 0x6d, 0x65)
	if err != nil {
		return
	}
	err = en.WriteString(z.Name)
	if err != nil {
		err = msgp.WrapError(err, "Name")
		return
	}
	// write "IsEmptyDir"
	err = en.Append(0xaa, 0x49, 0x73, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x44, 0x69, 0x72)
	if err != nil {
		return
	}
	err = en.WriteBool(z.IsEmptyDir)
	if err != nil {
		err = msgp.WrapError(err, "IsEmptyDir")
		return
	}
	// write "LatestModTime"
	err = en.Append(0xad, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x54, 0x69, 0x6d, 0x65)
	if err != nil {
		return
	}
	err = en.WriteTime(z.LatestModTime)
	if err != nil {
		err = msgp.WrapError(err, "LatestModTime")
		return
	}
	// write "Versions"
	err = en.Append(0xa8, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.Versions)))
	if err != nil {
		err = msgp.WrapError(err, "Versions")
		return
	}
	for za0001 := range z.Versions {
		err = z.Versions[za0001].EncodeMsg(en)
		if err != nil {
			err = msgp.WrapError(err, "Versions", za0001)
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *FileInfoVersions) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 5
	// string "Volume"
	o = append(o, 0x85, 0xa6, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Volume)
	// string "Name"
	o = append(o, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Name)
	// string "IsEmptyDir"
	o = append(o, 0xaa, 0x49, 0x73, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x44, 0x69, 0x72)
	o = msgp.AppendBool(o, z.IsEmptyDir)
	// string "LatestModTime"
	o = append(o, 0xad, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x54, 0x69, 0x6d, 0x65)
	o = msgp.AppendTime(o, z.LatestModTime)
	// string "Versions"
	o = append(o, 0xa8, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Versions)))
	for za0001 := range z.Versions {
		o, err = z.Versions[za0001].MarshalMsg(o)
		if err != nil {
			err = msgp.WrapError(err, "Versions", za0001)
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *FileInfoVersions) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "Volume":
			z.Volume, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Volume")
				return
			}
		case "Name":
			z.Name, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Name")
				return
			}
		case "IsEmptyDir":
			z.IsEmptyDir, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "IsEmptyDir")
				return
			}
		case "LatestModTime":
			z.LatestModTime, bts, err = msgp.ReadTimeBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "LatestModTime")
				return
			}
		case "Versions":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Versions")
				return
			}
			if cap(z.Versions) >= int(zb0002) {
				z.Versions = (z.Versions)[:zb0002]
			} else {
				z.Versions = make([]FileInfo, zb0002)
			}
			for za0001 := range z.Versions {
				bts, err = z.Versions[za0001].UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "Versions", za0001)
					return
				}
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
func (z *FileInfoVersions) Msgsize() (s int) {
	s = 1 + 7 + msgp.StringPrefixSize + len(z.Volume) + 5 + msgp.StringPrefixSize + len(z.Name) + 11 + msgp.BoolSize + 14 + msgp.TimeSize + 9 + msgp.ArrayHeaderSize
	for za0001 := range z.Versions {
		s += z.Versions[za0001].Msgsize()
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *FilesInfo) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "Files":
			var zb0002 uint32
			zb0002, err = dc.ReadArrayHeader()
			if err != nil {
				err = msgp.WrapError(err, "Files")
				return
			}
			if cap(z.Files) >= int(zb0002) {
				z.Files = (z.Files)[:zb0002]
			} else {
				z.Files = make([]FileInfo, zb0002)
			}
			for za0001 := range z.Files {
				err = z.Files[za0001].DecodeMsg(dc)
				if err != nil {
					err = msgp.WrapError(err, "Files", za0001)
					return
				}
			}
		case "IsTruncated":
			z.IsTruncated, err = dc.ReadBool()
			if err != nil {
				err = msgp.WrapError(err, "IsTruncated")
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
func (z *FilesInfo) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "Files"
	err = en.Append(0x82, 0xa5, 0x46, 0x69, 0x6c, 0x65, 0x73)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.Files)))
	if err != nil {
		err = msgp.WrapError(err, "Files")
		return
	}
	for za0001 := range z.Files {
		err = z.Files[za0001].EncodeMsg(en)
		if err != nil {
			err = msgp.WrapError(err, "Files", za0001)
			return
		}
	}
	// write "IsTruncated"
	err = en.Append(0xab, 0x49, 0x73, 0x54, 0x72, 0x75, 0x6e, 0x63, 0x61, 0x74, 0x65, 0x64)
	if err != nil {
		return
	}
	err = en.WriteBool(z.IsTruncated)
	if err != nil {
		err = msgp.WrapError(err, "IsTruncated")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *FilesInfo) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Files"
	o = append(o, 0x82, 0xa5, 0x46, 0x69, 0x6c, 0x65, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Files)))
	for za0001 := range z.Files {
		o, err = z.Files[za0001].MarshalMsg(o)
		if err != nil {
			err = msgp.WrapError(err, "Files", za0001)
			return
		}
	}
	// string "IsTruncated"
	o = append(o, 0xab, 0x49, 0x73, 0x54, 0x72, 0x75, 0x6e, 0x63, 0x61, 0x74, 0x65, 0x64)
	o = msgp.AppendBool(o, z.IsTruncated)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *FilesInfo) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "Files":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Files")
				return
			}
			if cap(z.Files) >= int(zb0002) {
				z.Files = (z.Files)[:zb0002]
			} else {
				z.Files = make([]FileInfo, zb0002)
			}
			for za0001 := range z.Files {
				bts, err = z.Files[za0001].UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "Files", za0001)
					return
				}
			}
		case "IsTruncated":
			z.IsTruncated, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "IsTruncated")
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
func (z *FilesInfo) Msgsize() (s int) {
	s = 1 + 6 + msgp.ArrayHeaderSize
	for za0001 := range z.Files {
		s += z.Files[za0001].Msgsize()
	}
	s += 12 + msgp.BoolSize
	return
}

// DecodeMsg implements msgp.Decodable
func (z *FilesInfoVersions) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "FilesVersions":
			var zb0002 uint32
			zb0002, err = dc.ReadArrayHeader()
			if err != nil {
				err = msgp.WrapError(err, "FilesVersions")
				return
			}
			if cap(z.FilesVersions) >= int(zb0002) {
				z.FilesVersions = (z.FilesVersions)[:zb0002]
			} else {
				z.FilesVersions = make([]FileInfoVersions, zb0002)
			}
			for za0001 := range z.FilesVersions {
				err = z.FilesVersions[za0001].DecodeMsg(dc)
				if err != nil {
					err = msgp.WrapError(err, "FilesVersions", za0001)
					return
				}
			}
		case "IsTruncated":
			z.IsTruncated, err = dc.ReadBool()
			if err != nil {
				err = msgp.WrapError(err, "IsTruncated")
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
func (z *FilesInfoVersions) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "FilesVersions"
	err = en.Append(0x82, 0xad, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.FilesVersions)))
	if err != nil {
		err = msgp.WrapError(err, "FilesVersions")
		return
	}
	for za0001 := range z.FilesVersions {
		err = z.FilesVersions[za0001].EncodeMsg(en)
		if err != nil {
			err = msgp.WrapError(err, "FilesVersions", za0001)
			return
		}
	}
	// write "IsTruncated"
	err = en.Append(0xab, 0x49, 0x73, 0x54, 0x72, 0x75, 0x6e, 0x63, 0x61, 0x74, 0x65, 0x64)
	if err != nil {
		return
	}
	err = en.WriteBool(z.IsTruncated)
	if err != nil {
		err = msgp.WrapError(err, "IsTruncated")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *FilesInfoVersions) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "FilesVersions"
	o = append(o, 0x82, 0xad, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.FilesVersions)))
	for za0001 := range z.FilesVersions {
		o, err = z.FilesVersions[za0001].MarshalMsg(o)
		if err != nil {
			err = msgp.WrapError(err, "FilesVersions", za0001)
			return
		}
	}
	// string "IsTruncated"
	o = append(o, 0xab, 0x49, 0x73, 0x54, 0x72, 0x75, 0x6e, 0x63, 0x61, 0x74, 0x65, 0x64)
	o = msgp.AppendBool(o, z.IsTruncated)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *FilesInfoVersions) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "FilesVersions":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "FilesVersions")
				return
			}
			if cap(z.FilesVersions) >= int(zb0002) {
				z.FilesVersions = (z.FilesVersions)[:zb0002]
			} else {
				z.FilesVersions = make([]FileInfoVersions, zb0002)
			}
			for za0001 := range z.FilesVersions {
				bts, err = z.FilesVersions[za0001].UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "FilesVersions", za0001)
					return
				}
			}
		case "IsTruncated":
			z.IsTruncated, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "IsTruncated")
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
func (z *FilesInfoVersions) Msgsize() (s int) {
	s = 1 + 14 + msgp.ArrayHeaderSize
	for za0001 := range z.FilesVersions {
		s += z.FilesVersions[za0001].Msgsize()
	}
	s += 12 + msgp.BoolSize
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VersionPurgeStatusType) DecodeMsg(dc *msgp.Reader) (err error) {
	{
		var zb0001 string
		zb0001, err = dc.ReadString()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		(*z) = VersionPurgeStatusType(zb0001)
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z VersionPurgeStatusType) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteString(string(z))
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z VersionPurgeStatusType) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendString(o, string(z))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VersionPurgeStatusType) UnmarshalMsg(bts []byte) (o []byte, err error) {
	{
		var zb0001 string
		zb0001, bts, err = msgp.ReadStringBytes(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		(*z) = VersionPurgeStatusType(zb0001)
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z VersionPurgeStatusType) Msgsize() (s int) {
	s = msgp.StringPrefixSize + len(string(z))
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VolInfo) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0001 uint32
	zb0001, err = dc.ReadArrayHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 2 {
		err = msgp.ArrayError{Wanted: 2, Got: zb0001}
		return
	}
	z.Name, err = dc.ReadString()
	if err != nil {
		err = msgp.WrapError(err, "Name")
		return
	}
	z.Created, err = dc.ReadTime()
	if err != nil {
		err = msgp.WrapError(err, "Created")
		return
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z VolInfo) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 2
	err = en.Append(0x92)
	if err != nil {
		return
	}
	err = en.WriteString(z.Name)
	if err != nil {
		err = msgp.WrapError(err, "Name")
		return
	}
	err = en.WriteTime(z.Created)
	if err != nil {
		err = msgp.WrapError(err, "Created")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z VolInfo) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 2
	o = append(o, 0x92)
	o = msgp.AppendString(o, z.Name)
	o = msgp.AppendTime(o, z.Created)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VolInfo) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 2 {
		err = msgp.ArrayError{Wanted: 2, Got: zb0001}
		return
	}
	z.Name, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "Name")
		return
	}
	z.Created, bts, err = msgp.ReadTimeBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "Created")
		return
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z VolInfo) Msgsize() (s int) {
	s = 1 + msgp.StringPrefixSize + len(z.Name) + msgp.TimeSize
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VolsInfo) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0002 uint32
	zb0002, err = dc.ReadArrayHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if cap((*z)) >= int(zb0002) {
		(*z) = (*z)[:zb0002]
	} else {
		(*z) = make(VolsInfo, zb0002)
	}
	for zb0001 := range *z {
		var zb0003 uint32
		zb0003, err = dc.ReadArrayHeader()
		if err != nil {
			err = msgp.WrapError(err, zb0001)
			return
		}
		if zb0003 != 2 {
			err = msgp.ArrayError{Wanted: 2, Got: zb0003}
			return
		}
		(*z)[zb0001].Name, err = dc.ReadString()
		if err != nil {
			err = msgp.WrapError(err, zb0001, "Name")
			return
		}
		(*z)[zb0001].Created, err = dc.ReadTime()
		if err != nil {
			err = msgp.WrapError(err, zb0001, "Created")
			return
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z VolsInfo) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteArrayHeader(uint32(len(z)))
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0004 := range z {
		// array header, size 2
		err = en.Append(0x92)
		if err != nil {
			return
		}
		err = en.WriteString(z[zb0004].Name)
		if err != nil {
			err = msgp.WrapError(err, zb0004, "Name")
			return
		}
		err = en.WriteTime(z[zb0004].Created)
		if err != nil {
			err = msgp.WrapError(err, zb0004, "Created")
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z VolsInfo) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendArrayHeader(o, uint32(len(z)))
	for zb0004 := range z {
		// array header, size 2
		o = append(o, 0x92)
		o = msgp.AppendString(o, z[zb0004].Name)
		o = msgp.AppendTime(o, z[zb0004].Created)
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VolsInfo) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0002 uint32
	zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if cap((*z)) >= int(zb0002) {
		(*z) = (*z)[:zb0002]
	} else {
		(*z) = make(VolsInfo, zb0002)
	}
	for zb0001 := range *z {
		var zb0003 uint32
		zb0003, bts, err = msgp.ReadArrayHeaderBytes(bts)
		if err != nil {
			err = msgp.WrapError(err, zb0001)
			return
		}
		if zb0003 != 2 {
			err = msgp.ArrayError{Wanted: 2, Got: zb0003}
			return
		}
		(*z)[zb0001].Name, bts, err = msgp.ReadStringBytes(bts)
		if err != nil {
			err = msgp.WrapError(err, zb0001, "Name")
			return
		}
		(*z)[zb0001].Created, bts, err = msgp.ReadTimeBytes(bts)
		if err != nil {
			err = msgp.WrapError(err, zb0001, "Created")
			return
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z VolsInfo) Msgsize() (s int) {
	s = msgp.ArrayHeaderSize
	for zb0004 := range z {
		s += 1 + msgp.StringPrefixSize + len(z[zb0004].Name) + msgp.TimeSize
	}
	return
}
