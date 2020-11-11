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

func GetCloudIdentityGroupMembershipCaiObject(d TerraformResourceData, config *Config) (Asset, error) {
	name, err := assetName(d, config, "//cloudidentity.googleapis.com/{{name}}")
	if err != nil {
		return Asset{}, err
	}
	if obj, err := GetCloudIdentityGroupMembershipApiObject(d, config); err == nil {
		return Asset{
			Name: name,
			Type: "cloudidentity.googleapis.com/GroupMembership",
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/cloudidentity/v1/rest",
				DiscoveryName:        "GroupMembership",
				Data:                 obj,
			},
		}, nil
	} else {
		return Asset{}, err
	}
}

func GetCloudIdentityGroupMembershipApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	preferredMemberKeyProp, err := expandCloudIdentityGroupMembershipPreferredMemberKey(d.Get("preferred_member_key"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("preferred_member_key"); !isEmptyValue(reflect.ValueOf(preferredMemberKeyProp)) && (ok || !reflect.DeepEqual(v, preferredMemberKeyProp)) {
		obj["preferredMemberKey"] = preferredMemberKeyProp
	}
	rolesProp, err := expandCloudIdentityGroupMembershipRoles(d.Get("roles"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("roles"); !isEmptyValue(reflect.ValueOf(rolesProp)) && (ok || !reflect.DeepEqual(v, rolesProp)) {
		obj["roles"] = rolesProp
	}

	return obj, nil
}

func expandCloudIdentityGroupMembershipPreferredMemberKey(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedId, err := expandCloudIdentityGroupMembershipPreferredMemberKeyId(original["id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedId); val.IsValid() && !isEmptyValue(val) {
		transformed["id"] = transformedId
	}

	transformedNamespace, err := expandCloudIdentityGroupMembershipPreferredMemberKeyNamespace(original["namespace"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNamespace); val.IsValid() && !isEmptyValue(val) {
		transformed["namespace"] = transformedNamespace
	}

	return transformed, nil
}

func expandCloudIdentityGroupMembershipPreferredMemberKeyId(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudIdentityGroupMembershipPreferredMemberKeyNamespace(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudIdentityGroupMembershipRoles(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedName, err := expandCloudIdentityGroupMembershipRolesName(original["name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedName); val.IsValid() && !isEmptyValue(val) {
			transformed["name"] = transformedName
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandCloudIdentityGroupMembershipRolesName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
