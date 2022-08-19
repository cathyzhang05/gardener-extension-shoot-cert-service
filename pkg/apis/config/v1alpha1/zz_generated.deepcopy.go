//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright (c) SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

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

package v1alpha1

import (
	configv1alpha1 "github.com/gardener/gardener/extensions/pkg/apis/config/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ACME) DeepCopyInto(out *ACME) {
	*out = *in
	if in.PrivateKey != nil {
		in, out := &in.PrivateKey, &out.PrivateKey
		*out = new(string)
		**out = **in
	}
	if in.PropagationTimeout != nil {
		in, out := &in.PropagationTimeout, &out.PropagationTimeout
		*out = new(v1.Duration)
		**out = **in
	}
	if in.PrecheckNameservers != nil {
		in, out := &in.PrecheckNameservers, &out.PrecheckNameservers
		*out = new(string)
		**out = **in
	}
	if in.CACertificates != nil {
		in, out := &in.CACertificates, &out.CACertificates
		*out = new(string)
		**out = **in
	}
	if in.DeactivateAuthorizations != nil {
		in, out := &in.DeactivateAuthorizations, &out.DeactivateAuthorizations
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ACME.
func (in *ACME) DeepCopy() *ACME {
	if in == nil {
		return nil
	}
	out := new(ACME)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Configuration) DeepCopyInto(out *Configuration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.RestrictIssuer != nil {
		in, out := &in.RestrictIssuer, &out.RestrictIssuer
		*out = new(bool)
		**out = **in
	}
	if in.DefaultRequestsPerDayQuota != nil {
		in, out := &in.DefaultRequestsPerDayQuota, &out.DefaultRequestsPerDayQuota
		*out = new(int32)
		**out = **in
	}
	if in.ShootIssuers != nil {
		in, out := &in.ShootIssuers, &out.ShootIssuers
		*out = new(ShootIssuers)
		**out = **in
	}
	in.ACME.DeepCopyInto(&out.ACME)
	if in.HealthCheckConfig != nil {
		in, out := &in.HealthCheckConfig, &out.HealthCheckConfig
		*out = new(configv1alpha1.HealthCheckConfig)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Configuration.
func (in *Configuration) DeepCopy() *Configuration {
	if in == nil {
		return nil
	}
	out := new(Configuration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Configuration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShootIssuers) DeepCopyInto(out *ShootIssuers) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShootIssuers.
func (in *ShootIssuers) DeepCopy() *ShootIssuers {
	if in == nil {
		return nil
	}
	out := new(ShootIssuers)
	in.DeepCopyInto(out)
	return out
}
