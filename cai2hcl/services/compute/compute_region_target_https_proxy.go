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
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/cai2hcl/common"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/caiasset"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const ComputeRegionTargetHttpsProxyAssetType string = "compute.googleapis.com/RegionTargetHttpsProxy"

const ComputeRegionTargetHttpsProxyAssetNameRegex string = "projects/(?P<project>[^/]+)/regions/(?P<region>[^/]+)/targetHttpsProxies"

type ComputeRegionTargetHttpsProxyConverter struct {
	name   string
	schema map[string]*schema.Schema
}

func NewComputeRegionTargetHttpsProxyConverter(name string, schema map[string]*schema.Schema) common.Converter {
	return &ComputeRegionTargetHttpsProxyConverter{
		name:   name,
		schema: schema,
	}
}

func (c *ComputeRegionTargetHttpsProxyConverter) Convert(assets []*caiasset.Asset) ([]*common.HCLResourceBlock, error) {
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

func (c *ComputeRegionTargetHttpsProxyConverter) convertResourceData(asset *caiasset.Asset, config *transport_tpg.Config) (*common.HCLResourceBlock, error) {
	if asset == nil || asset.Resource == nil || asset.Resource.Data == nil {
		return nil, fmt.Errorf("asset resource data is nil")
	}

	assetResourceData := asset.Resource.Data

	hcl, _ := resourceComputeRegionTargetHttpsProxyRead(assetResourceData, config)

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

func resourceComputeRegionTargetHttpsProxyRead(resource map[string]interface{}, config *transport_tpg.Config) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	var resource_data *schema.ResourceData = nil

	result["creation_timestamp"] = flattenComputeRegionTargetHttpsProxyCreationTimestamp(resource["creationTimestamp"], resource_data, config)
	result["description"] = flattenComputeRegionTargetHttpsProxyDescription(resource["description"], resource_data, config)
	result["proxy_id"] = flattenComputeRegionTargetHttpsProxyProxyId(resource["id"], resource_data, config)
	result["name"] = flattenComputeRegionTargetHttpsProxyName(resource["name"], resource_data, config)
	result["ssl_certificates"] = flattenComputeRegionTargetHttpsProxySslCertificates(resource["sslCertificates"], resource_data, config)
	result["ssl_policy"] = flattenComputeRegionTargetHttpsProxySslPolicy(resource["sslPolicy"], resource_data, config)
	result["url_map"] = flattenComputeRegionTargetHttpsProxyUrlMap(resource["urlMap"], resource_data, config)
	result["region"] = flattenComputeRegionTargetHttpsProxyRegion(resource["region"], resource_data, config)

	return result, nil
}

func flattenComputeRegionTargetHttpsProxyCreationTimestamp(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeRegionTargetHttpsProxyDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeRegionTargetHttpsProxyProxyId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenComputeRegionTargetHttpsProxyName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeRegionTargetHttpsProxySslCertificates(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.ConvertAndMapStringArr(v.([]interface{}), tpgresource.ConvertSelfLinkToV1)
}

func flattenComputeRegionTargetHttpsProxySslPolicy(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.ConvertSelfLinkToV1(v.(string))
}

func flattenComputeRegionTargetHttpsProxyUrlMap(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.ConvertSelfLinkToV1(v.(string))
}

func flattenComputeRegionTargetHttpsProxyRegion(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.NameFromSelfLinkStateFunc(v)
}
