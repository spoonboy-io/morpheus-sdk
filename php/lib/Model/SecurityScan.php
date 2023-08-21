<?php
/**
 * SecurityScan
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
 * SecurityScan Class Doc Comment
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<TKey, TValue>
 * @template TKey int|null
 * @template TValue mixed|null  
 */
class SecurityScan implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = 'securityScan';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'id' => 'int',
        'security_package' => '\OpenAPI\Client\Model\SecurityScanSecurityPackage',
        'status' => 'string',
        'scan_date' => '\DateTime',
        'scan_duration' => 'int',
        'test_count' => 'int',
        'run_count' => 'int',
        'pass_count' => 'int',
        'fail_count' => 'int',
        'other_count' => 'int',
        'scan_score' => 'float',
        'external_id' => 'string',
        'created_by' => 'string',
        'updated_by' => 'string',
        'date_created' => '\DateTime',
        'last_updated' => '\DateTime',
        'results' => 'object'
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
        'security_package' => null,
        'status' => null,
        'scan_date' => 'date-time',
        'scan_duration' => 'int64',
        'test_count' => 'int64',
        'run_count' => 'int64',
        'pass_count' => 'int64',
        'fail_count' => 'int64',
        'other_count' => 'int64',
        'scan_score' => null,
        'external_id' => null,
        'created_by' => null,
        'updated_by' => null,
        'date_created' => 'date-time',
        'last_updated' => 'date-time',
        'results' => null
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
        'security_package' => 'securityPackage',
        'status' => 'status',
        'scan_date' => 'scanDate',
        'scan_duration' => 'scanDuration',
        'test_count' => 'testCount',
        'run_count' => 'runCount',
        'pass_count' => 'passCount',
        'fail_count' => 'failCount',
        'other_count' => 'otherCount',
        'scan_score' => 'scanScore',
        'external_id' => 'externalId',
        'created_by' => 'createdBy',
        'updated_by' => 'updatedBy',
        'date_created' => 'dateCreated',
        'last_updated' => 'lastUpdated',
        'results' => 'results'
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'id' => 'setId',
        'security_package' => 'setSecurityPackage',
        'status' => 'setStatus',
        'scan_date' => 'setScanDate',
        'scan_duration' => 'setScanDuration',
        'test_count' => 'setTestCount',
        'run_count' => 'setRunCount',
        'pass_count' => 'setPassCount',
        'fail_count' => 'setFailCount',
        'other_count' => 'setOtherCount',
        'scan_score' => 'setScanScore',
        'external_id' => 'setExternalId',
        'created_by' => 'setCreatedBy',
        'updated_by' => 'setUpdatedBy',
        'date_created' => 'setDateCreated',
        'last_updated' => 'setLastUpdated',
        'results' => 'setResults'
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'id' => 'getId',
        'security_package' => 'getSecurityPackage',
        'status' => 'getStatus',
        'scan_date' => 'getScanDate',
        'scan_duration' => 'getScanDuration',
        'test_count' => 'getTestCount',
        'run_count' => 'getRunCount',
        'pass_count' => 'getPassCount',
        'fail_count' => 'getFailCount',
        'other_count' => 'getOtherCount',
        'scan_score' => 'getScanScore',
        'external_id' => 'getExternalId',
        'created_by' => 'getCreatedBy',
        'updated_by' => 'getUpdatedBy',
        'date_created' => 'getDateCreated',
        'last_updated' => 'getLastUpdated',
        'results' => 'getResults'
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
        $this->container['security_package'] = $data['security_package'] ?? null;
        $this->container['status'] = $data['status'] ?? null;
        $this->container['scan_date'] = $data['scan_date'] ?? null;
        $this->container['scan_duration'] = $data['scan_duration'] ?? null;
        $this->container['test_count'] = $data['test_count'] ?? null;
        $this->container['run_count'] = $data['run_count'] ?? null;
        $this->container['pass_count'] = $data['pass_count'] ?? null;
        $this->container['fail_count'] = $data['fail_count'] ?? null;
        $this->container['other_count'] = $data['other_count'] ?? null;
        $this->container['scan_score'] = $data['scan_score'] ?? null;
        $this->container['external_id'] = $data['external_id'] ?? null;
        $this->container['created_by'] = $data['created_by'] ?? null;
        $this->container['updated_by'] = $data['updated_by'] ?? null;
        $this->container['date_created'] = $data['date_created'] ?? null;
        $this->container['last_updated'] = $data['last_updated'] ?? null;
        $this->container['results'] = $data['results'] ?? null;
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
     * Gets security_package
     *
     * @return \OpenAPI\Client\Model\SecurityScanSecurityPackage|null
     */
    public function getSecurityPackage()
    {
        return $this->container['security_package'];
    }

    /**
     * Sets security_package
     *
     * @param \OpenAPI\Client\Model\SecurityScanSecurityPackage|null $security_package security_package
     *
     * @return self
     */
    public function setSecurityPackage($security_package)
    {
        $this->container['security_package'] = $security_package;

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
     * Gets scan_date
     *
     * @return \DateTime|null
     */
    public function getScanDate()
    {
        return $this->container['scan_date'];
    }

    /**
     * Sets scan_date
     *
     * @param \DateTime|null $scan_date scan_date
     *
     * @return self
     */
    public function setScanDate($scan_date)
    {
        $this->container['scan_date'] = $scan_date;

        return $this;
    }

    /**
     * Gets scan_duration
     *
     * @return int|null
     */
    public function getScanDuration()
    {
        return $this->container['scan_duration'];
    }

    /**
     * Sets scan_duration
     *
     * @param int|null $scan_duration scan_duration
     *
     * @return self
     */
    public function setScanDuration($scan_duration)
    {
        $this->container['scan_duration'] = $scan_duration;

        return $this;
    }

    /**
     * Gets test_count
     *
     * @return int|null
     */
    public function getTestCount()
    {
        return $this->container['test_count'];
    }

    /**
     * Sets test_count
     *
     * @param int|null $test_count test_count
     *
     * @return self
     */
    public function setTestCount($test_count)
    {
        $this->container['test_count'] = $test_count;

        return $this;
    }

    /**
     * Gets run_count
     *
     * @return int|null
     */
    public function getRunCount()
    {
        return $this->container['run_count'];
    }

    /**
     * Sets run_count
     *
     * @param int|null $run_count run_count
     *
     * @return self
     */
    public function setRunCount($run_count)
    {
        $this->container['run_count'] = $run_count;

        return $this;
    }

    /**
     * Gets pass_count
     *
     * @return int|null
     */
    public function getPassCount()
    {
        return $this->container['pass_count'];
    }

    /**
     * Sets pass_count
     *
     * @param int|null $pass_count pass_count
     *
     * @return self
     */
    public function setPassCount($pass_count)
    {
        $this->container['pass_count'] = $pass_count;

        return $this;
    }

    /**
     * Gets fail_count
     *
     * @return int|null
     */
    public function getFailCount()
    {
        return $this->container['fail_count'];
    }

    /**
     * Sets fail_count
     *
     * @param int|null $fail_count fail_count
     *
     * @return self
     */
    public function setFailCount($fail_count)
    {
        $this->container['fail_count'] = $fail_count;

        return $this;
    }

    /**
     * Gets other_count
     *
     * @return int|null
     */
    public function getOtherCount()
    {
        return $this->container['other_count'];
    }

    /**
     * Sets other_count
     *
     * @param int|null $other_count other_count
     *
     * @return self
     */
    public function setOtherCount($other_count)
    {
        $this->container['other_count'] = $other_count;

        return $this;
    }

    /**
     * Gets scan_score
     *
     * @return float|null
     */
    public function getScanScore()
    {
        return $this->container['scan_score'];
    }

    /**
     * Sets scan_score
     *
     * @param float|null $scan_score scan_score
     *
     * @return self
     */
    public function setScanScore($scan_score)
    {
        $this->container['scan_score'] = $scan_score;

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
     * Gets created_by
     *
     * @return string|null
     */
    public function getCreatedBy()
    {
        return $this->container['created_by'];
    }

    /**
     * Sets created_by
     *
     * @param string|null $created_by created_by
     *
     * @return self
     */
    public function setCreatedBy($created_by)
    {
        $this->container['created_by'] = $created_by;

        return $this;
    }

    /**
     * Gets updated_by
     *
     * @return string|null
     */
    public function getUpdatedBy()
    {
        return $this->container['updated_by'];
    }

    /**
     * Sets updated_by
     *
     * @param string|null $updated_by updated_by
     *
     * @return self
     */
    public function setUpdatedBy($updated_by)
    {
        $this->container['updated_by'] = $updated_by;

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
     * Gets results
     *
     * @return object|null
     */
    public function getResults()
    {
        return $this->container['results'];
    }

    /**
     * Sets results
     *
     * @param object|null $results Results Summary (only returned when using query parameter results=true)
     *
     * @return self
     */
    public function setResults($results)
    {
        $this->container['results'] = $results;

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

