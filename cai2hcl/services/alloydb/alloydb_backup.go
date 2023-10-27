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

package alloydb

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/cai2hcl/common"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/caiasset"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const AlloydbBackupAssetType string = "alloydb.googleapis.com/Backup"

const AlloydbBackupAssetNameRegex string = "projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/backups"

type AlloydbBackupConverter struct {
	name   string
	schema map[string]*schema.Schema
}

func NewAlloydbBackupConverter(name string, schema map[string]*schema.Schema) common.Converter {
	return &AlloydbBackupConverter{
		name:   name,
		schema: schema,
	}
}

func (c *AlloydbBackupConverter) Convert(assets []*caiasset.Asset) ([]*common.HCLResourceBlock, error) {
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

func (c *AlloydbBackupConverter) convertResourceData(asset *caiasset.Asset, config *transport_tpg.Config) (*common.HCLResourceBlock, error) {
	if asset == nil || asset.Resource == nil || asset.Resource.Data == nil {
		return nil, fmt.Errorf("asset resource data is nil")
	}

	assetResourceData := asset.Resource.Data

	hcl, _ := resourceAlloydbBackupRead(assetResourceData, config)

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

func resourceAlloydbBackupRead(resource map[string]interface{}, config *transport_tpg.Config) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	var resource_data *schema.ResourceData = nil

	result["name"] = flattenAlloydbBackupName(resource["name"], resource_data, config)
	result["display_name"] = flattenAlloydbBackupDisplayName(resource["displayName"], resource_data, config)
	result["uid"] = flattenAlloydbBackupUid(resource["uid"], resource_data, config)
	result["create_time"] = flattenAlloydbBackupCreateTime(resource["createTime"], resource_data, config)
	result["update_time"] = flattenAlloydbBackupUpdateTime(resource["updateTime"], resource_data, config)
	result["delete_time"] = flattenAlloydbBackupDeleteTime(resource["deleteTime"], resource_data, config)
	result["labels"] = flattenAlloydbBackupLabels(resource["labels"], resource_data, config)
	result["state"] = flattenAlloydbBackupState(resource["state"], resource_data, config)
	result["type"] = flattenAlloydbBackupType(resource["type"], resource_data, config)
	result["description"] = flattenAlloydbBackupDescription(resource["description"], resource_data, config)
	result["cluster_uid"] = flattenAlloydbBackupClusterUid(resource["clusterUid"], resource_data, config)
	result["cluster_name"] = flattenAlloydbBackupClusterName(resource["clusterName"], resource_data, config)
	result["reconciling"] = flattenAlloydbBackupReconciling(resource["reconciling"], resource_data, config)
	result["encryption_config"] = flattenAlloydbBackupEncryptionConfig(resource["encryptionConfig"], resource_data, config)
	result["encryption_info"] = flattenAlloydbBackupEncryptionInfo(resource["encryptionInfo"], resource_data, config)
	result["etag"] = flattenAlloydbBackupEtag(resource["etag"], resource_data, config)
	result["annotations"] = flattenAlloydbBackupAnnotations(resource["annotations"], resource_data, config)
	result["size_bytes"] = flattenAlloydbBackupSizeBytes(resource["sizeBytes"], resource_data, config)
	result["expiry_time"] = flattenAlloydbBackupExpiryTime(resource["expiryTime"], resource_data, config)
	result["expiry_quantity"] = flattenAlloydbBackupExpiryQuantity(resource["expiryQuantity"], resource_data, config)
	result["terraform_labels"] = flattenAlloydbBackupTerraformLabels(resource["labels"], resource_data, config)
	result["effective_labels"] = flattenAlloydbBackupEffectiveLabels(resource["labels"], resource_data, config)
	result["effective_annotations"] = flattenAlloydbBackupEffectiveAnnotations(resource["annotations"], resource_data, config)

	return result, nil
}

func flattenAlloydbBackupName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAlloydbBackupDisplayName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAlloydbBackupUid(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAlloydbBackupCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAlloydbBackupUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAlloydbBackupDeleteTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAlloydbBackupLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenAlloydbBackupState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAlloydbBackupType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAlloydbBackupDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAlloydbBackupClusterUid(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAlloydbBackupClusterName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAlloydbBackupReconciling(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAlloydbBackupEncryptionConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["kms_key_name"] =
		flattenAlloydbBackupEncryptionConfigKmsKeyName(original["kmsKeyName"], d, config)
	return []interface{}{transformed}
}
func flattenAlloydbBackupEncryptionConfigKmsKeyName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAlloydbBackupEncryptionInfo(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["encryption_type"] =
		flattenAlloydbBackupEncryptionInfoEncryptionType(original["encryptionType"], d, config)
	transformed["kms_key_versions"] =
		flattenAlloydbBackupEncryptionInfoKmsKeyVersions(original["kmsKeyVersions"], d, config)
	return []interface{}{transformed}
}
func flattenAlloydbBackupEncryptionInfoEncryptionType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAlloydbBackupEncryptionInfoKmsKeyVersions(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAlloydbBackupEtag(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAlloydbBackupAnnotations(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}

	transformed := make(map[string]interface{})
	if l, ok := d.GetOkExists("annotations"); ok {
		for k := range l.(map[string]interface{}) {
			transformed[k] = v.(map[string]interface{})[k]
		}
	}

	return transformed
}

func flattenAlloydbBackupSizeBytes(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAlloydbBackupExpiryTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAlloydbBackupExpiryQuantity(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["retention_count"] =
		flattenAlloydbBackupExpiryQuantityRetentionCount(original["retentionCount"], d, config)
	transformed["total_retention_count"] =
		flattenAlloydbBackupExpiryQuantityTotalRetentionCount(original["totalRetentionCount"], d, config)
	return []interface{}{transformed}
}
func flattenAlloydbBackupExpiryQuantityRetentionCount(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenAlloydbBackupExpiryQuantityTotalRetentionCount(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenAlloydbBackupTerraformLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenAlloydbBackupEffectiveLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAlloydbBackupEffectiveAnnotations(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}
