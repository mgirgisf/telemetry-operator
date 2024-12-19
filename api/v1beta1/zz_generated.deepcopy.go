//go:build !ignore_autogenerated

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

package v1beta1

import (
	networkv1beta1 "github.com/openstack-k8s-operators/infra-operator/apis/network/v1beta1"
	"github.com/openstack-k8s-operators/lib-common/modules/common/condition"
	"github.com/openstack-k8s-operators/lib-common/modules/common/service"
	"github.com/rhobs/observability-operator/pkg/apis/monitoring/v1alpha1"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIOverrideSpec) DeepCopyInto(out *APIOverrideSpec) {
	*out = *in
	if in.Service != nil {
		in, out := &in.Service, &out.Service
		*out = make(map[service.Endpoint]service.RoutedOverrideSpec, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIOverrideSpec.
func (in *APIOverrideSpec) DeepCopy() *APIOverrideSpec {
	if in == nil {
		return nil
	}
	out := new(APIOverrideSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Aodh) DeepCopyInto(out *Aodh) {
	*out = *in
	in.AodhCore.DeepCopyInto(&out.AodhCore)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Aodh.
func (in *Aodh) DeepCopy() *Aodh {
	if in == nil {
		return nil
	}
	out := new(Aodh)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AodhCore) DeepCopyInto(out *AodhCore) {
	*out = *in
	out.PasswordSelectors = in.PasswordSelectors
	if in.DefaultConfigOverwrite != nil {
		in, out := &in.DefaultConfigOverwrite, &out.DefaultConfigOverwrite
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.NetworkAttachmentDefinitions != nil {
		in, out := &in.NetworkAttachmentDefinitions, &out.NetworkAttachmentDefinitions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.Override.DeepCopyInto(&out.Override)
	in.TLS.DeepCopyInto(&out.TLS)
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = new(map[string]string)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[string]string, len(*in))
			for key, val := range *in {
				(*out)[key] = val
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AodhCore.
func (in *AodhCore) DeepCopy() *AodhCore {
	if in == nil {
		return nil
	}
	out := new(AodhCore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Autoscaling) DeepCopyInto(out *Autoscaling) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Autoscaling.
func (in *Autoscaling) DeepCopy() *Autoscaling {
	if in == nil {
		return nil
	}
	out := new(Autoscaling)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Autoscaling) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoscalingDefaults) DeepCopyInto(out *AutoscalingDefaults) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoscalingDefaults.
func (in *AutoscalingDefaults) DeepCopy() *AutoscalingDefaults {
	if in == nil {
		return nil
	}
	out := new(AutoscalingDefaults)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoscalingList) DeepCopyInto(out *AutoscalingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Autoscaling, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoscalingList.
func (in *AutoscalingList) DeepCopy() *AutoscalingList {
	if in == nil {
		return nil
	}
	out := new(AutoscalingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AutoscalingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoscalingSection) DeepCopyInto(out *AutoscalingSection) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	in.AutoscalingSpec.DeepCopyInto(&out.AutoscalingSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoscalingSection.
func (in *AutoscalingSection) DeepCopy() *AutoscalingSection {
	if in == nil {
		return nil
	}
	out := new(AutoscalingSection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoscalingSectionCore) DeepCopyInto(out *AutoscalingSectionCore) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	in.AutoscalingSpecCore.DeepCopyInto(&out.AutoscalingSpecCore)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoscalingSectionCore.
func (in *AutoscalingSectionCore) DeepCopy() *AutoscalingSectionCore {
	if in == nil {
		return nil
	}
	out := new(AutoscalingSectionCore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoscalingSpec) DeepCopyInto(out *AutoscalingSpec) {
	*out = *in
	in.AutoscalingSpecBase.DeepCopyInto(&out.AutoscalingSpecBase)
	in.Aodh.DeepCopyInto(&out.Aodh)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoscalingSpec.
func (in *AutoscalingSpec) DeepCopy() *AutoscalingSpec {
	if in == nil {
		return nil
	}
	out := new(AutoscalingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoscalingSpecBase) DeepCopyInto(out *AutoscalingSpecBase) {
	*out = *in
	if in.PrometheusTLSCaCertSecret != nil {
		in, out := &in.PrometheusTLSCaCertSecret, &out.PrometheusTLSCaCertSecret
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoscalingSpecBase.
func (in *AutoscalingSpecBase) DeepCopy() *AutoscalingSpecBase {
	if in == nil {
		return nil
	}
	out := new(AutoscalingSpecBase)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoscalingSpecCore) DeepCopyInto(out *AutoscalingSpecCore) {
	*out = *in
	in.AutoscalingSpecBase.DeepCopyInto(&out.AutoscalingSpecBase)
	in.Aodh.DeepCopyInto(&out.Aodh)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoscalingSpecCore.
func (in *AutoscalingSpecCore) DeepCopy() *AutoscalingSpecCore {
	if in == nil {
		return nil
	}
	out := new(AutoscalingSpecCore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoscalingStatus) DeepCopyInto(out *AutoscalingStatus) {
	*out = *in
	if in.Hash != nil {
		in, out := &in.Hash, &out.Hash
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(condition.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Networks != nil {
		in, out := &in.Networks, &out.Networks
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.APIEndpoints != nil {
		in, out := &in.APIEndpoints, &out.APIEndpoints
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoscalingStatus.
func (in *AutoscalingStatus) DeepCopy() *AutoscalingStatus {
	if in == nil {
		return nil
	}
	out := new(AutoscalingStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Ceilometer) DeepCopyInto(out *Ceilometer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.CeilometerStatus.DeepCopyInto(&out.CeilometerStatus)
	in.KSMStatus.DeepCopyInto(&out.KSMStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ceilometer.
func (in *Ceilometer) DeepCopy() *Ceilometer {
	if in == nil {
		return nil
	}
	out := new(Ceilometer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Ceilometer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CeilometerDefaults) DeepCopyInto(out *CeilometerDefaults) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CeilometerDefaults.
func (in *CeilometerDefaults) DeepCopy() *CeilometerDefaults {
	if in == nil {
		return nil
	}
	out := new(CeilometerDefaults)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CeilometerList) DeepCopyInto(out *CeilometerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Ceilometer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CeilometerList.
func (in *CeilometerList) DeepCopy() *CeilometerList {
	if in == nil {
		return nil
	}
	out := new(CeilometerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CeilometerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CeilometerSection) DeepCopyInto(out *CeilometerSection) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	in.CeilometerSpec.DeepCopyInto(&out.CeilometerSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CeilometerSection.
func (in *CeilometerSection) DeepCopy() *CeilometerSection {
	if in == nil {
		return nil
	}
	out := new(CeilometerSection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CeilometerSectionCore) DeepCopyInto(out *CeilometerSectionCore) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	in.CeilometerSpecCore.DeepCopyInto(&out.CeilometerSpecCore)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CeilometerSectionCore.
func (in *CeilometerSectionCore) DeepCopy() *CeilometerSectionCore {
	if in == nil {
		return nil
	}
	out := new(CeilometerSectionCore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CeilometerSpec) DeepCopyInto(out *CeilometerSpec) {
	*out = *in
	in.CeilometerSpecCore.DeepCopyInto(&out.CeilometerSpecCore)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CeilometerSpec.
func (in *CeilometerSpec) DeepCopy() *CeilometerSpec {
	if in == nil {
		return nil
	}
	out := new(CeilometerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CeilometerSpecCore) DeepCopyInto(out *CeilometerSpecCore) {
	*out = *in
	out.PasswordSelectors = in.PasswordSelectors
	if in.DefaultConfigOverwrite != nil {
		in, out := &in.DefaultConfigOverwrite, &out.DefaultConfigOverwrite
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.NetworkAttachmentDefinitions != nil {
		in, out := &in.NetworkAttachmentDefinitions, &out.NetworkAttachmentDefinitions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.MysqldExporterEnabled != nil {
		in, out := &in.MysqldExporterEnabled, &out.MysqldExporterEnabled
		*out = new(bool)
		**out = **in
	}
	in.TLS.DeepCopyInto(&out.TLS)
	in.KSMTLS.DeepCopyInto(&out.KSMTLS)
	in.MysqldExporterTLS.DeepCopyInto(&out.MysqldExporterTLS)
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = new(map[string]string)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[string]string, len(*in))
			for key, val := range *in {
				(*out)[key] = val
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CeilometerSpecCore.
func (in *CeilometerSpecCore) DeepCopy() *CeilometerSpecCore {
	if in == nil {
		return nil
	}
	out := new(CeilometerSpecCore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CeilometerStatus) DeepCopyInto(out *CeilometerStatus) {
	*out = *in
	if in.Hash != nil {
		in, out := &in.Hash, &out.Hash
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(condition.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Networks != nil {
		in, out := &in.Networks, &out.Networks
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.MysqldExporterHash != nil {
		in, out := &in.MysqldExporterHash, &out.MysqldExporterHash
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.MysqldExporterExportedGaleras != nil {
		in, out := &in.MysqldExporterExportedGaleras, &out.MysqldExporterExportedGaleras
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CeilometerStatus.
func (in *CeilometerStatus) DeepCopy() *CeilometerStatus {
	if in == nil {
		return nil
	}
	out := new(CeilometerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KSMStatus) DeepCopyInto(out *KSMStatus) {
	*out = *in
	if in.Hash != nil {
		in, out := &in.Hash, &out.Hash
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(condition.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KSMStatus.
func (in *KSMStatus) DeepCopy() *KSMStatus {
	if in == nil {
		return nil
	}
	out := new(KSMStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Logging) DeepCopyInto(out *Logging) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Logging.
func (in *Logging) DeepCopy() *Logging {
	if in == nil {
		return nil
	}
	out := new(Logging)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Logging) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoggingList) DeepCopyInto(out *LoggingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Logging, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoggingList.
func (in *LoggingList) DeepCopy() *LoggingList {
	if in == nil {
		return nil
	}
	out := new(LoggingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LoggingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoggingSection) DeepCopyInto(out *LoggingSection) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	in.LoggingSpec.DeepCopyInto(&out.LoggingSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoggingSection.
func (in *LoggingSection) DeepCopy() *LoggingSection {
	if in == nil {
		return nil
	}
	out := new(LoggingSection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoggingSpec) DeepCopyInto(out *LoggingSpec) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoggingSpec.
func (in *LoggingSpec) DeepCopy() *LoggingSpec {
	if in == nil {
		return nil
	}
	out := new(LoggingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoggingStatus) DeepCopyInto(out *LoggingStatus) {
	*out = *in
	if in.Hash != nil {
		in, out := &in.Hash, &out.Hash
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(condition.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoggingStatus.
func (in *LoggingStatus) DeepCopy() *LoggingStatus {
	if in == nil {
		return nil
	}
	out := new(LoggingStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricStorage) DeepCopyInto(out *MetricStorage) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricStorage.
func (in *MetricStorage) DeepCopy() *MetricStorage {
	if in == nil {
		return nil
	}
	out := new(MetricStorage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MetricStorage) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricStorageList) DeepCopyInto(out *MetricStorageList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MetricStorage, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricStorageList.
func (in *MetricStorageList) DeepCopy() *MetricStorageList {
	if in == nil {
		return nil
	}
	out := new(MetricStorageList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MetricStorageList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricStorageSection) DeepCopyInto(out *MetricStorageSection) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	in.MetricStorageSpec.DeepCopyInto(&out.MetricStorageSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricStorageSection.
func (in *MetricStorageSection) DeepCopy() *MetricStorageSection {
	if in == nil {
		return nil
	}
	out := new(MetricStorageSection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricStorageSpec) DeepCopyInto(out *MetricStorageSpec) {
	*out = *in
	if in.DataplaneNetwork != nil {
		in, out := &in.DataplaneNetwork, &out.DataplaneNetwork
		*out = new(networkv1beta1.NetNameStr)
		**out = **in
	}
	if in.MonitoringStack != nil {
		in, out := &in.MonitoringStack, &out.MonitoringStack
		*out = new(MonitoringStack)
		(*in).DeepCopyInto(*out)
	}
	if in.CustomMonitoringStack != nil {
		in, out := &in.CustomMonitoringStack, &out.CustomMonitoringStack
		*out = new(v1alpha1.MonitoringStackSpec)
		(*in).DeepCopyInto(*out)
	}
	in.PrometheusTLS.DeepCopyInto(&out.PrometheusTLS)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricStorageSpec.
func (in *MetricStorageSpec) DeepCopy() *MetricStorageSpec {
	if in == nil {
		return nil
	}
	out := new(MetricStorageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricStorageStatus) DeepCopyInto(out *MetricStorageStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(condition.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricStorageStatus.
func (in *MetricStorageStatus) DeepCopy() *MetricStorageStatus {
	if in == nil {
		return nil
	}
	out := new(MetricStorageStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitoringStack) DeepCopyInto(out *MonitoringStack) {
	*out = *in
	in.Storage.DeepCopyInto(&out.Storage)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitoringStack.
func (in *MonitoringStack) DeepCopy() *MonitoringStack {
	if in == nil {
		return nil
	}
	out := new(MonitoringStack)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PasswordsSelector) DeepCopyInto(out *PasswordsSelector) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PasswordsSelector.
func (in *PasswordsSelector) DeepCopy() *PasswordsSelector {
	if in == nil {
		return nil
	}
	out := new(PasswordsSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PersistentStorage) DeepCopyInto(out *PersistentStorage) {
	*out = *in
	in.PvcStorageSelector.DeepCopyInto(&out.PvcStorageSelector)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PersistentStorage.
func (in *PersistentStorage) DeepCopy() *PersistentStorage {
	if in == nil {
		return nil
	}
	out := new(PersistentStorage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Storage) DeepCopyInto(out *Storage) {
	*out = *in
	if in.Persistent != nil {
		in, out := &in.Persistent, &out.Persistent
		*out = new(PersistentStorage)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Storage.
func (in *Storage) DeepCopy() *Storage {
	if in == nil {
		return nil
	}
	out := new(Storage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Telemetry) DeepCopyInto(out *Telemetry) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Telemetry.
func (in *Telemetry) DeepCopy() *Telemetry {
	if in == nil {
		return nil
	}
	out := new(Telemetry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Telemetry) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TelemetryDefaults) DeepCopyInto(out *TelemetryDefaults) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TelemetryDefaults.
func (in *TelemetryDefaults) DeepCopy() *TelemetryDefaults {
	if in == nil {
		return nil
	}
	out := new(TelemetryDefaults)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TelemetryList) DeepCopyInto(out *TelemetryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Telemetry, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TelemetryList.
func (in *TelemetryList) DeepCopy() *TelemetryList {
	if in == nil {
		return nil
	}
	out := new(TelemetryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TelemetryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TelemetrySpec) DeepCopyInto(out *TelemetrySpec) {
	*out = *in
	in.TelemetrySpecBase.DeepCopyInto(&out.TelemetrySpecBase)
	in.Autoscaling.DeepCopyInto(&out.Autoscaling)
	in.Ceilometer.DeepCopyInto(&out.Ceilometer)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TelemetrySpec.
func (in *TelemetrySpec) DeepCopy() *TelemetrySpec {
	if in == nil {
		return nil
	}
	out := new(TelemetrySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TelemetrySpecBase) DeepCopyInto(out *TelemetrySpecBase) {
	*out = *in
	in.MetricStorage.DeepCopyInto(&out.MetricStorage)
	in.Logging.DeepCopyInto(&out.Logging)
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = new(map[string]string)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[string]string, len(*in))
			for key, val := range *in {
				(*out)[key] = val
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TelemetrySpecBase.
func (in *TelemetrySpecBase) DeepCopy() *TelemetrySpecBase {
	if in == nil {
		return nil
	}
	out := new(TelemetrySpecBase)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TelemetrySpecCore) DeepCopyInto(out *TelemetrySpecCore) {
	*out = *in
	in.TelemetrySpecBase.DeepCopyInto(&out.TelemetrySpecBase)
	in.Autoscaling.DeepCopyInto(&out.Autoscaling)
	in.Ceilometer.DeepCopyInto(&out.Ceilometer)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TelemetrySpecCore.
func (in *TelemetrySpecCore) DeepCopy() *TelemetrySpecCore {
	if in == nil {
		return nil
	}
	out := new(TelemetrySpecCore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TelemetryStatus) DeepCopyInto(out *TelemetryStatus) {
	*out = *in
	if in.Hash != nil {
		in, out := &in.Hash, &out.Hash
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(condition.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TelemetryStatus.
func (in *TelemetryStatus) DeepCopy() *TelemetryStatus {
	if in == nil {
		return nil
	}
	out := new(TelemetryStatus)
	in.DeepCopyInto(out)
	return out
}
