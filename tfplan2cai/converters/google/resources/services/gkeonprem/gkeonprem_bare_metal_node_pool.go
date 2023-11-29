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

package gkeonprem

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const GkeonpremBareMetalNodePoolAssetType string = "gkeonprem.googleapis.com/BareMetalNodePool"

func ResourceConverterGkeonpremBareMetalNodePool() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: GkeonpremBareMetalNodePoolAssetType,
		Convert:   GetGkeonpremBareMetalNodePoolCaiObject,
	}
}

func GetGkeonpremBareMetalNodePoolCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//gkeonprem.googleapis.com/projects/{{project}}/locations/{{location}}/bareMetalClusters/{{bare_metal_cluster}}/bareMetalNodePools/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetGkeonpremBareMetalNodePoolApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: GkeonpremBareMetalNodePoolAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/gkeonprem/v1/rest",
				DiscoveryName:        "BareMetalNodePool",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetGkeonpremBareMetalNodePoolApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	displayNameProp, err := expandGkeonpremBareMetalNodePoolDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	nodePoolConfigProp, err := expandGkeonpremBareMetalNodePoolNodePoolConfig(d.Get("node_pool_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("node_pool_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(nodePoolConfigProp)) && (ok || !reflect.DeepEqual(v, nodePoolConfigProp)) {
		obj["nodePoolConfig"] = nodePoolConfigProp
	}
	annotationsProp, err := expandGkeonpremBareMetalNodePoolEffectiveAnnotations(d.Get("effective_annotations"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("effective_annotations"); !tpgresource.IsEmptyValue(reflect.ValueOf(annotationsProp)) && (ok || !reflect.DeepEqual(v, annotationsProp)) {
		obj["annotations"] = annotationsProp
	}

	return obj, nil
}

func expandGkeonpremBareMetalNodePoolDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGkeonpremBareMetalNodePoolNodePoolConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedNodeConfigs, err := expandGkeonpremBareMetalNodePoolNodePoolConfigNodeConfigs(original["node_configs"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNodeConfigs); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["nodeConfigs"] = transformedNodeConfigs
	}

	transformedOperatingSystem, err := expandGkeonpremBareMetalNodePoolNodePoolConfigOperatingSystem(original["operating_system"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedOperatingSystem); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["operatingSystem"] = transformedOperatingSystem
	}

	transformedTaints, err := expandGkeonpremBareMetalNodePoolNodePoolConfigTaints(original["taints"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTaints); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["taints"] = transformedTaints
	}

	transformedLabels, err := expandGkeonpremBareMetalNodePoolNodePoolConfigLabels(original["labels"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLabels); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["labels"] = transformedLabels
	}

	return transformed, nil
}

func expandGkeonpremBareMetalNodePoolNodePoolConfigNodeConfigs(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedNodeIp, err := expandGkeonpremBareMetalNodePoolNodePoolConfigNodeConfigsNodeIp(original["node_ip"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedNodeIp); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["nodeIp"] = transformedNodeIp
		}

		transformedLabels, err := expandGkeonpremBareMetalNodePoolNodePoolConfigNodeConfigsLabels(original["labels"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedLabels); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["labels"] = transformedLabels
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandGkeonpremBareMetalNodePoolNodePoolConfigNodeConfigsNodeIp(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGkeonpremBareMetalNodePoolNodePoolConfigNodeConfigsLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandGkeonpremBareMetalNodePoolNodePoolConfigOperatingSystem(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGkeonpremBareMetalNodePoolNodePoolConfigTaints(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedKey, err := expandGkeonpremBareMetalNodePoolNodePoolConfigTaintsKey(original["key"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedKey); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["key"] = transformedKey
		}

		transformedValue, err := expandGkeonpremBareMetalNodePoolNodePoolConfigTaintsValue(original["value"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedValue); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["value"] = transformedValue
		}

		transformedEffect, err := expandGkeonpremBareMetalNodePoolNodePoolConfigTaintsEffect(original["effect"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedEffect); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["effect"] = transformedEffect
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandGkeonpremBareMetalNodePoolNodePoolConfigTaintsKey(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGkeonpremBareMetalNodePoolNodePoolConfigTaintsValue(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGkeonpremBareMetalNodePoolNodePoolConfigTaintsEffect(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGkeonpremBareMetalNodePoolNodePoolConfigLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandGkeonpremBareMetalNodePoolEffectiveAnnotations(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
