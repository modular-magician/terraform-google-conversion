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
	"fmt"
	"reflect"
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func compareSignatureAlgorithm(_, old, new string, _ *schema.ResourceData) bool {
	// See https://cloud.google.com/binary-authorization/docs/reference/rest/v1/projects.attestors#signaturealgorithm
	normalizedAlgorithms := map[string]string{
		"ECDSA_P256_SHA256":   "ECDSA_P256_SHA256",
		"EC_SIGN_P256_SHA256": "ECDSA_P256_SHA256",
		"ECDSA_P384_SHA384":   "ECDSA_P384_SHA384",
		"EC_SIGN_P384_SHA384": "ECDSA_P384_SHA384",
		"ECDSA_P521_SHA512":   "ECDSA_P521_SHA512",
		"EC_SIGN_P521_SHA512": "ECDSA_P521_SHA512",
	}

	normalizedOld := old
	normalizedNew := new

	if normalized, ok := normalizedAlgorithms[old]; ok {
		normalizedOld = normalized
	}
	if normalized, ok := normalizedAlgorithms[new]; ok {
		normalizedNew = normalized
	}

	if normalizedNew == normalizedOld {
		return true
	}

	return false
}

const BinaryAuthorizationAttestorAssetType string = "binaryauthorization.googleapis.com/Attestor"

func resourceConverterBinaryAuthorizationAttestor() ResourceConverter {
	return ResourceConverter{
		AssetType: BinaryAuthorizationAttestorAssetType,
		Convert:   GetBinaryAuthorizationAttestorCaiObject,
	}
}

func GetBinaryAuthorizationAttestorCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	name, err := assetName(d, config, "//binaryauthorization.googleapis.com/projects/{{project}}/attestors/{{name}}")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetBinaryAuthorizationAttestorApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: BinaryAuthorizationAttestorAssetType,
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/binaryauthorization/v1/rest",
				DiscoveryName:        "Attestor",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
	}
}

func GetBinaryAuthorizationAttestorApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandBinaryAuthorizationAttestorName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandBinaryAuthorizationAttestorDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	userOwnedGrafeasNoteProp, err := expandBinaryAuthorizationAttestorAttestationAuthorityNote(d.Get("attestation_authority_note"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("attestation_authority_note"); !isEmptyValue(reflect.ValueOf(userOwnedGrafeasNoteProp)) && (ok || !reflect.DeepEqual(v, userOwnedGrafeasNoteProp)) {
		obj["userOwnedGrafeasNote"] = userOwnedGrafeasNoteProp
	}

	return obj, nil
}

func expandBinaryAuthorizationAttestorName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return replaceVars(d, config, "projects/{{project}}/attestors/{{name}}")
}

func expandBinaryAuthorizationAttestorDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBinaryAuthorizationAttestorAttestationAuthorityNote(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedNoteReference, err := expandBinaryAuthorizationAttestorAttestationAuthorityNoteNoteReference(original["note_reference"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNoteReference); val.IsValid() && !isEmptyValue(val) {
		transformed["noteReference"] = transformedNoteReference
	}

	transformedPublicKeys, err := expandBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeys(original["public_keys"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPublicKeys); val.IsValid() && !isEmptyValue(val) {
		transformed["publicKeys"] = transformedPublicKeys
	}

	transformedDelegationServiceAccountEmail, err := expandBinaryAuthorizationAttestorAttestationAuthorityNoteDelegationServiceAccountEmail(original["delegation_service_account_email"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDelegationServiceAccountEmail); val.IsValid() && !isEmptyValue(val) {
		transformed["delegationServiceAccountEmail"] = transformedDelegationServiceAccountEmail
	}

	return transformed, nil
}

func expandBinaryAuthorizationAttestorAttestationAuthorityNoteNoteReference(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	r := regexp.MustCompile("projects/(.+)/notes/(.+)")
	if r.MatchString(v.(string)) {
		return v.(string), nil
	}

	project, err := getProject(d, config)
	if err != nil {
		return nil, err
	}

	return fmt.Sprintf("projects/%s/notes/%s", project, v.(string)), nil
}

func expandBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeys(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedComment, err := expandBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysComment(original["comment"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedComment); val.IsValid() && !isEmptyValue(val) {
			transformed["comment"] = transformedComment
		}

		transformedId, err := expandBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysId(original["id"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedId); val.IsValid() && !isEmptyValue(val) {
			transformed["id"] = transformedId
		}

		transformedAsciiArmoredPgpPublicKey, err := expandBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysAsciiArmoredPgpPublicKey(original["ascii_armored_pgp_public_key"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedAsciiArmoredPgpPublicKey); val.IsValid() && !isEmptyValue(val) {
			transformed["asciiArmoredPgpPublicKey"] = transformedAsciiArmoredPgpPublicKey
		}

		transformedPkixPublicKey, err := expandBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysPkixPublicKey(original["pkix_public_key"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedPkixPublicKey); val.IsValid() && !isEmptyValue(val) {
			transformed["pkixPublicKey"] = transformedPkixPublicKey
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysComment(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysId(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysAsciiArmoredPgpPublicKey(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysPkixPublicKey(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPublicKeyPem, err := expandBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysPkixPublicKeyPublicKeyPem(original["public_key_pem"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPublicKeyPem); val.IsValid() && !isEmptyValue(val) {
		transformed["publicKeyPem"] = transformedPublicKeyPem
	}

	transformedSignatureAlgorithm, err := expandBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysPkixPublicKeySignatureAlgorithm(original["signature_algorithm"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSignatureAlgorithm); val.IsValid() && !isEmptyValue(val) {
		transformed["signatureAlgorithm"] = transformedSignatureAlgorithm
	}

	return transformed, nil
}

func expandBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysPkixPublicKeyPublicKeyPem(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysPkixPublicKeySignatureAlgorithm(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBinaryAuthorizationAttestorAttestationAuthorityNoteDelegationServiceAccountEmail(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
