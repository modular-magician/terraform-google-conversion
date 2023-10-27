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

package datastore

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/cai2hcl/common"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/caiasset"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const DatastoreIndexAssetType string = "datastore.googleapis.com/Index"

const DatastoreIndexAssetNameRegex string = "projects/(?P<project>[^/]+)/indexes"

type DatastoreIndexConverter struct {
	name   string
	schema map[string]*schema.Schema
}

func NewDatastoreIndexConverter(name string, schema map[string]*schema.Schema) common.Converter {
	return &DatastoreIndexConverter{
		name:   name,
		schema: schema,
	}
}

func (c *DatastoreIndexConverter) Convert(assets []*caiasset.Asset) ([]*common.HCLResourceBlock, error) {
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

func (c *DatastoreIndexConverter) convertResourceData(asset *caiasset.Asset, config *transport_tpg.Config) (*common.HCLResourceBlock, error) {
	if asset == nil || asset.Resource == nil || asset.Resource.Data == nil {
		return nil, fmt.Errorf("asset resource data is nil")
	}

	assetResourceData := asset.Resource.Data

	hcl, _ := resourceDatastoreIndexRead(assetResourceData, config)

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

func resourceDatastoreIndexRead(resource map[string]interface{}, config *transport_tpg.Config) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	var resource_data *schema.ResourceData = nil

	result["index_id"] = flattenDatastoreIndexIndexId(resource["indexId"], resource_data, config)
	result["kind"] = flattenDatastoreIndexKind(resource["kind"], resource_data, config)
	result["ancestor"] = flattenDatastoreIndexAncestor(resource["ancestor"], resource_data, config)
	result["properties"] = flattenDatastoreIndexProperties(resource["properties"], resource_data, config)

	return result, nil
}

func flattenDatastoreIndexIndexId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDatastoreIndexKind(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDatastoreIndexAncestor(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDatastoreIndexProperties(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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
			"name":      flattenDatastoreIndexPropertiesName(original["name"], d, config),
			"direction": flattenDatastoreIndexPropertiesDirection(original["direction"], d, config),
		})
	}
	return transformed
}
func flattenDatastoreIndexPropertiesName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDatastoreIndexPropertiesDirection(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}
