# Helpers for Google's Protobufs

We provide methods to mimic null or omitted values for primitives in proto3 by using a combination of:
* Google's protobuf wrappers: github.com/golang/protobuf/ptypes/wrappers
* Getters methods (cf. file `./empty.go`) that return nil for an empty wrapper, a pointer to the primitive otherwise
