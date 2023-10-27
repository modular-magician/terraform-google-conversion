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

package gkehub2

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/cai2hcl/common"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/caiasset"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const GKEHub2MembershipRBACRoleBindingAssetType string = "gkehub.googleapis.com/MembershipRBACRoleBinding"

const GKEHub2MembershipRBACRoleBindingAssetNameRegex string = "projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/memberships/(?P<membership_id>[^/]+)/rbacrolebindings"

type GKEHub2MembershipRBACRoleBindingConverter struct {
	name   string
	schema map[string]*schema.Schema
}

func NewGKEHub2MembershipRBACRoleBindingConverter(name string, schema map[string]*schema.Schema) common.Converter {
	return &GKEHub2MembershipRBACRoleBindingConverter{
		name:   name,
		schema: schema,
	}
}

func (c *GKEHub2MembershipRBACRoleBindingConverter) Convert(assets []*caiasset.Asset) ([]*common.HCLResourceBlock, error) {
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

func (c *GKEHub2MembershipRBACRoleBindingConverter) convertResourceData(asset *caiasset.Asset, config *transport_tpg.Config) (*common.HCLResourceBlock, error) {
	if asset == nil || asset.Resource == nil || asset.Resource.Data == nil {
		return nil, fmt.Errorf("asset resource data is nil")
	}

	assetResourceData := asset.Resource.Data

	hcl, _ := resourceGKEHub2MembershipRBACRoleBindingRead(assetResourceData, config)

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

func resourceGKEHub2MembershipRBACRoleBindingRead(resource map[string]interface{}, config *transport_tpg.Config) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	var resource_data *schema.ResourceData = nil

	result["name"] = flattenGKEHub2MembershipRBACRoleBindingName(resource["name"], resource_data, config)
	result["uid"] = flattenGKEHub2MembershipRBACRoleBindingUid(resource["uid"], resource_data, config)
	result["create_time"] = flattenGKEHub2MembershipRBACRoleBindingCreateTime(resource["createTime"], resource_data, config)
	result["update_time"] = flattenGKEHub2MembershipRBACRoleBindingUpdateTime(resource["updateTime"], resource_data, config)
	result["delete_time"] = flattenGKEHub2MembershipRBACRoleBindingDeleteTime(resource["deleteTime"], resource_data, config)
	result["state"] = flattenGKEHub2MembershipRBACRoleBindingState(resource["state"], resource_data, config)
	result["user"] = flattenGKEHub2MembershipRBACRoleBindingUser(resource["user"], resource_data, config)
	result["role"] = flattenGKEHub2MembershipRBACRoleBindingRole(resource["role"], resource_data, config)

	return result, nil
}

func flattenGKEHub2MembershipRBACRoleBindingName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGKEHub2MembershipRBACRoleBindingUid(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGKEHub2MembershipRBACRoleBindingCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGKEHub2MembershipRBACRoleBindingUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGKEHub2MembershipRBACRoleBindingDeleteTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGKEHub2MembershipRBACRoleBindingState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["code"] =
		flattenGKEHub2MembershipRBACRoleBindingStateCode(original["code"], d, config)
	return []interface{}{transformed}
}
func flattenGKEHub2MembershipRBACRoleBindingStateCode(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGKEHub2MembershipRBACRoleBindingUser(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGKEHub2MembershipRBACRoleBindingRole(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["predefined_role"] =
		flattenGKEHub2MembershipRBACRoleBindingRolePredefinedRole(original["predefinedRole"], d, config)
	return []interface{}{transformed}
}
func flattenGKEHub2MembershipRBACRoleBindingRolePredefinedRole(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}
