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

package compute

import (
	"fmt"
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/id"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const ComputeRegionSslCertificateAssetType string = "compute.googleapis.com/RegionSslCertificate"

func ResourceConverterComputeRegionSslCertificate() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: ComputeRegionSslCertificateAssetType,
		Convert:   GetComputeRegionSslCertificateCaiObject,
	}
}

func GetComputeRegionSslCertificateCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//compute.googleapis.com/projects/{{project}}/regions/{{region}}/sslCertificates/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetComputeRegionSslCertificateApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: ComputeRegionSslCertificateAssetType,
			Resource: &cai.AssetResource{
				Version:              "beta",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/compute/beta/rest",
				DiscoveryName:        "RegionSslCertificate",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetComputeRegionSslCertificateApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	certificateProp, err := expandComputeRegionSslCertificateCertificate(d.Get("certificate"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("certificate"); !tpgresource.IsEmptyValue(reflect.ValueOf(certificateProp)) && (ok || !reflect.DeepEqual(v, certificateProp)) {
		obj["certificate"] = certificateProp
	}
	descriptionProp, err := expandComputeRegionSslCertificateDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	nameProp, err := expandComputeRegionSslCertificateName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	privateKeyProp, err := expandComputeRegionSslCertificatePrivateKey(d.Get("private_key"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("private_key"); !tpgresource.IsEmptyValue(reflect.ValueOf(privateKeyProp)) && (ok || !reflect.DeepEqual(v, privateKeyProp)) {
		obj["privateKey"] = privateKeyProp
	}
	regionProp, err := expandComputeRegionSslCertificateRegion(d.Get("region"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("region"); !tpgresource.IsEmptyValue(reflect.ValueOf(regionProp)) && (ok || !reflect.DeepEqual(v, regionProp)) {
		obj["region"] = regionProp
	}

	return obj, nil
}

func expandComputeRegionSslCertificateCertificate(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionSslCertificateDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionSslCertificateName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	var certName string
	if v, ok := d.GetOk("name"); ok {
		certName = v.(string)
	} else if v, ok := d.GetOk("name_prefix"); ok {
		certName = id.PrefixedUniqueId(v.(string))
	} else {
		certName = id.UniqueId()
	}

	// We need to get the {{name}} into schema to set the ID using tpgresource.ReplaceVars
	if err := d.Set("name", certName); err != nil {
		return nil, fmt.Errorf("Error setting name: %s", err)
	}

	return certName, nil
}

func expandComputeRegionSslCertificatePrivateKey(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionSslCertificateRegion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	f, err := tpgresource.ParseGlobalFieldValue("regions", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for region: %s", err)
	}
	return f.RelativeLink(), nil
}
