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

import "reflect"

const PrivatecaCertificateAssetType string = "privateca.googleapis.com/Certificate"

func resourceConverterPrivatecaCertificate() ResourceConverter {
	return ResourceConverter{
		AssetType: PrivatecaCertificateAssetType,
		Convert:   GetPrivatecaCertificateCaiObject,
	}
}

func GetPrivatecaCertificateCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	name, err := assetName(d, config, "//privateca.googleapis.com/projects/{{project}}/locations/{{location}}/caPools/{{pool}}/certificates/{{name}}")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetPrivatecaCertificateApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: PrivatecaCertificateAssetType,
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/privateca/v1/rest",
				DiscoveryName:        "Certificate",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
	}
}

func GetPrivatecaCertificateApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	lifetimeProp, err := expandPrivatecaCertificateLifetime(d.Get("lifetime"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("lifetime"); !isEmptyValue(reflect.ValueOf(lifetimeProp)) && (ok || !reflect.DeepEqual(v, lifetimeProp)) {
		obj["lifetime"] = lifetimeProp
	}
	certificateTemplateProp, err := expandPrivatecaCertificateCertificateTemplate(d.Get("certificate_template"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("certificate_template"); !isEmptyValue(reflect.ValueOf(certificateTemplateProp)) && (ok || !reflect.DeepEqual(v, certificateTemplateProp)) {
		obj["certificateTemplate"] = certificateTemplateProp
	}
	labelsProp, err := expandPrivatecaCertificateLabels(d.Get("labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	pemCsrProp, err := expandPrivatecaCertificatePemCsr(d.Get("pem_csr"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("pem_csr"); !isEmptyValue(reflect.ValueOf(pemCsrProp)) && (ok || !reflect.DeepEqual(v, pemCsrProp)) {
		obj["pemCsr"] = pemCsrProp
	}
	configProp, err := expandPrivatecaCertificateConfig(d.Get("config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("config"); !isEmptyValue(reflect.ValueOf(configProp)) && (ok || !reflect.DeepEqual(v, configProp)) {
		obj["config"] = configProp
	}

	return obj, nil
}

func expandPrivatecaCertificateLifetime(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateCertificateTemplate(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateLabels(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandPrivatecaCertificatePemCsr(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateConfig(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedX509Config, err := expandPrivatecaCertificateConfigX509Config(original["x509_config"], d, config)
	if err != nil {
		return nil, err
	} else {
		transformed["x509Config"] = transformedX509Config
	}

	transformedSubjectConfig, err := expandPrivatecaCertificateConfigSubjectConfig(original["subject_config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSubjectConfig); val.IsValid() && !isEmptyValue(val) {
		transformed["subjectConfig"] = transformedSubjectConfig
	}

	transformedPublicKey, err := expandPrivatecaCertificateConfigPublicKey(original["public_key"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPublicKey); val.IsValid() && !isEmptyValue(val) {
		transformed["publicKey"] = transformedPublicKey
	}

	return transformed, nil
}

func expandPrivatecaCertificateConfigX509Config(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedAdditionalExtensions, err := expandPrivatecaCertificateConfigX509ConfigAdditionalExtensions(original["additional_extensions"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAdditionalExtensions); val.IsValid() && !isEmptyValue(val) {
		transformed["additionalExtensions"] = transformedAdditionalExtensions
	}

	transformedPolicyIds, err := expandPrivatecaCertificateConfigX509ConfigPolicyIds(original["policy_ids"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPolicyIds); val.IsValid() && !isEmptyValue(val) {
		transformed["policyIds"] = transformedPolicyIds
	}

	transformedAiaOcspServers, err := expandPrivatecaCertificateConfigX509ConfigAiaOcspServers(original["aia_ocsp_servers"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAiaOcspServers); val.IsValid() && !isEmptyValue(val) {
		transformed["aiaOcspServers"] = transformedAiaOcspServers
	}

	transformedCaOptions, err := expandPrivatecaCertificateConfigX509ConfigCaOptions(original["ca_options"], d, config)
	if err != nil {
		return nil, err
	} else {
		transformed["caOptions"] = transformedCaOptions
	}

	transformedKeyUsage, err := expandPrivatecaCertificateConfigX509ConfigKeyUsage(original["key_usage"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKeyUsage); val.IsValid() && !isEmptyValue(val) {
		transformed["keyUsage"] = transformedKeyUsage
	}

	return transformed, nil
}

func expandPrivatecaCertificateConfigX509ConfigAdditionalExtensions(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedCritical, err := expandPrivatecaCertificateConfigX509ConfigAdditionalExtensionsCritical(original["critical"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedCritical); val.IsValid() && !isEmptyValue(val) {
			transformed["critical"] = transformedCritical
		}

		transformedValue, err := expandPrivatecaCertificateConfigX509ConfigAdditionalExtensionsValue(original["value"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedValue); val.IsValid() && !isEmptyValue(val) {
			transformed["value"] = transformedValue
		}

		transformedObjectId, err := expandPrivatecaCertificateConfigX509ConfigAdditionalExtensionsObjectId(original["object_id"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedObjectId); val.IsValid() && !isEmptyValue(val) {
			transformed["objectId"] = transformedObjectId
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandPrivatecaCertificateConfigX509ConfigAdditionalExtensionsCritical(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateConfigX509ConfigAdditionalExtensionsValue(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateConfigX509ConfigAdditionalExtensionsObjectId(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedObjectIdPath, err := expandPrivatecaCertificateConfigX509ConfigAdditionalExtensionsObjectIdObjectIdPath(original["object_id_path"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedObjectIdPath); val.IsValid() && !isEmptyValue(val) {
		transformed["objectIdPath"] = transformedObjectIdPath
	}

	return transformed, nil
}

func expandPrivatecaCertificateConfigX509ConfigAdditionalExtensionsObjectIdObjectIdPath(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateConfigX509ConfigPolicyIds(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedObjectIdPath, err := expandPrivatecaCertificateConfigX509ConfigPolicyIdsObjectIdPath(original["object_id_path"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedObjectIdPath); val.IsValid() && !isEmptyValue(val) {
			transformed["objectIdPath"] = transformedObjectIdPath
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandPrivatecaCertificateConfigX509ConfigPolicyIdsObjectIdPath(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateConfigX509ConfigAiaOcspServers(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateConfigX509ConfigCaOptions(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedIsCa, err := expandPrivatecaCertificateConfigX509ConfigCaOptionsIsCa(original["is_ca"], d, config)
	if err != nil {
		return nil, err
	} else {
		transformed["isCa"] = transformedIsCa
	}

	transformedMaxIssuerPathLength, err := expandPrivatecaCertificateConfigX509ConfigCaOptionsMaxIssuerPathLength(original["max_issuer_path_length"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxIssuerPathLength); val.IsValid() && !isEmptyValue(val) {
		transformed["maxIssuerPathLength"] = transformedMaxIssuerPathLength
	}

	return transformed, nil
}

func expandPrivatecaCertificateConfigX509ConfigCaOptionsIsCa(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateConfigX509ConfigCaOptionsMaxIssuerPathLength(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateConfigX509ConfigKeyUsage(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedBaseKeyUsage, err := expandPrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsage(original["base_key_usage"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedBaseKeyUsage); val.IsValid() && !isEmptyValue(val) {
		transformed["baseKeyUsage"] = transformedBaseKeyUsage
	}

	transformedExtendedKeyUsage, err := expandPrivatecaCertificateConfigX509ConfigKeyUsageExtendedKeyUsage(original["extended_key_usage"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedExtendedKeyUsage); val.IsValid() && !isEmptyValue(val) {
		transformed["extendedKeyUsage"] = transformedExtendedKeyUsage
	}

	transformedUnknownExtendedKeyUsages, err := expandPrivatecaCertificateConfigX509ConfigKeyUsageUnknownExtendedKeyUsages(original["unknown_extended_key_usages"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUnknownExtendedKeyUsages); val.IsValid() && !isEmptyValue(val) {
		transformed["unknownExtendedKeyUsages"] = transformedUnknownExtendedKeyUsages
	}

	return transformed, nil
}

func expandPrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsage(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedDigitalSignature, err := expandPrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsageDigitalSignature(original["digital_signature"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDigitalSignature); val.IsValid() && !isEmptyValue(val) {
		transformed["digitalSignature"] = transformedDigitalSignature
	}

	transformedContentCommitment, err := expandPrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsageContentCommitment(original["content_commitment"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedContentCommitment); val.IsValid() && !isEmptyValue(val) {
		transformed["contentCommitment"] = transformedContentCommitment
	}

	transformedKeyEncipherment, err := expandPrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsageKeyEncipherment(original["key_encipherment"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKeyEncipherment); val.IsValid() && !isEmptyValue(val) {
		transformed["keyEncipherment"] = transformedKeyEncipherment
	}

	transformedDataEncipherment, err := expandPrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsageDataEncipherment(original["data_encipherment"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDataEncipherment); val.IsValid() && !isEmptyValue(val) {
		transformed["dataEncipherment"] = transformedDataEncipherment
	}

	transformedKeyAgreement, err := expandPrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsageKeyAgreement(original["key_agreement"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKeyAgreement); val.IsValid() && !isEmptyValue(val) {
		transformed["keyAgreement"] = transformedKeyAgreement
	}

	transformedCertSign, err := expandPrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsageCertSign(original["cert_sign"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCertSign); val.IsValid() && !isEmptyValue(val) {
		transformed["certSign"] = transformedCertSign
	}

	transformedCrlSign, err := expandPrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsageCrlSign(original["crl_sign"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCrlSign); val.IsValid() && !isEmptyValue(val) {
		transformed["crlSign"] = transformedCrlSign
	}

	transformedEncipherOnly, err := expandPrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsageEncipherOnly(original["encipher_only"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEncipherOnly); val.IsValid() && !isEmptyValue(val) {
		transformed["encipherOnly"] = transformedEncipherOnly
	}

	transformedDecipherOnly, err := expandPrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsageDecipherOnly(original["decipher_only"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDecipherOnly); val.IsValid() && !isEmptyValue(val) {
		transformed["decipherOnly"] = transformedDecipherOnly
	}

	return transformed, nil
}

func expandPrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsageDigitalSignature(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsageContentCommitment(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsageKeyEncipherment(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsageDataEncipherment(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsageKeyAgreement(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsageCertSign(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsageCrlSign(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsageEncipherOnly(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsageDecipherOnly(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateConfigX509ConfigKeyUsageExtendedKeyUsage(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedServerAuth, err := expandPrivatecaCertificateConfigX509ConfigKeyUsageExtendedKeyUsageServerAuth(original["server_auth"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedServerAuth); val.IsValid() && !isEmptyValue(val) {
		transformed["serverAuth"] = transformedServerAuth
	}

	transformedClientAuth, err := expandPrivatecaCertificateConfigX509ConfigKeyUsageExtendedKeyUsageClientAuth(original["client_auth"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedClientAuth); val.IsValid() && !isEmptyValue(val) {
		transformed["clientAuth"] = transformedClientAuth
	}

	transformedCodeSigning, err := expandPrivatecaCertificateConfigX509ConfigKeyUsageExtendedKeyUsageCodeSigning(original["code_signing"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCodeSigning); val.IsValid() && !isEmptyValue(val) {
		transformed["codeSigning"] = transformedCodeSigning
	}

	transformedEmailProtection, err := expandPrivatecaCertificateConfigX509ConfigKeyUsageExtendedKeyUsageEmailProtection(original["email_protection"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEmailProtection); val.IsValid() && !isEmptyValue(val) {
		transformed["emailProtection"] = transformedEmailProtection
	}

	transformedTimeStamping, err := expandPrivatecaCertificateConfigX509ConfigKeyUsageExtendedKeyUsageTimeStamping(original["time_stamping"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTimeStamping); val.IsValid() && !isEmptyValue(val) {
		transformed["timeStamping"] = transformedTimeStamping
	}

	transformedOcspSigning, err := expandPrivatecaCertificateConfigX509ConfigKeyUsageExtendedKeyUsageOcspSigning(original["ocsp_signing"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedOcspSigning); val.IsValid() && !isEmptyValue(val) {
		transformed["ocspSigning"] = transformedOcspSigning
	}

	return transformed, nil
}

func expandPrivatecaCertificateConfigX509ConfigKeyUsageExtendedKeyUsageServerAuth(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateConfigX509ConfigKeyUsageExtendedKeyUsageClientAuth(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateConfigX509ConfigKeyUsageExtendedKeyUsageCodeSigning(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateConfigX509ConfigKeyUsageExtendedKeyUsageEmailProtection(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateConfigX509ConfigKeyUsageExtendedKeyUsageTimeStamping(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateConfigX509ConfigKeyUsageExtendedKeyUsageOcspSigning(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateConfigX509ConfigKeyUsageUnknownExtendedKeyUsages(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedObjectIdPath, err := expandPrivatecaCertificateConfigX509ConfigKeyUsageUnknownExtendedKeyUsagesObjectIdPath(original["object_id_path"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedObjectIdPath); val.IsValid() && !isEmptyValue(val) {
			transformed["objectIdPath"] = transformedObjectIdPath
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandPrivatecaCertificateConfigX509ConfigKeyUsageUnknownExtendedKeyUsagesObjectIdPath(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateConfigSubjectConfig(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedSubject, err := expandPrivatecaCertificateConfigSubjectConfigSubject(original["subject"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSubject); val.IsValid() && !isEmptyValue(val) {
		transformed["subject"] = transformedSubject
	}

	transformedSubjectAltName, err := expandPrivatecaCertificateConfigSubjectConfigSubjectAltName(original["subject_alt_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSubjectAltName); val.IsValid() && !isEmptyValue(val) {
		transformed["subjectAltName"] = transformedSubjectAltName
	}

	return transformed, nil
}

func expandPrivatecaCertificateConfigSubjectConfigSubject(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedCountryCode, err := expandPrivatecaCertificateConfigSubjectConfigSubjectCountryCode(original["country_code"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCountryCode); val.IsValid() && !isEmptyValue(val) {
		transformed["countryCode"] = transformedCountryCode
	}

	transformedOrganization, err := expandPrivatecaCertificateConfigSubjectConfigSubjectOrganization(original["organization"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedOrganization); val.IsValid() && !isEmptyValue(val) {
		transformed["organization"] = transformedOrganization
	}

	transformedOrganizationalUnit, err := expandPrivatecaCertificateConfigSubjectConfigSubjectOrganizationalUnit(original["organizational_unit"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedOrganizationalUnit); val.IsValid() && !isEmptyValue(val) {
		transformed["organizationalUnit"] = transformedOrganizationalUnit
	}

	transformedLocality, err := expandPrivatecaCertificateConfigSubjectConfigSubjectLocality(original["locality"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLocality); val.IsValid() && !isEmptyValue(val) {
		transformed["locality"] = transformedLocality
	}

	transformedProvince, err := expandPrivatecaCertificateConfigSubjectConfigSubjectProvince(original["province"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedProvince); val.IsValid() && !isEmptyValue(val) {
		transformed["province"] = transformedProvince
	}

	transformedStreetAddress, err := expandPrivatecaCertificateConfigSubjectConfigSubjectStreetAddress(original["street_address"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedStreetAddress); val.IsValid() && !isEmptyValue(val) {
		transformed["streetAddress"] = transformedStreetAddress
	}

	transformedPostalCode, err := expandPrivatecaCertificateConfigSubjectConfigSubjectPostalCode(original["postal_code"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPostalCode); val.IsValid() && !isEmptyValue(val) {
		transformed["postalCode"] = transformedPostalCode
	}

	transformedCommonName, err := expandPrivatecaCertificateConfigSubjectConfigSubjectCommonName(original["common_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCommonName); val.IsValid() && !isEmptyValue(val) {
		transformed["commonName"] = transformedCommonName
	}

	return transformed, nil
}

func expandPrivatecaCertificateConfigSubjectConfigSubjectCountryCode(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateConfigSubjectConfigSubjectOrganization(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateConfigSubjectConfigSubjectOrganizationalUnit(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateConfigSubjectConfigSubjectLocality(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateConfigSubjectConfigSubjectProvince(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateConfigSubjectConfigSubjectStreetAddress(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateConfigSubjectConfigSubjectPostalCode(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateConfigSubjectConfigSubjectCommonName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateConfigSubjectConfigSubjectAltName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedDnsNames, err := expandPrivatecaCertificateConfigSubjectConfigSubjectAltNameDnsNames(original["dns_names"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDnsNames); val.IsValid() && !isEmptyValue(val) {
		transformed["dnsNames"] = transformedDnsNames
	}

	transformedUris, err := expandPrivatecaCertificateConfigSubjectConfigSubjectAltNameUris(original["uris"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUris); val.IsValid() && !isEmptyValue(val) {
		transformed["uris"] = transformedUris
	}

	transformedEmailAddresses, err := expandPrivatecaCertificateConfigSubjectConfigSubjectAltNameEmailAddresses(original["email_addresses"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEmailAddresses); val.IsValid() && !isEmptyValue(val) {
		transformed["emailAddresses"] = transformedEmailAddresses
	}

	transformedIpAddresses, err := expandPrivatecaCertificateConfigSubjectConfigSubjectAltNameIpAddresses(original["ip_addresses"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedIpAddresses); val.IsValid() && !isEmptyValue(val) {
		transformed["ipAddresses"] = transformedIpAddresses
	}

	return transformed, nil
}

func expandPrivatecaCertificateConfigSubjectConfigSubjectAltNameDnsNames(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateConfigSubjectConfigSubjectAltNameUris(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateConfigSubjectConfigSubjectAltNameEmailAddresses(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateConfigSubjectConfigSubjectAltNameIpAddresses(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateConfigPublicKey(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedKey, err := expandPrivatecaCertificateConfigPublicKeyKey(original["key"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKey); val.IsValid() && !isEmptyValue(val) {
		transformed["key"] = transformedKey
	}

	transformedFormat, err := expandPrivatecaCertificateConfigPublicKeyFormat(original["format"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedFormat); val.IsValid() && !isEmptyValue(val) {
		transformed["format"] = transformedFormat
	}

	return transformed, nil
}

func expandPrivatecaCertificateConfigPublicKeyKey(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateConfigPublicKeyFormat(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
