// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package compute

import (
	"fmt"
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgresource"
	transport_tpg "github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/transport"
)

const ComputeRegionAutoscalerAssetType string = "compute.googleapis.com/RegionAutoscaler"

func ResourceConverterComputeRegionAutoscaler() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType: ComputeRegionAutoscalerAssetType,
		Convert:   GetComputeRegionAutoscalerCaiObject,
	}
}

func GetComputeRegionAutoscalerCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	name, err := tpgresource.AssetName(d, config, "//compute.googleapis.com/projects/{{project}}/regions/{{region}}/autoscalers/{{name}}")
	if err != nil {
		return []tpgresource.Asset{}, err
	}
	if obj, err := GetComputeRegionAutoscalerApiObject(d, config); err == nil {
		return []tpgresource.Asset{{
			Name: name,
			Type: ComputeRegionAutoscalerAssetType,
			Resource: &tpgresource.AssetResource{
				Version:              "beta",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/compute/beta/rest",
				DiscoveryName:        "RegionAutoscaler",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []tpgresource.Asset{}, err
	}
}

func GetComputeRegionAutoscalerApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandComputeRegionAutoscalerName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandComputeRegionAutoscalerDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	autoscalingPolicyProp, err := expandComputeRegionAutoscalerAutoscalingPolicy(d.Get("autoscaling_policy"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("autoscaling_policy"); !tpgresource.IsEmptyValue(reflect.ValueOf(autoscalingPolicyProp)) && (ok || !reflect.DeepEqual(v, autoscalingPolicyProp)) {
		obj["autoscalingPolicy"] = autoscalingPolicyProp
	}
	targetProp, err := expandComputeRegionAutoscalerTarget(d.Get("target"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("target"); !tpgresource.IsEmptyValue(reflect.ValueOf(targetProp)) && (ok || !reflect.DeepEqual(v, targetProp)) {
		obj["target"] = targetProp
	}
	regionProp, err := expandComputeRegionAutoscalerRegion(d.Get("region"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("region"); !tpgresource.IsEmptyValue(reflect.ValueOf(regionProp)) && (ok || !reflect.DeepEqual(v, regionProp)) {
		obj["region"] = regionProp
	}

	return obj, nil
}

func expandComputeRegionAutoscalerName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionAutoscalerDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionAutoscalerAutoscalingPolicy(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedMinReplicas, err := expandComputeRegionAutoscalerAutoscalingPolicyMinReplicas(original["min_replicas"], d, config)
	if err != nil {
		return nil, err
	} else {
		transformed["minNumReplicas"] = transformedMinReplicas
	}

	transformedMaxReplicas, err := expandComputeRegionAutoscalerAutoscalingPolicyMaxReplicas(original["max_replicas"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxReplicas); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["maxNumReplicas"] = transformedMaxReplicas
	}

	transformedCooldownPeriod, err := expandComputeRegionAutoscalerAutoscalingPolicyCooldownPeriod(original["cooldown_period"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCooldownPeriod); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["coolDownPeriodSec"] = transformedCooldownPeriod
	}

	transformedMode, err := expandComputeRegionAutoscalerAutoscalingPolicyMode(original["mode"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMode); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["mode"] = transformedMode
	}

	transformedScaleDownControl, err := expandComputeRegionAutoscalerAutoscalingPolicyScaleDownControl(original["scale_down_control"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedScaleDownControl); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["scaleDownControl"] = transformedScaleDownControl
	}

	transformedScaleInControl, err := expandComputeRegionAutoscalerAutoscalingPolicyScaleInControl(original["scale_in_control"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedScaleInControl); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["scaleInControl"] = transformedScaleInControl
	}

	transformedCpuUtilization, err := expandComputeRegionAutoscalerAutoscalingPolicyCpuUtilization(original["cpu_utilization"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCpuUtilization); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["cpuUtilization"] = transformedCpuUtilization
	}

	transformedMetric, err := expandComputeRegionAutoscalerAutoscalingPolicyMetric(original["metric"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMetric); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["customMetricUtilizations"] = transformedMetric
	}

	transformedLoadBalancingUtilization, err := expandComputeRegionAutoscalerAutoscalingPolicyLoadBalancingUtilization(original["load_balancing_utilization"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLoadBalancingUtilization); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["loadBalancingUtilization"] = transformedLoadBalancingUtilization
	}

	transformedScalingSchedules, err := expandComputeRegionAutoscalerAutoscalingPolicyScalingSchedules(original["scaling_schedules"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedScalingSchedules); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["scalingSchedules"] = transformedScalingSchedules
	}

	return transformed, nil
}

func expandComputeRegionAutoscalerAutoscalingPolicyMinReplicas(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionAutoscalerAutoscalingPolicyMaxReplicas(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionAutoscalerAutoscalingPolicyCooldownPeriod(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionAutoscalerAutoscalingPolicyMode(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionAutoscalerAutoscalingPolicyScaleDownControl(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedMaxScaledDownReplicas, err := expandComputeRegionAutoscalerAutoscalingPolicyScaleDownControlMaxScaledDownReplicas(original["max_scaled_down_replicas"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxScaledDownReplicas); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["maxScaledDownReplicas"] = transformedMaxScaledDownReplicas
	}

	transformedTimeWindowSec, err := expandComputeRegionAutoscalerAutoscalingPolicyScaleDownControlTimeWindowSec(original["time_window_sec"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTimeWindowSec); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["timeWindowSec"] = transformedTimeWindowSec
	}

	return transformed, nil
}

func expandComputeRegionAutoscalerAutoscalingPolicyScaleDownControlMaxScaledDownReplicas(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedFixed, err := expandComputeRegionAutoscalerAutoscalingPolicyScaleDownControlMaxScaledDownReplicasFixed(original["fixed"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedFixed); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["fixed"] = transformedFixed
	}

	transformedPercent, err := expandComputeRegionAutoscalerAutoscalingPolicyScaleDownControlMaxScaledDownReplicasPercent(original["percent"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPercent); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["percent"] = transformedPercent
	}

	return transformed, nil
}

func expandComputeRegionAutoscalerAutoscalingPolicyScaleDownControlMaxScaledDownReplicasFixed(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionAutoscalerAutoscalingPolicyScaleDownControlMaxScaledDownReplicasPercent(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionAutoscalerAutoscalingPolicyScaleDownControlTimeWindowSec(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionAutoscalerAutoscalingPolicyScaleInControl(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedMaxScaledInReplicas, err := expandComputeRegionAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas(original["max_scaled_in_replicas"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxScaledInReplicas); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["maxScaledInReplicas"] = transformedMaxScaledInReplicas
	}

	transformedTimeWindowSec, err := expandComputeRegionAutoscalerAutoscalingPolicyScaleInControlTimeWindowSec(original["time_window_sec"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTimeWindowSec); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["timeWindowSec"] = transformedTimeWindowSec
	}

	return transformed, nil
}

func expandComputeRegionAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedFixed, err := expandComputeRegionAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicasFixed(original["fixed"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedFixed); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["fixed"] = transformedFixed
	}

	transformedPercent, err := expandComputeRegionAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicasPercent(original["percent"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPercent); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["percent"] = transformedPercent
	}

	return transformed, nil
}

func expandComputeRegionAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicasFixed(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicasPercent(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionAutoscalerAutoscalingPolicyScaleInControlTimeWindowSec(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionAutoscalerAutoscalingPolicyCpuUtilization(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedTarget, err := expandComputeRegionAutoscalerAutoscalingPolicyCpuUtilizationTarget(original["target"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTarget); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["utilizationTarget"] = transformedTarget
	}

	transformedPredictiveMethod, err := expandComputeRegionAutoscalerAutoscalingPolicyCpuUtilizationPredictiveMethod(original["predictive_method"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPredictiveMethod); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["predictiveMethod"] = transformedPredictiveMethod
	}

	return transformed, nil
}

func expandComputeRegionAutoscalerAutoscalingPolicyCpuUtilizationTarget(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionAutoscalerAutoscalingPolicyCpuUtilizationPredictiveMethod(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionAutoscalerAutoscalingPolicyMetric(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedName, err := expandComputeRegionAutoscalerAutoscalingPolicyMetricName(original["name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["metric"] = transformedName
		}

		transformedSingleInstanceAssignment, err := expandComputeRegionAutoscalerAutoscalingPolicyMetricSingleInstanceAssignment(original["single_instance_assignment"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedSingleInstanceAssignment); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["singleInstanceAssignment"] = transformedSingleInstanceAssignment
		}

		transformedTarget, err := expandComputeRegionAutoscalerAutoscalingPolicyMetricTarget(original["target"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedTarget); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["utilizationTarget"] = transformedTarget
		}

		transformedType, err := expandComputeRegionAutoscalerAutoscalingPolicyMetricType(original["type"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedType); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["utilizationTargetType"] = transformedType
		}

		transformedFilter, err := expandComputeRegionAutoscalerAutoscalingPolicyMetricFilter(original["filter"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedFilter); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["filter"] = transformedFilter
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandComputeRegionAutoscalerAutoscalingPolicyMetricName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionAutoscalerAutoscalingPolicyMetricSingleInstanceAssignment(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionAutoscalerAutoscalingPolicyMetricTarget(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionAutoscalerAutoscalingPolicyMetricType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionAutoscalerAutoscalingPolicyMetricFilter(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionAutoscalerAutoscalingPolicyLoadBalancingUtilization(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedTarget, err := expandComputeRegionAutoscalerAutoscalingPolicyLoadBalancingUtilizationTarget(original["target"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTarget); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["utilizationTarget"] = transformedTarget
	}

	return transformed, nil
}

func expandComputeRegionAutoscalerAutoscalingPolicyLoadBalancingUtilizationTarget(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionAutoscalerAutoscalingPolicyScalingSchedules(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	if v == nil {
		return map[string]interface{}{}, nil
	}
	m := make(map[string]interface{})
	for _, raw := range v.(*schema.Set).List() {
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedMinRequiredReplicas, err := expandComputeRegionAutoscalerAutoscalingPolicyScalingSchedulesMinRequiredReplicas(original["min_required_replicas"], d, config)
		if err != nil {
			return nil, err
		} else {
			transformed["minRequiredReplicas"] = transformedMinRequiredReplicas
		}

		transformedSchedule, err := expandComputeRegionAutoscalerAutoscalingPolicyScalingSchedulesSchedule(original["schedule"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedSchedule); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["schedule"] = transformedSchedule
		}

		transformedTimeZone, err := expandComputeRegionAutoscalerAutoscalingPolicyScalingSchedulesTimeZone(original["time_zone"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedTimeZone); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["timeZone"] = transformedTimeZone
		}

		transformedDurationSec, err := expandComputeRegionAutoscalerAutoscalingPolicyScalingSchedulesDurationSec(original["duration_sec"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDurationSec); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["durationSec"] = transformedDurationSec
		}

		transformedDisabled, err := expandComputeRegionAutoscalerAutoscalingPolicyScalingSchedulesDisabled(original["disabled"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDisabled); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["disabled"] = transformedDisabled
		}

		transformedDescription, err := expandComputeRegionAutoscalerAutoscalingPolicyScalingSchedulesDescription(original["description"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDescription); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["description"] = transformedDescription
		}

		transformedName, err := tpgresource.ExpandString(original["name"], d, config)
		if err != nil {
			return nil, err
		}
		m[transformedName] = transformed
	}
	return m, nil
}

func expandComputeRegionAutoscalerAutoscalingPolicyScalingSchedulesMinRequiredReplicas(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionAutoscalerAutoscalingPolicyScalingSchedulesSchedule(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionAutoscalerAutoscalingPolicyScalingSchedulesTimeZone(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionAutoscalerAutoscalingPolicyScalingSchedulesDurationSec(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionAutoscalerAutoscalingPolicyScalingSchedulesDisabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionAutoscalerAutoscalingPolicyScalingSchedulesDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionAutoscalerTarget(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionAutoscalerRegion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	f, err := tpgresource.ParseGlobalFieldValue("regions", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for region: %s", err)
	}
	return f.RelativeLink(), nil
}
