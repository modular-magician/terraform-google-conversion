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

package beyondcorp

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/cai2hcl/common"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/caiasset"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const BeyondcorpAppGatewayAssetType string = "beyondcorp.googleapis.com/AppGateway"

const BeyondcorpAppGatewayAssetNameRegex string = "projects/(?P<project>[^/]+)/locations/(?P<region>[^/]+)/appGateways"

type BeyondcorpAppGatewayConverter struct {
	name   string
	schema map[string]*schema.Schema
}

func NewBeyondcorpAppGatewayConverter(name string, schema map[string]*schema.Schema) common.Converter {
	return &BeyondcorpAppGatewayConverter{
		name:   name,
		schema: schema,
	}
}

func (c *BeyondcorpAppGatewayConverter) Convert(assets []*caiasset.Asset) ([]*common.HCLResourceBlock, error) {
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

func (c *BeyondcorpAppGatewayConverter) convertResourceData(asset *caiasset.Asset, config *transport_tpg.Config) (*common.HCLResourceBlock, error) {
	if asset == nil || asset.Resource == nil || asset.Resource.Data == nil {
		return nil, fmt.Errorf("asset resource data is nil")
	}

	assetResourceData := asset.Resource.Data

	hcl, _ := resourceBeyondcorpAppGatewayRead(assetResourceData, config)

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

func resourceBeyondcorpAppGatewayRead(resource map[string]interface{}, config *transport_tpg.Config) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	var resource_data *schema.ResourceData = nil

	result["type"] = flattenBeyondcorpAppGatewayType(resource["type"], resource_data, config)
	result["host_type"] = flattenBeyondcorpAppGatewayHostType(resource["hostType"], resource_data, config)
	result["display_name"] = flattenBeyondcorpAppGatewayDisplayName(resource["displayName"], resource_data, config)
	result["labels"] = flattenBeyondcorpAppGatewayLabels(resource["labels"], resource_data, config)
	result["state"] = flattenBeyondcorpAppGatewayState(resource["state"], resource_data, config)
	result["uri"] = flattenBeyondcorpAppGatewayUri(resource["uri"], resource_data, config)
	result["allocated_connections"] = flattenBeyondcorpAppGatewayAllocatedConnections(resource["allocatedConnections"], resource_data, config)
	result["terraform_labels"] = flattenBeyondcorpAppGatewayTerraformLabels(resource["labels"], resource_data, config)
	result["effective_labels"] = flattenBeyondcorpAppGatewayEffectiveLabels(resource["labels"], resource_data, config)

	return result, nil
}

func flattenBeyondcorpAppGatewayType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBeyondcorpAppGatewayHostType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBeyondcorpAppGatewayDisplayName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBeyondcorpAppGatewayLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenBeyondcorpAppGatewayState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBeyondcorpAppGatewayUri(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBeyondcorpAppGatewayAllocatedConnections(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["psc_uri"] =
		flattenBeyondcorpAppGatewayAllocatedConnectionsPscUri(original["pscUri"], d, config)
	transformed["ingress_port"] =
		flattenBeyondcorpAppGatewayAllocatedConnectionsIngressPort(original["ingressPort"], d, config)
	return []interface{}{transformed}
}
func flattenBeyondcorpAppGatewayAllocatedConnectionsPscUri(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBeyondcorpAppGatewayAllocatedConnectionsIngressPort(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenBeyondcorpAppGatewayTerraformLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenBeyondcorpAppGatewayEffectiveLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}
