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

const ApigeeOrganizationAssetType string = "apigee.googleapis.com/Organization"

func resourceConverterApigeeOrganization() ResourceConverter {
	return ResourceConverter{
		AssetType: ApigeeOrganizationAssetType,
		Convert:   GetApigeeOrganizationCaiObject,
	}
}

func GetApigeeOrganizationCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]Asset, error) {
	name, err := assetName(d, config, "//apigee.googleapis.com/organizations/{{name}}")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetApigeeOrganizationApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: ApigeeOrganizationAssetType,
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/apigee/v1/rest",
				DiscoveryName:        "Organization",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
	}
}

func GetApigeeOrganizationApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	displayNameProp, err := expandApigeeOrganizationDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	descriptionProp, err := expandApigeeOrganizationDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	analyticsRegionProp, err := expandApigeeOrganizationAnalyticsRegion(d.Get("analytics_region"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("analytics_region"); !tpgresource.IsEmptyValue(reflect.ValueOf(analyticsRegionProp)) && (ok || !reflect.DeepEqual(v, analyticsRegionProp)) {
		obj["analyticsRegion"] = analyticsRegionProp
	}
	authorizedNetworkProp, err := expandApigeeOrganizationAuthorizedNetwork(d.Get("authorized_network"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("authorized_network"); !tpgresource.IsEmptyValue(reflect.ValueOf(authorizedNetworkProp)) && (ok || !reflect.DeepEqual(v, authorizedNetworkProp)) {
		obj["authorizedNetwork"] = authorizedNetworkProp
	}
	runtimeTypeProp, err := expandApigeeOrganizationRuntimeType(d.Get("runtime_type"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("runtime_type"); !tpgresource.IsEmptyValue(reflect.ValueOf(runtimeTypeProp)) && (ok || !reflect.DeepEqual(v, runtimeTypeProp)) {
		obj["runtimeType"] = runtimeTypeProp
	}
	billingTypeProp, err := expandApigeeOrganizationBillingType(d.Get("billing_type"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("billing_type"); !tpgresource.IsEmptyValue(reflect.ValueOf(billingTypeProp)) && (ok || !reflect.DeepEqual(v, billingTypeProp)) {
		obj["billingType"] = billingTypeProp
	}
	runtimeDatabaseEncryptionKeyNameProp, err := expandApigeeOrganizationRuntimeDatabaseEncryptionKeyName(d.Get("runtime_database_encryption_key_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("runtime_database_encryption_key_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(runtimeDatabaseEncryptionKeyNameProp)) && (ok || !reflect.DeepEqual(v, runtimeDatabaseEncryptionKeyNameProp)) {
		obj["runtimeDatabaseEncryptionKeyName"] = runtimeDatabaseEncryptionKeyNameProp
	}
	propertiesProp, err := expandApigeeOrganizationProperties(d.Get("properties"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("properties"); !tpgresource.IsEmptyValue(reflect.ValueOf(propertiesProp)) && (ok || !reflect.DeepEqual(v, propertiesProp)) {
		obj["properties"] = propertiesProp
	}

	return resourceApigeeOrganizationEncoder(d, config, obj)
}

func resourceApigeeOrganizationEncoder(d tpgresource.TerraformResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	obj["name"] = d.Get("project_id").(string)
	return obj, nil
}

func expandApigeeOrganizationDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeOrganizationDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeOrganizationAnalyticsRegion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeOrganizationAuthorizedNetwork(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeOrganizationRuntimeType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeOrganizationBillingType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeOrganizationRuntimeDatabaseEncryptionKeyName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeOrganizationProperties(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedProperty, err := expandApigeeOrganizationPropertiesProperty(original["property"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedProperty); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["property"] = transformedProperty
	}

	return transformed, nil
}

func expandApigeeOrganizationPropertiesProperty(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedName, err := expandApigeeOrganizationPropertiesPropertyName(original["name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["name"] = transformedName
		}

		transformedValue, err := expandApigeeOrganizationPropertiesPropertyValue(original["value"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedValue); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["value"] = transformedValue
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandApigeeOrganizationPropertiesPropertyName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeOrganizationPropertiesPropertyValue(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
