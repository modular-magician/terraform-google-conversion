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

package filestore

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgresource"
	transport_tpg "github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/transport"
)

const FilestoreSnapshotAssetType string = "file.googleapis.com/Snapshot"

func ResourceConverterFilestoreSnapshot() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType: FilestoreSnapshotAssetType,
		Convert:   GetFilestoreSnapshotCaiObject,
	}
}

func GetFilestoreSnapshotCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	name, err := tpgresource.AssetName(d, config, "//file.googleapis.com/projects/{{project}}/locations/{{location}}/instances/{{instance}}/snapshots/{{name}}")
	if err != nil {
		return []tpgresource.Asset{}, err
	}
	if obj, err := GetFilestoreSnapshotApiObject(d, config); err == nil {
		return []tpgresource.Asset{{
			Name: name,
			Type: FilestoreSnapshotAssetType,
			Resource: &tpgresource.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/file/v1/rest",
				DiscoveryName:        "Snapshot",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []tpgresource.Asset{}, err
	}
}

func GetFilestoreSnapshotApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	descriptionProp, err := expandFilestoreSnapshotDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	labelsProp, err := expandFilestoreSnapshotLabels(d.Get("labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	return obj, nil
}

func expandFilestoreSnapshotDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFilestoreSnapshotLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
