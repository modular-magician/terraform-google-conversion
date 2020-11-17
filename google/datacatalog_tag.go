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

func GetDataCatalogTagCaiObject(d TerraformResourceData, config *Config) (Asset, error) {
	name, err := assetName(d, config, "//datacatalog.googleapis.com/{{parent}}/tags")
	if err != nil {
		return Asset{}, err
	}
	if obj, err := GetDataCatalogTagApiObject(d, config); err == nil {
		return Asset{
			Name: name,
			Type: "datacatalog.googleapis.com/Tag",
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/datacatalog/v1/rest",
				DiscoveryName:        "Tag",
				Data:                 obj,
			},
		}, nil
	} else {
		return Asset{}, err
	}
}

func GetDataCatalogTagApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	templateProp, err := expandDataCatalogTagTemplate(d.Get("template"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("template"); !isEmptyValue(reflect.ValueOf(templateProp)) && (ok || !reflect.DeepEqual(v, templateProp)) {
		obj["template"] = templateProp
	}
	fieldsProp, err := expandDataCatalogTagFields(d.Get("fields"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("fields"); !isEmptyValue(reflect.ValueOf(fieldsProp)) && (ok || !reflect.DeepEqual(v, fieldsProp)) {
		obj["fields"] = fieldsProp
	}
	columnProp, err := expandDataCatalogTagColumn(d.Get("column"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("column"); !isEmptyValue(reflect.ValueOf(columnProp)) && (ok || !reflect.DeepEqual(v, columnProp)) {
		obj["column"] = columnProp
	}

	return resourceDataCatalogTagEncoder(d, config, obj)
}

func resourceDataCatalogTagEncoder(d TerraformResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	if obj["fields"] != nil {
		// isEmptyValue() does not work for a boolean as it shows
		// false when it is 'empty'. Filter boolValue here based on
		// the rule api does not take more than 1 'value'
		fields := obj["fields"].(map[string]interface{})
		for _, elements := range fields {
			values := elements.(map[string]interface{})
			if len(values) > 1 {
				for val := range values {
					if val == "boolValue" {
						delete(values, "boolValue")
					}
				}
			}
		}
	}
	return obj, nil
}

func expandDataCatalogTagTemplate(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataCatalogTagFields(v interface{}, d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	if v == nil {
		return map[string]interface{}{}, nil
	}
	m := make(map[string]interface{})
	for _, raw := range v.(*schema.Set).List() {
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedDisplayName, err := expandDataCatalogTagFieldsDisplayName(original["display_name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDisplayName); val.IsValid() && !isEmptyValue(val) {
			transformed["display_name"] = transformedDisplayName
		}

		transformedOrder, err := expandDataCatalogTagFieldsOrder(original["order"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedOrder); val.IsValid() && !isEmptyValue(val) {
			transformed["order"] = transformedOrder
		}

		transformedDoubleValue, err := expandDataCatalogTagFieldsDoubleValue(original["double_value"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDoubleValue); val.IsValid() && !isEmptyValue(val) {
			transformed["doubleValue"] = transformedDoubleValue
		}

		transformedStringValue, err := expandDataCatalogTagFieldsStringValue(original["string_value"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedStringValue); val.IsValid() && !isEmptyValue(val) {
			transformed["stringValue"] = transformedStringValue
		}

		transformedBoolValue, err := expandDataCatalogTagFieldsBoolValue(original["bool_value"], d, config)
		if err != nil {
			return nil, err
		} else {
			transformed["boolValue"] = transformedBoolValue
		}

		transformedTimestampValue, err := expandDataCatalogTagFieldsTimestampValue(original["timestamp_value"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedTimestampValue); val.IsValid() && !isEmptyValue(val) {
			transformed["timestampValue"] = transformedTimestampValue
		}

		transformedEnumValue, err := expandDataCatalogTagFieldsEnumValue(original["enum_value"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedEnumValue); val.IsValid() && !isEmptyValue(val) {
			transformed["enumValue"] = transformedEnumValue
		}

		transformedFieldName, err := expandString(original["field_name"], d, config)
		if err != nil {
			return nil, err
		}
		m[transformedFieldName] = transformed
	}
	return m, nil
}

func expandDataCatalogTagFieldsDisplayName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataCatalogTagFieldsOrder(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataCatalogTagFieldsDoubleValue(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataCatalogTagFieldsStringValue(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataCatalogTagFieldsBoolValue(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataCatalogTagFieldsTimestampValue(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataCatalogTagFieldsEnumValue(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	// we flattened the original["enum_value"]["display_name"] object to be just original["enum_value"] so here,
	// v is the value we want from the config
	transformed := make(map[string]interface{})
	if val := reflect.ValueOf(v); val.IsValid() && !isEmptyValue(val) {
		transformed["displayName"] = v
	}

	return transformed, nil
}

func expandDataCatalogTagColumn(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
