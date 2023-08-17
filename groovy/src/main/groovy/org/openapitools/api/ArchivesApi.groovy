package org.openapitools.api;

import org.openapitools.api.ApiUtils
import org.openapitools.model.AddArchiveBucket200Response
import org.openapitools.model.AddArchiveBucketRequest
import org.openapitools.model.AddArchiveFile200Response
import org.openapitools.model.AddArchiveFileLink200Response
import org.openapitools.model.DefaultError
import org.openapitools.model.GetArchiveBucket200Response
import org.openapitools.model.GetArchiveFileDetail200Response
import org.openapitools.model.GetArchiveFileLinks200Response
import org.openapitools.model.ListArchiveBuckets200Response
import org.openapitools.model.ListArchiveFiles200Response
import org.openapitools.model.Model200Success
import org.openapitools.model.UpdateArchiveBucketRequest

class ArchivesApi {
    String basePath = "https://CHANGEME"
    String versionPath = ""
    ApiUtils apiUtils = new ApiUtils();

    def addArchiveBucket ( AddArchiveBucketRequest addArchiveBucketRequest, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/archives/buckets"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType




        contentType = 'application/json';
        bodyParams = addArchiveBucketRequest


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    AddArchiveBucket200Response.class )

    }

    def addArchiveFile ( String bucket, String filepath, String filename, File _file, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/archives/buckets/${bucket}/files/${filepath}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (bucket == null) {
            throw new RuntimeException("missing required params bucket")
        }
        // verify required params are set
        if (filepath == null) {
            throw new RuntimeException("missing required params filepath")
        }

        if (filename != null) {
            queryParams.put("filename", filename)
        }



        contentType = 'multipart/form-data';
        bodyParams = _file

        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    AddArchiveFile200Response.class )

    }

    def addArchiveFileLink ( Long id, Long expireSeconds, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/archives/files/${id}/links"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }

        if (expireSeconds != null) {
            queryParams.put("expireSeconds", expireSeconds)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    AddArchiveFileLink200Response.class )

    }

    def deleteArchiveBucket ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/archives/buckets/${id}"

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

    def deleteArchiveFile ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/archives/files/${id}"

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

    def deleteArchiveFileLink ( Long id, Long linkId, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/archives/files/${id}/links/${linkId}"

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
        if (linkId == null) {
            throw new RuntimeException("missing required params linkId")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "DELETE", "",
                    Model200Success.class )

    }

    def getArchiveBucket ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/archives/buckets/${id}"

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
                    GetArchiveBucket200Response.class )

    }

    def getArchiveFile ( String bucket, String filepath, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/archives/download/${bucket}/${filepath}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (bucket == null) {
            throw new RuntimeException("missing required params bucket")
        }
        // verify required params are set
        if (filepath == null) {
            throw new RuntimeException("missing required params filepath")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    null )

    }

    def getArchiveFileDetail ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/archives/files/${id}"

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
                    GetArchiveFileDetail200Response.class )

    }

    def getArchiveFileLinks ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/archives/files/${id}/links"

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
                    GetArchiveFileLinks200Response.class )

    }

    def listArchiveBuckets ( String name, String phrase, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/archives/buckets"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType


        if (name != null) {
            queryParams.put("name", name)
        }
        if (phrase != null) {
            queryParams.put("phrase", phrase)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    ListArchiveBuckets200Response.class )

    }

    def listArchiveFiles ( String bucket, String filepath, String name, String phrase, Boolean fullTree, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/archives/buckets/${bucket}/files/${filepath}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (bucket == null) {
            throw new RuntimeException("missing required params bucket")
        }
        // verify required params are set
        if (filepath == null) {
            throw new RuntimeException("missing required params filepath")
        }

        if (name != null) {
            queryParams.put("name", name)
        }
        if (phrase != null) {
            queryParams.put("phrase", phrase)
        }
        if (fullTree != null) {
            queryParams.put("fullTree", fullTree)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    ListArchiveFiles200Response.class )

    }

    def updateArchiveBucket ( Long id, UpdateArchiveBucketRequest updateArchiveBucketRequest, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/archives/buckets/${id}"

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
        bodyParams = updateArchiveBucketRequest


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    AddArchiveBucket200Response.class )

    }

}
