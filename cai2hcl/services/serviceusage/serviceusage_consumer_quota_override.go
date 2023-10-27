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

package serviceusage

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/cai2hcl/common"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/caiasset"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const ServiceUsageConsumerQuotaOverrideAssetType string = "serviceusage.googleapis.com/ConsumerQuotaOverride"

const ServiceUsageConsumerQuotaOverrideAssetNameRegex string = "projects/(?P<project>[^/]+)/services/(?P<service>[^/]+)/consumerQuotaMetrics/(?P<metric>[^/]+)/limits/(?P<limit>[^/]+)/consumerOverrides"

type ServiceUsageConsumerQuotaOverrideConverter struct {
	name   string
	schema map[string]*schema.Schema
}

func NewServiceUsageConsumerQuotaOverrideConverter(name string, schema map[string]*schema.Schema) common.Converter {
	return &ServiceUsageConsumerQuotaOverrideConverter{
		name:   name,
		schema: schema,
	}
}

func (c *ServiceUsageConsumerQuotaOverrideConverter) Convert(assets []*caiasset.Asset) ([]*common.HCLResourceBlock, error) {
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

func (c *ServiceUsageConsumerQuotaOverrideConverter) convertResourceData(asset *caiasset.Asset, config *transport_tpg.Config) (*common.HCLResourceBlock, error) {
	if asset == nil || asset.Resource == nil || asset.Resource.Data == nil {
		return nil, fmt.Errorf("asset resource data is nil")
	}

	assetResourceData := asset.Resource.Data

	hcl, _ := resourceServiceUsageConsumerQuotaOverrideRead(assetResourceData, config)

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

func resourceServiceUsageConsumerQuotaOverrideRead(resource map[string]interface{}, config *transport_tpg.Config) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	var resource_data *schema.ResourceData = nil

	result["override_value"] = flattenNestedServiceUsageConsumerQuotaOverrideOverrideValue(resource["overrideValue"], resource_data, config)
	result["dimensions"] = flattenNestedServiceUsageConsumerQuotaOverrideDimensions(resource["dimensions"], resource_data, config)
	result["name"] = flattenNestedServiceUsageConsumerQuotaOverrideName(resource["name"], resource_data, config)

	return result, nil
}

func flattenNestedServiceUsageConsumerQuotaOverrideOverrideValue(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return "0"
	}
	return v
}

func flattenNestedServiceUsageConsumerQuotaOverrideDimensions(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNestedServiceUsageConsumerQuotaOverrideName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.NameFromSelfLinkStateFunc(v)
}
