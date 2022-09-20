//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021-2022 Red Hat, Inc.

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

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Application) DeepCopyInto(out *Application) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Application.
func (in *Application) DeepCopy() *Application {
	if in == nil {
		return nil
	}
	out := new(Application)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Application) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationGitRepository) DeepCopyInto(out *ApplicationGitRepository) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationGitRepository.
func (in *ApplicationGitRepository) DeepCopy() *ApplicationGitRepository {
	if in == nil {
		return nil
	}
	out := new(ApplicationGitRepository)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationList) DeepCopyInto(out *ApplicationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Application, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationList.
func (in *ApplicationList) DeepCopy() *ApplicationList {
	if in == nil {
		return nil
	}
	out := new(ApplicationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApplicationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationPromotionRun) DeepCopyInto(out *ApplicationPromotionRun) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationPromotionRun.
func (in *ApplicationPromotionRun) DeepCopy() *ApplicationPromotionRun {
	if in == nil {
		return nil
	}
	out := new(ApplicationPromotionRun)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApplicationPromotionRun) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationPromotionRunList) DeepCopyInto(out *ApplicationPromotionRunList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ApplicationPromotionRun, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationPromotionRunList.
func (in *ApplicationPromotionRunList) DeepCopy() *ApplicationPromotionRunList {
	if in == nil {
		return nil
	}
	out := new(ApplicationPromotionRunList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApplicationPromotionRunList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationPromotionRunSpec) DeepCopyInto(out *ApplicationPromotionRunSpec) {
	*out = *in
	out.ManualPromotion = in.ManualPromotion
	out.AutomatedPromotion = in.AutomatedPromotion
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationPromotionRunSpec.
func (in *ApplicationPromotionRunSpec) DeepCopy() *ApplicationPromotionRunSpec {
	if in == nil {
		return nil
	}
	out := new(ApplicationPromotionRunSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationPromotionRunStatus) DeepCopyInto(out *ApplicationPromotionRunStatus) {
	*out = *in
	if in.EnvironmentStatus != nil {
		in, out := &in.EnvironmentStatus, &out.EnvironmentStatus
		*out = make([]PromotionRunEnvironmentStatus, len(*in))
		copy(*out, *in)
	}
	if in.ActiveBindings != nil {
		in, out := &in.ActiveBindings, &out.ActiveBindings
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationPromotionRunStatus.
func (in *ApplicationPromotionRunStatus) DeepCopy() *ApplicationPromotionRunStatus {
	if in == nil {
		return nil
	}
	out := new(ApplicationPromotionRunStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationSnapshot) DeepCopyInto(out *ApplicationSnapshot) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationSnapshot.
func (in *ApplicationSnapshot) DeepCopy() *ApplicationSnapshot {
	if in == nil {
		return nil
	}
	out := new(ApplicationSnapshot)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApplicationSnapshot) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationSnapshotComponent) DeepCopyInto(out *ApplicationSnapshotComponent) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationSnapshotComponent.
func (in *ApplicationSnapshotComponent) DeepCopy() *ApplicationSnapshotComponent {
	if in == nil {
		return nil
	}
	out := new(ApplicationSnapshotComponent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationSnapshotEnvironmentBinding) DeepCopyInto(out *ApplicationSnapshotEnvironmentBinding) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationSnapshotEnvironmentBinding.
func (in *ApplicationSnapshotEnvironmentBinding) DeepCopy() *ApplicationSnapshotEnvironmentBinding {
	if in == nil {
		return nil
	}
	out := new(ApplicationSnapshotEnvironmentBinding)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApplicationSnapshotEnvironmentBinding) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationSnapshotEnvironmentBindingList) DeepCopyInto(out *ApplicationSnapshotEnvironmentBindingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ApplicationSnapshotEnvironmentBinding, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationSnapshotEnvironmentBindingList.
func (in *ApplicationSnapshotEnvironmentBindingList) DeepCopy() *ApplicationSnapshotEnvironmentBindingList {
	if in == nil {
		return nil
	}
	out := new(ApplicationSnapshotEnvironmentBindingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApplicationSnapshotEnvironmentBindingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationSnapshotEnvironmentBindingSpec) DeepCopyInto(out *ApplicationSnapshotEnvironmentBindingSpec) {
	*out = *in
	if in.Components != nil {
		in, out := &in.Components, &out.Components
		*out = make([]BindingComponent, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationSnapshotEnvironmentBindingSpec.
func (in *ApplicationSnapshotEnvironmentBindingSpec) DeepCopy() *ApplicationSnapshotEnvironmentBindingSpec {
	if in == nil {
		return nil
	}
	out := new(ApplicationSnapshotEnvironmentBindingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationSnapshotEnvironmentBindingStatus) DeepCopyInto(out *ApplicationSnapshotEnvironmentBindingStatus) {
	*out = *in
	if in.GitOpsDeployments != nil {
		in, out := &in.GitOpsDeployments, &out.GitOpsDeployments
		*out = make([]BindingStatusGitOpsDeployment, len(*in))
		copy(*out, *in)
	}
	if in.Components != nil {
		in, out := &in.Components, &out.Components
		*out = make([]BindingComponentStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.GitOpsRepoConditions != nil {
		in, out := &in.GitOpsRepoConditions, &out.GitOpsRepoConditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationSnapshotEnvironmentBindingStatus.
func (in *ApplicationSnapshotEnvironmentBindingStatus) DeepCopy() *ApplicationSnapshotEnvironmentBindingStatus {
	if in == nil {
		return nil
	}
	out := new(ApplicationSnapshotEnvironmentBindingStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationSnapshotList) DeepCopyInto(out *ApplicationSnapshotList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ApplicationSnapshot, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationSnapshotList.
func (in *ApplicationSnapshotList) DeepCopy() *ApplicationSnapshotList {
	if in == nil {
		return nil
	}
	out := new(ApplicationSnapshotList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApplicationSnapshotList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationSnapshotSpec) DeepCopyInto(out *ApplicationSnapshotSpec) {
	*out = *in
	if in.Components != nil {
		in, out := &in.Components, &out.Components
		*out = make([]ApplicationSnapshotComponent, len(*in))
		copy(*out, *in)
	}
	in.Artifacts.DeepCopyInto(&out.Artifacts)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationSnapshotSpec.
func (in *ApplicationSnapshotSpec) DeepCopy() *ApplicationSnapshotSpec {
	if in == nil {
		return nil
	}
	out := new(ApplicationSnapshotSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationSnapshotStatus) DeepCopyInto(out *ApplicationSnapshotStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationSnapshotStatus.
func (in *ApplicationSnapshotStatus) DeepCopy() *ApplicationSnapshotStatus {
	if in == nil {
		return nil
	}
	out := new(ApplicationSnapshotStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationSpec) DeepCopyInto(out *ApplicationSpec) {
	*out = *in
	out.AppModelRepository = in.AppModelRepository
	out.GitOpsRepository = in.GitOpsRepository
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationSpec.
func (in *ApplicationSpec) DeepCopy() *ApplicationSpec {
	if in == nil {
		return nil
	}
	out := new(ApplicationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationStatus) DeepCopyInto(out *ApplicationStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationStatus.
func (in *ApplicationStatus) DeepCopy() *ApplicationStatus {
	if in == nil {
		return nil
	}
	out := new(ApplicationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutomatedPromotionConfiguration) DeepCopyInto(out *AutomatedPromotionConfiguration) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutomatedPromotionConfiguration.
func (in *AutomatedPromotionConfiguration) DeepCopy() *AutomatedPromotionConfiguration {
	if in == nil {
		return nil
	}
	out := new(AutomatedPromotionConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BindingComponent) DeepCopyInto(out *BindingComponent) {
	*out = *in
	in.Configuration.DeepCopyInto(&out.Configuration)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BindingComponent.
func (in *BindingComponent) DeepCopy() *BindingComponent {
	if in == nil {
		return nil
	}
	out := new(BindingComponent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BindingComponentConfiguration) DeepCopyInto(out *BindingComponentConfiguration) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(corev1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]EnvVarPair, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BindingComponentConfiguration.
func (in *BindingComponentConfiguration) DeepCopy() *BindingComponentConfiguration {
	if in == nil {
		return nil
	}
	out := new(BindingComponentConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BindingComponentGitOpsRepository) DeepCopyInto(out *BindingComponentGitOpsRepository) {
	*out = *in
	if in.GeneratedResources != nil {
		in, out := &in.GeneratedResources, &out.GeneratedResources
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BindingComponentGitOpsRepository.
func (in *BindingComponentGitOpsRepository) DeepCopy() *BindingComponentGitOpsRepository {
	if in == nil {
		return nil
	}
	out := new(BindingComponentGitOpsRepository)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BindingComponentStatus) DeepCopyInto(out *BindingComponentStatus) {
	*out = *in
	in.GitOpsRepository.DeepCopyInto(&out.GitOpsRepository)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BindingComponentStatus.
func (in *BindingComponentStatus) DeepCopy() *BindingComponentStatus {
	if in == nil {
		return nil
	}
	out := new(BindingComponentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BindingStatusGitOpsDeployment) DeepCopyInto(out *BindingStatusGitOpsDeployment) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BindingStatusGitOpsDeployment.
func (in *BindingStatusGitOpsDeployment) DeepCopy() *BindingStatusGitOpsDeployment {
	if in == nil {
		return nil
	}
	out := new(BindingStatusGitOpsDeployment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Component) DeepCopyInto(out *Component) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Component.
func (in *Component) DeepCopy() *Component {
	if in == nil {
		return nil
	}
	out := new(Component)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Component) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComponentDetectionDescription) DeepCopyInto(out *ComponentDetectionDescription) {
	*out = *in
	in.ComponentStub.DeepCopyInto(&out.ComponentStub)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentDetectionDescription.
func (in *ComponentDetectionDescription) DeepCopy() *ComponentDetectionDescription {
	if in == nil {
		return nil
	}
	out := new(ComponentDetectionDescription)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in ComponentDetectionMap) DeepCopyInto(out *ComponentDetectionMap) {
	{
		in := &in
		*out = make(ComponentDetectionMap, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentDetectionMap.
func (in ComponentDetectionMap) DeepCopy() ComponentDetectionMap {
	if in == nil {
		return nil
	}
	out := new(ComponentDetectionMap)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComponentDetectionQuery) DeepCopyInto(out *ComponentDetectionQuery) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentDetectionQuery.
func (in *ComponentDetectionQuery) DeepCopy() *ComponentDetectionQuery {
	if in == nil {
		return nil
	}
	out := new(ComponentDetectionQuery)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ComponentDetectionQuery) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComponentDetectionQueryList) DeepCopyInto(out *ComponentDetectionQueryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ComponentDetectionQuery, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentDetectionQueryList.
func (in *ComponentDetectionQueryList) DeepCopy() *ComponentDetectionQueryList {
	if in == nil {
		return nil
	}
	out := new(ComponentDetectionQueryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ComponentDetectionQueryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComponentDetectionQuerySpec) DeepCopyInto(out *ComponentDetectionQuerySpec) {
	*out = *in
	out.GitSource = in.GitSource
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentDetectionQuerySpec.
func (in *ComponentDetectionQuerySpec) DeepCopy() *ComponentDetectionQuerySpec {
	if in == nil {
		return nil
	}
	out := new(ComponentDetectionQuerySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComponentDetectionQueryStatus) DeepCopyInto(out *ComponentDetectionQueryStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ComponentDetected != nil {
		in, out := &in.ComponentDetected, &out.ComponentDetected
		*out = make(ComponentDetectionMap, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentDetectionQueryStatus.
func (in *ComponentDetectionQueryStatus) DeepCopy() *ComponentDetectionQueryStatus {
	if in == nil {
		return nil
	}
	out := new(ComponentDetectionQueryStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComponentList) DeepCopyInto(out *ComponentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Component, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentList.
func (in *ComponentList) DeepCopy() *ComponentList {
	if in == nil {
		return nil
	}
	out := new(ComponentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ComponentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComponentSource) DeepCopyInto(out *ComponentSource) {
	*out = *in
	in.ComponentSourceUnion.DeepCopyInto(&out.ComponentSourceUnion)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentSource.
func (in *ComponentSource) DeepCopy() *ComponentSource {
	if in == nil {
		return nil
	}
	out := new(ComponentSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComponentSourceUnion) DeepCopyInto(out *ComponentSourceUnion) {
	*out = *in
	if in.GitSource != nil {
		in, out := &in.GitSource, &out.GitSource
		*out = new(GitSource)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentSourceUnion.
func (in *ComponentSourceUnion) DeepCopy() *ComponentSourceUnion {
	if in == nil {
		return nil
	}
	out := new(ComponentSourceUnion)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComponentSpec) DeepCopyInto(out *ComponentSpec) {
	*out = *in
	in.Source.DeepCopyInto(&out.Source)
	in.Resources.DeepCopyInto(&out.Resources)
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]corev1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentSpec.
func (in *ComponentSpec) DeepCopy() *ComponentSpec {
	if in == nil {
		return nil
	}
	out := new(ComponentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComponentStatus) DeepCopyInto(out *ComponentStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.GitOps = in.GitOps
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentStatus.
func (in *ComponentStatus) DeepCopy() *ComponentStatus {
	if in == nil {
		return nil
	}
	out := new(ComponentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvVarPair) DeepCopyInto(out *EnvVarPair) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvVarPair.
func (in *EnvVarPair) DeepCopy() *EnvVarPair {
	if in == nil {
		return nil
	}
	out := new(EnvVarPair)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Environment) DeepCopyInto(out *Environment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Environment.
func (in *Environment) DeepCopy() *Environment {
	if in == nil {
		return nil
	}
	out := new(Environment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Environment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentConfiguration) DeepCopyInto(out *EnvironmentConfiguration) {
	*out = *in
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]EnvVarPair, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentConfiguration.
func (in *EnvironmentConfiguration) DeepCopy() *EnvironmentConfiguration {
	if in == nil {
		return nil
	}
	out := new(EnvironmentConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentList) DeepCopyInto(out *EnvironmentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Environment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentList.
func (in *EnvironmentList) DeepCopy() *EnvironmentList {
	if in == nil {
		return nil
	}
	out := new(EnvironmentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EnvironmentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentSpec) DeepCopyInto(out *EnvironmentSpec) {
	*out = *in
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.Configuration.DeepCopyInto(&out.Configuration)
	if in.UnstableConfigurationFields != nil {
		in, out := &in.UnstableConfigurationFields, &out.UnstableConfigurationFields
		*out = new(UnstableEnvironmentConfiguration)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentSpec.
func (in *EnvironmentSpec) DeepCopy() *EnvironmentSpec {
	if in == nil {
		return nil
	}
	out := new(EnvironmentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentStatus) DeepCopyInto(out *EnvironmentStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentStatus.
func (in *EnvironmentStatus) DeepCopy() *EnvironmentStatus {
	if in == nil {
		return nil
	}
	out := new(EnvironmentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitOpsStatus) DeepCopyInto(out *GitOpsStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitOpsStatus.
func (in *GitOpsStatus) DeepCopy() *GitOpsStatus {
	if in == nil {
		return nil
	}
	out := new(GitOpsStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitSource) DeepCopyInto(out *GitSource) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitSource.
func (in *GitSource) DeepCopy() *GitSource {
	if in == nil {
		return nil
	}
	out := new(GitSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubernetesClusterCredentials) DeepCopyInto(out *KubernetesClusterCredentials) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubernetesClusterCredentials.
func (in *KubernetesClusterCredentials) DeepCopy() *KubernetesClusterCredentials {
	if in == nil {
		return nil
	}
	out := new(KubernetesClusterCredentials)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManualPromotionConfiguration) DeepCopyInto(out *ManualPromotionConfiguration) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManualPromotionConfiguration.
func (in *ManualPromotionConfiguration) DeepCopy() *ManualPromotionConfiguration {
	if in == nil {
		return nil
	}
	out := new(ManualPromotionConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PromotionRunEnvironmentStatus) DeepCopyInto(out *PromotionRunEnvironmentStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PromotionRunEnvironmentStatus.
func (in *PromotionRunEnvironmentStatus) DeepCopy() *PromotionRunEnvironmentStatus {
	if in == nil {
		return nil
	}
	out := new(PromotionRunEnvironmentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnapshotArtifacts) DeepCopyInto(out *SnapshotArtifacts) {
	*out = *in
	if in.UnstableFields != nil {
		in, out := &in.UnstableFields, &out.UnstableFields
		*out = new(apiextensionsv1.JSON)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnapshotArtifacts.
func (in *SnapshotArtifacts) DeepCopy() *SnapshotArtifacts {
	if in == nil {
		return nil
	}
	out := new(SnapshotArtifacts)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UnstableEnvironmentConfiguration) DeepCopyInto(out *UnstableEnvironmentConfiguration) {
	*out = *in
	out.KubernetesClusterCredentials = in.KubernetesClusterCredentials
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UnstableEnvironmentConfiguration.
func (in *UnstableEnvironmentConfiguration) DeepCopy() *UnstableEnvironmentConfiguration {
	if in == nil {
		return nil
	}
	out := new(UnstableEnvironmentConfiguration)
	in.DeepCopyInto(out)
	return out
}
