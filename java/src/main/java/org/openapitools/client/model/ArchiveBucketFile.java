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
import java.time.OffsetDateTime;
import org.openapitools.client.model.ArchiveBucketFileArchiveBucket;
import org.openapitools.client.model.ArchiveBucketFileCreatedBy;
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
 * ArchiveBucketFile
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-17T13:37:08.123540Z[Etc/UTC]")
public class ArchiveBucketFile {
  public static final String SERIALIZED_NAME_ID = "id";
  @SerializedName(SERIALIZED_NAME_ID)
  private Long id;

  public static final String SERIALIZED_NAME_NAME = "name";
  @SerializedName(SERIALIZED_NAME_NAME)
  private String name;

  public static final String SERIALIZED_NAME_FILE_PATH = "filePath";
  @SerializedName(SERIALIZED_NAME_FILE_PATH)
  private String filePath;

  public static final String SERIALIZED_NAME_ARCHIVE_BUCKET = "archiveBucket";
  @SerializedName(SERIALIZED_NAME_ARCHIVE_BUCKET)
  private ArchiveBucketFileArchiveBucket archiveBucket;

  public static final String SERIALIZED_NAME_CREATED_BY = "createdBy";
  @SerializedName(SERIALIZED_NAME_CREATED_BY)
  private ArchiveBucketFileCreatedBy createdBy;

  public static final String SERIALIZED_NAME_IS_DIRECTORY = "isDirectory";
  @SerializedName(SERIALIZED_NAME_IS_DIRECTORY)
  private Boolean isDirectory;

  public static final String SERIALIZED_NAME_STATUS = "status";
  @SerializedName(SERIALIZED_NAME_STATUS)
  private String status;

  public static final String SERIALIZED_NAME_RAW_SIZE = "rawSize";
  @SerializedName(SERIALIZED_NAME_RAW_SIZE)
  private Long rawSize;

  public static final String SERIALIZED_NAME_CONTENT_TYPE = "contentType";
  @SerializedName(SERIALIZED_NAME_CONTENT_TYPE)
  private String contentType;

  public static final String SERIALIZED_NAME_DOWNLOAD_COUNT = "downloadCount";
  @SerializedName(SERIALIZED_NAME_DOWNLOAD_COUNT)
  private Long downloadCount;

  public static final String SERIALIZED_NAME_DATE_CREATED = "dateCreated";
  @SerializedName(SERIALIZED_NAME_DATE_CREATED)
  private OffsetDateTime dateCreated;

  public static final String SERIALIZED_NAME_LAST_UPDATED = "lastUpdated";
  @SerializedName(SERIALIZED_NAME_LAST_UPDATED)
  private OffsetDateTime lastUpdated;

  public ArchiveBucketFile() {
  }

  public ArchiveBucketFile id(Long id) {
    
    this.id = id;
    return this;
  }

   /**
   * Get id
   * @return id
  **/
  @javax.annotation.Nullable
  public Long getId() {
    return id;
  }


  public void setId(Long id) {
    this.id = id;
  }


  public ArchiveBucketFile name(String name) {
    
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


  public ArchiveBucketFile filePath(String filePath) {
    
    this.filePath = filePath;
    return this;
  }

   /**
   * Get filePath
   * @return filePath
  **/
  @javax.annotation.Nullable
  public String getFilePath() {
    return filePath;
  }


  public void setFilePath(String filePath) {
    this.filePath = filePath;
  }


  public ArchiveBucketFile archiveBucket(ArchiveBucketFileArchiveBucket archiveBucket) {
    
    this.archiveBucket = archiveBucket;
    return this;
  }

   /**
   * Get archiveBucket
   * @return archiveBucket
  **/
  @javax.annotation.Nullable
  public ArchiveBucketFileArchiveBucket getArchiveBucket() {
    return archiveBucket;
  }


  public void setArchiveBucket(ArchiveBucketFileArchiveBucket archiveBucket) {
    this.archiveBucket = archiveBucket;
  }


  public ArchiveBucketFile createdBy(ArchiveBucketFileCreatedBy createdBy) {
    
    this.createdBy = createdBy;
    return this;
  }

   /**
   * Get createdBy
   * @return createdBy
  **/
  @javax.annotation.Nullable
  public ArchiveBucketFileCreatedBy getCreatedBy() {
    return createdBy;
  }


  public void setCreatedBy(ArchiveBucketFileCreatedBy createdBy) {
    this.createdBy = createdBy;
  }


  public ArchiveBucketFile isDirectory(Boolean isDirectory) {
    
    this.isDirectory = isDirectory;
    return this;
  }

   /**
   * Get isDirectory
   * @return isDirectory
  **/
  @javax.annotation.Nullable
  public Boolean getIsDirectory() {
    return isDirectory;
  }


  public void setIsDirectory(Boolean isDirectory) {
    this.isDirectory = isDirectory;
  }


  public ArchiveBucketFile status(String status) {
    
    this.status = status;
    return this;
  }

   /**
   * Get status
   * @return status
  **/
  @javax.annotation.Nullable
  public String getStatus() {
    return status;
  }


  public void setStatus(String status) {
    this.status = status;
  }


  public ArchiveBucketFile rawSize(Long rawSize) {
    
    this.rawSize = rawSize;
    return this;
  }

   /**
   * Get rawSize
   * @return rawSize
  **/
  @javax.annotation.Nullable
  public Long getRawSize() {
    return rawSize;
  }


  public void setRawSize(Long rawSize) {
    this.rawSize = rawSize;
  }


  public ArchiveBucketFile contentType(String contentType) {
    
    this.contentType = contentType;
    return this;
  }

   /**
   * Get contentType
   * @return contentType
  **/
  @javax.annotation.Nullable
  public String getContentType() {
    return contentType;
  }


  public void setContentType(String contentType) {
    this.contentType = contentType;
  }


  public ArchiveBucketFile downloadCount(Long downloadCount) {
    
    this.downloadCount = downloadCount;
    return this;
  }

   /**
   * Get downloadCount
   * @return downloadCount
  **/
  @javax.annotation.Nullable
  public Long getDownloadCount() {
    return downloadCount;
  }


  public void setDownloadCount(Long downloadCount) {
    this.downloadCount = downloadCount;
  }


  public ArchiveBucketFile dateCreated(OffsetDateTime dateCreated) {
    
    this.dateCreated = dateCreated;
    return this;
  }

   /**
   * Get dateCreated
   * @return dateCreated
  **/
  @javax.annotation.Nullable
  public OffsetDateTime getDateCreated() {
    return dateCreated;
  }


  public void setDateCreated(OffsetDateTime dateCreated) {
    this.dateCreated = dateCreated;
  }


  public ArchiveBucketFile lastUpdated(OffsetDateTime lastUpdated) {
    
    this.lastUpdated = lastUpdated;
    return this;
  }

   /**
   * Get lastUpdated
   * @return lastUpdated
  **/
  @javax.annotation.Nullable
  public OffsetDateTime getLastUpdated() {
    return lastUpdated;
  }


  public void setLastUpdated(OffsetDateTime lastUpdated) {
    this.lastUpdated = lastUpdated;
  }



  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    ArchiveBucketFile archiveBucketFile = (ArchiveBucketFile) o;
    return Objects.equals(this.id, archiveBucketFile.id) &&
        Objects.equals(this.name, archiveBucketFile.name) &&
        Objects.equals(this.filePath, archiveBucketFile.filePath) &&
        Objects.equals(this.archiveBucket, archiveBucketFile.archiveBucket) &&
        Objects.equals(this.createdBy, archiveBucketFile.createdBy) &&
        Objects.equals(this.isDirectory, archiveBucketFile.isDirectory) &&
        Objects.equals(this.status, archiveBucketFile.status) &&
        Objects.equals(this.rawSize, archiveBucketFile.rawSize) &&
        Objects.equals(this.contentType, archiveBucketFile.contentType) &&
        Objects.equals(this.downloadCount, archiveBucketFile.downloadCount) &&
        Objects.equals(this.dateCreated, archiveBucketFile.dateCreated) &&
        Objects.equals(this.lastUpdated, archiveBucketFile.lastUpdated);
  }

  private static <T> boolean equalsNullable(JsonNullable<T> a, JsonNullable<T> b) {
    return a == b || (a != null && b != null && a.isPresent() && b.isPresent() && Objects.deepEquals(a.get(), b.get()));
  }

  @Override
  public int hashCode() {
    return Objects.hash(id, name, filePath, archiveBucket, createdBy, isDirectory, status, rawSize, contentType, downloadCount, dateCreated, lastUpdated);
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
    sb.append("class ArchiveBucketFile {\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    filePath: ").append(toIndentedString(filePath)).append("\n");
    sb.append("    archiveBucket: ").append(toIndentedString(archiveBucket)).append("\n");
    sb.append("    createdBy: ").append(toIndentedString(createdBy)).append("\n");
    sb.append("    isDirectory: ").append(toIndentedString(isDirectory)).append("\n");
    sb.append("    status: ").append(toIndentedString(status)).append("\n");
    sb.append("    rawSize: ").append(toIndentedString(rawSize)).append("\n");
    sb.append("    contentType: ").append(toIndentedString(contentType)).append("\n");
    sb.append("    downloadCount: ").append(toIndentedString(downloadCount)).append("\n");
    sb.append("    dateCreated: ").append(toIndentedString(dateCreated)).append("\n");
    sb.append("    lastUpdated: ").append(toIndentedString(lastUpdated)).append("\n");
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
    openapiFields.add("id");
    openapiFields.add("name");
    openapiFields.add("filePath");
    openapiFields.add("archiveBucket");
    openapiFields.add("createdBy");
    openapiFields.add("isDirectory");
    openapiFields.add("status");
    openapiFields.add("rawSize");
    openapiFields.add("contentType");
    openapiFields.add("downloadCount");
    openapiFields.add("dateCreated");
    openapiFields.add("lastUpdated");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
  }

 /**
  * Validates the JSON Element and throws an exception if issues found
  *
  * @param jsonElement JSON Element
  * @throws IOException if the JSON Element is invalid with respect to ArchiveBucketFile
  */
  public static void validateJsonElement(JsonElement jsonElement) throws IOException {
      if (jsonElement == null) {
        if (!ArchiveBucketFile.openapiRequiredFields.isEmpty()) { // has required fields but JSON element is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in ArchiveBucketFile is not found in the empty JSON string", ArchiveBucketFile.openapiRequiredFields.toString()));
        }
      }

      Set<Entry<String, JsonElement>> entries = jsonElement.getAsJsonObject().entrySet();
      // check to see if the JSON string contains additional fields
      for (Entry<String, JsonElement> entry : entries) {
        if (!ArchiveBucketFile.openapiFields.contains(entry.getKey())) {
          throw new IllegalArgumentException(String.format("The field `%s` in the JSON string is not defined in the `ArchiveBucketFile` properties. JSON: %s", entry.getKey(), jsonElement.toString()));
        }
      }
        JsonObject jsonObj = jsonElement.getAsJsonObject();
      if ((jsonObj.get("name") != null && !jsonObj.get("name").isJsonNull()) && !jsonObj.get("name").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `name` to be a primitive type in the JSON string but got `%s`", jsonObj.get("name").toString()));
      }
      if ((jsonObj.get("filePath") != null && !jsonObj.get("filePath").isJsonNull()) && !jsonObj.get("filePath").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `filePath` to be a primitive type in the JSON string but got `%s`", jsonObj.get("filePath").toString()));
      }
      // validate the optional field `archiveBucket`
      if (jsonObj.get("archiveBucket") != null && !jsonObj.get("archiveBucket").isJsonNull()) {
        ArchiveBucketFileArchiveBucket.validateJsonElement(jsonObj.get("archiveBucket"));
      }
      // validate the optional field `createdBy`
      if (jsonObj.get("createdBy") != null && !jsonObj.get("createdBy").isJsonNull()) {
        ArchiveBucketFileCreatedBy.validateJsonElement(jsonObj.get("createdBy"));
      }
      if ((jsonObj.get("status") != null && !jsonObj.get("status").isJsonNull()) && !jsonObj.get("status").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `status` to be a primitive type in the JSON string but got `%s`", jsonObj.get("status").toString()));
      }
      if ((jsonObj.get("contentType") != null && !jsonObj.get("contentType").isJsonNull()) && !jsonObj.get("contentType").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `contentType` to be a primitive type in the JSON string but got `%s`", jsonObj.get("contentType").toString()));
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!ArchiveBucketFile.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'ArchiveBucketFile' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<ArchiveBucketFile> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(ArchiveBucketFile.class));

       return (TypeAdapter<T>) new TypeAdapter<ArchiveBucketFile>() {
           @Override
           public void write(JsonWriter out, ArchiveBucketFile value) throws IOException {
             JsonObject obj = thisAdapter.toJsonTree(value).getAsJsonObject();
             elementAdapter.write(out, obj);
           }

           @Override
           public ArchiveBucketFile read(JsonReader in) throws IOException {
             JsonElement jsonElement = elementAdapter.read(in);
             validateJsonElement(jsonElement);
             return thisAdapter.fromJsonTree(jsonElement);
           }

       }.nullSafe();
    }
  }

 /**
  * Create an instance of ArchiveBucketFile given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of ArchiveBucketFile
  * @throws IOException if the JSON string is invalid with respect to ArchiveBucketFile
  */
  public static ArchiveBucketFile fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, ArchiveBucketFile.class);
  }

 /**
  * Convert an instance of ArchiveBucketFile to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}
