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

package networkservices

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const NetworkServicesServiceLbPoliciesAssetType string = "networkservices.googleapis.com/ServiceLbPolicies"

func ResourceConverterNetworkServicesServiceLbPolicies() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: NetworkServicesServiceLbPoliciesAssetType,
		Convert:   GetNetworkServicesServiceLbPoliciesCaiObject,
	}
}

func GetNetworkServicesServiceLbPoliciesCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//networkservices.googleapis.com/projects/{{project}}/locations/{{location}}/serviceLbPolicies/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetNetworkServicesServiceLbPoliciesApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: NetworkServicesServiceLbPoliciesAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/networkservices/v1/rest",
				DiscoveryName:        "ServiceLbPolicies",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetNetworkServicesServiceLbPoliciesApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	descriptionProp, err := expandNetworkServicesServiceLbPoliciesDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	loadBalancingAlgorithmProp, err := expandNetworkServicesServiceLbPoliciesLoadBalancingAlgorithm(d.Get("load_balancing_algorithm"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("load_balancing_algorithm"); !tpgresource.IsEmptyValue(reflect.ValueOf(loadBalancingAlgorithmProp)) && (ok || !reflect.DeepEqual(v, loadBalancingAlgorithmProp)) {
		obj["loadBalancingAlgorithm"] = loadBalancingAlgorithmProp
	}
	autoCapacityDrainProp, err := expandNetworkServicesServiceLbPoliciesAutoCapacityDrain(d.Get("auto_capacity_drain"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("auto_capacity_drain"); !tpgresource.IsEmptyValue(reflect.ValueOf(autoCapacityDrainProp)) && (ok || !reflect.DeepEqual(v, autoCapacityDrainProp)) {
		obj["autoCapacityDrain"] = autoCapacityDrainProp
	}
	failoverConfigProp, err := expandNetworkServicesServiceLbPoliciesFailoverConfig(d.Get("failover_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("failover_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(failoverConfigProp)) && (ok || !reflect.DeepEqual(v, failoverConfigProp)) {
		obj["failoverConfig"] = failoverConfigProp
	}
	labelsProp, err := expandNetworkServicesServiceLbPoliciesEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	return obj, nil
}

func expandNetworkServicesServiceLbPoliciesDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesServiceLbPoliciesLoadBalancingAlgorithm(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesServiceLbPoliciesAutoCapacityDrain(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedEnable, err := expandNetworkServicesServiceLbPoliciesAutoCapacityDrainEnable(original["enable"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEnable); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["enable"] = transformedEnable
	}

	return transformed, nil
}

func expandNetworkServicesServiceLbPoliciesAutoCapacityDrainEnable(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesServiceLbPoliciesFailoverConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedFailoverHealthThreshold, err := expandNetworkServicesServiceLbPoliciesFailoverConfigFailoverHealthThreshold(original["failover_health_threshold"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedFailoverHealthThreshold); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["failoverHealthThreshold"] = transformedFailoverHealthThreshold
	}

	return transformed, nil
}

func expandNetworkServicesServiceLbPoliciesFailoverConfigFailoverHealthThreshold(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesServiceLbPoliciesEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
