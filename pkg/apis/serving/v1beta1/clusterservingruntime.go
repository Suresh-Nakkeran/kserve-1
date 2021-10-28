/*

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

package v1beta1

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NOTE: json tags are required. Any new fields you add must have json tags for the fields to be serialized.
type ModelType struct {
	// +required
	Name string `json:"name"`
	// +optional
	Version *string `json:"version,omitempty"`
}

// ClusterServingRuntimeSpec defines the desired state of ClusterServingRuntime
type ClusterServingRuntimeSpec struct {
	// Model formats and version supported by this runtime
	SupportedModelTypes []ModelType `json:"supportedModelTypes,omitempty"`
	// Set to true to disable use of this runtime
	// +optional
	Disabled *bool `json:"disabled,omitempty"`
	// List of containers belonging to the pod.
	// Containers cannot currently be added or removed.
	// There must be at least one container in a Pod.
	// +optional
	Containers []v1.Container `json:"containers" patchStrategy:"merge" patchMergeKey:"name" validate:"required"`
}

// ClusterServingRuntimeStatus defines the observed state of ClusterServingRuntime
type ClusterServingRuntimeStatus struct {
}

// ClusterServingRuntime is the Schema for the ClusterServingRuntimes API
// +k8s:openapi-gen=true
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Disabled",type="boolean",JSONPath=".spec.disabled"
// +kubebuilder:printcolumn:name="ModelType",type="string",JSONPath=".spec.supportedModelTypes[*].name"
// +kubebuilder:printcolumn:name="Containers",type="string",JSONPath=".spec.containers[*].name"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:path=clusterservingruntimes,shortName=csrt,singular=clusterservingruntime
type ClusterServingRuntime struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec ClusterServingRuntimeSpec `json:"spec,omitempty"`

	// +kubebuilder:pruning:PreserveUnknownFields
	Status ClusterServingRuntimeStatus `json:"status,omitempty"`
}

// ClusterServingRuntimeList contains a list of ClusterServingRuntime
// +k8s:openapi-gen=true
// +kubebuilder:object:root=true
type ClusterServingRuntimeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// +listType=set
	Items []ClusterServingRuntime `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ClusterServingRuntime{}, &ClusterServingRuntimeList{})
}

func (csrt ClusterServingRuntime) Disabled() bool {
	return csrt.Spec.Disabled != nil && *csrt.Spec.Disabled
}
