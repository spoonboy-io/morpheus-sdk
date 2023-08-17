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
import org.openapitools.client.model.BackupSettingsUpdate;

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
 * UpdateBackupSettingsRequest
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-17T13:37:08.123540Z[Etc/UTC]")
public class UpdateBackupSettingsRequest {
  public static final String SERIALIZED_NAME_BACKUP_SETTINGS = "backupSettings";
  @SerializedName(SERIALIZED_NAME_BACKUP_SETTINGS)
  private BackupSettingsUpdate backupSettings;

  public UpdateBackupSettingsRequest() {
  }

  public UpdateBackupSettingsRequest backupSettings(BackupSettingsUpdate backupSettings) {
    
    this.backupSettings = backupSettings;
    return this;
  }

   /**
   * Get backupSettings
   * @return backupSettings
  **/
  @javax.annotation.Nullable
  public BackupSettingsUpdate getBackupSettings() {
    return backupSettings;
  }


  public void setBackupSettings(BackupSettingsUpdate backupSettings) {
    this.backupSettings = backupSettings;
  }



  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    UpdateBackupSettingsRequest updateBackupSettingsRequest = (UpdateBackupSettingsRequest) o;
    return Objects.equals(this.backupSettings, updateBackupSettingsRequest.backupSettings);
  }

  @Override
  public int hashCode() {
    return Objects.hash(backupSettings);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class UpdateBackupSettingsRequest {\n");
    sb.append("    backupSettings: ").append(toIndentedString(backupSettings)).append("\n");
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
    openapiFields.add("backupSettings");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
  }

 /**
  * Validates the JSON Element and throws an exception if issues found
  *
  * @param jsonElement JSON Element
  * @throws IOException if the JSON Element is invalid with respect to UpdateBackupSettingsRequest
  */
  public static void validateJsonElement(JsonElement jsonElement) throws IOException {
      if (jsonElement == null) {
        if (!UpdateBackupSettingsRequest.openapiRequiredFields.isEmpty()) { // has required fields but JSON element is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in UpdateBackupSettingsRequest is not found in the empty JSON string", UpdateBackupSettingsRequest.openapiRequiredFields.toString()));
        }
      }

      Set<Entry<String, JsonElement>> entries = jsonElement.getAsJsonObject().entrySet();
      // check to see if the JSON string contains additional fields
      for (Entry<String, JsonElement> entry : entries) {
        if (!UpdateBackupSettingsRequest.openapiFields.contains(entry.getKey())) {
          throw new IllegalArgumentException(String.format("The field `%s` in the JSON string is not defined in the `UpdateBackupSettingsRequest` properties. JSON: %s", entry.getKey(), jsonElement.toString()));
        }
      }
        JsonObject jsonObj = jsonElement.getAsJsonObject();
      // validate the optional field `backupSettings`
      if (jsonObj.get("backupSettings") != null && !jsonObj.get("backupSettings").isJsonNull()) {
        BackupSettingsUpdate.validateJsonElement(jsonObj.get("backupSettings"));
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!UpdateBackupSettingsRequest.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'UpdateBackupSettingsRequest' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<UpdateBackupSettingsRequest> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(UpdateBackupSettingsRequest.class));

       return (TypeAdapter<T>) new TypeAdapter<UpdateBackupSettingsRequest>() {
           @Override
           public void write(JsonWriter out, UpdateBackupSettingsRequest value) throws IOException {
             JsonObject obj = thisAdapter.toJsonTree(value).getAsJsonObject();
             elementAdapter.write(out, obj);
           }

           @Override
           public UpdateBackupSettingsRequest read(JsonReader in) throws IOException {
             JsonElement jsonElement = elementAdapter.read(in);
             validateJsonElement(jsonElement);
             return thisAdapter.fromJsonTree(jsonElement);
           }

       }.nullSafe();
    }
  }

 /**
  * Create an instance of UpdateBackupSettingsRequest given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of UpdateBackupSettingsRequest
  * @throws IOException if the JSON string is invalid with respect to UpdateBackupSettingsRequest
  */
  public static UpdateBackupSettingsRequest fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, UpdateBackupSettingsRequest.class);
  }

 /**
  * Convert an instance of UpdateBackupSettingsRequest to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}
