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

package vmwareengine

import (
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgresource"
	transport_tpg "github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/transport"
)

const VmwareengineClusterAssetType string = "vmwareengine.googleapis.com/Cluster"

func ResourceConverterVmwareengineCluster() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType: VmwareengineClusterAssetType,
		Convert:   GetVmwareengineClusterCaiObject,
	}
}

func GetVmwareengineClusterCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	name, err := tpgresource.AssetName(d, config, "//vmwareengine.googleapis.com/{{parent}}/clusters/{{name}}")
	if err != nil {
		return []tpgresource.Asset{}, err
	}
	if obj, err := GetVmwareengineClusterApiObject(d, config); err == nil {
		return []tpgresource.Asset{{
			Name: name,
			Type: VmwareengineClusterAssetType,
			Resource: &tpgresource.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/vmwareengine/v1/rest",
				DiscoveryName:        "Cluster",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []tpgresource.Asset{}, err
	}
}

func GetVmwareengineClusterApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nodeTypeConfigsProp, err := expandVmwareengineClusterNodeTypeConfigs(d.Get("node_type_configs"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("node_type_configs"); !tpgresource.IsEmptyValue(reflect.ValueOf(nodeTypeConfigsProp)) && (ok || !reflect.DeepEqual(v, nodeTypeConfigsProp)) {
		obj["nodeTypeConfigs"] = nodeTypeConfigsProp
	}

	return obj, nil
}

func expandVmwareengineClusterNodeTypeConfigs(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	if v == nil {
		return map[string]interface{}{}, nil
	}
	m := make(map[string]interface{})
	for _, raw := range v.(*schema.Set).List() {
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedNodeCount, err := expandVmwareengineClusterNodeTypeConfigsNodeCount(original["node_count"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedNodeCount); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["nodeCount"] = transformedNodeCount
		}

		transformedCustomCoreCount, err := expandVmwareengineClusterNodeTypeConfigsCustomCoreCount(original["custom_core_count"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedCustomCoreCount); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["customCoreCount"] = transformedCustomCoreCount
		}

		transformedNodeTypeId, err := tpgresource.ExpandString(original["node_type_id"], d, config)
		if err != nil {
			return nil, err
		}
		m[transformedNodeTypeId] = transformed
	}
	return m, nil
}

func expandVmwareengineClusterNodeTypeConfigsNodeCount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVmwareengineClusterNodeTypeConfigsCustomCoreCount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
