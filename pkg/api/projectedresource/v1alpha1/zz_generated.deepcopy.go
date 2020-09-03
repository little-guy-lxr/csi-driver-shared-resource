// +build !ignore_autogenerated

/*
Copyright The OpenShift authors.

SPDX-License-Identifier: Apache-2.0
*/

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackingResource) DeepCopyInto(out *BackingResource) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackingResource.
func (in *BackingResource) DeepCopy() *BackingResource {
	if in == nil {
		return nil
	}
	out := new(BackingResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Share) DeepCopyInto(out *Share) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Share.
func (in *Share) DeepCopy() *Share {
	if in == nil {
		return nil
	}
	out := new(Share)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Share) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShareList) DeepCopyInto(out *ShareList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Share, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShareList.
func (in *ShareList) DeepCopy() *ShareList {
	if in == nil {
		return nil
	}
	out := new(ShareList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ShareList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShareSpec) DeepCopyInto(out *ShareSpec) {
	*out = *in
	out.BackingResource = in.BackingResource
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShareSpec.
func (in *ShareSpec) DeepCopy() *ShareSpec {
	if in == nil {
		return nil
	}
	out := new(ShareSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShareStatus) DeepCopyInto(out *ShareStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShareStatus.
func (in *ShareStatus) DeepCopy() *ShareStatus {
	if in == nil {
		return nil
	}
	out := new(ShareStatus)
	in.DeepCopyInto(out)
	return out
}
