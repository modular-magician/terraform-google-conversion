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

package kms

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/cai2hcl/common"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/caiasset"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const KMSCryptoKeyAssetType string = "cloudkms.googleapis.com/CryptoKey"

const KMSCryptoKeyAssetNameRegex string = "(?P<key_ring>[^/]+)/cryptoKeys"

type KMSCryptoKeyConverter struct {
	name   string
	schema map[string]*schema.Schema
}

func NewKMSCryptoKeyConverter(name string, schema map[string]*schema.Schema) common.Converter {
	return &KMSCryptoKeyConverter{
		name:   name,
		schema: schema,
	}
}

func (c *KMSCryptoKeyConverter) Convert(assets []*caiasset.Asset) ([]*common.HCLResourceBlock, error) {
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

func (c *KMSCryptoKeyConverter) convertResourceData(asset *caiasset.Asset, config *transport_tpg.Config) (*common.HCLResourceBlock, error) {
	if asset == nil || asset.Resource == nil || asset.Resource.Data == nil {
		return nil, fmt.Errorf("asset resource data is nil")
	}

	assetResourceData := asset.Resource.Data

	hcl, _ := resourceKMSCryptoKeyRead(assetResourceData, config)

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

func resourceKMSCryptoKeyRead(resource map[string]interface{}, config *transport_tpg.Config) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	var resource_data *schema.ResourceData = nil

	result["labels"] = flattenKMSCryptoKeyLabels(resource["labels"], resource_data, config)
	result["purpose"] = flattenKMSCryptoKeyPurpose(resource["purpose"], resource_data, config)
	result["rotation_period"] = flattenKMSCryptoKeyRotationPeriod(resource["rotationPeriod"], resource_data, config)
	result["version_template"] = flattenKMSCryptoKeyVersionTemplate(resource["versionTemplate"], resource_data, config)
	result["destroy_scheduled_duration"] = flattenKMSCryptoKeyDestroyScheduledDuration(resource["destroyScheduledDuration"], resource_data, config)
	result["import_only"] = flattenKMSCryptoKeyImportOnly(resource["importOnly"], resource_data, config)
	result["terraform_labels"] = flattenKMSCryptoKeyTerraformLabels(resource["labels"], resource_data, config)
	result["effective_labels"] = flattenKMSCryptoKeyEffectiveLabels(resource["labels"], resource_data, config)

	return result, nil
}

func flattenKMSCryptoKeyLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenKMSCryptoKeyPurpose(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenKMSCryptoKeyRotationPeriod(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenKMSCryptoKeyVersionTemplate(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["algorithm"] =
		flattenKMSCryptoKeyVersionTemplateAlgorithm(original["algorithm"], d, config)
	transformed["protection_level"] =
		flattenKMSCryptoKeyVersionTemplateProtectionLevel(original["protectionLevel"], d, config)
	return []interface{}{transformed}
}
func flattenKMSCryptoKeyVersionTemplateAlgorithm(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenKMSCryptoKeyVersionTemplateProtectionLevel(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenKMSCryptoKeyDestroyScheduledDuration(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenKMSCryptoKeyImportOnly(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenKMSCryptoKeyTerraformLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenKMSCryptoKeyEffectiveLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}
