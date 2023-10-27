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

package monitoring

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/cai2hcl/common"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/caiasset"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func ResourceMonitoringMonitoredProjectNameDiffSuppressFunc(k, old, new string, d tpgresource.TerraformResourceDataChange) bool {
	// Don't suppress if values are empty strings
	if old == "" || new == "" {
		return false
	}

	oldShort := tpgresource.GetResourceNameFromSelfLink(old)
	newShort := tpgresource.GetResourceNameFromSelfLink(new)

	// Suppress if short names are equal
	if oldShort == newShort {
		return true
	}

	_, isOldNumErr := tpgresource.StringToFixed64(oldShort)
	isOldNumber := isOldNumErr == nil
	_, isNewNumErr := tpgresource.StringToFixed64(newShort)
	isNewNumber := isNewNumErr == nil

	// Suppress if comparing a project number to project id
	return isOldNumber != isNewNumber
}

func resourceMonitoringMonitoredProjectNameDiffSuppress(k, old, new string, d *schema.ResourceData) bool {
	return ResourceMonitoringMonitoredProjectNameDiffSuppressFunc(k, old, new, d)
}

const MonitoringMonitoredProjectAssetType string = "monitoring.googleapis.com/MonitoredProject"

const MonitoringMonitoredProjectAssetNameRegex string = "v1/locations/global/metricsScopes"

type MonitoringMonitoredProjectConverter struct {
	name   string
	schema map[string]*schema.Schema
}

func NewMonitoringMonitoredProjectConverter(name string, schema map[string]*schema.Schema) common.Converter {
	return &MonitoringMonitoredProjectConverter{
		name:   name,
		schema: schema,
	}
}

func (c *MonitoringMonitoredProjectConverter) Convert(assets []*caiasset.Asset) ([]*common.HCLResourceBlock, error) {
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

func (c *MonitoringMonitoredProjectConverter) convertResourceData(asset *caiasset.Asset, config *transport_tpg.Config) (*common.HCLResourceBlock, error) {
	if asset == nil || asset.Resource == nil || asset.Resource.Data == nil {
		return nil, fmt.Errorf("asset resource data is nil")
	}

	assetResourceData := asset.Resource.Data

	hcl, _ := resourceMonitoringMonitoredProjectRead(assetResourceData, config)

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

func resourceMonitoringMonitoredProjectRead(resource map[string]interface{}, config *transport_tpg.Config) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	var resource_data *schema.ResourceData = nil

	result["name"] = flattenMonitoringMonitoredProjectName(resource["name"], resource_data, config)
	result["create_time"] = flattenMonitoringMonitoredProjectCreateTime(resource["createTime"], resource_data, config)

	return result, nil
}

func flattenMonitoringMonitoredProjectName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenMonitoringMonitoredProjectCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}
