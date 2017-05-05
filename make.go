package hipprotobuf

import (
	google_protobuf "github.com/golang/protobuf/ptypes/wrappers"
)

// MakeBool makes a wrapper
func MakeBool(in *bool) *google_protobuf.BoolValue {
	if in == nil {
		return nil
	}

	return &google_protobuf.BoolValue{
		Value: *in,
	}
}

// MakeBoolOrDefault makes a wrapper
func MakeBoolOrDefault(in *bool, defaultValue bool) *google_protobuf.BoolValue {
	if in == nil {
		return &google_protobuf.BoolValue{
			Value: defaultValue,
		}
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

// MakeInt32OrDefault makes a wrapper
func MakeInt32OrDefault(in *int32, defaultValue int32) *google_protobuf.Int32Value {
	if in == nil {
		return &google_protobuf.Int32Value{
			Value: defaultValue,
		}
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

// MakeInt64OrDefault makes a wrapper
func MakeInt64OrDefault(in *int64, defaultValue int64) *google_protobuf.Int64Value {
	if in == nil {
		return &google_protobuf.Int64Value{
			Value: defaultValue,
		}
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

// MakeFloatOrDefault makes a wrapper
func MakeFloatOrDefault(in *float32, defaultValue float32) *google_protobuf.FloatValue {
	if in == nil {
		return &google_protobuf.FloatValue{
			Value: defaultValue,
		}
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

// MakeDoubleOrDefault makes a wrapper
func MakeDoubleOrDefault(in *float64, defaultValue float64) *google_protobuf.DoubleValue {
	if in == nil {
		return &google_protobuf.DoubleValue{
			Value: defaultValue,
		}
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

// MakeBytesOrDefault makes a wrapper
func MakeBytesOrDefault(in *[]byte, defaultValue []byte) *google_protobuf.BytesValue {
	if in == nil {
		return &google_protobuf.BytesValue{
			Value: defaultValue,
		}
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

// MakeStringOrDefault makes a wrapper
func MakeStringOrDefault(in *string, defaultValue string) *google_protobuf.StringValue {
	if in == nil {
		return &google_protobuf.StringValue{
			Value: defaultValue,
		}
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

// MakeUInt32OrDefault makes a wrapper
func MakeUInt32OrDefault(in *uint32, defaultValue uint32) *google_protobuf.UInt32Value {
	if in == nil {
		return &google_protobuf.UInt32Value{
			Value: defaultValue,
		}
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

// MakeUInt64OrDefault makes a wrapper
func MakeUInt64OrDefault(in *uint64, defaultValue uint64) *google_protobuf.UInt64Value {
	if in == nil {
		return &google_protobuf.UInt64Value{
			Value: defaultValue,
		}
	}

	return &google_protobuf.UInt64Value{
		Value: *in,
	}
}
