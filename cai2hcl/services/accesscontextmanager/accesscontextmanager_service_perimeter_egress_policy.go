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

const AccessContextManagerServicePerimeterEgressPolicyAssetType string = "accesscontextmanager.googleapis.com/ServicePerimeterEgressPolicy"

const AccessContextManagerServicePerimeterEgressPolicyAssetNameRegex string = ""

type AccessContextManagerServicePerimeterEgressPolicyConverter struct {
	name   string
	schema map[string]*schema.Schema
}

func NewAccessContextManagerServicePerimeterEgressPolicyConverter(name string, schema map[string]*schema.Schema) common.Converter {
	return &AccessContextManagerServicePerimeterEgressPolicyConverter{
		name:   name,
		schema: schema,
	}
}

func (c *AccessContextManagerServicePerimeterEgressPolicyConverter) Convert(assets []*caiasset.Asset) ([]*common.HCLResourceBlock, error) {
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

func (c *AccessContextManagerServicePerimeterEgressPolicyConverter) convertResourceData(asset *caiasset.Asset, config *transport_tpg.Config) (*common.HCLResourceBlock, error) {
	if asset == nil || asset.Resource == nil || asset.Resource.Data == nil {
		return nil, fmt.Errorf("asset resource data is nil")
	}

	assetResourceData := asset.Resource.Data

	hcl, _ := resourceAccessContextManagerServicePerimeterEgressPolicyRead(assetResourceData, config)

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

func resourceAccessContextManagerServicePerimeterEgressPolicyRead(resource map[string]interface{}, config *transport_tpg.Config) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	var resource_data *schema.ResourceData = nil

	result["egress_from"] = flattenNestedAccessContextManagerServicePerimeterEgressPolicyEgressFrom(resource["egressFrom"], resource_data, config)
	result["egress_to"] = flattenNestedAccessContextManagerServicePerimeterEgressPolicyEgressTo(resource["egressTo"], resource_data, config)

	return result, nil
}

func flattenNestedAccessContextManagerServicePerimeterEgressPolicyEgressFrom(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["identity_type"] =
		flattenNestedAccessContextManagerServicePerimeterEgressPolicyEgressFromIdentityType(original["identityType"], d, config)
	transformed["identities"] =
		flattenNestedAccessContextManagerServicePerimeterEgressPolicyEgressFromIdentities(original["identities"], d, config)
	transformed["sources"] =
		flattenNestedAccessContextManagerServicePerimeterEgressPolicyEgressFromSources(original["sources"], d, config)
	transformed["source_restriction"] =
		flattenNestedAccessContextManagerServicePerimeterEgressPolicyEgressFromSourceRestriction(original["sourceRestriction"], d, config)
	return []interface{}{transformed}
}
func flattenNestedAccessContextManagerServicePerimeterEgressPolicyEgressFromIdentityType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNestedAccessContextManagerServicePerimeterEgressPolicyEgressFromIdentities(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNestedAccessContextManagerServicePerimeterEgressPolicyEgressFromSources(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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
			"access_level": flattenNestedAccessContextManagerServicePerimeterEgressPolicyEgressFromSourcesAccessLevel(original["accessLevel"], d, config),
		})
	}
	return transformed
}
func flattenNestedAccessContextManagerServicePerimeterEgressPolicyEgressFromSourcesAccessLevel(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNestedAccessContextManagerServicePerimeterEgressPolicyEgressFromSourceRestriction(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNestedAccessContextManagerServicePerimeterEgressPolicyEgressTo(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["resources"] =
		flattenNestedAccessContextManagerServicePerimeterEgressPolicyEgressToResources(original["resources"], d, config)
	transformed["external_resources"] =
		flattenNestedAccessContextManagerServicePerimeterEgressPolicyEgressToExternalResources(original["externalResources"], d, config)
	transformed["operations"] =
		flattenNestedAccessContextManagerServicePerimeterEgressPolicyEgressToOperations(original["operations"], d, config)
	return []interface{}{transformed}
}
func flattenNestedAccessContextManagerServicePerimeterEgressPolicyEgressToResources(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNestedAccessContextManagerServicePerimeterEgressPolicyEgressToExternalResources(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNestedAccessContextManagerServicePerimeterEgressPolicyEgressToOperations(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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
			"service_name":     flattenNestedAccessContextManagerServicePerimeterEgressPolicyEgressToOperationsServiceName(original["serviceName"], d, config),
			"method_selectors": flattenNestedAccessContextManagerServicePerimeterEgressPolicyEgressToOperationsMethodSelectors(original["methodSelectors"], d, config),
		})
	}
	return transformed
}
func flattenNestedAccessContextManagerServicePerimeterEgressPolicyEgressToOperationsServiceName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNestedAccessContextManagerServicePerimeterEgressPolicyEgressToOperationsMethodSelectors(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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
			"method":     flattenNestedAccessContextManagerServicePerimeterEgressPolicyEgressToOperationsMethodSelectorsMethod(original["method"], d, config),
			"permission": flattenNestedAccessContextManagerServicePerimeterEgressPolicyEgressToOperationsMethodSelectorsPermission(original["permission"], d, config),
		})
	}
	return transformed
}
func flattenNestedAccessContextManagerServicePerimeterEgressPolicyEgressToOperationsMethodSelectorsMethod(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNestedAccessContextManagerServicePerimeterEgressPolicyEgressToOperationsMethodSelectorsPermission(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}
