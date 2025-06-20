/*
Copyright The Kubernetes Authors.

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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	corev1 "k8s.io/api/core/v1"
)

// VolumeResourceRequirementsApplyConfiguration represents a declarative configuration of the VolumeResourceRequirements type for use
// with apply.
type VolumeResourceRequirementsApplyConfiguration struct {
	Limits   *corev1.ResourceList `json:"limits,omitempty"`
	Requests *corev1.ResourceList `json:"requests,omitempty"`
}

// VolumeResourceRequirementsApplyConfiguration constructs a declarative configuration of the VolumeResourceRequirements type for use with
// apply.
func VolumeResourceRequirements() *VolumeResourceRequirementsApplyConfiguration {
	return &VolumeResourceRequirementsApplyConfiguration{}
}
func (b VolumeResourceRequirementsApplyConfiguration) IsApplyConfiguration() {}

// WithLimits sets the Limits field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Limits field is set to the value of the last call.
func (b *VolumeResourceRequirementsApplyConfiguration) WithLimits(value corev1.ResourceList) *VolumeResourceRequirementsApplyConfiguration {
	b.Limits = &value
	return b
}

// WithRequests sets the Requests field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Requests field is set to the value of the last call.
func (b *VolumeResourceRequirementsApplyConfiguration) WithRequests(value corev1.ResourceList) *VolumeResourceRequirementsApplyConfiguration {
	b.Requests = &value
	return b
}
