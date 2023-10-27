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
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const ComputeOrganizationSecurityPolicyAssociationAssetType string = "compute.googleapis.com/OrganizationSecurityPolicyAssociation"

const ComputeOrganizationSecurityPolicyAssociationAssetNameRegex string = "(?P<policy_id>[^/]+)"

type ComputeOrganizationSecurityPolicyAssociationConverter struct {
	name   string
	schema map[string]*schema.Schema
}

func NewComputeOrganizationSecurityPolicyAssociationConverter(name string, schema map[string]*schema.Schema) common.Converter {
	return &ComputeOrganizationSecurityPolicyAssociationConverter{
		name:   name,
		schema: schema,
	}
}

func (c *ComputeOrganizationSecurityPolicyAssociationConverter) Convert(assets []*caiasset.Asset) ([]*common.HCLResourceBlock, error) {
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

func (c *ComputeOrganizationSecurityPolicyAssociationConverter) convertResourceData(asset *caiasset.Asset, config *transport_tpg.Config) (*common.HCLResourceBlock, error) {
	if asset == nil || asset.Resource == nil || asset.Resource.Data == nil {
		return nil, fmt.Errorf("asset resource data is nil")
	}

	assetResourceData := asset.Resource.Data

	hcl, _ := resourceComputeOrganizationSecurityPolicyAssociationRead(assetResourceData, config)

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

func resourceComputeOrganizationSecurityPolicyAssociationRead(resource map[string]interface{}, config *transport_tpg.Config) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	var resource_data *schema.ResourceData = nil

	result["name"] = flattenComputeOrganizationSecurityPolicyAssociationName(resource["name"], resource_data, config)
	result["attachment_id"] = flattenComputeOrganizationSecurityPolicyAssociationAttachmentId(resource["attachmentId"], resource_data, config)
	result["display_name"] = flattenComputeOrganizationSecurityPolicyAssociationDisplayName(resource["displayName"], resource_data, config)

	return result, nil
}

func flattenComputeOrganizationSecurityPolicyAssociationName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeOrganizationSecurityPolicyAssociationAttachmentId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeOrganizationSecurityPolicyAssociationDisplayName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}
