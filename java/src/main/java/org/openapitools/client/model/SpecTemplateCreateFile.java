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
import org.openapitools.client.model.SpecTemplateCreateFileRepository;

/**
 * File, object specifying file type and content
 */
@ApiModel(description = "File, object specifying file type and content")
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class SpecTemplateCreateFile {
  public static final String SERIALIZED_NAME_SOURCE_TYPE = "sourceType";
  @SerializedName(SERIALIZED_NAME_SOURCE_TYPE)
  private String sourceType = "local";

  public static final String SERIALIZED_NAME_CONTENT = "content";
  @SerializedName(SERIALIZED_NAME_CONTENT)
  private String content;

  public static final String SERIALIZED_NAME_CONTENT_PATH = "contentPath";
  @SerializedName(SERIALIZED_NAME_CONTENT_PATH)
  private String contentPath;

  public static final String SERIALIZED_NAME_CONTENT_REF = "contentRef";
  @SerializedName(SERIALIZED_NAME_CONTENT_REF)
  private String contentRef;

  public static final String SERIALIZED_NAME_REPOSITORY = "repository";
  @SerializedName(SERIALIZED_NAME_REPOSITORY)
  private SpecTemplateCreateFileRepository repository;


  public SpecTemplateCreateFile sourceType(String sourceType) {
    
    this.sourceType = sourceType;
    return this;
  }

   /**
   * File Source i.e. local, repository, url.
   * @return sourceType
  **/
  @ApiModelProperty(required = true, value = "File Source i.e. local, repository, url.")

  public String getSourceType() {
    return sourceType;
  }


  public void setSourceType(String sourceType) {
    this.sourceType = sourceType;
  }


  public SpecTemplateCreateFile content(String content) {
    
    this.content = content;
    return this;
  }

   /**
   * File content, the template text. Only required when sourceType is &#x60;local&#x60;.
   * @return content
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "File content, the template text. Only required when sourceType is `local`.")

  public String getContent() {
    return content;
  }


  public void setContent(String content) {
    this.content = content;
  }


  public SpecTemplateCreateFile contentPath(String contentPath) {
    
    this.contentPath = contentPath;
    return this;
  }

   /**
   * Content Path, the repo file location or url. Required when sourceType is repository or url.
   * @return contentPath
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Content Path, the repo file location or url. Required when sourceType is repository or url.")

  public String getContentPath() {
    return contentPath;
  }


  public void setContentPath(String contentPath) {
    this.contentPath = contentPath;
  }


  public SpecTemplateCreateFile contentRef(String contentRef) {
    
    this.contentRef = contentRef;
    return this;
  }

   /**
   * Content Ref, the branch/tag. Only used when sourceType is repo.
   * @return contentRef
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Content Ref, the branch/tag. Only used when sourceType is repo.")

  public String getContentRef() {
    return contentRef;
  }


  public void setContentRef(String contentRef) {
    this.contentRef = contentRef;
  }


  public SpecTemplateCreateFile repository(SpecTemplateCreateFileRepository repository) {
    
    this.repository = repository;
    return this;
  }

   /**
   * Get repository
   * @return repository
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public SpecTemplateCreateFileRepository getRepository() {
    return repository;
  }


  public void setRepository(SpecTemplateCreateFileRepository repository) {
    this.repository = repository;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    SpecTemplateCreateFile specTemplateCreateFile = (SpecTemplateCreateFile) o;
    return Objects.equals(this.sourceType, specTemplateCreateFile.sourceType) &&
        Objects.equals(this.content, specTemplateCreateFile.content) &&
        Objects.equals(this.contentPath, specTemplateCreateFile.contentPath) &&
        Objects.equals(this.contentRef, specTemplateCreateFile.contentRef) &&
        Objects.equals(this.repository, specTemplateCreateFile.repository);
  }

  @Override
  public int hashCode() {
    return Objects.hash(sourceType, content, contentPath, contentRef, repository);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class SpecTemplateCreateFile {\n");
    sb.append("    sourceType: ").append(toIndentedString(sourceType)).append("\n");
    sb.append("    content: ").append(toIndentedString(content)).append("\n");
    sb.append("    contentPath: ").append(toIndentedString(contentPath)).append("\n");
    sb.append("    contentRef: ").append(toIndentedString(contentRef)).append("\n");
    sb.append("    repository: ").append(toIndentedString(repository)).append("\n");
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

