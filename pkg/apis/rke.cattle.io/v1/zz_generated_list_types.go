/*
Copyright 2025 Rancher Labs, Inc.

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

// Code generated by main. DO NOT EDIT.

// +k8s:deepcopy-gen=package
// +groupName=rke.cattle.io
package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CustomMachineList is a list of CustomMachine resources
type CustomMachineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []CustomMachine `json:"items"`
}

func NewCustomMachine(namespace, name string, obj CustomMachine) *CustomMachine {
	obj.APIVersion, obj.Kind = SchemeGroupVersion.WithKind("CustomMachine").ToAPIVersionAndKind()
	obj.Name = name
	obj.Namespace = namespace
	return &obj
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ETCDSnapshotList is a list of ETCDSnapshot resources
type ETCDSnapshotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []ETCDSnapshot `json:"items"`
}

func NewETCDSnapshot(namespace, name string, obj ETCDSnapshot) *ETCDSnapshot {
	obj.APIVersion, obj.Kind = SchemeGroupVersion.WithKind("ETCDSnapshot").ToAPIVersionAndKind()
	obj.Name = name
	obj.Namespace = namespace
	return &obj
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RKEBootstrapList is a list of RKEBootstrap resources
type RKEBootstrapList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []RKEBootstrap `json:"items"`
}

func NewRKEBootstrap(namespace, name string, obj RKEBootstrap) *RKEBootstrap {
	obj.APIVersion, obj.Kind = SchemeGroupVersion.WithKind("RKEBootstrap").ToAPIVersionAndKind()
	obj.Name = name
	obj.Namespace = namespace
	return &obj
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RKEBootstrapTemplateList is a list of RKEBootstrapTemplate resources
type RKEBootstrapTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []RKEBootstrapTemplate `json:"items"`
}

func NewRKEBootstrapTemplate(namespace, name string, obj RKEBootstrapTemplate) *RKEBootstrapTemplate {
	obj.APIVersion, obj.Kind = SchemeGroupVersion.WithKind("RKEBootstrapTemplate").ToAPIVersionAndKind()
	obj.Name = name
	obj.Namespace = namespace
	return &obj
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RKEClusterList is a list of RKECluster resources
type RKEClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []RKECluster `json:"items"`
}

func NewRKECluster(namespace, name string, obj RKECluster) *RKECluster {
	obj.APIVersion, obj.Kind = SchemeGroupVersion.WithKind("RKECluster").ToAPIVersionAndKind()
	obj.Name = name
	obj.Namespace = namespace
	return &obj
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RKEControlPlaneList is a list of RKEControlPlane resources
type RKEControlPlaneList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []RKEControlPlane `json:"items"`
}

func NewRKEControlPlane(namespace, name string, obj RKEControlPlane) *RKEControlPlane {
	obj.APIVersion, obj.Kind = SchemeGroupVersion.WithKind("RKEControlPlane").ToAPIVersionAndKind()
	obj.Name = name
	obj.Namespace = namespace
	return &obj
}
