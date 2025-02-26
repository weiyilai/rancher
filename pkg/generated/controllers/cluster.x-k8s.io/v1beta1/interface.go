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

package v1beta1

import (
	"github.com/rancher/lasso/pkg/controller"
	"github.com/rancher/wrangler/v3/pkg/generic"
	"github.com/rancher/wrangler/v3/pkg/schemes"
	"k8s.io/apimachinery/pkg/runtime/schema"
	v1beta1 "sigs.k8s.io/cluster-api/api/v1beta1"
)

func init() {
	schemes.Register(v1beta1.AddToScheme)
}

type Interface interface {
	Cluster() ClusterController
	Machine() MachineController
	MachineDeployment() MachineDeploymentController
	MachineHealthCheck() MachineHealthCheckController
	MachineSet() MachineSetController
}

func New(controllerFactory controller.SharedControllerFactory) Interface {
	return &version{
		controllerFactory: controllerFactory,
	}
}

type version struct {
	controllerFactory controller.SharedControllerFactory
}

func (v *version) Cluster() ClusterController {
	return generic.NewController[*v1beta1.Cluster, *v1beta1.ClusterList](schema.GroupVersionKind{Group: "cluster.x-k8s.io", Version: "v1beta1", Kind: "Cluster"}, "clusters", true, v.controllerFactory)
}

func (v *version) Machine() MachineController {
	return generic.NewController[*v1beta1.Machine, *v1beta1.MachineList](schema.GroupVersionKind{Group: "cluster.x-k8s.io", Version: "v1beta1", Kind: "Machine"}, "machines", true, v.controllerFactory)
}

func (v *version) MachineDeployment() MachineDeploymentController {
	return generic.NewController[*v1beta1.MachineDeployment, *v1beta1.MachineDeploymentList](schema.GroupVersionKind{Group: "cluster.x-k8s.io", Version: "v1beta1", Kind: "MachineDeployment"}, "machinedeployments", true, v.controllerFactory)
}

func (v *version) MachineHealthCheck() MachineHealthCheckController {
	return generic.NewController[*v1beta1.MachineHealthCheck, *v1beta1.MachineHealthCheckList](schema.GroupVersionKind{Group: "cluster.x-k8s.io", Version: "v1beta1", Kind: "MachineHealthCheck"}, "machinehealthchecks", true, v.controllerFactory)
}

func (v *version) MachineSet() MachineSetController {
	return generic.NewController[*v1beta1.MachineSet, *v1beta1.MachineSetList](schema.GroupVersionKind{Group: "cluster.x-k8s.io", Version: "v1beta1", Kind: "MachineSet"}, "machinesets", true, v.controllerFactory)
}
