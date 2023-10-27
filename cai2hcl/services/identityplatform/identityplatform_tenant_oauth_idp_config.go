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

package identityplatform

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/cai2hcl/common"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/caiasset"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const IdentityPlatformTenantOauthIdpConfigAssetType string = "identitytoolkit.googleapis.com/TenantOauthIdpConfig"

const IdentityPlatformTenantOauthIdpConfigAssetNameRegex string = "projects/(?P<project>[^/]+)/tenants/(?P<tenant>[^/]+)/oauthIdpConfigs"

type IdentityPlatformTenantOauthIdpConfigConverter struct {
	name   string
	schema map[string]*schema.Schema
}

func NewIdentityPlatformTenantOauthIdpConfigConverter(name string, schema map[string]*schema.Schema) common.Converter {
	return &IdentityPlatformTenantOauthIdpConfigConverter{
		name:   name,
		schema: schema,
	}
}

func (c *IdentityPlatformTenantOauthIdpConfigConverter) Convert(assets []*caiasset.Asset) ([]*common.HCLResourceBlock, error) {
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

func (c *IdentityPlatformTenantOauthIdpConfigConverter) convertResourceData(asset *caiasset.Asset, config *transport_tpg.Config) (*common.HCLResourceBlock, error) {
	if asset == nil || asset.Resource == nil || asset.Resource.Data == nil {
		return nil, fmt.Errorf("asset resource data is nil")
	}

	assetResourceData := asset.Resource.Data

	hcl, _ := resourceIdentityPlatformTenantOauthIdpConfigRead(assetResourceData, config)

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

func resourceIdentityPlatformTenantOauthIdpConfigRead(resource map[string]interface{}, config *transport_tpg.Config) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	var resource_data *schema.ResourceData = nil

	result["name"] = flattenIdentityPlatformTenantOauthIdpConfigName(resource["name"], resource_data, config)
	result["display_name"] = flattenIdentityPlatformTenantOauthIdpConfigDisplayName(resource["displayName"], resource_data, config)
	result["enabled"] = flattenIdentityPlatformTenantOauthIdpConfigEnabled(resource["enabled"], resource_data, config)
	result["issuer"] = flattenIdentityPlatformTenantOauthIdpConfigIssuer(resource["issuer"], resource_data, config)
	result["client_id"] = flattenIdentityPlatformTenantOauthIdpConfigClientId(resource["clientId"], resource_data, config)
	result["client_secret"] = flattenIdentityPlatformTenantOauthIdpConfigClientSecret(resource["clientSecret"], resource_data, config)

	return result, nil
}

func flattenIdentityPlatformTenantOauthIdpConfigName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.NameFromSelfLinkStateFunc(v)
}

func flattenIdentityPlatformTenantOauthIdpConfigDisplayName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIdentityPlatformTenantOauthIdpConfigEnabled(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIdentityPlatformTenantOauthIdpConfigIssuer(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIdentityPlatformTenantOauthIdpConfigClientId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIdentityPlatformTenantOauthIdpConfigClientSecret(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}
