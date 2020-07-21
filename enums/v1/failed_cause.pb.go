// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: temporal/api/enums/v1/failed_cause.proto

package enums

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

type WorkflowTaskFailedCause int32

const (
	WORKFLOW_TASK_FAILED_CAUSE_UNSPECIFIED                                               WorkflowTaskFailedCause = 0
	WORKFLOW_TASK_FAILED_CAUSE_UNHANDLED_COMMAND                                         WorkflowTaskFailedCause = 1
	WORKFLOW_TASK_FAILED_CAUSE_BAD_SCHEDULE_ACTIVITY_ATTRIBUTES                          WorkflowTaskFailedCause = 2
	WORKFLOW_TASK_FAILED_CAUSE_BAD_REQUEST_CANCEL_ACTIVITY_ATTRIBUTES                    WorkflowTaskFailedCause = 3
	WORKFLOW_TASK_FAILED_CAUSE_BAD_START_TIMER_ATTRIBUTES                                WorkflowTaskFailedCause = 4
	WORKFLOW_TASK_FAILED_CAUSE_BAD_CANCEL_TIMER_ATTRIBUTES                               WorkflowTaskFailedCause = 5
	WORKFLOW_TASK_FAILED_CAUSE_BAD_RECORD_MARKER_ATTRIBUTES                              WorkflowTaskFailedCause = 6
	WORKFLOW_TASK_FAILED_CAUSE_BAD_COMPLETE_WORKFLOW_EXECUTION_ATTRIBUTES                WorkflowTaskFailedCause = 7
	WORKFLOW_TASK_FAILED_CAUSE_BAD_FAIL_WORKFLOW_EXECUTION_ATTRIBUTES                    WorkflowTaskFailedCause = 8
	WORKFLOW_TASK_FAILED_CAUSE_BAD_CANCEL_WORKFLOW_EXECUTION_ATTRIBUTES                  WorkflowTaskFailedCause = 9
	WORKFLOW_TASK_FAILED_CAUSE_BAD_REQUEST_CANCEL_EXTERNAL_WORKFLOW_EXECUTION_ATTRIBUTES WorkflowTaskFailedCause = 10
	WORKFLOW_TASK_FAILED_CAUSE_BAD_CONTINUE_AS_NEW_ATTRIBUTES                            WorkflowTaskFailedCause = 11
	WORKFLOW_TASK_FAILED_CAUSE_START_TIMER_DUPLICATE_ID                                  WorkflowTaskFailedCause = 12
	WORKFLOW_TASK_FAILED_CAUSE_RESET_STICKY_TASK_QUEUE                                   WorkflowTaskFailedCause = 13
	WORKFLOW_TASK_FAILED_CAUSE_WORKFLOW_WORKER_UNHANDLED_FAILURE                         WorkflowTaskFailedCause = 14
	WORKFLOW_TASK_FAILED_CAUSE_BAD_SIGNAL_WORKFLOW_EXECUTION_ATTRIBUTES                  WorkflowTaskFailedCause = 15
	WORKFLOW_TASK_FAILED_CAUSE_BAD_START_CHILD_EXECUTION_ATTRIBUTES                      WorkflowTaskFailedCause = 16
	WORKFLOW_TASK_FAILED_CAUSE_FORCE_CLOSE_COMMAND                                       WorkflowTaskFailedCause = 17
	WORKFLOW_TASK_FAILED_CAUSE_FAILOVER_CLOSE_COMMAND                                    WorkflowTaskFailedCause = 18
	WORKFLOW_TASK_FAILED_CAUSE_BAD_SIGNAL_INPUT_SIZE                                     WorkflowTaskFailedCause = 19
	WORKFLOW_TASK_FAILED_CAUSE_RESET_WORKFLOW                                            WorkflowTaskFailedCause = 20
	WORKFLOW_TASK_FAILED_CAUSE_BAD_BINARY                                                WorkflowTaskFailedCause = 21
	WORKFLOW_TASK_FAILED_CAUSE_SCHEDULE_ACTIVITY_DUPLICATE_ID                            WorkflowTaskFailedCause = 22
	WORKFLOW_TASK_FAILED_CAUSE_BAD_SEARCH_ATTRIBUTES                                     WorkflowTaskFailedCause = 23
)

var WorkflowTaskFailedCause_name = map[int32]string{
	0:  "Unspecified",
	1:  "UnhandledCommand",
	2:  "BadScheduleActivityAttributes",
	3:  "BadRequestCancelActivityAttributes",
	4:  "BadStartTimerAttributes",
	5:  "BadCancelTimerAttributes",
	6:  "BadRecordMarkerAttributes",
	7:  "BadCompleteWorkflowExecutionAttributes",
	8:  "BadFailWorkflowExecutionAttributes",
	9:  "BadCancelWorkflowExecutionAttributes",
	10: "BadRequestCancelExternalWorkflowExecutionAttributes",
	11: "BadContinueAsNewAttributes",
	12: "StartTimerDuplicateId",
	13: "ResetStickyTaskQueue",
	14: "WorkflowWorkerUnhandledFailure",
	15: "BadSignalWorkflowExecutionAttributes",
	16: "BadStartChildExecutionAttributes",
	17: "ForceCloseCommand",
	18: "FailoverCloseCommand",
	19: "BadSignalInputSize",
	20: "ResetWorkflow",
	21: "BadBinary",
	22: "ScheduleActivityDuplicateId",
	23: "BadSearchAttributes",
}

var WorkflowTaskFailedCause_value = map[string]int32{
	"Unspecified":                                         0,
	"UnhandledCommand":                                    1,
	"BadScheduleActivityAttributes":                       2,
	"BadRequestCancelActivityAttributes":                  3,
	"BadStartTimerAttributes":                             4,
	"BadCancelTimerAttributes":                            5,
	"BadRecordMarkerAttributes":                           6,
	"BadCompleteWorkflowExecutionAttributes":              7,
	"BadFailWorkflowExecutionAttributes":                  8,
	"BadCancelWorkflowExecutionAttributes":                9,
	"BadRequestCancelExternalWorkflowExecutionAttributes": 10,
	"BadContinueAsNewAttributes":                          11,
	"StartTimerDuplicateId":                               12,
	"ResetStickyTaskQueue":                                13,
	"WorkflowWorkerUnhandledFailure":                      14,
	"BadSignalWorkflowExecutionAttributes":                15,
	"BadStartChildExecutionAttributes":                    16,
	"ForceCloseCommand":                                   17,
	"FailoverCloseCommand":                                18,
	"BadSignalInputSize":                                  19,
	"ResetWorkflow":                                       20,
	"BadBinary":                                           21,
	"ScheduleActivityDuplicateId":                         22,
	"BadSearchAttributes":                                 23,
}

func (WorkflowTaskFailedCause) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b293cf8d1d965f2d, []int{0}
}

type StartChildWorkflowExecutionFailedCause int32

const (
	START_CHILD_WORKFLOW_EXECUTION_FAILED_CAUSE_UNSPECIFIED             StartChildWorkflowExecutionFailedCause = 0
	START_CHILD_WORKFLOW_EXECUTION_FAILED_CAUSE_WORKFLOW_ALREADY_EXISTS StartChildWorkflowExecutionFailedCause = 1
)

var StartChildWorkflowExecutionFailedCause_name = map[int32]string{
	0: "Unspecified",
	1: "WorkflowAlreadyExists",
}

var StartChildWorkflowExecutionFailedCause_value = map[string]int32{
	"Unspecified":           0,
	"WorkflowAlreadyExists": 1,
}

func (StartChildWorkflowExecutionFailedCause) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b293cf8d1d965f2d, []int{1}
}

type CancelExternalWorkflowExecutionFailedCause int32

const (
	CANCEL_EXTERNAL_WORKFLOW_EXECUTION_FAILED_CAUSE_UNSPECIFIED                           CancelExternalWorkflowExecutionFailedCause = 0
	CANCEL_EXTERNAL_WORKFLOW_EXECUTION_FAILED_CAUSE_EXTERNAL_WORKFLOW_EXECUTION_NOT_FOUND CancelExternalWorkflowExecutionFailedCause = 1
)

var CancelExternalWorkflowExecutionFailedCause_name = map[int32]string{
	0: "Unspecified",
	1: "ExternalWorkflowExecutionNotFound",
}

var CancelExternalWorkflowExecutionFailedCause_value = map[string]int32{
	"Unspecified":                       0,
	"ExternalWorkflowExecutionNotFound": 1,
}

func (CancelExternalWorkflowExecutionFailedCause) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b293cf8d1d965f2d, []int{2}
}

type SignalExternalWorkflowExecutionFailedCause int32

const (
	SIGNAL_EXTERNAL_WORKFLOW_EXECUTION_FAILED_CAUSE_UNSPECIFIED                           SignalExternalWorkflowExecutionFailedCause = 0
	SIGNAL_EXTERNAL_WORKFLOW_EXECUTION_FAILED_CAUSE_EXTERNAL_WORKFLOW_EXECUTION_NOT_FOUND SignalExternalWorkflowExecutionFailedCause = 1
)

var SignalExternalWorkflowExecutionFailedCause_name = map[int32]string{
	0: "Unspecified",
	1: "ExternalWorkflowExecutionNotFound",
}

var SignalExternalWorkflowExecutionFailedCause_value = map[string]int32{
	"Unspecified":                       0,
	"ExternalWorkflowExecutionNotFound": 1,
}

func (SignalExternalWorkflowExecutionFailedCause) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b293cf8d1d965f2d, []int{3}
}

func init() {
	proto.RegisterEnum("temporal.api.enums.v1.WorkflowTaskFailedCause", WorkflowTaskFailedCause_name, WorkflowTaskFailedCause_value)
	proto.RegisterEnum("temporal.api.enums.v1.StartChildWorkflowExecutionFailedCause", StartChildWorkflowExecutionFailedCause_name, StartChildWorkflowExecutionFailedCause_value)
	proto.RegisterEnum("temporal.api.enums.v1.CancelExternalWorkflowExecutionFailedCause", CancelExternalWorkflowExecutionFailedCause_name, CancelExternalWorkflowExecutionFailedCause_value)
	proto.RegisterEnum("temporal.api.enums.v1.SignalExternalWorkflowExecutionFailedCause", SignalExternalWorkflowExecutionFailedCause_name, SignalExternalWorkflowExecutionFailedCause_value)
}

func init() {
	proto.RegisterFile("temporal/api/enums/v1/failed_cause.proto", fileDescriptor_b293cf8d1d965f2d)
}

var fileDescriptor_b293cf8d1d965f2d = []byte{
	// 749 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x56, 0x4d, 0x53, 0xe3, 0x36,
	0x18, 0x8e, 0x69, 0x4b, 0x5b, 0xd1, 0x0f, 0x57, 0x2d, 0x85, 0x93, 0x67, 0x7a, 0x28, 0x53, 0x32,
	0x6d, 0x42, 0xa0, 0xc0, 0x90, 0xb4, 0x43, 0x15, 0xf9, 0x0d, 0xd1, 0xc4, 0xb1, 0x83, 0x24, 0x93,
	0x84, 0x8b, 0xc6, 0x85, 0x40, 0x3d, 0x84, 0x38, 0x13, 0x02, 0xe5, 0xd8, 0x9f, 0xd0, 0x9f, 0xb0,
	0xc7, 0x3d, 0xed, 0xef, 0xe0, 0xc8, 0x91, 0xe3, 0x12, 0x2e, 0x3b, 0x7b, 0xe2, 0x27, 0xec, 0x24,
	0x0b, 0x19, 0xb3, 0x64, 0x6c, 0x67, 0x6f, 0xb2, 0xf4, 0x3c, 0x8f, 0xde, 0x2f, 0xbd, 0x7e, 0xd1,
	0x2f, 0xfd, 0xd6, 0x69, 0x37, 0xe8, 0x79, 0xed, 0xac, 0xd7, 0xf5, 0xb3, 0xad, 0xce, 0xf9, 0xe9,
	0x59, 0xf6, 0x22, 0x97, 0x3d, 0xf2, 0xfc, 0x76, 0xeb, 0x50, 0x1d, 0x78, 0xe7, 0x67, 0xad, 0x4c,
	0xb7, 0x17, 0xf4, 0x03, 0x3c, 0xff, 0x88, 0xcc, 0x78, 0x5d, 0x3f, 0x33, 0x42, 0x66, 0x2e, 0x72,
	0xe9, 0x9b, 0x39, 0xb4, 0x50, 0x0f, 0x7a, 0x27, 0x47, 0xed, 0xe0, 0x5f, 0xe9, 0x9d, 0x9d, 0x94,
	0x46, 0x4c, 0x3a, 0x24, 0xe2, 0x34, 0x5a, 0xaa, 0x3b, 0xbc, 0x52, 0xb2, 0x9c, 0xba, 0x92, 0x44,
	0x54, 0x54, 0x89, 0x30, 0x0b, 0x4c, 0x45, 0x89, 0x2b, 0x40, 0xb9, 0xb6, 0xa8, 0x01, 0x65, 0x25,
	0x06, 0xa6, 0x9e, 0xc2, 0x2b, 0xe8, 0xd7, 0x48, 0x6c, 0x99, 0xd8, 0xe6, 0xe8, 0xdb, 0xa9, 0x56,
	0x89, 0x6d, 0xea, 0x1a, 0xde, 0x46, 0x85, 0x08, 0x46, 0x91, 0x98, 0x4a, 0xd0, 0x32, 0x98, 0xae,
	0x05, 0x8a, 0x50, 0xc9, 0xf6, 0x98, 0x6c, 0x2a, 0x22, 0x25, 0x67, 0x45, 0x57, 0x82, 0xd0, 0x67,
	0x30, 0x20, 0x12, 0x23, 0xc0, 0x61, 0xd7, 0x05, 0x21, 0x15, 0x25, 0x36, 0x05, 0x6b, 0xa2, 0xcc,
	0x27, 0x78, 0x0b, 0xad, 0xc7, 0xd9, 0x21, 0x09, 0x97, 0x4a, 0xb2, 0x2a, 0xf0, 0x30, 0xf5, 0x53,
	0x9c, 0x47, 0x1b, 0x31, 0xd4, 0x87, 0x9b, 0x9f, 0x71, 0x3f, 0xc3, 0x05, 0xb4, 0x19, 0x6b, 0x3d,
	0x75, 0xb8, 0xa9, 0xaa, 0x84, 0x57, 0x9e, 0x92, 0x67, 0x31, 0x43, 0x10, 0x77, 0xb1, 0x53, 0xad,
	0x59, 0x20, 0x41, 0x8d, 0x71, 0xd0, 0x00, 0xea, 0x4a, 0xe6, 0xd8, 0x61, 0xa9, 0xcf, 0x13, 0x44,
	0x71, 0xb8, 0x11, 0x23, 0xf3, 0x05, 0xde, 0x41, 0x34, 0x59, 0x28, 0xa2, 0x85, 0xbe, 0xc4, 0x0d,
	0x24, 0xa7, 0xcb, 0x2a, 0x34, 0x24, 0x70, 0x9b, 0xc4, 0x29, 0x23, 0xfc, 0x27, 0xda, 0x8a, 0x0d,
	0x9a, 0x2d, 0x99, 0xed, 0x82, 0x22, 0x42, 0xd9, 0x50, 0x0f, 0xd3, 0xe7, 0xf0, 0x26, 0x5a, 0x8b,
	0xa0, 0x87, 0x6b, 0xc4, 0x74, 0x6b, 0x16, 0xa3, 0x44, 0x82, 0x62, 0xa6, 0xfe, 0x15, 0xde, 0x40,
	0xab, 0x11, 0x44, 0x0e, 0x02, 0xa4, 0x12, 0x92, 0xd1, 0x4a, 0xf3, 0xfd, 0xf1, 0xae, 0x0b, 0x2e,
	0xe8, 0x5f, 0xe3, 0xbf, 0xd0, 0x1f, 0x11, 0xbc, 0xf1, 0xd1, 0x70, 0x01, 0x3c, 0xf4, 0xc4, 0x86,
	0x30, 0x97, 0x83, 0xfe, 0x4d, 0x82, 0xa4, 0x08, 0xb6, 0x13, 0x1f, 0xba, 0x6f, 0x31, 0x45, 0xdb,
	0x89, 0xde, 0x08, 0x2d, 0x33, 0xcb, 0x9c, 0x2c, 0xa2, 0xe3, 0x55, 0x94, 0x89, 0x10, 0x29, 0x39,
	0x9c, 0x82, 0xa2, 0x96, 0x23, 0x60, 0xdc, 0x24, 0xbe, 0xc3, 0xeb, 0x28, 0x17, 0xc5, 0x21, 0xcc,
	0x72, 0xf6, 0x80, 0x7f, 0x40, 0xc3, 0xf8, 0x77, 0xb4, 0x92, 0xcc, 0x71, 0x66, 0xd7, 0x5c, 0xa9,
	0x04, 0xdb, 0x07, 0xfd, 0x7b, 0xfc, 0x1b, 0x5a, 0x8e, 0x4d, 0xd4, 0x23, 0x40, 0xff, 0x01, 0x2f,
	0xa3, 0x9f, 0x63, 0x2e, 0x29, 0x32, 0x9b, 0xf0, 0xa6, 0x3e, 0x1f, 0x53, 0x7a, 0xcf, 0xfb, 0xdc,
	0x93, 0x0a, 0xfa, 0x31, 0x89, 0x3b, 0x40, 0x38, 0x2d, 0x87, 0xe3, 0xbd, 0x90, 0x7e, 0xa5, 0xa1,
	0x25, 0xd1, 0xf7, 0x7a, 0x7d, 0xfa, 0x8f, 0xdf, 0x3e, 0x7c, 0x6c, 0xf2, 0x70, 0xd9, 0x3a, 0x38,
	0xef, 0xfb, 0x41, 0x27, 0xdc, 0xe9, 0x0b, 0x68, 0x33, 0x9c, 0xc0, 0x09, 0xe5, 0x10, 0xd1, 0xfa,
	0x77, 0x10, 0x9d, 0x86, 0x3c, 0x3e, 0x27, 0x16, 0x07, 0x62, 0x36, 0x15, 0x34, 0x98, 0x90, 0x42,
	0xd7, 0xd2, 0x57, 0x1a, 0x4a, 0x53, 0xaf, 0x73, 0xd0, 0x6a, 0xc3, 0x65, 0xbf, 0xd5, 0xeb, 0x78,
	0xed, 0x48, 0xa3, 0xb7, 0x51, 0x21, 0x41, 0x0b, 0x88, 0x30, 0xbc, 0x89, 0xdc, 0x69, 0x05, 0xa2,
	0x80, 0xb6, 0x23, 0x55, 0xc9, 0x71, 0x87, 0x3f, 0xb7, 0x91, 0x2b, 0xc2, 0x3f, 0xee, 0x78, 0x89,
	0x5d, 0x79, 0x28, 0xc8, 0x8f, 0x77, 0x65, 0x5a, 0x81, 0x84, 0xae, 0x14, 0x5f, 0x68, 0xd7, 0xb7,
	0x46, 0xea, 0xe6, 0xd6, 0x48, 0xdd, 0xdf, 0x1a, 0xda, 0x7f, 0x03, 0x43, 0x7b, 0x39, 0x30, 0xb4,
	0xab, 0x81, 0xa1, 0x5d, 0x0f, 0x0c, 0xed, 0xf5, 0xc0, 0xd0, 0xde, 0x0c, 0x8c, 0xd4, 0xfd, 0xc0,
	0xd0, 0xfe, 0xbf, 0x33, 0x52, 0xd7, 0x77, 0x46, 0xea, 0xe6, 0xce, 0x48, 0xa1, 0x45, 0x3f, 0xc8,
	0x4c, 0x1c, 0x39, 0x8a, 0x7a, 0xc8, 0xf3, 0xda, 0x70, 0x36, 0xa9, 0x69, 0xfb, 0x3f, 0x1d, 0x87,
	0xd0, 0x7e, 0xf0, 0x64, 0x9a, 0x29, 0x8c, 0x16, 0x6f, 0x67, 0x16, 0xe5, 0x03, 0x20, 0x9f, 0x27,
	0x5d, 0x3f, 0x9f, 0x87, 0xe1, 0x76, 0x3e, 0xbf, 0x97, 0xfb, 0x7b, 0x76, 0x34, 0xe2, 0xac, 0xbd,
	0x0b, 0x00, 0x00, 0xff, 0xff, 0x6c, 0xa2, 0xf7, 0x34, 0x0e, 0x09, 0x00, 0x00,
}

func (x WorkflowTaskFailedCause) String() string {
	s, ok := WorkflowTaskFailedCause_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x StartChildWorkflowExecutionFailedCause) String() string {
	s, ok := StartChildWorkflowExecutionFailedCause_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x CancelExternalWorkflowExecutionFailedCause) String() string {
	s, ok := CancelExternalWorkflowExecutionFailedCause_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x SignalExternalWorkflowExecutionFailedCause) String() string {
	s, ok := SignalExternalWorkflowExecutionFailedCause_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
