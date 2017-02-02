// +build !ignore_autogenerated

/*
Copyright 2017 The Kubernetes Authors.

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

// This file was autogenerated by defaulter-gen. Do not edit it manually!

package v1

import (
	runtime "k8s.io/kubernetes/pkg/runtime"
)

// RegisterDefaults adds defaulters functions to the given scheme.
// Public to allow building arbitrary schemes.
// All generated defaulters are covering - they call all nested defaulters.
func RegisterDefaults(scheme *runtime.Scheme) error {
	scheme.AddTypeDefaultingFunc(&ConfigMap{}, func(obj interface{}) { SetObjectDefaults_ConfigMap(obj.(*ConfigMap)) })
	scheme.AddTypeDefaultingFunc(&ConfigMapList{}, func(obj interface{}) { SetObjectDefaults_ConfigMapList(obj.(*ConfigMapList)) })
	scheme.AddTypeDefaultingFunc(&Endpoints{}, func(obj interface{}) { SetObjectDefaults_Endpoints(obj.(*Endpoints)) })
	scheme.AddTypeDefaultingFunc(&EndpointsList{}, func(obj interface{}) { SetObjectDefaults_EndpointsList(obj.(*EndpointsList)) })
	scheme.AddTypeDefaultingFunc(&LimitRange{}, func(obj interface{}) { SetObjectDefaults_LimitRange(obj.(*LimitRange)) })
	scheme.AddTypeDefaultingFunc(&LimitRangeList{}, func(obj interface{}) { SetObjectDefaults_LimitRangeList(obj.(*LimitRangeList)) })
	scheme.AddTypeDefaultingFunc(&Namespace{}, func(obj interface{}) { SetObjectDefaults_Namespace(obj.(*Namespace)) })
	scheme.AddTypeDefaultingFunc(&NamespaceList{}, func(obj interface{}) { SetObjectDefaults_NamespaceList(obj.(*NamespaceList)) })
	scheme.AddTypeDefaultingFunc(&Node{}, func(obj interface{}) { SetObjectDefaults_Node(obj.(*Node)) })
	scheme.AddTypeDefaultingFunc(&NodeList{}, func(obj interface{}) { SetObjectDefaults_NodeList(obj.(*NodeList)) })
	scheme.AddTypeDefaultingFunc(&PersistentVolume{}, func(obj interface{}) { SetObjectDefaults_PersistentVolume(obj.(*PersistentVolume)) })
	scheme.AddTypeDefaultingFunc(&PersistentVolumeClaim{}, func(obj interface{}) { SetObjectDefaults_PersistentVolumeClaim(obj.(*PersistentVolumeClaim)) })
	scheme.AddTypeDefaultingFunc(&PersistentVolumeClaimList{}, func(obj interface{}) { SetObjectDefaults_PersistentVolumeClaimList(obj.(*PersistentVolumeClaimList)) })
	scheme.AddTypeDefaultingFunc(&PersistentVolumeList{}, func(obj interface{}) { SetObjectDefaults_PersistentVolumeList(obj.(*PersistentVolumeList)) })
	scheme.AddTypeDefaultingFunc(&Pod{}, func(obj interface{}) { SetObjectDefaults_Pod(obj.(*Pod)) })
	scheme.AddTypeDefaultingFunc(&PodAttachOptions{}, func(obj interface{}) { SetObjectDefaults_PodAttachOptions(obj.(*PodAttachOptions)) })
	scheme.AddTypeDefaultingFunc(&PodExecOptions{}, func(obj interface{}) { SetObjectDefaults_PodExecOptions(obj.(*PodExecOptions)) })
	scheme.AddTypeDefaultingFunc(&PodList{}, func(obj interface{}) { SetObjectDefaults_PodList(obj.(*PodList)) })
	scheme.AddTypeDefaultingFunc(&PodTemplate{}, func(obj interface{}) { SetObjectDefaults_PodTemplate(obj.(*PodTemplate)) })
	scheme.AddTypeDefaultingFunc(&PodTemplateList{}, func(obj interface{}) { SetObjectDefaults_PodTemplateList(obj.(*PodTemplateList)) })
	scheme.AddTypeDefaultingFunc(&ReplicationController{}, func(obj interface{}) { SetObjectDefaults_ReplicationController(obj.(*ReplicationController)) })
	scheme.AddTypeDefaultingFunc(&ReplicationControllerList{}, func(obj interface{}) { SetObjectDefaults_ReplicationControllerList(obj.(*ReplicationControllerList)) })
	scheme.AddTypeDefaultingFunc(&ResourceQuota{}, func(obj interface{}) { SetObjectDefaults_ResourceQuota(obj.(*ResourceQuota)) })
	scheme.AddTypeDefaultingFunc(&ResourceQuotaList{}, func(obj interface{}) { SetObjectDefaults_ResourceQuotaList(obj.(*ResourceQuotaList)) })
	scheme.AddTypeDefaultingFunc(&Secret{}, func(obj interface{}) { SetObjectDefaults_Secret(obj.(*Secret)) })
	scheme.AddTypeDefaultingFunc(&SecretList{}, func(obj interface{}) { SetObjectDefaults_SecretList(obj.(*SecretList)) })
	scheme.AddTypeDefaultingFunc(&Service{}, func(obj interface{}) { SetObjectDefaults_Service(obj.(*Service)) })
	scheme.AddTypeDefaultingFunc(&ServiceList{}, func(obj interface{}) { SetObjectDefaults_ServiceList(obj.(*ServiceList)) })
	return nil
}

func SetObjectDefaults_ConfigMap(in *ConfigMap) {
	SetDefaults_ConfigMap(in)
}

func SetObjectDefaults_ConfigMapList(in *ConfigMapList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_ConfigMap(a)
	}
}

func SetObjectDefaults_Endpoints(in *Endpoints) {
	SetDefaults_Endpoints(in)
}

func SetObjectDefaults_EndpointsList(in *EndpointsList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_Endpoints(a)
	}
}

func SetObjectDefaults_LimitRange(in *LimitRange) {
	for i := range in.Spec.Limits {
		a := &in.Spec.Limits[i]
		SetDefaults_LimitRangeItem(a)
		SetDefaults_ResourceList(&a.Max)
		SetDefaults_ResourceList(&a.Min)
		SetDefaults_ResourceList(&a.Default)
		SetDefaults_ResourceList(&a.DefaultRequest)
		SetDefaults_ResourceList(&a.MaxLimitRequestRatio)
	}
}

func SetObjectDefaults_LimitRangeList(in *LimitRangeList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_LimitRange(a)
	}
}

func SetObjectDefaults_Namespace(in *Namespace) {
	SetDefaults_NamespaceStatus(&in.Status)
}

func SetObjectDefaults_NamespaceList(in *NamespaceList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_Namespace(a)
	}
}

func SetObjectDefaults_Node(in *Node) {
	SetDefaults_Node(in)
	SetDefaults_NodeStatus(&in.Status)
	SetDefaults_ResourceList(&in.Status.Capacity)
	SetDefaults_ResourceList(&in.Status.Allocatable)
}

func SetObjectDefaults_NodeList(in *NodeList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_Node(a)
	}
}

func SetObjectDefaults_PersistentVolume(in *PersistentVolume) {
	SetDefaults_PersistentVolume(in)
	SetDefaults_ResourceList(&in.Spec.Capacity)
	if in.Spec.PersistentVolumeSource.RBD != nil {
		SetDefaults_RBDVolumeSource(in.Spec.PersistentVolumeSource.RBD)
	}
	if in.Spec.PersistentVolumeSource.ISCSI != nil {
		SetDefaults_ISCSIVolumeSource(in.Spec.PersistentVolumeSource.ISCSI)
	}
	if in.Spec.PersistentVolumeSource.AzureDisk != nil {
		SetDefaults_AzureDiskVolumeSource(in.Spec.PersistentVolumeSource.AzureDisk)
	}
}

func SetObjectDefaults_PersistentVolumeClaim(in *PersistentVolumeClaim) {
	SetDefaults_PersistentVolumeClaim(in)
	SetDefaults_ResourceList(&in.Spec.Resources.Limits)
	SetDefaults_ResourceList(&in.Spec.Resources.Requests)
	SetDefaults_ResourceList(&in.Status.Capacity)
}

func SetObjectDefaults_PersistentVolumeClaimList(in *PersistentVolumeClaimList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_PersistentVolumeClaim(a)
	}
}

func SetObjectDefaults_PersistentVolumeList(in *PersistentVolumeList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_PersistentVolume(a)
	}
}

func SetObjectDefaults_Pod(in *Pod) {
	SetDefaults_Pod(in)
	SetDefaults_PodSpec(&in.Spec)
	for i := range in.Spec.Volumes {
		a := &in.Spec.Volumes[i]
		SetDefaults_Volume(a)
		if a.VolumeSource.Secret != nil {
			SetDefaults_SecretVolumeSource(a.VolumeSource.Secret)
		}
		if a.VolumeSource.ISCSI != nil {
			SetDefaults_ISCSIVolumeSource(a.VolumeSource.ISCSI)
		}
		if a.VolumeSource.RBD != nil {
			SetDefaults_RBDVolumeSource(a.VolumeSource.RBD)
		}
		if a.VolumeSource.DownwardAPI != nil {
			SetDefaults_DownwardAPIVolumeSource(a.VolumeSource.DownwardAPI)
			for j := range a.VolumeSource.DownwardAPI.Items {
				b := &a.VolumeSource.DownwardAPI.Items[j]
				if b.FieldRef != nil {
					SetDefaults_ObjectFieldSelector(b.FieldRef)
				}
			}
		}
		if a.VolumeSource.ConfigMap != nil {
			SetDefaults_ConfigMapVolumeSource(a.VolumeSource.ConfigMap)
		}
		if a.VolumeSource.AzureDisk != nil {
			SetDefaults_AzureDiskVolumeSource(a.VolumeSource.AzureDisk)
		}
		if a.VolumeSource.Metadata != nil {
			SetDefaults_DeprecatedDownwardAPIVolumeSource(a.VolumeSource.Metadata)
			for j := range a.VolumeSource.Metadata.Items {
				b := &a.VolumeSource.Metadata.Items[j]
				if b.FieldRef != nil {
					SetDefaults_ObjectFieldSelector(b.FieldRef)
				}
			}
		}
	}
	for i := range in.Spec.InitContainers {
		a := &in.Spec.InitContainers[i]
		SetDefaults_Container(a)
		for j := range a.Ports {
			b := &a.Ports[j]
			SetDefaults_ContainerPort(b)
		}
		for j := range a.Env {
			b := &a.Env[j]
			if b.ValueFrom != nil {
				if b.ValueFrom.FieldRef != nil {
					SetDefaults_ObjectFieldSelector(b.ValueFrom.FieldRef)
				}
			}
		}
		SetDefaults_ResourceList(&a.Resources.Limits)
		SetDefaults_ResourceList(&a.Resources.Requests)
		if a.LivenessProbe != nil {
			SetDefaults_Probe(a.LivenessProbe)
			if a.LivenessProbe.Handler.HTTPGet != nil {
				SetDefaults_HTTPGetAction(a.LivenessProbe.Handler.HTTPGet)
			}
		}
		if a.ReadinessProbe != nil {
			SetDefaults_Probe(a.ReadinessProbe)
			if a.ReadinessProbe.Handler.HTTPGet != nil {
				SetDefaults_HTTPGetAction(a.ReadinessProbe.Handler.HTTPGet)
			}
		}
		if a.Lifecycle != nil {
			if a.Lifecycle.PostStart != nil {
				if a.Lifecycle.PostStart.HTTPGet != nil {
					SetDefaults_HTTPGetAction(a.Lifecycle.PostStart.HTTPGet)
				}
			}
			if a.Lifecycle.PreStop != nil {
				if a.Lifecycle.PreStop.HTTPGet != nil {
					SetDefaults_HTTPGetAction(a.Lifecycle.PreStop.HTTPGet)
				}
			}
		}
	}
	for i := range in.Spec.Containers {
		a := &in.Spec.Containers[i]
		SetDefaults_Container(a)
		for j := range a.Ports {
			b := &a.Ports[j]
			SetDefaults_ContainerPort(b)
		}
		for j := range a.Env {
			b := &a.Env[j]
			if b.ValueFrom != nil {
				if b.ValueFrom.FieldRef != nil {
					SetDefaults_ObjectFieldSelector(b.ValueFrom.FieldRef)
				}
			}
		}
		SetDefaults_ResourceList(&a.Resources.Limits)
		SetDefaults_ResourceList(&a.Resources.Requests)
		if a.LivenessProbe != nil {
			SetDefaults_Probe(a.LivenessProbe)
			if a.LivenessProbe.Handler.HTTPGet != nil {
				SetDefaults_HTTPGetAction(a.LivenessProbe.Handler.HTTPGet)
			}
		}
		if a.ReadinessProbe != nil {
			SetDefaults_Probe(a.ReadinessProbe)
			if a.ReadinessProbe.Handler.HTTPGet != nil {
				SetDefaults_HTTPGetAction(a.ReadinessProbe.Handler.HTTPGet)
			}
		}
		if a.Lifecycle != nil {
			if a.Lifecycle.PostStart != nil {
				if a.Lifecycle.PostStart.HTTPGet != nil {
					SetDefaults_HTTPGetAction(a.Lifecycle.PostStart.HTTPGet)
				}
			}
			if a.Lifecycle.PreStop != nil {
				if a.Lifecycle.PreStop.HTTPGet != nil {
					SetDefaults_HTTPGetAction(a.Lifecycle.PreStop.HTTPGet)
				}
			}
		}
	}
}

func SetObjectDefaults_PodAttachOptions(in *PodAttachOptions) {
	SetDefaults_PodAttachOptions(in)
}

func SetObjectDefaults_PodExecOptions(in *PodExecOptions) {
	SetDefaults_PodExecOptions(in)
}

func SetObjectDefaults_PodList(in *PodList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_Pod(a)
	}
}

func SetObjectDefaults_PodTemplate(in *PodTemplate) {
	SetDefaults_PodSpec(&in.Template.Spec)
	for i := range in.Template.Spec.Volumes {
		a := &in.Template.Spec.Volumes[i]
		SetDefaults_Volume(a)
		if a.VolumeSource.Secret != nil {
			SetDefaults_SecretVolumeSource(a.VolumeSource.Secret)
		}
		if a.VolumeSource.ISCSI != nil {
			SetDefaults_ISCSIVolumeSource(a.VolumeSource.ISCSI)
		}
		if a.VolumeSource.RBD != nil {
			SetDefaults_RBDVolumeSource(a.VolumeSource.RBD)
		}
		if a.VolumeSource.DownwardAPI != nil {
			SetDefaults_DownwardAPIVolumeSource(a.VolumeSource.DownwardAPI)
			for j := range a.VolumeSource.DownwardAPI.Items {
				b := &a.VolumeSource.DownwardAPI.Items[j]
				if b.FieldRef != nil {
					SetDefaults_ObjectFieldSelector(b.FieldRef)
				}
			}
		}
		if a.VolumeSource.ConfigMap != nil {
			SetDefaults_ConfigMapVolumeSource(a.VolumeSource.ConfigMap)
		}
		if a.VolumeSource.AzureDisk != nil {
			SetDefaults_AzureDiskVolumeSource(a.VolumeSource.AzureDisk)
		}
		if a.VolumeSource.Metadata != nil {
			SetDefaults_DeprecatedDownwardAPIVolumeSource(a.VolumeSource.Metadata)
			for j := range a.VolumeSource.Metadata.Items {
				b := &a.VolumeSource.Metadata.Items[j]
				if b.FieldRef != nil {
					SetDefaults_ObjectFieldSelector(b.FieldRef)
				}
			}
		}
	}
	for i := range in.Template.Spec.InitContainers {
		a := &in.Template.Spec.InitContainers[i]
		SetDefaults_Container(a)
		for j := range a.Ports {
			b := &a.Ports[j]
			SetDefaults_ContainerPort(b)
		}
		for j := range a.Env {
			b := &a.Env[j]
			if b.ValueFrom != nil {
				if b.ValueFrom.FieldRef != nil {
					SetDefaults_ObjectFieldSelector(b.ValueFrom.FieldRef)
				}
			}
		}
		SetDefaults_ResourceList(&a.Resources.Limits)
		SetDefaults_ResourceList(&a.Resources.Requests)
		if a.LivenessProbe != nil {
			SetDefaults_Probe(a.LivenessProbe)
			if a.LivenessProbe.Handler.HTTPGet != nil {
				SetDefaults_HTTPGetAction(a.LivenessProbe.Handler.HTTPGet)
			}
		}
		if a.ReadinessProbe != nil {
			SetDefaults_Probe(a.ReadinessProbe)
			if a.ReadinessProbe.Handler.HTTPGet != nil {
				SetDefaults_HTTPGetAction(a.ReadinessProbe.Handler.HTTPGet)
			}
		}
		if a.Lifecycle != nil {
			if a.Lifecycle.PostStart != nil {
				if a.Lifecycle.PostStart.HTTPGet != nil {
					SetDefaults_HTTPGetAction(a.Lifecycle.PostStart.HTTPGet)
				}
			}
			if a.Lifecycle.PreStop != nil {
				if a.Lifecycle.PreStop.HTTPGet != nil {
					SetDefaults_HTTPGetAction(a.Lifecycle.PreStop.HTTPGet)
				}
			}
		}
	}
	for i := range in.Template.Spec.Containers {
		a := &in.Template.Spec.Containers[i]
		SetDefaults_Container(a)
		for j := range a.Ports {
			b := &a.Ports[j]
			SetDefaults_ContainerPort(b)
		}
		for j := range a.Env {
			b := &a.Env[j]
			if b.ValueFrom != nil {
				if b.ValueFrom.FieldRef != nil {
					SetDefaults_ObjectFieldSelector(b.ValueFrom.FieldRef)
				}
			}
		}
		SetDefaults_ResourceList(&a.Resources.Limits)
		SetDefaults_ResourceList(&a.Resources.Requests)
		if a.LivenessProbe != nil {
			SetDefaults_Probe(a.LivenessProbe)
			if a.LivenessProbe.Handler.HTTPGet != nil {
				SetDefaults_HTTPGetAction(a.LivenessProbe.Handler.HTTPGet)
			}
		}
		if a.ReadinessProbe != nil {
			SetDefaults_Probe(a.ReadinessProbe)
			if a.ReadinessProbe.Handler.HTTPGet != nil {
				SetDefaults_HTTPGetAction(a.ReadinessProbe.Handler.HTTPGet)
			}
		}
		if a.Lifecycle != nil {
			if a.Lifecycle.PostStart != nil {
				if a.Lifecycle.PostStart.HTTPGet != nil {
					SetDefaults_HTTPGetAction(a.Lifecycle.PostStart.HTTPGet)
				}
			}
			if a.Lifecycle.PreStop != nil {
				if a.Lifecycle.PreStop.HTTPGet != nil {
					SetDefaults_HTTPGetAction(a.Lifecycle.PreStop.HTTPGet)
				}
			}
		}
	}
}

func SetObjectDefaults_PodTemplateList(in *PodTemplateList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_PodTemplate(a)
	}
}

func SetObjectDefaults_ReplicationController(in *ReplicationController) {
	SetDefaults_ReplicationController(in)
	if in.Spec.Template != nil {
		SetDefaults_PodSpec(&in.Spec.Template.Spec)
		for i := range in.Spec.Template.Spec.Volumes {
			a := &in.Spec.Template.Spec.Volumes[i]
			SetDefaults_Volume(a)
			if a.VolumeSource.Secret != nil {
				SetDefaults_SecretVolumeSource(a.VolumeSource.Secret)
			}
			if a.VolumeSource.ISCSI != nil {
				SetDefaults_ISCSIVolumeSource(a.VolumeSource.ISCSI)
			}
			if a.VolumeSource.RBD != nil {
				SetDefaults_RBDVolumeSource(a.VolumeSource.RBD)
			}
			if a.VolumeSource.DownwardAPI != nil {
				SetDefaults_DownwardAPIVolumeSource(a.VolumeSource.DownwardAPI)
				for j := range a.VolumeSource.DownwardAPI.Items {
					b := &a.VolumeSource.DownwardAPI.Items[j]
					if b.FieldRef != nil {
						SetDefaults_ObjectFieldSelector(b.FieldRef)
					}
				}
			}
			if a.VolumeSource.ConfigMap != nil {
				SetDefaults_ConfigMapVolumeSource(a.VolumeSource.ConfigMap)
			}
			if a.VolumeSource.AzureDisk != nil {
				SetDefaults_AzureDiskVolumeSource(a.VolumeSource.AzureDisk)
			}
			if a.VolumeSource.Metadata != nil {
				SetDefaults_DeprecatedDownwardAPIVolumeSource(a.VolumeSource.Metadata)
				for j := range a.VolumeSource.Metadata.Items {
					b := &a.VolumeSource.Metadata.Items[j]
					if b.FieldRef != nil {
						SetDefaults_ObjectFieldSelector(b.FieldRef)
					}
				}
			}
		}
		for i := range in.Spec.Template.Spec.InitContainers {
			a := &in.Spec.Template.Spec.InitContainers[i]
			SetDefaults_Container(a)
			for j := range a.Ports {
				b := &a.Ports[j]
				SetDefaults_ContainerPort(b)
			}
			for j := range a.Env {
				b := &a.Env[j]
				if b.ValueFrom != nil {
					if b.ValueFrom.FieldRef != nil {
						SetDefaults_ObjectFieldSelector(b.ValueFrom.FieldRef)
					}
				}
			}
			SetDefaults_ResourceList(&a.Resources.Limits)
			SetDefaults_ResourceList(&a.Resources.Requests)
			if a.LivenessProbe != nil {
				SetDefaults_Probe(a.LivenessProbe)
				if a.LivenessProbe.Handler.HTTPGet != nil {
					SetDefaults_HTTPGetAction(a.LivenessProbe.Handler.HTTPGet)
				}
			}
			if a.ReadinessProbe != nil {
				SetDefaults_Probe(a.ReadinessProbe)
				if a.ReadinessProbe.Handler.HTTPGet != nil {
					SetDefaults_HTTPGetAction(a.ReadinessProbe.Handler.HTTPGet)
				}
			}
			if a.Lifecycle != nil {
				if a.Lifecycle.PostStart != nil {
					if a.Lifecycle.PostStart.HTTPGet != nil {
						SetDefaults_HTTPGetAction(a.Lifecycle.PostStart.HTTPGet)
					}
				}
				if a.Lifecycle.PreStop != nil {
					if a.Lifecycle.PreStop.HTTPGet != nil {
						SetDefaults_HTTPGetAction(a.Lifecycle.PreStop.HTTPGet)
					}
				}
			}
		}
		for i := range in.Spec.Template.Spec.Containers {
			a := &in.Spec.Template.Spec.Containers[i]
			SetDefaults_Container(a)
			for j := range a.Ports {
				b := &a.Ports[j]
				SetDefaults_ContainerPort(b)
			}
			for j := range a.Env {
				b := &a.Env[j]
				if b.ValueFrom != nil {
					if b.ValueFrom.FieldRef != nil {
						SetDefaults_ObjectFieldSelector(b.ValueFrom.FieldRef)
					}
				}
			}
			SetDefaults_ResourceList(&a.Resources.Limits)
			SetDefaults_ResourceList(&a.Resources.Requests)
			if a.LivenessProbe != nil {
				SetDefaults_Probe(a.LivenessProbe)
				if a.LivenessProbe.Handler.HTTPGet != nil {
					SetDefaults_HTTPGetAction(a.LivenessProbe.Handler.HTTPGet)
				}
			}
			if a.ReadinessProbe != nil {
				SetDefaults_Probe(a.ReadinessProbe)
				if a.ReadinessProbe.Handler.HTTPGet != nil {
					SetDefaults_HTTPGetAction(a.ReadinessProbe.Handler.HTTPGet)
				}
			}
			if a.Lifecycle != nil {
				if a.Lifecycle.PostStart != nil {
					if a.Lifecycle.PostStart.HTTPGet != nil {
						SetDefaults_HTTPGetAction(a.Lifecycle.PostStart.HTTPGet)
					}
				}
				if a.Lifecycle.PreStop != nil {
					if a.Lifecycle.PreStop.HTTPGet != nil {
						SetDefaults_HTTPGetAction(a.Lifecycle.PreStop.HTTPGet)
					}
				}
			}
		}
	}
}

func SetObjectDefaults_ReplicationControllerList(in *ReplicationControllerList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_ReplicationController(a)
	}
}

func SetObjectDefaults_ResourceQuota(in *ResourceQuota) {
	SetDefaults_ResourceList(&in.Spec.Hard)
	SetDefaults_ResourceList(&in.Status.Hard)
	SetDefaults_ResourceList(&in.Status.Used)
}

func SetObjectDefaults_ResourceQuotaList(in *ResourceQuotaList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_ResourceQuota(a)
	}
}

func SetObjectDefaults_Secret(in *Secret) {
	SetDefaults_Secret(in)
}

func SetObjectDefaults_SecretList(in *SecretList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_Secret(a)
	}
}

func SetObjectDefaults_Service(in *Service) {
	SetDefaults_ServiceSpec(&in.Spec)
}

func SetObjectDefaults_ServiceList(in *ServiceList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_Service(a)
	}
}