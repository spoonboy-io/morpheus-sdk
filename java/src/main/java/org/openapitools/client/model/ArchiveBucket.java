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
import org.openapitools.client.model.ArchiveBucketCreatedBy;
import org.openapitools.client.model.InlineResponse20040AppDeployInstance;
import org.threeten.bp.OffsetDateTime;

/**
 * ArchiveBucket
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class ArchiveBucket {
  public static final String SERIALIZED_NAME_ID = "id";
  @SerializedName(SERIALIZED_NAME_ID)
  private Long id;

  public static final String SERIALIZED_NAME_NAME = "name";
  @SerializedName(SERIALIZED_NAME_NAME)
  private String name;

  public static final String SERIALIZED_NAME_DESCRIPTION = "description";
  @SerializedName(SERIALIZED_NAME_DESCRIPTION)
  private String description;

  public static final String SERIALIZED_NAME_STORAGE_PROVIDER = "storageProvider";
  @SerializedName(SERIALIZED_NAME_STORAGE_PROVIDER)
  private InlineResponse20040AppDeployInstance storageProvider;

  public static final String SERIALIZED_NAME_OWNER = "owner";
  @SerializedName(SERIALIZED_NAME_OWNER)
  private InlineResponse20040AppDeployInstance owner;

  public static final String SERIALIZED_NAME_CREATED_BY = "createdBy";
  @SerializedName(SERIALIZED_NAME_CREATED_BY)
  private ArchiveBucketCreatedBy createdBy;

  public static final String SERIALIZED_NAME_IS_PUBLIC = "isPublic";
  @SerializedName(SERIALIZED_NAME_IS_PUBLIC)
  private Boolean isPublic;

  public static final String SERIALIZED_NAME_VISIBILITY = "visibility";
  @SerializedName(SERIALIZED_NAME_VISIBILITY)
  private String visibility;

  public static final String SERIALIZED_NAME_CODE = "code";
  @SerializedName(SERIALIZED_NAME_CODE)
  private String code;

  public static final String SERIALIZED_NAME_FILE_PATH = "filePath";
  @SerializedName(SERIALIZED_NAME_FILE_PATH)
  private String filePath;

  public static final String SERIALIZED_NAME_RAW_SIZE = "rawSize";
  @SerializedName(SERIALIZED_NAME_RAW_SIZE)
  private Long rawSize;

  public static final String SERIALIZED_NAME_FILE_COUNT = "fileCount";
  @SerializedName(SERIALIZED_NAME_FILE_COUNT)
  private Long fileCount;

  public static final String SERIALIZED_NAME_ACCOUNTS = "accounts";
  @SerializedName(SERIALIZED_NAME_ACCOUNTS)
  private List<Object> accounts = null;

  public static final String SERIALIZED_NAME_DATE_CREATED = "dateCreated";
  @SerializedName(SERIALIZED_NAME_DATE_CREATED)
  private OffsetDateTime dateCreated;

  public static final String SERIALIZED_NAME_LAST_UPDATED = "lastUpdated";
  @SerializedName(SERIALIZED_NAME_LAST_UPDATED)
  private OffsetDateTime lastUpdated;


  public ArchiveBucket id(Long id) {
    
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


  public ArchiveBucket name(String name) {
    
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


  public ArchiveBucket description(String description) {
    
    this.description = description;
    return this;
  }

   /**
   * Get description
   * @return description
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getDescription() {
    return description;
  }


  public void setDescription(String description) {
    this.description = description;
  }


  public ArchiveBucket storageProvider(InlineResponse20040AppDeployInstance storageProvider) {
    
    this.storageProvider = storageProvider;
    return this;
  }

   /**
   * Get storageProvider
   * @return storageProvider
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public InlineResponse20040AppDeployInstance getStorageProvider() {
    return storageProvider;
  }


  public void setStorageProvider(InlineResponse20040AppDeployInstance storageProvider) {
    this.storageProvider = storageProvider;
  }


  public ArchiveBucket owner(InlineResponse20040AppDeployInstance owner) {
    
    this.owner = owner;
    return this;
  }

   /**
   * Get owner
   * @return owner
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public InlineResponse20040AppDeployInstance getOwner() {
    return owner;
  }


  public void setOwner(InlineResponse20040AppDeployInstance owner) {
    this.owner = owner;
  }


  public ArchiveBucket createdBy(ArchiveBucketCreatedBy createdBy) {
    
    this.createdBy = createdBy;
    return this;
  }

   /**
   * Get createdBy
   * @return createdBy
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public ArchiveBucketCreatedBy getCreatedBy() {
    return createdBy;
  }


  public void setCreatedBy(ArchiveBucketCreatedBy createdBy) {
    this.createdBy = createdBy;
  }


  public ArchiveBucket isPublic(Boolean isPublic) {
    
    this.isPublic = isPublic;
    return this;
  }

   /**
   * Get isPublic
   * @return isPublic
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getIsPublic() {
    return isPublic;
  }


  public void setIsPublic(Boolean isPublic) {
    this.isPublic = isPublic;
  }


  public ArchiveBucket visibility(String visibility) {
    
    this.visibility = visibility;
    return this;
  }

   /**
   * Get visibility
   * @return visibility
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getVisibility() {
    return visibility;
  }


  public void setVisibility(String visibility) {
    this.visibility = visibility;
  }


  public ArchiveBucket code(String code) {
    
    this.code = code;
    return this;
  }

   /**
   * Get code
   * @return code
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getCode() {
    return code;
  }


  public void setCode(String code) {
    this.code = code;
  }


  public ArchiveBucket filePath(String filePath) {
    
    this.filePath = filePath;
    return this;
  }

   /**
   * Get filePath
   * @return filePath
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getFilePath() {
    return filePath;
  }


  public void setFilePath(String filePath) {
    this.filePath = filePath;
  }


  public ArchiveBucket rawSize(Long rawSize) {
    
    this.rawSize = rawSize;
    return this;
  }

   /**
   * Get rawSize
   * @return rawSize
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getRawSize() {
    return rawSize;
  }


  public void setRawSize(Long rawSize) {
    this.rawSize = rawSize;
  }


  public ArchiveBucket fileCount(Long fileCount) {
    
    this.fileCount = fileCount;
    return this;
  }

   /**
   * Get fileCount
   * @return fileCount
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getFileCount() {
    return fileCount;
  }


  public void setFileCount(Long fileCount) {
    this.fileCount = fileCount;
  }


  public ArchiveBucket accounts(List<Object> accounts) {
    
    this.accounts = accounts;
    return this;
  }

  public ArchiveBucket addAccountsItem(Object accountsItem) {
    if (this.accounts == null) {
      this.accounts = new ArrayList<Object>();
    }
    this.accounts.add(accountsItem);
    return this;
  }

   /**
   * Get accounts
   * @return accounts
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<Object> getAccounts() {
    return accounts;
  }


  public void setAccounts(List<Object> accounts) {
    this.accounts = accounts;
  }


  public ArchiveBucket dateCreated(OffsetDateTime dateCreated) {
    
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


  public ArchiveBucket lastUpdated(OffsetDateTime lastUpdated) {
    
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
    ArchiveBucket archiveBucket = (ArchiveBucket) o;
    return Objects.equals(this.id, archiveBucket.id) &&
        Objects.equals(this.name, archiveBucket.name) &&
        Objects.equals(this.description, archiveBucket.description) &&
        Objects.equals(this.storageProvider, archiveBucket.storageProvider) &&
        Objects.equals(this.owner, archiveBucket.owner) &&
        Objects.equals(this.createdBy, archiveBucket.createdBy) &&
        Objects.equals(this.isPublic, archiveBucket.isPublic) &&
        Objects.equals(this.visibility, archiveBucket.visibility) &&
        Objects.equals(this.code, archiveBucket.code) &&
        Objects.equals(this.filePath, archiveBucket.filePath) &&
        Objects.equals(this.rawSize, archiveBucket.rawSize) &&
        Objects.equals(this.fileCount, archiveBucket.fileCount) &&
        Objects.equals(this.accounts, archiveBucket.accounts) &&
        Objects.equals(this.dateCreated, archiveBucket.dateCreated) &&
        Objects.equals(this.lastUpdated, archiveBucket.lastUpdated);
  }

  @Override
  public int hashCode() {
    return Objects.hash(id, name, description, storageProvider, owner, createdBy, isPublic, visibility, code, filePath, rawSize, fileCount, accounts, dateCreated, lastUpdated);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ArchiveBucket {\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    description: ").append(toIndentedString(description)).append("\n");
    sb.append("    storageProvider: ").append(toIndentedString(storageProvider)).append("\n");
    sb.append("    owner: ").append(toIndentedString(owner)).append("\n");
    sb.append("    createdBy: ").append(toIndentedString(createdBy)).append("\n");
    sb.append("    isPublic: ").append(toIndentedString(isPublic)).append("\n");
    sb.append("    visibility: ").append(toIndentedString(visibility)).append("\n");
    sb.append("    code: ").append(toIndentedString(code)).append("\n");
    sb.append("    filePath: ").append(toIndentedString(filePath)).append("\n");
    sb.append("    rawSize: ").append(toIndentedString(rawSize)).append("\n");
    sb.append("    fileCount: ").append(toIndentedString(fileCount)).append("\n");
    sb.append("    accounts: ").append(toIndentedString(accounts)).append("\n");
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

