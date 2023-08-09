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

package gameservices

import (
	"log"
	"reflect"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func suppressSuffixDiff(_, old, new string, _ *schema.ResourceData) bool {
	if strings.HasSuffix(old, new) {
		log.Printf("[INFO] suppressing diff as %s is the same as the full path of %s", new, old)
		return true
	}

	return false
}

const GameServicesGameServerClusterAssetType string = "gameservices.googleapis.com/GameServerCluster"

func ResourceConverterGameServicesGameServerCluster() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType: GameServicesGameServerClusterAssetType,
		Convert:   GetGameServicesGameServerClusterCaiObject,
	}
}

func GetGameServicesGameServerClusterCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	name, err := tpgresource.AssetName(d, config, "//gameservices.googleapis.com/projects/{{project}}/locations/{{location}}/realms/{{realm_id}}/gameServerClusters/{{cluster_id}}")
	if err != nil {
		return []tpgresource.Asset{}, err
	}
	if obj, err := GetGameServicesGameServerClusterApiObject(d, config); err == nil {
		return []tpgresource.Asset{{
			Name: name,
			Type: GameServicesGameServerClusterAssetType,
			Resource: &tpgresource.AssetResource{
				Version:              "v1beta",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/gameservices/v1beta/rest",
				DiscoveryName:        "GameServerCluster",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []tpgresource.Asset{}, err
	}
}

func GetGameServicesGameServerClusterApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	labelsProp, err := expandGameServicesGameServerClusterLabels(d.Get("labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	connectionInfoProp, err := expandGameServicesGameServerClusterConnectionInfo(d.Get("connection_info"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("connection_info"); !tpgresource.IsEmptyValue(reflect.ValueOf(connectionInfoProp)) && (ok || !reflect.DeepEqual(v, connectionInfoProp)) {
		obj["connectionInfo"] = connectionInfoProp
	}
	descriptionProp, err := expandGameServicesGameServerClusterDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}

	return obj, nil
}

func expandGameServicesGameServerClusterLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandGameServicesGameServerClusterConnectionInfo(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedGkeClusterReference, err := expandGameServicesGameServerClusterConnectionInfoGkeClusterReference(original["gke_cluster_reference"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedGkeClusterReference); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["gkeClusterReference"] = transformedGkeClusterReference
	}

	transformedNamespace, err := expandGameServicesGameServerClusterConnectionInfoNamespace(original["namespace"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNamespace); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["namespace"] = transformedNamespace
	}

	return transformed, nil
}

func expandGameServicesGameServerClusterConnectionInfoGkeClusterReference(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedCluster, err := expandGameServicesGameServerClusterConnectionInfoGkeClusterReferenceCluster(original["cluster"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCluster); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["cluster"] = transformedCluster
	}

	return transformed, nil
}

func expandGameServicesGameServerClusterConnectionInfoGkeClusterReferenceCluster(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGameServicesGameServerClusterConnectionInfoNamespace(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGameServicesGameServerClusterDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
