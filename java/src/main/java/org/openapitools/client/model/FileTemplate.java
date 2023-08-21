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
import org.openapitools.client.model.InlineResponse20082LoadBalancerInstanceSslCert;
import org.threeten.bp.OffsetDateTime;

/**
 * FileTemplate
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class FileTemplate {
  public static final String SERIALIZED_NAME_ID = "id";
  @SerializedName(SERIALIZED_NAME_ID)
  private Long id;

  public static final String SERIALIZED_NAME_CODE = "code";
  @SerializedName(SERIALIZED_NAME_CODE)
  private String code;

  public static final String SERIALIZED_NAME_ACCOUNT = "account";
  @SerializedName(SERIALIZED_NAME_ACCOUNT)
  private InlineResponse20082LoadBalancerInstanceSslCert account;

  public static final String SERIALIZED_NAME_NAME = "name";
  @SerializedName(SERIALIZED_NAME_NAME)
  private String name;

  public static final String SERIALIZED_NAME_LABELS = "labels";
  @SerializedName(SERIALIZED_NAME_LABELS)
  private List<String> labels = null;

  public static final String SERIALIZED_NAME_FILE_NAME = "fileName";
  @SerializedName(SERIALIZED_NAME_FILE_NAME)
  private String fileName;

  public static final String SERIALIZED_NAME_FILE_PATH = "filePath";
  @SerializedName(SERIALIZED_NAME_FILE_PATH)
  private String filePath;

  public static final String SERIALIZED_NAME_TEMPLATE_TYPE = "templateType";
  @SerializedName(SERIALIZED_NAME_TEMPLATE_TYPE)
  private String templateType;

  public static final String SERIALIZED_NAME_TEMPLATE_PHASE = "templatePhase";
  @SerializedName(SERIALIZED_NAME_TEMPLATE_PHASE)
  private String templatePhase;

  public static final String SERIALIZED_NAME_TEMPLATE = "template";
  @SerializedName(SERIALIZED_NAME_TEMPLATE)
  private String template;

  public static final String SERIALIZED_NAME_CATEGORY = "category";
  @SerializedName(SERIALIZED_NAME_CATEGORY)
  private String category;

  public static final String SERIALIZED_NAME_SETTING_CATEGORY = "settingCategory";
  @SerializedName(SERIALIZED_NAME_SETTING_CATEGORY)
  private String settingCategory;

  public static final String SERIALIZED_NAME_SETTING_NAME = "settingName";
  @SerializedName(SERIALIZED_NAME_SETTING_NAME)
  private String settingName;

  public static final String SERIALIZED_NAME_AUTO_RUN = "autoRun";
  @SerializedName(SERIALIZED_NAME_AUTO_RUN)
  private Boolean autoRun;

  public static final String SERIALIZED_NAME_RUN_ON_SCALE = "runOnScale";
  @SerializedName(SERIALIZED_NAME_RUN_ON_SCALE)
  private Boolean runOnScale;

  public static final String SERIALIZED_NAME_RUN_ON_DEPLOY = "runOnDeploy";
  @SerializedName(SERIALIZED_NAME_RUN_ON_DEPLOY)
  private Boolean runOnDeploy;

  public static final String SERIALIZED_NAME_FILE_OWNER = "fileOwner";
  @SerializedName(SERIALIZED_NAME_FILE_OWNER)
  private String fileOwner;

  public static final String SERIALIZED_NAME_FILE_GROUP = "fileGroup";
  @SerializedName(SERIALIZED_NAME_FILE_GROUP)
  private String fileGroup;

  public static final String SERIALIZED_NAME_PERMISSIONS = "permissions";
  @SerializedName(SERIALIZED_NAME_PERMISSIONS)
  private String permissions;

  public static final String SERIALIZED_NAME_DATE_CREATED = "dateCreated";
  @SerializedName(SERIALIZED_NAME_DATE_CREATED)
  private OffsetDateTime dateCreated;

  public static final String SERIALIZED_NAME_LAST_UPDATED = "lastUpdated";
  @SerializedName(SERIALIZED_NAME_LAST_UPDATED)
  private OffsetDateTime lastUpdated;


  public FileTemplate id(Long id) {
    
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


  public FileTemplate code(String code) {
    
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


  public FileTemplate account(InlineResponse20082LoadBalancerInstanceSslCert account) {
    
    this.account = account;
    return this;
  }

   /**
   * Get account
   * @return account
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public InlineResponse20082LoadBalancerInstanceSslCert getAccount() {
    return account;
  }


  public void setAccount(InlineResponse20082LoadBalancerInstanceSslCert account) {
    this.account = account;
  }


  public FileTemplate name(String name) {
    
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


  public FileTemplate labels(List<String> labels) {
    
    this.labels = labels;
    return this;
  }

  public FileTemplate addLabelsItem(String labelsItem) {
    if (this.labels == null) {
      this.labels = new ArrayList<String>();
    }
    this.labels.add(labelsItem);
    return this;
  }

   /**
   * Get labels
   * @return labels
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<String> getLabels() {
    return labels;
  }


  public void setLabels(List<String> labels) {
    this.labels = labels;
  }


  public FileTemplate fileName(String fileName) {
    
    this.fileName = fileName;
    return this;
  }

   /**
   * Get fileName
   * @return fileName
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getFileName() {
    return fileName;
  }


  public void setFileName(String fileName) {
    this.fileName = fileName;
  }


  public FileTemplate filePath(String filePath) {
    
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


  public FileTemplate templateType(String templateType) {
    
    this.templateType = templateType;
    return this;
  }

   /**
   * Get templateType
   * @return templateType
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getTemplateType() {
    return templateType;
  }


  public void setTemplateType(String templateType) {
    this.templateType = templateType;
  }


  public FileTemplate templatePhase(String templatePhase) {
    
    this.templatePhase = templatePhase;
    return this;
  }

   /**
   * Get templatePhase
   * @return templatePhase
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getTemplatePhase() {
    return templatePhase;
  }


  public void setTemplatePhase(String templatePhase) {
    this.templatePhase = templatePhase;
  }


  public FileTemplate template(String template) {
    
    this.template = template;
    return this;
  }

   /**
   * Get template
   * @return template
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getTemplate() {
    return template;
  }


  public void setTemplate(String template) {
    this.template = template;
  }


  public FileTemplate category(String category) {
    
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


  public FileTemplate settingCategory(String settingCategory) {
    
    this.settingCategory = settingCategory;
    return this;
  }

   /**
   * Get settingCategory
   * @return settingCategory
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getSettingCategory() {
    return settingCategory;
  }


  public void setSettingCategory(String settingCategory) {
    this.settingCategory = settingCategory;
  }


  public FileTemplate settingName(String settingName) {
    
    this.settingName = settingName;
    return this;
  }

   /**
   * Get settingName
   * @return settingName
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getSettingName() {
    return settingName;
  }


  public void setSettingName(String settingName) {
    this.settingName = settingName;
  }


  public FileTemplate autoRun(Boolean autoRun) {
    
    this.autoRun = autoRun;
    return this;
  }

   /**
   * Get autoRun
   * @return autoRun
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getAutoRun() {
    return autoRun;
  }


  public void setAutoRun(Boolean autoRun) {
    this.autoRun = autoRun;
  }


  public FileTemplate runOnScale(Boolean runOnScale) {
    
    this.runOnScale = runOnScale;
    return this;
  }

   /**
   * Get runOnScale
   * @return runOnScale
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getRunOnScale() {
    return runOnScale;
  }


  public void setRunOnScale(Boolean runOnScale) {
    this.runOnScale = runOnScale;
  }


  public FileTemplate runOnDeploy(Boolean runOnDeploy) {
    
    this.runOnDeploy = runOnDeploy;
    return this;
  }

   /**
   * Get runOnDeploy
   * @return runOnDeploy
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getRunOnDeploy() {
    return runOnDeploy;
  }


  public void setRunOnDeploy(Boolean runOnDeploy) {
    this.runOnDeploy = runOnDeploy;
  }


  public FileTemplate fileOwner(String fileOwner) {
    
    this.fileOwner = fileOwner;
    return this;
  }

   /**
   * Get fileOwner
   * @return fileOwner
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getFileOwner() {
    return fileOwner;
  }


  public void setFileOwner(String fileOwner) {
    this.fileOwner = fileOwner;
  }


  public FileTemplate fileGroup(String fileGroup) {
    
    this.fileGroup = fileGroup;
    return this;
  }

   /**
   * Get fileGroup
   * @return fileGroup
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getFileGroup() {
    return fileGroup;
  }


  public void setFileGroup(String fileGroup) {
    this.fileGroup = fileGroup;
  }


  public FileTemplate permissions(String permissions) {
    
    this.permissions = permissions;
    return this;
  }

   /**
   * Get permissions
   * @return permissions
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getPermissions() {
    return permissions;
  }


  public void setPermissions(String permissions) {
    this.permissions = permissions;
  }


  public FileTemplate dateCreated(OffsetDateTime dateCreated) {
    
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


  public FileTemplate lastUpdated(OffsetDateTime lastUpdated) {
    
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
    FileTemplate fileTemplate = (FileTemplate) o;
    return Objects.equals(this.id, fileTemplate.id) &&
        Objects.equals(this.code, fileTemplate.code) &&
        Objects.equals(this.account, fileTemplate.account) &&
        Objects.equals(this.name, fileTemplate.name) &&
        Objects.equals(this.labels, fileTemplate.labels) &&
        Objects.equals(this.fileName, fileTemplate.fileName) &&
        Objects.equals(this.filePath, fileTemplate.filePath) &&
        Objects.equals(this.templateType, fileTemplate.templateType) &&
        Objects.equals(this.templatePhase, fileTemplate.templatePhase) &&
        Objects.equals(this.template, fileTemplate.template) &&
        Objects.equals(this.category, fileTemplate.category) &&
        Objects.equals(this.settingCategory, fileTemplate.settingCategory) &&
        Objects.equals(this.settingName, fileTemplate.settingName) &&
        Objects.equals(this.autoRun, fileTemplate.autoRun) &&
        Objects.equals(this.runOnScale, fileTemplate.runOnScale) &&
        Objects.equals(this.runOnDeploy, fileTemplate.runOnDeploy) &&
        Objects.equals(this.fileOwner, fileTemplate.fileOwner) &&
        Objects.equals(this.fileGroup, fileTemplate.fileGroup) &&
        Objects.equals(this.permissions, fileTemplate.permissions) &&
        Objects.equals(this.dateCreated, fileTemplate.dateCreated) &&
        Objects.equals(this.lastUpdated, fileTemplate.lastUpdated);
  }

  @Override
  public int hashCode() {
    return Objects.hash(id, code, account, name, labels, fileName, filePath, templateType, templatePhase, template, category, settingCategory, settingName, autoRun, runOnScale, runOnDeploy, fileOwner, fileGroup, permissions, dateCreated, lastUpdated);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class FileTemplate {\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    code: ").append(toIndentedString(code)).append("\n");
    sb.append("    account: ").append(toIndentedString(account)).append("\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    labels: ").append(toIndentedString(labels)).append("\n");
    sb.append("    fileName: ").append(toIndentedString(fileName)).append("\n");
    sb.append("    filePath: ").append(toIndentedString(filePath)).append("\n");
    sb.append("    templateType: ").append(toIndentedString(templateType)).append("\n");
    sb.append("    templatePhase: ").append(toIndentedString(templatePhase)).append("\n");
    sb.append("    template: ").append(toIndentedString(template)).append("\n");
    sb.append("    category: ").append(toIndentedString(category)).append("\n");
    sb.append("    settingCategory: ").append(toIndentedString(settingCategory)).append("\n");
    sb.append("    settingName: ").append(toIndentedString(settingName)).append("\n");
    sb.append("    autoRun: ").append(toIndentedString(autoRun)).append("\n");
    sb.append("    runOnScale: ").append(toIndentedString(runOnScale)).append("\n");
    sb.append("    runOnDeploy: ").append(toIndentedString(runOnDeploy)).append("\n");
    sb.append("    fileOwner: ").append(toIndentedString(fileOwner)).append("\n");
    sb.append("    fileGroup: ").append(toIndentedString(fileGroup)).append("\n");
    sb.append("    permissions: ").append(toIndentedString(permissions)).append("\n");
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

