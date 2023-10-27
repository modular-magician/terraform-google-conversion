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

package compute

import (
	"bytes"
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/cai2hcl/common"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/caiasset"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func resourceComputeFirewallRuleHash(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	buf.WriteString(fmt.Sprintf("%s-", strings.ToLower(m["protocol"].(string))))

	// We need to make sure to sort the strings below so that we always
	// generate the same hash code no matter what is in the set.
	if v, ok := m["ports"]; ok && v != nil {
		s := tpgresource.ConvertStringArr(v.([]interface{}))
		sort.Strings(s)

		for _, v := range s {
			buf.WriteString(fmt.Sprintf("%s-", v))
		}
	}

	return tpgresource.Hashcode(buf.String())
}

func diffSuppressEnableLogging(k, old, new string, d *schema.ResourceData) bool {
	if k == "log_config.#" {
		if new == "0" && d.Get("enable_logging").(bool) {
			return true
		}
	}

	return false
}

func resourceComputeFirewallEnableLoggingCustomizeDiff(_ context.Context, diff *schema.ResourceDiff, v interface{}) error {
	enableLogging, enableExists := diff.GetOkExists("enable_logging")
	if !enableExists {
		return nil
	}

	logConfigExists := diff.Get("log_config.#").(int) != 0
	if logConfigExists && enableLogging == false {
		return fmt.Errorf("log_config cannot be defined when enable_logging is false")
	}

	return nil
}

// Per https://github.com/hashicorp/terraform-provider-google/issues/2924
// Make one of the source_ parameters Required in ingress google_compute_firewall
func resourceComputeFirewallSourceFieldsCustomizeDiff(_ context.Context, diff *schema.ResourceDiff, v interface{}) error {
	direction := diff.Get("direction").(string)

	if direction != "EGRESS" {
		_, tagsOk := diff.GetOk("source_tags")
		_, rangesOk := diff.GetOk("source_ranges")
		_, sasOk := diff.GetOk("source_service_accounts")

		_, tagsExist := diff.GetOkExists("source_tags")
		_, rangesExist := diff.GetOkExists("source_ranges")
		_, sasExist := diff.GetOkExists("source_service_accounts")

		if !tagsOk && !rangesOk && !sasOk && !tagsExist && !rangesExist && !sasExist {
			return fmt.Errorf("one of source_tags, source_ranges, or source_service_accounts must be defined")
		}
	}

	return nil
}

func diffSuppressSourceRanges(k, old, new string, d *schema.ResourceData) bool {
	if k == "source_ranges.#" {
		if old == "1" && new == "0" {
			// Allow diffing on the individual element if we are going from 1 -> 0
			// this allows for diff suppress on ["0.0.0.0/0"] -> []
			return true
		}
		// For any other source_ranges.# diff, don't suppress
		return false
	}
	kLength := "source_ranges.#"
	oldLength, newLength := d.GetChange(kLength)
	oldInt, ok := oldLength.(int)

	if !ok {
		return false
	}

	newInt, ok := newLength.(int)
	if !ok {
		return false
	}

	// Diff suppress only should suppress removing the default range
	// This should probably be newInt == 0, but due to Terraform core internals
	// (bug?) values found via GetChange may not have the correct new value
	// in some circumstances
	if oldInt == 1 && newInt == 1 {
		if old == "0.0.0.0/0" && new == "" {
			return true
		}
	}
	// For any other source_ranges value diff, don't suppress
	return false
}

const ComputeFirewallAssetType string = "compute.googleapis.com/Firewall"

const ComputeFirewallAssetNameRegex string = "projects/(?P<project>[^/]+)/global/firewalls"

type ComputeFirewallConverter struct {
	name   string
	schema map[string]*schema.Schema
}

func NewComputeFirewallConverter(name string, schema map[string]*schema.Schema) common.Converter {
	return &ComputeFirewallConverter{
		name:   name,
		schema: schema,
	}
}

func (c *ComputeFirewallConverter) Convert(assets []*caiasset.Asset) ([]*common.HCLResourceBlock, error) {
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

func (c *ComputeFirewallConverter) convertResourceData(asset *caiasset.Asset, config *transport_tpg.Config) (*common.HCLResourceBlock, error) {
	if asset == nil || asset.Resource == nil || asset.Resource.Data == nil {
		return nil, fmt.Errorf("asset resource data is nil")
	}

	assetResourceData := asset.Resource.Data

	hcl, _ := resourceComputeFirewallRead(assetResourceData, config)

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

func resourceComputeFirewallRead(resource map[string]interface{}, config *transport_tpg.Config) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	var resource_data *schema.ResourceData = nil

	result["allow"] = flattenComputeFirewallAllow(resource["allowed"], resource_data, config)
	result["creation_timestamp"] = flattenComputeFirewallCreationTimestamp(resource["creationTimestamp"], resource_data, config)
	result["deny"] = flattenComputeFirewallDeny(resource["denied"], resource_data, config)
	result["description"] = flattenComputeFirewallDescription(resource["description"], resource_data, config)
	result["destination_ranges"] = flattenComputeFirewallDestinationRanges(resource["destinationRanges"], resource_data, config)
	result["direction"] = flattenComputeFirewallDirection(resource["direction"], resource_data, config)
	result["disabled"] = flattenComputeFirewallDisabled(resource["disabled"], resource_data, config)
	result["log_config"] = flattenComputeFirewallLogConfig(resource["logConfig"], resource_data, config)
	result["name"] = flattenComputeFirewallName(resource["name"], resource_data, config)
	result["network"] = flattenComputeFirewallNetwork(resource["network"], resource_data, config)
	result["priority"] = flattenComputeFirewallPriority(resource["priority"], resource_data, config)
	result["source_ranges"] = flattenComputeFirewallSourceRanges(resource["sourceRanges"], resource_data, config)
	result["source_service_accounts"] = flattenComputeFirewallSourceServiceAccounts(resource["sourceServiceAccounts"], resource_data, config)
	result["source_tags"] = flattenComputeFirewallSourceTags(resource["sourceTags"], resource_data, config)
	result["target_service_accounts"] = flattenComputeFirewallTargetServiceAccounts(resource["targetServiceAccounts"], resource_data, config)
	result["target_tags"] = flattenComputeFirewallTargetTags(resource["targetTags"], resource_data, config)

	return result, nil
}

func flattenComputeFirewallAllow(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := schema.NewSet(resourceComputeFirewallRuleHash, []interface{}{})
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed.Add(map[string]interface{}{
			"protocol": flattenComputeFirewallAllowProtocol(original["IPProtocol"], d, config),
			"ports":    flattenComputeFirewallAllowPorts(original["ports"], d, config),
		})
	}
	return transformed
}
func flattenComputeFirewallAllowProtocol(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeFirewallAllowPorts(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeFirewallCreationTimestamp(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeFirewallDeny(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := schema.NewSet(resourceComputeFirewallRuleHash, []interface{}{})
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed.Add(map[string]interface{}{
			"protocol": flattenComputeFirewallDenyProtocol(original["IPProtocol"], d, config),
			"ports":    flattenComputeFirewallDenyPorts(original["ports"], d, config),
		})
	}
	return transformed
}
func flattenComputeFirewallDenyProtocol(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeFirewallDenyPorts(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeFirewallDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeFirewallDestinationRanges(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return schema.NewSet(schema.HashString, v.([]interface{}))
}

func flattenComputeFirewallDirection(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeFirewallDisabled(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeFirewallLogConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}

	v, ok := original["enable"]
	if ok && !v.(bool) {
		return nil
	}

	transformed := make(map[string]interface{})
	transformed["metadata"] = original["metadata"]
	return []interface{}{transformed}
}

func flattenComputeFirewallName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeFirewallNetwork(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.ConvertSelfLinkToV1(v.(string))
}

func flattenComputeFirewallPriority(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenComputeFirewallSourceRanges(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return schema.NewSet(schema.HashString, v.([]interface{}))
}

func flattenComputeFirewallSourceServiceAccounts(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return schema.NewSet(schema.HashString, v.([]interface{}))
}

func flattenComputeFirewallSourceTags(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return schema.NewSet(schema.HashString, v.([]interface{}))
}

func flattenComputeFirewallTargetServiceAccounts(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return schema.NewSet(schema.HashString, v.([]interface{}))
}

func flattenComputeFirewallTargetTags(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return schema.NewSet(schema.HashString, v.([]interface{}))
}
