//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CPUs) DeepCopyInto(out *CPUs) {
	*out = *in
	if in.Min != nil {
		in, out := &in.Min, &out.Min
		*out = new(MilliCPU)
		**out = **in
	}
	if in.Max != nil {
		in, out := &in.Max, &out.Max
		*out = new(MilliCPU)
		**out = **in
	}
	if in.Use != nil {
		in, out := &in.Use, &out.Use
		*out = new(MilliCPU)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CPUs.
func (in *CPUs) DeepCopy() *CPUs {
	if in == nil {
		return nil
	}
	out := new(CPUs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Disk) DeepCopyInto(out *Disk) {
	*out = *in
	if in.ReadOnly != nil {
		in, out := &in.ReadOnly, &out.ReadOnly
		*out = new(bool)
		**out = **in
	}
	in.DiskSource.DeepCopyInto(&out.DiskSource)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Disk.
func (in *Disk) DeepCopy() *Disk {
	if in == nil {
		return nil
	}
	out := new(Disk)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiskSource) DeepCopyInto(out *DiskSource) {
	*out = *in
	if in.EmptyDisk != nil {
		in, out := &in.EmptyDisk, &out.EmptyDisk
		*out = new(EmptyDiskSource)
		(*in).DeepCopyInto(*out)
	}
	if in.ConfigMap != nil {
		in, out := &in.ConfigMap, &out.ConfigMap
		*out = new(corev1.ConfigMapVolumeSource)
		(*in).DeepCopyInto(*out)
	}
	if in.Secret != nil {
		in, out := &in.Secret, &out.Secret
		*out = new(corev1.SecretVolumeSource)
		(*in).DeepCopyInto(*out)
	}
	if in.Tmpfs != nil {
		in, out := &in.Tmpfs, &out.Tmpfs
		*out = new(TmpfsDiskSource)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiskSource.
func (in *DiskSource) DeepCopy() *DiskSource {
	if in == nil {
		return nil
	}
	out := new(DiskSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EmptyDiskSource) DeepCopyInto(out *EmptyDiskSource) {
	*out = *in
	out.Size = in.Size.DeepCopy()
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EmptyDiskSource.
func (in *EmptyDiskSource) DeepCopy() *EmptyDiskSource {
	if in == nil {
		return nil
	}
	out := new(EmptyDiskSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvVar) DeepCopyInto(out *EnvVar) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvVar.
func (in *EnvVar) DeepCopy() *EnvVar {
	if in == nil {
		return nil
	}
	out := new(EnvVar)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExtraNetwork) DeepCopyInto(out *ExtraNetwork) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExtraNetwork.
func (in *ExtraNetwork) DeepCopy() *ExtraNetwork {
	if in == nil {
		return nil
	}
	out := new(ExtraNetwork)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Guest) DeepCopyInto(out *Guest) {
	*out = *in
	if in.KernelImage != nil {
		in, out := &in.KernelImage, &out.KernelImage
		*out = new(string)
		**out = **in
	}
	if in.AppendKernelCmdline != nil {
		in, out := &in.AppendKernelCmdline, &out.AppendKernelCmdline
		*out = new(string)
		**out = **in
	}
	in.CPUs.DeepCopyInto(&out.CPUs)
	out.MemorySlotSize = in.MemorySlotSize.DeepCopy()
	in.MemorySlots.DeepCopyInto(&out.MemorySlots)
	in.RootDisk.DeepCopyInto(&out.RootDisk)
	if in.Command != nil {
		in, out := &in.Command, &out.Command
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]EnvVar, len(*in))
		copy(*out, *in)
	}
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]Port, len(*in))
		copy(*out, *in)
	}
	if in.Settings != nil {
		in, out := &in.Settings, &out.Settings
		*out = new(GuestSettings)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Guest.
func (in *Guest) DeepCopy() *Guest {
	if in == nil {
		return nil
	}
	out := new(Guest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GuestSettings) DeepCopyInto(out *GuestSettings) {
	*out = *in
	if in.Sysctl != nil {
		in, out := &in.Sysctl, &out.Sysctl
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Swap != nil {
		in, out := &in.Swap, &out.Swap
		x := (*in).DeepCopy()
		*out = &x
	}
	if in.SwapInfo != nil {
		in, out := &in.SwapInfo, &out.SwapInfo
		*out = new(SwapInfo)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GuestSettings.
func (in *GuestSettings) DeepCopy() *GuestSettings {
	if in == nil {
		return nil
	}
	out := new(GuestSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPAllocation) DeepCopyInto(out *IPAllocation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPAllocation.
func (in *IPAllocation) DeepCopy() *IPAllocation {
	if in == nil {
		return nil
	}
	out := new(IPAllocation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPPool) DeepCopyInto(out *IPPool) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPPool.
func (in *IPPool) DeepCopy() *IPPool {
	if in == nil {
		return nil
	}
	out := new(IPPool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IPPool) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPPoolList) DeepCopyInto(out *IPPoolList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IPPool, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPPoolList.
func (in *IPPoolList) DeepCopy() *IPPoolList {
	if in == nil {
		return nil
	}
	out := new(IPPoolList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IPPoolList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPPoolSpec) DeepCopyInto(out *IPPoolSpec) {
	*out = *in
	if in.Allocations != nil {
		in, out := &in.Allocations, &out.Allocations
		*out = make(map[string]IPAllocation, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPPoolSpec.
func (in *IPPoolSpec) DeepCopy() *IPPoolSpec {
	if in == nil {
		return nil
	}
	out := new(IPPoolSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemorySlots) DeepCopyInto(out *MemorySlots) {
	*out = *in
	if in.Min != nil {
		in, out := &in.Min, &out.Min
		*out = new(int32)
		**out = **in
	}
	if in.Max != nil {
		in, out := &in.Max, &out.Max
		*out = new(int32)
		**out = **in
	}
	if in.Use != nil {
		in, out := &in.Use, &out.Use
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemorySlots.
func (in *MemorySlots) DeepCopy() *MemorySlots {
	if in == nil {
		return nil
	}
	out := new(MemorySlots)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MigrationInfo) DeepCopyInto(out *MigrationInfo) {
	*out = *in
	out.Ram = in.Ram
	out.Compression = in.Compression
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MigrationInfo.
func (in *MigrationInfo) DeepCopy() *MigrationInfo {
	if in == nil {
		return nil
	}
	out := new(MigrationInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MigrationInfoCompression) DeepCopyInto(out *MigrationInfoCompression) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MigrationInfoCompression.
func (in *MigrationInfoCompression) DeepCopy() *MigrationInfoCompression {
	if in == nil {
		return nil
	}
	out := new(MigrationInfoCompression)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MigrationInfoRam) DeepCopyInto(out *MigrationInfoRam) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MigrationInfoRam.
func (in *MigrationInfoRam) DeepCopy() *MigrationInfoRam {
	if in == nil {
		return nil
	}
	out := new(MigrationInfoRam)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Port) DeepCopyInto(out *Port) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Port.
func (in *Port) DeepCopy() *Port {
	if in == nil {
		return nil
	}
	out := new(Port)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RootDisk) DeepCopyInto(out *RootDisk) {
	*out = *in
	out.Size = in.Size.DeepCopy()
	if in.Execute != nil {
		in, out := &in.Execute, &out.Execute
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RootDisk.
func (in *RootDisk) DeepCopy() *RootDisk {
	if in == nil {
		return nil
	}
	out := new(RootDisk)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SwapInfo) DeepCopyInto(out *SwapInfo) {
	*out = *in
	out.Size = in.Size.DeepCopy()
	if in.SkipSwapon != nil {
		in, out := &in.SkipSwapon, &out.SkipSwapon
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SwapInfo.
func (in *SwapInfo) DeepCopy() *SwapInfo {
	if in == nil {
		return nil
	}
	out := new(SwapInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TmpfsDiskSource) DeepCopyInto(out *TmpfsDiskSource) {
	*out = *in
	out.Size = in.Size.DeepCopy()
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TmpfsDiskSource.
func (in *TmpfsDiskSource) DeepCopy() *TmpfsDiskSource {
	if in == nil {
		return nil
	}
	out := new(TmpfsDiskSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachine) DeepCopyInto(out *VirtualMachine) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachine.
func (in *VirtualMachine) DeepCopy() *VirtualMachine {
	if in == nil {
		return nil
	}
	out := new(VirtualMachine)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualMachine) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachineList) DeepCopyInto(out *VirtualMachineList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VirtualMachine, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachineList.
func (in *VirtualMachineList) DeepCopy() *VirtualMachineList {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualMachineList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachineMigration) DeepCopyInto(out *VirtualMachineMigration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachineMigration.
func (in *VirtualMachineMigration) DeepCopy() *VirtualMachineMigration {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineMigration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualMachineMigration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachineMigrationList) DeepCopyInto(out *VirtualMachineMigrationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VirtualMachineMigration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachineMigrationList.
func (in *VirtualMachineMigrationList) DeepCopy() *VirtualMachineMigrationList {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineMigrationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualMachineMigrationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachineMigrationSpec) DeepCopyInto(out *VirtualMachineMigrationSpec) {
	*out = *in
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.NodeAffinity != nil {
		in, out := &in.NodeAffinity, &out.NodeAffinity
		*out = new(corev1.NodeAffinity)
		(*in).DeepCopyInto(*out)
	}
	out.MaxBandwidth = in.MaxBandwidth.DeepCopy()
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachineMigrationSpec.
func (in *VirtualMachineMigrationSpec) DeepCopy() *VirtualMachineMigrationSpec {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineMigrationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachineMigrationStatus) DeepCopyInto(out *VirtualMachineMigrationStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.Info = in.Info
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachineMigrationStatus.
func (in *VirtualMachineMigrationStatus) DeepCopy() *VirtualMachineMigrationStatus {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineMigrationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachineResources) DeepCopyInto(out *VirtualMachineResources) {
	*out = *in
	in.CPUs.DeepCopyInto(&out.CPUs)
	in.MemorySlots.DeepCopyInto(&out.MemorySlots)
	out.MemorySlotSize = in.MemorySlotSize.DeepCopy()
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachineResources.
func (in *VirtualMachineResources) DeepCopy() *VirtualMachineResources {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineResources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachineSpec) DeepCopyInto(out *VirtualMachineSpec) {
	*out = *in
	if in.TerminationGracePeriodSeconds != nil {
		in, out := &in.TerminationGracePeriodSeconds, &out.TerminationGracePeriodSeconds
		*out = new(int64)
		**out = **in
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(corev1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]corev1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.PodResources.DeepCopyInto(&out.PodResources)
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]corev1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	in.Guest.DeepCopyInto(&out.Guest)
	if in.ExtraInitContainers != nil {
		in, out := &in.ExtraInitContainers, &out.ExtraInitContainers
		*out = make([]corev1.Container, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Disks != nil {
		in, out := &in.Disks, &out.Disks
		*out = make([]Disk, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ExtraNetwork != nil {
		in, out := &in.ExtraNetwork, &out.ExtraNetwork
		*out = new(ExtraNetwork)
		**out = **in
	}
	if in.ServiceLinks != nil {
		in, out := &in.ServiceLinks, &out.ServiceLinks
		*out = new(bool)
		**out = **in
	}
	if in.EnableAcceleration != nil {
		in, out := &in.EnableAcceleration, &out.EnableAcceleration
		*out = new(bool)
		**out = **in
	}
	if in.RunnerImage != nil {
		in, out := &in.RunnerImage, &out.RunnerImage
		*out = new(string)
		**out = **in
	}
	if in.EnableSSH != nil {
		in, out := &in.EnableSSH, &out.EnableSSH
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachineSpec.
func (in *VirtualMachineSpec) DeepCopy() *VirtualMachineSpec {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachineStatus) DeepCopyInto(out *VirtualMachineStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CPUs != nil {
		in, out := &in.CPUs, &out.CPUs
		*out = new(MilliCPU)
		**out = **in
	}
	if in.MemorySize != nil {
		in, out := &in.MemorySize, &out.MemorySize
		x := (*in).DeepCopy()
		*out = &x
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachineStatus.
func (in *VirtualMachineStatus) DeepCopy() *VirtualMachineStatus {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachineUsage) DeepCopyInto(out *VirtualMachineUsage) {
	*out = *in
	if in.CPU != nil {
		in, out := &in.CPU, &out.CPU
		x := (*in).DeepCopy()
		*out = &x
	}
	if in.Memory != nil {
		in, out := &in.Memory, &out.Memory
		x := (*in).DeepCopy()
		*out = &x
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachineUsage.
func (in *VirtualMachineUsage) DeepCopy() *VirtualMachineUsage {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineUsage)
	in.DeepCopyInto(out)
	return out
}
