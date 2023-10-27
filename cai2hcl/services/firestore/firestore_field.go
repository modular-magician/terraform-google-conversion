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

package firestore

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/cai2hcl/common"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/caiasset"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const FirestoreFieldAssetType string = "firestore.googleapis.com/Field"

const FirestoreFieldAssetNameRegex string = "projects/(?P<project>[^/]+)/databases/(?P<database>[^/]+)/collectionGroups/(?P<collection>[^/]+)/fields"

type FirestoreFieldConverter struct {
	name   string
	schema map[string]*schema.Schema
}

func NewFirestoreFieldConverter(name string, schema map[string]*schema.Schema) common.Converter {
	return &FirestoreFieldConverter{
		name:   name,
		schema: schema,
	}
}

func (c *FirestoreFieldConverter) Convert(assets []*caiasset.Asset) ([]*common.HCLResourceBlock, error) {
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

func (c *FirestoreFieldConverter) convertResourceData(asset *caiasset.Asset, config *transport_tpg.Config) (*common.HCLResourceBlock, error) {
	if asset == nil || asset.Resource == nil || asset.Resource.Data == nil {
		return nil, fmt.Errorf("asset resource data is nil")
	}

	assetResourceData := asset.Resource.Data

	hcl, _ := resourceFirestoreFieldRead(assetResourceData, config)

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

func resourceFirestoreFieldRead(resource map[string]interface{}, config *transport_tpg.Config) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	var resource_data *schema.ResourceData = nil

	result["name"] = flattenFirestoreFieldName(resource["name"], resource_data, config)
	result["index_config"] = flattenFirestoreFieldIndexConfig(resource["indexConfig"], resource_data, config)
	result["ttl_config"] = flattenFirestoreFieldTtlConfig(resource["ttlConfig"], resource_data, config)

	return result, nil
}

func flattenFirestoreFieldName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenFirestoreFieldIndexConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	indexConfig := v.(map[string]interface{})

	usesAncestorConfig := false
	if indexConfig["usesAncestorConfig"] != nil {
		usesAncestorConfig = indexConfig["usesAncestorConfig"].(bool)
	}

	if usesAncestorConfig {
		// The intent when uses_ancestor_config is no config.
		return []interface{}{}
	}

	if indexConfig["indexes"] == nil {
		// No indexes, return an existing, but empty index config.
		return [1]interface{}{nil}
	}

	// For Single field indexes, we put the field configuration on the index to avoid forced nesting.
	l := indexConfig["indexes"].([]interface{})
	transformed := make(map[string]interface{})
	transformedIndexes := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		fields := original["fields"].([]interface{})
		sfi := fields[0].(map[string]interface{})
		transformedIndexes = append(transformedIndexes, map[string]interface{}{
			"query_scope":  original["queryScope"],
			"order":        sfi["order"],
			"array_config": sfi["arrayConfig"],
		})
	}
	transformed["indexes"] = transformedIndexes
	return []interface{}{transformed}
}

func flattenFirestoreFieldTtlConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["state"] =
		flattenFirestoreFieldTtlConfigState(original["state"], d, config)
	return []interface{}{transformed}
}
func flattenFirestoreFieldTtlConfigState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}
