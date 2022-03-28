/*
Copyright 2021 The KServe Authors.

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

package v1alpha1

import (
	"github.com/kserve/kserve/pkg/constants"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:openapi-gen=true
type SupportedModelFormat struct {
	// Name of the model format.
	// +required
	Name string `json:"name"`
	// Version of the model format.
	// Used in validating that a predictor is supported by a runtime.
	// Can be "major", "major.minor" or "major.minor.patch".
	// +optional
	Version *string `json:"version,omitempty"`
	// Set to true to allow the ServingRuntime to be used for automatic model placement if
	// this model format is specified with no explicit runtime.
	// +optional
	AutoSelect *bool `json:"autoSelect,omitempty"`
}

// +k8s:openapi-gen=true
type StorageHelper struct {
	// +optional
	Disabled bool `json:"disabled,omitempty"`
}

// +k8s:openapi-gen=true
type ServingRuntimePodSpec struct {
	// List of containers belonging to the pod.
	// Containers cannot currently be added or removed.
	// There must be at least one container in a Pod.
	// Cannot be updated.
	// +patchMergeKey=name
	// +patchStrategy=merge
	Containers []corev1.Container `json:"containers" patchStrategy:"merge" patchMergeKey:"name" validate:"required"`

	// NodeSelector is a selector which must be true for the pod to fit on a node.
	// Selector which must match a node's labels for the pod to be scheduled on that node.
	// More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/
	// +optional
	NodeSelector map[string]string `json:"nodeSelector,omitempty"`

	// If specified, the pod's scheduling constraints
	// +optional
	Affinity *corev1.Affinity `json:"affinity,omitempty"`

	// If specified, the pod's tolerations.
	// +optional
	Tolerations []corev1.Toleration `json:"tolerations,omitempty"`

	// Possibly other things here
}

// ServingRuntimeSpec defines the desired state of ServingRuntime. This spec is currently provisional
// and are subject to change as details regarding single-model serving and multi-model serving
// are hammered out.
// +k8s:openapi-gen=true
type ServingRuntimeSpec struct {
	// Model formats and version supported by this runtime
	SupportedModelFormats []SupportedModelFormat `json:"supportedModelFormats,omitempty"`

	// Whether this ServingRuntime is intended for multi-model usage or not.
	// +optional
	MultiModel *bool `json:"multiModel,omitempty"`

	// Set to true to disable use of this runtime
	// +optional
	Disabled *bool `json:"disabled,omitempty"`

	// Supported protocol version (i.e. v1, v2 or empty for unknown protocol version).
	// +optional
	ProtocolVersion constants.InferenceServiceProtocol `json:"protocolVersion"`

	ServingRuntimePodSpec `json:",inline"`

	// The following fields apply to ModelMesh deployments.

	// Name for each of the Endpoint fields is either like "port:1234" or "unix:/tmp/kserve/grpc.sock"

	// Grpc endpoint for internal model-management (implementing mmesh.ModelRuntime gRPC service)
	// Assumed to be single-model runtime if omitted
	// +optional
	GrpcMultiModelManagementEndpoint *string `json:"grpcEndpoint,omitempty"`

	// Grpc endpoint for inferencing
	// +optional
	GrpcDataEndpoint *string `json:"grpcDataEndpoint,omitempty"`
	// HTTP endpoint for inferencing
	// +optional
	HTTPDataEndpoint *string `json:"httpDataEndpoint,omitempty"`

	// Configure the number of replicas in the Deployment generated by this ServingRuntime
	// If specified, this overrides the podsPerRuntime configuration value
	// +optional
	Replicas *uint16 `json:"replicas,omitempty"`

	// Configuration for this runtime's use of the storage helper (model puller)
	// It is enabled unless explicitly disabled
	// +optional
	StorageHelper *StorageHelper `json:"storageHelper,omitempty"`

	// Provide the details about built-in runtime adapter
	// +optional
	BuiltInAdapter *BuiltInAdapter `json:"builtInAdapter,omitempty"`
}

// ServingRuntimeStatus defines the observed state of ServingRuntime
// +k8s:openapi-gen=true
type ServingRuntimeStatus struct {
}

// ServerType constant for specifying the runtime name
// +k8s:openapi-gen=true
// +kubebuilder:validation:Enum=triton;mlserver
type ServerType string

// ServerType Enum
const (
	// Model server is Triton
	Triton ServerType = "triton"
	// Model server is MLServer
	MLServer ServerType = "mlserver"
)

// +k8s:openapi-gen=true
type BuiltInAdapter struct {
	// ServerType can be one of triton/mlserver and the runtime's container must have the same name
	ServerType ServerType `json:"serverType,omitempty"`
	// Port which the runtime server listens for model management requests
	RuntimeManagementPort int `json:"runtimeManagementPort,omitempty"`
	// Fixed memory overhead to subtract from runtime container's memory allocation to determine model capacity
	MemBufferBytes int `json:"memBufferBytes,omitempty"`
	// Timeout for model loading operations in milliseconds
	ModelLoadingTimeoutMillis int `json:"modelLoadingTimeoutMillis,omitempty"`
	// Environment variables used to control other aspects of the built-in adapter's behaviour (uncommon)
	Env []corev1.EnvVar `json:"env,omitempty"`
}

// ServingRuntime is the Schema for the servingruntimes API
// +k8s:openapi-gen=true
// +kubebuilder:object:root=true
// +kubebuilder:printcolumn:name="Disabled",type="boolean",JSONPath=".spec.disabled"
// +kubebuilder:printcolumn:name="ModelType",type="string",JSONPath=".spec.supportedModelFormats[*].name"
// +kubebuilder:printcolumn:name="Containers",type="string",JSONPath=".spec.containers[*].name"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
type ServingRuntime struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ServingRuntimeSpec   `json:"spec,omitempty"`
	Status ServingRuntimeStatus `json:"status,omitempty"`
}

// ServingRuntimeList contains a list of ServingRuntime
// +k8s:openapi-gen=true
// +kubebuilder:object:root=true
type ServingRuntimeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServingRuntime `json:"items"`
}

// ClusterServingRuntime is the Schema for the servingruntimes API
// +k8s:openapi-gen=true
// +kubebuilder:object:root=true
// +kubebuilder:resource:scope="Cluster"
// +kubebuilder:printcolumn:name="Disabled",type="boolean",JSONPath=".spec.disabled"
// +kubebuilder:printcolumn:name="ModelType",type="string",JSONPath=".spec.supportedModelFormats[*].name"
// +kubebuilder:printcolumn:name="Containers",type="string",JSONPath=".spec.containers[*].name"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
type ClusterServingRuntime struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ServingRuntimeSpec   `json:"spec,omitempty"`
	Status ServingRuntimeStatus `json:"status,omitempty"`
}

// ServingRuntimeList contains a list of ServingRuntime
// +k8s:openapi-gen=true
// +kubebuilder:object:root=true
type ClusterServingRuntimeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterServingRuntime `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ServingRuntime{}, &ServingRuntimeList{})
	SchemeBuilder.Register(&ClusterServingRuntime{}, &ClusterServingRuntimeList{})
}

func (srSpec *ServingRuntimeSpec) IsDisabled() bool {
	return srSpec.Disabled != nil && *srSpec.Disabled
}

func (srSpec *ServingRuntimeSpec) IsMultiModelRuntime() bool {
	return srSpec.MultiModel != nil && *srSpec.MultiModel
}

func (srSpec *ServingRuntimeSpec) IsProtocolVersionSupported(modelProtocolVersion constants.InferenceServiceProtocol) bool {
	return srSpec.ProtocolVersion == modelProtocolVersion || srSpec.ProtocolVersion == constants.ProtocolUnknown
}
