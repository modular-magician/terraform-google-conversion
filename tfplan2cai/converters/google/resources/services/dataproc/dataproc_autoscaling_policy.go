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

package dataproc

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const DataprocAutoscalingPolicyAssetType string = "dataproc.googleapis.com/AutoscalingPolicy"

func ResourceConverterDataprocAutoscalingPolicy() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType: DataprocAutoscalingPolicyAssetType,
		Convert:   GetDataprocAutoscalingPolicyCaiObject,
	}
}

func GetDataprocAutoscalingPolicyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	name, err := tpgresource.AssetName(d, config, "//dataproc.googleapis.com/projects/{{project}}/locations/{{location}}/autoscalingPolicies/{{policy_id}}")
	if err != nil {
		return []tpgresource.Asset{}, err
	}
	if obj, err := GetDataprocAutoscalingPolicyApiObject(d, config); err == nil {
		return []tpgresource.Asset{{
			Name: name,
			Type: DataprocAutoscalingPolicyAssetType,
			Resource: &tpgresource.AssetResource{
				Version:              "v1beta2",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/dataproc/v1beta2/rest",
				DiscoveryName:        "AutoscalingPolicy",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []tpgresource.Asset{}, err
	}
}

func GetDataprocAutoscalingPolicyApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	idProp, err := expandDataprocAutoscalingPolicyPolicyId(d.Get("policy_id"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("policy_id"); !tpgresource.IsEmptyValue(reflect.ValueOf(idProp)) && (ok || !reflect.DeepEqual(v, idProp)) {
		obj["id"] = idProp
	}
	workerConfigProp, err := expandDataprocAutoscalingPolicyWorkerConfig(d.Get("worker_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("worker_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(workerConfigProp)) && (ok || !reflect.DeepEqual(v, workerConfigProp)) {
		obj["workerConfig"] = workerConfigProp
	}
	secondaryWorkerConfigProp, err := expandDataprocAutoscalingPolicySecondaryWorkerConfig(d.Get("secondary_worker_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("secondary_worker_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(secondaryWorkerConfigProp)) && (ok || !reflect.DeepEqual(v, secondaryWorkerConfigProp)) {
		obj["secondaryWorkerConfig"] = secondaryWorkerConfigProp
	}
	basicAlgorithmProp, err := expandDataprocAutoscalingPolicyBasicAlgorithm(d.Get("basic_algorithm"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("basic_algorithm"); !tpgresource.IsEmptyValue(reflect.ValueOf(basicAlgorithmProp)) && (ok || !reflect.DeepEqual(v, basicAlgorithmProp)) {
		obj["basicAlgorithm"] = basicAlgorithmProp
	}

	return obj, nil
}

func expandDataprocAutoscalingPolicyPolicyId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocAutoscalingPolicyWorkerConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
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
	} else if val := reflect.ValueOf(transformedMinInstances); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["minInstances"] = transformedMinInstances
	}

	transformedMaxInstances, err := expandDataprocAutoscalingPolicyWorkerConfigMaxInstances(original["max_instances"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxInstances); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["maxInstances"] = transformedMaxInstances
	}

	transformedWeight, err := expandDataprocAutoscalingPolicyWorkerConfigWeight(original["weight"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedWeight); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["weight"] = transformedWeight
	}

	return transformed, nil
}

func expandDataprocAutoscalingPolicyWorkerConfigMinInstances(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocAutoscalingPolicyWorkerConfigMaxInstances(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocAutoscalingPolicyWorkerConfigWeight(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocAutoscalingPolicySecondaryWorkerConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
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
	} else if val := reflect.ValueOf(transformedMinInstances); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["minInstances"] = transformedMinInstances
	}

	transformedMaxInstances, err := expandDataprocAutoscalingPolicySecondaryWorkerConfigMaxInstances(original["max_instances"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxInstances); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["maxInstances"] = transformedMaxInstances
	}

	transformedWeight, err := expandDataprocAutoscalingPolicySecondaryWorkerConfigWeight(original["weight"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedWeight); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["weight"] = transformedWeight
	}

	return transformed, nil
}

func expandDataprocAutoscalingPolicySecondaryWorkerConfigMinInstances(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocAutoscalingPolicySecondaryWorkerConfigMaxInstances(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocAutoscalingPolicySecondaryWorkerConfigWeight(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocAutoscalingPolicyBasicAlgorithm(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
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
	} else if val := reflect.ValueOf(transformedCooldownPeriod); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["cooldownPeriod"] = transformedCooldownPeriod
	}

	transformedYarnConfig, err := expandDataprocAutoscalingPolicyBasicAlgorithmYarnConfig(original["yarn_config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedYarnConfig); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["yarnConfig"] = transformedYarnConfig
	}

	return transformed, nil
}

func expandDataprocAutoscalingPolicyBasicAlgorithmCooldownPeriod(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocAutoscalingPolicyBasicAlgorithmYarnConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
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
	} else if val := reflect.ValueOf(transformedGracefulDecommissionTimeout); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["gracefulDecommissionTimeout"] = transformedGracefulDecommissionTimeout
	}

	transformedScaleUpFactor, err := expandDataprocAutoscalingPolicyBasicAlgorithmYarnConfigScaleUpFactor(original["scale_up_factor"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedScaleUpFactor); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["scaleUpFactor"] = transformedScaleUpFactor
	}

	transformedScaleDownFactor, err := expandDataprocAutoscalingPolicyBasicAlgorithmYarnConfigScaleDownFactor(original["scale_down_factor"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedScaleDownFactor); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["scaleDownFactor"] = transformedScaleDownFactor
	}

	transformedScaleUpMinWorkerFraction, err := expandDataprocAutoscalingPolicyBasicAlgorithmYarnConfigScaleUpMinWorkerFraction(original["scale_up_min_worker_fraction"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedScaleUpMinWorkerFraction); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["scaleUpMinWorkerFraction"] = transformedScaleUpMinWorkerFraction
	}

	transformedScaleDownMinWorkerFraction, err := expandDataprocAutoscalingPolicyBasicAlgorithmYarnConfigScaleDownMinWorkerFraction(original["scale_down_min_worker_fraction"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedScaleDownMinWorkerFraction); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["scaleDownMinWorkerFraction"] = transformedScaleDownMinWorkerFraction
	}

	return transformed, nil
}

func expandDataprocAutoscalingPolicyBasicAlgorithmYarnConfigGracefulDecommissionTimeout(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocAutoscalingPolicyBasicAlgorithmYarnConfigScaleUpFactor(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocAutoscalingPolicyBasicAlgorithmYarnConfigScaleDownFactor(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocAutoscalingPolicyBasicAlgorithmYarnConfigScaleUpMinWorkerFraction(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocAutoscalingPolicyBasicAlgorithmYarnConfigScaleDownMinWorkerFraction(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
