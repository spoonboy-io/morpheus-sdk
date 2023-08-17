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
import org.openapitools.client.model.AppPrepareApplyDataGroup;
import org.openapitools.client.model.AppPrepareApplyDataTerraform;
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
 * AppPrepareApplyData
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-17T13:37:08.123540Z[Etc/UTC]")
public class AppPrepareApplyData {
  public static final String SERIALIZED_NAME_IMAGE = "image";
  @SerializedName(SERIALIZED_NAME_IMAGE)
  private String image;

  public static final String SERIALIZED_NAME_NAME = "name";
  @SerializedName(SERIALIZED_NAME_NAME)
  private String name;

  public static final String SERIALIZED_NAME_AUTO_VALIDATE = "autoValidate";
  @SerializedName(SERIALIZED_NAME_AUTO_VALIDATE)
  private Boolean autoValidate;

  public static final String SERIALIZED_NAME_TERRAFORM = "terraform";
  @SerializedName(SERIALIZED_NAME_TERRAFORM)
  private AppPrepareApplyDataTerraform terraform;

  public static final String SERIALIZED_NAME_TYPE = "type";
  @SerializedName(SERIALIZED_NAME_TYPE)
  private String type;

  public static final String SERIALIZED_NAME_CONFIG = "config";
  @SerializedName(SERIALIZED_NAME_CONFIG)
  private Object config;

  public static final String SERIALIZED_NAME_BLUEPRINT_NAME = "blueprintName";
  @SerializedName(SERIALIZED_NAME_BLUEPRINT_NAME)
  private String blueprintName;

  public static final String SERIALIZED_NAME_DESCRIPTION = "description";
  @SerializedName(SERIALIZED_NAME_DESCRIPTION)
  private String description;

  public static final String SERIALIZED_NAME_TEMPLATE_ID = "templateId";
  @SerializedName(SERIALIZED_NAME_TEMPLATE_ID)
  private Long templateId;

  public static final String SERIALIZED_NAME_BLUEPRINT_ID = "blueprintId";
  @SerializedName(SERIALIZED_NAME_BLUEPRINT_ID)
  private Long blueprintId;

  public static final String SERIALIZED_NAME_GROUP = "group";
  @SerializedName(SERIALIZED_NAME_GROUP)
  private AppPrepareApplyDataGroup group;

  public AppPrepareApplyData() {
  }

  public AppPrepareApplyData image(String image) {
    
    this.image = image;
    return this;
  }

   /**
   * Get image
   * @return image
  **/
  @javax.annotation.Nullable
  public String getImage() {
    return image;
  }


  public void setImage(String image) {
    this.image = image;
  }


  public AppPrepareApplyData name(String name) {
    
    this.name = name;
    return this;
  }

   /**
   * Get name
   * @return name
  **/
  @javax.annotation.Nullable
  public String getName() {
    return name;
  }


  public void setName(String name) {
    this.name = name;
  }


  public AppPrepareApplyData autoValidate(Boolean autoValidate) {
    
    this.autoValidate = autoValidate;
    return this;
  }

   /**
   * Get autoValidate
   * @return autoValidate
  **/
  @javax.annotation.Nullable
  public Boolean getAutoValidate() {
    return autoValidate;
  }


  public void setAutoValidate(Boolean autoValidate) {
    this.autoValidate = autoValidate;
  }


  public AppPrepareApplyData terraform(AppPrepareApplyDataTerraform terraform) {
    
    this.terraform = terraform;
    return this;
  }

   /**
   * Get terraform
   * @return terraform
  **/
  @javax.annotation.Nullable
  public AppPrepareApplyDataTerraform getTerraform() {
    return terraform;
  }


  public void setTerraform(AppPrepareApplyDataTerraform terraform) {
    this.terraform = terraform;
  }


  public AppPrepareApplyData type(String type) {
    
    this.type = type;
    return this;
  }

   /**
   * Get type
   * @return type
  **/
  @javax.annotation.Nullable
  public String getType() {
    return type;
  }


  public void setType(String type) {
    this.type = type;
  }


  public AppPrepareApplyData config(Object config) {
    
    this.config = config;
    return this;
  }

   /**
   * Get config
   * @return config
  **/
  @javax.annotation.Nullable
  public Object getConfig() {
    return config;
  }


  public void setConfig(Object config) {
    this.config = config;
  }


  public AppPrepareApplyData blueprintName(String blueprintName) {
    
    this.blueprintName = blueprintName;
    return this;
  }

   /**
   * Get blueprintName
   * @return blueprintName
  **/
  @javax.annotation.Nullable
  public String getBlueprintName() {
    return blueprintName;
  }


  public void setBlueprintName(String blueprintName) {
    this.blueprintName = blueprintName;
  }


  public AppPrepareApplyData description(String description) {
    
    this.description = description;
    return this;
  }

   /**
   * Get description
   * @return description
  **/
  @javax.annotation.Nullable
  public String getDescription() {
    return description;
  }


  public void setDescription(String description) {
    this.description = description;
  }


  public AppPrepareApplyData templateId(Long templateId) {
    
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


  public AppPrepareApplyData blueprintId(Long blueprintId) {
    
    this.blueprintId = blueprintId;
    return this;
  }

   /**
   * Get blueprintId
   * @return blueprintId
  **/
  @javax.annotation.Nullable
  public Long getBlueprintId() {
    return blueprintId;
  }


  public void setBlueprintId(Long blueprintId) {
    this.blueprintId = blueprintId;
  }


  public AppPrepareApplyData group(AppPrepareApplyDataGroup group) {
    
    this.group = group;
    return this;
  }

   /**
   * Get group
   * @return group
  **/
  @javax.annotation.Nullable
  public AppPrepareApplyDataGroup getGroup() {
    return group;
  }


  public void setGroup(AppPrepareApplyDataGroup group) {
    this.group = group;
  }



  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    AppPrepareApplyData appPrepareApplyData = (AppPrepareApplyData) o;
    return Objects.equals(this.image, appPrepareApplyData.image) &&
        Objects.equals(this.name, appPrepareApplyData.name) &&
        Objects.equals(this.autoValidate, appPrepareApplyData.autoValidate) &&
        Objects.equals(this.terraform, appPrepareApplyData.terraform) &&
        Objects.equals(this.type, appPrepareApplyData.type) &&
        Objects.equals(this.config, appPrepareApplyData.config) &&
        Objects.equals(this.blueprintName, appPrepareApplyData.blueprintName) &&
        Objects.equals(this.description, appPrepareApplyData.description) &&
        Objects.equals(this.templateId, appPrepareApplyData.templateId) &&
        Objects.equals(this.blueprintId, appPrepareApplyData.blueprintId) &&
        Objects.equals(this.group, appPrepareApplyData.group);
  }

  private static <T> boolean equalsNullable(JsonNullable<T> a, JsonNullable<T> b) {
    return a == b || (a != null && b != null && a.isPresent() && b.isPresent() && Objects.deepEquals(a.get(), b.get()));
  }

  @Override
  public int hashCode() {
    return Objects.hash(image, name, autoValidate, terraform, type, config, blueprintName, description, templateId, blueprintId, group);
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
    sb.append("class AppPrepareApplyData {\n");
    sb.append("    image: ").append(toIndentedString(image)).append("\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    autoValidate: ").append(toIndentedString(autoValidate)).append("\n");
    sb.append("    terraform: ").append(toIndentedString(terraform)).append("\n");
    sb.append("    type: ").append(toIndentedString(type)).append("\n");
    sb.append("    config: ").append(toIndentedString(config)).append("\n");
    sb.append("    blueprintName: ").append(toIndentedString(blueprintName)).append("\n");
    sb.append("    description: ").append(toIndentedString(description)).append("\n");
    sb.append("    templateId: ").append(toIndentedString(templateId)).append("\n");
    sb.append("    blueprintId: ").append(toIndentedString(blueprintId)).append("\n");
    sb.append("    group: ").append(toIndentedString(group)).append("\n");
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
    openapiFields.add("image");
    openapiFields.add("name");
    openapiFields.add("autoValidate");
    openapiFields.add("terraform");
    openapiFields.add("type");
    openapiFields.add("config");
    openapiFields.add("blueprintName");
    openapiFields.add("description");
    openapiFields.add("templateId");
    openapiFields.add("blueprintId");
    openapiFields.add("group");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
  }

 /**
  * Validates the JSON Element and throws an exception if issues found
  *
  * @param jsonElement JSON Element
  * @throws IOException if the JSON Element is invalid with respect to AppPrepareApplyData
  */
  public static void validateJsonElement(JsonElement jsonElement) throws IOException {
      if (jsonElement == null) {
        if (!AppPrepareApplyData.openapiRequiredFields.isEmpty()) { // has required fields but JSON element is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in AppPrepareApplyData is not found in the empty JSON string", AppPrepareApplyData.openapiRequiredFields.toString()));
        }
      }

      Set<Entry<String, JsonElement>> entries = jsonElement.getAsJsonObject().entrySet();
      // check to see if the JSON string contains additional fields
      for (Entry<String, JsonElement> entry : entries) {
        if (!AppPrepareApplyData.openapiFields.contains(entry.getKey())) {
          throw new IllegalArgumentException(String.format("The field `%s` in the JSON string is not defined in the `AppPrepareApplyData` properties. JSON: %s", entry.getKey(), jsonElement.toString()));
        }
      }
        JsonObject jsonObj = jsonElement.getAsJsonObject();
      if ((jsonObj.get("image") != null && !jsonObj.get("image").isJsonNull()) && !jsonObj.get("image").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `image` to be a primitive type in the JSON string but got `%s`", jsonObj.get("image").toString()));
      }
      if ((jsonObj.get("name") != null && !jsonObj.get("name").isJsonNull()) && !jsonObj.get("name").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `name` to be a primitive type in the JSON string but got `%s`", jsonObj.get("name").toString()));
      }
      // validate the optional field `terraform`
      if (jsonObj.get("terraform") != null && !jsonObj.get("terraform").isJsonNull()) {
        AppPrepareApplyDataTerraform.validateJsonElement(jsonObj.get("terraform"));
      }
      if ((jsonObj.get("type") != null && !jsonObj.get("type").isJsonNull()) && !jsonObj.get("type").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `type` to be a primitive type in the JSON string but got `%s`", jsonObj.get("type").toString()));
      }
      if ((jsonObj.get("blueprintName") != null && !jsonObj.get("blueprintName").isJsonNull()) && !jsonObj.get("blueprintName").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `blueprintName` to be a primitive type in the JSON string but got `%s`", jsonObj.get("blueprintName").toString()));
      }
      if ((jsonObj.get("description") != null && !jsonObj.get("description").isJsonNull()) && !jsonObj.get("description").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `description` to be a primitive type in the JSON string but got `%s`", jsonObj.get("description").toString()));
      }
      // validate the optional field `group`
      if (jsonObj.get("group") != null && !jsonObj.get("group").isJsonNull()) {
        AppPrepareApplyDataGroup.validateJsonElement(jsonObj.get("group"));
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!AppPrepareApplyData.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'AppPrepareApplyData' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<AppPrepareApplyData> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(AppPrepareApplyData.class));

       return (TypeAdapter<T>) new TypeAdapter<AppPrepareApplyData>() {
           @Override
           public void write(JsonWriter out, AppPrepareApplyData value) throws IOException {
             JsonObject obj = thisAdapter.toJsonTree(value).getAsJsonObject();
             elementAdapter.write(out, obj);
           }

           @Override
           public AppPrepareApplyData read(JsonReader in) throws IOException {
             JsonElement jsonElement = elementAdapter.read(in);
             validateJsonElement(jsonElement);
             return thisAdapter.fromJsonTree(jsonElement);
           }

       }.nullSafe();
    }
  }

 /**
  * Create an instance of AppPrepareApplyData given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of AppPrepareApplyData
  * @throws IOException if the JSON string is invalid with respect to AppPrepareApplyData
  */
  public static AppPrepareApplyData fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, AppPrepareApplyData.class);
  }

 /**
  * Convert an instance of AppPrepareApplyData to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}
