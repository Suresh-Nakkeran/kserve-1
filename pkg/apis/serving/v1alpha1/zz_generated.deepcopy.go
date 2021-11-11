// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"knative.dev/pkg/apis"
	duckv1 "knative.dev/pkg/apis/duck/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuiltInAdapter) DeepCopyInto(out *BuiltInAdapter) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuiltInAdapter.
func (in *BuiltInAdapter) DeepCopy() *BuiltInAdapter {
	if in == nil {
		return nil
	}
	out := new(BuiltInAdapter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterServingRuntime) DeepCopyInto(out *ClusterServingRuntime) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterServingRuntime.
func (in *ClusterServingRuntime) DeepCopy() *ClusterServingRuntime {
	if in == nil {
		return nil
	}
	out := new(ClusterServingRuntime)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterServingRuntime) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterServingRuntimeList) DeepCopyInto(out *ClusterServingRuntimeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterServingRuntime, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterServingRuntimeList.
func (in *ClusterServingRuntimeList) DeepCopy() *ClusterServingRuntimeList {
	if in == nil {
		return nil
	}
	out := new(ClusterServingRuntimeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterServingRuntimeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Container) DeepCopyInto(out *Container) {
	*out = *in
	if in.Command != nil {
		in, out := &in.Command, &out.Command
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.Resources.DeepCopyInto(&out.Resources)
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LivenessProbe != nil {
		in, out := &in.LivenessProbe, &out.LivenessProbe
		*out = new(v1.Probe)
		(*in).DeepCopyInto(*out)
	}
	if in.ReadinessProbe != nil {
		in, out := &in.ReadinessProbe, &out.ReadinessProbe
		*out = new(v1.Probe)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Container.
func (in *Container) DeepCopy() *Container {
	if in == nil {
		return nil
	}
	out := new(Container)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Framework) DeepCopyInto(out *Framework) {
	*out = *in
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Framework.
func (in *Framework) DeepCopy() *Framework {
	if in == nil {
		return nil
	}
	out := new(Framework)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModelSpec) DeepCopyInto(out *ModelSpec) {
	*out = *in
	out.Memory = in.Memory.DeepCopy()
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModelSpec.
func (in *ModelSpec) DeepCopy() *ModelSpec {
	if in == nil {
		return nil
	}
	out := new(ModelSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServingRuntime) DeepCopyInto(out *ServingRuntime) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServingRuntime.
func (in *ServingRuntime) DeepCopy() *ServingRuntime {
	if in == nil {
		return nil
	}
	out := new(ServingRuntime)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServingRuntime) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServingRuntimeList) DeepCopyInto(out *ServingRuntimeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ServingRuntime, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServingRuntimeList.
func (in *ServingRuntimeList) DeepCopy() *ServingRuntimeList {
	if in == nil {
		return nil
	}
	out := new(ServingRuntimeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServingRuntimeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServingRuntimePodSpec) DeepCopyInto(out *ServingRuntimePodSpec) {
	*out = *in
	if in.Containers != nil {
		in, out := &in.Containers, &out.Containers
		*out = make([]Container, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(v1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServingRuntimePodSpec.
func (in *ServingRuntimePodSpec) DeepCopy() *ServingRuntimePodSpec {
	if in == nil {
		return nil
	}
	out := new(ServingRuntimePodSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServingRuntimeSpec) DeepCopyInto(out *ServingRuntimeSpec) {
	*out = *in
	if in.SupportedModelTypes != nil {
		in, out := &in.SupportedModelTypes, &out.SupportedModelTypes
		*out = make([]Framework, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Disabled != nil {
		in, out := &in.Disabled, &out.Disabled
		*out = new(bool)
		**out = **in
	}
	in.ServingRuntimePodSpec.DeepCopyInto(&out.ServingRuntimePodSpec)
	if in.GrpcMultiModelManagementEndpoint != nil {
		in, out := &in.GrpcMultiModelManagementEndpoint, &out.GrpcMultiModelManagementEndpoint
		*out = new(string)
		**out = **in
	}
	if in.GrpcDataEndpoint != nil {
		in, out := &in.GrpcDataEndpoint, &out.GrpcDataEndpoint
		*out = new(string)
		**out = **in
	}
	if in.HTTPDataEndpoint != nil {
		in, out := &in.HTTPDataEndpoint, &out.HTTPDataEndpoint
		*out = new(string)
		**out = **in
	}
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(uint16)
		**out = **in
	}
	if in.StorageHelper != nil {
		in, out := &in.StorageHelper, &out.StorageHelper
		*out = new(StorageHelper)
		**out = **in
	}
	if in.BuiltInAdapter != nil {
		in, out := &in.BuiltInAdapter, &out.BuiltInAdapter
		*out = new(BuiltInAdapter)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServingRuntimeSpec.
func (in *ServingRuntimeSpec) DeepCopy() *ServingRuntimeSpec {
	if in == nil {
		return nil
	}
	out := new(ServingRuntimeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServingRuntimeStatus) DeepCopyInto(out *ServingRuntimeStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServingRuntimeStatus.
func (in *ServingRuntimeStatus) DeepCopy() *ServingRuntimeStatus {
	if in == nil {
		return nil
	}
	out := new(ServingRuntimeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageHelper) DeepCopyInto(out *StorageHelper) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageHelper.
func (in *StorageHelper) DeepCopy() *StorageHelper {
	if in == nil {
		return nil
	}
	out := new(StorageHelper)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrainedModel) DeepCopyInto(out *TrainedModel) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrainedModel.
func (in *TrainedModel) DeepCopy() *TrainedModel {
	if in == nil {
		return nil
	}
	out := new(TrainedModel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TrainedModel) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrainedModelList) DeepCopyInto(out *TrainedModelList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TrainedModel, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrainedModelList.
func (in *TrainedModelList) DeepCopy() *TrainedModelList {
	if in == nil {
		return nil
	}
	out := new(TrainedModelList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TrainedModelList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrainedModelSpec) DeepCopyInto(out *TrainedModelSpec) {
	*out = *in
	in.Model.DeepCopyInto(&out.Model)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrainedModelSpec.
func (in *TrainedModelSpec) DeepCopy() *TrainedModelSpec {
	if in == nil {
		return nil
	}
	out := new(TrainedModelSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrainedModelStatus) DeepCopyInto(out *TrainedModelStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
	if in.URL != nil {
		in, out := &in.URL, &out.URL
		*out = new(apis.URL)
		(*in).DeepCopyInto(*out)
	}
	if in.Address != nil {
		in, out := &in.Address, &out.Address
		*out = new(duckv1.Addressable)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrainedModelStatus.
func (in *TrainedModelStatus) DeepCopy() *TrainedModelStatus {
	if in == nil {
		return nil
	}
	out := new(TrainedModelStatus)
	in.DeepCopyInto(out)
	return out
}
