//go:build !ignore_autogenerated

// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocationTagBinding) DeepCopyInto(out *LocationTagBinding) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocationTagBinding.
func (in *LocationTagBinding) DeepCopy() *LocationTagBinding {
	if in == nil {
		return nil
	}
	out := new(LocationTagBinding)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LocationTagBinding) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocationTagBindingInitParameters) DeepCopyInto(out *LocationTagBindingInitParameters) {
	*out = *in
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.Parent != nil {
		in, out := &in.Parent, &out.Parent
		*out = new(string)
		**out = **in
	}
	if in.TagValue != nil {
		in, out := &in.TagValue, &out.TagValue
		*out = new(string)
		**out = **in
	}
	if in.TagValueRef != nil {
		in, out := &in.TagValueRef, &out.TagValueRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.TagValueSelector != nil {
		in, out := &in.TagValueSelector, &out.TagValueSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocationTagBindingInitParameters.
func (in *LocationTagBindingInitParameters) DeepCopy() *LocationTagBindingInitParameters {
	if in == nil {
		return nil
	}
	out := new(LocationTagBindingInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocationTagBindingList) DeepCopyInto(out *LocationTagBindingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LocationTagBinding, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocationTagBindingList.
func (in *LocationTagBindingList) DeepCopy() *LocationTagBindingList {
	if in == nil {
		return nil
	}
	out := new(LocationTagBindingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LocationTagBindingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocationTagBindingObservation) DeepCopyInto(out *LocationTagBindingObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Parent != nil {
		in, out := &in.Parent, &out.Parent
		*out = new(string)
		**out = **in
	}
	if in.TagValue != nil {
		in, out := &in.TagValue, &out.TagValue
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocationTagBindingObservation.
func (in *LocationTagBindingObservation) DeepCopy() *LocationTagBindingObservation {
	if in == nil {
		return nil
	}
	out := new(LocationTagBindingObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocationTagBindingParameters) DeepCopyInto(out *LocationTagBindingParameters) {
	*out = *in
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.Parent != nil {
		in, out := &in.Parent, &out.Parent
		*out = new(string)
		**out = **in
	}
	if in.TagValue != nil {
		in, out := &in.TagValue, &out.TagValue
		*out = new(string)
		**out = **in
	}
	if in.TagValueRef != nil {
		in, out := &in.TagValueRef, &out.TagValueRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.TagValueSelector != nil {
		in, out := &in.TagValueSelector, &out.TagValueSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocationTagBindingParameters.
func (in *LocationTagBindingParameters) DeepCopy() *LocationTagBindingParameters {
	if in == nil {
		return nil
	}
	out := new(LocationTagBindingParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocationTagBindingSpec) DeepCopyInto(out *LocationTagBindingSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocationTagBindingSpec.
func (in *LocationTagBindingSpec) DeepCopy() *LocationTagBindingSpec {
	if in == nil {
		return nil
	}
	out := new(LocationTagBindingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocationTagBindingStatus) DeepCopyInto(out *LocationTagBindingStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocationTagBindingStatus.
func (in *LocationTagBindingStatus) DeepCopy() *LocationTagBindingStatus {
	if in == nil {
		return nil
	}
	out := new(LocationTagBindingStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TagBinding) DeepCopyInto(out *TagBinding) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TagBinding.
func (in *TagBinding) DeepCopy() *TagBinding {
	if in == nil {
		return nil
	}
	out := new(TagBinding)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TagBinding) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TagBindingInitParameters) DeepCopyInto(out *TagBindingInitParameters) {
	*out = *in
	if in.Parent != nil {
		in, out := &in.Parent, &out.Parent
		*out = new(string)
		**out = **in
	}
	if in.TagValue != nil {
		in, out := &in.TagValue, &out.TagValue
		*out = new(string)
		**out = **in
	}
	if in.TagValueRef != nil {
		in, out := &in.TagValueRef, &out.TagValueRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.TagValueSelector != nil {
		in, out := &in.TagValueSelector, &out.TagValueSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TagBindingInitParameters.
func (in *TagBindingInitParameters) DeepCopy() *TagBindingInitParameters {
	if in == nil {
		return nil
	}
	out := new(TagBindingInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TagBindingList) DeepCopyInto(out *TagBindingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TagBinding, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TagBindingList.
func (in *TagBindingList) DeepCopy() *TagBindingList {
	if in == nil {
		return nil
	}
	out := new(TagBindingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TagBindingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TagBindingObservation) DeepCopyInto(out *TagBindingObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Parent != nil {
		in, out := &in.Parent, &out.Parent
		*out = new(string)
		**out = **in
	}
	if in.TagValue != nil {
		in, out := &in.TagValue, &out.TagValue
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TagBindingObservation.
func (in *TagBindingObservation) DeepCopy() *TagBindingObservation {
	if in == nil {
		return nil
	}
	out := new(TagBindingObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TagBindingParameters) DeepCopyInto(out *TagBindingParameters) {
	*out = *in
	if in.Parent != nil {
		in, out := &in.Parent, &out.Parent
		*out = new(string)
		**out = **in
	}
	if in.TagValue != nil {
		in, out := &in.TagValue, &out.TagValue
		*out = new(string)
		**out = **in
	}
	if in.TagValueRef != nil {
		in, out := &in.TagValueRef, &out.TagValueRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.TagValueSelector != nil {
		in, out := &in.TagValueSelector, &out.TagValueSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TagBindingParameters.
func (in *TagBindingParameters) DeepCopy() *TagBindingParameters {
	if in == nil {
		return nil
	}
	out := new(TagBindingParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TagBindingSpec) DeepCopyInto(out *TagBindingSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TagBindingSpec.
func (in *TagBindingSpec) DeepCopy() *TagBindingSpec {
	if in == nil {
		return nil
	}
	out := new(TagBindingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TagBindingStatus) DeepCopyInto(out *TagBindingStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TagBindingStatus.
func (in *TagBindingStatus) DeepCopy() *TagBindingStatus {
	if in == nil {
		return nil
	}
	out := new(TagBindingStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TagKey) DeepCopyInto(out *TagKey) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TagKey.
func (in *TagKey) DeepCopy() *TagKey {
	if in == nil {
		return nil
	}
	out := new(TagKey)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TagKey) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TagKeyInitParameters) DeepCopyInto(out *TagKeyInitParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Parent != nil {
		in, out := &in.Parent, &out.Parent
		*out = new(string)
		**out = **in
	}
	if in.Purpose != nil {
		in, out := &in.Purpose, &out.Purpose
		*out = new(string)
		**out = **in
	}
	if in.PurposeData != nil {
		in, out := &in.PurposeData, &out.PurposeData
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.ShortName != nil {
		in, out := &in.ShortName, &out.ShortName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TagKeyInitParameters.
func (in *TagKeyInitParameters) DeepCopy() *TagKeyInitParameters {
	if in == nil {
		return nil
	}
	out := new(TagKeyInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TagKeyList) DeepCopyInto(out *TagKeyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TagKey, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TagKeyList.
func (in *TagKeyList) DeepCopy() *TagKeyList {
	if in == nil {
		return nil
	}
	out := new(TagKeyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TagKeyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TagKeyObservation) DeepCopyInto(out *TagKeyObservation) {
	*out = *in
	if in.CreateTime != nil {
		in, out := &in.CreateTime, &out.CreateTime
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NamespacedName != nil {
		in, out := &in.NamespacedName, &out.NamespacedName
		*out = new(string)
		**out = **in
	}
	if in.Parent != nil {
		in, out := &in.Parent, &out.Parent
		*out = new(string)
		**out = **in
	}
	if in.Purpose != nil {
		in, out := &in.Purpose, &out.Purpose
		*out = new(string)
		**out = **in
	}
	if in.PurposeData != nil {
		in, out := &in.PurposeData, &out.PurposeData
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.ShortName != nil {
		in, out := &in.ShortName, &out.ShortName
		*out = new(string)
		**out = **in
	}
	if in.UpdateTime != nil {
		in, out := &in.UpdateTime, &out.UpdateTime
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TagKeyObservation.
func (in *TagKeyObservation) DeepCopy() *TagKeyObservation {
	if in == nil {
		return nil
	}
	out := new(TagKeyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TagKeyParameters) DeepCopyInto(out *TagKeyParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Parent != nil {
		in, out := &in.Parent, &out.Parent
		*out = new(string)
		**out = **in
	}
	if in.Purpose != nil {
		in, out := &in.Purpose, &out.Purpose
		*out = new(string)
		**out = **in
	}
	if in.PurposeData != nil {
		in, out := &in.PurposeData, &out.PurposeData
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.ShortName != nil {
		in, out := &in.ShortName, &out.ShortName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TagKeyParameters.
func (in *TagKeyParameters) DeepCopy() *TagKeyParameters {
	if in == nil {
		return nil
	}
	out := new(TagKeyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TagKeySpec) DeepCopyInto(out *TagKeySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TagKeySpec.
func (in *TagKeySpec) DeepCopy() *TagKeySpec {
	if in == nil {
		return nil
	}
	out := new(TagKeySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TagKeyStatus) DeepCopyInto(out *TagKeyStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TagKeyStatus.
func (in *TagKeyStatus) DeepCopy() *TagKeyStatus {
	if in == nil {
		return nil
	}
	out := new(TagKeyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TagValue) DeepCopyInto(out *TagValue) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TagValue.
func (in *TagValue) DeepCopy() *TagValue {
	if in == nil {
		return nil
	}
	out := new(TagValue)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TagValue) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TagValueInitParameters) DeepCopyInto(out *TagValueInitParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Parent != nil {
		in, out := &in.Parent, &out.Parent
		*out = new(string)
		**out = **in
	}
	if in.ParentRef != nil {
		in, out := &in.ParentRef, &out.ParentRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ParentSelector != nil {
		in, out := &in.ParentSelector, &out.ParentSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.ShortName != nil {
		in, out := &in.ShortName, &out.ShortName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TagValueInitParameters.
func (in *TagValueInitParameters) DeepCopy() *TagValueInitParameters {
	if in == nil {
		return nil
	}
	out := new(TagValueInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TagValueList) DeepCopyInto(out *TagValueList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TagValue, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TagValueList.
func (in *TagValueList) DeepCopy() *TagValueList {
	if in == nil {
		return nil
	}
	out := new(TagValueList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TagValueList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TagValueObservation) DeepCopyInto(out *TagValueObservation) {
	*out = *in
	if in.CreateTime != nil {
		in, out := &in.CreateTime, &out.CreateTime
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NamespacedName != nil {
		in, out := &in.NamespacedName, &out.NamespacedName
		*out = new(string)
		**out = **in
	}
	if in.Parent != nil {
		in, out := &in.Parent, &out.Parent
		*out = new(string)
		**out = **in
	}
	if in.ShortName != nil {
		in, out := &in.ShortName, &out.ShortName
		*out = new(string)
		**out = **in
	}
	if in.UpdateTime != nil {
		in, out := &in.UpdateTime, &out.UpdateTime
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TagValueObservation.
func (in *TagValueObservation) DeepCopy() *TagValueObservation {
	if in == nil {
		return nil
	}
	out := new(TagValueObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TagValueParameters) DeepCopyInto(out *TagValueParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Parent != nil {
		in, out := &in.Parent, &out.Parent
		*out = new(string)
		**out = **in
	}
	if in.ParentRef != nil {
		in, out := &in.ParentRef, &out.ParentRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ParentSelector != nil {
		in, out := &in.ParentSelector, &out.ParentSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.ShortName != nil {
		in, out := &in.ShortName, &out.ShortName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TagValueParameters.
func (in *TagValueParameters) DeepCopy() *TagValueParameters {
	if in == nil {
		return nil
	}
	out := new(TagValueParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TagValueSpec) DeepCopyInto(out *TagValueSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TagValueSpec.
func (in *TagValueSpec) DeepCopy() *TagValueSpec {
	if in == nil {
		return nil
	}
	out := new(TagValueSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TagValueStatus) DeepCopyInto(out *TagValueStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TagValueStatus.
func (in *TagValueStatus) DeepCopy() *TagValueStatus {
	if in == nil {
		return nil
	}
	out := new(TagValueStatus)
	in.DeepCopyInto(out)
	return out
}
