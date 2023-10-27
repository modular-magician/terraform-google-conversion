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

package gkehub

import (
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/cai2hcl/common"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/caiasset"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func suppressGkeHubEndpointSelfLinkDiff(_, old, new string, _ *schema.ResourceData) bool {
	// The custom expander injects //container.googleapis.com/ if a selflink is supplied.
	selfLink := strings.TrimPrefix(old, "//container.googleapis.com/")
	if selfLink == new {
		return true
	}

	return false
}

const GKEHubMembershipAssetType string = "gkehub.googleapis.com/Membership"

const GKEHubMembershipAssetNameRegex string = "projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/memberships"

type GKEHubMembershipConverter struct {
	name   string
	schema map[string]*schema.Schema
}

func NewGKEHubMembershipConverter(name string, schema map[string]*schema.Schema) common.Converter {
	return &GKEHubMembershipConverter{
		name:   name,
		schema: schema,
	}
}

func (c *GKEHubMembershipConverter) Convert(assets []*caiasset.Asset) ([]*common.HCLResourceBlock, error) {
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

func (c *GKEHubMembershipConverter) convertResourceData(asset *caiasset.Asset, config *transport_tpg.Config) (*common.HCLResourceBlock, error) {
	if asset == nil || asset.Resource == nil || asset.Resource.Data == nil {
		return nil, fmt.Errorf("asset resource data is nil")
	}

	assetResourceData := asset.Resource.Data

	hcl, _ := resourceGKEHubMembershipRead(assetResourceData, config)

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

func resourceGKEHubMembershipRead(resource map[string]interface{}, config *transport_tpg.Config) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	var resource_data *schema.ResourceData = nil

	result["name"] = flattenGKEHubMembershipName(resource["name"], resource_data, config)
	result["description"] = flattenGKEHubMembershipDescription(resource["description"], resource_data, config)
	result["labels"] = flattenGKEHubMembershipLabels(resource["labels"], resource_data, config)
	result["endpoint"] = flattenGKEHubMembershipEndpoint(resource["endpoint"], resource_data, config)
	result["authority"] = flattenGKEHubMembershipAuthority(resource["authority"], resource_data, config)
	result["terraform_labels"] = flattenGKEHubMembershipTerraformLabels(resource["labels"], resource_data, config)
	result["effective_labels"] = flattenGKEHubMembershipEffectiveLabels(resource["labels"], resource_data, config)

	return result, nil
}

func flattenGKEHubMembershipName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGKEHubMembershipDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGKEHubMembershipLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenGKEHubMembershipEndpoint(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["gke_cluster"] =
		flattenGKEHubMembershipEndpointGkeCluster(original["gkeCluster"], d, config)
	return []interface{}{transformed}
}
func flattenGKEHubMembershipEndpointGkeCluster(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["resource_link"] =
		flattenGKEHubMembershipEndpointGkeClusterResourceLink(original["resourceLink"], d, config)
	return []interface{}{transformed}
}
func flattenGKEHubMembershipEndpointGkeClusterResourceLink(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGKEHubMembershipAuthority(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["issuer"] =
		flattenGKEHubMembershipAuthorityIssuer(original["issuer"], d, config)
	return []interface{}{transformed}
}
func flattenGKEHubMembershipAuthorityIssuer(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGKEHubMembershipTerraformLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenGKEHubMembershipEffectiveLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}
