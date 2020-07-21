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
// source: temporal/api/enums/v1/workflow.proto

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

type WorkflowIdReusePolicy int32

const (
	WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED WorkflowIdReusePolicy = 0
	// Allow start a workflow execution using the same workflow Id, when workflow not running.
	WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE WorkflowIdReusePolicy = 1
	// Allow start a workflow execution using the same workflow Id, when workflow not running, and the last execution close state is in
	// [terminated, cancelled, timed out, failed].
	WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE_FAILED_ONLY WorkflowIdReusePolicy = 2
	// Do not allow start a workflow execution using the same workflow Id at all.
	WORKFLOW_ID_REUSE_POLICY_REJECT_DUPLICATE WorkflowIdReusePolicy = 3
)

var WorkflowIdReusePolicy_name = map[int32]string{
	0: "Unspecified",
	1: "AllowDuplicate",
	2: "AllowDuplicateFailedOnly",
	3: "RejectDuplicate",
}

var WorkflowIdReusePolicy_value = map[string]int32{
	"Unspecified":              0,
	"AllowDuplicate":           1,
	"AllowDuplicateFailedOnly": 2,
	"RejectDuplicate":          3,
}

func (WorkflowIdReusePolicy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_939fa9511cc117f0, []int{0}
}

type ParentClosePolicy int32

const (
	PARENT_CLOSE_POLICY_UNSPECIFIED ParentClosePolicy = 0
	// Terminate means terminating the child workflow.
	PARENT_CLOSE_POLICY_TERMINATE ParentClosePolicy = 1
	// Abandon means not doing anything on the child workflow.
	PARENT_CLOSE_POLICY_ABANDON ParentClosePolicy = 2
	// Cancel means requesting cancellation on the child workflow.
	PARENT_CLOSE_POLICY_REQUEST_CANCEL ParentClosePolicy = 3
)

var ParentClosePolicy_name = map[int32]string{
	0: "Unspecified",
	1: "Terminate",
	2: "Abandon",
	3: "RequestCancel",
}

var ParentClosePolicy_value = map[string]int32{
	"Unspecified":   0,
	"Terminate":     1,
	"Abandon":       2,
	"RequestCancel": 3,
}

func (ParentClosePolicy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_939fa9511cc117f0, []int{1}
}

type ContinueAsNewInitiator int32

const (
	CONTINUE_AS_NEW_INITIATOR_UNSPECIFIED   ContinueAsNewInitiator = 0
	CONTINUE_AS_NEW_INITIATOR_WORKFLOW      ContinueAsNewInitiator = 1
	CONTINUE_AS_NEW_INITIATOR_RETRY         ContinueAsNewInitiator = 2
	CONTINUE_AS_NEW_INITIATOR_CRON_SCHEDULE ContinueAsNewInitiator = 3
)

var ContinueAsNewInitiator_name = map[int32]string{
	0: "Unspecified",
	1: "Workflow",
	2: "Retry",
	3: "CronSchedule",
}

var ContinueAsNewInitiator_value = map[string]int32{
	"Unspecified":  0,
	"Workflow":     1,
	"Retry":        2,
	"CronSchedule": 3,
}

func (ContinueAsNewInitiator) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_939fa9511cc117f0, []int{2}
}

// (-- api-linter: core::0216::synonyms=disabled
//     aip.dev/not-precedent: There is WorkflowExecutionState already in another package. --)
type WorkflowExecutionStatus int32

const (
	WORKFLOW_EXECUTION_STATUS_UNSPECIFIED WorkflowExecutionStatus = 0
	// Value 1 is hardcoded in SQL persistence.
	WORKFLOW_EXECUTION_STATUS_RUNNING          WorkflowExecutionStatus = 1
	WORKFLOW_EXECUTION_STATUS_COMPLETED        WorkflowExecutionStatus = 2
	WORKFLOW_EXECUTION_STATUS_FAILED           WorkflowExecutionStatus = 3
	WORKFLOW_EXECUTION_STATUS_CANCELED         WorkflowExecutionStatus = 4
	WORKFLOW_EXECUTION_STATUS_TERMINATED       WorkflowExecutionStatus = 5
	WORKFLOW_EXECUTION_STATUS_CONTINUED_AS_NEW WorkflowExecutionStatus = 6
	WORKFLOW_EXECUTION_STATUS_TIMED_OUT        WorkflowExecutionStatus = 7
)

var WorkflowExecutionStatus_name = map[int32]string{
	0: "Unspecified",
	1: "Running",
	2: "Completed",
	3: "Failed",
	4: "Canceled",
	5: "Terminated",
	6: "ContinuedAsNew",
	7: "TimedOut",
}

var WorkflowExecutionStatus_value = map[string]int32{
	"Unspecified":    0,
	"Running":        1,
	"Completed":      2,
	"Failed":         3,
	"Canceled":       4,
	"Terminated":     5,
	"ContinuedAsNew": 6,
	"TimedOut":       7,
}

func (WorkflowExecutionStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_939fa9511cc117f0, []int{3}
}

type PendingActivityState int32

const (
	PENDING_ACTIVITY_STATE_UNSPECIFIED      PendingActivityState = 0
	PENDING_ACTIVITY_STATE_SCHEDULED        PendingActivityState = 1
	PENDING_ACTIVITY_STATE_STARTED          PendingActivityState = 2
	PENDING_ACTIVITY_STATE_CANCEL_REQUESTED PendingActivityState = 3
)

var PendingActivityState_name = map[int32]string{
	0: "Unspecified",
	1: "Scheduled",
	2: "Started",
	3: "CancelRequested",
}

var PendingActivityState_value = map[string]int32{
	"Unspecified":     0,
	"Scheduled":       1,
	"Started":         2,
	"CancelRequested": 3,
}

func (PendingActivityState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_939fa9511cc117f0, []int{4}
}

type HistoryEventFilterType int32

const (
	HISTORY_EVENT_FILTER_TYPE_UNSPECIFIED HistoryEventFilterType = 0
	HISTORY_EVENT_FILTER_TYPE_ALL_EVENT   HistoryEventFilterType = 1
	HISTORY_EVENT_FILTER_TYPE_CLOSE_EVENT HistoryEventFilterType = 2
)

var HistoryEventFilterType_name = map[int32]string{
	0: "Unspecified",
	1: "AllEvent",
	2: "CloseEvent",
}

var HistoryEventFilterType_value = map[string]int32{
	"Unspecified": 0,
	"AllEvent":    1,
	"CloseEvent":  2,
}

func (HistoryEventFilterType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_939fa9511cc117f0, []int{5}
}

type RetryState int32

const (
	RETRY_STATE_UNSPECIFIED              RetryState = 0
	RETRY_STATE_IN_PROGRESS              RetryState = 1
	RETRY_STATE_NON_RETRYABLE_FAILURE    RetryState = 2
	RETRY_STATE_TIMEOUT                  RetryState = 3
	RETRY_STATE_MAXIMUM_ATTEMPTS_REACHED RetryState = 4
	RETRY_STATE_RETRY_POLICY_NOT_SET     RetryState = 5
	RETRY_STATE_INTERNAL_SERVER_ERROR    RetryState = 6
	RETRY_STATE_CANCEL_REQUESTED         RetryState = 7
)

var RetryState_name = map[int32]string{
	0: "Unspecified",
	1: "InProgress",
	2: "NonRetryableFailure",
	3: "Timeout",
	4: "MaximumAttemptsReached",
	5: "RetryPolicyNotSet",
	6: "InternalServerError",
	7: "CancelRequested",
}

var RetryState_value = map[string]int32{
	"Unspecified":            0,
	"InProgress":             1,
	"NonRetryableFailure":    2,
	"Timeout":                3,
	"MaximumAttemptsReached": 4,
	"RetryPolicyNotSet":      5,
	"InternalServerError":    6,
	"CancelRequested":        7,
}

func (RetryState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_939fa9511cc117f0, []int{6}
}

type TimeoutType int32

const (
	TIMEOUT_TYPE_UNSPECIFIED       TimeoutType = 0
	TIMEOUT_TYPE_START_TO_CLOSE    TimeoutType = 1
	TIMEOUT_TYPE_SCHEDULE_TO_START TimeoutType = 2
	TIMEOUT_TYPE_SCHEDULE_TO_CLOSE TimeoutType = 3
	TIMEOUT_TYPE_HEARTBEAT         TimeoutType = 4
)

var TimeoutType_name = map[int32]string{
	0: "Unspecified",
	1: "StartToClose",
	2: "ScheduleToStart",
	3: "ScheduleToClose",
	4: "Heartbeat",
}

var TimeoutType_value = map[string]int32{
	"Unspecified":     0,
	"StartToClose":    1,
	"ScheduleToStart": 2,
	"ScheduleToClose": 3,
	"Heartbeat":       4,
}

func (TimeoutType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_939fa9511cc117f0, []int{7}
}

func init() {
	proto.RegisterEnum("temporal.api.enums.v1.WorkflowIdReusePolicy", WorkflowIdReusePolicy_name, WorkflowIdReusePolicy_value)
	proto.RegisterEnum("temporal.api.enums.v1.ParentClosePolicy", ParentClosePolicy_name, ParentClosePolicy_value)
	proto.RegisterEnum("temporal.api.enums.v1.ContinueAsNewInitiator", ContinueAsNewInitiator_name, ContinueAsNewInitiator_value)
	proto.RegisterEnum("temporal.api.enums.v1.WorkflowExecutionStatus", WorkflowExecutionStatus_name, WorkflowExecutionStatus_value)
	proto.RegisterEnum("temporal.api.enums.v1.PendingActivityState", PendingActivityState_name, PendingActivityState_value)
	proto.RegisterEnum("temporal.api.enums.v1.HistoryEventFilterType", HistoryEventFilterType_name, HistoryEventFilterType_value)
	proto.RegisterEnum("temporal.api.enums.v1.RetryState", RetryState_name, RetryState_value)
	proto.RegisterEnum("temporal.api.enums.v1.TimeoutType", TimeoutType_name, TimeoutType_value)
}

func init() {
	proto.RegisterFile("temporal/api/enums/v1/workflow.proto", fileDescriptor_939fa9511cc117f0)
}

var fileDescriptor_939fa9511cc117f0 = []byte{
	// 900 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x95, 0x4f, 0x6f, 0xe3, 0x44,
	0x18, 0xc6, 0x33, 0x49, 0xb7, 0x2b, 0x0d, 0x42, 0x1a, 0x0c, 0xdb, 0x56, 0x74, 0x71, 0xff, 0xee,
	0x76, 0x37, 0x0b, 0xa9, 0x2a, 0x38, 0xa0, 0x70, 0x9a, 0xd8, 0x6f, 0xda, 0x01, 0x67, 0x6c, 0xc6,
	0xe3, 0x76, 0xc3, 0x65, 0x14, 0xba, 0x66, 0x65, 0x91, 0xc6, 0x51, 0xe2, 0xb4, 0xf4, 0xc6, 0x47,
	0xe0, 0xc6, 0x09, 0x71, 0xe1, 0xc0, 0x91, 0x0b, 0x47, 0xee, 0x1c, 0x7b, 0xec, 0x91, 0xa6, 0x42,
	0x42, 0x9c, 0xf6, 0x23, 0xa0, 0xb1, 0x93, 0x2a, 0x49, 0xe3, 0x88, 0x9b, 0x35, 0xf3, 0x9b, 0xf1,
	0xbc, 0xcf, 0xf3, 0xbc, 0x33, 0x78, 0x37, 0x09, 0xcf, 0xba, 0x71, 0xaf, 0xd5, 0xde, 0x6f, 0x75,
	0xa3, 0xfd, 0xb0, 0x33, 0x38, 0xeb, 0xef, 0x9f, 0x1f, 0xec, 0x5f, 0xc4, 0xbd, 0x6f, 0xbf, 0x69,
	0xc7, 0x17, 0x95, 0x6e, 0x2f, 0x4e, 0x62, 0xe3, 0xd1, 0x98, 0xaa, 0xb4, 0xba, 0x51, 0x25, 0xa5,
	0x2a, 0xe7, 0x07, 0xe5, 0x6b, 0x84, 0x1f, 0x9d, 0x8c, 0x48, 0xf6, 0x4a, 0x84, 0x83, 0x7e, 0xe8,
	0xc5, 0xed, 0xe8, 0xf4, 0xd2, 0x78, 0x86, 0x77, 0x4f, 0x5c, 0xf1, 0x45, 0xdd, 0x71, 0x4f, 0x14,
	0xb3, 0x95, 0x80, 0xc0, 0x07, 0xe5, 0xb9, 0x0e, 0xb3, 0x9a, 0x2a, 0xe0, 0xbe, 0x07, 0x16, 0xab,
	0x33, 0xb0, 0x49, 0xc1, 0xf8, 0x10, 0x3f, 0xcb, 0x25, 0xa9, 0xa3, 0x47, 0xed, 0xc0, 0x73, 0x98,
	0x45, 0x25, 0x10, 0x64, 0x7c, 0x8a, 0x3f, 0xf9, 0xbf, 0xb4, 0xaa, 0x53, 0xe6, 0x80, 0xad, 0x5c,
	0xee, 0x34, 0x49, 0xd1, 0xf8, 0x08, 0x3f, 0xcf, 0x5d, 0x29, 0xe0, 0x73, 0xb0, 0xe4, 0xc4, 0x8f,
	0x4a, 0xe5, 0x5f, 0x10, 0x7e, 0xc7, 0x6b, 0xf5, 0xc2, 0x4e, 0x62, 0xb5, 0xe3, 0xbb, 0xb2, 0x76,
	0xf0, 0x86, 0x47, 0x05, 0x70, 0xa9, 0x2c, 0xc7, 0xcd, 0xab, 0x68, 0x0b, 0x7f, 0x30, 0x0f, 0x92,
	0x20, 0x1a, 0x8c, 0x67, 0x65, 0x6c, 0xe0, 0xf5, 0x79, 0x08, 0xad, 0x51, 0x6e, 0xbb, 0x9c, 0x14,
	0x8d, 0xa7, 0x78, 0x7b, 0x1e, 0x20, 0xe0, 0xcb, 0x00, 0x7c, 0xa9, 0x2c, 0xca, 0x2d, 0x70, 0x48,
	0xa9, 0xfc, 0x07, 0xc2, 0x2b, 0x56, 0xdc, 0x49, 0xa2, 0xce, 0x20, 0xa4, 0x7d, 0x1e, 0x5e, 0xb0,
	0x4e, 0x94, 0x44, 0xad, 0x24, 0xee, 0x19, 0xcf, 0xf1, 0x13, 0xcb, 0xe5, 0x92, 0xf1, 0x00, 0x14,
	0xf5, 0x15, 0x87, 0x13, 0xc5, 0x38, 0x93, 0x8c, 0x4a, 0x57, 0xcc, 0x9c, 0xf8, 0x29, 0xde, 0xce,
	0x47, 0xc7, 0xaa, 0x11, 0xa4, 0xcb, 0xcf, 0xe7, 0x04, 0x48, 0xa1, 0x85, 0x7e, 0x81, 0xf7, 0xf2,
	0x21, 0x4b, 0xb8, 0x5c, 0xf9, 0xd6, 0x11, 0xd8, 0x81, 0xa3, 0x65, 0xfe, 0xbb, 0x88, 0x57, 0xc7,
	0x09, 0x82, 0xef, 0xc2, 0xd3, 0x41, 0x12, 0xc5, 0x1d, 0x3f, 0x69, 0x25, 0x83, 0xbe, 0x2e, 0xe0,
	0xce, 0x31, 0x78, 0x09, 0x56, 0x20, 0x99, 0x5e, 0x2c, 0xa9, 0x0c, 0xfc, 0x99, 0x02, 0x9e, 0xe0,
	0xad, 0x7c, 0x54, 0x04, 0x9c, 0x33, 0x7e, 0x48, 0x90, 0xb1, 0x87, 0x77, 0xf2, 0x31, 0xcb, 0x6d,
	0x78, 0x0e, 0x48, 0xb0, 0x49, 0xd1, 0xd8, 0xc5, 0x9b, 0xf9, 0x60, 0x96, 0x2b, 0x52, 0xd2, 0xb2,
	0x2d, 0xd8, 0x2e, 0xb5, 0x08, 0x6c, 0xb2, 0x34, 0xd5, 0x0c, 0xf7, 0xb8, 0xbb, 0x58, 0xd8, 0xe4,
	0x81, 0x51, 0xc1, 0xe5, 0x45, 0x07, 0xcc, 0x54, 0xb5, 0x47, 0xb2, 0x92, 0xe5, 0xc5, 0x05, 0x49,
	0xd6, 0xd0, 0xf1, 0x0f, 0x24, 0x79, 0x58, 0xfe, 0x1d, 0xe1, 0xf7, 0xbc, 0xb0, 0xf3, 0x2a, 0xea,
	0xbc, 0xa6, 0xa7, 0x49, 0x74, 0x1e, 0x25, 0x97, 0x5a, 0xe5, 0x30, 0x0d, 0x1a, 0x70, 0x9b, 0xf1,
	0x43, 0x45, 0x2d, 0xc9, 0x8e, 0x99, 0x6c, 0xa6, 0xeb, 0x61, 0x46, 0xe1, 0x5d, 0xbc, 0x99, 0xc3,
	0x8d, 0xdd, 0xb4, 0x09, 0x32, 0xb6, 0xb1, 0x99, 0x47, 0x49, 0x2a, 0x32, 0x6d, 0x5f, 0xe0, 0xbd,
	0x1c, 0x26, 0x93, 0x6c, 0x1c, 0x72, 0x2d, 0x71, 0xf9, 0x47, 0x84, 0x57, 0x8e, 0xa2, 0x7e, 0x12,
	0xf7, 0x2e, 0xe1, 0x3c, 0xec, 0x24, 0xf5, 0xa8, 0x9d, 0x84, 0x3d, 0x79, 0xd9, 0x0d, 0x75, 0x3c,
	0x8e, 0x98, 0x2f, 0x5d, 0xd1, 0x54, 0x70, 0xac, 0x3b, 0xa5, 0xce, 0x1c, 0x09, 0x42, 0xc9, 0xa6,
	0x37, 0x7b, 0xf8, 0x3d, 0xbc, 0x93, 0x8f, 0x52, 0xc7, 0xc9, 0x46, 0x09, 0x5a, 0xbc, 0x67, 0xd6,
	0x89, 0x19, 0x5a, 0x2c, 0xff, 0x5c, 0xc4, 0x58, 0x84, 0x49, 0x6f, 0xa4, 0xe3, 0x3a, 0x5e, 0x4d,
	0x1b, 0x60, 0xae, 0x78, 0x33, 0x93, 0x8c, 0x2b, 0x4f, 0xb8, 0x87, 0x02, 0x7c, 0x9f, 0x20, 0x9d,
	0xdd, 0xc9, 0x49, 0xee, 0xf2, 0xac, 0x95, 0x68, 0xcd, 0xc9, 0xee, 0xb0, 0x40, 0x00, 0x29, 0x1a,
	0xab, 0xf8, 0xdd, 0x49, 0x4c, 0x9b, 0xab, 0xad, 0x2d, 0xe9, 0x74, 0x4d, 0x4e, 0x34, 0xe8, 0x4b,
	0xd6, 0x08, 0x1a, 0x8a, 0x4a, 0x09, 0x0d, 0x4f, 0xfa, 0x4a, 0x00, 0xd5, 0x0e, 0x91, 0x25, 0xed,
	0xe1, 0x24, 0x99, 0x7d, 0x8f, 0x6e, 0x16, 0xee, 0x4a, 0xe5, 0x83, 0x24, 0x0f, 0x66, 0xcf, 0xc3,
	0xb8, 0x04, 0xc1, 0xa9, 0xa3, 0x7c, 0x10, 0xc7, 0x20, 0x14, 0x08, 0xe1, 0x0a, 0xb2, 0x6c, 0x6c,
	0xe2, 0xc7, 0x93, 0xd8, 0x3d, 0xef, 0x1e, 0x96, 0x7f, 0x43, 0xf8, 0x2d, 0x19, 0x9d, 0x85, 0xf1,
	0x20, 0x49, 0x0d, 0x7b, 0x8c, 0xd7, 0x46, 0xa7, 0x9e, 0xe7, 0xd1, 0x06, 0x5e, 0x9f, 0x9a, 0x4d,
	0x03, 0xa3, 0xa4, 0x9b, 0xc9, 0x9e, 0x65, 0x6b, 0x1a, 0x18, 0xe5, 0x4e, 0x33, 0x29, 0x4c, 0x8a,
	0x0b, 0x99, 0x6c, 0x9f, 0x92, 0xf1, 0x3e, 0x5e, 0x99, 0x62, 0x8e, 0x80, 0x0a, 0x59, 0x03, 0x2a,
	0xc9, 0x52, 0xed, 0x27, 0x74, 0x75, 0x63, 0x16, 0xae, 0x6f, 0xcc, 0xc2, 0x9b, 0x1b, 0x13, 0x7d,
	0x3f, 0x34, 0xd1, 0xaf, 0x43, 0x13, 0xfd, 0x39, 0x34, 0xd1, 0xd5, 0xd0, 0x44, 0x7f, 0x0d, 0x4d,
	0xf4, 0xcf, 0xd0, 0x2c, 0xbc, 0x19, 0x9a, 0xe8, 0x87, 0x5b, 0xb3, 0x70, 0x75, 0x6b, 0x16, 0xae,
	0x6f, 0xcd, 0x02, 0x5e, 0x8b, 0xe2, 0xca, 0xdc, 0x17, 0xb2, 0xf6, 0xf6, 0xf8, 0x72, 0xf3, 0xf4,
	0x3b, 0xea, 0xa1, 0xaf, 0xb6, 0x5e, 0x4f, 0xa0, 0x51, 0x3c, 0xf5, 0xea, 0x7e, 0x96, 0x7e, 0xfc,
	0x5b, 0x5c, 0x93, 0x23, 0xa0, 0x5a, 0xa5, 0xdd, 0xa8, 0x5a, 0x05, 0x3d, 0x5c, 0xad, 0x1e, 0x1f,
	0x7c, 0xbd, 0x9c, 0x3e, 0xc7, 0x1f, 0xff, 0x17, 0x00, 0x00, 0xff, 0xff, 0x78, 0x18, 0x07, 0x6d,
	0xb6, 0x07, 0x00, 0x00,
}

func (x WorkflowIdReusePolicy) String() string {
	s, ok := WorkflowIdReusePolicy_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x ParentClosePolicy) String() string {
	s, ok := ParentClosePolicy_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x ContinueAsNewInitiator) String() string {
	s, ok := ContinueAsNewInitiator_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
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
func (x HistoryEventFilterType) String() string {
	s, ok := HistoryEventFilterType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x RetryState) String() string {
	s, ok := RetryState_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x TimeoutType) String() string {
	s, ok := TimeoutType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
