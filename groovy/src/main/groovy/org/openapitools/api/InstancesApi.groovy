package org.openapitools.api;

import org.openapitools.api.ApiUtils
import java.math.BigDecimal
import org.openapitools.model.DefaultError
import org.openapitools.model.InlineObject270
import org.openapitools.model.InlineObject91
import org.openapitools.model.InlineObject93
import org.openapitools.model.InlineObject94
import org.openapitools.model.InlineObject95
import org.openapitools.model.InlineObject96
import org.openapitools.model.InlineObject97
import org.openapitools.model.InlineObject98
import org.openapitools.model.InlineObject99
import org.openapitools.model.InlineResponse200168
import org.openapitools.model.InlineResponse20056
import org.openapitools.model.InlineResponse20057
import org.openapitools.model.InlineResponse20058
import org.openapitools.model.InlineResponse20059
import org.openapitools.model.InlineResponse20060
import org.openapitools.model.InstanceBackups
import org.openapitools.model.InstanceClone
import org.openapitools.model.InstanceCreate
import org.openapitools.model.InstanceResize
import org.openapitools.model.InstanceSnapshot
import org.openapitools.model.InstanceSnapshots
import org.openapitools.model.InstanceUpdate
import org.openapitools.model.InstanceWorkflow
import org.openapitools.model.InstancesCloneImage
import org.openapitools.model.Model200Success
import org.openapitools.model.NetworkInterfaceUpdate
import org.openapitools.model.Snapshot
import org.openapitools.model.SuccessMessage

class InstancesApi {
    String basePath = "https://CHANGEME"
    String versionPath = ""
    ApiUtils apiUtils = new ApiUtils();

    def addInstance ( InstanceCreate instanceCreate, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/instances"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType




        contentType = 'application/json';
        bodyParams = instanceCreate


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    Object.class )

    }

    def backupInstance ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/instances/${id}/backup"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Model200Success.class )

    }

    def backupsInstance ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/instances/${id}/backups"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    InstanceBackups.class )

    }

    def cancelExpirationInstance ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/instances/${id}/cancel-expiration"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Model200Success.class )

    }

    def cancelRemovalInstance ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/instances/${id}/cancel-removal"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Model200Success.class )

    }

    def cancelShutdownInstance ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/instances/${id}/cancel-shutdown"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Model200Success.class )

    }

    def cloneImageInstance ( Long id, InstancesCloneImage instancesCloneImage, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/instances/${id}/clone-image"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }



        contentType = 'application/json';
        bodyParams = instancesCloneImage


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Model200Success.class )

    }

    def cloneInstance ( Long id, InstanceClone instanceClone, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/instances/${id}/clone"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }



        contentType = 'application/json';
        bodyParams = instanceClone


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Model200Success.class )

    }

    def createInstanceSchedule ( Long id, InlineObject96 inlineObject96, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/instances/${id}/schedules"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }



        contentType = 'application/json';
        bodyParams = inlineObject96


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    Object.class )

    }

    def deleteAllSnapshotsInstance ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/instances/${id}/delete-all-snapshots"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "DELETE", "",
                    Model200Success.class )

    }

    def deleteAllSnapshotsInstanceContainer ( Long id, BigDecimal containerId, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/instances/${id}/delete-container-snapshots/${containerId}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }
        // verify required params are set
        if (containerId == null) {
            throw new RuntimeException("missing required params containerId")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "DELETE", "",
                    Model200Success.class )

    }

    def deleteInstance ( Long id, String preserveVolumes, String keepBackups, String releaseFloatingIps, String releaseEIPs, String force, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/instances/${id}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }

        if (preserveVolumes != null) {
            queryParams.put("preserveVolumes", preserveVolumes)
        }
        if (keepBackups != null) {
            queryParams.put("keepBackups", keepBackups)
        }
        if (releaseFloatingIps != null) {
            queryParams.put("releaseFloatingIps", releaseFloatingIps)
        }
        if (releaseEIPs != null) {
            queryParams.put("releaseEIPs", releaseEIPs)
        }
        if (force != null) {
            queryParams.put("force", force)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "DELETE", "",
                    Model200Success.class )

    }

    def deleteInstanceSchedule ( Long id, Long scheduleId, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/instances/${id}/schedules/${scheduleId}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }
        // verify required params are set
        if (scheduleId == null) {
            throw new RuntimeException("missing required params scheduleId")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "DELETE", "",
                    Model200Success.class )

    }

    def deleteSnapshotInstance ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/snapshots/${id}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "DELETE", "",
                    Model200Success.class )

    }

    def ejectInstance ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/instances/${id}/eject"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Model200Success.class )

    }

    def extendExpirationInstance ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/instances/${id}/extend-expiration"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Model200Success.class )

    }

    def extendShutdownInstance ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/instances/${id}/extend-shutdown"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Model200Success.class )

    }

    def getEnvVariables ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/instances/${id}/envs"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    InlineResponse20057.class )

    }

    def getInstance ( Long id, Boolean details, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/instances/${id}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }

        if (details != null) {
            queryParams.put("details", details)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    InlineResponse20056.class )

    }

    def getInstanceContainers ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/instances/${id}/containers"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def getInstanceHistory ( Long id, String phrase, Long containerId, Long serverId, Long zoneId, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/instances/${id}/history"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }

        if (phrase != null) {
            queryParams.put("phrase", phrase)
        }
        if (containerId != null) {
            queryParams.put("containerId", containerId)
        }
        if (serverId != null) {
            queryParams.put("serverId", serverId)
        }
        if (zoneId != null) {
            queryParams.put("zoneId", zoneId)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def getInstanceSchedule ( Long id, Long scheduleId, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/instances/${id}/schedules/${scheduleId}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }
        // verify required params are set
        if (scheduleId == null) {
            throw new RuntimeException("missing required params scheduleId")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    InlineResponse20059.class )

    }

    def getInstanceSchedules ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/instances/${id}/schedules"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def getInstanceThreshold ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/instances/${id}/threshold"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    InlineResponse20058.class )

    }

    def getInstanceTypeProvisioning ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/instance-types/${id}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def getPrepareApplyInstance ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/instances/${id}/prepare-apply"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def getSnapshotInstance ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/snapshots/${id}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Snapshot.class )

    }

    def getStateInstance ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/instances/${id}/state"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def getValidateApplyInstance ( Long id, InlineObject98 inlineObject98, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/instances/${id}/validate-apply"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }



        contentType = 'application/json';
        bodyParams = inlineObject98


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    Object.class )

    }

    def getWikiInstance ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/instances/${id}/wiki"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    InlineResponse200168.class )

    }

    def importSnapshotInstance ( Long id, InlineObject93 inlineObject93, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/instances/${id}/import-snapshot"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }



        contentType = 'application/json';
        bodyParams = inlineObject93


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Model200Success.class )

    }

    def linkedCloneSnapshotInstance ( Long id, BigDecimal snapshotId, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/instances/${id}/linked-clone/${snapshotId}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }
        // verify required params are set
        if (snapshotId == null) {
            throw new RuntimeException("missing required params snapshotId")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Model200Success.class )

    }

    def listInstanceServicePlans ( Long zoneId, Long layoutId, Long siteId, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/instances/service-plans"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType


        if (zoneId != null) {
            queryParams.put("zoneId", zoneId)
        }
        if (layoutId != null) {
            queryParams.put("layoutId", layoutId)
        }
        if (siteId != null) {
            queryParams.put("siteId", siteId)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    InlineResponse20060.class )

    }

    def listInstanceTypesProvisioning ( Long max, Long offset, String sort, String direction, String phrase, String name, String code, Boolean featured, Boolean details, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/instance-types"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType


        if (max != null) {
            queryParams.put("max", max)
        }
        if (offset != null) {
            queryParams.put("offset", offset)
        }
        if (sort != null) {
            queryParams.put("sort", sort)
        }
        if (direction != null) {
            queryParams.put("direction", direction)
        }
        if (phrase != null) {
            queryParams.put("phrase", phrase)
        }
        if (name != null) {
            queryParams.put("name", name)
        }
        if (code != null) {
            queryParams.put("code", code)
        }
        if (featured != null) {
            queryParams.put("featured", featured)
        }
        if (details != null) {
            queryParams.put("details", details)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def listInstances ( Long max, Long offset, String name, String phrase, String instanceType, Date lastUpdated, Long createdBy, Boolean agentInstalled, String status, String environment, Boolean showDeleted, Boolean deleted, String expireDate, String expireDateMin, String expireDays, String expireDaysMin, String shutdownDate, String shutdownDateMin, String shutdownDays, String shutdownDaysMin, String labels, String allLabels, String tags, String metadata, Boolean details, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/instances"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType


        if (max != null) {
            queryParams.put("max", max)
        }
        if (offset != null) {
            queryParams.put("offset", offset)
        }
        if (name != null) {
            queryParams.put("name", name)
        }
        if (phrase != null) {
            queryParams.put("phrase", phrase)
        }
        if (instanceType != null) {
            queryParams.put("instanceType", instanceType)
        }
        if (lastUpdated != null) {
            queryParams.put("lastUpdated", lastUpdated)
        }
        if (createdBy != null) {
            queryParams.put("createdBy", createdBy)
        }
        if (agentInstalled != null) {
            queryParams.put("agentInstalled", agentInstalled)
        }
        if (status != null) {
            queryParams.put("status", status)
        }
        if (environment != null) {
            queryParams.put("environment", environment)
        }
        if (showDeleted != null) {
            queryParams.put("showDeleted", showDeleted)
        }
        if (deleted != null) {
            queryParams.put("deleted", deleted)
        }
        if (expireDate != null) {
            queryParams.put("expireDate", expireDate)
        }
        if (expireDateMin != null) {
            queryParams.put("expireDateMin", expireDateMin)
        }
        if (expireDays != null) {
            queryParams.put("expireDays", expireDays)
        }
        if (expireDaysMin != null) {
            queryParams.put("expireDaysMin", expireDaysMin)
        }
        if (shutdownDate != null) {
            queryParams.put("shutdownDate", shutdownDate)
        }
        if (shutdownDateMin != null) {
            queryParams.put("shutdownDateMin", shutdownDateMin)
        }
        if (shutdownDays != null) {
            queryParams.put("shutdownDays", shutdownDays)
        }
        if (shutdownDaysMin != null) {
            queryParams.put("shutdownDaysMin", shutdownDaysMin)
        }
        if (labels != null) {
            queryParams.put("labels", labels)
        }
        if (allLabels != null) {
            queryParams.put("allLabels", allLabels)
        }
        if (tags != null) {
            queryParams.put("tags", tags)
        }
        if (metadata != null) {
            queryParams.put("metadata", metadata)
        }
        if (details != null) {
            queryParams.put("details", details)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def listSecurityGroupsInstance ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/instances/${id}/security-groups"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def lockInstance ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/instances/${id}/lock"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Model200Success.class )

    }

    def refreshStateInstance ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/instances/${id}/refresh"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    Model200Success.class )

    }

    def removeInstancesFromControl ( InlineObject99 inlineObject99, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/instances/remove-from-control"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType




        contentType = 'application/json';
        bodyParams = inlineObject99


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    SuccessMessage.class )

    }

    def resizeInstance ( Long id, InstanceResize instanceResize, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/instances/${id}/resize"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }



        contentType = 'application/json';
        bodyParams = instanceResize


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Object.class )

    }

    def restartInstance ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/instances/${id}/restart"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Object.class )

    }

    def revertSnapshotInstance ( Long id, BigDecimal snapshotId, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/instances/${id}/revert-snapshot/${snapshotId}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }
        // verify required params are set
        if (snapshotId == null) {
            throw new RuntimeException("missing required params snapshotId")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Model200Success.class )

    }

    def runWorkflowInstance ( Long id, Long workflowId, String workflowName, InstanceWorkflow instanceWorkflow, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/instances/${id}/workflow"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }

        if (workflowId != null) {
            queryParams.put("workflowId", workflowId)
        }
        if (workflowName != null) {
            queryParams.put("workflowName", workflowName)
        }


        contentType = 'application/json';
        bodyParams = instanceWorkflow


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Model200Success.class )

    }

    def setApplyInstance ( Long id, InlineObject91 inlineObject91, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/instances/${id}/apply"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }



        contentType = 'application/json';
        bodyParams = inlineObject91


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    Model200Success.class )

    }

    def setInstanceSecurityGroups ( Long id, InlineObject94 inlineObject94, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/instances/${id}/security-groups"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }



        contentType = 'application/json';
        bodyParams = inlineObject94


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    Object.class )

    }

    def snapshotInstance ( Long id, InstanceSnapshot instanceSnapshot, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/instances/${id}/snapshot"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }



        contentType = 'application/json';
        bodyParams = instanceSnapshot


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Model200Success.class )

    }

    def snapshotsInstance ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/instances/${id}/snapshots"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    InstanceSnapshots.class )

    }

    def startInstance ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/instances/${id}/start"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Object.class )

    }

    def stopInstance ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/instances/${id}/stop"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Object.class )

    }

    def suspendInstance ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/instances/${id}/suspend"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Model200Success.class )

    }

    def unlockInstance ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/instances/${id}/unlock"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Model200Success.class )

    }

    def updateInstance ( Long id, InstanceUpdate instanceUpdate, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/instances/${id}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }



        contentType = 'application/json';
        bodyParams = instanceUpdate


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Object.class )

    }

    def updateInstanceNetworkInterface ( Long id, BigDecimal networkInterfaceId, NetworkInterfaceUpdate networkInterfaceUpdate, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/instances/${id}/networkInterfaces/${networkInterfaceId}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }
        // verify required params are set
        if (networkInterfaceId == null) {
            throw new RuntimeException("missing required params networkInterfaceId")
        }



        contentType = 'application/json';
        bodyParams = networkInterfaceUpdate


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Object.class )

    }

    def updateInstanceSchedule ( Long id, Long scheduleId, InlineObject97 inlineObject97, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/instances/${id}/schedules/${scheduleId}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }
        // verify required params are set
        if (scheduleId == null) {
            throw new RuntimeException("missing required params scheduleId")
        }



        contentType = 'application/json';
        bodyParams = inlineObject97


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Object.class )

    }

    def updateInstanceThreshold ( Long id, InlineObject95 inlineObject95, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/instances/${id}/threshold"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }



        contentType = 'application/json';
        bodyParams = inlineObject95


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Object.class )

    }

    def updateWikiInstance ( Long id, InlineObject270 inlineObject270, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/instances/${id}/wiki"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }



        contentType = 'application/json';
        bodyParams = inlineObject270


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Object.class )

    }

}
