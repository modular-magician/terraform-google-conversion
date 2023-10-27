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
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/cai2hcl/common"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/caiasset"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func resourceNameSetFromSelfLinkSet(v interface{}) *schema.Set {
	if v == nil {
		return schema.NewSet(schema.HashString, nil)
	}
	vSet := v.(*schema.Set)
	ls := make([]interface{}, 0, vSet.Len())
	for _, v := range vSet.List() {
		if v == nil {
			continue
		}
		ls = append(ls, tpgresource.GetResourceNameFromSelfLink(v.(string)))
	}
	return schema.NewSet(schema.HashString, ls)
}

// drain_nat_ips MUST be set from (just set) previous values of nat_ips
// so this customizeDiff func makes sure drainNatIps values:
//   - aren't set at creation time
//   - are in old value of nat_ips but not in new values
func resourceComputeRouterNatDrainNatIpsCustomDiff(_ context.Context, diff *schema.ResourceDiff, meta interface{}) error {
	o, n := diff.GetChange("drain_nat_ips")
	oSet := resourceNameSetFromSelfLinkSet(o)
	nSet := resourceNameSetFromSelfLinkSet(n)
	addDrainIps := nSet.Difference(oSet)

	// We don't care if there are no new drainNatIps
	if addDrainIps.Len() == 0 {
		return nil
	}

	// Resource hasn't been created yet - return error
	if diff.Id() == "" {
		return fmt.Errorf("New RouterNat cannot have drain_nat_ips, got values %+v", addDrainIps.List())
	}
	//
	o, n = diff.GetChange("nat_ips")
	oNatSet := resourceNameSetFromSelfLinkSet(o)
	nNatSet := resourceNameSetFromSelfLinkSet(n)

	// Resource is being updated - make sure new drainNatIps were in natIps prior d and no longer are in natIps.
	for _, v := range addDrainIps.List() {
		if !oNatSet.Contains(v) {
			return fmt.Errorf("drain_nat_ip %q was not previously set in nat_ips %+v", v.(string), oNatSet.List())
		}
		if nNatSet.Contains(v) {
			return fmt.Errorf("drain_nat_ip %q cannot be drained if still set in nat_ips %+v", v.(string), nNatSet.List())
		}
	}
	return nil
}

func computeRouterNatSubnetworkHash(v interface{}) int {
	obj := v.(map[string]interface{})
	name := obj["name"]
	sourceIpRanges := obj["source_ip_ranges_to_nat"]
	sourceIpRangesHash := 0
	if sourceIpRanges != nil {
		sourceIpSet := sourceIpRanges.(*schema.Set)

		for _, ipRange := range sourceIpSet.List() {
			sourceIpRangesHash += schema.HashString(ipRange.(string))
		}
	}

	secondaryIpRangeNames := obj["secondary_ip_range_names"]
	secondaryIpRangeHash := 0
	if secondaryIpRangeNames != nil {
		secondaryIpRangeSet := secondaryIpRangeNames.(*schema.Set)

		for _, secondaryIp := range secondaryIpRangeSet.List() {
			secondaryIpRangeHash += schema.HashString(secondaryIp.(string))
		}
	}

	return schema.HashString(tpgresource.NameFromSelfLinkStateFunc(name)) + sourceIpRangesHash + secondaryIpRangeHash
}

func computeRouterNatIPsHash(v interface{}) int {
	val := (v.(string))
	newParts := strings.Split(val, "/")
	if len(newParts) == 1 {
		return schema.HashString(newParts[0])
	}
	return schema.HashString(tpgresource.GetResourceNameFromSelfLink(val))
}

func computeRouterNatRulesSubnetHash(v interface{}) int {
	return computeRouterNatIPsHash(v)
}

func computeRouterNatRulesHash(v interface{}) int {
	obj := v.(map[string]interface{})
	ruleNumber := obj["rule_number"].(int)

	description := obj["description"]
	descriptionHash := 0
	if description != nil {
		descriptionHash = schema.HashString(description.(string))
	}

	match := obj["match"].(string)

	sourceNatActiveIpHash := 0
	sourceNatDrainIpHash := 0
	sourceNatActiveRangeHash := 0
	sourceNatDrainRangeHash := 0
	routerNatRulesHash := 0

	if obj["action"] != nil {
		actions := obj["action"].([]interface{})
		if len(actions) != 0 && actions[0] != nil {
			action := actions[0].(map[string]interface{})

			sourceNatActiveIps := action["source_nat_active_ips"]
			if sourceNatActiveIps != nil {
				sourceNatActiveIpSet := sourceNatActiveIps.(*schema.Set)
				for _, sourceNatActiveIp := range sourceNatActiveIpSet.List() {
					sourceNatActiveIpStr := fmt.Sprintf("source_nat_active_ips-%d", computeRouterNatIPsHash(sourceNatActiveIp.(string)))
					sourceNatActiveIpHash += schema.HashString(sourceNatActiveIpStr)
				}
			}

			soureNatDrainIps := action["source_nat_drain_ips"]
			if soureNatDrainIps != nil {
				soureNatDrainIpSet := soureNatDrainIps.(*schema.Set)
				for _, soureNatDrainIp := range soureNatDrainIpSet.List() {
					sourceNatDrainIpStr := fmt.Sprintf("source_nat_drain_ips-%d", computeRouterNatIPsHash(soureNatDrainIp.(string)))
					sourceNatDrainIpHash += schema.HashString(sourceNatDrainIpStr)
				}
			}

			sourceNatActiveRanges := action["source_nat_active_ranges"]
			if sourceNatActiveRanges != nil {
				sourceNatActiveRangesSet := sourceNatActiveRanges.(*schema.Set)
				for _, sourceNatActiveRange := range sourceNatActiveRangesSet.List() {
					sourceNatActiveRangeStr := fmt.Sprintf("source_nat_active_ranges-%d", computeRouterNatRulesSubnetHash(sourceNatActiveRange.(string)))
					sourceNatActiveRangeHash += schema.HashString(sourceNatActiveRangeStr)
				}
			}

			sourceNatDrainRanges := action["source_nat_drain_ranges"]
			if sourceNatDrainRanges != nil {
				sourceNatDrainRangesSet := sourceNatDrainRanges.(*schema.Set)
				for _, sourceNatDrainRange := range sourceNatDrainRangesSet.List() {
					sourceNatDrainRangeStr := fmt.Sprintf("source_nat_drain_ranges-%d", computeRouterNatRulesSubnetHash(sourceNatDrainRange.(string)))
					sourceNatDrainRangeHash += schema.HashString(sourceNatDrainRangeStr)
				}
			}
		}
	}

	routerNatRulesHash = ruleNumber + descriptionHash + schema.HashString(match) + sourceNatActiveIpHash + sourceNatDrainIpHash
	routerNatRulesHash += sourceNatActiveRangeHash + sourceNatDrainRangeHash
	return routerNatRulesHash
}

const ComputeRouterNatAssetType string = "compute.googleapis.com/RouterNat"

const ComputeRouterNatAssetNameRegex string = "projects/(?P<project>[^/]+)/regions/(?P<region>[^/]+)/routers/(?P<router>[^/]+)"

type ComputeRouterNatConverter struct {
	name   string
	schema map[string]*schema.Schema
}

func NewComputeRouterNatConverter(name string, schema map[string]*schema.Schema) common.Converter {
	return &ComputeRouterNatConverter{
		name:   name,
		schema: schema,
	}
}

func (c *ComputeRouterNatConverter) Convert(assets []*caiasset.Asset) ([]*common.HCLResourceBlock, error) {
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

func (c *ComputeRouterNatConverter) convertResourceData(asset *caiasset.Asset, config *transport_tpg.Config) (*common.HCLResourceBlock, error) {
	if asset == nil || asset.Resource == nil || asset.Resource.Data == nil {
		return nil, fmt.Errorf("asset resource data is nil")
	}

	assetResourceData := asset.Resource.Data

	hcl, _ := resourceComputeRouterNatRead(assetResourceData, config)

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

func resourceComputeRouterNatRead(resource map[string]interface{}, config *transport_tpg.Config) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	var resource_data *schema.ResourceData = nil

	result["name"] = flattenNestedComputeRouterNatName(resource["name"], resource_data, config)
	result["nat_ip_allocate_option"] = flattenNestedComputeRouterNatNatIpAllocateOption(resource["natIpAllocateOption"], resource_data, config)
	result["nat_ips"] = flattenNestedComputeRouterNatNatIps(resource["natIps"], resource_data, config)
	result["drain_nat_ips"] = flattenNestedComputeRouterNatDrainNatIps(resource["drainNatIps"], resource_data, config)
	result["source_subnetwork_ip_ranges_to_nat"] = flattenNestedComputeRouterNatSourceSubnetworkIpRangesToNat(resource["sourceSubnetworkIpRangesToNat"], resource_data, config)
	result["subnetwork"] = flattenNestedComputeRouterNatSubnetwork(resource["subnetworks"], resource_data, config)
	result["min_ports_per_vm"] = flattenNestedComputeRouterNatMinPortsPerVm(resource["minPortsPerVm"], resource_data, config)
	result["max_ports_per_vm"] = flattenNestedComputeRouterNatMaxPortsPerVm(resource["maxPortsPerVm"], resource_data, config)
	result["enable_dynamic_port_allocation"] = flattenNestedComputeRouterNatEnableDynamicPortAllocation(resource["enableDynamicPortAllocation"], resource_data, config)
	result["udp_idle_timeout_sec"] = flattenNestedComputeRouterNatUdpIdleTimeoutSec(resource["udpIdleTimeoutSec"], resource_data, config)
	result["icmp_idle_timeout_sec"] = flattenNestedComputeRouterNatIcmpIdleTimeoutSec(resource["icmpIdleTimeoutSec"], resource_data, config)
	result["tcp_established_idle_timeout_sec"] = flattenNestedComputeRouterNatTcpEstablishedIdleTimeoutSec(resource["tcpEstablishedIdleTimeoutSec"], resource_data, config)
	result["tcp_transitory_idle_timeout_sec"] = flattenNestedComputeRouterNatTcpTransitoryIdleTimeoutSec(resource["tcpTransitoryIdleTimeoutSec"], resource_data, config)
	result["tcp_time_wait_timeout_sec"] = flattenNestedComputeRouterNatTcpTimeWaitTimeoutSec(resource["tcpTimeWaitTimeoutSec"], resource_data, config)
	result["log_config"] = flattenNestedComputeRouterNatLogConfig(resource["logConfig"], resource_data, config)
	result["rules"] = flattenNestedComputeRouterNatRules(resource["rules"], resource_data, config)
	result["enable_endpoint_independent_mapping"] = flattenNestedComputeRouterNatEnableEndpointIndependentMapping(resource["enableEndpointIndependentMapping"], resource_data, config)
	result["type"] = flattenNestedComputeRouterNatType(resource["type"], resource_data, config)

	return result, nil
}

func flattenNestedComputeRouterNatName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNestedComputeRouterNatNatIpAllocateOption(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNestedComputeRouterNatNatIps(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.ConvertAndMapStringArr(v.([]interface{}), tpgresource.ConvertSelfLinkToV1)
}

func flattenNestedComputeRouterNatDrainNatIps(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.ConvertAndMapStringArr(v.([]interface{}), tpgresource.ConvertSelfLinkToV1)
}

func flattenNestedComputeRouterNatSourceSubnetworkIpRangesToNat(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNestedComputeRouterNatSubnetwork(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := schema.NewSet(computeRouterNatSubnetworkHash, []interface{}{})
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed.Add(map[string]interface{}{
			"name":                     flattenNestedComputeRouterNatSubnetworkName(original["name"], d, config),
			"source_ip_ranges_to_nat":  flattenNestedComputeRouterNatSubnetworkSourceIpRangesToNat(original["sourceIpRangesToNat"], d, config),
			"secondary_ip_range_names": flattenNestedComputeRouterNatSubnetworkSecondaryIpRangeNames(original["secondaryIpRangeNames"], d, config),
		})
	}
	return transformed
}
func flattenNestedComputeRouterNatSubnetworkName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.ConvertSelfLinkToV1(v.(string))
}

func flattenNestedComputeRouterNatSubnetworkSourceIpRangesToNat(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return schema.NewSet(schema.HashString, v.([]interface{}))
}

func flattenNestedComputeRouterNatSubnetworkSecondaryIpRangeNames(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return schema.NewSet(schema.HashString, v.([]interface{}))
}

func flattenNestedComputeRouterNatMinPortsPerVm(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenNestedComputeRouterNatMaxPortsPerVm(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenNestedComputeRouterNatEnableDynamicPortAllocation(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNestedComputeRouterNatUdpIdleTimeoutSec(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil || tpgresource.IsEmptyValue(reflect.ValueOf(v)) {
		return 30
	}
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := tpgresource.StringToFixed64(strVal); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}

	return v
}

func flattenNestedComputeRouterNatIcmpIdleTimeoutSec(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil || tpgresource.IsEmptyValue(reflect.ValueOf(v)) {
		return 30
	}
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := tpgresource.StringToFixed64(strVal); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}

	return v
}

func flattenNestedComputeRouterNatTcpEstablishedIdleTimeoutSec(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil || tpgresource.IsEmptyValue(reflect.ValueOf(v)) {
		return 1200
	}
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := tpgresource.StringToFixed64(strVal); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}

	return v
}

func flattenNestedComputeRouterNatTcpTransitoryIdleTimeoutSec(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil || tpgresource.IsEmptyValue(reflect.ValueOf(v)) {
		return 30
	}
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := tpgresource.StringToFixed64(strVal); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}

	return v
}

func flattenNestedComputeRouterNatTcpTimeWaitTimeoutSec(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil || tpgresource.IsEmptyValue(reflect.ValueOf(v)) {
		return 120
	}
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := tpgresource.StringToFixed64(strVal); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}

	return v
}

func flattenNestedComputeRouterNatLogConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["enable"] =
		flattenNestedComputeRouterNatLogConfigEnable(original["enable"], d, config)
	transformed["filter"] =
		flattenNestedComputeRouterNatLogConfigFilter(original["filter"], d, config)
	return []interface{}{transformed}
}
func flattenNestedComputeRouterNatLogConfigEnable(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNestedComputeRouterNatLogConfigFilter(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNestedComputeRouterNatRules(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := schema.NewSet(computeRouterNatRulesHash, []interface{}{})
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed.Add(map[string]interface{}{
			"rule_number": flattenNestedComputeRouterNatRulesRuleNumber(original["ruleNumber"], d, config),
			"description": flattenNestedComputeRouterNatRulesDescription(original["description"], d, config),
			"match":       flattenNestedComputeRouterNatRulesMatch(original["match"], d, config),
			"action":      flattenNestedComputeRouterNatRulesAction(original["action"], d, config),
		})
	}
	return transformed
}
func flattenNestedComputeRouterNatRulesRuleNumber(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenNestedComputeRouterNatRulesDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNestedComputeRouterNatRulesMatch(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNestedComputeRouterNatRulesAction(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["source_nat_active_ips"] =
		flattenNestedComputeRouterNatRulesActionSourceNatActiveIps(original["sourceNatActiveIps"], d, config)
	transformed["source_nat_drain_ips"] =
		flattenNestedComputeRouterNatRulesActionSourceNatDrainIps(original["sourceNatDrainIps"], d, config)
	transformed["source_nat_active_ranges"] =
		flattenNestedComputeRouterNatRulesActionSourceNatActiveRanges(original["sourceNatActiveRanges"], d, config)
	transformed["source_nat_drain_ranges"] =
		flattenNestedComputeRouterNatRulesActionSourceNatDrainRanges(original["sourceNatDrainRanges"], d, config)
	return []interface{}{transformed}
}
func flattenNestedComputeRouterNatRulesActionSourceNatActiveIps(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return schema.NewSet(computeRouterNatIPsHash, tpgresource.ConvertStringArrToInterface(tpgresource.ConvertAndMapStringArr(v.([]interface{}), tpgresource.ConvertSelfLinkToV1)))
}

func flattenNestedComputeRouterNatRulesActionSourceNatDrainIps(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return schema.NewSet(computeRouterNatIPsHash, tpgresource.ConvertStringArrToInterface(tpgresource.ConvertAndMapStringArr(v.([]interface{}), tpgresource.ConvertSelfLinkToV1)))
}

func flattenNestedComputeRouterNatRulesActionSourceNatActiveRanges(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return schema.NewSet(computeRouterNatRulesSubnetHash, tpgresource.ConvertStringArrToInterface(tpgresource.ConvertAndMapStringArr(v.([]interface{}), tpgresource.ConvertSelfLinkToV1)))
}

func flattenNestedComputeRouterNatRulesActionSourceNatDrainRanges(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return schema.NewSet(computeRouterNatRulesSubnetHash, tpgresource.ConvertStringArrToInterface(tpgresource.ConvertAndMapStringArr(v.([]interface{}), tpgresource.ConvertSelfLinkToV1)))
}

func flattenNestedComputeRouterNatEnableEndpointIndependentMapping(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNestedComputeRouterNatType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}
