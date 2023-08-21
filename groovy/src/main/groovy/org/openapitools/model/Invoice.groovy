package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.math.BigDecimal;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.InlineResponse20040AppDeployInstance;
import org.openapitools.model.InvoiceCloud;
import org.openapitools.model.InvoiceLineItems;

@Canonical
class Invoice {
    
    Long id
    
    Long ownerId
    
    InlineResponse20040AppDeployInstance account
    
    Object group
    
    InvoiceCloud cloud
    
    Object instance
    
    String server
    
    String cluster
    
    Object user
    
    Object plan
    
    List<Object> tags = new ArrayList<Object>()
    
    String project
    
    String refType
    
    Long refId
    
    String refUuid
    
    String refName
    
    String refCategory
    
    String resourceId
    
    String resourceUuid
    
    String resourceType
    
    String resourceName
    
    String resourceExternalId
    
    String resourceInternalId
    
    String interval
    
    String period
    
    Boolean estimate
    
    Boolean summaryInvoice
    
    Boolean active
    
    Date startDate
    
    Date endDate
    
    Date refStart
    
    Date refEnd
    
    BigDecimal estimatedComputePrice
    
    BigDecimal estimatedComputeCost
    
    BigDecimal estimatedMemoryPrice
    
    BigDecimal estimatedMemoryCost
    
    BigDecimal estimatedStoragePrice
    
    BigDecimal estimatedStorageCost
    
    BigDecimal estimatedNetworkPrice
    
    BigDecimal estimatedNetworkCost
    
    BigDecimal estimatedLicensePrice
    
    BigDecimal estimatedLicenseCost
    
    BigDecimal estimatedExtraPrice
    
    BigDecimal estimatedExtraCost
    
    BigDecimal estimatedTotalPrice
    
    BigDecimal estimatedTotalCost
    
    BigDecimal estimatedRunningPrice
    
    BigDecimal estimatedRunningCost
    
    String estimatedCurrency
    
    BigDecimal estimatedConversionRate
    
    BigDecimal actualComputePrice
    
    BigDecimal actualComputeCost
    
    BigDecimal actualMemoryPrice
    
    BigDecimal actualMemoryCost
    
    BigDecimal actualStoragePrice
    
    BigDecimal actualStorageCost
    
    BigDecimal actualNetworkPrice
    
    BigDecimal actualNetworkCost
    
    BigDecimal actualLicensePrice
    
    BigDecimal actualLicenseCost
    
    BigDecimal actualExtraPrice
    
    BigDecimal actualExtraCost
    
    BigDecimal actualTotalPrice
    
    BigDecimal actualTotalCost
    
    BigDecimal actualRunningPrice
    
    BigDecimal actualRunningCost
    
    String actualCurrency
    
    BigDecimal actualConversionRate
    
    BigDecimal computePrice
    
    BigDecimal computeCost
    
    BigDecimal memoryPrice
    
    BigDecimal memoryCost
    
    BigDecimal storagePrice
    
    BigDecimal storageCost
    
    BigDecimal networkPrice
    
    BigDecimal networkCost
    
    BigDecimal licensePrice
    
    BigDecimal licenseCost
    
    BigDecimal extraPrice
    
    BigDecimal extraCost
    
    BigDecimal totalPrice
    
    BigDecimal totalCost
    
    BigDecimal runningPrice
    
    BigDecimal runningCost
    
    String currency
    
    BigDecimal conversionRate
    
    String costType
    
    Long offTime
    
    String powerState
    
    Date powerDate
    
    BigDecimal runningMultiplier
    
    String usageType
    
    String usageCategory
    
    Date lastCostDate
    
    Date lastActualDate
    
    Date dateCreated
    
    Date lastUpdated
    
    Long lineItemCount
    
    List<InvoiceLineItems> lineItems = new ArrayList<InvoiceLineItems>()
}
