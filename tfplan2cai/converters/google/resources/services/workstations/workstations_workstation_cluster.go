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

package workstations

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgresource"
	transport_tpg "github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/transport"
)

const WorkstationsWorkstationClusterAssetType string = "workstations.googleapis.com/WorkstationCluster"

func ResourceConverterWorkstationsWorkstationCluster() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType: WorkstationsWorkstationClusterAssetType,
		Convert:   GetWorkstationsWorkstationClusterCaiObject,
	}
}

func GetWorkstationsWorkstationClusterCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	name, err := tpgresource.AssetName(d, config, "//workstations.googleapis.com/projects/{{project}}/locations/{{location}}/workstationClusters/{{workstation_cluster_id}}")
	if err != nil {
		return []tpgresource.Asset{}, err
	}
	if obj, err := GetWorkstationsWorkstationClusterApiObject(d, config); err == nil {
		return []tpgresource.Asset{{
			Name: name,
			Type: WorkstationsWorkstationClusterAssetType,
			Resource: &tpgresource.AssetResource{
				Version:              "v1beta",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/workstations/v1beta/rest",
				DiscoveryName:        "WorkstationCluster",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []tpgresource.Asset{}, err
	}
}

func GetWorkstationsWorkstationClusterApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	labelsProp, err := expandWorkstationsWorkstationClusterLabels(d.Get("labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	networkProp, err := expandWorkstationsWorkstationClusterNetwork(d.Get("network"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("network"); !tpgresource.IsEmptyValue(reflect.ValueOf(networkProp)) && (ok || !reflect.DeepEqual(v, networkProp)) {
		obj["network"] = networkProp
	}
	subnetworkProp, err := expandWorkstationsWorkstationClusterSubnetwork(d.Get("subnetwork"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("subnetwork"); !tpgresource.IsEmptyValue(reflect.ValueOf(subnetworkProp)) && (ok || !reflect.DeepEqual(v, subnetworkProp)) {
		obj["subnetwork"] = subnetworkProp
	}
	displayNameProp, err := expandWorkstationsWorkstationClusterDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	annotationsProp, err := expandWorkstationsWorkstationClusterAnnotations(d.Get("annotations"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("annotations"); !tpgresource.IsEmptyValue(reflect.ValueOf(annotationsProp)) && (ok || !reflect.DeepEqual(v, annotationsProp)) {
		obj["annotations"] = annotationsProp
	}
	etagProp, err := expandWorkstationsWorkstationClusterEtag(d.Get("etag"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("etag"); !tpgresource.IsEmptyValue(reflect.ValueOf(etagProp)) && (ok || !reflect.DeepEqual(v, etagProp)) {
		obj["etag"] = etagProp
	}
	privateClusterConfigProp, err := expandWorkstationsWorkstationClusterPrivateClusterConfig(d.Get("private_cluster_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("private_cluster_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(privateClusterConfigProp)) && (ok || !reflect.DeepEqual(v, privateClusterConfigProp)) {
		obj["privateClusterConfig"] = privateClusterConfigProp
	}

	return obj, nil
}

func expandWorkstationsWorkstationClusterLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandWorkstationsWorkstationClusterNetwork(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandWorkstationsWorkstationClusterSubnetwork(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandWorkstationsWorkstationClusterDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandWorkstationsWorkstationClusterAnnotations(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandWorkstationsWorkstationClusterEtag(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandWorkstationsWorkstationClusterPrivateClusterConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedEnablePrivateEndpoint, err := expandWorkstationsWorkstationClusterPrivateClusterConfigEnablePrivateEndpoint(original["enable_private_endpoint"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEnablePrivateEndpoint); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["enablePrivateEndpoint"] = transformedEnablePrivateEndpoint
	}

	transformedClusterHostname, err := expandWorkstationsWorkstationClusterPrivateClusterConfigClusterHostname(original["cluster_hostname"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedClusterHostname); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["clusterHostname"] = transformedClusterHostname
	}

	transformedServiceAttachmentUri, err := expandWorkstationsWorkstationClusterPrivateClusterConfigServiceAttachmentUri(original["service_attachment_uri"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedServiceAttachmentUri); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["serviceAttachmentUri"] = transformedServiceAttachmentUri
	}

	return transformed, nil
}

func expandWorkstationsWorkstationClusterPrivateClusterConfigEnablePrivateEndpoint(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandWorkstationsWorkstationClusterPrivateClusterConfigClusterHostname(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandWorkstationsWorkstationClusterPrivateClusterConfigServiceAttachmentUri(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
