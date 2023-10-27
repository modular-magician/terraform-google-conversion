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

package databasemigrationservice

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/cai2hcl/common"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/caiasset"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const DatabaseMigrationServicePrivateConnectionAssetType string = "datamigration.googleapis.com/PrivateConnection"

const DatabaseMigrationServicePrivateConnectionAssetNameRegex string = "projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/privateConnections"

type DatabaseMigrationServicePrivateConnectionConverter struct {
	name   string
	schema map[string]*schema.Schema
}

func NewDatabaseMigrationServicePrivateConnectionConverter(name string, schema map[string]*schema.Schema) common.Converter {
	return &DatabaseMigrationServicePrivateConnectionConverter{
		name:   name,
		schema: schema,
	}
}

func (c *DatabaseMigrationServicePrivateConnectionConverter) Convert(assets []*caiasset.Asset) ([]*common.HCLResourceBlock, error) {
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

func (c *DatabaseMigrationServicePrivateConnectionConverter) convertResourceData(asset *caiasset.Asset, config *transport_tpg.Config) (*common.HCLResourceBlock, error) {
	if asset == nil || asset.Resource == nil || asset.Resource.Data == nil {
		return nil, fmt.Errorf("asset resource data is nil")
	}

	assetResourceData := asset.Resource.Data

	hcl, _ := resourceDatabaseMigrationServicePrivateConnectionRead(assetResourceData, config)

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

func resourceDatabaseMigrationServicePrivateConnectionRead(resource map[string]interface{}, config *transport_tpg.Config) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	var resource_data *schema.ResourceData = nil

	result["name"] = flattenDatabaseMigrationServicePrivateConnectionName(resource["name"], resource_data, config)
	result["labels"] = flattenDatabaseMigrationServicePrivateConnectionLabels(resource["labels"], resource_data, config)
	result["display_name"] = flattenDatabaseMigrationServicePrivateConnectionDisplayName(resource["displayName"], resource_data, config)
	result["state"] = flattenDatabaseMigrationServicePrivateConnectionState(resource["state"], resource_data, config)
	result["error"] = flattenDatabaseMigrationServicePrivateConnectionError(resource["error"], resource_data, config)
	result["vpc_peering_config"] = flattenDatabaseMigrationServicePrivateConnectionVpcPeeringConfig(resource["vpcPeeringConfig"], resource_data, config)
	result["terraform_labels"] = flattenDatabaseMigrationServicePrivateConnectionTerraformLabels(resource["labels"], resource_data, config)
	result["effective_labels"] = flattenDatabaseMigrationServicePrivateConnectionEffectiveLabels(resource["labels"], resource_data, config)

	return result, nil
}

func flattenDatabaseMigrationServicePrivateConnectionName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDatabaseMigrationServicePrivateConnectionLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenDatabaseMigrationServicePrivateConnectionDisplayName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDatabaseMigrationServicePrivateConnectionState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDatabaseMigrationServicePrivateConnectionError(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["message"] =
		flattenDatabaseMigrationServicePrivateConnectionErrorMessage(original["message"], d, config)
	transformed["details"] =
		flattenDatabaseMigrationServicePrivateConnectionErrorDetails(original["details"], d, config)
	return []interface{}{transformed}
}
func flattenDatabaseMigrationServicePrivateConnectionErrorMessage(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDatabaseMigrationServicePrivateConnectionErrorDetails(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDatabaseMigrationServicePrivateConnectionVpcPeeringConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["vpc_name"] =
		flattenDatabaseMigrationServicePrivateConnectionVpcPeeringConfigVpcName(original["vpcName"], d, config)
	transformed["subnet"] =
		flattenDatabaseMigrationServicePrivateConnectionVpcPeeringConfigSubnet(original["subnet"], d, config)
	return []interface{}{transformed}
}
func flattenDatabaseMigrationServicePrivateConnectionVpcPeeringConfigVpcName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDatabaseMigrationServicePrivateConnectionVpcPeeringConfigSubnet(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDatabaseMigrationServicePrivateConnectionTerraformLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenDatabaseMigrationServicePrivateConnectionEffectiveLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}
