/*
 * Morpheus API
 * Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.
 *
 * The version of the OpenAPI document: 6.2.1
 * Contact: dev@morpheusdata.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package org.openapitools.client.model;

import java.util.Objects;
import java.util.Arrays;
import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.annotations.SerializedName;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.io.IOException;

/**
 * NetworkRouterTypesOptionTypes
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class NetworkRouterTypesOptionTypes {
  public static final String SERIALIZED_NAME_ID = "id";
  @SerializedName(SERIALIZED_NAME_ID)
  private Long id;

  public static final String SERIALIZED_NAME_NAME = "name";
  @SerializedName(SERIALIZED_NAME_NAME)
  private String name;

  public static final String SERIALIZED_NAME_DESCRIPTION = "description";
  @SerializedName(SERIALIZED_NAME_DESCRIPTION)
  private String description;

  public static final String SERIALIZED_NAME_CODE = "code";
  @SerializedName(SERIALIZED_NAME_CODE)
  private String code;

  public static final String SERIALIZED_NAME_FIELD_NAME = "fieldName";
  @SerializedName(SERIALIZED_NAME_FIELD_NAME)
  private String fieldName;

  public static final String SERIALIZED_NAME_FIELD_LABEL = "fieldLabel";
  @SerializedName(SERIALIZED_NAME_FIELD_LABEL)
  private String fieldLabel;

  public static final String SERIALIZED_NAME_FIELD_CONTEXT = "fieldContext";
  @SerializedName(SERIALIZED_NAME_FIELD_CONTEXT)
  private String fieldContext;

  public static final String SERIALIZED_NAME_FIELD_GROUP = "fieldGroup";
  @SerializedName(SERIALIZED_NAME_FIELD_GROUP)
  private String fieldGroup;

  public static final String SERIALIZED_NAME_FIELD_CLASS = "fieldClass";
  @SerializedName(SERIALIZED_NAME_FIELD_CLASS)
  private String fieldClass;

  public static final String SERIALIZED_NAME_FIELD_ADD_ON = "fieldAddOn";
  @SerializedName(SERIALIZED_NAME_FIELD_ADD_ON)
  private String fieldAddOn;

  public static final String SERIALIZED_NAME_FIELD_COMPONENT = "fieldComponent";
  @SerializedName(SERIALIZED_NAME_FIELD_COMPONENT)
  private String fieldComponent;

  public static final String SERIALIZED_NAME_FIELD_INPUT = "fieldInput";
  @SerializedName(SERIALIZED_NAME_FIELD_INPUT)
  private String fieldInput;

  public static final String SERIALIZED_NAME_PLACE_HOLDER = "placeHolder";
  @SerializedName(SERIALIZED_NAME_PLACE_HOLDER)
  private String placeHolder;

  public static final String SERIALIZED_NAME_HELP_BLOCK = "helpBlock";
  @SerializedName(SERIALIZED_NAME_HELP_BLOCK)
  private String helpBlock;

  public static final String SERIALIZED_NAME_DEFAULT_VALUE = "defaultValue";
  @SerializedName(SERIALIZED_NAME_DEFAULT_VALUE)
  private String defaultValue;

  public static final String SERIALIZED_NAME_OPTION_SOURCE = "optionSource";
  @SerializedName(SERIALIZED_NAME_OPTION_SOURCE)
  private String optionSource;

  public static final String SERIALIZED_NAME_OPTION_LIST = "optionList";
  @SerializedName(SERIALIZED_NAME_OPTION_LIST)
  private String optionList;

  public static final String SERIALIZED_NAME_TYPE = "type";
  @SerializedName(SERIALIZED_NAME_TYPE)
  private String type;

  public static final String SERIALIZED_NAME_ADVANCED = "advanced";
  @SerializedName(SERIALIZED_NAME_ADVANCED)
  private Boolean advanced;

  public static final String SERIALIZED_NAME_REQUIRED = "required";
  @SerializedName(SERIALIZED_NAME_REQUIRED)
  private Boolean required;

  public static final String SERIALIZED_NAME_EDITABLE = "editable";
  @SerializedName(SERIALIZED_NAME_EDITABLE)
  private Boolean editable;

  public static final String SERIALIZED_NAME_CREATABLE = "creatable";
  @SerializedName(SERIALIZED_NAME_CREATABLE)
  private Boolean creatable;

  public static final String SERIALIZED_NAME_CONFIG = "config";
  @SerializedName(SERIALIZED_NAME_CONFIG)
  private Object config;

  public static final String SERIALIZED_NAME_DISPLAY_ORDER = "displayOrder";
  @SerializedName(SERIALIZED_NAME_DISPLAY_ORDER)
  private Long displayOrder;

  public static final String SERIALIZED_NAME_WRAPPER_CLASS = "wrapperClass";
  @SerializedName(SERIALIZED_NAME_WRAPPER_CLASS)
  private String wrapperClass;

  public static final String SERIALIZED_NAME_ENABLED = "enabled";
  @SerializedName(SERIALIZED_NAME_ENABLED)
  private Boolean enabled;

  public static final String SERIALIZED_NAME_NO_BLANK = "noBlank";
  @SerializedName(SERIALIZED_NAME_NO_BLANK)
  private Boolean noBlank;

  public static final String SERIALIZED_NAME_DEPENDS_ON_CODE = "dependsOnCode";
  @SerializedName(SERIALIZED_NAME_DEPENDS_ON_CODE)
  private String dependsOnCode;

  public static final String SERIALIZED_NAME_CONTEXTUAL_DEFAULT = "contextualDefault";
  @SerializedName(SERIALIZED_NAME_CONTEXTUAL_DEFAULT)
  private Boolean contextualDefault;


  public NetworkRouterTypesOptionTypes id(Long id) {
    
    this.id = id;
    return this;
  }

   /**
   * Get id
   * @return id
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getId() {
    return id;
  }


  public void setId(Long id) {
    this.id = id;
  }


  public NetworkRouterTypesOptionTypes name(String name) {
    
    this.name = name;
    return this;
  }

   /**
   * Get name
   * @return name
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getName() {
    return name;
  }


  public void setName(String name) {
    this.name = name;
  }


  public NetworkRouterTypesOptionTypes description(String description) {
    
    this.description = description;
    return this;
  }

   /**
   * Get description
   * @return description
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getDescription() {
    return description;
  }


  public void setDescription(String description) {
    this.description = description;
  }


  public NetworkRouterTypesOptionTypes code(String code) {
    
    this.code = code;
    return this;
  }

   /**
   * Get code
   * @return code
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getCode() {
    return code;
  }


  public void setCode(String code) {
    this.code = code;
  }


  public NetworkRouterTypesOptionTypes fieldName(String fieldName) {
    
    this.fieldName = fieldName;
    return this;
  }

   /**
   * Get fieldName
   * @return fieldName
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getFieldName() {
    return fieldName;
  }


  public void setFieldName(String fieldName) {
    this.fieldName = fieldName;
  }


  public NetworkRouterTypesOptionTypes fieldLabel(String fieldLabel) {
    
    this.fieldLabel = fieldLabel;
    return this;
  }

   /**
   * Get fieldLabel
   * @return fieldLabel
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getFieldLabel() {
    return fieldLabel;
  }


  public void setFieldLabel(String fieldLabel) {
    this.fieldLabel = fieldLabel;
  }


  public NetworkRouterTypesOptionTypes fieldContext(String fieldContext) {
    
    this.fieldContext = fieldContext;
    return this;
  }

   /**
   * Get fieldContext
   * @return fieldContext
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getFieldContext() {
    return fieldContext;
  }


  public void setFieldContext(String fieldContext) {
    this.fieldContext = fieldContext;
  }


  public NetworkRouterTypesOptionTypes fieldGroup(String fieldGroup) {
    
    this.fieldGroup = fieldGroup;
    return this;
  }

   /**
   * Get fieldGroup
   * @return fieldGroup
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getFieldGroup() {
    return fieldGroup;
  }


  public void setFieldGroup(String fieldGroup) {
    this.fieldGroup = fieldGroup;
  }


  public NetworkRouterTypesOptionTypes fieldClass(String fieldClass) {
    
    this.fieldClass = fieldClass;
    return this;
  }

   /**
   * Get fieldClass
   * @return fieldClass
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getFieldClass() {
    return fieldClass;
  }


  public void setFieldClass(String fieldClass) {
    this.fieldClass = fieldClass;
  }


  public NetworkRouterTypesOptionTypes fieldAddOn(String fieldAddOn) {
    
    this.fieldAddOn = fieldAddOn;
    return this;
  }

   /**
   * Get fieldAddOn
   * @return fieldAddOn
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getFieldAddOn() {
    return fieldAddOn;
  }


  public void setFieldAddOn(String fieldAddOn) {
    this.fieldAddOn = fieldAddOn;
  }


  public NetworkRouterTypesOptionTypes fieldComponent(String fieldComponent) {
    
    this.fieldComponent = fieldComponent;
    return this;
  }

   /**
   * Get fieldComponent
   * @return fieldComponent
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getFieldComponent() {
    return fieldComponent;
  }


  public void setFieldComponent(String fieldComponent) {
    this.fieldComponent = fieldComponent;
  }


  public NetworkRouterTypesOptionTypes fieldInput(String fieldInput) {
    
    this.fieldInput = fieldInput;
    return this;
  }

   /**
   * Get fieldInput
   * @return fieldInput
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getFieldInput() {
    return fieldInput;
  }


  public void setFieldInput(String fieldInput) {
    this.fieldInput = fieldInput;
  }


  public NetworkRouterTypesOptionTypes placeHolder(String placeHolder) {
    
    this.placeHolder = placeHolder;
    return this;
  }

   /**
   * Get placeHolder
   * @return placeHolder
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getPlaceHolder() {
    return placeHolder;
  }


  public void setPlaceHolder(String placeHolder) {
    this.placeHolder = placeHolder;
  }


  public NetworkRouterTypesOptionTypes helpBlock(String helpBlock) {
    
    this.helpBlock = helpBlock;
    return this;
  }

   /**
   * Get helpBlock
   * @return helpBlock
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getHelpBlock() {
    return helpBlock;
  }


  public void setHelpBlock(String helpBlock) {
    this.helpBlock = helpBlock;
  }


  public NetworkRouterTypesOptionTypes defaultValue(String defaultValue) {
    
    this.defaultValue = defaultValue;
    return this;
  }

   /**
   * Get defaultValue
   * @return defaultValue
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getDefaultValue() {
    return defaultValue;
  }


  public void setDefaultValue(String defaultValue) {
    this.defaultValue = defaultValue;
  }


  public NetworkRouterTypesOptionTypes optionSource(String optionSource) {
    
    this.optionSource = optionSource;
    return this;
  }

   /**
   * Get optionSource
   * @return optionSource
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getOptionSource() {
    return optionSource;
  }


  public void setOptionSource(String optionSource) {
    this.optionSource = optionSource;
  }


  public NetworkRouterTypesOptionTypes optionList(String optionList) {
    
    this.optionList = optionList;
    return this;
  }

   /**
   * Get optionList
   * @return optionList
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getOptionList() {
    return optionList;
  }


  public void setOptionList(String optionList) {
    this.optionList = optionList;
  }


  public NetworkRouterTypesOptionTypes type(String type) {
    
    this.type = type;
    return this;
  }

   /**
   * Get type
   * @return type
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getType() {
    return type;
  }


  public void setType(String type) {
    this.type = type;
  }


  public NetworkRouterTypesOptionTypes advanced(Boolean advanced) {
    
    this.advanced = advanced;
    return this;
  }

   /**
   * Get advanced
   * @return advanced
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getAdvanced() {
    return advanced;
  }


  public void setAdvanced(Boolean advanced) {
    this.advanced = advanced;
  }


  public NetworkRouterTypesOptionTypes required(Boolean required) {
    
    this.required = required;
    return this;
  }

   /**
   * Get required
   * @return required
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getRequired() {
    return required;
  }


  public void setRequired(Boolean required) {
    this.required = required;
  }


  public NetworkRouterTypesOptionTypes editable(Boolean editable) {
    
    this.editable = editable;
    return this;
  }

   /**
   * Get editable
   * @return editable
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getEditable() {
    return editable;
  }


  public void setEditable(Boolean editable) {
    this.editable = editable;
  }


  public NetworkRouterTypesOptionTypes creatable(Boolean creatable) {
    
    this.creatable = creatable;
    return this;
  }

   /**
   * Get creatable
   * @return creatable
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getCreatable() {
    return creatable;
  }


  public void setCreatable(Boolean creatable) {
    this.creatable = creatable;
  }


  public NetworkRouterTypesOptionTypes config(Object config) {
    
    this.config = config;
    return this;
  }

   /**
   * Get config
   * @return config
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Object getConfig() {
    return config;
  }


  public void setConfig(Object config) {
    this.config = config;
  }


  public NetworkRouterTypesOptionTypes displayOrder(Long displayOrder) {
    
    this.displayOrder = displayOrder;
    return this;
  }

   /**
   * Get displayOrder
   * @return displayOrder
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getDisplayOrder() {
    return displayOrder;
  }


  public void setDisplayOrder(Long displayOrder) {
    this.displayOrder = displayOrder;
  }


  public NetworkRouterTypesOptionTypes wrapperClass(String wrapperClass) {
    
    this.wrapperClass = wrapperClass;
    return this;
  }

   /**
   * Get wrapperClass
   * @return wrapperClass
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getWrapperClass() {
    return wrapperClass;
  }


  public void setWrapperClass(String wrapperClass) {
    this.wrapperClass = wrapperClass;
  }


  public NetworkRouterTypesOptionTypes enabled(Boolean enabled) {
    
    this.enabled = enabled;
    return this;
  }

   /**
   * Get enabled
   * @return enabled
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getEnabled() {
    return enabled;
  }


  public void setEnabled(Boolean enabled) {
    this.enabled = enabled;
  }


  public NetworkRouterTypesOptionTypes noBlank(Boolean noBlank) {
    
    this.noBlank = noBlank;
    return this;
  }

   /**
   * Get noBlank
   * @return noBlank
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getNoBlank() {
    return noBlank;
  }


  public void setNoBlank(Boolean noBlank) {
    this.noBlank = noBlank;
  }


  public NetworkRouterTypesOptionTypes dependsOnCode(String dependsOnCode) {
    
    this.dependsOnCode = dependsOnCode;
    return this;
  }

   /**
   * Get dependsOnCode
   * @return dependsOnCode
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getDependsOnCode() {
    return dependsOnCode;
  }


  public void setDependsOnCode(String dependsOnCode) {
    this.dependsOnCode = dependsOnCode;
  }


  public NetworkRouterTypesOptionTypes contextualDefault(Boolean contextualDefault) {
    
    this.contextualDefault = contextualDefault;
    return this;
  }

   /**
   * Get contextualDefault
   * @return contextualDefault
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getContextualDefault() {
    return contextualDefault;
  }


  public void setContextualDefault(Boolean contextualDefault) {
    this.contextualDefault = contextualDefault;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    NetworkRouterTypesOptionTypes networkRouterTypesOptionTypes = (NetworkRouterTypesOptionTypes) o;
    return Objects.equals(this.id, networkRouterTypesOptionTypes.id) &&
        Objects.equals(this.name, networkRouterTypesOptionTypes.name) &&
        Objects.equals(this.description, networkRouterTypesOptionTypes.description) &&
        Objects.equals(this.code, networkRouterTypesOptionTypes.code) &&
        Objects.equals(this.fieldName, networkRouterTypesOptionTypes.fieldName) &&
        Objects.equals(this.fieldLabel, networkRouterTypesOptionTypes.fieldLabel) &&
        Objects.equals(this.fieldContext, networkRouterTypesOptionTypes.fieldContext) &&
        Objects.equals(this.fieldGroup, networkRouterTypesOptionTypes.fieldGroup) &&
        Objects.equals(this.fieldClass, networkRouterTypesOptionTypes.fieldClass) &&
        Objects.equals(this.fieldAddOn, networkRouterTypesOptionTypes.fieldAddOn) &&
        Objects.equals(this.fieldComponent, networkRouterTypesOptionTypes.fieldComponent) &&
        Objects.equals(this.fieldInput, networkRouterTypesOptionTypes.fieldInput) &&
        Objects.equals(this.placeHolder, networkRouterTypesOptionTypes.placeHolder) &&
        Objects.equals(this.helpBlock, networkRouterTypesOptionTypes.helpBlock) &&
        Objects.equals(this.defaultValue, networkRouterTypesOptionTypes.defaultValue) &&
        Objects.equals(this.optionSource, networkRouterTypesOptionTypes.optionSource) &&
        Objects.equals(this.optionList, networkRouterTypesOptionTypes.optionList) &&
        Objects.equals(this.type, networkRouterTypesOptionTypes.type) &&
        Objects.equals(this.advanced, networkRouterTypesOptionTypes.advanced) &&
        Objects.equals(this.required, networkRouterTypesOptionTypes.required) &&
        Objects.equals(this.editable, networkRouterTypesOptionTypes.editable) &&
        Objects.equals(this.creatable, networkRouterTypesOptionTypes.creatable) &&
        Objects.equals(this.config, networkRouterTypesOptionTypes.config) &&
        Objects.equals(this.displayOrder, networkRouterTypesOptionTypes.displayOrder) &&
        Objects.equals(this.wrapperClass, networkRouterTypesOptionTypes.wrapperClass) &&
        Objects.equals(this.enabled, networkRouterTypesOptionTypes.enabled) &&
        Objects.equals(this.noBlank, networkRouterTypesOptionTypes.noBlank) &&
        Objects.equals(this.dependsOnCode, networkRouterTypesOptionTypes.dependsOnCode) &&
        Objects.equals(this.contextualDefault, networkRouterTypesOptionTypes.contextualDefault);
  }

  @Override
  public int hashCode() {
    return Objects.hash(id, name, description, code, fieldName, fieldLabel, fieldContext, fieldGroup, fieldClass, fieldAddOn, fieldComponent, fieldInput, placeHolder, helpBlock, defaultValue, optionSource, optionList, type, advanced, required, editable, creatable, config, displayOrder, wrapperClass, enabled, noBlank, dependsOnCode, contextualDefault);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class NetworkRouterTypesOptionTypes {\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    description: ").append(toIndentedString(description)).append("\n");
    sb.append("    code: ").append(toIndentedString(code)).append("\n");
    sb.append("    fieldName: ").append(toIndentedString(fieldName)).append("\n");
    sb.append("    fieldLabel: ").append(toIndentedString(fieldLabel)).append("\n");
    sb.append("    fieldContext: ").append(toIndentedString(fieldContext)).append("\n");
    sb.append("    fieldGroup: ").append(toIndentedString(fieldGroup)).append("\n");
    sb.append("    fieldClass: ").append(toIndentedString(fieldClass)).append("\n");
    sb.append("    fieldAddOn: ").append(toIndentedString(fieldAddOn)).append("\n");
    sb.append("    fieldComponent: ").append(toIndentedString(fieldComponent)).append("\n");
    sb.append("    fieldInput: ").append(toIndentedString(fieldInput)).append("\n");
    sb.append("    placeHolder: ").append(toIndentedString(placeHolder)).append("\n");
    sb.append("    helpBlock: ").append(toIndentedString(helpBlock)).append("\n");
    sb.append("    defaultValue: ").append(toIndentedString(defaultValue)).append("\n");
    sb.append("    optionSource: ").append(toIndentedString(optionSource)).append("\n");
    sb.append("    optionList: ").append(toIndentedString(optionList)).append("\n");
    sb.append("    type: ").append(toIndentedString(type)).append("\n");
    sb.append("    advanced: ").append(toIndentedString(advanced)).append("\n");
    sb.append("    required: ").append(toIndentedString(required)).append("\n");
    sb.append("    editable: ").append(toIndentedString(editable)).append("\n");
    sb.append("    creatable: ").append(toIndentedString(creatable)).append("\n");
    sb.append("    config: ").append(toIndentedString(config)).append("\n");
    sb.append("    displayOrder: ").append(toIndentedString(displayOrder)).append("\n");
    sb.append("    wrapperClass: ").append(toIndentedString(wrapperClass)).append("\n");
    sb.append("    enabled: ").append(toIndentedString(enabled)).append("\n");
    sb.append("    noBlank: ").append(toIndentedString(noBlank)).append("\n");
    sb.append("    dependsOnCode: ").append(toIndentedString(dependsOnCode)).append("\n");
    sb.append("    contextualDefault: ").append(toIndentedString(contextualDefault)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  /**
   * Convert the given object to string with each line indented by 4 spaces
   * (except the first line).
   */
  private String toIndentedString(Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }

}
