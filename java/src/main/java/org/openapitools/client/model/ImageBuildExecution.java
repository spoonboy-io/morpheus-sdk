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
import org.openapitools.client.model.ArchiveBucketFileCreatedBy;
import org.openapitools.client.model.InlineResponse20040AppDeployInstance;
import org.threeten.bp.OffsetDateTime;

/**
 * ImageBuildExecution
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class ImageBuildExecution {
  public static final String SERIALIZED_NAME_ID = "id";
  @SerializedName(SERIALIZED_NAME_ID)
  private Long id;

  public static final String SERIALIZED_NAME_IMAGE_BUILD = "imageBuild";
  @SerializedName(SERIALIZED_NAME_IMAGE_BUILD)
  private InlineResponse20040AppDeployInstance imageBuild;

  public static final String SERIALIZED_NAME_BUILD_NUMBER = "buildNumber";
  @SerializedName(SERIALIZED_NAME_BUILD_NUMBER)
  private Long buildNumber;

  public static final String SERIALIZED_NAME_START_DATE = "startDate";
  @SerializedName(SERIALIZED_NAME_START_DATE)
  private OffsetDateTime startDate;

  public static final String SERIALIZED_NAME_END_DATE = "endDate";
  @SerializedName(SERIALIZED_NAME_END_DATE)
  private OffsetDateTime endDate;

  public static final String SERIALIZED_NAME_STATUS_MESSAGE = "statusMessage";
  @SerializedName(SERIALIZED_NAME_STATUS_MESSAGE)
  private String statusMessage;

  public static final String SERIALIZED_NAME_STATUS_PERCENT = "statusPercent";
  @SerializedName(SERIALIZED_NAME_STATUS_PERCENT)
  private Long statusPercent;

  public static final String SERIALIZED_NAME_STATUS_ETA = "statusEta";
  @SerializedName(SERIALIZED_NAME_STATUS_ETA)
  private String statusEta;

  public static final String SERIALIZED_NAME_STATUS = "status";
  @SerializedName(SERIALIZED_NAME_STATUS)
  private String status;

  public static final String SERIALIZED_NAME_ERROR_MESSAGE = "errorMessage";
  @SerializedName(SERIALIZED_NAME_ERROR_MESSAGE)
  private String errorMessage;

  public static final String SERIALIZED_NAME_CREATED_BY = "createdBy";
  @SerializedName(SERIALIZED_NAME_CREATED_BY)
  private ArchiveBucketFileCreatedBy createdBy;

  public static final String SERIALIZED_NAME_TEMP_INSTANCE = "tempInstance";
  @SerializedName(SERIALIZED_NAME_TEMP_INSTANCE)
  private String tempInstance;

  public static final String SERIALIZED_NAME_VIRTUAL_IMAGES = "virtualImages";
  @SerializedName(SERIALIZED_NAME_VIRTUAL_IMAGES)
  private List<Object> virtualImages = null;


  public ImageBuildExecution id(Long id) {
    
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


  public ImageBuildExecution imageBuild(InlineResponse20040AppDeployInstance imageBuild) {
    
    this.imageBuild = imageBuild;
    return this;
  }

   /**
   * Get imageBuild
   * @return imageBuild
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public InlineResponse20040AppDeployInstance getImageBuild() {
    return imageBuild;
  }


  public void setImageBuild(InlineResponse20040AppDeployInstance imageBuild) {
    this.imageBuild = imageBuild;
  }


  public ImageBuildExecution buildNumber(Long buildNumber) {
    
    this.buildNumber = buildNumber;
    return this;
  }

   /**
   * Get buildNumber
   * @return buildNumber
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getBuildNumber() {
    return buildNumber;
  }


  public void setBuildNumber(Long buildNumber) {
    this.buildNumber = buildNumber;
  }


  public ImageBuildExecution startDate(OffsetDateTime startDate) {
    
    this.startDate = startDate;
    return this;
  }

   /**
   * Get startDate
   * @return startDate
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public OffsetDateTime getStartDate() {
    return startDate;
  }


  public void setStartDate(OffsetDateTime startDate) {
    this.startDate = startDate;
  }


  public ImageBuildExecution endDate(OffsetDateTime endDate) {
    
    this.endDate = endDate;
    return this;
  }

   /**
   * Get endDate
   * @return endDate
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public OffsetDateTime getEndDate() {
    return endDate;
  }


  public void setEndDate(OffsetDateTime endDate) {
    this.endDate = endDate;
  }


  public ImageBuildExecution statusMessage(String statusMessage) {
    
    this.statusMessage = statusMessage;
    return this;
  }

   /**
   * Get statusMessage
   * @return statusMessage
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getStatusMessage() {
    return statusMessage;
  }


  public void setStatusMessage(String statusMessage) {
    this.statusMessage = statusMessage;
  }


  public ImageBuildExecution statusPercent(Long statusPercent) {
    
    this.statusPercent = statusPercent;
    return this;
  }

   /**
   * Get statusPercent
   * @return statusPercent
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getStatusPercent() {
    return statusPercent;
  }


  public void setStatusPercent(Long statusPercent) {
    this.statusPercent = statusPercent;
  }


  public ImageBuildExecution statusEta(String statusEta) {
    
    this.statusEta = statusEta;
    return this;
  }

   /**
   * Get statusEta
   * @return statusEta
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getStatusEta() {
    return statusEta;
  }


  public void setStatusEta(String statusEta) {
    this.statusEta = statusEta;
  }


  public ImageBuildExecution status(String status) {
    
    this.status = status;
    return this;
  }

   /**
   * Get status
   * @return status
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getStatus() {
    return status;
  }


  public void setStatus(String status) {
    this.status = status;
  }


  public ImageBuildExecution errorMessage(String errorMessage) {
    
    this.errorMessage = errorMessage;
    return this;
  }

   /**
   * Get errorMessage
   * @return errorMessage
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getErrorMessage() {
    return errorMessage;
  }


  public void setErrorMessage(String errorMessage) {
    this.errorMessage = errorMessage;
  }


  public ImageBuildExecution createdBy(ArchiveBucketFileCreatedBy createdBy) {
    
    this.createdBy = createdBy;
    return this;
  }

   /**
   * Get createdBy
   * @return createdBy
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public ArchiveBucketFileCreatedBy getCreatedBy() {
    return createdBy;
  }


  public void setCreatedBy(ArchiveBucketFileCreatedBy createdBy) {
    this.createdBy = createdBy;
  }


  public ImageBuildExecution tempInstance(String tempInstance) {
    
    this.tempInstance = tempInstance;
    return this;
  }

   /**
   * Get tempInstance
   * @return tempInstance
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getTempInstance() {
    return tempInstance;
  }


  public void setTempInstance(String tempInstance) {
    this.tempInstance = tempInstance;
  }


  public ImageBuildExecution virtualImages(List<Object> virtualImages) {
    
    this.virtualImages = virtualImages;
    return this;
  }

  public ImageBuildExecution addVirtualImagesItem(Object virtualImagesItem) {
    if (this.virtualImages == null) {
      this.virtualImages = new ArrayList<Object>();
    }
    this.virtualImages.add(virtualImagesItem);
    return this;
  }

   /**
   * Get virtualImages
   * @return virtualImages
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<Object> getVirtualImages() {
    return virtualImages;
  }


  public void setVirtualImages(List<Object> virtualImages) {
    this.virtualImages = virtualImages;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    ImageBuildExecution imageBuildExecution = (ImageBuildExecution) o;
    return Objects.equals(this.id, imageBuildExecution.id) &&
        Objects.equals(this.imageBuild, imageBuildExecution.imageBuild) &&
        Objects.equals(this.buildNumber, imageBuildExecution.buildNumber) &&
        Objects.equals(this.startDate, imageBuildExecution.startDate) &&
        Objects.equals(this.endDate, imageBuildExecution.endDate) &&
        Objects.equals(this.statusMessage, imageBuildExecution.statusMessage) &&
        Objects.equals(this.statusPercent, imageBuildExecution.statusPercent) &&
        Objects.equals(this.statusEta, imageBuildExecution.statusEta) &&
        Objects.equals(this.status, imageBuildExecution.status) &&
        Objects.equals(this.errorMessage, imageBuildExecution.errorMessage) &&
        Objects.equals(this.createdBy, imageBuildExecution.createdBy) &&
        Objects.equals(this.tempInstance, imageBuildExecution.tempInstance) &&
        Objects.equals(this.virtualImages, imageBuildExecution.virtualImages);
  }

  @Override
  public int hashCode() {
    return Objects.hash(id, imageBuild, buildNumber, startDate, endDate, statusMessage, statusPercent, statusEta, status, errorMessage, createdBy, tempInstance, virtualImages);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ImageBuildExecution {\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    imageBuild: ").append(toIndentedString(imageBuild)).append("\n");
    sb.append("    buildNumber: ").append(toIndentedString(buildNumber)).append("\n");
    sb.append("    startDate: ").append(toIndentedString(startDate)).append("\n");
    sb.append("    endDate: ").append(toIndentedString(endDate)).append("\n");
    sb.append("    statusMessage: ").append(toIndentedString(statusMessage)).append("\n");
    sb.append("    statusPercent: ").append(toIndentedString(statusPercent)).append("\n");
    sb.append("    statusEta: ").append(toIndentedString(statusEta)).append("\n");
    sb.append("    status: ").append(toIndentedString(status)).append("\n");
    sb.append("    errorMessage: ").append(toIndentedString(errorMessage)).append("\n");
    sb.append("    createdBy: ").append(toIndentedString(createdBy)).append("\n");
    sb.append("    tempInstance: ").append(toIndentedString(tempInstance)).append("\n");
    sb.append("    virtualImages: ").append(toIndentedString(virtualImages)).append("\n");
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
