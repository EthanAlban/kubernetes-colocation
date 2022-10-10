/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// VirtualNodeSpec defines the desired state of VirtualNode
type VirtualNodeSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of VirtualNode. Edit virtualnode_types.go to remove/update
	NodeName string `json:"nodename,omitempty"`
	Detail   string `json:"detail,omitempty"`
}

// VirtualNodeStatus defines the observed state of VirtualNode
type VirtualNodeStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Created      bool `json:"created,omitempty"`
	CpuCapacity  int  `json:"cpu_capacity,omitempty"`
	CpuUsage     int  `json:"cpu_usage,omitempty"`
	MemUsage     int  `json:"mem_usage,omitempty"`
	MemCapacity  int  `json:"mem_capacity,omitempty"`
	DiskUsage    int  `json:"disk_usage,omitempty"`
	DiskCapacity int  `json:"disk_capacity,omitempty"`
}

//+genclient
//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// VirtualNode is the Schema for the virtualnodes API
type VirtualNode struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VirtualNodeSpec   `json:"spec,omitempty"`
	Status VirtualNodeStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// VirtualNodeList contains a list of VirtualNode
type VirtualNodeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualNode `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VirtualNode{}, &VirtualNodeList{})
}
