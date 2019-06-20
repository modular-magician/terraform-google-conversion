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
	"fmt"
	"reflect"
	"time"
)

func GetKmsCryptoKeyCaiObject(d TerraformResourceData, config *Config) (Asset, error) {
	name, err := assetName(d, config, "//kms.googleapis.com/projects/{{projects}}/keyRings/{{key_ring}}/cryptoKeys/{{name}}")
	if err != nil {
		return Asset{}, err
	}
	if obj, err := GetKmsCryptoKeyApiObject(d, config); err == nil {
		return Asset{
			Name: name,
			Type: "kms.googleapis.com/CryptoKey",
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/kms/v1/rest",
				DiscoveryName:        "CryptoKey",
				Data:                 obj,
			},
		}, nil
	} else {
		return Asset{}, err
	}
}

func GetKmsCryptoKeyApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	purposeProp, err := expandKmsCryptoKeyPurpose(d.Get("purpose"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("purpose"); !isEmptyValue(reflect.ValueOf(purposeProp)) && (ok || !reflect.DeepEqual(v, purposeProp)) {
		obj["purpose"] = purposeProp
	}
	rotationPeriodProp, err := expandKmsCryptoKeyRotationPeriod(d.Get("rotation_period"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("rotation_period"); !isEmptyValue(reflect.ValueOf(rotationPeriodProp)) && (ok || !reflect.DeepEqual(v, rotationPeriodProp)) {
		obj["rotationPeriod"] = rotationPeriodProp
	}
	versionTemplateProp, err := expandKmsCryptoKeyVersionTemplate(d.Get("version_template"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("version_template"); !isEmptyValue(reflect.ValueOf(versionTemplateProp)) && (ok || !reflect.DeepEqual(v, versionTemplateProp)) {
		obj["versionTemplate"] = versionTemplateProp
	}

	return resourceKmsCryptoKeyEncoder(d, config, obj)
}

func resourceKmsCryptoKeyEncoder(d TerraformResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	// if rotationPeriod is set, nextRotationTime must also be set.
	if d.Get("rotation_period") != "" {
		rotationPeriod := d.Get("rotation_period").(string)
		nextRotation, err := kmsCryptoKeyNextRotation(time.Now(), rotationPeriod)

		if err != nil {
			return nil, fmt.Errorf("Error setting CryptoKey rotation period: %s", err.Error())
		}

		obj["nextRotationTime"] = nextRotation
	}

	return obj, nil
}

func expandKmsCryptoKeyPurpose(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandKmsCryptoKeyRotationPeriod(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandKmsCryptoKeyVersionTemplate(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedAlgorithm, err := expandKmsCryptoKeyVersionTemplateAlgorithm(original["algorithm"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAlgorithm); val.IsValid() && !isEmptyValue(val) {
		transformed["algorithm"] = transformedAlgorithm
	}

	transformedProtectionLevel, err := expandKmsCryptoKeyVersionTemplateProtectionLevel(original["protection_level"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedProtectionLevel); val.IsValid() && !isEmptyValue(val) {
		transformed["protectionLevel"] = transformedProtectionLevel
	}

	return transformed, nil
}

func expandKmsCryptoKeyVersionTemplateAlgorithm(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandKmsCryptoKeyVersionTemplateProtectionLevel(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
