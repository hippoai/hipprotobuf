package hipprotobuf

import (
	google_protobuf "github.com/golang/protobuf/ptypes/wrappers"
)

// GetBool returns nil or a pointer to the primitive value
func GetBool(in *google_protobuf.BoolValue) *bool {
	if in == nil {
		return nil
	}

	return &in.Value
}

// GetInt32 returns nil or a pointer to the primitive value
func GetInt32(in *google_protobuf.Int32Value) *int32 {
	if in == nil {
		return nil
	}

	return &in.Value
}

// GetInt64 returns nil or a pointer to the primitive value
func GetInt64(in *google_protobuf.Int64Value) *int64 {
	if in == nil {
		return nil
	}

	return &in.Value
}

// GetFloat returns nil or a pointer to the primitive value
func GetFloat(in *google_protobuf.FloatValue) *float32 {
	if in == nil {
		return nil
	}

	return &in.Value
}

// GetDouble returns nil or a pointer to the primitive value
func GetDouble(in *google_protobuf.DoubleValue) *float64 {
	if in == nil {
		return nil
	}

	return &in.Value
}

// GetBytes returns nil or a pointer to the primitive value
func GetBytes(in *google_protobuf.BytesValue) *[]byte {
	if in == nil {
		return nil
	}

	return &in.Value
}

// GetString returns nil or a pointer to the primitive value
func GetString(in *google_protobuf.StringValue) *string {
	if in == nil {
		return nil
	}

	return &in.Value
}

// GetUInt32 returns nil or a pointer to the primitive value
func GetUInt32(in *google_protobuf.UInt32Value) *uint32 {
	if in == nil {
		return nil
	}

	return &in.Value
}

// GetUInt64 returns nil or a pointer to the primitive value
func GetUInt64(in *google_protobuf.UInt32Value) *uint32 {
	if in == nil {
		return nil
	}

	return &in.Value
}

// MakeBool makes a wrapper
func MakeBool(in *bool) *google_protobuf.BoolValue {
	if in == nil {
		return nil
	}

	return &google_protobuf.BoolValue{
		Value: *in,
	}
}

// MakeInt32 makes a wrapper
func MakeInt32(in *int32) *google_protobuf.Int32Value {
	if in == nil {
		return nil
	}

	return &google_protobuf.Int32Value{
		Value: *in,
	}
}

// MakeInt64 makes a wrapper
func MakeInt64(in *int64) *google_protobuf.Int64Value {
	if in == nil {
		return nil
	}

	return &google_protobuf.Int64Value{
		Value: *in,
	}
}

// MakeFloat makes a wrapper
func MakeFloat(in *float32) *google_protobuf.FloatValue {
	if in == nil {
		return nil
	}

	return &google_protobuf.FloatValue{
		Value: *in,
	}
}

// MakeDouble makes a wrapper
func MakeDouble(in *float64) *google_protobuf.DoubleValue {
	if in == nil {
		return nil
	}

	return &google_protobuf.DoubleValue{
		Value: *in,
	}
}

// MakeBytes makes a wrapper
func MakeBytes(in *[]byte) *google_protobuf.BytesValue {
	if in == nil {
		return nil
	}

	return &google_protobuf.BytesValue{
		Value: *in,
	}
}

// MakeString makes a wrapper
func MakeString(in *string) *google_protobuf.StringValue {
	if in == nil {
		return nil
	}

	return &google_protobuf.StringValue{
		Value: *in,
	}
}

// MakeUInt32 makes a wrapper
func MakeUInt32(in *uint32) *google_protobuf.UInt32Value {
	if in == nil {
		return nil
	}

	return &google_protobuf.UInt32Value{
		Value: *in,
	}
}

// MakeUInt64 makes a wrapper
func MakeUInt64(in *uint64) *google_protobuf.UInt64Value {
	if in == nil {
		return nil
	}

	return &google_protobuf.UInt64Value{
		Value: *in,
	}
}
