// +build !ignore_autogenerated

/*
Copyright 2018 The Kubernetes Authors.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1beta1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackendConfig) DeepCopyInto(out *BackendConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackendConfig.
func (in *BackendConfig) DeepCopy() *BackendConfig {
	if in == nil {
		return nil
	}
	out := new(BackendConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BackendConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackendConfigList) DeepCopyInto(out *BackendConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BackendConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackendConfigList.
func (in *BackendConfigList) DeepCopy() *BackendConfigList {
	if in == nil {
		return nil
	}
	out := new(BackendConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BackendConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackendConfigSpec) DeepCopyInto(out *BackendConfigSpec) {
	*out = *in
	if in.Iap != nil {
		in, out := &in.Iap, &out.Iap
		if *in == nil {
			*out = nil
		} else {
			*out = new(IAPConfig)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Cdn != nil {
		in, out := &in.Cdn, &out.Cdn
		if *in == nil {
			*out = nil
		} else {
			*out = new(CDNConfig)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.SecurityPolicy != nil {
		in, out := &in.SecurityPolicy, &out.SecurityPolicy
		if *in == nil {
			*out = nil
		} else {
			*out = new(SecurityPolicyConfig)
			**out = **in
		}
	}
	if in.ConnectionDraining != nil {
		in, out := &in.ConnectionDraining, &out.ConnectionDraining
		if *in == nil {
			*out = nil
		} else {
			*out = new(ConnectionDrainingConfig)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackendConfigSpec.
func (in *BackendConfigSpec) DeepCopy() *BackendConfigSpec {
	if in == nil {
		return nil
	}
	out := new(BackendConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackendConfigStatus) DeepCopyInto(out *BackendConfigStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackendConfigStatus.
func (in *BackendConfigStatus) DeepCopy() *BackendConfigStatus {
	if in == nil {
		return nil
	}
	out := new(BackendConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CDNConfig) DeepCopyInto(out *CDNConfig) {
	*out = *in
	if in.CachePolicy != nil {
		in, out := &in.CachePolicy, &out.CachePolicy
		if *in == nil {
			*out = nil
		} else {
			*out = new(CacheKeyPolicy)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CDNConfig.
func (in *CDNConfig) DeepCopy() *CDNConfig {
	if in == nil {
		return nil
	}
	out := new(CDNConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CacheKeyPolicy) DeepCopyInto(out *CacheKeyPolicy) {
	*out = *in
	if in.QueryStringBlacklist != nil {
		in, out := &in.QueryStringBlacklist, &out.QueryStringBlacklist
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.QueryStringWhitelist != nil {
		in, out := &in.QueryStringWhitelist, &out.QueryStringWhitelist
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CacheKeyPolicy.
func (in *CacheKeyPolicy) DeepCopy() *CacheKeyPolicy {
	if in == nil {
		return nil
	}
	out := new(CacheKeyPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConnectionDrainingConfig) DeepCopyInto(out *ConnectionDrainingConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnectionDrainingConfig.
func (in *ConnectionDrainingConfig) DeepCopy() *ConnectionDrainingConfig {
	if in == nil {
		return nil
	}
	out := new(ConnectionDrainingConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAPConfig) DeepCopyInto(out *IAPConfig) {
	*out = *in
	if in.OAuthClientCredentials != nil {
		in, out := &in.OAuthClientCredentials, &out.OAuthClientCredentials
		if *in == nil {
			*out = nil
		} else {
			*out = new(OAuthClientCredentials)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAPConfig.
func (in *IAPConfig) DeepCopy() *IAPConfig {
	if in == nil {
		return nil
	}
	out := new(IAPConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OAuthClientCredentials) DeepCopyInto(out *OAuthClientCredentials) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OAuthClientCredentials.
func (in *OAuthClientCredentials) DeepCopy() *OAuthClientCredentials {
	if in == nil {
		return nil
	}
	out := new(OAuthClientCredentials)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityPolicyConfig) DeepCopyInto(out *SecurityPolicyConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityPolicyConfig.
func (in *SecurityPolicyConfig) DeepCopy() *SecurityPolicyConfig {
	if in == nil {
		return nil
	}
	out := new(SecurityPolicyConfig)
	in.DeepCopyInto(out)
	return out
}
