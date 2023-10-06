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

package edgecontainer

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const EdgecontainerNodePoolAssetType string = "edgecontainer.googleapis.com/NodePool"

func ResourceConverterEdgecontainerNodePool() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: EdgecontainerNodePoolAssetType,
		Convert:   GetEdgecontainerNodePoolCaiObject,
	}
}

func GetEdgecontainerNodePoolCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//edgecontainer.googleapis.com/projects/{{project}}/locations/{{location}}/clusters/{{cluster}}/nodePools/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetEdgecontainerNodePoolApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: EdgecontainerNodePoolAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/edgecontainer/v1/rest",
				DiscoveryName:        "NodePool",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetEdgecontainerNodePoolApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nodeLocationProp, err := expandEdgecontainerNodePoolNodeLocation(d.Get("node_location"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("node_location"); !tpgresource.IsEmptyValue(reflect.ValueOf(nodeLocationProp)) && (ok || !reflect.DeepEqual(v, nodeLocationProp)) {
		obj["nodeLocation"] = nodeLocationProp
	}
	nodeCountProp, err := expandEdgecontainerNodePoolNodeCount(d.Get("node_count"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("node_count"); !tpgresource.IsEmptyValue(reflect.ValueOf(nodeCountProp)) && (ok || !reflect.DeepEqual(v, nodeCountProp)) {
		obj["nodeCount"] = nodeCountProp
	}
	machineFilterProp, err := expandEdgecontainerNodePoolMachineFilter(d.Get("machine_filter"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("machine_filter"); !tpgresource.IsEmptyValue(reflect.ValueOf(machineFilterProp)) && (ok || !reflect.DeepEqual(v, machineFilterProp)) {
		obj["machineFilter"] = machineFilterProp
	}
	localDiskEncryptionProp, err := expandEdgecontainerNodePoolLocalDiskEncryption(d.Get("local_disk_encryption"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("local_disk_encryption"); !tpgresource.IsEmptyValue(reflect.ValueOf(localDiskEncryptionProp)) && (ok || !reflect.DeepEqual(v, localDiskEncryptionProp)) {
		obj["localDiskEncryption"] = localDiskEncryptionProp
	}
	nodeConfigProp, err := expandEdgecontainerNodePoolNodeConfig(d.Get("node_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("node_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(nodeConfigProp)) && (ok || !reflect.DeepEqual(v, nodeConfigProp)) {
		obj["nodeConfig"] = nodeConfigProp
	}
	labelsProp, err := expandEdgecontainerNodePoolEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	return obj, nil
}

func expandEdgecontainerNodePoolNodeLocation(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandEdgecontainerNodePoolNodeCount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandEdgecontainerNodePoolMachineFilter(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandEdgecontainerNodePoolLocalDiskEncryption(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedKmsKey, err := expandEdgecontainerNodePoolLocalDiskEncryptionKmsKey(original["kms_key"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKmsKey); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["kmsKey"] = transformedKmsKey
	}

	transformedKmsKeyActiveVersion, err := expandEdgecontainerNodePoolLocalDiskEncryptionKmsKeyActiveVersion(original["kms_key_active_version"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKmsKeyActiveVersion); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["kmsKeyActiveVersion"] = transformedKmsKeyActiveVersion
	}

	transformedKmsKeyState, err := expandEdgecontainerNodePoolLocalDiskEncryptionKmsKeyState(original["kms_key_state"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKmsKeyState); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["kmsKeyState"] = transformedKmsKeyState
	}

	return transformed, nil
}

func expandEdgecontainerNodePoolLocalDiskEncryptionKmsKey(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandEdgecontainerNodePoolLocalDiskEncryptionKmsKeyActiveVersion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandEdgecontainerNodePoolLocalDiskEncryptionKmsKeyState(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandEdgecontainerNodePoolNodeConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedLabels, err := expandEdgecontainerNodePoolNodeConfigLabels(original["labels"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLabels); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["labels"] = transformedLabels
	}

	return transformed, nil
}

func expandEdgecontainerNodePoolNodeConfigLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandEdgecontainerNodePoolEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
