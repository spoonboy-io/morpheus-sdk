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
import org.openapitools.client.model.InlineResponse20082LoadBalancerInstanceSslCert;

/**
 * TaskAnsiblePlaybookConfigFile
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class TaskAnsiblePlaybookConfigFile {
  public static final String SERIALIZED_NAME_ID = "id";
  @SerializedName(SERIALIZED_NAME_ID)
  private Long id;

  public static final String SERIALIZED_NAME_SOURCE_TYPE = "sourceType";
  @SerializedName(SERIALIZED_NAME_SOURCE_TYPE)
  private String sourceType;

  public static final String SERIALIZED_NAME_CONTENT_REF = "contentRef";
  @SerializedName(SERIALIZED_NAME_CONTENT_REF)
  private String contentRef;

  public static final String SERIALIZED_NAME_CONTENT_PATH = "contentPath";
  @SerializedName(SERIALIZED_NAME_CONTENT_PATH)
  private String contentPath;

  public static final String SERIALIZED_NAME_REPOSITORY = "repository";
  @SerializedName(SERIALIZED_NAME_REPOSITORY)
  private InlineResponse20082LoadBalancerInstanceSslCert repository;

  public static final String SERIALIZED_NAME_CONTENT = "content";
  @SerializedName(SERIALIZED_NAME_CONTENT)
  private String content;


  public TaskAnsiblePlaybookConfigFile id(Long id) {
    
    this.id = id;
    return this;
  }

   /**
   * Get id
   * @return id
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getId() {
    return id;
  }


  public void setId(Long id) {
    this.id = id;
  }


  public TaskAnsiblePlaybookConfigFile sourceType(String sourceType) {
    
    this.sourceType = sourceType;
    return this;
  }

   /**
   * Get sourceType
   * @return sourceType
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getSourceType() {
    return sourceType;
  }


  public void setSourceType(String sourceType) {
    this.sourceType = sourceType;
  }


  public TaskAnsiblePlaybookConfigFile contentRef(String contentRef) {
    
    this.contentRef = contentRef;
    return this;
  }

   /**
   * Get contentRef
   * @return contentRef
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getContentRef() {
    return contentRef;
  }


  public void setContentRef(String contentRef) {
    this.contentRef = contentRef;
  }


  public TaskAnsiblePlaybookConfigFile contentPath(String contentPath) {
    
    this.contentPath = contentPath;
    return this;
  }

   /**
   * Get contentPath
   * @return contentPath
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getContentPath() {
    return contentPath;
  }


  public void setContentPath(String contentPath) {
    this.contentPath = contentPath;
  }


  public TaskAnsiblePlaybookConfigFile repository(InlineResponse20082LoadBalancerInstanceSslCert repository) {
    
    this.repository = repository;
    return this;
  }

   /**
   * Get repository
   * @return repository
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public InlineResponse20082LoadBalancerInstanceSslCert getRepository() {
    return repository;
  }


  public void setRepository(InlineResponse20082LoadBalancerInstanceSslCert repository) {
    this.repository = repository;
  }


  public TaskAnsiblePlaybookConfigFile content(String content) {
    
    this.content = content;
    return this;
  }

   /**
   * Get content
   * @return content
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getContent() {
    return content;
  }


  public void setContent(String content) {
    this.content = content;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    TaskAnsiblePlaybookConfigFile taskAnsiblePlaybookConfigFile = (TaskAnsiblePlaybookConfigFile) o;
    return Objects.equals(this.id, taskAnsiblePlaybookConfigFile.id) &&
        Objects.equals(this.sourceType, taskAnsiblePlaybookConfigFile.sourceType) &&
        Objects.equals(this.contentRef, taskAnsiblePlaybookConfigFile.contentRef) &&
        Objects.equals(this.contentPath, taskAnsiblePlaybookConfigFile.contentPath) &&
        Objects.equals(this.repository, taskAnsiblePlaybookConfigFile.repository) &&
        Objects.equals(this.content, taskAnsiblePlaybookConfigFile.content);
  }

  @Override
  public int hashCode() {
    return Objects.hash(id, sourceType, contentRef, contentPath, repository, content);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class TaskAnsiblePlaybookConfigFile {\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    sourceType: ").append(toIndentedString(sourceType)).append("\n");
    sb.append("    contentRef: ").append(toIndentedString(contentRef)).append("\n");
    sb.append("    contentPath: ").append(toIndentedString(contentPath)).append("\n");
    sb.append("    repository: ").append(toIndentedString(repository)).append("\n");
    sb.append("    content: ").append(toIndentedString(content)).append("\n");
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

