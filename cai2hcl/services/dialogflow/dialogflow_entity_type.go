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

package dialogflow

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/cai2hcl/common"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/caiasset"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const DialogflowEntityTypeAssetType string = "dialogflow.googleapis.com/EntityType"

const DialogflowEntityTypeAssetNameRegex string = "projects/(?P<project>[^/]+)/agent/entityTypes/"

type DialogflowEntityTypeConverter struct {
	name   string
	schema map[string]*schema.Schema
}

func NewDialogflowEntityTypeConverter(name string, schema map[string]*schema.Schema) common.Converter {
	return &DialogflowEntityTypeConverter{
		name:   name,
		schema: schema,
	}
}

func (c *DialogflowEntityTypeConverter) Convert(assets []*caiasset.Asset) ([]*common.HCLResourceBlock, error) {
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

func (c *DialogflowEntityTypeConverter) convertResourceData(asset *caiasset.Asset, config *transport_tpg.Config) (*common.HCLResourceBlock, error) {
	if asset == nil || asset.Resource == nil || asset.Resource.Data == nil {
		return nil, fmt.Errorf("asset resource data is nil")
	}

	assetResourceData := asset.Resource.Data

	hcl, _ := resourceDialogflowEntityTypeRead(assetResourceData, config)

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

func resourceDialogflowEntityTypeRead(resource map[string]interface{}, config *transport_tpg.Config) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	var resource_data *schema.ResourceData = nil

	result["name"] = flattenDialogflowEntityTypeName(resource["name"], resource_data, config)
	result["display_name"] = flattenDialogflowEntityTypeDisplayName(resource["displayName"], resource_data, config)
	result["kind"] = flattenDialogflowEntityTypeKind(resource["kind"], resource_data, config)
	result["enable_fuzzy_extraction"] = flattenDialogflowEntityTypeEnableFuzzyExtraction(resource["enableFuzzyExtraction"], resource_data, config)
	result["entities"] = flattenDialogflowEntityTypeEntities(resource["entities"], resource_data, config)

	return result, nil
}

func flattenDialogflowEntityTypeName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDialogflowEntityTypeDisplayName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDialogflowEntityTypeKind(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDialogflowEntityTypeEnableFuzzyExtraction(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDialogflowEntityTypeEntities(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed = append(transformed, map[string]interface{}{
			"value":    flattenDialogflowEntityTypeEntitiesValue(original["value"], d, config),
			"synonyms": flattenDialogflowEntityTypeEntitiesSynonyms(original["synonyms"], d, config),
		})
	}
	return transformed
}
func flattenDialogflowEntityTypeEntitiesValue(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDialogflowEntityTypeEntitiesSynonyms(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}
