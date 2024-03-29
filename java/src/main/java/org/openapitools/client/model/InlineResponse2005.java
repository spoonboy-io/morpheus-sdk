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
import org.openapitools.client.model.ArchiveBucketFile;

/**
 * InlineResponse2005
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class InlineResponse2005 {
  public static final String SERIALIZED_NAME_ARCHIVE_FILE = "archiveFile";
  @SerializedName(SERIALIZED_NAME_ARCHIVE_FILE)
  private ArchiveBucketFile archiveFile;

  public static final String SERIALIZED_NAME_IS_OWNER = "isOwner";
  @SerializedName(SERIALIZED_NAME_IS_OWNER)
  private Boolean isOwner;


  public InlineResponse2005 archiveFile(ArchiveBucketFile archiveFile) {
    
    this.archiveFile = archiveFile;
    return this;
  }

   /**
   * Get archiveFile
   * @return archiveFile
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public ArchiveBucketFile getArchiveFile() {
    return archiveFile;
  }


  public void setArchiveFile(ArchiveBucketFile archiveFile) {
    this.archiveFile = archiveFile;
  }


  public InlineResponse2005 isOwner(Boolean isOwner) {
    
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


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    InlineResponse2005 inlineResponse2005 = (InlineResponse2005) o;
    return Objects.equals(this.archiveFile, inlineResponse2005.archiveFile) &&
        Objects.equals(this.isOwner, inlineResponse2005.isOwner);
  }

  @Override
  public int hashCode() {
    return Objects.hash(archiveFile, isOwner);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class InlineResponse2005 {\n");
    sb.append("    archiveFile: ").append(toIndentedString(archiveFile)).append("\n");
    sb.append("    isOwner: ").append(toIndentedString(isOwner)).append("\n");
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

