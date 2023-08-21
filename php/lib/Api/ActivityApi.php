<?php
/**
 * ActivityApi
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
 * ActivityApi Class Doc Comment
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 */
class ActivityApi
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
     * Operation listActivity
     *
     * Retrieves Activity
     *
     * @param  int $max Maximum number of records to return, -1 can be used to fetch all records (optional, default to 25)
     * @param  int $offset Offset records, the number of records to skip, for paginating requests (optional, default to 0)
     * @param  string $sort Sort order, the name of the property to sort by (optional, default to 'name')
     * @param  string $order Sort direction, use &#39;desc&#39; to reverse sort (optional, default to 'asc')
     * @param  string $phrase Search phrase for partial matches on name or description (optional)
     * @param  string $name Filter by name, wildcard may be specified as %, eg. example-% (optional)
     * @param  int $user_id Filter by User ID (optional)
     * @param  float $tenant_id Filter by Tenant ID. Only available to the master account. (optional)
     * @param  string $timeframe Filter by a timeframe (ex - today, yesterday, week, month, 3months) (optional, default to 'month')
     * @param  \DateTime $start Filter by activity on or after a date(time). Default is 1 month prior (optional)
     * @param  \DateTime $end Filter by activity on or before a date(time). Default is current date (optional)
     *
     * @throws \OpenAPI\Client\ApiException on non-2xx response
     * @throws \InvalidArgumentException
     * @return object|\OpenAPI\Client\Model\DefaultError|\OpenAPI\Client\Model\DefaultError
     */
    public function listActivity($max = 25, $offset = 0, $sort = 'name', $order = 'asc', $phrase = null, $name = null, $user_id = null, $tenant_id = null, $timeframe = 'month', $start = null, $end = null)
    {
        list($response) = $this->listActivityWithHttpInfo($max, $offset, $sort, $order, $phrase, $name, $user_id, $tenant_id, $timeframe, $start, $end);
        return $response;
    }

    /**
     * Operation listActivityWithHttpInfo
     *
     * Retrieves Activity
     *
     * @param  int $max Maximum number of records to return, -1 can be used to fetch all records (optional, default to 25)
     * @param  int $offset Offset records, the number of records to skip, for paginating requests (optional, default to 0)
     * @param  string $sort Sort order, the name of the property to sort by (optional, default to 'name')
     * @param  string $order Sort direction, use &#39;desc&#39; to reverse sort (optional, default to 'asc')
     * @param  string $phrase Search phrase for partial matches on name or description (optional)
     * @param  string $name Filter by name, wildcard may be specified as %, eg. example-% (optional)
     * @param  int $user_id Filter by User ID (optional)
     * @param  float $tenant_id Filter by Tenant ID. Only available to the master account. (optional)
     * @param  string $timeframe Filter by a timeframe (ex - today, yesterday, week, month, 3months) (optional, default to 'month')
     * @param  \DateTime $start Filter by activity on or after a date(time). Default is 1 month prior (optional)
     * @param  \DateTime $end Filter by activity on or before a date(time). Default is current date (optional)
     *
     * @throws \OpenAPI\Client\ApiException on non-2xx response
     * @throws \InvalidArgumentException
     * @return array of object|\OpenAPI\Client\Model\DefaultError|\OpenAPI\Client\Model\DefaultError, HTTP status code, HTTP response headers (array of strings)
     */
    public function listActivityWithHttpInfo($max = 25, $offset = 0, $sort = 'name', $order = 'asc', $phrase = null, $name = null, $user_id = null, $tenant_id = null, $timeframe = 'month', $start = null, $end = null)
    {
        $request = $this->listActivityRequest($max, $offset, $sort, $order, $phrase, $name, $user_id, $tenant_id, $timeframe, $start, $end);

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
                    if ('object' === '\SplFileObject') {
                        $content = $responseBody; //stream goes to serializer
                    } else {
                        $content = (string) $responseBody;
                    }

                    return [
                        ObjectSerializer::deserialize($content, 'object', []),
                        $response->getStatusCode(),
                        $response->getHeaders()
                    ];
                case 4XX:
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
                case 5XX:
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

            $returnType = 'object';
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
                        'object',
                        $e->getResponseHeaders()
                    );
                    $e->setResponseObject($data);
                    break;
                case 4XX:
                    $data = ObjectSerializer::deserialize(
                        $e->getResponseBody(),
                        '\OpenAPI\Client\Model\DefaultError',
                        $e->getResponseHeaders()
                    );
                    $e->setResponseObject($data);
                    break;
                case 5XX:
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
     * Operation listActivityAsync
     *
     * Retrieves Activity
     *
     * @param  int $max Maximum number of records to return, -1 can be used to fetch all records (optional, default to 25)
     * @param  int $offset Offset records, the number of records to skip, for paginating requests (optional, default to 0)
     * @param  string $sort Sort order, the name of the property to sort by (optional, default to 'name')
     * @param  string $order Sort direction, use &#39;desc&#39; to reverse sort (optional, default to 'asc')
     * @param  string $phrase Search phrase for partial matches on name or description (optional)
     * @param  string $name Filter by name, wildcard may be specified as %, eg. example-% (optional)
     * @param  int $user_id Filter by User ID (optional)
     * @param  float $tenant_id Filter by Tenant ID. Only available to the master account. (optional)
     * @param  string $timeframe Filter by a timeframe (ex - today, yesterday, week, month, 3months) (optional, default to 'month')
     * @param  \DateTime $start Filter by activity on or after a date(time). Default is 1 month prior (optional)
     * @param  \DateTime $end Filter by activity on or before a date(time). Default is current date (optional)
     *
     * @throws \InvalidArgumentException
     * @return \GuzzleHttp\Promise\PromiseInterface
     */
    public function listActivityAsync($max = 25, $offset = 0, $sort = 'name', $order = 'asc', $phrase = null, $name = null, $user_id = null, $tenant_id = null, $timeframe = 'month', $start = null, $end = null)
    {
        return $this->listActivityAsyncWithHttpInfo($max, $offset, $sort, $order, $phrase, $name, $user_id, $tenant_id, $timeframe, $start, $end)
            ->then(
                function ($response) {
                    return $response[0];
                }
            );
    }

    /**
     * Operation listActivityAsyncWithHttpInfo
     *
     * Retrieves Activity
     *
     * @param  int $max Maximum number of records to return, -1 can be used to fetch all records (optional, default to 25)
     * @param  int $offset Offset records, the number of records to skip, for paginating requests (optional, default to 0)
     * @param  string $sort Sort order, the name of the property to sort by (optional, default to 'name')
     * @param  string $order Sort direction, use &#39;desc&#39; to reverse sort (optional, default to 'asc')
     * @param  string $phrase Search phrase for partial matches on name or description (optional)
     * @param  string $name Filter by name, wildcard may be specified as %, eg. example-% (optional)
     * @param  int $user_id Filter by User ID (optional)
     * @param  float $tenant_id Filter by Tenant ID. Only available to the master account. (optional)
     * @param  string $timeframe Filter by a timeframe (ex - today, yesterday, week, month, 3months) (optional, default to 'month')
     * @param  \DateTime $start Filter by activity on or after a date(time). Default is 1 month prior (optional)
     * @param  \DateTime $end Filter by activity on or before a date(time). Default is current date (optional)
     *
     * @throws \InvalidArgumentException
     * @return \GuzzleHttp\Promise\PromiseInterface
     */
    public function listActivityAsyncWithHttpInfo($max = 25, $offset = 0, $sort = 'name', $order = 'asc', $phrase = null, $name = null, $user_id = null, $tenant_id = null, $timeframe = 'month', $start = null, $end = null)
    {
        $returnType = 'object';
        $request = $this->listActivityRequest($max, $offset, $sort, $order, $phrase, $name, $user_id, $tenant_id, $timeframe, $start, $end);

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
     * Create request for operation 'listActivity'
     *
     * @param  int $max Maximum number of records to return, -1 can be used to fetch all records (optional, default to 25)
     * @param  int $offset Offset records, the number of records to skip, for paginating requests (optional, default to 0)
     * @param  string $sort Sort order, the name of the property to sort by (optional, default to 'name')
     * @param  string $order Sort direction, use &#39;desc&#39; to reverse sort (optional, default to 'asc')
     * @param  string $phrase Search phrase for partial matches on name or description (optional)
     * @param  string $name Filter by name, wildcard may be specified as %, eg. example-% (optional)
     * @param  int $user_id Filter by User ID (optional)
     * @param  float $tenant_id Filter by Tenant ID. Only available to the master account. (optional)
     * @param  string $timeframe Filter by a timeframe (ex - today, yesterday, week, month, 3months) (optional, default to 'month')
     * @param  \DateTime $start Filter by activity on or after a date(time). Default is 1 month prior (optional)
     * @param  \DateTime $end Filter by activity on or before a date(time). Default is current date (optional)
     *
     * @throws \InvalidArgumentException
     * @return \GuzzleHttp\Psr7\Request
     */
    public function listActivityRequest($max = 25, $offset = 0, $sort = 'name', $order = 'asc', $phrase = null, $name = null, $user_id = null, $tenant_id = null, $timeframe = 'month', $start = null, $end = null)
    {

        $resourcePath = '/api/activity';
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
        if ($phrase !== null) {
            if('form' === 'form' && is_array($phrase)) {
                foreach($phrase as $key => $value) {
                    $queryParams[$key] = $value;
                }
            }
            else {
                $queryParams['phrase'] = $phrase;
            }
        }
        // query params
        if ($name !== null) {
            if('form' === 'form' && is_array($name)) {
                foreach($name as $key => $value) {
                    $queryParams[$key] = $value;
                }
            }
            else {
                $queryParams['name'] = $name;
            }
        }
        // query params
        if ($user_id !== null) {
            if('form' === 'form' && is_array($user_id)) {
                foreach($user_id as $key => $value) {
                    $queryParams[$key] = $value;
                }
            }
            else {
                $queryParams['userId'] = $user_id;
            }
        }
        // query params
        if ($tenant_id !== null) {
            if('form' === 'form' && is_array($tenant_id)) {
                foreach($tenant_id as $key => $value) {
                    $queryParams[$key] = $value;
                }
            }
            else {
                $queryParams['tenantId'] = $tenant_id;
            }
        }
        // query params
        if ($timeframe !== null) {
            if('form' === 'form' && is_array($timeframe)) {
                foreach($timeframe as $key => $value) {
                    $queryParams[$key] = $value;
                }
            }
            else {
                $queryParams['timeframe'] = $timeframe;
            }
        }
        // query params
        if ($start !== null) {
            if('form' === 'form' && is_array($start)) {
                foreach($start as $key => $value) {
                    $queryParams[$key] = $value;
                }
            }
            else {
                $queryParams['start'] = $start;
            }
        }
        // query params
        if ($end !== null) {
            if('form' === 'form' && is_array($end)) {
                foreach($end as $key => $value) {
                    $queryParams[$key] = $value;
                }
            }
            else {
                $queryParams['end'] = $end;
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