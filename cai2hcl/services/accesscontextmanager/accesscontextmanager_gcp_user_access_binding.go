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

const AccessContextManagerGcpUserAccessBindingAssetType string = "accesscontextmanager.googleapis.com/GcpUserAccessBinding"

const AccessContextManagerGcpUserAccessBindingAssetNameRegex string = "organizations/(?P<organization_id>[^/]+)/gcpUserAccessBindings"

type AccessContextManagerGcpUserAccessBindingConverter struct {
	name   string
	schema map[string]*schema.Schema
}

func NewAccessContextManagerGcpUserAccessBindingConverter(name string, schema map[string]*schema.Schema) common.Converter {
	return &AccessContextManagerGcpUserAccessBindingConverter{
		name:   name,
		schema: schema,
	}
}

func (c *AccessContextManagerGcpUserAccessBindingConverter) Convert(assets []*caiasset.Asset) ([]*common.HCLResourceBlock, error) {
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

func (c *AccessContextManagerGcpUserAccessBindingConverter) convertResourceData(asset *caiasset.Asset, config *transport_tpg.Config) (*common.HCLResourceBlock, error) {
	if asset == nil || asset.Resource == nil || asset.Resource.Data == nil {
		return nil, fmt.Errorf("asset resource data is nil")
	}

	assetResourceData := asset.Resource.Data

	hcl, _ := resourceAccessContextManagerGcpUserAccessBindingRead(assetResourceData, config)

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

func resourceAccessContextManagerGcpUserAccessBindingRead(resource map[string]interface{}, config *transport_tpg.Config) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	var resource_data *schema.ResourceData = nil

	result["name"] = flattenAccessContextManagerGcpUserAccessBindingName(resource["name"], resource_data, config)
	result["group_key"] = flattenAccessContextManagerGcpUserAccessBindingGroupKey(resource["groupKey"], resource_data, config)
	result["access_levels"] = flattenAccessContextManagerGcpUserAccessBindingAccessLevels(resource["accessLevels"], resource_data, config)

	return result, nil
}

func flattenAccessContextManagerGcpUserAccessBindingName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAccessContextManagerGcpUserAccessBindingGroupKey(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAccessContextManagerGcpUserAccessBindingAccessLevels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}
