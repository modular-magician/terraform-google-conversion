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

package apigee

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const ApigeeEnvironmentAssetType string = "apigee.googleapis.com/Environment"

func ResourceConverterApigeeEnvironment() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType: ApigeeEnvironmentAssetType,
		Convert:   GetApigeeEnvironmentCaiObject,
	}
}

func GetApigeeEnvironmentCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	name, err := tpgresource.AssetName(d, config, "//apigee.googleapis.com/{{org_id}}/environments/{{name}}")
	if err != nil {
		return []tpgresource.Asset{}, err
	}
	if obj, err := GetApigeeEnvironmentApiObject(d, config); err == nil {
		return []tpgresource.Asset{{
			Name: name,
			Type: ApigeeEnvironmentAssetType,
			Resource: &tpgresource.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/apigee/v1/rest",
				DiscoveryName:        "Environment",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []tpgresource.Asset{}, err
	}
}

func GetApigeeEnvironmentApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandApigeeEnvironmentName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	displayNameProp, err := expandApigeeEnvironmentDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	descriptionProp, err := expandApigeeEnvironmentDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	deploymentTypeProp, err := expandApigeeEnvironmentDeploymentType(d.Get("deployment_type"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("deployment_type"); !tpgresource.IsEmptyValue(reflect.ValueOf(deploymentTypeProp)) && (ok || !reflect.DeepEqual(v, deploymentTypeProp)) {
		obj["deploymentType"] = deploymentTypeProp
	}
	apiProxyTypeProp, err := expandApigeeEnvironmentApiProxyType(d.Get("api_proxy_type"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("api_proxy_type"); !tpgresource.IsEmptyValue(reflect.ValueOf(apiProxyTypeProp)) && (ok || !reflect.DeepEqual(v, apiProxyTypeProp)) {
		obj["apiProxyType"] = apiProxyTypeProp
	}
	nodeConfigProp, err := expandApigeeEnvironmentNodeConfig(d.Get("node_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("node_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(nodeConfigProp)) && (ok || !reflect.DeepEqual(v, nodeConfigProp)) {
		obj["nodeConfig"] = nodeConfigProp
	}

	return obj, nil
}

func expandApigeeEnvironmentName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeEnvironmentDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeEnvironmentDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeEnvironmentDeploymentType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeEnvironmentApiProxyType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeEnvironmentNodeConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedMinNodeCount, err := expandApigeeEnvironmentNodeConfigMinNodeCount(original["min_node_count"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMinNodeCount); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["minNodeCount"] = transformedMinNodeCount
	}

	transformedMaxNodeCount, err := expandApigeeEnvironmentNodeConfigMaxNodeCount(original["max_node_count"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxNodeCount); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["maxNodeCount"] = transformedMaxNodeCount
	}

	transformedCurrentAggregateNodeCount, err := expandApigeeEnvironmentNodeConfigCurrentAggregateNodeCount(original["current_aggregate_node_count"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCurrentAggregateNodeCount); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["currentAggregateNodeCount"] = transformedCurrentAggregateNodeCount
	}

	return transformed, nil
}

func expandApigeeEnvironmentNodeConfigMinNodeCount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeEnvironmentNodeConfigMaxNodeCount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeEnvironmentNodeConfigCurrentAggregateNodeCount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
