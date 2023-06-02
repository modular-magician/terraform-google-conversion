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

package google

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgresource"
	transport_tpg "github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/transport"
)

const CertificateManagerCertificateIssuanceConfigAssetType string = "certificatemanager.googleapis.com/CertificateIssuanceConfig"

func resourceConverterCertificateManagerCertificateIssuanceConfig() ResourceConverter {
	return ResourceConverter{
		AssetType: CertificateManagerCertificateIssuanceConfigAssetType,
		Convert:   GetCertificateManagerCertificateIssuanceConfigCaiObject,
	}
}

func GetCertificateManagerCertificateIssuanceConfigCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]Asset, error) {
	name, err := assetName(d, config, "//certificatemanager.googleapis.com/projects/{{project}}/locations/{{location}}/certificateIssuanceConfigs/{{name}}")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetCertificateManagerCertificateIssuanceConfigApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: CertificateManagerCertificateIssuanceConfigAssetType,
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/certificatemanager/v1/rest",
				DiscoveryName:        "CertificateIssuanceConfig",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
	}
}

func GetCertificateManagerCertificateIssuanceConfigApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	descriptionProp, err := expandCertificateManagerCertificateIssuanceConfigDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	rotationWindowPercentageProp, err := expandCertificateManagerCertificateIssuanceConfigRotationWindowPercentage(d.Get("rotation_window_percentage"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("rotation_window_percentage"); !tpgresource.IsEmptyValue(reflect.ValueOf(rotationWindowPercentageProp)) && (ok || !reflect.DeepEqual(v, rotationWindowPercentageProp)) {
		obj["rotationWindowPercentage"] = rotationWindowPercentageProp
	}
	keyAlgorithmProp, err := expandCertificateManagerCertificateIssuanceConfigKeyAlgorithm(d.Get("key_algorithm"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("key_algorithm"); !tpgresource.IsEmptyValue(reflect.ValueOf(keyAlgorithmProp)) && (ok || !reflect.DeepEqual(v, keyAlgorithmProp)) {
		obj["keyAlgorithm"] = keyAlgorithmProp
	}
	lifetimeProp, err := expandCertificateManagerCertificateIssuanceConfigLifetime(d.Get("lifetime"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("lifetime"); !tpgresource.IsEmptyValue(reflect.ValueOf(lifetimeProp)) && (ok || !reflect.DeepEqual(v, lifetimeProp)) {
		obj["lifetime"] = lifetimeProp
	}
	labelsProp, err := expandCertificateManagerCertificateIssuanceConfigLabels(d.Get("labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	certificateAuthorityConfigProp, err := expandCertificateManagerCertificateIssuanceConfigCertificateAuthorityConfig(d.Get("certificate_authority_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("certificate_authority_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(certificateAuthorityConfigProp)) && (ok || !reflect.DeepEqual(v, certificateAuthorityConfigProp)) {
		obj["certificateAuthorityConfig"] = certificateAuthorityConfigProp
	}

	return obj, nil
}

func expandCertificateManagerCertificateIssuanceConfigDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCertificateManagerCertificateIssuanceConfigRotationWindowPercentage(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCertificateManagerCertificateIssuanceConfigKeyAlgorithm(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCertificateManagerCertificateIssuanceConfigLifetime(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCertificateManagerCertificateIssuanceConfigLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandCertificateManagerCertificateIssuanceConfigCertificateAuthorityConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedCertificateAuthorityServiceConfig, err := expandCertificateManagerCertificateIssuanceConfigCertificateAuthorityConfigCertificateAuthorityServiceConfig(original["certificate_authority_service_config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCertificateAuthorityServiceConfig); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["certificateAuthorityServiceConfig"] = transformedCertificateAuthorityServiceConfig
	}

	return transformed, nil
}

func expandCertificateManagerCertificateIssuanceConfigCertificateAuthorityConfigCertificateAuthorityServiceConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedCaPool, err := expandCertificateManagerCertificateIssuanceConfigCertificateAuthorityConfigCertificateAuthorityServiceConfigCaPool(original["ca_pool"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCaPool); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["caPool"] = transformedCaPool
	}

	return transformed, nil
}

func expandCertificateManagerCertificateIssuanceConfigCertificateAuthorityConfigCertificateAuthorityServiceConfigCaPool(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
