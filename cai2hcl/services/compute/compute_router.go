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

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/cai2hcl/common"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/caiasset"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

// customizeDiff func for additional checks on google_compute_router properties:
func resourceComputeRouterCustomDiff(_ context.Context, diff *schema.ResourceDiff, meta interface{}) error {

	block := diff.Get("bgp.0").(map[string]interface{})
	advertiseMode := block["advertise_mode"]
	advertisedGroups := block["advertised_groups"].([]interface{})
	advertisedIPRanges := block["advertised_ip_ranges"].([]interface{})

	if advertiseMode == "DEFAULT" && len(advertisedGroups) != 0 {
		return fmt.Errorf("Error in bgp: advertised_groups cannot be specified when using advertise_mode DEFAULT")
	}
	if advertiseMode == "DEFAULT" && len(advertisedIPRanges) != 0 {
		return fmt.Errorf("Error in bgp: advertised_ip_ranges cannot be specified when using advertise_mode DEFAULT")
	}

	return nil
}

const ComputeRouterAssetType string = "compute.googleapis.com/Router"

const ComputeRouterAssetNameRegex string = "projects/(?P<project>[^/]+)/regions/(?P<region>[^/]+)/routers"

type ComputeRouterConverter struct {
	name   string
	schema map[string]*schema.Schema
}

func NewComputeRouterConverter(name string, schema map[string]*schema.Schema) common.Converter {
	return &ComputeRouterConverter{
		name:   name,
		schema: schema,
	}
}

func (c *ComputeRouterConverter) Convert(assets []*caiasset.Asset) ([]*common.HCLResourceBlock, error) {
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

func (c *ComputeRouterConverter) convertResourceData(asset *caiasset.Asset, config *transport_tpg.Config) (*common.HCLResourceBlock, error) {
	if asset == nil || asset.Resource == nil || asset.Resource.Data == nil {
		return nil, fmt.Errorf("asset resource data is nil")
	}

	assetResourceData := asset.Resource.Data

	hcl, _ := resourceComputeRouterRead(assetResourceData, config)

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

func resourceComputeRouterRead(resource map[string]interface{}, config *transport_tpg.Config) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	var resource_data *schema.ResourceData = nil

	result["creation_timestamp"] = flattenComputeRouterCreationTimestamp(resource["creationTimestamp"], resource_data, config)
	result["name"] = flattenComputeRouterName(resource["name"], resource_data, config)
	result["description"] = flattenComputeRouterDescription(resource["description"], resource_data, config)
	result["network"] = flattenComputeRouterNetwork(resource["network"], resource_data, config)
	result["bgp"] = flattenComputeRouterBgp(resource["bgp"], resource_data, config)
	result["encrypted_interconnect_router"] = flattenComputeRouterEncryptedInterconnectRouter(resource["encryptedInterconnectRouter"], resource_data, config)
	result["region"] = flattenComputeRouterRegion(resource["region"], resource_data, config)

	return result, nil
}

func flattenComputeRouterCreationTimestamp(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeRouterName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeRouterDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeRouterNetwork(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.ConvertSelfLinkToV1(v.(string))
}

func flattenComputeRouterBgp(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["asn"] =
		flattenComputeRouterBgpAsn(original["asn"], d, config)
	transformed["advertise_mode"] =
		flattenComputeRouterBgpAdvertiseMode(original["advertiseMode"], d, config)
	transformed["advertised_groups"] =
		flattenComputeRouterBgpAdvertisedGroups(original["advertisedGroups"], d, config)
	transformed["advertised_ip_ranges"] =
		flattenComputeRouterBgpAdvertisedIpRanges(original["advertisedIpRanges"], d, config)
	transformed["keepalive_interval"] =
		flattenComputeRouterBgpKeepaliveInterval(original["keepaliveInterval"], d, config)
	return []interface{}{transformed}
}
func flattenComputeRouterBgpAsn(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenComputeRouterBgpAdvertiseMode(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeRouterBgpAdvertisedGroups(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeRouterBgpAdvertisedIpRanges(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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
			"range":       flattenComputeRouterBgpAdvertisedIpRangesRange(original["range"], d, config),
			"description": flattenComputeRouterBgpAdvertisedIpRangesDescription(original["description"], d, config),
		})
	}
	return transformed
}
func flattenComputeRouterBgpAdvertisedIpRangesRange(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeRouterBgpAdvertisedIpRangesDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeRouterBgpKeepaliveInterval(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenComputeRouterEncryptedInterconnectRouter(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeRouterRegion(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.NameFromSelfLinkStateFunc(v)
}
