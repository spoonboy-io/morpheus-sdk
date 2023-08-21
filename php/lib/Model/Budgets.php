<?php
/**
 * Budgets
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
 * Budgets Class Doc Comment
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<TKey, TValue>
 * @template TKey int|null
 * @template TValue mixed|null  
 */
class Budgets implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = 'budgets';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'id' => 'int',
        'name' => 'string',
        'description' => 'string',
        'account' => '\OpenAPI\Client\Model\InlineResponse20040AppDeployInstance',
        'enabled' => 'bool',
        'ref_scope' => 'string',
        'ref_type' => 'string',
        'ref_id' => 'int',
        'ref_name' => 'string',
        'period' => 'string',
        'year' => 'string',
        'resource_type' => 'string',
        'timezone' => 'string',
        'start_date' => '\DateTime',
        'end_date' => '\DateTime',
        'interval' => 'string',
        'costs' => 'int[]',
        'is_fiscal' => 'bool',
        'average_cost' => 'int',
        'total_cost' => 'int',
        'currency' => 'string',
        'rollover' => 'bool',
        'warning_limit' => 'string',
        'over_limit' => 'string',
        'external_id' => 'string',
        'internal_id' => 'string',
        'created_by_id' => 'int',
        'created_by_name' => 'string',
        'updated_by_id' => 'string',
        'updated_by_name' => 'string',
        'date_created' => '\DateTime',
        'last_updated' => '\DateTime'
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
        'description' => null,
        'account' => null,
        'enabled' => null,
        'ref_scope' => null,
        'ref_type' => null,
        'ref_id' => 'int64',
        'ref_name' => null,
        'period' => null,
        'year' => null,
        'resource_type' => null,
        'timezone' => null,
        'start_date' => 'date-time',
        'end_date' => 'date-time',
        'interval' => null,
        'costs' => 'int64',
        'is_fiscal' => null,
        'average_cost' => 'int64',
        'total_cost' => 'int64',
        'currency' => null,
        'rollover' => null,
        'warning_limit' => null,
        'over_limit' => null,
        'external_id' => null,
        'internal_id' => null,
        'created_by_id' => 'int64',
        'created_by_name' => null,
        'updated_by_id' => null,
        'updated_by_name' => null,
        'date_created' => 'date-time',
        'last_updated' => 'date-time'
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
        'description' => 'description',
        'account' => 'account',
        'enabled' => 'enabled',
        'ref_scope' => 'refScope',
        'ref_type' => 'refType',
        'ref_id' => 'refId',
        'ref_name' => 'refName',
        'period' => 'period',
        'year' => 'year',
        'resource_type' => 'resourceType',
        'timezone' => 'timezone',
        'start_date' => 'startDate',
        'end_date' => 'endDate',
        'interval' => 'interval',
        'costs' => 'costs',
        'is_fiscal' => 'isFiscal',
        'average_cost' => 'averageCost',
        'total_cost' => 'totalCost',
        'currency' => 'currency',
        'rollover' => 'rollover',
        'warning_limit' => 'warningLimit',
        'over_limit' => 'overLimit',
        'external_id' => 'externalId',
        'internal_id' => 'internalId',
        'created_by_id' => 'createdById',
        'created_by_name' => 'createdByName',
        'updated_by_id' => 'updatedById',
        'updated_by_name' => 'updatedByName',
        'date_created' => 'dateCreated',
        'last_updated' => 'lastUpdated'
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'id' => 'setId',
        'name' => 'setName',
        'description' => 'setDescription',
        'account' => 'setAccount',
        'enabled' => 'setEnabled',
        'ref_scope' => 'setRefScope',
        'ref_type' => 'setRefType',
        'ref_id' => 'setRefId',
        'ref_name' => 'setRefName',
        'period' => 'setPeriod',
        'year' => 'setYear',
        'resource_type' => 'setResourceType',
        'timezone' => 'setTimezone',
        'start_date' => 'setStartDate',
        'end_date' => 'setEndDate',
        'interval' => 'setInterval',
        'costs' => 'setCosts',
        'is_fiscal' => 'setIsFiscal',
        'average_cost' => 'setAverageCost',
        'total_cost' => 'setTotalCost',
        'currency' => 'setCurrency',
        'rollover' => 'setRollover',
        'warning_limit' => 'setWarningLimit',
        'over_limit' => 'setOverLimit',
        'external_id' => 'setExternalId',
        'internal_id' => 'setInternalId',
        'created_by_id' => 'setCreatedById',
        'created_by_name' => 'setCreatedByName',
        'updated_by_id' => 'setUpdatedById',
        'updated_by_name' => 'setUpdatedByName',
        'date_created' => 'setDateCreated',
        'last_updated' => 'setLastUpdated'
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'id' => 'getId',
        'name' => 'getName',
        'description' => 'getDescription',
        'account' => 'getAccount',
        'enabled' => 'getEnabled',
        'ref_scope' => 'getRefScope',
        'ref_type' => 'getRefType',
        'ref_id' => 'getRefId',
        'ref_name' => 'getRefName',
        'period' => 'getPeriod',
        'year' => 'getYear',
        'resource_type' => 'getResourceType',
        'timezone' => 'getTimezone',
        'start_date' => 'getStartDate',
        'end_date' => 'getEndDate',
        'interval' => 'getInterval',
        'costs' => 'getCosts',
        'is_fiscal' => 'getIsFiscal',
        'average_cost' => 'getAverageCost',
        'total_cost' => 'getTotalCost',
        'currency' => 'getCurrency',
        'rollover' => 'getRollover',
        'warning_limit' => 'getWarningLimit',
        'over_limit' => 'getOverLimit',
        'external_id' => 'getExternalId',
        'internal_id' => 'getInternalId',
        'created_by_id' => 'getCreatedById',
        'created_by_name' => 'getCreatedByName',
        'updated_by_id' => 'getUpdatedById',
        'updated_by_name' => 'getUpdatedByName',
        'date_created' => 'getDateCreated',
        'last_updated' => 'getLastUpdated'
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
        $this->container['description'] = $data['description'] ?? null;
        $this->container['account'] = $data['account'] ?? null;
        $this->container['enabled'] = $data['enabled'] ?? null;
        $this->container['ref_scope'] = $data['ref_scope'] ?? null;
        $this->container['ref_type'] = $data['ref_type'] ?? null;
        $this->container['ref_id'] = $data['ref_id'] ?? null;
        $this->container['ref_name'] = $data['ref_name'] ?? null;
        $this->container['period'] = $data['period'] ?? null;
        $this->container['year'] = $data['year'] ?? null;
        $this->container['resource_type'] = $data['resource_type'] ?? null;
        $this->container['timezone'] = $data['timezone'] ?? null;
        $this->container['start_date'] = $data['start_date'] ?? null;
        $this->container['end_date'] = $data['end_date'] ?? null;
        $this->container['interval'] = $data['interval'] ?? null;
        $this->container['costs'] = $data['costs'] ?? null;
        $this->container['is_fiscal'] = $data['is_fiscal'] ?? null;
        $this->container['average_cost'] = $data['average_cost'] ?? null;
        $this->container['total_cost'] = $data['total_cost'] ?? null;
        $this->container['currency'] = $data['currency'] ?? null;
        $this->container['rollover'] = $data['rollover'] ?? null;
        $this->container['warning_limit'] = $data['warning_limit'] ?? null;
        $this->container['over_limit'] = $data['over_limit'] ?? null;
        $this->container['external_id'] = $data['external_id'] ?? null;
        $this->container['internal_id'] = $data['internal_id'] ?? null;
        $this->container['created_by_id'] = $data['created_by_id'] ?? null;
        $this->container['created_by_name'] = $data['created_by_name'] ?? null;
        $this->container['updated_by_id'] = $data['updated_by_id'] ?? null;
        $this->container['updated_by_name'] = $data['updated_by_name'] ?? null;
        $this->container['date_created'] = $data['date_created'] ?? null;
        $this->container['last_updated'] = $data['last_updated'] ?? null;
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
     * Gets account
     *
     * @return \OpenAPI\Client\Model\InlineResponse20040AppDeployInstance|null
     */
    public function getAccount()
    {
        return $this->container['account'];
    }

    /**
     * Sets account
     *
     * @param \OpenAPI\Client\Model\InlineResponse20040AppDeployInstance|null $account account
     *
     * @return self
     */
    public function setAccount($account)
    {
        $this->container['account'] = $account;

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
     * Gets ref_scope
     *
     * @return string|null
     */
    public function getRefScope()
    {
        return $this->container['ref_scope'];
    }

    /**
     * Sets ref_scope
     *
     * @param string|null $ref_scope ref_scope
     *
     * @return self
     */
    public function setRefScope($ref_scope)
    {
        $this->container['ref_scope'] = $ref_scope;

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
     * Gets ref_name
     *
     * @return string|null
     */
    public function getRefName()
    {
        return $this->container['ref_name'];
    }

    /**
     * Sets ref_name
     *
     * @param string|null $ref_name ref_name
     *
     * @return self
     */
    public function setRefName($ref_name)
    {
        $this->container['ref_name'] = $ref_name;

        return $this;
    }

    /**
     * Gets period
     *
     * @return string|null
     */
    public function getPeriod()
    {
        return $this->container['period'];
    }

    /**
     * Sets period
     *
     * @param string|null $period period
     *
     * @return self
     */
    public function setPeriod($period)
    {
        $this->container['period'] = $period;

        return $this;
    }

    /**
     * Gets year
     *
     * @return string|null
     */
    public function getYear()
    {
        return $this->container['year'];
    }

    /**
     * Sets year
     *
     * @param string|null $year year
     *
     * @return self
     */
    public function setYear($year)
    {
        $this->container['year'] = $year;

        return $this;
    }

    /**
     * Gets resource_type
     *
     * @return string|null
     */
    public function getResourceType()
    {
        return $this->container['resource_type'];
    }

    /**
     * Sets resource_type
     *
     * @param string|null $resource_type resource_type
     *
     * @return self
     */
    public function setResourceType($resource_type)
    {
        $this->container['resource_type'] = $resource_type;

        return $this;
    }

    /**
     * Gets timezone
     *
     * @return string|null
     */
    public function getTimezone()
    {
        return $this->container['timezone'];
    }

    /**
     * Sets timezone
     *
     * @param string|null $timezone timezone
     *
     * @return self
     */
    public function setTimezone($timezone)
    {
        $this->container['timezone'] = $timezone;

        return $this;
    }

    /**
     * Gets start_date
     *
     * @return \DateTime|null
     */
    public function getStartDate()
    {
        return $this->container['start_date'];
    }

    /**
     * Sets start_date
     *
     * @param \DateTime|null $start_date start_date
     *
     * @return self
     */
    public function setStartDate($start_date)
    {
        $this->container['start_date'] = $start_date;

        return $this;
    }

    /**
     * Gets end_date
     *
     * @return \DateTime|null
     */
    public function getEndDate()
    {
        return $this->container['end_date'];
    }

    /**
     * Sets end_date
     *
     * @param \DateTime|null $end_date end_date
     *
     * @return self
     */
    public function setEndDate($end_date)
    {
        $this->container['end_date'] = $end_date;

        return $this;
    }

    /**
     * Gets interval
     *
     * @return string|null
     */
    public function getInterval()
    {
        return $this->container['interval'];
    }

    /**
     * Sets interval
     *
     * @param string|null $interval interval
     *
     * @return self
     */
    public function setInterval($interval)
    {
        $this->container['interval'] = $interval;

        return $this;
    }

    /**
     * Gets costs
     *
     * @return int[]|null
     */
    public function getCosts()
    {
        return $this->container['costs'];
    }

    /**
     * Sets costs
     *
     * @param int[]|null $costs costs
     *
     * @return self
     */
    public function setCosts($costs)
    {
        $this->container['costs'] = $costs;

        return $this;
    }

    /**
     * Gets is_fiscal
     *
     * @return bool|null
     */
    public function getIsFiscal()
    {
        return $this->container['is_fiscal'];
    }

    /**
     * Sets is_fiscal
     *
     * @param bool|null $is_fiscal is_fiscal
     *
     * @return self
     */
    public function setIsFiscal($is_fiscal)
    {
        $this->container['is_fiscal'] = $is_fiscal;

        return $this;
    }

    /**
     * Gets average_cost
     *
     * @return int|null
     */
    public function getAverageCost()
    {
        return $this->container['average_cost'];
    }

    /**
     * Sets average_cost
     *
     * @param int|null $average_cost average_cost
     *
     * @return self
     */
    public function setAverageCost($average_cost)
    {
        $this->container['average_cost'] = $average_cost;

        return $this;
    }

    /**
     * Gets total_cost
     *
     * @return int|null
     */
    public function getTotalCost()
    {
        return $this->container['total_cost'];
    }

    /**
     * Sets total_cost
     *
     * @param int|null $total_cost total_cost
     *
     * @return self
     */
    public function setTotalCost($total_cost)
    {
        $this->container['total_cost'] = $total_cost;

        return $this;
    }

    /**
     * Gets currency
     *
     * @return string|null
     */
    public function getCurrency()
    {
        return $this->container['currency'];
    }

    /**
     * Sets currency
     *
     * @param string|null $currency currency
     *
     * @return self
     */
    public function setCurrency($currency)
    {
        $this->container['currency'] = $currency;

        return $this;
    }

    /**
     * Gets rollover
     *
     * @return bool|null
     */
    public function getRollover()
    {
        return $this->container['rollover'];
    }

    /**
     * Sets rollover
     *
     * @param bool|null $rollover rollover
     *
     * @return self
     */
    public function setRollover($rollover)
    {
        $this->container['rollover'] = $rollover;

        return $this;
    }

    /**
     * Gets warning_limit
     *
     * @return string|null
     */
    public function getWarningLimit()
    {
        return $this->container['warning_limit'];
    }

    /**
     * Sets warning_limit
     *
     * @param string|null $warning_limit warning_limit
     *
     * @return self
     */
    public function setWarningLimit($warning_limit)
    {
        $this->container['warning_limit'] = $warning_limit;

        return $this;
    }

    /**
     * Gets over_limit
     *
     * @return string|null
     */
    public function getOverLimit()
    {
        return $this->container['over_limit'];
    }

    /**
     * Sets over_limit
     *
     * @param string|null $over_limit over_limit
     *
     * @return self
     */
    public function setOverLimit($over_limit)
    {
        $this->container['over_limit'] = $over_limit;

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
     * Gets created_by_id
     *
     * @return int|null
     */
    public function getCreatedById()
    {
        return $this->container['created_by_id'];
    }

    /**
     * Sets created_by_id
     *
     * @param int|null $created_by_id created_by_id
     *
     * @return self
     */
    public function setCreatedById($created_by_id)
    {
        $this->container['created_by_id'] = $created_by_id;

        return $this;
    }

    /**
     * Gets created_by_name
     *
     * @return string|null
     */
    public function getCreatedByName()
    {
        return $this->container['created_by_name'];
    }

    /**
     * Sets created_by_name
     *
     * @param string|null $created_by_name created_by_name
     *
     * @return self
     */
    public function setCreatedByName($created_by_name)
    {
        $this->container['created_by_name'] = $created_by_name;

        return $this;
    }

    /**
     * Gets updated_by_id
     *
     * @return string|null
     */
    public function getUpdatedById()
    {
        return $this->container['updated_by_id'];
    }

    /**
     * Sets updated_by_id
     *
     * @param string|null $updated_by_id updated_by_id
     *
     * @return self
     */
    public function setUpdatedById($updated_by_id)
    {
        $this->container['updated_by_id'] = $updated_by_id;

        return $this;
    }

    /**
     * Gets updated_by_name
     *
     * @return string|null
     */
    public function getUpdatedByName()
    {
        return $this->container['updated_by_name'];
    }

    /**
     * Sets updated_by_name
     *
     * @param string|null $updated_by_name updated_by_name
     *
     * @return self
     */
    public function setUpdatedByName($updated_by_name)
    {
        $this->container['updated_by_name'] = $updated_by_name;

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


