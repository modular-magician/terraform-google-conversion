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
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func GetAppEngineStandardAppVersionCaiObject(d TerraformResourceData, config *Config) (Asset, error) {
	name, err := assetName(d, config, "//appengine.googleapis.com/apps/{{project}}/services/{{service}}/versions/{{version_id}}")
	if err != nil {
		return Asset{}, err
	}
	if obj, err := GetAppEngineStandardAppVersionApiObject(d, config); err == nil {
		return Asset{
			Name: name,
			Type: "appengine.googleapis.com/StandardAppVersion",
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/appengine/v1/rest",
				DiscoveryName:        "StandardAppVersion",
				Data:                 obj,
			},
		}, nil
	} else {
		return Asset{}, err
	}
}

func GetAppEngineStandardAppVersionApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	idProp, err := expandAppEngineStandardAppVersionVersionId(d.Get("version_id"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("version_id"); !isEmptyValue(reflect.ValueOf(idProp)) && (ok || !reflect.DeepEqual(v, idProp)) {
		obj["id"] = idProp
	}
	runtimeProp, err := expandAppEngineStandardAppVersionRuntime(d.Get("runtime"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("runtime"); !isEmptyValue(reflect.ValueOf(runtimeProp)) && (ok || !reflect.DeepEqual(v, runtimeProp)) {
		obj["runtime"] = runtimeProp
	}
	threadsafeProp, err := expandAppEngineStandardAppVersionThreadsafe(d.Get("threadsafe"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("threadsafe"); !isEmptyValue(reflect.ValueOf(threadsafeProp)) && (ok || !reflect.DeepEqual(v, threadsafeProp)) {
		obj["threadsafe"] = threadsafeProp
	}
	runtimeApiVersionProp, err := expandAppEngineStandardAppVersionRuntimeApiVersion(d.Get("runtime_api_version"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("runtime_api_version"); !isEmptyValue(reflect.ValueOf(runtimeApiVersionProp)) && (ok || !reflect.DeepEqual(v, runtimeApiVersionProp)) {
		obj["runtimeApiVersion"] = runtimeApiVersionProp
	}
	handlersProp, err := expandAppEngineStandardAppVersionHandlers(d.Get("handlers"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("handlers"); !isEmptyValue(reflect.ValueOf(handlersProp)) && (ok || !reflect.DeepEqual(v, handlersProp)) {
		obj["handlers"] = handlersProp
	}
	librariesProp, err := expandAppEngineStandardAppVersionLibraries(d.Get("libraries"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("libraries"); !isEmptyValue(reflect.ValueOf(librariesProp)) && (ok || !reflect.DeepEqual(v, librariesProp)) {
		obj["libraries"] = librariesProp
	}
	envVariablesProp, err := expandAppEngineStandardAppVersionEnvVariables(d.Get("env_variables"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("env_variables"); !isEmptyValue(reflect.ValueOf(envVariablesProp)) && (ok || !reflect.DeepEqual(v, envVariablesProp)) {
		obj["envVariables"] = envVariablesProp
	}
	deploymentProp, err := expandAppEngineStandardAppVersionDeployment(d.Get("deployment"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("deployment"); !isEmptyValue(reflect.ValueOf(deploymentProp)) && (ok || !reflect.DeepEqual(v, deploymentProp)) {
		obj["deployment"] = deploymentProp
	}
	entrypointProp, err := expandAppEngineStandardAppVersionEntrypoint(d.Get("entrypoint"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("entrypoint"); !isEmptyValue(reflect.ValueOf(entrypointProp)) && (ok || !reflect.DeepEqual(v, entrypointProp)) {
		obj["entrypoint"] = entrypointProp
	}
	instanceClassProp, err := expandAppEngineStandardAppVersionInstanceClass(d.Get("instance_class"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("instance_class"); !isEmptyValue(reflect.ValueOf(instanceClassProp)) && (ok || !reflect.DeepEqual(v, instanceClassProp)) {
		obj["instanceClass"] = instanceClassProp
	}

	return obj, nil
}

func expandAppEngineStandardAppVersionVersionId(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineStandardAppVersionRuntime(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineStandardAppVersionThreadsafe(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineStandardAppVersionRuntimeApiVersion(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineStandardAppVersionHandlers(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedUrlRegex, err := expandAppEngineStandardAppVersionHandlersUrlRegex(original["url_regex"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedUrlRegex); val.IsValid() && !isEmptyValue(val) {
			transformed["urlRegex"] = transformedUrlRegex
		}

		transformedSecurityLevel, err := expandAppEngineStandardAppVersionHandlersSecurityLevel(original["security_level"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedSecurityLevel); val.IsValid() && !isEmptyValue(val) {
			transformed["securityLevel"] = transformedSecurityLevel
		}

		transformedLogin, err := expandAppEngineStandardAppVersionHandlersLogin(original["login"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedLogin); val.IsValid() && !isEmptyValue(val) {
			transformed["login"] = transformedLogin
		}

		transformedAuthFailAction, err := expandAppEngineStandardAppVersionHandlersAuthFailAction(original["auth_fail_action"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedAuthFailAction); val.IsValid() && !isEmptyValue(val) {
			transformed["authFailAction"] = transformedAuthFailAction
		}

		transformedRedirectHttpResponseCode, err := expandAppEngineStandardAppVersionHandlersRedirectHttpResponseCode(original["redirect_http_response_code"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedRedirectHttpResponseCode); val.IsValid() && !isEmptyValue(val) {
			transformed["redirectHttpResponseCode"] = transformedRedirectHttpResponseCode
		}

		transformedScript, err := expandAppEngineStandardAppVersionHandlersScript(original["script"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedScript); val.IsValid() && !isEmptyValue(val) {
			transformed["script"] = transformedScript
		}

		transformedStaticFiles, err := expandAppEngineStandardAppVersionHandlersStaticFiles(original["static_files"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedStaticFiles); val.IsValid() && !isEmptyValue(val) {
			transformed["staticFiles"] = transformedStaticFiles
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandAppEngineStandardAppVersionHandlersUrlRegex(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineStandardAppVersionHandlersSecurityLevel(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineStandardAppVersionHandlersLogin(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineStandardAppVersionHandlersAuthFailAction(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineStandardAppVersionHandlersRedirectHttpResponseCode(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineStandardAppVersionHandlersScript(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedScriptPath, err := expandAppEngineStandardAppVersionHandlersScriptScriptPath(original["script_path"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedScriptPath); val.IsValid() && !isEmptyValue(val) {
		transformed["scriptPath"] = transformedScriptPath
	}

	return transformed, nil
}

func expandAppEngineStandardAppVersionHandlersScriptScriptPath(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineStandardAppVersionHandlersStaticFiles(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPath, err := expandAppEngineStandardAppVersionHandlersStaticFilesPath(original["path"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPath); val.IsValid() && !isEmptyValue(val) {
		transformed["path"] = transformedPath
	}

	transformedUploadPathRegex, err := expandAppEngineStandardAppVersionHandlersStaticFilesUploadPathRegex(original["upload_path_regex"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUploadPathRegex); val.IsValid() && !isEmptyValue(val) {
		transformed["uploadPathRegex"] = transformedUploadPathRegex
	}

	transformedHttpHeaders, err := expandAppEngineStandardAppVersionHandlersStaticFilesHttpHeaders(original["http_headers"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedHttpHeaders); val.IsValid() && !isEmptyValue(val) {
		transformed["httpHeaders"] = transformedHttpHeaders
	}

	transformedMimeType, err := expandAppEngineStandardAppVersionHandlersStaticFilesMimeType(original["mime_type"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMimeType); val.IsValid() && !isEmptyValue(val) {
		transformed["mimeType"] = transformedMimeType
	}

	transformedExpiration, err := expandAppEngineStandardAppVersionHandlersStaticFilesExpiration(original["expiration"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedExpiration); val.IsValid() && !isEmptyValue(val) {
		transformed["expiration"] = transformedExpiration
	}

	transformedRequireMatchingFile, err := expandAppEngineStandardAppVersionHandlersStaticFilesRequireMatchingFile(original["require_matching_file"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRequireMatchingFile); val.IsValid() && !isEmptyValue(val) {
		transformed["requireMatchingFile"] = transformedRequireMatchingFile
	}

	transformedApplicationReadable, err := expandAppEngineStandardAppVersionHandlersStaticFilesApplicationReadable(original["application_readable"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedApplicationReadable); val.IsValid() && !isEmptyValue(val) {
		transformed["applicationReadable"] = transformedApplicationReadable
	}

	return transformed, nil
}

func expandAppEngineStandardAppVersionHandlersStaticFilesPath(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineStandardAppVersionHandlersStaticFilesUploadPathRegex(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineStandardAppVersionHandlersStaticFilesHttpHeaders(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandAppEngineStandardAppVersionHandlersStaticFilesMimeType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineStandardAppVersionHandlersStaticFilesExpiration(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineStandardAppVersionHandlersStaticFilesRequireMatchingFile(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineStandardAppVersionHandlersStaticFilesApplicationReadable(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineStandardAppVersionLibraries(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedName, err := expandAppEngineStandardAppVersionLibrariesName(original["name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedName); val.IsValid() && !isEmptyValue(val) {
			transformed["name"] = transformedName
		}

		transformedVersion, err := expandAppEngineStandardAppVersionLibrariesVersion(original["version"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedVersion); val.IsValid() && !isEmptyValue(val) {
			transformed["version"] = transformedVersion
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandAppEngineStandardAppVersionLibrariesName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineStandardAppVersionLibrariesVersion(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineStandardAppVersionEnvVariables(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandAppEngineStandardAppVersionDeployment(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedZip, err := expandAppEngineStandardAppVersionDeploymentZip(original["zip"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedZip); val.IsValid() && !isEmptyValue(val) {
		transformed["zip"] = transformedZip
	}

	transformedFiles, err := expandAppEngineStandardAppVersionDeploymentFiles(original["files"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedFiles); val.IsValid() && !isEmptyValue(val) {
		transformed["files"] = transformedFiles
	}

	return transformed, nil
}

func expandAppEngineStandardAppVersionDeploymentZip(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedSourceUrl, err := expandAppEngineStandardAppVersionDeploymentZipSourceUrl(original["source_url"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSourceUrl); val.IsValid() && !isEmptyValue(val) {
		transformed["sourceUrl"] = transformedSourceUrl
	}

	transformedFilesCount, err := expandAppEngineStandardAppVersionDeploymentZipFilesCount(original["files_count"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedFilesCount); val.IsValid() && !isEmptyValue(val) {
		transformed["filesCount"] = transformedFilesCount
	}

	return transformed, nil
}

func expandAppEngineStandardAppVersionDeploymentZipSourceUrl(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineStandardAppVersionDeploymentZipFilesCount(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineStandardAppVersionDeploymentFiles(v interface{}, d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	if v == nil {
		return map[string]interface{}{}, nil
	}
	m := make(map[string]interface{})
	for _, raw := range v.(*schema.Set).List() {
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedSha1Sum, err := expandAppEngineStandardAppVersionDeploymentFilesSha1Sum(original["sha1_sum"], d, config)
		if err != nil {
			return nil, err
		}
		transformed["sha1Sum"] = transformedSha1Sum
		transformedSourceUrl, err := expandAppEngineStandardAppVersionDeploymentFilesSourceUrl(original["source_url"], d, config)
		if err != nil {
			return nil, err
		}
		transformed["sourceUrl"] = transformedSourceUrl

		m[original["name"].(string)] = transformed
	}
	return m, nil
}

func expandAppEngineStandardAppVersionDeploymentFilesSha1Sum(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineStandardAppVersionDeploymentFilesSourceUrl(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineStandardAppVersionEntrypoint(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedShell, err := expandAppEngineStandardAppVersionEntrypointShell(original["shell"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedShell); val.IsValid() && !isEmptyValue(val) {
		transformed["shell"] = transformedShell
	}

	return transformed, nil
}

func expandAppEngineStandardAppVersionEntrypointShell(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineStandardAppVersionInstanceClass(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
