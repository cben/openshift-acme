// +build !ignore_autogenerated_openshift

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	intstr "k8s.io/apimachinery/pkg/util/intstr"
	api_v1 "k8s.io/kubernetes/pkg/api/v1"
	reflect "reflect"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_CustomDeploymentStrategyParams, InType: reflect.TypeOf(&CustomDeploymentStrategyParams{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_DeploymentCause, InType: reflect.TypeOf(&DeploymentCause{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_DeploymentCauseImageTrigger, InType: reflect.TypeOf(&DeploymentCauseImageTrigger{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_DeploymentCondition, InType: reflect.TypeOf(&DeploymentCondition{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_DeploymentConfig, InType: reflect.TypeOf(&DeploymentConfig{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_DeploymentConfigList, InType: reflect.TypeOf(&DeploymentConfigList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_DeploymentConfigRollback, InType: reflect.TypeOf(&DeploymentConfigRollback{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_DeploymentConfigRollbackSpec, InType: reflect.TypeOf(&DeploymentConfigRollbackSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_DeploymentConfigSpec, InType: reflect.TypeOf(&DeploymentConfigSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_DeploymentConfigStatus, InType: reflect.TypeOf(&DeploymentConfigStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_DeploymentDetails, InType: reflect.TypeOf(&DeploymentDetails{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_DeploymentLog, InType: reflect.TypeOf(&DeploymentLog{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_DeploymentLogOptions, InType: reflect.TypeOf(&DeploymentLogOptions{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_DeploymentRequest, InType: reflect.TypeOf(&DeploymentRequest{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_DeploymentStrategy, InType: reflect.TypeOf(&DeploymentStrategy{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_DeploymentTriggerImageChangeParams, InType: reflect.TypeOf(&DeploymentTriggerImageChangeParams{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_DeploymentTriggerPolicy, InType: reflect.TypeOf(&DeploymentTriggerPolicy{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_ExecNewPodHook, InType: reflect.TypeOf(&ExecNewPodHook{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_LifecycleHook, InType: reflect.TypeOf(&LifecycleHook{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_RecreateDeploymentStrategyParams, InType: reflect.TypeOf(&RecreateDeploymentStrategyParams{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_RollingDeploymentStrategyParams, InType: reflect.TypeOf(&RollingDeploymentStrategyParams{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_TagImageHook, InType: reflect.TypeOf(&TagImageHook{})},
	)
}

func DeepCopy_v1_CustomDeploymentStrategyParams(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*CustomDeploymentStrategyParams)
		out := out.(*CustomDeploymentStrategyParams)
		*out = *in
		if in.Environment != nil {
			in, out := &in.Environment, &out.Environment
			*out = make([]api_v1.EnvVar, len(*in))
			for i := range *in {
				if err := api_v1.DeepCopy_v1_EnvVar(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		if in.Command != nil {
			in, out := &in.Command, &out.Command
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		return nil
	}
}

func DeepCopy_v1_DeploymentCause(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*DeploymentCause)
		out := out.(*DeploymentCause)
		*out = *in
		if in.ImageTrigger != nil {
			in, out := &in.ImageTrigger, &out.ImageTrigger
			*out = new(DeploymentCauseImageTrigger)
			**out = **in
		}
		return nil
	}
}

func DeepCopy_v1_DeploymentCauseImageTrigger(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*DeploymentCauseImageTrigger)
		out := out.(*DeploymentCauseImageTrigger)
		*out = *in
		return nil
	}
}

func DeepCopy_v1_DeploymentCondition(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*DeploymentCondition)
		out := out.(*DeploymentCondition)
		*out = *in
		out.LastUpdateTime = in.LastUpdateTime.DeepCopy()
		out.LastTransitionTime = in.LastTransitionTime.DeepCopy()
		return nil
	}
}

func DeepCopy_v1_DeploymentConfig(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*DeploymentConfig)
		out := out.(*DeploymentConfig)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*meta_v1.ObjectMeta)
		}
		if err := DeepCopy_v1_DeploymentConfigSpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		if err := DeepCopy_v1_DeploymentConfigStatus(&in.Status, &out.Status, c); err != nil {
			return err
		}
		return nil
	}
}

func DeepCopy_v1_DeploymentConfigList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*DeploymentConfigList)
		out := out.(*DeploymentConfigList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]DeploymentConfig, len(*in))
			for i := range *in {
				if err := DeepCopy_v1_DeploymentConfig(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

func DeepCopy_v1_DeploymentConfigRollback(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*DeploymentConfigRollback)
		out := out.(*DeploymentConfigRollback)
		*out = *in
		if in.UpdatedAnnotations != nil {
			in, out := &in.UpdatedAnnotations, &out.UpdatedAnnotations
			*out = make(map[string]string)
			for key, val := range *in {
				(*out)[key] = val
			}
		}
		return nil
	}
}

func DeepCopy_v1_DeploymentConfigRollbackSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*DeploymentConfigRollbackSpec)
		out := out.(*DeploymentConfigRollbackSpec)
		*out = *in
		return nil
	}
}

func DeepCopy_v1_DeploymentConfigSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*DeploymentConfigSpec)
		out := out.(*DeploymentConfigSpec)
		*out = *in
		if err := DeepCopy_v1_DeploymentStrategy(&in.Strategy, &out.Strategy, c); err != nil {
			return err
		}
		if in.Triggers != nil {
			in, out := &in.Triggers, &out.Triggers
			*out = make(DeploymentTriggerPolicies, len(*in))
			for i := range *in {
				if err := DeepCopy_v1_DeploymentTriggerPolicy(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		if in.RevisionHistoryLimit != nil {
			in, out := &in.RevisionHistoryLimit, &out.RevisionHistoryLimit
			*out = new(int32)
			**out = **in
		}
		if in.Selector != nil {
			in, out := &in.Selector, &out.Selector
			*out = make(map[string]string)
			for key, val := range *in {
				(*out)[key] = val
			}
		}
		if in.Template != nil {
			in, out := &in.Template, &out.Template
			*out = new(api_v1.PodTemplateSpec)
			if err := api_v1.DeepCopy_v1_PodTemplateSpec(*in, *out, c); err != nil {
				return err
			}
		}
		return nil
	}
}

func DeepCopy_v1_DeploymentConfigStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*DeploymentConfigStatus)
		out := out.(*DeploymentConfigStatus)
		*out = *in
		if in.Details != nil {
			in, out := &in.Details, &out.Details
			*out = new(DeploymentDetails)
			if err := DeepCopy_v1_DeploymentDetails(*in, *out, c); err != nil {
				return err
			}
		}
		if in.Conditions != nil {
			in, out := &in.Conditions, &out.Conditions
			*out = make([]DeploymentCondition, len(*in))
			for i := range *in {
				if err := DeepCopy_v1_DeploymentCondition(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

func DeepCopy_v1_DeploymentDetails(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*DeploymentDetails)
		out := out.(*DeploymentDetails)
		*out = *in
		if in.Causes != nil {
			in, out := &in.Causes, &out.Causes
			*out = make([]DeploymentCause, len(*in))
			for i := range *in {
				if err := DeepCopy_v1_DeploymentCause(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

func DeepCopy_v1_DeploymentLog(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*DeploymentLog)
		out := out.(*DeploymentLog)
		*out = *in
		return nil
	}
}

func DeepCopy_v1_DeploymentLogOptions(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*DeploymentLogOptions)
		out := out.(*DeploymentLogOptions)
		*out = *in
		if in.SinceSeconds != nil {
			in, out := &in.SinceSeconds, &out.SinceSeconds
			*out = new(int64)
			**out = **in
		}
		if in.SinceTime != nil {
			in, out := &in.SinceTime, &out.SinceTime
			*out = new(meta_v1.Time)
			**out = (*in).DeepCopy()
		}
		if in.TailLines != nil {
			in, out := &in.TailLines, &out.TailLines
			*out = new(int64)
			**out = **in
		}
		if in.LimitBytes != nil {
			in, out := &in.LimitBytes, &out.LimitBytes
			*out = new(int64)
			**out = **in
		}
		if in.Version != nil {
			in, out := &in.Version, &out.Version
			*out = new(int64)
			**out = **in
		}
		return nil
	}
}

func DeepCopy_v1_DeploymentRequest(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*DeploymentRequest)
		out := out.(*DeploymentRequest)
		*out = *in
		if in.ExcludeTriggers != nil {
			in, out := &in.ExcludeTriggers, &out.ExcludeTriggers
			*out = make([]DeploymentTriggerType, len(*in))
			copy(*out, *in)
		}
		return nil
	}
}

func DeepCopy_v1_DeploymentStrategy(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*DeploymentStrategy)
		out := out.(*DeploymentStrategy)
		*out = *in
		if in.CustomParams != nil {
			in, out := &in.CustomParams, &out.CustomParams
			*out = new(CustomDeploymentStrategyParams)
			if err := DeepCopy_v1_CustomDeploymentStrategyParams(*in, *out, c); err != nil {
				return err
			}
		}
		if in.RecreateParams != nil {
			in, out := &in.RecreateParams, &out.RecreateParams
			*out = new(RecreateDeploymentStrategyParams)
			if err := DeepCopy_v1_RecreateDeploymentStrategyParams(*in, *out, c); err != nil {
				return err
			}
		}
		if in.RollingParams != nil {
			in, out := &in.RollingParams, &out.RollingParams
			*out = new(RollingDeploymentStrategyParams)
			if err := DeepCopy_v1_RollingDeploymentStrategyParams(*in, *out, c); err != nil {
				return err
			}
		}
		if err := api_v1.DeepCopy_v1_ResourceRequirements(&in.Resources, &out.Resources, c); err != nil {
			return err
		}
		if in.Labels != nil {
			in, out := &in.Labels, &out.Labels
			*out = make(map[string]string)
			for key, val := range *in {
				(*out)[key] = val
			}
		}
		if in.Annotations != nil {
			in, out := &in.Annotations, &out.Annotations
			*out = make(map[string]string)
			for key, val := range *in {
				(*out)[key] = val
			}
		}
		if in.ActiveDeadlineSeconds != nil {
			in, out := &in.ActiveDeadlineSeconds, &out.ActiveDeadlineSeconds
			*out = new(int64)
			**out = **in
		}
		return nil
	}
}

func DeepCopy_v1_DeploymentTriggerImageChangeParams(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*DeploymentTriggerImageChangeParams)
		out := out.(*DeploymentTriggerImageChangeParams)
		*out = *in
		if in.ContainerNames != nil {
			in, out := &in.ContainerNames, &out.ContainerNames
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		return nil
	}
}

func DeepCopy_v1_DeploymentTriggerPolicy(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*DeploymentTriggerPolicy)
		out := out.(*DeploymentTriggerPolicy)
		*out = *in
		if in.ImageChangeParams != nil {
			in, out := &in.ImageChangeParams, &out.ImageChangeParams
			*out = new(DeploymentTriggerImageChangeParams)
			if err := DeepCopy_v1_DeploymentTriggerImageChangeParams(*in, *out, c); err != nil {
				return err
			}
		}
		return nil
	}
}

func DeepCopy_v1_ExecNewPodHook(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ExecNewPodHook)
		out := out.(*ExecNewPodHook)
		*out = *in
		if in.Command != nil {
			in, out := &in.Command, &out.Command
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		if in.Env != nil {
			in, out := &in.Env, &out.Env
			*out = make([]api_v1.EnvVar, len(*in))
			for i := range *in {
				if err := api_v1.DeepCopy_v1_EnvVar(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		if in.Volumes != nil {
			in, out := &in.Volumes, &out.Volumes
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		return nil
	}
}

func DeepCopy_v1_LifecycleHook(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*LifecycleHook)
		out := out.(*LifecycleHook)
		*out = *in
		if in.ExecNewPod != nil {
			in, out := &in.ExecNewPod, &out.ExecNewPod
			*out = new(ExecNewPodHook)
			if err := DeepCopy_v1_ExecNewPodHook(*in, *out, c); err != nil {
				return err
			}
		}
		if in.TagImages != nil {
			in, out := &in.TagImages, &out.TagImages
			*out = make([]TagImageHook, len(*in))
			copy(*out, *in)
		}
		return nil
	}
}

func DeepCopy_v1_RecreateDeploymentStrategyParams(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*RecreateDeploymentStrategyParams)
		out := out.(*RecreateDeploymentStrategyParams)
		*out = *in
		if in.TimeoutSeconds != nil {
			in, out := &in.TimeoutSeconds, &out.TimeoutSeconds
			*out = new(int64)
			**out = **in
		}
		if in.Pre != nil {
			in, out := &in.Pre, &out.Pre
			*out = new(LifecycleHook)
			if err := DeepCopy_v1_LifecycleHook(*in, *out, c); err != nil {
				return err
			}
		}
		if in.Mid != nil {
			in, out := &in.Mid, &out.Mid
			*out = new(LifecycleHook)
			if err := DeepCopy_v1_LifecycleHook(*in, *out, c); err != nil {
				return err
			}
		}
		if in.Post != nil {
			in, out := &in.Post, &out.Post
			*out = new(LifecycleHook)
			if err := DeepCopy_v1_LifecycleHook(*in, *out, c); err != nil {
				return err
			}
		}
		return nil
	}
}

func DeepCopy_v1_RollingDeploymentStrategyParams(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*RollingDeploymentStrategyParams)
		out := out.(*RollingDeploymentStrategyParams)
		*out = *in
		if in.UpdatePeriodSeconds != nil {
			in, out := &in.UpdatePeriodSeconds, &out.UpdatePeriodSeconds
			*out = new(int64)
			**out = **in
		}
		if in.IntervalSeconds != nil {
			in, out := &in.IntervalSeconds, &out.IntervalSeconds
			*out = new(int64)
			**out = **in
		}
		if in.TimeoutSeconds != nil {
			in, out := &in.TimeoutSeconds, &out.TimeoutSeconds
			*out = new(int64)
			**out = **in
		}
		if in.MaxUnavailable != nil {
			in, out := &in.MaxUnavailable, &out.MaxUnavailable
			*out = new(intstr.IntOrString)
			**out = **in
		}
		if in.MaxSurge != nil {
			in, out := &in.MaxSurge, &out.MaxSurge
			*out = new(intstr.IntOrString)
			**out = **in
		}
		if in.Pre != nil {
			in, out := &in.Pre, &out.Pre
			*out = new(LifecycleHook)
			if err := DeepCopy_v1_LifecycleHook(*in, *out, c); err != nil {
				return err
			}
		}
		if in.Post != nil {
			in, out := &in.Post, &out.Post
			*out = new(LifecycleHook)
			if err := DeepCopy_v1_LifecycleHook(*in, *out, c); err != nil {
				return err
			}
		}
		return nil
	}
}

func DeepCopy_v1_TagImageHook(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*TagImageHook)
		out := out.(*TagImageHook)
		*out = *in
		return nil
	}
}
