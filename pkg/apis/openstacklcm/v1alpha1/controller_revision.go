// Copyright 2019 The OpenstackLcm Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ControllerRevision implements an immutable snapshot of state data. Clients
// are responsible for serializing and deserializing the objects that contain
// their internal state.
// Once a ControllerRevision has been successfully created, it can not be updated.
// The API Server will fail validation of all requests that attempt to mutate
// the Data field. ControllerRevisions may, however, be deleted. Note that, due to its use by both
// the DaemonSet and StatefulSet controllers for update and rollback, this object is beta. However,
// it may be subject to name and representation changes in future releases, and clients should not
// depend on its stability. It is primarily for internal use by controllers.
type ControllerRevision struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Data is the serialized representation of the state.
	Data runtime.RawExtension `json:"data,omitempty" protobuf:"bytes,2,opt,name=data"`

	// Revision indicates the revision of the state represented by Data.
	Revision int64 `json:"revision" protobuf:"varint,3,opt,name=revision"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ControllerRevisionList is a resource containing a list of ControllerRevision objects.
type ControllerRevisionList struct {
	metav1.TypeMeta `json:",inline"`

	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
	// +optional
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Items is the list of ControllerRevisions
	Items []ControllerRevision `json:"items" protobuf:"bytes,2,rep,name=items"`
}
