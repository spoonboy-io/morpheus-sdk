<?php
/**
 * StorageServer
 *
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

namespace OpenAPI\Client\Model;

use \ArrayAccess;
use \OpenAPI\Client\ObjectSerializer;

/**
 * StorageServer Class Doc Comment
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<TKey, TValue>
 * @template TKey int|null
 * @template TValue mixed|null  
 */
class StorageServer implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = 'storageServer';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'id' => 'int',
        'name' => 'string',
        'type' => '\OpenAPI\Client\Model\InlineResponse20079LoadBalancerMonitorLoadBalancerType',
        'chassis' => 'string',
        'visibility' => 'string',
        'description' => 'string',
        'internal_id' => 'string',
        'external_id' => 'string',
        'service_url' => 'string',
        'service_host' => 'string',
        'service_path' => 'string',
        'service_token' => 'string',
        'service_token_hash' => 'string',
        'service_version' => 'string',
        'service_username' => 'string',
        'service_password' => 'string',
        'service_password_hash' => 'string',
        'internal_ip' => 'string',
        'external_ip' => 'string',
        'api_port' => 'string',
        'admin_port' => 'string',
        'config' => 'object',
        'ref_type' => 'string',
        'ref_id' => 'int',
        'category' => 'string',
        'server_vendor' => 'string',
        'server_model' => 'string',
        'serial_number' => 'string',
        'status' => 'string',
        'status_message' => 'string',
        'status_date' => '\DateTime',
        'error_message' => 'string',
        'max_storage' => 'string',
        'used_storage' => 'string',
        'disk_count' => 'string',
        'date_created' => '\DateTime',
        'last_updated' => '\DateTime',
        'enabled' => 'bool',
        'groups' => 'object[]',
        'host_groups' => 'object[]',
        'hosts' => 'object[]',
        'tenants' => 'object[]',
        'owner' => '\OpenAPI\Client\Model\InlineResponse20040AppDeployInstance',
        'credential' => 'object'
    ];

    /**
      * Array of property to format mappings. Used for (de)serialization
      *
      * @var string[]
      * @phpstan-var array<string, string|null>
      * @psalm-var array<string, string|null>
      */
    protected static $openAPIFormats = [
        'id' => 'int64',
        'name' => null,
        'type' => null,
        'chassis' => null,
        'visibility' => null,
        'description' => null,
        'internal_id' => null,
        'external_id' => null,
        'service_url' => null,
        'service_host' => null,
        'service_path' => null,
        'service_token' => null,
        'service_token_hash' => null,
        'service_version' => null,
        'service_username' => null,
        'service_password' => null,
        'service_password_hash' => null,
        'internal_ip' => null,
        'external_ip' => null,
        'api_port' => null,
        'admin_port' => null,
        'config' => null,
        'ref_type' => null,
        'ref_id' => 'int64',
        'category' => null,
        'server_vendor' => null,
        'server_model' => null,
        'serial_number' => null,
        'status' => null,
        'status_message' => null,
        'status_date' => 'date-time',
        'error_message' => null,
        'max_storage' => null,
        'used_storage' => null,
        'disk_count' => null,
        'date_created' => 'date-time',
        'last_updated' => 'date-time',
        'enabled' => null,
        'groups' => null,
        'host_groups' => null,
        'hosts' => null,
        'tenants' => null,
        'owner' => null,
        'credential' => null
    ];

    /**
     * Array of property to type mappings. Used for (de)serialization
     *
     * @return array
     */
    public static function openAPITypes()
    {
        return self::$openAPITypes;
    }

    /**
     * Array of property to format mappings. Used for (de)serialization
     *
     * @return array
     */
    public static function openAPIFormats()
    {
        return self::$openAPIFormats;
    }

    /**
     * Array of attributes where the key is the local name,
     * and the value is the original name
     *
     * @var string[]
     */
    protected static $attributeMap = [
        'id' => 'id',
        'name' => 'name',
        'type' => 'type',
        'chassis' => 'chassis',
        'visibility' => 'visibility',
        'description' => 'description',
        'internal_id' => 'internalId',
        'external_id' => 'externalId',
        'service_url' => 'serviceUrl',
        'service_host' => 'serviceHost',
        'service_path' => 'servicePath',
        'service_token' => 'serviceToken',
        'service_token_hash' => 'serviceTokenHash',
        'service_version' => 'serviceVersion',
        'service_username' => 'serviceUsername',
        'service_password' => 'servicePassword',
        'service_password_hash' => 'servicePasswordHash',
        'internal_ip' => 'internalIp',
        'external_ip' => 'externalIp',
        'api_port' => 'apiPort',
        'admin_port' => 'adminPort',
        'config' => 'config',
        'ref_type' => 'refType',
        'ref_id' => 'refId',
        'category' => 'category',
        'server_vendor' => 'serverVendor',
        'server_model' => 'serverModel',
        'serial_number' => 'serialNumber',
        'status' => 'status',
        'status_message' => 'statusMessage',
        'status_date' => 'statusDate',
        'error_message' => 'errorMessage',
        'max_storage' => 'maxStorage',
        'used_storage' => 'usedStorage',
        'disk_count' => 'diskCount',
        'date_created' => 'dateCreated',
        'last_updated' => 'lastUpdated',
        'enabled' => 'enabled',
        'groups' => 'groups',
        'host_groups' => 'hostGroups',
        'hosts' => 'hosts',
        'tenants' => 'tenants',
        'owner' => 'owner',
        'credential' => 'credential'
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'id' => 'setId',
        'name' => 'setName',
        'type' => 'setType',
        'chassis' => 'setChassis',
        'visibility' => 'setVisibility',
        'description' => 'setDescription',
        'internal_id' => 'setInternalId',
        'external_id' => 'setExternalId',
        'service_url' => 'setServiceUrl',
        'service_host' => 'setServiceHost',
        'service_path' => 'setServicePath',
        'service_token' => 'setServiceToken',
        'service_token_hash' => 'setServiceTokenHash',
        'service_version' => 'setServiceVersion',
        'service_username' => 'setServiceUsername',
        'service_password' => 'setServicePassword',
        'service_password_hash' => 'setServicePasswordHash',
        'internal_ip' => 'setInternalIp',
        'external_ip' => 'setExternalIp',
        'api_port' => 'setApiPort',
        'admin_port' => 'setAdminPort',
        'config' => 'setConfig',
        'ref_type' => 'setRefType',
        'ref_id' => 'setRefId',
        'category' => 'setCategory',
        'server_vendor' => 'setServerVendor',
        'server_model' => 'setServerModel',
        'serial_number' => 'setSerialNumber',
        'status' => 'setStatus',
        'status_message' => 'setStatusMessage',
        'status_date' => 'setStatusDate',
        'error_message' => 'setErrorMessage',
        'max_storage' => 'setMaxStorage',
        'used_storage' => 'setUsedStorage',
        'disk_count' => 'setDiskCount',
        'date_created' => 'setDateCreated',
        'last_updated' => 'setLastUpdated',
        'enabled' => 'setEnabled',
        'groups' => 'setGroups',
        'host_groups' => 'setHostGroups',
        'hosts' => 'setHosts',
        'tenants' => 'setTenants',
        'owner' => 'setOwner',
        'credential' => 'setCredential'
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'id' => 'getId',
        'name' => 'getName',
        'type' => 'getType',
        'chassis' => 'getChassis',
        'visibility' => 'getVisibility',
        'description' => 'getDescription',
        'internal_id' => 'getInternalId',
        'external_id' => 'getExternalId',
        'service_url' => 'getServiceUrl',
        'service_host' => 'getServiceHost',
        'service_path' => 'getServicePath',
        'service_token' => 'getServiceToken',
        'service_token_hash' => 'getServiceTokenHash',
        'service_version' => 'getServiceVersion',
        'service_username' => 'getServiceUsername',
        'service_password' => 'getServicePassword',
        'service_password_hash' => 'getServicePasswordHash',
        'internal_ip' => 'getInternalIp',
        'external_ip' => 'getExternalIp',
        'api_port' => 'getApiPort',
        'admin_port' => 'getAdminPort',
        'config' => 'getConfig',
        'ref_type' => 'getRefType',
        'ref_id' => 'getRefId',
        'category' => 'getCategory',
        'server_vendor' => 'getServerVendor',
        'server_model' => 'getServerModel',
        'serial_number' => 'getSerialNumber',
        'status' => 'getStatus',
        'status_message' => 'getStatusMessage',
        'status_date' => 'getStatusDate',
        'error_message' => 'getErrorMessage',
        'max_storage' => 'getMaxStorage',
        'used_storage' => 'getUsedStorage',
        'disk_count' => 'getDiskCount',
        'date_created' => 'getDateCreated',
        'last_updated' => 'getLastUpdated',
        'enabled' => 'getEnabled',
        'groups' => 'getGroups',
        'host_groups' => 'getHostGroups',
        'hosts' => 'getHosts',
        'tenants' => 'getTenants',
        'owner' => 'getOwner',
        'credential' => 'getCredential'
    ];

    /**
     * Array of attributes where the key is the local name,
     * and the value is the original name
     *
     * @return array
     */
    public static function attributeMap()
    {
        return self::$attributeMap;
    }

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @return array
     */
    public static function setters()
    {
        return self::$setters;
    }

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @return array
     */
    public static function getters()
    {
        return self::$getters;
    }

    /**
     * The original name of the model.
     *
     * @return string
     */
    public function getModelName()
    {
        return self::$openAPIModelName;
    }

    

    

    /**
     * Associative array for storing property values
     *
     * @var mixed[]
     */
    protected $container = [];

    /**
     * Constructor
     *
     * @param mixed[] $data Associated array of property values
     *                      initializing the model
     */
    public function __construct(array $data = null)
    {
        $this->container['id'] = $data['id'] ?? null;
        $this->container['name'] = $data['name'] ?? null;
        $this->container['type'] = $data['type'] ?? null;
        $this->container['chassis'] = $data['chassis'] ?? null;
        $this->container['visibility'] = $data['visibility'] ?? null;
        $this->container['description'] = $data['description'] ?? null;
        $this->container['internal_id'] = $data['internal_id'] ?? null;
        $this->container['external_id'] = $data['external_id'] ?? null;
        $this->container['service_url'] = $data['service_url'] ?? null;
        $this->container['service_host'] = $data['service_host'] ?? null;
        $this->container['service_path'] = $data['service_path'] ?? null;
        $this->container['service_token'] = $data['service_token'] ?? null;
        $this->container['service_token_hash'] = $data['service_token_hash'] ?? null;
        $this->container['service_version'] = $data['service_version'] ?? null;
        $this->container['service_username'] = $data['service_username'] ?? null;
        $this->container['service_password'] = $data['service_password'] ?? null;
        $this->container['service_password_hash'] = $data['service_password_hash'] ?? null;
        $this->container['internal_ip'] = $data['internal_ip'] ?? null;
        $this->container['external_ip'] = $data['external_ip'] ?? null;
        $this->container['api_port'] = $data['api_port'] ?? null;
        $this->container['admin_port'] = $data['admin_port'] ?? null;
        $this->container['config'] = $data['config'] ?? null;
        $this->container['ref_type'] = $data['ref_type'] ?? null;
        $this->container['ref_id'] = $data['ref_id'] ?? null;
        $this->container['category'] = $data['category'] ?? null;
        $this->container['server_vendor'] = $data['server_vendor'] ?? null;
        $this->container['server_model'] = $data['server_model'] ?? null;
        $this->container['serial_number'] = $data['serial_number'] ?? null;
        $this->container['status'] = $data['status'] ?? null;
        $this->container['status_message'] = $data['status_message'] ?? null;
        $this->container['status_date'] = $data['status_date'] ?? null;
        $this->container['error_message'] = $data['error_message'] ?? null;
        $this->container['max_storage'] = $data['max_storage'] ?? null;
        $this->container['used_storage'] = $data['used_storage'] ?? null;
        $this->container['disk_count'] = $data['disk_count'] ?? null;
        $this->container['date_created'] = $data['date_created'] ?? null;
        $this->container['last_updated'] = $data['last_updated'] ?? null;
        $this->container['enabled'] = $data['enabled'] ?? null;
        $this->container['groups'] = $data['groups'] ?? null;
        $this->container['host_groups'] = $data['host_groups'] ?? null;
        $this->container['hosts'] = $data['hosts'] ?? null;
        $this->container['tenants'] = $data['tenants'] ?? null;
        $this->container['owner'] = $data['owner'] ?? null;
        $this->container['credential'] = $data['credential'] ?? null;
    }

    /**
     * Show all the invalid properties with reasons.
     *
     * @return array invalid properties with reasons
     */
    public function listInvalidProperties()
    {
        $invalidProperties = [];

        return $invalidProperties;
    }

    /**
     * Validate all the properties in the model
     * return true if all passed
     *
     * @return bool True if all properties are valid
     */
    public function valid()
    {
        return count($this->listInvalidProperties()) === 0;
    }


    /**
     * Gets id
     *
     * @return int|null
     */
    public function getId()
    {
        return $this->container['id'];
    }

    /**
     * Sets id
     *
     * @param int|null $id id
     *
     * @return self
     */
    public function setId($id)
    {
        $this->container['id'] = $id;

        return $this;
    }

    /**
     * Gets name
     *
     * @return string|null
     */
    public function getName()
    {
        return $this->container['name'];
    }

    /**
     * Sets name
     *
     * @param string|null $name name
     *
     * @return self
     */
    public function setName($name)
    {
        $this->container['name'] = $name;

        return $this;
    }

    /**
     * Gets type
     *
     * @return \OpenAPI\Client\Model\InlineResponse20079LoadBalancerMonitorLoadBalancerType|null
     */
    public function getType()
    {
        return $this->container['type'];
    }

    /**
     * Sets type
     *
     * @param \OpenAPI\Client\Model\InlineResponse20079LoadBalancerMonitorLoadBalancerType|null $type type
     *
     * @return self
     */
    public function setType($type)
    {
        $this->container['type'] = $type;

        return $this;
    }

    /**
     * Gets chassis
     *
     * @return string|null
     */
    public function getChassis()
    {
        return $this->container['chassis'];
    }

    /**
     * Sets chassis
     *
     * @param string|null $chassis chassis
     *
     * @return self
     */
    public function setChassis($chassis)
    {
        $this->container['chassis'] = $chassis;

        return $this;
    }

    /**
     * Gets visibility
     *
     * @return string|null
     */
    public function getVisibility()
    {
        return $this->container['visibility'];
    }

    /**
     * Sets visibility
     *
     * @param string|null $visibility visibility
     *
     * @return self
     */
    public function setVisibility($visibility)
    {
        $this->container['visibility'] = $visibility;

        return $this;
    }

    /**
     * Gets description
     *
     * @return string|null
     */
    public function getDescription()
    {
        return $this->container['description'];
    }

    /**
     * Sets description
     *
     * @param string|null $description description
     *
     * @return self
     */
    public function setDescription($description)
    {
        $this->container['description'] = $description;

        return $this;
    }

    /**
     * Gets internal_id
     *
     * @return string|null
     */
    public function getInternalId()
    {
        return $this->container['internal_id'];
    }

    /**
     * Sets internal_id
     *
     * @param string|null $internal_id internal_id
     *
     * @return self
     */
    public function setInternalId($internal_id)
    {
        $this->container['internal_id'] = $internal_id;

        return $this;
    }

    /**
     * Gets external_id
     *
     * @return string|null
     */
    public function getExternalId()
    {
        return $this->container['external_id'];
    }

    /**
     * Sets external_id
     *
     * @param string|null $external_id external_id
     *
     * @return self
     */
    public function setExternalId($external_id)
    {
        $this->container['external_id'] = $external_id;

        return $this;
    }

    /**
     * Gets service_url
     *
     * @return string|null
     */
    public function getServiceUrl()
    {
        return $this->container['service_url'];
    }

    /**
     * Sets service_url
     *
     * @param string|null $service_url service_url
     *
     * @return self
     */
    public function setServiceUrl($service_url)
    {
        $this->container['service_url'] = $service_url;

        return $this;
    }

    /**
     * Gets service_host
     *
     * @return string|null
     */
    public function getServiceHost()
    {
        return $this->container['service_host'];
    }

    /**
     * Sets service_host
     *
     * @param string|null $service_host service_host
     *
     * @return self
     */
    public function setServiceHost($service_host)
    {
        $this->container['service_host'] = $service_host;

        return $this;
    }

    /**
     * Gets service_path
     *
     * @return string|null
     */
    public function getServicePath()
    {
        return $this->container['service_path'];
    }

    /**
     * Sets service_path
     *
     * @param string|null $service_path service_path
     *
     * @return self
     */
    public function setServicePath($service_path)
    {
        $this->container['service_path'] = $service_path;

        return $this;
    }

    /**
     * Gets service_token
     *
     * @return string|null
     */
    public function getServiceToken()
    {
        return $this->container['service_token'];
    }

    /**
     * Sets service_token
     *
     * @param string|null $service_token service_token
     *
     * @return self
     */
    public function setServiceToken($service_token)
    {
        $this->container['service_token'] = $service_token;

        return $this;
    }

    /**
     * Gets service_token_hash
     *
     * @return string|null
     */
    public function getServiceTokenHash()
    {
        return $this->container['service_token_hash'];
    }

    /**
     * Sets service_token_hash
     *
     * @param string|null $service_token_hash service_token_hash
     *
     * @return self
     */
    public function setServiceTokenHash($service_token_hash)
    {
        $this->container['service_token_hash'] = $service_token_hash;

        return $this;
    }

    /**
     * Gets service_version
     *
     * @return string|null
     */
    public function getServiceVersion()
    {
        return $this->container['service_version'];
    }

    /**
     * Sets service_version
     *
     * @param string|null $service_version service_version
     *
     * @return self
     */
    public function setServiceVersion($service_version)
    {
        $this->container['service_version'] = $service_version;

        return $this;
    }

    /**
     * Gets service_username
     *
     * @return string|null
     */
    public function getServiceUsername()
    {
        return $this->container['service_username'];
    }

    /**
     * Sets service_username
     *
     * @param string|null $service_username service_username
     *
     * @return self
     */
    public function setServiceUsername($service_username)
    {
        $this->container['service_username'] = $service_username;

        return $this;
    }

    /**
     * Gets service_password
     *
     * @return string|null
     */
    public function getServicePassword()
    {
        return $this->container['service_password'];
    }

    /**
     * Sets service_password
     *
     * @param string|null $service_password service_password
     *
     * @return self
     */
    public function setServicePassword($service_password)
    {
        $this->container['service_password'] = $service_password;

        return $this;
    }

    /**
     * Gets service_password_hash
     *
     * @return string|null
     */
    public function getServicePasswordHash()
    {
        return $this->container['service_password_hash'];
    }

    /**
     * Sets service_password_hash
     *
     * @param string|null $service_password_hash service_password_hash
     *
     * @return self
     */
    public function setServicePasswordHash($service_password_hash)
    {
        $this->container['service_password_hash'] = $service_password_hash;

        return $this;
    }

    /**
     * Gets internal_ip
     *
     * @return string|null
     */
    public function getInternalIp()
    {
        return $this->container['internal_ip'];
    }

    /**
     * Sets internal_ip
     *
     * @param string|null $internal_ip internal_ip
     *
     * @return self
     */
    public function setInternalIp($internal_ip)
    {
        $this->container['internal_ip'] = $internal_ip;

        return $this;
    }

    /**
     * Gets external_ip
     *
     * @return string|null
     */
    public function getExternalIp()
    {
        return $this->container['external_ip'];
    }

    /**
     * Sets external_ip
     *
     * @param string|null $external_ip external_ip
     *
     * @return self
     */
    public function setExternalIp($external_ip)
    {
        $this->container['external_ip'] = $external_ip;

        return $this;
    }

    /**
     * Gets api_port
     *
     * @return string|null
     */
    public function getApiPort()
    {
        return $this->container['api_port'];
    }

    /**
     * Sets api_port
     *
     * @param string|null $api_port api_port
     *
     * @return self
     */
    public function setApiPort($api_port)
    {
        $this->container['api_port'] = $api_port;

        return $this;
    }

    /**
     * Gets admin_port
     *
     * @return string|null
     */
    public function getAdminPort()
    {
        return $this->container['admin_port'];
    }

    /**
     * Sets admin_port
     *
     * @param string|null $admin_port admin_port
     *
     * @return self
     */
    public function setAdminPort($admin_port)
    {
        $this->container['admin_port'] = $admin_port;

        return $this;
    }

    /**
     * Gets config
     *
     * @return object|null
     */
    public function getConfig()
    {
        return $this->container['config'];
    }

    /**
     * Sets config
     *
     * @param object|null $config config
     *
     * @return self
     */
    public function setConfig($config)
    {
        $this->container['config'] = $config;

        return $this;
    }

    /**
     * Gets ref_type
     *
     * @return string|null
     */
    public function getRefType()
    {
        return $this->container['ref_type'];
    }

    /**
     * Sets ref_type
     *
     * @param string|null $ref_type ref_type
     *
     * @return self
     */
    public function setRefType($ref_type)
    {
        $this->container['ref_type'] = $ref_type;

        return $this;
    }

    /**
     * Gets ref_id
     *
     * @return int|null
     */
    public function getRefId()
    {
        return $this->container['ref_id'];
    }

    /**
     * Sets ref_id
     *
     * @param int|null $ref_id ref_id
     *
     * @return self
     */
    public function setRefId($ref_id)
    {
        $this->container['ref_id'] = $ref_id;

        return $this;
    }

    /**
     * Gets category
     *
     * @return string|null
     */
    public function getCategory()
    {
        return $this->container['category'];
    }

    /**
     * Sets category
     *
     * @param string|null $category category
     *
     * @return self
     */
    public function setCategory($category)
    {
        $this->container['category'] = $category;

        return $this;
    }

    /**
     * Gets server_vendor
     *
     * @return string|null
     */
    public function getServerVendor()
    {
        return $this->container['server_vendor'];
    }

    /**
     * Sets server_vendor
     *
     * @param string|null $server_vendor server_vendor
     *
     * @return self
     */
    public function setServerVendor($server_vendor)
    {
        $this->container['server_vendor'] = $server_vendor;

        return $this;
    }

    /**
     * Gets server_model
     *
     * @return string|null
     */
    public function getServerModel()
    {
        return $this->container['server_model'];
    }

    /**
     * Sets server_model
     *
     * @param string|null $server_model server_model
     *
     * @return self
     */
    public function setServerModel($server_model)
    {
        $this->container['server_model'] = $server_model;

        return $this;
    }

    /**
     * Gets serial_number
     *
     * @return string|null
     */
    public function getSerialNumber()
    {
        return $this->container['serial_number'];
    }

    /**
     * Sets serial_number
     *
     * @param string|null $serial_number serial_number
     *
     * @return self
     */
    public function setSerialNumber($serial_number)
    {
        $this->container['serial_number'] = $serial_number;

        return $this;
    }

    /**
     * Gets status
     *
     * @return string|null
     */
    public function getStatus()
    {
        return $this->container['status'];
    }

    /**
     * Sets status
     *
     * @param string|null $status status
     *
     * @return self
     */
    public function setStatus($status)
    {
        $this->container['status'] = $status;

        return $this;
    }

    /**
     * Gets status_message
     *
     * @return string|null
     */
    public function getStatusMessage()
    {
        return $this->container['status_message'];
    }

    /**
     * Sets status_message
     *
     * @param string|null $status_message status_message
     *
     * @return self
     */
    public function setStatusMessage($status_message)
    {
        $this->container['status_message'] = $status_message;

        return $this;
    }

    /**
     * Gets status_date
     *
     * @return \DateTime|null
     */
    public function getStatusDate()
    {
        return $this->container['status_date'];
    }

    /**
     * Sets status_date
     *
     * @param \DateTime|null $status_date status_date
     *
     * @return self
     */
    public function setStatusDate($status_date)
    {
        $this->container['status_date'] = $status_date;

        return $this;
    }

    /**
     * Gets error_message
     *
     * @return string|null
     */
    public function getErrorMessage()
    {
        return $this->container['error_message'];
    }

    /**
     * Sets error_message
     *
     * @param string|null $error_message error_message
     *
     * @return self
     */
    public function setErrorMessage($error_message)
    {
        $this->container['error_message'] = $error_message;

        return $this;
    }

    /**
     * Gets max_storage
     *
     * @return string|null
     */
    public function getMaxStorage()
    {
        return $this->container['max_storage'];
    }

    /**
     * Sets max_storage
     *
     * @param string|null $max_storage max_storage
     *
     * @return self
     */
    public function setMaxStorage($max_storage)
    {
        $this->container['max_storage'] = $max_storage;

        return $this;
    }

    /**
     * Gets used_storage
     *
     * @return string|null
     */
    public function getUsedStorage()
    {
        return $this->container['used_storage'];
    }

    /**
     * Sets used_storage
     *
     * @param string|null $used_storage used_storage
     *
     * @return self
     */
    public function setUsedStorage($used_storage)
    {
        $this->container['used_storage'] = $used_storage;

        return $this;
    }

    /**
     * Gets disk_count
     *
     * @return string|null
     */
    public function getDiskCount()
    {
        return $this->container['disk_count'];
    }

    /**
     * Sets disk_count
     *
     * @param string|null $disk_count disk_count
     *
     * @return self
     */
    public function setDiskCount($disk_count)
    {
        $this->container['disk_count'] = $disk_count;

        return $this;
    }

    /**
     * Gets date_created
     *
     * @return \DateTime|null
     */
    public function getDateCreated()
    {
        return $this->container['date_created'];
    }

    /**
     * Sets date_created
     *
     * @param \DateTime|null $date_created date_created
     *
     * @return self
     */
    public function setDateCreated($date_created)
    {
        $this->container['date_created'] = $date_created;

        return $this;
    }

    /**
     * Gets last_updated
     *
     * @return \DateTime|null
     */
    public function getLastUpdated()
    {
        return $this->container['last_updated'];
    }

    /**
     * Sets last_updated
     *
     * @param \DateTime|null $last_updated last_updated
     *
     * @return self
     */
    public function setLastUpdated($last_updated)
    {
        $this->container['last_updated'] = $last_updated;

        return $this;
    }

    /**
     * Gets enabled
     *
     * @return bool|null
     */
    public function getEnabled()
    {
        return $this->container['enabled'];
    }

    /**
     * Sets enabled
     *
     * @param bool|null $enabled enabled
     *
     * @return self
     */
    public function setEnabled($enabled)
    {
        $this->container['enabled'] = $enabled;

        return $this;
    }

    /**
     * Gets groups
     *
     * @return object[]|null
     */
    public function getGroups()
    {
        return $this->container['groups'];
    }

    /**
     * Sets groups
     *
     * @param object[]|null $groups groups
     *
     * @return self
     */
    public function setGroups($groups)
    {
        $this->container['groups'] = $groups;

        return $this;
    }

    /**
     * Gets host_groups
     *
     * @return object[]|null
     */
    public function getHostGroups()
    {
        return $this->container['host_groups'];
    }

    /**
     * Sets host_groups
     *
     * @param object[]|null $host_groups host_groups
     *
     * @return self
     */
    public function setHostGroups($host_groups)
    {
        $this->container['host_groups'] = $host_groups;

        return $this;
    }

    /**
     * Gets hosts
     *
     * @return object[]|null
     */
    public function getHosts()
    {
        return $this->container['hosts'];
    }

    /**
     * Sets hosts
     *
     * @param object[]|null $hosts hosts
     *
     * @return self
     */
    public function setHosts($hosts)
    {
        $this->container['hosts'] = $hosts;

        return $this;
    }

    /**
     * Gets tenants
     *
     * @return object[]|null
     */
    public function getTenants()
    {
        return $this->container['tenants'];
    }

    /**
     * Sets tenants
     *
     * @param object[]|null $tenants tenants
     *
     * @return self
     */
    public function setTenants($tenants)
    {
        $this->container['tenants'] = $tenants;

        return $this;
    }

    /**
     * Gets owner
     *
     * @return \OpenAPI\Client\Model\InlineResponse20040AppDeployInstance|null
     */
    public function getOwner()
    {
        return $this->container['owner'];
    }

    /**
     * Sets owner
     *
     * @param \OpenAPI\Client\Model\InlineResponse20040AppDeployInstance|null $owner owner
     *
     * @return self
     */
    public function setOwner($owner)
    {
        $this->container['owner'] = $owner;

        return $this;
    }

    /**
     * Gets credential
     *
     * @return object|null
     */
    public function getCredential()
    {
        return $this->container['credential'];
    }

    /**
     * Sets credential
     *
     * @param object|null $credential credential
     *
     * @return self
     */
    public function setCredential($credential)
    {
        $this->container['credential'] = $credential;

        return $this;
    }
    /**
     * Returns true if offset exists. False otherwise.
     *
     * @param integer $offset Offset
     *
     * @return boolean
     */
    public function offsetExists($offset)
    {
        return isset($this->container[$offset]);
    }

    /**
     * Gets offset.
     *
     * @param integer $offset Offset
     *
     * @return mixed|null
     */
    public function offsetGet($offset)
    {
        return $this->container[$offset] ?? null;
    }

    /**
     * Sets value based on offset.
     *
     * @param int|null $offset Offset
     * @param mixed    $value  Value to be set
     *
     * @return void
     */
    public function offsetSet($offset, $value)
    {
        if (is_null($offset)) {
            $this->container[] = $value;
        } else {
            $this->container[$offset] = $value;
        }
    }

    /**
     * Unsets offset.
     *
     * @param integer $offset Offset
     *
     * @return void
     */
    public function offsetUnset($offset)
    {
        unset($this->container[$offset]);
    }

    /**
     * Serializes the object to a value that can be serialized natively by json_encode().
     * @link https://www.php.net/manual/en/jsonserializable.jsonserialize.php
     *
     * @return mixed Returns data which can be serialized by json_encode(), which is a value
     * of any type other than a resource.
     */
    public function jsonSerialize()
    {
       return ObjectSerializer::sanitizeForSerialization($this);
    }

    /**
     * Gets the string presentation of the object
     *
     * @return string
     */
    public function __toString()
    {
        return json_encode(
            ObjectSerializer::sanitizeForSerialization($this),
            JSON_PRETTY_PRINT
        );
    }

    /**
     * Gets a header-safe presentation of the object
     *
     * @return string
     */
    public function toHeaderValue()
    {
        return json_encode(ObjectSerializer::sanitizeForSerialization($this));
    }
}

