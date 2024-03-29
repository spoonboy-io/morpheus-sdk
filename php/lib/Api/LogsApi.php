<?php
/**
 * LogsApi
 * PHP version 7.2
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 */

/**
 * Morpheus API
 *
 * Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.
 *
 * The version of the OpenAPI document: 6.2.1
 * Contact: dev@morpheusdata.com
 * Generated by: https://openapi-generator.tech
 * OpenAPI Generator version: 5.0.0
 */

/**
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

namespace OpenAPI\Client\Api;

use GuzzleHttp\Client;
use GuzzleHttp\ClientInterface;
use GuzzleHttp\Exception\RequestException;
use GuzzleHttp\Psr7\MultipartStream;
use GuzzleHttp\Psr7\Request;
use GuzzleHttp\RequestOptions;
use OpenAPI\Client\ApiException;
use OpenAPI\Client\Configuration;
use OpenAPI\Client\HeaderSelector;
use OpenAPI\Client\ObjectSerializer;

/**
 * LogsApi Class Doc Comment
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 */
class LogsApi
{
    /**
     * @var ClientInterface
     */
    protected $client;

    /**
     * @var Configuration
     */
    protected $config;

    /**
     * @var HeaderSelector
     */
    protected $headerSelector;

    /**
     * @var int Host index
     */
    protected $hostIndex;

    /**
     * @param ClientInterface $client
     * @param Configuration   $config
     * @param HeaderSelector  $selector
     * @param int             $hostIndex (Optional) host index to select the list of hosts if defined in the OpenAPI spec
     */
    public function __construct(
        ClientInterface $client = null,
        Configuration $config = null,
        HeaderSelector $selector = null,
        $hostIndex = 0
    ) {
        $this->client = $client ?: new Client();
        $this->config = $config ?: new Configuration();
        $this->headerSelector = $selector ?: new HeaderSelector();
        $this->hostIndex = $hostIndex;
    }

    /**
     * Set the host index
     *
     * @param int $hostIndex Host index (required)
     */
    public function setHostIndex($hostIndex)
    {
        $this->hostIndex = $hostIndex;
    }

    /**
     * Get the host index
     *
     * @return int Host index
     */
    public function getHostIndex()
    {
        return $this->hostIndex;
    }

    /**
     * @return Configuration
     */
    public function getConfig()
    {
        return $this->config;
    }

    /**
     * Operation listLogs
     *
     * Retrieves Logs
     *
     * @param  int $max Maximum number of records to return, -1 can be used to fetch all records (optional, default to 25)
     * @param  int $offset Offset records, the number of records to skip, for paginating requests (optional, default to 0)
     * @param  string $sort Sort order, the name of the property to sort by (optional, default to 'name')
     * @param  string $order Sort direction, use &#39;desc&#39; to reverse sort (optional, default to 'asc')
     * @param  string $query Alias for phrase (optional)
     * @param  string $message Filter by message (optional)
     * @param  string $source_type Filter by source type (optional)
     * @param  string $type_code Filter by code type (optional)
     * @param  int $object_id Filter by objectId (optional)
     * @param  string $token Filter by token (optional)
     * @param  string $level Filter by log level. Multiple values can be passed pipe delimited. (optional)
     * @param  int $start_ms Date filter in milliseconds (unix epoch), restricts query to only load logs updated more recently than the time specified. (optional)
     * @param  int $end_ms Date filter in milliseconds (unix epoch), restricts query to only load logs updated before the time specified. (optional)
     * @param  \DateTime $start_date_time Start Date timestamp (ISO 8601) (optional)
     * @param  \DateTime $end_date_time End Date timestamp (ISO 8601) (optional)
     * @param  int $containers The Container ID(s) for filtering. Accepts multiple values. (optional)
     * @param  int $servers The Server ID(s) for filtering. Accepts multiple values. (optional)
     * @param  int $cluster_id The Cluster ID(s) for filtering. Accepts multiple values. (optional)
     *
     * @throws \OpenAPI\Client\ApiException on non-2xx response
     * @throws \InvalidArgumentException
     * @return Log|\OpenAPI\Client\Model\DefaultError|\OpenAPI\Client\Model\DefaultError
     */
    public function listLogs($max = 25, $offset = 0, $sort = 'name', $order = 'asc', $query = null, $message = null, $source_type = null, $type_code = null, $object_id = null, $token = null, $level = null, $start_ms = null, $end_ms = null, $start_date_time = null, $end_date_time = null, $containers = null, $servers = null, $cluster_id = null)
    {
        list($response) = $this->listLogsWithHttpInfo($max, $offset, $sort, $order, $query, $message, $source_type, $type_code, $object_id, $token, $level, $start_ms, $end_ms, $start_date_time, $end_date_time, $containers, $servers, $cluster_id);
        return $response;
    }

    /**
     * Operation listLogsWithHttpInfo
     *
     * Retrieves Logs
     *
     * @param  int $max Maximum number of records to return, -1 can be used to fetch all records (optional, default to 25)
     * @param  int $offset Offset records, the number of records to skip, for paginating requests (optional, default to 0)
     * @param  string $sort Sort order, the name of the property to sort by (optional, default to 'name')
     * @param  string $order Sort direction, use &#39;desc&#39; to reverse sort (optional, default to 'asc')
     * @param  string $query Alias for phrase (optional)
     * @param  string $message Filter by message (optional)
     * @param  string $source_type Filter by source type (optional)
     * @param  string $type_code Filter by code type (optional)
     * @param  int $object_id Filter by objectId (optional)
     * @param  string $token Filter by token (optional)
     * @param  string $level Filter by log level. Multiple values can be passed pipe delimited. (optional)
     * @param  int $start_ms Date filter in milliseconds (unix epoch), restricts query to only load logs updated more recently than the time specified. (optional)
     * @param  int $end_ms Date filter in milliseconds (unix epoch), restricts query to only load logs updated before the time specified. (optional)
     * @param  \DateTime $start_date_time Start Date timestamp (ISO 8601) (optional)
     * @param  \DateTime $end_date_time End Date timestamp (ISO 8601) (optional)
     * @param  int $containers The Container ID(s) for filtering. Accepts multiple values. (optional)
     * @param  int $servers The Server ID(s) for filtering. Accepts multiple values. (optional)
     * @param  int $cluster_id The Cluster ID(s) for filtering. Accepts multiple values. (optional)
     *
     * @throws \OpenAPI\Client\ApiException on non-2xx response
     * @throws \InvalidArgumentException
     * @return array of Log|\OpenAPI\Client\Model\DefaultError|\OpenAPI\Client\Model\DefaultError, HTTP status code, HTTP response headers (array of strings)
     */
    public function listLogsWithHttpInfo($max = 25, $offset = 0, $sort = 'name', $order = 'asc', $query = null, $message = null, $source_type = null, $type_code = null, $object_id = null, $token = null, $level = null, $start_ms = null, $end_ms = null, $start_date_time = null, $end_date_time = null, $containers = null, $servers = null, $cluster_id = null)
    {
        $request = $this->listLogsRequest($max, $offset, $sort, $order, $query, $message, $source_type, $type_code, $object_id, $token, $level, $start_ms, $end_ms, $start_date_time, $end_date_time, $containers, $servers, $cluster_id);

        try {
            $options = $this->createHttpClientOption();
            try {
                $response = $this->client->send($request, $options);
            } catch (RequestException $e) {
                throw new ApiException(
                    "[{$e->getCode()}] {$e->getMessage()}",
                    $e->getCode(),
                    $e->getResponse() ? $e->getResponse()->getHeaders() : null,
                    $e->getResponse() ? (string) $e->getResponse()->getBody() : null
                );
            }

            $statusCode = $response->getStatusCode();

            if ($statusCode < 200 || $statusCode > 299) {
                throw new ApiException(
                    sprintf(
                        '[%d] Error connecting to the API (%s)',
                        $statusCode,
                        $request->getUri()
                    ),
                    $statusCode,
                    $response->getHeaders(),
                    $response->getBody()
                );
            }

            $responseBody = $response->getBody();
            switch($statusCode) {
                case 200:
                    if ('Log' === '\SplFileObject') {
                        $content = $responseBody; //stream goes to serializer
                    } else {
                        $content = (string) $responseBody;
                    }

                    return [
                        ObjectSerializer::deserialize($content, 'Log', []),
                        $response->getStatusCode(),
                        $response->getHeaders()
                    ];
                case 400,401,402,403,404,405,406,407,408,409,410,411,412,413,414,415,416,417:
                    if ('\OpenAPI\Client\Model\DefaultError' === '\SplFileObject') {
                        $content = $responseBody; //stream goes to serializer
                    } else {
                        $content = (string) $responseBody;
                    }

                    return [
                        ObjectSerializer::deserialize($content, '\OpenAPI\Client\Model\DefaultError', []),
                        $response->getStatusCode(),
                        $response->getHeaders()
                    ];
                case 500,501,502,503,504,505,506,507,509,510:
                    if ('\OpenAPI\Client\Model\DefaultError' === '\SplFileObject') {
                        $content = $responseBody; //stream goes to serializer
                    } else {
                        $content = (string) $responseBody;
                    }

                    return [
                        ObjectSerializer::deserialize($content, '\OpenAPI\Client\Model\DefaultError', []),
                        $response->getStatusCode(),
                        $response->getHeaders()
                    ];
            }

            $returnType = 'Log';
            $responseBody = $response->getBody();
            if ($returnType === '\SplFileObject') {
                $content = $responseBody; //stream goes to serializer
            } else {
                $content = (string) $responseBody;
            }

            return [
                ObjectSerializer::deserialize($content, $returnType, []),
                $response->getStatusCode(),
                $response->getHeaders()
            ];

        } catch (ApiException $e) {
            switch ($e->getCode()) {
                case 200:
                    $data = ObjectSerializer::deserialize(
                        $e->getResponseBody(),
                        'Log',
                        $e->getResponseHeaders()
                    );
                    $e->setResponseObject($data);
                    break;
                case 400,401,402,403,404,405,406,407,408,409,410,411,412,413,414,415,416,417:
                    $data = ObjectSerializer::deserialize(
                        $e->getResponseBody(),
                        '\OpenAPI\Client\Model\DefaultError',
                        $e->getResponseHeaders()
                    );
                    $e->setResponseObject($data);
                    break;
                case 500,501,502,503,504,505,506,507,509,510:
                    $data = ObjectSerializer::deserialize(
                        $e->getResponseBody(),
                        '\OpenAPI\Client\Model\DefaultError',
                        $e->getResponseHeaders()
                    );
                    $e->setResponseObject($data);
                    break;
            }
            throw $e;
        }
    }

    /**
     * Operation listLogsAsync
     *
     * Retrieves Logs
     *
     * @param  int $max Maximum number of records to return, -1 can be used to fetch all records (optional, default to 25)
     * @param  int $offset Offset records, the number of records to skip, for paginating requests (optional, default to 0)
     * @param  string $sort Sort order, the name of the property to sort by (optional, default to 'name')
     * @param  string $order Sort direction, use &#39;desc&#39; to reverse sort (optional, default to 'asc')
     * @param  string $query Alias for phrase (optional)
     * @param  string $message Filter by message (optional)
     * @param  string $source_type Filter by source type (optional)
     * @param  string $type_code Filter by code type (optional)
     * @param  int $object_id Filter by objectId (optional)
     * @param  string $token Filter by token (optional)
     * @param  string $level Filter by log level. Multiple values can be passed pipe delimited. (optional)
     * @param  int $start_ms Date filter in milliseconds (unix epoch), restricts query to only load logs updated more recently than the time specified. (optional)
     * @param  int $end_ms Date filter in milliseconds (unix epoch), restricts query to only load logs updated before the time specified. (optional)
     * @param  \DateTime $start_date_time Start Date timestamp (ISO 8601) (optional)
     * @param  \DateTime $end_date_time End Date timestamp (ISO 8601) (optional)
     * @param  int $containers The Container ID(s) for filtering. Accepts multiple values. (optional)
     * @param  int $servers The Server ID(s) for filtering. Accepts multiple values. (optional)
     * @param  int $cluster_id The Cluster ID(s) for filtering. Accepts multiple values. (optional)
     *
     * @throws \InvalidArgumentException
     * @return \GuzzleHttp\Promise\PromiseInterface
     */
    public function listLogsAsync($max = 25, $offset = 0, $sort = 'name', $order = 'asc', $query = null, $message = null, $source_type = null, $type_code = null, $object_id = null, $token = null, $level = null, $start_ms = null, $end_ms = null, $start_date_time = null, $end_date_time = null, $containers = null, $servers = null, $cluster_id = null)
    {
        return $this->listLogsAsyncWithHttpInfo($max, $offset, $sort, $order, $query, $message, $source_type, $type_code, $object_id, $token, $level, $start_ms, $end_ms, $start_date_time, $end_date_time, $containers, $servers, $cluster_id)
            ->then(
                function ($response) {
                    return $response[0];
                }
            );
    }

    /**
     * Operation listLogsAsyncWithHttpInfo
     *
     * Retrieves Logs
     *
     * @param  int $max Maximum number of records to return, -1 can be used to fetch all records (optional, default to 25)
     * @param  int $offset Offset records, the number of records to skip, for paginating requests (optional, default to 0)
     * @param  string $sort Sort order, the name of the property to sort by (optional, default to 'name')
     * @param  string $order Sort direction, use &#39;desc&#39; to reverse sort (optional, default to 'asc')
     * @param  string $query Alias for phrase (optional)
     * @param  string $message Filter by message (optional)
     * @param  string $source_type Filter by source type (optional)
     * @param  string $type_code Filter by code type (optional)
     * @param  int $object_id Filter by objectId (optional)
     * @param  string $token Filter by token (optional)
     * @param  string $level Filter by log level. Multiple values can be passed pipe delimited. (optional)
     * @param  int $start_ms Date filter in milliseconds (unix epoch), restricts query to only load logs updated more recently than the time specified. (optional)
     * @param  int $end_ms Date filter in milliseconds (unix epoch), restricts query to only load logs updated before the time specified. (optional)
     * @param  \DateTime $start_date_time Start Date timestamp (ISO 8601) (optional)
     * @param  \DateTime $end_date_time End Date timestamp (ISO 8601) (optional)
     * @param  int $containers The Container ID(s) for filtering. Accepts multiple values. (optional)
     * @param  int $servers The Server ID(s) for filtering. Accepts multiple values. (optional)
     * @param  int $cluster_id The Cluster ID(s) for filtering. Accepts multiple values. (optional)
     *
     * @throws \InvalidArgumentException
     * @return \GuzzleHttp\Promise\PromiseInterface
     */
    public function listLogsAsyncWithHttpInfo($max = 25, $offset = 0, $sort = 'name', $order = 'asc', $query = null, $message = null, $source_type = null, $type_code = null, $object_id = null, $token = null, $level = null, $start_ms = null, $end_ms = null, $start_date_time = null, $end_date_time = null, $containers = null, $servers = null, $cluster_id = null)
    {
        $returnType = 'Log';
        $request = $this->listLogsRequest($max, $offset, $sort, $order, $query, $message, $source_type, $type_code, $object_id, $token, $level, $start_ms, $end_ms, $start_date_time, $end_date_time, $containers, $servers, $cluster_id);

        return $this->client
            ->sendAsync($request, $this->createHttpClientOption())
            ->then(
                function ($response) use ($returnType) {
                    $responseBody = $response->getBody();
                    if ($returnType === '\SplFileObject') {
                        $content = $responseBody; //stream goes to serializer
                    } else {
                        $content = (string) $responseBody;
                    }

                    return [
                        ObjectSerializer::deserialize($content, $returnType, []),
                        $response->getStatusCode(),
                        $response->getHeaders()
                    ];
                },
                function ($exception) {
                    $response = $exception->getResponse();
                    $statusCode = $response->getStatusCode();
                    throw new ApiException(
                        sprintf(
                            '[%d] Error connecting to the API (%s)',
                            $statusCode,
                            $exception->getRequest()->getUri()
                        ),
                        $statusCode,
                        $response->getHeaders(),
                        $response->getBody()
                    );
                }
            );
    }

    /**
     * Create request for operation 'listLogs'
     *
     * @param  int $max Maximum number of records to return, -1 can be used to fetch all records (optional, default to 25)
     * @param  int $offset Offset records, the number of records to skip, for paginating requests (optional, default to 0)
     * @param  string $sort Sort order, the name of the property to sort by (optional, default to 'name')
     * @param  string $order Sort direction, use &#39;desc&#39; to reverse sort (optional, default to 'asc')
     * @param  string $query Alias for phrase (optional)
     * @param  string $message Filter by message (optional)
     * @param  string $source_type Filter by source type (optional)
     * @param  string $type_code Filter by code type (optional)
     * @param  int $object_id Filter by objectId (optional)
     * @param  string $token Filter by token (optional)
     * @param  string $level Filter by log level. Multiple values can be passed pipe delimited. (optional)
     * @param  int $start_ms Date filter in milliseconds (unix epoch), restricts query to only load logs updated more recently than the time specified. (optional)
     * @param  int $end_ms Date filter in milliseconds (unix epoch), restricts query to only load logs updated before the time specified. (optional)
     * @param  \DateTime $start_date_time Start Date timestamp (ISO 8601) (optional)
     * @param  \DateTime $end_date_time End Date timestamp (ISO 8601) (optional)
     * @param  int $containers The Container ID(s) for filtering. Accepts multiple values. (optional)
     * @param  int $servers The Server ID(s) for filtering. Accepts multiple values. (optional)
     * @param  int $cluster_id The Cluster ID(s) for filtering. Accepts multiple values. (optional)
     *
     * @throws \InvalidArgumentException
     * @return \GuzzleHttp\Psr7\Request
     */
    public function listLogsRequest($max = 25, $offset = 0, $sort = 'name', $order = 'asc', $query = null, $message = null, $source_type = null, $type_code = null, $object_id = null, $token = null, $level = null, $start_ms = null, $end_ms = null, $start_date_time = null, $end_date_time = null, $containers = null, $servers = null, $cluster_id = null)
    {

        $resourcePath = '/api/logs';
        $formParams = [];
        $queryParams = [];
        $headerParams = [];
        $httpBody = '';
        $multipart = false;

        // query params
        if ($max !== null) {
            if('form' === 'form' && is_array($max)) {
                foreach($max as $key => $value) {
                    $queryParams[$key] = $value;
                }
            }
            else {
                $queryParams['max'] = $max;
            }
        }
        // query params
        if ($offset !== null) {
            if('form' === 'form' && is_array($offset)) {
                foreach($offset as $key => $value) {
                    $queryParams[$key] = $value;
                }
            }
            else {
                $queryParams['offset'] = $offset;
            }
        }
        // query params
        if ($sort !== null) {
            if('form' === 'form' && is_array($sort)) {
                foreach($sort as $key => $value) {
                    $queryParams[$key] = $value;
                }
            }
            else {
                $queryParams['sort'] = $sort;
            }
        }
        // query params
        if ($order !== null) {
            if('form' === 'form' && is_array($order)) {
                foreach($order as $key => $value) {
                    $queryParams[$key] = $value;
                }
            }
            else {
                $queryParams['order'] = $order;
            }
        }
        // query params
        if ($query !== null) {
            if('form' === 'form' && is_array($query)) {
                foreach($query as $key => $value) {
                    $queryParams[$key] = $value;
                }
            }
            else {
                $queryParams['query'] = $query;
            }
        }
        // query params
        if ($message !== null) {
            if('form' === 'form' && is_array($message)) {
                foreach($message as $key => $value) {
                    $queryParams[$key] = $value;
                }
            }
            else {
                $queryParams['message'] = $message;
            }
        }
        // query params
        if ($source_type !== null) {
            if('form' === 'form' && is_array($source_type)) {
                foreach($source_type as $key => $value) {
                    $queryParams[$key] = $value;
                }
            }
            else {
                $queryParams['sourceType'] = $source_type;
            }
        }
        // query params
        if ($type_code !== null) {
            if('form' === 'form' && is_array($type_code)) {
                foreach($type_code as $key => $value) {
                    $queryParams[$key] = $value;
                }
            }
            else {
                $queryParams['typeCode'] = $type_code;
            }
        }
        // query params
        if ($object_id !== null) {
            if('form' === 'form' && is_array($object_id)) {
                foreach($object_id as $key => $value) {
                    $queryParams[$key] = $value;
                }
            }
            else {
                $queryParams['objectId'] = $object_id;
            }
        }
        // query params
        if ($token !== null) {
            if('form' === 'form' && is_array($token)) {
                foreach($token as $key => $value) {
                    $queryParams[$key] = $value;
                }
            }
            else {
                $queryParams['token'] = $token;
            }
        }
        // query params
        if ($level !== null) {
            if('form' === 'form' && is_array($level)) {
                foreach($level as $key => $value) {
                    $queryParams[$key] = $value;
                }
            }
            else {
                $queryParams['level'] = $level;
            }
        }
        // query params
        if ($start_ms !== null) {
            if('form' === 'form' && is_array($start_ms)) {
                foreach($start_ms as $key => $value) {
                    $queryParams[$key] = $value;
                }
            }
            else {
                $queryParams['startMs'] = $start_ms;
            }
        }
        // query params
        if ($end_ms !== null) {
            if('form' === 'form' && is_array($end_ms)) {
                foreach($end_ms as $key => $value) {
                    $queryParams[$key] = $value;
                }
            }
            else {
                $queryParams['endMs'] = $end_ms;
            }
        }
        // query params
        if ($start_date_time !== null) {
            if('form' === 'form' && is_array($start_date_time)) {
                foreach($start_date_time as $key => $value) {
                    $queryParams[$key] = $value;
                }
            }
            else {
                $queryParams['startDateTime'] = $start_date_time;
            }
        }
        // query params
        if ($end_date_time !== null) {
            if('form' === 'form' && is_array($end_date_time)) {
                foreach($end_date_time as $key => $value) {
                    $queryParams[$key] = $value;
                }
            }
            else {
                $queryParams['endDateTime'] = $end_date_time;
            }
        }
        // query params
        if ($containers !== null) {
            if('form' === 'form' && is_array($containers)) {
                foreach($containers as $key => $value) {
                    $queryParams[$key] = $value;
                }
            }
            else {
                $queryParams['containers'] = $containers;
            }
        }
        // query params
        if ($servers !== null) {
            if('form' === 'form' && is_array($servers)) {
                foreach($servers as $key => $value) {
                    $queryParams[$key] = $value;
                }
            }
            else {
                $queryParams['servers'] = $servers;
            }
        }
        // query params
        if ($cluster_id !== null) {
            if('form' === 'form' && is_array($cluster_id)) {
                foreach($cluster_id as $key => $value) {
                    $queryParams[$key] = $value;
                }
            }
            else {
                $queryParams['clusterId'] = $cluster_id;
            }
        }




        if ($multipart) {
            $headers = $this->headerSelector->selectHeadersForMultipart(
                ['application/json']
            );
        } else {
            $headers = $this->headerSelector->selectHeaders(
                ['application/json'],
                []
            );
        }

        // for model (json/xml)
        if (count($formParams) > 0) {
            if ($multipart) {
                $multipartContents = [];
                foreach ($formParams as $formParamName => $formParamValue) {
                    $formParamValueItems = is_array($formParamValue) ? $formParamValue : [$formParamValue];
                    foreach ($formParamValueItems as $formParamValueItem) {
                        $multipartContents[] = [
                            'name' => $formParamName,
                            'contents' => $formParamValueItem
                        ];
                    }
                }
                // for HTTP post (form)
                $httpBody = new MultipartStream($multipartContents);

            } elseif ($headers['Content-Type'] === 'application/json') {
                $httpBody = \GuzzleHttp\json_encode($formParams);

            } else {
                // for HTTP post (form)
                $httpBody = \GuzzleHttp\Psr7\build_query($formParams);
            }
        }

        // this endpoint requires Bearer authentication (access token)
        if ($this->config->getAccessToken() !== null) {
            $headers['Authorization'] = 'Bearer ' . $this->config->getAccessToken();
        }

        $defaultHeaders = [];
        if ($this->config->getUserAgent()) {
            $defaultHeaders['User-Agent'] = $this->config->getUserAgent();
        }

        $headers = array_merge(
            $defaultHeaders,
            $headerParams,
            $headers
        );

        $query = \GuzzleHttp\Psr7\build_query($queryParams);
        return new Request(
            'GET',
            $this->config->getHost() . $resourcePath . ($query ? "?{$query}" : ''),
            $headers,
            $httpBody
        );
    }

    /**
     * Create http client option
     *
     * @throws \RuntimeException on file opening failure
     * @return array of http client options
     */
    protected function createHttpClientOption()
    {
        $options = [];
        if ($this->config->getDebug()) {
            $options[RequestOptions::DEBUG] = fopen($this->config->getDebugFile(), 'a');
            if (!$options[RequestOptions::DEBUG]) {
                throw new \RuntimeException('Failed to open the debug file: ' . $this->config->getDebugFile());
            }
        }

        return $options;
    }
}
