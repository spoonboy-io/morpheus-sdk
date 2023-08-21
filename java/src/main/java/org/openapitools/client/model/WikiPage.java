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
import org.openapitools.client.model.InlineResponse200107NetworkPoolCreatedBy;
import org.threeten.bp.OffsetDateTime;

/**
 * WikiPage
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class WikiPage {
  public static final String SERIALIZED_NAME_ID = "id";
  @SerializedName(SERIALIZED_NAME_ID)
  private Long id;

  public static final String SERIALIZED_NAME_NAME = "name";
  @SerializedName(SERIALIZED_NAME_NAME)
  private String name;

  public static final String SERIALIZED_NAME_URL_NAME = "urlName";
  @SerializedName(SERIALIZED_NAME_URL_NAME)
  private String urlName;

  public static final String SERIALIZED_NAME_CATEGORY = "category";
  @SerializedName(SERIALIZED_NAME_CATEGORY)
  private String category;

  public static final String SERIALIZED_NAME_REF_ID = "refId";
  @SerializedName(SERIALIZED_NAME_REF_ID)
  private String refId;

  public static final String SERIALIZED_NAME_REF_TYPE = "refType";
  @SerializedName(SERIALIZED_NAME_REF_TYPE)
  private String refType;

  public static final String SERIALIZED_NAME_FORMAT = "format";
  @SerializedName(SERIALIZED_NAME_FORMAT)
  private String format;

  public static final String SERIALIZED_NAME_CONTENT = "content";
  @SerializedName(SERIALIZED_NAME_CONTENT)
  private String content;

  public static final String SERIALIZED_NAME_CREATED_BY = "createdBy";
  @SerializedName(SERIALIZED_NAME_CREATED_BY)
  private InlineResponse200107NetworkPoolCreatedBy createdBy;

  public static final String SERIALIZED_NAME_UPDATED_BY = "updatedBy";
  @SerializedName(SERIALIZED_NAME_UPDATED_BY)
  private InlineResponse200107NetworkPoolCreatedBy updatedBy;

  public static final String SERIALIZED_NAME_DATE_CREATED = "dateCreated";
  @SerializedName(SERIALIZED_NAME_DATE_CREATED)
  private OffsetDateTime dateCreated;

  public static final String SERIALIZED_NAME_LAST_UPDATED = "lastUpdated";
  @SerializedName(SERIALIZED_NAME_LAST_UPDATED)
  private OffsetDateTime lastUpdated;


  public WikiPage id(Long id) {
    
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


  public WikiPage name(String name) {
    
    this.name = name;
    return this;
  }

   /**
   * Get name
   * @return name
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getName() {
    return name;
  }


  public void setName(String name) {
    this.name = name;
  }


  public WikiPage urlName(String urlName) {
    
    this.urlName = urlName;
    return this;
  }

   /**
   * Get urlName
   * @return urlName
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getUrlName() {
    return urlName;
  }


  public void setUrlName(String urlName) {
    this.urlName = urlName;
  }


  public WikiPage category(String category) {
    
    this.category = category;
    return this;
  }

   /**
   * Get category
   * @return category
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getCategory() {
    return category;
  }


  public void setCategory(String category) {
    this.category = category;
  }


  public WikiPage refId(String refId) {
    
    this.refId = refId;
    return this;
  }

   /**
   * Get refId
   * @return refId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getRefId() {
    return refId;
  }


  public void setRefId(String refId) {
    this.refId = refId;
  }


  public WikiPage refType(String refType) {
    
    this.refType = refType;
    return this;
  }

   /**
   * Get refType
   * @return refType
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getRefType() {
    return refType;
  }


  public void setRefType(String refType) {
    this.refType = refType;
  }


  public WikiPage format(String format) {
    
    this.format = format;
    return this;
  }

   /**
   * Get format
   * @return format
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getFormat() {
    return format;
  }


  public void setFormat(String format) {
    this.format = format;
  }


  public WikiPage content(String content) {
    
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


  public WikiPage createdBy(InlineResponse200107NetworkPoolCreatedBy createdBy) {
    
    this.createdBy = createdBy;
    return this;
  }

   /**
   * Get createdBy
   * @return createdBy
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public InlineResponse200107NetworkPoolCreatedBy getCreatedBy() {
    return createdBy;
  }


  public void setCreatedBy(InlineResponse200107NetworkPoolCreatedBy createdBy) {
    this.createdBy = createdBy;
  }


  public WikiPage updatedBy(InlineResponse200107NetworkPoolCreatedBy updatedBy) {
    
    this.updatedBy = updatedBy;
    return this;
  }

   /**
   * Get updatedBy
   * @return updatedBy
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public InlineResponse200107NetworkPoolCreatedBy getUpdatedBy() {
    return updatedBy;
  }


  public void setUpdatedBy(InlineResponse200107NetworkPoolCreatedBy updatedBy) {
    this.updatedBy = updatedBy;
  }


  public WikiPage dateCreated(OffsetDateTime dateCreated) {
    
    this.dateCreated = dateCreated;
    return this;
  }

   /**
   * Get dateCreated
   * @return dateCreated
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public OffsetDateTime getDateCreated() {
    return dateCreated;
  }


  public void setDateCreated(OffsetDateTime dateCreated) {
    this.dateCreated = dateCreated;
  }


  public WikiPage lastUpdated(OffsetDateTime lastUpdated) {
    
    this.lastUpdated = lastUpdated;
    return this;
  }

   /**
   * Get lastUpdated
   * @return lastUpdated
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

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
    WikiPage wikiPage = (WikiPage) o;
    return Objects.equals(this.id, wikiPage.id) &&
        Objects.equals(this.name, wikiPage.name) &&
        Objects.equals(this.urlName, wikiPage.urlName) &&
        Objects.equals(this.category, wikiPage.category) &&
        Objects.equals(this.refId, wikiPage.refId) &&
        Objects.equals(this.refType, wikiPage.refType) &&
        Objects.equals(this.format, wikiPage.format) &&
        Objects.equals(this.content, wikiPage.content) &&
        Objects.equals(this.createdBy, wikiPage.createdBy) &&
        Objects.equals(this.updatedBy, wikiPage.updatedBy) &&
        Objects.equals(this.dateCreated, wikiPage.dateCreated) &&
        Objects.equals(this.lastUpdated, wikiPage.lastUpdated);
  }

  @Override
  public int hashCode() {
    return Objects.hash(id, name, urlName, category, refId, refType, format, content, createdBy, updatedBy, dateCreated, lastUpdated);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class WikiPage {\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    urlName: ").append(toIndentedString(urlName)).append("\n");
    sb.append("    category: ").append(toIndentedString(category)).append("\n");
    sb.append("    refId: ").append(toIndentedString(refId)).append("\n");
    sb.append("    refType: ").append(toIndentedString(refType)).append("\n");
    sb.append("    format: ").append(toIndentedString(format)).append("\n");
    sb.append("    content: ").append(toIndentedString(content)).append("\n");
    sb.append("    createdBy: ").append(toIndentedString(createdBy)).append("\n");
    sb.append("    updatedBy: ").append(toIndentedString(updatedBy)).append("\n");
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

}

