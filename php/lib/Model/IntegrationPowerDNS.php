<?php
/**
 * IntegrationPowerDNS
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
 * IntegrationPowerDNS Class Doc Comment
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<TKey, TValue>
 * @template TKey int|null
 * @template TValue mixed|null  
 */
class IntegrationPowerDNS implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = 'integrationPowerDNS';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'id' => 'int',
        'name' => 'string',
        'enabled' => 'bool',
        'type' => 'string',
        'integration_type' => '\OpenAPI\Client\Model\InlineResponse20079LoadBalancerMonitorLoadBalancerType',
        'url' => 'string',
        'version' => 'string',
        'service_flag' => 'bool',
        'is_plugin' => 'bool',
        'config' => 'object',
        'status' => 'string',
        'status_date' => '\DateTime',
        'status_message' => 'string',
        'last_sync' => 'string',
        'last_sync_duration' => 'string',
        'credential' => '\OpenAPI\Client\Model\Creds'
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
        'enabled' => null,
        'type' => null,
        'integration_type' => null,
        'url' => null,
        'version' => null,
        'service_flag' => null,
        'is_plugin' => null,
        'config' => null,
        'status' => null,
        'status_date' => 'date-time',
        'status_message' => null,
        'last_sync' => null,
        'last_sync_duration' => null,
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
        'enabled' => 'enabled',
        'type' => 'type',
        'integration_type' => 'integrationType',
        'url' => 'url',
        'version' => 'version',
        'service_flag' => 'serviceFlag',
        'is_plugin' => 'isPlugin',
        'config' => 'config',
        'status' => 'status',
        'status_date' => 'statusDate',
        'status_message' => 'statusMessage',
        'last_sync' => 'lastSync',
        'last_sync_duration' => 'lastSyncDuration',
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
        'enabled' => 'setEnabled',
        'type' => 'setType',
        'integration_type' => 'setIntegrationType',
        'url' => 'setUrl',
        'version' => 'setVersion',
        'service_flag' => 'setServiceFlag',
        'is_plugin' => 'setIsPlugin',
        'config' => 'setConfig',
        'status' => 'setStatus',
        'status_date' => 'setStatusDate',
        'status_message' => 'setStatusMessage',
        'last_sync' => 'setLastSync',
        'last_sync_duration' => 'setLastSyncDuration',
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
        'enabled' => 'getEnabled',
        'type' => 'getType',
        'integration_type' => 'getIntegrationType',
        'url' => 'getUrl',
        'version' => 'getVersion',
        'service_flag' => 'getServiceFlag',
        'is_plugin' => 'getIsPlugin',
        'config' => 'getConfig',
        'status' => 'getStatus',
        'status_date' => 'getStatusDate',
        'status_message' => 'getStatusMessage',
        'last_sync' => 'getLastSync',
        'last_sync_duration' => 'getLastSyncDuration',
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

    const TYPE_POWER_DNS = 'powerDns';
    

    
    /**
     * Gets allowable values of the enum
     *
     * @return string[]
     */
    public function getTypeAllowableValues()
    {
        return [
            self::TYPE_POWER_DNS,
        ];
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
        $this->container['enabled'] = $data['enabled'] ?? null;
        $this->container['type'] = $data['type'] ?? null;
        $this->container['integration_type'] = $data['integration_type'] ?? null;
        $this->container['url'] = $data['url'] ?? null;
        $this->container['version'] = $data['version'] ?? null;
        $this->container['service_flag'] = $data['service_flag'] ?? null;
        $this->container['is_plugin'] = $data['is_plugin'] ?? null;
        $this->container['config'] = $data['config'] ?? null;
        $this->container['status'] = $data['status'] ?? null;
        $this->container['status_date'] = $data['status_date'] ?? null;
        $this->container['status_message'] = $data['status_message'] ?? null;
        $this->container['last_sync'] = $data['last_sync'] ?? null;
        $this->container['last_sync_duration'] = $data['last_sync_duration'] ?? null;
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

        $allowedValues = $this->getTypeAllowableValues();
        if (!is_null($this->container['type']) && !in_array($this->container['type'], $allowedValues, true)) {
            $invalidProperties[] = sprintf(
                "invalid value '%s' for 'type', must be one of '%s'",
                $this->container['type'],
                implode("', '", $allowedValues)
            );
        }

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
     * Gets type
     *
     * @return string|null
     */
    public function getType()
    {
        return $this->container['type'];
    }

    /**
     * Sets type
     *
     * @param string|null $type type
     *
     * @return self
     */
    public function setType($type)
    {
        $allowedValues = $this->getTypeAllowableValues();
        if (!is_null($type) && !in_array($type, $allowedValues, true)) {
            throw new \InvalidArgumentException(
                sprintf(
                    "Invalid value '%s' for 'type', must be one of '%s'",
                    $type,
                    implode("', '", $allowedValues)
                )
            );
        }
        $this->container['type'] = $type;

        return $this;
    }

    /**
     * Gets integration_type
     *
     * @return \OpenAPI\Client\Model\InlineResponse20079LoadBalancerMonitorLoadBalancerType|null
     */
    public function getIntegrationType()
    {
        return $this->container['integration_type'];
    }

    /**
     * Sets integration_type
     *
     * @param \OpenAPI\Client\Model\InlineResponse20079LoadBalancerMonitorLoadBalancerType|null $integration_type integration_type
     *
     * @return self
     */
    public function setIntegrationType($integration_type)
    {
        $this->container['integration_type'] = $integration_type;

        return $this;
    }

    /**
     * Gets url
     *
     * @return string|null
     */
    public function getUrl()
    {
        return $this->container['url'];
    }

    /**
     * Sets url
     *
     * @param string|null $url url
     *
     * @return self
     */
    public function setUrl($url)
    {
        $this->container['url'] = $url;

        return $this;
    }

    /**
     * Gets version
     *
     * @return string|null
     */
    public function getVersion()
    {
        return $this->container['version'];
    }

    /**
     * Sets version
     *
     * @param string|null $version version
     *
     * @return self
     */
    public function setVersion($version)
    {
        $this->container['version'] = $version;

        return $this;
    }

    /**
     * Gets service_flag
     *
     * @return bool|null
     */
    public function getServiceFlag()
    {
        return $this->container['service_flag'];
    }

    /**
     * Sets service_flag
     *
     * @param bool|null $service_flag service_flag
     *
     * @return self
     */
    public function setServiceFlag($service_flag)
    {
        $this->container['service_flag'] = $service_flag;

        return $this;
    }

    /**
     * Gets is_plugin
     *
     * @return bool|null
     */
    public function getIsPlugin()
    {
        return $this->container['is_plugin'];
    }

    /**
     * Sets is_plugin
     *
     * @param bool|null $is_plugin is_plugin
     *
     * @return self
     */
    public function setIsPlugin($is_plugin)
    {
        $this->container['is_plugin'] = $is_plugin;

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
     * Gets last_sync
     *
     * @return string|null
     */
    public function getLastSync()
    {
        return $this->container['last_sync'];
    }

    /**
     * Sets last_sync
     *
     * @param string|null $last_sync last_sync
     *
     * @return self
     */
    public function setLastSync($last_sync)
    {
        $this->container['last_sync'] = $last_sync;

        return $this;
    }

    /**
     * Gets last_sync_duration
     *
     * @return string|null
     */
    public function getLastSyncDuration()
    {
        return $this->container['last_sync_duration'];
    }

    /**
     * Sets last_sync_duration
     *
     * @param string|null $last_sync_duration last_sync_duration
     *
     * @return self
     */
    public function setLastSyncDuration($last_sync_duration)
    {
        $this->container['last_sync_duration'] = $last_sync_duration;

        return $this;
    }

    /**
     * Gets credential
     *
     * @return \OpenAPI\Client\Model\Creds|null
     */
    public function getCredential()
    {
        return $this->container['credential'];
    }

    /**
     * Sets credential
     *
     * @param \OpenAPI\Client\Model\Creds|null $credential credential
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


