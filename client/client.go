// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AllocateInstancePublicConnectionRequest struct {
	// The network type of the endpoint. Valid values:
	//
	// *   **primary**: primary endpoint
	// *   **cluster**: instance endpoint. This value is supported only for an instance that contains multiple coordinator nodes.
	//
	// >  The default value is primary.
	AddressType *string `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	// The prefix of the endpoint.
	//
	// Specify a prefix for the endpoint. Example: `gp-bp12ga6v69h86****`. In this example, the endpoint is `gp-bp12ga6v69h86****.gpdb.rds.aliyuncs.com`.
	ConnectionStringPrefix *string `json:"ConnectionStringPrefix,omitempty" xml:"ConnectionStringPrefix,omitempty"`
	// The ID of the instance.
	//
	// >  You can call the [DescribeDBInstances](~~86911~~) operation to query the details of all AnalyticDB for PostgreSQL instances in a specific region, including instance IDs.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The port number. Example: 5432.
	Port                 *string `json:"Port,omitempty" xml:"Port,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s AllocateInstancePublicConnectionRequest) String() string {
	return tea.Prettify(s)
}

func (s AllocateInstancePublicConnectionRequest) GoString() string {
	return s.String()
}

func (s *AllocateInstancePublicConnectionRequest) SetAddressType(v string) *AllocateInstancePublicConnectionRequest {
	s.AddressType = &v
	return s
}

func (s *AllocateInstancePublicConnectionRequest) SetConnectionStringPrefix(v string) *AllocateInstancePublicConnectionRequest {
	s.ConnectionStringPrefix = &v
	return s
}

func (s *AllocateInstancePublicConnectionRequest) SetDBInstanceId(v string) *AllocateInstancePublicConnectionRequest {
	s.DBInstanceId = &v
	return s
}

func (s *AllocateInstancePublicConnectionRequest) SetOwnerId(v int64) *AllocateInstancePublicConnectionRequest {
	s.OwnerId = &v
	return s
}

func (s *AllocateInstancePublicConnectionRequest) SetPort(v string) *AllocateInstancePublicConnectionRequest {
	s.Port = &v
	return s
}

func (s *AllocateInstancePublicConnectionRequest) SetResourceOwnerAccount(v string) *AllocateInstancePublicConnectionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AllocateInstancePublicConnectionRequest) SetResourceOwnerId(v int64) *AllocateInstancePublicConnectionRequest {
	s.ResourceOwnerId = &v
	return s
}

type AllocateInstancePublicConnectionResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AllocateInstancePublicConnectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AllocateInstancePublicConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *AllocateInstancePublicConnectionResponseBody) SetRequestId(v string) *AllocateInstancePublicConnectionResponseBody {
	s.RequestId = &v
	return s
}

type AllocateInstancePublicConnectionResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AllocateInstancePublicConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AllocateInstancePublicConnectionResponse) String() string {
	return tea.Prettify(s)
}

func (s AllocateInstancePublicConnectionResponse) GoString() string {
	return s.String()
}

func (s *AllocateInstancePublicConnectionResponse) SetHeaders(v map[string]*string) *AllocateInstancePublicConnectionResponse {
	s.Headers = v
	return s
}

func (s *AllocateInstancePublicConnectionResponse) SetStatusCode(v int32) *AllocateInstancePublicConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *AllocateInstancePublicConnectionResponse) SetBody(v *AllocateInstancePublicConnectionResponseBody) *AllocateInstancePublicConnectionResponse {
	s.Body = v
	return s
}

type CheckServiceLinkedRoleRequest struct {
	// The ID of the region. You can call the [DescribeRegions](~~86912~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CheckServiceLinkedRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckServiceLinkedRoleRequest) GoString() string {
	return s.String()
}

func (s *CheckServiceLinkedRoleRequest) SetRegionId(v string) *CheckServiceLinkedRoleRequest {
	s.RegionId = &v
	return s
}

type CheckServiceLinkedRoleResponseBody struct {
	// Indicates whether an SLR is created.
	HasServiceLinkedRole *string `json:"HasServiceLinkedRole,omitempty" xml:"HasServiceLinkedRole,omitempty"`
	// The ID of the region.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckServiceLinkedRoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckServiceLinkedRoleResponseBody) GoString() string {
	return s.String()
}

func (s *CheckServiceLinkedRoleResponseBody) SetHasServiceLinkedRole(v string) *CheckServiceLinkedRoleResponseBody {
	s.HasServiceLinkedRole = &v
	return s
}

func (s *CheckServiceLinkedRoleResponseBody) SetRegionId(v string) *CheckServiceLinkedRoleResponseBody {
	s.RegionId = &v
	return s
}

func (s *CheckServiceLinkedRoleResponseBody) SetRequestId(v string) *CheckServiceLinkedRoleResponseBody {
	s.RequestId = &v
	return s
}

type CheckServiceLinkedRoleResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CheckServiceLinkedRoleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckServiceLinkedRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckServiceLinkedRoleResponse) GoString() string {
	return s.String()
}

func (s *CheckServiceLinkedRoleResponse) SetHeaders(v map[string]*string) *CheckServiceLinkedRoleResponse {
	s.Headers = v
	return s
}

func (s *CheckServiceLinkedRoleResponse) SetStatusCode(v int32) *CheckServiceLinkedRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckServiceLinkedRoleResponse) SetBody(v *CheckServiceLinkedRoleResponseBody) *CheckServiceLinkedRoleResponse {
	s.Body = v
	return s
}

type CreateAccountRequest struct {
	// The description of the privileged account.
	AccountDescription *string `json:"AccountDescription,omitempty" xml:"AccountDescription,omitempty"`
	// The name of the privileged account.
	//
	// *   The name can contain lowercase letters, digits, and underscores (\_).
	// *   The name must start with a lowercase letter and end with a lowercase letter or a digit.
	// *   The name cannot start with gp.
	// *   The name must be 2 to 16 characters in length.
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The password of the privileged account.
	//
	// *   The password must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters.
	// *   Special characters include `! @ # $ % ^ & * ( ) _ + - =`
	// *   The password must be 8 to 32 characters in length.
	AccountPassword *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
	// The ID of the instance.
	//
	// >  You can call the [DescribeDBInstances](~~86911~~) operation to query the details of all AnalyticDB for PostgreSQL instances in a specific region, including instance IDs.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The name of the database.
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the resource group to which the instance belongs.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s CreateAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAccountRequest) GoString() string {
	return s.String()
}

func (s *CreateAccountRequest) SetAccountDescription(v string) *CreateAccountRequest {
	s.AccountDescription = &v
	return s
}

func (s *CreateAccountRequest) SetAccountName(v string) *CreateAccountRequest {
	s.AccountName = &v
	return s
}

func (s *CreateAccountRequest) SetAccountPassword(v string) *CreateAccountRequest {
	s.AccountPassword = &v
	return s
}

func (s *CreateAccountRequest) SetDBInstanceId(v string) *CreateAccountRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CreateAccountRequest) SetDatabaseName(v string) *CreateAccountRequest {
	s.DatabaseName = &v
	return s
}

func (s *CreateAccountRequest) SetOwnerId(v int64) *CreateAccountRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateAccountRequest) SetResourceGroupId(v string) *CreateAccountRequest {
	s.ResourceGroupId = &v
	return s
}

type CreateAccountResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAccountResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAccountResponseBody) SetRequestId(v string) *CreateAccountResponseBody {
	s.RequestId = &v
	return s
}

type CreateAccountResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateAccountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAccountResponse) GoString() string {
	return s.String()
}

func (s *CreateAccountResponse) SetHeaders(v map[string]*string) *CreateAccountResponse {
	s.Headers = v
	return s
}

func (s *CreateAccountResponse) SetStatusCode(v int32) *CreateAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAccountResponse) SetBody(v *CreateAccountResponseBody) *CreateAccountResponse {
	s.Body = v
	return s
}

type CreateCollectionRequest struct {
	Collection              *string `json:"Collection,omitempty" xml:"Collection,omitempty"`
	DBInstanceId            *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	Dimension               *int64  `json:"Dimension,omitempty" xml:"Dimension,omitempty"`
	FullTextRetrievalFields *string `json:"FullTextRetrievalFields,omitempty" xml:"FullTextRetrievalFields,omitempty"`
	ManagerAccount          *string `json:"ManagerAccount,omitempty" xml:"ManagerAccount,omitempty"`
	ManagerAccountPassword  *string `json:"ManagerAccountPassword,omitempty" xml:"ManagerAccountPassword,omitempty"`
	Metadata                *string `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	Namespace               *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	OwnerId                 *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Parser                  *string `json:"Parser,omitempty" xml:"Parser,omitempty"`
	RegionId                *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateCollectionRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCollectionRequest) GoString() string {
	return s.String()
}

func (s *CreateCollectionRequest) SetCollection(v string) *CreateCollectionRequest {
	s.Collection = &v
	return s
}

func (s *CreateCollectionRequest) SetDBInstanceId(v string) *CreateCollectionRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CreateCollectionRequest) SetDimension(v int64) *CreateCollectionRequest {
	s.Dimension = &v
	return s
}

func (s *CreateCollectionRequest) SetFullTextRetrievalFields(v string) *CreateCollectionRequest {
	s.FullTextRetrievalFields = &v
	return s
}

func (s *CreateCollectionRequest) SetManagerAccount(v string) *CreateCollectionRequest {
	s.ManagerAccount = &v
	return s
}

func (s *CreateCollectionRequest) SetManagerAccountPassword(v string) *CreateCollectionRequest {
	s.ManagerAccountPassword = &v
	return s
}

func (s *CreateCollectionRequest) SetMetadata(v string) *CreateCollectionRequest {
	s.Metadata = &v
	return s
}

func (s *CreateCollectionRequest) SetNamespace(v string) *CreateCollectionRequest {
	s.Namespace = &v
	return s
}

func (s *CreateCollectionRequest) SetOwnerId(v int64) *CreateCollectionRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateCollectionRequest) SetParser(v string) *CreateCollectionRequest {
	s.Parser = &v
	return s
}

func (s *CreateCollectionRequest) SetRegionId(v string) *CreateCollectionRequest {
	s.RegionId = &v
	return s
}

type CreateCollectionResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateCollectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCollectionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCollectionResponseBody) SetMessage(v string) *CreateCollectionResponseBody {
	s.Message = &v
	return s
}

func (s *CreateCollectionResponseBody) SetRequestId(v string) *CreateCollectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCollectionResponseBody) SetStatus(v string) *CreateCollectionResponseBody {
	s.Status = &v
	return s
}

type CreateCollectionResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateCollectionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateCollectionResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCollectionResponse) GoString() string {
	return s.String()
}

func (s *CreateCollectionResponse) SetHeaders(v map[string]*string) *CreateCollectionResponse {
	s.Headers = v
	return s
}

func (s *CreateCollectionResponse) SetStatusCode(v int32) *CreateCollectionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCollectionResponse) SetBody(v *CreateCollectionResponseBody) *CreateCollectionResponse {
	s.Body = v
	return s
}

type CreateDBInstanceRequest struct {
	// The client token that is used to ensure the idempotence of the request. For more information, see [Ensure idempotence](~~327176~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to load a sample dataset after the instance is created. Valid values:
	//
	// - **true**
	// - **false**
	//
	// If you do not specify this parameter, no sample dataset is loaded.
	CreateSampleData *bool `json:"CreateSampleData,omitempty" xml:"CreateSampleData,omitempty"`
	// The edition of the instance. Valid values:
	//
	// - **HighAvailability**: High-availability Edition.
	// - **Basic**: Basic Edition.
	//
	// > This parameter must be specified when you create an instance in elastic storage mode.
	DBInstanceCategory *string `json:"DBInstanceCategory,omitempty" xml:"DBInstanceCategory,omitempty"`
	// The instance type of the instance. For information, see [Instance types](~~86942~~).
	//
	// > This parameter must be specified when you create an instance in reserved storage mode.
	DBInstanceClass *string `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty"`
	// The description of the instance.
	DBInstanceDescription *string `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	// The number of compute groups. Valid values: 2, 4, 8, 12, 16, 24, 32, 64, 96, and 128.
	//
	// > This parameter must be specified when you create an instance in reserved storage mode.
	DBInstanceGroupCount *string `json:"DBInstanceGroupCount,omitempty" xml:"DBInstanceGroupCount,omitempty"`
	// The resource type of the instance. Valid values:
	//
	// - **StorageElastic**: elastic storage mode.
	// - **Serverless**: Serverless mode.
	// - **Classic**: reserved storage mode.
	//
	// > This parameter must be specified.
	DBInstanceMode *string `json:"DBInstanceMode,omitempty" xml:"DBInstanceMode,omitempty"`
	// The ID of the encryption key.
	//
	// > If EncryptionType is set to CloudDisk, you must specify an encryption key that resides in the same region as the cloud disk that is specified by EncryptionType. Otherwise, leave this parameter empty.
	EncryptionKey *string `json:"EncryptionKey,omitempty" xml:"EncryptionKey,omitempty"`
	// The encryption type. Valid values:
	//
	// - **NULL** (default): Encryption is disabled.
	// - **CloudDisk**: Encryption is enabled on cloud disks, and EncryptionKey is used to specify an encryption key.
	//
	// > Disk encryption cannot be disabled after it is enabled.
	EncryptionType *string `json:"EncryptionType,omitempty" xml:"EncryptionType,omitempty"`
	// The database engine of the instance. Set the value to gpdb.
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The version of the database engine. Valid values:
	//
	// - 6.0
	// - 7.0
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The wait time for the instance that has no traffic to become idle. Minimum value: 60. Default value: 600. Unit: seconds.
	//
	// > This parameter must be specified only when you create an instance in automatic Serverless mode.
	IdleTime *int32 `json:"IdleTime,omitempty" xml:"IdleTime,omitempty"`
	// The network type of the instance. Set the value to VPC.
	//
	// >
	// - Only the Virtual Private Cloud (VPC) type is supported.
	// - If you do not specify this parameter, VPC is used.
	InstanceNetworkType *string `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
	// The specifications of compute nodes.
	//
	// Valid values for High-availability Edition instances in elastic storage mode:
	//
	// - **2C16G**
	// - **4C32G**
	// - **16C128G**
	//
	// Valid values for Basic Edition instances in elastic storage mode:
	//
	// - **2C8G**
	// - **4C16G**
	// - **8C32G**
	// - **16C64G**
	//
	// Valid values for instances in Serverless mode:
	//
	// - **4C16G**
	// - **8C32G**
	//
	// > This parameter must be specified when you create an instance in elastic storage mode or Serverless mode.
	InstanceSpec *string `json:"InstanceSpec,omitempty" xml:"InstanceSpec,omitempty"`
	// The number of coordinator nodes. Valid values: 1 and 2.
	//
	// > If you do not specify this parameter, 1 is used.
	MasterNodeNum *string `json:"MasterNodeNum,omitempty" xml:"MasterNodeNum,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// - **Postpaid**: pay-as-you-go.
	// - **Prepaid**: subscription.
	// >
	// - If you do not specify this parameter, Postpaid is used.
	// - You can obtain more cost savings if you create a subscription instance for one year or longer. We recommend that you select the billing method that best suits your needs.
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The unit of the subscription duration. Valid values:
	//
	// - **Month**
	// - **Year**
	// > This parameter must be specified when PayType is set to Prepaid.
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The private IP address of the instance.
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	// The ID of the region. You can call the [DescribeRegions](~~86912~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the instance belongs.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The IP address whitelist of the instance.
	//
	// A value of 127.0.0.1 specifies that no IP address is allowed for external access. You can call the [ModifySecurityIps](~~86928~~) operation to modify the IP address whitelist after you create an instance.
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	// The performance level of ESSDs. Valid values:
	//
	// - **pl0**
	// - **pl1**
	// - **pl2**
	// >
	// - This parameter takes effect only when SegStorageType is set to cloud_essd.
	// - If you do not specify this parameter, pl1 is used.
	SegDiskPerformanceLevel *string `json:"SegDiskPerformanceLevel,omitempty" xml:"SegDiskPerformanceLevel,omitempty"`
	// The number of compute nodes.
	//
	// - Valid values for High-availability Edition instances in elastic storage mode: multiples of 4 in the range of 4 to 512.
	// - Valid values for Basic Edition instances in elastic storage mode: multiples of 2 in the range of 2 to 512.
	// - Valid values for instances in Serverless mode: multiples of 2 in the range of 2 to 512.
	//
	// > This parameter must be specified when you create an instance in elastic storage mode or Serverless mode.
	SegNodeNum *string `json:"SegNodeNum,omitempty" xml:"SegNodeNum,omitempty"`
	// The disk storage type of the instance. Only enhanced SSDs (ESSDs) are supported. Set the value to cloud_essd.
	//
	// > This parameter must be specified when you create an instance in elastic storage mode.
	SegStorageType *string `json:"SegStorageType,omitempty" xml:"SegStorageType,omitempty"`
	// The type of the Serverless mode. Valid values:
	//
	// - **Manual** (default): manual scheduling.
	// - **Auto**: automatic scheduling.
	//
	// > This parameter must be specified only when you create an instance in Serverless mode.
	ServerlessMode *string `json:"ServerlessMode,omitempty" xml:"ServerlessMode,omitempty"`
	// The threshold of computing resources. Unit: AnalyticDB compute unit (ACU). Valid values: 8 to 32. The value must be in increments of 8 ACUs. Default value: 32.
	//
	// > This parameter must be specified only when you create an instance in automatic Serverless mode.
	ServerlessResource *int32 `json:"ServerlessResource,omitempty" xml:"ServerlessResource,omitempty"`
	// The storage capacity of the instance. Unit: GB. Valid values: 50 to 4000.
	//
	// > This parameter must be specified when you create an instance in elastic storage mode.
	StorageSize *int64 `json:"StorageSize,omitempty" xml:"StorageSize,omitempty"`
	// This parameter is no longer used.
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// The list of tags.
	Tag []*CreateDBInstanceRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The subscription duration.
	//
	// - Valid values when Period is set to Month: 1 to 9.
	// - Valid values when Period is set to Year: 1 to 3.
	// > This parameter must be specified when PayType is set to Prepaid.
	UsedTime *string `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	// The VPC ID of the instance.
	//
	// >
	// - This parameter must be specified.
	// - The region where the VPC resides must be the same as the region that is specified by RegionId.
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// The vSwitch ID of the instance.
	//
	// >
	// - This parameter must be specified.
	// - The zone where the vSwitch resides must be the same as the zone that is specified by ZoneId.
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// Specifies whether to enable vector engine optimization. Valid values:
	//
	// - **enabled**
	// - **disabled** (default)
	//
	// >
	// - We recommend that you do not enable vector engine optimization in mainstream analysis and real-time data warehousing scenarios.
	// - We recommend that you enable vector engine optimization in AI Generated Content (AIGC) and vector retrieval scenarios that require the vector analysis engine.
	VectorConfigurationStatus *string `json:"VectorConfigurationStatus,omitempty" xml:"VectorConfigurationStatus,omitempty"`
	// The zone ID of the read-only instance. You can call the [DescribeRegions](~~86912~~) operation to query the most recent zone list.
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateDBInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDBInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateDBInstanceRequest) SetClientToken(v string) *CreateDBInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDBInstanceRequest) SetCreateSampleData(v bool) *CreateDBInstanceRequest {
	s.CreateSampleData = &v
	return s
}

func (s *CreateDBInstanceRequest) SetDBInstanceCategory(v string) *CreateDBInstanceRequest {
	s.DBInstanceCategory = &v
	return s
}

func (s *CreateDBInstanceRequest) SetDBInstanceClass(v string) *CreateDBInstanceRequest {
	s.DBInstanceClass = &v
	return s
}

func (s *CreateDBInstanceRequest) SetDBInstanceDescription(v string) *CreateDBInstanceRequest {
	s.DBInstanceDescription = &v
	return s
}

func (s *CreateDBInstanceRequest) SetDBInstanceGroupCount(v string) *CreateDBInstanceRequest {
	s.DBInstanceGroupCount = &v
	return s
}

func (s *CreateDBInstanceRequest) SetDBInstanceMode(v string) *CreateDBInstanceRequest {
	s.DBInstanceMode = &v
	return s
}

func (s *CreateDBInstanceRequest) SetEncryptionKey(v string) *CreateDBInstanceRequest {
	s.EncryptionKey = &v
	return s
}

func (s *CreateDBInstanceRequest) SetEncryptionType(v string) *CreateDBInstanceRequest {
	s.EncryptionType = &v
	return s
}

func (s *CreateDBInstanceRequest) SetEngine(v string) *CreateDBInstanceRequest {
	s.Engine = &v
	return s
}

func (s *CreateDBInstanceRequest) SetEngineVersion(v string) *CreateDBInstanceRequest {
	s.EngineVersion = &v
	return s
}

func (s *CreateDBInstanceRequest) SetIdleTime(v int32) *CreateDBInstanceRequest {
	s.IdleTime = &v
	return s
}

func (s *CreateDBInstanceRequest) SetInstanceNetworkType(v string) *CreateDBInstanceRequest {
	s.InstanceNetworkType = &v
	return s
}

func (s *CreateDBInstanceRequest) SetInstanceSpec(v string) *CreateDBInstanceRequest {
	s.InstanceSpec = &v
	return s
}

func (s *CreateDBInstanceRequest) SetMasterNodeNum(v string) *CreateDBInstanceRequest {
	s.MasterNodeNum = &v
	return s
}

func (s *CreateDBInstanceRequest) SetOwnerId(v int64) *CreateDBInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetPayType(v string) *CreateDBInstanceRequest {
	s.PayType = &v
	return s
}

func (s *CreateDBInstanceRequest) SetPeriod(v string) *CreateDBInstanceRequest {
	s.Period = &v
	return s
}

func (s *CreateDBInstanceRequest) SetPrivateIpAddress(v string) *CreateDBInstanceRequest {
	s.PrivateIpAddress = &v
	return s
}

func (s *CreateDBInstanceRequest) SetRegionId(v string) *CreateDBInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetResourceGroupId(v string) *CreateDBInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetSecurityIPList(v string) *CreateDBInstanceRequest {
	s.SecurityIPList = &v
	return s
}

func (s *CreateDBInstanceRequest) SetSegDiskPerformanceLevel(v string) *CreateDBInstanceRequest {
	s.SegDiskPerformanceLevel = &v
	return s
}

func (s *CreateDBInstanceRequest) SetSegNodeNum(v string) *CreateDBInstanceRequest {
	s.SegNodeNum = &v
	return s
}

func (s *CreateDBInstanceRequest) SetSegStorageType(v string) *CreateDBInstanceRequest {
	s.SegStorageType = &v
	return s
}

func (s *CreateDBInstanceRequest) SetServerlessMode(v string) *CreateDBInstanceRequest {
	s.ServerlessMode = &v
	return s
}

func (s *CreateDBInstanceRequest) SetServerlessResource(v int32) *CreateDBInstanceRequest {
	s.ServerlessResource = &v
	return s
}

func (s *CreateDBInstanceRequest) SetStorageSize(v int64) *CreateDBInstanceRequest {
	s.StorageSize = &v
	return s
}

func (s *CreateDBInstanceRequest) SetStorageType(v string) *CreateDBInstanceRequest {
	s.StorageType = &v
	return s
}

func (s *CreateDBInstanceRequest) SetTag(v []*CreateDBInstanceRequestTag) *CreateDBInstanceRequest {
	s.Tag = v
	return s
}

func (s *CreateDBInstanceRequest) SetUsedTime(v string) *CreateDBInstanceRequest {
	s.UsedTime = &v
	return s
}

func (s *CreateDBInstanceRequest) SetVPCId(v string) *CreateDBInstanceRequest {
	s.VPCId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetVSwitchId(v string) *CreateDBInstanceRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetVectorConfigurationStatus(v string) *CreateDBInstanceRequest {
	s.VectorConfigurationStatus = &v
	return s
}

func (s *CreateDBInstanceRequest) SetZoneId(v string) *CreateDBInstanceRequest {
	s.ZoneId = &v
	return s
}

type CreateDBInstanceRequestTag struct {
	// The key of tag N. Take note of the following requirements:
	//
	// - The tag key cannot be an empty string.
	// - The tag key can be up to 128 characters in length.
	// - The tag key cannot start with `aliyun` or `acs:`, and contain `http://` or `https://`.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N. Take note of the following requirements:
	//
	// - The tag key cannot be an empty string.
	// - The tag key can be up to 128 characters in length.
	// - The tag key cannot start with `aliyun` or `acs:`, and contain `http://` or `https://`.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateDBInstanceRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateDBInstanceRequestTag) GoString() string {
	return s.String()
}

func (s *CreateDBInstanceRequestTag) SetKey(v string) *CreateDBInstanceRequestTag {
	s.Key = &v
	return s
}

func (s *CreateDBInstanceRequestTag) SetValue(v string) *CreateDBInstanceRequestTag {
	s.Value = &v
	return s
}

type CreateDBInstanceResponseBody struct {
	// An invalid parameter. It is no longer returned when you call this operation.
	//
	// You can call the [DescribeDBInstanceAttribute](~~86910~~) operation to query the endpoint that is used to connect to the instance.
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// The instance ID.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The order ID.
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// An invalid parameter. It is no longer returned when you call this operation.
	//
	// You can call the [DescribeDBInstanceAttribute](~~86910~~) operation to query the port number that is used to connect to the instance.
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDBInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDBInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDBInstanceResponseBody) SetConnectionString(v string) *CreateDBInstanceResponseBody {
	s.ConnectionString = &v
	return s
}

func (s *CreateDBInstanceResponseBody) SetDBInstanceId(v string) *CreateDBInstanceResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *CreateDBInstanceResponseBody) SetOrderId(v string) *CreateDBInstanceResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateDBInstanceResponseBody) SetPort(v string) *CreateDBInstanceResponseBody {
	s.Port = &v
	return s
}

func (s *CreateDBInstanceResponseBody) SetRequestId(v string) *CreateDBInstanceResponseBody {
	s.RequestId = &v
	return s
}

type CreateDBInstanceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateDBInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDBInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDBInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateDBInstanceResponse) SetHeaders(v map[string]*string) *CreateDBInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateDBInstanceResponse) SetStatusCode(v int32) *CreateDBInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDBInstanceResponse) SetBody(v *CreateDBInstanceResponseBody) *CreateDBInstanceResponse {
	s.Body = v
	return s
}

type CreateDBInstancePlanRequest struct {
	// The ID of instance.
	//
	// >  You can call the [DescribeDBInstances](~~86911~~) operation to query the IDs of all AnalyticDB for PostgreSQL instances in a specific region.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The execution information of the plan. Specify the parameter in the JSON format. The parameter value varies based on the values of the **PlanType** and **PlanScheduleType** parameters. The following section describes the PlanConfig parameter.
	PlanConfig *string `json:"PlanConfig,omitempty" xml:"PlanConfig,omitempty"`
	// The description of the plan.
	PlanDesc *string `json:"PlanDesc,omitempty" xml:"PlanDesc,omitempty"`
	// The end time of the plan. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm*Z format. The time must be in UTC. The end time must be later than the start time.
	//
	// > *   This parameter is required only if the **PlanScheduleType** parameter is set to **Regular**.
	// > *   If you do not specify this parameter, the plan does not end.
	PlanEndDate *string `json:"PlanEndDate,omitempty" xml:"PlanEndDate,omitempty"`
	// The name of the plan.
	PlanName *string `json:"PlanName,omitempty" xml:"PlanName,omitempty"`
	// The execution mode of the plan. Valid values:
	//
	// *   **Postpone**: The plan is executed later.
	// *   **Regular**: The plan is executed periodically.
	PlanScheduleType *string `json:"PlanScheduleType,omitempty" xml:"PlanScheduleType,omitempty"`
	// The start time of the plan. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm*Z format. The time must be in UTC.
	//
	// >  *   This parameter is required only if the **PlanScheduleType** parameter is set to **Regular**.
	// >  *   If you do not specify this parameter, the plan is executed immediately.
	PlanStartDate *string `json:"PlanStartDate,omitempty" xml:"PlanStartDate,omitempty"`
	// The type of the plan. Valid values:
	//
	// *   **PauseResume**: pauses and resumes an instance.
	// *   **Resize**: changes the number of compute nodes.
	// *   **ModifySpec**: changes compute node specifications.
	//
	// > *   You can specify the value to Resize only for instances in Serverless mode.
	// > *   You can specify the value to ModifySpec only for instances in elastic storage mode.
	PlanType *string `json:"PlanType,omitempty" xml:"PlanType,omitempty"`
}

func (s CreateDBInstancePlanRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDBInstancePlanRequest) GoString() string {
	return s.String()
}

func (s *CreateDBInstancePlanRequest) SetDBInstanceId(v string) *CreateDBInstancePlanRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CreateDBInstancePlanRequest) SetOwnerId(v int64) *CreateDBInstancePlanRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateDBInstancePlanRequest) SetPlanConfig(v string) *CreateDBInstancePlanRequest {
	s.PlanConfig = &v
	return s
}

func (s *CreateDBInstancePlanRequest) SetPlanDesc(v string) *CreateDBInstancePlanRequest {
	s.PlanDesc = &v
	return s
}

func (s *CreateDBInstancePlanRequest) SetPlanEndDate(v string) *CreateDBInstancePlanRequest {
	s.PlanEndDate = &v
	return s
}

func (s *CreateDBInstancePlanRequest) SetPlanName(v string) *CreateDBInstancePlanRequest {
	s.PlanName = &v
	return s
}

func (s *CreateDBInstancePlanRequest) SetPlanScheduleType(v string) *CreateDBInstancePlanRequest {
	s.PlanScheduleType = &v
	return s
}

func (s *CreateDBInstancePlanRequest) SetPlanStartDate(v string) *CreateDBInstancePlanRequest {
	s.PlanStartDate = &v
	return s
}

func (s *CreateDBInstancePlanRequest) SetPlanType(v string) *CreateDBInstancePlanRequest {
	s.PlanType = &v
	return s
}

type CreateDBInstancePlanResponseBody struct {
	// The ID of instance.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The error message.
	//
	// This parameter is returned only if the operation fails.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the plan.
	PlanId *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The state of the operation.
	//
	// If the operation is successful, **success** is returned. If the operation fails, this parameter is not returned.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateDBInstancePlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDBInstancePlanResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDBInstancePlanResponseBody) SetDBInstanceId(v string) *CreateDBInstancePlanResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *CreateDBInstancePlanResponseBody) SetErrorMessage(v string) *CreateDBInstancePlanResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateDBInstancePlanResponseBody) SetPlanId(v string) *CreateDBInstancePlanResponseBody {
	s.PlanId = &v
	return s
}

func (s *CreateDBInstancePlanResponseBody) SetRequestId(v string) *CreateDBInstancePlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDBInstancePlanResponseBody) SetStatus(v string) *CreateDBInstancePlanResponseBody {
	s.Status = &v
	return s
}

type CreateDBInstancePlanResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateDBInstancePlanResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDBInstancePlanResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDBInstancePlanResponse) GoString() string {
	return s.String()
}

func (s *CreateDBInstancePlanResponse) SetHeaders(v map[string]*string) *CreateDBInstancePlanResponse {
	s.Headers = v
	return s
}

func (s *CreateDBInstancePlanResponse) SetStatusCode(v int32) *CreateDBInstancePlanResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDBInstancePlanResponse) SetBody(v *CreateDBInstancePlanResponseBody) *CreateDBInstancePlanResponse {
	s.Body = v
	return s
}

type CreateNamespaceRequest struct {
	DBInstanceId           *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	ManagerAccount         *string `json:"ManagerAccount,omitempty" xml:"ManagerAccount,omitempty"`
	ManagerAccountPassword *string `json:"ManagerAccountPassword,omitempty" xml:"ManagerAccountPassword,omitempty"`
	Namespace              *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	NamespacePassword      *string `json:"NamespacePassword,omitempty" xml:"NamespacePassword,omitempty"`
	OwnerId                *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId               *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateNamespaceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateNamespaceRequest) GoString() string {
	return s.String()
}

func (s *CreateNamespaceRequest) SetDBInstanceId(v string) *CreateNamespaceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CreateNamespaceRequest) SetManagerAccount(v string) *CreateNamespaceRequest {
	s.ManagerAccount = &v
	return s
}

func (s *CreateNamespaceRequest) SetManagerAccountPassword(v string) *CreateNamespaceRequest {
	s.ManagerAccountPassword = &v
	return s
}

func (s *CreateNamespaceRequest) SetNamespace(v string) *CreateNamespaceRequest {
	s.Namespace = &v
	return s
}

func (s *CreateNamespaceRequest) SetNamespacePassword(v string) *CreateNamespaceRequest {
	s.NamespacePassword = &v
	return s
}

func (s *CreateNamespaceRequest) SetOwnerId(v int64) *CreateNamespaceRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateNamespaceRequest) SetRegionId(v string) *CreateNamespaceRequest {
	s.RegionId = &v
	return s
}

type CreateNamespaceResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateNamespaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateNamespaceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateNamespaceResponseBody) SetMessage(v string) *CreateNamespaceResponseBody {
	s.Message = &v
	return s
}

func (s *CreateNamespaceResponseBody) SetRequestId(v string) *CreateNamespaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateNamespaceResponseBody) SetStatus(v string) *CreateNamespaceResponseBody {
	s.Status = &v
	return s
}

type CreateNamespaceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateNamespaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateNamespaceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateNamespaceResponse) GoString() string {
	return s.String()
}

func (s *CreateNamespaceResponse) SetHeaders(v map[string]*string) *CreateNamespaceResponse {
	s.Headers = v
	return s
}

func (s *CreateNamespaceResponse) SetStatusCode(v int32) *CreateNamespaceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateNamespaceResponse) SetBody(v *CreateNamespaceResponseBody) *CreateNamespaceResponse {
	s.Body = v
	return s
}

type CreateSampleDataRequest struct {
	// The ID of the instance.
	//
	// >  You can call the [DescribeDBInstances](~~86911~~) operation to query the IDs of all AnalyticDB for PostgreSQL instances in a specific region.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s CreateSampleDataRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSampleDataRequest) GoString() string {
	return s.String()
}

func (s *CreateSampleDataRequest) SetDBInstanceId(v string) *CreateSampleDataRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CreateSampleDataRequest) SetOwnerId(v int64) *CreateSampleDataRequest {
	s.OwnerId = &v
	return s
}

type CreateSampleDataResponseBody struct {
	// The ID of the instance.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The error message returned if an error occurs. This message does not affect the execution of the operation.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The execution state of the operation. Valid values:
	//
	// *   **false**: The operation fails.
	// *   **true**: The operation is successful.
	Status *bool `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateSampleDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSampleDataResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSampleDataResponseBody) SetDBInstanceId(v string) *CreateSampleDataResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *CreateSampleDataResponseBody) SetErrorMessage(v string) *CreateSampleDataResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateSampleDataResponseBody) SetRequestId(v string) *CreateSampleDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSampleDataResponseBody) SetStatus(v bool) *CreateSampleDataResponseBody {
	s.Status = &v
	return s
}

type CreateSampleDataResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateSampleDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSampleDataResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSampleDataResponse) GoString() string {
	return s.String()
}

func (s *CreateSampleDataResponse) SetHeaders(v map[string]*string) *CreateSampleDataResponse {
	s.Headers = v
	return s
}

func (s *CreateSampleDataResponse) SetStatusCode(v int32) *CreateSampleDataResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSampleDataResponse) SetBody(v *CreateSampleDataResponseBody) *CreateSampleDataResponse {
	s.Body = v
	return s
}

type CreateServiceLinkedRoleRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateServiceLinkedRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceLinkedRoleRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceLinkedRoleRequest) SetOwnerId(v int64) *CreateServiceLinkedRoleRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateServiceLinkedRoleRequest) SetRegionId(v string) *CreateServiceLinkedRoleRequest {
	s.RegionId = &v
	return s
}

type CreateServiceLinkedRoleResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateServiceLinkedRoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceLinkedRoleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServiceLinkedRoleResponseBody) SetRequestId(v string) *CreateServiceLinkedRoleResponseBody {
	s.RequestId = &v
	return s
}

type CreateServiceLinkedRoleResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateServiceLinkedRoleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateServiceLinkedRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceLinkedRoleResponse) GoString() string {
	return s.String()
}

func (s *CreateServiceLinkedRoleResponse) SetHeaders(v map[string]*string) *CreateServiceLinkedRoleResponse {
	s.Headers = v
	return s
}

func (s *CreateServiceLinkedRoleResponse) SetStatusCode(v int32) *CreateServiceLinkedRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateServiceLinkedRoleResponse) SetBody(v *CreateServiceLinkedRoleResponseBody) *CreateServiceLinkedRoleResponse {
	s.Body = v
	return s
}

type CreateVectorIndexRequest struct {
	Collection             *string `json:"Collection,omitempty" xml:"Collection,omitempty"`
	DBInstanceId           *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	Dimension              *int32  `json:"Dimension,omitempty" xml:"Dimension,omitempty"`
	ManagerAccount         *string `json:"ManagerAccount,omitempty" xml:"ManagerAccount,omitempty"`
	ManagerAccountPassword *string `json:"ManagerAccountPassword,omitempty" xml:"ManagerAccountPassword,omitempty"`
	// Distance Metrics。
	Metrics   *string `json:"Metrics,omitempty" xml:"Metrics,omitempty"`
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateVectorIndexRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateVectorIndexRequest) GoString() string {
	return s.String()
}

func (s *CreateVectorIndexRequest) SetCollection(v string) *CreateVectorIndexRequest {
	s.Collection = &v
	return s
}

func (s *CreateVectorIndexRequest) SetDBInstanceId(v string) *CreateVectorIndexRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CreateVectorIndexRequest) SetDimension(v int32) *CreateVectorIndexRequest {
	s.Dimension = &v
	return s
}

func (s *CreateVectorIndexRequest) SetManagerAccount(v string) *CreateVectorIndexRequest {
	s.ManagerAccount = &v
	return s
}

func (s *CreateVectorIndexRequest) SetManagerAccountPassword(v string) *CreateVectorIndexRequest {
	s.ManagerAccountPassword = &v
	return s
}

func (s *CreateVectorIndexRequest) SetMetrics(v string) *CreateVectorIndexRequest {
	s.Metrics = &v
	return s
}

func (s *CreateVectorIndexRequest) SetNamespace(v string) *CreateVectorIndexRequest {
	s.Namespace = &v
	return s
}

func (s *CreateVectorIndexRequest) SetOwnerId(v int64) *CreateVectorIndexRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateVectorIndexRequest) SetRegionId(v string) *CreateVectorIndexRequest {
	s.RegionId = &v
	return s
}

type CreateVectorIndexResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateVectorIndexResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateVectorIndexResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVectorIndexResponseBody) SetMessage(v string) *CreateVectorIndexResponseBody {
	s.Message = &v
	return s
}

func (s *CreateVectorIndexResponseBody) SetRequestId(v string) *CreateVectorIndexResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVectorIndexResponseBody) SetStatus(v string) *CreateVectorIndexResponseBody {
	s.Status = &v
	return s
}

type CreateVectorIndexResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateVectorIndexResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateVectorIndexResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateVectorIndexResponse) GoString() string {
	return s.String()
}

func (s *CreateVectorIndexResponse) SetHeaders(v map[string]*string) *CreateVectorIndexResponse {
	s.Headers = v
	return s
}

func (s *CreateVectorIndexResponse) SetStatusCode(v int32) *CreateVectorIndexResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVectorIndexResponse) SetBody(v *CreateVectorIndexResponseBody) *CreateVectorIndexResponse {
	s.Body = v
	return s
}

type DeleteCollectionRequest struct {
	Collection        *string `json:"Collection,omitempty" xml:"Collection,omitempty"`
	DBInstanceId      *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	Namespace         *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	NamespacePassword *string `json:"NamespacePassword,omitempty" xml:"NamespacePassword,omitempty"`
	OwnerId           *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteCollectionRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCollectionRequest) GoString() string {
	return s.String()
}

func (s *DeleteCollectionRequest) SetCollection(v string) *DeleteCollectionRequest {
	s.Collection = &v
	return s
}

func (s *DeleteCollectionRequest) SetDBInstanceId(v string) *DeleteCollectionRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DeleteCollectionRequest) SetNamespace(v string) *DeleteCollectionRequest {
	s.Namespace = &v
	return s
}

func (s *DeleteCollectionRequest) SetNamespacePassword(v string) *DeleteCollectionRequest {
	s.NamespacePassword = &v
	return s
}

func (s *DeleteCollectionRequest) SetOwnerId(v int64) *DeleteCollectionRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteCollectionRequest) SetRegionId(v string) *DeleteCollectionRequest {
	s.RegionId = &v
	return s
}

type DeleteCollectionResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DeleteCollectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteCollectionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCollectionResponseBody) SetMessage(v string) *DeleteCollectionResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteCollectionResponseBody) SetRequestId(v string) *DeleteCollectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCollectionResponseBody) SetStatus(v string) *DeleteCollectionResponseBody {
	s.Status = &v
	return s
}

type DeleteCollectionResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteCollectionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteCollectionResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCollectionResponse) GoString() string {
	return s.String()
}

func (s *DeleteCollectionResponse) SetHeaders(v map[string]*string) *DeleteCollectionResponse {
	s.Headers = v
	return s
}

func (s *DeleteCollectionResponse) SetStatusCode(v int32) *DeleteCollectionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCollectionResponse) SetBody(v *DeleteCollectionResponseBody) *DeleteCollectionResponse {
	s.Body = v
	return s
}

type DeleteCollectionDataRequest struct {
	Collection           *string `json:"Collection,omitempty" xml:"Collection,omitempty"`
	CollectionData       *string `json:"CollectionData,omitempty" xml:"CollectionData,omitempty"`
	CollectionDataFilter *string `json:"CollectionDataFilter,omitempty" xml:"CollectionDataFilter,omitempty"`
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	Namespace            *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	NamespacePassword    *string `json:"NamespacePassword,omitempty" xml:"NamespacePassword,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteCollectionDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCollectionDataRequest) GoString() string {
	return s.String()
}

func (s *DeleteCollectionDataRequest) SetCollection(v string) *DeleteCollectionDataRequest {
	s.Collection = &v
	return s
}

func (s *DeleteCollectionDataRequest) SetCollectionData(v string) *DeleteCollectionDataRequest {
	s.CollectionData = &v
	return s
}

func (s *DeleteCollectionDataRequest) SetCollectionDataFilter(v string) *DeleteCollectionDataRequest {
	s.CollectionDataFilter = &v
	return s
}

func (s *DeleteCollectionDataRequest) SetDBInstanceId(v string) *DeleteCollectionDataRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DeleteCollectionDataRequest) SetNamespace(v string) *DeleteCollectionDataRequest {
	s.Namespace = &v
	return s
}

func (s *DeleteCollectionDataRequest) SetNamespacePassword(v string) *DeleteCollectionDataRequest {
	s.NamespacePassword = &v
	return s
}

func (s *DeleteCollectionDataRequest) SetOwnerId(v int64) *DeleteCollectionDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteCollectionDataRequest) SetRegionId(v string) *DeleteCollectionDataRequest {
	s.RegionId = &v
	return s
}

type DeleteCollectionDataResponseBody struct {
	AppliedRows *int64  `json:"AppliedRows,omitempty" xml:"AppliedRows,omitempty"`
	Message     *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DeleteCollectionDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteCollectionDataResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCollectionDataResponseBody) SetAppliedRows(v int64) *DeleteCollectionDataResponseBody {
	s.AppliedRows = &v
	return s
}

func (s *DeleteCollectionDataResponseBody) SetMessage(v string) *DeleteCollectionDataResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteCollectionDataResponseBody) SetRequestId(v string) *DeleteCollectionDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCollectionDataResponseBody) SetStatus(v string) *DeleteCollectionDataResponseBody {
	s.Status = &v
	return s
}

type DeleteCollectionDataResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteCollectionDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteCollectionDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCollectionDataResponse) GoString() string {
	return s.String()
}

func (s *DeleteCollectionDataResponse) SetHeaders(v map[string]*string) *DeleteCollectionDataResponse {
	s.Headers = v
	return s
}

func (s *DeleteCollectionDataResponse) SetStatusCode(v int32) *DeleteCollectionDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCollectionDataResponse) SetBody(v *DeleteCollectionDataResponseBody) *DeleteCollectionDataResponse {
	s.Body = v
	return s
}

type DeleteDBInstanceRequest struct {
	// The client token that is used to ensure the idempotence of the request. For more information, see [How to ensure idempotence](~~327176~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the instance.
	//
	// >  You can call the [DescribeDBInstances](~~86911~~) operation to query the details of all AnalyticDB for PostgreSQL instances in a specific region, including instance IDs.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the resource group to which the instance belongs. For more information about how to obtain the ID of a resource group, see [View basic information of a resource group](~~151181~~).
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DeleteDBInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDBInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeleteDBInstanceRequest) SetClientToken(v string) *DeleteDBInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteDBInstanceRequest) SetDBInstanceId(v string) *DeleteDBInstanceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DeleteDBInstanceRequest) SetOwnerId(v int64) *DeleteDBInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteDBInstanceRequest) SetResourceGroupId(v string) *DeleteDBInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

type DeleteDBInstanceResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDBInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDBInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDBInstanceResponseBody) SetRequestId(v string) *DeleteDBInstanceResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDBInstanceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteDBInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDBInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDBInstanceResponse) GoString() string {
	return s.String()
}

func (s *DeleteDBInstanceResponse) SetHeaders(v map[string]*string) *DeleteDBInstanceResponse {
	s.Headers = v
	return s
}

func (s *DeleteDBInstanceResponse) SetStatusCode(v int32) *DeleteDBInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDBInstanceResponse) SetBody(v *DeleteDBInstanceResponseBody) *DeleteDBInstanceResponse {
	s.Body = v
	return s
}

type DeleteDBInstancePlanRequest struct {
	// The ID of the instance.
	//
	// >  You can call the [DescribeDBInstances](~~86911~~) operation to query the details of all AnalyticDB for PostgreSQL instances in a specific region, including instance IDs.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the plan.
	//
	// >  You can call the [DescribeDBInstancePlans](~~449398~~) operation to query the details of plans, including plan IDs.
	PlanId *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
}

func (s DeleteDBInstancePlanRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDBInstancePlanRequest) GoString() string {
	return s.String()
}

func (s *DeleteDBInstancePlanRequest) SetDBInstanceId(v string) *DeleteDBInstancePlanRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DeleteDBInstancePlanRequest) SetOwnerId(v int64) *DeleteDBInstancePlanRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteDBInstancePlanRequest) SetPlanId(v string) *DeleteDBInstancePlanRequest {
	s.PlanId = &v
	return s
}

type DeleteDBInstancePlanResponseBody struct {
	// The ID of the instance.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The error message returned.
	//
	// This parameter is returned only when the operation fails.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the plan.
	PlanId *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The state of the operation.
	//
	// If the operation is successful, **success** is returned. If the operation fails, this parameter is not returned.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DeleteDBInstancePlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDBInstancePlanResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDBInstancePlanResponseBody) SetDBInstanceId(v string) *DeleteDBInstancePlanResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *DeleteDBInstancePlanResponseBody) SetErrorMessage(v string) *DeleteDBInstancePlanResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteDBInstancePlanResponseBody) SetPlanId(v string) *DeleteDBInstancePlanResponseBody {
	s.PlanId = &v
	return s
}

func (s *DeleteDBInstancePlanResponseBody) SetRequestId(v string) *DeleteDBInstancePlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDBInstancePlanResponseBody) SetStatus(v string) *DeleteDBInstancePlanResponseBody {
	s.Status = &v
	return s
}

type DeleteDBInstancePlanResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteDBInstancePlanResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDBInstancePlanResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDBInstancePlanResponse) GoString() string {
	return s.String()
}

func (s *DeleteDBInstancePlanResponse) SetHeaders(v map[string]*string) *DeleteDBInstancePlanResponse {
	s.Headers = v
	return s
}

func (s *DeleteDBInstancePlanResponse) SetStatusCode(v int32) *DeleteDBInstancePlanResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDBInstancePlanResponse) SetBody(v *DeleteDBInstancePlanResponseBody) *DeleteDBInstancePlanResponse {
	s.Body = v
	return s
}

type DeleteNamespaceRequest struct {
	DBInstanceId           *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	ManagerAccount         *string `json:"ManagerAccount,omitempty" xml:"ManagerAccount,omitempty"`
	ManagerAccountPassword *string `json:"ManagerAccountPassword,omitempty" xml:"ManagerAccountPassword,omitempty"`
	Namespace              *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	OwnerId                *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId               *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteNamespaceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteNamespaceRequest) GoString() string {
	return s.String()
}

func (s *DeleteNamespaceRequest) SetDBInstanceId(v string) *DeleteNamespaceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DeleteNamespaceRequest) SetManagerAccount(v string) *DeleteNamespaceRequest {
	s.ManagerAccount = &v
	return s
}

func (s *DeleteNamespaceRequest) SetManagerAccountPassword(v string) *DeleteNamespaceRequest {
	s.ManagerAccountPassword = &v
	return s
}

func (s *DeleteNamespaceRequest) SetNamespace(v string) *DeleteNamespaceRequest {
	s.Namespace = &v
	return s
}

func (s *DeleteNamespaceRequest) SetOwnerId(v int64) *DeleteNamespaceRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteNamespaceRequest) SetRegionId(v string) *DeleteNamespaceRequest {
	s.RegionId = &v
	return s
}

type DeleteNamespaceResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DeleteNamespaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteNamespaceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteNamespaceResponseBody) SetMessage(v string) *DeleteNamespaceResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteNamespaceResponseBody) SetRequestId(v string) *DeleteNamespaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteNamespaceResponseBody) SetStatus(v string) *DeleteNamespaceResponseBody {
	s.Status = &v
	return s
}

type DeleteNamespaceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteNamespaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteNamespaceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteNamespaceResponse) GoString() string {
	return s.String()
}

func (s *DeleteNamespaceResponse) SetHeaders(v map[string]*string) *DeleteNamespaceResponse {
	s.Headers = v
	return s
}

func (s *DeleteNamespaceResponse) SetStatusCode(v int32) *DeleteNamespaceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteNamespaceResponse) SetBody(v *DeleteNamespaceResponseBody) *DeleteNamespaceResponse {
	s.Body = v
	return s
}

type DeleteVectorIndexRequest struct {
	Collection             *string `json:"Collection,omitempty" xml:"Collection,omitempty"`
	DBInstanceId           *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	ManagerAccount         *string `json:"ManagerAccount,omitempty" xml:"ManagerAccount,omitempty"`
	ManagerAccountPassword *string `json:"ManagerAccountPassword,omitempty" xml:"ManagerAccountPassword,omitempty"`
	Namespace              *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	OwnerId                *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId               *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteVectorIndexRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteVectorIndexRequest) GoString() string {
	return s.String()
}

func (s *DeleteVectorIndexRequest) SetCollection(v string) *DeleteVectorIndexRequest {
	s.Collection = &v
	return s
}

func (s *DeleteVectorIndexRequest) SetDBInstanceId(v string) *DeleteVectorIndexRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DeleteVectorIndexRequest) SetManagerAccount(v string) *DeleteVectorIndexRequest {
	s.ManagerAccount = &v
	return s
}

func (s *DeleteVectorIndexRequest) SetManagerAccountPassword(v string) *DeleteVectorIndexRequest {
	s.ManagerAccountPassword = &v
	return s
}

func (s *DeleteVectorIndexRequest) SetNamespace(v string) *DeleteVectorIndexRequest {
	s.Namespace = &v
	return s
}

func (s *DeleteVectorIndexRequest) SetOwnerId(v int64) *DeleteVectorIndexRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteVectorIndexRequest) SetRegionId(v string) *DeleteVectorIndexRequest {
	s.RegionId = &v
	return s
}

type DeleteVectorIndexResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DeleteVectorIndexResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteVectorIndexResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVectorIndexResponseBody) SetMessage(v string) *DeleteVectorIndexResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteVectorIndexResponseBody) SetRequestId(v string) *DeleteVectorIndexResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteVectorIndexResponseBody) SetStatus(v string) *DeleteVectorIndexResponseBody {
	s.Status = &v
	return s
}

type DeleteVectorIndexResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteVectorIndexResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteVectorIndexResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteVectorIndexResponse) GoString() string {
	return s.String()
}

func (s *DeleteVectorIndexResponse) SetHeaders(v map[string]*string) *DeleteVectorIndexResponse {
	s.Headers = v
	return s
}

func (s *DeleteVectorIndexResponse) SetStatusCode(v int32) *DeleteVectorIndexResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVectorIndexResponse) SetBody(v *DeleteVectorIndexResponseBody) *DeleteVectorIndexResponse {
	s.Body = v
	return s
}

type DescribeAccountsRequest struct {
	// The name of the account.
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The ID of the instance.
	//
	// >  You can call the [DescribeDBInstances](~~86911~~) operation to query the details of all AnalyticDB for PostgreSQL instances in a specific region, including instance IDs.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s DescribeAccountsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAccountsRequest) SetAccountName(v string) *DescribeAccountsRequest {
	s.AccountName = &v
	return s
}

func (s *DescribeAccountsRequest) SetDBInstanceId(v string) *DescribeAccountsRequest {
	s.DBInstanceId = &v
	return s
}

type DescribeAccountsResponseBody struct {
	// Details of the account.
	Accounts *DescribeAccountsResponseBodyAccounts `json:"Accounts,omitempty" xml:"Accounts,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAccountsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAccountsResponseBody) SetAccounts(v *DescribeAccountsResponseBodyAccounts) *DescribeAccountsResponseBody {
	s.Accounts = v
	return s
}

func (s *DescribeAccountsResponseBody) SetRequestId(v string) *DescribeAccountsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeAccountsResponseBodyAccounts struct {
	DBInstanceAccount []*DescribeAccountsResponseBodyAccountsDBInstanceAccount `json:"DBInstanceAccount,omitempty" xml:"DBInstanceAccount,omitempty" type:"Repeated"`
}

func (s DescribeAccountsResponseBodyAccounts) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountsResponseBodyAccounts) GoString() string {
	return s.String()
}

func (s *DescribeAccountsResponseBodyAccounts) SetDBInstanceAccount(v []*DescribeAccountsResponseBodyAccountsDBInstanceAccount) *DescribeAccountsResponseBodyAccounts {
	s.DBInstanceAccount = v
	return s
}

type DescribeAccountsResponseBodyAccountsDBInstanceAccount struct {
	// The description of the account.
	AccountDescription *string `json:"AccountDescription,omitempty" xml:"AccountDescription,omitempty"`
	// The name of the account.
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The state of the account.
	//
	// *   **0**: The account is being created.
	// *   **1**: The account is in use.
	// *   **3**: The account is being deleted.
	AccountStatus *string `json:"AccountStatus,omitempty" xml:"AccountStatus,omitempty"`
	// The ID of the instance.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s DescribeAccountsResponseBodyAccountsDBInstanceAccount) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountsResponseBodyAccountsDBInstanceAccount) GoString() string {
	return s.String()
}

func (s *DescribeAccountsResponseBodyAccountsDBInstanceAccount) SetAccountDescription(v string) *DescribeAccountsResponseBodyAccountsDBInstanceAccount {
	s.AccountDescription = &v
	return s
}

func (s *DescribeAccountsResponseBodyAccountsDBInstanceAccount) SetAccountName(v string) *DescribeAccountsResponseBodyAccountsDBInstanceAccount {
	s.AccountName = &v
	return s
}

func (s *DescribeAccountsResponseBodyAccountsDBInstanceAccount) SetAccountStatus(v string) *DescribeAccountsResponseBodyAccountsDBInstanceAccount {
	s.AccountStatus = &v
	return s
}

func (s *DescribeAccountsResponseBodyAccountsDBInstanceAccount) SetDBInstanceId(v string) *DescribeAccountsResponseBodyAccountsDBInstanceAccount {
	s.DBInstanceId = &v
	return s
}

type DescribeAccountsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeAccountsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAccountsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAccountsResponse) SetHeaders(v map[string]*string) *DescribeAccountsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAccountsResponse) SetStatusCode(v int32) *DescribeAccountsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAccountsResponse) SetBody(v *DescribeAccountsResponseBody) *DescribeAccountsResponse {
	s.Body = v
	return s
}

type DescribeAvailableResourcesRequest struct {
	// The billing method. Valid values:
	//
	// *   **Postpaid**: pay-as-you-go
	// *   **Prepaid**: subscription
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The ID of the region.
	//
	// >  You can call the [DescribeRegions](~~86912~~) operation to query the most recent region list.
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The ID of the zone.
	//
	// >  You can call the [DescribeRegions](~~86912~~) operation to query the most recent zone list.
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeAvailableResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourcesRequest) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourcesRequest) SetChargeType(v string) *DescribeAvailableResourcesRequest {
	s.ChargeType = &v
	return s
}

func (s *DescribeAvailableResourcesRequest) SetRegion(v string) *DescribeAvailableResourcesRequest {
	s.Region = &v
	return s
}

func (s *DescribeAvailableResourcesRequest) SetZoneId(v string) *DescribeAvailableResourcesRequest {
	s.ZoneId = &v
	return s
}

type DescribeAvailableResourcesResponseBody struct {
	// The ID of the region.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Details of the available resources.
	Resources []*DescribeAvailableResourcesResponseBodyResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
}

func (s DescribeAvailableResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourcesResponseBody) SetRegionId(v string) *DescribeAvailableResourcesResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeAvailableResourcesResponseBody) SetRequestId(v string) *DescribeAvailableResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAvailableResourcesResponseBody) SetResources(v []*DescribeAvailableResourcesResponseBodyResources) *DescribeAvailableResourcesResponseBody {
	s.Resources = v
	return s
}

type DescribeAvailableResourcesResponseBodyResources struct {
	// The available engine version and specifications.
	SupportedEngines []*DescribeAvailableResourcesResponseBodyResourcesSupportedEngines `json:"SupportedEngines,omitempty" xml:"SupportedEngines,omitempty" type:"Repeated"`
	// The ID of the zone.
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeAvailableResourcesResponseBodyResources) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourcesResponseBodyResources) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourcesResponseBodyResources) SetSupportedEngines(v []*DescribeAvailableResourcesResponseBodyResourcesSupportedEngines) *DescribeAvailableResourcesResponseBodyResources {
	s.SupportedEngines = v
	return s
}

func (s *DescribeAvailableResourcesResponseBodyResources) SetZoneId(v string) *DescribeAvailableResourcesResponseBodyResources {
	s.ZoneId = &v
	return s
}

type DescribeAvailableResourcesResponseBodyResourcesSupportedEngines struct {
	// The instance resource type. Valid values:
	//
	// *   **ecs**: elastic storage mode
	// *   **serverless**: Serverless mode
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The available engine version.
	SupportedEngineVersion *string `json:"SupportedEngineVersion,omitempty" xml:"SupportedEngineVersion,omitempty"`
	// The available specifications.
	SupportedInstanceClasses []*DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClasses `json:"SupportedInstanceClasses,omitempty" xml:"SupportedInstanceClasses,omitempty" type:"Repeated"`
}

func (s DescribeAvailableResourcesResponseBodyResourcesSupportedEngines) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourcesResponseBodyResourcesSupportedEngines) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourcesResponseBodyResourcesSupportedEngines) SetMode(v string) *DescribeAvailableResourcesResponseBodyResourcesSupportedEngines {
	s.Mode = &v
	return s
}

func (s *DescribeAvailableResourcesResponseBodyResourcesSupportedEngines) SetSupportedEngineVersion(v string) *DescribeAvailableResourcesResponseBodyResourcesSupportedEngines {
	s.SupportedEngineVersion = &v
	return s
}

func (s *DescribeAvailableResourcesResponseBodyResourcesSupportedEngines) SetSupportedInstanceClasses(v []*DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClasses) *DescribeAvailableResourcesResponseBodyResourcesSupportedEngines {
	s.SupportedInstanceClasses = v
	return s
}

type DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClasses struct {
	// The instance edition. Valid values:
	//
	// *   **HighAvailability**: High-availability Edition
	// *   **Basic**: Basic Edition
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The description of compute node specifications.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The specifications of each compute node.
	DisplayClass *string `json:"DisplayClass,omitempty" xml:"DisplayClass,omitempty"`
	// The specifications of each compute node.
	InstanceClass *string `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	// Details about the compute nodes.
	NodeCount *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesNodeCount `json:"NodeCount,omitempty" xml:"NodeCount,omitempty" type:"Struct"`
	// Details about the storage capacity of compute nodes.
	StorageSize *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesStorageSize `json:"StorageSize,omitempty" xml:"StorageSize,omitempty" type:"Struct"`
	// The storage type. Valid values:
	//
	// *   **cloud_essd**: enhanced SSD (ESSD)
	// *   **cloud_efficiency**: ultra disk
	// *   **oss**: Object Storage Service (OSS)
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
}

func (s DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClasses) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClasses) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClasses) SetCategory(v string) *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClasses {
	s.Category = &v
	return s
}

func (s *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClasses) SetDescription(v string) *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClasses {
	s.Description = &v
	return s
}

func (s *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClasses) SetDisplayClass(v string) *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClasses {
	s.DisplayClass = &v
	return s
}

func (s *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClasses) SetInstanceClass(v string) *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClasses {
	s.InstanceClass = &v
	return s
}

func (s *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClasses) SetNodeCount(v *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesNodeCount) *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClasses {
	s.NodeCount = v
	return s
}

func (s *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClasses) SetStorageSize(v *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesStorageSize) *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClasses {
	s.StorageSize = v
	return s
}

func (s *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClasses) SetStorageType(v string) *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClasses {
	s.StorageType = &v
	return s
}

type DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesNodeCount struct {
	// The maximum number of compute nodes.
	MaxCount *string `json:"MaxCount,omitempty" xml:"MaxCount,omitempty"`
	// The minimum number of compute nodes.
	MinCount *string `json:"MinCount,omitempty" xml:"MinCount,omitempty"`
	// The step size for adding compute nodes.
	//
	// For example, if the value of this parameter is 4, compute nodes must be added by multiples of 4.
	Step *string `json:"Step,omitempty" xml:"Step,omitempty"`
}

func (s DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesNodeCount) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesNodeCount) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesNodeCount) SetMaxCount(v string) *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesNodeCount {
	s.MaxCount = &v
	return s
}

func (s *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesNodeCount) SetMinCount(v string) *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesNodeCount {
	s.MinCount = &v
	return s
}

func (s *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesNodeCount) SetStep(v string) *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesNodeCount {
	s.Step = &v
	return s
}

type DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesStorageSize struct {
	// The maximum storage capacity of each compute node.
	MaxCount *string `json:"MaxCount,omitempty" xml:"MaxCount,omitempty"`
	// The minimum storage capacity of each compute node.
	MinCount *string `json:"MinCount,omitempty" xml:"MinCount,omitempty"`
	// The step size for adding storage capacity for compute nodes.
	Step *string `json:"Step,omitempty" xml:"Step,omitempty"`
}

func (s DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesStorageSize) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesStorageSize) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesStorageSize) SetMaxCount(v string) *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesStorageSize {
	s.MaxCount = &v
	return s
}

func (s *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesStorageSize) SetMinCount(v string) *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesStorageSize {
	s.MinCount = &v
	return s
}

func (s *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesStorageSize) SetStep(v string) *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesStorageSize {
	s.Step = &v
	return s
}

type DescribeAvailableResourcesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeAvailableResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAvailableResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourcesResponse) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourcesResponse) SetHeaders(v map[string]*string) *DescribeAvailableResourcesResponse {
	s.Headers = v
	return s
}

func (s *DescribeAvailableResourcesResponse) SetStatusCode(v int32) *DescribeAvailableResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAvailableResourcesResponse) SetBody(v *DescribeAvailableResourcesResponseBody) *DescribeAvailableResourcesResponse {
	s.Body = v
	return s
}

type DescribeBackupPolicyRequest struct {
	// The ID of the instance.
	//
	// >  You can call the [DescribeDBInstances](~~86911~~) operation to query the details of all AnalyticDB for PostgreSQL instances in a specific region, including instance IDs.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s DescribeBackupPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupPolicyRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupPolicyRequest) SetDBInstanceId(v string) *DescribeBackupPolicyRequest {
	s.DBInstanceId = &v
	return s
}

type DescribeBackupPolicyResponseBody struct {
	// The number of days for which data backup files are retained.
	BackupRetentionPeriod *int32 `json:"BackupRetentionPeriod,omitempty" xml:"BackupRetentionPeriod,omitempty"`
	// Indicates whether automatic point-in-time backup is enabled. Valid values:
	//
	// *   **true**: Automatic point-in-time backup is enabled.
	// *   **false**: Automatic point-in-time backup is disabled.
	EnableRecoveryPoint *bool `json:"EnableRecoveryPoint,omitempty" xml:"EnableRecoveryPoint,omitempty"`
	// The cycle based on which backups are performed. If more than one day of the week is specified, the days of the week are separated by commas (,). Valid values:
	//
	// *   **Monday**
	// *   **Tuesday**
	// *   **Wednesday**
	// *   **Thursday**
	// *   **Friday**
	// *   **Saturday**
	// *   **Sunday**
	PreferredBackupPeriod *string `json:"PreferredBackupPeriod,omitempty" xml:"PreferredBackupPeriod,omitempty"`
	// The backup time. The time is in the HH:mmZ-HH:mmZ format. The time is displayed in UTC.
	PreferredBackupTime *string `json:"PreferredBackupTime,omitempty" xml:"PreferredBackupTime,omitempty"`
	// The frequency of the point-in-time backup. Valid values:
	//
	// *   **1**: per hour
	// *   **2**: per 2 hours
	// *   **4**: per 4 hours
	// *   **8**: per 8 hours
	RecoveryPointPeriod *string `json:"RecoveryPointPeriod,omitempty" xml:"RecoveryPointPeriod,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeBackupPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupPolicyResponseBody) SetBackupRetentionPeriod(v int32) *DescribeBackupPolicyResponseBody {
	s.BackupRetentionPeriod = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetEnableRecoveryPoint(v bool) *DescribeBackupPolicyResponseBody {
	s.EnableRecoveryPoint = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetPreferredBackupPeriod(v string) *DescribeBackupPolicyResponseBody {
	s.PreferredBackupPeriod = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetPreferredBackupTime(v string) *DescribeBackupPolicyResponseBody {
	s.PreferredBackupTime = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetRecoveryPointPeriod(v string) *DescribeBackupPolicyResponseBody {
	s.RecoveryPointPeriod = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetRequestId(v string) *DescribeBackupPolicyResponseBody {
	s.RequestId = &v
	return s
}

type DescribeBackupPolicyResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeBackupPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeBackupPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupPolicyResponse) GoString() string {
	return s.String()
}

func (s *DescribeBackupPolicyResponse) SetHeaders(v map[string]*string) *DescribeBackupPolicyResponse {
	s.Headers = v
	return s
}

func (s *DescribeBackupPolicyResponse) SetStatusCode(v int32) *DescribeBackupPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBackupPolicyResponse) SetBody(v *DescribeBackupPolicyResponseBody) *DescribeBackupPolicyResponse {
	s.Body = v
	return s
}

type DescribeCollectionRequest struct {
	Collection        *string `json:"Collection,omitempty" xml:"Collection,omitempty"`
	DBInstanceId      *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	Namespace         *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	NamespacePassword *string `json:"NamespacePassword,omitempty" xml:"NamespacePassword,omitempty"`
	OwnerId           *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeCollectionRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCollectionRequest) GoString() string {
	return s.String()
}

func (s *DescribeCollectionRequest) SetCollection(v string) *DescribeCollectionRequest {
	s.Collection = &v
	return s
}

func (s *DescribeCollectionRequest) SetDBInstanceId(v string) *DescribeCollectionRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeCollectionRequest) SetNamespace(v string) *DescribeCollectionRequest {
	s.Namespace = &v
	return s
}

func (s *DescribeCollectionRequest) SetNamespacePassword(v string) *DescribeCollectionRequest {
	s.NamespacePassword = &v
	return s
}

func (s *DescribeCollectionRequest) SetOwnerId(v int64) *DescribeCollectionRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCollectionRequest) SetRegionId(v string) *DescribeCollectionRequest {
	s.RegionId = &v
	return s
}

type DescribeCollectionResponseBody struct {
	DBInstanceId            *string            `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	Dimension               *int32             `json:"Dimension,omitempty" xml:"Dimension,omitempty"`
	FullTextRetrievalFields *string            `json:"FullTextRetrievalFields,omitempty" xml:"FullTextRetrievalFields,omitempty"`
	Message                 *string            `json:"Message,omitempty" xml:"Message,omitempty"`
	Metadata                map[string]*string `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// Distance Metrics。
	Metrics   *string `json:"Metrics,omitempty" xml:"Metrics,omitempty"`
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	Parser    *string `json:"Parser,omitempty" xml:"Parser,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeCollectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCollectionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCollectionResponseBody) SetDBInstanceId(v string) *DescribeCollectionResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeCollectionResponseBody) SetDimension(v int32) *DescribeCollectionResponseBody {
	s.Dimension = &v
	return s
}

func (s *DescribeCollectionResponseBody) SetFullTextRetrievalFields(v string) *DescribeCollectionResponseBody {
	s.FullTextRetrievalFields = &v
	return s
}

func (s *DescribeCollectionResponseBody) SetMessage(v string) *DescribeCollectionResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeCollectionResponseBody) SetMetadata(v map[string]*string) *DescribeCollectionResponseBody {
	s.Metadata = v
	return s
}

func (s *DescribeCollectionResponseBody) SetMetrics(v string) *DescribeCollectionResponseBody {
	s.Metrics = &v
	return s
}

func (s *DescribeCollectionResponseBody) SetNamespace(v string) *DescribeCollectionResponseBody {
	s.Namespace = &v
	return s
}

func (s *DescribeCollectionResponseBody) SetParser(v string) *DescribeCollectionResponseBody {
	s.Parser = &v
	return s
}

func (s *DescribeCollectionResponseBody) SetRegionId(v string) *DescribeCollectionResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeCollectionResponseBody) SetRequestId(v string) *DescribeCollectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCollectionResponseBody) SetStatus(v string) *DescribeCollectionResponseBody {
	s.Status = &v
	return s
}

type DescribeCollectionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeCollectionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCollectionResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCollectionResponse) GoString() string {
	return s.String()
}

func (s *DescribeCollectionResponse) SetHeaders(v map[string]*string) *DescribeCollectionResponse {
	s.Headers = v
	return s
}

func (s *DescribeCollectionResponse) SetStatusCode(v int32) *DescribeCollectionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCollectionResponse) SetBody(v *DescribeCollectionResponseBody) *DescribeCollectionResponse {
	s.Body = v
	return s
}

type DescribeDBClusterNodeRequest struct {
	// The ID of the instance.
	//
	// >  You can call the [DescribeDBInstances](~~86911~~) operation to query the details of all AnalyticDB for PostgreSQL instances in a specific region, including instance IDs.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The node type. Valid values:
	//
	// *   **master**: coordinator node
	// *   **segment**: compute node
	//
	// >  If you do not specify this parameter, the information of all nodes is returned.
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
}

func (s DescribeDBClusterNodeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterNodeRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterNodeRequest) SetDBInstanceId(v string) *DescribeDBClusterNodeRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBClusterNodeRequest) SetNodeType(v string) *DescribeDBClusterNodeRequest {
	s.NodeType = &v
	return s
}

type DescribeDBClusterNodeResponseBody struct {
	// The ID of the instance.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The information of nodes.
	Nodes []*DescribeDBClusterNodeResponseBodyNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBClusterNodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterNodeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterNodeResponseBody) SetDBClusterId(v string) *DescribeDBClusterNodeResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterNodeResponseBody) SetNodes(v []*DescribeDBClusterNodeResponseBodyNodes) *DescribeDBClusterNodeResponseBody {
	s.Nodes = v
	return s
}

func (s *DescribeDBClusterNodeResponseBody) SetRequestId(v string) *DescribeDBClusterNodeResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDBClusterNodeResponseBodyNodes struct {
	// The name of the node.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeDBClusterNodeResponseBodyNodes) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterNodeResponseBodyNodes) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterNodeResponseBodyNodes) SetName(v string) *DescribeDBClusterNodeResponseBodyNodes {
	s.Name = &v
	return s
}

type DescribeDBClusterNodeResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDBClusterNodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDBClusterNodeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterNodeResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterNodeResponse) SetHeaders(v map[string]*string) *DescribeDBClusterNodeResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBClusterNodeResponse) SetStatusCode(v int32) *DescribeDBClusterNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBClusterNodeResponse) SetBody(v *DescribeDBClusterNodeResponseBody) *DescribeDBClusterNodeResponse {
	s.Body = v
	return s
}

type DescribeDBClusterPerformanceRequest struct {
	// The ID of the instance.
	//
	// >  You can call the [DescribeDBInstances](~~86911~~) operation to query the details of all AnalyticDB for PostgreSQL instances in a specific region, including instance IDs.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the `YYYY-MM-DDTHH:mmZ` format.
	//
	// >  The end time must be later than the start time. The interval cannot be more than seven days.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The performance metric that you want to query. Separate multiple values with commas (,). For more information, see [Performance parameters](~~86943~~).
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The node type. Valid values:
	//
	// *   **master**: coordinator node
	// *   **segment**: compute node
	//
	// >  If you do not specify this parameter, the performance metrics of all nodes are returned.
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	// The nodes for which you want to query performance metrics. Separate multiple values with commas (,). Example: `master-10******1,master-10******2`. You can call the [DescribeDBClusterNode](~~390136~~) operation to query the names of nodes.
	//
	// The nodes can also be filtered based on their metric values. Valid values:
	//
	// *   **top10**: the 10 nodes that have the highest metric values
	// *   **top20**: the 20 nodes that have the highest metric values
	// *   **bottom10**: the 10 nodes that have the lowest metric values
	// *   **bottom20**: the 20 nodes that have the lowest metric values
	Nodes *string `json:"Nodes,omitempty" xml:"Nodes,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the `YYYY-MM-DDTHH:mmZ` format.
	//
	// >  You can query monitoring information only within the last 30 days.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDBClusterPerformanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterPerformanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterPerformanceRequest) SetDBInstanceId(v string) *DescribeDBClusterPerformanceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBClusterPerformanceRequest) SetEndTime(v string) *DescribeDBClusterPerformanceRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDBClusterPerformanceRequest) SetKey(v string) *DescribeDBClusterPerformanceRequest {
	s.Key = &v
	return s
}

func (s *DescribeDBClusterPerformanceRequest) SetNodeType(v string) *DescribeDBClusterPerformanceRequest {
	s.NodeType = &v
	return s
}

func (s *DescribeDBClusterPerformanceRequest) SetNodes(v string) *DescribeDBClusterPerformanceRequest {
	s.Nodes = &v
	return s
}

func (s *DescribeDBClusterPerformanceRequest) SetStartTime(v string) *DescribeDBClusterPerformanceRequest {
	s.StartTime = &v
	return s
}

type DescribeDBClusterPerformanceResponseBody struct {
	// The ID of the instance.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The end time of the query. The time follows the ISO 8601 standard in the `YYYY-MM-DDTHH:mmZ` format. The time is displayed in UTC.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Details of the performance metrics of the instance.
	PerformanceKeys []*DescribeDBClusterPerformanceResponseBodyPerformanceKeys `json:"PerformanceKeys,omitempty" xml:"PerformanceKeys,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The start time of the query. The time follows the ISO 8601 standard in the `YYYY-MM-DDTHH:mmZ` format. The time is displayed in UTC.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDBClusterPerformanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterPerformanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterPerformanceResponseBody) SetDBClusterId(v string) *DescribeDBClusterPerformanceResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBody) SetEndTime(v string) *DescribeDBClusterPerformanceResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBody) SetPerformanceKeys(v []*DescribeDBClusterPerformanceResponseBodyPerformanceKeys) *DescribeDBClusterPerformanceResponseBody {
	s.PerformanceKeys = v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBody) SetRequestId(v string) *DescribeDBClusterPerformanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBody) SetStartTime(v string) *DescribeDBClusterPerformanceResponseBody {
	s.StartTime = &v
	return s
}

type DescribeDBClusterPerformanceResponseBodyPerformanceKeys struct {
	// The name of the performance metric. For more information, see [Performance parameters](~~86943~~).
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Details of the performance metric of a node.
	Series []*DescribeDBClusterPerformanceResponseBodyPerformanceKeysSeries `json:"Series,omitempty" xml:"Series,omitempty" type:"Repeated"`
	// The unit of the performance metric.
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
}

func (s DescribeDBClusterPerformanceResponseBodyPerformanceKeys) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterPerformanceResponseBodyPerformanceKeys) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformanceKeys) SetName(v string) *DescribeDBClusterPerformanceResponseBodyPerformanceKeys {
	s.Name = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformanceKeys) SetSeries(v []*DescribeDBClusterPerformanceResponseBodyPerformanceKeysSeries) *DescribeDBClusterPerformanceResponseBodyPerformanceKeys {
	s.Series = v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformanceKeys) SetUnit(v string) *DescribeDBClusterPerformanceResponseBodyPerformanceKeys {
	s.Unit = &v
	return s
}

type DescribeDBClusterPerformanceResponseBodyPerformanceKeysSeries struct {
	// The name of the compute node or compute group.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The role of the node. Valid values:
	//
	// *   **master**: primary coordinator node
	// *   **standby**: standby coordinator node
	// *   **segment**: compute node
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// The value of the performance metric collected at a point in time.
	Values []*DescribeDBClusterPerformanceResponseBodyPerformanceKeysSeriesValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s DescribeDBClusterPerformanceResponseBodyPerformanceKeysSeries) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterPerformanceResponseBodyPerformanceKeysSeries) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformanceKeysSeries) SetName(v string) *DescribeDBClusterPerformanceResponseBodyPerformanceKeysSeries {
	s.Name = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformanceKeysSeries) SetRole(v string) *DescribeDBClusterPerformanceResponseBodyPerformanceKeysSeries {
	s.Role = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformanceKeysSeries) SetValues(v []*DescribeDBClusterPerformanceResponseBodyPerformanceKeysSeriesValues) *DescribeDBClusterPerformanceResponseBodyPerformanceKeysSeries {
	s.Values = v
	return s
}

type DescribeDBClusterPerformanceResponseBodyPerformanceKeysSeriesValues struct {
	// The value of the performance metric and the time when the metric value was collected.
	Point []*string `json:"Point,omitempty" xml:"Point,omitempty" type:"Repeated"`
}

func (s DescribeDBClusterPerformanceResponseBodyPerformanceKeysSeriesValues) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterPerformanceResponseBodyPerformanceKeysSeriesValues) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformanceKeysSeriesValues) SetPoint(v []*string) *DescribeDBClusterPerformanceResponseBodyPerformanceKeysSeriesValues {
	s.Point = v
	return s
}

type DescribeDBClusterPerformanceResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDBClusterPerformanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDBClusterPerformanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterPerformanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterPerformanceResponse) SetHeaders(v map[string]*string) *DescribeDBClusterPerformanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBClusterPerformanceResponse) SetStatusCode(v int32) *DescribeDBClusterPerformanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponse) SetBody(v *DescribeDBClusterPerformanceResponseBody) *DescribeDBClusterPerformanceResponse {
	s.Body = v
	return s
}

type DescribeDBInstanceAttributeRequest struct {
	// The instance ID.
	//
	// >  You can call the [DescribeDBInstances](~~86911~~) operation to query the IDs of all AnalyticDB for PostgreSQL instances in a region.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the resource group to which the instance belongs. For more information about how to obtain the ID of a resource group, see [View basic information of a resource group](~~151181~~).
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeDBInstanceAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeRequest) SetDBInstanceId(v string) *DescribeDBInstanceAttributeRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceAttributeRequest) SetOwnerId(v int64) *DescribeDBInstanceAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBInstanceAttributeRequest) SetResourceGroupId(v string) *DescribeDBInstanceAttributeRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeDBInstanceAttributeResponseBody struct {
	// The queried instance.
	Items *DescribeDBInstanceAttributeResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBInstanceAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponseBody) SetItems(v *DescribeDBInstanceAttributeResponseBodyItems) *DescribeDBInstanceAttributeResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBody) SetRequestId(v string) *DescribeDBInstanceAttributeResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDBInstanceAttributeResponseBodyItems struct {
	DBInstanceAttribute []*DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute `json:"DBInstanceAttribute,omitempty" xml:"DBInstanceAttribute,omitempty" type:"Repeated"`
}

func (s DescribeDBInstanceAttributeResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponseBodyItems) SetDBInstanceAttribute(v []*DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) *DescribeDBInstanceAttributeResponseBodyItems {
	s.DBInstanceAttribute = v
	return s
}

type DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute struct {
	// The service availability of the instance. Unit: %.
	//
	// >  This parameter is available only for instances in reserved storage mode.
	AvailabilityValue *string `json:"AvailabilityValue,omitempty" xml:"AvailabilityValue,omitempty"`
	// The access mode of the instance. Valid values:
	//
	// *   **Performance**: standard mode.
	// *   **Safety**: safe mode.
	// *   **LVS**: Linux Virtual Server (LVS) mode.
	ConnectionMode *string `json:"ConnectionMode,omitempty" xml:"ConnectionMode,omitempty"`
	// The endpoint that is used to connect to the instance.
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// The number of the minor version.
	CoreVersion *string `json:"CoreVersion,omitempty" xml:"CoreVersion,omitempty"`
	// The number of CPU cores per compute node.
	CpuCores *int32 `json:"CpuCores,omitempty" xml:"CpuCores,omitempty"`
	// The number of CPU cores per node.
	//
	// >  This parameter is available only for instances in reserved storage mode.
	CpuCoresPerNode *int32 `json:"CpuCoresPerNode,omitempty" xml:"CpuCoresPerNode,omitempty"`
	// The time when the instance was created.
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The edition of the instance. Valid values:
	//
	// *   **Basic**: Basic Edition.
	// *   **HighAvailability**: High-availability Edition.
	DBInstanceCategory *string `json:"DBInstanceCategory,omitempty" xml:"DBInstanceCategory,omitempty"`
	// The instance type of the instance.
	//
	// >  This parameter is available only for instances in reserved storage mode.
	DBInstanceClass *string `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty"`
	// The instance family of the instance. Valid values:
	//
	// *   **s**: shared.
	// *   **x**: general-purpose.
	// *   **d**: dedicated.
	// *   **h**: dedicated host.
	DBInstanceClassType *string `json:"DBInstanceClassType,omitempty" xml:"DBInstanceClassType,omitempty"`
	// The number of CPU cores.
	DBInstanceCpuCores *int32 `json:"DBInstanceCpuCores,omitempty" xml:"DBInstanceCpuCores,omitempty"`
	// The description of the instance.
	DBInstanceDescription *string `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	// The maximum disk throughput of the compute group. Unit: Mbit/s.
	//
	// >  This parameter is available only for instances in reserved storage mode.
	DBInstanceDiskMBPS *int64 `json:"DBInstanceDiskMBPS,omitempty" xml:"DBInstanceDiskMBPS,omitempty"`
	// The number of compute groups.
	//
	// >  This parameter is available only for instances in reserved storage mode.
	DBInstanceGroupCount *string `json:"DBInstanceGroupCount,omitempty" xml:"DBInstanceGroupCount,omitempty"`
	// The instance ID.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The memory capacity per compute node.
	//
	// >  For instances in reserved storage mode, the unit of this parameter is MB. For instances in elastic storage mode and Serverless mode, the unit of this parameter is GB.
	DBInstanceMemory *int64 `json:"DBInstanceMemory,omitempty" xml:"DBInstanceMemory,omitempty"`
	// The resource type of the instance. Valid values:
	//
	// *   **Serverless**: Serverless mode.
	// *   **StorageElastic**: elastic storage mode.
	// *   **Classic**: reserved storage mode.
	DBInstanceMode *string `json:"DBInstanceMode,omitempty" xml:"DBInstanceMode,omitempty"`
	// The type of the network interface card (NIC) that is used by the instance. Valid values:
	//
	// *   **0**: Internet.
	// *   **1**: internal network.
	// *   **2**: VPC.
	DBInstanceNetType *string `json:"DBInstanceNetType,omitempty" xml:"DBInstanceNetType,omitempty"`
	// The state of the instance. For more information, see [Instance statuses](~~86944~~).
	DBInstanceStatus *string `json:"DBInstanceStatus,omitempty" xml:"DBInstanceStatus,omitempty"`
	// The maximum storage capacity per node. Unit: GB.
	DBInstanceStorage *int64 `json:"DBInstanceStorage,omitempty" xml:"DBInstanceStorage,omitempty"`
	// The encryption key.
	//
	// >  This parameter is returned only for instances for which disk encryption is enabled.
	EncryptionKey *string `json:"EncryptionKey,omitempty" xml:"EncryptionKey,omitempty"`
	// The encryption type. Valid values:
	//
	// *   **CloudDisk**: disk encryption.
	//
	// >  This parameter is returned only for instances for which disk encryption is enabled.
	EncryptionType *string `json:"EncryptionType,omitempty" xml:"EncryptionType,omitempty"`
	// The database engine of the instance.
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The version of the database engine.
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The expiration time of the instance. The time is displayed in UTC.
	//
	// >  For pay-as-you-go instances, `2999-09-08T16:00:00Z` is returned.
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The disk type of the compute group. Valid values:
	//
	// *   **0**: SSD.
	// *   **1**: HDD.
	//
	// >  This parameter is available only for instances in reserved storage mode.
	HostType *string `json:"HostType,omitempty" xml:"HostType,omitempty"`
	// The wait period for the instance that has no traffic to become idle. Unit: seconds.
	//
	// >  This parameter is returned only for instances in automatic Serverless mode.
	IdleTime *int32 `json:"IdleTime,omitempty" xml:"IdleTime,omitempty"`
	// The network type of the instance. Valid values:
	//
	// *   **Classic**: classic network.
	// *   **VPC**: VPC.
	InstanceNetworkType *string `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
	// The lock mode of the instance. Valid values:
	//
	// *   **Unlock**: The instance is not locked.
	// *   **ManualLock**: The cluster is manually locked.
	// *   **LockByExpiration**: The instance is automatically locked due to instance expiration.
	// *   **LockByRestoration**: The instance is automatically locked due to instance restoration.
	// *   **LockByDiskQuota**: The instance is automatically locked due to exhausted storage.
	LockMode *string `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	// An invalid parameter. It is no longer returned when you call this operation.
	LockReason *string `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	// The end time of the maintenance window.
	MaintainEndTime *string `json:"MaintainEndTime,omitempty" xml:"MaintainEndTime,omitempty"`
	// The start time of the maintenance window.
	MaintainStartTime *string `json:"MaintainStartTime,omitempty" xml:"MaintainStartTime,omitempty"`
	// The number of coordinator nodes.
	MasterNodeNum *int32 `json:"MasterNodeNum,omitempty" xml:"MasterNodeNum,omitempty"`
	// The maximum number of concurrent connections to the instance.
	//
	// >  This parameter is available only for instances in reserved storage mode.
	MaxConnections *int32 `json:"MaxConnections,omitempty" xml:"MaxConnections,omitempty"`
	// The memory capacity per node. The unit can be one of the valid values of the **MemoryUnit** parameter.
	//
	// >  This parameter is available only for instances in reserved storage mode.
	MemoryPerNode *int32 `json:"MemoryPerNode,omitempty" xml:"MemoryPerNode,omitempty"`
	// The memory capacity per compute node.
	//
	// >  For instances in reserved storage mode, the unit of this parameter is MB. For instances in elastic storage mode and Serverless mode, the unit of this parameter is GB.
	MemorySize *int64 `json:"MemorySize,omitempty" xml:"MemorySize,omitempty"`
	// The unit of the memory capacity.
	//
	// >  This parameter is available only for instances in reserved storage mode.
	MemoryUnit *string `json:"MemoryUnit,omitempty" xml:"MemoryUnit,omitempty"`
	// The minor version of the instance.
	MinorVersion *string `json:"MinorVersion,omitempty" xml:"MinorVersion,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// *   **Postpaid**: pay-as-you-go.
	// *   **Prepaid**: subscription.
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The port number that is used to connect to the instance.
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// An invalid parameter. It is no longer returned when you call this operation.
	ReadDelayTime *string `json:"ReadDelayTime,omitempty" xml:"ReadDelayTime,omitempty"`
	// The region ID of the instance.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the instance belongs.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The running duration of the instance.
	RunningTime *string `json:"RunningTime,omitempty" xml:"RunningTime,omitempty"`
	// The IP address whitelist of the instance.
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	// The performance level of ESSDs. Only **PL1** is supported.
	SegDiskPerformanceLevel *string `json:"SegDiskPerformanceLevel,omitempty" xml:"SegDiskPerformanceLevel,omitempty"`
	// The number of compute nodes.
	//
	// >  This parameter is available only for instances in elastic storage mode or manual Serverless mode.
	SegNodeNum *int32 `json:"SegNodeNum,omitempty" xml:"SegNodeNum,omitempty"`
	// The number of compute groups.
	//
	// >  This parameter is available only for instances in reserved storage mode.
	SegmentCounts *int32 `json:"SegmentCounts,omitempty" xml:"SegmentCounts,omitempty"`
	// The type of the Serverless mode. Valid values:
	//
	// *   **Manual**: manual scheduling.
	// *   **Auto**: automatic scheduling.
	//
	// >  This parameter is returned only for instances in Serverless mode.
	ServerlessMode *string `json:"ServerlessMode,omitempty" xml:"ServerlessMode,omitempty"`
	// The threshold of computing resources. Unit: AnalyticDB compute unit (ACU).
	//
	// >  This parameter is returned only for instances in automatic Serverless mode.
	ServerlessResource *int32 `json:"ServerlessResource,omitempty" xml:"ServerlessResource,omitempty"`
	// The time when the instance started to run.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The storage capacity per node. The unit can be one of the valid values of the **StorageUnit** parameter.
	//
	// >  This parameter is available only for instances in reserved storage mode.
	StoragePerNode *int32 `json:"StoragePerNode,omitempty" xml:"StoragePerNode,omitempty"`
	// The storage capacity. Unit: GB.
	StorageSize *int64 `json:"StorageSize,omitempty" xml:"StorageSize,omitempty"`
	// The storage type of the instance. Valid values:
	//
	// - **cloud_essd**: enhanced SSD (ESSD).
	// - **cloud_efficiency**: ultra disk.
	//
	// >  This parameter is available only for instances in elastic storage mode.
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// The unit of the storage capacity. Valid values:
	//
	// *   **GB SSD**
	// *   **TB SSD**
	// *   **GB HDD**
	//
	// >  This parameter is available only for instances in reserved storage mode or Serverless mode.
	StorageUnit *string `json:"StorageUnit,omitempty" xml:"StorageUnit,omitempty"`
	// Indicates whether the instance supports backup and restoration.
	//
	// *   **true**
	// *   **false**
	SupportRestore *bool `json:"SupportRestore,omitempty" xml:"SupportRestore,omitempty"`
	// The tags of the instance. Each tag is a key-value pair.
	Tags *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The vSwitch ID of the instance.
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// Indicates whether vector engine optimization is enabled. Valid values:
	//
	// *   **enabled**
	// *   **disabled**
	VectorConfigurationStatus *string `json:"VectorConfigurationStatus,omitempty" xml:"VectorConfigurationStatus,omitempty"`
	// The virtual private cloud (VPC) ID of the instance.
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The zone ID of the instance.
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetAvailabilityValue(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.AvailabilityValue = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetConnectionMode(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.ConnectionMode = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetConnectionString(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.ConnectionString = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetCoreVersion(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.CoreVersion = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetCpuCores(v int32) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.CpuCores = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetCpuCoresPerNode(v int32) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.CpuCoresPerNode = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetCreationTime(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.CreationTime = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetDBInstanceCategory(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.DBInstanceCategory = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetDBInstanceClass(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.DBInstanceClass = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetDBInstanceClassType(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.DBInstanceClassType = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetDBInstanceCpuCores(v int32) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.DBInstanceCpuCores = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetDBInstanceDescription(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.DBInstanceDescription = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetDBInstanceDiskMBPS(v int64) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.DBInstanceDiskMBPS = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetDBInstanceGroupCount(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.DBInstanceGroupCount = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetDBInstanceId(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetDBInstanceMemory(v int64) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.DBInstanceMemory = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetDBInstanceMode(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.DBInstanceMode = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetDBInstanceNetType(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.DBInstanceNetType = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetDBInstanceStatus(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.DBInstanceStatus = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetDBInstanceStorage(v int64) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.DBInstanceStorage = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetEncryptionKey(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.EncryptionKey = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetEncryptionType(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.EncryptionType = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetEngine(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.Engine = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetEngineVersion(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.EngineVersion = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetExpireTime(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetHostType(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.HostType = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetIdleTime(v int32) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.IdleTime = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetInstanceNetworkType(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.InstanceNetworkType = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetLockMode(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.LockMode = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetLockReason(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.LockReason = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetMaintainEndTime(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.MaintainEndTime = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetMaintainStartTime(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.MaintainStartTime = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetMasterNodeNum(v int32) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.MasterNodeNum = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetMaxConnections(v int32) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.MaxConnections = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetMemoryPerNode(v int32) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.MemoryPerNode = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetMemorySize(v int64) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.MemorySize = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetMemoryUnit(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.MemoryUnit = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetMinorVersion(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.MinorVersion = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetPayType(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.PayType = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetPort(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.Port = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetReadDelayTime(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.ReadDelayTime = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetRegionId(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.RegionId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetResourceGroupId(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetRunningTime(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.RunningTime = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetSecurityIPList(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.SecurityIPList = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetSegDiskPerformanceLevel(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.SegDiskPerformanceLevel = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetSegNodeNum(v int32) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.SegNodeNum = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetSegmentCounts(v int32) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.SegmentCounts = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetServerlessMode(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.ServerlessMode = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetServerlessResource(v int32) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.ServerlessResource = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetStartTime(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.StartTime = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetStoragePerNode(v int32) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.StoragePerNode = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetStorageSize(v int64) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.StorageSize = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetStorageType(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.StorageType = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetStorageUnit(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.StorageUnit = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetSupportRestore(v bool) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.SupportRestore = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetTags(v *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeTags) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.Tags = v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetVSwitchId(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetVectorConfigurationStatus(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.VectorConfigurationStatus = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetVpcId(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.VpcId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetZoneId(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.ZoneId = &v
	return s
}

type DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeTags struct {
	Tag []*DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeTags) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeTags) SetTag(v []*DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeTagsTag) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeTags {
	s.Tag = v
	return s
}

type DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeTagsTag struct {
	// The tag key.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeTagsTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeTagsTag) SetKey(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeTagsTag) SetValue(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeTagsTag {
	s.Value = &v
	return s
}

type DescribeDBInstanceAttributeResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDBInstanceAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDBInstanceAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponse) SetHeaders(v map[string]*string) *DescribeDBInstanceAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInstanceAttributeResponse) SetStatusCode(v int32) *DescribeDBInstanceAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponse) SetBody(v *DescribeDBInstanceAttributeResponseBody) *DescribeDBInstanceAttributeResponse {
	s.Body = v
	return s
}

type DescribeDBInstanceDataBloatRequest struct {
	// The ID of the instance.
	//
	// >  You can call the [DescribeDBInstances](~~86911~~) operation to query the details of all AnalyticDB for PostgreSQL instances in a specific region, including instance IDs.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The number of the page to return. The value must be an integer that is greater than 0. Default value: **1**.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values:
	//
	// *   **30**
	// *   **50**
	// *   **100**
	//
	// Default value: **30**.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeDBInstanceDataBloatRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceDataBloatRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceDataBloatRequest) SetDBInstanceId(v string) *DescribeDBInstanceDataBloatRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceDataBloatRequest) SetPageNumber(v int32) *DescribeDBInstanceDataBloatRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBInstanceDataBloatRequest) SetPageSize(v int32) *DescribeDBInstanceDataBloatRequest {
	s.PageSize = &v
	return s
}

type DescribeDBInstanceDataBloatResponseBody struct {
	// Details of data bloat.
	Items []*DescribeDBInstanceDataBloatResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The page number of the returned page.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDBInstanceDataBloatResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceDataBloatResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceDataBloatResponseBody) SetItems(v []*DescribeDBInstanceDataBloatResponseBodyItems) *DescribeDBInstanceDataBloatResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDBInstanceDataBloatResponseBody) SetPageNumber(v int32) *DescribeDBInstanceDataBloatResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBInstanceDataBloatResponseBody) SetRequestId(v string) *DescribeDBInstanceDataBloatResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstanceDataBloatResponseBody) SetTotalCount(v int32) *DescribeDBInstanceDataBloatResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeDBInstanceDataBloatResponseBodyItems struct {
	// The coefficient of data bloat. It is calculated by using the following formula:
	//
	// Bloat coefficient = Number of dead rows/Number of active rows.
	BloatCeoff *string `json:"BloatCeoff,omitempty" xml:"BloatCeoff,omitempty"`
	// The bloat size of the table. It indicates the amount of space that can be released.
	BloatSize *string `json:"BloatSize,omitempty" xml:"BloatSize,omitempty"`
	// The name of the database.
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// The expected size of the table.
	//
	// It indicates the size of the table that has no data bloat.
	ExpectTableSize *string `json:"ExpectTableSize,omitempty" xml:"ExpectTableSize,omitempty"`
	// The actual size of the table.
	RealTableSize *string `json:"RealTableSize,omitempty" xml:"RealTableSize,omitempty"`
	// The name of the schema.
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The sequence number.
	Sequence *int32 `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// The storage type of the table. Valid values:
	//
	// *   **Heap**: heap table
	// *   **AO**: append-optimized (AO) table
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// This parameter is not returned.
	SuggestedAction *string `json:"SuggestedAction,omitempty" xml:"SuggestedAction,omitempty"`
	// The name of the table.
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The time when the table was last deleted, inserted, or updated.
	TimeLastUpdated *string `json:"TimeLastUpdated,omitempty" xml:"TimeLastUpdated,omitempty"`
	// The time when the table was last vacuumed. The time is displayed in UTC.
	TimeLastVacuumed *string `json:"TimeLastVacuumed,omitempty" xml:"TimeLastVacuumed,omitempty"`
}

func (s DescribeDBInstanceDataBloatResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceDataBloatResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceDataBloatResponseBodyItems) SetBloatCeoff(v string) *DescribeDBInstanceDataBloatResponseBodyItems {
	s.BloatCeoff = &v
	return s
}

func (s *DescribeDBInstanceDataBloatResponseBodyItems) SetBloatSize(v string) *DescribeDBInstanceDataBloatResponseBodyItems {
	s.BloatSize = &v
	return s
}

func (s *DescribeDBInstanceDataBloatResponseBodyItems) SetDatabaseName(v string) *DescribeDBInstanceDataBloatResponseBodyItems {
	s.DatabaseName = &v
	return s
}

func (s *DescribeDBInstanceDataBloatResponseBodyItems) SetExpectTableSize(v string) *DescribeDBInstanceDataBloatResponseBodyItems {
	s.ExpectTableSize = &v
	return s
}

func (s *DescribeDBInstanceDataBloatResponseBodyItems) SetRealTableSize(v string) *DescribeDBInstanceDataBloatResponseBodyItems {
	s.RealTableSize = &v
	return s
}

func (s *DescribeDBInstanceDataBloatResponseBodyItems) SetSchemaName(v string) *DescribeDBInstanceDataBloatResponseBodyItems {
	s.SchemaName = &v
	return s
}

func (s *DescribeDBInstanceDataBloatResponseBodyItems) SetSequence(v int32) *DescribeDBInstanceDataBloatResponseBodyItems {
	s.Sequence = &v
	return s
}

func (s *DescribeDBInstanceDataBloatResponseBodyItems) SetStorageType(v string) *DescribeDBInstanceDataBloatResponseBodyItems {
	s.StorageType = &v
	return s
}

func (s *DescribeDBInstanceDataBloatResponseBodyItems) SetSuggestedAction(v string) *DescribeDBInstanceDataBloatResponseBodyItems {
	s.SuggestedAction = &v
	return s
}

func (s *DescribeDBInstanceDataBloatResponseBodyItems) SetTableName(v string) *DescribeDBInstanceDataBloatResponseBodyItems {
	s.TableName = &v
	return s
}

func (s *DescribeDBInstanceDataBloatResponseBodyItems) SetTimeLastUpdated(v string) *DescribeDBInstanceDataBloatResponseBodyItems {
	s.TimeLastUpdated = &v
	return s
}

func (s *DescribeDBInstanceDataBloatResponseBodyItems) SetTimeLastVacuumed(v string) *DescribeDBInstanceDataBloatResponseBodyItems {
	s.TimeLastVacuumed = &v
	return s
}

type DescribeDBInstanceDataBloatResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDBInstanceDataBloatResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDBInstanceDataBloatResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceDataBloatResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceDataBloatResponse) SetHeaders(v map[string]*string) *DescribeDBInstanceDataBloatResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInstanceDataBloatResponse) SetStatusCode(v int32) *DescribeDBInstanceDataBloatResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBInstanceDataBloatResponse) SetBody(v *DescribeDBInstanceDataBloatResponseBody) *DescribeDBInstanceDataBloatResponse {
	s.Body = v
	return s
}

type DescribeDBInstanceDataSkewRequest struct {
	// The ID of the instance.
	//
	// >  You can call the [DescribeDBInstances](~~86911~~) operation to query the details of all AnalyticDB for PostgreSQL instances in a specific region, including instance IDs.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The number of the page to return. The value must be an integer that is greater than 0. Default value: **1**.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values:
	//
	// *   **20**
	// *   **50**
	// *   **100**
	//
	// Default value: **20**.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeDBInstanceDataSkewRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceDataSkewRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceDataSkewRequest) SetDBInstanceId(v string) *DescribeDBInstanceDataSkewRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceDataSkewRequest) SetPageNumber(v int32) *DescribeDBInstanceDataSkewRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBInstanceDataSkewRequest) SetPageSize(v int32) *DescribeDBInstanceDataSkewRequest {
	s.PageSize = &v
	return s
}

type DescribeDBInstanceDataSkewResponseBody struct {
	// Details about data skew.
	Items []*DescribeDBInstanceDataSkewResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The page number of the returned page.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDBInstanceDataSkewResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceDataSkewResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceDataSkewResponseBody) SetItems(v []*DescribeDBInstanceDataSkewResponseBodyItems) *DescribeDBInstanceDataSkewResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDBInstanceDataSkewResponseBody) SetPageNumber(v int32) *DescribeDBInstanceDataSkewResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBInstanceDataSkewResponseBody) SetRequestId(v string) *DescribeDBInstanceDataSkewResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstanceDataSkewResponseBody) SetTotalCount(v int32) *DescribeDBInstanceDataSkewResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeDBInstanceDataSkewResponseBodyItems struct {
	// The name of the database.
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// The distribution key of the table.
	DistributeKey *string `json:"DistributeKey,omitempty" xml:"DistributeKey,omitempty"`
	// The owner of the table.
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The name of the schema.
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The sequence number of the data skew case. All data skew cases are sorted by severity in descending order.
	Sequence *int32 `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// The name of the table.
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The total number of rows in the table.
	TableSize *string `json:"TableSize,omitempty" xml:"TableSize,omitempty"`
	// The skew ratio of the table. Valid values: 0 to 100. Unit: %. A greater value indicates that the table is more severely skewed. A smaller value indicates less impact on query performance. A value of 0 indicates no data skew.
	TableSkew *string `json:"TableSkew,omitempty" xml:"TableSkew,omitempty"`
	// The time when the table was last deleted, inserted, or updated.
	TimeLastUpdated *string `json:"TimeLastUpdated,omitempty" xml:"TimeLastUpdated,omitempty"`
}

func (s DescribeDBInstanceDataSkewResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceDataSkewResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceDataSkewResponseBodyItems) SetDatabaseName(v string) *DescribeDBInstanceDataSkewResponseBodyItems {
	s.DatabaseName = &v
	return s
}

func (s *DescribeDBInstanceDataSkewResponseBodyItems) SetDistributeKey(v string) *DescribeDBInstanceDataSkewResponseBodyItems {
	s.DistributeKey = &v
	return s
}

func (s *DescribeDBInstanceDataSkewResponseBodyItems) SetOwner(v string) *DescribeDBInstanceDataSkewResponseBodyItems {
	s.Owner = &v
	return s
}

func (s *DescribeDBInstanceDataSkewResponseBodyItems) SetSchemaName(v string) *DescribeDBInstanceDataSkewResponseBodyItems {
	s.SchemaName = &v
	return s
}

func (s *DescribeDBInstanceDataSkewResponseBodyItems) SetSequence(v int32) *DescribeDBInstanceDataSkewResponseBodyItems {
	s.Sequence = &v
	return s
}

func (s *DescribeDBInstanceDataSkewResponseBodyItems) SetTableName(v string) *DescribeDBInstanceDataSkewResponseBodyItems {
	s.TableName = &v
	return s
}

func (s *DescribeDBInstanceDataSkewResponseBodyItems) SetTableSize(v string) *DescribeDBInstanceDataSkewResponseBodyItems {
	s.TableSize = &v
	return s
}

func (s *DescribeDBInstanceDataSkewResponseBodyItems) SetTableSkew(v string) *DescribeDBInstanceDataSkewResponseBodyItems {
	s.TableSkew = &v
	return s
}

func (s *DescribeDBInstanceDataSkewResponseBodyItems) SetTimeLastUpdated(v string) *DescribeDBInstanceDataSkewResponseBodyItems {
	s.TimeLastUpdated = &v
	return s
}

type DescribeDBInstanceDataSkewResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDBInstanceDataSkewResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDBInstanceDataSkewResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceDataSkewResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceDataSkewResponse) SetHeaders(v map[string]*string) *DescribeDBInstanceDataSkewResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInstanceDataSkewResponse) SetStatusCode(v int32) *DescribeDBInstanceDataSkewResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBInstanceDataSkewResponse) SetBody(v *DescribeDBInstanceDataSkewResponseBody) *DescribeDBInstanceDataSkewResponse {
	s.Body = v
	return s
}

type DescribeDBInstanceDiagnosisSummaryRequest struct {
	// The ID of the instance.
	//
	// >  You can call the [DescribeDBInstances](~~86911~~) operation to query the details of all AnalyticDB for PostgreSQL instances in a specific region, including instance IDs.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The number of the page to return. The value must be an integer that is greater than 0. Default value: **1**.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values:
	//
	// *   **20**
	// *   **50**
	// *   **100**
	//
	// Default value: **20**.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The role state of the node. It indicates whether a primary/secondary switchover occurs. Valid values:
	//
	// *   **normal**: The node role is normal. No primary/secondary switchover occurs.
	// *   **reverse**: The node role is reversed. A primary/secondary switchover occurs.
	RolePreferd *string `json:"RolePreferd,omitempty" xml:"RolePreferd,omitempty"`
	// The running state of the node. Valid values:
	//
	// *   **UP**: The node is running.
	// *   **DOWN**: The node is faulty.
	//
	// If this parameter is not specified, information of nodes in all running states is returned.
	StartStatus *string `json:"StartStatus,omitempty" xml:"StartStatus,omitempty"`
	// The data synchronization state of the node. Valid values:
	//
	// *   **synced**: The node data is synchronized.
	// *   **notSyncing**: The node data is not synchronized.
	//
	// If this parameter is not specified, information of nodes in all synchronization states is returned.
	SyncMode *string `json:"SyncMode,omitempty" xml:"SyncMode,omitempty"`
}

func (s DescribeDBInstanceDiagnosisSummaryRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceDiagnosisSummaryRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceDiagnosisSummaryRequest) SetDBInstanceId(v string) *DescribeDBInstanceDiagnosisSummaryRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryRequest) SetPageNumber(v int32) *DescribeDBInstanceDiagnosisSummaryRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryRequest) SetPageSize(v int32) *DescribeDBInstanceDiagnosisSummaryRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryRequest) SetRolePreferd(v string) *DescribeDBInstanceDiagnosisSummaryRequest {
	s.RolePreferd = &v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryRequest) SetStartStatus(v string) *DescribeDBInstanceDiagnosisSummaryRequest {
	s.StartStatus = &v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryRequest) SetSyncMode(v string) *DescribeDBInstanceDiagnosisSummaryRequest {
	s.SyncMode = &v
	return s
}

type DescribeDBInstanceDiagnosisSummaryResponseBody struct {
	// Details of instance nodes.
	Items []*DescribeDBInstanceDiagnosisSummaryResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// State statistics of the coordinator node.
	MasterStatusInfo *DescribeDBInstanceDiagnosisSummaryResponseBodyMasterStatusInfo `json:"MasterStatusInfo,omitempty" xml:"MasterStatusInfo,omitempty" type:"Struct"`
	// The page number of the returned page.
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// State statistics of compute nodes.
	SegmentStatusInfo *DescribeDBInstanceDiagnosisSummaryResponseBodySegmentStatusInfo `json:"SegmentStatusInfo,omitempty" xml:"SegmentStatusInfo,omitempty" type:"Struct"`
	// The total number of entries returned.
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDBInstanceDiagnosisSummaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceDiagnosisSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBody) SetItems(v []*DescribeDBInstanceDiagnosisSummaryResponseBodyItems) *DescribeDBInstanceDiagnosisSummaryResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBody) SetMasterStatusInfo(v *DescribeDBInstanceDiagnosisSummaryResponseBodyMasterStatusInfo) *DescribeDBInstanceDiagnosisSummaryResponseBody {
	s.MasterStatusInfo = v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBody) SetPageNumber(v string) *DescribeDBInstanceDiagnosisSummaryResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBody) SetRequestId(v string) *DescribeDBInstanceDiagnosisSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBody) SetSegmentStatusInfo(v *DescribeDBInstanceDiagnosisSummaryResponseBodySegmentStatusInfo) *DescribeDBInstanceDiagnosisSummaryResponseBody {
	s.SegmentStatusInfo = v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBody) SetTotalCount(v string) *DescribeDBInstanceDiagnosisSummaryResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeDBInstanceDiagnosisSummaryResponseBodyItems struct {
	// The name of the node.
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// The IP address of the node.
	NodeAddress *string `json:"NodeAddress,omitempty" xml:"NodeAddress,omitempty"`
	// The ID of the node group.
	NodeCID *string `json:"NodeCID,omitempty" xml:"NodeCID,omitempty"`
	// The ID of the node.
	NodeID *string `json:"NodeID,omitempty" xml:"NodeID,omitempty"`
	// The name of the host where the node resides.
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The port number of the node.
	NodePort *string `json:"NodePort,omitempty" xml:"NodePort,omitempty"`
	// The initial role of the node. Valid values:
	//
	// *   **primary**: primary node
	// *   **mirror**: secondary node
	//
	// If the value of this parameter is the same as that of **NodeRole**, no primary/secondary switchover occurs. If the value of this parameter is not the same as that of **NodeRole**, a primary/secondary switchover occurs.
	NodePreferredRole *string `json:"NodePreferredRole,omitempty" xml:"NodePreferredRole,omitempty"`
	// The data synchronization state of the node. Valid values:
	//
	// *   **Synced**: The node data is synchronized.
	// *   **Not Syncing**: The node data is not synchronized.
	// *   **No sync required**: Data synchronization is not required. This value may be returned only for the coordinator node.
	NodeReplicationMode *string `json:"NodeReplicationMode,omitempty" xml:"NodeReplicationMode,omitempty"`
	// The current role of the node. Valid values:
	//
	// *   **primary**: primary node
	// *   **mirror**: secondary node
	NodeRole *string `json:"NodeRole,omitempty" xml:"NodeRole,omitempty"`
	// The running state of the node. Valid values:
	//
	// *   **UP**: The node is running.
	// *   **DOWN**: The node is faulty.
	NodeStatus *string `json:"NodeStatus,omitempty" xml:"NodeStatus,omitempty"`
	// The type of the node. Valid values:
	//
	// *   **master**: primary coordinator node
	// *   **slave**: standby coordinator node
	// *   **segment**: compute node
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
}

func (s DescribeDBInstanceDiagnosisSummaryResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceDiagnosisSummaryResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodyItems) SetHostname(v string) *DescribeDBInstanceDiagnosisSummaryResponseBodyItems {
	s.Hostname = &v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodyItems) SetNodeAddress(v string) *DescribeDBInstanceDiagnosisSummaryResponseBodyItems {
	s.NodeAddress = &v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodyItems) SetNodeCID(v string) *DescribeDBInstanceDiagnosisSummaryResponseBodyItems {
	s.NodeCID = &v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodyItems) SetNodeID(v string) *DescribeDBInstanceDiagnosisSummaryResponseBodyItems {
	s.NodeID = &v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodyItems) SetNodeName(v string) *DescribeDBInstanceDiagnosisSummaryResponseBodyItems {
	s.NodeName = &v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodyItems) SetNodePort(v string) *DescribeDBInstanceDiagnosisSummaryResponseBodyItems {
	s.NodePort = &v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodyItems) SetNodePreferredRole(v string) *DescribeDBInstanceDiagnosisSummaryResponseBodyItems {
	s.NodePreferredRole = &v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodyItems) SetNodeReplicationMode(v string) *DescribeDBInstanceDiagnosisSummaryResponseBodyItems {
	s.NodeReplicationMode = &v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodyItems) SetNodeRole(v string) *DescribeDBInstanceDiagnosisSummaryResponseBodyItems {
	s.NodeRole = &v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodyItems) SetNodeStatus(v string) *DescribeDBInstanceDiagnosisSummaryResponseBodyItems {
	s.NodeStatus = &v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodyItems) SetNodeType(v string) *DescribeDBInstanceDiagnosisSummaryResponseBodyItems {
	s.NodeType = &v
	return s
}

type DescribeDBInstanceDiagnosisSummaryResponseBodyMasterStatusInfo struct {
	// The number of abnormal nodes.
	ExceptionNodeNum *int32 `json:"ExceptionNodeNum,omitempty" xml:"ExceptionNodeNum,omitempty"`
	// The number of normal nodes.
	NormalNodeNum *int32 `json:"NormalNodeNum,omitempty" xml:"NormalNodeNum,omitempty"`
	// The number of nodes whose roles are reversed.
	NotPreferredNodeNum *int32 `json:"NotPreferredNodeNum,omitempty" xml:"NotPreferredNodeNum,omitempty"`
	// The number of unsynchronized nodes.
	NotSyncingNodeNum *int32 `json:"NotSyncingNodeNum,omitempty" xml:"NotSyncingNodeNum,omitempty"`
	// The number of nodes whose roles are normal.
	PreferredNodeNum *int32 `json:"PreferredNodeNum,omitempty" xml:"PreferredNodeNum,omitempty"`
	// The number of synchronized nodes.
	SyncedNodeNum *int32 `json:"SyncedNodeNum,omitempty" xml:"SyncedNodeNum,omitempty"`
}

func (s DescribeDBInstanceDiagnosisSummaryResponseBodyMasterStatusInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceDiagnosisSummaryResponseBodyMasterStatusInfo) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodyMasterStatusInfo) SetExceptionNodeNum(v int32) *DescribeDBInstanceDiagnosisSummaryResponseBodyMasterStatusInfo {
	s.ExceptionNodeNum = &v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodyMasterStatusInfo) SetNormalNodeNum(v int32) *DescribeDBInstanceDiagnosisSummaryResponseBodyMasterStatusInfo {
	s.NormalNodeNum = &v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodyMasterStatusInfo) SetNotPreferredNodeNum(v int32) *DescribeDBInstanceDiagnosisSummaryResponseBodyMasterStatusInfo {
	s.NotPreferredNodeNum = &v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodyMasterStatusInfo) SetNotSyncingNodeNum(v int32) *DescribeDBInstanceDiagnosisSummaryResponseBodyMasterStatusInfo {
	s.NotSyncingNodeNum = &v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodyMasterStatusInfo) SetPreferredNodeNum(v int32) *DescribeDBInstanceDiagnosisSummaryResponseBodyMasterStatusInfo {
	s.PreferredNodeNum = &v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodyMasterStatusInfo) SetSyncedNodeNum(v int32) *DescribeDBInstanceDiagnosisSummaryResponseBodyMasterStatusInfo {
	s.SyncedNodeNum = &v
	return s
}

type DescribeDBInstanceDiagnosisSummaryResponseBodySegmentStatusInfo struct {
	// The number of abnormal nodes.
	ExceptionNodeNum *int32 `json:"ExceptionNodeNum,omitempty" xml:"ExceptionNodeNum,omitempty"`
	// The number of normal nodes.
	NormalNodeNum *int32 `json:"NormalNodeNum,omitempty" xml:"NormalNodeNum,omitempty"`
	// The number of nodes whose roles are reversed.
	NotPreferredNodeNum *int32 `json:"NotPreferredNodeNum,omitempty" xml:"NotPreferredNodeNum,omitempty"`
	// The number of unsynchronized nodes.
	NotSyncingNodeNum *int32 `json:"NotSyncingNodeNum,omitempty" xml:"NotSyncingNodeNum,omitempty"`
	// The number of nodes whose roles are normal.
	PreferredNodeNum *int32 `json:"PreferredNodeNum,omitempty" xml:"PreferredNodeNum,omitempty"`
	// The number of synchronized nodes.
	SyncedNodeNum *int32 `json:"SyncedNodeNum,omitempty" xml:"SyncedNodeNum,omitempty"`
}

func (s DescribeDBInstanceDiagnosisSummaryResponseBodySegmentStatusInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceDiagnosisSummaryResponseBodySegmentStatusInfo) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodySegmentStatusInfo) SetExceptionNodeNum(v int32) *DescribeDBInstanceDiagnosisSummaryResponseBodySegmentStatusInfo {
	s.ExceptionNodeNum = &v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodySegmentStatusInfo) SetNormalNodeNum(v int32) *DescribeDBInstanceDiagnosisSummaryResponseBodySegmentStatusInfo {
	s.NormalNodeNum = &v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodySegmentStatusInfo) SetNotPreferredNodeNum(v int32) *DescribeDBInstanceDiagnosisSummaryResponseBodySegmentStatusInfo {
	s.NotPreferredNodeNum = &v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodySegmentStatusInfo) SetNotSyncingNodeNum(v int32) *DescribeDBInstanceDiagnosisSummaryResponseBodySegmentStatusInfo {
	s.NotSyncingNodeNum = &v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodySegmentStatusInfo) SetPreferredNodeNum(v int32) *DescribeDBInstanceDiagnosisSummaryResponseBodySegmentStatusInfo {
	s.PreferredNodeNum = &v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodySegmentStatusInfo) SetSyncedNodeNum(v int32) *DescribeDBInstanceDiagnosisSummaryResponseBodySegmentStatusInfo {
	s.SyncedNodeNum = &v
	return s
}

type DescribeDBInstanceDiagnosisSummaryResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDBInstanceDiagnosisSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDBInstanceDiagnosisSummaryResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceDiagnosisSummaryResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceDiagnosisSummaryResponse) SetHeaders(v map[string]*string) *DescribeDBInstanceDiagnosisSummaryResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryResponse) SetStatusCode(v int32) *DescribeDBInstanceDiagnosisSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryResponse) SetBody(v *DescribeDBInstanceDiagnosisSummaryResponseBody) *DescribeDBInstanceDiagnosisSummaryResponse {
	s.Body = v
	return s
}

type DescribeDBInstanceErrorLogRequest struct {
	// The ID of the instance.
	//
	// >  You can call the [DescribeDBInstances](~~86911~~) operation to query the details of all AnalyticDB for PostgreSQL instances in a specific region, including instance IDs.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The name of the database.
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	// The end of the time range to query. Specify the time in the *yyyy-MM-dd*T*HH:mm*Z format. The time must be in UTC. The end time must be later than the start time.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is not supported in Alibaba Cloud public cloud.
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// One or more keywords that can be used to query error logs.
	Keywords *string `json:"Keywords,omitempty" xml:"Keywords,omitempty"`
	// The level of the logs to query. Valid values:
	//
	// *   **ALL**: queries all error logs.
	// *   **PANIC**: queries only abnormal-level logs.
	// *   **FATAL**: queries only critical-level logs.
	// *   **ERROR**: queries only error-level logs.
	LogLevel *string `json:"LogLevel,omitempty" xml:"LogLevel,omitempty"`
	// The number of the page to return. The value must be an integer that is greater than 0. Default value: **1**.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values:
	//
	// *   **20**
	// *   **50**
	// *   **100**
	//
	// Default value: **20**.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The beginning of the time range to query. Specify the time in the *yyyy-MM-dd*T*HH:mm*Z format. The time must be in UTC.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The name of the database account.
	User *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DescribeDBInstanceErrorLogRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceErrorLogRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceErrorLogRequest) SetDBInstanceId(v string) *DescribeDBInstanceErrorLogRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceErrorLogRequest) SetDatabase(v string) *DescribeDBInstanceErrorLogRequest {
	s.Database = &v
	return s
}

func (s *DescribeDBInstanceErrorLogRequest) SetEndTime(v string) *DescribeDBInstanceErrorLogRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDBInstanceErrorLogRequest) SetHost(v string) *DescribeDBInstanceErrorLogRequest {
	s.Host = &v
	return s
}

func (s *DescribeDBInstanceErrorLogRequest) SetKeywords(v string) *DescribeDBInstanceErrorLogRequest {
	s.Keywords = &v
	return s
}

func (s *DescribeDBInstanceErrorLogRequest) SetLogLevel(v string) *DescribeDBInstanceErrorLogRequest {
	s.LogLevel = &v
	return s
}

func (s *DescribeDBInstanceErrorLogRequest) SetPageNumber(v int32) *DescribeDBInstanceErrorLogRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBInstanceErrorLogRequest) SetPageSize(v int32) *DescribeDBInstanceErrorLogRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDBInstanceErrorLogRequest) SetStartTime(v string) *DescribeDBInstanceErrorLogRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDBInstanceErrorLogRequest) SetUser(v string) *DescribeDBInstanceErrorLogRequest {
	s.User = &v
	return s
}

type DescribeDBInstanceErrorLogResponseBody struct {
	// Details of the error logs.
	Items []*DescribeDBInstanceErrorLogResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The page number of the returned page.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDBInstanceErrorLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceErrorLogResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceErrorLogResponseBody) SetItems(v []*DescribeDBInstanceErrorLogResponseBodyItems) *DescribeDBInstanceErrorLogResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDBInstanceErrorLogResponseBody) SetPageNumber(v int32) *DescribeDBInstanceErrorLogResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBInstanceErrorLogResponseBody) SetRequestId(v string) *DescribeDBInstanceErrorLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstanceErrorLogResponseBody) SetTotalCount(v int32) *DescribeDBInstanceErrorLogResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeDBInstanceErrorLogResponseBodyItems struct {
	// The name of the database.
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	// This parameter is not supported.
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The content of the error log.
	LogContext *string `json:"LogContext,omitempty" xml:"LogContext,omitempty"`
	// The level of the queried log.
	LogLevel *string `json:"LogLevel,omitempty" xml:"LogLevel,omitempty"`
	// The time when the log was generated. The time is displayed in UTC.
	Time *int64 `json:"Time,omitempty" xml:"Time,omitempty"`
	// The name of the database account.
	User *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DescribeDBInstanceErrorLogResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceErrorLogResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceErrorLogResponseBodyItems) SetDatabase(v string) *DescribeDBInstanceErrorLogResponseBodyItems {
	s.Database = &v
	return s
}

func (s *DescribeDBInstanceErrorLogResponseBodyItems) SetHost(v string) *DescribeDBInstanceErrorLogResponseBodyItems {
	s.Host = &v
	return s
}

func (s *DescribeDBInstanceErrorLogResponseBodyItems) SetLogContext(v string) *DescribeDBInstanceErrorLogResponseBodyItems {
	s.LogContext = &v
	return s
}

func (s *DescribeDBInstanceErrorLogResponseBodyItems) SetLogLevel(v string) *DescribeDBInstanceErrorLogResponseBodyItems {
	s.LogLevel = &v
	return s
}

func (s *DescribeDBInstanceErrorLogResponseBodyItems) SetTime(v int64) *DescribeDBInstanceErrorLogResponseBodyItems {
	s.Time = &v
	return s
}

func (s *DescribeDBInstanceErrorLogResponseBodyItems) SetUser(v string) *DescribeDBInstanceErrorLogResponseBodyItems {
	s.User = &v
	return s
}

type DescribeDBInstanceErrorLogResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDBInstanceErrorLogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDBInstanceErrorLogResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceErrorLogResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceErrorLogResponse) SetHeaders(v map[string]*string) *DescribeDBInstanceErrorLogResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInstanceErrorLogResponse) SetStatusCode(v int32) *DescribeDBInstanceErrorLogResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBInstanceErrorLogResponse) SetBody(v *DescribeDBInstanceErrorLogResponseBody) *DescribeDBInstanceErrorLogResponse {
	s.Body = v
	return s
}

type DescribeDBInstanceIPArrayListRequest struct {
	// The ID of the instance.
	//
	// >  You can call the [DescribeDBInstances](~~86911~~) operation to query the details of all AnalyticDB for PostgreSQL instances in a specific region, including instance IDs.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The ID of the resource group to which the instance belongs. For more information about how to obtain the ID of a resource group, see [View basic information of a resource group](~~151181~~).
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeDBInstanceIPArrayListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceIPArrayListRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceIPArrayListRequest) SetDBInstanceId(v string) *DescribeDBInstanceIPArrayListRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceIPArrayListRequest) SetResourceGroupId(v string) *DescribeDBInstanceIPArrayListRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeDBInstanceIPArrayListResponseBody struct {
	// Details of the IP address whitelists.
	Items *DescribeDBInstanceIPArrayListResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBInstanceIPArrayListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceIPArrayListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceIPArrayListResponseBody) SetItems(v *DescribeDBInstanceIPArrayListResponseBodyItems) *DescribeDBInstanceIPArrayListResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDBInstanceIPArrayListResponseBody) SetRequestId(v string) *DescribeDBInstanceIPArrayListResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDBInstanceIPArrayListResponseBodyItems struct {
	DBInstanceIPArray []*DescribeDBInstanceIPArrayListResponseBodyItemsDBInstanceIPArray `json:"DBInstanceIPArray,omitempty" xml:"DBInstanceIPArray,omitempty" type:"Repeated"`
}

func (s DescribeDBInstanceIPArrayListResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceIPArrayListResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceIPArrayListResponseBodyItems) SetDBInstanceIPArray(v []*DescribeDBInstanceIPArrayListResponseBodyItemsDBInstanceIPArray) *DescribeDBInstanceIPArrayListResponseBodyItems {
	s.DBInstanceIPArray = v
	return s
}

type DescribeDBInstanceIPArrayListResponseBodyItemsDBInstanceIPArray struct {
	// The attribute of the IP address whitelist. By default, this parameter is empty. A whitelist with the `hidden` attribute does not appear in the console.
	DBInstanceIPArrayAttribute *string `json:"DBInstanceIPArrayAttribute,omitempty" xml:"DBInstanceIPArrayAttribute,omitempty"`
	// The name of the IP address whitelist.
	DBInstanceIPArrayName *string `json:"DBInstanceIPArrayName,omitempty" xml:"DBInstanceIPArrayName,omitempty"`
	// The IP addresses listed in the whitelist. You can add up to 1,000 IP addresses to the whitelist. Separate multiple IP addresses with commas (,). The IP addresses must use one of the following formats:
	//
	// *   0.0.0.0/0
	// *   10.23.12.24. This is a standard IP address.
	// *   10.23.12.24/24. This is a CIDR block. The value `/24` indicates that the prefix of the CIDR block is 24-bit long. You can replace 24 with a value in the range of `1 to 32`.
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
}

func (s DescribeDBInstanceIPArrayListResponseBodyItemsDBInstanceIPArray) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceIPArrayListResponseBodyItemsDBInstanceIPArray) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceIPArrayListResponseBodyItemsDBInstanceIPArray) SetDBInstanceIPArrayAttribute(v string) *DescribeDBInstanceIPArrayListResponseBodyItemsDBInstanceIPArray {
	s.DBInstanceIPArrayAttribute = &v
	return s
}

func (s *DescribeDBInstanceIPArrayListResponseBodyItemsDBInstanceIPArray) SetDBInstanceIPArrayName(v string) *DescribeDBInstanceIPArrayListResponseBodyItemsDBInstanceIPArray {
	s.DBInstanceIPArrayName = &v
	return s
}

func (s *DescribeDBInstanceIPArrayListResponseBodyItemsDBInstanceIPArray) SetSecurityIPList(v string) *DescribeDBInstanceIPArrayListResponseBodyItemsDBInstanceIPArray {
	s.SecurityIPList = &v
	return s
}

type DescribeDBInstanceIPArrayListResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDBInstanceIPArrayListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDBInstanceIPArrayListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceIPArrayListResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceIPArrayListResponse) SetHeaders(v map[string]*string) *DescribeDBInstanceIPArrayListResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInstanceIPArrayListResponse) SetStatusCode(v int32) *DescribeDBInstanceIPArrayListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBInstanceIPArrayListResponse) SetBody(v *DescribeDBInstanceIPArrayListResponseBody) *DescribeDBInstanceIPArrayListResponse {
	s.Body = v
	return s
}

type DescribeDBInstanceIndexUsageRequest struct {
	// The ID of the instance.
	//
	// >  You can call the [DescribeDBInstances](~~86911~~) operation to query the details of all AnalyticDB for PostgreSQL instances in a specific region, including instance IDs.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The number of the page to return. The value must be an integer that is greater than 0. Default value: **1**.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values:
	//
	// *   **20**
	// *   **50**
	// *   **100**
	//
	// Default value: **20**.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeDBInstanceIndexUsageRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceIndexUsageRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceIndexUsageRequest) SetDBInstanceId(v string) *DescribeDBInstanceIndexUsageRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceIndexUsageRequest) SetPageNumber(v int32) *DescribeDBInstanceIndexUsageRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBInstanceIndexUsageRequest) SetPageSize(v int32) *DescribeDBInstanceIndexUsageRequest {
	s.PageSize = &v
	return s
}

type DescribeDBInstanceIndexUsageResponseBody struct {
	// Details of index usage.
	Items []*DescribeDBInstanceIndexUsageResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The page number of the returned page.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDBInstanceIndexUsageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceIndexUsageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceIndexUsageResponseBody) SetItems(v []*DescribeDBInstanceIndexUsageResponseBodyItems) *DescribeDBInstanceIndexUsageResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDBInstanceIndexUsageResponseBody) SetPageNumber(v int32) *DescribeDBInstanceIndexUsageResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBInstanceIndexUsageResponseBody) SetRequestId(v string) *DescribeDBInstanceIndexUsageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstanceIndexUsageResponseBody) SetTotalCount(v int32) *DescribeDBInstanceIndexUsageResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeDBInstanceIndexUsageResponseBodyItems struct {
	// The name of the database.
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// The definition of the index.
	IndexDef *string `json:"IndexDef,omitempty" xml:"IndexDef,omitempty"`
	// The name of the index.
	IndexName *string `json:"IndexName,omitempty" xml:"IndexName,omitempty"`
	// The number of index scans.
	IndexScanTimes *int32 `json:"IndexScanTimes,omitempty" xml:"IndexScanTimes,omitempty"`
	// The size of the index. Unit: bytes.
	IndexSize *string `json:"IndexSize,omitempty" xml:"IndexSize,omitempty"`
	// Indicates whether the table is a partitioned table. Valid values:
	//
	// *   **true**: The table is a partitioned table.
	// *   **false**: The table is not a partitioned table.
	IsPartitionTable *bool `json:"IsPartitionTable,omitempty" xml:"IsPartitionTable,omitempty"`
	// The name of the parent table.
	ParentTableName *string `json:"ParentTableName,omitempty" xml:"ParentTableName,omitempty"`
	// The name of the schema.
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The name of the table.
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The time when the table was last deleted, inserted, or updated.
	TimeLastUpdated *string `json:"TimeLastUpdated,omitempty" xml:"TimeLastUpdated,omitempty"`
}

func (s DescribeDBInstanceIndexUsageResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceIndexUsageResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceIndexUsageResponseBodyItems) SetDatabaseName(v string) *DescribeDBInstanceIndexUsageResponseBodyItems {
	s.DatabaseName = &v
	return s
}

func (s *DescribeDBInstanceIndexUsageResponseBodyItems) SetIndexDef(v string) *DescribeDBInstanceIndexUsageResponseBodyItems {
	s.IndexDef = &v
	return s
}

func (s *DescribeDBInstanceIndexUsageResponseBodyItems) SetIndexName(v string) *DescribeDBInstanceIndexUsageResponseBodyItems {
	s.IndexName = &v
	return s
}

func (s *DescribeDBInstanceIndexUsageResponseBodyItems) SetIndexScanTimes(v int32) *DescribeDBInstanceIndexUsageResponseBodyItems {
	s.IndexScanTimes = &v
	return s
}

func (s *DescribeDBInstanceIndexUsageResponseBodyItems) SetIndexSize(v string) *DescribeDBInstanceIndexUsageResponseBodyItems {
	s.IndexSize = &v
	return s
}

func (s *DescribeDBInstanceIndexUsageResponseBodyItems) SetIsPartitionTable(v bool) *DescribeDBInstanceIndexUsageResponseBodyItems {
	s.IsPartitionTable = &v
	return s
}

func (s *DescribeDBInstanceIndexUsageResponseBodyItems) SetParentTableName(v string) *DescribeDBInstanceIndexUsageResponseBodyItems {
	s.ParentTableName = &v
	return s
}

func (s *DescribeDBInstanceIndexUsageResponseBodyItems) SetSchemaName(v string) *DescribeDBInstanceIndexUsageResponseBodyItems {
	s.SchemaName = &v
	return s
}

func (s *DescribeDBInstanceIndexUsageResponseBodyItems) SetTableName(v string) *DescribeDBInstanceIndexUsageResponseBodyItems {
	s.TableName = &v
	return s
}

func (s *DescribeDBInstanceIndexUsageResponseBodyItems) SetTimeLastUpdated(v string) *DescribeDBInstanceIndexUsageResponseBodyItems {
	s.TimeLastUpdated = &v
	return s
}

type DescribeDBInstanceIndexUsageResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDBInstanceIndexUsageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDBInstanceIndexUsageResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceIndexUsageResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceIndexUsageResponse) SetHeaders(v map[string]*string) *DescribeDBInstanceIndexUsageResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInstanceIndexUsageResponse) SetStatusCode(v int32) *DescribeDBInstanceIndexUsageResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBInstanceIndexUsageResponse) SetBody(v *DescribeDBInstanceIndexUsageResponseBody) *DescribeDBInstanceIndexUsageResponse {
	s.Body = v
	return s
}

type DescribeDBInstanceNetInfoRequest struct {
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// The ID of the instance.
	//
	// >  You can call the [DescribeDBInstances](~~86911~~) operation to query details about all AnalyticDB for PostgreSQL instances in a specific region, including instance IDs.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s DescribeDBInstanceNetInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceNetInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceNetInfoRequest) SetConnectionString(v string) *DescribeDBInstanceNetInfoRequest {
	s.ConnectionString = &v
	return s
}

func (s *DescribeDBInstanceNetInfoRequest) SetDBInstanceId(v string) *DescribeDBInstanceNetInfoRequest {
	s.DBInstanceId = &v
	return s
}

type DescribeDBInstanceNetInfoResponseBody struct {
	// The connection information of the instance.
	DBInstanceNetInfos *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos `json:"DBInstanceNetInfos,omitempty" xml:"DBInstanceNetInfos,omitempty" type:"Struct"`
	// The network type of the instance.
	//
	// *   **VPC**: a virtual private cloud (VPC)
	// *   **Classic**: classic network
	InstanceNetworkType *string `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBInstanceNetInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceNetInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceNetInfoResponseBody) SetDBInstanceNetInfos(v *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos) *DescribeDBInstanceNetInfoResponseBody {
	s.DBInstanceNetInfos = v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBody) SetInstanceNetworkType(v string) *DescribeDBInstanceNetInfoResponseBody {
	s.InstanceNetworkType = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBody) SetRequestId(v string) *DescribeDBInstanceNetInfoResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos struct {
	DBInstanceNetInfo []*DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo `json:"DBInstanceNetInfo,omitempty" xml:"DBInstanceNetInfo,omitempty" type:"Repeated"`
}

func (s DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos) SetDBInstanceNetInfo(v []*DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo) *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos {
	s.DBInstanceNetInfo = v
	return s
}

type DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo struct {
	// The IP address type of the instance.
	AddressType *string `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	// The endpoint used to connect to the instance.
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// The IP address of the instance.
	IPAddress *string `json:"IPAddress,omitempty" xml:"IPAddress,omitempty"`
	// The type of IP address.
	//
	// *   Valid values for instances in the classic network: Inner and Public
	// *   Valid values for instances in a virtual private cloud (VPC): Private and Public
	IPType *string `json:"IPType,omitempty" xml:"IPType,omitempty"`
	// The port number used to connect to the instance.
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The ID of the VPC.
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// The ID of the vSwitch. Multiple IDs are separated by commas (,).
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the VPC.
	VpcInstanceId *string `json:"VpcInstanceId,omitempty" xml:"VpcInstanceId,omitempty"`
}

func (s DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo) SetAddressType(v string) *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo {
	s.AddressType = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo) SetConnectionString(v string) *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo {
	s.ConnectionString = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo) SetIPAddress(v string) *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo {
	s.IPAddress = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo) SetIPType(v string) *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo {
	s.IPType = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo) SetPort(v string) *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo {
	s.Port = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo) SetVPCId(v string) *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo {
	s.VPCId = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo) SetVSwitchId(v string) *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo) SetVpcInstanceId(v string) *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo {
	s.VpcInstanceId = &v
	return s
}

type DescribeDBInstanceNetInfoResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDBInstanceNetInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDBInstanceNetInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceNetInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceNetInfoResponse) SetHeaders(v map[string]*string) *DescribeDBInstanceNetInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInstanceNetInfoResponse) SetStatusCode(v int32) *DescribeDBInstanceNetInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponse) SetBody(v *DescribeDBInstanceNetInfoResponseBody) *DescribeDBInstanceNetInfoResponse {
	s.Body = v
	return s
}

type DescribeDBInstancePerformanceRequest struct {
	// The ID of the instance.
	//
	// >  You can call the [DescribeDBInstances](~~86911~~) operation to query the instance IDs of all AnalyticDB for PostgreSQL instances in a specific region.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm*Z format. The time must be in UTC. The end time must be later than the start time.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The performance metric. Separate multiple values with commas (,). For more information, see [Performance parameters](~~86943~~).
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The ID of the resource group to which the instance belongs. For more information about how to obtain the ID of a resource group, see [View basic information of a resource group](~~151181~~).
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm*Z format. The time must be in UTC.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDBInstancePerformanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancePerformanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancePerformanceRequest) SetDBInstanceId(v string) *DescribeDBInstancePerformanceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstancePerformanceRequest) SetEndTime(v string) *DescribeDBInstancePerformanceRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDBInstancePerformanceRequest) SetKey(v string) *DescribeDBInstancePerformanceRequest {
	s.Key = &v
	return s
}

func (s *DescribeDBInstancePerformanceRequest) SetResourceGroupId(v string) *DescribeDBInstancePerformanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDBInstancePerformanceRequest) SetStartTime(v string) *DescribeDBInstancePerformanceRequest {
	s.StartTime = &v
	return s
}

type DescribeDBInstancePerformanceResponseBody struct {
	// The ID of the instance.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The end time of the query.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The database engine of the instance.
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// Details of the performance metrics. Format: {perf1, perf2, perf3, …}.
	PerformanceKeys []*string `json:"PerformanceKeys,omitempty" xml:"PerformanceKeys,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The start time of the query.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDBInstancePerformanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancePerformanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancePerformanceResponseBody) SetDBInstanceId(v string) *DescribeDBInstancePerformanceResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstancePerformanceResponseBody) SetEndTime(v string) *DescribeDBInstancePerformanceResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDBInstancePerformanceResponseBody) SetEngine(v string) *DescribeDBInstancePerformanceResponseBody {
	s.Engine = &v
	return s
}

func (s *DescribeDBInstancePerformanceResponseBody) SetPerformanceKeys(v []*string) *DescribeDBInstancePerformanceResponseBody {
	s.PerformanceKeys = v
	return s
}

func (s *DescribeDBInstancePerformanceResponseBody) SetRequestId(v string) *DescribeDBInstancePerformanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstancePerformanceResponseBody) SetStartTime(v string) *DescribeDBInstancePerformanceResponseBody {
	s.StartTime = &v
	return s
}

type DescribeDBInstancePerformanceResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDBInstancePerformanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDBInstancePerformanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancePerformanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancePerformanceResponse) SetHeaders(v map[string]*string) *DescribeDBInstancePerformanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInstancePerformanceResponse) SetStatusCode(v int32) *DescribeDBInstancePerformanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBInstancePerformanceResponse) SetBody(v *DescribeDBInstancePerformanceResponseBody) *DescribeDBInstancePerformanceResponse {
	s.Body = v
	return s
}

type DescribeDBInstancePlansRequest struct {
	// The ID of the instance.
	//
	// >  You can call the [DescribeDBInstances](~~86911~~) operation to query the details of all AnalyticDB for PostgreSQL instances in a specific region, including instance IDs.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The time used to filter plans. If you specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm*Z format, the plans created before this time are returned. The time must be in UTC. If you do not specify this parameter, all plans are returned.
	PlanCreateDate *string `json:"PlanCreateDate,omitempty" xml:"PlanCreateDate,omitempty"`
	// The description of the plan.
	PlanDesc *string `json:"PlanDesc,omitempty" xml:"PlanDesc,omitempty"`
	// The ID of the plan.
	//
	// >  You can call the [DescribeDBInstancePlans](~~449398~~) operation to query the details of plans, including plan IDs.
	PlanId *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	// The execution mode of the plan. Valid values:
	//
	// *   **Postpone**: The plan is executed later.
	// *   **Regular**: The plan is executed periodically.
	PlanScheduleType *string `json:"PlanScheduleType,omitempty" xml:"PlanScheduleType,omitempty"`
	// The type of the plan. Valid values:
	//
	// *   **PauseResume**: pauses and resumes an instance.
	// *   **Resize**: scales an instance.
	PlanType *string `json:"PlanType,omitempty" xml:"PlanType,omitempty"`
}

func (s DescribeDBInstancePlansRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancePlansRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancePlansRequest) SetDBInstanceId(v string) *DescribeDBInstancePlansRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstancePlansRequest) SetOwnerId(v int64) *DescribeDBInstancePlansRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBInstancePlansRequest) SetPlanCreateDate(v string) *DescribeDBInstancePlansRequest {
	s.PlanCreateDate = &v
	return s
}

func (s *DescribeDBInstancePlansRequest) SetPlanDesc(v string) *DescribeDBInstancePlansRequest {
	s.PlanDesc = &v
	return s
}

func (s *DescribeDBInstancePlansRequest) SetPlanId(v string) *DescribeDBInstancePlansRequest {
	s.PlanId = &v
	return s
}

func (s *DescribeDBInstancePlansRequest) SetPlanScheduleType(v string) *DescribeDBInstancePlansRequest {
	s.PlanScheduleType = &v
	return s
}

func (s *DescribeDBInstancePlansRequest) SetPlanType(v string) *DescribeDBInstancePlansRequest {
	s.PlanType = &v
	return s
}

type DescribeDBInstancePlansResponseBody struct {
	// The error message returned.
	//
	// This parameter is returned only when the operation fails.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// Details of the plans.
	Items *DescribeDBInstancePlansResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The page number of the returned page.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned on the current page.
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The state of the operation.
	//
	// If the operation is successful, **success** is returned. If the operation fails, this parameter is not returned.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The total number of entries.
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeDBInstancePlansResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancePlansResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancePlansResponseBody) SetErrorMessage(v string) *DescribeDBInstancePlansResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeDBInstancePlansResponseBody) SetItems(v *DescribeDBInstancePlansResponseBodyItems) *DescribeDBInstancePlansResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDBInstancePlansResponseBody) SetPageNumber(v int32) *DescribeDBInstancePlansResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBInstancePlansResponseBody) SetPageRecordCount(v int32) *DescribeDBInstancePlansResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeDBInstancePlansResponseBody) SetRequestId(v string) *DescribeDBInstancePlansResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstancePlansResponseBody) SetStatus(v string) *DescribeDBInstancePlansResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeDBInstancePlansResponseBody) SetTotalRecordCount(v int32) *DescribeDBInstancePlansResponseBody {
	s.TotalRecordCount = &v
	return s
}

type DescribeDBInstancePlansResponseBodyItems struct {
	PlanList []*DescribeDBInstancePlansResponseBodyItemsPlanList `json:"PlanList,omitempty" xml:"PlanList,omitempty" type:"Repeated"`
}

func (s DescribeDBInstancePlansResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancePlansResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancePlansResponseBodyItems) SetPlanList(v []*DescribeDBInstancePlansResponseBodyItemsPlanList) *DescribeDBInstancePlansResponseBodyItems {
	s.PlanList = v
	return s
}

type DescribeDBInstancePlansResponseBodyItemsPlanList struct {
	// The ID of the instance.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The execution information of the plan.
	PlanConfig *string `json:"PlanConfig,omitempty" xml:"PlanConfig,omitempty"`
	// The description of the plan.
	PlanDesc *string `json:"PlanDesc,omitempty" xml:"PlanDesc,omitempty"`
	// The end time of the plan. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm*Z format. The time is displayed in UTC.
	//
	// >  This parameter is returned only for periodically executed plans.
	PlanEndDate *string `json:"PlanEndDate,omitempty" xml:"PlanEndDate,omitempty"`
	// The ID of the plan.
	PlanId *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	// The name of the plan.
	PlanName *string `json:"PlanName,omitempty" xml:"PlanName,omitempty"`
	// The execution mode of the plan. Valid values:
	//
	// *   **Postpone**: The plan is executed later.
	// *   **Regular**: The plan is executed periodically.
	PlanScheduleType *string `json:"PlanScheduleType,omitempty" xml:"PlanScheduleType,omitempty"`
	// The start time of the plan. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm*Z format. The time is displayed in UTC.
	//
	// >  This parameter is returned only for periodically executed plans.
	PlanStartDate *string `json:"PlanStartDate,omitempty" xml:"PlanStartDate,omitempty"`
	// The state of the plan. Valid values:
	//
	// *   **active**: The plan is running.
	// *   **cancel**: The plan is canceled.
	// *   **deleted**: The plan is deleted.
	// *   **finished**: The plan execution is complete.
	PlanStatus *string `json:"PlanStatus,omitempty" xml:"PlanStatus,omitempty"`
	// The type of the plan. Valid values:
	//
	// *   **PauseResume**: pauses and resumes an instance.
	// *   **Resize**: scales an instance.
	PlanType *string `json:"PlanType,omitempty" xml:"PlanType,omitempty"`
}

func (s DescribeDBInstancePlansResponseBodyItemsPlanList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancePlansResponseBodyItemsPlanList) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancePlansResponseBodyItemsPlanList) SetDBInstanceId(v string) *DescribeDBInstancePlansResponseBodyItemsPlanList {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstancePlansResponseBodyItemsPlanList) SetPlanConfig(v string) *DescribeDBInstancePlansResponseBodyItemsPlanList {
	s.PlanConfig = &v
	return s
}

func (s *DescribeDBInstancePlansResponseBodyItemsPlanList) SetPlanDesc(v string) *DescribeDBInstancePlansResponseBodyItemsPlanList {
	s.PlanDesc = &v
	return s
}

func (s *DescribeDBInstancePlansResponseBodyItemsPlanList) SetPlanEndDate(v string) *DescribeDBInstancePlansResponseBodyItemsPlanList {
	s.PlanEndDate = &v
	return s
}

func (s *DescribeDBInstancePlansResponseBodyItemsPlanList) SetPlanId(v string) *DescribeDBInstancePlansResponseBodyItemsPlanList {
	s.PlanId = &v
	return s
}

func (s *DescribeDBInstancePlansResponseBodyItemsPlanList) SetPlanName(v string) *DescribeDBInstancePlansResponseBodyItemsPlanList {
	s.PlanName = &v
	return s
}

func (s *DescribeDBInstancePlansResponseBodyItemsPlanList) SetPlanScheduleType(v string) *DescribeDBInstancePlansResponseBodyItemsPlanList {
	s.PlanScheduleType = &v
	return s
}

func (s *DescribeDBInstancePlansResponseBodyItemsPlanList) SetPlanStartDate(v string) *DescribeDBInstancePlansResponseBodyItemsPlanList {
	s.PlanStartDate = &v
	return s
}

func (s *DescribeDBInstancePlansResponseBodyItemsPlanList) SetPlanStatus(v string) *DescribeDBInstancePlansResponseBodyItemsPlanList {
	s.PlanStatus = &v
	return s
}

func (s *DescribeDBInstancePlansResponseBodyItemsPlanList) SetPlanType(v string) *DescribeDBInstancePlansResponseBodyItemsPlanList {
	s.PlanType = &v
	return s
}

type DescribeDBInstancePlansResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDBInstancePlansResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDBInstancePlansResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancePlansResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancePlansResponse) SetHeaders(v map[string]*string) *DescribeDBInstancePlansResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInstancePlansResponse) SetStatusCode(v int32) *DescribeDBInstancePlansResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBInstancePlansResponse) SetBody(v *DescribeDBInstancePlansResponseBody) *DescribeDBInstancePlansResponse {
	s.Body = v
	return s
}

type DescribeDBInstanceSSLRequest struct {
	// The ID of the instance.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s DescribeDBInstanceSSLRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceSSLRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceSSLRequest) SetDBInstanceId(v string) *DescribeDBInstanceSSLRequest {
	s.DBInstanceId = &v
	return s
}

type DescribeDBInstanceSSLResponseBody struct {
	// The name of the SSL certificate.
	CertCommonName *string `json:"CertCommonName,omitempty" xml:"CertCommonName,omitempty"`
	// The ID of the instance.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The name of the instance.
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether SSL encryption is enabled.
	SSLEnabled *bool `json:"SSLEnabled,omitempty" xml:"SSLEnabled,omitempty"`
	// The expiration time of the SSL certificate.
	SSLExpiredTime *string `json:"SSLExpiredTime,omitempty" xml:"SSLExpiredTime,omitempty"`
}

func (s DescribeDBInstanceSSLResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceSSLResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceSSLResponseBody) SetCertCommonName(v string) *DescribeDBInstanceSSLResponseBody {
	s.CertCommonName = &v
	return s
}

func (s *DescribeDBInstanceSSLResponseBody) SetDBInstanceId(v string) *DescribeDBInstanceSSLResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceSSLResponseBody) SetDBInstanceName(v string) *DescribeDBInstanceSSLResponseBody {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeDBInstanceSSLResponseBody) SetRequestId(v string) *DescribeDBInstanceSSLResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstanceSSLResponseBody) SetSSLEnabled(v bool) *DescribeDBInstanceSSLResponseBody {
	s.SSLEnabled = &v
	return s
}

func (s *DescribeDBInstanceSSLResponseBody) SetSSLExpiredTime(v string) *DescribeDBInstanceSSLResponseBody {
	s.SSLExpiredTime = &v
	return s
}

type DescribeDBInstanceSSLResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDBInstanceSSLResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDBInstanceSSLResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceSSLResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceSSLResponse) SetHeaders(v map[string]*string) *DescribeDBInstanceSSLResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInstanceSSLResponse) SetStatusCode(v int32) *DescribeDBInstanceSSLResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBInstanceSSLResponse) SetBody(v *DescribeDBInstanceSSLResponseBody) *DescribeDBInstanceSSLResponse {
	s.Body = v
	return s
}

type DescribeDBInstancesRequest struct {
	// The edition of the instance. Separate multiple values with commas (,). Valid values:
	//
	// *   **basic**: Basic Edition
	// *   **highavailability**: High-availability Edition
	// *   **finance**: Enterprise Edition
	DBInstanceCategories []*string `json:"DBInstanceCategories,omitempty" xml:"DBInstanceCategories,omitempty" type:"Repeated"`
	// The description of the instance.
	DBInstanceDescription *string `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	// The ID of the instance. Separate multiple IDs with commas (,).
	DBInstanceIds *string `json:"DBInstanceIds,omitempty" xml:"DBInstanceIds,omitempty"`
	// The resource type of the instance. Separate multiple values with commas (,). Valid values:
	//
	// *   **serverless**: Serverless mode
	// *   **storageelastic**: elastic storage mode
	// *   **classic**: reserved storage mode
	DBInstanceModes []*string `json:"DBInstanceModes,omitempty" xml:"DBInstanceModes,omitempty" type:"Repeated"`
	// The state of the instance. Separate multiple values with commas (,). For more information, see [Instance statuses](~~86944~~).
	//
	// >  The value of this parameter must be in lowercase.
	DBInstanceStatuses []*string `json:"DBInstanceStatuses,omitempty" xml:"DBInstanceStatuses,omitempty" type:"Repeated"`
	// This parameter is no longer used.
	InstanceDeployTypes []*string `json:"InstanceDeployTypes,omitempty" xml:"InstanceDeployTypes,omitempty" type:"Repeated"`
	// The network type of the instance. Valid values:
	//
	// *   **VPC**
	// *   **Classic**
	//
	// >  If you do not specify this parameter, instances of both network types are returned.
	InstanceNetworkType *string `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
	OwnerId             *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the page to return. The value must be an integer that is greater than 0. Default value: **1**.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values:
	//
	// *   **30**
	// *   **50**
	// *   **100**
	//
	// Default value: **30**.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the instance.
	//
	// >  You can call the [DescribeRegions](~~86912~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the instance belongs.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The list of tags.
	Tag []*DescribeDBInstancesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeDBInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesRequest) SetDBInstanceCategories(v []*string) *DescribeDBInstancesRequest {
	s.DBInstanceCategories = v
	return s
}

func (s *DescribeDBInstancesRequest) SetDBInstanceDescription(v string) *DescribeDBInstancesRequest {
	s.DBInstanceDescription = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetDBInstanceIds(v string) *DescribeDBInstancesRequest {
	s.DBInstanceIds = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetDBInstanceModes(v []*string) *DescribeDBInstancesRequest {
	s.DBInstanceModes = v
	return s
}

func (s *DescribeDBInstancesRequest) SetDBInstanceStatuses(v []*string) *DescribeDBInstancesRequest {
	s.DBInstanceStatuses = v
	return s
}

func (s *DescribeDBInstancesRequest) SetInstanceDeployTypes(v []*string) *DescribeDBInstancesRequest {
	s.InstanceDeployTypes = v
	return s
}

func (s *DescribeDBInstancesRequest) SetInstanceNetworkType(v string) *DescribeDBInstancesRequest {
	s.InstanceNetworkType = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetOwnerId(v int64) *DescribeDBInstancesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetPageNumber(v int32) *DescribeDBInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetPageSize(v int32) *DescribeDBInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetRegionId(v string) *DescribeDBInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetResourceGroupId(v string) *DescribeDBInstancesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetTag(v []*DescribeDBInstancesRequestTag) *DescribeDBInstancesRequest {
	s.Tag = v
	return s
}

type DescribeDBInstancesRequestTag struct {
	// The key of tag N.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDBInstancesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancesRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesRequestTag) SetKey(v string) *DescribeDBInstancesRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeDBInstancesRequestTag) SetValue(v string) *DescribeDBInstancesRequestTag {
	s.Value = &v
	return s
}

type DescribeDBInstancesShrinkRequest struct {
	// The edition of the instance. Separate multiple values with commas (,). Valid values:
	//
	// *   **basic**: Basic Edition
	// *   **highavailability**: High-availability Edition
	// *   **finance**: Enterprise Edition
	DBInstanceCategoriesShrink *string `json:"DBInstanceCategories,omitempty" xml:"DBInstanceCategories,omitempty"`
	// The description of the instance.
	DBInstanceDescription *string `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	// The ID of the instance. Separate multiple IDs with commas (,).
	DBInstanceIds *string `json:"DBInstanceIds,omitempty" xml:"DBInstanceIds,omitempty"`
	// The resource type of the instance. Separate multiple values with commas (,). Valid values:
	//
	// *   **serverless**: Serverless mode
	// *   **storageelastic**: elastic storage mode
	// *   **classic**: reserved storage mode
	DBInstanceModesShrink *string `json:"DBInstanceModes,omitempty" xml:"DBInstanceModes,omitempty"`
	// The state of the instance. Separate multiple values with commas (,). For more information, see [Instance statuses](~~86944~~).
	//
	// >  The value of this parameter must be in lowercase.
	DBInstanceStatusesShrink *string `json:"DBInstanceStatuses,omitempty" xml:"DBInstanceStatuses,omitempty"`
	// This parameter is no longer used.
	InstanceDeployTypesShrink *string `json:"InstanceDeployTypes,omitempty" xml:"InstanceDeployTypes,omitempty"`
	// The network type of the instance. Valid values:
	//
	// *   **VPC**
	// *   **Classic**
	//
	// >  If you do not specify this parameter, instances of both network types are returned.
	InstanceNetworkType *string `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
	OwnerId             *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the page to return. The value must be an integer that is greater than 0. Default value: **1**.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values:
	//
	// *   **30**
	// *   **50**
	// *   **100**
	//
	// Default value: **30**.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the instance.
	//
	// >  You can call the [DescribeRegions](~~86912~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the instance belongs.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The list of tags.
	Tag []*DescribeDBInstancesShrinkRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeDBInstancesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancesShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesShrinkRequest) SetDBInstanceCategoriesShrink(v string) *DescribeDBInstancesShrinkRequest {
	s.DBInstanceCategoriesShrink = &v
	return s
}

func (s *DescribeDBInstancesShrinkRequest) SetDBInstanceDescription(v string) *DescribeDBInstancesShrinkRequest {
	s.DBInstanceDescription = &v
	return s
}

func (s *DescribeDBInstancesShrinkRequest) SetDBInstanceIds(v string) *DescribeDBInstancesShrinkRequest {
	s.DBInstanceIds = &v
	return s
}

func (s *DescribeDBInstancesShrinkRequest) SetDBInstanceModesShrink(v string) *DescribeDBInstancesShrinkRequest {
	s.DBInstanceModesShrink = &v
	return s
}

func (s *DescribeDBInstancesShrinkRequest) SetDBInstanceStatusesShrink(v string) *DescribeDBInstancesShrinkRequest {
	s.DBInstanceStatusesShrink = &v
	return s
}

func (s *DescribeDBInstancesShrinkRequest) SetInstanceDeployTypesShrink(v string) *DescribeDBInstancesShrinkRequest {
	s.InstanceDeployTypesShrink = &v
	return s
}

func (s *DescribeDBInstancesShrinkRequest) SetInstanceNetworkType(v string) *DescribeDBInstancesShrinkRequest {
	s.InstanceNetworkType = &v
	return s
}

func (s *DescribeDBInstancesShrinkRequest) SetOwnerId(v int64) *DescribeDBInstancesShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBInstancesShrinkRequest) SetPageNumber(v int32) *DescribeDBInstancesShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBInstancesShrinkRequest) SetPageSize(v int32) *DescribeDBInstancesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDBInstancesShrinkRequest) SetRegionId(v string) *DescribeDBInstancesShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBInstancesShrinkRequest) SetResourceGroupId(v string) *DescribeDBInstancesShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDBInstancesShrinkRequest) SetTag(v []*DescribeDBInstancesShrinkRequestTag) *DescribeDBInstancesShrinkRequest {
	s.Tag = v
	return s
}

type DescribeDBInstancesShrinkRequestTag struct {
	// The key of tag N.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDBInstancesShrinkRequestTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancesShrinkRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesShrinkRequestTag) SetKey(v string) *DescribeDBInstancesShrinkRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeDBInstancesShrinkRequestTag) SetValue(v string) *DescribeDBInstancesShrinkRequestTag {
	s.Value = &v
	return s
}

type DescribeDBInstancesResponseBody struct {
	// Details of the instance.
	Items *DescribeDBInstancesResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The page number of the returned page.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned on the current page.
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeDBInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesResponseBody) SetItems(v *DescribeDBInstancesResponseBodyItems) *DescribeDBInstancesResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDBInstancesResponseBody) SetPageNumber(v int32) *DescribeDBInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBInstancesResponseBody) SetPageRecordCount(v int32) *DescribeDBInstancesResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeDBInstancesResponseBody) SetRequestId(v string) *DescribeDBInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstancesResponseBody) SetTotalRecordCount(v int32) *DescribeDBInstancesResponseBody {
	s.TotalRecordCount = &v
	return s
}

type DescribeDBInstancesResponseBodyItems struct {
	DBInstance []*DescribeDBInstancesResponseBodyItemsDBInstance `json:"DBInstance,omitempty" xml:"DBInstance,omitempty" type:"Repeated"`
}

func (s DescribeDBInstancesResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesResponseBodyItems) SetDBInstance(v []*DescribeDBInstancesResponseBodyItemsDBInstance) *DescribeDBInstancesResponseBodyItems {
	s.DBInstance = v
	return s
}

type DescribeDBInstancesResponseBodyItemsDBInstance struct {
	// An invalid parameter. It is no longer returned when you call this operation.
	//
	// You can call the [DescribeDBInstanceAttribute](~~86910~~) operation to query the access mode of an instance.
	ConnectionMode *string `json:"ConnectionMode,omitempty" xml:"ConnectionMode,omitempty"`
	// The time when the instance was created. The time is displayed in UTC.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The edition of the instance. Valid values:
	//
	// *   **Basic**: Basic Edition
	// *   **HighAvailability**: High-availability Edition
	// *   **Finance**: Enterprise Edition
	DBInstanceCategory *string `json:"DBInstanceCategory,omitempty" xml:"DBInstanceCategory,omitempty"`
	// The description of the instance.
	DBInstanceDescription *string `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	// The ID of the instance.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The resource type of the instance. Valid values:
	//
	// *   **Serverless**: Serverless mode
	// *   **StorageElastic**: elastic storage mode
	// *   **Classic**: reserved storage mode
	DBInstanceMode *string `json:"DBInstanceMode,omitempty" xml:"DBInstanceMode,omitempty"`
	// The type of the network interface card (NIC) that is used by the instance. Valid values:
	//
	// *   **0**: Internet
	// *   **1**: internal network
	// *   **2**: VPC
	DBInstanceNetType *string `json:"DBInstanceNetType,omitempty" xml:"DBInstanceNetType,omitempty"`
	// The state of the instance. For more information, see [Instance statuses](~~86944~~).
	DBInstanceStatus *string `json:"DBInstanceStatus,omitempty" xml:"DBInstanceStatus,omitempty"`
	// The database engine that the instance runs.
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The version of the database engine.
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The expiration time of the instance. The time is displayed in UTC.
	//
	// >  For pay-as-you-go instances, `2999-09-08T16:00:00Z` is returned.
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The resource type of the instance. Valid values:
	//
	// *   **cluster**: elastic storage mode or Serverless mode
	// *   **replicaSet**: reserved storage mode
	InstanceDeployType *string `json:"InstanceDeployType,omitempty" xml:"InstanceDeployType,omitempty"`
	// The network type of the instance. Valid values:
	//
	// *   **Classic**
	// *   **VPC**
	InstanceNetworkType *string `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
	// The lock mode of the instance. Valid values:
	//
	// *   **Unlock**: The instance is not locked.
	// *   **ManualLock**: The instance is manually locked.
	// *   **LockByExpiration**: The instance is automatically locked due to instance expiration.
	// *   **LockByRestoration**: The instance is automatically locked due to instance restoration.
	// *   **LockByDiskQuota**: The instance is automatically locked due to exhausted storage.
	// *   **LockReadInstanceByDiskQuota**: The instance is a read-only instance and is automatically locked due to exhausted storage.
	LockMode *string `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	// The reason why the cluster is locked.
	//
	// >  This parameter is returned only when the cluster is locked. The value is **instance_expire**.
	LockReason *string `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	// The number of coordinator nodes.
	MasterNodeNum *int32 `json:"MasterNodeNum,omitempty" xml:"MasterNodeNum,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// *   **Postpaid**: pay-as-you-go
	// *   **Prepaid**: subscription
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The region ID of the instance.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the instance belongs.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The number of compute nodes.
	SegNodeNum *string `json:"SegNodeNum,omitempty" xml:"SegNodeNum,omitempty"`
	// The type of the Serverless mode. Valid values:
	//
	// *   **Manual**: manual scheduling
	// *   **Auto**: automatic scheduling
	//
	// >  This parameter is returned only for instances in Serverless mode.
	ServerlessMode *string `json:"ServerlessMode,omitempty" xml:"ServerlessMode,omitempty"`
	// The storage capacity. Unit: GB.
	StorageSize *string `json:"StorageSize,omitempty" xml:"StorageSize,omitempty"`
	// The storage type of the instance. Valid values:
	//
	// *   **cloud_essd**: enhanced SSD (ESSD)
	// *   **cloud_efficiency**: ultra disk
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// The tags of the instance. Each tag is a key-value pair.
	Tags *DescribeDBInstancesResponseBodyItemsDBInstanceTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The ID of the vSwitch.
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of virtual private cloud (VPC).
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The zone ID of the instance.
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDBInstancesResponseBodyItemsDBInstance) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancesResponseBodyItemsDBInstance) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetConnectionMode(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.ConnectionMode = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetCreateTime(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.CreateTime = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetDBInstanceCategory(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.DBInstanceCategory = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetDBInstanceDescription(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.DBInstanceDescription = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetDBInstanceId(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetDBInstanceMode(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.DBInstanceMode = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetDBInstanceNetType(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.DBInstanceNetType = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetDBInstanceStatus(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.DBInstanceStatus = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetEngine(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.Engine = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetEngineVersion(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.EngineVersion = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetExpireTime(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetInstanceDeployType(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.InstanceDeployType = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetInstanceNetworkType(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.InstanceNetworkType = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetLockMode(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.LockMode = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetLockReason(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.LockReason = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetMasterNodeNum(v int32) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.MasterNodeNum = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetPayType(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.PayType = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetRegionId(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.RegionId = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetResourceGroupId(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetSegNodeNum(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.SegNodeNum = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetServerlessMode(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.ServerlessMode = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetStorageSize(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.StorageSize = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetStorageType(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.StorageType = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetTags(v *DescribeDBInstancesResponseBodyItemsDBInstanceTags) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.Tags = v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetVSwitchId(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetVpcId(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.VpcId = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetZoneId(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.ZoneId = &v
	return s
}

type DescribeDBInstancesResponseBodyItemsDBInstanceTags struct {
	Tag []*DescribeDBInstancesResponseBodyItemsDBInstanceTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeDBInstancesResponseBodyItemsDBInstanceTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancesResponseBodyItemsDBInstanceTags) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstanceTags) SetTag(v []*DescribeDBInstancesResponseBodyItemsDBInstanceTagsTag) *DescribeDBInstancesResponseBodyItemsDBInstanceTags {
	s.Tag = v
	return s
}

type DescribeDBInstancesResponseBodyItemsDBInstanceTagsTag struct {
	// The key of the tag.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDBInstancesResponseBodyItemsDBInstanceTagsTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancesResponseBodyItemsDBInstanceTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstanceTagsTag) SetKey(v string) *DescribeDBInstancesResponseBodyItemsDBInstanceTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstanceTagsTag) SetValue(v string) *DescribeDBInstancesResponseBodyItemsDBInstanceTagsTag {
	s.Value = &v
	return s
}

type DescribeDBInstancesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDBInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDBInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesResponse) SetHeaders(v map[string]*string) *DescribeDBInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInstancesResponse) SetStatusCode(v int32) *DescribeDBInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBInstancesResponse) SetBody(v *DescribeDBInstancesResponseBody) *DescribeDBInstancesResponse {
	s.Body = v
	return s
}

type DescribeDataBackupsRequest struct {
	// The ID of the backup set. If you specify the BackupId parameter, the details of the backup set are returned.
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The backup mode. Valid values:
	//
	// *   Automated: automatic backup
	// *   Manual: manual backup
	//
	// If you do not specify this parameter, the records of the backup sets in all modes are returned.
	BackupMode *string `json:"BackupMode,omitempty" xml:"BackupMode,omitempty"`
	// The status of the backup set. Valid values:
	//
	// *   Success: The backup is complete.
	// *   Failed: The backup task fails.
	//
	// If you do not specify this parameter, the records of the backup sets in all states are returned.
	BackupStatus *string `json:"BackupStatus,omitempty" xml:"BackupStatus,omitempty"`
	// The ID of the instance.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The type of the backup. Valid values:
	//
	// *   DATA: full backup
	// *   RESTOREPOI: point-in-time backup
	//
	// If you do not specify this parameter, the records of the full backup set are returned.
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// The end of the time range to query. The end time must be later than the start time. Specify the time in the yyyy-MM-ddTHH:mmZ format. The time must be in UTC.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The number of the page to return. The value must be an integer that is larger than 0. Default value: 1.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values:
	//
	// *   30
	// *   50
	// *   100
	//
	// Default value: 30.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The beginning of the time range to query. Specify the time in the yyyy-MM-ddTHH:mmZ format. The time must be in UTC.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDataBackupsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataBackupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDataBackupsRequest) SetBackupId(v string) *DescribeDataBackupsRequest {
	s.BackupId = &v
	return s
}

func (s *DescribeDataBackupsRequest) SetBackupMode(v string) *DescribeDataBackupsRequest {
	s.BackupMode = &v
	return s
}

func (s *DescribeDataBackupsRequest) SetBackupStatus(v string) *DescribeDataBackupsRequest {
	s.BackupStatus = &v
	return s
}

func (s *DescribeDataBackupsRequest) SetDBInstanceId(v string) *DescribeDataBackupsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDataBackupsRequest) SetDataType(v string) *DescribeDataBackupsRequest {
	s.DataType = &v
	return s
}

func (s *DescribeDataBackupsRequest) SetEndTime(v string) *DescribeDataBackupsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDataBackupsRequest) SetPageNumber(v int32) *DescribeDataBackupsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDataBackupsRequest) SetPageSize(v int32) *DescribeDataBackupsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDataBackupsRequest) SetStartTime(v string) *DescribeDataBackupsRequest {
	s.StartTime = &v
	return s
}

type DescribeDataBackupsResponseBody struct {
	// Details about the backup sets.
	Items []*DescribeDataBackupsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The page number of the returned page.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of backup sets on the page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDataBackupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataBackupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDataBackupsResponseBody) SetItems(v []*DescribeDataBackupsResponseBodyItems) *DescribeDataBackupsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDataBackupsResponseBody) SetPageNumber(v int32) *DescribeDataBackupsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDataBackupsResponseBody) SetPageSize(v int32) *DescribeDataBackupsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDataBackupsResponseBody) SetRequestId(v string) *DescribeDataBackupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDataBackupsResponseBody) SetTotalCount(v int32) *DescribeDataBackupsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeDataBackupsResponseBodyItems struct {
	// The UTC time when the backup ended. The time is in the yyyy-MM-ddTHH:mmZ format. The time is displayed in UTC.
	BackupEndTime *string `json:"BackupEndTime,omitempty" xml:"BackupEndTime,omitempty"`
	// The local time when the backup ended. The time is in the yyyy-MM-dd HH:mm:ss format. The time is your local time.
	BackupEndTimeLocal *string `json:"BackupEndTimeLocal,omitempty" xml:"BackupEndTimeLocal,omitempty"`
	// The backup mode.
	//
	// Valid values for full backup:
	//
	// *   Automated: automatic backup
	// *   Manual: manual backup
	//
	// Valid values for point-in-time backup:
	//
	// *   Automated: point-in-time backup after full backup
	// *   Manual: manual point-in-time backup
	// *   Period: point-in-time backup that is triggered by a backup policy
	BackupMode *string `json:"BackupMode,omitempty" xml:"BackupMode,omitempty"`
	// The ID of the backup set.
	BackupSetId *string `json:"BackupSetId,omitempty" xml:"BackupSetId,omitempty"`
	// The size of the backup file. Unit: bytes.
	BackupSize *int64 `json:"BackupSize,omitempty" xml:"BackupSize,omitempty"`
	// The UTC time when the backup started. The time is in the yyyy-MM-ddTHH:mmZ format. The time is displayed in UTC.
	BackupStartTime *string `json:"BackupStartTime,omitempty" xml:"BackupStartTime,omitempty"`
	// The local time when the backup started. The time is in the yyyy-MM-dd HH:mm:ss format. The time is your local time.
	BackupStartTimeLocal *string `json:"BackupStartTimeLocal,omitempty" xml:"BackupStartTimeLocal,omitempty"`
	// The status of the backup set. Valid values:
	//
	// *   Success
	// *   Failure
	BackupStatus *string `json:"BackupStatus,omitempty" xml:"BackupStatus,omitempty"`
	// The name of a point-in-time backup set or the full backup set.
	BaksetName *string `json:"BaksetName,omitempty" xml:"BaksetName,omitempty"`
	// *   For full backup, this parameter indicates the point in time at which the data in the data backup file is consistent.
	// *   For point-in-time backup, this parameter indicates that the returned point in time is a timestamp.
	ConsistentTime *int64 `json:"ConsistentTime,omitempty" xml:"ConsistentTime,omitempty"`
	// The ID of the instance.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The type of the backup. Valid values:
	//
	// *   DATA: full backup
	// *   RESTOREPOI: point-in-time backup
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
}

func (s DescribeDataBackupsResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataBackupsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDataBackupsResponseBodyItems) SetBackupEndTime(v string) *DescribeDataBackupsResponseBodyItems {
	s.BackupEndTime = &v
	return s
}

func (s *DescribeDataBackupsResponseBodyItems) SetBackupEndTimeLocal(v string) *DescribeDataBackupsResponseBodyItems {
	s.BackupEndTimeLocal = &v
	return s
}

func (s *DescribeDataBackupsResponseBodyItems) SetBackupMode(v string) *DescribeDataBackupsResponseBodyItems {
	s.BackupMode = &v
	return s
}

func (s *DescribeDataBackupsResponseBodyItems) SetBackupSetId(v string) *DescribeDataBackupsResponseBodyItems {
	s.BackupSetId = &v
	return s
}

func (s *DescribeDataBackupsResponseBodyItems) SetBackupSize(v int64) *DescribeDataBackupsResponseBodyItems {
	s.BackupSize = &v
	return s
}

func (s *DescribeDataBackupsResponseBodyItems) SetBackupStartTime(v string) *DescribeDataBackupsResponseBodyItems {
	s.BackupStartTime = &v
	return s
}

func (s *DescribeDataBackupsResponseBodyItems) SetBackupStartTimeLocal(v string) *DescribeDataBackupsResponseBodyItems {
	s.BackupStartTimeLocal = &v
	return s
}

func (s *DescribeDataBackupsResponseBodyItems) SetBackupStatus(v string) *DescribeDataBackupsResponseBodyItems {
	s.BackupStatus = &v
	return s
}

func (s *DescribeDataBackupsResponseBodyItems) SetBaksetName(v string) *DescribeDataBackupsResponseBodyItems {
	s.BaksetName = &v
	return s
}

func (s *DescribeDataBackupsResponseBodyItems) SetConsistentTime(v int64) *DescribeDataBackupsResponseBodyItems {
	s.ConsistentTime = &v
	return s
}

func (s *DescribeDataBackupsResponseBodyItems) SetDBInstanceId(v string) *DescribeDataBackupsResponseBodyItems {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDataBackupsResponseBodyItems) SetDataType(v string) *DescribeDataBackupsResponseBodyItems {
	s.DataType = &v
	return s
}

type DescribeDataBackupsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDataBackupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDataBackupsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataBackupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDataBackupsResponse) SetHeaders(v map[string]*string) *DescribeDataBackupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDataBackupsResponse) SetStatusCode(v int32) *DescribeDataBackupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDataBackupsResponse) SetBody(v *DescribeDataBackupsResponseBody) *DescribeDataBackupsResponse {
	s.Body = v
	return s
}

type DescribeDataReDistributeInfoRequest struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeDataReDistributeInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataReDistributeInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeDataReDistributeInfoRequest) SetDBInstanceId(v string) *DescribeDataReDistributeInfoRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDataReDistributeInfoRequest) SetOwnerId(v int64) *DescribeDataReDistributeInfoRequest {
	s.OwnerId = &v
	return s
}

type DescribeDataReDistributeInfoResponseBody struct {
	DataReDistributeInfo *DescribeDataReDistributeInfoResponseBodyDataReDistributeInfo `json:"DataReDistributeInfo,omitempty" xml:"DataReDistributeInfo,omitempty" type:"Struct"`
	RequestId            *string                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDataReDistributeInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataReDistributeInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDataReDistributeInfoResponseBody) SetDataReDistributeInfo(v *DescribeDataReDistributeInfoResponseBodyDataReDistributeInfo) *DescribeDataReDistributeInfoResponseBody {
	s.DataReDistributeInfo = v
	return s
}

func (s *DescribeDataReDistributeInfoResponseBody) SetRequestId(v string) *DescribeDataReDistributeInfoResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDataReDistributeInfoResponseBodyDataReDistributeInfo struct {
	Message    *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Progress   *int64  `json:"Progress,omitempty" xml:"Progress,omitempty"`
	RemainTime *string `json:"RemainTime,omitempty" xml:"RemainTime,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Type       *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeDataReDistributeInfoResponseBodyDataReDistributeInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataReDistributeInfoResponseBodyDataReDistributeInfo) GoString() string {
	return s.String()
}

func (s *DescribeDataReDistributeInfoResponseBodyDataReDistributeInfo) SetMessage(v string) *DescribeDataReDistributeInfoResponseBodyDataReDistributeInfo {
	s.Message = &v
	return s
}

func (s *DescribeDataReDistributeInfoResponseBodyDataReDistributeInfo) SetProgress(v int64) *DescribeDataReDistributeInfoResponseBodyDataReDistributeInfo {
	s.Progress = &v
	return s
}

func (s *DescribeDataReDistributeInfoResponseBodyDataReDistributeInfo) SetRemainTime(v string) *DescribeDataReDistributeInfoResponseBodyDataReDistributeInfo {
	s.RemainTime = &v
	return s
}

func (s *DescribeDataReDistributeInfoResponseBodyDataReDistributeInfo) SetStartTime(v string) *DescribeDataReDistributeInfoResponseBodyDataReDistributeInfo {
	s.StartTime = &v
	return s
}

func (s *DescribeDataReDistributeInfoResponseBodyDataReDistributeInfo) SetStatus(v string) *DescribeDataReDistributeInfoResponseBodyDataReDistributeInfo {
	s.Status = &v
	return s
}

func (s *DescribeDataReDistributeInfoResponseBodyDataReDistributeInfo) SetType(v string) *DescribeDataReDistributeInfoResponseBodyDataReDistributeInfo {
	s.Type = &v
	return s
}

type DescribeDataReDistributeInfoResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDataReDistributeInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDataReDistributeInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataReDistributeInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeDataReDistributeInfoResponse) SetHeaders(v map[string]*string) *DescribeDataReDistributeInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeDataReDistributeInfoResponse) SetStatusCode(v int32) *DescribeDataReDistributeInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDataReDistributeInfoResponse) SetBody(v *DescribeDataReDistributeInfoResponseBody) *DescribeDataReDistributeInfoResponse {
	s.Body = v
	return s
}

type DescribeDataShareInstancesRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the page to return. The value must be an integer that is greater than 0. Default value: 1.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values:
	//
	// *   **30**
	// *   **50**
	// *   **100**
	//
	// Default value: 30.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the instance.
	//
	// >  You can call the [DescribeRegions](~~86912~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the instance belongs. For more information about how to obtain the ID of a resource group, see [View basic information of a resource group](~~151181~~).
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The keyword used to filter instances, which can be an instance ID or instance description.
	//
	// >  You can call the [DescribeDBInstances](~~86911~~) operation to query the details of all AnalyticDB for PostgreSQL instances in a specific region, including instance IDs and instance descriptions.
	SearchValue *string `json:"SearchValue,omitempty" xml:"SearchValue,omitempty"`
}

func (s DescribeDataShareInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataShareInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDataShareInstancesRequest) SetOwnerId(v int64) *DescribeDataShareInstancesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDataShareInstancesRequest) SetPageNumber(v int32) *DescribeDataShareInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDataShareInstancesRequest) SetPageSize(v int32) *DescribeDataShareInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDataShareInstancesRequest) SetRegionId(v string) *DescribeDataShareInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDataShareInstancesRequest) SetResourceGroupId(v string) *DescribeDataShareInstancesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDataShareInstancesRequest) SetSearchValue(v string) *DescribeDataShareInstancesRequest {
	s.SearchValue = &v
	return s
}

type DescribeDataShareInstancesResponseBody struct {
	// Details of the instances.
	Items *DescribeDataShareInstancesResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The page number of the returned page.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeDataShareInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataShareInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDataShareInstancesResponseBody) SetItems(v *DescribeDataShareInstancesResponseBodyItems) *DescribeDataShareInstancesResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDataShareInstancesResponseBody) SetPageNumber(v int32) *DescribeDataShareInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDataShareInstancesResponseBody) SetPageRecordCount(v int32) *DescribeDataShareInstancesResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeDataShareInstancesResponseBody) SetRequestId(v string) *DescribeDataShareInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDataShareInstancesResponseBody) SetTotalRecordCount(v int32) *DescribeDataShareInstancesResponseBody {
	s.TotalRecordCount = &v
	return s
}

type DescribeDataShareInstancesResponseBodyItems struct {
	DBInstance []*DescribeDataShareInstancesResponseBodyItemsDBInstance `json:"DBInstance,omitempty" xml:"DBInstance,omitempty" type:"Repeated"`
}

func (s DescribeDataShareInstancesResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataShareInstancesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDataShareInstancesResponseBodyItems) SetDBInstance(v []*DescribeDataShareInstancesResponseBodyItemsDBInstance) *DescribeDataShareInstancesResponseBodyItems {
	s.DBInstance = v
	return s
}

type DescribeDataShareInstancesResponseBodyItemsDBInstance struct {
	// The ID of the instance.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The resource type of the instance. Valid values:
	//
	// *   **Serverless**: Serverless mode
	// *   **StorageElasic**: elastic storage mode
	// *   **Classic**: reserved storage mode
	DBInstanceMode *string `json:"DBInstanceMode,omitempty" xml:"DBInstanceMode,omitempty"`
	// The state of data sharing. Valid values:
	//
	// *   **opening**: Data sharing is being enabled.
	// *   **opened**: Data sharing is enabled.
	// *   **closing**: Data sharing is being disabled.
	// *   **closed**: Data sharing is disabled.
	DataShareStatus *string `json:"DataShareStatus,omitempty" xml:"DataShareStatus,omitempty"`
	// The description of the instance.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The region ID of the instance.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The zone ID of the instance.
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDataShareInstancesResponseBodyItemsDBInstance) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataShareInstancesResponseBodyItemsDBInstance) GoString() string {
	return s.String()
}

func (s *DescribeDataShareInstancesResponseBodyItemsDBInstance) SetDBInstanceId(v string) *DescribeDataShareInstancesResponseBodyItemsDBInstance {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDataShareInstancesResponseBodyItemsDBInstance) SetDBInstanceMode(v string) *DescribeDataShareInstancesResponseBodyItemsDBInstance {
	s.DBInstanceMode = &v
	return s
}

func (s *DescribeDataShareInstancesResponseBodyItemsDBInstance) SetDataShareStatus(v string) *DescribeDataShareInstancesResponseBodyItemsDBInstance {
	s.DataShareStatus = &v
	return s
}

func (s *DescribeDataShareInstancesResponseBodyItemsDBInstance) SetDescription(v string) *DescribeDataShareInstancesResponseBodyItemsDBInstance {
	s.Description = &v
	return s
}

func (s *DescribeDataShareInstancesResponseBodyItemsDBInstance) SetRegionId(v string) *DescribeDataShareInstancesResponseBodyItemsDBInstance {
	s.RegionId = &v
	return s
}

func (s *DescribeDataShareInstancesResponseBodyItemsDBInstance) SetZoneId(v string) *DescribeDataShareInstancesResponseBodyItemsDBInstance {
	s.ZoneId = &v
	return s
}

type DescribeDataShareInstancesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDataShareInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDataShareInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataShareInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDataShareInstancesResponse) SetHeaders(v map[string]*string) *DescribeDataShareInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDataShareInstancesResponse) SetStatusCode(v int32) *DescribeDataShareInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDataShareInstancesResponse) SetBody(v *DescribeDataShareInstancesResponseBody) *DescribeDataShareInstancesResponse {
	s.Body = v
	return s
}

type DescribeDataSharePerformanceRequest struct {
	// The end of the time range to query. Specify the time in the *yyyy-MM-dd*T*HH:mm*Z format. The time must be in UTC.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The name of the performance metric. Separate multiple values with commas (,). Valid values:
	//
	// *   **adbpg_datashare_topic_count**: the number of shared topics.
	// *   **adbpg_datashare_data_size_mb**: the amount of data shared.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The ID of the region.
	//
	// >  You can call the [DescribeRegions](~~86912~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the instance belongs. For more information about how to obtain the ID of a resource group, see [View basic information of a resource group](~~151181~~).
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The beginning of the time range to query. Specify the time in the *yyyy-MM-dd*T*HH:mm*Z format. The time must be in UTC.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDataSharePerformanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataSharePerformanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeDataSharePerformanceRequest) SetEndTime(v string) *DescribeDataSharePerformanceRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDataSharePerformanceRequest) SetKey(v string) *DescribeDataSharePerformanceRequest {
	s.Key = &v
	return s
}

func (s *DescribeDataSharePerformanceRequest) SetRegionId(v string) *DescribeDataSharePerformanceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDataSharePerformanceRequest) SetResourceGroupId(v string) *DescribeDataSharePerformanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDataSharePerformanceRequest) SetStartTime(v string) *DescribeDataSharePerformanceRequest {
	s.StartTime = &v
	return s
}

type DescribeDataSharePerformanceResponseBody struct {
	// The ID of the instance.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The end time of the query.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Details of data sharing performance metrics.
	PerformanceKeys []*DescribeDataSharePerformanceResponseBodyPerformanceKeys `json:"PerformanceKeys,omitempty" xml:"PerformanceKeys,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The start time of the query.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDataSharePerformanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataSharePerformanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDataSharePerformanceResponseBody) SetDBClusterId(v string) *DescribeDataSharePerformanceResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDataSharePerformanceResponseBody) SetEndTime(v string) *DescribeDataSharePerformanceResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDataSharePerformanceResponseBody) SetPerformanceKeys(v []*DescribeDataSharePerformanceResponseBodyPerformanceKeys) *DescribeDataSharePerformanceResponseBody {
	s.PerformanceKeys = v
	return s
}

func (s *DescribeDataSharePerformanceResponseBody) SetRequestId(v string) *DescribeDataSharePerformanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDataSharePerformanceResponseBody) SetStartTime(v string) *DescribeDataSharePerformanceResponseBody {
	s.StartTime = &v
	return s
}

type DescribeDataSharePerformanceResponseBodyPerformanceKeys struct {
	// The name of the performance metric.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Details of the performance metric.
	Series []*DescribeDataSharePerformanceResponseBodyPerformanceKeysSeries `json:"Series,omitempty" xml:"Series,omitempty" type:"Repeated"`
	// The unit of the performance metric.
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
}

func (s DescribeDataSharePerformanceResponseBodyPerformanceKeys) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataSharePerformanceResponseBodyPerformanceKeys) GoString() string {
	return s.String()
}

func (s *DescribeDataSharePerformanceResponseBodyPerformanceKeys) SetName(v string) *DescribeDataSharePerformanceResponseBodyPerformanceKeys {
	s.Name = &v
	return s
}

func (s *DescribeDataSharePerformanceResponseBodyPerformanceKeys) SetSeries(v []*DescribeDataSharePerformanceResponseBodyPerformanceKeysSeries) *DescribeDataSharePerformanceResponseBodyPerformanceKeys {
	s.Series = v
	return s
}

func (s *DescribeDataSharePerformanceResponseBodyPerformanceKeys) SetUnit(v string) *DescribeDataSharePerformanceResponseBodyPerformanceKeys {
	s.Unit = &v
	return s
}

type DescribeDataSharePerformanceResponseBodyPerformanceKeysSeries struct {
	// The name of the performance metric.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// One or more values of the performance metric.
	Values []*DescribeDataSharePerformanceResponseBodyPerformanceKeysSeriesValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s DescribeDataSharePerformanceResponseBodyPerformanceKeysSeries) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataSharePerformanceResponseBodyPerformanceKeysSeries) GoString() string {
	return s.String()
}

func (s *DescribeDataSharePerformanceResponseBodyPerformanceKeysSeries) SetName(v string) *DescribeDataSharePerformanceResponseBodyPerformanceKeysSeries {
	s.Name = &v
	return s
}

func (s *DescribeDataSharePerformanceResponseBodyPerformanceKeysSeries) SetValues(v []*DescribeDataSharePerformanceResponseBodyPerformanceKeysSeriesValues) *DescribeDataSharePerformanceResponseBodyPerformanceKeysSeries {
	s.Values = v
	return s
}

type DescribeDataSharePerformanceResponseBodyPerformanceKeysSeriesValues struct {
	// The value of the performance metric at a point in time.
	Point []*string `json:"Point,omitempty" xml:"Point,omitempty" type:"Repeated"`
}

func (s DescribeDataSharePerformanceResponseBodyPerformanceKeysSeriesValues) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataSharePerformanceResponseBodyPerformanceKeysSeriesValues) GoString() string {
	return s.String()
}

func (s *DescribeDataSharePerformanceResponseBodyPerformanceKeysSeriesValues) SetPoint(v []*string) *DescribeDataSharePerformanceResponseBodyPerformanceKeysSeriesValues {
	s.Point = v
	return s
}

type DescribeDataSharePerformanceResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDataSharePerformanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDataSharePerformanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataSharePerformanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeDataSharePerformanceResponse) SetHeaders(v map[string]*string) *DescribeDataSharePerformanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeDataSharePerformanceResponse) SetStatusCode(v int32) *DescribeDataSharePerformanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDataSharePerformanceResponse) SetBody(v *DescribeDataSharePerformanceResponseBody) *DescribeDataSharePerformanceResponse {
	s.Body = v
	return s
}

type DescribeDiagnosisDimensionsRequest struct {
	// The ID of the instance.
	//
	// >  You can call the [DescribeDBInstances](~~86911~~) operation to query the details of all AnalyticDB for PostgreSQL instances in a specific region, including instance IDs.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s DescribeDiagnosisDimensionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiagnosisDimensionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisDimensionsRequest) SetDBInstanceId(v string) *DescribeDiagnosisDimensionsRequest {
	s.DBInstanceId = &v
	return s
}

type DescribeDiagnosisDimensionsResponseBody struct {
	// The name of the database.
	Databases []*string `json:"Databases,omitempty" xml:"Databases,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The name of the database account.
	UserNames []*string `json:"UserNames,omitempty" xml:"UserNames,omitempty" type:"Repeated"`
}

func (s DescribeDiagnosisDimensionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiagnosisDimensionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisDimensionsResponseBody) SetDatabases(v []*string) *DescribeDiagnosisDimensionsResponseBody {
	s.Databases = v
	return s
}

func (s *DescribeDiagnosisDimensionsResponseBody) SetRequestId(v string) *DescribeDiagnosisDimensionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDiagnosisDimensionsResponseBody) SetUserNames(v []*string) *DescribeDiagnosisDimensionsResponseBody {
	s.UserNames = v
	return s
}

type DescribeDiagnosisDimensionsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDiagnosisDimensionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDiagnosisDimensionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiagnosisDimensionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisDimensionsResponse) SetHeaders(v map[string]*string) *DescribeDiagnosisDimensionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDiagnosisDimensionsResponse) SetStatusCode(v int32) *DescribeDiagnosisDimensionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDiagnosisDimensionsResponse) SetBody(v *DescribeDiagnosisDimensionsResponseBody) *DescribeDiagnosisDimensionsResponse {
	s.Body = v
	return s
}

type DescribeDiagnosisMonitorPerformanceRequest struct {
	// The ID of the instance.
	//
	// >  You can call the [DescribeDBInstances](~~86911~~) operation to query the details of all AnalyticDB for PostgreSQL instances in a specific region, including instance IDs.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The name of the database.
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	// The end of the time range to query. Specify the time in the *yyyy-MM-dd*T*HH:mm*Z format. The time must be in UTC. The end time must be later than the start time.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The filter condition on queries. Specify the value in the JSON format. Valid values:
	//
	// *   `{"Type":"maxCost", "Value":"100"}`: filters the top 100 queries that are the most time-consuming.
	//
	// *   `{"Type":"status","Value":"finished"}`: filters completed queries.
	//
	// *   `{"Type":"status","Value":"running"}`: filters running queries.
	//
	// *   `{"Type":"cost","Min":"30","Max":"50"}`: filters the queries that consume 30 milliseconds or more and less than 50 milliseconds. You can customize a filter condition by setting **Min** and **Max**.
	//
	//     *   If only **Min** is specified, the queries that consume a period of time that is greater than or equal to the Min value are filtered.
	//     *   If only **Max** is specified, the queries that consume a period of time that is less than the Max value are filtered.
	//     *   If both **Min** and **Max** are specified, the queries that consume a period of time that is greater than or equal to the **Min** value and less than the **Max** value are filtered.
	QueryCondition *string `json:"QueryCondition,omitempty" xml:"QueryCondition,omitempty"`
	// The beginning of the time range to query. Specify the time in the *yyyy-MM-dd*T*HH:mm*Z format. The time must be in UTC.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The name of the database account.
	User *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DescribeDiagnosisMonitorPerformanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiagnosisMonitorPerformanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisMonitorPerformanceRequest) SetDBInstanceId(v string) *DescribeDiagnosisMonitorPerformanceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDiagnosisMonitorPerformanceRequest) SetDatabase(v string) *DescribeDiagnosisMonitorPerformanceRequest {
	s.Database = &v
	return s
}

func (s *DescribeDiagnosisMonitorPerformanceRequest) SetEndTime(v string) *DescribeDiagnosisMonitorPerformanceRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDiagnosisMonitorPerformanceRequest) SetQueryCondition(v string) *DescribeDiagnosisMonitorPerformanceRequest {
	s.QueryCondition = &v
	return s
}

func (s *DescribeDiagnosisMonitorPerformanceRequest) SetStartTime(v string) *DescribeDiagnosisMonitorPerformanceRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDiagnosisMonitorPerformanceRequest) SetUser(v string) *DescribeDiagnosisMonitorPerformanceRequest {
	s.User = &v
	return s
}

type DescribeDiagnosisMonitorPerformanceResponseBody struct {
	// Details of query execution.
	Performances []*DescribeDiagnosisMonitorPerformanceResponseBodyPerformances `json:"Performances,omitempty" xml:"Performances,omitempty" type:"Repeated"`
	// The threshold for the number of queries.
	PerformancesThreshold *int32 `json:"PerformancesThreshold,omitempty" xml:"PerformancesThreshold,omitempty"`
	// Indicates whether the queries are truncated when the number of queries exceeds the threshold. Valid values:
	//
	// *   **true**: The queries are truncated.
	// *   **false**: The queries are not truncated.
	PerformancesTruncated *bool `json:"PerformancesTruncated,omitempty" xml:"PerformancesTruncated,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDiagnosisMonitorPerformanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiagnosisMonitorPerformanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisMonitorPerformanceResponseBody) SetPerformances(v []*DescribeDiagnosisMonitorPerformanceResponseBodyPerformances) *DescribeDiagnosisMonitorPerformanceResponseBody {
	s.Performances = v
	return s
}

func (s *DescribeDiagnosisMonitorPerformanceResponseBody) SetPerformancesThreshold(v int32) *DescribeDiagnosisMonitorPerformanceResponseBody {
	s.PerformancesThreshold = &v
	return s
}

func (s *DescribeDiagnosisMonitorPerformanceResponseBody) SetPerformancesTruncated(v bool) *DescribeDiagnosisMonitorPerformanceResponseBody {
	s.PerformancesTruncated = &v
	return s
}

func (s *DescribeDiagnosisMonitorPerformanceResponseBody) SetRequestId(v string) *DescribeDiagnosisMonitorPerformanceResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDiagnosisMonitorPerformanceResponseBodyPerformances struct {
	// The execution duration of the query. Unit: milliseconds.
	Cost *int32 `json:"Cost,omitempty" xml:"Cost,omitempty"`
	// The name of the database.
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	// The ID of the query. It is a unique identifier of the query.
	QueryID *string `json:"QueryID,omitempty" xml:"QueryID,omitempty"`
	// The start time of the query. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The execution state of the query. Valid values:
	//
	// *   **running**: The query is being executed.
	// *   **finished**: The query is complete.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The name of the database account.
	User *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DescribeDiagnosisMonitorPerformanceResponseBodyPerformances) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiagnosisMonitorPerformanceResponseBodyPerformances) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisMonitorPerformanceResponseBodyPerformances) SetCost(v int32) *DescribeDiagnosisMonitorPerformanceResponseBodyPerformances {
	s.Cost = &v
	return s
}

func (s *DescribeDiagnosisMonitorPerformanceResponseBodyPerformances) SetDatabase(v string) *DescribeDiagnosisMonitorPerformanceResponseBodyPerformances {
	s.Database = &v
	return s
}

func (s *DescribeDiagnosisMonitorPerformanceResponseBodyPerformances) SetQueryID(v string) *DescribeDiagnosisMonitorPerformanceResponseBodyPerformances {
	s.QueryID = &v
	return s
}

func (s *DescribeDiagnosisMonitorPerformanceResponseBodyPerformances) SetStartTime(v int64) *DescribeDiagnosisMonitorPerformanceResponseBodyPerformances {
	s.StartTime = &v
	return s
}

func (s *DescribeDiagnosisMonitorPerformanceResponseBodyPerformances) SetStatus(v string) *DescribeDiagnosisMonitorPerformanceResponseBodyPerformances {
	s.Status = &v
	return s
}

func (s *DescribeDiagnosisMonitorPerformanceResponseBodyPerformances) SetUser(v string) *DescribeDiagnosisMonitorPerformanceResponseBodyPerformances {
	s.User = &v
	return s
}

type DescribeDiagnosisMonitorPerformanceResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDiagnosisMonitorPerformanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDiagnosisMonitorPerformanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiagnosisMonitorPerformanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisMonitorPerformanceResponse) SetHeaders(v map[string]*string) *DescribeDiagnosisMonitorPerformanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeDiagnosisMonitorPerformanceResponse) SetStatusCode(v int32) *DescribeDiagnosisMonitorPerformanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDiagnosisMonitorPerformanceResponse) SetBody(v *DescribeDiagnosisMonitorPerformanceResponseBody) *DescribeDiagnosisMonitorPerformanceResponse {
	s.Body = v
	return s
}

type DescribeDiagnosisRecordsRequest struct {
	// The ID of the instance.
	//
	// >  You can call the [DescribeDBInstances](~~86911~~) operation to query the details of all AnalyticDB for PostgreSQL instances in a specific region, including instance IDs.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The name of the database.
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	// The end of the time range to query. Specify the time in the *yyyy-MM-dd*T*HH:mm*Z format. The time must be in UTC. The end time must be later than the start time.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The keyword of the SQL statement.
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The order of fields in the console. You do not need to specify this parameter.
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The number of the page to return. The value must be an integer that is greater than 0. Default value: **1**.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values:
	//
	// *   **30**
	// *   **50**
	// *   **100**
	//
	// Default value: **30**.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The filter condition on queries. Specify the value in the JSON format. Valid values:
	//
	// *   `{"Type":"maxCost", "Value":"100"}`: filters the top 100 queries that are the most time-consuming.
	//
	// *   `{"Type":"status","Value":"finished"}`: filters completed queries.
	//
	// *   `{"Type":"status","Value":"running"}`: filters running queries.
	//
	// *   `{"Type":"cost","Min":"30","Max":"50"}`: filters the queries that consume 30 milliseconds or more and less than 50 milliseconds. You can customize a filter condition by setting **Min** and **Max**.
	//
	//     *   If only **Min** is specified, the queries that consume a period of time that is greater than or equal to the Min value are filtered.
	//     *   If only **Max** is specified, the queries that consume a period of time that is less than the Max value are filtered.
	//     *   If both **Min** and **Max** are specified, the queries that consume a period of time that is greater than or equal to the **Min** value and less than the **Max** value are filtered.
	QueryCondition *string `json:"QueryCondition,omitempty" xml:"QueryCondition,omitempty"`
	// The beginning of the time range to query. Specify the time in the *yyyy-MM-dd*T*HH:mm*Z format. The time must be in UTC.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The name of the database account.
	User *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DescribeDiagnosisRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiagnosisRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisRecordsRequest) SetDBInstanceId(v string) *DescribeDiagnosisRecordsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDiagnosisRecordsRequest) SetDatabase(v string) *DescribeDiagnosisRecordsRequest {
	s.Database = &v
	return s
}

func (s *DescribeDiagnosisRecordsRequest) SetEndTime(v string) *DescribeDiagnosisRecordsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDiagnosisRecordsRequest) SetKeyword(v string) *DescribeDiagnosisRecordsRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeDiagnosisRecordsRequest) SetOrder(v string) *DescribeDiagnosisRecordsRequest {
	s.Order = &v
	return s
}

func (s *DescribeDiagnosisRecordsRequest) SetPageNumber(v int32) *DescribeDiagnosisRecordsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDiagnosisRecordsRequest) SetPageSize(v int32) *DescribeDiagnosisRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDiagnosisRecordsRequest) SetQueryCondition(v string) *DescribeDiagnosisRecordsRequest {
	s.QueryCondition = &v
	return s
}

func (s *DescribeDiagnosisRecordsRequest) SetStartTime(v string) *DescribeDiagnosisRecordsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDiagnosisRecordsRequest) SetUser(v string) *DescribeDiagnosisRecordsRequest {
	s.User = &v
	return s
}

type DescribeDiagnosisRecordsResponseBody struct {
	// Details of SQL queries.
	Items []*DescribeDiagnosisRecordsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The page number of the returned page.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDiagnosisRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiagnosisRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisRecordsResponseBody) SetItems(v []*DescribeDiagnosisRecordsResponseBodyItems) *DescribeDiagnosisRecordsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBody) SetPageNumber(v int32) *DescribeDiagnosisRecordsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBody) SetRequestId(v string) *DescribeDiagnosisRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBody) SetTotalCount(v int32) *DescribeDiagnosisRecordsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeDiagnosisRecordsResponseBodyItems struct {
	// The name of the database.
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	// The execution duration of the query. Unit: seconds.
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The ID of the query. It is a unique identifier of the query.
	QueryID *string `json:"QueryID,omitempty" xml:"QueryID,omitempty"`
	// The SQL statement.
	SQLStmt *string `json:"SQLStmt,omitempty" xml:"SQLStmt,omitempty"`
	// Indicates whether the SQL statement needs to be truncated. Valid values:
	//
	// *   **true**: The SQL statement needs to be truncated.
	// *   **false**: The SQL statement does not need to be truncated.
	SQLTruncated *bool `json:"SQLTruncated,omitempty" xml:"SQLTruncated,omitempty"`
	// The threshold used to determine whether an SQL statement must be truncated. The value is the number of characters.
	SQLTruncatedThreshold *int32 `json:"SQLTruncatedThreshold,omitempty" xml:"SQLTruncatedThreshold,omitempty"`
	// The ID of the session that contains the query.
	SessionID *string `json:"SessionID,omitempty" xml:"SessionID,omitempty"`
	// The start time of the query. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The execution state of the query. Valid values:
	//
	// *   **running**: The query is being executed.
	// *   **finished**: The query is complete.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The name of the database account.
	User *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DescribeDiagnosisRecordsResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiagnosisRecordsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisRecordsResponseBodyItems) SetDatabase(v string) *DescribeDiagnosisRecordsResponseBodyItems {
	s.Database = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyItems) SetDuration(v int32) *DescribeDiagnosisRecordsResponseBodyItems {
	s.Duration = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyItems) SetQueryID(v string) *DescribeDiagnosisRecordsResponseBodyItems {
	s.QueryID = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyItems) SetSQLStmt(v string) *DescribeDiagnosisRecordsResponseBodyItems {
	s.SQLStmt = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyItems) SetSQLTruncated(v bool) *DescribeDiagnosisRecordsResponseBodyItems {
	s.SQLTruncated = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyItems) SetSQLTruncatedThreshold(v int32) *DescribeDiagnosisRecordsResponseBodyItems {
	s.SQLTruncatedThreshold = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyItems) SetSessionID(v string) *DescribeDiagnosisRecordsResponseBodyItems {
	s.SessionID = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyItems) SetStartTime(v int64) *DescribeDiagnosisRecordsResponseBodyItems {
	s.StartTime = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyItems) SetStatus(v string) *DescribeDiagnosisRecordsResponseBodyItems {
	s.Status = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyItems) SetUser(v string) *DescribeDiagnosisRecordsResponseBodyItems {
	s.User = &v
	return s
}

type DescribeDiagnosisRecordsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDiagnosisRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDiagnosisRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiagnosisRecordsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisRecordsResponse) SetHeaders(v map[string]*string) *DescribeDiagnosisRecordsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDiagnosisRecordsResponse) SetStatusCode(v int32) *DescribeDiagnosisRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponse) SetBody(v *DescribeDiagnosisRecordsResponseBody) *DescribeDiagnosisRecordsResponse {
	s.Body = v
	return s
}

type DescribeDiagnosisSQLInfoRequest struct {
	// The ID of the instance.
	//
	// >  You can call the [DescribeDBInstances](~~86911~~) operation to query the details of all AnalyticDB for PostgreSQL instances in a specific region, including instance IDs.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The name of the database.
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	// The ID of the query. It is a unique identifier of the query.
	//
	// >  You can call the [DescribeDiagnosisRecords](~~450511~~) operation to query the query ID.
	QueryID *string `json:"QueryID,omitempty" xml:"QueryID,omitempty"`
}

func (s DescribeDiagnosisSQLInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiagnosisSQLInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisSQLInfoRequest) SetDBInstanceId(v string) *DescribeDiagnosisSQLInfoRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoRequest) SetDatabase(v string) *DescribeDiagnosisSQLInfoRequest {
	s.Database = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoRequest) SetQueryID(v string) *DescribeDiagnosisSQLInfoRequest {
	s.QueryID = &v
	return s
}

type DescribeDiagnosisSQLInfoResponseBody struct {
	// The name of the database.
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	// The execution duration of the query. Unit: seconds.
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The maximum number of output rows.
	MaxOutputRows *string `json:"MaxOutputRows,omitempty" xml:"MaxOutputRows,omitempty"`
	// The ID of the query.
	QueryID *string `json:"QueryID,omitempty" xml:"QueryID,omitempty"`
	// The information of the operator.
	QueryPlan *string `json:"QueryPlan,omitempty" xml:"QueryPlan,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The SQL statement.
	SQLStmt *string `json:"SQLStmt,omitempty" xml:"SQLStmt,omitempty"`
	// The ID of the session that contains the query.
	SessionID *string `json:"SessionID,omitempty" xml:"SessionID,omitempty"`
	// The sequence of metrics.
	SortedMetrics *string `json:"SortedMetrics,omitempty" xml:"SortedMetrics,omitempty"`
	// The start time of the query. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The execution state of the query. Valid values:
	//
	// *   **running**: The query is being executed.
	// *   **finished**: The query execution is complete.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The information of the execution plan.
	TextPlan *string `json:"TextPlan,omitempty" xml:"TextPlan,omitempty"`
	// The name of the database account.
	User *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DescribeDiagnosisSQLInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiagnosisSQLInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisSQLInfoResponseBody) SetDatabase(v string) *DescribeDiagnosisSQLInfoResponseBody {
	s.Database = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoResponseBody) SetDuration(v int32) *DescribeDiagnosisSQLInfoResponseBody {
	s.Duration = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoResponseBody) SetMaxOutputRows(v string) *DescribeDiagnosisSQLInfoResponseBody {
	s.MaxOutputRows = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoResponseBody) SetQueryID(v string) *DescribeDiagnosisSQLInfoResponseBody {
	s.QueryID = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoResponseBody) SetQueryPlan(v string) *DescribeDiagnosisSQLInfoResponseBody {
	s.QueryPlan = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoResponseBody) SetRequestId(v string) *DescribeDiagnosisSQLInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoResponseBody) SetSQLStmt(v string) *DescribeDiagnosisSQLInfoResponseBody {
	s.SQLStmt = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoResponseBody) SetSessionID(v string) *DescribeDiagnosisSQLInfoResponseBody {
	s.SessionID = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoResponseBody) SetSortedMetrics(v string) *DescribeDiagnosisSQLInfoResponseBody {
	s.SortedMetrics = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoResponseBody) SetStartTime(v int64) *DescribeDiagnosisSQLInfoResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoResponseBody) SetStatus(v string) *DescribeDiagnosisSQLInfoResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoResponseBody) SetTextPlan(v string) *DescribeDiagnosisSQLInfoResponseBody {
	s.TextPlan = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoResponseBody) SetUser(v string) *DescribeDiagnosisSQLInfoResponseBody {
	s.User = &v
	return s
}

type DescribeDiagnosisSQLInfoResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDiagnosisSQLInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDiagnosisSQLInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiagnosisSQLInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisSQLInfoResponse) SetHeaders(v map[string]*string) *DescribeDiagnosisSQLInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeDiagnosisSQLInfoResponse) SetStatusCode(v int32) *DescribeDiagnosisSQLInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoResponse) SetBody(v *DescribeDiagnosisSQLInfoResponseBody) *DescribeDiagnosisSQLInfoResponse {
	s.Body = v
	return s
}

type DescribeDownloadRecordsRequest struct {
	// The ID of the instance.
	//
	// >  You can call the [DescribeDBInstances](~~86911~~) operation to query the details of all AnalyticDB for PostgreSQL instances in a specific region, including instance IDs.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s DescribeDownloadRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDownloadRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDownloadRecordsRequest) SetDBInstanceId(v string) *DescribeDownloadRecordsRequest {
	s.DBInstanceId = &v
	return s
}

type DescribeDownloadRecordsResponseBody struct {
	// Details of the download records.
	Records []*DescribeDownloadRecordsResponseBodyRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDownloadRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDownloadRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDownloadRecordsResponseBody) SetRecords(v []*DescribeDownloadRecordsResponseBodyRecords) *DescribeDownloadRecordsResponseBody {
	s.Records = v
	return s
}

func (s *DescribeDownloadRecordsResponseBody) SetRequestId(v string) *DescribeDownloadRecordsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDownloadRecordsResponseBodyRecords struct {
	// The ID of the download record.
	DownloadId *int64 `json:"DownloadId,omitempty" xml:"DownloadId,omitempty"`
	// The URL that can be used to download the file.
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	// The error message returned.
	ExceptionMsg *string `json:"ExceptionMsg,omitempty" xml:"ExceptionMsg,omitempty"`
	// The name of the file.
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The state of the upload task. After you call the DownloadDiagnosisRecords operation, query diagnostic information is first uploaded to Object Storage Service (OSS). After the upload task is complete, the query diagnostic information can be downloaded. Valid values:
	//
	// *   **running**: uploading
	// *   **finished**: uploaded
	// *   **failed**: failed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDownloadRecordsResponseBodyRecords) String() string {
	return tea.Prettify(s)
}

func (s DescribeDownloadRecordsResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *DescribeDownloadRecordsResponseBodyRecords) SetDownloadId(v int64) *DescribeDownloadRecordsResponseBodyRecords {
	s.DownloadId = &v
	return s
}

func (s *DescribeDownloadRecordsResponseBodyRecords) SetDownloadUrl(v string) *DescribeDownloadRecordsResponseBodyRecords {
	s.DownloadUrl = &v
	return s
}

func (s *DescribeDownloadRecordsResponseBodyRecords) SetExceptionMsg(v string) *DescribeDownloadRecordsResponseBodyRecords {
	s.ExceptionMsg = &v
	return s
}

func (s *DescribeDownloadRecordsResponseBodyRecords) SetFileName(v string) *DescribeDownloadRecordsResponseBodyRecords {
	s.FileName = &v
	return s
}

func (s *DescribeDownloadRecordsResponseBodyRecords) SetStatus(v string) *DescribeDownloadRecordsResponseBodyRecords {
	s.Status = &v
	return s
}

type DescribeDownloadRecordsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDownloadRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDownloadRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDownloadRecordsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDownloadRecordsResponse) SetHeaders(v map[string]*string) *DescribeDownloadRecordsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDownloadRecordsResponse) SetStatusCode(v int32) *DescribeDownloadRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDownloadRecordsResponse) SetBody(v *DescribeDownloadRecordsResponseBody) *DescribeDownloadRecordsResponse {
	s.Body = v
	return s
}

type DescribeDownloadSQLLogsRequest struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s DescribeDownloadSQLLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDownloadSQLLogsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDownloadSQLLogsRequest) SetDBInstanceId(v string) *DescribeDownloadSQLLogsRequest {
	s.DBInstanceId = &v
	return s
}

type DescribeDownloadSQLLogsResponseBody struct {
	Records   []*DescribeDownloadSQLLogsResponseBodyRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDownloadSQLLogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDownloadSQLLogsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDownloadSQLLogsResponseBody) SetRecords(v []*DescribeDownloadSQLLogsResponseBodyRecords) *DescribeDownloadSQLLogsResponseBody {
	s.Records = v
	return s
}

func (s *DescribeDownloadSQLLogsResponseBody) SetRequestId(v string) *DescribeDownloadSQLLogsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDownloadSQLLogsResponseBodyRecords struct {
	DownloadId   *int64  `json:"DownloadId,omitempty" xml:"DownloadId,omitempty"`
	DownloadUrl  *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	ExceptionMsg *string `json:"ExceptionMsg,omitempty" xml:"ExceptionMsg,omitempty"`
	FileName     *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDownloadSQLLogsResponseBodyRecords) String() string {
	return tea.Prettify(s)
}

func (s DescribeDownloadSQLLogsResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *DescribeDownloadSQLLogsResponseBodyRecords) SetDownloadId(v int64) *DescribeDownloadSQLLogsResponseBodyRecords {
	s.DownloadId = &v
	return s
}

func (s *DescribeDownloadSQLLogsResponseBodyRecords) SetDownloadUrl(v string) *DescribeDownloadSQLLogsResponseBodyRecords {
	s.DownloadUrl = &v
	return s
}

func (s *DescribeDownloadSQLLogsResponseBodyRecords) SetExceptionMsg(v string) *DescribeDownloadSQLLogsResponseBodyRecords {
	s.ExceptionMsg = &v
	return s
}

func (s *DescribeDownloadSQLLogsResponseBodyRecords) SetFileName(v string) *DescribeDownloadSQLLogsResponseBodyRecords {
	s.FileName = &v
	return s
}

func (s *DescribeDownloadSQLLogsResponseBodyRecords) SetStatus(v string) *DescribeDownloadSQLLogsResponseBodyRecords {
	s.Status = &v
	return s
}

type DescribeDownloadSQLLogsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDownloadSQLLogsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDownloadSQLLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDownloadSQLLogsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDownloadSQLLogsResponse) SetHeaders(v map[string]*string) *DescribeDownloadSQLLogsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDownloadSQLLogsResponse) SetStatusCode(v int32) *DescribeDownloadSQLLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDownloadSQLLogsResponse) SetBody(v *DescribeDownloadSQLLogsResponseBody) *DescribeDownloadSQLLogsResponse {
	s.Body = v
	return s
}

type DescribeHealthStatusRequest struct {
	// The ID of the instance.
	//
	// >  You can call the [DescribeDBInstances](~~86911~~) operation to query the details of all AnalyticDB for PostgreSQL instances in a specific region, including instance IDs.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The performance metric that you want to query. Separate multiple values with commas (,). For more information, see [Performance parameters](~~86943~~).
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
}

func (s DescribeHealthStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeHealthStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeHealthStatusRequest) SetDBInstanceId(v string) *DescribeHealthStatusRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeHealthStatusRequest) SetKey(v string) *DescribeHealthStatusRequest {
	s.Key = &v
	return s
}

type DescribeHealthStatusResponseBody struct {
	// The ID of instance.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information of performance metrics. Each performance metric consists of the parameter name, status, and metric value. The metric information is returned only for the performance parameters specified by **Key**. For example, if you set **Key** to **adbpg_status**, only the metric information of **adbpg_status** is returned.
	//
	// For more information about performance parameters, see [Performance parameters](~~86943~~).
	Status *DescribeHealthStatusResponseBodyStatus `json:"Status,omitempty" xml:"Status,omitempty" type:"Struct"`
}

func (s DescribeHealthStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeHealthStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHealthStatusResponseBody) SetDBClusterId(v string) *DescribeHealthStatusResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeHealthStatusResponseBody) SetRequestId(v string) *DescribeHealthStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHealthStatusResponseBody) SetStatus(v *DescribeHealthStatusResponseBodyStatus) *DescribeHealthStatusResponseBody {
	s.Status = v
	return s
}

type DescribeHealthStatusResponseBodyStatus struct {
	// The information of maximum compute node storage usage.
	//
	// >  This parameter value is returned only for instances in elastic storage mode.
	AdbgpSegmentDiskUsagePercentMax *DescribeHealthStatusResponseBodyStatusAdbgpSegmentDiskUsagePercentMax `json:"adbgp_segment_disk_usage_percent_max,omitempty" xml:"adbgp_segment_disk_usage_percent_max,omitempty" type:"Struct"`
	// The information of instance connection health status.
	AdbpgConnectionStatus *DescribeHealthStatusResponseBodyStatusAdbpgConnectionStatus `json:"adbpg_connection_status,omitempty" xml:"adbpg_connection_status,omitempty" type:"Struct"`
	// The information of instance storage status.
	//
	// >  This parameter value is returned only for instances in elastic storage mode.
	AdbpgDiskStatus *DescribeHealthStatusResponseBodyStatusAdbpgDiskStatus `json:"adbpg_disk_status,omitempty" xml:"adbpg_disk_status,omitempty" type:"Struct"`
	// The information of instance storage usage.
	//
	// >  This parameter value is returned only for instances in elastic storage mode.
	AdbpgDiskUsagePercent    *DescribeHealthStatusResponseBodyStatusAdbpgDiskUsagePercent    `json:"adbpg_disk_usage_percent,omitempty" xml:"adbpg_disk_usage_percent,omitempty" type:"Struct"`
	AdbpgInstanceColdDataGb  *DescribeHealthStatusResponseBodyStatusAdbpgInstanceColdDataGb  `json:"adbpg_instance_cold_data_gb,omitempty" xml:"adbpg_instance_cold_data_gb,omitempty" type:"Struct"`
	AdbpgInstanceHotDataGb   *DescribeHealthStatusResponseBodyStatusAdbpgInstanceHotDataGb   `json:"adbpg_instance_hot_data_gb,omitempty" xml:"adbpg_instance_hot_data_gb,omitempty" type:"Struct"`
	AdbpgInstanceTotalDataGb *DescribeHealthStatusResponseBodyStatusAdbpgInstanceTotalDataGb `json:"adbpg_instance_total_data_gb,omitempty" xml:"adbpg_instance_total_data_gb,omitempty" type:"Struct"`
	// The information of maximum coordinator node storage usage.
	//
	// >  This parameter value is returned only for instances in elastic storage mode.
	AdbpgMasterDiskUsagePercentMax *DescribeHealthStatusResponseBodyStatusAdbpgMasterDiskUsagePercentMax `json:"adbpg_master_disk_usage_percent_max,omitempty" xml:"adbpg_master_disk_usage_percent_max,omitempty" type:"Struct"`
	// The information of coordinator node availability status.
	AdbpgMasterStatus *DescribeHealthStatusResponseBodyStatusAdbpgMasterStatus `json:"adbpg_master_status,omitempty" xml:"adbpg_master_status,omitempty" type:"Struct"`
	// The information of compute node availability status.
	AdbpgSegmentStatus *DescribeHealthStatusResponseBodyStatusAdbpgSegmentStatus `json:"adbpg_segment_status,omitempty" xml:"adbpg_segment_status,omitempty" type:"Struct"`
	// The information of instance health status.
	AdbpgStatus *DescribeHealthStatusResponseBodyStatusAdbpgStatus `json:"adbpg_status,omitempty" xml:"adbpg_status,omitempty" type:"Struct"`
	// The information of coordinator node connection health status.
	NodeMasterConnectionStatus *DescribeHealthStatusResponseBodyStatusNodeMasterConnectionStatus `json:"node_master_connection_status,omitempty" xml:"node_master_connection_status,omitempty" type:"Struct"`
	// The information of coordinator node health status.
	NodeMasterStatus *DescribeHealthStatusResponseBodyStatusNodeMasterStatus `json:"node_master_status,omitempty" xml:"node_master_status,omitempty" type:"Struct"`
	// The information of compute node connection health status.
	NodeSegmentConnectionStatus *DescribeHealthStatusResponseBodyStatusNodeSegmentConnectionStatus `json:"node_segment_connection_status,omitempty" xml:"node_segment_connection_status,omitempty" type:"Struct"`
	// The information of compute node storage status.
	//
	// >  This parameter value is returned only for instances in elastic storage mode.
	NodeSegmentDiskStatus *DescribeHealthStatusResponseBodyStatusNodeSegmentDiskStatus `json:"node_segment_disk_status,omitempty" xml:"node_segment_disk_status,omitempty" type:"Struct"`
}

func (s DescribeHealthStatusResponseBodyStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeHealthStatusResponseBodyStatus) GoString() string {
	return s.String()
}

func (s *DescribeHealthStatusResponseBodyStatus) SetAdbgpSegmentDiskUsagePercentMax(v *DescribeHealthStatusResponseBodyStatusAdbgpSegmentDiskUsagePercentMax) *DescribeHealthStatusResponseBodyStatus {
	s.AdbgpSegmentDiskUsagePercentMax = v
	return s
}

func (s *DescribeHealthStatusResponseBodyStatus) SetAdbpgConnectionStatus(v *DescribeHealthStatusResponseBodyStatusAdbpgConnectionStatus) *DescribeHealthStatusResponseBodyStatus {
	s.AdbpgConnectionStatus = v
	return s
}

func (s *DescribeHealthStatusResponseBodyStatus) SetAdbpgDiskStatus(v *DescribeHealthStatusResponseBodyStatusAdbpgDiskStatus) *DescribeHealthStatusResponseBodyStatus {
	s.AdbpgDiskStatus = v
	return s
}

func (s *DescribeHealthStatusResponseBodyStatus) SetAdbpgDiskUsagePercent(v *DescribeHealthStatusResponseBodyStatusAdbpgDiskUsagePercent) *DescribeHealthStatusResponseBodyStatus {
	s.AdbpgDiskUsagePercent = v
	return s
}

func (s *DescribeHealthStatusResponseBodyStatus) SetAdbpgInstanceColdDataGb(v *DescribeHealthStatusResponseBodyStatusAdbpgInstanceColdDataGb) *DescribeHealthStatusResponseBodyStatus {
	s.AdbpgInstanceColdDataGb = v
	return s
}

func (s *DescribeHealthStatusResponseBodyStatus) SetAdbpgInstanceHotDataGb(v *DescribeHealthStatusResponseBodyStatusAdbpgInstanceHotDataGb) *DescribeHealthStatusResponseBodyStatus {
	s.AdbpgInstanceHotDataGb = v
	return s
}

func (s *DescribeHealthStatusResponseBodyStatus) SetAdbpgInstanceTotalDataGb(v *DescribeHealthStatusResponseBodyStatusAdbpgInstanceTotalDataGb) *DescribeHealthStatusResponseBodyStatus {
	s.AdbpgInstanceTotalDataGb = v
	return s
}

func (s *DescribeHealthStatusResponseBodyStatus) SetAdbpgMasterDiskUsagePercentMax(v *DescribeHealthStatusResponseBodyStatusAdbpgMasterDiskUsagePercentMax) *DescribeHealthStatusResponseBodyStatus {
	s.AdbpgMasterDiskUsagePercentMax = v
	return s
}

func (s *DescribeHealthStatusResponseBodyStatus) SetAdbpgMasterStatus(v *DescribeHealthStatusResponseBodyStatusAdbpgMasterStatus) *DescribeHealthStatusResponseBodyStatus {
	s.AdbpgMasterStatus = v
	return s
}

func (s *DescribeHealthStatusResponseBodyStatus) SetAdbpgSegmentStatus(v *DescribeHealthStatusResponseBodyStatusAdbpgSegmentStatus) *DescribeHealthStatusResponseBodyStatus {
	s.AdbpgSegmentStatus = v
	return s
}

func (s *DescribeHealthStatusResponseBodyStatus) SetAdbpgStatus(v *DescribeHealthStatusResponseBodyStatusAdbpgStatus) *DescribeHealthStatusResponseBodyStatus {
	s.AdbpgStatus = v
	return s
}

func (s *DescribeHealthStatusResponseBodyStatus) SetNodeMasterConnectionStatus(v *DescribeHealthStatusResponseBodyStatusNodeMasterConnectionStatus) *DescribeHealthStatusResponseBodyStatus {
	s.NodeMasterConnectionStatus = v
	return s
}

func (s *DescribeHealthStatusResponseBodyStatus) SetNodeMasterStatus(v *DescribeHealthStatusResponseBodyStatusNodeMasterStatus) *DescribeHealthStatusResponseBodyStatus {
	s.NodeMasterStatus = v
	return s
}

func (s *DescribeHealthStatusResponseBodyStatus) SetNodeSegmentConnectionStatus(v *DescribeHealthStatusResponseBodyStatusNodeSegmentConnectionStatus) *DescribeHealthStatusResponseBodyStatus {
	s.NodeSegmentConnectionStatus = v
	return s
}

func (s *DescribeHealthStatusResponseBodyStatus) SetNodeSegmentDiskStatus(v *DescribeHealthStatusResponseBodyStatusNodeSegmentDiskStatus) *DescribeHealthStatusResponseBodyStatus {
	s.NodeSegmentDiskStatus = v
	return s
}

type DescribeHealthStatusResponseBodyStatusAdbgpSegmentDiskUsagePercentMax struct {
	// The status corresponding to the maximum storage usage among all compute nodes. Valid values:
	//
	// *   **critical**: The compute node storage usage is greater than or equal to 90%. In this case, the instance is locked.
	// *   **warning**: The compute node storage usage is greater than or equal to 80% and less than 90%.
	// *   **healthy**: The compute node storage usage is less than 80%.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The metric value of maximum compute node storage usage.
	//
	// Unit: %.
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeHealthStatusResponseBodyStatusAdbgpSegmentDiskUsagePercentMax) String() string {
	return tea.Prettify(s)
}

func (s DescribeHealthStatusResponseBodyStatusAdbgpSegmentDiskUsagePercentMax) GoString() string {
	return s.String()
}

func (s *DescribeHealthStatusResponseBodyStatusAdbgpSegmentDiskUsagePercentMax) SetStatus(v string) *DescribeHealthStatusResponseBodyStatusAdbgpSegmentDiskUsagePercentMax {
	s.Status = &v
	return s
}

func (s *DescribeHealthStatusResponseBodyStatusAdbgpSegmentDiskUsagePercentMax) SetValue(v float32) *DescribeHealthStatusResponseBodyStatusAdbgpSegmentDiskUsagePercentMax {
	s.Value = &v
	return s
}

type DescribeHealthStatusResponseBodyStatusAdbpgConnectionStatus struct {
	// The connection health status of the instance. Valid values:
	//
	// *   **critical**: The instance connection usage is greater than 95%. In this case, this metric is marked in red in the console.
	// *   **warning**: The instance connection usage is greater than 90% and less than or equal to 95%. In this case, this metric is marked in yellow in the console.
	// *   **healthy**: The instance connection usage is less than or equal to 90%. In this case, this metric is marked in green in the console.
	//
	// >  The instance connection usage is the maximum connection usage among all the coordinator and compute nodes.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The metric value of instance connection usage.
	//
	// Unit: %.
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeHealthStatusResponseBodyStatusAdbpgConnectionStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeHealthStatusResponseBodyStatusAdbpgConnectionStatus) GoString() string {
	return s.String()
}

func (s *DescribeHealthStatusResponseBodyStatusAdbpgConnectionStatus) SetStatus(v string) *DescribeHealthStatusResponseBodyStatusAdbpgConnectionStatus {
	s.Status = &v
	return s
}

func (s *DescribeHealthStatusResponseBodyStatusAdbpgConnectionStatus) SetValue(v float32) *DescribeHealthStatusResponseBodyStatusAdbpgConnectionStatus {
	s.Value = &v
	return s
}

type DescribeHealthStatusResponseBodyStatusAdbpgDiskStatus struct {
	// The storage status of the instance. Valid values:
	//
	// *   **critical**: The instance storage usage is greater than or equal to 90%. In this case, this metric is marked in red in the console and the instance is locked.
	// *   **warning**: The instance storage usage is greater than or equal to 70% and less than 90%. In this case, this metric is marked in yellow in the console.
	// *   **healthy**: The instance storage usage is less than 70%. In this case, this metric is marked in green in the console.
	//
	// >  The instance storage usage is the average storage usage of all compute nodes.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The metric value of instance storage usage.
	//
	// Unit: %.
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeHealthStatusResponseBodyStatusAdbpgDiskStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeHealthStatusResponseBodyStatusAdbpgDiskStatus) GoString() string {
	return s.String()
}

func (s *DescribeHealthStatusResponseBodyStatusAdbpgDiskStatus) SetStatus(v string) *DescribeHealthStatusResponseBodyStatusAdbpgDiskStatus {
	s.Status = &v
	return s
}

func (s *DescribeHealthStatusResponseBodyStatusAdbpgDiskStatus) SetValue(v float32) *DescribeHealthStatusResponseBodyStatusAdbpgDiskStatus {
	s.Value = &v
	return s
}

type DescribeHealthStatusResponseBodyStatusAdbpgDiskUsagePercent struct {
	// The status corresponding to the storage usage of the instance. Valid values:
	//
	// *   **critical**: The instance storage usage is greater than or equal to 90%. In this case, the instance is locked.
	// *   **warning**: The instance storage usage is greater than or equal to 70% and less than 90%.
	// *   **healthy**: The instance storage usage is less than 70%.
	//
	// >  The instance storage usage is the average storage usage of all compute nodes.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The metric value of instance storage usage.
	//
	// Unit: %.
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeHealthStatusResponseBodyStatusAdbpgDiskUsagePercent) String() string {
	return tea.Prettify(s)
}

func (s DescribeHealthStatusResponseBodyStatusAdbpgDiskUsagePercent) GoString() string {
	return s.String()
}

func (s *DescribeHealthStatusResponseBodyStatusAdbpgDiskUsagePercent) SetStatus(v string) *DescribeHealthStatusResponseBodyStatusAdbpgDiskUsagePercent {
	s.Status = &v
	return s
}

func (s *DescribeHealthStatusResponseBodyStatusAdbpgDiskUsagePercent) SetValue(v float32) *DescribeHealthStatusResponseBodyStatusAdbpgDiskUsagePercent {
	s.Value = &v
	return s
}

type DescribeHealthStatusResponseBodyStatusAdbpgInstanceColdDataGb struct {
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeHealthStatusResponseBodyStatusAdbpgInstanceColdDataGb) String() string {
	return tea.Prettify(s)
}

func (s DescribeHealthStatusResponseBodyStatusAdbpgInstanceColdDataGb) GoString() string {
	return s.String()
}

func (s *DescribeHealthStatusResponseBodyStatusAdbpgInstanceColdDataGb) SetValue(v float32) *DescribeHealthStatusResponseBodyStatusAdbpgInstanceColdDataGb {
	s.Value = &v
	return s
}

type DescribeHealthStatusResponseBodyStatusAdbpgInstanceHotDataGb struct {
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeHealthStatusResponseBodyStatusAdbpgInstanceHotDataGb) String() string {
	return tea.Prettify(s)
}

func (s DescribeHealthStatusResponseBodyStatusAdbpgInstanceHotDataGb) GoString() string {
	return s.String()
}

func (s *DescribeHealthStatusResponseBodyStatusAdbpgInstanceHotDataGb) SetValue(v float32) *DescribeHealthStatusResponseBodyStatusAdbpgInstanceHotDataGb {
	s.Value = &v
	return s
}

type DescribeHealthStatusResponseBodyStatusAdbpgInstanceTotalDataGb struct {
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeHealthStatusResponseBodyStatusAdbpgInstanceTotalDataGb) String() string {
	return tea.Prettify(s)
}

func (s DescribeHealthStatusResponseBodyStatusAdbpgInstanceTotalDataGb) GoString() string {
	return s.String()
}

func (s *DescribeHealthStatusResponseBodyStatusAdbpgInstanceTotalDataGb) SetValue(v float32) *DescribeHealthStatusResponseBodyStatusAdbpgInstanceTotalDataGb {
	s.Value = &v
	return s
}

type DescribeHealthStatusResponseBodyStatusAdbpgMasterDiskUsagePercentMax struct {
	// The status corresponding to the maximum storage usage of the coordinator node. Valid values:
	//
	// *   **critical**: The coordinator node storage usage is greater than or equal to 90%. In this case, the instance is locked.
	// *   **warning**: The coordinator node storage usage is greater than or equal to 70% and less than 90%.
	// *   **healthy**: The coordinator node storage usage is less than 70%.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The metric value of maximum coordinator node storage usage.
	//
	// Unit: %.
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeHealthStatusResponseBodyStatusAdbpgMasterDiskUsagePercentMax) String() string {
	return tea.Prettify(s)
}

func (s DescribeHealthStatusResponseBodyStatusAdbpgMasterDiskUsagePercentMax) GoString() string {
	return s.String()
}

func (s *DescribeHealthStatusResponseBodyStatusAdbpgMasterDiskUsagePercentMax) SetStatus(v string) *DescribeHealthStatusResponseBodyStatusAdbpgMasterDiskUsagePercentMax {
	s.Status = &v
	return s
}

func (s *DescribeHealthStatusResponseBodyStatusAdbpgMasterDiskUsagePercentMax) SetValue(v float32) *DescribeHealthStatusResponseBodyStatusAdbpgMasterDiskUsagePercentMax {
	s.Value = &v
	return s
}

type DescribeHealthStatusResponseBodyStatusAdbpgMasterStatus struct {
	// The availability status of the coordinator node. Valid values:
	//
	// *   **critical**: Both the primary and standby coordinator nodes are unavailable. In this case, this metric is marked in red in the console.
	// *   **warning**: The primary or standby coordinator node is unavailable. In this case, this metric is marked in yellow in the console.
	// *   **healthy**: Both the primary and standby coordinator nodes are available. In this case, this metric is marked in green in the console.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The metric value of coordinator node availability status. Valid values:
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeHealthStatusResponseBodyStatusAdbpgMasterStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeHealthStatusResponseBodyStatusAdbpgMasterStatus) GoString() string {
	return s.String()
}

func (s *DescribeHealthStatusResponseBodyStatusAdbpgMasterStatus) SetStatus(v string) *DescribeHealthStatusResponseBodyStatusAdbpgMasterStatus {
	s.Status = &v
	return s
}

func (s *DescribeHealthStatusResponseBodyStatusAdbpgMasterStatus) SetValue(v float32) *DescribeHealthStatusResponseBodyStatusAdbpgMasterStatus {
	s.Value = &v
	return s
}

type DescribeHealthStatusResponseBodyStatusAdbpgSegmentStatus struct {
	// The availability status of compute nodes. Valid values:
	//
	// *   **critical**: All the primary and secondary compute nodes are unavailable. In this case, this metric is marked in red in the console.
	// *   **warning**: Fifty percent or more than fifty percent of compute nodes are unavailable. In this case, this metric is marked in yellow in the console.
	// *   **healthy**: All compute nodes are available. In this case, this metric is marked in green in the console.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The metric value of compute node availability status.
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeHealthStatusResponseBodyStatusAdbpgSegmentStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeHealthStatusResponseBodyStatusAdbpgSegmentStatus) GoString() string {
	return s.String()
}

func (s *DescribeHealthStatusResponseBodyStatusAdbpgSegmentStatus) SetStatus(v string) *DescribeHealthStatusResponseBodyStatusAdbpgSegmentStatus {
	s.Status = &v
	return s
}

func (s *DescribeHealthStatusResponseBodyStatusAdbpgSegmentStatus) SetValue(v float32) *DescribeHealthStatusResponseBodyStatusAdbpgSegmentStatus {
	s.Value = &v
	return s
}

type DescribeHealthStatusResponseBodyStatusAdbpgStatus struct {
	// The health status of the instance. Valid values:
	//
	// *   **critical**: The coordinator node or a compute node is unavailable. In this case, this metric is marked in red in the console.
	// *   **healthy**: All nodes are available. In this case, this metric is marked in green in the console.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The metric value of instance health status. Valid values:
	//
	// *   **1**: healthy
	// *   **0**: critical
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeHealthStatusResponseBodyStatusAdbpgStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeHealthStatusResponseBodyStatusAdbpgStatus) GoString() string {
	return s.String()
}

func (s *DescribeHealthStatusResponseBodyStatusAdbpgStatus) SetStatus(v string) *DescribeHealthStatusResponseBodyStatusAdbpgStatus {
	s.Status = &v
	return s
}

func (s *DescribeHealthStatusResponseBodyStatusAdbpgStatus) SetValue(v float32) *DescribeHealthStatusResponseBodyStatusAdbpgStatus {
	s.Value = &v
	return s
}

type DescribeHealthStatusResponseBodyStatusNodeMasterConnectionStatus struct {
	// The connection health status of the coordinator node. Valid values:
	//
	// *   **critical**: The coordinator node connection usage is greater than 95%. In this case, this metric is marked in red in the console.
	// *   **warning**: The coordinator node connection usage is greater than or equal to 90% and less than 95%. In this case, this metric is marked in yellow in the console.
	// *   **healthy**: The coordinator node connection usage is less than 90%. In this case, this metric is marked in green in the console.
	//
	// >  The coordinator node connection usage is the maximum connection usage of the coordinator node.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The metric value of coordinator node connection usage.
	//
	// Unit: %.
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeHealthStatusResponseBodyStatusNodeMasterConnectionStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeHealthStatusResponseBodyStatusNodeMasterConnectionStatus) GoString() string {
	return s.String()
}

func (s *DescribeHealthStatusResponseBodyStatusNodeMasterConnectionStatus) SetStatus(v string) *DescribeHealthStatusResponseBodyStatusNodeMasterConnectionStatus {
	s.Status = &v
	return s
}

func (s *DescribeHealthStatusResponseBodyStatusNodeMasterConnectionStatus) SetValue(v float32) *DescribeHealthStatusResponseBodyStatusNodeMasterConnectionStatus {
	s.Value = &v
	return s
}

type DescribeHealthStatusResponseBodyStatusNodeMasterStatus struct {
	// The health status of the coordinator node. Valid values:
	//
	// *   **critical**: The primary or standby coordinator node is unavailable. In this case, this metric is marked in red in the console.
	// *   **healthy**: Both the primary and standby coordinator nodes are available. In this case, this metric is marked in green in the console.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The metric value of coordinator node health status. Valid values:
	//
	// *   **1**: healthy
	// *   **0**: critical
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeHealthStatusResponseBodyStatusNodeMasterStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeHealthStatusResponseBodyStatusNodeMasterStatus) GoString() string {
	return s.String()
}

func (s *DescribeHealthStatusResponseBodyStatusNodeMasterStatus) SetStatus(v string) *DescribeHealthStatusResponseBodyStatusNodeMasterStatus {
	s.Status = &v
	return s
}

func (s *DescribeHealthStatusResponseBodyStatusNodeMasterStatus) SetValue(v float32) *DescribeHealthStatusResponseBodyStatusNodeMasterStatus {
	s.Value = &v
	return s
}

type DescribeHealthStatusResponseBodyStatusNodeSegmentConnectionStatus struct {
	// The connection health status of compute nodes. Valid values:
	//
	// *   **critical**: The compute node connection usage is greater than or equal to 95%. In this case, this metric is marked in red in the console.
	// *   **warning**: The compute node connection usage is greater than or equal to 90% and less than 95%. In this case, this metric is marked in yellow in the console.
	// *   **healthy**: The compute node connection usage is less than 90%. In this case, this metric is marked in green in the console.
	//
	// >  The compute node connection usage is the maximum connection usage among all compute nodes.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The metric value of maximum compute node connection usage.
	//
	// Unit: %.
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeHealthStatusResponseBodyStatusNodeSegmentConnectionStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeHealthStatusResponseBodyStatusNodeSegmentConnectionStatus) GoString() string {
	return s.String()
}

func (s *DescribeHealthStatusResponseBodyStatusNodeSegmentConnectionStatus) SetStatus(v string) *DescribeHealthStatusResponseBodyStatusNodeSegmentConnectionStatus {
	s.Status = &v
	return s
}

func (s *DescribeHealthStatusResponseBodyStatusNodeSegmentConnectionStatus) SetValue(v float32) *DescribeHealthStatusResponseBodyStatusNodeSegmentConnectionStatus {
	s.Value = &v
	return s
}

type DescribeHealthStatusResponseBodyStatusNodeSegmentDiskStatus struct {
	// The storage status of compute nodes. Valid values:
	//
	// *   **critical**: The compute node storage usage is greater than or equal to 90%. In this case, this metric is marked in red in the console and the instance is locked.
	// *   **warning**: The compute node storage usage is greater than or equal to 80% and less than 90%. In this case, this metric is marked in yellow in the console.
	// *   **healthy**: The compute node storage usage is less than 80%. In this case, this metric is marked in green in the console.
	//
	// >  The compute node storage usage is the maximum storage usage among all compute nodes.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The metric value of maximum compute node storage usage.
	//
	// Unit: %.
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeHealthStatusResponseBodyStatusNodeSegmentDiskStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeHealthStatusResponseBodyStatusNodeSegmentDiskStatus) GoString() string {
	return s.String()
}

func (s *DescribeHealthStatusResponseBodyStatusNodeSegmentDiskStatus) SetStatus(v string) *DescribeHealthStatusResponseBodyStatusNodeSegmentDiskStatus {
	s.Status = &v
	return s
}

func (s *DescribeHealthStatusResponseBodyStatusNodeSegmentDiskStatus) SetValue(v float32) *DescribeHealthStatusResponseBodyStatusNodeSegmentDiskStatus {
	s.Value = &v
	return s
}

type DescribeHealthStatusResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeHealthStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeHealthStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeHealthStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeHealthStatusResponse) SetHeaders(v map[string]*string) *DescribeHealthStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeHealthStatusResponse) SetStatusCode(v int32) *DescribeHealthStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHealthStatusResponse) SetBody(v *DescribeHealthStatusResponseBody) *DescribeHealthStatusResponse {
	s.Body = v
	return s
}

type DescribeLogBackupsRequest struct {
	// The ID of the instance.
	//
	// >  You can call the [DescribeDBInstances](~~86911~~) operation to query the IDs of all AnalyticDB for PostgreSQL instances in a specific region.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The end of the time range to query. The end time must be later than the start time. Specify the time in the yyyy-MM-ddTHH:mmZ format. The time must be in UTC.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The number of the page to return. The value must be an integer that is greater than 0. Default value: **1**.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values:
	//
	// *   **30**
	// *   **50**
	// *   **100**
	//
	// Default value: **30**.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The beginning of the time range to query. Specify the time in the yyyy-MM-ddTHH:mmZ format. The time must be in UTC.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeLogBackupsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogBackupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeLogBackupsRequest) SetDBInstanceId(v string) *DescribeLogBackupsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeLogBackupsRequest) SetEndTime(v string) *DescribeLogBackupsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeLogBackupsRequest) SetPageNumber(v int32) *DescribeLogBackupsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeLogBackupsRequest) SetPageSize(v int32) *DescribeLogBackupsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeLogBackupsRequest) SetStartTime(v string) *DescribeLogBackupsRequest {
	s.StartTime = &v
	return s
}

type DescribeLogBackupsResponseBody struct {
	// Details of the backup sets.
	Items []*DescribeLogBackupsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The page number of the returned page.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of backup sets on the current page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The total size of logs in the time range. Unit: bytes.
	TotalLogSize *int64 `json:"TotalLogSize,omitempty" xml:"TotalLogSize,omitempty"`
}

func (s DescribeLogBackupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogBackupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLogBackupsResponseBody) SetItems(v []*DescribeLogBackupsResponseBodyItems) *DescribeLogBackupsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeLogBackupsResponseBody) SetPageNumber(v int32) *DescribeLogBackupsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeLogBackupsResponseBody) SetPageSize(v int32) *DescribeLogBackupsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeLogBackupsResponseBody) SetRequestId(v string) *DescribeLogBackupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLogBackupsResponseBody) SetTotalCount(v int32) *DescribeLogBackupsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeLogBackupsResponseBody) SetTotalLogSize(v int64) *DescribeLogBackupsResponseBody {
	s.TotalLogSize = &v
	return s
}

type DescribeLogBackupsResponseBodyItems struct {
	// The ID of the backup set.
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The ID of the coordinator node.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The name of the log backup set that is stored in Object Storage Service (OSS).
	LogFileName *string `json:"LogFileName,omitempty" xml:"LogFileName,omitempty"`
	// The size of the log backup set. Unit: bytes.
	LogFileSize *int64 `json:"LogFileSize,omitempty" xml:"LogFileSize,omitempty"`
	// The timestamp of the log.
	LogTime *string `json:"LogTime,omitempty" xml:"LogTime,omitempty"`
	// The name of the compute node.
	SegmentName *string `json:"SegmentName,omitempty" xml:"SegmentName,omitempty"`
}

func (s DescribeLogBackupsResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogBackupsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeLogBackupsResponseBodyItems) SetBackupId(v string) *DescribeLogBackupsResponseBodyItems {
	s.BackupId = &v
	return s
}

func (s *DescribeLogBackupsResponseBodyItems) SetDBInstanceId(v string) *DescribeLogBackupsResponseBodyItems {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeLogBackupsResponseBodyItems) SetLogFileName(v string) *DescribeLogBackupsResponseBodyItems {
	s.LogFileName = &v
	return s
}

func (s *DescribeLogBackupsResponseBodyItems) SetLogFileSize(v int64) *DescribeLogBackupsResponseBodyItems {
	s.LogFileSize = &v
	return s
}

func (s *DescribeLogBackupsResponseBodyItems) SetLogTime(v string) *DescribeLogBackupsResponseBodyItems {
	s.LogTime = &v
	return s
}

func (s *DescribeLogBackupsResponseBodyItems) SetSegmentName(v string) *DescribeLogBackupsResponseBodyItems {
	s.SegmentName = &v
	return s
}

type DescribeLogBackupsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeLogBackupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeLogBackupsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogBackupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeLogBackupsResponse) SetHeaders(v map[string]*string) *DescribeLogBackupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeLogBackupsResponse) SetStatusCode(v int32) *DescribeLogBackupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLogBackupsResponse) SetBody(v *DescribeLogBackupsResponseBody) *DescribeLogBackupsResponse {
	s.Body = v
	return s
}

type DescribeModifyParameterLogRequest struct {
	// The ID of the instance.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The end of the time range to query.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The beginning of the time range to query.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeModifyParameterLogRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeModifyParameterLogRequest) GoString() string {
	return s.String()
}

func (s *DescribeModifyParameterLogRequest) SetDBInstanceId(v string) *DescribeModifyParameterLogRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeModifyParameterLogRequest) SetEndTime(v string) *DescribeModifyParameterLogRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeModifyParameterLogRequest) SetStartTime(v string) *DescribeModifyParameterLogRequest {
	s.StartTime = &v
	return s
}

type DescribeModifyParameterLogResponseBody struct {
	// Details about the parameter reconfiguration logs.
	Changelogs []*DescribeModifyParameterLogResponseBodyChangelogs `json:"Changelogs,omitempty" xml:"Changelogs,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeModifyParameterLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeModifyParameterLogResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeModifyParameterLogResponseBody) SetChangelogs(v []*DescribeModifyParameterLogResponseBodyChangelogs) *DescribeModifyParameterLogResponseBody {
	s.Changelogs = v
	return s
}

func (s *DescribeModifyParameterLogResponseBody) SetRequestId(v string) *DescribeModifyParameterLogResponseBody {
	s.RequestId = &v
	return s
}

type DescribeModifyParameterLogResponseBodyChangelogs struct {
	// The time when the configuration change takes effect.
	EffectTime *string `json:"EffectTime,omitempty" xml:"EffectTime,omitempty"`
	// The name of the parameter.
	ParameterName *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	// Indicates whether the configuration change takes effect.
	ParameterValid *string `json:"ParameterValid,omitempty" xml:"ParameterValid,omitempty"`
	// The original value of the parameter.
	ParameterValueAfter *string `json:"ParameterValueAfter,omitempty" xml:"ParameterValueAfter,omitempty"`
	// The new value of the parameter.
	ParameterValueBefore *string `json:"ParameterValueBefore,omitempty" xml:"ParameterValueBefore,omitempty"`
}

func (s DescribeModifyParameterLogResponseBodyChangelogs) String() string {
	return tea.Prettify(s)
}

func (s DescribeModifyParameterLogResponseBodyChangelogs) GoString() string {
	return s.String()
}

func (s *DescribeModifyParameterLogResponseBodyChangelogs) SetEffectTime(v string) *DescribeModifyParameterLogResponseBodyChangelogs {
	s.EffectTime = &v
	return s
}

func (s *DescribeModifyParameterLogResponseBodyChangelogs) SetParameterName(v string) *DescribeModifyParameterLogResponseBodyChangelogs {
	s.ParameterName = &v
	return s
}

func (s *DescribeModifyParameterLogResponseBodyChangelogs) SetParameterValid(v string) *DescribeModifyParameterLogResponseBodyChangelogs {
	s.ParameterValid = &v
	return s
}

func (s *DescribeModifyParameterLogResponseBodyChangelogs) SetParameterValueAfter(v string) *DescribeModifyParameterLogResponseBodyChangelogs {
	s.ParameterValueAfter = &v
	return s
}

func (s *DescribeModifyParameterLogResponseBodyChangelogs) SetParameterValueBefore(v string) *DescribeModifyParameterLogResponseBodyChangelogs {
	s.ParameterValueBefore = &v
	return s
}

type DescribeModifyParameterLogResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeModifyParameterLogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeModifyParameterLogResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeModifyParameterLogResponse) GoString() string {
	return s.String()
}

func (s *DescribeModifyParameterLogResponse) SetHeaders(v map[string]*string) *DescribeModifyParameterLogResponse {
	s.Headers = v
	return s
}

func (s *DescribeModifyParameterLogResponse) SetStatusCode(v int32) *DescribeModifyParameterLogResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeModifyParameterLogResponse) SetBody(v *DescribeModifyParameterLogResponseBody) *DescribeModifyParameterLogResponse {
	s.Body = v
	return s
}

type DescribeNamespaceRequest struct {
	DBInstanceId           *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	ManagerAccount         *string `json:"ManagerAccount,omitempty" xml:"ManagerAccount,omitempty"`
	ManagerAccountPassword *string `json:"ManagerAccountPassword,omitempty" xml:"ManagerAccountPassword,omitempty"`
	Namespace              *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	OwnerId                *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId               *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeNamespaceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeNamespaceRequest) GoString() string {
	return s.String()
}

func (s *DescribeNamespaceRequest) SetDBInstanceId(v string) *DescribeNamespaceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeNamespaceRequest) SetManagerAccount(v string) *DescribeNamespaceRequest {
	s.ManagerAccount = &v
	return s
}

func (s *DescribeNamespaceRequest) SetManagerAccountPassword(v string) *DescribeNamespaceRequest {
	s.ManagerAccountPassword = &v
	return s
}

func (s *DescribeNamespaceRequest) SetNamespace(v string) *DescribeNamespaceRequest {
	s.Namespace = &v
	return s
}

func (s *DescribeNamespaceRequest) SetOwnerId(v int64) *DescribeNamespaceRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeNamespaceRequest) SetRegionId(v string) *DescribeNamespaceRequest {
	s.RegionId = &v
	return s
}

type DescribeNamespaceResponseBody struct {
	DBInstanceId  *string            `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	Message       *string            `json:"Message,omitempty" xml:"Message,omitempty"`
	Namespace     *string            `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	NamespaceInfo map[string]*string `json:"NamespaceInfo,omitempty" xml:"NamespaceInfo,omitempty"`
	RegionId      *string            `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RequestId     *string            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status        *string            `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeNamespaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeNamespaceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNamespaceResponseBody) SetDBInstanceId(v string) *DescribeNamespaceResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeNamespaceResponseBody) SetMessage(v string) *DescribeNamespaceResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeNamespaceResponseBody) SetNamespace(v string) *DescribeNamespaceResponseBody {
	s.Namespace = &v
	return s
}

func (s *DescribeNamespaceResponseBody) SetNamespaceInfo(v map[string]*string) *DescribeNamespaceResponseBody {
	s.NamespaceInfo = v
	return s
}

func (s *DescribeNamespaceResponseBody) SetRegionId(v string) *DescribeNamespaceResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeNamespaceResponseBody) SetRequestId(v string) *DescribeNamespaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNamespaceResponseBody) SetStatus(v string) *DescribeNamespaceResponseBody {
	s.Status = &v
	return s
}

type DescribeNamespaceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeNamespaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeNamespaceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeNamespaceResponse) GoString() string {
	return s.String()
}

func (s *DescribeNamespaceResponse) SetHeaders(v map[string]*string) *DescribeNamespaceResponse {
	s.Headers = v
	return s
}

func (s *DescribeNamespaceResponse) SetStatusCode(v int32) *DescribeNamespaceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNamespaceResponse) SetBody(v *DescribeNamespaceResponseBody) *DescribeNamespaceResponse {
	s.Body = v
	return s
}

type DescribeParametersRequest struct {
	// The ID of the instance.
	//
	// >  You can call the [DescribeDBInstances](~~86911~~) operation to query the details of all AnalyticDB for PostgreSQL instances in a specific region, including instance IDs.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s DescribeParametersRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeParametersRequest) GoString() string {
	return s.String()
}

func (s *DescribeParametersRequest) SetDBInstanceId(v string) *DescribeParametersRequest {
	s.DBInstanceId = &v
	return s
}

type DescribeParametersResponseBody struct {
	// Details of the parameters.
	Parameters []*DescribeParametersResponseBodyParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeParametersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeParametersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeParametersResponseBody) SetParameters(v []*DescribeParametersResponseBodyParameters) *DescribeParametersResponseBody {
	s.Parameters = v
	return s
}

func (s *DescribeParametersResponseBody) SetRequestId(v string) *DescribeParametersResponseBody {
	s.RequestId = &v
	return s
}

type DescribeParametersResponseBodyParameters struct {
	// The current value of the parameter.
	CurrentValue *string `json:"CurrentValue,omitempty" xml:"CurrentValue,omitempty"`
	// Indicates whether a restart is required for parameter modifications to take effect. Valid values:
	//
	// *   **true**
	// *   **false**
	ForceRestartInstance *string `json:"ForceRestartInstance,omitempty" xml:"ForceRestartInstance,omitempty"`
	// Indicates whether the parameter can be modified. Valid values:
	//
	// *   **true**
	// *   **false**
	IsChangeableConfig *string `json:"IsChangeableConfig,omitempty" xml:"IsChangeableConfig,omitempty"`
	// The valid values of the parameter.
	OptionalRange *string `json:"OptionalRange,omitempty" xml:"OptionalRange,omitempty"`
	// The description of the parameter.
	ParameterDescription *string `json:"ParameterDescription,omitempty" xml:"ParameterDescription,omitempty"`
	// The name of the parameter.
	ParameterName *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	// The default value of the parameter.
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s DescribeParametersResponseBodyParameters) String() string {
	return tea.Prettify(s)
}

func (s DescribeParametersResponseBodyParameters) GoString() string {
	return s.String()
}

func (s *DescribeParametersResponseBodyParameters) SetCurrentValue(v string) *DescribeParametersResponseBodyParameters {
	s.CurrentValue = &v
	return s
}

func (s *DescribeParametersResponseBodyParameters) SetForceRestartInstance(v string) *DescribeParametersResponseBodyParameters {
	s.ForceRestartInstance = &v
	return s
}

func (s *DescribeParametersResponseBodyParameters) SetIsChangeableConfig(v string) *DescribeParametersResponseBodyParameters {
	s.IsChangeableConfig = &v
	return s
}

func (s *DescribeParametersResponseBodyParameters) SetOptionalRange(v string) *DescribeParametersResponseBodyParameters {
	s.OptionalRange = &v
	return s
}

func (s *DescribeParametersResponseBodyParameters) SetParameterDescription(v string) *DescribeParametersResponseBodyParameters {
	s.ParameterDescription = &v
	return s
}

func (s *DescribeParametersResponseBodyParameters) SetParameterName(v string) *DescribeParametersResponseBodyParameters {
	s.ParameterName = &v
	return s
}

func (s *DescribeParametersResponseBodyParameters) SetParameterValue(v string) *DescribeParametersResponseBodyParameters {
	s.ParameterValue = &v
	return s
}

type DescribeParametersResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeParametersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeParametersResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeParametersResponse) GoString() string {
	return s.String()
}

func (s *DescribeParametersResponse) SetHeaders(v map[string]*string) *DescribeParametersResponse {
	s.Headers = v
	return s
}

func (s *DescribeParametersResponse) SetStatusCode(v int32) *DescribeParametersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeParametersResponse) SetBody(v *DescribeParametersResponseBody) *DescribeParametersResponse {
	s.Body = v
	return s
}

type DescribeRdsVSwitchsRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region.
	//
	// >  You can call the [DescribeRegions](~~86912~~) operation to query the most recent region list and zone list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the instance belongs.
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The ID of virtual private cloud (VPC).
	//
	// > *   You can call the [DescribeRdsVpcs](~~208327~~) operation to query the available VPCs.
	// > *   This parameter is required.
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The ID of the zone.
	//
	// >  You can call the [DescribeRegions](~~86912~~) operation to query the most recent region list and zone list.
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeRdsVSwitchsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRdsVSwitchsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRdsVSwitchsRequest) SetOwnerAccount(v string) *DescribeRdsVSwitchsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeRdsVSwitchsRequest) SetOwnerId(v int64) *DescribeRdsVSwitchsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRdsVSwitchsRequest) SetRegionId(v string) *DescribeRdsVSwitchsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRdsVSwitchsRequest) SetResourceGroupId(v string) *DescribeRdsVSwitchsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeRdsVSwitchsRequest) SetResourceOwnerAccount(v string) *DescribeRdsVSwitchsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeRdsVSwitchsRequest) SetResourceOwnerId(v int64) *DescribeRdsVSwitchsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeRdsVSwitchsRequest) SetSecurityToken(v string) *DescribeRdsVSwitchsRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeRdsVSwitchsRequest) SetVpcId(v string) *DescribeRdsVSwitchsRequest {
	s.VpcId = &v
	return s
}

func (s *DescribeRdsVSwitchsRequest) SetZoneId(v string) *DescribeRdsVSwitchsRequest {
	s.ZoneId = &v
	return s
}

type DescribeRdsVSwitchsResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Details of the vSwitches.
	VSwitches *DescribeRdsVSwitchsResponseBodyVSwitches `json:"VSwitches,omitempty" xml:"VSwitches,omitempty" type:"Struct"`
}

func (s DescribeRdsVSwitchsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRdsVSwitchsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRdsVSwitchsResponseBody) SetRequestId(v string) *DescribeRdsVSwitchsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRdsVSwitchsResponseBody) SetVSwitches(v *DescribeRdsVSwitchsResponseBodyVSwitches) *DescribeRdsVSwitchsResponseBody {
	s.VSwitches = v
	return s
}

type DescribeRdsVSwitchsResponseBodyVSwitches struct {
	// Details of the vSwitch.
	VSwitch []*DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch `json:"VSwitch,omitempty" xml:"VSwitch,omitempty" type:"Repeated"`
}

func (s DescribeRdsVSwitchsResponseBodyVSwitches) String() string {
	return tea.Prettify(s)
}

func (s DescribeRdsVSwitchsResponseBodyVSwitches) GoString() string {
	return s.String()
}

func (s *DescribeRdsVSwitchsResponseBodyVSwitches) SetVSwitch(v []*DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch) *DescribeRdsVSwitchsResponseBodyVSwitches {
	s.VSwitch = v
	return s
}

type DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch struct {
	// An invalid parameter. It is no longer returned when you call this operation.
	AliUid *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// An invalid parameter. It is no longer returned when you call this operation.
	Bid *string `json:"Bid,omitempty" xml:"Bid,omitempty"`
	// The CIDR block of the vSwitch.
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	// An invalid parameter. It is no longer returned when you call this operation.
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// An invalid parameter. It is no longer returned when you call this operation.
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// Indicates whether the vSwitch is the default vSwitch. Valid values:
	//
	// *   **true**
	// *   **false**
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// The ID of the zone.
	IzNo *string `json:"IzNo,omitempty" xml:"IzNo,omitempty"`
	// An invalid parameter. It is no longer returned when you call this operation.
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// The state of the vSwitch. If **Available** is returned, the vSwitch is available.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the vSwitch.
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The name of the vSwitch.
	VSwitchName *string `json:"VSwitchName,omitempty" xml:"VSwitchName,omitempty"`
}

func (s DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch) String() string {
	return tea.Prettify(s)
}

func (s DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch) GoString() string {
	return s.String()
}

func (s *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch) SetAliUid(v string) *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch {
	s.AliUid = &v
	return s
}

func (s *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch) SetBid(v string) *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch {
	s.Bid = &v
	return s
}

func (s *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch) SetCidrBlock(v string) *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch {
	s.CidrBlock = &v
	return s
}

func (s *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch) SetGmtCreate(v string) *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch {
	s.GmtCreate = &v
	return s
}

func (s *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch) SetGmtModified(v string) *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch {
	s.GmtModified = &v
	return s
}

func (s *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch) SetIsDefault(v bool) *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch {
	s.IsDefault = &v
	return s
}

func (s *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch) SetIzNo(v string) *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch {
	s.IzNo = &v
	return s
}

func (s *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch) SetRegionNo(v string) *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch {
	s.RegionNo = &v
	return s
}

func (s *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch) SetStatus(v string) *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch {
	s.Status = &v
	return s
}

func (s *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch) SetVSwitchId(v string) *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch {
	s.VSwitchId = &v
	return s
}

func (s *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch) SetVSwitchName(v string) *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch {
	s.VSwitchName = &v
	return s
}

type DescribeRdsVSwitchsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeRdsVSwitchsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRdsVSwitchsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRdsVSwitchsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRdsVSwitchsResponse) SetHeaders(v map[string]*string) *DescribeRdsVSwitchsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRdsVSwitchsResponse) SetStatusCode(v int32) *DescribeRdsVSwitchsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRdsVSwitchsResponse) SetBody(v *DescribeRdsVSwitchsResponseBody) *DescribeRdsVSwitchsResponse {
	s.Body = v
	return s
}

type DescribeRdsVpcsRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region.
	//
	// >  You can call the [DescribeRegions](~~86912~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the instance belongs.
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The ID of the zone.
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeRdsVpcsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRdsVpcsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRdsVpcsRequest) SetOwnerAccount(v string) *DescribeRdsVpcsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeRdsVpcsRequest) SetOwnerId(v int64) *DescribeRdsVpcsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRdsVpcsRequest) SetRegionId(v string) *DescribeRdsVpcsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRdsVpcsRequest) SetResourceGroupId(v string) *DescribeRdsVpcsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeRdsVpcsRequest) SetResourceOwnerAccount(v string) *DescribeRdsVpcsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeRdsVpcsRequest) SetResourceOwnerId(v int64) *DescribeRdsVpcsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeRdsVpcsRequest) SetSecurityToken(v string) *DescribeRdsVpcsRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeRdsVpcsRequest) SetZoneId(v string) *DescribeRdsVpcsRequest {
	s.ZoneId = &v
	return s
}

type DescribeRdsVpcsResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Details of the VPCs.
	Vpcs *DescribeRdsVpcsResponseBodyVpcs `json:"Vpcs,omitempty" xml:"Vpcs,omitempty" type:"Struct"`
}

func (s DescribeRdsVpcsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRdsVpcsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRdsVpcsResponseBody) SetRequestId(v string) *DescribeRdsVpcsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRdsVpcsResponseBody) SetVpcs(v *DescribeRdsVpcsResponseBodyVpcs) *DescribeRdsVpcsResponseBody {
	s.Vpcs = v
	return s
}

type DescribeRdsVpcsResponseBodyVpcs struct {
	// Details of the VPC.
	Vpc []*DescribeRdsVpcsResponseBodyVpcsVpc `json:"Vpc,omitempty" xml:"Vpc,omitempty" type:"Repeated"`
}

func (s DescribeRdsVpcsResponseBodyVpcs) String() string {
	return tea.Prettify(s)
}

func (s DescribeRdsVpcsResponseBodyVpcs) GoString() string {
	return s.String()
}

func (s *DescribeRdsVpcsResponseBodyVpcs) SetVpc(v []*DescribeRdsVpcsResponseBodyVpcsVpc) *DescribeRdsVpcsResponseBodyVpcs {
	s.Vpc = v
	return s
}

type DescribeRdsVpcsResponseBodyVpcsVpc struct {
	// An invalid parameter. It is no longer returned when you call this operation.
	AliUid *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// An invalid parameter. It is no longer returned when you call this operation.
	Bid *string `json:"Bid,omitempty" xml:"Bid,omitempty"`
	// The CIDR block of the VPC.
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	// An invalid parameter. It is no longer returned when you call this operation.
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// An invalid parameter. It is no longer returned when you call this operation.
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// Indicates whether the VPC is the default VPC. Valid values:
	//
	// *   **true**
	// *   **false**
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// The ID of the region.
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// The state of the VPC. If **Available** is returned, the VPC is available.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Details of the vSwitches.
	VSwitchs []*DescribeRdsVpcsResponseBodyVpcsVpcVSwitchs `json:"VSwitchs,omitempty" xml:"VSwitchs,omitempty" type:"Repeated"`
	// The ID of VPC.
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The name of the VPC.
	VpcName *string `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
}

func (s DescribeRdsVpcsResponseBodyVpcsVpc) String() string {
	return tea.Prettify(s)
}

func (s DescribeRdsVpcsResponseBodyVpcsVpc) GoString() string {
	return s.String()
}

func (s *DescribeRdsVpcsResponseBodyVpcsVpc) SetAliUid(v string) *DescribeRdsVpcsResponseBodyVpcsVpc {
	s.AliUid = &v
	return s
}

func (s *DescribeRdsVpcsResponseBodyVpcsVpc) SetBid(v string) *DescribeRdsVpcsResponseBodyVpcsVpc {
	s.Bid = &v
	return s
}

func (s *DescribeRdsVpcsResponseBodyVpcsVpc) SetCidrBlock(v string) *DescribeRdsVpcsResponseBodyVpcsVpc {
	s.CidrBlock = &v
	return s
}

func (s *DescribeRdsVpcsResponseBodyVpcsVpc) SetGmtCreate(v string) *DescribeRdsVpcsResponseBodyVpcsVpc {
	s.GmtCreate = &v
	return s
}

func (s *DescribeRdsVpcsResponseBodyVpcsVpc) SetGmtModified(v string) *DescribeRdsVpcsResponseBodyVpcsVpc {
	s.GmtModified = &v
	return s
}

func (s *DescribeRdsVpcsResponseBodyVpcsVpc) SetIsDefault(v bool) *DescribeRdsVpcsResponseBodyVpcsVpc {
	s.IsDefault = &v
	return s
}

func (s *DescribeRdsVpcsResponseBodyVpcsVpc) SetRegionNo(v string) *DescribeRdsVpcsResponseBodyVpcsVpc {
	s.RegionNo = &v
	return s
}

func (s *DescribeRdsVpcsResponseBodyVpcsVpc) SetStatus(v string) *DescribeRdsVpcsResponseBodyVpcsVpc {
	s.Status = &v
	return s
}

func (s *DescribeRdsVpcsResponseBodyVpcsVpc) SetVSwitchs(v []*DescribeRdsVpcsResponseBodyVpcsVpcVSwitchs) *DescribeRdsVpcsResponseBodyVpcsVpc {
	s.VSwitchs = v
	return s
}

func (s *DescribeRdsVpcsResponseBodyVpcsVpc) SetVpcId(v string) *DescribeRdsVpcsResponseBodyVpcsVpc {
	s.VpcId = &v
	return s
}

func (s *DescribeRdsVpcsResponseBodyVpcsVpc) SetVpcName(v string) *DescribeRdsVpcsResponseBodyVpcsVpc {
	s.VpcName = &v
	return s
}

type DescribeRdsVpcsResponseBodyVpcsVpcVSwitchs struct {
	// The CIDR block of the vSwitch.
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	// An invalid parameter. It is no longer returned when you call this operation.
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// An invalid parameter. It is no longer returned when you call this operation.
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// Indicates whether the vSwitch is the default vSwitch. Valid values:
	//
	// *   **true**
	// *   **false**
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// The ID of the zone to which the vSwitch belongs.
	IzNo *string `json:"IzNo,omitempty" xml:"IzNo,omitempty"`
	// The state of the vSwitch. If **Available** is returned, the vSwitch is available.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the vSwitch.
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The name of the vSwitch.
	VSwitchName *string `json:"VSwitchName,omitempty" xml:"VSwitchName,omitempty"`
}

func (s DescribeRdsVpcsResponseBodyVpcsVpcVSwitchs) String() string {
	return tea.Prettify(s)
}

func (s DescribeRdsVpcsResponseBodyVpcsVpcVSwitchs) GoString() string {
	return s.String()
}

func (s *DescribeRdsVpcsResponseBodyVpcsVpcVSwitchs) SetCidrBlock(v string) *DescribeRdsVpcsResponseBodyVpcsVpcVSwitchs {
	s.CidrBlock = &v
	return s
}

func (s *DescribeRdsVpcsResponseBodyVpcsVpcVSwitchs) SetGmtCreate(v string) *DescribeRdsVpcsResponseBodyVpcsVpcVSwitchs {
	s.GmtCreate = &v
	return s
}

func (s *DescribeRdsVpcsResponseBodyVpcsVpcVSwitchs) SetGmtModified(v string) *DescribeRdsVpcsResponseBodyVpcsVpcVSwitchs {
	s.GmtModified = &v
	return s
}

func (s *DescribeRdsVpcsResponseBodyVpcsVpcVSwitchs) SetIsDefault(v bool) *DescribeRdsVpcsResponseBodyVpcsVpcVSwitchs {
	s.IsDefault = &v
	return s
}

func (s *DescribeRdsVpcsResponseBodyVpcsVpcVSwitchs) SetIzNo(v string) *DescribeRdsVpcsResponseBodyVpcsVpcVSwitchs {
	s.IzNo = &v
	return s
}

func (s *DescribeRdsVpcsResponseBodyVpcsVpcVSwitchs) SetStatus(v string) *DescribeRdsVpcsResponseBodyVpcsVpcVSwitchs {
	s.Status = &v
	return s
}

func (s *DescribeRdsVpcsResponseBodyVpcsVpcVSwitchs) SetVSwitchId(v string) *DescribeRdsVpcsResponseBodyVpcsVpcVSwitchs {
	s.VSwitchId = &v
	return s
}

func (s *DescribeRdsVpcsResponseBodyVpcsVpcVSwitchs) SetVSwitchName(v string) *DescribeRdsVpcsResponseBodyVpcsVpcVSwitchs {
	s.VSwitchName = &v
	return s
}

type DescribeRdsVpcsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeRdsVpcsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRdsVpcsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRdsVpcsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRdsVpcsResponse) SetHeaders(v map[string]*string) *DescribeRdsVpcsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRdsVpcsResponse) SetStatusCode(v int32) *DescribeRdsVpcsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRdsVpcsResponse) SetBody(v *DescribeRdsVpcsResponseBody) *DescribeRdsVpcsResponse {
	s.Body = v
	return s
}

type DescribeRegionsRequest struct {
	// The ID of the region.
	//
	// >  You can call the [DescribeRegions](~~86912~~) operation to query the most recent region list.
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) SetRegion(v string) *DescribeRegionsRequest {
	s.Region = &v
	return s
}

type DescribeRegionsResponseBody struct {
	// Details of the regions.
	Regions *DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) SetRegions(v *DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeRegionsResponseBodyRegions struct {
	Region []*DescribeRegionsResponseBodyRegionsRegion `json:"Region,omitempty" xml:"Region,omitempty" type:"Repeated"`
}

func (s DescribeRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegions) SetRegion(v []*DescribeRegionsResponseBodyRegionsRegion) *DescribeRegionsResponseBodyRegions {
	s.Region = v
	return s
}

type DescribeRegionsResponseBodyRegionsRegion struct {
	// The ID of the region.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Details of the zones.
	Zones *DescribeRegionsResponseBodyRegionsRegionZones `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Struct"`
}

func (s DescribeRegionsResponseBodyRegionsRegion) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegionsRegion) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetRegionId(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.RegionId = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetZones(v *DescribeRegionsResponseBodyRegionsRegionZones) *DescribeRegionsResponseBodyRegionsRegion {
	s.Zones = v
	return s
}

type DescribeRegionsResponseBodyRegionsRegionZones struct {
	Zone []*DescribeRegionsResponseBodyRegionsRegionZonesZone `json:"Zone,omitempty" xml:"Zone,omitempty" type:"Repeated"`
}

func (s DescribeRegionsResponseBodyRegionsRegionZones) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegionsRegionZones) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegionsRegionZones) SetZone(v []*DescribeRegionsResponseBodyRegionsRegionZonesZone) *DescribeRegionsResponseBodyRegionsRegionZones {
	s.Zone = v
	return s
}

type DescribeRegionsResponseBodyRegionsRegionZonesZone struct {
	// Indicates whether Virtual Private Cloud (VPC) is available.
	//
	// *   **true**: VPC is available.
	// *   **false**: VPC is unavailable.
	VpcEnabled *bool `json:"VpcEnabled,omitempty" xml:"VpcEnabled,omitempty"`
	// The ID of the zone.
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeRegionsResponseBodyRegionsRegionZonesZone) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegionsRegionZonesZone) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegionsRegionZonesZone) SetVpcEnabled(v bool) *DescribeRegionsResponseBodyRegionsRegionZonesZone {
	s.VpcEnabled = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegionZonesZone) SetZoneId(v string) *DescribeRegionsResponseBodyRegionsRegionZonesZone {
	s.ZoneId = &v
	return s
}

type DescribeRegionsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponse) SetHeaders(v map[string]*string) *DescribeRegionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRegionsResponse) SetStatusCode(v int32) *DescribeRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRegionsResponse) SetBody(v *DescribeRegionsResponseBody) *DescribeRegionsResponse {
	s.Body = v
	return s
}

type DescribeSQLLogCountRequest struct {
	// The ID of the instance.
	//
	// >  You can call the [DescribeDBInstances](~~86911~~) operation to query details of all AnalyticDB for PostgreSQL instances in a specific region, including instance IDs.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The name of the database.
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the `yyyy-MM-ddTHH:mmZ` format. The time must be in UTC.
	//
	// >  The end time must be later than the start time. Their interval cannot be more than seven days.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The execution duration of the query. Unit: seconds.
	ExecuteCost *string `json:"ExecuteCost,omitempty" xml:"ExecuteCost,omitempty"`
	// The execution state of the query. Valid values:
	//
	// *   **success**
	// *   **fail**
	ExecuteState *string `json:"ExecuteState,omitempty" xml:"ExecuteState,omitempty"`
	// The maximum amount of time consumed by a slow query. Minimum value: 0. Unit: seconds.
	MaxExecuteCost *string `json:"MaxExecuteCost,omitempty" xml:"MaxExecuteCost,omitempty"`
	// The minimum amount of time consumed by a slow query. Minimum value: 0. Unit: seconds.
	MinExecuteCost *string `json:"MinExecuteCost,omitempty" xml:"MinExecuteCost,omitempty"`
	// The type of the query language. Valid values:
	//
	// *   **DQL**
	// *   **DML**
	// *   **DDL**
	// *   **DCL**
	// *   **TCL**
	OperationClass *string `json:"OperationClass,omitempty" xml:"OperationClass,omitempty"`
	// The type of the SQL statement.
	//
	// > *   If the **OperationClass** parameter is specified, the **OperationType** value must belong to the corresponding query language. For example, if the **OperationClass** value is **DQL**, the **OperationType** value must be a **DQL** SQL statement such as **SELECT**.
	// >*   If the **OperationClass** parameter is not specified, the **OperationType** value can be an SQL statement of all query languages.
	// >*   If neither of the **OperationClass** and **OperationType** parameters is specified, all types of SQL statements are returned.
	OperationType *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	// The keywords used to query.
	QueryKeywords *string `json:"QueryKeywords,omitempty" xml:"QueryKeywords,omitempty"`
	// The source IP address.
	SourceIP *string `json:"SourceIP,omitempty" xml:"SourceIP,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the `yyyy-MM-ddTHH:mmZ` format. The time must be in UTC.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The username that is used to log on to the database.
	User *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DescribeSQLLogCountRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLLogCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogCountRequest) SetDBInstanceId(v string) *DescribeSQLLogCountRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeSQLLogCountRequest) SetDatabase(v string) *DescribeSQLLogCountRequest {
	s.Database = &v
	return s
}

func (s *DescribeSQLLogCountRequest) SetEndTime(v string) *DescribeSQLLogCountRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSQLLogCountRequest) SetExecuteCost(v string) *DescribeSQLLogCountRequest {
	s.ExecuteCost = &v
	return s
}

func (s *DescribeSQLLogCountRequest) SetExecuteState(v string) *DescribeSQLLogCountRequest {
	s.ExecuteState = &v
	return s
}

func (s *DescribeSQLLogCountRequest) SetMaxExecuteCost(v string) *DescribeSQLLogCountRequest {
	s.MaxExecuteCost = &v
	return s
}

func (s *DescribeSQLLogCountRequest) SetMinExecuteCost(v string) *DescribeSQLLogCountRequest {
	s.MinExecuteCost = &v
	return s
}

func (s *DescribeSQLLogCountRequest) SetOperationClass(v string) *DescribeSQLLogCountRequest {
	s.OperationClass = &v
	return s
}

func (s *DescribeSQLLogCountRequest) SetOperationType(v string) *DescribeSQLLogCountRequest {
	s.OperationType = &v
	return s
}

func (s *DescribeSQLLogCountRequest) SetQueryKeywords(v string) *DescribeSQLLogCountRequest {
	s.QueryKeywords = &v
	return s
}

func (s *DescribeSQLLogCountRequest) SetSourceIP(v string) *DescribeSQLLogCountRequest {
	s.SourceIP = &v
	return s
}

func (s *DescribeSQLLogCountRequest) SetStartTime(v string) *DescribeSQLLogCountRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSQLLogCountRequest) SetUser(v string) *DescribeSQLLogCountRequest {
	s.User = &v
	return s
}

type DescribeSQLLogCountResponseBody struct {
	// The ID of the instance.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The end time of the query.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Details of the audit logs of the instance.
	Items []*DescribeSQLLogCountResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The start time of the query.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeSQLLogCountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLLogCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogCountResponseBody) SetDBClusterId(v string) *DescribeSQLLogCountResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeSQLLogCountResponseBody) SetEndTime(v string) *DescribeSQLLogCountResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeSQLLogCountResponseBody) SetItems(v []*DescribeSQLLogCountResponseBodyItems) *DescribeSQLLogCountResponseBody {
	s.Items = v
	return s
}

func (s *DescribeSQLLogCountResponseBody) SetRequestId(v string) *DescribeSQLLogCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSQLLogCountResponseBody) SetStartTime(v string) *DescribeSQLLogCountResponseBody {
	s.StartTime = &v
	return s
}

type DescribeSQLLogCountResponseBodyItems struct {
	// The name of the table.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Details of the audit logs.
	Series []*DescribeSQLLogCountResponseBodyItemsSeries `json:"Series,omitempty" xml:"Series,omitempty" type:"Repeated"`
}

func (s DescribeSQLLogCountResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLLogCountResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogCountResponseBodyItems) SetName(v string) *DescribeSQLLogCountResponseBodyItems {
	s.Name = &v
	return s
}

func (s *DescribeSQLLogCountResponseBodyItems) SetSeries(v []*DescribeSQLLogCountResponseBodyItemsSeries) *DescribeSQLLogCountResponseBodyItems {
	s.Series = v
	return s
}

type DescribeSQLLogCountResponseBodyItemsSeries struct {
	// Details of the audit logs.
	Values []*DescribeSQLLogCountResponseBodyItemsSeriesValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s DescribeSQLLogCountResponseBodyItemsSeries) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLLogCountResponseBodyItemsSeries) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogCountResponseBodyItemsSeries) SetValues(v []*DescribeSQLLogCountResponseBodyItemsSeriesValues) *DescribeSQLLogCountResponseBodyItemsSeries {
	s.Values = v
	return s
}

type DescribeSQLLogCountResponseBodyItemsSeriesValues struct {
	// The time when the audit logs were generated and the number of the audit logs.
	Point []*string `json:"Point,omitempty" xml:"Point,omitempty" type:"Repeated"`
}

func (s DescribeSQLLogCountResponseBodyItemsSeriesValues) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLLogCountResponseBodyItemsSeriesValues) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogCountResponseBodyItemsSeriesValues) SetPoint(v []*string) *DescribeSQLLogCountResponseBodyItemsSeriesValues {
	s.Point = v
	return s
}

type DescribeSQLLogCountResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeSQLLogCountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSQLLogCountResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLLogCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogCountResponse) SetHeaders(v map[string]*string) *DescribeSQLLogCountResponse {
	s.Headers = v
	return s
}

func (s *DescribeSQLLogCountResponse) SetStatusCode(v int32) *DescribeSQLLogCountResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSQLLogCountResponse) SetBody(v *DescribeSQLLogCountResponseBody) *DescribeSQLLogCountResponse {
	s.Body = v
	return s
}

type DescribeSQLLogsV2Request struct {
	// The ID of instance.
	//
	// >  You can call the [DescribeDBInstances](~~86911~~) operation to query details of all AnalyticDB for PostgreSQL instances in a specific region, including instance IDs.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The name of the database.
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-ddTHH:mmZ* format. The time must be in UTC.
	//
	// >  The end time must be later than the start time. The interval cannot be more than 24 hours.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The execution duration of the query. Unit: seconds.
	ExecuteCost *string `json:"ExecuteCost,omitempty" xml:"ExecuteCost,omitempty"`
	// The execution state of the query. Valid values:
	//
	// *   **success**
	// *   **fail**
	ExecuteState *string `json:"ExecuteState,omitempty" xml:"ExecuteState,omitempty"`
	// The maximum amount of time consumed by a slow query. Minimum value: 0. Unit: seconds.
	MaxExecuteCost *string `json:"MaxExecuteCost,omitempty" xml:"MaxExecuteCost,omitempty"`
	// The minimum amount of time consumed by a slow query. Minimum value: 0. Unit: seconds.
	MinExecuteCost *string `json:"MinExecuteCost,omitempty" xml:"MinExecuteCost,omitempty"`
	// The type of the query language. Valid values:
	//
	// *   **DQL**
	// *   **DML**
	// *   **DDL**
	// *   **DCL**
	// *   **TCL**
	OperationClass *string `json:"OperationClass,omitempty" xml:"OperationClass,omitempty"`
	// The type of the SQL statement.
	//
	// > *   If the **OperationClass** parameter is specified, the **OperationType** value must belong to the corresponding query language. For example, if the **OperationClass** value is **DQL**, the **OperationType** value must be a **DQL** SQL statement such as **SELECT**.
	// >*   If the **OperationClass** parameter is not specified, the **OperationType** value can be an SQL statement of all query languages.
	// >*   If neither of the **OperationClass** and **OperationType** parameters is specified, all types of SQL statements are returned.
	OperationType *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	// The number of entries to return on each page.
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of the page to return. The maximum value is 200.
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The keywords of the SQL statement.
	QueryKeywords *string `json:"QueryKeywords,omitempty" xml:"QueryKeywords,omitempty"`
	// The region ID of the instance.
	//
	// >  You can call the [DescribeRegions](~~86912~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the instance belongs.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The source IP address.
	SourceIP *string `json:"SourceIP,omitempty" xml:"SourceIP,omitempty"`
	// The beginning of the time range. Specify the time in the ISO 8601 standard in the *yyyy-MM-ddTHH:mmZ* format. The time must be in UTC.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The name of the database account.
	User *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DescribeSQLLogsV2Request) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLLogsV2Request) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogsV2Request) SetDBInstanceId(v string) *DescribeSQLLogsV2Request {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeSQLLogsV2Request) SetDatabase(v string) *DescribeSQLLogsV2Request {
	s.Database = &v
	return s
}

func (s *DescribeSQLLogsV2Request) SetEndTime(v string) *DescribeSQLLogsV2Request {
	s.EndTime = &v
	return s
}

func (s *DescribeSQLLogsV2Request) SetExecuteCost(v string) *DescribeSQLLogsV2Request {
	s.ExecuteCost = &v
	return s
}

func (s *DescribeSQLLogsV2Request) SetExecuteState(v string) *DescribeSQLLogsV2Request {
	s.ExecuteState = &v
	return s
}

func (s *DescribeSQLLogsV2Request) SetMaxExecuteCost(v string) *DescribeSQLLogsV2Request {
	s.MaxExecuteCost = &v
	return s
}

func (s *DescribeSQLLogsV2Request) SetMinExecuteCost(v string) *DescribeSQLLogsV2Request {
	s.MinExecuteCost = &v
	return s
}

func (s *DescribeSQLLogsV2Request) SetOperationClass(v string) *DescribeSQLLogsV2Request {
	s.OperationClass = &v
	return s
}

func (s *DescribeSQLLogsV2Request) SetOperationType(v string) *DescribeSQLLogsV2Request {
	s.OperationType = &v
	return s
}

func (s *DescribeSQLLogsV2Request) SetPageNumber(v string) *DescribeSQLLogsV2Request {
	s.PageNumber = &v
	return s
}

func (s *DescribeSQLLogsV2Request) SetPageSize(v string) *DescribeSQLLogsV2Request {
	s.PageSize = &v
	return s
}

func (s *DescribeSQLLogsV2Request) SetQueryKeywords(v string) *DescribeSQLLogsV2Request {
	s.QueryKeywords = &v
	return s
}

func (s *DescribeSQLLogsV2Request) SetRegionId(v string) *DescribeSQLLogsV2Request {
	s.RegionId = &v
	return s
}

func (s *DescribeSQLLogsV2Request) SetResourceGroupId(v string) *DescribeSQLLogsV2Request {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeSQLLogsV2Request) SetSourceIP(v string) *DescribeSQLLogsV2Request {
	s.SourceIP = &v
	return s
}

func (s *DescribeSQLLogsV2Request) SetStartTime(v string) *DescribeSQLLogsV2Request {
	s.StartTime = &v
	return s
}

func (s *DescribeSQLLogsV2Request) SetUser(v string) *DescribeSQLLogsV2Request {
	s.User = &v
	return s
}

type DescribeSQLLogsV2ResponseBody struct {
	// Details of the SQL logs.
	Items []*DescribeSQLLogsV2ResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The page number of the returned page.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeSQLLogsV2ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLLogsV2ResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogsV2ResponseBody) SetItems(v []*DescribeSQLLogsV2ResponseBodyItems) *DescribeSQLLogsV2ResponseBody {
	s.Items = v
	return s
}

func (s *DescribeSQLLogsV2ResponseBody) SetPageNumber(v int32) *DescribeSQLLogsV2ResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSQLLogsV2ResponseBody) SetPageRecordCount(v int32) *DescribeSQLLogsV2ResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeSQLLogsV2ResponseBody) SetRequestId(v string) *DescribeSQLLogsV2ResponseBody {
	s.RequestId = &v
	return s
}

type DescribeSQLLogsV2ResponseBodyItems struct {
	// The database account that executes the SQL statement.
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The name of the database.
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// The role of the database.
	DBRole *string `json:"DBRole,omitempty" xml:"DBRole,omitempty"`
	// The execution duration of the query.
	ExecuteCost *float32 `json:"ExecuteCost,omitempty" xml:"ExecuteCost,omitempty"`
	// The execution state of the query. Valid values:
	//
	// *   **success**
	// *   **fail**
	ExecuteState *string `json:"ExecuteState,omitempty" xml:"ExecuteState,omitempty"`
	// The type of the query language.
	OperationClass *string `json:"OperationClass,omitempty" xml:"OperationClass,omitempty"`
	// The time when the SQL statement was executed.
	OperationExecuteTime *string `json:"OperationExecuteTime,omitempty" xml:"OperationExecuteTime,omitempty"`
	// The type of the SQL statement.
	OperationType *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	// The number of entries returned.
	ReturnRowCounts *int64 `json:"ReturnRowCounts,omitempty" xml:"ReturnRowCounts,omitempty"`
	// The SQL statement.
	SQLText *string `json:"SQLText,omitempty" xml:"SQLText,omitempty"`
	// The number of entries scanned.
	ScanRowCounts *int64 `json:"ScanRowCounts,omitempty" xml:"ScanRowCounts,omitempty"`
	// The source IP address.
	SourceIP *string `json:"SourceIP,omitempty" xml:"SourceIP,omitempty"`
	// The number of the source port.
	SourcePort *int32 `json:"SourcePort,omitempty" xml:"SourcePort,omitempty"`
}

func (s DescribeSQLLogsV2ResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLLogsV2ResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogsV2ResponseBodyItems) SetAccountName(v string) *DescribeSQLLogsV2ResponseBodyItems {
	s.AccountName = &v
	return s
}

func (s *DescribeSQLLogsV2ResponseBodyItems) SetDBName(v string) *DescribeSQLLogsV2ResponseBodyItems {
	s.DBName = &v
	return s
}

func (s *DescribeSQLLogsV2ResponseBodyItems) SetDBRole(v string) *DescribeSQLLogsV2ResponseBodyItems {
	s.DBRole = &v
	return s
}

func (s *DescribeSQLLogsV2ResponseBodyItems) SetExecuteCost(v float32) *DescribeSQLLogsV2ResponseBodyItems {
	s.ExecuteCost = &v
	return s
}

func (s *DescribeSQLLogsV2ResponseBodyItems) SetExecuteState(v string) *DescribeSQLLogsV2ResponseBodyItems {
	s.ExecuteState = &v
	return s
}

func (s *DescribeSQLLogsV2ResponseBodyItems) SetOperationClass(v string) *DescribeSQLLogsV2ResponseBodyItems {
	s.OperationClass = &v
	return s
}

func (s *DescribeSQLLogsV2ResponseBodyItems) SetOperationExecuteTime(v string) *DescribeSQLLogsV2ResponseBodyItems {
	s.OperationExecuteTime = &v
	return s
}

func (s *DescribeSQLLogsV2ResponseBodyItems) SetOperationType(v string) *DescribeSQLLogsV2ResponseBodyItems {
	s.OperationType = &v
	return s
}

func (s *DescribeSQLLogsV2ResponseBodyItems) SetReturnRowCounts(v int64) *DescribeSQLLogsV2ResponseBodyItems {
	s.ReturnRowCounts = &v
	return s
}

func (s *DescribeSQLLogsV2ResponseBodyItems) SetSQLText(v string) *DescribeSQLLogsV2ResponseBodyItems {
	s.SQLText = &v
	return s
}

func (s *DescribeSQLLogsV2ResponseBodyItems) SetScanRowCounts(v int64) *DescribeSQLLogsV2ResponseBodyItems {
	s.ScanRowCounts = &v
	return s
}

func (s *DescribeSQLLogsV2ResponseBodyItems) SetSourceIP(v string) *DescribeSQLLogsV2ResponseBodyItems {
	s.SourceIP = &v
	return s
}

func (s *DescribeSQLLogsV2ResponseBodyItems) SetSourcePort(v int32) *DescribeSQLLogsV2ResponseBodyItems {
	s.SourcePort = &v
	return s
}

type DescribeSQLLogsV2Response struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeSQLLogsV2ResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSQLLogsV2Response) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLLogsV2Response) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogsV2Response) SetHeaders(v map[string]*string) *DescribeSQLLogsV2Response {
	s.Headers = v
	return s
}

func (s *DescribeSQLLogsV2Response) SetStatusCode(v int32) *DescribeSQLLogsV2Response {
	s.StatusCode = &v
	return s
}

func (s *DescribeSQLLogsV2Response) SetBody(v *DescribeSQLLogsV2ResponseBody) *DescribeSQLLogsV2Response {
	s.Body = v
	return s
}

type DescribeSampleDataRequest struct {
	// The ID of the instance.
	//
	// >  You can call the [DescribeDBInstances](~~86911~~) operation to query the IDs of all AnalyticDB for PostgreSQL instances in a specific region.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeSampleDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSampleDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeSampleDataRequest) SetDBInstanceId(v string) *DescribeSampleDataRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeSampleDataRequest) SetOwnerId(v int64) *DescribeSampleDataRequest {
	s.OwnerId = &v
	return s
}

type DescribeSampleDataResponseBody struct {
	// The ID of the instance.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The error message returned if an error occurs. This message does not affect the execution of the operation.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// Indicates whether a sample dataset is loaded to the instance. Valid values:
	//
	// *   **true**: A sample dataset is loaded.
	// *   **false**: No sample dataset is loaded.
	HasSampleData *bool `json:"HasSampleData,omitempty" xml:"HasSampleData,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeSampleDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSampleDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSampleDataResponseBody) SetDBInstanceId(v string) *DescribeSampleDataResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeSampleDataResponseBody) SetErrorMessage(v string) *DescribeSampleDataResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeSampleDataResponseBody) SetHasSampleData(v bool) *DescribeSampleDataResponseBody {
	s.HasSampleData = &v
	return s
}

func (s *DescribeSampleDataResponseBody) SetRequestId(v string) *DescribeSampleDataResponseBody {
	s.RequestId = &v
	return s
}

type DescribeSampleDataResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeSampleDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSampleDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSampleDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeSampleDataResponse) SetHeaders(v map[string]*string) *DescribeSampleDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeSampleDataResponse) SetStatusCode(v int32) *DescribeSampleDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSampleDataResponse) SetBody(v *DescribeSampleDataResponseBody) *DescribeSampleDataResponse {
	s.Body = v
	return s
}

type DescribeSupportFeaturesRequest struct {
	// The ID of the instance.
	//
	// >  You can call the [DescribeDBInstances](~~86911~~) operation to query the instance IDs of all AnalyticDB for PostgreSQL instances in a specific region.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeSupportFeaturesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSupportFeaturesRequest) GoString() string {
	return s.String()
}

func (s *DescribeSupportFeaturesRequest) SetDBInstanceId(v string) *DescribeSupportFeaturesRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeSupportFeaturesRequest) SetOwnerId(v int64) *DescribeSupportFeaturesRequest {
	s.OwnerId = &v
	return s
}

type DescribeSupportFeaturesResponseBody struct {
	// The ID of the instance.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The features supported by the instance. Valid values:
	//
	// *   sample_data: sample dataset. For more information, see [Sample dataset](~~452278~~).
	// *   diagnose_and_optimize: diagnostics and optimization. For more information, see [Diagnostics and optimization](~~323453~~).
	SupportFeatureList *string `json:"SupportFeatureList,omitempty" xml:"SupportFeatureList,omitempty"`
}

func (s DescribeSupportFeaturesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSupportFeaturesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSupportFeaturesResponseBody) SetDBInstanceId(v string) *DescribeSupportFeaturesResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeSupportFeaturesResponseBody) SetRequestId(v string) *DescribeSupportFeaturesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSupportFeaturesResponseBody) SetSupportFeatureList(v string) *DescribeSupportFeaturesResponseBody {
	s.SupportFeatureList = &v
	return s
}

type DescribeSupportFeaturesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeSupportFeaturesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSupportFeaturesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSupportFeaturesResponse) GoString() string {
	return s.String()
}

func (s *DescribeSupportFeaturesResponse) SetHeaders(v map[string]*string) *DescribeSupportFeaturesResponse {
	s.Headers = v
	return s
}

func (s *DescribeSupportFeaturesResponse) SetStatusCode(v int32) *DescribeSupportFeaturesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSupportFeaturesResponse) SetBody(v *DescribeSupportFeaturesResponseBody) *DescribeSupportFeaturesResponse {
	s.Body = v
	return s
}

type DescribeTagsRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region.
	//
	// >  You can call the [DescribeRegions](~~86912~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the instance belongs. For more information about how to obtain the ID of a resource group, see [View basic information of a resource group](~~151181~~).
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The type of the resource. Set the value to **instance**.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s DescribeTagsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTagsRequest) GoString() string {
	return s.String()
}

func (s *DescribeTagsRequest) SetOwnerAccount(v string) *DescribeTagsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeTagsRequest) SetOwnerId(v int64) *DescribeTagsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeTagsRequest) SetRegionId(v string) *DescribeTagsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeTagsRequest) SetResourceGroupId(v string) *DescribeTagsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeTagsRequest) SetResourceOwnerAccount(v string) *DescribeTagsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeTagsRequest) SetResourceOwnerId(v int64) *DescribeTagsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeTagsRequest) SetResourceType(v string) *DescribeTagsRequest {
	s.ResourceType = &v
	return s
}

type DescribeTagsResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Details of the tags.
	Tags []*DescribeTagsResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s DescribeTagsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTagsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTagsResponseBody) SetRequestId(v string) *DescribeTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTagsResponseBody) SetTags(v []*DescribeTagsResponseBodyTags) *DescribeTagsResponseBody {
	s.Tags = v
	return s
}

type DescribeTagsResponseBodyTags struct {
	// The key of the tag.
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The value of the tag.
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeTagsResponseBodyTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeTagsResponseBodyTags) GoString() string {
	return s.String()
}

func (s *DescribeTagsResponseBodyTags) SetTagKey(v string) *DescribeTagsResponseBodyTags {
	s.TagKey = &v
	return s
}

func (s *DescribeTagsResponseBodyTags) SetTagValue(v string) *DescribeTagsResponseBodyTags {
	s.TagValue = &v
	return s
}

type DescribeTagsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeTagsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeTagsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTagsResponse) GoString() string {
	return s.String()
}

func (s *DescribeTagsResponse) SetHeaders(v map[string]*string) *DescribeTagsResponse {
	s.Headers = v
	return s
}

func (s *DescribeTagsResponse) SetStatusCode(v int32) *DescribeTagsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTagsResponse) SetBody(v *DescribeTagsResponseBody) *DescribeTagsResponse {
	s.Body = v
	return s
}

type DescribeUserEncryptionKeyListRequest struct {
	// The number of the page to return. Default value: 1.
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of KMS keys to return on each page. Default value: 10.
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeUserEncryptionKeyListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserEncryptionKeyListRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserEncryptionKeyListRequest) SetPageNumber(v string) *DescribeUserEncryptionKeyListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeUserEncryptionKeyListRequest) SetPageSize(v string) *DescribeUserEncryptionKeyListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeUserEncryptionKeyListRequest) SetRegionId(v string) *DescribeUserEncryptionKeyListRequest {
	s.RegionId = &v
	return s
}

type DescribeUserEncryptionKeyListResponseBody struct {
	// Details about the KMS keys.
	KmsKeys []*DescribeUserEncryptionKeyListResponseBodyKmsKeys `json:"KmsKeys,omitempty" xml:"KmsKeys,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeUserEncryptionKeyListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserEncryptionKeyListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserEncryptionKeyListResponseBody) SetKmsKeys(v []*DescribeUserEncryptionKeyListResponseBodyKmsKeys) *DescribeUserEncryptionKeyListResponseBody {
	s.KmsKeys = v
	return s
}

func (s *DescribeUserEncryptionKeyListResponseBody) SetRequestId(v string) *DescribeUserEncryptionKeyListResponseBody {
	s.RequestId = &v
	return s
}

type DescribeUserEncryptionKeyListResponseBodyKmsKeys struct {
	// The ID of the KMS key.
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
}

func (s DescribeUserEncryptionKeyListResponseBodyKmsKeys) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserEncryptionKeyListResponseBodyKmsKeys) GoString() string {
	return s.String()
}

func (s *DescribeUserEncryptionKeyListResponseBodyKmsKeys) SetKeyId(v string) *DescribeUserEncryptionKeyListResponseBodyKmsKeys {
	s.KeyId = &v
	return s
}

type DescribeUserEncryptionKeyListResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeUserEncryptionKeyListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeUserEncryptionKeyListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserEncryptionKeyListResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserEncryptionKeyListResponse) SetHeaders(v map[string]*string) *DescribeUserEncryptionKeyListResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserEncryptionKeyListResponse) SetStatusCode(v int32) *DescribeUserEncryptionKeyListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUserEncryptionKeyListResponse) SetBody(v *DescribeUserEncryptionKeyListResponseBody) *DescribeUserEncryptionKeyListResponse {
	s.Body = v
	return s
}

type DescribeWaitingSQLInfoRequest struct {
	// The ID of the instance.
	//
	// >  You can call the [DescribeDBInstances](~~86911~~) operation to query the instance IDs of all AnalyticDB for PostgreSQL instances in a specific region.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The name of the database.
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	// The ID of the process that uniquely identifies the query.
	//
	// >  You can call the [DescribeWaitingSQLRecords](~~461735~~) operation to obtain the process IDs of lock-waiting queries.
	PID *string `json:"PID,omitempty" xml:"PID,omitempty"`
}

func (s DescribeWaitingSQLInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeWaitingSQLInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeWaitingSQLInfoRequest) SetDBInstanceId(v string) *DescribeWaitingSQLInfoRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeWaitingSQLInfoRequest) SetDatabase(v string) *DescribeWaitingSQLInfoRequest {
	s.Database = &v
	return s
}

func (s *DescribeWaitingSQLInfoRequest) SetPID(v string) *DescribeWaitingSQLInfoRequest {
	s.PID = &v
	return s
}

type DescribeWaitingSQLInfoResponseBody struct {
	// The name of the database.
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	// Details of the lock-waiting query.
	Items []*DescribeWaitingSQLInfoResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeWaitingSQLInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeWaitingSQLInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWaitingSQLInfoResponseBody) SetDatabase(v string) *DescribeWaitingSQLInfoResponseBody {
	s.Database = &v
	return s
}

func (s *DescribeWaitingSQLInfoResponseBody) SetItems(v []*DescribeWaitingSQLInfoResponseBodyItems) *DescribeWaitingSQLInfoResponseBody {
	s.Items = v
	return s
}

func (s *DescribeWaitingSQLInfoResponseBody) SetRequestId(v string) *DescribeWaitingSQLInfoResponseBody {
	s.RequestId = &v
	return s
}

type DescribeWaitingSQLInfoResponseBodyItems struct {
	// The application that sent the query.
	Application *string `json:"Application,omitempty" xml:"Application,omitempty"`
	// The application that sent the blocking query.
	BlockedByApplication *string `json:"BlockedByApplication,omitempty" xml:"BlockedByApplication,omitempty"`
	// The process ID of the blocking query.
	BlockedByPID *string `json:"BlockedByPID,omitempty" xml:"BlockedByPID,omitempty"`
	// The SQL statement of the blocking query.
	BlockedBySQLStmt *string `json:"BlockedBySQLStmt,omitempty" xml:"BlockedBySQLStmt,omitempty"`
	// The database account that is used to perform the blocking query.
	BlockedByUser *string `json:"BlockedByUser,omitempty" xml:"BlockedByUser,omitempty"`
	// The authorized locks.
	GrantLocks *string `json:"GrantLocks,omitempty" xml:"GrantLocks,omitempty"`
	// The unauthorized locks.
	NotGrantLocks *string `json:"NotGrantLocks,omitempty" xml:"NotGrantLocks,omitempty"`
	// The ID of the process that uniquely identifies the query.
	PID *string `json:"PID,omitempty" xml:"PID,omitempty"`
	// The SQL statement of the query.
	SQLStmt *string `json:"SQLStmt,omitempty" xml:"SQLStmt,omitempty"`
	// The database account that is used to perform the query.
	User *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DescribeWaitingSQLInfoResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeWaitingSQLInfoResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeWaitingSQLInfoResponseBodyItems) SetApplication(v string) *DescribeWaitingSQLInfoResponseBodyItems {
	s.Application = &v
	return s
}

func (s *DescribeWaitingSQLInfoResponseBodyItems) SetBlockedByApplication(v string) *DescribeWaitingSQLInfoResponseBodyItems {
	s.BlockedByApplication = &v
	return s
}

func (s *DescribeWaitingSQLInfoResponseBodyItems) SetBlockedByPID(v string) *DescribeWaitingSQLInfoResponseBodyItems {
	s.BlockedByPID = &v
	return s
}

func (s *DescribeWaitingSQLInfoResponseBodyItems) SetBlockedBySQLStmt(v string) *DescribeWaitingSQLInfoResponseBodyItems {
	s.BlockedBySQLStmt = &v
	return s
}

func (s *DescribeWaitingSQLInfoResponseBodyItems) SetBlockedByUser(v string) *DescribeWaitingSQLInfoResponseBodyItems {
	s.BlockedByUser = &v
	return s
}

func (s *DescribeWaitingSQLInfoResponseBodyItems) SetGrantLocks(v string) *DescribeWaitingSQLInfoResponseBodyItems {
	s.GrantLocks = &v
	return s
}

func (s *DescribeWaitingSQLInfoResponseBodyItems) SetNotGrantLocks(v string) *DescribeWaitingSQLInfoResponseBodyItems {
	s.NotGrantLocks = &v
	return s
}

func (s *DescribeWaitingSQLInfoResponseBodyItems) SetPID(v string) *DescribeWaitingSQLInfoResponseBodyItems {
	s.PID = &v
	return s
}

func (s *DescribeWaitingSQLInfoResponseBodyItems) SetSQLStmt(v string) *DescribeWaitingSQLInfoResponseBodyItems {
	s.SQLStmt = &v
	return s
}

func (s *DescribeWaitingSQLInfoResponseBodyItems) SetUser(v string) *DescribeWaitingSQLInfoResponseBodyItems {
	s.User = &v
	return s
}

type DescribeWaitingSQLInfoResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeWaitingSQLInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeWaitingSQLInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeWaitingSQLInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeWaitingSQLInfoResponse) SetHeaders(v map[string]*string) *DescribeWaitingSQLInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeWaitingSQLInfoResponse) SetStatusCode(v int32) *DescribeWaitingSQLInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeWaitingSQLInfoResponse) SetBody(v *DescribeWaitingSQLInfoResponseBody) *DescribeWaitingSQLInfoResponse {
	s.Body = v
	return s
}

type DescribeWaitingSQLRecordsRequest struct {
	// The ID of the instance.
	//
	// >  You can call the [DescribeDBInstances](~~86911~~) operation to query the instance IDs of all AnalyticDB for PostgreSQL instances in a specific region.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The name of the database.
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm*Z format. The time must be in UTC. The end time must be later than the start time.
	//
	// If this parameter is not specified, all lock diagnostics records that are generated after the query start time are returned. If the query start time is not specified either, all lock diagnostics records are returned.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The keyword used to filter queries.
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The field used to sort lock diagnostics records and the sorting order.
	//
	// Default value: `{"Field":"StartTime","Type":"Desc"}`, which indicates that lock diagnostics records are sorted by the start time in descending order. No other values are supported.
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The number of the page to return. The value must be an integer that is greater than 0. Default value: **1**.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values:
	//
	// *   **30**
	// *   **50**
	// *   **100**
	//
	// Default value: **30**.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The filter condition on queries. Valid values:
	//
	// *   `{"Type":"maxCost","Value":"10"}`: filters the top 10 longest-waiting queries.
	// *   `{"Type":"status","Value":"LockWaiting"}`: filters lock-waiting queries.
	// *   `{"Type":"status","Value":"ResourceWaiting"}`: filters resource-waiting queries.
	QueryCondition *string `json:"QueryCondition,omitempty" xml:"QueryCondition,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm*Z format. The time must be in UTC.
	//
	// If this parameter is not specified, all lock diagnostics records that are generated before the query end time are returned. If the query end time is not specified either, all lock diagnostics records are returned.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The name of the database account. If this parameter is not specified, the lock diagnostics records of all database accounts are queried.
	User *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DescribeWaitingSQLRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeWaitingSQLRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeWaitingSQLRecordsRequest) SetDBInstanceId(v string) *DescribeWaitingSQLRecordsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeWaitingSQLRecordsRequest) SetDatabase(v string) *DescribeWaitingSQLRecordsRequest {
	s.Database = &v
	return s
}

func (s *DescribeWaitingSQLRecordsRequest) SetEndTime(v string) *DescribeWaitingSQLRecordsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeWaitingSQLRecordsRequest) SetKeyword(v string) *DescribeWaitingSQLRecordsRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeWaitingSQLRecordsRequest) SetOrder(v string) *DescribeWaitingSQLRecordsRequest {
	s.Order = &v
	return s
}

func (s *DescribeWaitingSQLRecordsRequest) SetPageNumber(v int32) *DescribeWaitingSQLRecordsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeWaitingSQLRecordsRequest) SetPageSize(v int32) *DescribeWaitingSQLRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeWaitingSQLRecordsRequest) SetQueryCondition(v string) *DescribeWaitingSQLRecordsRequest {
	s.QueryCondition = &v
	return s
}

func (s *DescribeWaitingSQLRecordsRequest) SetStartTime(v string) *DescribeWaitingSQLRecordsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeWaitingSQLRecordsRequest) SetUser(v string) *DescribeWaitingSQLRecordsRequest {
	s.User = &v
	return s
}

type DescribeWaitingSQLRecordsResponseBody struct {
	// The list of lock diagnostics records.
	Items []*DescribeWaitingSQLRecordsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The page number of the returned page.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeWaitingSQLRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeWaitingSQLRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWaitingSQLRecordsResponseBody) SetItems(v []*DescribeWaitingSQLRecordsResponseBodyItems) *DescribeWaitingSQLRecordsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeWaitingSQLRecordsResponseBody) SetPageNumber(v int32) *DescribeWaitingSQLRecordsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeWaitingSQLRecordsResponseBody) SetRequestId(v string) *DescribeWaitingSQLRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeWaitingSQLRecordsResponseBody) SetTotalCount(v int32) *DescribeWaitingSQLRecordsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeWaitingSQLRecordsResponseBodyItems struct {
	// The name of the database.
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	// The ID of the process that uniquely identifies the query.
	PID *string `json:"PID,omitempty" xml:"PID,omitempty"`
	// The SQL statement of the query.
	SQLStmt *string `json:"SQLStmt,omitempty" xml:"SQLStmt,omitempty"`
	// The ID of the session that contains the query.
	SessionID *string `json:"SessionID,omitempty" xml:"SessionID,omitempty"`
	// The start time of the query. This value is in the timestamp format. Unit: milliseconds.
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The waiting state of the query. Valid values:
	//
	// *   **LockWaiting**
	// *   **ResourceWaiting**
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The database account that is used to perform the query.
	User *string `json:"User,omitempty" xml:"User,omitempty"`
	// The waiting period of the query. Unit: milliseconds.
	WaitingTime *int64 `json:"WaitingTime,omitempty" xml:"WaitingTime,omitempty"`
}

func (s DescribeWaitingSQLRecordsResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeWaitingSQLRecordsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeWaitingSQLRecordsResponseBodyItems) SetDatabase(v string) *DescribeWaitingSQLRecordsResponseBodyItems {
	s.Database = &v
	return s
}

func (s *DescribeWaitingSQLRecordsResponseBodyItems) SetPID(v string) *DescribeWaitingSQLRecordsResponseBodyItems {
	s.PID = &v
	return s
}

func (s *DescribeWaitingSQLRecordsResponseBodyItems) SetSQLStmt(v string) *DescribeWaitingSQLRecordsResponseBodyItems {
	s.SQLStmt = &v
	return s
}

func (s *DescribeWaitingSQLRecordsResponseBodyItems) SetSessionID(v string) *DescribeWaitingSQLRecordsResponseBodyItems {
	s.SessionID = &v
	return s
}

func (s *DescribeWaitingSQLRecordsResponseBodyItems) SetStartTime(v int64) *DescribeWaitingSQLRecordsResponseBodyItems {
	s.StartTime = &v
	return s
}

func (s *DescribeWaitingSQLRecordsResponseBodyItems) SetStatus(v string) *DescribeWaitingSQLRecordsResponseBodyItems {
	s.Status = &v
	return s
}

func (s *DescribeWaitingSQLRecordsResponseBodyItems) SetUser(v string) *DescribeWaitingSQLRecordsResponseBodyItems {
	s.User = &v
	return s
}

func (s *DescribeWaitingSQLRecordsResponseBodyItems) SetWaitingTime(v int64) *DescribeWaitingSQLRecordsResponseBodyItems {
	s.WaitingTime = &v
	return s
}

type DescribeWaitingSQLRecordsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeWaitingSQLRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeWaitingSQLRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeWaitingSQLRecordsResponse) GoString() string {
	return s.String()
}

func (s *DescribeWaitingSQLRecordsResponse) SetHeaders(v map[string]*string) *DescribeWaitingSQLRecordsResponse {
	s.Headers = v
	return s
}

func (s *DescribeWaitingSQLRecordsResponse) SetStatusCode(v int32) *DescribeWaitingSQLRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeWaitingSQLRecordsResponse) SetBody(v *DescribeWaitingSQLRecordsResponseBody) *DescribeWaitingSQLRecordsResponse {
	s.Body = v
	return s
}

type DownloadDiagnosisRecordsRequest struct {
	// The ID of the instance.
	//
	// >  You can call the [DescribeDBInstances](~~86911~~) operation to query the details of all AnalyticDB for PostgreSQL instances in a specific region, including instance IDs.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The name of the database.
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	// The end of the time range to query. Specify the time in the *yyyy-MM-dd*T*HH:mm*Z format. The time must be in UTC.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The language of the file that contains the query diagnostic information. Valid values:
	//
	// *   **zh**: simplified Chinese
	// *   **en**: English
	// *   **ja**: Japanese
	// *   **zh-tw**: traditional Chinese
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The filter condition on queries. The value is in the JSON format. Valid values:
	//
	// *   `{"Type":"maxCost", "Value":"100"}`: filters the top 100 queries that are the most time-consuming.
	// *   `{"Type":"status","Value":"finished"}`: filters completed queries.
	// *   `{"Type":"status","Value":"running"}`: filters running queries.
	// *   `{"Type":"cost","Max":"200"}`: filters the queries that consume less than 200 milliseconds.
	// *   `{"Type":"cost","Min":"200","Max":"60000"}`: filters the queries that consume 200 milliseconds or more and less than 1 minute.
	// *   `{"Type":"cost","Min":"60000"}`: filters the queries that consume 1 minute or more.
	// *   `{"Type":"cost","Min":"30","Max":"50"}`: filters the queries that consume 30 milliseconds or more and less than 50 milliseconds. You can customize a filter condition by setting **Min** and **Max**.
	QueryCondition *string `json:"QueryCondition,omitempty" xml:"QueryCondition,omitempty"`
	// The ID of the resource group to which the instance belongs. For more information about how to obtain the ID of a resource group, see [View basic information of a resource group](~~151181~~).
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The beginning of the time range to query. Specify the time in the *yyyy-MM-dd*T*HH:mm*Z format. The time must be in UTC.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The name of the database account.
	User *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DownloadDiagnosisRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s DownloadDiagnosisRecordsRequest) GoString() string {
	return s.String()
}

func (s *DownloadDiagnosisRecordsRequest) SetDBInstanceId(v string) *DownloadDiagnosisRecordsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DownloadDiagnosisRecordsRequest) SetDatabase(v string) *DownloadDiagnosisRecordsRequest {
	s.Database = &v
	return s
}

func (s *DownloadDiagnosisRecordsRequest) SetEndTime(v string) *DownloadDiagnosisRecordsRequest {
	s.EndTime = &v
	return s
}

func (s *DownloadDiagnosisRecordsRequest) SetLang(v string) *DownloadDiagnosisRecordsRequest {
	s.Lang = &v
	return s
}

func (s *DownloadDiagnosisRecordsRequest) SetQueryCondition(v string) *DownloadDiagnosisRecordsRequest {
	s.QueryCondition = &v
	return s
}

func (s *DownloadDiagnosisRecordsRequest) SetResourceGroupId(v string) *DownloadDiagnosisRecordsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DownloadDiagnosisRecordsRequest) SetStartTime(v string) *DownloadDiagnosisRecordsRequest {
	s.StartTime = &v
	return s
}

func (s *DownloadDiagnosisRecordsRequest) SetUser(v string) *DownloadDiagnosisRecordsRequest {
	s.User = &v
	return s
}

type DownloadDiagnosisRecordsResponseBody struct {
	// The ID of the instance.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The ID of the download task.
	DownloadId *string `json:"DownloadId,omitempty" xml:"DownloadId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DownloadDiagnosisRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DownloadDiagnosisRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DownloadDiagnosisRecordsResponseBody) SetDBInstanceId(v string) *DownloadDiagnosisRecordsResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *DownloadDiagnosisRecordsResponseBody) SetDownloadId(v string) *DownloadDiagnosisRecordsResponseBody {
	s.DownloadId = &v
	return s
}

func (s *DownloadDiagnosisRecordsResponseBody) SetRequestId(v string) *DownloadDiagnosisRecordsResponseBody {
	s.RequestId = &v
	return s
}

type DownloadDiagnosisRecordsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DownloadDiagnosisRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DownloadDiagnosisRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s DownloadDiagnosisRecordsResponse) GoString() string {
	return s.String()
}

func (s *DownloadDiagnosisRecordsResponse) SetHeaders(v map[string]*string) *DownloadDiagnosisRecordsResponse {
	s.Headers = v
	return s
}

func (s *DownloadDiagnosisRecordsResponse) SetStatusCode(v int32) *DownloadDiagnosisRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *DownloadDiagnosisRecordsResponse) SetBody(v *DownloadDiagnosisRecordsResponseBody) *DownloadDiagnosisRecordsResponse {
	s.Body = v
	return s
}

type DownloadSQLLogsRecordsRequest struct {
	DBInstanceId   *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	Database       *string `json:"Database,omitempty" xml:"Database,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ExecuteCost    *string `json:"ExecuteCost,omitempty" xml:"ExecuteCost,omitempty"`
	ExecuteState   *string `json:"ExecuteState,omitempty" xml:"ExecuteState,omitempty"`
	Lang           *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	MaxExecuteCost *string `json:"MaxExecuteCost,omitempty" xml:"MaxExecuteCost,omitempty"`
	MinExecuteCost *string `json:"MinExecuteCost,omitempty" xml:"MinExecuteCost,omitempty"`
	OperationClass *string `json:"OperationClass,omitempty" xml:"OperationClass,omitempty"`
	OperationType  *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	PageNumber     *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	QueryKeywords  *string `json:"QueryKeywords,omitempty" xml:"QueryKeywords,omitempty"`
	SourceIP       *string `json:"SourceIP,omitempty" xml:"SourceIP,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	User           *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DownloadSQLLogsRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s DownloadSQLLogsRecordsRequest) GoString() string {
	return s.String()
}

func (s *DownloadSQLLogsRecordsRequest) SetDBInstanceId(v string) *DownloadSQLLogsRecordsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DownloadSQLLogsRecordsRequest) SetDatabase(v string) *DownloadSQLLogsRecordsRequest {
	s.Database = &v
	return s
}

func (s *DownloadSQLLogsRecordsRequest) SetEndTime(v string) *DownloadSQLLogsRecordsRequest {
	s.EndTime = &v
	return s
}

func (s *DownloadSQLLogsRecordsRequest) SetExecuteCost(v string) *DownloadSQLLogsRecordsRequest {
	s.ExecuteCost = &v
	return s
}

func (s *DownloadSQLLogsRecordsRequest) SetExecuteState(v string) *DownloadSQLLogsRecordsRequest {
	s.ExecuteState = &v
	return s
}

func (s *DownloadSQLLogsRecordsRequest) SetLang(v string) *DownloadSQLLogsRecordsRequest {
	s.Lang = &v
	return s
}

func (s *DownloadSQLLogsRecordsRequest) SetMaxExecuteCost(v string) *DownloadSQLLogsRecordsRequest {
	s.MaxExecuteCost = &v
	return s
}

func (s *DownloadSQLLogsRecordsRequest) SetMinExecuteCost(v string) *DownloadSQLLogsRecordsRequest {
	s.MinExecuteCost = &v
	return s
}

func (s *DownloadSQLLogsRecordsRequest) SetOperationClass(v string) *DownloadSQLLogsRecordsRequest {
	s.OperationClass = &v
	return s
}

func (s *DownloadSQLLogsRecordsRequest) SetOperationType(v string) *DownloadSQLLogsRecordsRequest {
	s.OperationType = &v
	return s
}

func (s *DownloadSQLLogsRecordsRequest) SetPageNumber(v int32) *DownloadSQLLogsRecordsRequest {
	s.PageNumber = &v
	return s
}

func (s *DownloadSQLLogsRecordsRequest) SetPageSize(v int32) *DownloadSQLLogsRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *DownloadSQLLogsRecordsRequest) SetQueryKeywords(v string) *DownloadSQLLogsRecordsRequest {
	s.QueryKeywords = &v
	return s
}

func (s *DownloadSQLLogsRecordsRequest) SetSourceIP(v string) *DownloadSQLLogsRecordsRequest {
	s.SourceIP = &v
	return s
}

func (s *DownloadSQLLogsRecordsRequest) SetStartTime(v string) *DownloadSQLLogsRecordsRequest {
	s.StartTime = &v
	return s
}

func (s *DownloadSQLLogsRecordsRequest) SetUser(v string) *DownloadSQLLogsRecordsRequest {
	s.User = &v
	return s
}

type DownloadSQLLogsRecordsResponseBody struct {
	DownloadId *int64  `json:"DownloadId,omitempty" xml:"DownloadId,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DownloadSQLLogsRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DownloadSQLLogsRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DownloadSQLLogsRecordsResponseBody) SetDownloadId(v int64) *DownloadSQLLogsRecordsResponseBody {
	s.DownloadId = &v
	return s
}

func (s *DownloadSQLLogsRecordsResponseBody) SetRequestId(v string) *DownloadSQLLogsRecordsResponseBody {
	s.RequestId = &v
	return s
}

type DownloadSQLLogsRecordsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DownloadSQLLogsRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DownloadSQLLogsRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s DownloadSQLLogsRecordsResponse) GoString() string {
	return s.String()
}

func (s *DownloadSQLLogsRecordsResponse) SetHeaders(v map[string]*string) *DownloadSQLLogsRecordsResponse {
	s.Headers = v
	return s
}

func (s *DownloadSQLLogsRecordsResponse) SetStatusCode(v int32) *DownloadSQLLogsRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *DownloadSQLLogsRecordsResponse) SetBody(v *DownloadSQLLogsRecordsResponseBody) *DownloadSQLLogsRecordsResponse {
	s.Body = v
	return s
}

type GrantCollectionRequest struct {
	Collection             *string `json:"Collection,omitempty" xml:"Collection,omitempty"`
	DBInstanceId           *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	GrantToNamespace       *string `json:"GrantToNamespace,omitempty" xml:"GrantToNamespace,omitempty"`
	GrantType              *string `json:"GrantType,omitempty" xml:"GrantType,omitempty"`
	ManagerAccount         *string `json:"ManagerAccount,omitempty" xml:"ManagerAccount,omitempty"`
	ManagerAccountPassword *string `json:"ManagerAccountPassword,omitempty" xml:"ManagerAccountPassword,omitempty"`
	Namespace              *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	OwnerId                *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId               *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GrantCollectionRequest) String() string {
	return tea.Prettify(s)
}

func (s GrantCollectionRequest) GoString() string {
	return s.String()
}

func (s *GrantCollectionRequest) SetCollection(v string) *GrantCollectionRequest {
	s.Collection = &v
	return s
}

func (s *GrantCollectionRequest) SetDBInstanceId(v string) *GrantCollectionRequest {
	s.DBInstanceId = &v
	return s
}

func (s *GrantCollectionRequest) SetGrantToNamespace(v string) *GrantCollectionRequest {
	s.GrantToNamespace = &v
	return s
}

func (s *GrantCollectionRequest) SetGrantType(v string) *GrantCollectionRequest {
	s.GrantType = &v
	return s
}

func (s *GrantCollectionRequest) SetManagerAccount(v string) *GrantCollectionRequest {
	s.ManagerAccount = &v
	return s
}

func (s *GrantCollectionRequest) SetManagerAccountPassword(v string) *GrantCollectionRequest {
	s.ManagerAccountPassword = &v
	return s
}

func (s *GrantCollectionRequest) SetNamespace(v string) *GrantCollectionRequest {
	s.Namespace = &v
	return s
}

func (s *GrantCollectionRequest) SetOwnerId(v int64) *GrantCollectionRequest {
	s.OwnerId = &v
	return s
}

func (s *GrantCollectionRequest) SetRegionId(v string) *GrantCollectionRequest {
	s.RegionId = &v
	return s
}

type GrantCollectionResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GrantCollectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GrantCollectionResponseBody) GoString() string {
	return s.String()
}

func (s *GrantCollectionResponseBody) SetMessage(v string) *GrantCollectionResponseBody {
	s.Message = &v
	return s
}

func (s *GrantCollectionResponseBody) SetRequestId(v string) *GrantCollectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GrantCollectionResponseBody) SetStatus(v string) *GrantCollectionResponseBody {
	s.Status = &v
	return s
}

type GrantCollectionResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GrantCollectionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GrantCollectionResponse) String() string {
	return tea.Prettify(s)
}

func (s GrantCollectionResponse) GoString() string {
	return s.String()
}

func (s *GrantCollectionResponse) SetHeaders(v map[string]*string) *GrantCollectionResponse {
	s.Headers = v
	return s
}

func (s *GrantCollectionResponse) SetStatusCode(v int32) *GrantCollectionResponse {
	s.StatusCode = &v
	return s
}

func (s *GrantCollectionResponse) SetBody(v *GrantCollectionResponseBody) *GrantCollectionResponse {
	s.Body = v
	return s
}

type InitVectorDatabaseRequest struct {
	DBInstanceId           *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	ManagerAccount         *string `json:"ManagerAccount,omitempty" xml:"ManagerAccount,omitempty"`
	ManagerAccountPassword *string `json:"ManagerAccountPassword,omitempty" xml:"ManagerAccountPassword,omitempty"`
	OwnerId                *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId               *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s InitVectorDatabaseRequest) String() string {
	return tea.Prettify(s)
}

func (s InitVectorDatabaseRequest) GoString() string {
	return s.String()
}

func (s *InitVectorDatabaseRequest) SetDBInstanceId(v string) *InitVectorDatabaseRequest {
	s.DBInstanceId = &v
	return s
}

func (s *InitVectorDatabaseRequest) SetManagerAccount(v string) *InitVectorDatabaseRequest {
	s.ManagerAccount = &v
	return s
}

func (s *InitVectorDatabaseRequest) SetManagerAccountPassword(v string) *InitVectorDatabaseRequest {
	s.ManagerAccountPassword = &v
	return s
}

func (s *InitVectorDatabaseRequest) SetOwnerId(v int64) *InitVectorDatabaseRequest {
	s.OwnerId = &v
	return s
}

func (s *InitVectorDatabaseRequest) SetRegionId(v string) *InitVectorDatabaseRequest {
	s.RegionId = &v
	return s
}

type InitVectorDatabaseResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s InitVectorDatabaseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InitVectorDatabaseResponseBody) GoString() string {
	return s.String()
}

func (s *InitVectorDatabaseResponseBody) SetMessage(v string) *InitVectorDatabaseResponseBody {
	s.Message = &v
	return s
}

func (s *InitVectorDatabaseResponseBody) SetRequestId(v string) *InitVectorDatabaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *InitVectorDatabaseResponseBody) SetStatus(v string) *InitVectorDatabaseResponseBody {
	s.Status = &v
	return s
}

type InitVectorDatabaseResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *InitVectorDatabaseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InitVectorDatabaseResponse) String() string {
	return tea.Prettify(s)
}

func (s InitVectorDatabaseResponse) GoString() string {
	return s.String()
}

func (s *InitVectorDatabaseResponse) SetHeaders(v map[string]*string) *InitVectorDatabaseResponse {
	s.Headers = v
	return s
}

func (s *InitVectorDatabaseResponse) SetStatusCode(v int32) *InitVectorDatabaseResponse {
	s.StatusCode = &v
	return s
}

func (s *InitVectorDatabaseResponse) SetBody(v *InitVectorDatabaseResponseBody) *InitVectorDatabaseResponse {
	s.Body = v
	return s
}

type ListCollectionsRequest struct {
	DBInstanceId      *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	Namespace         *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	NamespacePassword *string `json:"NamespacePassword,omitempty" xml:"NamespacePassword,omitempty"`
	OwnerId           *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListCollectionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCollectionsRequest) GoString() string {
	return s.String()
}

func (s *ListCollectionsRequest) SetDBInstanceId(v string) *ListCollectionsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ListCollectionsRequest) SetNamespace(v string) *ListCollectionsRequest {
	s.Namespace = &v
	return s
}

func (s *ListCollectionsRequest) SetNamespacePassword(v string) *ListCollectionsRequest {
	s.NamespacePassword = &v
	return s
}

func (s *ListCollectionsRequest) SetOwnerId(v int64) *ListCollectionsRequest {
	s.OwnerId = &v
	return s
}

func (s *ListCollectionsRequest) SetRegionId(v string) *ListCollectionsRequest {
	s.RegionId = &v
	return s
}

type ListCollectionsResponseBody struct {
	Collections  *ListCollectionsResponseBodyCollections `json:"Collections,omitempty" xml:"Collections,omitempty" type:"Struct"`
	Count        *int32                                  `json:"Count,omitempty" xml:"Count,omitempty"`
	DBInstanceId *string                                 `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	Message      *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	Namespace    *string                                 `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	RegionId     *string                                 `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RequestId    *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status       *string                                 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListCollectionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCollectionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCollectionsResponseBody) SetCollections(v *ListCollectionsResponseBodyCollections) *ListCollectionsResponseBody {
	s.Collections = v
	return s
}

func (s *ListCollectionsResponseBody) SetCount(v int32) *ListCollectionsResponseBody {
	s.Count = &v
	return s
}

func (s *ListCollectionsResponseBody) SetDBInstanceId(v string) *ListCollectionsResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *ListCollectionsResponseBody) SetMessage(v string) *ListCollectionsResponseBody {
	s.Message = &v
	return s
}

func (s *ListCollectionsResponseBody) SetNamespace(v string) *ListCollectionsResponseBody {
	s.Namespace = &v
	return s
}

func (s *ListCollectionsResponseBody) SetRegionId(v string) *ListCollectionsResponseBody {
	s.RegionId = &v
	return s
}

func (s *ListCollectionsResponseBody) SetRequestId(v string) *ListCollectionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCollectionsResponseBody) SetStatus(v string) *ListCollectionsResponseBody {
	s.Status = &v
	return s
}

type ListCollectionsResponseBodyCollections struct {
	Collection []*string `json:"Collection,omitempty" xml:"Collection,omitempty" type:"Repeated"`
}

func (s ListCollectionsResponseBodyCollections) String() string {
	return tea.Prettify(s)
}

func (s ListCollectionsResponseBodyCollections) GoString() string {
	return s.String()
}

func (s *ListCollectionsResponseBodyCollections) SetCollection(v []*string) *ListCollectionsResponseBodyCollections {
	s.Collection = v
	return s
}

type ListCollectionsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListCollectionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListCollectionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCollectionsResponse) GoString() string {
	return s.String()
}

func (s *ListCollectionsResponse) SetHeaders(v map[string]*string) *ListCollectionsResponse {
	s.Headers = v
	return s
}

func (s *ListCollectionsResponse) SetStatusCode(v int32) *ListCollectionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCollectionsResponse) SetBody(v *ListCollectionsResponseBody) *ListCollectionsResponse {
	s.Body = v
	return s
}

type ListNamespacesRequest struct {
	DBInstanceId           *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	ManagerAccount         *string `json:"ManagerAccount,omitempty" xml:"ManagerAccount,omitempty"`
	ManagerAccountPassword *string `json:"ManagerAccountPassword,omitempty" xml:"ManagerAccountPassword,omitempty"`
	OwnerId                *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId               *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListNamespacesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListNamespacesRequest) GoString() string {
	return s.String()
}

func (s *ListNamespacesRequest) SetDBInstanceId(v string) *ListNamespacesRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ListNamespacesRequest) SetManagerAccount(v string) *ListNamespacesRequest {
	s.ManagerAccount = &v
	return s
}

func (s *ListNamespacesRequest) SetManagerAccountPassword(v string) *ListNamespacesRequest {
	s.ManagerAccountPassword = &v
	return s
}

func (s *ListNamespacesRequest) SetOwnerId(v int64) *ListNamespacesRequest {
	s.OwnerId = &v
	return s
}

func (s *ListNamespacesRequest) SetRegionId(v string) *ListNamespacesRequest {
	s.RegionId = &v
	return s
}

type ListNamespacesResponseBody struct {
	Count        *int32                                `json:"Count,omitempty" xml:"Count,omitempty"`
	DBInstanceId *string                               `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	Message      *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	Namespaces   *ListNamespacesResponseBodyNamespaces `json:"Namespaces,omitempty" xml:"Namespaces,omitempty" type:"Struct"`
	RegionId     *string                               `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RequestId    *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status       *string                               `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListNamespacesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListNamespacesResponseBody) GoString() string {
	return s.String()
}

func (s *ListNamespacesResponseBody) SetCount(v int32) *ListNamespacesResponseBody {
	s.Count = &v
	return s
}

func (s *ListNamespacesResponseBody) SetDBInstanceId(v string) *ListNamespacesResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *ListNamespacesResponseBody) SetMessage(v string) *ListNamespacesResponseBody {
	s.Message = &v
	return s
}

func (s *ListNamespacesResponseBody) SetNamespaces(v *ListNamespacesResponseBodyNamespaces) *ListNamespacesResponseBody {
	s.Namespaces = v
	return s
}

func (s *ListNamespacesResponseBody) SetRegionId(v string) *ListNamespacesResponseBody {
	s.RegionId = &v
	return s
}

func (s *ListNamespacesResponseBody) SetRequestId(v string) *ListNamespacesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNamespacesResponseBody) SetStatus(v string) *ListNamespacesResponseBody {
	s.Status = &v
	return s
}

type ListNamespacesResponseBodyNamespaces struct {
	Namespace []*string `json:"Namespace,omitempty" xml:"Namespace,omitempty" type:"Repeated"`
}

func (s ListNamespacesResponseBodyNamespaces) String() string {
	return tea.Prettify(s)
}

func (s ListNamespacesResponseBodyNamespaces) GoString() string {
	return s.String()
}

func (s *ListNamespacesResponseBodyNamespaces) SetNamespace(v []*string) *ListNamespacesResponseBodyNamespaces {
	s.Namespace = v
	return s
}

type ListNamespacesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListNamespacesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListNamespacesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListNamespacesResponse) GoString() string {
	return s.String()
}

func (s *ListNamespacesResponse) SetHeaders(v map[string]*string) *ListNamespacesResponse {
	s.Headers = v
	return s
}

func (s *ListNamespacesResponse) SetStatusCode(v int32) *ListNamespacesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNamespacesResponse) SetBody(v *ListNamespacesResponseBody) *ListNamespacesResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	// The token used to perform the next query.
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the instance. You can call the [DescribeRegions](~~86912~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of instance N. Valid values of N: 1 to 50.
	ResourceId           []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The storage mode of the instance. Valid values:
	//
	// *   `instance`: reserved storage mode
	// *   `ALIYUN::GPDB::INSTANCE`: elastic storage mode
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The list of tags.
	Tag []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) SetNextToken(v string) *ListTagResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesRequest) SetOwnerAccount(v string) *ListTagResourcesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListTagResourcesRequest) SetOwnerId(v int64) *ListTagResourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *ListTagResourcesRequest) SetRegionId(v string) *ListTagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceId(v []*string) *ListTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *ListTagResourcesRequest) SetResourceOwnerAccount(v string) *ListTagResourcesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceOwnerId(v int64) *ListTagResourcesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceType(v string) *ListTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesRequest) SetTag(v []*ListTagResourcesRequestTag) *ListTagResourcesRequest {
	s.Tag = v
	return s
}

type ListTagResourcesRequestTag struct {
	// The key of tag N. The key must be 1 to 128 characters in length. Valid values of N: 1 to 20.
	//
	// You can use `Tag.N.Key and Tag.N.Value` to query AnalyticDB for PostgreSQL instances to which specific tags are bound.
	//
	// *   If you specify only `Tag.N.Key`, the instances whose tags contain the specified tag keys are returned.
	// *   If you specify only `Tag.N.Value`, `InvalidParameter.TagValue` is returned.
	// *   If you specify multiple tag key-value pairs at a time, the instances to which all the specified tags are bound are returned.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N. The value must be 1 to 128 characters in length. Valid values of N: 1 to 20.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequestTag) SetKey(v string) *ListTagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *ListTagResourcesRequestTag) SetValue(v string) *ListTagResourcesRequestTag {
	s.Value = &v
	return s
}

type ListTagResourcesResponseBody struct {
	// The token used to perform the next query.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Details about the instances and tags, including the instance IDs, instance modes, and tag key-value pairs.
	TagResources *ListTagResourcesResponseBodyTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Struct"`
}

func (s ListTagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBody) SetNextToken(v string) *ListTagResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetRequestId(v string) *ListTagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetTagResources(v *ListTagResourcesResponseBodyTagResources) *ListTagResourcesResponseBody {
	s.TagResources = v
	return s
}

type ListTagResourcesResponseBodyTagResources struct {
	TagResource []*ListTagResourcesResponseBodyTagResourcesTagResource `json:"TagResource,omitempty" xml:"TagResource,omitempty" type:"Repeated"`
}

func (s ListTagResourcesResponseBodyTagResources) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagResource(v []*ListTagResourcesResponseBodyTagResourcesTagResource) *ListTagResourcesResponseBodyTagResources {
	s.TagResource = v
	return s
}

type ListTagResourcesResponseBodyTagResourcesTagResource struct {
	// The ID of the instance.
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The storage mode of the instance.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tag key.
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListTagResourcesResponseBodyTagResourcesTagResource) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResourcesTagResource) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetResourceId(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetResourceType(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetTagKey(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.TagKey = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetTagValue(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.TagValue = &v
	return s
}

type ListTagResourcesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponse) SetHeaders(v map[string]*string) *ListTagResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListTagResourcesResponse) SetStatusCode(v int32) *ListTagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagResourcesResponse) SetBody(v *ListTagResourcesResponseBody) *ListTagResourcesResponse {
	s.Body = v
	return s
}

type ModifyAccountDescriptionRequest struct {
	// The description of the account. The description must meet the following requirements:
	//
	// *   The description must start with a letter.
	// *   The description can contain letters, digits, underscores (\_), and hyphens (-).
	// *   The description cannot start with `http://` or `https://`.
	// *   The description must be 2 to 256 characters in length.
	AccountDescription *string `json:"AccountDescription,omitempty" xml:"AccountDescription,omitempty"`
	// The name of the database account.
	//
	// > You can call the [DescribeAccounts](~~~~) operation to query the information about database accounts in a cluster, including the database account name.
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The ID of the instance.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s ModifyAccountDescriptionRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccountDescriptionRequest) GoString() string {
	return s.String()
}

func (s *ModifyAccountDescriptionRequest) SetAccountDescription(v string) *ModifyAccountDescriptionRequest {
	s.AccountDescription = &v
	return s
}

func (s *ModifyAccountDescriptionRequest) SetAccountName(v string) *ModifyAccountDescriptionRequest {
	s.AccountName = &v
	return s
}

func (s *ModifyAccountDescriptionRequest) SetDBInstanceId(v string) *ModifyAccountDescriptionRequest {
	s.DBInstanceId = &v
	return s
}

type ModifyAccountDescriptionResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAccountDescriptionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccountDescriptionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAccountDescriptionResponseBody) SetRequestId(v string) *ModifyAccountDescriptionResponseBody {
	s.RequestId = &v
	return s
}

type ModifyAccountDescriptionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyAccountDescriptionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyAccountDescriptionResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccountDescriptionResponse) GoString() string {
	return s.String()
}

func (s *ModifyAccountDescriptionResponse) SetHeaders(v map[string]*string) *ModifyAccountDescriptionResponse {
	s.Headers = v
	return s
}

func (s *ModifyAccountDescriptionResponse) SetStatusCode(v int32) *ModifyAccountDescriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAccountDescriptionResponse) SetBody(v *ModifyAccountDescriptionResponseBody) *ModifyAccountDescriptionResponse {
	s.Body = v
	return s
}

type ModifyBackupPolicyRequest struct {
	// The number of days for which data backup files are retained. Default value: 7. Maximum value: 7. Valid values: 1 to 7.
	BackupRetentionPeriod *int32 `json:"BackupRetentionPeriod,omitempty" xml:"BackupRetentionPeriod,omitempty"`
	// The ID of the instance.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// Specifies whether to enable automatic point-in-time backup.
	//
	// *   true
	// *   false
	//
	// Default value: true.
	EnableRecoveryPoint *bool `json:"EnableRecoveryPoint,omitempty" xml:"EnableRecoveryPoint,omitempty"`
	// The cycle based on which you want to perform a backup. Separate multiple values with commas (,). Valid values:
	//
	// *   Monday
	// *   Tuesday
	// *   Wednesday
	// *   Thursday
	// *   Friday
	// *   Saturday
	// *   Sunday
	PreferredBackupPeriod *string `json:"PreferredBackupPeriod,omitempty" xml:"PreferredBackupPeriod,omitempty"`
	// The backup window. Specify the backup window in the HH:mmZ-HH:mmZ format. The backup window must be in UTC. Default value: 00:00-01:00.
	PreferredBackupTime *string `json:"PreferredBackupTime,omitempty" xml:"PreferredBackupTime,omitempty"`
	// The frequency of point-in-time backup.
	//
	// *   1: per hour
	// *   2: per 2 hours
	// *   4: per 4 hours
	// *   8: per 8 hours
	//
	// Default value: 8.
	RecoveryPointPeriod *string `json:"RecoveryPointPeriod,omitempty" xml:"RecoveryPointPeriod,omitempty"`
}

func (s ModifyBackupPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyBackupPolicyRequest) GoString() string {
	return s.String()
}

func (s *ModifyBackupPolicyRequest) SetBackupRetentionPeriod(v int32) *ModifyBackupPolicyRequest {
	s.BackupRetentionPeriod = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetDBInstanceId(v string) *ModifyBackupPolicyRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetEnableRecoveryPoint(v bool) *ModifyBackupPolicyRequest {
	s.EnableRecoveryPoint = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetPreferredBackupPeriod(v string) *ModifyBackupPolicyRequest {
	s.PreferredBackupPeriod = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetPreferredBackupTime(v string) *ModifyBackupPolicyRequest {
	s.PreferredBackupTime = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetRecoveryPointPeriod(v string) *ModifyBackupPolicyRequest {
	s.RecoveryPointPeriod = &v
	return s
}

type ModifyBackupPolicyResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyBackupPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyBackupPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyBackupPolicyResponseBody) SetRequestId(v string) *ModifyBackupPolicyResponseBody {
	s.RequestId = &v
	return s
}

type ModifyBackupPolicyResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyBackupPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyBackupPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyBackupPolicyResponse) GoString() string {
	return s.String()
}

func (s *ModifyBackupPolicyResponse) SetHeaders(v map[string]*string) *ModifyBackupPolicyResponse {
	s.Headers = v
	return s
}

func (s *ModifyBackupPolicyResponse) SetStatusCode(v int32) *ModifyBackupPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyBackupPolicyResponse) SetBody(v *ModifyBackupPolicyResponseBody) *ModifyBackupPolicyResponse {
	s.Body = v
	return s
}

type ModifyDBInstanceConfigRequest struct {
	// The description of the instance.
	DBInstanceDescription *string `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	// The ID of the instance.
	//
	// >  You can call the [DescribeDBInstances](~~86911~~) operation to query the instance IDs of all AnalyticDB for PostgreSQL instances in a specific region.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The wait period for the instance that has no traffic to become idle. Minimum value: 60. Default value: 600. Unit: seconds.
	IdleTime *int32 `json:"IdleTime,omitempty" xml:"IdleTime,omitempty"`
	// The ID of the resource group to which the instance belongs. For more information about how to obtain the ID of a resource group, see [View basic information of a resource group](~~151181~~).
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The threshold of computing resources. Valid values: 8 to 32. Unit: AnalyticDB Compute Units (ACUs).
	ServerlessResource *int32 `json:"ServerlessResource,omitempty" xml:"ServerlessResource,omitempty"`
}

func (s ModifyDBInstanceConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceConfigRequest) SetDBInstanceDescription(v string) *ModifyDBInstanceConfigRequest {
	s.DBInstanceDescription = &v
	return s
}

func (s *ModifyDBInstanceConfigRequest) SetDBInstanceId(v string) *ModifyDBInstanceConfigRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceConfigRequest) SetIdleTime(v int32) *ModifyDBInstanceConfigRequest {
	s.IdleTime = &v
	return s
}

func (s *ModifyDBInstanceConfigRequest) SetResourceGroupId(v string) *ModifyDBInstanceConfigRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyDBInstanceConfigRequest) SetServerlessResource(v int32) *ModifyDBInstanceConfigRequest {
	s.ServerlessResource = &v
	return s
}

type ModifyDBInstanceConfigResponseBody struct {
	// The ID of the instance.
	DbInstanceId *string `json:"DbInstanceId,omitempty" xml:"DbInstanceId,omitempty"`
	// The error message returned if the operation fails.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The state of the operation. Valid values:
	//
	// *   **0**: The operation failed.
	// *   **1**: The operation is successful.
	Status *bool `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ModifyDBInstanceConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceConfigResponseBody) SetDbInstanceId(v string) *ModifyDBInstanceConfigResponseBody {
	s.DbInstanceId = &v
	return s
}

func (s *ModifyDBInstanceConfigResponseBody) SetErrorMessage(v string) *ModifyDBInstanceConfigResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ModifyDBInstanceConfigResponseBody) SetRequestId(v string) *ModifyDBInstanceConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBInstanceConfigResponseBody) SetStatus(v bool) *ModifyDBInstanceConfigResponseBody {
	s.Status = &v
	return s
}

type ModifyDBInstanceConfigResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyDBInstanceConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDBInstanceConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceConfigResponse) SetHeaders(v map[string]*string) *ModifyDBInstanceConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBInstanceConfigResponse) SetStatusCode(v int32) *ModifyDBInstanceConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBInstanceConfigResponse) SetBody(v *ModifyDBInstanceConfigResponseBody) *ModifyDBInstanceConfigResponse {
	s.Body = v
	return s
}

type ModifyDBInstanceConnectionStringRequest struct {
	// The new endpoint of the instance.
	ConnectionStringPrefix *string `json:"ConnectionStringPrefix,omitempty" xml:"ConnectionStringPrefix,omitempty"`
	// The original endpoint of the instance.
	CurrentConnectionString *string `json:"CurrentConnectionString,omitempty" xml:"CurrentConnectionString,omitempty"`
	// The ID of the instance.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The new port number of the instance.
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s ModifyDBInstanceConnectionStringRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceConnectionStringRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceConnectionStringRequest) SetConnectionStringPrefix(v string) *ModifyDBInstanceConnectionStringRequest {
	s.ConnectionStringPrefix = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringRequest) SetCurrentConnectionString(v string) *ModifyDBInstanceConnectionStringRequest {
	s.CurrentConnectionString = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringRequest) SetDBInstanceId(v string) *ModifyDBInstanceConnectionStringRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringRequest) SetPort(v string) *ModifyDBInstanceConnectionStringRequest {
	s.Port = &v
	return s
}

type ModifyDBInstanceConnectionStringResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBInstanceConnectionStringResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceConnectionStringResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceConnectionStringResponseBody) SetRequestId(v string) *ModifyDBInstanceConnectionStringResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDBInstanceConnectionStringResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyDBInstanceConnectionStringResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDBInstanceConnectionStringResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceConnectionStringResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceConnectionStringResponse) SetHeaders(v map[string]*string) *ModifyDBInstanceConnectionStringResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBInstanceConnectionStringResponse) SetStatusCode(v int32) *ModifyDBInstanceConnectionStringResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringResponse) SetBody(v *ModifyDBInstanceConnectionStringResponseBody) *ModifyDBInstanceConnectionStringResponse {
	s.Body = v
	return s
}

type ModifyDBInstanceDescriptionRequest struct {
	// The description of the instance.
	//
	// The description must be 2 to 256 characters in length. It cannot start with http:// or https://.
	DBInstanceDescription *string `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	// The ID of the instance.
	//
	// >  You can call the [DescribeDBInstances](~~86911~~) operation to query the instance IDs of all AnalyticDB for PostgreSQL instances in a specific region.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The ID of the resource group to which the instance belongs. For more information about how to obtain the ID of a resource group, see [View basic information of a resource group](~~151181~~).
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ModifyDBInstanceDescriptionRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceDescriptionRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceDescriptionRequest) SetDBInstanceDescription(v string) *ModifyDBInstanceDescriptionRequest {
	s.DBInstanceDescription = &v
	return s
}

func (s *ModifyDBInstanceDescriptionRequest) SetDBInstanceId(v string) *ModifyDBInstanceDescriptionRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceDescriptionRequest) SetResourceGroupId(v string) *ModifyDBInstanceDescriptionRequest {
	s.ResourceGroupId = &v
	return s
}

type ModifyDBInstanceDescriptionResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBInstanceDescriptionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceDescriptionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceDescriptionResponseBody) SetRequestId(v string) *ModifyDBInstanceDescriptionResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDBInstanceDescriptionResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyDBInstanceDescriptionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDBInstanceDescriptionResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceDescriptionResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceDescriptionResponse) SetHeaders(v map[string]*string) *ModifyDBInstanceDescriptionResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBInstanceDescriptionResponse) SetStatusCode(v int32) *ModifyDBInstanceDescriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBInstanceDescriptionResponse) SetBody(v *ModifyDBInstanceDescriptionResponseBody) *ModifyDBInstanceDescriptionResponse {
	s.Body = v
	return s
}

type ModifyDBInstanceMaintainTimeRequest struct {
	// The ID of the instance.
	//
	// >  You can call the [DescribeDBInstances](~~86911~~) operation to query the details of all AnalyticDB for PostgreSQL instances in a specific region, including instance IDs.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The end time of the maintenance window. The end time must be later than the start time. Specify the time in the HH:mmZ format. The time must be in UTC.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the resource group to which the instance belongs. For more information about how to obtain the ID of a resource group, see [View basic information of a resource group](~~151181~~).
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The start time of the maintenance window. Specify the time in the HH:mmZ format. The time must be in UTC.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ModifyDBInstanceMaintainTimeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceMaintainTimeRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceMaintainTimeRequest) SetDBInstanceId(v string) *ModifyDBInstanceMaintainTimeRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceMaintainTimeRequest) SetEndTime(v string) *ModifyDBInstanceMaintainTimeRequest {
	s.EndTime = &v
	return s
}

func (s *ModifyDBInstanceMaintainTimeRequest) SetResourceGroupId(v string) *ModifyDBInstanceMaintainTimeRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyDBInstanceMaintainTimeRequest) SetStartTime(v string) *ModifyDBInstanceMaintainTimeRequest {
	s.StartTime = &v
	return s
}

type ModifyDBInstanceMaintainTimeResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBInstanceMaintainTimeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceMaintainTimeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceMaintainTimeResponseBody) SetRequestId(v string) *ModifyDBInstanceMaintainTimeResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDBInstanceMaintainTimeResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyDBInstanceMaintainTimeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDBInstanceMaintainTimeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceMaintainTimeResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceMaintainTimeResponse) SetHeaders(v map[string]*string) *ModifyDBInstanceMaintainTimeResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBInstanceMaintainTimeResponse) SetStatusCode(v int32) *ModifyDBInstanceMaintainTimeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBInstanceMaintainTimeResponse) SetBody(v *ModifyDBInstanceMaintainTimeResponseBody) *ModifyDBInstanceMaintainTimeResponse {
	s.Body = v
	return s
}

type ModifyDBInstanceResourceGroupRequest struct {
	// The ID of the instance.
	//
	// >  You can call the [DescribeDBInstances](~~86911~~) operation to query the instance IDs of all AnalyticDB for PostgreSQL instances in a specific region.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The ID of the resource group to which you want to move the instance. For more information about how to obtain the ID of a resource group, see [View basic information of a resource group](~~151181~~).
	NewResourceGroupId *string `json:"NewResourceGroupId,omitempty" xml:"NewResourceGroupId,omitempty"`
	OwnerAccount       *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId            *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the resource group to which the instance belongs. For more information about how to obtain the ID of a resource group, see [View basic information of a resource group](~~151181~~).
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyDBInstanceResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceResourceGroupRequest) SetDBInstanceId(v string) *ModifyDBInstanceResourceGroupRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceResourceGroupRequest) SetNewResourceGroupId(v string) *ModifyDBInstanceResourceGroupRequest {
	s.NewResourceGroupId = &v
	return s
}

func (s *ModifyDBInstanceResourceGroupRequest) SetOwnerAccount(v string) *ModifyDBInstanceResourceGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBInstanceResourceGroupRequest) SetOwnerId(v int64) *ModifyDBInstanceResourceGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBInstanceResourceGroupRequest) SetResourceGroupId(v string) *ModifyDBInstanceResourceGroupRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyDBInstanceResourceGroupRequest) SetResourceOwnerAccount(v string) *ModifyDBInstanceResourceGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBInstanceResourceGroupRequest) SetResourceOwnerId(v int64) *ModifyDBInstanceResourceGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

type ModifyDBInstanceResourceGroupResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBInstanceResourceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceResourceGroupResponseBody) SetRequestId(v string) *ModifyDBInstanceResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDBInstanceResourceGroupResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyDBInstanceResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDBInstanceResourceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceResourceGroupResponse) SetHeaders(v map[string]*string) *ModifyDBInstanceResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBInstanceResourceGroupResponse) SetStatusCode(v int32) *ModifyDBInstanceResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBInstanceResourceGroupResponse) SetBody(v *ModifyDBInstanceResourceGroupResponseBody) *ModifyDBInstanceResourceGroupResponse {
	s.Body = v
	return s
}

type ModifyDBInstanceSSLRequest struct {
	// The encrypted endpoint. By default, the wildcards are used for instances that are hosted on ECS instances. This way, the endpoints that can be resolved to the same IP address are encrypted.
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// The ID of the instance.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The status of SSL encryption. Valid values:
	//
	// *   0: disables SSL encryption.
	// *   1: enables SSL encryption.
	// *   2: updates SSL encryption.
	SSLEnabled *int32 `json:"SSLEnabled,omitempty" xml:"SSLEnabled,omitempty"`
}

func (s ModifyDBInstanceSSLRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceSSLRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceSSLRequest) SetConnectionString(v string) *ModifyDBInstanceSSLRequest {
	s.ConnectionString = &v
	return s
}

func (s *ModifyDBInstanceSSLRequest) SetDBInstanceId(v string) *ModifyDBInstanceSSLRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceSSLRequest) SetSSLEnabled(v int32) *ModifyDBInstanceSSLRequest {
	s.SSLEnabled = &v
	return s
}

type ModifyDBInstanceSSLResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBInstanceSSLResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceSSLResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceSSLResponseBody) SetRequestId(v string) *ModifyDBInstanceSSLResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDBInstanceSSLResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyDBInstanceSSLResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDBInstanceSSLResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceSSLResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceSSLResponse) SetHeaders(v map[string]*string) *ModifyDBInstanceSSLResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBInstanceSSLResponse) SetStatusCode(v int32) *ModifyDBInstanceSSLResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBInstanceSSLResponse) SetBody(v *ModifyDBInstanceSSLResponseBody) *ModifyDBInstanceSSLResponse {
	s.Body = v
	return s
}

type ModifyParametersRequest struct {
	// The ID of the instance.
	//
	// >  You can call the [DescribeDBInstances](~~86911~~) operation to query the details of all AnalyticDB for PostgreSQL instances in a specific region, including instance IDs.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// Specifies whether to forcibly restart the instance. Valid values:
	//
	// *   **true**
	// *   **false**
	ForceRestartInstance *bool `json:"ForceRestartInstance,omitempty" xml:"ForceRestartInstance,omitempty"`
	// The name and value of the parameter to be modified. Specify the parameter in the `<Parameter name>:<Parameter value>` format.
	//
	// You can call the [DescribeParameters](~~208310~~) operation to query the parameters that can be modified.
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
}

func (s ModifyParametersRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyParametersRequest) GoString() string {
	return s.String()
}

func (s *ModifyParametersRequest) SetDBInstanceId(v string) *ModifyParametersRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyParametersRequest) SetForceRestartInstance(v bool) *ModifyParametersRequest {
	s.ForceRestartInstance = &v
	return s
}

func (s *ModifyParametersRequest) SetParameters(v string) *ModifyParametersRequest {
	s.Parameters = &v
	return s
}

type ModifyParametersResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyParametersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyParametersResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyParametersResponseBody) SetRequestId(v string) *ModifyParametersResponseBody {
	s.RequestId = &v
	return s
}

type ModifyParametersResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyParametersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyParametersResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyParametersResponse) GoString() string {
	return s.String()
}

func (s *ModifyParametersResponse) SetHeaders(v map[string]*string) *ModifyParametersResponse {
	s.Headers = v
	return s
}

func (s *ModifyParametersResponse) SetStatusCode(v int32) *ModifyParametersResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyParametersResponse) SetBody(v *ModifyParametersResponseBody) *ModifyParametersResponse {
	s.Body = v
	return s
}

type ModifySQLCollectorPolicyRequest struct {
	// The ID of the instance.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// Specifies whether to enable or disable SQL collection.
	//
	// *   Enable: enables SQL collection.
	// *   Disabled: disables SQL collection.
	SQLCollectorStatus *string `json:"SQLCollectorStatus,omitempty" xml:"SQLCollectorStatus,omitempty"`
}

func (s ModifySQLCollectorPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifySQLCollectorPolicyRequest) GoString() string {
	return s.String()
}

func (s *ModifySQLCollectorPolicyRequest) SetDBInstanceId(v string) *ModifySQLCollectorPolicyRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifySQLCollectorPolicyRequest) SetSQLCollectorStatus(v string) *ModifySQLCollectorPolicyRequest {
	s.SQLCollectorStatus = &v
	return s
}

type ModifySQLCollectorPolicyResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifySQLCollectorPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifySQLCollectorPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySQLCollectorPolicyResponseBody) SetRequestId(v string) *ModifySQLCollectorPolicyResponseBody {
	s.RequestId = &v
	return s
}

type ModifySQLCollectorPolicyResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifySQLCollectorPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifySQLCollectorPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifySQLCollectorPolicyResponse) GoString() string {
	return s.String()
}

func (s *ModifySQLCollectorPolicyResponse) SetHeaders(v map[string]*string) *ModifySQLCollectorPolicyResponse {
	s.Headers = v
	return s
}

func (s *ModifySQLCollectorPolicyResponse) SetStatusCode(v int32) *ModifySQLCollectorPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySQLCollectorPolicyResponse) SetBody(v *ModifySQLCollectorPolicyResponseBody) *ModifySQLCollectorPolicyResponse {
	s.Body = v
	return s
}

type ModifySecurityIpsRequest struct {
	// The attribute of the IP address whitelist. By default, this parameter is empty. A whitelist with the `hidden` attribute does not appear in the console.
	DBInstanceIPArrayAttribute *string `json:"DBInstanceIPArrayAttribute,omitempty" xml:"DBInstanceIPArrayAttribute,omitempty"`
	// The name of the whitelist. If you do not enter a name, IP addresses are added to the default whitelist.
	//
	// >  You can create up to 50 whitelists for an instance.
	DBInstanceIPArrayName *string `json:"DBInstanceIPArrayName,omitempty" xml:"DBInstanceIPArrayName,omitempty"`
	// The ID of the instance.
	//
	// >  You can call the [DescribeDBInstances](~~86911~~) operation to query the instance IDs of all AnalyticDB for PostgreSQL instances in a specific region.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The method of modification. Valid values:
	//
	// *   **Cover**: overwrites the whitelist.
	// *   **Append**: appends data to the whitelist.
	// *   **Delete**: deletes the whitelist.
	ModifyMode *string `json:"ModifyMode,omitempty" xml:"ModifyMode,omitempty"`
	// The ID of the resource group to which the instance belongs. For more information about how to obtain the ID of a resource group, see [View basic information of a resource group](~~151181~~).
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The IP addresses listed in the whitelist. You can add up to 1,000 IP addresses to the whitelist. Separate multiple IP addresses with commas (,). The IP addresses must use one of the following formats:
	//
	// *   0.0.0.0/0
	// *   10.23.12.24. This is a standard IP address.
	// *   10.23.12.24/24. This is a CIDR block. The value `/24` indicates that the prefix of the CIDR block is 24-bit long. You can replace 24 with a value in the range of `1 to 32`.
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
}

func (s ModifySecurityIpsRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifySecurityIpsRequest) GoString() string {
	return s.String()
}

func (s *ModifySecurityIpsRequest) SetDBInstanceIPArrayAttribute(v string) *ModifySecurityIpsRequest {
	s.DBInstanceIPArrayAttribute = &v
	return s
}

func (s *ModifySecurityIpsRequest) SetDBInstanceIPArrayName(v string) *ModifySecurityIpsRequest {
	s.DBInstanceIPArrayName = &v
	return s
}

func (s *ModifySecurityIpsRequest) SetDBInstanceId(v string) *ModifySecurityIpsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifySecurityIpsRequest) SetModifyMode(v string) *ModifySecurityIpsRequest {
	s.ModifyMode = &v
	return s
}

func (s *ModifySecurityIpsRequest) SetResourceGroupId(v string) *ModifySecurityIpsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifySecurityIpsRequest) SetSecurityIPList(v string) *ModifySecurityIpsRequest {
	s.SecurityIPList = &v
	return s
}

type ModifySecurityIpsResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifySecurityIpsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifySecurityIpsResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySecurityIpsResponseBody) SetRequestId(v string) *ModifySecurityIpsResponseBody {
	s.RequestId = &v
	return s
}

type ModifySecurityIpsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifySecurityIpsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifySecurityIpsResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifySecurityIpsResponse) GoString() string {
	return s.String()
}

func (s *ModifySecurityIpsResponse) SetHeaders(v map[string]*string) *ModifySecurityIpsResponse {
	s.Headers = v
	return s
}

func (s *ModifySecurityIpsResponse) SetStatusCode(v int32) *ModifySecurityIpsResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySecurityIpsResponse) SetBody(v *ModifySecurityIpsResponseBody) *ModifySecurityIpsResponse {
	s.Body = v
	return s
}

type ModifyVectorConfigurationRequest struct {
	// The instance ID.
	//
	// >  You can call the [DescribeDBInstances](~~86911~~) operation to query the IDs of all AnalyticDB for PostgreSQL instances in a region.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Specifies whether to enable vector engine optimization. Valid values:
	//
	// *   **enabled**
	// *   **disabled**
	//
	// > *   We recommend that you **do not enable** vector engine optimization in mainstream analysis and real-time data warehousing scenarios.
	// > *   We recommend that you **enable** vector engine optimization in AI Generated Content (AIGC) and vector retrieval scenarios that require the vector analysis engine.
	VectorConfigurationStatus *string `json:"VectorConfigurationStatus,omitempty" xml:"VectorConfigurationStatus,omitempty"`
}

func (s ModifyVectorConfigurationRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyVectorConfigurationRequest) GoString() string {
	return s.String()
}

func (s *ModifyVectorConfigurationRequest) SetDBInstanceId(v string) *ModifyVectorConfigurationRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyVectorConfigurationRequest) SetOwnerId(v int64) *ModifyVectorConfigurationRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyVectorConfigurationRequest) SetVectorConfigurationStatus(v string) *ModifyVectorConfigurationRequest {
	s.VectorConfigurationStatus = &v
	return s
}

type ModifyVectorConfigurationResponseBody struct {
	// The instance ID.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The error message that is returned.
	//
	// This parameter is returned only if the request fails.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// *   **true**
	// *   **false**
	Status *bool `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ModifyVectorConfigurationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyVectorConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyVectorConfigurationResponseBody) SetDBInstanceId(v string) *ModifyVectorConfigurationResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyVectorConfigurationResponseBody) SetErrorMessage(v string) *ModifyVectorConfigurationResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ModifyVectorConfigurationResponseBody) SetRequestId(v string) *ModifyVectorConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyVectorConfigurationResponseBody) SetStatus(v bool) *ModifyVectorConfigurationResponseBody {
	s.Status = &v
	return s
}

type ModifyVectorConfigurationResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyVectorConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyVectorConfigurationResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyVectorConfigurationResponse) GoString() string {
	return s.String()
}

func (s *ModifyVectorConfigurationResponse) SetHeaders(v map[string]*string) *ModifyVectorConfigurationResponse {
	s.Headers = v
	return s
}

func (s *ModifyVectorConfigurationResponse) SetStatusCode(v int32) *ModifyVectorConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyVectorConfigurationResponse) SetBody(v *ModifyVectorConfigurationResponseBody) *ModifyVectorConfigurationResponse {
	s.Body = v
	return s
}

type PauseInstanceRequest struct {
	// The ID of the instance.
	//
	// >  You can call the [DescribeDBInstances](~~86911~~) operation to query the details of all AnalyticDB for PostgreSQL instances in a specific region, including instance IDs.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s PauseInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s PauseInstanceRequest) GoString() string {
	return s.String()
}

func (s *PauseInstanceRequest) SetDBInstanceId(v string) *PauseInstanceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *PauseInstanceRequest) SetOwnerId(v int64) *PauseInstanceRequest {
	s.OwnerId = &v
	return s
}

type PauseInstanceResponseBody struct {
	// The ID of the instance.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The error message returned.
	//
	// This parameter is returned only if **false** is returned for the **Status** parameter.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// *   **false**: The request failed.
	// *   **true**: The request was successful.
	Status *bool `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s PauseInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PauseInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *PauseInstanceResponseBody) SetDBInstanceId(v string) *PauseInstanceResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *PauseInstanceResponseBody) SetErrorMessage(v string) *PauseInstanceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *PauseInstanceResponseBody) SetRequestId(v string) *PauseInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *PauseInstanceResponseBody) SetStatus(v bool) *PauseInstanceResponseBody {
	s.Status = &v
	return s
}

type PauseInstanceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PauseInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PauseInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s PauseInstanceResponse) GoString() string {
	return s.String()
}

func (s *PauseInstanceResponse) SetHeaders(v map[string]*string) *PauseInstanceResponse {
	s.Headers = v
	return s
}

func (s *PauseInstanceResponse) SetStatusCode(v int32) *PauseInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *PauseInstanceResponse) SetBody(v *PauseInstanceResponseBody) *PauseInstanceResponse {
	s.Body = v
	return s
}

type QueryCollectionDataRequest struct {
	Collection        *string    `json:"Collection,omitempty" xml:"Collection,omitempty"`
	Content           *string    `json:"Content,omitempty" xml:"Content,omitempty"`
	DBInstanceId      *string    `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	Filter            *string    `json:"Filter,omitempty" xml:"Filter,omitempty"`
	Namespace         *string    `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	NamespacePassword *string    `json:"NamespacePassword,omitempty" xml:"NamespacePassword,omitempty"`
	OwnerId           *int64     `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId          *string    `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TopK              *int64     `json:"TopK,omitempty" xml:"TopK,omitempty"`
	Vector            []*float64 `json:"Vector,omitempty" xml:"Vector,omitempty" type:"Repeated"`
}

func (s QueryCollectionDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCollectionDataRequest) GoString() string {
	return s.String()
}

func (s *QueryCollectionDataRequest) SetCollection(v string) *QueryCollectionDataRequest {
	s.Collection = &v
	return s
}

func (s *QueryCollectionDataRequest) SetContent(v string) *QueryCollectionDataRequest {
	s.Content = &v
	return s
}

func (s *QueryCollectionDataRequest) SetDBInstanceId(v string) *QueryCollectionDataRequest {
	s.DBInstanceId = &v
	return s
}

func (s *QueryCollectionDataRequest) SetFilter(v string) *QueryCollectionDataRequest {
	s.Filter = &v
	return s
}

func (s *QueryCollectionDataRequest) SetNamespace(v string) *QueryCollectionDataRequest {
	s.Namespace = &v
	return s
}

func (s *QueryCollectionDataRequest) SetNamespacePassword(v string) *QueryCollectionDataRequest {
	s.NamespacePassword = &v
	return s
}

func (s *QueryCollectionDataRequest) SetOwnerId(v int64) *QueryCollectionDataRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryCollectionDataRequest) SetRegionId(v string) *QueryCollectionDataRequest {
	s.RegionId = &v
	return s
}

func (s *QueryCollectionDataRequest) SetTopK(v int64) *QueryCollectionDataRequest {
	s.TopK = &v
	return s
}

func (s *QueryCollectionDataRequest) SetVector(v []*float64) *QueryCollectionDataRequest {
	s.Vector = v
	return s
}

type QueryCollectionDataShrinkRequest struct {
	Collection        *string `json:"Collection,omitempty" xml:"Collection,omitempty"`
	Content           *string `json:"Content,omitempty" xml:"Content,omitempty"`
	DBInstanceId      *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	Filter            *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	Namespace         *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	NamespacePassword *string `json:"NamespacePassword,omitempty" xml:"NamespacePassword,omitempty"`
	OwnerId           *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TopK              *int64  `json:"TopK,omitempty" xml:"TopK,omitempty"`
	VectorShrink      *string `json:"Vector,omitempty" xml:"Vector,omitempty"`
}

func (s QueryCollectionDataShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCollectionDataShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryCollectionDataShrinkRequest) SetCollection(v string) *QueryCollectionDataShrinkRequest {
	s.Collection = &v
	return s
}

func (s *QueryCollectionDataShrinkRequest) SetContent(v string) *QueryCollectionDataShrinkRequest {
	s.Content = &v
	return s
}

func (s *QueryCollectionDataShrinkRequest) SetDBInstanceId(v string) *QueryCollectionDataShrinkRequest {
	s.DBInstanceId = &v
	return s
}

func (s *QueryCollectionDataShrinkRequest) SetFilter(v string) *QueryCollectionDataShrinkRequest {
	s.Filter = &v
	return s
}

func (s *QueryCollectionDataShrinkRequest) SetNamespace(v string) *QueryCollectionDataShrinkRequest {
	s.Namespace = &v
	return s
}

func (s *QueryCollectionDataShrinkRequest) SetNamespacePassword(v string) *QueryCollectionDataShrinkRequest {
	s.NamespacePassword = &v
	return s
}

func (s *QueryCollectionDataShrinkRequest) SetOwnerId(v int64) *QueryCollectionDataShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryCollectionDataShrinkRequest) SetRegionId(v string) *QueryCollectionDataShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *QueryCollectionDataShrinkRequest) SetTopK(v int64) *QueryCollectionDataShrinkRequest {
	s.TopK = &v
	return s
}

func (s *QueryCollectionDataShrinkRequest) SetVectorShrink(v string) *QueryCollectionDataShrinkRequest {
	s.VectorShrink = &v
	return s
}

type QueryCollectionDataResponseBody struct {
	Matches   *QueryCollectionDataResponseBodyMatches `json:"Matches,omitempty" xml:"Matches,omitempty" type:"Struct"`
	Message   *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status    *string                                 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s QueryCollectionDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryCollectionDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCollectionDataResponseBody) SetMatches(v *QueryCollectionDataResponseBodyMatches) *QueryCollectionDataResponseBody {
	s.Matches = v
	return s
}

func (s *QueryCollectionDataResponseBody) SetMessage(v string) *QueryCollectionDataResponseBody {
	s.Message = &v
	return s
}

func (s *QueryCollectionDataResponseBody) SetRequestId(v string) *QueryCollectionDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryCollectionDataResponseBody) SetStatus(v string) *QueryCollectionDataResponseBody {
	s.Status = &v
	return s
}

type QueryCollectionDataResponseBodyMatches struct {
	Match []*QueryCollectionDataResponseBodyMatchesMatch `json:"match,omitempty" xml:"match,omitempty" type:"Repeated"`
}

func (s QueryCollectionDataResponseBodyMatches) String() string {
	return tea.Prettify(s)
}

func (s QueryCollectionDataResponseBodyMatches) GoString() string {
	return s.String()
}

func (s *QueryCollectionDataResponseBodyMatches) SetMatch(v []*QueryCollectionDataResponseBodyMatchesMatch) *QueryCollectionDataResponseBodyMatches {
	s.Match = v
	return s
}

type QueryCollectionDataResponseBodyMatchesMatch struct {
	Id         *string                                            `json:"Id,omitempty" xml:"Id,omitempty"`
	Metadata   map[string]*string                                 `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	Similarity *float64                                           `json:"Similarity,omitempty" xml:"Similarity,omitempty"`
	Values     *QueryCollectionDataResponseBodyMatchesMatchValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Struct"`
}

func (s QueryCollectionDataResponseBodyMatchesMatch) String() string {
	return tea.Prettify(s)
}

func (s QueryCollectionDataResponseBodyMatchesMatch) GoString() string {
	return s.String()
}

func (s *QueryCollectionDataResponseBodyMatchesMatch) SetId(v string) *QueryCollectionDataResponseBodyMatchesMatch {
	s.Id = &v
	return s
}

func (s *QueryCollectionDataResponseBodyMatchesMatch) SetMetadata(v map[string]*string) *QueryCollectionDataResponseBodyMatchesMatch {
	s.Metadata = v
	return s
}

func (s *QueryCollectionDataResponseBodyMatchesMatch) SetSimilarity(v float64) *QueryCollectionDataResponseBodyMatchesMatch {
	s.Similarity = &v
	return s
}

func (s *QueryCollectionDataResponseBodyMatchesMatch) SetValues(v *QueryCollectionDataResponseBodyMatchesMatchValues) *QueryCollectionDataResponseBodyMatchesMatch {
	s.Values = v
	return s
}

type QueryCollectionDataResponseBodyMatchesMatchValues struct {
	Value []*float64 `json:"value,omitempty" xml:"value,omitempty" type:"Repeated"`
}

func (s QueryCollectionDataResponseBodyMatchesMatchValues) String() string {
	return tea.Prettify(s)
}

func (s QueryCollectionDataResponseBodyMatchesMatchValues) GoString() string {
	return s.String()
}

func (s *QueryCollectionDataResponseBodyMatchesMatchValues) SetValue(v []*float64) *QueryCollectionDataResponseBodyMatchesMatchValues {
	s.Value = v
	return s
}

type QueryCollectionDataResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryCollectionDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryCollectionDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCollectionDataResponse) GoString() string {
	return s.String()
}

func (s *QueryCollectionDataResponse) SetHeaders(v map[string]*string) *QueryCollectionDataResponse {
	s.Headers = v
	return s
}

func (s *QueryCollectionDataResponse) SetStatusCode(v int32) *QueryCollectionDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCollectionDataResponse) SetBody(v *QueryCollectionDataResponseBody) *QueryCollectionDataResponse {
	s.Body = v
	return s
}

type RebalanceDBInstanceRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must make sure that it is unique among different requests.
	//
	// The token can be up to 64 characters in length and can contain letters, digits, hyphens (-), and underscores (\_).
	//
	// For more information, see [How to ensure idempotence](~~134212~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The instance ID.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s RebalanceDBInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RebalanceDBInstanceRequest) GoString() string {
	return s.String()
}

func (s *RebalanceDBInstanceRequest) SetClientToken(v string) *RebalanceDBInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *RebalanceDBInstanceRequest) SetDBInstanceId(v string) *RebalanceDBInstanceRequest {
	s.DBInstanceId = &v
	return s
}

type RebalanceDBInstanceResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RebalanceDBInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RebalanceDBInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RebalanceDBInstanceResponseBody) SetRequestId(v string) *RebalanceDBInstanceResponseBody {
	s.RequestId = &v
	return s
}

type RebalanceDBInstanceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RebalanceDBInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RebalanceDBInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s RebalanceDBInstanceResponse) GoString() string {
	return s.String()
}

func (s *RebalanceDBInstanceResponse) SetHeaders(v map[string]*string) *RebalanceDBInstanceResponse {
	s.Headers = v
	return s
}

func (s *RebalanceDBInstanceResponse) SetStatusCode(v int32) *RebalanceDBInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *RebalanceDBInstanceResponse) SetBody(v *RebalanceDBInstanceResponseBody) *RebalanceDBInstanceResponse {
	s.Body = v
	return s
}

type ReleaseInstancePublicConnectionRequest struct {
	// The type of the endpoint. Default value: primary. Valid values:
	//
	// *   **primary**: primary endpoint.
	// *   **cluster**: cluster endpoint. This type of endpoints can be created only for instances that have multiple coordinator nodes.
	AddressType *string `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	// The public endpoint of the instance.
	//
	// You can log on to the AnalyticDB for PostgreSQL console and go to the **Basic Information** page of the instance to view the **public endpoint** in the **Database Connection** section.
	CurrentConnectionString *string `json:"CurrentConnectionString,omitempty" xml:"CurrentConnectionString,omitempty"`
	// The ID of the instance.
	//
	// >  You can call the [DescribeDBInstances](~~86911~~) operation to query the details of all AnalyticDB for PostgreSQL instances in a specific region, including instance IDs.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s ReleaseInstancePublicConnectionRequest) String() string {
	return tea.Prettify(s)
}

func (s ReleaseInstancePublicConnectionRequest) GoString() string {
	return s.String()
}

func (s *ReleaseInstancePublicConnectionRequest) SetAddressType(v string) *ReleaseInstancePublicConnectionRequest {
	s.AddressType = &v
	return s
}

func (s *ReleaseInstancePublicConnectionRequest) SetCurrentConnectionString(v string) *ReleaseInstancePublicConnectionRequest {
	s.CurrentConnectionString = &v
	return s
}

func (s *ReleaseInstancePublicConnectionRequest) SetDBInstanceId(v string) *ReleaseInstancePublicConnectionRequest {
	s.DBInstanceId = &v
	return s
}

type ReleaseInstancePublicConnectionResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReleaseInstancePublicConnectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReleaseInstancePublicConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseInstancePublicConnectionResponseBody) SetRequestId(v string) *ReleaseInstancePublicConnectionResponseBody {
	s.RequestId = &v
	return s
}

type ReleaseInstancePublicConnectionResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ReleaseInstancePublicConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReleaseInstancePublicConnectionResponse) String() string {
	return tea.Prettify(s)
}

func (s ReleaseInstancePublicConnectionResponse) GoString() string {
	return s.String()
}

func (s *ReleaseInstancePublicConnectionResponse) SetHeaders(v map[string]*string) *ReleaseInstancePublicConnectionResponse {
	s.Headers = v
	return s
}

func (s *ReleaseInstancePublicConnectionResponse) SetStatusCode(v int32) *ReleaseInstancePublicConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleaseInstancePublicConnectionResponse) SetBody(v *ReleaseInstancePublicConnectionResponseBody) *ReleaseInstancePublicConnectionResponse {
	s.Body = v
	return s
}

type ResetAccountPasswordRequest struct {
	// The ID of the instance.
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The name of the account.
	AccountPassword *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
	// Before you call this operation, make sure that the following requirements are met:
	//
	// *   The instance is in the running state.
	// *   The instance is not locked.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s ResetAccountPasswordRequest) String() string {
	return tea.Prettify(s)
}

func (s ResetAccountPasswordRequest) GoString() string {
	return s.String()
}

func (s *ResetAccountPasswordRequest) SetAccountName(v string) *ResetAccountPasswordRequest {
	s.AccountName = &v
	return s
}

func (s *ResetAccountPasswordRequest) SetAccountPassword(v string) *ResetAccountPasswordRequest {
	s.AccountPassword = &v
	return s
}

func (s *ResetAccountPasswordRequest) SetDBInstanceId(v string) *ResetAccountPasswordRequest {
	s.DBInstanceId = &v
	return s
}

type ResetAccountPasswordResponseBody struct {
	// The new password for the account. The password must be 8 to 32 characters in length and contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include `! @ # $ % ^ & * ( ) _ + - =`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetAccountPasswordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResetAccountPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *ResetAccountPasswordResponseBody) SetRequestId(v string) *ResetAccountPasswordResponseBody {
	s.RequestId = &v
	return s
}

type ResetAccountPasswordResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ResetAccountPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ResetAccountPasswordResponse) String() string {
	return tea.Prettify(s)
}

func (s ResetAccountPasswordResponse) GoString() string {
	return s.String()
}

func (s *ResetAccountPasswordResponse) SetHeaders(v map[string]*string) *ResetAccountPasswordResponse {
	s.Headers = v
	return s
}

func (s *ResetAccountPasswordResponse) SetStatusCode(v int32) *ResetAccountPasswordResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetAccountPasswordResponse) SetBody(v *ResetAccountPasswordResponseBody) *ResetAccountPasswordResponse {
	s.Body = v
	return s
}

type RestartDBInstanceRequest struct {
	// The client token that is used to ensure the idempotence of the request. For more information, see [How to ensure idempotence](~~327176~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the instance.
	//
	// >  You can call the [DescribeDBInstances](~~86911~~) operation to query the details of all AnalyticDB for PostgreSQL instances in a specific region, including instance IDs.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s RestartDBInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RestartDBInstanceRequest) GoString() string {
	return s.String()
}

func (s *RestartDBInstanceRequest) SetClientToken(v string) *RestartDBInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *RestartDBInstanceRequest) SetDBInstanceId(v string) *RestartDBInstanceRequest {
	s.DBInstanceId = &v
	return s
}

type RestartDBInstanceResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RestartDBInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RestartDBInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RestartDBInstanceResponseBody) SetRequestId(v string) *RestartDBInstanceResponseBody {
	s.RequestId = &v
	return s
}

type RestartDBInstanceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RestartDBInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RestartDBInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s RestartDBInstanceResponse) GoString() string {
	return s.String()
}

func (s *RestartDBInstanceResponse) SetHeaders(v map[string]*string) *RestartDBInstanceResponse {
	s.Headers = v
	return s
}

func (s *RestartDBInstanceResponse) SetStatusCode(v int32) *RestartDBInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *RestartDBInstanceResponse) SetBody(v *RestartDBInstanceResponseBody) *RestartDBInstanceResponse {
	s.Body = v
	return s
}

type ResumeInstanceRequest struct {
	// The ID of the instance.
	//
	// >  You can call the [DescribeDBInstances](~~86911~~) operation to query the details of all AnalyticDB for PostgreSQL instances in a specific region, including instance IDs.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s ResumeInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ResumeInstanceRequest) GoString() string {
	return s.String()
}

func (s *ResumeInstanceRequest) SetDBInstanceId(v string) *ResumeInstanceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ResumeInstanceRequest) SetOwnerId(v int64) *ResumeInstanceRequest {
	s.OwnerId = &v
	return s
}

type ResumeInstanceResponseBody struct {
	// The ID of the instance.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The error message returned.
	//
	// This parameter is returned only if **false** is returned for the **Status** parameter.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// *   **false**: The request failed.
	// *   **true**: The request was successful.
	Status *bool `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ResumeInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResumeInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ResumeInstanceResponseBody) SetDBInstanceId(v string) *ResumeInstanceResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *ResumeInstanceResponseBody) SetErrorMessage(v string) *ResumeInstanceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ResumeInstanceResponseBody) SetRequestId(v string) *ResumeInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResumeInstanceResponseBody) SetStatus(v bool) *ResumeInstanceResponseBody {
	s.Status = &v
	return s
}

type ResumeInstanceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ResumeInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ResumeInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s ResumeInstanceResponse) GoString() string {
	return s.String()
}

func (s *ResumeInstanceResponse) SetHeaders(v map[string]*string) *ResumeInstanceResponse {
	s.Headers = v
	return s
}

func (s *ResumeInstanceResponse) SetStatusCode(v int32) *ResumeInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ResumeInstanceResponse) SetBody(v *ResumeInstanceResponseBody) *ResumeInstanceResponse {
	s.Body = v
	return s
}

type SetDBInstancePlanStatusRequest struct {
	// The ID of the instance.
	//
	// >  You can call the [DescribeDBInstances](~~86911~~) operation to query the details of all AnalyticDB for PostgreSQL instances in a specific region, including instance IDs.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the plan.
	//
	// >  You can call the [DescribeDBInstancePlans](~~449398~~) operation to query the details of plans, including plan IDs.
	PlanId *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	// Specifies whether to enable or disable the plan. Valid values:
	//
	// *   **disable**: disables the plan.
	// *   **enable**: enables the plan.
	PlanStatus *string `json:"PlanStatus,omitempty" xml:"PlanStatus,omitempty"`
}

func (s SetDBInstancePlanStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s SetDBInstancePlanStatusRequest) GoString() string {
	return s.String()
}

func (s *SetDBInstancePlanStatusRequest) SetDBInstanceId(v string) *SetDBInstancePlanStatusRequest {
	s.DBInstanceId = &v
	return s
}

func (s *SetDBInstancePlanStatusRequest) SetOwnerId(v int64) *SetDBInstancePlanStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *SetDBInstancePlanStatusRequest) SetPlanId(v string) *SetDBInstancePlanStatusRequest {
	s.PlanId = &v
	return s
}

func (s *SetDBInstancePlanStatusRequest) SetPlanStatus(v string) *SetDBInstancePlanStatusRequest {
	s.PlanStatus = &v
	return s
}

type SetDBInstancePlanStatusResponseBody struct {
	// The error message returned.
	//
	// This parameter is returned only when the operation fails.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the plan.
	PlanId *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The state of the operation.
	//
	// If the operation is successful, **success** is returned. If the operation fails, this parameter is not returned.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s SetDBInstancePlanStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetDBInstancePlanStatusResponseBody) GoString() string {
	return s.String()
}

func (s *SetDBInstancePlanStatusResponseBody) SetErrorMessage(v string) *SetDBInstancePlanStatusResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *SetDBInstancePlanStatusResponseBody) SetPlanId(v string) *SetDBInstancePlanStatusResponseBody {
	s.PlanId = &v
	return s
}

func (s *SetDBInstancePlanStatusResponseBody) SetRequestId(v string) *SetDBInstancePlanStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetDBInstancePlanStatusResponseBody) SetStatus(v string) *SetDBInstancePlanStatusResponseBody {
	s.Status = &v
	return s
}

type SetDBInstancePlanStatusResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetDBInstancePlanStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetDBInstancePlanStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s SetDBInstancePlanStatusResponse) GoString() string {
	return s.String()
}

func (s *SetDBInstancePlanStatusResponse) SetHeaders(v map[string]*string) *SetDBInstancePlanStatusResponse {
	s.Headers = v
	return s
}

func (s *SetDBInstancePlanStatusResponse) SetStatusCode(v int32) *SetDBInstancePlanStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *SetDBInstancePlanStatusResponse) SetBody(v *SetDBInstancePlanStatusResponseBody) *SetDBInstancePlanStatusResponse {
	s.Body = v
	return s
}

type SetDataShareInstanceRequest struct {
	// The ID of the AnalyticDB for PostgreSQL instance in Serverless mode.
	//
	// >  You can call the [DescribeDBInstances](~~86911~~) operation to query the details of all AnalyticDB for PostgreSQL instances in a specific region, including instance IDs.
	InstanceList []*string `json:"InstanceList,omitempty" xml:"InstanceList,omitempty" type:"Repeated"`
	// Specifies whether to enable or disable data sharing. Valid values:
	//
	// *   **add**: enables data sharing.
	// *   **remove**: disables data sharing.
	OperationType *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region.
	//
	// >  You can call the [DescribeRegions](~~86912~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s SetDataShareInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s SetDataShareInstanceRequest) GoString() string {
	return s.String()
}

func (s *SetDataShareInstanceRequest) SetInstanceList(v []*string) *SetDataShareInstanceRequest {
	s.InstanceList = v
	return s
}

func (s *SetDataShareInstanceRequest) SetOperationType(v string) *SetDataShareInstanceRequest {
	s.OperationType = &v
	return s
}

func (s *SetDataShareInstanceRequest) SetOwnerId(v int64) *SetDataShareInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *SetDataShareInstanceRequest) SetRegionId(v string) *SetDataShareInstanceRequest {
	s.RegionId = &v
	return s
}

type SetDataShareInstanceShrinkRequest struct {
	// The ID of the AnalyticDB for PostgreSQL instance in Serverless mode.
	//
	// >  You can call the [DescribeDBInstances](~~86911~~) operation to query the details of all AnalyticDB for PostgreSQL instances in a specific region, including instance IDs.
	InstanceListShrink *string `json:"InstanceList,omitempty" xml:"InstanceList,omitempty"`
	// Specifies whether to enable or disable data sharing. Valid values:
	//
	// *   **add**: enables data sharing.
	// *   **remove**: disables data sharing.
	OperationType *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region.
	//
	// >  You can call the [DescribeRegions](~~86912~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s SetDataShareInstanceShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SetDataShareInstanceShrinkRequest) GoString() string {
	return s.String()
}

func (s *SetDataShareInstanceShrinkRequest) SetInstanceListShrink(v string) *SetDataShareInstanceShrinkRequest {
	s.InstanceListShrink = &v
	return s
}

func (s *SetDataShareInstanceShrinkRequest) SetOperationType(v string) *SetDataShareInstanceShrinkRequest {
	s.OperationType = &v
	return s
}

func (s *SetDataShareInstanceShrinkRequest) SetOwnerId(v int64) *SetDataShareInstanceShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *SetDataShareInstanceShrinkRequest) SetRegionId(v string) *SetDataShareInstanceShrinkRequest {
	s.RegionId = &v
	return s
}

type SetDataShareInstanceResponseBody struct {
	// The error message returned if the operation fails.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The state of the operation. Valid values:
	//
	// *   **success**: The operation is successful.
	// *   **failed**: The operation fails.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s SetDataShareInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetDataShareInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *SetDataShareInstanceResponseBody) SetErrMessage(v string) *SetDataShareInstanceResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *SetDataShareInstanceResponseBody) SetRequestId(v string) *SetDataShareInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetDataShareInstanceResponseBody) SetStatus(v string) *SetDataShareInstanceResponseBody {
	s.Status = &v
	return s
}

type SetDataShareInstanceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetDataShareInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetDataShareInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s SetDataShareInstanceResponse) GoString() string {
	return s.String()
}

func (s *SetDataShareInstanceResponse) SetHeaders(v map[string]*string) *SetDataShareInstanceResponse {
	s.Headers = v
	return s
}

func (s *SetDataShareInstanceResponse) SetStatusCode(v int32) *SetDataShareInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *SetDataShareInstanceResponse) SetBody(v *SetDataShareInstanceResponseBody) *SetDataShareInstanceResponse {
	s.Body = v
	return s
}

type SwitchDBInstanceNetTypeRequest struct {
	// The prefix of the custom endpoint. The prefix must be 8 to 64 characters in length and can contain letters and digits. It must start with a lowercase letter. A valid endpoint is in the following format: Prefix.Database engine.rds.aliyuncs.com. Example: test1234.mysql.rds.aliyuncs.com.
	ConnectionStringPrefix *string `json:"ConnectionStringPrefix,omitempty" xml:"ConnectionStringPrefix,omitempty"`
	// The ID of the instance.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The port number. Valid values: 3000 to 5999.
	//
	// >
	// *   Only ApsaraDB PolarDB MySQL-compatible edition clusters support this parameter. If you leave this parameter empty, the default port 3306 is used.
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s SwitchDBInstanceNetTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s SwitchDBInstanceNetTypeRequest) GoString() string {
	return s.String()
}

func (s *SwitchDBInstanceNetTypeRequest) SetConnectionStringPrefix(v string) *SwitchDBInstanceNetTypeRequest {
	s.ConnectionStringPrefix = &v
	return s
}

func (s *SwitchDBInstanceNetTypeRequest) SetDBInstanceId(v string) *SwitchDBInstanceNetTypeRequest {
	s.DBInstanceId = &v
	return s
}

func (s *SwitchDBInstanceNetTypeRequest) SetPort(v string) *SwitchDBInstanceNetTypeRequest {
	s.Port = &v
	return s
}

type SwitchDBInstanceNetTypeResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SwitchDBInstanceNetTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SwitchDBInstanceNetTypeResponseBody) GoString() string {
	return s.String()
}

func (s *SwitchDBInstanceNetTypeResponseBody) SetRequestId(v string) *SwitchDBInstanceNetTypeResponseBody {
	s.RequestId = &v
	return s
}

type SwitchDBInstanceNetTypeResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SwitchDBInstanceNetTypeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SwitchDBInstanceNetTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s SwitchDBInstanceNetTypeResponse) GoString() string {
	return s.String()
}

func (s *SwitchDBInstanceNetTypeResponse) SetHeaders(v map[string]*string) *SwitchDBInstanceNetTypeResponse {
	s.Headers = v
	return s
}

func (s *SwitchDBInstanceNetTypeResponse) SetStatusCode(v int32) *SwitchDBInstanceNetTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *SwitchDBInstanceNetTypeResponse) SetBody(v *SwitchDBInstanceNetTypeResponseBody) *SwitchDBInstanceNetTypeResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region. You can call the [DescribeRegions](~~86912~~) operation to query region IDs.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of an instance. Valid values of N: 1 to 50.
	ResourceId           []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The mode of the instance. Valid values:
	//
	// *   `instance`: reserved storage mode
	// *   `ALIYUN::GPDB::INSTANCE`: elastic storage mode
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The list of tags.
	Tag []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s TagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) SetOwnerAccount(v string) *TagResourcesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *TagResourcesRequest) SetOwnerId(v int64) *TagResourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *TagResourcesRequest) SetRegionId(v string) *TagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceId(v []*string) *TagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *TagResourcesRequest) SetResourceOwnerAccount(v string) *TagResourcesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *TagResourcesRequest) SetResourceOwnerId(v int64) *TagResourcesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceType(v string) *TagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesRequest) SetTag(v []*TagResourcesRequestTag) *TagResourcesRequest {
	s.Tag = v
	return s
}

type TagResourcesRequestTag struct {
	// The key of a tag. Valid values of N: 1 to 20. This parameter value cannot be an empty string. A tag key can contain a maximum of 128 characters. It cannot start with `aliyun` or`  acs: ` and cannot contain `http://` or`  https:// `.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of a tag. Valid values of N: 1 to 20. This parameter value can be an empty string. A tag value can contain a maximum of 128 characters. It cannot start with `acs:` and cannot contain `http://` or `https://`.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s TagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *TagResourcesRequestTag) SetKey(v string) *TagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *TagResourcesRequestTag) SetValue(v string) *TagResourcesRequestTag {
	s.Value = &v
	return s
}

type TagResourcesResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *TagResourcesResponseBody) SetRequestId(v string) *TagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type TagResourcesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *TagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponse) GoString() string {
	return s.String()
}

func (s *TagResourcesResponse) SetHeaders(v map[string]*string) *TagResourcesResponse {
	s.Headers = v
	return s
}

func (s *TagResourcesResponse) SetStatusCode(v int32) *TagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *TagResourcesResponse) SetBody(v *TagResourcesResponseBody) *TagResourcesResponse {
	s.Body = v
	return s
}

type UnloadSampleDataRequest struct {
	// The ID of the instance.
	//
	// >  You can call the [DescribeDBInstances](~~86911~~) operation to query the IDs of all AnalyticDB for PostgreSQL instances in a specific region.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s UnloadSampleDataRequest) String() string {
	return tea.Prettify(s)
}

func (s UnloadSampleDataRequest) GoString() string {
	return s.String()
}

func (s *UnloadSampleDataRequest) SetDBInstanceId(v string) *UnloadSampleDataRequest {
	s.DBInstanceId = &v
	return s
}

func (s *UnloadSampleDataRequest) SetOwnerId(v int64) *UnloadSampleDataRequest {
	s.OwnerId = &v
	return s
}

type UnloadSampleDataResponseBody struct {
	// The ID of the instance.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The error message returned if an error occurs. This message does not affect the execution of the operation.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The execution state of the operation. Valid values:
	//
	// *   **false**: The operation fails.
	// *   **true**: The operation is successful.
	Status *bool `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UnloadSampleDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnloadSampleDataResponseBody) GoString() string {
	return s.String()
}

func (s *UnloadSampleDataResponseBody) SetDBInstanceId(v string) *UnloadSampleDataResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *UnloadSampleDataResponseBody) SetErrorMessage(v string) *UnloadSampleDataResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UnloadSampleDataResponseBody) SetRequestId(v string) *UnloadSampleDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnloadSampleDataResponseBody) SetStatus(v bool) *UnloadSampleDataResponseBody {
	s.Status = &v
	return s
}

type UnloadSampleDataResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UnloadSampleDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnloadSampleDataResponse) String() string {
	return tea.Prettify(s)
}

func (s UnloadSampleDataResponse) GoString() string {
	return s.String()
}

func (s *UnloadSampleDataResponse) SetHeaders(v map[string]*string) *UnloadSampleDataResponse {
	s.Headers = v
	return s
}

func (s *UnloadSampleDataResponse) SetStatusCode(v int32) *UnloadSampleDataResponse {
	s.StatusCode = &v
	return s
}

func (s *UnloadSampleDataResponse) SetBody(v *UnloadSampleDataResponseBody) *UnloadSampleDataResponse {
	s.Body = v
	return s
}

type UntagResourcesRequest struct {
	// Specifies whether to unbind all tags from an instance. This parameter is valid only when the TagKey.N parameter is not specified. Valid values:
	//
	// *   true
	// *   false
	//
	// Default value: false.
	All          *bool   `json:"All,omitempty" xml:"All,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the instance. You can call the [DescribeRegions](~~86912~~) operation to query the most recent region list.
	RegionId             *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceId           []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The storage mode of the instance. Valid values:
	//
	// *   `instance`: reserved storage mode
	// *   `ALIYUN::GPDB::INSTANCE`: elastic storage mode
	ResourceType *string   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagKey       []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s UntagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesRequest) SetAll(v bool) *UntagResourcesRequest {
	s.All = &v
	return s
}

func (s *UntagResourcesRequest) SetOwnerAccount(v string) *UntagResourcesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UntagResourcesRequest) SetOwnerId(v int64) *UntagResourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *UntagResourcesRequest) SetRegionId(v string) *UntagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceId(v []*string) *UntagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *UntagResourcesRequest) SetResourceOwnerAccount(v string) *UntagResourcesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceOwnerId(v int64) *UntagResourcesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceType(v string) *UntagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesRequest) SetTagKey(v []*string) *UntagResourcesRequest {
	s.TagKey = v
	return s
}

type UntagResourcesResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UntagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponseBody) SetRequestId(v string) *UntagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type UntagResourcesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UntagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UntagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponse) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponse) SetHeaders(v map[string]*string) *UntagResourcesResponse {
	s.Headers = v
	return s
}

func (s *UntagResourcesResponse) SetStatusCode(v int32) *UntagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *UntagResourcesResponse) SetBody(v *UntagResourcesResponseBody) *UntagResourcesResponse {
	s.Body = v
	return s
}

type UpdateDBInstancePlanRequest struct {
	// The ID of the instance.
	//
	// >  You can call the [DescribeDBInstances](~~86911~~) operation to query the details of all AnalyticDB for PostgreSQL instances in a specific region, including instance IDs.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The execution information of the plan. Specify the parameter in the JSON format. The parameter value varies based on the values of **PlanType** and **PlanScheduleType**. The following section describes the PlanConfig parameter.
	PlanConfig *string `json:"PlanConfig,omitempty" xml:"PlanConfig,omitempty"`
	// The description of the plan.
	PlanDesc *string `json:"PlanDesc,omitempty" xml:"PlanDesc,omitempty"`
	// The end time of the plan. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm*Z format. The time must be in UTC. The end time must be later than the start time.
	//
	// >  This parameter is required only for **periodically executed** plans.
	PlanEndDate *string `json:"PlanEndDate,omitempty" xml:"PlanEndDate,omitempty"`
	// The ID of the plan.
	//
	// >  You can call the [DescribeDBInstancePlans](~~449398~~) operation to query the details of plans, including plan IDs.
	PlanId *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	// The name of the plan.
	PlanName *string `json:"PlanName,omitempty" xml:"PlanName,omitempty"`
	// The start time of the plan. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm*Z format. The time is displayed in UTC.
	//
	// >  This parameter is required only for **periodically executed** plans.
	PlanStartDate *string `json:"PlanStartDate,omitempty" xml:"PlanStartDate,omitempty"`
}

func (s UpdateDBInstancePlanRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDBInstancePlanRequest) GoString() string {
	return s.String()
}

func (s *UpdateDBInstancePlanRequest) SetDBInstanceId(v string) *UpdateDBInstancePlanRequest {
	s.DBInstanceId = &v
	return s
}

func (s *UpdateDBInstancePlanRequest) SetOwnerId(v int64) *UpdateDBInstancePlanRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateDBInstancePlanRequest) SetPlanConfig(v string) *UpdateDBInstancePlanRequest {
	s.PlanConfig = &v
	return s
}

func (s *UpdateDBInstancePlanRequest) SetPlanDesc(v string) *UpdateDBInstancePlanRequest {
	s.PlanDesc = &v
	return s
}

func (s *UpdateDBInstancePlanRequest) SetPlanEndDate(v string) *UpdateDBInstancePlanRequest {
	s.PlanEndDate = &v
	return s
}

func (s *UpdateDBInstancePlanRequest) SetPlanId(v string) *UpdateDBInstancePlanRequest {
	s.PlanId = &v
	return s
}

func (s *UpdateDBInstancePlanRequest) SetPlanName(v string) *UpdateDBInstancePlanRequest {
	s.PlanName = &v
	return s
}

func (s *UpdateDBInstancePlanRequest) SetPlanStartDate(v string) *UpdateDBInstancePlanRequest {
	s.PlanStartDate = &v
	return s
}

type UpdateDBInstancePlanResponseBody struct {
	// The ID of the instance.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The error message returned.
	//
	// This parameter is returned only when the operation fails.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the plan.
	PlanId *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The state of the operation.
	//
	// If the operation is successful, **success** is returned. If the operation fails, this parameter is not returned.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateDBInstancePlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDBInstancePlanResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDBInstancePlanResponseBody) SetDBInstanceId(v string) *UpdateDBInstancePlanResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *UpdateDBInstancePlanResponseBody) SetErrorMessage(v string) *UpdateDBInstancePlanResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateDBInstancePlanResponseBody) SetPlanId(v string) *UpdateDBInstancePlanResponseBody {
	s.PlanId = &v
	return s
}

func (s *UpdateDBInstancePlanResponseBody) SetRequestId(v string) *UpdateDBInstancePlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDBInstancePlanResponseBody) SetStatus(v string) *UpdateDBInstancePlanResponseBody {
	s.Status = &v
	return s
}

type UpdateDBInstancePlanResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateDBInstancePlanResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateDBInstancePlanResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDBInstancePlanResponse) GoString() string {
	return s.String()
}

func (s *UpdateDBInstancePlanResponse) SetHeaders(v map[string]*string) *UpdateDBInstancePlanResponse {
	s.Headers = v
	return s
}

func (s *UpdateDBInstancePlanResponse) SetStatusCode(v int32) *UpdateDBInstancePlanResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDBInstancePlanResponse) SetBody(v *UpdateDBInstancePlanResponseBody) *UpdateDBInstancePlanResponse {
	s.Body = v
	return s
}

type UpgradeDBInstanceRequest struct {
	// This parameter is no longer used.
	DBInstanceClass *string `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty"`
	// This parameter is no longer used.
	DBInstanceGroupCount *string `json:"DBInstanceGroupCount,omitempty" xml:"DBInstanceGroupCount,omitempty"`
	// The instance ID.
	//
	// >  You can call the [DescribeDBInstances](~~86911~~) operation to query the IDs of all AnalyticDB for PostgreSQL instances in a region.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The specifications of each compute node. For information about the supported specifications, see [Instance specifications](~~35406~~).
	//
	// >  This parameter is available only for instances in elastic storage mode.
	InstanceSpec *string `json:"InstanceSpec,omitempty" xml:"InstanceSpec,omitempty"`
	// The number of coordinator nodes. Valid values: 1 and 2.
	//
	// >  This parameter is available only on the China site (aliyun.com).
	MasterNodeNum *string `json:"MasterNodeNum,omitempty" xml:"MasterNodeNum,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is no longer used.
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The region ID of the instance.
	//
	// >  You can call the [DescribeRegions](~~86912~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the instance belongs. For more information about how to obtain the ID of a resource group, see [View basic information of a resource group](~~151181~~).
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The performance level of enhanced SSDs (ESSDs). Valid values:
	//
	// *   **pl0**
	// *   **pl1**
	// *   **pl2**
	SegDiskPerformanceLevel *string `json:"SegDiskPerformanceLevel,omitempty" xml:"SegDiskPerformanceLevel,omitempty"`
	// The number of compute nodes. The number of compute nodes varies based on the instance resource type and edition.
	//
	// *   Valid values for High-availability Edition instances in elastic storage mode: 4 to 512, in 4 increments
	// *   Valid values for High-performance Edition instances in elastic storage mode: 2 to 512, in 2 increments
	// *   Valid values for instances in manual Serverless mode: 2 to 512, in 2 increments
	SegNodeNum *string `json:"SegNodeNum,omitempty" xml:"SegNodeNum,omitempty"`
	// The disk storage type of the instance after the change. The disk storage type can be changed only to ESSD. Set the value to **cloud_essd**.
	SegStorageType *string `json:"SegStorageType,omitempty" xml:"SegStorageType,omitempty"`
	// The storage capacity of each compute node. Unit: GB. Valid values: 50 to 6000, in 50 increments.
	//
	// >  This parameter is available only for instances in elastic storage mode.
	StorageSize *string `json:"StorageSize,omitempty" xml:"StorageSize,omitempty"`
	// The type of the instance configuration change. Valid values:
	//
	// *   **0** (default): changes the number of compute nodes.
	// *   **1**: changes the specifications and storage capacity of each compute node.
	// *   **2**: changes the number of coordinator nodes.
	// *   **3**: changes the disk storage type and ESSD performance level of the instance.
	//
	// > *   The supported changes to compute node configurations vary based on the instance resource type. For more information, see the "[Precautions](~~50956~~)" section of the Change compute node configurations topic.
	// > *   After you specify a change type, only the corresponding parameters take effect. For example, if you set **UpgradeType** to 0, the parameter that is used to change the number of compute nodes takes effect, but the parameter that is used to change the number of coordinator nodes does not.
	// > *   The number of coordinator nodes can be changed only on the China site (aliyun.com).
	// > *   The disk storage type can be changed only from ultra disks to ESSDs.
	UpgradeType *int64 `json:"UpgradeType,omitempty" xml:"UpgradeType,omitempty"`
}

func (s UpgradeDBInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpgradeDBInstanceRequest) GoString() string {
	return s.String()
}

func (s *UpgradeDBInstanceRequest) SetDBInstanceClass(v string) *UpgradeDBInstanceRequest {
	s.DBInstanceClass = &v
	return s
}

func (s *UpgradeDBInstanceRequest) SetDBInstanceGroupCount(v string) *UpgradeDBInstanceRequest {
	s.DBInstanceGroupCount = &v
	return s
}

func (s *UpgradeDBInstanceRequest) SetDBInstanceId(v string) *UpgradeDBInstanceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *UpgradeDBInstanceRequest) SetInstanceSpec(v string) *UpgradeDBInstanceRequest {
	s.InstanceSpec = &v
	return s
}

func (s *UpgradeDBInstanceRequest) SetMasterNodeNum(v string) *UpgradeDBInstanceRequest {
	s.MasterNodeNum = &v
	return s
}

func (s *UpgradeDBInstanceRequest) SetOwnerId(v int64) *UpgradeDBInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *UpgradeDBInstanceRequest) SetPayType(v string) *UpgradeDBInstanceRequest {
	s.PayType = &v
	return s
}

func (s *UpgradeDBInstanceRequest) SetRegionId(v string) *UpgradeDBInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *UpgradeDBInstanceRequest) SetResourceGroupId(v string) *UpgradeDBInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UpgradeDBInstanceRequest) SetSegDiskPerformanceLevel(v string) *UpgradeDBInstanceRequest {
	s.SegDiskPerformanceLevel = &v
	return s
}

func (s *UpgradeDBInstanceRequest) SetSegNodeNum(v string) *UpgradeDBInstanceRequest {
	s.SegNodeNum = &v
	return s
}

func (s *UpgradeDBInstanceRequest) SetSegStorageType(v string) *UpgradeDBInstanceRequest {
	s.SegStorageType = &v
	return s
}

func (s *UpgradeDBInstanceRequest) SetStorageSize(v string) *UpgradeDBInstanceRequest {
	s.StorageSize = &v
	return s
}

func (s *UpgradeDBInstanceRequest) SetUpgradeType(v int64) *UpgradeDBInstanceRequest {
	s.UpgradeType = &v
	return s
}

type UpgradeDBInstanceResponseBody struct {
	// The instance ID.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The order ID.
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpgradeDBInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpgradeDBInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeDBInstanceResponseBody) SetDBInstanceId(v string) *UpgradeDBInstanceResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *UpgradeDBInstanceResponseBody) SetOrderId(v string) *UpgradeDBInstanceResponseBody {
	s.OrderId = &v
	return s
}

func (s *UpgradeDBInstanceResponseBody) SetRequestId(v string) *UpgradeDBInstanceResponseBody {
	s.RequestId = &v
	return s
}

type UpgradeDBInstanceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpgradeDBInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpgradeDBInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpgradeDBInstanceResponse) GoString() string {
	return s.String()
}

func (s *UpgradeDBInstanceResponse) SetHeaders(v map[string]*string) *UpgradeDBInstanceResponse {
	s.Headers = v
	return s
}

func (s *UpgradeDBInstanceResponse) SetStatusCode(v int32) *UpgradeDBInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeDBInstanceResponse) SetBody(v *UpgradeDBInstanceResponseBody) *UpgradeDBInstanceResponse {
	s.Body = v
	return s
}

type UpgradeDBVersionRequest struct {
	// The ID of the instance.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The major version of the instance.
	MajorVersion *string `json:"MajorVersion,omitempty" xml:"MajorVersion,omitempty"`
	// The minor version of the instance.
	MinorVersion *string `json:"MinorVersion,omitempty" xml:"MinorVersion,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the instance.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The upgrade time.
	SwitchTime *string `json:"SwitchTime,omitempty" xml:"SwitchTime,omitempty"`
	// The upgrade method.
	SwitchTimeMode *string `json:"SwitchTimeMode,omitempty" xml:"SwitchTimeMode,omitempty"`
}

func (s UpgradeDBVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpgradeDBVersionRequest) GoString() string {
	return s.String()
}

func (s *UpgradeDBVersionRequest) SetDBInstanceId(v string) *UpgradeDBVersionRequest {
	s.DBInstanceId = &v
	return s
}

func (s *UpgradeDBVersionRequest) SetMajorVersion(v string) *UpgradeDBVersionRequest {
	s.MajorVersion = &v
	return s
}

func (s *UpgradeDBVersionRequest) SetMinorVersion(v string) *UpgradeDBVersionRequest {
	s.MinorVersion = &v
	return s
}

func (s *UpgradeDBVersionRequest) SetOwnerId(v int64) *UpgradeDBVersionRequest {
	s.OwnerId = &v
	return s
}

func (s *UpgradeDBVersionRequest) SetRegionId(v string) *UpgradeDBVersionRequest {
	s.RegionId = &v
	return s
}

func (s *UpgradeDBVersionRequest) SetSwitchTime(v string) *UpgradeDBVersionRequest {
	s.SwitchTime = &v
	return s
}

func (s *UpgradeDBVersionRequest) SetSwitchTimeMode(v string) *UpgradeDBVersionRequest {
	s.SwitchTimeMode = &v
	return s
}

type UpgradeDBVersionResponseBody struct {
	// The ID of the instance.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The name of the instance.
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the task.
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s UpgradeDBVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpgradeDBVersionResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeDBVersionResponseBody) SetDBInstanceId(v string) *UpgradeDBVersionResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *UpgradeDBVersionResponseBody) SetDBInstanceName(v string) *UpgradeDBVersionResponseBody {
	s.DBInstanceName = &v
	return s
}

func (s *UpgradeDBVersionResponseBody) SetRequestId(v string) *UpgradeDBVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeDBVersionResponseBody) SetTaskId(v string) *UpgradeDBVersionResponseBody {
	s.TaskId = &v
	return s
}

type UpgradeDBVersionResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpgradeDBVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpgradeDBVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpgradeDBVersionResponse) GoString() string {
	return s.String()
}

func (s *UpgradeDBVersionResponse) SetHeaders(v map[string]*string) *UpgradeDBVersionResponse {
	s.Headers = v
	return s
}

func (s *UpgradeDBVersionResponse) SetStatusCode(v int32) *UpgradeDBVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeDBVersionResponse) SetBody(v *UpgradeDBVersionResponseBody) *UpgradeDBVersionResponse {
	s.Body = v
	return s
}

type UpsertCollectionDataRequest struct {
	Collection        *string                            `json:"Collection,omitempty" xml:"Collection,omitempty"`
	DBInstanceId      *string                            `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	Namespace         *string                            `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	NamespacePassword *string                            `json:"NamespacePassword,omitempty" xml:"NamespacePassword,omitempty"`
	OwnerId           *int64                             `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId          *string                            `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Rows              []*UpsertCollectionDataRequestRows `json:"Rows,omitempty" xml:"Rows,omitempty" type:"Repeated"`
}

func (s UpsertCollectionDataRequest) String() string {
	return tea.Prettify(s)
}

func (s UpsertCollectionDataRequest) GoString() string {
	return s.String()
}

func (s *UpsertCollectionDataRequest) SetCollection(v string) *UpsertCollectionDataRequest {
	s.Collection = &v
	return s
}

func (s *UpsertCollectionDataRequest) SetDBInstanceId(v string) *UpsertCollectionDataRequest {
	s.DBInstanceId = &v
	return s
}

func (s *UpsertCollectionDataRequest) SetNamespace(v string) *UpsertCollectionDataRequest {
	s.Namespace = &v
	return s
}

func (s *UpsertCollectionDataRequest) SetNamespacePassword(v string) *UpsertCollectionDataRequest {
	s.NamespacePassword = &v
	return s
}

func (s *UpsertCollectionDataRequest) SetOwnerId(v int64) *UpsertCollectionDataRequest {
	s.OwnerId = &v
	return s
}

func (s *UpsertCollectionDataRequest) SetRegionId(v string) *UpsertCollectionDataRequest {
	s.RegionId = &v
	return s
}

func (s *UpsertCollectionDataRequest) SetRows(v []*UpsertCollectionDataRequestRows) *UpsertCollectionDataRequest {
	s.Rows = v
	return s
}

type UpsertCollectionDataRequestRows struct {
	Id       *string            `json:"Id,omitempty" xml:"Id,omitempty"`
	Metadata map[string]*string `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	Vector   []*float64         `json:"Vector,omitempty" xml:"Vector,omitempty" type:"Repeated"`
}

func (s UpsertCollectionDataRequestRows) String() string {
	return tea.Prettify(s)
}

func (s UpsertCollectionDataRequestRows) GoString() string {
	return s.String()
}

func (s *UpsertCollectionDataRequestRows) SetId(v string) *UpsertCollectionDataRequestRows {
	s.Id = &v
	return s
}

func (s *UpsertCollectionDataRequestRows) SetMetadata(v map[string]*string) *UpsertCollectionDataRequestRows {
	s.Metadata = v
	return s
}

func (s *UpsertCollectionDataRequestRows) SetVector(v []*float64) *UpsertCollectionDataRequestRows {
	s.Vector = v
	return s
}

type UpsertCollectionDataShrinkRequest struct {
	Collection        *string `json:"Collection,omitempty" xml:"Collection,omitempty"`
	DBInstanceId      *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	Namespace         *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	NamespacePassword *string `json:"NamespacePassword,omitempty" xml:"NamespacePassword,omitempty"`
	OwnerId           *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RowsShrink        *string `json:"Rows,omitempty" xml:"Rows,omitempty"`
}

func (s UpsertCollectionDataShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpsertCollectionDataShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpsertCollectionDataShrinkRequest) SetCollection(v string) *UpsertCollectionDataShrinkRequest {
	s.Collection = &v
	return s
}

func (s *UpsertCollectionDataShrinkRequest) SetDBInstanceId(v string) *UpsertCollectionDataShrinkRequest {
	s.DBInstanceId = &v
	return s
}

func (s *UpsertCollectionDataShrinkRequest) SetNamespace(v string) *UpsertCollectionDataShrinkRequest {
	s.Namespace = &v
	return s
}

func (s *UpsertCollectionDataShrinkRequest) SetNamespacePassword(v string) *UpsertCollectionDataShrinkRequest {
	s.NamespacePassword = &v
	return s
}

func (s *UpsertCollectionDataShrinkRequest) SetOwnerId(v int64) *UpsertCollectionDataShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *UpsertCollectionDataShrinkRequest) SetRegionId(v string) *UpsertCollectionDataShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *UpsertCollectionDataShrinkRequest) SetRowsShrink(v string) *UpsertCollectionDataShrinkRequest {
	s.RowsShrink = &v
	return s
}

type UpsertCollectionDataResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpsertCollectionDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpsertCollectionDataResponseBody) GoString() string {
	return s.String()
}

func (s *UpsertCollectionDataResponseBody) SetMessage(v string) *UpsertCollectionDataResponseBody {
	s.Message = &v
	return s
}

func (s *UpsertCollectionDataResponseBody) SetRequestId(v string) *UpsertCollectionDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpsertCollectionDataResponseBody) SetStatus(v string) *UpsertCollectionDataResponseBody {
	s.Status = &v
	return s
}

type UpsertCollectionDataResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpsertCollectionDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpsertCollectionDataResponse) String() string {
	return tea.Prettify(s)
}

func (s UpsertCollectionDataResponse) GoString() string {
	return s.String()
}

func (s *UpsertCollectionDataResponse) SetHeaders(v map[string]*string) *UpsertCollectionDataResponse {
	s.Headers = v
	return s
}

func (s *UpsertCollectionDataResponse) SetStatusCode(v int32) *UpsertCollectionDataResponse {
	s.StatusCode = &v
	return s
}

func (s *UpsertCollectionDataResponse) SetBody(v *UpsertCollectionDataResponseBody) *UpsertCollectionDataResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	client.EndpointMap = map[string]*string{
		"cn-beijing":            tea.String("gpdb.aliyuncs.com"),
		"cn-hangzhou":           tea.String("gpdb.aliyuncs.com"),
		"cn-shanghai":           tea.String("gpdb.aliyuncs.com"),
		"cn-shenzhen":           tea.String("gpdb.aliyuncs.com"),
		"cn-hongkong":           tea.String("gpdb.aliyuncs.com"),
		"ap-southeast-1":        tea.String("gpdb.aliyuncs.com"),
		"us-west-1":             tea.String("gpdb.aliyuncs.com"),
		"us-east-1":             tea.String("gpdb.aliyuncs.com"),
		"cn-hangzhou-finance":   tea.String("gpdb.aliyuncs.com"),
		"cn-shanghai-finance-1": tea.String("gpdb.aliyuncs.com"),
		"cn-shenzhen-finance-1": tea.String("gpdb.aliyuncs.com"),
		"cn-qingdao":            tea.String("gpdb.aliyuncs.com"),
		"cn-north-2-gov-1":      tea.String("gpdb.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("gpdb"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this operation to apply for a public endpoint for an AnalyticDB for PostgreSQL instance. Both the primary and instance endpoints of an AnalyticDB for PostgreSQL instance can be public endpoints. For more information, see [Endpoints of an instance and its primary coordinator node](~~204879~~).
 * ## Limits
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request AllocateInstancePublicConnectionRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return AllocateInstancePublicConnectionResponse
 */
func (client *Client) AllocateInstancePublicConnectionWithOptions(request *AllocateInstancePublicConnectionRequest, runtime *util.RuntimeOptions) (_result *AllocateInstancePublicConnectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AddressType)) {
		query["AddressType"] = request.AddressType
	}

	if !tea.BoolValue(util.IsUnset(request.ConnectionStringPrefix)) {
		query["ConnectionStringPrefix"] = request.ConnectionStringPrefix
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Port)) {
		query["Port"] = request.Port
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AllocateInstancePublicConnection"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AllocateInstancePublicConnectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this operation to apply for a public endpoint for an AnalyticDB for PostgreSQL instance. Both the primary and instance endpoints of an AnalyticDB for PostgreSQL instance can be public endpoints. For more information, see [Endpoints of an instance and its primary coordinator node](~~204879~~).
 * ## Limits
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request AllocateInstancePublicConnectionRequest
 * @return AllocateInstancePublicConnectionResponse
 */
func (client *Client) AllocateInstancePublicConnection(request *AllocateInstancePublicConnectionRequest) (_result *AllocateInstancePublicConnectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AllocateInstancePublicConnectionResponse{}
	_body, _err := client.AllocateInstancePublicConnectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckServiceLinkedRoleWithOptions(request *CheckServiceLinkedRoleRequest, runtime *util.RuntimeOptions) (_result *CheckServiceLinkedRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckServiceLinkedRole"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckServiceLinkedRoleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckServiceLinkedRole(request *CheckServiceLinkedRoleRequest) (_result *CheckServiceLinkedRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckServiceLinkedRoleResponse{}
	_body, _err := client.CheckServiceLinkedRoleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * *   Before you can use an AnalyticDB for PostgreSQL instance, you must create a privileged account for the instance.
 * *   You can call this operation to create only privileged accounts. For information about how to create other types of accounts, see [Create a database account](~~50206~~).
 * ## Limits
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request CreateAccountRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateAccountResponse
 */
func (client *Client) CreateAccountWithOptions(request *CreateAccountRequest, runtime *util.RuntimeOptions) (_result *CreateAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountDescription)) {
		query["AccountDescription"] = request.AccountDescription
	}

	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		query["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.AccountPassword)) {
		query["AccountPassword"] = request.AccountPassword
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.DatabaseName)) {
		query["DatabaseName"] = request.DatabaseName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAccount"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * *   Before you can use an AnalyticDB for PostgreSQL instance, you must create a privileged account for the instance.
 * *   You can call this operation to create only privileged accounts. For information about how to create other types of accounts, see [Create a database account](~~50206~~).
 * ## Limits
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request CreateAccountRequest
 * @return CreateAccountResponse
 */
func (client *Client) CreateAccount(request *CreateAccountRequest) (_result *CreateAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAccountResponse{}
	_body, _err := client.CreateAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateCollectionWithOptions(request *CreateCollectionRequest, runtime *util.RuntimeOptions) (_result *CreateCollectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Collection)) {
		query["Collection"] = request.Collection
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Dimension)) {
		query["Dimension"] = request.Dimension
	}

	if !tea.BoolValue(util.IsUnset(request.FullTextRetrievalFields)) {
		query["FullTextRetrievalFields"] = request.FullTextRetrievalFields
	}

	if !tea.BoolValue(util.IsUnset(request.ManagerAccount)) {
		query["ManagerAccount"] = request.ManagerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ManagerAccountPassword)) {
		query["ManagerAccountPassword"] = request.ManagerAccountPassword
	}

	if !tea.BoolValue(util.IsUnset(request.Metadata)) {
		query["Metadata"] = request.Metadata
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["Namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Parser)) {
		query["Parser"] = request.Parser
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateCollection"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateCollectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCollection(request *CreateCollectionRequest) (_result *CreateCollectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateCollectionResponse{}
	_body, _err := client.CreateCollectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this operation when you need to create AnalyticDB for PostgreSQL instances to meet the requirements of new applications or services.
 * Before you call this operation, make sure that you are familiar with the billing of AnalyticDB for PostgreSQL instances. For more information, see [Billing methods](~~35406~~) and [AnalyticDB for PostgreSQL pricing](https://www.alibabacloud.com/zh/product/hybriddb-postgresql/pricing).
 * ## Limits
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds a limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limits when you call this operation.
 *
 * @param request CreateDBInstanceRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateDBInstanceResponse
 */
func (client *Client) CreateDBInstanceWithOptions(request *CreateDBInstanceRequest, runtime *util.RuntimeOptions) (_result *CreateDBInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.CreateSampleData)) {
		query["CreateSampleData"] = request.CreateSampleData
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceCategory)) {
		query["DBInstanceCategory"] = request.DBInstanceCategory
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceClass)) {
		query["DBInstanceClass"] = request.DBInstanceClass
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceDescription)) {
		query["DBInstanceDescription"] = request.DBInstanceDescription
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceGroupCount)) {
		query["DBInstanceGroupCount"] = request.DBInstanceGroupCount
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceMode)) {
		query["DBInstanceMode"] = request.DBInstanceMode
	}

	if !tea.BoolValue(util.IsUnset(request.EncryptionKey)) {
		query["EncryptionKey"] = request.EncryptionKey
	}

	if !tea.BoolValue(util.IsUnset(request.EncryptionType)) {
		query["EncryptionType"] = request.EncryptionType
	}

	if !tea.BoolValue(util.IsUnset(request.Engine)) {
		query["Engine"] = request.Engine
	}

	if !tea.BoolValue(util.IsUnset(request.EngineVersion)) {
		query["EngineVersion"] = request.EngineVersion
	}

	if !tea.BoolValue(util.IsUnset(request.IdleTime)) {
		query["IdleTime"] = request.IdleTime
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceNetworkType)) {
		query["InstanceNetworkType"] = request.InstanceNetworkType
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceSpec)) {
		query["InstanceSpec"] = request.InstanceSpec
	}

	if !tea.BoolValue(util.IsUnset(request.MasterNodeNum)) {
		query["MasterNodeNum"] = request.MasterNodeNum
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PayType)) {
		query["PayType"] = request.PayType
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		query["Period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.PrivateIpAddress)) {
		query["PrivateIpAddress"] = request.PrivateIpAddress
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityIPList)) {
		query["SecurityIPList"] = request.SecurityIPList
	}

	if !tea.BoolValue(util.IsUnset(request.SegDiskPerformanceLevel)) {
		query["SegDiskPerformanceLevel"] = request.SegDiskPerformanceLevel
	}

	if !tea.BoolValue(util.IsUnset(request.SegNodeNum)) {
		query["SegNodeNum"] = request.SegNodeNum
	}

	if !tea.BoolValue(util.IsUnset(request.SegStorageType)) {
		query["SegStorageType"] = request.SegStorageType
	}

	if !tea.BoolValue(util.IsUnset(request.ServerlessMode)) {
		query["ServerlessMode"] = request.ServerlessMode
	}

	if !tea.BoolValue(util.IsUnset(request.ServerlessResource)) {
		query["ServerlessResource"] = request.ServerlessResource
	}

	if !tea.BoolValue(util.IsUnset(request.StorageSize)) {
		query["StorageSize"] = request.StorageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StorageType)) {
		query["StorageType"] = request.StorageType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.UsedTime)) {
		query["UsedTime"] = request.UsedTime
	}

	if !tea.BoolValue(util.IsUnset(request.VPCId)) {
		query["VPCId"] = request.VPCId
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchId)) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.VectorConfigurationStatus)) {
		query["VectorConfigurationStatus"] = request.VectorConfigurationStatus
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDBInstance"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDBInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this operation when you need to create AnalyticDB for PostgreSQL instances to meet the requirements of new applications or services.
 * Before you call this operation, make sure that you are familiar with the billing of AnalyticDB for PostgreSQL instances. For more information, see [Billing methods](~~35406~~) and [AnalyticDB for PostgreSQL pricing](https://www.alibabacloud.com/zh/product/hybriddb-postgresql/pricing).
 * ## Limits
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds a limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limits when you call this operation.
 *
 * @param request CreateDBInstanceRequest
 * @return CreateDBInstanceResponse
 */
func (client *Client) CreateDBInstance(request *CreateDBInstanceRequest) (_result *CreateDBInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDBInstanceResponse{}
	_body, _err := client.CreateDBInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this operation to create a plan for an AnalyticDB for PostgreSQL instance. For example, you can create a plan to pause and resume an instance, change the number of compute nodes, or change compute node specifications.
 * >  This operation is applicable only to pay-as-you-go instances.
 * ## Limits
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request CreateDBInstancePlanRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateDBInstancePlanResponse
 */
func (client *Client) CreateDBInstancePlanWithOptions(request *CreateDBInstancePlanRequest, runtime *util.RuntimeOptions) (_result *CreateDBInstancePlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PlanConfig)) {
		query["PlanConfig"] = request.PlanConfig
	}

	if !tea.BoolValue(util.IsUnset(request.PlanDesc)) {
		query["PlanDesc"] = request.PlanDesc
	}

	if !tea.BoolValue(util.IsUnset(request.PlanEndDate)) {
		query["PlanEndDate"] = request.PlanEndDate
	}

	if !tea.BoolValue(util.IsUnset(request.PlanName)) {
		query["PlanName"] = request.PlanName
	}

	if !tea.BoolValue(util.IsUnset(request.PlanScheduleType)) {
		query["PlanScheduleType"] = request.PlanScheduleType
	}

	if !tea.BoolValue(util.IsUnset(request.PlanStartDate)) {
		query["PlanStartDate"] = request.PlanStartDate
	}

	if !tea.BoolValue(util.IsUnset(request.PlanType)) {
		query["PlanType"] = request.PlanType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDBInstancePlan"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDBInstancePlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this operation to create a plan for an AnalyticDB for PostgreSQL instance. For example, you can create a plan to pause and resume an instance, change the number of compute nodes, or change compute node specifications.
 * >  This operation is applicable only to pay-as-you-go instances.
 * ## Limits
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request CreateDBInstancePlanRequest
 * @return CreateDBInstancePlanResponse
 */
func (client *Client) CreateDBInstancePlan(request *CreateDBInstancePlanRequest) (_result *CreateDBInstancePlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDBInstancePlanResponse{}
	_body, _err := client.CreateDBInstancePlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateNamespaceWithOptions(request *CreateNamespaceRequest, runtime *util.RuntimeOptions) (_result *CreateNamespaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ManagerAccount)) {
		query["ManagerAccount"] = request.ManagerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ManagerAccountPassword)) {
		query["ManagerAccountPassword"] = request.ManagerAccountPassword
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["Namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.NamespacePassword)) {
		query["NamespacePassword"] = request.NamespacePassword
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateNamespace"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateNamespaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateNamespace(request *CreateNamespaceRequest) (_result *CreateNamespaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateNamespaceResponse{}
	_body, _err := client.CreateNamespaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this operation to load a sample dataset to an AnalyticDB for PostgreSQL instance. Then, you can execute query statements on the sample dataset to experience or test your instance. For more information about query statements, see [Dataset information and query examples](~~452277~~).
 * ## Precautions
 * - If your instance is in elastic storage mode, the sample dataset is supported only for V6.3.10.3 or later. If your instance is in Serverless mode, the sample dataset is supported only for V1.0.4.0 or later. For more information about how to update the minor engine version of an instance, see [Update the minor engine version](/help/en/analyticdb-for-postgresql/latest/upgrade-the-engine-version).
 * - The sample dataset is about 10 GB in size. Make sure that your instance has sufficient storage space.
 * - The sample dataset contains a database named `ADB_SampleData_TPCH`. Make sure that your instance does not have a database with the same name. Otherwise, the dataset may fail to be loaded.
 * - It may take 6 to 8 minutes to load the sample dataset. During this period, operations on your instance such as adding nodes or changing node specifications may be affected.
 * ## Limits
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request CreateSampleDataRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateSampleDataResponse
 */
func (client *Client) CreateSampleDataWithOptions(request *CreateSampleDataRequest, runtime *util.RuntimeOptions) (_result *CreateSampleDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSampleData"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSampleDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this operation to load a sample dataset to an AnalyticDB for PostgreSQL instance. Then, you can execute query statements on the sample dataset to experience or test your instance. For more information about query statements, see [Dataset information and query examples](~~452277~~).
 * ## Precautions
 * - If your instance is in elastic storage mode, the sample dataset is supported only for V6.3.10.3 or later. If your instance is in Serverless mode, the sample dataset is supported only for V1.0.4.0 or later. For more information about how to update the minor engine version of an instance, see [Update the minor engine version](/help/en/analyticdb-for-postgresql/latest/upgrade-the-engine-version).
 * - The sample dataset is about 10 GB in size. Make sure that your instance has sufficient storage space.
 * - The sample dataset contains a database named `ADB_SampleData_TPCH`. Make sure that your instance does not have a database with the same name. Otherwise, the dataset may fail to be loaded.
 * - It may take 6 to 8 minutes to load the sample dataset. During this period, operations on your instance such as adding nodes or changing node specifications may be affected.
 * ## Limits
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request CreateSampleDataRequest
 * @return CreateSampleDataResponse
 */
func (client *Client) CreateSampleData(request *CreateSampleDataRequest) (_result *CreateSampleDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSampleDataResponse{}
	_body, _err := client.CreateSampleDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateServiceLinkedRoleWithOptions(request *CreateServiceLinkedRoleRequest, runtime *util.RuntimeOptions) (_result *CreateServiceLinkedRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateServiceLinkedRole"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateServiceLinkedRoleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateServiceLinkedRole(request *CreateServiceLinkedRoleRequest) (_result *CreateServiceLinkedRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateServiceLinkedRoleResponse{}
	_body, _err := client.CreateServiceLinkedRoleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateVectorIndexWithOptions(request *CreateVectorIndexRequest, runtime *util.RuntimeOptions) (_result *CreateVectorIndexResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Collection)) {
		query["Collection"] = request.Collection
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Dimension)) {
		query["Dimension"] = request.Dimension
	}

	if !tea.BoolValue(util.IsUnset(request.ManagerAccount)) {
		query["ManagerAccount"] = request.ManagerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ManagerAccountPassword)) {
		query["ManagerAccountPassword"] = request.ManagerAccountPassword
	}

	if !tea.BoolValue(util.IsUnset(request.Metrics)) {
		query["Metrics"] = request.Metrics
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["Namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateVectorIndex"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateVectorIndexResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateVectorIndex(request *CreateVectorIndexRequest) (_result *CreateVectorIndexResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateVectorIndexResponse{}
	_body, _err := client.CreateVectorIndexWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteCollectionWithOptions(request *DeleteCollectionRequest, runtime *util.RuntimeOptions) (_result *DeleteCollectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Collection)) {
		query["Collection"] = request.Collection
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["Namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.NamespacePassword)) {
		query["NamespacePassword"] = request.NamespacePassword
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteCollection"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteCollectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteCollection(request *DeleteCollectionRequest) (_result *DeleteCollectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteCollectionResponse{}
	_body, _err := client.DeleteCollectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteCollectionDataWithOptions(request *DeleteCollectionDataRequest, runtime *util.RuntimeOptions) (_result *DeleteCollectionDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Collection)) {
		query["Collection"] = request.Collection
	}

	if !tea.BoolValue(util.IsUnset(request.CollectionData)) {
		query["CollectionData"] = request.CollectionData
	}

	if !tea.BoolValue(util.IsUnset(request.CollectionDataFilter)) {
		query["CollectionDataFilter"] = request.CollectionDataFilter
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["Namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.NamespacePassword)) {
		query["NamespacePassword"] = request.NamespacePassword
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteCollectionData"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteCollectionDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteCollectionData(request *DeleteCollectionDataRequest) (_result *DeleteCollectionDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteCollectionDataResponse{}
	_body, _err := client.DeleteCollectionDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * *   Subscription instances cannot be manually released. They are automatically released when they expire.
 * *   You can call this operation to release pay-as-you-go instances only when they are in the **Running** state.
 * ## Limits
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DeleteDBInstanceRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteDBInstanceResponse
 */
func (client *Client) DeleteDBInstanceWithOptions(request *DeleteDBInstanceRequest, runtime *util.RuntimeOptions) (_result *DeleteDBInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDBInstance"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDBInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * *   Subscription instances cannot be manually released. They are automatically released when they expire.
 * *   You can call this operation to release pay-as-you-go instances only when they are in the **Running** state.
 * ## Limits
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DeleteDBInstanceRequest
 * @return DeleteDBInstanceResponse
 */
func (client *Client) DeleteDBInstance(request *DeleteDBInstanceRequest) (_result *DeleteDBInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDBInstanceResponse{}
	_body, _err := client.DeleteDBInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * If you no longer need a plan, you can call this operation to delete the plan. The plan management feature is supported only for AnalyticDB for PostgreSQL instances in Serverless mode.
 * ## Limits
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DeleteDBInstancePlanRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteDBInstancePlanResponse
 */
func (client *Client) DeleteDBInstancePlanWithOptions(request *DeleteDBInstancePlanRequest, runtime *util.RuntimeOptions) (_result *DeleteDBInstancePlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PlanId)) {
		query["PlanId"] = request.PlanId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDBInstancePlan"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDBInstancePlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * If you no longer need a plan, you can call this operation to delete the plan. The plan management feature is supported only for AnalyticDB for PostgreSQL instances in Serverless mode.
 * ## Limits
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DeleteDBInstancePlanRequest
 * @return DeleteDBInstancePlanResponse
 */
func (client *Client) DeleteDBInstancePlan(request *DeleteDBInstancePlanRequest) (_result *DeleteDBInstancePlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDBInstancePlanResponse{}
	_body, _err := client.DeleteDBInstancePlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteNamespaceWithOptions(request *DeleteNamespaceRequest, runtime *util.RuntimeOptions) (_result *DeleteNamespaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ManagerAccount)) {
		query["ManagerAccount"] = request.ManagerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ManagerAccountPassword)) {
		query["ManagerAccountPassword"] = request.ManagerAccountPassword
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["Namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteNamespace"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteNamespaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteNamespace(request *DeleteNamespaceRequest) (_result *DeleteNamespaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteNamespaceResponse{}
	_body, _err := client.DeleteNamespaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteVectorIndexWithOptions(request *DeleteVectorIndexRequest, runtime *util.RuntimeOptions) (_result *DeleteVectorIndexResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Collection)) {
		query["Collection"] = request.Collection
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ManagerAccount)) {
		query["ManagerAccount"] = request.ManagerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ManagerAccountPassword)) {
		query["ManagerAccountPassword"] = request.ManagerAccountPassword
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["Namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteVectorIndex"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteVectorIndexResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteVectorIndex(request *DeleteVectorIndexRequest) (_result *DeleteVectorIndexResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteVectorIndexResponse{}
	_body, _err := client.DeleteVectorIndexWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * This operation is called to query the information of the privileged account in an AnalyticDB for PostgreSQL instance, such as its state, description, and the instance.
 * ## Limit
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered and may affect your business. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeAccountsRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeAccountsResponse
 */
func (client *Client) DescribeAccountsWithOptions(request *DescribeAccountsRequest, runtime *util.RuntimeOptions) (_result *DescribeAccountsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		query["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAccounts"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAccountsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * This operation is called to query the information of the privileged account in an AnalyticDB for PostgreSQL instance, such as its state, description, and the instance.
 * ## Limit
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered and may affect your business. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeAccountsRequest
 * @return DescribeAccountsResponse
 */
func (client *Client) DescribeAccounts(request *DescribeAccountsRequest) (_result *DescribeAccountsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAccountsResponse{}
	_body, _err := client.DescribeAccountsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * When you create an AnalyticDB for PostgreSQL instance, you can call this operation to query the available resources within a specific zone.
 * ## Limits
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeAvailableResourcesRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeAvailableResourcesResponse
 */
func (client *Client) DescribeAvailableResourcesWithOptions(request *DescribeAvailableResourcesRequest, runtime *util.RuntimeOptions) (_result *DescribeAvailableResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChargeType)) {
		query["ChargeType"] = request.ChargeType
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["Region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAvailableResources"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAvailableResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * When you create an AnalyticDB for PostgreSQL instance, you can call this operation to query the available resources within a specific zone.
 * ## Limits
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeAvailableResourcesRequest
 * @return DescribeAvailableResourcesResponse
 */
func (client *Client) DescribeAvailableResources(request *DescribeAvailableResourcesRequest) (_result *DescribeAvailableResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAvailableResourcesResponse{}
	_body, _err := client.DescribeAvailableResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this operation to query the backup settings of an AnalyticDB for PostgreSQL instance in elastic storage mode. Periodically backing data can prevent data loss. For more information about how to modify backup policies, see [ModifyBackupPolicy](~~210095~~).
 * ## Limits
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeBackupPolicyRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeBackupPolicyResponse
 */
func (client *Client) DescribeBackupPolicyWithOptions(request *DescribeBackupPolicyRequest, runtime *util.RuntimeOptions) (_result *DescribeBackupPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeBackupPolicy"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeBackupPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this operation to query the backup settings of an AnalyticDB for PostgreSQL instance in elastic storage mode. Periodically backing data can prevent data loss. For more information about how to modify backup policies, see [ModifyBackupPolicy](~~210095~~).
 * ## Limits
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeBackupPolicyRequest
 * @return DescribeBackupPolicyResponse
 */
func (client *Client) DescribeBackupPolicy(request *DescribeBackupPolicyRequest) (_result *DescribeBackupPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBackupPolicyResponse{}
	_body, _err := client.DescribeBackupPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCollectionWithOptions(request *DescribeCollectionRequest, runtime *util.RuntimeOptions) (_result *DescribeCollectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Collection)) {
		query["Collection"] = request.Collection
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["Namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.NamespacePassword)) {
		query["NamespacePassword"] = request.NamespacePassword
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCollection"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCollectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCollection(request *DescribeCollectionRequest) (_result *DescribeCollectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCollectionResponse{}
	_body, _err := client.DescribeCollectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * This operation is called to query the information of coordinator and compute nodes in an AnalyticDB for PostgreSQL instance in elastic storage mode or Serverless mode.
 * ## Limits
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered and may affect your business. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeDBClusterNodeRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeDBClusterNodeResponse
 */
func (client *Client) DescribeDBClusterNodeWithOptions(request *DescribeDBClusterNodeRequest, runtime *util.RuntimeOptions) (_result *DescribeDBClusterNodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.NodeType)) {
		query["NodeType"] = request.NodeType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDBClusterNode"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDBClusterNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * This operation is called to query the information of coordinator and compute nodes in an AnalyticDB for PostgreSQL instance in elastic storage mode or Serverless mode.
 * ## Limits
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered and may affect your business. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeDBClusterNodeRequest
 * @return DescribeDBClusterNodeResponse
 */
func (client *Client) DescribeDBClusterNode(request *DescribeDBClusterNodeRequest) (_result *DescribeDBClusterNodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBClusterNodeResponse{}
	_body, _err := client.DescribeDBClusterNodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * This operation is called to query the performance metrics of an AnalyticDB for PostgreSQL instance, such as the number of connections, memory usage, CPU utilization, I/O throughput, read IOPS, write IOPS, and disk space usage.
 * You can query monitoring information only within the last 30 days.
 * ## Limits
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered and may affect your business. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeDBClusterPerformanceRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeDBClusterPerformanceResponse
 */
func (client *Client) DescribeDBClusterPerformanceWithOptions(request *DescribeDBClusterPerformanceRequest, runtime *util.RuntimeOptions) (_result *DescribeDBClusterPerformanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Key)) {
		query["Key"] = request.Key
	}

	if !tea.BoolValue(util.IsUnset(request.NodeType)) {
		query["NodeType"] = request.NodeType
	}

	if !tea.BoolValue(util.IsUnset(request.Nodes)) {
		query["Nodes"] = request.Nodes
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDBClusterPerformance"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDBClusterPerformanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * This operation is called to query the performance metrics of an AnalyticDB for PostgreSQL instance, such as the number of connections, memory usage, CPU utilization, I/O throughput, read IOPS, write IOPS, and disk space usage.
 * You can query monitoring information only within the last 30 days.
 * ## Limits
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered and may affect your business. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeDBClusterPerformanceRequest
 * @return DescribeDBClusterPerformanceResponse
 */
func (client *Client) DescribeDBClusterPerformance(request *DescribeDBClusterPerformanceRequest) (_result *DescribeDBClusterPerformanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBClusterPerformanceResponse{}
	_body, _err := client.DescribeDBClusterPerformanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this operation to query the information about an AnalyticDB for PostgreSQL instance, such as the instance type, network type, and instance state.
 * ## Limits
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds a limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limits when you call this operation.
 *
 * @param request DescribeDBInstanceAttributeRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeDBInstanceAttributeResponse
 */
func (client *Client) DescribeDBInstanceAttributeWithOptions(request *DescribeDBInstanceAttributeRequest, runtime *util.RuntimeOptions) (_result *DescribeDBInstanceAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDBInstanceAttribute"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDBInstanceAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this operation to query the information about an AnalyticDB for PostgreSQL instance, such as the instance type, network type, and instance state.
 * ## Limits
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds a limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limits when you call this operation.
 *
 * @param request DescribeDBInstanceAttributeRequest
 * @return DescribeDBInstanceAttributeResponse
 */
func (client *Client) DescribeDBInstanceAttribute(request *DescribeDBInstanceAttributeRequest) (_result *DescribeDBInstanceAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBInstanceAttributeResponse{}
	_body, _err := client.DescribeDBInstanceAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this operation to query the details of data bloat on an AnalyticDB for PostgreSQL instance in elastic storage mode. The minor version of the instance must be V6.3.10.1 or later. For more information about how to view and update the minor version of an instance, see [View the minor engine version](~~277424~~) and [Update the minor engine version](~~139271~~).
 * ## Limits
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeDBInstanceDataBloatRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeDBInstanceDataBloatResponse
 */
func (client *Client) DescribeDBInstanceDataBloatWithOptions(request *DescribeDBInstanceDataBloatRequest, runtime *util.RuntimeOptions) (_result *DescribeDBInstanceDataBloatResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDBInstanceDataBloat"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDBInstanceDataBloatResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this operation to query the details of data bloat on an AnalyticDB for PostgreSQL instance in elastic storage mode. The minor version of the instance must be V6.3.10.1 or later. For more information about how to view and update the minor version of an instance, see [View the minor engine version](~~277424~~) and [Update the minor engine version](~~139271~~).
 * ## Limits
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeDBInstanceDataBloatRequest
 * @return DescribeDBInstanceDataBloatResponse
 */
func (client *Client) DescribeDBInstanceDataBloat(request *DescribeDBInstanceDataBloatRequest) (_result *DescribeDBInstanceDataBloatResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBInstanceDataBloatResponse{}
	_body, _err := client.DescribeDBInstanceDataBloatWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * To prevent data skew from affecting your database service, you can call this operation to query the details about data skew on an AnalyticDB for PostgreSQL instance.
 * ## Limits
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeDBInstanceDataSkewRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeDBInstanceDataSkewResponse
 */
func (client *Client) DescribeDBInstanceDataSkewWithOptions(request *DescribeDBInstanceDataSkewRequest, runtime *util.RuntimeOptions) (_result *DescribeDBInstanceDataSkewResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDBInstanceDataSkew"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDBInstanceDataSkewResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * To prevent data skew from affecting your database service, you can call this operation to query the details about data skew on an AnalyticDB for PostgreSQL instance.
 * ## Limits
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeDBInstanceDataSkewRequest
 * @return DescribeDBInstanceDataSkewResponse
 */
func (client *Client) DescribeDBInstanceDataSkew(request *DescribeDBInstanceDataSkewRequest) (_result *DescribeDBInstanceDataSkewResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBInstanceDataSkewResponse{}
	_body, _err := client.DescribeDBInstanceDataSkewWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this operation to query the distribution and states of coordinator and compute nodes on an AnalyticDB for PostgreSQL instance.
 * ## Limits
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeDBInstanceDiagnosisSummaryRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeDBInstanceDiagnosisSummaryResponse
 */
func (client *Client) DescribeDBInstanceDiagnosisSummaryWithOptions(request *DescribeDBInstanceDiagnosisSummaryRequest, runtime *util.RuntimeOptions) (_result *DescribeDBInstanceDiagnosisSummaryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RolePreferd)) {
		query["RolePreferd"] = request.RolePreferd
	}

	if !tea.BoolValue(util.IsUnset(request.StartStatus)) {
		query["StartStatus"] = request.StartStatus
	}

	if !tea.BoolValue(util.IsUnset(request.SyncMode)) {
		query["SyncMode"] = request.SyncMode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDBInstanceDiagnosisSummary"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDBInstanceDiagnosisSummaryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this operation to query the distribution and states of coordinator and compute nodes on an AnalyticDB for PostgreSQL instance.
 * ## Limits
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeDBInstanceDiagnosisSummaryRequest
 * @return DescribeDBInstanceDiagnosisSummaryResponse
 */
func (client *Client) DescribeDBInstanceDiagnosisSummary(request *DescribeDBInstanceDiagnosisSummaryRequest) (_result *DescribeDBInstanceDiagnosisSummaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBInstanceDiagnosisSummaryResponse{}
	_body, _err := client.DescribeDBInstanceDiagnosisSummaryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this operation to query the error logs of an AnalyticDB for PostgreSQL instance.
 * ## Limits
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeDBInstanceErrorLogRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeDBInstanceErrorLogResponse
 */
func (client *Client) DescribeDBInstanceErrorLogWithOptions(request *DescribeDBInstanceErrorLogRequest, runtime *util.RuntimeOptions) (_result *DescribeDBInstanceErrorLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Database)) {
		query["Database"] = request.Database
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Host)) {
		query["Host"] = request.Host
	}

	if !tea.BoolValue(util.IsUnset(request.Keywords)) {
		query["Keywords"] = request.Keywords
	}

	if !tea.BoolValue(util.IsUnset(request.LogLevel)) {
		query["LogLevel"] = request.LogLevel
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.User)) {
		query["User"] = request.User
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDBInstanceErrorLog"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDBInstanceErrorLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this operation to query the error logs of an AnalyticDB for PostgreSQL instance.
 * ## Limits
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeDBInstanceErrorLogRequest
 * @return DescribeDBInstanceErrorLogResponse
 */
func (client *Client) DescribeDBInstanceErrorLog(request *DescribeDBInstanceErrorLogRequest) (_result *DescribeDBInstanceErrorLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBInstanceErrorLogResponse{}
	_body, _err := client.DescribeDBInstanceErrorLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this operation to query the whitelists of IP addresses that are allowed to access an AnalyticDB for PostgreSQL instance.
 * ## Limits
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeDBInstanceIPArrayListRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeDBInstanceIPArrayListResponse
 */
func (client *Client) DescribeDBInstanceIPArrayListWithOptions(request *DescribeDBInstanceIPArrayListRequest, runtime *util.RuntimeOptions) (_result *DescribeDBInstanceIPArrayListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDBInstanceIPArrayList"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDBInstanceIPArrayListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this operation to query the whitelists of IP addresses that are allowed to access an AnalyticDB for PostgreSQL instance.
 * ## Limits
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeDBInstanceIPArrayListRequest
 * @return DescribeDBInstanceIPArrayListResponse
 */
func (client *Client) DescribeDBInstanceIPArrayList(request *DescribeDBInstanceIPArrayListRequest) (_result *DescribeDBInstanceIPArrayListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBInstanceIPArrayListResponse{}
	_body, _err := client.DescribeDBInstanceIPArrayListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Appropriate indexes can improve the database query speed. You can call this operation to query the details of index usage on an AnalyticDB for PostgreSQL instance.
 * This operation is available only for instances of V6.3.10.1 or later in elastic storage mode. For more information about how to view and update the minor engine version of an instance, see [View the minor engine version](~~277424~~) and [Update the minor engine version](~~139271~~).
 * ## Limits
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeDBInstanceIndexUsageRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeDBInstanceIndexUsageResponse
 */
func (client *Client) DescribeDBInstanceIndexUsageWithOptions(request *DescribeDBInstanceIndexUsageRequest, runtime *util.RuntimeOptions) (_result *DescribeDBInstanceIndexUsageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDBInstanceIndexUsage"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDBInstanceIndexUsageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Appropriate indexes can improve the database query speed. You can call this operation to query the details of index usage on an AnalyticDB for PostgreSQL instance.
 * This operation is available only for instances of V6.3.10.1 or later in elastic storage mode. For more information about how to view and update the minor engine version of an instance, see [View the minor engine version](~~277424~~) and [Update the minor engine version](~~139271~~).
 * ## Limits
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeDBInstanceIndexUsageRequest
 * @return DescribeDBInstanceIndexUsageResponse
 */
func (client *Client) DescribeDBInstanceIndexUsage(request *DescribeDBInstanceIndexUsageRequest) (_result *DescribeDBInstanceIndexUsageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBInstanceIndexUsageResponse{}
	_body, _err := client.DescribeDBInstanceIndexUsageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDBInstanceNetInfoWithOptions(request *DescribeDBInstanceNetInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeDBInstanceNetInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConnectionString)) {
		query["ConnectionString"] = request.ConnectionString
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDBInstanceNetInfo"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDBInstanceNetInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBInstanceNetInfo(request *DescribeDBInstanceNetInfoRequest) (_result *DescribeDBInstanceNetInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBInstanceNetInfoResponse{}
	_body, _err := client.DescribeDBInstanceNetInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDBInstancePerformanceWithOptions(request *DescribeDBInstancePerformanceRequest, runtime *util.RuntimeOptions) (_result *DescribeDBInstancePerformanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Key)) {
		query["Key"] = request.Key
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDBInstancePerformance"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDBInstancePerformanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBInstancePerformance(request *DescribeDBInstancePerformanceRequest) (_result *DescribeDBInstancePerformanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBInstancePerformanceResponse{}
	_body, _err := client.DescribeDBInstancePerformanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this operation to query the details of plans for an AnalyticDB for PostgreSQL instance in Serverless mode.
 * ## Limits
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeDBInstancePlansRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeDBInstancePlansResponse
 */
func (client *Client) DescribeDBInstancePlansWithOptions(request *DescribeDBInstancePlansRequest, runtime *util.RuntimeOptions) (_result *DescribeDBInstancePlansResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PlanCreateDate)) {
		query["PlanCreateDate"] = request.PlanCreateDate
	}

	if !tea.BoolValue(util.IsUnset(request.PlanDesc)) {
		query["PlanDesc"] = request.PlanDesc
	}

	if !tea.BoolValue(util.IsUnset(request.PlanId)) {
		query["PlanId"] = request.PlanId
	}

	if !tea.BoolValue(util.IsUnset(request.PlanScheduleType)) {
		query["PlanScheduleType"] = request.PlanScheduleType
	}

	if !tea.BoolValue(util.IsUnset(request.PlanType)) {
		query["PlanType"] = request.PlanType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDBInstancePlans"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDBInstancePlansResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this operation to query the details of plans for an AnalyticDB for PostgreSQL instance in Serverless mode.
 * ## Limits
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeDBInstancePlansRequest
 * @return DescribeDBInstancePlansResponse
 */
func (client *Client) DescribeDBInstancePlans(request *DescribeDBInstancePlansRequest) (_result *DescribeDBInstancePlansResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBInstancePlansResponse{}
	_body, _err := client.DescribeDBInstancePlansWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDBInstanceSSLWithOptions(request *DescribeDBInstanceSSLRequest, runtime *util.RuntimeOptions) (_result *DescribeDBInstanceSSLResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDBInstanceSSL"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDBInstanceSSLResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBInstanceSSL(request *DescribeDBInstanceSSLRequest) (_result *DescribeDBInstanceSSLResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBInstanceSSLResponse{}
	_body, _err := client.DescribeDBInstanceSSLWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this operation to query the instance types, network types, and states of AnalyticDB for PostgreSQL instances within a specific region.
 * ## Limits
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param tmpReq DescribeDBInstancesRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeDBInstancesResponse
 */
func (client *Client) DescribeDBInstancesWithOptions(tmpReq *DescribeDBInstancesRequest, runtime *util.RuntimeOptions) (_result *DescribeDBInstancesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DescribeDBInstancesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.DBInstanceCategories)) {
		request.DBInstanceCategoriesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DBInstanceCategories, tea.String("DBInstanceCategories"), tea.String("simple"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.DBInstanceModes)) {
		request.DBInstanceModesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DBInstanceModes, tea.String("DBInstanceModes"), tea.String("simple"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.DBInstanceStatuses)) {
		request.DBInstanceStatusesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DBInstanceStatuses, tea.String("DBInstanceStatuses"), tea.String("simple"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.InstanceDeployTypes)) {
		request.InstanceDeployTypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InstanceDeployTypes, tea.String("InstanceDeployTypes"), tea.String("simple"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceCategoriesShrink)) {
		query["DBInstanceCategories"] = request.DBInstanceCategoriesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceDescription)) {
		query["DBInstanceDescription"] = request.DBInstanceDescription
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceIds)) {
		query["DBInstanceIds"] = request.DBInstanceIds
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceModesShrink)) {
		query["DBInstanceModes"] = request.DBInstanceModesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceStatusesShrink)) {
		query["DBInstanceStatuses"] = request.DBInstanceStatusesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceDeployTypesShrink)) {
		query["InstanceDeployTypes"] = request.InstanceDeployTypesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceNetworkType)) {
		query["InstanceNetworkType"] = request.InstanceNetworkType
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDBInstances"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDBInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this operation to query the instance types, network types, and states of AnalyticDB for PostgreSQL instances within a specific region.
 * ## Limits
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeDBInstancesRequest
 * @return DescribeDBInstancesResponse
 */
func (client *Client) DescribeDBInstances(request *DescribeDBInstancesRequest) (_result *DescribeDBInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBInstancesResponse{}
	_body, _err := client.DescribeDBInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDataBackupsWithOptions(request *DescribeDataBackupsRequest, runtime *util.RuntimeOptions) (_result *DescribeDataBackupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BackupId)) {
		query["BackupId"] = request.BackupId
	}

	if !tea.BoolValue(util.IsUnset(request.BackupMode)) {
		query["BackupMode"] = request.BackupMode
	}

	if !tea.BoolValue(util.IsUnset(request.BackupStatus)) {
		query["BackupStatus"] = request.BackupStatus
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.DataType)) {
		query["DataType"] = request.DataType
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDataBackups"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDataBackupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDataBackups(request *DescribeDataBackupsRequest) (_result *DescribeDataBackupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDataBackupsResponse{}
	_body, _err := client.DescribeDataBackupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDataReDistributeInfoWithOptions(request *DescribeDataReDistributeInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeDataReDistributeInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDataReDistributeInfo"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDataReDistributeInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDataReDistributeInfo(request *DescribeDataReDistributeInfoRequest) (_result *DescribeDataReDistributeInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDataReDistributeInfoResponse{}
	_body, _err := client.DescribeDataReDistributeInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this operation to query the state of data sharing for AnalyticDB for PostgreSQL instances in Serverless mode.
 * ## Limits
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeDataShareInstancesRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeDataShareInstancesResponse
 */
func (client *Client) DescribeDataShareInstancesWithOptions(request *DescribeDataShareInstancesRequest, runtime *util.RuntimeOptions) (_result *DescribeDataShareInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SearchValue)) {
		query["SearchValue"] = request.SearchValue
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDataShareInstances"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDataShareInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this operation to query the state of data sharing for AnalyticDB for PostgreSQL instances in Serverless mode.
 * ## Limits
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeDataShareInstancesRequest
 * @return DescribeDataShareInstancesResponse
 */
func (client *Client) DescribeDataShareInstances(request *DescribeDataShareInstancesRequest) (_result *DescribeDataShareInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDataShareInstancesResponse{}
	_body, _err := client.DescribeDataShareInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this operation to query the details of data sharing performance metrics for an AnalyticDB for PostgreSQL instance in Serverless mode, such as the number of shared topics and the amount of data shared.
 * ## Limits
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeDataSharePerformanceRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeDataSharePerformanceResponse
 */
func (client *Client) DescribeDataSharePerformanceWithOptions(request *DescribeDataSharePerformanceRequest, runtime *util.RuntimeOptions) (_result *DescribeDataSharePerformanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Key)) {
		query["Key"] = request.Key
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDataSharePerformance"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDataSharePerformanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this operation to query the details of data sharing performance metrics for an AnalyticDB for PostgreSQL instance in Serverless mode, such as the number of shared topics and the amount of data shared.
 * ## Limits
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeDataSharePerformanceRequest
 * @return DescribeDataSharePerformanceResponse
 */
func (client *Client) DescribeDataSharePerformance(request *DescribeDataSharePerformanceRequest) (_result *DescribeDataSharePerformanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDataSharePerformanceResponse{}
	_body, _err := client.DescribeDataSharePerformanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * To facilitate management, you can call this operation to query all databases and database accounts on an AnalyticDB for PostgreSQL instance.
 * ## Limits
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeDiagnosisDimensionsRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeDiagnosisDimensionsResponse
 */
func (client *Client) DescribeDiagnosisDimensionsWithOptions(request *DescribeDiagnosisDimensionsRequest, runtime *util.RuntimeOptions) (_result *DescribeDiagnosisDimensionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDiagnosisDimensions"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDiagnosisDimensionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * To facilitate management, you can call this operation to query all databases and database accounts on an AnalyticDB for PostgreSQL instance.
 * ## Limits
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeDiagnosisDimensionsRequest
 * @return DescribeDiagnosisDimensionsResponse
 */
func (client *Client) DescribeDiagnosisDimensions(request *DescribeDiagnosisDimensionsRequest) (_result *DescribeDiagnosisDimensionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDiagnosisDimensionsResponse{}
	_body, _err := client.DescribeDiagnosisDimensionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this operation to query the details of query execution on an AnalyticDB for PostgreSQL instance in elastic storage mode within a specified time range.
 * ## Limits
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeDiagnosisMonitorPerformanceRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeDiagnosisMonitorPerformanceResponse
 */
func (client *Client) DescribeDiagnosisMonitorPerformanceWithOptions(request *DescribeDiagnosisMonitorPerformanceRequest, runtime *util.RuntimeOptions) (_result *DescribeDiagnosisMonitorPerformanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Database)) {
		query["Database"] = request.Database
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.QueryCondition)) {
		query["QueryCondition"] = request.QueryCondition
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.User)) {
		query["User"] = request.User
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDiagnosisMonitorPerformance"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDiagnosisMonitorPerformanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this operation to query the details of query execution on an AnalyticDB for PostgreSQL instance in elastic storage mode within a specified time range.
 * ## Limits
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeDiagnosisMonitorPerformanceRequest
 * @return DescribeDiagnosisMonitorPerformanceResponse
 */
func (client *Client) DescribeDiagnosisMonitorPerformance(request *DescribeDiagnosisMonitorPerformanceRequest) (_result *DescribeDiagnosisMonitorPerformanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDiagnosisMonitorPerformanceResponse{}
	_body, _err := client.DescribeDiagnosisMonitorPerformanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this operation to query the details of SQL queries on an AnalyticDB for PostgreSQL instance within a specified time range.
 * ## Limits
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeDiagnosisRecordsRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeDiagnosisRecordsResponse
 */
func (client *Client) DescribeDiagnosisRecordsWithOptions(request *DescribeDiagnosisRecordsRequest, runtime *util.RuntimeOptions) (_result *DescribeDiagnosisRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Database)) {
		query["Database"] = request.Database
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["Keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.QueryCondition)) {
		query["QueryCondition"] = request.QueryCondition
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.User)) {
		query["User"] = request.User
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDiagnosisRecords"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDiagnosisRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this operation to query the details of SQL queries on an AnalyticDB for PostgreSQL instance within a specified time range.
 * ## Limits
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeDiagnosisRecordsRequest
 * @return DescribeDiagnosisRecordsResponse
 */
func (client *Client) DescribeDiagnosisRecords(request *DescribeDiagnosisRecordsRequest) (_result *DescribeDiagnosisRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDiagnosisRecordsResponse{}
	_body, _err := client.DescribeDiagnosisRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this operation to query the details of a specified query on an AnalyticDB for PostgreSQL instance, including the SQL statement, execution plan text, and execution plan tree.
 * This operation is available only for instances of V6.3.10.1 or later in elastic storage mode. For more information about how to view and update the minor engine version of an instance, see [View the minor engine version](~~277424~~) and [Update the minor engine version](~~139271~~).
 * ## Limits
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeDiagnosisSQLInfoRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeDiagnosisSQLInfoResponse
 */
func (client *Client) DescribeDiagnosisSQLInfoWithOptions(request *DescribeDiagnosisSQLInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeDiagnosisSQLInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Database)) {
		query["Database"] = request.Database
	}

	if !tea.BoolValue(util.IsUnset(request.QueryID)) {
		query["QueryID"] = request.QueryID
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDiagnosisSQLInfo"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDiagnosisSQLInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this operation to query the details of a specified query on an AnalyticDB for PostgreSQL instance, including the SQL statement, execution plan text, and execution plan tree.
 * This operation is available only for instances of V6.3.10.1 or later in elastic storage mode. For more information about how to view and update the minor engine version of an instance, see [View the minor engine version](~~277424~~) and [Update the minor engine version](~~139271~~).
 * ## Limits
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeDiagnosisSQLInfoRequest
 * @return DescribeDiagnosisSQLInfoResponse
 */
func (client *Client) DescribeDiagnosisSQLInfo(request *DescribeDiagnosisSQLInfoRequest) (_result *DescribeDiagnosisSQLInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDiagnosisSQLInfoResponse{}
	_body, _err := client.DescribeDiagnosisSQLInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You must call the [DownloadDiagnosisRecords](~~447700~~) operation to obtain a download record before you can call this operation to query and download the query diagnostic information.
 * This operation is available only for instances of V6.3.10.1 or later in elastic storage mode. For more information about how to view and update the minor engine version of an instance, see [View the minor engine version](~~277424~~) and [Update the minor engine version](~~139271~~).
 * ## Limits
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeDownloadRecordsRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeDownloadRecordsResponse
 */
func (client *Client) DescribeDownloadRecordsWithOptions(request *DescribeDownloadRecordsRequest, runtime *util.RuntimeOptions) (_result *DescribeDownloadRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDownloadRecords"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDownloadRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You must call the [DownloadDiagnosisRecords](~~447700~~) operation to obtain a download record before you can call this operation to query and download the query diagnostic information.
 * This operation is available only for instances of V6.3.10.1 or later in elastic storage mode. For more information about how to view and update the minor engine version of an instance, see [View the minor engine version](~~277424~~) and [Update the minor engine version](~~139271~~).
 * ## Limits
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeDownloadRecordsRequest
 * @return DescribeDownloadRecordsResponse
 */
func (client *Client) DescribeDownloadRecords(request *DescribeDownloadRecordsRequest) (_result *DescribeDownloadRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDownloadRecordsResponse{}
	_body, _err := client.DescribeDownloadRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDownloadSQLLogsWithOptions(request *DescribeDownloadSQLLogsRequest, runtime *util.RuntimeOptions) (_result *DescribeDownloadSQLLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDownloadSQLLogs"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDownloadSQLLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDownloadSQLLogs(request *DescribeDownloadSQLLogsRequest) (_result *DescribeDownloadSQLLogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDownloadSQLLogsResponse{}
	_body, _err := client.DescribeDownloadSQLLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * This operation is called to query the health status of an AnalyticDB for PostgreSQL instance in elastic storage mode or Serverless mode and its coordinator and compute nodes.
 * ## Limits
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeHealthStatusRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeHealthStatusResponse
 */
func (client *Client) DescribeHealthStatusWithOptions(request *DescribeHealthStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeHealthStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Key)) {
		query["Key"] = request.Key
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeHealthStatus"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeHealthStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * This operation is called to query the health status of an AnalyticDB for PostgreSQL instance in elastic storage mode or Serverless mode and its coordinator and compute nodes.
 * ## Limits
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeHealthStatusRequest
 * @return DescribeHealthStatusResponse
 */
func (client *Client) DescribeHealthStatus(request *DescribeHealthStatusRequest) (_result *DescribeHealthStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeHealthStatusResponse{}
	_body, _err := client.DescribeHealthStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLogBackupsWithOptions(request *DescribeLogBackupsRequest, runtime *util.RuntimeOptions) (_result *DescribeLogBackupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeLogBackups"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeLogBackupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLogBackups(request *DescribeLogBackupsRequest) (_result *DescribeLogBackupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLogBackupsResponse{}
	_body, _err := client.DescribeLogBackupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeModifyParameterLogWithOptions(request *DescribeModifyParameterLogRequest, runtime *util.RuntimeOptions) (_result *DescribeModifyParameterLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeModifyParameterLog"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeModifyParameterLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeModifyParameterLog(request *DescribeModifyParameterLogRequest) (_result *DescribeModifyParameterLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeModifyParameterLogResponse{}
	_body, _err := client.DescribeModifyParameterLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeNamespaceWithOptions(request *DescribeNamespaceRequest, runtime *util.RuntimeOptions) (_result *DescribeNamespaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ManagerAccount)) {
		query["ManagerAccount"] = request.ManagerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ManagerAccountPassword)) {
		query["ManagerAccountPassword"] = request.ManagerAccountPassword
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["Namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeNamespace"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeNamespaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeNamespace(request *DescribeNamespaceRequest) (_result *DescribeNamespaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeNamespaceResponse{}
	_body, _err := client.DescribeNamespaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * This operation can be called to query the details of parameters in an AnalyticDB for PostgreSQL instance in elastic storage mode or Serverless mode.
 * ## Limits
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered and may affect your business. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeParametersRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeParametersResponse
 */
func (client *Client) DescribeParametersWithOptions(request *DescribeParametersRequest, runtime *util.RuntimeOptions) (_result *DescribeParametersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeParameters"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeParametersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * This operation can be called to query the details of parameters in an AnalyticDB for PostgreSQL instance in elastic storage mode or Serverless mode.
 * ## Limits
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered and may affect your business. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeParametersRequest
 * @return DescribeParametersResponse
 */
func (client *Client) DescribeParameters(request *DescribeParametersRequest) (_result *DescribeParametersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeParametersResponse{}
	_body, _err := client.DescribeParametersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * When you create AnalyticDB for PostgreSQL instances, you can call this operation to query the details of vSwitches within a specified region or zone.
 * ## Limits
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeRdsVSwitchsRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeRdsVSwitchsResponse
 */
func (client *Client) DescribeRdsVSwitchsWithOptions(request *DescribeRdsVSwitchsRequest, runtime *util.RuntimeOptions) (_result *DescribeRdsVSwitchsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		query["VpcId"] = request.VpcId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRdsVSwitchs"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRdsVSwitchsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * When you create AnalyticDB for PostgreSQL instances, you can call this operation to query the details of vSwitches within a specified region or zone.
 * ## Limits
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeRdsVSwitchsRequest
 * @return DescribeRdsVSwitchsResponse
 */
func (client *Client) DescribeRdsVSwitchs(request *DescribeRdsVSwitchsRequest) (_result *DescribeRdsVSwitchsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRdsVSwitchsResponse{}
	_body, _err := client.DescribeRdsVSwitchsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * When you create an AnalyticDB for PostgreSQL instance, you can call this operation to query the available VPCs within a specified region or zone.
 * ## Limits
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeRdsVpcsRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeRdsVpcsResponse
 */
func (client *Client) DescribeRdsVpcsWithOptions(request *DescribeRdsVpcsRequest, runtime *util.RuntimeOptions) (_result *DescribeRdsVpcsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRdsVpcs"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRdsVpcsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * When you create an AnalyticDB for PostgreSQL instance, you can call this operation to query the available VPCs within a specified region or zone.
 * ## Limits
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeRdsVpcsRequest
 * @return DescribeRdsVpcsResponse
 */
func (client *Client) DescribeRdsVpcs(request *DescribeRdsVpcsRequest) (_result *DescribeRdsVpcsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRdsVpcsResponse{}
	_body, _err := client.DescribeRdsVpcsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Before you create an AnalyticDB for PostgreSQL instance, you must call this operation to query available regions and zones.
 * ## Limit
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered and may affect your business. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeRegionsRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeRegionsResponse
 */
func (client *Client) DescribeRegionsWithOptions(request *DescribeRegionsRequest, runtime *util.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["Region"] = request.Region
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRegions"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Before you create an AnalyticDB for PostgreSQL instance, you must call this operation to query available regions and zones.
 * ## Limit
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered and may affect your business. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeRegionsRequest
 * @return DescribeRegionsResponse
 */
func (client *Client) DescribeRegions(request *DescribeRegionsRequest) (_result *DescribeRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSQLLogCountWithOptions(request *DescribeSQLLogCountRequest, runtime *util.RuntimeOptions) (_result *DescribeSQLLogCountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Database)) {
		query["Database"] = request.Database
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.ExecuteCost)) {
		query["ExecuteCost"] = request.ExecuteCost
	}

	if !tea.BoolValue(util.IsUnset(request.ExecuteState)) {
		query["ExecuteState"] = request.ExecuteState
	}

	if !tea.BoolValue(util.IsUnset(request.MaxExecuteCost)) {
		query["MaxExecuteCost"] = request.MaxExecuteCost
	}

	if !tea.BoolValue(util.IsUnset(request.MinExecuteCost)) {
		query["MinExecuteCost"] = request.MinExecuteCost
	}

	if !tea.BoolValue(util.IsUnset(request.OperationClass)) {
		query["OperationClass"] = request.OperationClass
	}

	if !tea.BoolValue(util.IsUnset(request.OperationType)) {
		query["OperationType"] = request.OperationType
	}

	if !tea.BoolValue(util.IsUnset(request.QueryKeywords)) {
		query["QueryKeywords"] = request.QueryKeywords
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIP)) {
		query["SourceIP"] = request.SourceIP
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.User)) {
		query["User"] = request.User
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSQLLogCount"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSQLLogCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSQLLogCount(request *DescribeSQLLogCountRequest) (_result *DescribeSQLLogCountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSQLLogCountResponse{}
	_body, _err := client.DescribeSQLLogCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this operation to query SQL logs of an AnalyticDB for PostgreSQL instance within a specific time range.
 * ## Limits
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeSQLLogsV2Request
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeSQLLogsV2Response
 */
func (client *Client) DescribeSQLLogsV2WithOptions(request *DescribeSQLLogsV2Request, runtime *util.RuntimeOptions) (_result *DescribeSQLLogsV2Response, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Database)) {
		query["Database"] = request.Database
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.ExecuteCost)) {
		query["ExecuteCost"] = request.ExecuteCost
	}

	if !tea.BoolValue(util.IsUnset(request.ExecuteState)) {
		query["ExecuteState"] = request.ExecuteState
	}

	if !tea.BoolValue(util.IsUnset(request.MaxExecuteCost)) {
		query["MaxExecuteCost"] = request.MaxExecuteCost
	}

	if !tea.BoolValue(util.IsUnset(request.MinExecuteCost)) {
		query["MinExecuteCost"] = request.MinExecuteCost
	}

	if !tea.BoolValue(util.IsUnset(request.OperationClass)) {
		query["OperationClass"] = request.OperationClass
	}

	if !tea.BoolValue(util.IsUnset(request.OperationType)) {
		query["OperationType"] = request.OperationType
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.QueryKeywords)) {
		query["QueryKeywords"] = request.QueryKeywords
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIP)) {
		query["SourceIP"] = request.SourceIP
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.User)) {
		query["User"] = request.User
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSQLLogsV2"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSQLLogsV2Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this operation to query SQL logs of an AnalyticDB for PostgreSQL instance within a specific time range.
 * ## Limits
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeSQLLogsV2Request
 * @return DescribeSQLLogsV2Response
 */
func (client *Client) DescribeSQLLogsV2(request *DescribeSQLLogsV2Request) (_result *DescribeSQLLogsV2Response, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSQLLogsV2Response{}
	_body, _err := client.DescribeSQLLogsV2WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeSampleDataRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeSampleDataResponse
 */
func (client *Client) DescribeSampleDataWithOptions(request *DescribeSampleDataRequest, runtime *util.RuntimeOptions) (_result *DescribeSampleDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSampleData"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSampleDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeSampleDataRequest
 * @return DescribeSampleDataResponse
 */
func (client *Client) DescribeSampleData(request *DescribeSampleDataRequest) (_result *DescribeSampleDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSampleDataResponse{}
	_body, _err := client.DescribeSampleDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSupportFeaturesWithOptions(request *DescribeSupportFeaturesRequest, runtime *util.RuntimeOptions) (_result *DescribeSupportFeaturesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSupportFeatures"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSupportFeaturesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSupportFeatures(request *DescribeSupportFeaturesRequest) (_result *DescribeSupportFeaturesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSupportFeaturesResponse{}
	_body, _err := client.DescribeSupportFeaturesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeTagsWithOptions(request *DescribeTagsRequest, runtime *util.RuntimeOptions) (_result *DescribeTagsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeTags"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeTagsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTags(request *DescribeTagsRequest) (_result *DescribeTagsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTagsResponse{}
	_body, _err := client.DescribeTagsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeUserEncryptionKeyListWithOptions(request *DescribeUserEncryptionKeyListRequest, runtime *util.RuntimeOptions) (_result *DescribeUserEncryptionKeyListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeUserEncryptionKeyList"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeUserEncryptionKeyListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeUserEncryptionKeyList(request *DescribeUserEncryptionKeyListRequest) (_result *DescribeUserEncryptionKeyListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUserEncryptionKeyListResponse{}
	_body, _err := client.DescribeUserEncryptionKeyListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this operation to query the details of a lock-waiting query only for an AnalyticDB for PostgreSQL V6.0 instance in elastic storage mode.
 * ## Limits
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeWaitingSQLInfoRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeWaitingSQLInfoResponse
 */
func (client *Client) DescribeWaitingSQLInfoWithOptions(request *DescribeWaitingSQLInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeWaitingSQLInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Database)) {
		query["Database"] = request.Database
	}

	if !tea.BoolValue(util.IsUnset(request.PID)) {
		query["PID"] = request.PID
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeWaitingSQLInfo"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeWaitingSQLInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this operation to query the details of a lock-waiting query only for an AnalyticDB for PostgreSQL V6.0 instance in elastic storage mode.
 * ## Limits
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeWaitingSQLInfoRequest
 * @return DescribeWaitingSQLInfoResponse
 */
func (client *Client) DescribeWaitingSQLInfo(request *DescribeWaitingSQLInfoRequest) (_result *DescribeWaitingSQLInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeWaitingSQLInfoResponse{}
	_body, _err := client.DescribeWaitingSQLInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this operation to query the lock diagnostics records only for an AnalyticDB for PostgreSQL V6.0 instance in elastic storage mode.
 * ## Limits
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeWaitingSQLRecordsRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeWaitingSQLRecordsResponse
 */
func (client *Client) DescribeWaitingSQLRecordsWithOptions(request *DescribeWaitingSQLRecordsRequest, runtime *util.RuntimeOptions) (_result *DescribeWaitingSQLRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Database)) {
		query["Database"] = request.Database
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["Keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.QueryCondition)) {
		query["QueryCondition"] = request.QueryCondition
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.User)) {
		query["User"] = request.User
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeWaitingSQLRecords"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeWaitingSQLRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this operation to query the lock diagnostics records only for an AnalyticDB for PostgreSQL V6.0 instance in elastic storage mode.
 * ## Limits
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeWaitingSQLRecordsRequest
 * @return DescribeWaitingSQLRecordsResponse
 */
func (client *Client) DescribeWaitingSQLRecords(request *DescribeWaitingSQLRecordsRequest) (_result *DescribeWaitingSQLRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeWaitingSQLRecordsResponse{}
	_body, _err := client.DescribeWaitingSQLRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this operation to download the query diagnostic information of an AnalyticDB for PostgreSQL instance. After the download is complete, you can call the [DescribeDownloadRecords](~~447712~~) operation to query download records and download URLs.
 * This operation is available only for instances of V6.3.10.1 or later in elastic storage mode. For more information about how to view and update the minor engine version of an instance, see [View the minor engine version](~~277424~~) and [Update the minor engine version](~~139271~~).
 * ## Limits
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DownloadDiagnosisRecordsRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DownloadDiagnosisRecordsResponse
 */
func (client *Client) DownloadDiagnosisRecordsWithOptions(request *DownloadDiagnosisRecordsRequest, runtime *util.RuntimeOptions) (_result *DownloadDiagnosisRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Database)) {
		query["Database"] = request.Database
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.QueryCondition)) {
		query["QueryCondition"] = request.QueryCondition
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.User)) {
		query["User"] = request.User
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DownloadDiagnosisRecords"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DownloadDiagnosisRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this operation to download the query diagnostic information of an AnalyticDB for PostgreSQL instance. After the download is complete, you can call the [DescribeDownloadRecords](~~447712~~) operation to query download records and download URLs.
 * This operation is available only for instances of V6.3.10.1 or later in elastic storage mode. For more information about how to view and update the minor engine version of an instance, see [View the minor engine version](~~277424~~) and [Update the minor engine version](~~139271~~).
 * ## Limits
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DownloadDiagnosisRecordsRequest
 * @return DownloadDiagnosisRecordsResponse
 */
func (client *Client) DownloadDiagnosisRecords(request *DownloadDiagnosisRecordsRequest) (_result *DownloadDiagnosisRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DownloadDiagnosisRecordsResponse{}
	_body, _err := client.DownloadDiagnosisRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DownloadSQLLogsRecordsWithOptions(request *DownloadSQLLogsRecordsRequest, runtime *util.RuntimeOptions) (_result *DownloadSQLLogsRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Database)) {
		query["Database"] = request.Database
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.ExecuteCost)) {
		query["ExecuteCost"] = request.ExecuteCost
	}

	if !tea.BoolValue(util.IsUnset(request.ExecuteState)) {
		query["ExecuteState"] = request.ExecuteState
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.MaxExecuteCost)) {
		query["MaxExecuteCost"] = request.MaxExecuteCost
	}

	if !tea.BoolValue(util.IsUnset(request.MinExecuteCost)) {
		query["MinExecuteCost"] = request.MinExecuteCost
	}

	if !tea.BoolValue(util.IsUnset(request.OperationClass)) {
		query["OperationClass"] = request.OperationClass
	}

	if !tea.BoolValue(util.IsUnset(request.OperationType)) {
		query["OperationType"] = request.OperationType
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.QueryKeywords)) {
		query["QueryKeywords"] = request.QueryKeywords
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIP)) {
		query["SourceIP"] = request.SourceIP
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.User)) {
		query["User"] = request.User
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DownloadSQLLogsRecords"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DownloadSQLLogsRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DownloadSQLLogsRecords(request *DownloadSQLLogsRecordsRequest) (_result *DownloadSQLLogsRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DownloadSQLLogsRecordsResponse{}
	_body, _err := client.DownloadSQLLogsRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GrantCollectionWithOptions(request *GrantCollectionRequest, runtime *util.RuntimeOptions) (_result *GrantCollectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Collection)) {
		query["Collection"] = request.Collection
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.GrantToNamespace)) {
		query["GrantToNamespace"] = request.GrantToNamespace
	}

	if !tea.BoolValue(util.IsUnset(request.GrantType)) {
		query["GrantType"] = request.GrantType
	}

	if !tea.BoolValue(util.IsUnset(request.ManagerAccount)) {
		query["ManagerAccount"] = request.ManagerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ManagerAccountPassword)) {
		query["ManagerAccountPassword"] = request.ManagerAccountPassword
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["Namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GrantCollection"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GrantCollectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GrantCollection(request *GrantCollectionRequest) (_result *GrantCollectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GrantCollectionResponse{}
	_body, _err := client.GrantCollectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InitVectorDatabaseWithOptions(request *InitVectorDatabaseRequest, runtime *util.RuntimeOptions) (_result *InitVectorDatabaseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ManagerAccount)) {
		query["ManagerAccount"] = request.ManagerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ManagerAccountPassword)) {
		query["ManagerAccountPassword"] = request.ManagerAccountPassword
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("InitVectorDatabase"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &InitVectorDatabaseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InitVectorDatabase(request *InitVectorDatabaseRequest) (_result *InitVectorDatabaseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InitVectorDatabaseResponse{}
	_body, _err := client.InitVectorDatabaseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListCollectionsWithOptions(request *ListCollectionsRequest, runtime *util.RuntimeOptions) (_result *ListCollectionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["Namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.NamespacePassword)) {
		query["NamespacePassword"] = request.NamespacePassword
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListCollections"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListCollectionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListCollections(request *ListCollectionsRequest) (_result *ListCollectionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListCollectionsResponse{}
	_body, _err := client.ListCollectionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListNamespacesWithOptions(request *ListNamespacesRequest, runtime *util.RuntimeOptions) (_result *ListNamespacesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ManagerAccount)) {
		query["ManagerAccount"] = request.ManagerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ManagerAccountPassword)) {
		query["ManagerAccountPassword"] = request.ManagerAccountPassword
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListNamespaces"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListNamespacesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListNamespaces(request *ListNamespacesRequest) (_result *ListNamespacesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListNamespacesResponse{}
	_body, _err := client.ListNamespacesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, runtime *util.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagResources"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTagResources(request *ListTagResourcesRequest) (_result *ListTagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.ListTagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyAccountDescriptionWithOptions(request *ModifyAccountDescriptionRequest, runtime *util.RuntimeOptions) (_result *ModifyAccountDescriptionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountDescription)) {
		query["AccountDescription"] = request.AccountDescription
	}

	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		query["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyAccountDescription"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyAccountDescriptionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyAccountDescription(request *ModifyAccountDescriptionRequest) (_result *ModifyAccountDescriptionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyAccountDescriptionResponse{}
	_body, _err := client.ModifyAccountDescriptionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyBackupPolicyWithOptions(request *ModifyBackupPolicyRequest, runtime *util.RuntimeOptions) (_result *ModifyBackupPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BackupRetentionPeriod)) {
		query["BackupRetentionPeriod"] = request.BackupRetentionPeriod
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.EnableRecoveryPoint)) {
		query["EnableRecoveryPoint"] = request.EnableRecoveryPoint
	}

	if !tea.BoolValue(util.IsUnset(request.PreferredBackupPeriod)) {
		query["PreferredBackupPeriod"] = request.PreferredBackupPeriod
	}

	if !tea.BoolValue(util.IsUnset(request.PreferredBackupTime)) {
		query["PreferredBackupTime"] = request.PreferredBackupTime
	}

	if !tea.BoolValue(util.IsUnset(request.RecoveryPointPeriod)) {
		query["RecoveryPointPeriod"] = request.RecoveryPointPeriod
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyBackupPolicy"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyBackupPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyBackupPolicy(request *ModifyBackupPolicyRequest) (_result *ModifyBackupPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyBackupPolicyResponse{}
	_body, _err := client.ModifyBackupPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDBInstanceConfigWithOptions(request *ModifyDBInstanceConfigRequest, runtime *util.RuntimeOptions) (_result *ModifyDBInstanceConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceDescription)) {
		query["DBInstanceDescription"] = request.DBInstanceDescription
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.IdleTime)) {
		query["IdleTime"] = request.IdleTime
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ServerlessResource)) {
		query["ServerlessResource"] = request.ServerlessResource
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDBInstanceConfig"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDBInstanceConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDBInstanceConfig(request *ModifyDBInstanceConfigRequest) (_result *ModifyDBInstanceConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDBInstanceConfigResponse{}
	_body, _err := client.ModifyDBInstanceConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDBInstanceConnectionStringWithOptions(request *ModifyDBInstanceConnectionStringRequest, runtime *util.RuntimeOptions) (_result *ModifyDBInstanceConnectionStringResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConnectionStringPrefix)) {
		query["ConnectionStringPrefix"] = request.ConnectionStringPrefix
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentConnectionString)) {
		query["CurrentConnectionString"] = request.CurrentConnectionString
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Port)) {
		query["Port"] = request.Port
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDBInstanceConnectionString"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDBInstanceConnectionStringResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDBInstanceConnectionString(request *ModifyDBInstanceConnectionStringRequest) (_result *ModifyDBInstanceConnectionStringResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDBInstanceConnectionStringResponse{}
	_body, _err := client.ModifyDBInstanceConnectionStringWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * To make it easy to identify AnalyticDB for PostgreSQL instances, you can call this operation to modify the description of instances.
 * ## Limits
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request ModifyDBInstanceDescriptionRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ModifyDBInstanceDescriptionResponse
 */
func (client *Client) ModifyDBInstanceDescriptionWithOptions(request *ModifyDBInstanceDescriptionRequest, runtime *util.RuntimeOptions) (_result *ModifyDBInstanceDescriptionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceDescription)) {
		query["DBInstanceDescription"] = request.DBInstanceDescription
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDBInstanceDescription"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDBInstanceDescriptionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * To make it easy to identify AnalyticDB for PostgreSQL instances, you can call this operation to modify the description of instances.
 * ## Limits
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request ModifyDBInstanceDescriptionRequest
 * @return ModifyDBInstanceDescriptionResponse
 */
func (client *Client) ModifyDBInstanceDescription(request *ModifyDBInstanceDescriptionRequest) (_result *ModifyDBInstanceDescriptionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDBInstanceDescriptionResponse{}
	_body, _err := client.ModifyDBInstanceDescriptionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * The system maintains AnalyticDB for PostgreSQL instances during the maintenance window that you specify. We recommend that you set the maintenance window to off-peak hours to minimize the impact on your business.
 * ## Limits
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request ModifyDBInstanceMaintainTimeRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ModifyDBInstanceMaintainTimeResponse
 */
func (client *Client) ModifyDBInstanceMaintainTimeWithOptions(request *ModifyDBInstanceMaintainTimeRequest, runtime *util.RuntimeOptions) (_result *ModifyDBInstanceMaintainTimeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDBInstanceMaintainTime"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDBInstanceMaintainTimeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * The system maintains AnalyticDB for PostgreSQL instances during the maintenance window that you specify. We recommend that you set the maintenance window to off-peak hours to minimize the impact on your business.
 * ## Limits
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request ModifyDBInstanceMaintainTimeRequest
 * @return ModifyDBInstanceMaintainTimeResponse
 */
func (client *Client) ModifyDBInstanceMaintainTime(request *ModifyDBInstanceMaintainTimeRequest) (_result *ModifyDBInstanceMaintainTimeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDBInstanceMaintainTimeResponse{}
	_body, _err := client.ModifyDBInstanceMaintainTimeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Resource Management allows you to build an organizational structure for resources based on your business requirements. You can use resource directories, folders, accounts, and resource groups to hierarchically organize and manage resources. For more information, see [What is Resource Management?](~~94475~~)
 *
 * @param request ModifyDBInstanceResourceGroupRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ModifyDBInstanceResourceGroupResponse
 */
func (client *Client) ModifyDBInstanceResourceGroupWithOptions(request *ModifyDBInstanceResourceGroupRequest, runtime *util.RuntimeOptions) (_result *ModifyDBInstanceResourceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.NewResourceGroupId)) {
		query["NewResourceGroupId"] = request.NewResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDBInstanceResourceGroup"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDBInstanceResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Resource Management allows you to build an organizational structure for resources based on your business requirements. You can use resource directories, folders, accounts, and resource groups to hierarchically organize and manage resources. For more information, see [What is Resource Management?](~~94475~~)
 *
 * @param request ModifyDBInstanceResourceGroupRequest
 * @return ModifyDBInstanceResourceGroupResponse
 */
func (client *Client) ModifyDBInstanceResourceGroup(request *ModifyDBInstanceResourceGroupRequest) (_result *ModifyDBInstanceResourceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDBInstanceResourceGroupResponse{}
	_body, _err := client.ModifyDBInstanceResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDBInstanceSSLWithOptions(request *ModifyDBInstanceSSLRequest, runtime *util.RuntimeOptions) (_result *ModifyDBInstanceSSLResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConnectionString)) {
		query["ConnectionString"] = request.ConnectionString
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.SSLEnabled)) {
		query["SSLEnabled"] = request.SSLEnabled
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDBInstanceSSL"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDBInstanceSSLResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDBInstanceSSL(request *ModifyDBInstanceSSLRequest) (_result *ModifyDBInstanceSSLResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDBInstanceSSLResponse{}
	_body, _err := client.ModifyDBInstanceSSLWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * This operation can be called to modify parameters of an AnalyticDB for PostgreSQL instance in elastic storage mode or Serverless mode.
 * ## Limits
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered and may affect your business. We recommend that you take note of the limit when you call this operation.
 *
 * @param request ModifyParametersRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ModifyParametersResponse
 */
func (client *Client) ModifyParametersWithOptions(request *ModifyParametersRequest, runtime *util.RuntimeOptions) (_result *ModifyParametersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ForceRestartInstance)) {
		query["ForceRestartInstance"] = request.ForceRestartInstance
	}

	if !tea.BoolValue(util.IsUnset(request.Parameters)) {
		query["Parameters"] = request.Parameters
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyParameters"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyParametersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * This operation can be called to modify parameters of an AnalyticDB for PostgreSQL instance in elastic storage mode or Serverless mode.
 * ## Limits
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered and may affect your business. We recommend that you take note of the limit when you call this operation.
 *
 * @param request ModifyParametersRequest
 * @return ModifyParametersResponse
 */
func (client *Client) ModifyParameters(request *ModifyParametersRequest) (_result *ModifyParametersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyParametersResponse{}
	_body, _err := client.ModifyParametersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifySQLCollectorPolicyWithOptions(request *ModifySQLCollectorPolicyRequest, runtime *util.RuntimeOptions) (_result *ModifySQLCollectorPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.SQLCollectorStatus)) {
		query["SQLCollectorStatus"] = request.SQLCollectorStatus
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifySQLCollectorPolicy"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifySQLCollectorPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifySQLCollectorPolicy(request *ModifySQLCollectorPolicyRequest) (_result *ModifySQLCollectorPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifySQLCollectorPolicyResponse{}
	_body, _err := client.ModifySQLCollectorPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * To ensure the security and stability of AnalyticDB for PostgreSQL instances, the system denies all external IP addresses to access AnalyticDB for PostgreSQL instances by default. Before you can use an AnalyticDB for PostgreSQL instance, you must add the IP address or CIDR block of your client to the IP address whitelist of the instance.
 * ## Limits
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request ModifySecurityIpsRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ModifySecurityIpsResponse
 */
func (client *Client) ModifySecurityIpsWithOptions(request *ModifySecurityIpsRequest, runtime *util.RuntimeOptions) (_result *ModifySecurityIpsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceIPArrayAttribute)) {
		query["DBInstanceIPArrayAttribute"] = request.DBInstanceIPArrayAttribute
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceIPArrayName)) {
		query["DBInstanceIPArrayName"] = request.DBInstanceIPArrayName
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ModifyMode)) {
		query["ModifyMode"] = request.ModifyMode
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityIPList)) {
		query["SecurityIPList"] = request.SecurityIPList
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifySecurityIps"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifySecurityIpsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * To ensure the security and stability of AnalyticDB for PostgreSQL instances, the system denies all external IP addresses to access AnalyticDB for PostgreSQL instances by default. Before you can use an AnalyticDB for PostgreSQL instance, you must add the IP address or CIDR block of your client to the IP address whitelist of the instance.
 * ## Limits
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request ModifySecurityIpsRequest
 * @return ModifySecurityIpsResponse
 */
func (client *Client) ModifySecurityIps(request *ModifySecurityIpsRequest) (_result *ModifySecurityIpsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifySecurityIpsResponse{}
	_body, _err := client.ModifySecurityIpsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyVectorConfigurationWithOptions(request *ModifyVectorConfigurationRequest, runtime *util.RuntimeOptions) (_result *ModifyVectorConfigurationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.VectorConfigurationStatus)) {
		query["VectorConfigurationStatus"] = request.VectorConfigurationStatus
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyVectorConfiguration"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyVectorConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyVectorConfiguration(request *ModifyVectorConfigurationRequest) (_result *ModifyVectorConfigurationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyVectorConfigurationResponse{}
	_body, _err := client.ModifyVectorConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this operation to pause an AnalyticDB for PostgreSQL instance that is in the **Running** state.
 * This operation is available only for AnalyticDB for PostgreSQL instances in Serverless mode that run V1.0.2.1 or later. For more information about how to view and update the minor engine version of an instance, see [View the minor engine version](~~277424~~) and [Update the minor engine version](~~139271~~).
 * >  Before you call this operation, make sure that you are familiar with the billing methods and pricing of AnalyticDB for PostgreSQL instances. For more information, see [Billing methods](~~35406~~) and [AnalyticDB for PostgreSQL pricing](https://www.alibabacloud.com/zh/product/hybriddb-postgresql/pricing).
 * ## Limits
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request PauseInstanceRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return PauseInstanceResponse
 */
func (client *Client) PauseInstanceWithOptions(request *PauseInstanceRequest, runtime *util.RuntimeOptions) (_result *PauseInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PauseInstance"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PauseInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this operation to pause an AnalyticDB for PostgreSQL instance that is in the **Running** state.
 * This operation is available only for AnalyticDB for PostgreSQL instances in Serverless mode that run V1.0.2.1 or later. For more information about how to view and update the minor engine version of an instance, see [View the minor engine version](~~277424~~) and [Update the minor engine version](~~139271~~).
 * >  Before you call this operation, make sure that you are familiar with the billing methods and pricing of AnalyticDB for PostgreSQL instances. For more information, see [Billing methods](~~35406~~) and [AnalyticDB for PostgreSQL pricing](https://www.alibabacloud.com/zh/product/hybriddb-postgresql/pricing).
 * ## Limits
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request PauseInstanceRequest
 * @return PauseInstanceResponse
 */
func (client *Client) PauseInstance(request *PauseInstanceRequest) (_result *PauseInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PauseInstanceResponse{}
	_body, _err := client.PauseInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryCollectionDataWithOptions(tmpReq *QueryCollectionDataRequest, runtime *util.RuntimeOptions) (_result *QueryCollectionDataResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &QueryCollectionDataShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Vector)) {
		request.VectorShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Vector, tea.String("Vector"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Collection)) {
		query["Collection"] = request.Collection
	}

	if !tea.BoolValue(util.IsUnset(request.Content)) {
		query["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Filter)) {
		query["Filter"] = request.Filter
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["Namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.NamespacePassword)) {
		query["NamespacePassword"] = request.NamespacePassword
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TopK)) {
		query["TopK"] = request.TopK
	}

	if !tea.BoolValue(util.IsUnset(request.VectorShrink)) {
		query["Vector"] = request.VectorShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryCollectionData"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryCollectionDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryCollectionData(request *QueryCollectionDataRequest) (_result *QueryCollectionDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryCollectionDataResponse{}
	_body, _err := client.QueryCollectionDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RebalanceDBInstanceWithOptions(request *RebalanceDBInstanceRequest, runtime *util.RuntimeOptions) (_result *RebalanceDBInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RebalanceDBInstance"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RebalanceDBInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RebalanceDBInstance(request *RebalanceDBInstanceRequest) (_result *RebalanceDBInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RebalanceDBInstanceResponse{}
	_body, _err := client.RebalanceDBInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReleaseInstancePublicConnectionWithOptions(request *ReleaseInstancePublicConnectionRequest, runtime *util.RuntimeOptions) (_result *ReleaseInstancePublicConnectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AddressType)) {
		query["AddressType"] = request.AddressType
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentConnectionString)) {
		query["CurrentConnectionString"] = request.CurrentConnectionString
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ReleaseInstancePublicConnection"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReleaseInstancePublicConnectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReleaseInstancePublicConnection(request *ReleaseInstancePublicConnectionRequest) (_result *ReleaseInstancePublicConnectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReleaseInstancePublicConnectionResponse{}
	_body, _err := client.ReleaseInstancePublicConnectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResetAccountPasswordWithOptions(request *ResetAccountPasswordRequest, runtime *util.RuntimeOptions) (_result *ResetAccountPasswordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		query["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.AccountPassword)) {
		query["AccountPassword"] = request.AccountPassword
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ResetAccountPassword"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ResetAccountPasswordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResetAccountPassword(request *ResetAccountPasswordRequest) (_result *ResetAccountPasswordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResetAccountPasswordResponse{}
	_body, _err := client.ResetAccountPasswordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * A restart takes about 3 to 30 minutes. During the restart, services are unavailable. We recommend that you restart the instance during off-peak hours. After the instance is restarted and enters the running state, you can access the instance.
 * ## Limit
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered and may affect your business. We recommend that you take note of the limit when you call this operation.
 *
 * @param request RestartDBInstanceRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return RestartDBInstanceResponse
 */
func (client *Client) RestartDBInstanceWithOptions(request *RestartDBInstanceRequest, runtime *util.RuntimeOptions) (_result *RestartDBInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RestartDBInstance"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RestartDBInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * A restart takes about 3 to 30 minutes. During the restart, services are unavailable. We recommend that you restart the instance during off-peak hours. After the instance is restarted and enters the running state, you can access the instance.
 * ## Limit
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered and may affect your business. We recommend that you take note of the limit when you call this operation.
 *
 * @param request RestartDBInstanceRequest
 * @return RestartDBInstanceResponse
 */
func (client *Client) RestartDBInstance(request *RestartDBInstanceRequest) (_result *RestartDBInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RestartDBInstanceResponse{}
	_body, _err := client.RestartDBInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this operation to resume an AnalyticDB for PostgreSQL instance that is in the **Paused** state.
 * This operation is available only for AnalyticDB for PostgreSQL instances in Serverless mode that run V1.0.2.1 or later. For more information about how to view and update the minor engine version of an instance, see [View the minor engine version](~~277424~~) and [Update the minor engine version](~~139271~~).
 * >  Before you call this operation, make sure that you are familiar with the billing methods and pricing of AnalyticDB for PostgreSQL instances. For more information, see [Billing methods](~~35406~~) and [AnalyticDB for PostgreSQL pricing](https://www.alibabacloud.com/zh/product/hybriddb-postgresql/pricing).
 * ## Limits
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request ResumeInstanceRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ResumeInstanceResponse
 */
func (client *Client) ResumeInstanceWithOptions(request *ResumeInstanceRequest, runtime *util.RuntimeOptions) (_result *ResumeInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ResumeInstance"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ResumeInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this operation to resume an AnalyticDB for PostgreSQL instance that is in the **Paused** state.
 * This operation is available only for AnalyticDB for PostgreSQL instances in Serverless mode that run V1.0.2.1 or later. For more information about how to view and update the minor engine version of an instance, see [View the minor engine version](~~277424~~) and [Update the minor engine version](~~139271~~).
 * >  Before you call this operation, make sure that you are familiar with the billing methods and pricing of AnalyticDB for PostgreSQL instances. For more information, see [Billing methods](~~35406~~) and [AnalyticDB for PostgreSQL pricing](https://www.alibabacloud.com/zh/product/hybriddb-postgresql/pricing).
 * ## Limits
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request ResumeInstanceRequest
 * @return ResumeInstanceResponse
 */
func (client *Client) ResumeInstance(request *ResumeInstanceRequest) (_result *ResumeInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResumeInstanceResponse{}
	_body, _err := client.ResumeInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this operation to enable or disable a specified plan. The plan management feature is supported only for AnalyticDB for PostgreSQL instances in Serverless mode.
 * ## Limits
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request SetDBInstancePlanStatusRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return SetDBInstancePlanStatusResponse
 */
func (client *Client) SetDBInstancePlanStatusWithOptions(request *SetDBInstancePlanStatusRequest, runtime *util.RuntimeOptions) (_result *SetDBInstancePlanStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PlanId)) {
		query["PlanId"] = request.PlanId
	}

	if !tea.BoolValue(util.IsUnset(request.PlanStatus)) {
		query["PlanStatus"] = request.PlanStatus
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetDBInstancePlanStatus"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetDBInstancePlanStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this operation to enable or disable a specified plan. The plan management feature is supported only for AnalyticDB for PostgreSQL instances in Serverless mode.
 * ## Limits
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request SetDBInstancePlanStatusRequest
 * @return SetDBInstancePlanStatusResponse
 */
func (client *Client) SetDBInstancePlanStatus(request *SetDBInstancePlanStatusRequest) (_result *SetDBInstancePlanStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetDBInstancePlanStatusResponse{}
	_body, _err := client.SetDBInstancePlanStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * This operation is called to enable or disable data sharing for an AnalyticDB for PostgreSQL instance in Serverless mode.
 * ## Limits
 * You can call this operation up to 1,000 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation.
 *
 * @param tmpReq SetDataShareInstanceRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return SetDataShareInstanceResponse
 */
func (client *Client) SetDataShareInstanceWithOptions(tmpReq *SetDataShareInstanceRequest, runtime *util.RuntimeOptions) (_result *SetDataShareInstanceResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SetDataShareInstanceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.InstanceList)) {
		request.InstanceListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InstanceList, tea.String("InstanceList"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceListShrink)) {
		query["InstanceList"] = request.InstanceListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.OperationType)) {
		query["OperationType"] = request.OperationType
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetDataShareInstance"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetDataShareInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * This operation is called to enable or disable data sharing for an AnalyticDB for PostgreSQL instance in Serverless mode.
 * ## Limits
 * You can call this operation up to 1,000 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation.
 *
 * @param request SetDataShareInstanceRequest
 * @return SetDataShareInstanceResponse
 */
func (client *Client) SetDataShareInstance(request *SetDataShareInstanceRequest) (_result *SetDataShareInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetDataShareInstanceResponse{}
	_body, _err := client.SetDataShareInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SwitchDBInstanceNetTypeWithOptions(request *SwitchDBInstanceNetTypeRequest, runtime *util.RuntimeOptions) (_result *SwitchDBInstanceNetTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConnectionStringPrefix)) {
		query["ConnectionStringPrefix"] = request.ConnectionStringPrefix
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Port)) {
		query["Port"] = request.Port
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SwitchDBInstanceNetType"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SwitchDBInstanceNetTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SwitchDBInstanceNetType(request *SwitchDBInstanceNetTypeRequest) (_result *SwitchDBInstanceNetTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SwitchDBInstanceNetTypeResponse{}
	_body, _err := client.SwitchDBInstanceNetTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, runtime *util.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TagResources"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TagResources(request *TagResourcesRequest) (_result *TagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TagResourcesResponse{}
	_body, _err := client.TagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this operation to release a sample dataset from an AnalyticDB for PostgreSQL instance. You must have already loaded the sample dataset.
 * ## Limits
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request UnloadSampleDataRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return UnloadSampleDataResponse
 */
func (client *Client) UnloadSampleDataWithOptions(request *UnloadSampleDataRequest, runtime *util.RuntimeOptions) (_result *UnloadSampleDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UnloadSampleData"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UnloadSampleDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this operation to release a sample dataset from an AnalyticDB for PostgreSQL instance. You must have already loaded the sample dataset.
 * ## Limits
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request UnloadSampleDataRequest
 * @return UnloadSampleDataResponse
 */
func (client *Client) UnloadSampleData(request *UnloadSampleDataRequest) (_result *UnloadSampleDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnloadSampleDataResponse{}
	_body, _err := client.UnloadSampleDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UntagResourcesWithOptions(request *UntagResourcesRequest, runtime *util.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.All)) {
		query["All"] = request.All
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.TagKey)) {
		query["TagKey"] = request.TagKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UntagResources"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UntagResources(request *UntagResourcesRequest) (_result *UntagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UntagResourcesResponse{}
	_body, _err := client.UntagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this operation to modify a plan for an AnalyticDB for PostgreSQL instance in Serverless mode. For example, you can modify a plan for periodically pausing and resuming an instance or scaling an instance.
 * ## Limits
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request UpdateDBInstancePlanRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return UpdateDBInstancePlanResponse
 */
func (client *Client) UpdateDBInstancePlanWithOptions(request *UpdateDBInstancePlanRequest, runtime *util.RuntimeOptions) (_result *UpdateDBInstancePlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PlanConfig)) {
		query["PlanConfig"] = request.PlanConfig
	}

	if !tea.BoolValue(util.IsUnset(request.PlanDesc)) {
		query["PlanDesc"] = request.PlanDesc
	}

	if !tea.BoolValue(util.IsUnset(request.PlanEndDate)) {
		query["PlanEndDate"] = request.PlanEndDate
	}

	if !tea.BoolValue(util.IsUnset(request.PlanId)) {
		query["PlanId"] = request.PlanId
	}

	if !tea.BoolValue(util.IsUnset(request.PlanName)) {
		query["PlanName"] = request.PlanName
	}

	if !tea.BoolValue(util.IsUnset(request.PlanStartDate)) {
		query["PlanStartDate"] = request.PlanStartDate
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateDBInstancePlan"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateDBInstancePlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this operation to modify a plan for an AnalyticDB for PostgreSQL instance in Serverless mode. For example, you can modify a plan for periodically pausing and resuming an instance or scaling an instance.
 * ## Limits
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request UpdateDBInstancePlanRequest
 * @return UpdateDBInstancePlanResponse
 */
func (client *Client) UpdateDBInstancePlan(request *UpdateDBInstancePlanRequest) (_result *UpdateDBInstancePlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateDBInstancePlanResponse{}
	_body, _err := client.UpdateDBInstancePlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this operation to change the configurations of an AnalyticDB for PostgreSQL instance.
 * >  This operation is not supported for instances in reserved storage mode.
 * Before you call this operation, make sure that you are familiar with the billing of AnalyticDB for PostgreSQL. For more information, see [Billing methods](~~35406~~) and [AnalyticDB for PostgreSQL pricing](https://www.alibabacloud.com/zh/product/hybriddb-postgresql/pricing).
 * ## Limits
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request UpgradeDBInstanceRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return UpgradeDBInstanceResponse
 */
func (client *Client) UpgradeDBInstanceWithOptions(request *UpgradeDBInstanceRequest, runtime *util.RuntimeOptions) (_result *UpgradeDBInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceClass)) {
		query["DBInstanceClass"] = request.DBInstanceClass
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceGroupCount)) {
		query["DBInstanceGroupCount"] = request.DBInstanceGroupCount
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceSpec)) {
		query["InstanceSpec"] = request.InstanceSpec
	}

	if !tea.BoolValue(util.IsUnset(request.MasterNodeNum)) {
		query["MasterNodeNum"] = request.MasterNodeNum
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PayType)) {
		query["PayType"] = request.PayType
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SegDiskPerformanceLevel)) {
		query["SegDiskPerformanceLevel"] = request.SegDiskPerformanceLevel
	}

	if !tea.BoolValue(util.IsUnset(request.SegNodeNum)) {
		query["SegNodeNum"] = request.SegNodeNum
	}

	if !tea.BoolValue(util.IsUnset(request.SegStorageType)) {
		query["SegStorageType"] = request.SegStorageType
	}

	if !tea.BoolValue(util.IsUnset(request.StorageSize)) {
		query["StorageSize"] = request.StorageSize
	}

	if !tea.BoolValue(util.IsUnset(request.UpgradeType)) {
		query["UpgradeType"] = request.UpgradeType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpgradeDBInstance"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpgradeDBInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this operation to change the configurations of an AnalyticDB for PostgreSQL instance.
 * >  This operation is not supported for instances in reserved storage mode.
 * Before you call this operation, make sure that you are familiar with the billing of AnalyticDB for PostgreSQL. For more information, see [Billing methods](~~35406~~) and [AnalyticDB for PostgreSQL pricing](https://www.alibabacloud.com/zh/product/hybriddb-postgresql/pricing).
 * ## Limits
 * You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request UpgradeDBInstanceRequest
 * @return UpgradeDBInstanceResponse
 */
func (client *Client) UpgradeDBInstance(request *UpgradeDBInstanceRequest) (_result *UpgradeDBInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpgradeDBInstanceResponse{}
	_body, _err := client.UpgradeDBInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpgradeDBVersionWithOptions(request *UpgradeDBVersionRequest, runtime *util.RuntimeOptions) (_result *UpgradeDBVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.MajorVersion)) {
		query["MajorVersion"] = request.MajorVersion
	}

	if !tea.BoolValue(util.IsUnset(request.MinorVersion)) {
		query["MinorVersion"] = request.MinorVersion
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SwitchTime)) {
		query["SwitchTime"] = request.SwitchTime
	}

	if !tea.BoolValue(util.IsUnset(request.SwitchTimeMode)) {
		query["SwitchTimeMode"] = request.SwitchTimeMode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpgradeDBVersion"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpgradeDBVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpgradeDBVersion(request *UpgradeDBVersionRequest) (_result *UpgradeDBVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpgradeDBVersionResponse{}
	_body, _err := client.UpgradeDBVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpsertCollectionDataWithOptions(tmpReq *UpsertCollectionDataRequest, runtime *util.RuntimeOptions) (_result *UpsertCollectionDataResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpsertCollectionDataShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Rows)) {
		request.RowsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Rows, tea.String("Rows"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Collection)) {
		query["Collection"] = request.Collection
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["Namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.NamespacePassword)) {
		query["NamespacePassword"] = request.NamespacePassword
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RowsShrink)) {
		query["Rows"] = request.RowsShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpsertCollectionData"),
		Version:     tea.String("2016-05-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpsertCollectionDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpsertCollectionData(request *UpsertCollectionDataRequest) (_result *UpsertCollectionDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpsertCollectionDataResponse{}
	_body, _err := client.UpsertCollectionDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
