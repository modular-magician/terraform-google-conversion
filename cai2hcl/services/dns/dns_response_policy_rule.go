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

package dns

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/cai2hcl/common"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/caiasset"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const DNSResponsePolicyRuleAssetType string = "dns.googleapis.com/ResponsePolicyRule"

const DNSResponsePolicyRuleAssetNameRegex string = "projects/(?P<project>[^/]+)/responsePolicies/(?P<response_policy>[^/]+)/rules"

type DNSResponsePolicyRuleConverter struct {
	name   string
	schema map[string]*schema.Schema
}

func NewDNSResponsePolicyRuleConverter(name string, schema map[string]*schema.Schema) common.Converter {
	return &DNSResponsePolicyRuleConverter{
		name:   name,
		schema: schema,
	}
}

func (c *DNSResponsePolicyRuleConverter) Convert(assets []*caiasset.Asset) ([]*common.HCLResourceBlock, error) {
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

func (c *DNSResponsePolicyRuleConverter) convertResourceData(asset *caiasset.Asset, config *transport_tpg.Config) (*common.HCLResourceBlock, error) {
	if asset == nil || asset.Resource == nil || asset.Resource.Data == nil {
		return nil, fmt.Errorf("asset resource data is nil")
	}

	assetResourceData := asset.Resource.Data

	hcl, _ := resourceDNSResponsePolicyRuleRead(assetResourceData, config)

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

func resourceDNSResponsePolicyRuleRead(resource map[string]interface{}, config *transport_tpg.Config) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	var resource_data *schema.ResourceData = nil

	result["rule_name"] = flattenDNSResponsePolicyRuleRuleName(resource["ruleName"], resource_data, config)
	result["dns_name"] = flattenDNSResponsePolicyRuleDnsName(resource["dnsName"], resource_data, config)
	result["local_data"] = flattenDNSResponsePolicyRuleLocalData(resource["localData"], resource_data, config)
	result["behavior"] = flattenDNSResponsePolicyRuleBehavior(resource["behavior"], resource_data, config)

	return result, nil
}

func flattenDNSResponsePolicyRuleRuleName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDNSResponsePolicyRuleDnsName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDNSResponsePolicyRuleLocalData(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["local_datas"] =
		flattenDNSResponsePolicyRuleLocalDataLocalDatas(original["localDatas"], d, config)
	return []interface{}{transformed}
}
func flattenDNSResponsePolicyRuleLocalDataLocalDatas(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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
			"name":    flattenDNSResponsePolicyRuleLocalDataLocalDatasName(original["name"], d, config),
			"type":    flattenDNSResponsePolicyRuleLocalDataLocalDatasType(original["type"], d, config),
			"ttl":     flattenDNSResponsePolicyRuleLocalDataLocalDatasTtl(original["ttl"], d, config),
			"rrdatas": flattenDNSResponsePolicyRuleLocalDataLocalDatasRrdatas(original["rrdatas"], d, config),
		})
	}
	return transformed
}
func flattenDNSResponsePolicyRuleLocalDataLocalDatasName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDNSResponsePolicyRuleLocalDataLocalDatasType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDNSResponsePolicyRuleLocalDataLocalDatasTtl(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenDNSResponsePolicyRuleLocalDataLocalDatasRrdatas(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDNSResponsePolicyRuleBehavior(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}
