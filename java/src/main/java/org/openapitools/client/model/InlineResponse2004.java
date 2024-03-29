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
import java.util.ArrayList;
import java.util.List;
import org.openapitools.client.model.ArchiveBucket;
import org.openapitools.client.model.ArchiveBucketFile;

/**
 * InlineResponse2004
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class InlineResponse2004 {
  public static final String SERIALIZED_NAME_ARCHIVE_BUCKET = "archiveBucket";
  @SerializedName(SERIALIZED_NAME_ARCHIVE_BUCKET)
  private ArchiveBucket archiveBucket;

  public static final String SERIALIZED_NAME_IS_OWNER = "isOwner";
  @SerializedName(SERIALIZED_NAME_IS_OWNER)
  private Boolean isOwner;

  public static final String SERIALIZED_NAME_PARENT_DIRECTORY = "parentDirectory";
  @SerializedName(SERIALIZED_NAME_PARENT_DIRECTORY)
  private String parentDirectory;

  public static final String SERIALIZED_NAME_ARCHIVE_FILES = "archiveFiles";
  @SerializedName(SERIALIZED_NAME_ARCHIVE_FILES)
  private List<ArchiveBucketFile> archiveFiles = null;

  public static final String SERIALIZED_NAME_ARCHIVE_FILE_COUNT = "archiveFileCount";
  @SerializedName(SERIALIZED_NAME_ARCHIVE_FILE_COUNT)
  private Long archiveFileCount;


  public InlineResponse2004 archiveBucket(ArchiveBucket archiveBucket) {
    
    this.archiveBucket = archiveBucket;
    return this;
  }

   /**
   * Get archiveBucket
   * @return archiveBucket
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public ArchiveBucket getArchiveBucket() {
    return archiveBucket;
  }


  public void setArchiveBucket(ArchiveBucket archiveBucket) {
    this.archiveBucket = archiveBucket;
  }


  public InlineResponse2004 isOwner(Boolean isOwner) {
    
    this.isOwner = isOwner;
    return this;
  }

   /**
   * Get isOwner
   * @return isOwner
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getIsOwner() {
    return isOwner;
  }


  public void setIsOwner(Boolean isOwner) {
    this.isOwner = isOwner;
  }


  public InlineResponse2004 parentDirectory(String parentDirectory) {
    
    this.parentDirectory = parentDirectory;
    return this;
  }

   /**
   * Get parentDirectory
   * @return parentDirectory
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getParentDirectory() {
    return parentDirectory;
  }


  public void setParentDirectory(String parentDirectory) {
    this.parentDirectory = parentDirectory;
  }


  public InlineResponse2004 archiveFiles(List<ArchiveBucketFile> archiveFiles) {
    
    this.archiveFiles = archiveFiles;
    return this;
  }

  public InlineResponse2004 addArchiveFilesItem(ArchiveBucketFile archiveFilesItem) {
    if (this.archiveFiles == null) {
      this.archiveFiles = new ArrayList<ArchiveBucketFile>();
    }
    this.archiveFiles.add(archiveFilesItem);
    return this;
  }

   /**
   * Get archiveFiles
   * @return archiveFiles
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<ArchiveBucketFile> getArchiveFiles() {
    return archiveFiles;
  }


  public void setArchiveFiles(List<ArchiveBucketFile> archiveFiles) {
    this.archiveFiles = archiveFiles;
  }


  public InlineResponse2004 archiveFileCount(Long archiveFileCount) {
    
    this.archiveFileCount = archiveFileCount;
    return this;
  }

   /**
   * Get archiveFileCount
   * @return archiveFileCount
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getArchiveFileCount() {
    return archiveFileCount;
  }


  public void setArchiveFileCount(Long archiveFileCount) {
    this.archiveFileCount = archiveFileCount;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    InlineResponse2004 inlineResponse2004 = (InlineResponse2004) o;
    return Objects.equals(this.archiveBucket, inlineResponse2004.archiveBucket) &&
        Objects.equals(this.isOwner, inlineResponse2004.isOwner) &&
        Objects.equals(this.parentDirectory, inlineResponse2004.parentDirectory) &&
        Objects.equals(this.archiveFiles, inlineResponse2004.archiveFiles) &&
        Objects.equals(this.archiveFileCount, inlineResponse2004.archiveFileCount);
  }

  @Override
  public int hashCode() {
    return Objects.hash(archiveBucket, isOwner, parentDirectory, archiveFiles, archiveFileCount);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class InlineResponse2004 {\n");
    sb.append("    archiveBucket: ").append(toIndentedString(archiveBucket)).append("\n");
    sb.append("    isOwner: ").append(toIndentedString(isOwner)).append("\n");
    sb.append("    parentDirectory: ").append(toIndentedString(parentDirectory)).append("\n");
    sb.append("    archiveFiles: ").append(toIndentedString(archiveFiles)).append("\n");
    sb.append("    archiveFileCount: ").append(toIndentedString(archiveFileCount)).append("\n");
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

