package org.openapitools.api;

import org.openapitools.api.ApiUtils
import org.openapitools.model.DefaultError
import org.openapitools.model.InlineObject29
import org.openapitools.model.InlineObject30
import org.openapitools.model.InlineObject31
import org.openapitools.model.InlineObject32
import org.openapitools.model.InlineObject33
import org.openapitools.model.InlineObject34
import org.openapitools.model.InlineObject35
import org.openapitools.model.InlineObject36
import org.openapitools.model.InlineObject37
import org.openapitools.model.InlineObject38
import org.openapitools.model.InlineObject39
import org.openapitools.model.InlineObject40
import org.openapitools.model.InlineResponse20016
import org.openapitools.model.InlineResponse20017
import org.openapitools.model.InlineResponse20018
import org.openapitools.model.InlineResponse20019
import org.openapitools.model.Model200Success

class ChecksApi {
    String basePath = "https://CHANGEME"
    String versionPath = ""
    ApiUtils apiUtils = new ApiUtils();

    def addCheckApps ( InlineObject29 inlineObject29, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/monitoring/apps"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType




        contentType = 'application/json';
        bodyParams = inlineObject29


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    Object.class )

    }

    def addCheckGroups ( InlineObject37 inlineObject37, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/monitoring/groups"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType




        contentType = 'application/json';
        bodyParams = inlineObject37


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    Object.class )

    }

    def addChecks ( InlineObject33 inlineObject33, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/monitoring/checks"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType




        contentType = 'application/json';
        bodyParams = inlineObject33


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    Object.class )

    }

    def deleteCheckApps ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/monitoring/apps/${id}"

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

    def deleteCheckGroups ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/monitoring/groups/${id}"

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

    def deleteChecks ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/monitoring/checks/${id}"

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

    def getCheckApps ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/monitoring/apps/${id}"

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
                    InlineResponse20016.class )

    }

    def getCheckGroups ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/monitoring/groups/${id}"

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
                    InlineResponse20019.class )

    }

    def getCheckTypes ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/monitoring/check-types/${id}"

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
                    InlineResponse20018.class )

    }

    def getChecks ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/monitoring/checks/${id}"

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
                    InlineResponse20017.class )

    }

    def listCheckApps ( Long max, Long offset, String sort, String name, String phrase, String status, Date lastUpdated, Boolean deleted, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/monitoring/apps"

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
        if (name != null) {
            queryParams.put("name", name)
        }
        if (phrase != null) {
            queryParams.put("phrase", phrase)
        }
        if (status != null) {
            queryParams.put("status", status)
        }
        if (lastUpdated != null) {
            queryParams.put("lastUpdated", lastUpdated)
        }
        if (deleted != null) {
            queryParams.put("deleted", deleted)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def listCheckGroups ( Long max, Long offset, String sort, String name, String phrase, String status, Date lastUpdated, Boolean deleted, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/monitoring/groups"

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
        if (name != null) {
            queryParams.put("name", name)
        }
        if (phrase != null) {
            queryParams.put("phrase", phrase)
        }
        if (status != null) {
            queryParams.put("status", status)
        }
        if (lastUpdated != null) {
            queryParams.put("lastUpdated", lastUpdated)
        }
        if (deleted != null) {
            queryParams.put("deleted", deleted)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def listCheckTypes ( Long max, Long offset, String sort, String name, String phrase, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/monitoring/check-types"

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
        if (name != null) {
            queryParams.put("name", name)
        }
        if (phrase != null) {
            queryParams.put("phrase", phrase)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def listChecks ( Long max, Long offset, String sort, String name, String phrase, String status, Date lastUpdated, Boolean deleted, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/monitoring/checks"

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
        if (name != null) {
            queryParams.put("name", name)
        }
        if (phrase != null) {
            queryParams.put("phrase", phrase)
        }
        if (status != null) {
            queryParams.put("status", status)
        }
        if (lastUpdated != null) {
            queryParams.put("lastUpdated", lastUpdated)
        }
        if (deleted != null) {
            queryParams.put("deleted", deleted)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def updateCheckApps ( Long id, InlineObject31 inlineObject31, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/monitoring/apps/${id}"

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
        bodyParams = inlineObject31


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Object.class )

    }

    def updateCheckGroups ( Long id, InlineObject38 inlineObject38, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/monitoring/groups/${id}"

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
        bodyParams = inlineObject38


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Object.class )

    }

    def updateChecks ( Long id, InlineObject35 inlineObject35, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/monitoring/checks/${id}"

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
        bodyParams = inlineObject35


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Object.class )

    }

    def updateMuteAllCheckApps ( InlineObject30 inlineObject30, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/monitoring/apps/mute-all"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType




        contentType = 'application/json';
        bodyParams = inlineObject30


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Object.class )

    }

    def updateMuteAllCheckGroups ( InlineObject40 inlineObject40, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/monitoring/groups/mute-all"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType




        contentType = 'application/json';
        bodyParams = inlineObject40


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Object.class )

    }

    def updateMuteAllChecks ( InlineObject34 inlineObject34, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/monitoring/checks/mute-all"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType




        contentType = 'application/json';
        bodyParams = inlineObject34


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Object.class )

    }

    def updateMuteCheckApps ( Long id, InlineObject32 inlineObject32, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/monitoring/apps/${id}/mute"

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
        bodyParams = inlineObject32


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Object.class )

    }

    def updateMuteCheckGroups ( Long id, InlineObject39 inlineObject39, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/monitoring/groups/${id}/mute"

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
        bodyParams = inlineObject39


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Object.class )

    }

    def updateMuteChecks ( Long id, InlineObject36 inlineObject36, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/monitoring/checks/${id}/mute"

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
        bodyParams = inlineObject36


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Object.class )

    }

}
