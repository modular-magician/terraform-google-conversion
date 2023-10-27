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

package accessapproval

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/cai2hcl/common"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/caiasset"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const AccessApprovalProjectSettingsAssetType string = "accessapproval.googleapis.com/ProjectSettings"

const AccessApprovalProjectSettingsAssetNameRegex string = "projects/(?P<project_id>[^/]+)/accessApprovalSettings"

type AccessApprovalProjectSettingsConverter struct {
	name   string
	schema map[string]*schema.Schema
}

func NewAccessApprovalProjectSettingsConverter(name string, schema map[string]*schema.Schema) common.Converter {
	return &AccessApprovalProjectSettingsConverter{
		name:   name,
		schema: schema,
	}
}

func (c *AccessApprovalProjectSettingsConverter) Convert(assets []*caiasset.Asset) ([]*common.HCLResourceBlock, error) {
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

func (c *AccessApprovalProjectSettingsConverter) convertResourceData(asset *caiasset.Asset, config *transport_tpg.Config) (*common.HCLResourceBlock, error) {
	if asset == nil || asset.Resource == nil || asset.Resource.Data == nil {
		return nil, fmt.Errorf("asset resource data is nil")
	}

	assetResourceData := asset.Resource.Data

	hcl, _ := resourceAccessApprovalProjectSettingsRead(assetResourceData, config)

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

func resourceAccessApprovalProjectSettingsRead(resource map[string]interface{}, config *transport_tpg.Config) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	var resource_data *schema.ResourceData = nil

	result["name"] = flattenAccessApprovalProjectSettingsName(resource["name"], resource_data, config)
	result["notification_emails"] = flattenAccessApprovalProjectSettingsNotificationEmails(resource["notificationEmails"], resource_data, config)
	result["enrolled_services"] = flattenAccessApprovalProjectSettingsEnrolledServices(resource["enrolledServices"], resource_data, config)
	result["enrolled_ancestor"] = flattenAccessApprovalProjectSettingsEnrolledAncestor(resource["enrolledAncestor"], resource_data, config)
	result["active_key_version"] = flattenAccessApprovalProjectSettingsActiveKeyVersion(resource["activeKeyVersion"], resource_data, config)
	result["ancestor_has_active_key_version"] = flattenAccessApprovalProjectSettingsAncestorHasActiveKeyVersion(resource["ancestorHasActiveKeyVersion"], resource_data, config)
	result["invalid_key_version"] = flattenAccessApprovalProjectSettingsInvalidKeyVersion(resource["invalidKeyVersion"], resource_data, config)
	result["project"] = flattenAccessApprovalProjectSettingsProject(resource["project"], resource_data, config)

	return result, nil
}

func flattenAccessApprovalProjectSettingsName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAccessApprovalProjectSettingsNotificationEmails(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return schema.NewSet(schema.HashString, v.([]interface{}))
}

func flattenAccessApprovalProjectSettingsEnrolledServices(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := schema.NewSet(accessApprovalEnrolledServicesHash, []interface{}{})
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed.Add(map[string]interface{}{
			"cloud_product":    flattenAccessApprovalProjectSettingsEnrolledServicesCloudProduct(original["cloudProduct"], d, config),
			"enrollment_level": flattenAccessApprovalProjectSettingsEnrolledServicesEnrollmentLevel(original["enrollmentLevel"], d, config),
		})
	}
	return transformed
}
func flattenAccessApprovalProjectSettingsEnrolledServicesCloudProduct(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAccessApprovalProjectSettingsEnrolledServicesEnrollmentLevel(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAccessApprovalProjectSettingsEnrolledAncestor(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAccessApprovalProjectSettingsActiveKeyVersion(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAccessApprovalProjectSettingsAncestorHasActiveKeyVersion(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAccessApprovalProjectSettingsInvalidKeyVersion(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAccessApprovalProjectSettingsProject(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}
