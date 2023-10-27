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

package secretmanager

import (
	"encoding/base64"
	"fmt"
	"regexp"
	"strings"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/cai2hcl/common"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/caiasset"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
	"google.golang.org/api/googleapi"
)

func resourceSecretManagerSecretVersionUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)

	_, err := expandSecretManagerSecretVersionEnabled(d.Get("enabled"), d, config)
	if err != nil {
		return err
	}

	return resourceSecretManagerSecretVersionRead(d, meta)
}

const SecretManagerSecretVersionAssetType string = "secretmanager.googleapis.com/SecretVersion"

const SecretManagerSecretVersionAssetNameRegex string = "(?P<name>[^/]+)"

type SecretManagerSecretVersionConverter struct {
	name   string
	schema map[string]*schema.Schema
}

func NewSecretManagerSecretVersionConverter(name string, schema map[string]*schema.Schema) common.Converter {
	return &SecretManagerSecretVersionConverter{
		name:   name,
		schema: schema,
	}
}

func (c *SecretManagerSecretVersionConverter) Convert(assets []*caiasset.Asset) ([]*common.HCLResourceBlock, error) {
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

func (c *SecretManagerSecretVersionConverter) convertResourceData(asset *caiasset.Asset, config *transport_tpg.Config) (*common.HCLResourceBlock, error) {
	if asset == nil || asset.Resource == nil || asset.Resource.Data == nil {
		return nil, fmt.Errorf("asset resource data is nil")
	}

	assetResourceData := asset.Resource.Data

	hcl, _ := resourceSecretManagerSecretVersionRead(assetResourceData, config)

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

func resourceSecretManagerSecretVersionRead(resource map[string]interface{}, config *transport_tpg.Config) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	var resource_data *schema.ResourceData = nil

	result["enabled"] = flattenSecretManagerSecretVersionEnabled(resource["state"], resource_data, config)
	result["name"] = flattenSecretManagerSecretVersionName(resource["name"], resource_data, config)
	result["version"] = flattenSecretManagerSecretVersionVersion(resource["version"], resource_data, config)
	result["create_time"] = flattenSecretManagerSecretVersionCreateTime(resource["createTime"], resource_data, config)
	result["destroy_time"] = flattenSecretManagerSecretVersionDestroyTime(resource["destroyTime"], resource_data, config)
	if flattenedProp := flattenSecretManagerSecretVersionPayload(resource["payload"], resource_data, config); flattenedProp != nil {
		if gerr, ok := flattenedProp.(*googleapi.Error); ok {
			return nil, fmt.Errorf("Error reading SecretVersion: %s", gerr)
		}
		casted := flattenedProp.([]interface{})[0]
		if casted != nil {
			for k, v := range casted.(map[string]interface{}) {
				result[k] = v
			}
		}
	}

	return result, nil
}

func flattenSecretManagerSecretVersionEnabled(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v.(string) == "ENABLED" {
		return true
	}

	return false
}

func flattenSecretManagerSecretVersionName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecretManagerSecretVersionVersion(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	name := d.Get("name").(string)
	secretRegex := regexp.MustCompile("projects/(.+)/secrets/(.+)/versions/(.+)$")

	parts := secretRegex.FindStringSubmatch(name)
	if len(parts) != 4 {
		panic(fmt.Sprintf("Version name does not fit the format `projects/{{project}}/secrets/{{secret}}/versions/{{version}}`"))
	}

	return parts[3]
}

func flattenSecretManagerSecretVersionCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecretManagerSecretVersionDestroyTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecretManagerSecretVersionPayload(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	transformed := make(map[string]interface{})

	// if this secret version is disabled, the api will return an error, as the value cannot be accessed, return what we have
	if d.Get("enabled").(bool) == false {
		transformed["secret_data"] = d.Get("secret_data")
		return []interface{}{transformed}
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{SecretManagerBasePath}}{{name}}:access")
	if err != nil {
		return err
	}

	parts := strings.Split(d.Get("name").(string), "/")
	project := parts[1]

	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	accessRes, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   project,
		RawURL:    url,
		UserAgent: userAgent,
	})
	if err != nil {
		return err
	}

	if d.Get("is_secret_data_base64").(bool) {
		transformed["secret_data"] = accessRes["payload"].(map[string]interface{})["data"].(string)
	} else {
		data, err := base64.StdEncoding.DecodeString(accessRes["payload"].(map[string]interface{})["data"].(string))
		if err != nil {
			return err
		}
		transformed["secret_data"] = string(data)
	}
	return []interface{}{transformed}
}
