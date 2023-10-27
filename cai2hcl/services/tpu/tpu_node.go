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

package tpu

import (
	"context"
	"fmt"
	"regexp"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/cai2hcl/common"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/caiasset"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

// compareTpuNodeSchedulingConfig diff suppresses for the default
// scheduling, i.e. if preemptible is false, the API may either return no
// schedulingConfig or an empty schedulingConfig.
func compareTpuNodeSchedulingConfig(k, old, new string, d *schema.ResourceData) bool {
	if k == "scheduling_config.0.preemptible" {
		return old == "" && new == "false"
	}
	if k == "scheduling_config.#" {
		o, n := d.GetChange("scheduling_config.0.preemptible")
		return o.(bool) == n.(bool)
	}
	return false
}

func tpuNodeCustomizeDiff(_ context.Context, diff *schema.ResourceDiff, meta interface{}) error {
	old, new := diff.GetChange("network")
	config := meta.(*transport_tpg.Config)

	networkLinkRegex := regexp.MustCompile("projects/(.+)/global/networks/(.+)")

	var pid string

	if networkLinkRegex.MatchString(new.(string)) {
		parts := networkLinkRegex.FindStringSubmatch(new.(string))
		pid = parts[1]
	}

	if pid == "" {
		return nil
	}

	project, err := config.NewResourceManagerClient(config.UserAgent).Projects.Get(pid).Do()
	if err != nil {
		return fmt.Errorf("Failed to retrieve project, pid: %s, err: %s", pid, err)
	}

	if networkLinkRegex.MatchString(old.(string)) {
		parts := networkLinkRegex.FindStringSubmatch(old.(string))
		i, err := tpgresource.StringToFixed64(parts[1])
		if err == nil {
			if project.ProjectNumber == i {
				if err := diff.SetNew("network", old); err != nil {
					return err
				}
				return nil
			}
		}
	}
	return nil
}

const TPUNodeAssetType string = "tpu.googleapis.com/Node"

const TPUNodeAssetNameRegex string = "projects/(?P<project>[^/]+)/locations/(?P<zone>[^/]+)/nodes"

type TPUNodeConverter struct {
	name   string
	schema map[string]*schema.Schema
}

func NewTPUNodeConverter(name string, schema map[string]*schema.Schema) common.Converter {
	return &TPUNodeConverter{
		name:   name,
		schema: schema,
	}
}

func (c *TPUNodeConverter) Convert(assets []*caiasset.Asset) ([]*common.HCLResourceBlock, error) {
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

func (c *TPUNodeConverter) convertResourceData(asset *caiasset.Asset, config *transport_tpg.Config) (*common.HCLResourceBlock, error) {
	if asset == nil || asset.Resource == nil || asset.Resource.Data == nil {
		return nil, fmt.Errorf("asset resource data is nil")
	}

	assetResourceData := asset.Resource.Data

	hcl, _ := resourceTPUNodeRead(assetResourceData, config)

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

func resourceTPUNodeRead(resource map[string]interface{}, config *transport_tpg.Config) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	var resource_data *schema.ResourceData = nil

	result["name"] = flattenTPUNodeName(resource["name"], resource_data, config)
	result["description"] = flattenTPUNodeDescription(resource["description"], resource_data, config)
	result["accelerator_type"] = flattenTPUNodeAcceleratorType(resource["acceleratorType"], resource_data, config)
	result["tensorflow_version"] = flattenTPUNodeTensorflowVersion(resource["tensorflowVersion"], resource_data, config)
	result["network"] = flattenTPUNodeNetwork(resource["network"], resource_data, config)
	result["cidr_block"] = flattenTPUNodeCidrBlock(resource["cidrBlock"], resource_data, config)
	result["service_account"] = flattenTPUNodeServiceAccount(resource["serviceAccount"], resource_data, config)
	result["use_service_networking"] = flattenTPUNodeUseServiceNetworking(resource["useServiceNetworking"], resource_data, config)
	result["scheduling_config"] = flattenTPUNodeSchedulingConfig(resource["schedulingConfig"], resource_data, config)
	result["network_endpoints"] = flattenTPUNodeNetworkEndpoints(resource["networkEndpoints"], resource_data, config)
	result["labels"] = flattenTPUNodeLabels(resource["labels"], resource_data, config)
	result["terraform_labels"] = flattenTPUNodeTerraformLabels(resource["labels"], resource_data, config)
	result["effective_labels"] = flattenTPUNodeEffectiveLabels(resource["labels"], resource_data, config)

	return result, nil
}

func flattenTPUNodeName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.NameFromSelfLinkStateFunc(v)
}

func flattenTPUNodeDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenTPUNodeAcceleratorType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenTPUNodeTensorflowVersion(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenTPUNodeNetwork(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenTPUNodeCidrBlock(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenTPUNodeServiceAccount(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenTPUNodeUseServiceNetworking(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenTPUNodeSchedulingConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["preemptible"] =
		flattenTPUNodeSchedulingConfigPreemptible(original["preemptible"], d, config)
	return []interface{}{transformed}
}
func flattenTPUNodeSchedulingConfigPreemptible(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenTPUNodeNetworkEndpoints(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed = append(transformed, map[string]interface{}{
			"ip_address": flattenTPUNodeNetworkEndpointsIpAddress(original["ipAddress"], d, config),
			"port":       flattenTPUNodeNetworkEndpointsPort(original["port"], d, config),
		})
	}
	return transformed
}
func flattenTPUNodeNetworkEndpointsIpAddress(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenTPUNodeNetworkEndpointsPort(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := tpgresource.StringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenTPUNodeLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenTPUNodeTerraformLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenTPUNodeEffectiveLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}
