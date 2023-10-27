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

package firebase

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/cai2hcl/common"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/caiasset"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const FirebaseAppleAppAssetType string = "firebase.googleapis.com/AppleApp"

const FirebaseAppleAppAssetNameRegex string = "projects/(?P<project>[^/]+)/iosApps"

type FirebaseAppleAppConverter struct {
	name   string
	schema map[string]*schema.Schema
}

func NewFirebaseAppleAppConverter(name string, schema map[string]*schema.Schema) common.Converter {
	return &FirebaseAppleAppConverter{
		name:   name,
		schema: schema,
	}
}

func (c *FirebaseAppleAppConverter) Convert(assets []*caiasset.Asset) ([]*common.HCLResourceBlock, error) {
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

func (c *FirebaseAppleAppConverter) convertResourceData(asset *caiasset.Asset, config *transport_tpg.Config) (*common.HCLResourceBlock, error) {
	if asset == nil || asset.Resource == nil || asset.Resource.Data == nil {
		return nil, fmt.Errorf("asset resource data is nil")
	}

	assetResourceData := asset.Resource.Data

	hcl, _ := resourceFirebaseAppleAppRead(assetResourceData, config)

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

func resourceFirebaseAppleAppRead(resource map[string]interface{}, config *transport_tpg.Config) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	var resource_data *schema.ResourceData = nil

	result["name"] = flattenFirebaseAppleAppName(resource["name"], resource_data, config)
	result["display_name"] = flattenFirebaseAppleAppDisplayName(resource["displayName"], resource_data, config)
	result["app_id"] = flattenFirebaseAppleAppAppId(resource["appId"], resource_data, config)
	result["bundle_id"] = flattenFirebaseAppleAppBundleId(resource["bundleId"], resource_data, config)
	result["app_store_id"] = flattenFirebaseAppleAppAppStoreId(resource["appStoreId"], resource_data, config)
	result["team_id"] = flattenFirebaseAppleAppTeamId(resource["teamId"], resource_data, config)
	result["api_key_id"] = flattenFirebaseAppleAppApiKeyId(resource["apiKeyId"], resource_data, config)

	return result, nil
}

func flattenFirebaseAppleAppName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenFirebaseAppleAppDisplayName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenFirebaseAppleAppAppId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenFirebaseAppleAppBundleId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenFirebaseAppleAppAppStoreId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenFirebaseAppleAppTeamId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenFirebaseAppleAppApiKeyId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}
