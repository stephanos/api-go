// The MIT License
//
// Copyright (c) 2022 Temporal Technologies Inc.  All rights reserved.
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

package history

import (
	"google.golang.org/protobuf/proto"
)

// Marshal an object of type WorkflowExecutionStartedEventAttributes to the protobuf v3 wire format
func (val *WorkflowExecutionStartedEventAttributes) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type WorkflowExecutionStartedEventAttributes from the protobuf v3 wire format
func (val *WorkflowExecutionStartedEventAttributes) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *WorkflowExecutionStartedEventAttributes) Size() int {
	return proto.Size(val)
}

// Equal returns whether two WorkflowExecutionStartedEventAttributes values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *WorkflowExecutionStartedEventAttributes) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *WorkflowExecutionStartedEventAttributes
	switch t := that.(type) {
	case *WorkflowExecutionStartedEventAttributes:
		that1 = t
	case WorkflowExecutionStartedEventAttributes:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type WorkflowExecutionCompletedEventAttributes to the protobuf v3 wire format
func (val *WorkflowExecutionCompletedEventAttributes) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type WorkflowExecutionCompletedEventAttributes from the protobuf v3 wire format
func (val *WorkflowExecutionCompletedEventAttributes) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *WorkflowExecutionCompletedEventAttributes) Size() int {
	return proto.Size(val)
}

// Equal returns whether two WorkflowExecutionCompletedEventAttributes values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *WorkflowExecutionCompletedEventAttributes) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *WorkflowExecutionCompletedEventAttributes
	switch t := that.(type) {
	case *WorkflowExecutionCompletedEventAttributes:
		that1 = t
	case WorkflowExecutionCompletedEventAttributes:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type WorkflowExecutionFailedEventAttributes to the protobuf v3 wire format
func (val *WorkflowExecutionFailedEventAttributes) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type WorkflowExecutionFailedEventAttributes from the protobuf v3 wire format
func (val *WorkflowExecutionFailedEventAttributes) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *WorkflowExecutionFailedEventAttributes) Size() int {
	return proto.Size(val)
}

// Equal returns whether two WorkflowExecutionFailedEventAttributes values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *WorkflowExecutionFailedEventAttributes) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *WorkflowExecutionFailedEventAttributes
	switch t := that.(type) {
	case *WorkflowExecutionFailedEventAttributes:
		that1 = t
	case WorkflowExecutionFailedEventAttributes:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type WorkflowExecutionTimedOutEventAttributes to the protobuf v3 wire format
func (val *WorkflowExecutionTimedOutEventAttributes) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type WorkflowExecutionTimedOutEventAttributes from the protobuf v3 wire format
func (val *WorkflowExecutionTimedOutEventAttributes) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *WorkflowExecutionTimedOutEventAttributes) Size() int {
	return proto.Size(val)
}

// Equal returns whether two WorkflowExecutionTimedOutEventAttributes values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *WorkflowExecutionTimedOutEventAttributes) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *WorkflowExecutionTimedOutEventAttributes
	switch t := that.(type) {
	case *WorkflowExecutionTimedOutEventAttributes:
		that1 = t
	case WorkflowExecutionTimedOutEventAttributes:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type WorkflowExecutionContinuedAsNewEventAttributes to the protobuf v3 wire format
func (val *WorkflowExecutionContinuedAsNewEventAttributes) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type WorkflowExecutionContinuedAsNewEventAttributes from the protobuf v3 wire format
func (val *WorkflowExecutionContinuedAsNewEventAttributes) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *WorkflowExecutionContinuedAsNewEventAttributes) Size() int {
	return proto.Size(val)
}

// Equal returns whether two WorkflowExecutionContinuedAsNewEventAttributes values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *WorkflowExecutionContinuedAsNewEventAttributes) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *WorkflowExecutionContinuedAsNewEventAttributes
	switch t := that.(type) {
	case *WorkflowExecutionContinuedAsNewEventAttributes:
		that1 = t
	case WorkflowExecutionContinuedAsNewEventAttributes:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type WorkflowTaskScheduledEventAttributes to the protobuf v3 wire format
func (val *WorkflowTaskScheduledEventAttributes) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type WorkflowTaskScheduledEventAttributes from the protobuf v3 wire format
func (val *WorkflowTaskScheduledEventAttributes) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *WorkflowTaskScheduledEventAttributes) Size() int {
	return proto.Size(val)
}

// Equal returns whether two WorkflowTaskScheduledEventAttributes values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *WorkflowTaskScheduledEventAttributes) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *WorkflowTaskScheduledEventAttributes
	switch t := that.(type) {
	case *WorkflowTaskScheduledEventAttributes:
		that1 = t
	case WorkflowTaskScheduledEventAttributes:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type WorkflowTaskStartedEventAttributes to the protobuf v3 wire format
func (val *WorkflowTaskStartedEventAttributes) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type WorkflowTaskStartedEventAttributes from the protobuf v3 wire format
func (val *WorkflowTaskStartedEventAttributes) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *WorkflowTaskStartedEventAttributes) Size() int {
	return proto.Size(val)
}

// Equal returns whether two WorkflowTaskStartedEventAttributes values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *WorkflowTaskStartedEventAttributes) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *WorkflowTaskStartedEventAttributes
	switch t := that.(type) {
	case *WorkflowTaskStartedEventAttributes:
		that1 = t
	case WorkflowTaskStartedEventAttributes:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type WorkflowTaskCompletedEventAttributes to the protobuf v3 wire format
func (val *WorkflowTaskCompletedEventAttributes) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type WorkflowTaskCompletedEventAttributes from the protobuf v3 wire format
func (val *WorkflowTaskCompletedEventAttributes) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *WorkflowTaskCompletedEventAttributes) Size() int {
	return proto.Size(val)
}

// Equal returns whether two WorkflowTaskCompletedEventAttributes values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *WorkflowTaskCompletedEventAttributes) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *WorkflowTaskCompletedEventAttributes
	switch t := that.(type) {
	case *WorkflowTaskCompletedEventAttributes:
		that1 = t
	case WorkflowTaskCompletedEventAttributes:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type WorkflowTaskTimedOutEventAttributes to the protobuf v3 wire format
func (val *WorkflowTaskTimedOutEventAttributes) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type WorkflowTaskTimedOutEventAttributes from the protobuf v3 wire format
func (val *WorkflowTaskTimedOutEventAttributes) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *WorkflowTaskTimedOutEventAttributes) Size() int {
	return proto.Size(val)
}

// Equal returns whether two WorkflowTaskTimedOutEventAttributes values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *WorkflowTaskTimedOutEventAttributes) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *WorkflowTaskTimedOutEventAttributes
	switch t := that.(type) {
	case *WorkflowTaskTimedOutEventAttributes:
		that1 = t
	case WorkflowTaskTimedOutEventAttributes:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type WorkflowTaskFailedEventAttributes to the protobuf v3 wire format
func (val *WorkflowTaskFailedEventAttributes) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type WorkflowTaskFailedEventAttributes from the protobuf v3 wire format
func (val *WorkflowTaskFailedEventAttributes) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *WorkflowTaskFailedEventAttributes) Size() int {
	return proto.Size(val)
}

// Equal returns whether two WorkflowTaskFailedEventAttributes values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *WorkflowTaskFailedEventAttributes) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *WorkflowTaskFailedEventAttributes
	switch t := that.(type) {
	case *WorkflowTaskFailedEventAttributes:
		that1 = t
	case WorkflowTaskFailedEventAttributes:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type ActivityTaskScheduledEventAttributes to the protobuf v3 wire format
func (val *ActivityTaskScheduledEventAttributes) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type ActivityTaskScheduledEventAttributes from the protobuf v3 wire format
func (val *ActivityTaskScheduledEventAttributes) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *ActivityTaskScheduledEventAttributes) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ActivityTaskScheduledEventAttributes values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ActivityTaskScheduledEventAttributes) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ActivityTaskScheduledEventAttributes
	switch t := that.(type) {
	case *ActivityTaskScheduledEventAttributes:
		that1 = t
	case ActivityTaskScheduledEventAttributes:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type ActivityTaskStartedEventAttributes to the protobuf v3 wire format
func (val *ActivityTaskStartedEventAttributes) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type ActivityTaskStartedEventAttributes from the protobuf v3 wire format
func (val *ActivityTaskStartedEventAttributes) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *ActivityTaskStartedEventAttributes) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ActivityTaskStartedEventAttributes values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ActivityTaskStartedEventAttributes) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ActivityTaskStartedEventAttributes
	switch t := that.(type) {
	case *ActivityTaskStartedEventAttributes:
		that1 = t
	case ActivityTaskStartedEventAttributes:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type ActivityTaskCompletedEventAttributes to the protobuf v3 wire format
func (val *ActivityTaskCompletedEventAttributes) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type ActivityTaskCompletedEventAttributes from the protobuf v3 wire format
func (val *ActivityTaskCompletedEventAttributes) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *ActivityTaskCompletedEventAttributes) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ActivityTaskCompletedEventAttributes values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ActivityTaskCompletedEventAttributes) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ActivityTaskCompletedEventAttributes
	switch t := that.(type) {
	case *ActivityTaskCompletedEventAttributes:
		that1 = t
	case ActivityTaskCompletedEventAttributes:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type ActivityTaskFailedEventAttributes to the protobuf v3 wire format
func (val *ActivityTaskFailedEventAttributes) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type ActivityTaskFailedEventAttributes from the protobuf v3 wire format
func (val *ActivityTaskFailedEventAttributes) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *ActivityTaskFailedEventAttributes) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ActivityTaskFailedEventAttributes values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ActivityTaskFailedEventAttributes) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ActivityTaskFailedEventAttributes
	switch t := that.(type) {
	case *ActivityTaskFailedEventAttributes:
		that1 = t
	case ActivityTaskFailedEventAttributes:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type ActivityTaskTimedOutEventAttributes to the protobuf v3 wire format
func (val *ActivityTaskTimedOutEventAttributes) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type ActivityTaskTimedOutEventAttributes from the protobuf v3 wire format
func (val *ActivityTaskTimedOutEventAttributes) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *ActivityTaskTimedOutEventAttributes) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ActivityTaskTimedOutEventAttributes values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ActivityTaskTimedOutEventAttributes) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ActivityTaskTimedOutEventAttributes
	switch t := that.(type) {
	case *ActivityTaskTimedOutEventAttributes:
		that1 = t
	case ActivityTaskTimedOutEventAttributes:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type ActivityTaskCancelRequestedEventAttributes to the protobuf v3 wire format
func (val *ActivityTaskCancelRequestedEventAttributes) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type ActivityTaskCancelRequestedEventAttributes from the protobuf v3 wire format
func (val *ActivityTaskCancelRequestedEventAttributes) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *ActivityTaskCancelRequestedEventAttributes) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ActivityTaskCancelRequestedEventAttributes values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ActivityTaskCancelRequestedEventAttributes) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ActivityTaskCancelRequestedEventAttributes
	switch t := that.(type) {
	case *ActivityTaskCancelRequestedEventAttributes:
		that1 = t
	case ActivityTaskCancelRequestedEventAttributes:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type ActivityTaskCanceledEventAttributes to the protobuf v3 wire format
func (val *ActivityTaskCanceledEventAttributes) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type ActivityTaskCanceledEventAttributes from the protobuf v3 wire format
func (val *ActivityTaskCanceledEventAttributes) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *ActivityTaskCanceledEventAttributes) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ActivityTaskCanceledEventAttributes values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ActivityTaskCanceledEventAttributes) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ActivityTaskCanceledEventAttributes
	switch t := that.(type) {
	case *ActivityTaskCanceledEventAttributes:
		that1 = t
	case ActivityTaskCanceledEventAttributes:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type TimerStartedEventAttributes to the protobuf v3 wire format
func (val *TimerStartedEventAttributes) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type TimerStartedEventAttributes from the protobuf v3 wire format
func (val *TimerStartedEventAttributes) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *TimerStartedEventAttributes) Size() int {
	return proto.Size(val)
}

// Equal returns whether two TimerStartedEventAttributes values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *TimerStartedEventAttributes) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *TimerStartedEventAttributes
	switch t := that.(type) {
	case *TimerStartedEventAttributes:
		that1 = t
	case TimerStartedEventAttributes:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type TimerFiredEventAttributes to the protobuf v3 wire format
func (val *TimerFiredEventAttributes) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type TimerFiredEventAttributes from the protobuf v3 wire format
func (val *TimerFiredEventAttributes) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *TimerFiredEventAttributes) Size() int {
	return proto.Size(val)
}

// Equal returns whether two TimerFiredEventAttributes values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *TimerFiredEventAttributes) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *TimerFiredEventAttributes
	switch t := that.(type) {
	case *TimerFiredEventAttributes:
		that1 = t
	case TimerFiredEventAttributes:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type TimerCanceledEventAttributes to the protobuf v3 wire format
func (val *TimerCanceledEventAttributes) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type TimerCanceledEventAttributes from the protobuf v3 wire format
func (val *TimerCanceledEventAttributes) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *TimerCanceledEventAttributes) Size() int {
	return proto.Size(val)
}

// Equal returns whether two TimerCanceledEventAttributes values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *TimerCanceledEventAttributes) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *TimerCanceledEventAttributes
	switch t := that.(type) {
	case *TimerCanceledEventAttributes:
		that1 = t
	case TimerCanceledEventAttributes:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type WorkflowExecutionCancelRequestedEventAttributes to the protobuf v3 wire format
func (val *WorkflowExecutionCancelRequestedEventAttributes) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type WorkflowExecutionCancelRequestedEventAttributes from the protobuf v3 wire format
func (val *WorkflowExecutionCancelRequestedEventAttributes) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *WorkflowExecutionCancelRequestedEventAttributes) Size() int {
	return proto.Size(val)
}

// Equal returns whether two WorkflowExecutionCancelRequestedEventAttributes values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *WorkflowExecutionCancelRequestedEventAttributes) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *WorkflowExecutionCancelRequestedEventAttributes
	switch t := that.(type) {
	case *WorkflowExecutionCancelRequestedEventAttributes:
		that1 = t
	case WorkflowExecutionCancelRequestedEventAttributes:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type WorkflowExecutionCanceledEventAttributes to the protobuf v3 wire format
func (val *WorkflowExecutionCanceledEventAttributes) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type WorkflowExecutionCanceledEventAttributes from the protobuf v3 wire format
func (val *WorkflowExecutionCanceledEventAttributes) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *WorkflowExecutionCanceledEventAttributes) Size() int {
	return proto.Size(val)
}

// Equal returns whether two WorkflowExecutionCanceledEventAttributes values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *WorkflowExecutionCanceledEventAttributes) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *WorkflowExecutionCanceledEventAttributes
	switch t := that.(type) {
	case *WorkflowExecutionCanceledEventAttributes:
		that1 = t
	case WorkflowExecutionCanceledEventAttributes:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type MarkerRecordedEventAttributes to the protobuf v3 wire format
func (val *MarkerRecordedEventAttributes) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type MarkerRecordedEventAttributes from the protobuf v3 wire format
func (val *MarkerRecordedEventAttributes) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *MarkerRecordedEventAttributes) Size() int {
	return proto.Size(val)
}

// Equal returns whether two MarkerRecordedEventAttributes values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *MarkerRecordedEventAttributes) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *MarkerRecordedEventAttributes
	switch t := that.(type) {
	case *MarkerRecordedEventAttributes:
		that1 = t
	case MarkerRecordedEventAttributes:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type WorkflowExecutionSignaledEventAttributes to the protobuf v3 wire format
func (val *WorkflowExecutionSignaledEventAttributes) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type WorkflowExecutionSignaledEventAttributes from the protobuf v3 wire format
func (val *WorkflowExecutionSignaledEventAttributes) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *WorkflowExecutionSignaledEventAttributes) Size() int {
	return proto.Size(val)
}

// Equal returns whether two WorkflowExecutionSignaledEventAttributes values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *WorkflowExecutionSignaledEventAttributes) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *WorkflowExecutionSignaledEventAttributes
	switch t := that.(type) {
	case *WorkflowExecutionSignaledEventAttributes:
		that1 = t
	case WorkflowExecutionSignaledEventAttributes:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type WorkflowExecutionTerminatedEventAttributes to the protobuf v3 wire format
func (val *WorkflowExecutionTerminatedEventAttributes) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type WorkflowExecutionTerminatedEventAttributes from the protobuf v3 wire format
func (val *WorkflowExecutionTerminatedEventAttributes) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *WorkflowExecutionTerminatedEventAttributes) Size() int {
	return proto.Size(val)
}

// Equal returns whether two WorkflowExecutionTerminatedEventAttributes values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *WorkflowExecutionTerminatedEventAttributes) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *WorkflowExecutionTerminatedEventAttributes
	switch t := that.(type) {
	case *WorkflowExecutionTerminatedEventAttributes:
		that1 = t
	case WorkflowExecutionTerminatedEventAttributes:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type RequestCancelExternalWorkflowExecutionInitiatedEventAttributes to the protobuf v3 wire format
func (val *RequestCancelExternalWorkflowExecutionInitiatedEventAttributes) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type RequestCancelExternalWorkflowExecutionInitiatedEventAttributes from the protobuf v3 wire format
func (val *RequestCancelExternalWorkflowExecutionInitiatedEventAttributes) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *RequestCancelExternalWorkflowExecutionInitiatedEventAttributes) Size() int {
	return proto.Size(val)
}

// Equal returns whether two RequestCancelExternalWorkflowExecutionInitiatedEventAttributes values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *RequestCancelExternalWorkflowExecutionInitiatedEventAttributes) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *RequestCancelExternalWorkflowExecutionInitiatedEventAttributes
	switch t := that.(type) {
	case *RequestCancelExternalWorkflowExecutionInitiatedEventAttributes:
		that1 = t
	case RequestCancelExternalWorkflowExecutionInitiatedEventAttributes:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type RequestCancelExternalWorkflowExecutionFailedEventAttributes to the protobuf v3 wire format
func (val *RequestCancelExternalWorkflowExecutionFailedEventAttributes) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type RequestCancelExternalWorkflowExecutionFailedEventAttributes from the protobuf v3 wire format
func (val *RequestCancelExternalWorkflowExecutionFailedEventAttributes) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *RequestCancelExternalWorkflowExecutionFailedEventAttributes) Size() int {
	return proto.Size(val)
}

// Equal returns whether two RequestCancelExternalWorkflowExecutionFailedEventAttributes values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *RequestCancelExternalWorkflowExecutionFailedEventAttributes) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *RequestCancelExternalWorkflowExecutionFailedEventAttributes
	switch t := that.(type) {
	case *RequestCancelExternalWorkflowExecutionFailedEventAttributes:
		that1 = t
	case RequestCancelExternalWorkflowExecutionFailedEventAttributes:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type ExternalWorkflowExecutionCancelRequestedEventAttributes to the protobuf v3 wire format
func (val *ExternalWorkflowExecutionCancelRequestedEventAttributes) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type ExternalWorkflowExecutionCancelRequestedEventAttributes from the protobuf v3 wire format
func (val *ExternalWorkflowExecutionCancelRequestedEventAttributes) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *ExternalWorkflowExecutionCancelRequestedEventAttributes) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ExternalWorkflowExecutionCancelRequestedEventAttributes values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ExternalWorkflowExecutionCancelRequestedEventAttributes) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ExternalWorkflowExecutionCancelRequestedEventAttributes
	switch t := that.(type) {
	case *ExternalWorkflowExecutionCancelRequestedEventAttributes:
		that1 = t
	case ExternalWorkflowExecutionCancelRequestedEventAttributes:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type SignalExternalWorkflowExecutionInitiatedEventAttributes to the protobuf v3 wire format
func (val *SignalExternalWorkflowExecutionInitiatedEventAttributes) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type SignalExternalWorkflowExecutionInitiatedEventAttributes from the protobuf v3 wire format
func (val *SignalExternalWorkflowExecutionInitiatedEventAttributes) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *SignalExternalWorkflowExecutionInitiatedEventAttributes) Size() int {
	return proto.Size(val)
}

// Equal returns whether two SignalExternalWorkflowExecutionInitiatedEventAttributes values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *SignalExternalWorkflowExecutionInitiatedEventAttributes) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *SignalExternalWorkflowExecutionInitiatedEventAttributes
	switch t := that.(type) {
	case *SignalExternalWorkflowExecutionInitiatedEventAttributes:
		that1 = t
	case SignalExternalWorkflowExecutionInitiatedEventAttributes:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type SignalExternalWorkflowExecutionFailedEventAttributes to the protobuf v3 wire format
func (val *SignalExternalWorkflowExecutionFailedEventAttributes) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type SignalExternalWorkflowExecutionFailedEventAttributes from the protobuf v3 wire format
func (val *SignalExternalWorkflowExecutionFailedEventAttributes) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *SignalExternalWorkflowExecutionFailedEventAttributes) Size() int {
	return proto.Size(val)
}

// Equal returns whether two SignalExternalWorkflowExecutionFailedEventAttributes values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *SignalExternalWorkflowExecutionFailedEventAttributes) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *SignalExternalWorkflowExecutionFailedEventAttributes
	switch t := that.(type) {
	case *SignalExternalWorkflowExecutionFailedEventAttributes:
		that1 = t
	case SignalExternalWorkflowExecutionFailedEventAttributes:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type ExternalWorkflowExecutionSignaledEventAttributes to the protobuf v3 wire format
func (val *ExternalWorkflowExecutionSignaledEventAttributes) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type ExternalWorkflowExecutionSignaledEventAttributes from the protobuf v3 wire format
func (val *ExternalWorkflowExecutionSignaledEventAttributes) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *ExternalWorkflowExecutionSignaledEventAttributes) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ExternalWorkflowExecutionSignaledEventAttributes values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ExternalWorkflowExecutionSignaledEventAttributes) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ExternalWorkflowExecutionSignaledEventAttributes
	switch t := that.(type) {
	case *ExternalWorkflowExecutionSignaledEventAttributes:
		that1 = t
	case ExternalWorkflowExecutionSignaledEventAttributes:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type UpsertWorkflowSearchAttributesEventAttributes to the protobuf v3 wire format
func (val *UpsertWorkflowSearchAttributesEventAttributes) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type UpsertWorkflowSearchAttributesEventAttributes from the protobuf v3 wire format
func (val *UpsertWorkflowSearchAttributesEventAttributes) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *UpsertWorkflowSearchAttributesEventAttributes) Size() int {
	return proto.Size(val)
}

// Equal returns whether two UpsertWorkflowSearchAttributesEventAttributes values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *UpsertWorkflowSearchAttributesEventAttributes) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *UpsertWorkflowSearchAttributesEventAttributes
	switch t := that.(type) {
	case *UpsertWorkflowSearchAttributesEventAttributes:
		that1 = t
	case UpsertWorkflowSearchAttributesEventAttributes:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type WorkflowPropertiesModifiedEventAttributes to the protobuf v3 wire format
func (val *WorkflowPropertiesModifiedEventAttributes) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type WorkflowPropertiesModifiedEventAttributes from the protobuf v3 wire format
func (val *WorkflowPropertiesModifiedEventAttributes) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *WorkflowPropertiesModifiedEventAttributes) Size() int {
	return proto.Size(val)
}

// Equal returns whether two WorkflowPropertiesModifiedEventAttributes values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *WorkflowPropertiesModifiedEventAttributes) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *WorkflowPropertiesModifiedEventAttributes
	switch t := that.(type) {
	case *WorkflowPropertiesModifiedEventAttributes:
		that1 = t
	case WorkflowPropertiesModifiedEventAttributes:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type StartChildWorkflowExecutionInitiatedEventAttributes to the protobuf v3 wire format
func (val *StartChildWorkflowExecutionInitiatedEventAttributes) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type StartChildWorkflowExecutionInitiatedEventAttributes from the protobuf v3 wire format
func (val *StartChildWorkflowExecutionInitiatedEventAttributes) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *StartChildWorkflowExecutionInitiatedEventAttributes) Size() int {
	return proto.Size(val)
}

// Equal returns whether two StartChildWorkflowExecutionInitiatedEventAttributes values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *StartChildWorkflowExecutionInitiatedEventAttributes) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *StartChildWorkflowExecutionInitiatedEventAttributes
	switch t := that.(type) {
	case *StartChildWorkflowExecutionInitiatedEventAttributes:
		that1 = t
	case StartChildWorkflowExecutionInitiatedEventAttributes:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type StartChildWorkflowExecutionFailedEventAttributes to the protobuf v3 wire format
func (val *StartChildWorkflowExecutionFailedEventAttributes) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type StartChildWorkflowExecutionFailedEventAttributes from the protobuf v3 wire format
func (val *StartChildWorkflowExecutionFailedEventAttributes) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *StartChildWorkflowExecutionFailedEventAttributes) Size() int {
	return proto.Size(val)
}

// Equal returns whether two StartChildWorkflowExecutionFailedEventAttributes values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *StartChildWorkflowExecutionFailedEventAttributes) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *StartChildWorkflowExecutionFailedEventAttributes
	switch t := that.(type) {
	case *StartChildWorkflowExecutionFailedEventAttributes:
		that1 = t
	case StartChildWorkflowExecutionFailedEventAttributes:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type ChildWorkflowExecutionStartedEventAttributes to the protobuf v3 wire format
func (val *ChildWorkflowExecutionStartedEventAttributes) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type ChildWorkflowExecutionStartedEventAttributes from the protobuf v3 wire format
func (val *ChildWorkflowExecutionStartedEventAttributes) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *ChildWorkflowExecutionStartedEventAttributes) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ChildWorkflowExecutionStartedEventAttributes values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ChildWorkflowExecutionStartedEventAttributes) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ChildWorkflowExecutionStartedEventAttributes
	switch t := that.(type) {
	case *ChildWorkflowExecutionStartedEventAttributes:
		that1 = t
	case ChildWorkflowExecutionStartedEventAttributes:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type ChildWorkflowExecutionCompletedEventAttributes to the protobuf v3 wire format
func (val *ChildWorkflowExecutionCompletedEventAttributes) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type ChildWorkflowExecutionCompletedEventAttributes from the protobuf v3 wire format
func (val *ChildWorkflowExecutionCompletedEventAttributes) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *ChildWorkflowExecutionCompletedEventAttributes) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ChildWorkflowExecutionCompletedEventAttributes values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ChildWorkflowExecutionCompletedEventAttributes) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ChildWorkflowExecutionCompletedEventAttributes
	switch t := that.(type) {
	case *ChildWorkflowExecutionCompletedEventAttributes:
		that1 = t
	case ChildWorkflowExecutionCompletedEventAttributes:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type ChildWorkflowExecutionFailedEventAttributes to the protobuf v3 wire format
func (val *ChildWorkflowExecutionFailedEventAttributes) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type ChildWorkflowExecutionFailedEventAttributes from the protobuf v3 wire format
func (val *ChildWorkflowExecutionFailedEventAttributes) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *ChildWorkflowExecutionFailedEventAttributes) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ChildWorkflowExecutionFailedEventAttributes values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ChildWorkflowExecutionFailedEventAttributes) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ChildWorkflowExecutionFailedEventAttributes
	switch t := that.(type) {
	case *ChildWorkflowExecutionFailedEventAttributes:
		that1 = t
	case ChildWorkflowExecutionFailedEventAttributes:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type ChildWorkflowExecutionCanceledEventAttributes to the protobuf v3 wire format
func (val *ChildWorkflowExecutionCanceledEventAttributes) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type ChildWorkflowExecutionCanceledEventAttributes from the protobuf v3 wire format
func (val *ChildWorkflowExecutionCanceledEventAttributes) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *ChildWorkflowExecutionCanceledEventAttributes) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ChildWorkflowExecutionCanceledEventAttributes values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ChildWorkflowExecutionCanceledEventAttributes) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ChildWorkflowExecutionCanceledEventAttributes
	switch t := that.(type) {
	case *ChildWorkflowExecutionCanceledEventAttributes:
		that1 = t
	case ChildWorkflowExecutionCanceledEventAttributes:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type ChildWorkflowExecutionTimedOutEventAttributes to the protobuf v3 wire format
func (val *ChildWorkflowExecutionTimedOutEventAttributes) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type ChildWorkflowExecutionTimedOutEventAttributes from the protobuf v3 wire format
func (val *ChildWorkflowExecutionTimedOutEventAttributes) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *ChildWorkflowExecutionTimedOutEventAttributes) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ChildWorkflowExecutionTimedOutEventAttributes values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ChildWorkflowExecutionTimedOutEventAttributes) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ChildWorkflowExecutionTimedOutEventAttributes
	switch t := that.(type) {
	case *ChildWorkflowExecutionTimedOutEventAttributes:
		that1 = t
	case ChildWorkflowExecutionTimedOutEventAttributes:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type ChildWorkflowExecutionTerminatedEventAttributes to the protobuf v3 wire format
func (val *ChildWorkflowExecutionTerminatedEventAttributes) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type ChildWorkflowExecutionTerminatedEventAttributes from the protobuf v3 wire format
func (val *ChildWorkflowExecutionTerminatedEventAttributes) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *ChildWorkflowExecutionTerminatedEventAttributes) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ChildWorkflowExecutionTerminatedEventAttributes values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ChildWorkflowExecutionTerminatedEventAttributes) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ChildWorkflowExecutionTerminatedEventAttributes
	switch t := that.(type) {
	case *ChildWorkflowExecutionTerminatedEventAttributes:
		that1 = t
	case ChildWorkflowExecutionTerminatedEventAttributes:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type WorkflowPropertiesModifiedExternallyEventAttributes to the protobuf v3 wire format
func (val *WorkflowPropertiesModifiedExternallyEventAttributes) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type WorkflowPropertiesModifiedExternallyEventAttributes from the protobuf v3 wire format
func (val *WorkflowPropertiesModifiedExternallyEventAttributes) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *WorkflowPropertiesModifiedExternallyEventAttributes) Size() int {
	return proto.Size(val)
}

// Equal returns whether two WorkflowPropertiesModifiedExternallyEventAttributes values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *WorkflowPropertiesModifiedExternallyEventAttributes) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *WorkflowPropertiesModifiedExternallyEventAttributes
	switch t := that.(type) {
	case *WorkflowPropertiesModifiedExternallyEventAttributes:
		that1 = t
	case WorkflowPropertiesModifiedExternallyEventAttributes:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type ActivityPropertiesModifiedExternallyEventAttributes to the protobuf v3 wire format
func (val *ActivityPropertiesModifiedExternallyEventAttributes) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type ActivityPropertiesModifiedExternallyEventAttributes from the protobuf v3 wire format
func (val *ActivityPropertiesModifiedExternallyEventAttributes) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *ActivityPropertiesModifiedExternallyEventAttributes) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ActivityPropertiesModifiedExternallyEventAttributes values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ActivityPropertiesModifiedExternallyEventAttributes) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ActivityPropertiesModifiedExternallyEventAttributes
	switch t := that.(type) {
	case *ActivityPropertiesModifiedExternallyEventAttributes:
		that1 = t
	case ActivityPropertiesModifiedExternallyEventAttributes:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type WorkflowExecutionUpdateAcceptedEventAttributes to the protobuf v3 wire format
func (val *WorkflowExecutionUpdateAcceptedEventAttributes) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type WorkflowExecutionUpdateAcceptedEventAttributes from the protobuf v3 wire format
func (val *WorkflowExecutionUpdateAcceptedEventAttributes) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *WorkflowExecutionUpdateAcceptedEventAttributes) Size() int {
	return proto.Size(val)
}

// Equal returns whether two WorkflowExecutionUpdateAcceptedEventAttributes values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *WorkflowExecutionUpdateAcceptedEventAttributes) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *WorkflowExecutionUpdateAcceptedEventAttributes
	switch t := that.(type) {
	case *WorkflowExecutionUpdateAcceptedEventAttributes:
		that1 = t
	case WorkflowExecutionUpdateAcceptedEventAttributes:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type WorkflowExecutionUpdateCompletedEventAttributes to the protobuf v3 wire format
func (val *WorkflowExecutionUpdateCompletedEventAttributes) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type WorkflowExecutionUpdateCompletedEventAttributes from the protobuf v3 wire format
func (val *WorkflowExecutionUpdateCompletedEventAttributes) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *WorkflowExecutionUpdateCompletedEventAttributes) Size() int {
	return proto.Size(val)
}

// Equal returns whether two WorkflowExecutionUpdateCompletedEventAttributes values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *WorkflowExecutionUpdateCompletedEventAttributes) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *WorkflowExecutionUpdateCompletedEventAttributes
	switch t := that.(type) {
	case *WorkflowExecutionUpdateCompletedEventAttributes:
		that1 = t
	case WorkflowExecutionUpdateCompletedEventAttributes:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type WorkflowExecutionUpdateRejectedEventAttributes to the protobuf v3 wire format
func (val *WorkflowExecutionUpdateRejectedEventAttributes) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type WorkflowExecutionUpdateRejectedEventAttributes from the protobuf v3 wire format
func (val *WorkflowExecutionUpdateRejectedEventAttributes) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *WorkflowExecutionUpdateRejectedEventAttributes) Size() int {
	return proto.Size(val)
}

// Equal returns whether two WorkflowExecutionUpdateRejectedEventAttributes values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *WorkflowExecutionUpdateRejectedEventAttributes) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *WorkflowExecutionUpdateRejectedEventAttributes
	switch t := that.(type) {
	case *WorkflowExecutionUpdateRejectedEventAttributes:
		that1 = t
	case WorkflowExecutionUpdateRejectedEventAttributes:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type HistoryEvent to the protobuf v3 wire format
func (val *HistoryEvent) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type HistoryEvent from the protobuf v3 wire format
func (val *HistoryEvent) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *HistoryEvent) Size() int {
	return proto.Size(val)
}

// Equal returns whether two HistoryEvent values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *HistoryEvent) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *HistoryEvent
	switch t := that.(type) {
	case *HistoryEvent:
		that1 = t
	case HistoryEvent:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type History to the protobuf v3 wire format
func (val *History) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type History from the protobuf v3 wire format
func (val *History) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *History) Size() int {
	return proto.Size(val)
}

// Equal returns whether two History values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *History) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *History
	switch t := that.(type) {
	case *History:
		that1 = t
	case History:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
