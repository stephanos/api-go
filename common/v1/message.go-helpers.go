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

package common

import (
	"google.golang.org/protobuf/proto"
)

// Marshal an object of type DataBlob to the protobuf v3 wire format
func (val *DataBlob) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type DataBlob from the protobuf v3 wire format
func (val *DataBlob) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *DataBlob) Size() int {
	return proto.Size(val)
}

// Equal returns whether two DataBlob values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *DataBlob) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *DataBlob
	switch t := that.(type) {
	case *DataBlob:
		that1 = t
	case DataBlob:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type Payloads to the protobuf v3 wire format
func (val *Payloads) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type Payloads from the protobuf v3 wire format
func (val *Payloads) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *Payloads) Size() int {
	return proto.Size(val)
}

// Equal returns whether two Payloads values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *Payloads) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *Payloads
	switch t := that.(type) {
	case *Payloads:
		that1 = t
	case Payloads:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type Payload to the protobuf v3 wire format
func (val *Payload) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type Payload from the protobuf v3 wire format
func (val *Payload) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *Payload) Size() int {
	return proto.Size(val)
}

// Equal returns whether two Payload values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *Payload) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *Payload
	switch t := that.(type) {
	case *Payload:
		that1 = t
	case Payload:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type SearchAttributes to the protobuf v3 wire format
func (val *SearchAttributes) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type SearchAttributes from the protobuf v3 wire format
func (val *SearchAttributes) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *SearchAttributes) Size() int {
	return proto.Size(val)
}

// Equal returns whether two SearchAttributes values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *SearchAttributes) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *SearchAttributes
	switch t := that.(type) {
	case *SearchAttributes:
		that1 = t
	case SearchAttributes:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type Memo to the protobuf v3 wire format
func (val *Memo) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type Memo from the protobuf v3 wire format
func (val *Memo) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *Memo) Size() int {
	return proto.Size(val)
}

// Equal returns whether two Memo values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *Memo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *Memo
	switch t := that.(type) {
	case *Memo:
		that1 = t
	case Memo:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type Header to the protobuf v3 wire format
func (val *Header) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type Header from the protobuf v3 wire format
func (val *Header) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *Header) Size() int {
	return proto.Size(val)
}

// Equal returns whether two Header values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *Header) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *Header
	switch t := that.(type) {
	case *Header:
		that1 = t
	case Header:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type WorkflowExecution to the protobuf v3 wire format
func (val *WorkflowExecution) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type WorkflowExecution from the protobuf v3 wire format
func (val *WorkflowExecution) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *WorkflowExecution) Size() int {
	return proto.Size(val)
}

// Equal returns whether two WorkflowExecution values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *WorkflowExecution) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *WorkflowExecution
	switch t := that.(type) {
	case *WorkflowExecution:
		that1 = t
	case WorkflowExecution:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type WorkflowType to the protobuf v3 wire format
func (val *WorkflowType) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type WorkflowType from the protobuf v3 wire format
func (val *WorkflowType) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *WorkflowType) Size() int {
	return proto.Size(val)
}

// Equal returns whether two WorkflowType values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *WorkflowType) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *WorkflowType
	switch t := that.(type) {
	case *WorkflowType:
		that1 = t
	case WorkflowType:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type ActivityType to the protobuf v3 wire format
func (val *ActivityType) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type ActivityType from the protobuf v3 wire format
func (val *ActivityType) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *ActivityType) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ActivityType values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ActivityType) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ActivityType
	switch t := that.(type) {
	case *ActivityType:
		that1 = t
	case ActivityType:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type RetryPolicy to the protobuf v3 wire format
func (val *RetryPolicy) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type RetryPolicy from the protobuf v3 wire format
func (val *RetryPolicy) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *RetryPolicy) Size() int {
	return proto.Size(val)
}

// Equal returns whether two RetryPolicy values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *RetryPolicy) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *RetryPolicy
	switch t := that.(type) {
	case *RetryPolicy:
		that1 = t
	case RetryPolicy:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type MeteringMetadata to the protobuf v3 wire format
func (val *MeteringMetadata) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type MeteringMetadata from the protobuf v3 wire format
func (val *MeteringMetadata) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *MeteringMetadata) Size() int {
	return proto.Size(val)
}

// Equal returns whether two MeteringMetadata values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *MeteringMetadata) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *MeteringMetadata
	switch t := that.(type) {
	case *MeteringMetadata:
		that1 = t
	case MeteringMetadata:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type WorkerVersionStamp to the protobuf v3 wire format
func (val *WorkerVersionStamp) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type WorkerVersionStamp from the protobuf v3 wire format
func (val *WorkerVersionStamp) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *WorkerVersionStamp) Size() int {
	return proto.Size(val)
}

// Equal returns whether two WorkerVersionStamp values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *WorkerVersionStamp) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *WorkerVersionStamp
	switch t := that.(type) {
	case *WorkerVersionStamp:
		that1 = t
	case WorkerVersionStamp:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type WorkerVersionCapabilities to the protobuf v3 wire format
func (val *WorkerVersionCapabilities) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type WorkerVersionCapabilities from the protobuf v3 wire format
func (val *WorkerVersionCapabilities) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *WorkerVersionCapabilities) Size() int {
	return proto.Size(val)
}

// Equal returns whether two WorkerVersionCapabilities values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *WorkerVersionCapabilities) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *WorkerVersionCapabilities
	switch t := that.(type) {
	case *WorkerVersionCapabilities:
		that1 = t
	case WorkerVersionCapabilities:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
