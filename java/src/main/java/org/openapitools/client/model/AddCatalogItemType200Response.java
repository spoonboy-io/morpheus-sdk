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
import org.openapitools.client.model.UpdateBlueprintPermissionsRequestResourcePermissionSitesInner;

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
 * AddCatalogItemType200Response
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-17T13:37:08.123540Z[Etc/UTC]")
public class AddCatalogItemType200Response {
  public static final String SERIALIZED_NAME_SUCCESS = "success";
  @SerializedName(SERIALIZED_NAME_SUCCESS)
  private Boolean success;

  public static final String SERIALIZED_NAME_CATALOG_ITEM_TYPE = "catalogItemType";
  @SerializedName(SERIALIZED_NAME_CATALOG_ITEM_TYPE)
  private UpdateBlueprintPermissionsRequestResourcePermissionSitesInner catalogItemType;

  public AddCatalogItemType200Response() {
  }

  public AddCatalogItemType200Response success(Boolean success) {
    
    this.success = success;
    return this;
  }

   /**
   * Get success
   * @return success
  **/
  @javax.annotation.Nullable
  public Boolean getSuccess() {
    return success;
  }


  public void setSuccess(Boolean success) {
    this.success = success;
  }


  public AddCatalogItemType200Response catalogItemType(UpdateBlueprintPermissionsRequestResourcePermissionSitesInner catalogItemType) {
    
    this.catalogItemType = catalogItemType;
    return this;
  }

   /**
   * Get catalogItemType
   * @return catalogItemType
  **/
  @javax.annotation.Nullable
  public UpdateBlueprintPermissionsRequestResourcePermissionSitesInner getCatalogItemType() {
    return catalogItemType;
  }


  public void setCatalogItemType(UpdateBlueprintPermissionsRequestResourcePermissionSitesInner catalogItemType) {
    this.catalogItemType = catalogItemType;
  }



  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    AddCatalogItemType200Response addCatalogItemType200Response = (AddCatalogItemType200Response) o;
    return Objects.equals(this.success, addCatalogItemType200Response.success) &&
        Objects.equals(this.catalogItemType, addCatalogItemType200Response.catalogItemType);
  }

  @Override
  public int hashCode() {
    return Objects.hash(success, catalogItemType);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class AddCatalogItemType200Response {\n");
    sb.append("    success: ").append(toIndentedString(success)).append("\n");
    sb.append("    catalogItemType: ").append(toIndentedString(catalogItemType)).append("\n");
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
    openapiFields.add("success");
    openapiFields.add("catalogItemType");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
  }

 /**
  * Validates the JSON Element and throws an exception if issues found
  *
  * @param jsonElement JSON Element
  * @throws IOException if the JSON Element is invalid with respect to AddCatalogItemType200Response
  */
  public static void validateJsonElement(JsonElement jsonElement) throws IOException {
      if (jsonElement == null) {
        if (!AddCatalogItemType200Response.openapiRequiredFields.isEmpty()) { // has required fields but JSON element is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in AddCatalogItemType200Response is not found in the empty JSON string", AddCatalogItemType200Response.openapiRequiredFields.toString()));
        }
      }

      Set<Entry<String, JsonElement>> entries = jsonElement.getAsJsonObject().entrySet();
      // check to see if the JSON string contains additional fields
      for (Entry<String, JsonElement> entry : entries) {
        if (!AddCatalogItemType200Response.openapiFields.contains(entry.getKey())) {
          throw new IllegalArgumentException(String.format("The field `%s` in the JSON string is not defined in the `AddCatalogItemType200Response` properties. JSON: %s", entry.getKey(), jsonElement.toString()));
        }
      }
        JsonObject jsonObj = jsonElement.getAsJsonObject();
      // validate the optional field `catalogItemType`
      if (jsonObj.get("catalogItemType") != null && !jsonObj.get("catalogItemType").isJsonNull()) {
        UpdateBlueprintPermissionsRequestResourcePermissionSitesInner.validateJsonElement(jsonObj.get("catalogItemType"));
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!AddCatalogItemType200Response.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'AddCatalogItemType200Response' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<AddCatalogItemType200Response> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(AddCatalogItemType200Response.class));

       return (TypeAdapter<T>) new TypeAdapter<AddCatalogItemType200Response>() {
           @Override
           public void write(JsonWriter out, AddCatalogItemType200Response value) throws IOException {
             JsonObject obj = thisAdapter.toJsonTree(value).getAsJsonObject();
             elementAdapter.write(out, obj);
           }

           @Override
           public AddCatalogItemType200Response read(JsonReader in) throws IOException {
             JsonElement jsonElement = elementAdapter.read(in);
             validateJsonElement(jsonElement);
             return thisAdapter.fromJsonTree(jsonElement);
           }

       }.nullSafe();
    }
  }

 /**
  * Create an instance of AddCatalogItemType200Response given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of AddCatalogItemType200Response
  * @throws IOException if the JSON string is invalid with respect to AddCatalogItemType200Response
  */
  public static AddCatalogItemType200Response fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, AddCatalogItemType200Response.class);
  }

 /**
  * Convert an instance of AddCatalogItemType200Response to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}

