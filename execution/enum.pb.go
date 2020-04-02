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
// source: execution/enum.proto

package execution

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

type WorkflowExecutionStatus int32

const (
	WorkflowExecutionStatusUnknown        WorkflowExecutionStatus = 0
	WorkflowExecutionStatusRunning        WorkflowExecutionStatus = 1
	WorkflowExecutionStatusCompleted      WorkflowExecutionStatus = 2
	WorkflowExecutionStatusFailed         WorkflowExecutionStatus = 3
	WorkflowExecutionStatusCanceled       WorkflowExecutionStatus = 4
	WorkflowExecutionStatusTerminated     WorkflowExecutionStatus = 5
	WorkflowExecutionStatusContinuedAsNew WorkflowExecutionStatus = 6
	WorkflowExecutionStatusTimedOut       WorkflowExecutionStatus = 7
)

var WorkflowExecutionStatus_name = map[int32]string{
	0: "WorkflowExecutionStatusUnknown",
	1: "WorkflowExecutionStatusRunning",
	2: "WorkflowExecutionStatusCompleted",
	3: "WorkflowExecutionStatusFailed",
	4: "WorkflowExecutionStatusCanceled",
	5: "WorkflowExecutionStatusTerminated",
	6: "WorkflowExecutionStatusContinuedAsNew",
	7: "WorkflowExecutionStatusTimedOut",
}

var WorkflowExecutionStatus_value = map[string]int32{
	"WorkflowExecutionStatusUnknown":        0,
	"WorkflowExecutionStatusRunning":        1,
	"WorkflowExecutionStatusCompleted":      2,
	"WorkflowExecutionStatusFailed":         3,
	"WorkflowExecutionStatusCanceled":       4,
	"WorkflowExecutionStatusTerminated":     5,
	"WorkflowExecutionStatusContinuedAsNew": 6,
	"WorkflowExecutionStatusTimedOut":       7,
}

func (WorkflowExecutionStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2297211ce2f9cb26, []int{0}
}

type PendingActivityState int32

const (
	PendingActivityStateScheduled       PendingActivityState = 0
	PendingActivityStateStarted         PendingActivityState = 1
	PendingActivityStateCancelRequested PendingActivityState = 2
)

var PendingActivityState_name = map[int32]string{
	0: "PendingActivityStateScheduled",
	1: "PendingActivityStateStarted",
	2: "PendingActivityStateCancelRequested",
}

var PendingActivityState_value = map[string]int32{
	"PendingActivityStateScheduled":       0,
	"PendingActivityStateStarted":         1,
	"PendingActivityStateCancelRequested": 2,
}

func (PendingActivityState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2297211ce2f9cb26, []int{1}
}

func init() {
	proto.RegisterEnum("execution.WorkflowExecutionStatus", WorkflowExecutionStatus_name, WorkflowExecutionStatus_value)
	proto.RegisterEnum("execution.PendingActivityState", PendingActivityState_name, PendingActivityState_value)
}

func init() { proto.RegisterFile("execution/enum.proto", fileDescriptor_2297211ce2f9cb26) }

var fileDescriptor_2297211ce2f9cb26 = []byte{
	// 347 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0xd2, 0xb1, 0x4e, 0x32, 0x41,
	0x10, 0x07, 0xf0, 0x5d, 0xbe, 0xef, 0xe3, 0x8b, 0x5b, 0x5d, 0x36, 0x24, 0x16, 0xc4, 0x41, 0x44,
	0x42, 0x24, 0x11, 0x0a, 0x9f, 0x00, 0x8d, 0x96, 0x4a, 0x00, 0x63, 0x62, 0x77, 0xde, 0x8d, 0xb8,
	0xe1, 0x6e, 0x16, 0x8f, 0x5d, 0xd1, 0xc4, 0xc2, 0xc4, 0x17, 0xf0, 0x31, 0x7c, 0x12, 0x63, 0x49,
	0x49, 0x29, 0x47, 0x63, 0xc9, 0x23, 0x98, 0x83, 0x40, 0x2c, 0x38, 0xba, 0xcb, 0x7f, 0x7e, 0xb9,
	0x99, 0xd9, 0x8c, 0xc8, 0xe1, 0x23, 0x7a, 0xd6, 0x28, 0x4d, 0x75, 0x24, 0x1b, 0xd6, 0xfa, 0x91,
	0x36, 0x5a, 0x6e, 0xad, 0xd2, 0xea, 0x47, 0x46, 0x6c, 0x5f, 0xe9, 0xa8, 0x77, 0x1b, 0xe8, 0xe1,
	0xe9, 0x32, 0x6d, 0x1b, 0xd7, 0xd8, 0x81, 0xdc, 0x13, 0x90, 0x52, 0xba, 0xa4, 0x1e, 0xe9, 0x21,
	0x39, 0x6c, 0x83, 0x69, 0x59, 0x22, 0x45, 0x5d, 0x87, 0xcb, 0x7d, 0xb1, 0x9b, 0x62, 0x4e, 0x74,
	0xd8, 0x0f, 0xd0, 0xa0, 0xef, 0x64, 0x64, 0x51, 0xec, 0xa4, 0xa8, 0x33, 0x57, 0x05, 0xe8, 0x3b,
	0x7f, 0x64, 0x49, 0x14, 0xd2, 0x7e, 0xe4, 0x92, 0x87, 0x09, 0xfa, 0x2b, 0xcb, 0xa2, 0x98, 0x82,
	0x3a, 0x18, 0x85, 0x8a, 0xdc, 0xa4, 0xdd, 0x3f, 0x79, 0x20, 0xca, 0xa9, 0x43, 0x91, 0x51, 0x64,
	0xd1, 0x6f, 0x0c, 0xce, 0x71, 0xe8, 0x64, 0x37, 0xb4, 0xed, 0xa8, 0x10, 0xfd, 0x0b, 0x6b, 0x9c,
	0xff, 0xd5, 0x57, 0x2e, 0x72, 0x4d, 0x24, 0x5f, 0x51, 0xb7, 0xe1, 0x19, 0xf5, 0xa0, 0xcc, 0x53,
	0x62, 0x30, 0xd9, 0x6b, 0x5d, 0xde, 0xf6, 0xee, 0xd0, 0xb7, 0xc9, 0xc8, 0x4c, 0x16, 0x44, 0x7e,
	0x2d, 0x31, 0x6e, 0x94, 0x0c, 0xcb, 0x65, 0x45, 0x94, 0xd6, 0x81, 0xc5, 0xd6, 0x2d, 0xbc, 0xb7,
	0x38, 0x98, 0x3f, 0xe2, 0xf1, 0xf3, 0x68, 0x02, 0x6c, 0x3c, 0x01, 0x36, 0x9b, 0x00, 0x7f, 0x89,
	0x81, 0xbf, 0xc7, 0xc0, 0x3f, 0x63, 0xe0, 0xa3, 0x18, 0xf8, 0x57, 0x0c, 0xfc, 0x3b, 0x06, 0x36,
	0x8b, 0x81, 0xbf, 0x4d, 0x81, 0x8d, 0xa6, 0xc0, 0xc6, 0x53, 0x60, 0x22, 0xaf, 0x74, 0xcd, 0x60,
	0xd8, 0xd7, 0x91, 0x1b, 0x2c, 0x4e, 0xa4, 0xb6, 0xba, 0x90, 0x26, 0xbf, 0xae, 0x74, 0x7f, 0x95,
	0x95, 0xae, 0x2f, 0xbf, 0x0f, 0xe7, 0xb4, 0xbe, 0xa2, 0x37, 0xd9, 0x79, 0x70, 0xf4, 0x13, 0x00,
	0x00, 0xff, 0xff, 0xce, 0x8b, 0x0c, 0xeb, 0x76, 0x02, 0x00, 0x00,
}

func (x WorkflowExecutionStatus) String() string {
	s, ok := WorkflowExecutionStatus_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x PendingActivityState) String() string {
	s, ok := PendingActivityState_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}