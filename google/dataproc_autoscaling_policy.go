// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Nathan is editing this to generate diffs in lots of files.
//     He won't submit this change.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import "reflect"

func GetDataprocAutoscalingPolicyCaiObject(d TerraformResourceData, config *Config) (Asset, error) {
	name, err := assetName(d, config, "//dataproc.googleapis.com/projects/{{project}}/locations/{{location}}/autoscalingPolicies/{{policy_id}}")
	if err != nil {
		return Asset{}, err
	}
	if obj, err := GetDataprocAutoscalingPolicyApiObject(d, config); err == nil {
		return Asset{
			Name: name,
			Type: "dataproc.googleapis.com/AutoscalingPolicy",
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/dataproc/v1/rest",
				DiscoveryName:        "AutoscalingPolicy",
				Data:                 obj,
			},
		}, nil
	} else {
		return Asset{}, err
	}
}

func GetDataprocAutoscalingPolicyApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	idProp, err := expandDataprocAutoscalingPolicyPolicyId(d.Get("policy_id"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("policy_id"); !isEmptyValue(reflect.ValueOf(idProp)) && (ok || !reflect.DeepEqual(v, idProp)) {
		obj["id"] = idProp
	}
	workerConfigProp, err := expandDataprocAutoscalingPolicyWorkerConfig(d.Get("worker_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("worker_config"); !isEmptyValue(reflect.ValueOf(workerConfigProp)) && (ok || !reflect.DeepEqual(v, workerConfigProp)) {
		obj["workerConfig"] = workerConfigProp
	}
	secondaryWorkerConfigProp, err := expandDataprocAutoscalingPolicySecondaryWorkerConfig(d.Get("secondary_worker_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("secondary_worker_config"); !isEmptyValue(reflect.ValueOf(secondaryWorkerConfigProp)) && (ok || !reflect.DeepEqual(v, secondaryWorkerConfigProp)) {
		obj["secondaryWorkerConfig"] = secondaryWorkerConfigProp
	}
	basicAlgorithmProp, err := expandDataprocAutoscalingPolicyBasicAlgorithm(d.Get("basic_algorithm"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("basic_algorithm"); !isEmptyValue(reflect.ValueOf(basicAlgorithmProp)) && (ok || !reflect.DeepEqual(v, basicAlgorithmProp)) {
		obj["basicAlgorithm"] = basicAlgorithmProp
	}

	return obj, nil
}

func expandDataprocAutoscalingPolicyPolicyId(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataprocAutoscalingPolicyWorkerConfig(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedMinInstances, err := expandDataprocAutoscalingPolicyWorkerConfigMinInstances(original["min_instances"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMinInstances); val.IsValid() && !isEmptyValue(val) {
		transformed["minInstances"] = transformedMinInstances
	}

	transformedMaxInstances, err := expandDataprocAutoscalingPolicyWorkerConfigMaxInstances(original["max_instances"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxInstances); val.IsValid() && !isEmptyValue(val) {
		transformed["maxInstances"] = transformedMaxInstances
	}

	transformedWeight, err := expandDataprocAutoscalingPolicyWorkerConfigWeight(original["weight"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedWeight); val.IsValid() && !isEmptyValue(val) {
		transformed["weight"] = transformedWeight
	}

	return transformed, nil
}

func expandDataprocAutoscalingPolicyWorkerConfigMinInstances(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataprocAutoscalingPolicyWorkerConfigMaxInstances(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataprocAutoscalingPolicyWorkerConfigWeight(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataprocAutoscalingPolicySecondaryWorkerConfig(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedMinInstances, err := expandDataprocAutoscalingPolicySecondaryWorkerConfigMinInstances(original["min_instances"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMinInstances); val.IsValid() && !isEmptyValue(val) {
		transformed["minInstances"] = transformedMinInstances
	}

	transformedMaxInstances, err := expandDataprocAutoscalingPolicySecondaryWorkerConfigMaxInstances(original["max_instances"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxInstances); val.IsValid() && !isEmptyValue(val) {
		transformed["maxInstances"] = transformedMaxInstances
	}

	transformedWeight, err := expandDataprocAutoscalingPolicySecondaryWorkerConfigWeight(original["weight"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedWeight); val.IsValid() && !isEmptyValue(val) {
		transformed["weight"] = transformedWeight
	}

	return transformed, nil
}

func expandDataprocAutoscalingPolicySecondaryWorkerConfigMinInstances(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataprocAutoscalingPolicySecondaryWorkerConfigMaxInstances(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataprocAutoscalingPolicySecondaryWorkerConfigWeight(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataprocAutoscalingPolicyBasicAlgorithm(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedCooldownPeriod, err := expandDataprocAutoscalingPolicyBasicAlgorithmCooldownPeriod(original["cooldown_period"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCooldownPeriod); val.IsValid() && !isEmptyValue(val) {
		transformed["cooldownPeriod"] = transformedCooldownPeriod
	}

	transformedYarnConfig, err := expandDataprocAutoscalingPolicyBasicAlgorithmYarnConfig(original["yarn_config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedYarnConfig); val.IsValid() && !isEmptyValue(val) {
		transformed["yarnConfig"] = transformedYarnConfig
	}

	return transformed, nil
}

func expandDataprocAutoscalingPolicyBasicAlgorithmCooldownPeriod(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataprocAutoscalingPolicyBasicAlgorithmYarnConfig(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedGracefulDecommissionTimeout, err := expandDataprocAutoscalingPolicyBasicAlgorithmYarnConfigGracefulDecommissionTimeout(original["graceful_decommission_timeout"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedGracefulDecommissionTimeout); val.IsValid() && !isEmptyValue(val) {
		transformed["gracefulDecommissionTimeout"] = transformedGracefulDecommissionTimeout
	}

	transformedScaleUpFactor, err := expandDataprocAutoscalingPolicyBasicAlgorithmYarnConfigScaleUpFactor(original["scale_up_factor"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedScaleUpFactor); val.IsValid() && !isEmptyValue(val) {
		transformed["scaleUpFactor"] = transformedScaleUpFactor
	}

	transformedScaleDownFactor, err := expandDataprocAutoscalingPolicyBasicAlgorithmYarnConfigScaleDownFactor(original["scale_down_factor"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedScaleDownFactor); val.IsValid() && !isEmptyValue(val) {
		transformed["scaleDownFactor"] = transformedScaleDownFactor
	}

	transformedScaleUpMinWorkerFraction, err := expandDataprocAutoscalingPolicyBasicAlgorithmYarnConfigScaleUpMinWorkerFraction(original["scale_up_min_worker_fraction"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedScaleUpMinWorkerFraction); val.IsValid() && !isEmptyValue(val) {
		transformed["scaleUpMinWorkerFraction"] = transformedScaleUpMinWorkerFraction
	}

	transformedScaleDownMinWorkerFraction, err := expandDataprocAutoscalingPolicyBasicAlgorithmYarnConfigScaleDownMinWorkerFraction(original["scale_down_min_worker_fraction"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedScaleDownMinWorkerFraction); val.IsValid() && !isEmptyValue(val) {
		transformed["scaleDownMinWorkerFraction"] = transformedScaleDownMinWorkerFraction
	}

	return transformed, nil
}

func expandDataprocAutoscalingPolicyBasicAlgorithmYarnConfigGracefulDecommissionTimeout(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataprocAutoscalingPolicyBasicAlgorithmYarnConfigScaleUpFactor(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataprocAutoscalingPolicyBasicAlgorithmYarnConfigScaleDownFactor(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataprocAutoscalingPolicyBasicAlgorithmYarnConfigScaleUpMinWorkerFraction(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataprocAutoscalingPolicyBasicAlgorithmYarnConfigScaleDownMinWorkerFraction(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
