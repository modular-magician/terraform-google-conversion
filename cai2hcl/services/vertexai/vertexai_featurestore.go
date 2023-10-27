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

package vertexai

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/cai2hcl/common"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/caiasset"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const VertexAIFeaturestoreAssetType string = "{{region}}-aiplatform.googleapis.com/Featurestore"

const VertexAIFeaturestoreAssetNameRegex string = "projects/(?P<project>[^/]+)/locations/(?P<region>[^/]+)/featurestores"

type VertexAIFeaturestoreConverter struct {
	name   string
	schema map[string]*schema.Schema
}

func NewVertexAIFeaturestoreConverter(name string, schema map[string]*schema.Schema) common.Converter {
	return &VertexAIFeaturestoreConverter{
		name:   name,
		schema: schema,
	}
}

func (c *VertexAIFeaturestoreConverter) Convert(assets []*caiasset.Asset) ([]*common.HCLResourceBlock, error) {
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

func (c *VertexAIFeaturestoreConverter) convertResourceData(asset *caiasset.Asset, config *transport_tpg.Config) (*common.HCLResourceBlock, error) {
	if asset == nil || asset.Resource == nil || asset.Resource.Data == nil {
		return nil, fmt.Errorf("asset resource data is nil")
	}

	assetResourceData := asset.Resource.Data

	hcl, _ := resourceVertexAIFeaturestoreRead(assetResourceData, config)

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

func resourceVertexAIFeaturestoreRead(resource map[string]interface{}, config *transport_tpg.Config) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	var resource_data *schema.ResourceData = nil

	result["create_time"] = flattenVertexAIFeaturestoreCreateTime(resource["createTime"], resource_data, config)
	result["update_time"] = flattenVertexAIFeaturestoreUpdateTime(resource["updateTime"], resource_data, config)
	result["labels"] = flattenVertexAIFeaturestoreLabels(resource["labels"], resource_data, config)
	result["online_serving_config"] = flattenVertexAIFeaturestoreOnlineServingConfig(resource["onlineServingConfig"], resource_data, config)
	result["online_storage_ttl_days"] = flattenVertexAIFeaturestoreOnlineStorageTtlDays(resource["onlineStorageTtlDays"], resource_data, config)
	result["encryption_spec"] = flattenVertexAIFeaturestoreEncryptionSpec(resource["encryptionSpec"], resource_data, config)
	result["terraform_labels"] = flattenVertexAIFeaturestoreTerraformLabels(resource["labels"], resource_data, config)
	result["effective_labels"] = flattenVertexAIFeaturestoreEffectiveLabels(resource["labels"], resource_data, config)

	return result, nil
}

func flattenVertexAIFeaturestoreCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenVertexAIFeaturestoreUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenVertexAIFeaturestoreLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenVertexAIFeaturestoreOnlineServingConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["fixed_node_count"] =
		flattenVertexAIFeaturestoreOnlineServingConfigFixedNodeCount(original["fixedNodeCount"], d, config)
	transformed["scaling"] =
		flattenVertexAIFeaturestoreOnlineServingConfigScaling(original["scaling"], d, config)
	return []interface{}{transformed}
}
func flattenVertexAIFeaturestoreOnlineServingConfigFixedNodeCount(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenVertexAIFeaturestoreOnlineServingConfigScaling(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["min_node_count"] =
		flattenVertexAIFeaturestoreOnlineServingConfigScalingMinNodeCount(original["minNodeCount"], d, config)
	transformed["max_node_count"] =
		flattenVertexAIFeaturestoreOnlineServingConfigScalingMaxNodeCount(original["maxNodeCount"], d, config)
	return []interface{}{transformed}
}
func flattenVertexAIFeaturestoreOnlineServingConfigScalingMinNodeCount(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenVertexAIFeaturestoreOnlineServingConfigScalingMaxNodeCount(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenVertexAIFeaturestoreOnlineStorageTtlDays(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenVertexAIFeaturestoreEncryptionSpec(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["kms_key_name"] =
		flattenVertexAIFeaturestoreEncryptionSpecKmsKeyName(original["kmsKeyName"], d, config)
	return []interface{}{transformed}
}
func flattenVertexAIFeaturestoreEncryptionSpecKmsKeyName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenVertexAIFeaturestoreTerraformLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenVertexAIFeaturestoreEffectiveLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}
