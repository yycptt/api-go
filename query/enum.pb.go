// The MIT License (MIT)
//
// Copyright (c) 2020 Temporal Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: query/enum.proto

package query

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	math "math"
	strconv "strconv"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type QueryTaskCompletedType int32

const (
	QueryTaskCompletedTypeCompleted QueryTaskCompletedType = 0
	QueryTaskCompletedTypeFailed    QueryTaskCompletedType = 1
)

var QueryTaskCompletedType_name = map[int32]string{
	0: "QueryTaskCompletedTypeCompleted",
	1: "QueryTaskCompletedTypeFailed",
}

var QueryTaskCompletedType_value = map[string]int32{
	"QueryTaskCompletedTypeCompleted": 0,
	"QueryTaskCompletedTypeFailed":    1,
}

func (QueryTaskCompletedType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4b1c9789806fe1e4, []int{0}
}

type QueryResultType int32

const (
	QueryResultTypeAnswered QueryResultType = 0
	QueryResultTypeFailed   QueryResultType = 1
)

var QueryResultType_name = map[int32]string{
	0: "QueryResultTypeAnswered",
	1: "QueryResultTypeFailed",
}

var QueryResultType_value = map[string]int32{
	"QueryResultTypeAnswered": 0,
	"QueryResultTypeFailed":   1,
}

func (QueryResultType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4b1c9789806fe1e4, []int{1}
}

type QueryRejectCondition int32

const (
	// None indicates that query should not be rejected.
	QueryRejectConditionNone QueryRejectCondition = 0
	// NotOpen indicates that query should be rejected if workflow is not open.
	QueryRejectConditionNotOpen QueryRejectCondition = 1
	// NotCompletedCleanly indicates that query should be rejected if workflow did not complete cleanly.
	QueryRejectConditionNotCompletedCleanly QueryRejectCondition = 2
)

var QueryRejectCondition_name = map[int32]string{
	0: "QueryRejectConditionNone",
	1: "QueryRejectConditionNotOpen",
	2: "QueryRejectConditionNotCompletedCleanly",
}

var QueryRejectCondition_value = map[string]int32{
	"QueryRejectConditionNone":                0,
	"QueryRejectConditionNotOpen":             1,
	"QueryRejectConditionNotCompletedCleanly": 2,
}

func (QueryRejectCondition) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4b1c9789806fe1e4, []int{2}
}

type QueryConsistencyLevel int32

const (
	// Eventual indicates that query should be eventually consistent.
	QueryConsistencyLevelEventual QueryConsistencyLevel = 0
	// Strong indicates that any events that came before query should be reflected in workflow state before running query.
	QueryConsistencyLevelStrong QueryConsistencyLevel = 1
)

var QueryConsistencyLevel_name = map[int32]string{
	0: "QueryConsistencyLevelEventual",
	1: "QueryConsistencyLevelStrong",
}

var QueryConsistencyLevel_value = map[string]int32{
	"QueryConsistencyLevelEventual": 0,
	"QueryConsistencyLevelStrong":   1,
}

func (QueryConsistencyLevel) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4b1c9789806fe1e4, []int{3}
}

func init() {
	proto.RegisterEnum("query.QueryTaskCompletedType", QueryTaskCompletedType_name, QueryTaskCompletedType_value)
	proto.RegisterEnum("query.QueryResultType", QueryResultType_name, QueryResultType_value)
	proto.RegisterEnum("query.QueryRejectCondition", QueryRejectCondition_name, QueryRejectCondition_value)
	proto.RegisterEnum("query.QueryConsistencyLevel", QueryConsistencyLevel_name, QueryConsistencyLevel_value)
}

func init() { proto.RegisterFile("query/enum.proto", fileDescriptor_4b1c9789806fe1e4) }

var fileDescriptor_4b1c9789806fe1e4 = []byte{
	// 327 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xbf, 0x4a, 0x2b, 0x41,
	0x14, 0xc6, 0x67, 0x2e, 0xdc, 0x5b, 0x9c, 0xe6, 0x0e, 0x83, 0x1a, 0x25, 0xf1, 0x44, 0x49, 0x21,
	0x44, 0x4c, 0x0a, 0x9f, 0x40, 0x17, 0x05, 0x41, 0xfc, 0x9b, 0x4a, 0x0b, 0x59, 0x93, 0x43, 0x58,
	0x9d, 0xcc, 0xac, 0xbb, 0xb3, 0x91, 0xed, 0xc4, 0x27, 0xf0, 0x31, 0x7c, 0x14, 0xcb, 0x94, 0x29,
	0xcd, 0xa4, 0xb1, 0xcc, 0x23, 0x88, 0x63, 0x0c, 0x21, 0xac, 0xdd, 0x9c, 0xef, 0xf7, 0xe3, 0xe3,
	0x0c, 0x07, 0xc4, 0x43, 0x46, 0x49, 0xde, 0x24, 0x9d, 0xf5, 0x1a, 0x71, 0x62, 0xac, 0x91, 0x7f,
	0x7d, 0x52, 0xbf, 0x81, 0x95, 0xf3, 0xaf, 0x47, 0x2b, 0x4c, 0xef, 0x03, 0xd3, 0x8b, 0x15, 0x59,
	0xea, 0xb4, 0xf2, 0x98, 0x64, 0x0d, 0xaa, 0xc5, 0x64, 0x36, 0x08, 0x26, 0x37, 0xa0, 0x52, 0x2c,
	0x1d, 0x86, 0x91, 0xa2, 0x8e, 0xe0, 0xf5, 0x23, 0xf8, 0xef, 0x8d, 0x0b, 0x4a, 0x33, 0x65, 0x7d,
	0x73, 0x19, 0x4a, 0x0b, 0xd1, 0x9e, 0x4e, 0x1f, 0x29, 0xf1, 0x8d, 0x6b, 0xb0, 0xbc, 0x00, 0x67,
	0x55, 0xcf, 0x1c, 0x96, 0xa6, 0xec, 0x8e, 0xda, 0x36, 0x30, 0xba, 0x13, 0xd9, 0xc8, 0x68, 0x59,
	0x81, 0xd5, 0xa2, 0xfc, 0xc4, 0x68, 0x12, 0x4c, 0x56, 0xa1, 0x5c, 0x4c, 0xed, 0x69, 0x4c, 0x5a,
	0x70, 0xb9, 0x0d, 0x5b, 0xbf, 0x08, 0xb3, 0x2f, 0x05, 0x8a, 0x42, 0xad, 0x72, 0xf1, 0xa7, 0x7e,
	0x3d, 0xdd, 0x2f, 0x30, 0x3a, 0x8d, 0x52, 0x4b, 0xba, 0x9d, 0x1f, 0x53, 0x9f, 0x94, 0xdc, 0x84,
	0xf5, 0x42, 0x70, 0xd0, 0x27, 0x6d, 0xb3, 0x50, 0xcd, 0x6d, 0xb2, 0xa8, 0x5c, 0xda, 0xc4, 0xe8,
	0xae, 0xe0, 0xfb, 0x76, 0x30, 0x42, 0x36, 0x1c, 0x21, 0x9b, 0x8c, 0x90, 0x3f, 0x39, 0xe4, 0xaf,
	0x0e, 0xf9, 0x9b, 0x43, 0x3e, 0x70, 0xc8, 0xdf, 0x1d, 0xf2, 0x0f, 0x87, 0x6c, 0xe2, 0x90, 0xbf,
	0x8c, 0x91, 0x0d, 0xc6, 0xc8, 0x86, 0x63, 0x64, 0x50, 0x8a, 0x4c, 0xc3, 0x52, 0x2f, 0x36, 0x49,
	0xa8, 0xbe, 0xaf, 0xdb, 0xf0, 0xc7, 0x3d, 0xe3, 0x57, 0xb5, 0xee, 0x1c, 0x8a, 0x4c, 0xf3, 0xe7,
	0xbd, 0xe3, 0xb5, 0xa6, 0xd7, 0x6e, 0xff, 0xf9, 0x61, 0xf7, 0x33, 0x00, 0x00, 0xff, 0xff, 0xa9,
	0xa4, 0xd3, 0x1f, 0x25, 0x02, 0x00, 0x00,
}

func (x QueryTaskCompletedType) String() string {
	s, ok := QueryTaskCompletedType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x QueryResultType) String() string {
	s, ok := QueryResultType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x QueryRejectCondition) String() string {
	s, ok := QueryRejectCondition_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x QueryConsistencyLevel) String() string {
	s, ok := QueryConsistencyLevel_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}