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

import "reflect"

func GetPrivatecaCertificateAuthorityCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	name, err := assetName(d, config, "//privateca.googleapis.com/projects/{{project}}/locations/{{location}}/caPools/{{pool}}/certificateAuthorities/{{certificate_authority_id}}")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetPrivatecaCertificateAuthorityApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: "privateca.googleapis.com/CertificateAuthority",
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/privateca/v1/rest",
				DiscoveryName:        "CertificateAuthority",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
	}
}

func GetPrivatecaCertificateAuthorityApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	typeProp, err := expandPrivatecaCertificateAuthorityType(d.Get("type"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("type"); !isEmptyValue(reflect.ValueOf(typeProp)) && (ok || !reflect.DeepEqual(v, typeProp)) {
		obj["type"] = typeProp
	}
	tierProp, err := expandPrivatecaCertificateAuthorityTier(d.Get("tier"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("tier"); !isEmptyValue(reflect.ValueOf(tierProp)) && (ok || !reflect.DeepEqual(v, tierProp)) {
		obj["tier"] = tierProp
	}
	configProp, err := expandPrivatecaCertificateAuthorityConfig(d.Get("config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("config"); !isEmptyValue(reflect.ValueOf(configProp)) && (ok || !reflect.DeepEqual(v, configProp)) {
		obj["config"] = configProp
	}
	lifetimeProp, err := expandPrivatecaCertificateAuthorityLifetime(d.Get("lifetime"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("lifetime"); !isEmptyValue(reflect.ValueOf(lifetimeProp)) && (ok || !reflect.DeepEqual(v, lifetimeProp)) {
		obj["lifetime"] = lifetimeProp
	}
	keySpecProp, err := expandPrivatecaCertificateAuthorityKeySpec(d.Get("key_spec"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("key_spec"); !isEmptyValue(reflect.ValueOf(keySpecProp)) && (ok || !reflect.DeepEqual(v, keySpecProp)) {
		obj["keySpec"] = keySpecProp
	}
	gcsBucketProp, err := expandPrivatecaCertificateAuthorityGcsBucket(d.Get("gcs_bucket"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("gcs_bucket"); !isEmptyValue(reflect.ValueOf(gcsBucketProp)) && (ok || !reflect.DeepEqual(v, gcsBucketProp)) {
		obj["gcsBucket"] = gcsBucketProp
	}
	labelsProp, err := expandPrivatecaCertificateAuthorityLabels(d.Get("labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	return obj, nil
}

func expandPrivatecaCertificateAuthorityType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateAuthorityTier(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateAuthorityConfig(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedX509Config, err := expandPrivatecaCertificateAuthorityConfigX509Config(original["x509_config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedX509Config); val.IsValid() && !isEmptyValue(val) {
		transformed["x509Config"] = transformedX509Config
	}

	transformedSubjectConfig, err := expandPrivatecaCertificateAuthorityConfigSubjectConfig(original["subject_config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSubjectConfig); val.IsValid() && !isEmptyValue(val) {
		transformed["subjectConfig"] = transformedSubjectConfig
	}

	return transformed, nil
}

func expandPrivatecaCertificateAuthorityConfigX509Config(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedAdditionalExtensions, err := expandPrivatecaCertificateAuthorityConfigX509ConfigAdditionalExtensions(original["additional_extensions"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAdditionalExtensions); val.IsValid() && !isEmptyValue(val) {
		transformed["additionalExtensions"] = transformedAdditionalExtensions
	}

	transformedPolicyIds, err := expandPrivatecaCertificateAuthorityConfigX509ConfigPolicyIds(original["policy_ids"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPolicyIds); val.IsValid() && !isEmptyValue(val) {
		transformed["policyIds"] = transformedPolicyIds
	}

	transformedAiaOcspServers, err := expandPrivatecaCertificateAuthorityConfigX509ConfigAiaOcspServers(original["aia_ocsp_servers"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAiaOcspServers); val.IsValid() && !isEmptyValue(val) {
		transformed["aiaOcspServers"] = transformedAiaOcspServers
	}

	transformedCaOptions, err := expandPrivatecaCertificateAuthorityConfigX509ConfigCaOptions(original["ca_options"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCaOptions); val.IsValid() && !isEmptyValue(val) {
		transformed["caOptions"] = transformedCaOptions
	}

	transformedKeyUsage, err := expandPrivatecaCertificateAuthorityConfigX509ConfigKeyUsage(original["key_usage"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKeyUsage); val.IsValid() && !isEmptyValue(val) {
		transformed["keyUsage"] = transformedKeyUsage
	}

	return transformed, nil
}

func expandPrivatecaCertificateAuthorityConfigX509ConfigAdditionalExtensions(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedCritical, err := expandPrivatecaCertificateAuthorityConfigX509ConfigAdditionalExtensionsCritical(original["critical"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedCritical); val.IsValid() && !isEmptyValue(val) {
			transformed["critical"] = transformedCritical
		}

		transformedValue, err := expandPrivatecaCertificateAuthorityConfigX509ConfigAdditionalExtensionsValue(original["value"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedValue); val.IsValid() && !isEmptyValue(val) {
			transformed["value"] = transformedValue
		}

		transformedObjectId, err := expandPrivatecaCertificateAuthorityConfigX509ConfigAdditionalExtensionsObjectId(original["object_id"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedObjectId); val.IsValid() && !isEmptyValue(val) {
			transformed["objectId"] = transformedObjectId
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandPrivatecaCertificateAuthorityConfigX509ConfigAdditionalExtensionsCritical(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateAuthorityConfigX509ConfigAdditionalExtensionsValue(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateAuthorityConfigX509ConfigAdditionalExtensionsObjectId(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedObjectIdPath, err := expandPrivatecaCertificateAuthorityConfigX509ConfigAdditionalExtensionsObjectIdObjectIdPath(original["object_id_path"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedObjectIdPath); val.IsValid() && !isEmptyValue(val) {
		transformed["objectIdPath"] = transformedObjectIdPath
	}

	return transformed, nil
}

func expandPrivatecaCertificateAuthorityConfigX509ConfigAdditionalExtensionsObjectIdObjectIdPath(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateAuthorityConfigX509ConfigPolicyIds(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedObjectIdPath, err := expandPrivatecaCertificateAuthorityConfigX509ConfigPolicyIdsObjectIdPath(original["object_id_path"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedObjectIdPath); val.IsValid() && !isEmptyValue(val) {
			transformed["objectIdPath"] = transformedObjectIdPath
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandPrivatecaCertificateAuthorityConfigX509ConfigPolicyIdsObjectIdPath(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateAuthorityConfigX509ConfigAiaOcspServers(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateAuthorityConfigX509ConfigCaOptions(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedIsCa, err := expandPrivatecaCertificateAuthorityConfigX509ConfigCaOptionsIsCa(original["is_ca"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedIsCa); val.IsValid() && !isEmptyValue(val) {
		transformed["isCa"] = transformedIsCa
	}

	transformedMaxIssuerPathLength, err := expandPrivatecaCertificateAuthorityConfigX509ConfigCaOptionsMaxIssuerPathLength(original["max_issuer_path_length"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxIssuerPathLength); val.IsValid() && !isEmptyValue(val) {
		transformed["maxIssuerPathLength"] = transformedMaxIssuerPathLength
	}

	return transformed, nil
}

func expandPrivatecaCertificateAuthorityConfigX509ConfigCaOptionsIsCa(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateAuthorityConfigX509ConfigCaOptionsMaxIssuerPathLength(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateAuthorityConfigX509ConfigKeyUsage(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedBaseKeyUsage, err := expandPrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsage(original["base_key_usage"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedBaseKeyUsage); val.IsValid() && !isEmptyValue(val) {
		transformed["baseKeyUsage"] = transformedBaseKeyUsage
	}

	transformedExtendedKeyUsage, err := expandPrivatecaCertificateAuthorityConfigX509ConfigKeyUsageExtendedKeyUsage(original["extended_key_usage"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedExtendedKeyUsage); val.IsValid() && !isEmptyValue(val) {
		transformed["extendedKeyUsage"] = transformedExtendedKeyUsage
	}

	transformedUnknownExtendedKeyUsages, err := expandPrivatecaCertificateAuthorityConfigX509ConfigKeyUsageUnknownExtendedKeyUsages(original["unknown_extended_key_usages"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUnknownExtendedKeyUsages); val.IsValid() && !isEmptyValue(val) {
		transformed["unknownExtendedKeyUsages"] = transformedUnknownExtendedKeyUsages
	}

	return transformed, nil
}

func expandPrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsage(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedDigitalSignature, err := expandPrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageDigitalSignature(original["digital_signature"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDigitalSignature); val.IsValid() && !isEmptyValue(val) {
		transformed["digitalSignature"] = transformedDigitalSignature
	}

	transformedContentCommitment, err := expandPrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageContentCommitment(original["content_commitment"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedContentCommitment); val.IsValid() && !isEmptyValue(val) {
		transformed["contentCommitment"] = transformedContentCommitment
	}

	transformedKeyEncipherment, err := expandPrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageKeyEncipherment(original["key_encipherment"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKeyEncipherment); val.IsValid() && !isEmptyValue(val) {
		transformed["keyEncipherment"] = transformedKeyEncipherment
	}

	transformedDataEncipherment, err := expandPrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageDataEncipherment(original["data_encipherment"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDataEncipherment); val.IsValid() && !isEmptyValue(val) {
		transformed["dataEncipherment"] = transformedDataEncipherment
	}

	transformedKeyAgreement, err := expandPrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageKeyAgreement(original["key_agreement"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKeyAgreement); val.IsValid() && !isEmptyValue(val) {
		transformed["keyAgreement"] = transformedKeyAgreement
	}

	transformedCertSign, err := expandPrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageCertSign(original["cert_sign"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCertSign); val.IsValid() && !isEmptyValue(val) {
		transformed["certSign"] = transformedCertSign
	}

	transformedCrlSign, err := expandPrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageCrlSign(original["crl_sign"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCrlSign); val.IsValid() && !isEmptyValue(val) {
		transformed["crlSign"] = transformedCrlSign
	}

	transformedEncipherOnly, err := expandPrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageEncipherOnly(original["encipher_only"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEncipherOnly); val.IsValid() && !isEmptyValue(val) {
		transformed["encipherOnly"] = transformedEncipherOnly
	}

	transformedDecipherOnly, err := expandPrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageDecipherOnly(original["decipher_only"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDecipherOnly); val.IsValid() && !isEmptyValue(val) {
		transformed["decipherOnly"] = transformedDecipherOnly
	}

	return transformed, nil
}

func expandPrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageDigitalSignature(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageContentCommitment(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageKeyEncipherment(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageDataEncipherment(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageKeyAgreement(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageCertSign(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageCrlSign(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageEncipherOnly(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageDecipherOnly(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateAuthorityConfigX509ConfigKeyUsageExtendedKeyUsage(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedServerAuth, err := expandPrivatecaCertificateAuthorityConfigX509ConfigKeyUsageExtendedKeyUsageServerAuth(original["server_auth"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedServerAuth); val.IsValid() && !isEmptyValue(val) {
		transformed["serverAuth"] = transformedServerAuth
	}

	transformedClientAuth, err := expandPrivatecaCertificateAuthorityConfigX509ConfigKeyUsageExtendedKeyUsageClientAuth(original["client_auth"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedClientAuth); val.IsValid() && !isEmptyValue(val) {
		transformed["clientAuth"] = transformedClientAuth
	}

	transformedCodeSigning, err := expandPrivatecaCertificateAuthorityConfigX509ConfigKeyUsageExtendedKeyUsageCodeSigning(original["code_signing"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCodeSigning); val.IsValid() && !isEmptyValue(val) {
		transformed["codeSigning"] = transformedCodeSigning
	}

	transformedEmailProtection, err := expandPrivatecaCertificateAuthorityConfigX509ConfigKeyUsageExtendedKeyUsageEmailProtection(original["email_protection"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEmailProtection); val.IsValid() && !isEmptyValue(val) {
		transformed["emailProtection"] = transformedEmailProtection
	}

	transformedTimeStamping, err := expandPrivatecaCertificateAuthorityConfigX509ConfigKeyUsageExtendedKeyUsageTimeStamping(original["time_stamping"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTimeStamping); val.IsValid() && !isEmptyValue(val) {
		transformed["timeStamping"] = transformedTimeStamping
	}

	transformedOcspSigning, err := expandPrivatecaCertificateAuthorityConfigX509ConfigKeyUsageExtendedKeyUsageOcspSigning(original["ocsp_signing"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedOcspSigning); val.IsValid() && !isEmptyValue(val) {
		transformed["ocspSigning"] = transformedOcspSigning
	}

	return transformed, nil
}

func expandPrivatecaCertificateAuthorityConfigX509ConfigKeyUsageExtendedKeyUsageServerAuth(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateAuthorityConfigX509ConfigKeyUsageExtendedKeyUsageClientAuth(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateAuthorityConfigX509ConfigKeyUsageExtendedKeyUsageCodeSigning(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateAuthorityConfigX509ConfigKeyUsageExtendedKeyUsageEmailProtection(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateAuthorityConfigX509ConfigKeyUsageExtendedKeyUsageTimeStamping(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateAuthorityConfigX509ConfigKeyUsageExtendedKeyUsageOcspSigning(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateAuthorityConfigX509ConfigKeyUsageUnknownExtendedKeyUsages(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedObjectIdPath, err := expandPrivatecaCertificateAuthorityConfigX509ConfigKeyUsageUnknownExtendedKeyUsagesObjectIdPath(original["object_id_path"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedObjectIdPath); val.IsValid() && !isEmptyValue(val) {
			transformed["objectIdPath"] = transformedObjectIdPath
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandPrivatecaCertificateAuthorityConfigX509ConfigKeyUsageUnknownExtendedKeyUsagesObjectIdPath(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateAuthorityConfigSubjectConfig(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedSubject, err := expandPrivatecaCertificateAuthorityConfigSubjectConfigSubject(original["subject"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSubject); val.IsValid() && !isEmptyValue(val) {
		transformed["subject"] = transformedSubject
	}

	transformedSubjectAltName, err := expandPrivatecaCertificateAuthorityConfigSubjectConfigSubjectAltName(original["subject_alt_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSubjectAltName); val.IsValid() && !isEmptyValue(val) {
		transformed["subjectAltName"] = transformedSubjectAltName
	}

	return transformed, nil
}

func expandPrivatecaCertificateAuthorityConfigSubjectConfigSubject(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedCountryCode, err := expandPrivatecaCertificateAuthorityConfigSubjectConfigSubjectCountryCode(original["country_code"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCountryCode); val.IsValid() && !isEmptyValue(val) {
		transformed["countryCode"] = transformedCountryCode
	}

	transformedOrganization, err := expandPrivatecaCertificateAuthorityConfigSubjectConfigSubjectOrganization(original["organization"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedOrganization); val.IsValid() && !isEmptyValue(val) {
		transformed["organization"] = transformedOrganization
	}

	transformedOrganizationalUnit, err := expandPrivatecaCertificateAuthorityConfigSubjectConfigSubjectOrganizationalUnit(original["organizational_unit"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedOrganizationalUnit); val.IsValid() && !isEmptyValue(val) {
		transformed["organizationalUnit"] = transformedOrganizationalUnit
	}

	transformedLocality, err := expandPrivatecaCertificateAuthorityConfigSubjectConfigSubjectLocality(original["locality"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLocality); val.IsValid() && !isEmptyValue(val) {
		transformed["locality"] = transformedLocality
	}

	transformedProvince, err := expandPrivatecaCertificateAuthorityConfigSubjectConfigSubjectProvince(original["province"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedProvince); val.IsValid() && !isEmptyValue(val) {
		transformed["province"] = transformedProvince
	}

	transformedStreetAddress, err := expandPrivatecaCertificateAuthorityConfigSubjectConfigSubjectStreetAddress(original["street_address"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedStreetAddress); val.IsValid() && !isEmptyValue(val) {
		transformed["streetAddress"] = transformedStreetAddress
	}

	transformedPostalCode, err := expandPrivatecaCertificateAuthorityConfigSubjectConfigSubjectPostalCode(original["postal_code"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPostalCode); val.IsValid() && !isEmptyValue(val) {
		transformed["postalCode"] = transformedPostalCode
	}

	transformedCommonName, err := expandPrivatecaCertificateAuthorityConfigSubjectConfigSubjectCommonName(original["common_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCommonName); val.IsValid() && !isEmptyValue(val) {
		transformed["commonName"] = transformedCommonName
	}

	return transformed, nil
}

func expandPrivatecaCertificateAuthorityConfigSubjectConfigSubjectCountryCode(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateAuthorityConfigSubjectConfigSubjectOrganization(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateAuthorityConfigSubjectConfigSubjectOrganizationalUnit(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateAuthorityConfigSubjectConfigSubjectLocality(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateAuthorityConfigSubjectConfigSubjectProvince(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateAuthorityConfigSubjectConfigSubjectStreetAddress(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateAuthorityConfigSubjectConfigSubjectPostalCode(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateAuthorityConfigSubjectConfigSubjectCommonName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateAuthorityConfigSubjectConfigSubjectAltName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedDnsNames, err := expandPrivatecaCertificateAuthorityConfigSubjectConfigSubjectAltNameDnsNames(original["dns_names"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDnsNames); val.IsValid() && !isEmptyValue(val) {
		transformed["dnsNames"] = transformedDnsNames
	}

	transformedUris, err := expandPrivatecaCertificateAuthorityConfigSubjectConfigSubjectAltNameUris(original["uris"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUris); val.IsValid() && !isEmptyValue(val) {
		transformed["uris"] = transformedUris
	}

	transformedEmailAddresses, err := expandPrivatecaCertificateAuthorityConfigSubjectConfigSubjectAltNameEmailAddresses(original["email_addresses"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEmailAddresses); val.IsValid() && !isEmptyValue(val) {
		transformed["emailAddresses"] = transformedEmailAddresses
	}

	transformedIpAddresses, err := expandPrivatecaCertificateAuthorityConfigSubjectConfigSubjectAltNameIpAddresses(original["ip_addresses"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedIpAddresses); val.IsValid() && !isEmptyValue(val) {
		transformed["ipAddresses"] = transformedIpAddresses
	}

	return transformed, nil
}

func expandPrivatecaCertificateAuthorityConfigSubjectConfigSubjectAltNameDnsNames(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateAuthorityConfigSubjectConfigSubjectAltNameUris(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateAuthorityConfigSubjectConfigSubjectAltNameEmailAddresses(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateAuthorityConfigSubjectConfigSubjectAltNameIpAddresses(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateAuthorityLifetime(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateAuthorityKeySpec(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedCloudKmsKeyVersion, err := expandPrivatecaCertificateAuthorityKeySpecCloudKmsKeyVersion(original["cloud_kms_key_version"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCloudKmsKeyVersion); val.IsValid() && !isEmptyValue(val) {
		transformed["cloudKmsKeyVersion"] = transformedCloudKmsKeyVersion
	}

	transformedAlgorithm, err := expandPrivatecaCertificateAuthorityKeySpecAlgorithm(original["algorithm"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAlgorithm); val.IsValid() && !isEmptyValue(val) {
		transformed["algorithm"] = transformedAlgorithm
	}

	return transformed, nil
}

func expandPrivatecaCertificateAuthorityKeySpecCloudKmsKeyVersion(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateAuthorityKeySpecAlgorithm(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateAuthorityGcsBucket(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateAuthorityLabels(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
