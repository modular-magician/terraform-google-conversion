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

package accesscontextmanager

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/cai2hcl/common"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/caiasset"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const AccessContextManagerServicePerimeterResourceAssetType string = "accesscontextmanager.googleapis.com/ServicePerimeterResource"

const AccessContextManagerServicePerimeterResourceAssetNameRegex string = ""

type AccessContextManagerServicePerimeterResourceConverter struct {
	name   string
	schema map[string]*schema.Schema
}

func NewAccessContextManagerServicePerimeterResourceConverter(name string, schema map[string]*schema.Schema) common.Converter {
	return &AccessContextManagerServicePerimeterResourceConverter{
		name:   name,
		schema: schema,
	}
}

func (c *AccessContextManagerServicePerimeterResourceConverter) Convert(assets []*caiasset.Asset) ([]*common.HCLResourceBlock, error) {
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

func (c *AccessContextManagerServicePerimeterResourceConverter) convertResourceData(asset *caiasset.Asset, config *transport_tpg.Config) (*common.HCLResourceBlock, error) {
	if asset == nil || asset.Resource == nil || asset.Resource.Data == nil {
		return nil, fmt.Errorf("asset resource data is nil")
	}

	assetResourceData := asset.Resource.Data

	hcl, _ := resourceAccessContextManagerServicePerimeterResourceRead(assetResourceData, config)

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

func resourceAccessContextManagerServicePerimeterResourceRead(resource map[string]interface{}, config *transport_tpg.Config) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	var resource_data *schema.ResourceData = nil

	result["resource"] = flattenNestedAccessContextManagerServicePerimeterResourceResource(resource["resource"], resource_data, config)

	return result, nil
}

func flattenNestedAccessContextManagerServicePerimeterResourceResource(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}
