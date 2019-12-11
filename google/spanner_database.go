// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Nathan is editing this to generate diffs in lots of files.
//     He won't submit this change.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"reflect"
)

func GetSpannerDatabaseCaiObject(d TerraformResourceData, config *Config) (Asset, error) {
	name, err := assetName(d, config, "//spanner.googleapis.com/projects/{{project}}/instances/{{instance}}/databases/{{name}}")
	if err != nil {
		return Asset{}, err
	}
	if obj, err := GetSpannerDatabaseApiObject(d, config); err == nil {
		return Asset{
			Name: name,
			Type: "spanner.googleapis.com/Database",
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/spanner/v1/rest",
				DiscoveryName:        "Database",
				Data:                 obj,
			},
		}, nil
	} else {
		return Asset{}, err
	}
}

func GetSpannerDatabaseApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandSpannerDatabaseName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	extraStatementsProp, err := expandSpannerDatabaseDdl(d.Get("ddl"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("ddl"); !isEmptyValue(reflect.ValueOf(extraStatementsProp)) && (ok || !reflect.DeepEqual(v, extraStatementsProp)) {
		obj["extraStatements"] = extraStatementsProp
	}
	instanceProp, err := expandSpannerDatabaseInstance(d.Get("instance"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("instance"); !isEmptyValue(reflect.ValueOf(instanceProp)) && (ok || !reflect.DeepEqual(v, instanceProp)) {
		obj["instance"] = instanceProp
	}

	return resourceSpannerDatabaseEncoder(d, config, obj)
}

func resourceSpannerDatabaseEncoder(d TerraformResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	obj["createStatement"] = fmt.Sprintf("CREATE DATABASE `%s`", obj["name"])
	delete(obj, "name")
	delete(obj, "instance")
	return obj, nil
}

func expandSpannerDatabaseName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandSpannerDatabaseDdl(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandSpannerDatabaseInstance(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseGlobalFieldValue("instances", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for instance: %s", err)
	}
	return f.RelativeLink(), nil
}
