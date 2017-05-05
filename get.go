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

// GetBoolOrDefault returns nil or a pointer to the primitive value
func GetBoolOrDefault(in *google_protobuf.BoolValue, defaultValue bool) *bool {
	if in == nil {
		return &defaultValue
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

// GetInt32OrDefault returns nil or a pointer to the primitive value
func GetInt32OrDefault(in *google_protobuf.Int32Value, defaultValue int32) *int32 {
	if in == nil {
		return &defaultValue
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

// GetInt64OrDefault returns nil or a pointer to the primitive value
func GetInt64OrDefault(in *google_protobuf.Int64Value, defaultValue int64) *int64 {
	if in == nil {
		return &defaultValue
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

// GetFloatOrDefault returns nil or a pointer to the primitive value
func GetFloatOrDefault(in *google_protobuf.FloatValue, defaultValue float32) *float32 {
	if in == nil {
		return &defaultValue
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

// GetDoubleOrDefault returns nil or a pointer to the primitive value
func GetDoubleOrDefault(in *google_protobuf.DoubleValue, defaultValue float64) *float64 {
	if in == nil {
		return &defaultValue
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

// GetBytesOrDefault returns nil or a pointer to the primitive value
func GetBytesOrDefault(in *google_protobuf.BytesValue, defaultValue []byte) *[]byte {
	if in == nil {
		return &defaultValue
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

// GetStringOrDefault returns nil or a pointer to the primitive value
func GetStringOrDefault(in *google_protobuf.StringValue, defaultValue string) *string {
	if in == nil {
		return &defaultValue
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

// GetUInt32OrDefault returns nil or a pointer to the primitive value
func GetUInt32OrDefault(in *google_protobuf.UInt32Value, defaultValue uint32) *uint32 {
	if in == nil {
		return &defaultValue
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

// GetUInt64OrDefault returns nil or a pointer to the primitive value
func GetUInt64OrDefault(in *google_protobuf.UInt64Value, defaultValue uint64) *uint64 {
	if in == nil {
		return &defaultValue
	}

	return &in.Value
}
