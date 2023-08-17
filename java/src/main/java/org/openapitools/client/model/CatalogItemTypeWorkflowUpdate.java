/*
 * Morpheus API
 * Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.
 *
 * The version of the OpenAPI document: 6.1.1
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
import java.io.IOException;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import org.openapitools.client.model.UpdateBlueprintPermissionsRequestResourcePermissionSitesInner;
import org.openapitools.jackson.nullable.JsonNullable;

import com.google.gson.Gson;
import com.google.gson.GsonBuilder;
import com.google.gson.JsonArray;
import com.google.gson.JsonDeserializationContext;
import com.google.gson.JsonDeserializer;
import com.google.gson.JsonElement;
import com.google.gson.JsonObject;
import com.google.gson.JsonParseException;
import com.google.gson.TypeAdapterFactory;
import com.google.gson.reflect.TypeToken;
import com.google.gson.TypeAdapter;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import java.io.IOException;

import java.lang.reflect.Type;
import java.util.HashMap;
import java.util.HashSet;
import java.util.List;
import java.util.Map;
import java.util.Map.Entry;
import java.util.Set;

import org.openapitools.client.JSON;

/**
 * CatalogItemTypeWorkflowUpdate
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-17T13:37:08.123540Z[Etc/UTC]")
public class CatalogItemTypeWorkflowUpdate {
  public static final String SERIALIZED_NAME_NAME = "name";
  @SerializedName(SERIALIZED_NAME_NAME)
  private String name;

  public static final String SERIALIZED_NAME_CODE = "code";
  @SerializedName(SERIALIZED_NAME_CODE)
  private String code;

  public static final String SERIALIZED_NAME_CATEGORY = "category";
  @SerializedName(SERIALIZED_NAME_CATEGORY)
  private String category;

  public static final String SERIALIZED_NAME_DESCRIPTION = "description";
  @SerializedName(SERIALIZED_NAME_DESCRIPTION)
  private String description;

  public static final String SERIALIZED_NAME_LABELS = "labels";
  @SerializedName(SERIALIZED_NAME_LABELS)
  private List<String> labels;

  /**
   * Type, &#x60;instance&#x60;, &#x60;blueprint&#x60; or &#x60;workflow&#x60;. This determines whether an Instance or App will be provisioned. Instance types require a config and blueprint requires a blueprint and appSpec, while workflow types requires a workflow and context.
   */
  @JsonAdapter(TypeEnum.Adapter.class)
  public enum TypeEnum {
    WORKFLOW("workflow");

    private String value;

    TypeEnum(String value) {
      this.value = value;
    }

    public String getValue() {
      return value;
    }

    @Override
    public String toString() {
      return String.valueOf(value);
    }

    public static TypeEnum fromValue(String value) {
      for (TypeEnum b : TypeEnum.values()) {
        if (b.value.equals(value)) {
          return b;
        }
      }
      throw new IllegalArgumentException("Unexpected value '" + value + "'");
    }

    public static class Adapter extends TypeAdapter<TypeEnum> {
      @Override
      public void write(final JsonWriter jsonWriter, final TypeEnum enumeration) throws IOException {
        jsonWriter.value(enumeration.getValue());
      }

      @Override
      public TypeEnum read(final JsonReader jsonReader) throws IOException {
        String value =  jsonReader.nextString();
        return TypeEnum.fromValue(value);
      }
    }
  }

  public static final String SERIALIZED_NAME_TYPE = "type";
  @SerializedName(SERIALIZED_NAME_TYPE)
  private TypeEnum type;

  public static final String SERIALIZED_NAME_VISIBILITY = "visibility";
  @SerializedName(SERIALIZED_NAME_VISIBILITY)
  private String visibility = "private";

  public static final String SERIALIZED_NAME_LAYOUT_CODE = "layoutCode";
  @SerializedName(SERIALIZED_NAME_LAYOUT_CODE)
  private String layoutCode;

  public static final String SERIALIZED_NAME_ICON_PATH = "iconPath";
  @SerializedName(SERIALIZED_NAME_ICON_PATH)
  private String iconPath;

  public static final String SERIALIZED_NAME_ENABLED = "enabled";
  @SerializedName(SERIALIZED_NAME_ENABLED)
  private Boolean enabled = true;

  public static final String SERIALIZED_NAME_FEATURED = "featured";
  @SerializedName(SERIALIZED_NAME_FEATURED)
  private Boolean featured = false;

  public static final String SERIALIZED_NAME_ALLOW_QUANTITY = "allowQuantity";
  @SerializedName(SERIALIZED_NAME_ALLOW_QUANTITY)
  private Boolean allowQuantity = false;

  public static final String SERIALIZED_NAME_WORKFLOW = "workflow";
  @SerializedName(SERIALIZED_NAME_WORKFLOW)
  private UpdateBlueprintPermissionsRequestResourcePermissionSitesInner workflow;

  public static final String SERIALIZED_NAME_CONTEXT = "context";
  @SerializedName(SERIALIZED_NAME_CONTEXT)
  private String context;

  public static final String SERIALIZED_NAME_WORKFLOW_CONFIG = "workflowConfig";
  @SerializedName(SERIALIZED_NAME_WORKFLOW_CONFIG)
  private String workflowConfig;

  public CatalogItemTypeWorkflowUpdate() {
  }

  public CatalogItemTypeWorkflowUpdate name(String name) {
    
    this.name = name;
    return this;
  }

   /**
   * Catalog Item Type name
   * @return name
  **/
  @javax.annotation.Nullable
  public String getName() {
    return name;
  }


  public void setName(String name) {
    this.name = name;
  }


  public CatalogItemTypeWorkflowUpdate code(String code) {
    
    this.code = code;
    return this;
  }

   /**
   * Useful shortcode for provisioning naming schemes and export reference.
   * @return code
  **/
  @javax.annotation.Nullable
  public String getCode() {
    return code;
  }


  public void setCode(String code) {
    this.code = code;
  }


  public CatalogItemTypeWorkflowUpdate category(String category) {
    
    this.category = category;
    return this;
  }

   /**
   * Catalog Item Type category
   * @return category
  **/
  @javax.annotation.Nullable
  public String getCategory() {
    return category;
  }


  public void setCategory(String category) {
    this.category = category;
  }


  public CatalogItemTypeWorkflowUpdate description(String description) {
    
    this.description = description;
    return this;
  }

   /**
   * Catalog Item Type description
   * @return description
  **/
  @javax.annotation.Nullable
  public String getDescription() {
    return description;
  }


  public void setDescription(String description) {
    this.description = description;
  }


  public CatalogItemTypeWorkflowUpdate labels(List<String> labels) {
    
    this.labels = labels;
    return this;
  }

  public CatalogItemTypeWorkflowUpdate addLabelsItem(String labelsItem) {
    if (this.labels == null) {
      this.labels = new ArrayList<>();
    }
    this.labels.add(labelsItem);
    return this;
  }

   /**
   * Array of label strings, can be used for filtering.
   * @return labels
  **/
  @javax.annotation.Nullable
  public List<String> getLabels() {
    return labels;
  }


  public void setLabels(List<String> labels) {
    this.labels = labels;
  }


  public CatalogItemTypeWorkflowUpdate type(TypeEnum type) {
    
    this.type = type;
    return this;
  }

   /**
   * Type, &#x60;instance&#x60;, &#x60;blueprint&#x60; or &#x60;workflow&#x60;. This determines whether an Instance or App will be provisioned. Instance types require a config and blueprint requires a blueprint and appSpec, while workflow types requires a workflow and context.
   * @return type
  **/
  @javax.annotation.Nullable
  public TypeEnum getType() {
    return type;
  }


  public void setType(TypeEnum type) {
    this.type = type;
  }


  public CatalogItemTypeWorkflowUpdate visibility(String visibility) {
    
    this.visibility = visibility;
    return this;
  }

   /**
   * Visibility - Set to public to allow all tenants
   * @return visibility
  **/
  @javax.annotation.Nullable
  public String getVisibility() {
    return visibility;
  }


  public void setVisibility(String visibility) {
    this.visibility = visibility;
  }


  public CatalogItemTypeWorkflowUpdate layoutCode(String layoutCode) {
    
    this.layoutCode = layoutCode;
    return this;
  }

   /**
   * Identifier primarily used for Plugin Catalog Item Types
   * @return layoutCode
  **/
  @javax.annotation.Nullable
  public String getLayoutCode() {
    return layoutCode;
  }


  public void setLayoutCode(String layoutCode) {
    this.layoutCode = layoutCode;
  }


  public CatalogItemTypeWorkflowUpdate iconPath(String iconPath) {
    
    this.iconPath = iconPath;
    return this;
  }

   /**
   * Icon Path, relative location of an icon image, eg. /assets/containers-png/nginx.png.
   * @return iconPath
  **/
  @javax.annotation.Nullable
  public String getIconPath() {
    return iconPath;
  }


  public void setIconPath(String iconPath) {
    this.iconPath = iconPath;
  }


  public CatalogItemTypeWorkflowUpdate enabled(Boolean enabled) {
    
    this.enabled = enabled;
    return this;
  }

   /**
   * Can be used to enable / disable the catalog item type.
   * @return enabled
  **/
  @javax.annotation.Nullable
  public Boolean getEnabled() {
    return enabled;
  }


  public void setEnabled(Boolean enabled) {
    this.enabled = enabled;
  }


  public CatalogItemTypeWorkflowUpdate featured(Boolean featured) {
    
    this.featured = featured;
    return this;
  }

   /**
   * Can be used to feature the catalog item type.
   * @return featured
  **/
  @javax.annotation.Nullable
  public Boolean getFeatured() {
    return featured;
  }


  public void setFeatured(Boolean featured) {
    this.featured = featured;
  }


  public CatalogItemTypeWorkflowUpdate allowQuantity(Boolean allowQuantity) {
    
    this.allowQuantity = allowQuantity;
    return this;
  }

   /**
   * Can users order more than one of this item at a time.
   * @return allowQuantity
  **/
  @javax.annotation.Nullable
  public Boolean getAllowQuantity() {
    return allowQuantity;
  }


  public void setAllowQuantity(Boolean allowQuantity) {
    this.allowQuantity = allowQuantity;
  }


  public CatalogItemTypeWorkflowUpdate workflow(UpdateBlueprintPermissionsRequestResourcePermissionSitesInner workflow) {
    
    this.workflow = workflow;
    return this;
  }

   /**
   * Get workflow
   * @return workflow
  **/
  @javax.annotation.Nullable
  public UpdateBlueprintPermissionsRequestResourcePermissionSitesInner getWorkflow() {
    return workflow;
  }


  public void setWorkflow(UpdateBlueprintPermissionsRequestResourcePermissionSitesInner workflow) {
    this.workflow = workflow;
  }


  public CatalogItemTypeWorkflowUpdate context(String context) {
    
    this.context = context;
    return this;
  }

   /**
   * Get context
   * @return context
  **/
  @javax.annotation.Nullable
  public String getContext() {
    return context;
  }


  public void setContext(String context) {
    this.context = context;
  }


  public CatalogItemTypeWorkflowUpdate workflowConfig(String workflowConfig) {
    
    this.workflowConfig = workflowConfig;
    return this;
  }

   /**
   * Configuration object that contains settings for the workflow.
   * @return workflowConfig
  **/
  @javax.annotation.Nullable
  public String getWorkflowConfig() {
    return workflowConfig;
  }


  public void setWorkflowConfig(String workflowConfig) {
    this.workflowConfig = workflowConfig;
  }



  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    CatalogItemTypeWorkflowUpdate catalogItemTypeWorkflowUpdate = (CatalogItemTypeWorkflowUpdate) o;
    return Objects.equals(this.name, catalogItemTypeWorkflowUpdate.name) &&
        Objects.equals(this.code, catalogItemTypeWorkflowUpdate.code) &&
        Objects.equals(this.category, catalogItemTypeWorkflowUpdate.category) &&
        Objects.equals(this.description, catalogItemTypeWorkflowUpdate.description) &&
        Objects.equals(this.labels, catalogItemTypeWorkflowUpdate.labels) &&
        Objects.equals(this.type, catalogItemTypeWorkflowUpdate.type) &&
        Objects.equals(this.visibility, catalogItemTypeWorkflowUpdate.visibility) &&
        Objects.equals(this.layoutCode, catalogItemTypeWorkflowUpdate.layoutCode) &&
        Objects.equals(this.iconPath, catalogItemTypeWorkflowUpdate.iconPath) &&
        Objects.equals(this.enabled, catalogItemTypeWorkflowUpdate.enabled) &&
        Objects.equals(this.featured, catalogItemTypeWorkflowUpdate.featured) &&
        Objects.equals(this.allowQuantity, catalogItemTypeWorkflowUpdate.allowQuantity) &&
        Objects.equals(this.workflow, catalogItemTypeWorkflowUpdate.workflow) &&
        Objects.equals(this.context, catalogItemTypeWorkflowUpdate.context) &&
        Objects.equals(this.workflowConfig, catalogItemTypeWorkflowUpdate.workflowConfig);
  }

  private static <T> boolean equalsNullable(JsonNullable<T> a, JsonNullable<T> b) {
    return a == b || (a != null && b != null && a.isPresent() && b.isPresent() && Objects.deepEquals(a.get(), b.get()));
  }

  @Override
  public int hashCode() {
    return Objects.hash(name, code, category, description, labels, type, visibility, layoutCode, iconPath, enabled, featured, allowQuantity, workflow, context, workflowConfig);
  }

  private static <T> int hashCodeNullable(JsonNullable<T> a) {
    if (a == null) {
      return 1;
    }
    return a.isPresent() ? Arrays.deepHashCode(new Object[]{a.get()}) : 31;
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class CatalogItemTypeWorkflowUpdate {\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    code: ").append(toIndentedString(code)).append("\n");
    sb.append("    category: ").append(toIndentedString(category)).append("\n");
    sb.append("    description: ").append(toIndentedString(description)).append("\n");
    sb.append("    labels: ").append(toIndentedString(labels)).append("\n");
    sb.append("    type: ").append(toIndentedString(type)).append("\n");
    sb.append("    visibility: ").append(toIndentedString(visibility)).append("\n");
    sb.append("    layoutCode: ").append(toIndentedString(layoutCode)).append("\n");
    sb.append("    iconPath: ").append(toIndentedString(iconPath)).append("\n");
    sb.append("    enabled: ").append(toIndentedString(enabled)).append("\n");
    sb.append("    featured: ").append(toIndentedString(featured)).append("\n");
    sb.append("    allowQuantity: ").append(toIndentedString(allowQuantity)).append("\n");
    sb.append("    workflow: ").append(toIndentedString(workflow)).append("\n");
    sb.append("    context: ").append(toIndentedString(context)).append("\n");
    sb.append("    workflowConfig: ").append(toIndentedString(workflowConfig)).append("\n");
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


  public static HashSet<String> openapiFields;
  public static HashSet<String> openapiRequiredFields;

  static {
    // a set of all properties/fields (JSON key names)
    openapiFields = new HashSet<String>();
    openapiFields.add("name");
    openapiFields.add("code");
    openapiFields.add("category");
    openapiFields.add("description");
    openapiFields.add("labels");
    openapiFields.add("type");
    openapiFields.add("visibility");
    openapiFields.add("layoutCode");
    openapiFields.add("iconPath");
    openapiFields.add("enabled");
    openapiFields.add("featured");
    openapiFields.add("allowQuantity");
    openapiFields.add("workflow");
    openapiFields.add("context");
    openapiFields.add("workflowConfig");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
  }

 /**
  * Validates the JSON Element and throws an exception if issues found
  *
  * @param jsonElement JSON Element
  * @throws IOException if the JSON Element is invalid with respect to CatalogItemTypeWorkflowUpdate
  */
  public static void validateJsonElement(JsonElement jsonElement) throws IOException {
      if (jsonElement == null) {
        if (!CatalogItemTypeWorkflowUpdate.openapiRequiredFields.isEmpty()) { // has required fields but JSON element is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in CatalogItemTypeWorkflowUpdate is not found in the empty JSON string", CatalogItemTypeWorkflowUpdate.openapiRequiredFields.toString()));
        }
      }

      Set<Entry<String, JsonElement>> entries = jsonElement.getAsJsonObject().entrySet();
      // check to see if the JSON string contains additional fields
      for (Entry<String, JsonElement> entry : entries) {
        if (!CatalogItemTypeWorkflowUpdate.openapiFields.contains(entry.getKey())) {
          throw new IllegalArgumentException(String.format("The field `%s` in the JSON string is not defined in the `CatalogItemTypeWorkflowUpdate` properties. JSON: %s", entry.getKey(), jsonElement.toString()));
        }
      }
        JsonObject jsonObj = jsonElement.getAsJsonObject();
      if ((jsonObj.get("name") != null && !jsonObj.get("name").isJsonNull()) && !jsonObj.get("name").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `name` to be a primitive type in the JSON string but got `%s`", jsonObj.get("name").toString()));
      }
      if ((jsonObj.get("code") != null && !jsonObj.get("code").isJsonNull()) && !jsonObj.get("code").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `code` to be a primitive type in the JSON string but got `%s`", jsonObj.get("code").toString()));
      }
      if ((jsonObj.get("category") != null && !jsonObj.get("category").isJsonNull()) && !jsonObj.get("category").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `category` to be a primitive type in the JSON string but got `%s`", jsonObj.get("category").toString()));
      }
      if ((jsonObj.get("description") != null && !jsonObj.get("description").isJsonNull()) && !jsonObj.get("description").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `description` to be a primitive type in the JSON string but got `%s`", jsonObj.get("description").toString()));
      }
      // ensure the optional json data is an array if present
      if (jsonObj.get("labels") != null && !jsonObj.get("labels").isJsonNull() && !jsonObj.get("labels").isJsonArray()) {
        throw new IllegalArgumentException(String.format("Expected the field `labels` to be an array in the JSON string but got `%s`", jsonObj.get("labels").toString()));
      }
      if ((jsonObj.get("type") != null && !jsonObj.get("type").isJsonNull()) && !jsonObj.get("type").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `type` to be a primitive type in the JSON string but got `%s`", jsonObj.get("type").toString()));
      }
      if ((jsonObj.get("visibility") != null && !jsonObj.get("visibility").isJsonNull()) && !jsonObj.get("visibility").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `visibility` to be a primitive type in the JSON string but got `%s`", jsonObj.get("visibility").toString()));
      }
      if ((jsonObj.get("layoutCode") != null && !jsonObj.get("layoutCode").isJsonNull()) && !jsonObj.get("layoutCode").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `layoutCode` to be a primitive type in the JSON string but got `%s`", jsonObj.get("layoutCode").toString()));
      }
      if ((jsonObj.get("iconPath") != null && !jsonObj.get("iconPath").isJsonNull()) && !jsonObj.get("iconPath").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `iconPath` to be a primitive type in the JSON string but got `%s`", jsonObj.get("iconPath").toString()));
      }
      // validate the optional field `workflow`
      if (jsonObj.get("workflow") != null && !jsonObj.get("workflow").isJsonNull()) {
        UpdateBlueprintPermissionsRequestResourcePermissionSitesInner.validateJsonElement(jsonObj.get("workflow"));
      }
      if ((jsonObj.get("context") != null && !jsonObj.get("context").isJsonNull()) && !jsonObj.get("context").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `context` to be a primitive type in the JSON string but got `%s`", jsonObj.get("context").toString()));
      }
      if ((jsonObj.get("workflowConfig") != null && !jsonObj.get("workflowConfig").isJsonNull()) && !jsonObj.get("workflowConfig").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `workflowConfig` to be a primitive type in the JSON string but got `%s`", jsonObj.get("workflowConfig").toString()));
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!CatalogItemTypeWorkflowUpdate.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'CatalogItemTypeWorkflowUpdate' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<CatalogItemTypeWorkflowUpdate> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(CatalogItemTypeWorkflowUpdate.class));

       return (TypeAdapter<T>) new TypeAdapter<CatalogItemTypeWorkflowUpdate>() {
           @Override
           public void write(JsonWriter out, CatalogItemTypeWorkflowUpdate value) throws IOException {
             JsonObject obj = thisAdapter.toJsonTree(value).getAsJsonObject();
             elementAdapter.write(out, obj);
           }

           @Override
           public CatalogItemTypeWorkflowUpdate read(JsonReader in) throws IOException {
             JsonElement jsonElement = elementAdapter.read(in);
             validateJsonElement(jsonElement);
             return thisAdapter.fromJsonTree(jsonElement);
           }

       }.nullSafe();
    }
  }

 /**
  * Create an instance of CatalogItemTypeWorkflowUpdate given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of CatalogItemTypeWorkflowUpdate
  * @throws IOException if the JSON string is invalid with respect to CatalogItemTypeWorkflowUpdate
  */
  public static CatalogItemTypeWorkflowUpdate fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, CatalogItemTypeWorkflowUpdate.class);
  }

 /**
  * Convert an instance of CatalogItemTypeWorkflowUpdate to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}
