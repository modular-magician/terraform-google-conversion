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

const GkeonpremVmwareNodePoolAssetType string = "gkeonprem.googleapis.com/VmwareNodePool"

func ResourceConverterGkeonpremVmwareNodePool() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: GkeonpremVmwareNodePoolAssetType,
		Convert:   GetGkeonpremVmwareNodePoolCaiObject,
	}
}

func GetGkeonpremVmwareNodePoolCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//gkeonprem.googleapis.com/projects/{{project}}/locations/{{location}}/vmwareClusters/{{vmware_cluster}}/vmwareNodePools/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetGkeonpremVmwareNodePoolApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: GkeonpremVmwareNodePoolAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/gkeonprem/v1/rest",
				DiscoveryName:        "VmwareNodePool",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetGkeonpremVmwareNodePoolApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	displayNameProp, err := expandGkeonpremVmwareNodePoolDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	nodePoolAutoscalingProp, err := expandGkeonpremVmwareNodePoolNodePoolAutoscaling(d.Get("node_pool_autoscaling"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("node_pool_autoscaling"); !tpgresource.IsEmptyValue(reflect.ValueOf(nodePoolAutoscalingProp)) && (ok || !reflect.DeepEqual(v, nodePoolAutoscalingProp)) {
		obj["nodePoolAutoscaling"] = nodePoolAutoscalingProp
	}
	configProp, err := expandGkeonpremVmwareNodePoolConfig(d.Get("config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("config"); !tpgresource.IsEmptyValue(reflect.ValueOf(configProp)) && (ok || !reflect.DeepEqual(v, configProp)) {
		obj["config"] = configProp
	}
	annotationsProp, err := expandGkeonpremVmwareNodePoolEffectiveAnnotations(d.Get("effective_annotations"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("effective_annotations"); !tpgresource.IsEmptyValue(reflect.ValueOf(annotationsProp)) && (ok || !reflect.DeepEqual(v, annotationsProp)) {
		obj["annotations"] = annotationsProp
	}

	return obj, nil
}

func expandGkeonpremVmwareNodePoolDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGkeonpremVmwareNodePoolNodePoolAutoscaling(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedMinReplicas, err := expandGkeonpremVmwareNodePoolNodePoolAutoscalingMinReplicas(original["min_replicas"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMinReplicas); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["minReplicas"] = transformedMinReplicas
	}

	transformedMaxReplicas, err := expandGkeonpremVmwareNodePoolNodePoolAutoscalingMaxReplicas(original["max_replicas"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxReplicas); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["maxReplicas"] = transformedMaxReplicas
	}

	return transformed, nil
}

func expandGkeonpremVmwareNodePoolNodePoolAutoscalingMinReplicas(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGkeonpremVmwareNodePoolNodePoolAutoscalingMaxReplicas(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGkeonpremVmwareNodePoolConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedCpus, err := expandGkeonpremVmwareNodePoolConfigCpus(original["cpus"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCpus); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["cpus"] = transformedCpus
	}

	transformedMemoryMb, err := expandGkeonpremVmwareNodePoolConfigMemoryMb(original["memory_mb"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMemoryMb); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["memoryMb"] = transformedMemoryMb
	}

	transformedReplicas, err := expandGkeonpremVmwareNodePoolConfigReplicas(original["replicas"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedReplicas); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["replicas"] = transformedReplicas
	}

	transformedImageType, err := expandGkeonpremVmwareNodePoolConfigImageType(original["image_type"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedImageType); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["imageType"] = transformedImageType
	}

	transformedImage, err := expandGkeonpremVmwareNodePoolConfigImage(original["image"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedImage); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["image"] = transformedImage
	}

	transformedBootDiskSizeGb, err := expandGkeonpremVmwareNodePoolConfigBootDiskSizeGb(original["boot_disk_size_gb"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedBootDiskSizeGb); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["bootDiskSizeGb"] = transformedBootDiskSizeGb
	}

	transformedTaints, err := expandGkeonpremVmwareNodePoolConfigTaints(original["taints"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTaints); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["taints"] = transformedTaints
	}

	transformedLabels, err := expandGkeonpremVmwareNodePoolConfigLabels(original["labels"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLabels); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["labels"] = transformedLabels
	}

	transformedVsphereConfig, err := expandGkeonpremVmwareNodePoolConfigVsphereConfig(original["vsphere_config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedVsphereConfig); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["vsphereConfig"] = transformedVsphereConfig
	}

	transformedEnableLoadBalancer, err := expandGkeonpremVmwareNodePoolConfigEnableLoadBalancer(original["enable_load_balancer"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEnableLoadBalancer); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["enableLoadBalancer"] = transformedEnableLoadBalancer
	}

	return transformed, nil
}

func expandGkeonpremVmwareNodePoolConfigCpus(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGkeonpremVmwareNodePoolConfigMemoryMb(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGkeonpremVmwareNodePoolConfigReplicas(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGkeonpremVmwareNodePoolConfigImageType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGkeonpremVmwareNodePoolConfigImage(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGkeonpremVmwareNodePoolConfigBootDiskSizeGb(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGkeonpremVmwareNodePoolConfigTaints(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedKey, err := expandGkeonpremVmwareNodePoolConfigTaintsKey(original["key"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedKey); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["key"] = transformedKey
		}

		transformedValue, err := expandGkeonpremVmwareNodePoolConfigTaintsValue(original["value"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedValue); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["value"] = transformedValue
		}

		transformedEffect, err := expandGkeonpremVmwareNodePoolConfigTaintsEffect(original["effect"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedEffect); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["effect"] = transformedEffect
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandGkeonpremVmwareNodePoolConfigTaintsKey(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGkeonpremVmwareNodePoolConfigTaintsValue(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGkeonpremVmwareNodePoolConfigTaintsEffect(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGkeonpremVmwareNodePoolConfigLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandGkeonpremVmwareNodePoolConfigVsphereConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedDatastore, err := expandGkeonpremVmwareNodePoolConfigVsphereConfigDatastore(original["datastore"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDatastore); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["datastore"] = transformedDatastore
	}

	transformedTags, err := expandGkeonpremVmwareNodePoolConfigVsphereConfigTags(original["tags"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTags); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["tags"] = transformedTags
	}

	transformedHostGroups, err := expandGkeonpremVmwareNodePoolConfigVsphereConfigHostGroups(original["host_groups"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedHostGroups); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["hostGroups"] = transformedHostGroups
	}

	return transformed, nil
}

func expandGkeonpremVmwareNodePoolConfigVsphereConfigDatastore(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGkeonpremVmwareNodePoolConfigVsphereConfigTags(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedCategory, err := expandGkeonpremVmwareNodePoolConfigVsphereConfigTagsCategory(original["category"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedCategory); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["category"] = transformedCategory
		}

		transformedTag, err := expandGkeonpremVmwareNodePoolConfigVsphereConfigTagsTag(original["tag"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedTag); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["tag"] = transformedTag
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandGkeonpremVmwareNodePoolConfigVsphereConfigTagsCategory(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGkeonpremVmwareNodePoolConfigVsphereConfigTagsTag(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGkeonpremVmwareNodePoolConfigVsphereConfigHostGroups(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGkeonpremVmwareNodePoolConfigEnableLoadBalancer(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGkeonpremVmwareNodePoolEffectiveAnnotations(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
