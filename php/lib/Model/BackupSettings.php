<?php
/**
 * BackupSettings
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
 * BackupSettings Class Doc Comment
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<TKey, TValue>
 * @template TKey int|null
 * @template TValue mixed|null  
 */
class BackupSettings implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = 'backupSettings';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'backups_enabled' => 'bool',
        'create_backups' => 'bool',
        'backup_appliance' => 'bool',
        'default_storage_bucket' => '\OpenAPI\Client\Model\InlineResponse20040AppDeployInstance',
        'default_schedule' => '\OpenAPI\Client\Model\InlineResponse20079LoadBalancerMonitorLoadBalancerType',
        'retention_count' => 'int'
    ];

    /**
      * Array of property to format mappings. Used for (de)serialization
      *
      * @var string[]
      * @phpstan-var array<string, string|null>
      * @psalm-var array<string, string|null>
      */
    protected static $openAPIFormats = [
        'backups_enabled' => null,
        'create_backups' => null,
        'backup_appliance' => null,
        'default_storage_bucket' => null,
        'default_schedule' => null,
        'retention_count' => 'int64'
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
        'backups_enabled' => 'backupsEnabled',
        'create_backups' => 'createBackups',
        'backup_appliance' => 'backupAppliance',
        'default_storage_bucket' => 'defaultStorageBucket',
        'default_schedule' => 'defaultSchedule',
        'retention_count' => 'retentionCount'
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'backups_enabled' => 'setBackupsEnabled',
        'create_backups' => 'setCreateBackups',
        'backup_appliance' => 'setBackupAppliance',
        'default_storage_bucket' => 'setDefaultStorageBucket',
        'default_schedule' => 'setDefaultSchedule',
        'retention_count' => 'setRetentionCount'
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'backups_enabled' => 'getBackupsEnabled',
        'create_backups' => 'getCreateBackups',
        'backup_appliance' => 'getBackupAppliance',
        'default_storage_bucket' => 'getDefaultStorageBucket',
        'default_schedule' => 'getDefaultSchedule',
        'retention_count' => 'getRetentionCount'
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
        $this->container['backups_enabled'] = $data['backups_enabled'] ?? null;
        $this->container['create_backups'] = $data['create_backups'] ?? null;
        $this->container['backup_appliance'] = $data['backup_appliance'] ?? null;
        $this->container['default_storage_bucket'] = $data['default_storage_bucket'] ?? null;
        $this->container['default_schedule'] = $data['default_schedule'] ?? null;
        $this->container['retention_count'] = $data['retention_count'] ?? null;
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
     * Gets backups_enabled
     *
     * @return bool|null
     */
    public function getBackupsEnabled()
    {
        return $this->container['backups_enabled'];
    }

    /**
     * Sets backups_enabled
     *
     * @param bool|null $backups_enabled backups_enabled
     *
     * @return self
     */
    public function setBackupsEnabled($backups_enabled)
    {
        $this->container['backups_enabled'] = $backups_enabled;

        return $this;
    }

    /**
     * Gets create_backups
     *
     * @return bool|null
     */
    public function getCreateBackups()
    {
        return $this->container['create_backups'];
    }

    /**
     * Sets create_backups
     *
     * @param bool|null $create_backups create_backups
     *
     * @return self
     */
    public function setCreateBackups($create_backups)
    {
        $this->container['create_backups'] = $create_backups;

        return $this;
    }

    /**
     * Gets backup_appliance
     *
     * @return bool|null
     */
    public function getBackupAppliance()
    {
        return $this->container['backup_appliance'];
    }

    /**
     * Sets backup_appliance
     *
     * @param bool|null $backup_appliance backup_appliance
     *
     * @return self
     */
    public function setBackupAppliance($backup_appliance)
    {
        $this->container['backup_appliance'] = $backup_appliance;

        return $this;
    }

    /**
     * Gets default_storage_bucket
     *
     * @return \OpenAPI\Client\Model\InlineResponse20040AppDeployInstance|null
     */
    public function getDefaultStorageBucket()
    {
        return $this->container['default_storage_bucket'];
    }

    /**
     * Sets default_storage_bucket
     *
     * @param \OpenAPI\Client\Model\InlineResponse20040AppDeployInstance|null $default_storage_bucket default_storage_bucket
     *
     * @return self
     */
    public function setDefaultStorageBucket($default_storage_bucket)
    {
        $this->container['default_storage_bucket'] = $default_storage_bucket;

        return $this;
    }

    /**
     * Gets default_schedule
     *
     * @return \OpenAPI\Client\Model\InlineResponse20079LoadBalancerMonitorLoadBalancerType|null
     */
    public function getDefaultSchedule()
    {
        return $this->container['default_schedule'];
    }

    /**
     * Sets default_schedule
     *
     * @param \OpenAPI\Client\Model\InlineResponse20079LoadBalancerMonitorLoadBalancerType|null $default_schedule default_schedule
     *
     * @return self
     */
    public function setDefaultSchedule($default_schedule)
    {
        $this->container['default_schedule'] = $default_schedule;

        return $this;
    }

    /**
     * Gets retention_count
     *
     * @return int|null
     */
    public function getRetentionCount()
    {
        return $this->container['retention_count'];
    }

    /**
     * Sets retention_count
     *
     * @param int|null $retention_count retention_count
     *
     * @return self
     */
    public function setRetentionCount($retention_count)
    {
        $this->container['retention_count'] = $retention_count;

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


