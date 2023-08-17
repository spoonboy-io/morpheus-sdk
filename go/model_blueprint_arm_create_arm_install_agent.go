/*
Morpheus API

Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.

API version: 6.1.1
Contact: dev@morpheusdata.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// BlueprintARMCreateArmInstallAgent - Install Morpheus Agent
type BlueprintARMCreateArmInstallAgent struct {
	Bool *bool
	String *string
}

// boolAsBlueprintARMCreateArmInstallAgent is a convenience function that returns bool wrapped in BlueprintARMCreateArmInstallAgent
func BoolAsBlueprintARMCreateArmInstallAgent(v *bool) BlueprintARMCreateArmInstallAgent {
	return BlueprintARMCreateArmInstallAgent{
		Bool: v,
	}
}

// stringAsBlueprintARMCreateArmInstallAgent is a convenience function that returns string wrapped in BlueprintARMCreateArmInstallAgent
func StringAsBlueprintARMCreateArmInstallAgent(v *string) BlueprintARMCreateArmInstallAgent {
	return BlueprintARMCreateArmInstallAgent{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *BlueprintARMCreateArmInstallAgent) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Bool
	err = newStrictDecoder(data).Decode(&dst.Bool)
	if err == nil {
		jsonBool, _ := json.Marshal(dst.Bool)
		if string(jsonBool) == "{}" { // empty struct
			dst.Bool = nil
		} else {
			match++
		}
	} else {
		dst.Bool = nil
	}

	// try to unmarshal data into String
	err = newStrictDecoder(data).Decode(&dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			match++
		}
	} else {
		dst.String = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Bool = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(BlueprintARMCreateArmInstallAgent)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(BlueprintARMCreateArmInstallAgent)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BlueprintARMCreateArmInstallAgent) MarshalJSON() ([]byte, error) {
	if src.Bool != nil {
		return json.Marshal(&src.Bool)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *BlueprintARMCreateArmInstallAgent) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Bool != nil {
		return obj.Bool
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableBlueprintARMCreateArmInstallAgent struct {
	value *BlueprintARMCreateArmInstallAgent
	isSet bool
}

func (v NullableBlueprintARMCreateArmInstallAgent) Get() *BlueprintARMCreateArmInstallAgent {
	return v.value
}

func (v *NullableBlueprintARMCreateArmInstallAgent) Set(val *BlueprintARMCreateArmInstallAgent) {
	v.value = val
	v.isSet = true
}

func (v NullableBlueprintARMCreateArmInstallAgent) IsSet() bool {
	return v.isSet
}

func (v *NullableBlueprintARMCreateArmInstallAgent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlueprintARMCreateArmInstallAgent(val *BlueprintARMCreateArmInstallAgent) *NullableBlueprintARMCreateArmInstallAgent {
	return &NullableBlueprintARMCreateArmInstallAgent{value: val, isSet: true}
}

func (v NullableBlueprintARMCreateArmInstallAgent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlueprintARMCreateArmInstallAgent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

