package hipprotobuf

import (
	google_protobuf "github.com/golang/protobuf/ptypes/wrappers"
)

// GetBool returns nil or the value
func GetBool(in *google_protobuf.BoolValue) *bool {
	if in == nil {
		return nil
	}

	return &in.Value
}

// GetInt32 returns nil or the value
func GetInt32(in *google_protobuf.Int32Value) *int32 {
	if in == nil {
		return nil
	}

	return &in.Value
}

// GetInt64 returns nil or the value
func GetInt64(in *google_protobuf.Int64Value) *int64 {
	if in == nil {
		return nil
	}

	return &in.Value
}

// GetFloat returns nil or the value
func GetFloat(in *google_protobuf.FloatValue) *float32 {
	if in == nil {
		return nil
	}

	return &in.Value
}

// GetString returns nil or the value
func GetString(in *google_protobuf.StringValue) *string {
	if in == nil {
		return nil
	}

	return &in.Value
}

// GetUInt32 returns nil or the value
func GetUInt32(in *google_protobuf.UInt32Value) *uint32 {
	if in == nil {
		return nil
	}

	return &in.Value
}

// GetUInt64 returns nil or the value
func GetUInt64(in *google_protobuf.UInt32Value) *uint32 {
	if in == nil {
		return nil
	}

	return &in.Value
}

// GetDouble returns nil or the value
func GetDouble(in *google_protobuf.DoubleValue) *float64 {
	if in == nil {
		return nil
	}

	return &in.Value
}

// GetBytes returns nil or the value
func GetBytes(in *google_protobuf.BytesValue) *[]byte {
	if in == nil {
		return nil
	}

	return &in.Value
}
