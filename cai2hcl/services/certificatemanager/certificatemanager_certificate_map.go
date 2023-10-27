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

package certificatemanager

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/cai2hcl/common"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/caiasset"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const CertificateManagerCertificateMapAssetType string = "certificatemanager.googleapis.com/CertificateMap"

const CertificateManagerCertificateMapAssetNameRegex string = "projects/(?P<project>[^/]+)/locations/global/certificateMaps"

type CertificateManagerCertificateMapConverter struct {
	name   string
	schema map[string]*schema.Schema
}

func NewCertificateManagerCertificateMapConverter(name string, schema map[string]*schema.Schema) common.Converter {
	return &CertificateManagerCertificateMapConverter{
		name:   name,
		schema: schema,
	}
}

func (c *CertificateManagerCertificateMapConverter) Convert(assets []*caiasset.Asset) ([]*common.HCLResourceBlock, error) {
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

func (c *CertificateManagerCertificateMapConverter) convertResourceData(asset *caiasset.Asset, config *transport_tpg.Config) (*common.HCLResourceBlock, error) {
	if asset == nil || asset.Resource == nil || asset.Resource.Data == nil {
		return nil, fmt.Errorf("asset resource data is nil")
	}

	assetResourceData := asset.Resource.Data

	hcl, _ := resourceCertificateManagerCertificateMapRead(assetResourceData, config)

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

func resourceCertificateManagerCertificateMapRead(resource map[string]interface{}, config *transport_tpg.Config) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	var resource_data *schema.ResourceData = nil

	result["description"] = flattenCertificateManagerCertificateMapDescription(resource["description"], resource_data, config)
	result["create_time"] = flattenCertificateManagerCertificateMapCreateTime(resource["createTime"], resource_data, config)
	result["update_time"] = flattenCertificateManagerCertificateMapUpdateTime(resource["updateTime"], resource_data, config)
	result["labels"] = flattenCertificateManagerCertificateMapLabels(resource["labels"], resource_data, config)
	result["gclb_targets"] = flattenCertificateManagerCertificateMapGclbTargets(resource["gclbTargets"], resource_data, config)
	result["terraform_labels"] = flattenCertificateManagerCertificateMapTerraformLabels(resource["labels"], resource_data, config)
	result["effective_labels"] = flattenCertificateManagerCertificateMapEffectiveLabels(resource["labels"], resource_data, config)

	return result, nil
}

func flattenCertificateManagerCertificateMapDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenCertificateManagerCertificateMapCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenCertificateManagerCertificateMapUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenCertificateManagerCertificateMapLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenCertificateManagerCertificateMapGclbTargets(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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
			"ip_configs":         flattenCertificateManagerCertificateMapGclbTargetsIpConfigs(original["ipConfigs"], d, config),
			"target_https_proxy": flattenCertificateManagerCertificateMapGclbTargetsTargetHttpsProxy(original["targetHttpsProxy"], d, config),
			"target_ssl_proxy":   flattenCertificateManagerCertificateMapGclbTargetsTargetSslProxy(original["targetSslProxy"], d, config),
		})
	}
	return transformed
}
func flattenCertificateManagerCertificateMapGclbTargetsIpConfigs(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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
			"ip_address": flattenCertificateManagerCertificateMapGclbTargetsIpConfigsIpAddress(original["ipAddress"], d, config),
			"ports":      flattenCertificateManagerCertificateMapGclbTargetsIpConfigsPorts(original["ports"], d, config),
		})
	}
	return transformed
}
func flattenCertificateManagerCertificateMapGclbTargetsIpConfigsIpAddress(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenCertificateManagerCertificateMapGclbTargetsIpConfigsPorts(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenCertificateManagerCertificateMapGclbTargetsTargetHttpsProxy(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenCertificateManagerCertificateMapGclbTargetsTargetSslProxy(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenCertificateManagerCertificateMapTerraformLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenCertificateManagerCertificateMapEffectiveLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}
