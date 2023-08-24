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

package certificatemanager

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const CertificateManagerCertificateMapEntryAssetType string = "certificatemanager.googleapis.com/CertificateMapEntry"

func ResourceConverterCertificateManagerCertificateMapEntry() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: CertificateManagerCertificateMapEntryAssetType,
		Convert:   GetCertificateManagerCertificateMapEntryCaiObject,
	}
}

func GetCertificateManagerCertificateMapEntryCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//certificatemanager.googleapis.com/projects/{{project}}/locations/global/certificateMaps/{{map}}/certificateMapEntries/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetCertificateManagerCertificateMapEntryApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: CertificateManagerCertificateMapEntryAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/certificatemanager/v1/rest",
				DiscoveryName:        "CertificateMapEntry",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetCertificateManagerCertificateMapEntryApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	descriptionProp, err := expandCertificateManagerCertificateMapEntryDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	labelsProp, err := expandCertificateManagerCertificateMapEntryLabels(d.Get("labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	certificatesProp, err := expandCertificateManagerCertificateMapEntryCertificates(d.Get("certificates"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("certificates"); !tpgresource.IsEmptyValue(reflect.ValueOf(certificatesProp)) && (ok || !reflect.DeepEqual(v, certificatesProp)) {
		obj["certificates"] = certificatesProp
	}
	hostnameProp, err := expandCertificateManagerCertificateMapEntryHostname(d.Get("hostname"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("hostname"); !tpgresource.IsEmptyValue(reflect.ValueOf(hostnameProp)) && (ok || !reflect.DeepEqual(v, hostnameProp)) {
		obj["hostname"] = hostnameProp
	}
	matcherProp, err := expandCertificateManagerCertificateMapEntryMatcher(d.Get("matcher"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("matcher"); !tpgresource.IsEmptyValue(reflect.ValueOf(matcherProp)) && (ok || !reflect.DeepEqual(v, matcherProp)) {
		obj["matcher"] = matcherProp
	}
	nameProp, err := expandCertificateManagerCertificateMapEntryName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}

	return obj, nil
}

func expandCertificateManagerCertificateMapEntryDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCertificateManagerCertificateMapEntryLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandCertificateManagerCertificateMapEntryCertificates(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCertificateManagerCertificateMapEntryHostname(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCertificateManagerCertificateMapEntryMatcher(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCertificateManagerCertificateMapEntryName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return tpgresource.GetResourceNameFromSelfLink(v.(string)), nil
}
