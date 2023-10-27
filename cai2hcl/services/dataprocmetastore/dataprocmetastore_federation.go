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

package dataprocmetastore

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/cai2hcl/common"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/caiasset"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const DataprocMetastoreFederationAssetType string = "metastore.googleapis.com/Federation"

const DataprocMetastoreFederationAssetNameRegex string = "projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/federations"

type DataprocMetastoreFederationConverter struct {
	name   string
	schema map[string]*schema.Schema
}

func NewDataprocMetastoreFederationConverter(name string, schema map[string]*schema.Schema) common.Converter {
	return &DataprocMetastoreFederationConverter{
		name:   name,
		schema: schema,
	}
}

func (c *DataprocMetastoreFederationConverter) Convert(assets []*caiasset.Asset) ([]*common.HCLResourceBlock, error) {
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

func (c *DataprocMetastoreFederationConverter) convertResourceData(asset *caiasset.Asset, config *transport_tpg.Config) (*common.HCLResourceBlock, error) {
	if asset == nil || asset.Resource == nil || asset.Resource.Data == nil {
		return nil, fmt.Errorf("asset resource data is nil")
	}

	assetResourceData := asset.Resource.Data

	hcl, _ := resourceDataprocMetastoreFederationRead(assetResourceData, config)

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

func resourceDataprocMetastoreFederationRead(resource map[string]interface{}, config *transport_tpg.Config) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	var resource_data *schema.ResourceData = nil

	result["name"] = flattenDataprocMetastoreFederationName(resource["name"], resource_data, config)
	result["labels"] = flattenDataprocMetastoreFederationLabels(resource["labels"], resource_data, config)
	result["endpoint_uri"] = flattenDataprocMetastoreFederationEndpointUri(resource["endpointUri"], resource_data, config)
	result["state"] = flattenDataprocMetastoreFederationState(resource["state"], resource_data, config)
	result["state_message"] = flattenDataprocMetastoreFederationStateMessage(resource["stateMessage"], resource_data, config)
	result["uid"] = flattenDataprocMetastoreFederationUid(resource["uid"], resource_data, config)
	result["version"] = flattenDataprocMetastoreFederationVersion(resource["version"], resource_data, config)
	result["backend_metastores"] = flattenDataprocMetastoreFederationBackendMetastores(resource["backendMetastores"], resource_data, config)
	result["terraform_labels"] = flattenDataprocMetastoreFederationTerraformLabels(resource["labels"], resource_data, config)
	result["effective_labels"] = flattenDataprocMetastoreFederationEffectiveLabels(resource["labels"], resource_data, config)

	return result, nil
}

func flattenDataprocMetastoreFederationName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataprocMetastoreFederationLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenDataprocMetastoreFederationEndpointUri(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataprocMetastoreFederationState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataprocMetastoreFederationStateMessage(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataprocMetastoreFederationUid(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataprocMetastoreFederationVersion(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataprocMetastoreFederationBackendMetastores(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	l := v.(map[string]interface{})
	transformed := make([]interface{}, 0, len(l))
	for k, raw := range l {
		original := raw.(map[string]interface{})
		transformed = append(transformed, map[string]interface{}{
			"rank":           k,
			"name":           flattenDataprocMetastoreFederationBackendMetastoresName(original["name"], d, config),
			"metastore_type": flattenDataprocMetastoreFederationBackendMetastoresMetastoreType(original["metastoreType"], d, config),
		})
	}
	return transformed
}
func flattenDataprocMetastoreFederationBackendMetastoresName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataprocMetastoreFederationBackendMetastoresMetastoreType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataprocMetastoreFederationTerraformLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenDataprocMetastoreFederationEffectiveLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}
