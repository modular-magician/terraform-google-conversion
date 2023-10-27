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

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/cai2hcl/common"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/caiasset"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const ComputeSnapshotAssetType string = "compute.googleapis.com/Snapshot"

const ComputeSnapshotAssetNameRegex string = "projects/(?P<project>[^/]+)/global/snapshots"

type ComputeSnapshotConverter struct {
	name   string
	schema map[string]*schema.Schema
}

func NewComputeSnapshotConverter(name string, schema map[string]*schema.Schema) common.Converter {
	return &ComputeSnapshotConverter{
		name:   name,
		schema: schema,
	}
}

func (c *ComputeSnapshotConverter) Convert(assets []*caiasset.Asset) ([]*common.HCLResourceBlock, error) {
	var blocks []*common.HCLResourceBlock
	config := common.NewConfig()

	for _, asset := range assets {
		if asset == nil {
			continue
		}
		if asset.Resource != nil && asset.Resource.Data != nil {
			block, err := c.convertResourceData(asset, config)
			if err != nil {
				return nil, err
			}
			blocks = append(blocks, block)
		}
	}
	return blocks, nil
}

func (c *ComputeSnapshotConverter) convertResourceData(asset *caiasset.Asset, config *transport_tpg.Config) (*common.HCLResourceBlock, error) {
	if asset == nil || asset.Resource == nil || asset.Resource.Data == nil {
		return nil, fmt.Errorf("asset resource data is nil")
	}

	assetResourceData := asset.Resource.Data

	hcl, _ := resourceComputeSnapshotRead(assetResourceData, config)

	ctyVal, err := common.MapToCtyValWithSchema(hcl, c.schema)
	if err != nil {
		return nil, err
	}

	resourceName := assetResourceData["name"].(string)

	return &common.HCLResourceBlock{
		Labels: []string{c.name, resourceName},
		Value:  ctyVal,
	}, nil
}

func resourceComputeSnapshotRead(resource map[string]interface{}, config *transport_tpg.Config) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	var resource_data *schema.ResourceData = nil

	result["creation_timestamp"] = flattenComputeSnapshotCreationTimestamp(resource["creationTimestamp"], resource_data, config)
	result["snapshot_id"] = flattenComputeSnapshotSnapshotId(resource["id"], resource_data, config)
	result["disk_size_gb"] = flattenComputeSnapshotDiskSizeGb(resource["diskSizeGb"], resource_data, config)
	result["chain_name"] = flattenComputeSnapshotChainName(resource["chainName"], resource_data, config)
	result["name"] = flattenComputeSnapshotName(resource["name"], resource_data, config)
	result["description"] = flattenComputeSnapshotDescription(resource["description"], resource_data, config)
	result["storage_bytes"] = flattenComputeSnapshotStorageBytes(resource["storageBytes"], resource_data, config)
	result["storage_locations"] = flattenComputeSnapshotStorageLocations(resource["storageLocations"], resource_data, config)
	result["licenses"] = flattenComputeSnapshotLicenses(resource["licenses"], resource_data, config)
	result["labels"] = flattenComputeSnapshotLabels(resource["labels"], resource_data, config)
	result["label_fingerprint"] = flattenComputeSnapshotLabelFingerprint(resource["labelFingerprint"], resource_data, config)
	result["terraform_labels"] = flattenComputeSnapshotTerraformLabels(resource["labels"], resource_data, config)
	result["effective_labels"] = flattenComputeSnapshotEffectiveLabels(resource["labels"], resource_data, config)
	result["source_disk"] = flattenComputeSnapshotSourceDisk(resource["sourceDisk"], resource_data, config)
	result["snapshot_encryption_key"] = flattenComputeSnapshotSnapshotEncryptionKey(resource["snapshotEncryptionKey"], resource_data, config)

	return result, nil
}

func flattenComputeSnapshotCreationTimestamp(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeSnapshotSnapshotId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := tpgresource.StringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenComputeSnapshotDiskSizeGb(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := tpgresource.StringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenComputeSnapshotChainName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeSnapshotName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeSnapshotDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeSnapshotStorageBytes(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := tpgresource.StringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenComputeSnapshotStorageLocations(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeSnapshotLicenses(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.ConvertAndMapStringArr(v.([]interface{}), tpgresource.ConvertSelfLinkToV1)
}

func flattenComputeSnapshotLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}

	transformed := make(map[string]interface{})
	if l, ok := d.GetOkExists("labels"); ok {
		for k := range l.(map[string]interface{}) {
			transformed[k] = v.(map[string]interface{})[k]
		}
	}

	return transformed
}

func flattenComputeSnapshotLabelFingerprint(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeSnapshotTerraformLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}

	transformed := make(map[string]interface{})
	if l, ok := d.GetOkExists("terraform_labels"); ok {
		for k := range l.(map[string]interface{}) {
			transformed[k] = v.(map[string]interface{})[k]
		}
	}

	return transformed
}

func flattenComputeSnapshotEffectiveLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeSnapshotSourceDisk(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.ConvertSelfLinkToV1(v.(string))
}

func flattenComputeSnapshotSnapshotEncryptionKey(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["raw_key"] =
		flattenComputeSnapshotSnapshotEncryptionKeyRawKey(original["rawKey"], d, config)
	transformed["sha256"] =
		flattenComputeSnapshotSnapshotEncryptionKeySha256(original["sha256"], d, config)
	transformed["kms_key_self_link"] =
		flattenComputeSnapshotSnapshotEncryptionKeyKmsKeySelfLink(original["kmsKeyName"], d, config)
	transformed["kms_key_service_account"] =
		flattenComputeSnapshotSnapshotEncryptionKeyKmsKeyServiceAccount(original["kmsKeyServiceAccount"], d, config)
	return []interface{}{transformed}
}
func flattenComputeSnapshotSnapshotEncryptionKeyRawKey(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return d.Get("snapshot_encryption_key.0.raw_key")
}

func flattenComputeSnapshotSnapshotEncryptionKeySha256(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeSnapshotSnapshotEncryptionKeyKmsKeySelfLink(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeSnapshotSnapshotEncryptionKeyKmsKeyServiceAccount(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}
