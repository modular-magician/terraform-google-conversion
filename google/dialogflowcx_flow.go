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

func GetDialogflowCXFlowCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	name, err := assetName(d, config, "//dialogflowcx.googleapis.com/{{parent}}/flows/{{name}}")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetDialogflowCXFlowApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: "dialogflowcx.googleapis.com/Flow",
			Resource: &AssetResource{
				Version:              "v3",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/dialogflowcx/v3/rest",
				DiscoveryName:        "Flow",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
	}
}

func GetDialogflowCXFlowApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	displayNameProp, err := expandDialogflowCXFlowDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	descriptionProp, err := expandDialogflowCXFlowDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	transitionRoutesProp, err := expandDialogflowCXFlowTransitionRoutes(d.Get("transition_routes"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("transition_routes"); !isEmptyValue(reflect.ValueOf(transitionRoutesProp)) && (ok || !reflect.DeepEqual(v, transitionRoutesProp)) {
		obj["transitionRoutes"] = transitionRoutesProp
	}
	eventHandlersProp, err := expandDialogflowCXFlowEventHandlers(d.Get("event_handlers"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("event_handlers"); !isEmptyValue(reflect.ValueOf(eventHandlersProp)) && (ok || !reflect.DeepEqual(v, eventHandlersProp)) {
		obj["eventHandlers"] = eventHandlersProp
	}
	transitionRouteGroupsProp, err := expandDialogflowCXFlowTransitionRouteGroups(d.Get("transition_route_groups"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("transition_route_groups"); !isEmptyValue(reflect.ValueOf(transitionRouteGroupsProp)) && (ok || !reflect.DeepEqual(v, transitionRouteGroupsProp)) {
		obj["transitionRouteGroups"] = transitionRouteGroupsProp
	}
	nluSettingsProp, err := expandDialogflowCXFlowNluSettings(d.Get("nlu_settings"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("nlu_settings"); !isEmptyValue(reflect.ValueOf(nluSettingsProp)) && (ok || !reflect.DeepEqual(v, nluSettingsProp)) {
		obj["nluSettings"] = nluSettingsProp
	}
	languageCodeProp, err := expandDialogflowCXFlowLanguageCode(d.Get("language_code"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("language_code"); !isEmptyValue(reflect.ValueOf(languageCodeProp)) && (ok || !reflect.DeepEqual(v, languageCodeProp)) {
		obj["languageCode"] = languageCodeProp
	}

	return obj, nil
}

func expandDialogflowCXFlowDisplayName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXFlowDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXFlowTransitionRoutes(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedName, err := expandDialogflowCXFlowTransitionRoutesName(original["name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedName); val.IsValid() && !isEmptyValue(val) {
			transformed["name"] = transformedName
		}

		transformedIntent, err := expandDialogflowCXFlowTransitionRoutesIntent(original["intent"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedIntent); val.IsValid() && !isEmptyValue(val) {
			transformed["intent"] = transformedIntent
		}

		transformedCondition, err := expandDialogflowCXFlowTransitionRoutesCondition(original["condition"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedCondition); val.IsValid() && !isEmptyValue(val) {
			transformed["condition"] = transformedCondition
		}

		transformedTriggerFulfillment, err := expandDialogflowCXFlowTransitionRoutesTriggerFulfillment(original["trigger_fulfillment"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedTriggerFulfillment); val.IsValid() && !isEmptyValue(val) {
			transformed["triggerFulfillment"] = transformedTriggerFulfillment
		}

		transformedTargetPage, err := expandDialogflowCXFlowTransitionRoutesTargetPage(original["target_page"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedTargetPage); val.IsValid() && !isEmptyValue(val) {
			transformed["targetPage"] = transformedTargetPage
		}

		transformedTargetFlow, err := expandDialogflowCXFlowTransitionRoutesTargetFlow(original["target_flow"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedTargetFlow); val.IsValid() && !isEmptyValue(val) {
			transformed["targetFlow"] = transformedTargetFlow
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandDialogflowCXFlowTransitionRoutesName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXFlowTransitionRoutesIntent(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXFlowTransitionRoutesCondition(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXFlowTransitionRoutesTriggerFulfillment(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedMessages, err := expandDialogflowCXFlowTransitionRoutesTriggerFulfillmentMessages(original["messages"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMessages); val.IsValid() && !isEmptyValue(val) {
		transformed["messages"] = transformedMessages
	}

	transformedWebhook, err := expandDialogflowCXFlowTransitionRoutesTriggerFulfillmentWebhook(original["webhook"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedWebhook); val.IsValid() && !isEmptyValue(val) {
		transformed["webhook"] = transformedWebhook
	}

	transformedReturnPartialResponses, err := expandDialogflowCXFlowTransitionRoutesTriggerFulfillmentReturnPartialResponses(original["return_partial_responses"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedReturnPartialResponses); val.IsValid() && !isEmptyValue(val) {
		transformed["returnPartialResponses"] = transformedReturnPartialResponses
	}

	transformedTag, err := expandDialogflowCXFlowTransitionRoutesTriggerFulfillmentTag(original["tag"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTag); val.IsValid() && !isEmptyValue(val) {
		transformed["tag"] = transformedTag
	}

	return transformed, nil
}

func expandDialogflowCXFlowTransitionRoutesTriggerFulfillmentMessages(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedText, err := expandDialogflowCXFlowTransitionRoutesTriggerFulfillmentMessagesText(original["text"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedText); val.IsValid() && !isEmptyValue(val) {
			transformed["text"] = transformedText
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandDialogflowCXFlowTransitionRoutesTriggerFulfillmentMessagesText(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedText, err := expandDialogflowCXFlowTransitionRoutesTriggerFulfillmentMessagesTextText(original["text"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedText); val.IsValid() && !isEmptyValue(val) {
		transformed["text"] = transformedText
	}

	transformedAllowPlaybackInterruption, err := expandDialogflowCXFlowTransitionRoutesTriggerFulfillmentMessagesTextAllowPlaybackInterruption(original["allow_playback_interruption"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAllowPlaybackInterruption); val.IsValid() && !isEmptyValue(val) {
		transformed["allowPlaybackInterruption"] = transformedAllowPlaybackInterruption
	}

	return transformed, nil
}

func expandDialogflowCXFlowTransitionRoutesTriggerFulfillmentMessagesTextText(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXFlowTransitionRoutesTriggerFulfillmentMessagesTextAllowPlaybackInterruption(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXFlowTransitionRoutesTriggerFulfillmentWebhook(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXFlowTransitionRoutesTriggerFulfillmentReturnPartialResponses(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXFlowTransitionRoutesTriggerFulfillmentTag(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXFlowTransitionRoutesTargetPage(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXFlowTransitionRoutesTargetFlow(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXFlowEventHandlers(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedName, err := expandDialogflowCXFlowEventHandlersName(original["name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedName); val.IsValid() && !isEmptyValue(val) {
			transformed["name"] = transformedName
		}

		transformedEvent, err := expandDialogflowCXFlowEventHandlersEvent(original["event"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedEvent); val.IsValid() && !isEmptyValue(val) {
			transformed["event"] = transformedEvent
		}

		transformedTriggerFulfillment, err := expandDialogflowCXFlowEventHandlersTriggerFulfillment(original["trigger_fulfillment"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedTriggerFulfillment); val.IsValid() && !isEmptyValue(val) {
			transformed["triggerFulfillment"] = transformedTriggerFulfillment
		}

		transformedTargetPage, err := expandDialogflowCXFlowEventHandlersTargetPage(original["target_page"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedTargetPage); val.IsValid() && !isEmptyValue(val) {
			transformed["targetPage"] = transformedTargetPage
		}

		transformedTargetFlow, err := expandDialogflowCXFlowEventHandlersTargetFlow(original["target_flow"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedTargetFlow); val.IsValid() && !isEmptyValue(val) {
			transformed["targetFlow"] = transformedTargetFlow
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandDialogflowCXFlowEventHandlersName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXFlowEventHandlersEvent(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXFlowEventHandlersTriggerFulfillment(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedMessages, err := expandDialogflowCXFlowEventHandlersTriggerFulfillmentMessages(original["messages"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMessages); val.IsValid() && !isEmptyValue(val) {
		transformed["messages"] = transformedMessages
	}

	transformedWebhook, err := expandDialogflowCXFlowEventHandlersTriggerFulfillmentWebhook(original["webhook"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedWebhook); val.IsValid() && !isEmptyValue(val) {
		transformed["webhook"] = transformedWebhook
	}

	transformedReturnPartialResponses, err := expandDialogflowCXFlowEventHandlersTriggerFulfillmentReturnPartialResponses(original["return_partial_responses"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedReturnPartialResponses); val.IsValid() && !isEmptyValue(val) {
		transformed["returnPartialResponses"] = transformedReturnPartialResponses
	}

	transformedTag, err := expandDialogflowCXFlowEventHandlersTriggerFulfillmentTag(original["tag"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTag); val.IsValid() && !isEmptyValue(val) {
		transformed["tag"] = transformedTag
	}

	return transformed, nil
}

func expandDialogflowCXFlowEventHandlersTriggerFulfillmentMessages(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedText, err := expandDialogflowCXFlowEventHandlersTriggerFulfillmentMessagesText(original["text"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedText); val.IsValid() && !isEmptyValue(val) {
			transformed["text"] = transformedText
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandDialogflowCXFlowEventHandlersTriggerFulfillmentMessagesText(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedText, err := expandDialogflowCXFlowEventHandlersTriggerFulfillmentMessagesTextText(original["text"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedText); val.IsValid() && !isEmptyValue(val) {
		transformed["text"] = transformedText
	}

	transformedAllowPlaybackInterruption, err := expandDialogflowCXFlowEventHandlersTriggerFulfillmentMessagesTextAllowPlaybackInterruption(original["allow_playback_interruption"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAllowPlaybackInterruption); val.IsValid() && !isEmptyValue(val) {
		transformed["allowPlaybackInterruption"] = transformedAllowPlaybackInterruption
	}

	return transformed, nil
}

func expandDialogflowCXFlowEventHandlersTriggerFulfillmentMessagesTextText(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXFlowEventHandlersTriggerFulfillmentMessagesTextAllowPlaybackInterruption(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXFlowEventHandlersTriggerFulfillmentWebhook(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXFlowEventHandlersTriggerFulfillmentReturnPartialResponses(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXFlowEventHandlersTriggerFulfillmentTag(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXFlowEventHandlersTargetPage(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXFlowEventHandlersTargetFlow(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXFlowTransitionRouteGroups(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXFlowNluSettings(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedModelType, err := expandDialogflowCXFlowNluSettingsModelType(original["model_type"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedModelType); val.IsValid() && !isEmptyValue(val) {
		transformed["modelType"] = transformedModelType
	}

	transformedClassificationThreshold, err := expandDialogflowCXFlowNluSettingsClassificationThreshold(original["classification_threshold"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedClassificationThreshold); val.IsValid() && !isEmptyValue(val) {
		transformed["classificationThreshold"] = transformedClassificationThreshold
	}

	transformedModelTrainingMode, err := expandDialogflowCXFlowNluSettingsModelTrainingMode(original["model_training_mode"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedModelTrainingMode); val.IsValid() && !isEmptyValue(val) {
		transformed["modelTrainingMode"] = transformedModelTrainingMode
	}

	return transformed, nil
}

func expandDialogflowCXFlowNluSettingsModelType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXFlowNluSettingsClassificationThreshold(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXFlowNluSettingsModelTrainingMode(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXFlowLanguageCode(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
