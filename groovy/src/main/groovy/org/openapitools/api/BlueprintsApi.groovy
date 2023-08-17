package org.openapitools.api;

import org.openapitools.api.ApiUtils
import org.openapitools.model.AddBlueprint200Response
import org.openapitools.model.AddBlueprintRequest
import org.openapitools.model.DefaultError
import org.openapitools.model.GetBlueprint200Response
import org.openapitools.model.ListBlueprints200Response
import org.openapitools.model.Model200Success
import org.openapitools.model.UpdateBlueprintPermissionsRequest

class BlueprintsApi {
    String basePath = "https://CHANGEME"
    String versionPath = ""
    ApiUtils apiUtils = new ApiUtils();

    def addBlueprint ( AddBlueprintRequest addBlueprintRequest, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/blueprints"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType




        contentType = 'application/json';
        bodyParams = addBlueprintRequest


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    AddBlueprint200Response.class )

    }

    def deleteBlueprint ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/blueprints/${id}"

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

    def getBlueprint ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/blueprints/${id}"

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
                    GetBlueprint200Response.class )

    }

    def listBlueprints ( Long max, Long offset, String name, String phrase, String labels, String allLabels, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/blueprints"

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
        if (labels != null) {
            queryParams.put("labels", labels)
        }
        if (allLabels != null) {
            queryParams.put("allLabels", allLabels)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    ListBlueprints200Response.class )

    }

    def updateBlueprint ( Long id, AddBlueprintRequest addBlueprintRequest, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/blueprints/${id}"

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
        bodyParams = addBlueprintRequest


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    GetBlueprint200Response.class )

    }

    def updateBlueprintImage ( Long id, File templateImage, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/blueprints/${id}/image"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }




        contentType = 'multipart/form-data';
        bodyParams = templateImage

        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    GetBlueprint200Response.class )

    }

    def updateBlueprintPermissions ( Long id, UpdateBlueprintPermissionsRequest updateBlueprintPermissionsRequest, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/blueprints/${id}/update-permissions"

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
        bodyParams = updateBlueprintPermissionsRequest


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    GetBlueprint200Response.class )

    }

}
