// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutomationDef) DeepCopyInto(out *AutomationDef) {
	*out = *in
	if in.ExtraVars != nil {
		in, out := &in.ExtraVars, &out.ExtraVars
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutomationDef.
func (in *AutomationDef) DeepCopy() *AutomationDef {
	if in == nil {
		return nil
	}
	out := new(AutomationDef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyAutomation) DeepCopyInto(out *PolicyAutomation) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyAutomation.
func (in *PolicyAutomation) DeepCopy() *PolicyAutomation {
	if in == nil {
		return nil
	}
	out := new(PolicyAutomation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PolicyAutomation) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyAutomationList) DeepCopyInto(out *PolicyAutomationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PolicyAutomation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyAutomationList.
func (in *PolicyAutomationList) DeepCopy() *PolicyAutomationList {
	if in == nil {
		return nil
	}
	out := new(PolicyAutomationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PolicyAutomationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyAutomationSpec) DeepCopyInto(out *PolicyAutomationSpec) {
	*out = *in
	in.Automation.DeepCopyInto(&out.Automation)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyAutomationSpec.
func (in *PolicyAutomationSpec) DeepCopy() *PolicyAutomationSpec {
	if in == nil {
		return nil
	}
	out := new(PolicyAutomationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyAutomationStatus) DeepCopyInto(out *PolicyAutomationStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyAutomationStatus.
func (in *PolicyAutomationStatus) DeepCopy() *PolicyAutomationStatus {
	if in == nil {
		return nil
	}
	out := new(PolicyAutomationStatus)
	in.DeepCopyInto(out)
	return out
}
