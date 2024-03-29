// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseClaim) DeepCopyInto(out *DatabaseClaim) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseClaim.
func (in *DatabaseClaim) DeepCopy() *DatabaseClaim {
	if in == nil {
		return nil
	}
	out := new(DatabaseClaim)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DatabaseClaim) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseClaimList) DeepCopyInto(out *DatabaseClaimList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DatabaseClaim, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseClaimList.
func (in *DatabaseClaimList) DeepCopy() *DatabaseClaimList {
	if in == nil {
		return nil
	}
	out := new(DatabaseClaimList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DatabaseClaimList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseClaimSpec) DeepCopyInto(out *DatabaseClaimSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseClaimSpec.
func (in *DatabaseClaimSpec) DeepCopy() *DatabaseClaimSpec {
	if in == nil {
		return nil
	}
	out := new(DatabaseClaimSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseClaimStatus) DeepCopyInto(out *DatabaseClaimStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseClaimStatus.
func (in *DatabaseClaimStatus) DeepCopy() *DatabaseClaimStatus {
	if in == nil {
		return nil
	}
	out := new(DatabaseClaimStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseInstance) DeepCopyInto(out *DatabaseInstance) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseInstance.
func (in *DatabaseInstance) DeepCopy() *DatabaseInstance {
	if in == nil {
		return nil
	}
	out := new(DatabaseInstance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DatabaseInstance) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseInstanceList) DeepCopyInto(out *DatabaseInstanceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DatabaseInstance, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseInstanceList.
func (in *DatabaseInstanceList) DeepCopy() *DatabaseInstanceList {
	if in == nil {
		return nil
	}
	out := new(DatabaseInstanceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DatabaseInstanceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseInstanceSpec) DeepCopyInto(out *DatabaseInstanceSpec) {
	*out = *in
	out.UnoperatedInstance = in.UnoperatedInstance
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseInstanceSpec.
func (in *DatabaseInstanceSpec) DeepCopy() *DatabaseInstanceSpec {
	if in == nil {
		return nil
	}
	out := new(DatabaseInstanceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseInstanceStatus) DeepCopyInto(out *DatabaseInstanceStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseInstanceStatus.
func (in *DatabaseInstanceStatus) DeepCopy() *DatabaseInstanceStatus {
	if in == nil {
		return nil
	}
	out := new(DatabaseInstanceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UnoperatedInstance) DeepCopyInto(out *UnoperatedInstance) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UnoperatedInstance.
func (in *UnoperatedInstance) DeepCopy() *UnoperatedInstance {
	if in == nil {
		return nil
	}
	out := new(UnoperatedInstance)
	in.DeepCopyInto(out)
	return out
}
