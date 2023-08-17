package org.openapitools.api;

import org.openapitools.api.ApiUtils
import org.openapitools.model.AddClient200Response
import org.openapitools.model.AddClientRequest
import org.openapitools.model.DefaultError
import org.openapitools.model.GetClients200Response
import org.openapitools.model.ListClients200Response
import org.openapitools.model.Model200Success
import org.openapitools.model.UpdateClientsRequest

class ClientsApi {
    String basePath = "https://CHANGEME"
    String versionPath = ""
    ApiUtils apiUtils = new ApiUtils();

    def addClient ( AddClientRequest addClientRequest, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/clients"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType




        contentType = 'application/json';
        bodyParams = addClientRequest


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    AddClient200Response.class )

    }

    def getClients ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/clients/${id}"

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
                    GetClients200Response.class )

    }

    def listClients ( Long max, Long offset, String sort, String direction, String phrase, String clientId, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/clients"

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
        if (clientId != null) {
            queryParams.put("clientId", clientId)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    ListClients200Response.class )

    }

    def removeClients ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/clients/${id}"

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

    def updateClients ( Long id, UpdateClientsRequest updateClientsRequest, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/clients/${id}"

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
        bodyParams = updateClientsRequest


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    AddClient200Response.class )

    }

}
