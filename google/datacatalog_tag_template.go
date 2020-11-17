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
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func GetDataCatalogTagTemplateCaiObject(d TerraformResourceData, config *Config) (Asset, error) {
	name, err := assetName(d, config, "//datacatalog.googleapis.com/{{name}}")
	if err != nil {
		return Asset{}, err
	}
	if obj, err := GetDataCatalogTagTemplateApiObject(d, config); err == nil {
		return Asset{
			Name: name,
			Type: "datacatalog.googleapis.com/TagTemplate",
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/datacatalog/v1/rest",
				DiscoveryName:        "TagTemplate",
				Data:                 obj,
			},
		}, nil
	} else {
		return Asset{}, err
	}
}

func GetDataCatalogTagTemplateApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	displayNameProp, err := expandDataCatalogTagTemplateDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	fieldsProp, err := expandDataCatalogTagTemplateFields(d.Get("fields"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("fields"); !isEmptyValue(reflect.ValueOf(fieldsProp)) && (ok || !reflect.DeepEqual(v, fieldsProp)) {
		obj["fields"] = fieldsProp
	}

	return obj, nil
}

func expandDataCatalogTagTemplateDisplayName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataCatalogTagTemplateFields(v interface{}, d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	if v == nil {
		return map[string]interface{}{}, nil
	}
	m := make(map[string]interface{})
	for _, raw := range v.(*schema.Set).List() {
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedName, err := expandDataCatalogTagTemplateFieldsName(original["name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedName); val.IsValid() && !isEmptyValue(val) {
			transformed["name"] = transformedName
		}

		transformedDisplayName, err := expandDataCatalogTagTemplateFieldsDisplayName(original["display_name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDisplayName); val.IsValid() && !isEmptyValue(val) {
			transformed["displayName"] = transformedDisplayName
		}

		transformedType, err := expandDataCatalogTagTemplateFieldsType(original["type"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedType); val.IsValid() && !isEmptyValue(val) {
			transformed["type"] = transformedType
		}

		transformedIsRequired, err := expandDataCatalogTagTemplateFieldsIsRequired(original["is_required"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedIsRequired); val.IsValid() && !isEmptyValue(val) {
			transformed["isRequired"] = transformedIsRequired
		}

		transformedOrder, err := expandDataCatalogTagTemplateFieldsOrder(original["order"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedOrder); val.IsValid() && !isEmptyValue(val) {
			transformed["order"] = transformedOrder
		}

		transformedFieldId, err := expandString(original["field_id"], d, config)
		if err != nil {
			return nil, err
		}
		m[transformedFieldId] = transformed
	}
	return m, nil
}

func expandDataCatalogTagTemplateFieldsName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataCatalogTagTemplateFieldsDisplayName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataCatalogTagTemplateFieldsType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPrimitiveType, err := expandDataCatalogTagTemplateFieldsTypePrimitiveType(original["primitive_type"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPrimitiveType); val.IsValid() && !isEmptyValue(val) {
		transformed["primitiveType"] = transformedPrimitiveType
	}

	transformedEnumType, err := expandDataCatalogTagTemplateFieldsTypeEnumType(original["enum_type"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEnumType); val.IsValid() && !isEmptyValue(val) {
		transformed["enumType"] = transformedEnumType
	}

	return transformed, nil
}

func expandDataCatalogTagTemplateFieldsTypePrimitiveType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataCatalogTagTemplateFieldsTypeEnumType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedAllowedValues, err := expandDataCatalogTagTemplateFieldsTypeEnumTypeAllowedValues(original["allowed_values"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAllowedValues); val.IsValid() && !isEmptyValue(val) {
		transformed["allowedValues"] = transformedAllowedValues
	}

	return transformed, nil
}

func expandDataCatalogTagTemplateFieldsTypeEnumTypeAllowedValues(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedDisplayName, err := expandDataCatalogTagTemplateFieldsTypeEnumTypeAllowedValuesDisplayName(original["display_name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDisplayName); val.IsValid() && !isEmptyValue(val) {
			transformed["displayName"] = transformedDisplayName
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandDataCatalogTagTemplateFieldsTypeEnumTypeAllowedValuesDisplayName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataCatalogTagTemplateFieldsIsRequired(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataCatalogTagTemplateFieldsOrder(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
