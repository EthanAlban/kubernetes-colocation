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

// KeepQueueSpec defines the desired state of KeepQueue
type KeepQueueSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of KeepQueue. Edit keepqueue_types.go to remove/update
	Foo       string `json:"foo,omitempty"`
	QueueName string `json:"queueName,omitempty"`
	// 监管的任务的 namespace/jobName
	OwnJobs []KeepJob `json:"ownJobs,omitempty"`
	// 运行在本队列中的任务是否可以在资源不足时被抢占
	Reclaimable bool `json:"reclaimable,omitempty"`
	// 运行在本队列中的任务所使用的时间片的比重，比重越大其使用资源的代价越大
	Weight int `json:"weight,omitempty"`
}

// KeepQueueStatus defines the observed state of KeepQueue
type KeepQueueStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +genclient
// +genclient:nonNamespaced
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
// KeepQueue is the Schema for the keepqueues API
type KeepQueue struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KeepQueueSpec   `json:"spec,omitempty"`
	Status KeepQueueStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// KeepQueueList contains a list of KeepQueue
type KeepQueueList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KeepQueue `json:"items"`
}

func init() {
	SchemeBuilder.Register(&KeepQueue{}, &KeepQueueList{})
}
