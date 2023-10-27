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

package gkehub2

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/cai2hcl/common"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/caiasset"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const GKEHub2FleetAssetType string = "gkehub.googleapis.com/Fleet"

const GKEHub2FleetAssetNameRegex string = "projects/(?P<project>[^/]+)/locations/global/fleets"

type GKEHub2FleetConverter struct {
	name   string
	schema map[string]*schema.Schema
}

func NewGKEHub2FleetConverter(name string, schema map[string]*schema.Schema) common.Converter {
	return &GKEHub2FleetConverter{
		name:   name,
		schema: schema,
	}
}

func (c *GKEHub2FleetConverter) Convert(assets []*caiasset.Asset) ([]*common.HCLResourceBlock, error) {
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

func (c *GKEHub2FleetConverter) convertResourceData(asset *caiasset.Asset, config *transport_tpg.Config) (*common.HCLResourceBlock, error) {
	if asset == nil || asset.Resource == nil || asset.Resource.Data == nil {
		return nil, fmt.Errorf("asset resource data is nil")
	}

	assetResourceData := asset.Resource.Data

	hcl, _ := resourceGKEHub2FleetRead(assetResourceData, config)

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

func resourceGKEHub2FleetRead(resource map[string]interface{}, config *transport_tpg.Config) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	var resource_data *schema.ResourceData = nil

	result["display_name"] = flattenGKEHub2FleetDisplayName(resource["displayName"], resource_data, config)
	result["create_time"] = flattenGKEHub2FleetCreateTime(resource["createTime"], resource_data, config)
	result["update_time"] = flattenGKEHub2FleetUpdateTime(resource["updateTime"], resource_data, config)
	result["delete_time"] = flattenGKEHub2FleetDeleteTime(resource["deleteTime"], resource_data, config)
	result["uid"] = flattenGKEHub2FleetUid(resource["uid"], resource_data, config)
	result["state"] = flattenGKEHub2FleetState(resource["state"], resource_data, config)

	return result, nil
}

func flattenGKEHub2FleetDisplayName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGKEHub2FleetCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGKEHub2FleetUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGKEHub2FleetDeleteTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGKEHub2FleetUid(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGKEHub2FleetState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["code"] =
		flattenGKEHub2FleetStateCode(original["code"], d, config)
	return []interface{}{transformed}
}
func flattenGKEHub2FleetStateCode(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}
