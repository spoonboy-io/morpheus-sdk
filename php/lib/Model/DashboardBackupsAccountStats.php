<?php
/**
 * DashboardBackupsAccountStats
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
 * DashboardBackupsAccountStats Class Doc Comment
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<TKey, TValue>
 * @template TKey int|null
 * @template TValue mixed|null  
 */
class DashboardBackupsAccountStats implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = 'dashboard_backups_accountStats';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'total_size_by_day' => 'float[]',
        'total_size_by_day7_days' => 'float[]',
        'formatted_total_size' => '\OpenAPI\Client\Model\DashboardBackupsAccountStatsFormattedTotalSize',
        'backup_count' => 'float',
        'total_size' => 'float',
        'success' => 'float',
        'failed' => 'float',
        'total_completed' => 'float',
        'last_seven_days' => '\OpenAPI\Client\Model\DashboardBackupsAccountStatsLastSevenDays',
        'avg_size' => 'float',
        'failed_rate' => 'float',
        'success_rate' => 'float',
        'next_fire_total' => 'float',
        'backup_day_count' => 'float[]',
        'backup_day_count_total' => 'float'
    ];

    /**
      * Array of property to format mappings. Used for (de)serialization
      *
      * @var string[]
      * @phpstan-var array<string, string|null>
      * @psalm-var array<string, string|null>
      */
    protected static $openAPIFormats = [
        'total_size_by_day' => null,
        'total_size_by_day7_days' => null,
        'formatted_total_size' => null,
        'backup_count' => null,
        'total_size' => null,
        'success' => null,
        'failed' => null,
        'total_completed' => null,
        'last_seven_days' => null,
        'avg_size' => null,
        'failed_rate' => null,
        'success_rate' => null,
        'next_fire_total' => null,
        'backup_day_count' => null,
        'backup_day_count_total' => null
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
        'total_size_by_day' => 'totalSizeByDay',
        'total_size_by_day7_days' => 'totalSizeByDay7Days',
        'formatted_total_size' => 'formattedTotalSize',
        'backup_count' => 'backupCount',
        'total_size' => 'totalSize',
        'success' => 'success',
        'failed' => 'failed',
        'total_completed' => 'totalCompleted',
        'last_seven_days' => 'lastSevenDays',
        'avg_size' => 'avgSize',
        'failed_rate' => 'failedRate',
        'success_rate' => 'successRate',
        'next_fire_total' => 'nextFireTotal',
        'backup_day_count' => 'backupDayCount',
        'backup_day_count_total' => 'backupDayCountTotal'
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'total_size_by_day' => 'setTotalSizeByDay',
        'total_size_by_day7_days' => 'setTotalSizeByDay7Days',
        'formatted_total_size' => 'setFormattedTotalSize',
        'backup_count' => 'setBackupCount',
        'total_size' => 'setTotalSize',
        'success' => 'setSuccess',
        'failed' => 'setFailed',
        'total_completed' => 'setTotalCompleted',
        'last_seven_days' => 'setLastSevenDays',
        'avg_size' => 'setAvgSize',
        'failed_rate' => 'setFailedRate',
        'success_rate' => 'setSuccessRate',
        'next_fire_total' => 'setNextFireTotal',
        'backup_day_count' => 'setBackupDayCount',
        'backup_day_count_total' => 'setBackupDayCountTotal'
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'total_size_by_day' => 'getTotalSizeByDay',
        'total_size_by_day7_days' => 'getTotalSizeByDay7Days',
        'formatted_total_size' => 'getFormattedTotalSize',
        'backup_count' => 'getBackupCount',
        'total_size' => 'getTotalSize',
        'success' => 'getSuccess',
        'failed' => 'getFailed',
        'total_completed' => 'getTotalCompleted',
        'last_seven_days' => 'getLastSevenDays',
        'avg_size' => 'getAvgSize',
        'failed_rate' => 'getFailedRate',
        'success_rate' => 'getSuccessRate',
        'next_fire_total' => 'getNextFireTotal',
        'backup_day_count' => 'getBackupDayCount',
        'backup_day_count_total' => 'getBackupDayCountTotal'
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
        $this->container['total_size_by_day'] = $data['total_size_by_day'] ?? null;
        $this->container['total_size_by_day7_days'] = $data['total_size_by_day7_days'] ?? null;
        $this->container['formatted_total_size'] = $data['formatted_total_size'] ?? null;
        $this->container['backup_count'] = $data['backup_count'] ?? null;
        $this->container['total_size'] = $data['total_size'] ?? null;
        $this->container['success'] = $data['success'] ?? null;
        $this->container['failed'] = $data['failed'] ?? null;
        $this->container['total_completed'] = $data['total_completed'] ?? null;
        $this->container['last_seven_days'] = $data['last_seven_days'] ?? null;
        $this->container['avg_size'] = $data['avg_size'] ?? null;
        $this->container['failed_rate'] = $data['failed_rate'] ?? null;
        $this->container['success_rate'] = $data['success_rate'] ?? null;
        $this->container['next_fire_total'] = $data['next_fire_total'] ?? null;
        $this->container['backup_day_count'] = $data['backup_day_count'] ?? null;
        $this->container['backup_day_count_total'] = $data['backup_day_count_total'] ?? null;
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
     * Gets total_size_by_day
     *
     * @return float[]|null
     */
    public function getTotalSizeByDay()
    {
        return $this->container['total_size_by_day'];
    }

    /**
     * Sets total_size_by_day
     *
     * @param float[]|null $total_size_by_day total_size_by_day
     *
     * @return self
     */
    public function setTotalSizeByDay($total_size_by_day)
    {
        $this->container['total_size_by_day'] = $total_size_by_day;

        return $this;
    }

    /**
     * Gets total_size_by_day7_days
     *
     * @return float[]|null
     */
    public function getTotalSizeByDay7Days()
    {
        return $this->container['total_size_by_day7_days'];
    }

    /**
     * Sets total_size_by_day7_days
     *
     * @param float[]|null $total_size_by_day7_days total_size_by_day7_days
     *
     * @return self
     */
    public function setTotalSizeByDay7Days($total_size_by_day7_days)
    {
        $this->container['total_size_by_day7_days'] = $total_size_by_day7_days;

        return $this;
    }

    /**
     * Gets formatted_total_size
     *
     * @return \OpenAPI\Client\Model\DashboardBackupsAccountStatsFormattedTotalSize|null
     */
    public function getFormattedTotalSize()
    {
        return $this->container['formatted_total_size'];
    }

    /**
     * Sets formatted_total_size
     *
     * @param \OpenAPI\Client\Model\DashboardBackupsAccountStatsFormattedTotalSize|null $formatted_total_size formatted_total_size
     *
     * @return self
     */
    public function setFormattedTotalSize($formatted_total_size)
    {
        $this->container['formatted_total_size'] = $formatted_total_size;

        return $this;
    }

    /**
     * Gets backup_count
     *
     * @return float|null
     */
    public function getBackupCount()
    {
        return $this->container['backup_count'];
    }

    /**
     * Sets backup_count
     *
     * @param float|null $backup_count backup_count
     *
     * @return self
     */
    public function setBackupCount($backup_count)
    {
        $this->container['backup_count'] = $backup_count;

        return $this;
    }

    /**
     * Gets total_size
     *
     * @return float|null
     */
    public function getTotalSize()
    {
        return $this->container['total_size'];
    }

    /**
     * Sets total_size
     *
     * @param float|null $total_size total_size
     *
     * @return self
     */
    public function setTotalSize($total_size)
    {
        $this->container['total_size'] = $total_size;

        return $this;
    }

    /**
     * Gets success
     *
     * @return float|null
     */
    public function getSuccess()
    {
        return $this->container['success'];
    }

    /**
     * Sets success
     *
     * @param float|null $success success
     *
     * @return self
     */
    public function setSuccess($success)
    {
        $this->container['success'] = $success;

        return $this;
    }

    /**
     * Gets failed
     *
     * @return float|null
     */
    public function getFailed()
    {
        return $this->container['failed'];
    }

    /**
     * Sets failed
     *
     * @param float|null $failed failed
     *
     * @return self
     */
    public function setFailed($failed)
    {
        $this->container['failed'] = $failed;

        return $this;
    }

    /**
     * Gets total_completed
     *
     * @return float|null
     */
    public function getTotalCompleted()
    {
        return $this->container['total_completed'];
    }

    /**
     * Sets total_completed
     *
     * @param float|null $total_completed total_completed
     *
     * @return self
     */
    public function setTotalCompleted($total_completed)
    {
        $this->container['total_completed'] = $total_completed;

        return $this;
    }

    /**
     * Gets last_seven_days
     *
     * @return \OpenAPI\Client\Model\DashboardBackupsAccountStatsLastSevenDays|null
     */
    public function getLastSevenDays()
    {
        return $this->container['last_seven_days'];
    }

    /**
     * Sets last_seven_days
     *
     * @param \OpenAPI\Client\Model\DashboardBackupsAccountStatsLastSevenDays|null $last_seven_days last_seven_days
     *
     * @return self
     */
    public function setLastSevenDays($last_seven_days)
    {
        $this->container['last_seven_days'] = $last_seven_days;

        return $this;
    }

    /**
     * Gets avg_size
     *
     * @return float|null
     */
    public function getAvgSize()
    {
        return $this->container['avg_size'];
    }

    /**
     * Sets avg_size
     *
     * @param float|null $avg_size avg_size
     *
     * @return self
     */
    public function setAvgSize($avg_size)
    {
        $this->container['avg_size'] = $avg_size;

        return $this;
    }

    /**
     * Gets failed_rate
     *
     * @return float|null
     */
    public function getFailedRate()
    {
        return $this->container['failed_rate'];
    }

    /**
     * Sets failed_rate
     *
     * @param float|null $failed_rate failed_rate
     *
     * @return self
     */
    public function setFailedRate($failed_rate)
    {
        $this->container['failed_rate'] = $failed_rate;

        return $this;
    }

    /**
     * Gets success_rate
     *
     * @return float|null
     */
    public function getSuccessRate()
    {
        return $this->container['success_rate'];
    }

    /**
     * Sets success_rate
     *
     * @param float|null $success_rate success_rate
     *
     * @return self
     */
    public function setSuccessRate($success_rate)
    {
        $this->container['success_rate'] = $success_rate;

        return $this;
    }

    /**
     * Gets next_fire_total
     *
     * @return float|null
     */
    public function getNextFireTotal()
    {
        return $this->container['next_fire_total'];
    }

    /**
     * Sets next_fire_total
     *
     * @param float|null $next_fire_total next_fire_total
     *
     * @return self
     */
    public function setNextFireTotal($next_fire_total)
    {
        $this->container['next_fire_total'] = $next_fire_total;

        return $this;
    }

    /**
     * Gets backup_day_count
     *
     * @return float[]|null
     */
    public function getBackupDayCount()
    {
        return $this->container['backup_day_count'];
    }

    /**
     * Sets backup_day_count
     *
     * @param float[]|null $backup_day_count backup_day_count
     *
     * @return self
     */
    public function setBackupDayCount($backup_day_count)
    {
        $this->container['backup_day_count'] = $backup_day_count;

        return $this;
    }

    /**
     * Gets backup_day_count_total
     *
     * @return float|null
     */
    public function getBackupDayCountTotal()
    {
        return $this->container['backup_day_count_total'];
    }

    /**
     * Sets backup_day_count_total
     *
     * @param float|null $backup_day_count_total backup_day_count_total
     *
     * @return self
     */
    public function setBackupDayCountTotal($backup_day_count_total)
    {
        $this->container['backup_day_count_total'] = $backup_day_count_total;

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

