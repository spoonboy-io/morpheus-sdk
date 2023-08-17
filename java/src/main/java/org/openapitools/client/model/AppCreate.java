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
import org.openapitools.client.model.AppCreateBlueprintId;
import org.openapitools.client.model.AppCreateDefaultCloud;
import org.openapitools.client.model.AppCreateGroup;
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
 * AppCreate
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-17T13:37:08.123540Z[Etc/UTC]")
public class AppCreate {
  public static final String SERIALIZED_NAME_TEMPLATE_ID = "templateId";
  @SerializedName(SERIALIZED_NAME_TEMPLATE_ID)
  private Long templateId;

  public static final String SERIALIZED_NAME_BLUEPRINT_ID = "blueprintId";
  @SerializedName(SERIALIZED_NAME_BLUEPRINT_ID)
  private AppCreateBlueprintId blueprintId;

  public static final String SERIALIZED_NAME_NAME = "name";
  @SerializedName(SERIALIZED_NAME_NAME)
  private String name;

  public static final String SERIALIZED_NAME_DESCRIPTION = "description";
  @SerializedName(SERIALIZED_NAME_DESCRIPTION)
  private String description;

  public static final String SERIALIZED_NAME_LABELS = "labels";
  @SerializedName(SERIALIZED_NAME_LABELS)
  private List<String> labels;

  public static final String SERIALIZED_NAME_GROUP = "group";
  @SerializedName(SERIALIZED_NAME_GROUP)
  private AppCreateGroup group;

  public static final String SERIALIZED_NAME_DEFAULT_CLOUD = "defaultCloud";
  @SerializedName(SERIALIZED_NAME_DEFAULT_CLOUD)
  private AppCreateDefaultCloud defaultCloud;

  public static final String SERIALIZED_NAME_ENVIRONMENT = "environment";
  @SerializedName(SERIALIZED_NAME_ENVIRONMENT)
  private String environment;

  public static final String SERIALIZED_NAME_TIERS = "tiers";
  @SerializedName(SERIALIZED_NAME_TIERS)
  private Object tiers;

  public AppCreate() {
  }

  public AppCreate templateId(Long templateId) {
    
    this.templateId = templateId;
    return this;
  }

   /**
   * Get templateId
   * @return templateId
  **/
  @javax.annotation.Nullable
  public Long getTemplateId() {
    return templateId;
  }


  public void setTemplateId(Long templateId) {
    this.templateId = templateId;
  }


  public AppCreate blueprintId(AppCreateBlueprintId blueprintId) {
    
    this.blueprintId = blueprintId;
    return this;
  }

   /**
   * Get blueprintId
   * @return blueprintId
  **/
  @javax.annotation.Nonnull
  public AppCreateBlueprintId getBlueprintId() {
    return blueprintId;
  }


  public void setBlueprintId(AppCreateBlueprintId blueprintId) {
    this.blueprintId = blueprintId;
  }


  public AppCreate name(String name) {
    
    this.name = name;
    return this;
  }

   /**
   * A unique name for the app
   * @return name
  **/
  @javax.annotation.Nonnull
  public String getName() {
    return name;
  }


  public void setName(String name) {
    this.name = name;
  }


  public AppCreate description(String description) {
    
    this.description = description;
    return this;
  }

   /**
   * Description
   * @return description
  **/
  @javax.annotation.Nullable
  public String getDescription() {
    return description;
  }


  public void setDescription(String description) {
    this.description = description;
  }


  public AppCreate labels(List<String> labels) {
    
    this.labels = labels;
    return this;
  }

  public AppCreate addLabelsItem(String labelsItem) {
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


  public AppCreate group(AppCreateGroup group) {
    
    this.group = group;
    return this;
  }

   /**
   * Get group
   * @return group
  **/
  @javax.annotation.Nullable
  public AppCreateGroup getGroup() {
    return group;
  }


  public void setGroup(AppCreateGroup group) {
    this.group = group;
  }


  public AppCreate defaultCloud(AppCreateDefaultCloud defaultCloud) {
    
    this.defaultCloud = defaultCloud;
    return this;
  }

   /**
   * Get defaultCloud
   * @return defaultCloud
  **/
  @javax.annotation.Nullable
  public AppCreateDefaultCloud getDefaultCloud() {
    return defaultCloud;
  }


  public void setDefaultCloud(AppCreateDefaultCloud defaultCloud) {
    this.defaultCloud = defaultCloud;
  }


  public AppCreate environment(String environment) {
    
    this.environment = environment;
    return this;
  }

   /**
   * Environment code (appContext)
   * @return environment
  **/
  @javax.annotation.Nullable
  public String getEnvironment() {
    return environment;
  }


  public void setEnvironment(String environment) {
    this.environment = environment;
  }


  public AppCreate tiers(Object tiers) {
    
    this.tiers = tiers;
    return this;
  }

   /**
   * Configuration of app elements
   * @return tiers
  **/
  @javax.annotation.Nullable
  public Object getTiers() {
    return tiers;
  }


  public void setTiers(Object tiers) {
    this.tiers = tiers;
  }



  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    AppCreate appCreate = (AppCreate) o;
    return Objects.equals(this.templateId, appCreate.templateId) &&
        Objects.equals(this.blueprintId, appCreate.blueprintId) &&
        Objects.equals(this.name, appCreate.name) &&
        Objects.equals(this.description, appCreate.description) &&
        Objects.equals(this.labels, appCreate.labels) &&
        Objects.equals(this.group, appCreate.group) &&
        Objects.equals(this.defaultCloud, appCreate.defaultCloud) &&
        Objects.equals(this.environment, appCreate.environment) &&
        Objects.equals(this.tiers, appCreate.tiers);
  }

  private static <T> boolean equalsNullable(JsonNullable<T> a, JsonNullable<T> b) {
    return a == b || (a != null && b != null && a.isPresent() && b.isPresent() && Objects.deepEquals(a.get(), b.get()));
  }

  @Override
  public int hashCode() {
    return Objects.hash(templateId, blueprintId, name, description, labels, group, defaultCloud, environment, tiers);
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
    sb.append("class AppCreate {\n");
    sb.append("    templateId: ").append(toIndentedString(templateId)).append("\n");
    sb.append("    blueprintId: ").append(toIndentedString(blueprintId)).append("\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    description: ").append(toIndentedString(description)).append("\n");
    sb.append("    labels: ").append(toIndentedString(labels)).append("\n");
    sb.append("    group: ").append(toIndentedString(group)).append("\n");
    sb.append("    defaultCloud: ").append(toIndentedString(defaultCloud)).append("\n");
    sb.append("    environment: ").append(toIndentedString(environment)).append("\n");
    sb.append("    tiers: ").append(toIndentedString(tiers)).append("\n");
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
    openapiFields.add("templateId");
    openapiFields.add("blueprintId");
    openapiFields.add("name");
    openapiFields.add("description");
    openapiFields.add("labels");
    openapiFields.add("group");
    openapiFields.add("defaultCloud");
    openapiFields.add("environment");
    openapiFields.add("tiers");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
    openapiRequiredFields.add("blueprintId");
    openapiRequiredFields.add("name");
  }

 /**
  * Validates the JSON Element and throws an exception if issues found
  *
  * @param jsonElement JSON Element
  * @throws IOException if the JSON Element is invalid with respect to AppCreate
  */
  public static void validateJsonElement(JsonElement jsonElement) throws IOException {
      if (jsonElement == null) {
        if (!AppCreate.openapiRequiredFields.isEmpty()) { // has required fields but JSON element is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in AppCreate is not found in the empty JSON string", AppCreate.openapiRequiredFields.toString()));
        }
      }

      Set<Entry<String, JsonElement>> entries = jsonElement.getAsJsonObject().entrySet();
      // check to see if the JSON string contains additional fields
      for (Entry<String, JsonElement> entry : entries) {
        if (!AppCreate.openapiFields.contains(entry.getKey())) {
          throw new IllegalArgumentException(String.format("The field `%s` in the JSON string is not defined in the `AppCreate` properties. JSON: %s", entry.getKey(), jsonElement.toString()));
        }
      }

      // check to make sure all required properties/fields are present in the JSON string
      for (String requiredField : AppCreate.openapiRequiredFields) {
        if (jsonElement.getAsJsonObject().get(requiredField) == null) {
          throw new IllegalArgumentException(String.format("The required field `%s` is not found in the JSON string: %s", requiredField, jsonElement.toString()));
        }
      }
        JsonObject jsonObj = jsonElement.getAsJsonObject();
      // validate the required field `blueprintId`
      AppCreateBlueprintId.validateJsonElement(jsonObj.get("blueprintId"));
      if (!jsonObj.get("name").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `name` to be a primitive type in the JSON string but got `%s`", jsonObj.get("name").toString()));
      }
      if ((jsonObj.get("description") != null && !jsonObj.get("description").isJsonNull()) && !jsonObj.get("description").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `description` to be a primitive type in the JSON string but got `%s`", jsonObj.get("description").toString()));
      }
      // ensure the optional json data is an array if present
      if (jsonObj.get("labels") != null && !jsonObj.get("labels").isJsonNull() && !jsonObj.get("labels").isJsonArray()) {
        throw new IllegalArgumentException(String.format("Expected the field `labels` to be an array in the JSON string but got `%s`", jsonObj.get("labels").toString()));
      }
      // validate the optional field `group`
      if (jsonObj.get("group") != null && !jsonObj.get("group").isJsonNull()) {
        AppCreateGroup.validateJsonElement(jsonObj.get("group"));
      }
      // validate the optional field `defaultCloud`
      if (jsonObj.get("defaultCloud") != null && !jsonObj.get("defaultCloud").isJsonNull()) {
        AppCreateDefaultCloud.validateJsonElement(jsonObj.get("defaultCloud"));
      }
      if ((jsonObj.get("environment") != null && !jsonObj.get("environment").isJsonNull()) && !jsonObj.get("environment").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `environment` to be a primitive type in the JSON string but got `%s`", jsonObj.get("environment").toString()));
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!AppCreate.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'AppCreate' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<AppCreate> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(AppCreate.class));

       return (TypeAdapter<T>) new TypeAdapter<AppCreate>() {
           @Override
           public void write(JsonWriter out, AppCreate value) throws IOException {
             JsonObject obj = thisAdapter.toJsonTree(value).getAsJsonObject();
             elementAdapter.write(out, obj);
           }

           @Override
           public AppCreate read(JsonReader in) throws IOException {
             JsonElement jsonElement = elementAdapter.read(in);
             validateJsonElement(jsonElement);
             return thisAdapter.fromJsonTree(jsonElement);
           }

       }.nullSafe();
    }
  }

 /**
  * Create an instance of AppCreate given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of AppCreate
  * @throws IOException if the JSON string is invalid with respect to AppCreate
  */
  public static AppCreate fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, AppCreate.class);
  }

 /**
  * Convert an instance of AppCreate to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}

