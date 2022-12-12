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

// KeepjobGroupSpec defines the desired state of KeepjobGroup
type KeepjobGroupSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// MinMember  表示在这个group中最少需要的任务数，当其中的任务数量没有达到这个量级，没有任务会启动，当达到这个量级之后集群中的资源满足所有任务同时启动，那么任务就会被启动
	MinMember string    `json:"minmember,omitempty"`
	Jobs      []KeepJob `json:"jobs,omitempty"`
}

// KeepjobGroupStatus defines the observed state of KeepjobGroup
type KeepjobGroupStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+genclient
//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// KeepjobGroup is the Schema for the keepjobgroups API
type KeepjobGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KeepjobGroupSpec   `json:"spec,omitempty"`
	Status KeepjobGroupStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// KeepjobGroupList contains a list of KeepjobGroup
type KeepjobGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KeepjobGroup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&KeepjobGroup{}, &KeepjobGroupList{})
}
