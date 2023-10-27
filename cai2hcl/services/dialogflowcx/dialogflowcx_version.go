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

package dialogflowcx

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/cai2hcl/common"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/caiasset"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const DialogflowCXVersionAssetType string = "{{location}}-dialogflow.googleapis.com/Version"

const DialogflowCXVersionAssetNameRegex string = "(?P<parent>[^/]+)/versions"

type DialogflowCXVersionConverter struct {
	name   string
	schema map[string]*schema.Schema
}

func NewDialogflowCXVersionConverter(name string, schema map[string]*schema.Schema) common.Converter {
	return &DialogflowCXVersionConverter{
		name:   name,
		schema: schema,
	}
}

func (c *DialogflowCXVersionConverter) Convert(assets []*caiasset.Asset) ([]*common.HCLResourceBlock, error) {
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

func (c *DialogflowCXVersionConverter) convertResourceData(asset *caiasset.Asset, config *transport_tpg.Config) (*common.HCLResourceBlock, error) {
	if asset == nil || asset.Resource == nil || asset.Resource.Data == nil {
		return nil, fmt.Errorf("asset resource data is nil")
	}

	assetResourceData := asset.Resource.Data

	hcl, _ := resourceDialogflowCXVersionRead(assetResourceData, config)

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

func resourceDialogflowCXVersionRead(resource map[string]interface{}, config *transport_tpg.Config) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	var resource_data *schema.ResourceData = nil

	result["name"] = flattenDialogflowCXVersionName(resource["name"], resource_data, config)
	result["display_name"] = flattenDialogflowCXVersionDisplayName(resource["displayName"], resource_data, config)
	result["description"] = flattenDialogflowCXVersionDescription(resource["description"], resource_data, config)
	result["nlu_settings"] = flattenDialogflowCXVersionNluSettings(resource["nluSettings"], resource_data, config)
	result["create_time"] = flattenDialogflowCXVersionCreateTime(resource["createTime"], resource_data, config)
	result["state"] = flattenDialogflowCXVersionState(resource["state"], resource_data, config)

	return result, nil
}

func flattenDialogflowCXVersionName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.NameFromSelfLinkStateFunc(v)
}

func flattenDialogflowCXVersionDisplayName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDialogflowCXVersionDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDialogflowCXVersionNluSettings(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["model_type"] =
		flattenDialogflowCXVersionNluSettingsModelType(original["modelType"], d, config)
	transformed["classification_threshold"] =
		flattenDialogflowCXVersionNluSettingsClassificationThreshold(original["classificationThreshold"], d, config)
	transformed["model_training_mode"] =
		flattenDialogflowCXVersionNluSettingsModelTrainingMode(original["modelTrainingMode"], d, config)
	return []interface{}{transformed}
}
func flattenDialogflowCXVersionNluSettingsModelType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDialogflowCXVersionNluSettingsClassificationThreshold(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDialogflowCXVersionNluSettingsModelTrainingMode(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDialogflowCXVersionCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDialogflowCXVersionState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}
