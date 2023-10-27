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

package networksecurity

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/cai2hcl/common"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/caiasset"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const NetworkSecurityUrlListsAssetType string = "networksecurity.googleapis.com/UrlLists"

const NetworkSecurityUrlListsAssetNameRegex string = "projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/urlLists"

type NetworkSecurityUrlListsConverter struct {
	name   string
	schema map[string]*schema.Schema
}

func NewNetworkSecurityUrlListsConverter(name string, schema map[string]*schema.Schema) common.Converter {
	return &NetworkSecurityUrlListsConverter{
		name:   name,
		schema: schema,
	}
}

func (c *NetworkSecurityUrlListsConverter) Convert(assets []*caiasset.Asset) ([]*common.HCLResourceBlock, error) {
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

func (c *NetworkSecurityUrlListsConverter) convertResourceData(asset *caiasset.Asset, config *transport_tpg.Config) (*common.HCLResourceBlock, error) {
	if asset == nil || asset.Resource == nil || asset.Resource.Data == nil {
		return nil, fmt.Errorf("asset resource data is nil")
	}

	assetResourceData := asset.Resource.Data

	hcl, _ := resourceNetworkSecurityUrlListsRead(assetResourceData, config)

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

func resourceNetworkSecurityUrlListsRead(resource map[string]interface{}, config *transport_tpg.Config) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	var resource_data *schema.ResourceData = nil

	result["create_time"] = flattenNetworkSecurityUrlListsCreateTime(resource["createTime"], resource_data, config)
	result["update_time"] = flattenNetworkSecurityUrlListsUpdateTime(resource["updateTime"], resource_data, config)
	result["description"] = flattenNetworkSecurityUrlListsDescription(resource["description"], resource_data, config)
	result["values"] = flattenNetworkSecurityUrlListsValues(resource["values"], resource_data, config)

	return result, nil
}

func flattenNetworkSecurityUrlListsCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityUrlListsUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityUrlListsDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityUrlListsValues(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}
