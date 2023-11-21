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

package update

import (
	"google.golang.org/protobuf/proto"
)

// Marshal an object of type WaitPolicy to the protobuf v3 wire format
func (val *WaitPolicy) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type WaitPolicy from the protobuf v3 wire format
func (val *WaitPolicy) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *WaitPolicy) Size() int {
	return proto.Size(val)
}

// Equal returns whether two WaitPolicy values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *WaitPolicy) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *WaitPolicy
	switch t := that.(type) {
	case *WaitPolicy:
		that1 = t
	case WaitPolicy:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type UpdateRef to the protobuf v3 wire format
func (val *UpdateRef) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type UpdateRef from the protobuf v3 wire format
func (val *UpdateRef) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *UpdateRef) Size() int {
	return proto.Size(val)
}

// Equal returns whether two UpdateRef values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *UpdateRef) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *UpdateRef
	switch t := that.(type) {
	case *UpdateRef:
		that1 = t
	case UpdateRef:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type Outcome to the protobuf v3 wire format
func (val *Outcome) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type Outcome from the protobuf v3 wire format
func (val *Outcome) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *Outcome) Size() int {
	return proto.Size(val)
}

// Equal returns whether two Outcome values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *Outcome) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *Outcome
	switch t := that.(type) {
	case *Outcome:
		that1 = t
	case Outcome:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type Meta to the protobuf v3 wire format
func (val *Meta) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type Meta from the protobuf v3 wire format
func (val *Meta) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *Meta) Size() int {
	return proto.Size(val)
}

// Equal returns whether two Meta values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *Meta) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *Meta
	switch t := that.(type) {
	case *Meta:
		that1 = t
	case Meta:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type Input to the protobuf v3 wire format
func (val *Input) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type Input from the protobuf v3 wire format
func (val *Input) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *Input) Size() int {
	return proto.Size(val)
}

// Equal returns whether two Input values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *Input) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *Input
	switch t := that.(type) {
	case *Input:
		that1 = t
	case Input:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type Request to the protobuf v3 wire format
func (val *Request) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type Request from the protobuf v3 wire format
func (val *Request) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *Request) Size() int {
	return proto.Size(val)
}

// Equal returns whether two Request values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *Request) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *Request
	switch t := that.(type) {
	case *Request:
		that1 = t
	case Request:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type Rejection to the protobuf v3 wire format
func (val *Rejection) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type Rejection from the protobuf v3 wire format
func (val *Rejection) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *Rejection) Size() int {
	return proto.Size(val)
}

// Equal returns whether two Rejection values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *Rejection) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *Rejection
	switch t := that.(type) {
	case *Rejection:
		that1 = t
	case Rejection:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type Acceptance to the protobuf v3 wire format
func (val *Acceptance) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type Acceptance from the protobuf v3 wire format
func (val *Acceptance) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *Acceptance) Size() int {
	return proto.Size(val)
}

// Equal returns whether two Acceptance values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *Acceptance) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *Acceptance
	switch t := that.(type) {
	case *Acceptance:
		that1 = t
	case Acceptance:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type Response to the protobuf v3 wire format
func (val *Response) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type Response from the protobuf v3 wire format
func (val *Response) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *Response) Size() int {
	return proto.Size(val)
}

// Equal returns whether two Response values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *Response) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *Response
	switch t := that.(type) {
	case *Response:
		that1 = t
	case Response:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
