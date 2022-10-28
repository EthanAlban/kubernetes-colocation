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

// KeepJobSpec defines the desired state of KeepJob
type KeepJobSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	JobName      string      `json:"job_name,omitempty"`
	Namespace    string      `json:"namespace,omitempty"`
	Image        string      `json:"image,omitempty"`
	JobType      string      `json:"job_type,omitempty"`
	Replica      int         `json:"replica,omitempty"`
	Weight       int         `json:"weight,omitempty"`
	Plugins      []string    `json:"plugins,omitempty"`
	JobQueueName string      `json:"job_queue_name,omitempty"`
	CreatingTime metav1.Time `json:"creating_time,omitempty"`
}

// KeepJobStatus defines the observed state of KeepJob
type KeepJobStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+genclient
//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// KeepJob is the Schema for the keepjobs API
type KeepJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KeepJobSpec   `json:"spec,omitempty"`
	Status KeepJobStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// KeepJobList contains a list of KeepJob
type KeepJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KeepJob `json:"items"`
}

func init() {
	SchemeBuilder.Register(&KeepJob{}, &KeepJobList{})
}
