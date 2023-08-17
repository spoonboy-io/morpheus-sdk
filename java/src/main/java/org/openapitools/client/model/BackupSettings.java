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
import org.openapitools.client.model.ApplianceSettingsEnabledZoneTypesInner;
import org.openapitools.client.model.BackupSettingsDefaultSchedule;

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
 * BackupSettings
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-17T13:37:08.123540Z[Etc/UTC]")
public class BackupSettings {
  public static final String SERIALIZED_NAME_BACKUPS_ENABLED = "backupsEnabled";
  @SerializedName(SERIALIZED_NAME_BACKUPS_ENABLED)
  private Boolean backupsEnabled;

  public static final String SERIALIZED_NAME_CREATE_BACKUPS = "createBackups";
  @SerializedName(SERIALIZED_NAME_CREATE_BACKUPS)
  private Boolean createBackups;

  public static final String SERIALIZED_NAME_BACKUP_APPLIANCE = "backupAppliance";
  @SerializedName(SERIALIZED_NAME_BACKUP_APPLIANCE)
  private Boolean backupAppliance;

  public static final String SERIALIZED_NAME_DEFAULT_STORAGE_BUCKET = "defaultStorageBucket";
  @SerializedName(SERIALIZED_NAME_DEFAULT_STORAGE_BUCKET)
  private ApplianceSettingsEnabledZoneTypesInner defaultStorageBucket;

  public static final String SERIALIZED_NAME_DEFAULT_SCHEDULE = "defaultSchedule";
  @SerializedName(SERIALIZED_NAME_DEFAULT_SCHEDULE)
  private BackupSettingsDefaultSchedule defaultSchedule;

  public static final String SERIALIZED_NAME_RETENTION_COUNT = "retentionCount";
  @SerializedName(SERIALIZED_NAME_RETENTION_COUNT)
  private Long retentionCount;

  public BackupSettings() {
  }

  public BackupSettings backupsEnabled(Boolean backupsEnabled) {
    
    this.backupsEnabled = backupsEnabled;
    return this;
  }

   /**
   * Get backupsEnabled
   * @return backupsEnabled
  **/
  @javax.annotation.Nullable
  public Boolean getBackupsEnabled() {
    return backupsEnabled;
  }


  public void setBackupsEnabled(Boolean backupsEnabled) {
    this.backupsEnabled = backupsEnabled;
  }


  public BackupSettings createBackups(Boolean createBackups) {
    
    this.createBackups = createBackups;
    return this;
  }

   /**
   * Get createBackups
   * @return createBackups
  **/
  @javax.annotation.Nullable
  public Boolean getCreateBackups() {
    return createBackups;
  }


  public void setCreateBackups(Boolean createBackups) {
    this.createBackups = createBackups;
  }


  public BackupSettings backupAppliance(Boolean backupAppliance) {
    
    this.backupAppliance = backupAppliance;
    return this;
  }

   /**
   * Get backupAppliance
   * @return backupAppliance
  **/
  @javax.annotation.Nullable
  public Boolean getBackupAppliance() {
    return backupAppliance;
  }


  public void setBackupAppliance(Boolean backupAppliance) {
    this.backupAppliance = backupAppliance;
  }


  public BackupSettings defaultStorageBucket(ApplianceSettingsEnabledZoneTypesInner defaultStorageBucket) {
    
    this.defaultStorageBucket = defaultStorageBucket;
    return this;
  }

   /**
   * Get defaultStorageBucket
   * @return defaultStorageBucket
  **/
  @javax.annotation.Nullable
  public ApplianceSettingsEnabledZoneTypesInner getDefaultStorageBucket() {
    return defaultStorageBucket;
  }


  public void setDefaultStorageBucket(ApplianceSettingsEnabledZoneTypesInner defaultStorageBucket) {
    this.defaultStorageBucket = defaultStorageBucket;
  }


  public BackupSettings defaultSchedule(BackupSettingsDefaultSchedule defaultSchedule) {
    
    this.defaultSchedule = defaultSchedule;
    return this;
  }

   /**
   * Get defaultSchedule
   * @return defaultSchedule
  **/
  @javax.annotation.Nullable
  public BackupSettingsDefaultSchedule getDefaultSchedule() {
    return defaultSchedule;
  }


  public void setDefaultSchedule(BackupSettingsDefaultSchedule defaultSchedule) {
    this.defaultSchedule = defaultSchedule;
  }


  public BackupSettings retentionCount(Long retentionCount) {
    
    this.retentionCount = retentionCount;
    return this;
  }

   /**
   * Get retentionCount
   * @return retentionCount
  **/
  @javax.annotation.Nullable
  public Long getRetentionCount() {
    return retentionCount;
  }


  public void setRetentionCount(Long retentionCount) {
    this.retentionCount = retentionCount;
  }



  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    BackupSettings backupSettings = (BackupSettings) o;
    return Objects.equals(this.backupsEnabled, backupSettings.backupsEnabled) &&
        Objects.equals(this.createBackups, backupSettings.createBackups) &&
        Objects.equals(this.backupAppliance, backupSettings.backupAppliance) &&
        Objects.equals(this.defaultStorageBucket, backupSettings.defaultStorageBucket) &&
        Objects.equals(this.defaultSchedule, backupSettings.defaultSchedule) &&
        Objects.equals(this.retentionCount, backupSettings.retentionCount);
  }

  @Override
  public int hashCode() {
    return Objects.hash(backupsEnabled, createBackups, backupAppliance, defaultStorageBucket, defaultSchedule, retentionCount);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class BackupSettings {\n");
    sb.append("    backupsEnabled: ").append(toIndentedString(backupsEnabled)).append("\n");
    sb.append("    createBackups: ").append(toIndentedString(createBackups)).append("\n");
    sb.append("    backupAppliance: ").append(toIndentedString(backupAppliance)).append("\n");
    sb.append("    defaultStorageBucket: ").append(toIndentedString(defaultStorageBucket)).append("\n");
    sb.append("    defaultSchedule: ").append(toIndentedString(defaultSchedule)).append("\n");
    sb.append("    retentionCount: ").append(toIndentedString(retentionCount)).append("\n");
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
    openapiFields.add("backupsEnabled");
    openapiFields.add("createBackups");
    openapiFields.add("backupAppliance");
    openapiFields.add("defaultStorageBucket");
    openapiFields.add("defaultSchedule");
    openapiFields.add("retentionCount");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
  }

 /**
  * Validates the JSON Element and throws an exception if issues found
  *
  * @param jsonElement JSON Element
  * @throws IOException if the JSON Element is invalid with respect to BackupSettings
  */
  public static void validateJsonElement(JsonElement jsonElement) throws IOException {
      if (jsonElement == null) {
        if (!BackupSettings.openapiRequiredFields.isEmpty()) { // has required fields but JSON element is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in BackupSettings is not found in the empty JSON string", BackupSettings.openapiRequiredFields.toString()));
        }
      }

      Set<Entry<String, JsonElement>> entries = jsonElement.getAsJsonObject().entrySet();
      // check to see if the JSON string contains additional fields
      for (Entry<String, JsonElement> entry : entries) {
        if (!BackupSettings.openapiFields.contains(entry.getKey())) {
          throw new IllegalArgumentException(String.format("The field `%s` in the JSON string is not defined in the `BackupSettings` properties. JSON: %s", entry.getKey(), jsonElement.toString()));
        }
      }
        JsonObject jsonObj = jsonElement.getAsJsonObject();
      // validate the optional field `defaultStorageBucket`
      if (jsonObj.get("defaultStorageBucket") != null && !jsonObj.get("defaultStorageBucket").isJsonNull()) {
        ApplianceSettingsEnabledZoneTypesInner.validateJsonElement(jsonObj.get("defaultStorageBucket"));
      }
      // validate the optional field `defaultSchedule`
      if (jsonObj.get("defaultSchedule") != null && !jsonObj.get("defaultSchedule").isJsonNull()) {
        BackupSettingsDefaultSchedule.validateJsonElement(jsonObj.get("defaultSchedule"));
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!BackupSettings.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'BackupSettings' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<BackupSettings> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(BackupSettings.class));

       return (TypeAdapter<T>) new TypeAdapter<BackupSettings>() {
           @Override
           public void write(JsonWriter out, BackupSettings value) throws IOException {
             JsonObject obj = thisAdapter.toJsonTree(value).getAsJsonObject();
             elementAdapter.write(out, obj);
           }

           @Override
           public BackupSettings read(JsonReader in) throws IOException {
             JsonElement jsonElement = elementAdapter.read(in);
             validateJsonElement(jsonElement);
             return thisAdapter.fromJsonTree(jsonElement);
           }

       }.nullSafe();
    }
  }

 /**
  * Create an instance of BackupSettings given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of BackupSettings
  * @throws IOException if the JSON string is invalid with respect to BackupSettings
  */
  public static BackupSettings fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, BackupSettings.class);
  }

 /**
  * Convert an instance of BackupSettings to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}

