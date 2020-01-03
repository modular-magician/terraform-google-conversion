// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
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
	"encoding/base64"
	"reflect"
)

func GetKMSSecretCiphertextCaiObject(d TerraformResourceData, config *Config) (Asset, error) {
	name, err := assetName(d, config, "//kms.googleapis.com/{{crypto_key}}")
	if err != nil {
		return Asset{}, err
	}
	if obj, err := GetKMSSecretCiphertextApiObject(d, config); err == nil {
		return Asset{
			Name: name,
			Type: "kms.googleapis.com/SecretCiphertext",
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/kms/v1/rest",
				DiscoveryName:        "SecretCiphertext",
				Data:                 obj,
			},
		}, nil
	} else {
		return Asset{}, err
	}
}

func GetKMSSecretCiphertextApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	plaintextProp, err := expandKMSSecretCiphertextPlaintext(d.Get("plaintext"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("plaintext"); !isEmptyValue(reflect.ValueOf(plaintextProp)) && (ok || !reflect.DeepEqual(v, plaintextProp)) {
		obj["plaintext"] = plaintextProp
	}

	return obj, nil
}

func expandKMSSecretCiphertextPlaintext(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	if v == nil {
		return nil, nil
	}

	return base64.StdEncoding.EncodeToString([]byte(v.(string))), nil
}
