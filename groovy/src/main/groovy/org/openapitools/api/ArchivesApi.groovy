package org.openapitools.api;

import org.openapitools.api.ApiUtils
import org.openapitools.model.DefaultError
import org.openapitools.model.InlineObject7
import org.openapitools.model.InlineObject8
import org.openapitools.model.InlineResponse2004
import org.openapitools.model.InlineResponse2005
import org.openapitools.model.Model200Success

class ArchivesApi {
    String basePath = "https://CHANGEME"
    String versionPath = ""
    ApiUtils apiUtils = new ApiUtils();

    def addArchiveBucket ( InlineObject7 inlineObject7, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/archives/buckets"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType




        contentType = 'application/json';
        bodyParams = inlineObject7


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    Object.class )

    }

    def addArchiveFile ( String bucket, String filepath, String filename, File file, Closure onSuccess, Closure onFailure)  {
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
        bodyParams = file

        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    Object.class )

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
                    Model200Success.class )

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
                    InlineResponse2004.class )

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
                    InlineResponse2005.class )

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
                    Object.class )

    }

    def getArchivePublicFile ( String bucket, String filepath, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/public-archives/download/${bucket}/${filepath}"

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

    def getArchivePublicFileLink ( String s, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/public-archives/link"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (s == null) {
            throw new RuntimeException("missing required params s")
        }

        if (s != null) {
            queryParams.put("s", s)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    null )

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
                    Object.class )

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
                    Object.class )

    }

    def updateArchiveBucket ( Long id, InlineObject8 inlineObject8, Closure onSuccess, Closure onFailure)  {
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
        bodyParams = inlineObject8


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Object.class )

    }

}
