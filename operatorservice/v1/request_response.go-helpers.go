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

package operatorservice

import (
	"google.golang.org/protobuf/proto"
)

// Marshal an object of type AddSearchAttributesRequest to the protobuf v3 wire format
func (val *AddSearchAttributesRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type AddSearchAttributesRequest from the protobuf v3 wire format
func (val *AddSearchAttributesRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *AddSearchAttributesRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two AddSearchAttributesRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *AddSearchAttributesRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *AddSearchAttributesRequest
	switch t := that.(type) {
	case *AddSearchAttributesRequest:
		that1 = t
	case AddSearchAttributesRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type AddSearchAttributesResponse to the protobuf v3 wire format
func (val *AddSearchAttributesResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type AddSearchAttributesResponse from the protobuf v3 wire format
func (val *AddSearchAttributesResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *AddSearchAttributesResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two AddSearchAttributesResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *AddSearchAttributesResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *AddSearchAttributesResponse
	switch t := that.(type) {
	case *AddSearchAttributesResponse:
		that1 = t
	case AddSearchAttributesResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type RemoveSearchAttributesRequest to the protobuf v3 wire format
func (val *RemoveSearchAttributesRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type RemoveSearchAttributesRequest from the protobuf v3 wire format
func (val *RemoveSearchAttributesRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *RemoveSearchAttributesRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two RemoveSearchAttributesRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *RemoveSearchAttributesRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *RemoveSearchAttributesRequest
	switch t := that.(type) {
	case *RemoveSearchAttributesRequest:
		that1 = t
	case RemoveSearchAttributesRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type RemoveSearchAttributesResponse to the protobuf v3 wire format
func (val *RemoveSearchAttributesResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type RemoveSearchAttributesResponse from the protobuf v3 wire format
func (val *RemoveSearchAttributesResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *RemoveSearchAttributesResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two RemoveSearchAttributesResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *RemoveSearchAttributesResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *RemoveSearchAttributesResponse
	switch t := that.(type) {
	case *RemoveSearchAttributesResponse:
		that1 = t
	case RemoveSearchAttributesResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type ListSearchAttributesRequest to the protobuf v3 wire format
func (val *ListSearchAttributesRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type ListSearchAttributesRequest from the protobuf v3 wire format
func (val *ListSearchAttributesRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *ListSearchAttributesRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ListSearchAttributesRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ListSearchAttributesRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ListSearchAttributesRequest
	switch t := that.(type) {
	case *ListSearchAttributesRequest:
		that1 = t
	case ListSearchAttributesRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type ListSearchAttributesResponse to the protobuf v3 wire format
func (val *ListSearchAttributesResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type ListSearchAttributesResponse from the protobuf v3 wire format
func (val *ListSearchAttributesResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *ListSearchAttributesResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ListSearchAttributesResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ListSearchAttributesResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ListSearchAttributesResponse
	switch t := that.(type) {
	case *ListSearchAttributesResponse:
		that1 = t
	case ListSearchAttributesResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type DeleteNamespaceRequest to the protobuf v3 wire format
func (val *DeleteNamespaceRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type DeleteNamespaceRequest from the protobuf v3 wire format
func (val *DeleteNamespaceRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *DeleteNamespaceRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two DeleteNamespaceRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *DeleteNamespaceRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *DeleteNamespaceRequest
	switch t := that.(type) {
	case *DeleteNamespaceRequest:
		that1 = t
	case DeleteNamespaceRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type DeleteNamespaceResponse to the protobuf v3 wire format
func (val *DeleteNamespaceResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type DeleteNamespaceResponse from the protobuf v3 wire format
func (val *DeleteNamespaceResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *DeleteNamespaceResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two DeleteNamespaceResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *DeleteNamespaceResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *DeleteNamespaceResponse
	switch t := that.(type) {
	case *DeleteNamespaceResponse:
		that1 = t
	case DeleteNamespaceResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type AddOrUpdateRemoteClusterRequest to the protobuf v3 wire format
func (val *AddOrUpdateRemoteClusterRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type AddOrUpdateRemoteClusterRequest from the protobuf v3 wire format
func (val *AddOrUpdateRemoteClusterRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *AddOrUpdateRemoteClusterRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two AddOrUpdateRemoteClusterRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *AddOrUpdateRemoteClusterRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *AddOrUpdateRemoteClusterRequest
	switch t := that.(type) {
	case *AddOrUpdateRemoteClusterRequest:
		that1 = t
	case AddOrUpdateRemoteClusterRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type AddOrUpdateRemoteClusterResponse to the protobuf v3 wire format
func (val *AddOrUpdateRemoteClusterResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type AddOrUpdateRemoteClusterResponse from the protobuf v3 wire format
func (val *AddOrUpdateRemoteClusterResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *AddOrUpdateRemoteClusterResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two AddOrUpdateRemoteClusterResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *AddOrUpdateRemoteClusterResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *AddOrUpdateRemoteClusterResponse
	switch t := that.(type) {
	case *AddOrUpdateRemoteClusterResponse:
		that1 = t
	case AddOrUpdateRemoteClusterResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type RemoveRemoteClusterRequest to the protobuf v3 wire format
func (val *RemoveRemoteClusterRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type RemoveRemoteClusterRequest from the protobuf v3 wire format
func (val *RemoveRemoteClusterRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *RemoveRemoteClusterRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two RemoveRemoteClusterRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *RemoveRemoteClusterRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *RemoveRemoteClusterRequest
	switch t := that.(type) {
	case *RemoveRemoteClusterRequest:
		that1 = t
	case RemoveRemoteClusterRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type RemoveRemoteClusterResponse to the protobuf v3 wire format
func (val *RemoveRemoteClusterResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type RemoveRemoteClusterResponse from the protobuf v3 wire format
func (val *RemoveRemoteClusterResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *RemoveRemoteClusterResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two RemoveRemoteClusterResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *RemoveRemoteClusterResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *RemoveRemoteClusterResponse
	switch t := that.(type) {
	case *RemoveRemoteClusterResponse:
		that1 = t
	case RemoveRemoteClusterResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type ListClustersRequest to the protobuf v3 wire format
func (val *ListClustersRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type ListClustersRequest from the protobuf v3 wire format
func (val *ListClustersRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *ListClustersRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ListClustersRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ListClustersRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ListClustersRequest
	switch t := that.(type) {
	case *ListClustersRequest:
		that1 = t
	case ListClustersRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type ListClustersResponse to the protobuf v3 wire format
func (val *ListClustersResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type ListClustersResponse from the protobuf v3 wire format
func (val *ListClustersResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *ListClustersResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ListClustersResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ListClustersResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ListClustersResponse
	switch t := that.(type) {
	case *ListClustersResponse:
		that1 = t
	case ListClustersResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type ClusterMetadata to the protobuf v3 wire format
func (val *ClusterMetadata) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type ClusterMetadata from the protobuf v3 wire format
func (val *ClusterMetadata) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *ClusterMetadata) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ClusterMetadata values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ClusterMetadata) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ClusterMetadata
	switch t := that.(type) {
	case *ClusterMetadata:
		that1 = t
	case ClusterMetadata:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
