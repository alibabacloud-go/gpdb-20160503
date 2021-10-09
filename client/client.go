// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddBuDBInstanceRelationRequest struct {
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	BusinessUnit *string `json:"BusinessUnit,omitempty" xml:"BusinessUnit,omitempty"`
}

func (s AddBuDBInstanceRelationRequest) String() string {
	return tea.Prettify(s)
}

func (s AddBuDBInstanceRelationRequest) GoString() string {
	return s.String()
}

func (s *AddBuDBInstanceRelationRequest) SetOwnerId(v int64) *AddBuDBInstanceRelationRequest {
	s.OwnerId = &v
	return s
}

func (s *AddBuDBInstanceRelationRequest) SetDBInstanceId(v string) *AddBuDBInstanceRelationRequest {
	s.DBInstanceId = &v
	return s
}

func (s *AddBuDBInstanceRelationRequest) SetBusinessUnit(v string) *AddBuDBInstanceRelationRequest {
	s.BusinessUnit = &v
	return s
}

type AddBuDBInstanceRelationResponseBody struct {
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	BusinessUnit   *string `json:"BusinessUnit,omitempty" xml:"BusinessUnit,omitempty"`
}

func (s AddBuDBInstanceRelationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddBuDBInstanceRelationResponseBody) GoString() string {
	return s.String()
}

func (s *AddBuDBInstanceRelationResponseBody) SetDBInstanceName(v string) *AddBuDBInstanceRelationResponseBody {
	s.DBInstanceName = &v
	return s
}

func (s *AddBuDBInstanceRelationResponseBody) SetRequestId(v string) *AddBuDBInstanceRelationResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddBuDBInstanceRelationResponseBody) SetBusinessUnit(v string) *AddBuDBInstanceRelationResponseBody {
	s.BusinessUnit = &v
	return s
}

type AddBuDBInstanceRelationResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddBuDBInstanceRelationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddBuDBInstanceRelationResponse) String() string {
	return tea.Prettify(s)
}

func (s AddBuDBInstanceRelationResponse) GoString() string {
	return s.String()
}

func (s *AddBuDBInstanceRelationResponse) SetHeaders(v map[string]*string) *AddBuDBInstanceRelationResponse {
	s.Headers = v
	return s
}

func (s *AddBuDBInstanceRelationResponse) SetBody(v *AddBuDBInstanceRelationResponseBody) *AddBuDBInstanceRelationResponse {
	s.Body = v
	return s
}

type AllocateInstancePublicConnectionRequest struct {
	OwnerId                *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount   *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId        *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	DBInstanceId           *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	ConnectionStringPrefix *string `json:"ConnectionStringPrefix,omitempty" xml:"ConnectionStringPrefix,omitempty"`
	Port                   *string `json:"Port,omitempty" xml:"Port,omitempty"`
	AddressType            *string `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
}

func (s AllocateInstancePublicConnectionRequest) String() string {
	return tea.Prettify(s)
}

func (s AllocateInstancePublicConnectionRequest) GoString() string {
	return s.String()
}

func (s *AllocateInstancePublicConnectionRequest) SetOwnerId(v int64) *AllocateInstancePublicConnectionRequest {
	s.OwnerId = &v
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

func (s *AllocateInstancePublicConnectionRequest) SetDBInstanceId(v string) *AllocateInstancePublicConnectionRequest {
	s.DBInstanceId = &v
	return s
}

func (s *AllocateInstancePublicConnectionRequest) SetConnectionStringPrefix(v string) *AllocateInstancePublicConnectionRequest {
	s.ConnectionStringPrefix = &v
	return s
}

func (s *AllocateInstancePublicConnectionRequest) SetPort(v string) *AllocateInstancePublicConnectionRequest {
	s.Port = &v
	return s
}

func (s *AllocateInstancePublicConnectionRequest) SetAddressType(v string) *AllocateInstancePublicConnectionRequest {
	s.AddressType = &v
	return s
}

type AllocateInstancePublicConnectionResponseBody struct {
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
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AllocateInstancePublicConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *AllocateInstancePublicConnectionResponse) SetBody(v *AllocateInstancePublicConnectionResponseBody) *AllocateInstancePublicConnectionResponse {
	s.Body = v
	return s
}

type CheckServiceLinkedRoleRequest struct {
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
	RequestId            *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	HasServiceLinkedRole *string `json:"HasServiceLinkedRole,omitempty" xml:"HasServiceLinkedRole,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CheckServiceLinkedRoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckServiceLinkedRoleResponseBody) GoString() string {
	return s.String()
}

func (s *CheckServiceLinkedRoleResponseBody) SetRequestId(v string) *CheckServiceLinkedRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckServiceLinkedRoleResponseBody) SetHasServiceLinkedRole(v string) *CheckServiceLinkedRoleResponseBody {
	s.HasServiceLinkedRole = &v
	return s
}

func (s *CheckServiceLinkedRoleResponseBody) SetRegionId(v string) *CheckServiceLinkedRoleResponseBody {
	s.RegionId = &v
	return s
}

type CheckServiceLinkedRoleResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CheckServiceLinkedRoleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CheckServiceLinkedRoleResponse) SetBody(v *CheckServiceLinkedRoleResponseBody) *CheckServiceLinkedRoleResponse {
	s.Body = v
	return s
}

type CreateAccountRequest struct {
	OwnerId            *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	DBInstanceId       *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	DatabaseName       *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	AccountName        *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	AccountPassword    *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
	AccountDescription *string `json:"AccountDescription,omitempty" xml:"AccountDescription,omitempty"`
}

func (s CreateAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAccountRequest) GoString() string {
	return s.String()
}

func (s *CreateAccountRequest) SetOwnerId(v int64) *CreateAccountRequest {
	s.OwnerId = &v
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

func (s *CreateAccountRequest) SetAccountName(v string) *CreateAccountRequest {
	s.AccountName = &v
	return s
}

func (s *CreateAccountRequest) SetAccountPassword(v string) *CreateAccountRequest {
	s.AccountPassword = &v
	return s
}

func (s *CreateAccountRequest) SetAccountDescription(v string) *CreateAccountRequest {
	s.AccountDescription = &v
	return s
}

type CreateAccountResponseBody struct {
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
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateAccountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateAccountResponse) SetBody(v *CreateAccountResponseBody) *CreateAccountResponse {
	s.Body = v
	return s
}

type CreateDBInstanceRequest struct {
	OwnerId               *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId              *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ZoneId                *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	EngineVersion         *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	Engine                *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	DBInstanceClass       *string `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty"`
	DBInstanceGroupCount  *string `json:"DBInstanceGroupCount,omitempty" xml:"DBInstanceGroupCount,omitempty"`
	DBInstanceDescription *string `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	SecurityIPList        *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	PayType               *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	Period                *string `json:"Period,omitempty" xml:"Period,omitempty"`
	UsedTime              *string `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	ClientToken           *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceNetworkType   *string `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
	VPCId                 *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	VSwitchId             *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	PrivateIpAddress      *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
}

func (s CreateDBInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDBInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateDBInstanceRequest) SetOwnerId(v int64) *CreateDBInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetRegionId(v string) *CreateDBInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetZoneId(v string) *CreateDBInstanceRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetEngineVersion(v string) *CreateDBInstanceRequest {
	s.EngineVersion = &v
	return s
}

func (s *CreateDBInstanceRequest) SetEngine(v string) *CreateDBInstanceRequest {
	s.Engine = &v
	return s
}

func (s *CreateDBInstanceRequest) SetDBInstanceClass(v string) *CreateDBInstanceRequest {
	s.DBInstanceClass = &v
	return s
}

func (s *CreateDBInstanceRequest) SetDBInstanceGroupCount(v string) *CreateDBInstanceRequest {
	s.DBInstanceGroupCount = &v
	return s
}

func (s *CreateDBInstanceRequest) SetDBInstanceDescription(v string) *CreateDBInstanceRequest {
	s.DBInstanceDescription = &v
	return s
}

func (s *CreateDBInstanceRequest) SetSecurityIPList(v string) *CreateDBInstanceRequest {
	s.SecurityIPList = &v
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

func (s *CreateDBInstanceRequest) SetUsedTime(v string) *CreateDBInstanceRequest {
	s.UsedTime = &v
	return s
}

func (s *CreateDBInstanceRequest) SetClientToken(v string) *CreateDBInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDBInstanceRequest) SetInstanceNetworkType(v string) *CreateDBInstanceRequest {
	s.InstanceNetworkType = &v
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

func (s *CreateDBInstanceRequest) SetPrivateIpAddress(v string) *CreateDBInstanceRequest {
	s.PrivateIpAddress = &v
	return s
}

type CreateDBInstanceResponseBody struct {
	DBInstanceId     *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Port             *string `json:"Port,omitempty" xml:"Port,omitempty"`
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	OrderId          *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s CreateDBInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDBInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDBInstanceResponseBody) SetDBInstanceId(v string) *CreateDBInstanceResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *CreateDBInstanceResponseBody) SetRequestId(v string) *CreateDBInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDBInstanceResponseBody) SetPort(v string) *CreateDBInstanceResponseBody {
	s.Port = &v
	return s
}

func (s *CreateDBInstanceResponseBody) SetConnectionString(v string) *CreateDBInstanceResponseBody {
	s.ConnectionString = &v
	return s
}

func (s *CreateDBInstanceResponseBody) SetOrderId(v string) *CreateDBInstanceResponseBody {
	s.OrderId = &v
	return s
}

type CreateDBInstanceResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDBInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateDBInstanceResponse) SetBody(v *CreateDBInstanceResponseBody) *CreateDBInstanceResponse {
	s.Body = v
	return s
}

type CreateECSDBInstanceRequest struct {
	OwnerId               *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId              *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ZoneId                *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	EngineVersion         *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	Engine                *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	InstanceSpec          *string `json:"InstanceSpec,omitempty" xml:"InstanceSpec,omitempty"`
	SegNodeNum            *int32  `json:"SegNodeNum,omitempty" xml:"SegNodeNum,omitempty"`
	SegStorageType        *string `json:"SegStorageType,omitempty" xml:"SegStorageType,omitempty"`
	StorageSize           *int32  `json:"StorageSize,omitempty" xml:"StorageSize,omitempty"`
	DBInstanceDescription *string `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	SecurityIPList        *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	PayType               *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	Period                *string `json:"Period,omitempty" xml:"Period,omitempty"`
	UsedTime              *string `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	ClientToken           *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceNetworkType   *string `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
	VPCId                 *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	VSwitchId             *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	PrivateIpAddress      *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	EncryptionKey         *string `json:"EncryptionKey,omitempty" xml:"EncryptionKey,omitempty"`
	EncryptionType        *string `json:"EncryptionType,omitempty" xml:"EncryptionType,omitempty"`
	MasterNodeNum         *int32  `json:"MasterNodeNum,omitempty" xml:"MasterNodeNum,omitempty"`
	SrcDbInstanceName     *string `json:"SrcDbInstanceName,omitempty" xml:"SrcDbInstanceName,omitempty"`
	BackupId              *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	DBInstanceCategory    *string `json:"DBInstanceCategory,omitempty" xml:"DBInstanceCategory,omitempty"`
}

func (s CreateECSDBInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateECSDBInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateECSDBInstanceRequest) SetOwnerId(v int64) *CreateECSDBInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateECSDBInstanceRequest) SetRegionId(v string) *CreateECSDBInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateECSDBInstanceRequest) SetZoneId(v string) *CreateECSDBInstanceRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateECSDBInstanceRequest) SetEngineVersion(v string) *CreateECSDBInstanceRequest {
	s.EngineVersion = &v
	return s
}

func (s *CreateECSDBInstanceRequest) SetEngine(v string) *CreateECSDBInstanceRequest {
	s.Engine = &v
	return s
}

func (s *CreateECSDBInstanceRequest) SetInstanceSpec(v string) *CreateECSDBInstanceRequest {
	s.InstanceSpec = &v
	return s
}

func (s *CreateECSDBInstanceRequest) SetSegNodeNum(v int32) *CreateECSDBInstanceRequest {
	s.SegNodeNum = &v
	return s
}

func (s *CreateECSDBInstanceRequest) SetSegStorageType(v string) *CreateECSDBInstanceRequest {
	s.SegStorageType = &v
	return s
}

func (s *CreateECSDBInstanceRequest) SetStorageSize(v int32) *CreateECSDBInstanceRequest {
	s.StorageSize = &v
	return s
}

func (s *CreateECSDBInstanceRequest) SetDBInstanceDescription(v string) *CreateECSDBInstanceRequest {
	s.DBInstanceDescription = &v
	return s
}

func (s *CreateECSDBInstanceRequest) SetSecurityIPList(v string) *CreateECSDBInstanceRequest {
	s.SecurityIPList = &v
	return s
}

func (s *CreateECSDBInstanceRequest) SetPayType(v string) *CreateECSDBInstanceRequest {
	s.PayType = &v
	return s
}

func (s *CreateECSDBInstanceRequest) SetPeriod(v string) *CreateECSDBInstanceRequest {
	s.Period = &v
	return s
}

func (s *CreateECSDBInstanceRequest) SetUsedTime(v string) *CreateECSDBInstanceRequest {
	s.UsedTime = &v
	return s
}

func (s *CreateECSDBInstanceRequest) SetClientToken(v string) *CreateECSDBInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateECSDBInstanceRequest) SetInstanceNetworkType(v string) *CreateECSDBInstanceRequest {
	s.InstanceNetworkType = &v
	return s
}

func (s *CreateECSDBInstanceRequest) SetVPCId(v string) *CreateECSDBInstanceRequest {
	s.VPCId = &v
	return s
}

func (s *CreateECSDBInstanceRequest) SetVSwitchId(v string) *CreateECSDBInstanceRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateECSDBInstanceRequest) SetPrivateIpAddress(v string) *CreateECSDBInstanceRequest {
	s.PrivateIpAddress = &v
	return s
}

func (s *CreateECSDBInstanceRequest) SetEncryptionKey(v string) *CreateECSDBInstanceRequest {
	s.EncryptionKey = &v
	return s
}

func (s *CreateECSDBInstanceRequest) SetEncryptionType(v string) *CreateECSDBInstanceRequest {
	s.EncryptionType = &v
	return s
}

func (s *CreateECSDBInstanceRequest) SetMasterNodeNum(v int32) *CreateECSDBInstanceRequest {
	s.MasterNodeNum = &v
	return s
}

func (s *CreateECSDBInstanceRequest) SetSrcDbInstanceName(v string) *CreateECSDBInstanceRequest {
	s.SrcDbInstanceName = &v
	return s
}

func (s *CreateECSDBInstanceRequest) SetBackupId(v string) *CreateECSDBInstanceRequest {
	s.BackupId = &v
	return s
}

func (s *CreateECSDBInstanceRequest) SetDBInstanceCategory(v string) *CreateECSDBInstanceRequest {
	s.DBInstanceCategory = &v
	return s
}

type CreateECSDBInstanceResponseBody struct {
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DBInstanceId     *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	Port             *string `json:"Port,omitempty" xml:"Port,omitempty"`
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	OrderId          *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s CreateECSDBInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateECSDBInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateECSDBInstanceResponseBody) SetRequestId(v string) *CreateECSDBInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateECSDBInstanceResponseBody) SetDBInstanceId(v string) *CreateECSDBInstanceResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *CreateECSDBInstanceResponseBody) SetPort(v string) *CreateECSDBInstanceResponseBody {
	s.Port = &v
	return s
}

func (s *CreateECSDBInstanceResponseBody) SetConnectionString(v string) *CreateECSDBInstanceResponseBody {
	s.ConnectionString = &v
	return s
}

func (s *CreateECSDBInstanceResponseBody) SetOrderId(v string) *CreateECSDBInstanceResponseBody {
	s.OrderId = &v
	return s
}

type CreateECSDBInstanceResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateECSDBInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateECSDBInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateECSDBInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateECSDBInstanceResponse) SetHeaders(v map[string]*string) *CreateECSDBInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateECSDBInstanceResponse) SetBody(v *CreateECSDBInstanceResponseBody) *CreateECSDBInstanceResponse {
	s.Body = v
	return s
}

type CreateServiceLinkedRoleRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s CreateServiceLinkedRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceLinkedRoleRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceLinkedRoleRequest) SetRegionId(v string) *CreateServiceLinkedRoleRequest {
	s.RegionId = &v
	return s
}

func (s *CreateServiceLinkedRoleRequest) SetOwnerId(v int64) *CreateServiceLinkedRoleRequest {
	s.OwnerId = &v
	return s
}

type CreateServiceLinkedRoleResponseBody struct {
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
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateServiceLinkedRoleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateServiceLinkedRoleResponse) SetBody(v *CreateServiceLinkedRoleResponseBody) *CreateServiceLinkedRoleResponse {
	s.Body = v
	return s
}

type DeleteDatabaseRequest struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	DBName       *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
}

func (s DeleteDatabaseRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDatabaseRequest) GoString() string {
	return s.String()
}

func (s *DeleteDatabaseRequest) SetDBInstanceId(v string) *DeleteDatabaseRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DeleteDatabaseRequest) SetDBName(v string) *DeleteDatabaseRequest {
	s.DBName = &v
	return s
}

type DeleteDatabaseResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDatabaseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDatabaseResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDatabaseResponseBody) SetRequestId(v string) *DeleteDatabaseResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDatabaseResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDatabaseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDatabaseResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDatabaseResponse) GoString() string {
	return s.String()
}

func (s *DeleteDatabaseResponse) SetHeaders(v map[string]*string) *DeleteDatabaseResponse {
	s.Headers = v
	return s
}

func (s *DeleteDatabaseResponse) SetBody(v *DeleteDatabaseResponseBody) *DeleteDatabaseResponse {
	s.Body = v
	return s
}

type DeleteDBInstanceRequest struct {
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s DeleteDBInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDBInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeleteDBInstanceRequest) SetOwnerId(v int64) *DeleteDBInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteDBInstanceRequest) SetClientToken(v string) *DeleteDBInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteDBInstanceRequest) SetDBInstanceId(v string) *DeleteDBInstanceRequest {
	s.DBInstanceId = &v
	return s
}

type DeleteDBInstanceResponseBody struct {
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
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDBInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteDBInstanceResponse) SetBody(v *DeleteDBInstanceResponseBody) *DeleteDBInstanceResponse {
	s.Body = v
	return s
}

type DescribeAccountsRequest struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	AccountName  *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
}

func (s DescribeAccountsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAccountsRequest) SetDBInstanceId(v string) *DescribeAccountsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeAccountsRequest) SetAccountName(v string) *DescribeAccountsRequest {
	s.AccountName = &v
	return s
}

type DescribeAccountsResponseBody struct {
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Accounts  *DescribeAccountsResponseBodyAccounts `json:"Accounts,omitempty" xml:"Accounts,omitempty" type:"Struct"`
}

func (s DescribeAccountsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAccountsResponseBody) SetRequestId(v string) *DescribeAccountsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAccountsResponseBody) SetAccounts(v *DescribeAccountsResponseBodyAccounts) *DescribeAccountsResponseBody {
	s.Accounts = v
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
	AccountDescription *string `json:"AccountDescription,omitempty" xml:"AccountDescription,omitempty"`
	DBInstanceId       *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	AccountStatus      *string `json:"AccountStatus,omitempty" xml:"AccountStatus,omitempty"`
	AccountName        *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
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

func (s *DescribeAccountsResponseBodyAccountsDBInstanceAccount) SetDBInstanceId(v string) *DescribeAccountsResponseBodyAccountsDBInstanceAccount {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeAccountsResponseBodyAccountsDBInstanceAccount) SetAccountStatus(v string) *DescribeAccountsResponseBodyAccountsDBInstanceAccount {
	s.AccountStatus = &v
	return s
}

func (s *DescribeAccountsResponseBodyAccountsDBInstanceAccount) SetAccountName(v string) *DescribeAccountsResponseBodyAccountsDBInstanceAccount {
	s.AccountName = &v
	return s
}

type DescribeAccountsResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAccountsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeAccountsResponse) SetBody(v *DescribeAccountsResponseBody) *DescribeAccountsResponse {
	s.Body = v
	return s
}

type DescribeAvailableResourcesRequest struct {
	Region     *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ZoneId     *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
}

func (s DescribeAvailableResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourcesRequest) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourcesRequest) SetRegion(v string) *DescribeAvailableResourcesRequest {
	s.Region = &v
	return s
}

func (s *DescribeAvailableResourcesRequest) SetZoneId(v string) *DescribeAvailableResourcesRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeAvailableResourcesRequest) SetChargeType(v string) *DescribeAvailableResourcesRequest {
	s.ChargeType = &v
	return s
}

type DescribeAvailableResourcesResponseBody struct {
	RegionId  *string                                            `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RequestId *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	ZoneId           *string                                                            `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	SupportedEngines []*DescribeAvailableResourcesResponseBodyResourcesSupportedEngines `json:"SupportedEngines,omitempty" xml:"SupportedEngines,omitempty" type:"Repeated"`
}

func (s DescribeAvailableResourcesResponseBodyResources) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourcesResponseBodyResources) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourcesResponseBodyResources) SetZoneId(v string) *DescribeAvailableResourcesResponseBodyResources {
	s.ZoneId = &v
	return s
}

func (s *DescribeAvailableResourcesResponseBodyResources) SetSupportedEngines(v []*DescribeAvailableResourcesResponseBodyResourcesSupportedEngines) *DescribeAvailableResourcesResponseBodyResources {
	s.SupportedEngines = v
	return s
}

type DescribeAvailableResourcesResponseBodyResourcesSupportedEngines struct {
	SupportedEngineVersion   *string                                                                                    `json:"SupportedEngineVersion,omitempty" xml:"SupportedEngineVersion,omitempty"`
	Mode                     *string                                                                                    `json:"Mode,omitempty" xml:"Mode,omitempty"`
	SupportedInstanceClasses []*DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClasses `json:"SupportedInstanceClasses,omitempty" xml:"SupportedInstanceClasses,omitempty" type:"Repeated"`
}

func (s DescribeAvailableResourcesResponseBodyResourcesSupportedEngines) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourcesResponseBodyResourcesSupportedEngines) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourcesResponseBodyResourcesSupportedEngines) SetSupportedEngineVersion(v string) *DescribeAvailableResourcesResponseBodyResourcesSupportedEngines {
	s.SupportedEngineVersion = &v
	return s
}

func (s *DescribeAvailableResourcesResponseBodyResourcesSupportedEngines) SetMode(v string) *DescribeAvailableResourcesResponseBodyResourcesSupportedEngines {
	s.Mode = &v
	return s
}

func (s *DescribeAvailableResourcesResponseBodyResourcesSupportedEngines) SetSupportedInstanceClasses(v []*DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClasses) *DescribeAvailableResourcesResponseBodyResourcesSupportedEngines {
	s.SupportedInstanceClasses = v
	return s
}

type DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClasses struct {
	StorageType   *string                                                                                             `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	Description   *string                                                                                             `json:"Description,omitempty" xml:"Description,omitempty"`
	DisplayClass  *string                                                                                             `json:"DisplayClass,omitempty" xml:"DisplayClass,omitempty"`
	InstanceClass *string                                                                                             `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	NodeCount     *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesNodeCount   `json:"NodeCount,omitempty" xml:"NodeCount,omitempty" type:"Struct"`
	StorageSize   *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesStorageSize `json:"StorageSize,omitempty" xml:"StorageSize,omitempty" type:"Struct"`
}

func (s DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClasses) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClasses) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClasses) SetStorageType(v string) *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClasses {
	s.StorageType = &v
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

type DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesNodeCount struct {
	Step     *string `json:"Step,omitempty" xml:"Step,omitempty"`
	MinCount *string `json:"MinCount,omitempty" xml:"MinCount,omitempty"`
	MaxCount *string `json:"MaxCount,omitempty" xml:"MaxCount,omitempty"`
}

func (s DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesNodeCount) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesNodeCount) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesNodeCount) SetStep(v string) *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesNodeCount {
	s.Step = &v
	return s
}

func (s *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesNodeCount) SetMinCount(v string) *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesNodeCount {
	s.MinCount = &v
	return s
}

func (s *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesNodeCount) SetMaxCount(v string) *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesNodeCount {
	s.MaxCount = &v
	return s
}

type DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesStorageSize struct {
	Step     *string `json:"Step,omitempty" xml:"Step,omitempty"`
	MinCount *string `json:"MinCount,omitempty" xml:"MinCount,omitempty"`
	MaxCount *string `json:"MaxCount,omitempty" xml:"MaxCount,omitempty"`
}

func (s DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesStorageSize) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesStorageSize) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesStorageSize) SetStep(v string) *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesStorageSize {
	s.Step = &v
	return s
}

func (s *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesStorageSize) SetMinCount(v string) *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesStorageSize {
	s.MinCount = &v
	return s
}

func (s *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesStorageSize) SetMaxCount(v string) *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesStorageSize {
	s.MaxCount = &v
	return s
}

type DescribeAvailableResourcesResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAvailableResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeAvailableResourcesResponse) SetBody(v *DescribeAvailableResourcesResponseBody) *DescribeAvailableResourcesResponse {
	s.Body = v
	return s
}

type DescribeBackupPolicyRequest struct {
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
	BackupRetentionPeriod *int32  `json:"BackupRetentionPeriod,omitempty" xml:"BackupRetentionPeriod,omitempty"`
	RequestId             *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PreferredBackupPeriod *string `json:"PreferredBackupPeriod,omitempty" xml:"PreferredBackupPeriod,omitempty"`
	PreferredBackupTime   *string `json:"PreferredBackupTime,omitempty" xml:"PreferredBackupTime,omitempty"`
	RecoveryPointPeriod   *string `json:"RecoveryPointPeriod,omitempty" xml:"RecoveryPointPeriod,omitempty"`
	EnableRecoveryPoint   *bool   `json:"EnableRecoveryPoint,omitempty" xml:"EnableRecoveryPoint,omitempty"`
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

func (s *DescribeBackupPolicyResponseBody) SetRequestId(v string) *DescribeBackupPolicyResponseBody {
	s.RequestId = &v
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

func (s *DescribeBackupPolicyResponseBody) SetEnableRecoveryPoint(v bool) *DescribeBackupPolicyResponseBody {
	s.EnableRecoveryPoint = &v
	return s
}

type DescribeBackupPolicyResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeBackupPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeBackupPolicyResponse) SetBody(v *DescribeBackupPolicyResponseBody) *DescribeBackupPolicyResponse {
	s.Body = v
	return s
}

type DescribeDataBackupsRequest struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	BackupId     *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber   *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	DataType     *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	BackupMode   *string `json:"BackupMode,omitempty" xml:"BackupMode,omitempty"`
	BackupStatus *string `json:"BackupStatus,omitempty" xml:"BackupStatus,omitempty"`
}

func (s DescribeDataBackupsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataBackupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDataBackupsRequest) SetDBInstanceId(v string) *DescribeDataBackupsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDataBackupsRequest) SetBackupId(v string) *DescribeDataBackupsRequest {
	s.BackupId = &v
	return s
}

func (s *DescribeDataBackupsRequest) SetStartTime(v string) *DescribeDataBackupsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDataBackupsRequest) SetEndTime(v string) *DescribeDataBackupsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDataBackupsRequest) SetPageSize(v int32) *DescribeDataBackupsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDataBackupsRequest) SetPageNumber(v int32) *DescribeDataBackupsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDataBackupsRequest) SetDataType(v string) *DescribeDataBackupsRequest {
	s.DataType = &v
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

type DescribeDataBackupsResponseBody struct {
	PageSize   *int32                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *int32                                  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	TotalCount *int32                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Items      []*DescribeDataBackupsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
}

func (s DescribeDataBackupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataBackupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDataBackupsResponseBody) SetPageSize(v int32) *DescribeDataBackupsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDataBackupsResponseBody) SetRequestId(v string) *DescribeDataBackupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDataBackupsResponseBody) SetPageNumber(v int32) *DescribeDataBackupsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDataBackupsResponseBody) SetTotalCount(v int32) *DescribeDataBackupsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDataBackupsResponseBody) SetItems(v []*DescribeDataBackupsResponseBodyItems) *DescribeDataBackupsResponseBody {
	s.Items = v
	return s
}

type DescribeDataBackupsResponseBodyItems struct {
	DataType             *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	ConsistentTime       *int64  `json:"ConsistentTime,omitempty" xml:"ConsistentTime,omitempty"`
	BackupStatus         *string `json:"BackupStatus,omitempty" xml:"BackupStatus,omitempty"`
	BackupStartTime      *string `json:"BackupStartTime,omitempty" xml:"BackupStartTime,omitempty"`
	BackupEndTime        *string `json:"BackupEndTime,omitempty" xml:"BackupEndTime,omitempty"`
	BackupEndTimeLocal   *string `json:"BackupEndTimeLocal,omitempty" xml:"BackupEndTimeLocal,omitempty"`
	BackupSetId          *string `json:"BackupSetId,omitempty" xml:"BackupSetId,omitempty"`
	BaksetName           *string `json:"BaksetName,omitempty" xml:"BaksetName,omitempty"`
	BackupSize           *int64  `json:"BackupSize,omitempty" xml:"BackupSize,omitempty"`
	BackupMode           *string `json:"BackupMode,omitempty" xml:"BackupMode,omitempty"`
	BackupStartTimeLocal *string `json:"BackupStartTimeLocal,omitempty" xml:"BackupStartTimeLocal,omitempty"`
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s DescribeDataBackupsResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataBackupsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDataBackupsResponseBodyItems) SetDataType(v string) *DescribeDataBackupsResponseBodyItems {
	s.DataType = &v
	return s
}

func (s *DescribeDataBackupsResponseBodyItems) SetConsistentTime(v int64) *DescribeDataBackupsResponseBodyItems {
	s.ConsistentTime = &v
	return s
}

func (s *DescribeDataBackupsResponseBodyItems) SetBackupStatus(v string) *DescribeDataBackupsResponseBodyItems {
	s.BackupStatus = &v
	return s
}

func (s *DescribeDataBackupsResponseBodyItems) SetBackupStartTime(v string) *DescribeDataBackupsResponseBodyItems {
	s.BackupStartTime = &v
	return s
}

func (s *DescribeDataBackupsResponseBodyItems) SetBackupEndTime(v string) *DescribeDataBackupsResponseBodyItems {
	s.BackupEndTime = &v
	return s
}

func (s *DescribeDataBackupsResponseBodyItems) SetBackupEndTimeLocal(v string) *DescribeDataBackupsResponseBodyItems {
	s.BackupEndTimeLocal = &v
	return s
}

func (s *DescribeDataBackupsResponseBodyItems) SetBackupSetId(v string) *DescribeDataBackupsResponseBodyItems {
	s.BackupSetId = &v
	return s
}

func (s *DescribeDataBackupsResponseBodyItems) SetBaksetName(v string) *DescribeDataBackupsResponseBodyItems {
	s.BaksetName = &v
	return s
}

func (s *DescribeDataBackupsResponseBodyItems) SetBackupSize(v int64) *DescribeDataBackupsResponseBodyItems {
	s.BackupSize = &v
	return s
}

func (s *DescribeDataBackupsResponseBodyItems) SetBackupMode(v string) *DescribeDataBackupsResponseBodyItems {
	s.BackupMode = &v
	return s
}

func (s *DescribeDataBackupsResponseBodyItems) SetBackupStartTimeLocal(v string) *DescribeDataBackupsResponseBodyItems {
	s.BackupStartTimeLocal = &v
	return s
}

func (s *DescribeDataBackupsResponseBodyItems) SetDBInstanceId(v string) *DescribeDataBackupsResponseBodyItems {
	s.DBInstanceId = &v
	return s
}

type DescribeDataBackupsResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDataBackupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeDataBackupsResponse) SetBody(v *DescribeDataBackupsResponseBody) *DescribeDataBackupsResponse {
	s.Body = v
	return s
}

type DescribeDBClusterPerformanceRequest struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	Key          *string `json:"Key,omitempty" xml:"Key,omitempty"`
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
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

func (s *DescribeDBClusterPerformanceRequest) SetKey(v string) *DescribeDBClusterPerformanceRequest {
	s.Key = &v
	return s
}

func (s *DescribeDBClusterPerformanceRequest) SetStartTime(v string) *DescribeDBClusterPerformanceRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDBClusterPerformanceRequest) SetEndTime(v string) *DescribeDBClusterPerformanceRequest {
	s.EndTime = &v
	return s
}

type DescribeDBClusterPerformanceResponseBody struct {
	EndTime         *string                                                    `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime       *string                                                    `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	RequestId       *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DBClusterId     *string                                                    `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	PerformanceKeys []*DescribeDBClusterPerformanceResponseBodyPerformanceKeys `json:"PerformanceKeys,omitempty" xml:"PerformanceKeys,omitempty" type:"Repeated"`
}

func (s DescribeDBClusterPerformanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterPerformanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterPerformanceResponseBody) SetEndTime(v string) *DescribeDBClusterPerformanceResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBody) SetStartTime(v string) *DescribeDBClusterPerformanceResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBody) SetRequestId(v string) *DescribeDBClusterPerformanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBody) SetDBClusterId(v string) *DescribeDBClusterPerformanceResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBody) SetPerformanceKeys(v []*DescribeDBClusterPerformanceResponseBodyPerformanceKeys) *DescribeDBClusterPerformanceResponseBody {
	s.PerformanceKeys = v
	return s
}

type DescribeDBClusterPerformanceResponseBodyPerformanceKeys struct {
	Name   *string                                                          `json:"Name,omitempty" xml:"Name,omitempty"`
	Unit   *string                                                          `json:"Unit,omitempty" xml:"Unit,omitempty"`
	Series []*DescribeDBClusterPerformanceResponseBodyPerformanceKeysSeries `json:"Series,omitempty" xml:"Series,omitempty" type:"Repeated"`
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

func (s *DescribeDBClusterPerformanceResponseBodyPerformanceKeys) SetUnit(v string) *DescribeDBClusterPerformanceResponseBodyPerformanceKeys {
	s.Unit = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformanceKeys) SetSeries(v []*DescribeDBClusterPerformanceResponseBodyPerformanceKeysSeries) *DescribeDBClusterPerformanceResponseBodyPerformanceKeys {
	s.Series = v
	return s
}

type DescribeDBClusterPerformanceResponseBodyPerformanceKeysSeries struct {
	Role   *string                                                                `json:"Role,omitempty" xml:"Role,omitempty"`
	Name   *string                                                                `json:"Name,omitempty" xml:"Name,omitempty"`
	Values []*DescribeDBClusterPerformanceResponseBodyPerformanceKeysSeriesValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s DescribeDBClusterPerformanceResponseBodyPerformanceKeysSeries) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterPerformanceResponseBodyPerformanceKeysSeries) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformanceKeysSeries) SetRole(v string) *DescribeDBClusterPerformanceResponseBodyPerformanceKeysSeries {
	s.Role = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformanceKeysSeries) SetName(v string) *DescribeDBClusterPerformanceResponseBodyPerformanceKeysSeries {
	s.Name = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformanceKeysSeries) SetValues(v []*DescribeDBClusterPerformanceResponseBodyPerformanceKeysSeriesValues) *DescribeDBClusterPerformanceResponseBodyPerformanceKeysSeries {
	s.Values = v
	return s
}

type DescribeDBClusterPerformanceResponseBodyPerformanceKeysSeriesValues struct {
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
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDBClusterPerformanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeDBClusterPerformanceResponse) SetBody(v *DescribeDBClusterPerformanceResponseBody) *DescribeDBClusterPerformanceResponse {
	s.Body = v
	return s
}

type DescribeDBInstanceAttributeRequest struct {
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s DescribeDBInstanceAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeRequest) SetOwnerId(v int64) *DescribeDBInstanceAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBInstanceAttributeRequest) SetDBInstanceId(v string) *DescribeDBInstanceAttributeRequest {
	s.DBInstanceId = &v
	return s
}

type DescribeDBInstanceAttributeResponseBody struct {
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Items     *DescribeDBInstanceAttributeResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
}

func (s DescribeDBInstanceAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponseBody) SetRequestId(v string) *DescribeDBInstanceAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBody) SetItems(v *DescribeDBInstanceAttributeResponseBodyItems) *DescribeDBInstanceAttributeResponseBody {
	s.Items = v
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
	VpcId                 *string                                                              `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	CreationTime          *string                                                              `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	DBInstanceCpuCores    *int32                                                               `json:"DBInstanceCpuCores,omitempty" xml:"DBInstanceCpuCores,omitempty"`
	SegmentCounts         *int32                                                               `json:"SegmentCounts,omitempty" xml:"SegmentCounts,omitempty"`
	StoragePerNode        *int32                                                               `json:"StoragePerNode,omitempty" xml:"StoragePerNode,omitempty"`
	DBInstanceMemory      *int64                                                               `json:"DBInstanceMemory,omitempty" xml:"DBInstanceMemory,omitempty"`
	HostType              *string                                                              `json:"HostType,omitempty" xml:"HostType,omitempty"`
	PayType               *string                                                              `json:"PayType,omitempty" xml:"PayType,omitempty"`
	Tags                  *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	StorageType           *string                                                              `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	AvailabilityValue     *string                                                              `json:"AvailabilityValue,omitempty" xml:"AvailabilityValue,omitempty"`
	ReadDelayTime         *string                                                              `json:"ReadDelayTime,omitempty" xml:"ReadDelayTime,omitempty"`
	CpuCoresPerNode       *int32                                                               `json:"CpuCoresPerNode,omitempty" xml:"CpuCoresPerNode,omitempty"`
	Port                  *string                                                              `json:"Port,omitempty" xml:"Port,omitempty"`
	ConnectionMode        *string                                                              `json:"ConnectionMode,omitempty" xml:"ConnectionMode,omitempty"`
	LockMode              *string                                                              `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	EngineVersion         *string                                                              `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	StorageUnit           *string                                                              `json:"StorageUnit,omitempty" xml:"StorageUnit,omitempty"`
	MemoryPerNode         *int32                                                               `json:"MemoryPerNode,omitempty" xml:"MemoryPerNode,omitempty"`
	ConnectionString      *string                                                              `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	InstanceNetworkType   *string                                                              `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
	SecurityIPList        *string                                                              `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	MemoryUnit            *string                                                              `json:"MemoryUnit,omitempty" xml:"MemoryUnit,omitempty"`
	DBInstanceClassType   *string                                                              `json:"DBInstanceClassType,omitempty" xml:"DBInstanceClassType,omitempty"`
	DBInstanceDescription *string                                                              `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	DBInstanceGroupCount  *string                                                              `json:"DBInstanceGroupCount,omitempty" xml:"DBInstanceGroupCount,omitempty"`
	ExpireTime            *string                                                              `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	DBInstanceNetType     *string                                                              `json:"DBInstanceNetType,omitempty" xml:"DBInstanceNetType,omitempty"`
	MaintainStartTime     *string                                                              `json:"MaintainStartTime,omitempty" xml:"MaintainStartTime,omitempty"`
	MaintainEndTime       *string                                                              `json:"MaintainEndTime,omitempty" xml:"MaintainEndTime,omitempty"`
	LockReason            *string                                                              `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	DBInstanceStatus      *string                                                              `json:"DBInstanceStatus,omitempty" xml:"DBInstanceStatus,omitempty"`
	RegionId              *string                                                              `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DBInstanceDiskMBPS    *int64                                                               `json:"DBInstanceDiskMBPS,omitempty" xml:"DBInstanceDiskMBPS,omitempty"`
	DBInstanceStorage     *int64                                                               `json:"DBInstanceStorage,omitempty" xml:"DBInstanceStorage,omitempty"`
	ZoneId                *string                                                              `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	MaxConnections        *int32                                                               `json:"MaxConnections,omitempty" xml:"MaxConnections,omitempty"`
	DBInstanceId          *string                                                              `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	DBInstanceClass       *string                                                              `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty"`
	Engine                *string                                                              `json:"Engine,omitempty" xml:"Engine,omitempty"`
	DBInstanceCategory    *string                                                              `json:"DBInstanceCategory,omitempty" xml:"DBInstanceCategory,omitempty"`
}

func (s DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetVpcId(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.VpcId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetCreationTime(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.CreationTime = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetDBInstanceCpuCores(v int32) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.DBInstanceCpuCores = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetSegmentCounts(v int32) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.SegmentCounts = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetStoragePerNode(v int32) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.StoragePerNode = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetDBInstanceMemory(v int64) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.DBInstanceMemory = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetHostType(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.HostType = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetPayType(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.PayType = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetTags(v *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeTags) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.Tags = v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetStorageType(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.StorageType = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetAvailabilityValue(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.AvailabilityValue = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetReadDelayTime(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.ReadDelayTime = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetCpuCoresPerNode(v int32) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.CpuCoresPerNode = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetPort(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.Port = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetConnectionMode(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.ConnectionMode = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetLockMode(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.LockMode = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetEngineVersion(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.EngineVersion = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetStorageUnit(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.StorageUnit = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetMemoryPerNode(v int32) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.MemoryPerNode = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetConnectionString(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.ConnectionString = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetInstanceNetworkType(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.InstanceNetworkType = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetSecurityIPList(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.SecurityIPList = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetMemoryUnit(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.MemoryUnit = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetDBInstanceClassType(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.DBInstanceClassType = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetDBInstanceDescription(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.DBInstanceDescription = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetDBInstanceGroupCount(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.DBInstanceGroupCount = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetExpireTime(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetDBInstanceNetType(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.DBInstanceNetType = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetMaintainStartTime(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.MaintainStartTime = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetMaintainEndTime(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.MaintainEndTime = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetLockReason(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.LockReason = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetDBInstanceStatus(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.DBInstanceStatus = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetRegionId(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.RegionId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetDBInstanceDiskMBPS(v int64) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.DBInstanceDiskMBPS = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetDBInstanceStorage(v int64) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.DBInstanceStorage = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetZoneId(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.ZoneId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetMaxConnections(v int32) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.MaxConnections = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetDBInstanceId(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetDBInstanceClass(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.DBInstanceClass = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetEngine(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.Engine = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetDBInstanceCategory(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.DBInstanceCategory = &v
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
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDBInstanceAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeDBInstanceAttributeResponse) SetBody(v *DescribeDBInstanceAttributeResponseBody) *DescribeDBInstanceAttributeResponse {
	s.Body = v
	return s
}

type DescribeDBInstanceIPArrayListRequest struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
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

type DescribeDBInstanceIPArrayListResponseBody struct {
	RequestId *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Items     *DescribeDBInstanceIPArrayListResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
}

func (s DescribeDBInstanceIPArrayListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceIPArrayListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceIPArrayListResponseBody) SetRequestId(v string) *DescribeDBInstanceIPArrayListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstanceIPArrayListResponseBody) SetItems(v *DescribeDBInstanceIPArrayListResponseBodyItems) *DescribeDBInstanceIPArrayListResponseBody {
	s.Items = v
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
	DBInstanceIPArrayAttribute *string `json:"DBInstanceIPArrayAttribute,omitempty" xml:"DBInstanceIPArrayAttribute,omitempty"`
	DBInstanceIPArrayName      *string `json:"DBInstanceIPArrayName,omitempty" xml:"DBInstanceIPArrayName,omitempty"`
	SecurityIPList             *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
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
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDBInstanceIPArrayListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeDBInstanceIPArrayListResponse) SetBody(v *DescribeDBInstanceIPArrayListResponseBody) *DescribeDBInstanceIPArrayListResponse {
	s.Body = v
	return s
}

type DescribeDBInstanceNetInfoRequest struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s DescribeDBInstanceNetInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceNetInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceNetInfoRequest) SetDBInstanceId(v string) *DescribeDBInstanceNetInfoRequest {
	s.DBInstanceId = &v
	return s
}

type DescribeDBInstanceNetInfoResponseBody struct {
	RequestId           *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DBInstanceNetInfos  *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos `json:"DBInstanceNetInfos,omitempty" xml:"DBInstanceNetInfos,omitempty" type:"Struct"`
	InstanceNetworkType *string                                                  `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
}

func (s DescribeDBInstanceNetInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceNetInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceNetInfoResponseBody) SetRequestId(v string) *DescribeDBInstanceNetInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBody) SetDBInstanceNetInfos(v *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos) *DescribeDBInstanceNetInfoResponseBody {
	s.DBInstanceNetInfos = v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBody) SetInstanceNetworkType(v string) *DescribeDBInstanceNetInfoResponseBody {
	s.InstanceNetworkType = &v
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
	VSwitchId        *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	IPType           *string `json:"IPType,omitempty" xml:"IPType,omitempty"`
	Port             *string `json:"Port,omitempty" xml:"Port,omitempty"`
	VpcInstanceId    *string `json:"VpcInstanceId,omitempty" xml:"VpcInstanceId,omitempty"`
	VPCId            *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	IPAddress        *string `json:"IPAddress,omitempty" xml:"IPAddress,omitempty"`
	AddressType      *string `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
}

func (s DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo) SetVSwitchId(v string) *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo) SetConnectionString(v string) *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo {
	s.ConnectionString = &v
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

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo) SetVpcInstanceId(v string) *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo {
	s.VpcInstanceId = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo) SetVPCId(v string) *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo {
	s.VPCId = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo) SetIPAddress(v string) *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo {
	s.IPAddress = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo) SetAddressType(v string) *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo {
	s.AddressType = &v
	return s
}

type DescribeDBInstanceNetInfoResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDBInstanceNetInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeDBInstanceNetInfoResponse) SetBody(v *DescribeDBInstanceNetInfoResponseBody) *DescribeDBInstanceNetInfoResponse {
	s.Body = v
	return s
}

type DescribeDBInstanceOnECSAttributeRequest struct {
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s DescribeDBInstanceOnECSAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceOnECSAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceOnECSAttributeRequest) SetOwnerId(v int64) *DescribeDBInstanceOnECSAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeRequest) SetDBInstanceId(v string) *DescribeDBInstanceOnECSAttributeRequest {
	s.DBInstanceId = &v
	return s
}

type DescribeDBInstanceOnECSAttributeResponseBody struct {
	RequestId *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Items     *DescribeDBInstanceOnECSAttributeResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
}

func (s DescribeDBInstanceOnECSAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceOnECSAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceOnECSAttributeResponseBody) SetRequestId(v string) *DescribeDBInstanceOnECSAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponseBody) SetItems(v *DescribeDBInstanceOnECSAttributeResponseBodyItems) *DescribeDBInstanceOnECSAttributeResponseBody {
	s.Items = v
	return s
}

type DescribeDBInstanceOnECSAttributeResponseBodyItems struct {
	DBInstanceAttribute []*DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute `json:"DBInstanceAttribute,omitempty" xml:"DBInstanceAttribute,omitempty" type:"Repeated"`
}

func (s DescribeDBInstanceOnECSAttributeResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceOnECSAttributeResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceOnECSAttributeResponseBodyItems) SetDBInstanceAttribute(v []*DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute) *DescribeDBInstanceOnECSAttributeResponseBodyItems {
	s.DBInstanceAttribute = v
	return s
}

type DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute struct {
	CreationTime          *string                                                                   `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	VpcId                 *string                                                                   `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	EncryptionType        *string                                                                   `json:"EncryptionType,omitempty" xml:"EncryptionType,omitempty"`
	InstanceDeployType    *string                                                                   `json:"InstanceDeployType,omitempty" xml:"InstanceDeployType,omitempty"`
	PayType               *string                                                                   `json:"PayType,omitempty" xml:"PayType,omitempty"`
	Tags                  *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttributeTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	StorageType           *string                                                                   `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	ConnectionMode        *string                                                                   `json:"ConnectionMode,omitempty" xml:"ConnectionMode,omitempty"`
	Port                  *string                                                                   `json:"Port,omitempty" xml:"Port,omitempty"`
	LockMode              *string                                                                   `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	EngineVersion         *string                                                                   `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	MemorySize            *int32                                                                    `json:"MemorySize,omitempty" xml:"MemorySize,omitempty"`
	SegNodeNum            *int32                                                                    `json:"SegNodeNum,omitempty" xml:"SegNodeNum,omitempty"`
	ConnectionString      *string                                                                   `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	InstanceNetworkType   *string                                                                   `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
	EncryptionKey         *string                                                                   `json:"EncryptionKey,omitempty" xml:"EncryptionKey,omitempty"`
	DBInstanceDescription *string                                                                   `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	CpuCores              *int32                                                                    `json:"CpuCores,omitempty" xml:"CpuCores,omitempty"`
	ExpireTime            *string                                                                   `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	DBInstanceStatus      *string                                                                   `json:"DBInstanceStatus,omitempty" xml:"DBInstanceStatus,omitempty"`
	StorageSize           *int32                                                                    `json:"StorageSize,omitempty" xml:"StorageSize,omitempty"`
	RegionId              *string                                                                   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	VSwitchId             *string                                                                   `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	ZoneId                *string                                                                   `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	DBInstanceId          *string                                                                   `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	Engine                *string                                                                   `json:"Engine,omitempty" xml:"Engine,omitempty"`
	DBInstanceClass       *string                                                                   `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty"`
	SupportRestore        *bool                                                                     `json:"SupportRestore,omitempty" xml:"SupportRestore,omitempty"`
	MinorVersion          *string                                                                   `json:"MinorVersion,omitempty" xml:"MinorVersion,omitempty"`
	MasterNodeNum         *int32                                                                    `json:"MasterNodeNum,omitempty" xml:"MasterNodeNum,omitempty"`
	DBInstanceCategory    *string                                                                   `json:"DBInstanceCategory,omitempty" xml:"DBInstanceCategory,omitempty"`
}

func (s DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute) SetCreationTime(v string) *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute {
	s.CreationTime = &v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute) SetVpcId(v string) *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute {
	s.VpcId = &v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute) SetEncryptionType(v string) *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute {
	s.EncryptionType = &v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute) SetInstanceDeployType(v string) *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute {
	s.InstanceDeployType = &v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute) SetPayType(v string) *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute {
	s.PayType = &v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute) SetTags(v *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttributeTags) *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute {
	s.Tags = v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute) SetStorageType(v string) *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute {
	s.StorageType = &v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute) SetConnectionMode(v string) *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute {
	s.ConnectionMode = &v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute) SetPort(v string) *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute {
	s.Port = &v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute) SetLockMode(v string) *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute {
	s.LockMode = &v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute) SetEngineVersion(v string) *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute {
	s.EngineVersion = &v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute) SetMemorySize(v int32) *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute {
	s.MemorySize = &v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute) SetSegNodeNum(v int32) *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute {
	s.SegNodeNum = &v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute) SetConnectionString(v string) *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute {
	s.ConnectionString = &v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute) SetInstanceNetworkType(v string) *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute {
	s.InstanceNetworkType = &v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute) SetEncryptionKey(v string) *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute {
	s.EncryptionKey = &v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute) SetDBInstanceDescription(v string) *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute {
	s.DBInstanceDescription = &v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute) SetCpuCores(v int32) *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute {
	s.CpuCores = &v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute) SetExpireTime(v string) *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute) SetDBInstanceStatus(v string) *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute {
	s.DBInstanceStatus = &v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute) SetStorageSize(v int32) *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute {
	s.StorageSize = &v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute) SetRegionId(v string) *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute {
	s.RegionId = &v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute) SetVSwitchId(v string) *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute) SetZoneId(v string) *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute {
	s.ZoneId = &v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute) SetDBInstanceId(v string) *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute) SetEngine(v string) *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute {
	s.Engine = &v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute) SetDBInstanceClass(v string) *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute {
	s.DBInstanceClass = &v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute) SetSupportRestore(v bool) *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute {
	s.SupportRestore = &v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute) SetMinorVersion(v string) *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute {
	s.MinorVersion = &v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute) SetMasterNodeNum(v int32) *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute {
	s.MasterNodeNum = &v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute) SetDBInstanceCategory(v string) *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttribute {
	s.DBInstanceCategory = &v
	return s
}

type DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttributeTags struct {
	Tag []*DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttributeTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttributeTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttributeTags) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttributeTags) SetTag(v []*DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttributeTagsTag) *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttributeTags {
	s.Tag = v
	return s
}

type DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttributeTagsTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttributeTagsTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttributeTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttributeTagsTag) SetKey(v string) *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttributeTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttributeTagsTag) SetValue(v string) *DescribeDBInstanceOnECSAttributeResponseBodyItemsDBInstanceAttributeTagsTag {
	s.Value = &v
	return s
}

type DescribeDBInstanceOnECSAttributeResponse struct {
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDBInstanceOnECSAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDBInstanceOnECSAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceOnECSAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceOnECSAttributeResponse) SetHeaders(v map[string]*string) *DescribeDBInstanceOnECSAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponse) SetBody(v *DescribeDBInstanceOnECSAttributeResponseBody) *DescribeDBInstanceOnECSAttributeResponse {
	s.Body = v
	return s
}

type DescribeDBInstancePerformanceRequest struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	Key          *string `json:"Key,omitempty" xml:"Key,omitempty"`
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
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

func (s *DescribeDBInstancePerformanceRequest) SetKey(v string) *DescribeDBInstancePerformanceRequest {
	s.Key = &v
	return s
}

func (s *DescribeDBInstancePerformanceRequest) SetStartTime(v string) *DescribeDBInstancePerformanceRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDBInstancePerformanceRequest) SetEndTime(v string) *DescribeDBInstancePerformanceRequest {
	s.EndTime = &v
	return s
}

type DescribeDBInstancePerformanceResponseBody struct {
	EndTime         *string   `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime       *string   `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	DBInstanceId    *string   `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	Engine          *string   `json:"Engine,omitempty" xml:"Engine,omitempty"`
	RequestId       *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PerformanceKeys []*string `json:"PerformanceKeys,omitempty" xml:"PerformanceKeys,omitempty" type:"Repeated"`
}

func (s DescribeDBInstancePerformanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancePerformanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancePerformanceResponseBody) SetEndTime(v string) *DescribeDBInstancePerformanceResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDBInstancePerformanceResponseBody) SetStartTime(v string) *DescribeDBInstancePerformanceResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDBInstancePerformanceResponseBody) SetDBInstanceId(v string) *DescribeDBInstancePerformanceResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstancePerformanceResponseBody) SetEngine(v string) *DescribeDBInstancePerformanceResponseBody {
	s.Engine = &v
	return s
}

func (s *DescribeDBInstancePerformanceResponseBody) SetRequestId(v string) *DescribeDBInstancePerformanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstancePerformanceResponseBody) SetPerformanceKeys(v []*string) *DescribeDBInstancePerformanceResponseBody {
	s.PerformanceKeys = v
	return s
}

type DescribeDBInstancePerformanceResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDBInstancePerformanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeDBInstancePerformanceResponse) SetBody(v *DescribeDBInstancePerformanceResponseBody) *DescribeDBInstancePerformanceResponse {
	s.Body = v
	return s
}

type DescribeDBInstancesRequest struct {
	OwnerId               *int64                           `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId              *string                          `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DBInstanceDescription *string                          `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	InstanceNetworkType   *string                          `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
	DBInstanceIds         *string                          `json:"DBInstanceIds,omitempty" xml:"DBInstanceIds,omitempty"`
	PageSize              *int32                           `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber            *int32                           `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Tag                   []*DescribeDBInstancesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	InstanceDeployTypes   []*string                        `json:"InstanceDeployTypes,omitempty" xml:"InstanceDeployTypes,omitempty" type:"Repeated"`
	DBInstanceStatuses    []*string                        `json:"DBInstanceStatuses,omitempty" xml:"DBInstanceStatuses,omitempty" type:"Repeated"`
	DBInstanceCategories  []*string                        `json:"DBInstanceCategories,omitempty" xml:"DBInstanceCategories,omitempty" type:"Repeated"`
}

func (s DescribeDBInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesRequest) SetOwnerId(v int64) *DescribeDBInstancesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetRegionId(v string) *DescribeDBInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetDBInstanceDescription(v string) *DescribeDBInstancesRequest {
	s.DBInstanceDescription = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetInstanceNetworkType(v string) *DescribeDBInstancesRequest {
	s.InstanceNetworkType = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetDBInstanceIds(v string) *DescribeDBInstancesRequest {
	s.DBInstanceIds = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetPageSize(v int32) *DescribeDBInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetPageNumber(v int32) *DescribeDBInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetTag(v []*DescribeDBInstancesRequestTag) *DescribeDBInstancesRequest {
	s.Tag = v
	return s
}

func (s *DescribeDBInstancesRequest) SetInstanceDeployTypes(v []*string) *DescribeDBInstancesRequest {
	s.InstanceDeployTypes = v
	return s
}

func (s *DescribeDBInstancesRequest) SetDBInstanceStatuses(v []*string) *DescribeDBInstancesRequest {
	s.DBInstanceStatuses = v
	return s
}

func (s *DescribeDBInstancesRequest) SetDBInstanceCategories(v []*string) *DescribeDBInstancesRequest {
	s.DBInstanceCategories = v
	return s
}

type DescribeDBInstancesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
	OwnerId                    *int64                                 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId                   *string                                `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DBInstanceDescription      *string                                `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	InstanceNetworkType        *string                                `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
	DBInstanceIds              *string                                `json:"DBInstanceIds,omitempty" xml:"DBInstanceIds,omitempty"`
	PageSize                   *int32                                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber                 *int32                                 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Tag                        []*DescribeDBInstancesShrinkRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	InstanceDeployTypesShrink  *string                                `json:"InstanceDeployTypes,omitempty" xml:"InstanceDeployTypes,omitempty"`
	DBInstanceStatusesShrink   *string                                `json:"DBInstanceStatuses,omitempty" xml:"DBInstanceStatuses,omitempty"`
	DBInstanceCategoriesShrink *string                                `json:"DBInstanceCategories,omitempty" xml:"DBInstanceCategories,omitempty"`
}

func (s DescribeDBInstancesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancesShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesShrinkRequest) SetOwnerId(v int64) *DescribeDBInstancesShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBInstancesShrinkRequest) SetRegionId(v string) *DescribeDBInstancesShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBInstancesShrinkRequest) SetDBInstanceDescription(v string) *DescribeDBInstancesShrinkRequest {
	s.DBInstanceDescription = &v
	return s
}

func (s *DescribeDBInstancesShrinkRequest) SetInstanceNetworkType(v string) *DescribeDBInstancesShrinkRequest {
	s.InstanceNetworkType = &v
	return s
}

func (s *DescribeDBInstancesShrinkRequest) SetDBInstanceIds(v string) *DescribeDBInstancesShrinkRequest {
	s.DBInstanceIds = &v
	return s
}

func (s *DescribeDBInstancesShrinkRequest) SetPageSize(v int32) *DescribeDBInstancesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDBInstancesShrinkRequest) SetPageNumber(v int32) *DescribeDBInstancesShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBInstancesShrinkRequest) SetTag(v []*DescribeDBInstancesShrinkRequestTag) *DescribeDBInstancesShrinkRequest {
	s.Tag = v
	return s
}

func (s *DescribeDBInstancesShrinkRequest) SetInstanceDeployTypesShrink(v string) *DescribeDBInstancesShrinkRequest {
	s.InstanceDeployTypesShrink = &v
	return s
}

func (s *DescribeDBInstancesShrinkRequest) SetDBInstanceStatusesShrink(v string) *DescribeDBInstancesShrinkRequest {
	s.DBInstanceStatusesShrink = &v
	return s
}

func (s *DescribeDBInstancesShrinkRequest) SetDBInstanceCategoriesShrink(v string) *DescribeDBInstancesShrinkRequest {
	s.DBInstanceCategoriesShrink = &v
	return s
}

type DescribeDBInstancesShrinkRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
	TotalRecordCount *int32                                `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
	PageRecordCount  *int32                                `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	RequestId        *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber       *int32                                `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Items            *DescribeDBInstancesResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
}

func (s DescribeDBInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesResponseBody) SetTotalRecordCount(v int32) *DescribeDBInstancesResponseBody {
	s.TotalRecordCount = &v
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

func (s *DescribeDBInstancesResponseBody) SetPageNumber(v int32) *DescribeDBInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBInstancesResponseBody) SetItems(v *DescribeDBInstancesResponseBodyItems) *DescribeDBInstancesResponseBody {
	s.Items = v
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
	VpcId                 *string                                             `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	ExpireTime            *string                                             `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	DBInstanceNetType     *string                                             `json:"DBInstanceNetType,omitempty" xml:"DBInstanceNetType,omitempty"`
	InstanceDeployType    *string                                             `json:"InstanceDeployType,omitempty" xml:"InstanceDeployType,omitempty"`
	StorageType           *string                                             `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	CreateTime            *string                                             `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	PayType               *string                                             `json:"PayType,omitempty" xml:"PayType,omitempty"`
	Tags                  *DescribeDBInstancesResponseBodyItemsDBInstanceTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	LockReason            *string                                             `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	DBInstanceStatus      *string                                             `json:"DBInstanceStatus,omitempty" xml:"DBInstanceStatus,omitempty"`
	ConnectionMode        *string                                             `json:"ConnectionMode,omitempty" xml:"ConnectionMode,omitempty"`
	LockMode              *string                                             `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	EngineVersion         *string                                             `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	RegionId              *string                                             `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	VSwitchId             *string                                             `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	InstanceNetworkType   *string                                             `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
	ZoneId                *string                                             `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	DBInstanceId          *string                                             `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	Engine                *string                                             `json:"Engine,omitempty" xml:"Engine,omitempty"`
	DBInstanceDescription *string                                             `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	SegNodeNum            *string                                             `json:"SegNodeNum,omitempty" xml:"SegNodeNum,omitempty"`
	StorageSize           *string                                             `json:"StorageSize,omitempty" xml:"StorageSize,omitempty"`
	MasterNodeNum         *int32                                              `json:"MasterNodeNum,omitempty" xml:"MasterNodeNum,omitempty"`
	DBInstanceCategory    *string                                             `json:"DBInstanceCategory,omitempty" xml:"DBInstanceCategory,omitempty"`
}

func (s DescribeDBInstancesResponseBodyItemsDBInstance) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancesResponseBodyItemsDBInstance) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetVpcId(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.VpcId = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetExpireTime(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetDBInstanceNetType(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.DBInstanceNetType = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetInstanceDeployType(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.InstanceDeployType = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetStorageType(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.StorageType = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetCreateTime(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.CreateTime = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetPayType(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.PayType = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetTags(v *DescribeDBInstancesResponseBodyItemsDBInstanceTags) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.Tags = v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetLockReason(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.LockReason = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetDBInstanceStatus(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.DBInstanceStatus = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetConnectionMode(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.ConnectionMode = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetLockMode(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.LockMode = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetEngineVersion(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.EngineVersion = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetRegionId(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.RegionId = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetVSwitchId(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetInstanceNetworkType(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.InstanceNetworkType = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetZoneId(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.ZoneId = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetDBInstanceId(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetEngine(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.Engine = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetDBInstanceDescription(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.DBInstanceDescription = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetSegNodeNum(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.SegNodeNum = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetStorageSize(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.StorageSize = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetMasterNodeNum(v int32) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.MasterNodeNum = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsDBInstance) SetDBInstanceCategory(v string) *DescribeDBInstancesResponseBodyItemsDBInstance {
	s.DBInstanceCategory = &v
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
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDBInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeDBInstancesResponse) SetBody(v *DescribeDBInstancesResponseBody) *DescribeDBInstancesResponse {
	s.Body = v
	return s
}

type DescribeDBInstanceSQLPatternsRequest struct {
	DBInstanceId  *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	QueryKeywords *string `json:"QueryKeywords,omitempty" xml:"QueryKeywords,omitempty"`
	StartTime     *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Database      *string `json:"Database,omitempty" xml:"Database,omitempty"`
	User          *string `json:"User,omitempty" xml:"User,omitempty"`
	EndTime       *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	SourceIP      *string `json:"SourceIP,omitempty" xml:"SourceIP,omitempty"`
}

func (s DescribeDBInstanceSQLPatternsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceSQLPatternsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceSQLPatternsRequest) SetDBInstanceId(v string) *DescribeDBInstanceSQLPatternsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceSQLPatternsRequest) SetQueryKeywords(v string) *DescribeDBInstanceSQLPatternsRequest {
	s.QueryKeywords = &v
	return s
}

func (s *DescribeDBInstanceSQLPatternsRequest) SetStartTime(v string) *DescribeDBInstanceSQLPatternsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDBInstanceSQLPatternsRequest) SetDatabase(v string) *DescribeDBInstanceSQLPatternsRequest {
	s.Database = &v
	return s
}

func (s *DescribeDBInstanceSQLPatternsRequest) SetUser(v string) *DescribeDBInstanceSQLPatternsRequest {
	s.User = &v
	return s
}

func (s *DescribeDBInstanceSQLPatternsRequest) SetEndTime(v string) *DescribeDBInstanceSQLPatternsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDBInstanceSQLPatternsRequest) SetSourceIP(v string) *DescribeDBInstanceSQLPatternsRequest {
	s.SourceIP = &v
	return s
}

type DescribeDBInstanceSQLPatternsResponseBody struct {
	EndTime     *string                                              `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RequestId   *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime   *string                                              `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	DBClusterId *string                                              `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	Patterns    []*DescribeDBInstanceSQLPatternsResponseBodyPatterns `json:"Patterns,omitempty" xml:"Patterns,omitempty" type:"Repeated"`
}

func (s DescribeDBInstanceSQLPatternsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceSQLPatternsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceSQLPatternsResponseBody) SetEndTime(v string) *DescribeDBInstanceSQLPatternsResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDBInstanceSQLPatternsResponseBody) SetRequestId(v string) *DescribeDBInstanceSQLPatternsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstanceSQLPatternsResponseBody) SetStartTime(v string) *DescribeDBInstanceSQLPatternsResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDBInstanceSQLPatternsResponseBody) SetDBClusterId(v string) *DescribeDBInstanceSQLPatternsResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBInstanceSQLPatternsResponseBody) SetPatterns(v []*DescribeDBInstanceSQLPatternsResponseBodyPatterns) *DescribeDBInstanceSQLPatternsResponseBody {
	s.Patterns = v
	return s
}

type DescribeDBInstanceSQLPatternsResponseBodyPatterns struct {
	Name   *string                `json:"Name,omitempty" xml:"Name,omitempty"`
	Values map[string]interface{} `json:"Values,omitempty" xml:"Values,omitempty"`
}

func (s DescribeDBInstanceSQLPatternsResponseBodyPatterns) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceSQLPatternsResponseBodyPatterns) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceSQLPatternsResponseBodyPatterns) SetName(v string) *DescribeDBInstanceSQLPatternsResponseBodyPatterns {
	s.Name = &v
	return s
}

func (s *DescribeDBInstanceSQLPatternsResponseBodyPatterns) SetValues(v map[string]interface{}) *DescribeDBInstanceSQLPatternsResponseBodyPatterns {
	s.Values = v
	return s
}

type DescribeDBInstanceSQLPatternsResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDBInstanceSQLPatternsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDBInstanceSQLPatternsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceSQLPatternsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceSQLPatternsResponse) SetHeaders(v map[string]*string) *DescribeDBInstanceSQLPatternsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInstanceSQLPatternsResponse) SetBody(v *DescribeDBInstanceSQLPatternsResponseBody) *DescribeDBInstanceSQLPatternsResponse {
	s.Body = v
	return s
}

type DescribeDBInstanceSSLRequest struct {
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
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SSLEnabled     *bool   `json:"SSLEnabled,omitempty" xml:"SSLEnabled,omitempty"`
	DBInstanceId   *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	CertCommonName *string `json:"CertCommonName,omitempty" xml:"CertCommonName,omitempty"`
	SSLExpiredTime *string `json:"SSLExpiredTime,omitempty" xml:"SSLExpiredTime,omitempty"`
}

func (s DescribeDBInstanceSSLResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceSSLResponseBody) GoString() string {
	return s.String()
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

func (s *DescribeDBInstanceSSLResponseBody) SetDBInstanceId(v string) *DescribeDBInstanceSSLResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceSSLResponseBody) SetCertCommonName(v string) *DescribeDBInstanceSSLResponseBody {
	s.CertCommonName = &v
	return s
}

func (s *DescribeDBInstanceSSLResponseBody) SetSSLExpiredTime(v string) *DescribeDBInstanceSSLResponseBody {
	s.SSLExpiredTime = &v
	return s
}

type DescribeDBInstanceSSLResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDBInstanceSSLResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeDBInstanceSSLResponse) SetBody(v *DescribeDBInstanceSSLResponseBody) *DescribeDBInstanceSSLResponse {
	s.Body = v
	return s
}

type DescribeLogBackupsRequest struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber   *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
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

func (s *DescribeLogBackupsRequest) SetStartTime(v string) *DescribeLogBackupsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeLogBackupsRequest) SetEndTime(v string) *DescribeLogBackupsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeLogBackupsRequest) SetPageSize(v int32) *DescribeLogBackupsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeLogBackupsRequest) SetPageNumber(v int32) *DescribeLogBackupsRequest {
	s.PageNumber = &v
	return s
}

type DescribeLogBackupsResponseBody struct {
	TotalLogSize *int64                                 `json:"TotalLogSize,omitempty" xml:"TotalLogSize,omitempty"`
	PageSize     *int32                                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber   *int32                                 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	RequestId    *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount   *int32                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Items        []*DescribeLogBackupsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
}

func (s DescribeLogBackupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogBackupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLogBackupsResponseBody) SetTotalLogSize(v int64) *DescribeLogBackupsResponseBody {
	s.TotalLogSize = &v
	return s
}

func (s *DescribeLogBackupsResponseBody) SetPageSize(v int32) *DescribeLogBackupsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeLogBackupsResponseBody) SetPageNumber(v int32) *DescribeLogBackupsResponseBody {
	s.PageNumber = &v
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

func (s *DescribeLogBackupsResponseBody) SetItems(v []*DescribeLogBackupsResponseBodyItems) *DescribeLogBackupsResponseBody {
	s.Items = v
	return s
}

type DescribeLogBackupsResponseBodyItems struct {
	LogFileSize  *int64  `json:"LogFileSize,omitempty" xml:"LogFileSize,omitempty"`
	LogTime      *string `json:"LogTime,omitempty" xml:"LogTime,omitempty"`
	SegmentName  *string `json:"SegmentName,omitempty" xml:"SegmentName,omitempty"`
	LogFileName  *string `json:"LogFileName,omitempty" xml:"LogFileName,omitempty"`
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	BackupId     *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
}

func (s DescribeLogBackupsResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogBackupsResponseBodyItems) GoString() string {
	return s.String()
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

func (s *DescribeLogBackupsResponseBodyItems) SetLogFileName(v string) *DescribeLogBackupsResponseBodyItems {
	s.LogFileName = &v
	return s
}

func (s *DescribeLogBackupsResponseBodyItems) SetDBInstanceId(v string) *DescribeLogBackupsResponseBodyItems {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeLogBackupsResponseBodyItems) SetBackupId(v string) *DescribeLogBackupsResponseBodyItems {
	s.BackupId = &v
	return s
}

type DescribeLogBackupsResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeLogBackupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeLogBackupsResponse) SetBody(v *DescribeLogBackupsResponseBody) *DescribeLogBackupsResponse {
	s.Body = v
	return s
}

type DescribeModifyParameterLogRequest struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
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

func (s *DescribeModifyParameterLogRequest) SetStartTime(v string) *DescribeModifyParameterLogRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeModifyParameterLogRequest) SetEndTime(v string) *DescribeModifyParameterLogRequest {
	s.EndTime = &v
	return s
}

type DescribeModifyParameterLogResponseBody struct {
	RequestId  *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Changelogs []*DescribeModifyParameterLogResponseBodyChangelogs `json:"Changelogs,omitempty" xml:"Changelogs,omitempty" type:"Repeated"`
}

func (s DescribeModifyParameterLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeModifyParameterLogResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeModifyParameterLogResponseBody) SetRequestId(v string) *DescribeModifyParameterLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeModifyParameterLogResponseBody) SetChangelogs(v []*DescribeModifyParameterLogResponseBodyChangelogs) *DescribeModifyParameterLogResponseBody {
	s.Changelogs = v
	return s
}

type DescribeModifyParameterLogResponseBodyChangelogs struct {
	ParameterValueAfter  *string `json:"ParameterValueAfter,omitempty" xml:"ParameterValueAfter,omitempty"`
	ParameterValueBefore *string `json:"ParameterValueBefore,omitempty" xml:"ParameterValueBefore,omitempty"`
	ParameterName        *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	ParameterValid       *string `json:"ParameterValid,omitempty" xml:"ParameterValid,omitempty"`
	EffectTime           *string `json:"EffectTime,omitempty" xml:"EffectTime,omitempty"`
}

func (s DescribeModifyParameterLogResponseBodyChangelogs) String() string {
	return tea.Prettify(s)
}

func (s DescribeModifyParameterLogResponseBodyChangelogs) GoString() string {
	return s.String()
}

func (s *DescribeModifyParameterLogResponseBodyChangelogs) SetParameterValueAfter(v string) *DescribeModifyParameterLogResponseBodyChangelogs {
	s.ParameterValueAfter = &v
	return s
}

func (s *DescribeModifyParameterLogResponseBodyChangelogs) SetParameterValueBefore(v string) *DescribeModifyParameterLogResponseBodyChangelogs {
	s.ParameterValueBefore = &v
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

func (s *DescribeModifyParameterLogResponseBodyChangelogs) SetEffectTime(v string) *DescribeModifyParameterLogResponseBodyChangelogs {
	s.EffectTime = &v
	return s
}

type DescribeModifyParameterLogResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeModifyParameterLogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeModifyParameterLogResponse) SetBody(v *DescribeModifyParameterLogResponseBody) *DescribeModifyParameterLogResponse {
	s.Body = v
	return s
}

type DescribeParametersRequest struct {
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
	RequestId  *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Parameters []*DescribeParametersResponseBodyParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
}

func (s DescribeParametersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeParametersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeParametersResponseBody) SetRequestId(v string) *DescribeParametersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeParametersResponseBody) SetParameters(v []*DescribeParametersResponseBodyParameters) *DescribeParametersResponseBody {
	s.Parameters = v
	return s
}

type DescribeParametersResponseBodyParameters struct {
	IsChangeableConfig   *string `json:"IsChangeableConfig,omitempty" xml:"IsChangeableConfig,omitempty"`
	ForceRestartInstance *string `json:"ForceRestartInstance,omitempty" xml:"ForceRestartInstance,omitempty"`
	OptionalRange        *string `json:"OptionalRange,omitempty" xml:"OptionalRange,omitempty"`
	ParameterName        *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	ParameterValue       *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
	ParameterDescription *string `json:"ParameterDescription,omitempty" xml:"ParameterDescription,omitempty"`
	CurrentValue         *string `json:"CurrentValue,omitempty" xml:"CurrentValue,omitempty"`
}

func (s DescribeParametersResponseBodyParameters) String() string {
	return tea.Prettify(s)
}

func (s DescribeParametersResponseBodyParameters) GoString() string {
	return s.String()
}

func (s *DescribeParametersResponseBodyParameters) SetIsChangeableConfig(v string) *DescribeParametersResponseBodyParameters {
	s.IsChangeableConfig = &v
	return s
}

func (s *DescribeParametersResponseBodyParameters) SetForceRestartInstance(v string) *DescribeParametersResponseBodyParameters {
	s.ForceRestartInstance = &v
	return s
}

func (s *DescribeParametersResponseBodyParameters) SetOptionalRange(v string) *DescribeParametersResponseBodyParameters {
	s.OptionalRange = &v
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

func (s *DescribeParametersResponseBodyParameters) SetParameterDescription(v string) *DescribeParametersResponseBodyParameters {
	s.ParameterDescription = &v
	return s
}

func (s *DescribeParametersResponseBodyParameters) SetCurrentValue(v string) *DescribeParametersResponseBodyParameters {
	s.CurrentValue = &v
	return s
}

type DescribeParametersResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeParametersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeParametersResponse) SetBody(v *DescribeParametersResponseBody) *DescribeParametersResponse {
	s.Body = v
	return s
}

type DescribeRdsVpcsRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeRdsVpcsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRdsVpcsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRdsVpcsRequest) SetSecurityToken(v string) *DescribeRdsVpcsRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeRdsVpcsRequest) SetOwnerId(v int64) *DescribeRdsVpcsRequest {
	s.OwnerId = &v
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

func (s *DescribeRdsVpcsRequest) SetOwnerAccount(v string) *DescribeRdsVpcsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeRdsVpcsRequest) SetRegionId(v string) *DescribeRdsVpcsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRdsVpcsRequest) SetZoneId(v string) *DescribeRdsVpcsRequest {
	s.ZoneId = &v
	return s
}

type DescribeRdsVpcsResponseBody struct {
	Vpcs      *DescribeRdsVpcsResponseBodyVpcs `json:"Vpcs,omitempty" xml:"Vpcs,omitempty" type:"Struct"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRdsVpcsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRdsVpcsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRdsVpcsResponseBody) SetVpcs(v *DescribeRdsVpcsResponseBodyVpcs) *DescribeRdsVpcsResponseBody {
	s.Vpcs = v
	return s
}

func (s *DescribeRdsVpcsResponseBody) SetRequestId(v string) *DescribeRdsVpcsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeRdsVpcsResponseBodyVpcs struct {
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
	Status      *string                                       `json:"Status,omitempty" xml:"Status,omitempty"`
	VpcName     *string                                       `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
	VpcId       *string                                       `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	IsDefault   *bool                                         `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	CidrBlock   *string                                       `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	RegionNo    *string                                       `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	GmtCreate   *string                                       `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	AliUid      *string                                       `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	VSwitchs    []*DescribeRdsVpcsResponseBodyVpcsVpcVSwitchs `json:"VSwitchs,omitempty" xml:"VSwitchs,omitempty" type:"Repeated"`
	GmtModified *string                                       `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Bid         *string                                       `json:"Bid,omitempty" xml:"Bid,omitempty"`
}

func (s DescribeRdsVpcsResponseBodyVpcsVpc) String() string {
	return tea.Prettify(s)
}

func (s DescribeRdsVpcsResponseBodyVpcsVpc) GoString() string {
	return s.String()
}

func (s *DescribeRdsVpcsResponseBodyVpcsVpc) SetStatus(v string) *DescribeRdsVpcsResponseBodyVpcsVpc {
	s.Status = &v
	return s
}

func (s *DescribeRdsVpcsResponseBodyVpcsVpc) SetVpcName(v string) *DescribeRdsVpcsResponseBodyVpcsVpc {
	s.VpcName = &v
	return s
}

func (s *DescribeRdsVpcsResponseBodyVpcsVpc) SetVpcId(v string) *DescribeRdsVpcsResponseBodyVpcsVpc {
	s.VpcId = &v
	return s
}

func (s *DescribeRdsVpcsResponseBodyVpcsVpc) SetIsDefault(v bool) *DescribeRdsVpcsResponseBodyVpcsVpc {
	s.IsDefault = &v
	return s
}

func (s *DescribeRdsVpcsResponseBodyVpcsVpc) SetCidrBlock(v string) *DescribeRdsVpcsResponseBodyVpcsVpc {
	s.CidrBlock = &v
	return s
}

func (s *DescribeRdsVpcsResponseBodyVpcsVpc) SetRegionNo(v string) *DescribeRdsVpcsResponseBodyVpcsVpc {
	s.RegionNo = &v
	return s
}

func (s *DescribeRdsVpcsResponseBodyVpcsVpc) SetGmtCreate(v string) *DescribeRdsVpcsResponseBodyVpcsVpc {
	s.GmtCreate = &v
	return s
}

func (s *DescribeRdsVpcsResponseBodyVpcsVpc) SetAliUid(v string) *DescribeRdsVpcsResponseBodyVpcsVpc {
	s.AliUid = &v
	return s
}

func (s *DescribeRdsVpcsResponseBodyVpcsVpc) SetVSwitchs(v []*DescribeRdsVpcsResponseBodyVpcsVpcVSwitchs) *DescribeRdsVpcsResponseBodyVpcsVpc {
	s.VSwitchs = v
	return s
}

func (s *DescribeRdsVpcsResponseBodyVpcsVpc) SetGmtModified(v string) *DescribeRdsVpcsResponseBodyVpcsVpc {
	s.GmtModified = &v
	return s
}

func (s *DescribeRdsVpcsResponseBodyVpcsVpc) SetBid(v string) *DescribeRdsVpcsResponseBodyVpcsVpc {
	s.Bid = &v
	return s
}

type DescribeRdsVpcsResponseBodyVpcsVpcVSwitchs struct {
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
	VSwitchId   *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	IsDefault   *bool   `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	CidrBlock   *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	GmtCreate   *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	IzNo        *string `json:"IzNo,omitempty" xml:"IzNo,omitempty"`
	VSwitchName *string `json:"VSwitchName,omitempty" xml:"VSwitchName,omitempty"`
}

func (s DescribeRdsVpcsResponseBodyVpcsVpcVSwitchs) String() string {
	return tea.Prettify(s)
}

func (s DescribeRdsVpcsResponseBodyVpcsVpcVSwitchs) GoString() string {
	return s.String()
}

func (s *DescribeRdsVpcsResponseBodyVpcsVpcVSwitchs) SetStatus(v string) *DescribeRdsVpcsResponseBodyVpcsVpcVSwitchs {
	s.Status = &v
	return s
}

func (s *DescribeRdsVpcsResponseBodyVpcsVpcVSwitchs) SetVSwitchId(v string) *DescribeRdsVpcsResponseBodyVpcsVpcVSwitchs {
	s.VSwitchId = &v
	return s
}

func (s *DescribeRdsVpcsResponseBodyVpcsVpcVSwitchs) SetIsDefault(v bool) *DescribeRdsVpcsResponseBodyVpcsVpcVSwitchs {
	s.IsDefault = &v
	return s
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

func (s *DescribeRdsVpcsResponseBodyVpcsVpcVSwitchs) SetIzNo(v string) *DescribeRdsVpcsResponseBodyVpcsVpcVSwitchs {
	s.IzNo = &v
	return s
}

func (s *DescribeRdsVpcsResponseBodyVpcsVpcVSwitchs) SetVSwitchName(v string) *DescribeRdsVpcsResponseBodyVpcsVpcVSwitchs {
	s.VSwitchName = &v
	return s
}

type DescribeRdsVpcsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRdsVpcsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeRdsVpcsResponse) SetBody(v *DescribeRdsVpcsResponseBody) *DescribeRdsVpcsResponse {
	s.Body = v
	return s
}

type DescribeRdsVSwitchsRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	VpcId                *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeRdsVSwitchsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRdsVSwitchsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRdsVSwitchsRequest) SetSecurityToken(v string) *DescribeRdsVSwitchsRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeRdsVSwitchsRequest) SetOwnerId(v int64) *DescribeRdsVSwitchsRequest {
	s.OwnerId = &v
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

func (s *DescribeRdsVSwitchsRequest) SetOwnerAccount(v string) *DescribeRdsVSwitchsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeRdsVSwitchsRequest) SetRegionId(v string) *DescribeRdsVSwitchsRequest {
	s.RegionId = &v
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
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
	IsDefault   *bool   `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	VSwitchId   *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	CidrBlock   *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	RegionNo    *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	GmtCreate   *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	AliUid      *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Bid         *string `json:"Bid,omitempty" xml:"Bid,omitempty"`
	IzNo        *string `json:"IzNo,omitempty" xml:"IzNo,omitempty"`
	VSwitchName *string `json:"VSwitchName,omitempty" xml:"VSwitchName,omitempty"`
}

func (s DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch) String() string {
	return tea.Prettify(s)
}

func (s DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch) GoString() string {
	return s.String()
}

func (s *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch) SetStatus(v string) *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch {
	s.Status = &v
	return s
}

func (s *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch) SetIsDefault(v bool) *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch {
	s.IsDefault = &v
	return s
}

func (s *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch) SetVSwitchId(v string) *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch {
	s.VSwitchId = &v
	return s
}

func (s *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch) SetCidrBlock(v string) *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch {
	s.CidrBlock = &v
	return s
}

func (s *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch) SetRegionNo(v string) *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch {
	s.RegionNo = &v
	return s
}

func (s *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch) SetGmtCreate(v string) *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch {
	s.GmtCreate = &v
	return s
}

func (s *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch) SetAliUid(v string) *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch {
	s.AliUid = &v
	return s
}

func (s *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch) SetGmtModified(v string) *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch {
	s.GmtModified = &v
	return s
}

func (s *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch) SetBid(v string) *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch {
	s.Bid = &v
	return s
}

func (s *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch) SetIzNo(v string) *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch {
	s.IzNo = &v
	return s
}

func (s *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch) SetVSwitchName(v string) *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch {
	s.VSwitchName = &v
	return s
}

type DescribeRdsVSwitchsResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRdsVSwitchsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeRdsVSwitchsResponse) SetBody(v *DescribeRdsVSwitchsResponseBody) *DescribeRdsVSwitchsResponse {
	s.Body = v
	return s
}

type DescribeRegionsRequest struct {
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
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Regions   *DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Struct"`
}

func (s DescribeRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetRegions(v *DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
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
	RegionId *string                                        `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Zones    *DescribeRegionsResponseBodyRegionsRegionZones `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Struct"`
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
	ZoneId     *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	VpcEnabled *bool   `json:"VpcEnabled,omitempty" xml:"VpcEnabled,omitempty"`
}

func (s DescribeRegionsResponseBodyRegionsRegionZonesZone) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegionsRegionZonesZone) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegionsRegionZonesZone) SetZoneId(v string) *DescribeRegionsResponseBodyRegionsRegionZonesZone {
	s.ZoneId = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegionZonesZone) SetVpcEnabled(v bool) *DescribeRegionsResponseBodyRegionsRegionZonesZone {
	s.VpcEnabled = &v
	return s
}

type DescribeRegionsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeRegionsResponse) SetBody(v *DescribeRegionsResponseBody) *DescribeRegionsResponse {
	s.Body = v
	return s
}

type DescribeResourceUsageRequest struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s DescribeResourceUsageRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourceUsageRequest) GoString() string {
	return s.String()
}

func (s *DescribeResourceUsageRequest) SetDBInstanceId(v string) *DescribeResourceUsageRequest {
	s.DBInstanceId = &v
	return s
}

type DescribeResourceUsageResponseBody struct {
	DiskUsed     *int64  `json:"DiskUsed,omitempty" xml:"DiskUsed,omitempty"`
	DataSize     *int64  `json:"DataSize,omitempty" xml:"DataSize,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	BackupSize   *int64  `json:"BackupSize,omitempty" xml:"BackupSize,omitempty"`
	LogSize      *int64  `json:"LogSize,omitempty" xml:"LogSize,omitempty"`
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	Engine       *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
}

func (s DescribeResourceUsageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourceUsageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeResourceUsageResponseBody) SetDiskUsed(v int64) *DescribeResourceUsageResponseBody {
	s.DiskUsed = &v
	return s
}

func (s *DescribeResourceUsageResponseBody) SetDataSize(v int64) *DescribeResourceUsageResponseBody {
	s.DataSize = &v
	return s
}

func (s *DescribeResourceUsageResponseBody) SetRequestId(v string) *DescribeResourceUsageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeResourceUsageResponseBody) SetBackupSize(v int64) *DescribeResourceUsageResponseBody {
	s.BackupSize = &v
	return s
}

func (s *DescribeResourceUsageResponseBody) SetLogSize(v int64) *DescribeResourceUsageResponseBody {
	s.LogSize = &v
	return s
}

func (s *DescribeResourceUsageResponseBody) SetDBInstanceId(v string) *DescribeResourceUsageResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeResourceUsageResponseBody) SetEngine(v string) *DescribeResourceUsageResponseBody {
	s.Engine = &v
	return s
}

type DescribeResourceUsageResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeResourceUsageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeResourceUsageResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourceUsageResponse) GoString() string {
	return s.String()
}

func (s *DescribeResourceUsageResponse) SetHeaders(v map[string]*string) *DescribeResourceUsageResponse {
	s.Headers = v
	return s
}

func (s *DescribeResourceUsageResponse) SetBody(v *DescribeResourceUsageResponseBody) *DescribeResourceUsageResponse {
	s.Body = v
	return s
}

type DescribeSlowLogRecordsRequest struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	SQLId        *int64  `json:"SQLId,omitempty" xml:"SQLId,omitempty"`
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	DBName       *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber   *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s DescribeSlowLogRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowLogRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogRecordsRequest) SetDBInstanceId(v string) *DescribeSlowLogRecordsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetSQLId(v int64) *DescribeSlowLogRecordsRequest {
	s.SQLId = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetStartTime(v string) *DescribeSlowLogRecordsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetEndTime(v string) *DescribeSlowLogRecordsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetDBName(v string) *DescribeSlowLogRecordsRequest {
	s.DBName = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetPageSize(v int32) *DescribeSlowLogRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetPageNumber(v int32) *DescribeSlowLogRecordsRequest {
	s.PageNumber = &v
	return s
}

type DescribeSlowLogRecordsResponseBody struct {
	PageNumber       *int32                                   `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Engine           *string                                  `json:"Engine,omitempty" xml:"Engine,omitempty"`
	RequestId        *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageRecordCount  *int32                                   `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	TotalRecordCount *int32                                   `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
	Items            *DescribeSlowLogRecordsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
}

func (s DescribeSlowLogRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowLogRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogRecordsResponseBody) SetPageNumber(v int32) *DescribeSlowLogRecordsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBody) SetEngine(v string) *DescribeSlowLogRecordsResponseBody {
	s.Engine = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBody) SetRequestId(v string) *DescribeSlowLogRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBody) SetPageRecordCount(v int32) *DescribeSlowLogRecordsResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBody) SetTotalRecordCount(v int32) *DescribeSlowLogRecordsResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBody) SetItems(v *DescribeSlowLogRecordsResponseBodyItems) *DescribeSlowLogRecordsResponseBody {
	s.Items = v
	return s
}

type DescribeSlowLogRecordsResponseBodyItems struct {
	SQLSlowRecord []*DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord `json:"SQLSlowRecord,omitempty" xml:"SQLSlowRecord,omitempty" type:"Repeated"`
}

func (s DescribeSlowLogRecordsResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowLogRecordsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogRecordsResponseBodyItems) SetSQLSlowRecord(v []*DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord) *DescribeSlowLogRecordsResponseBodyItems {
	s.SQLSlowRecord = v
	return s
}

type DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord struct {
	ExecutionStartTime *string `json:"ExecutionStartTime,omitempty" xml:"ExecutionStartTime,omitempty"`
	HostAddress        *string `json:"HostAddress,omitempty" xml:"HostAddress,omitempty"`
	QueryTimes         *int64  `json:"QueryTimes,omitempty" xml:"QueryTimes,omitempty"`
	SQLText            *string `json:"SQLText,omitempty" xml:"SQLText,omitempty"`
	ReturnRowCounts    *int64  `json:"ReturnRowCounts,omitempty" xml:"ReturnRowCounts,omitempty"`
	ParseRowCounts     *int64  `json:"ParseRowCounts,omitempty" xml:"ParseRowCounts,omitempty"`
	DBName             *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	LockTimes          *int64  `json:"LockTimes,omitempty" xml:"LockTimes,omitempty"`
}

func (s DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord) SetExecutionStartTime(v string) *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord {
	s.ExecutionStartTime = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord) SetHostAddress(v string) *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord {
	s.HostAddress = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord) SetQueryTimes(v int64) *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord {
	s.QueryTimes = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord) SetSQLText(v string) *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord {
	s.SQLText = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord) SetReturnRowCounts(v int64) *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord {
	s.ReturnRowCounts = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord) SetParseRowCounts(v int64) *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord {
	s.ParseRowCounts = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord) SetDBName(v string) *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord {
	s.DBName = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord) SetLockTimes(v int64) *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord {
	s.LockTimes = &v
	return s
}

type DescribeSlowLogRecordsResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSlowLogRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSlowLogRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowLogRecordsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogRecordsResponse) SetHeaders(v map[string]*string) *DescribeSlowLogRecordsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSlowLogRecordsResponse) SetBody(v *DescribeSlowLogRecordsResponseBody) *DescribeSlowLogRecordsResponse {
	s.Body = v
	return s
}

type DescribeSlowSQLLogsRequest struct {
	DBInstanceId   *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	QueryKeywords  *string `json:"QueryKeywords,omitempty" xml:"QueryKeywords,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Database       *string `json:"Database,omitempty" xml:"Database,omitempty"`
	User           *string `json:"User,omitempty" xml:"User,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber     *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	SourceIP       *string `json:"SourceIP,omitempty" xml:"SourceIP,omitempty"`
	ExecuteState   *string `json:"ExecuteState,omitempty" xml:"ExecuteState,omitempty"`
	OperationClass *string `json:"OperationClass,omitempty" xml:"OperationClass,omitempty"`
	OperationType  *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	MinExecuteCost *string `json:"MinExecuteCost,omitempty" xml:"MinExecuteCost,omitempty"`
	MaxExecuteCost *string `json:"MaxExecuteCost,omitempty" xml:"MaxExecuteCost,omitempty"`
}

func (s DescribeSlowSQLLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowSQLLogsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSlowSQLLogsRequest) SetDBInstanceId(v string) *DescribeSlowSQLLogsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeSlowSQLLogsRequest) SetQueryKeywords(v string) *DescribeSlowSQLLogsRequest {
	s.QueryKeywords = &v
	return s
}

func (s *DescribeSlowSQLLogsRequest) SetStartTime(v string) *DescribeSlowSQLLogsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSlowSQLLogsRequest) SetDatabase(v string) *DescribeSlowSQLLogsRequest {
	s.Database = &v
	return s
}

func (s *DescribeSlowSQLLogsRequest) SetUser(v string) *DescribeSlowSQLLogsRequest {
	s.User = &v
	return s
}

func (s *DescribeSlowSQLLogsRequest) SetEndTime(v string) *DescribeSlowSQLLogsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSlowSQLLogsRequest) SetPageSize(v int32) *DescribeSlowSQLLogsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSlowSQLLogsRequest) SetPageNumber(v int32) *DescribeSlowSQLLogsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSlowSQLLogsRequest) SetSourceIP(v string) *DescribeSlowSQLLogsRequest {
	s.SourceIP = &v
	return s
}

func (s *DescribeSlowSQLLogsRequest) SetExecuteState(v string) *DescribeSlowSQLLogsRequest {
	s.ExecuteState = &v
	return s
}

func (s *DescribeSlowSQLLogsRequest) SetOperationClass(v string) *DescribeSlowSQLLogsRequest {
	s.OperationClass = &v
	return s
}

func (s *DescribeSlowSQLLogsRequest) SetOperationType(v string) *DescribeSlowSQLLogsRequest {
	s.OperationType = &v
	return s
}

func (s *DescribeSlowSQLLogsRequest) SetMinExecuteCost(v string) *DescribeSlowSQLLogsRequest {
	s.MinExecuteCost = &v
	return s
}

func (s *DescribeSlowSQLLogsRequest) SetMaxExecuteCost(v string) *DescribeSlowSQLLogsRequest {
	s.MaxExecuteCost = &v
	return s
}

type DescribeSlowSQLLogsResponseBody struct {
	PageRecordCount *int32                                  `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	RequestId       *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber      *int32                                  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Items           []*DescribeSlowSQLLogsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
}

func (s DescribeSlowSQLLogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowSQLLogsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSlowSQLLogsResponseBody) SetPageRecordCount(v int32) *DescribeSlowSQLLogsResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeSlowSQLLogsResponseBody) SetRequestId(v string) *DescribeSlowSQLLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSlowSQLLogsResponseBody) SetPageNumber(v int32) *DescribeSlowSQLLogsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSlowSQLLogsResponseBody) SetItems(v []*DescribeSlowSQLLogsResponseBodyItems) *DescribeSlowSQLLogsResponseBody {
	s.Items = v
	return s
}

type DescribeSlowSQLLogsResponseBodyItems struct {
	OperationClass       *string  `json:"OperationClass,omitempty" xml:"OperationClass,omitempty"`
	ExecuteState         *string  `json:"ExecuteState,omitempty" xml:"ExecuteState,omitempty"`
	ExecuteCost          *float32 `json:"ExecuteCost,omitempty" xml:"ExecuteCost,omitempty"`
	SQLText              *string  `json:"SQLText,omitempty" xml:"SQLText,omitempty"`
	SourcePort           *int32   `json:"SourcePort,omitempty" xml:"SourcePort,omitempty"`
	DBRole               *string  `json:"DBRole,omitempty" xml:"DBRole,omitempty"`
	OperationType        *string  `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	SourceIP             *string  `json:"SourceIP,omitempty" xml:"SourceIP,omitempty"`
	SQLPlan              *string  `json:"SQLPlan,omitempty" xml:"SQLPlan,omitempty"`
	ReturnRowCounts      *int64   `json:"ReturnRowCounts,omitempty" xml:"ReturnRowCounts,omitempty"`
	DBName               *string  `json:"DBName,omitempty" xml:"DBName,omitempty"`
	OperationExecuteTime *string  `json:"OperationExecuteTime,omitempty" xml:"OperationExecuteTime,omitempty"`
	ScanRowCounts        *int64   `json:"ScanRowCounts,omitempty" xml:"ScanRowCounts,omitempty"`
	AccountName          *string  `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	QueryId              *string  `json:"QueryId,omitempty" xml:"QueryId,omitempty"`
}

func (s DescribeSlowSQLLogsResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowSQLLogsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeSlowSQLLogsResponseBodyItems) SetOperationClass(v string) *DescribeSlowSQLLogsResponseBodyItems {
	s.OperationClass = &v
	return s
}

func (s *DescribeSlowSQLLogsResponseBodyItems) SetExecuteState(v string) *DescribeSlowSQLLogsResponseBodyItems {
	s.ExecuteState = &v
	return s
}

func (s *DescribeSlowSQLLogsResponseBodyItems) SetExecuteCost(v float32) *DescribeSlowSQLLogsResponseBodyItems {
	s.ExecuteCost = &v
	return s
}

func (s *DescribeSlowSQLLogsResponseBodyItems) SetSQLText(v string) *DescribeSlowSQLLogsResponseBodyItems {
	s.SQLText = &v
	return s
}

func (s *DescribeSlowSQLLogsResponseBodyItems) SetSourcePort(v int32) *DescribeSlowSQLLogsResponseBodyItems {
	s.SourcePort = &v
	return s
}

func (s *DescribeSlowSQLLogsResponseBodyItems) SetDBRole(v string) *DescribeSlowSQLLogsResponseBodyItems {
	s.DBRole = &v
	return s
}

func (s *DescribeSlowSQLLogsResponseBodyItems) SetOperationType(v string) *DescribeSlowSQLLogsResponseBodyItems {
	s.OperationType = &v
	return s
}

func (s *DescribeSlowSQLLogsResponseBodyItems) SetSourceIP(v string) *DescribeSlowSQLLogsResponseBodyItems {
	s.SourceIP = &v
	return s
}

func (s *DescribeSlowSQLLogsResponseBodyItems) SetSQLPlan(v string) *DescribeSlowSQLLogsResponseBodyItems {
	s.SQLPlan = &v
	return s
}

func (s *DescribeSlowSQLLogsResponseBodyItems) SetReturnRowCounts(v int64) *DescribeSlowSQLLogsResponseBodyItems {
	s.ReturnRowCounts = &v
	return s
}

func (s *DescribeSlowSQLLogsResponseBodyItems) SetDBName(v string) *DescribeSlowSQLLogsResponseBodyItems {
	s.DBName = &v
	return s
}

func (s *DescribeSlowSQLLogsResponseBodyItems) SetOperationExecuteTime(v string) *DescribeSlowSQLLogsResponseBodyItems {
	s.OperationExecuteTime = &v
	return s
}

func (s *DescribeSlowSQLLogsResponseBodyItems) SetScanRowCounts(v int64) *DescribeSlowSQLLogsResponseBodyItems {
	s.ScanRowCounts = &v
	return s
}

func (s *DescribeSlowSQLLogsResponseBodyItems) SetAccountName(v string) *DescribeSlowSQLLogsResponseBodyItems {
	s.AccountName = &v
	return s
}

func (s *DescribeSlowSQLLogsResponseBodyItems) SetQueryId(v string) *DescribeSlowSQLLogsResponseBodyItems {
	s.QueryId = &v
	return s
}

type DescribeSlowSQLLogsResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSlowSQLLogsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSlowSQLLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowSQLLogsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSlowSQLLogsResponse) SetHeaders(v map[string]*string) *DescribeSlowSQLLogsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSlowSQLLogsResponse) SetBody(v *DescribeSlowSQLLogsResponseBody) *DescribeSlowSQLLogsResponse {
	s.Body = v
	return s
}

type DescribeSpecificationRequest struct {
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	StorageType  *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	CpuCores     *int32  `json:"CpuCores,omitempty" xml:"CpuCores,omitempty"`
	TotalNodeNum *int32  `json:"TotalNodeNum,omitempty" xml:"TotalNodeNum,omitempty"`
}

func (s DescribeSpecificationRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSpecificationRequest) GoString() string {
	return s.String()
}

func (s *DescribeSpecificationRequest) SetOwnerId(v int64) *DescribeSpecificationRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSpecificationRequest) SetDBInstanceId(v string) *DescribeSpecificationRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeSpecificationRequest) SetStorageType(v string) *DescribeSpecificationRequest {
	s.StorageType = &v
	return s
}

func (s *DescribeSpecificationRequest) SetCpuCores(v int32) *DescribeSpecificationRequest {
	s.CpuCores = &v
	return s
}

func (s *DescribeSpecificationRequest) SetTotalNodeNum(v int32) *DescribeSpecificationRequest {
	s.TotalNodeNum = &v
	return s
}

type DescribeSpecificationResponseBody struct {
	RequestId            *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DBInstanceClass      []*DescribeSpecificationResponseBodyDBInstanceClass      `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty" type:"Repeated"`
	DBInstanceGroupCount []*DescribeSpecificationResponseBodyDBInstanceGroupCount `json:"DBInstanceGroupCount,omitempty" xml:"DBInstanceGroupCount,omitempty" type:"Repeated"`
	StorageNotice        []*DescribeSpecificationResponseBodyStorageNotice        `json:"StorageNotice,omitempty" xml:"StorageNotice,omitempty" type:"Repeated"`
}

func (s DescribeSpecificationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSpecificationResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSpecificationResponseBody) SetRequestId(v string) *DescribeSpecificationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSpecificationResponseBody) SetDBInstanceClass(v []*DescribeSpecificationResponseBodyDBInstanceClass) *DescribeSpecificationResponseBody {
	s.DBInstanceClass = v
	return s
}

func (s *DescribeSpecificationResponseBody) SetDBInstanceGroupCount(v []*DescribeSpecificationResponseBodyDBInstanceGroupCount) *DescribeSpecificationResponseBody {
	s.DBInstanceGroupCount = v
	return s
}

func (s *DescribeSpecificationResponseBody) SetStorageNotice(v []*DescribeSpecificationResponseBodyStorageNotice) *DescribeSpecificationResponseBody {
	s.StorageNotice = v
	return s
}

type DescribeSpecificationResponseBodyDBInstanceClass struct {
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	Text  *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s DescribeSpecificationResponseBodyDBInstanceClass) String() string {
	return tea.Prettify(s)
}

func (s DescribeSpecificationResponseBodyDBInstanceClass) GoString() string {
	return s.String()
}

func (s *DescribeSpecificationResponseBodyDBInstanceClass) SetValue(v string) *DescribeSpecificationResponseBodyDBInstanceClass {
	s.Value = &v
	return s
}

func (s *DescribeSpecificationResponseBodyDBInstanceClass) SetText(v string) *DescribeSpecificationResponseBodyDBInstanceClass {
	s.Text = &v
	return s
}

type DescribeSpecificationResponseBodyDBInstanceGroupCount struct {
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	Text  *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s DescribeSpecificationResponseBodyDBInstanceGroupCount) String() string {
	return tea.Prettify(s)
}

func (s DescribeSpecificationResponseBodyDBInstanceGroupCount) GoString() string {
	return s.String()
}

func (s *DescribeSpecificationResponseBodyDBInstanceGroupCount) SetValue(v string) *DescribeSpecificationResponseBodyDBInstanceGroupCount {
	s.Value = &v
	return s
}

func (s *DescribeSpecificationResponseBodyDBInstanceGroupCount) SetText(v string) *DescribeSpecificationResponseBodyDBInstanceGroupCount {
	s.Text = &v
	return s
}

type DescribeSpecificationResponseBodyStorageNotice struct {
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	Text  *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s DescribeSpecificationResponseBodyStorageNotice) String() string {
	return tea.Prettify(s)
}

func (s DescribeSpecificationResponseBodyStorageNotice) GoString() string {
	return s.String()
}

func (s *DescribeSpecificationResponseBodyStorageNotice) SetValue(v string) *DescribeSpecificationResponseBodyStorageNotice {
	s.Value = &v
	return s
}

func (s *DescribeSpecificationResponseBodyStorageNotice) SetText(v string) *DescribeSpecificationResponseBodyStorageNotice {
	s.Text = &v
	return s
}

type DescribeSpecificationResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSpecificationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSpecificationResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSpecificationResponse) GoString() string {
	return s.String()
}

func (s *DescribeSpecificationResponse) SetHeaders(v map[string]*string) *DescribeSpecificationResponse {
	s.Headers = v
	return s
}

func (s *DescribeSpecificationResponse) SetBody(v *DescribeSpecificationResponseBody) *DescribeSpecificationResponse {
	s.Body = v
	return s
}

type DescribeSQLCollectorPolicyRequest struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s DescribeSQLCollectorPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLCollectorPolicyRequest) GoString() string {
	return s.String()
}

func (s *DescribeSQLCollectorPolicyRequest) SetDBInstanceId(v string) *DescribeSQLCollectorPolicyRequest {
	s.DBInstanceId = &v
	return s
}

type DescribeSQLCollectorPolicyResponseBody struct {
	SQLCollectorStatus *string `json:"SQLCollectorStatus,omitempty" xml:"SQLCollectorStatus,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeSQLCollectorPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLCollectorPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSQLCollectorPolicyResponseBody) SetSQLCollectorStatus(v string) *DescribeSQLCollectorPolicyResponseBody {
	s.SQLCollectorStatus = &v
	return s
}

func (s *DescribeSQLCollectorPolicyResponseBody) SetRequestId(v string) *DescribeSQLCollectorPolicyResponseBody {
	s.RequestId = &v
	return s
}

type DescribeSQLCollectorPolicyResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSQLCollectorPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSQLCollectorPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLCollectorPolicyResponse) GoString() string {
	return s.String()
}

func (s *DescribeSQLCollectorPolicyResponse) SetHeaders(v map[string]*string) *DescribeSQLCollectorPolicyResponse {
	s.Headers = v
	return s
}

func (s *DescribeSQLCollectorPolicyResponse) SetBody(v *DescribeSQLCollectorPolicyResponseBody) *DescribeSQLCollectorPolicyResponse {
	s.Body = v
	return s
}

type DescribeSQLLogByQueryIdRequest struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	QueryId      *string `json:"QueryId,omitempty" xml:"QueryId,omitempty"`
}

func (s DescribeSQLLogByQueryIdRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLLogByQueryIdRequest) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogByQueryIdRequest) SetDBInstanceId(v string) *DescribeSQLLogByQueryIdRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeSQLLogByQueryIdRequest) SetQueryId(v string) *DescribeSQLLogByQueryIdRequest {
	s.QueryId = &v
	return s
}

type DescribeSQLLogByQueryIdResponseBody struct {
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Items     []*DescribeSQLLogByQueryIdResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
}

func (s DescribeSQLLogByQueryIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLLogByQueryIdResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogByQueryIdResponseBody) SetRequestId(v string) *DescribeSQLLogByQueryIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSQLLogByQueryIdResponseBody) SetItems(v []*DescribeSQLLogByQueryIdResponseBodyItems) *DescribeSQLLogByQueryIdResponseBody {
	s.Items = v
	return s
}

type DescribeSQLLogByQueryIdResponseBodyItems struct {
	OperationClass       *string   `json:"OperationClass,omitempty" xml:"OperationClass,omitempty"`
	ExecuteState         *string   `json:"ExecuteState,omitempty" xml:"ExecuteState,omitempty"`
	ExecuteCost          *float32  `json:"ExecuteCost,omitempty" xml:"ExecuteCost,omitempty"`
	SQLText              *string   `json:"SQLText,omitempty" xml:"SQLText,omitempty"`
	SourcePort           *int32    `json:"SourcePort,omitempty" xml:"SourcePort,omitempty"`
	DBRole               *string   `json:"DBRole,omitempty" xml:"DBRole,omitempty"`
	OperationType        *string   `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	SourceIP             *string   `json:"SourceIP,omitempty" xml:"SourceIP,omitempty"`
	SQLPlan              *string   `json:"SQLPlan,omitempty" xml:"SQLPlan,omitempty"`
	ReturnRowCounts      *int64    `json:"ReturnRowCounts,omitempty" xml:"ReturnRowCounts,omitempty"`
	DBName               *string   `json:"DBName,omitempty" xml:"DBName,omitempty"`
	OperationExecuteTime *string   `json:"OperationExecuteTime,omitempty" xml:"OperationExecuteTime,omitempty"`
	ScanRowCounts        *int64    `json:"ScanRowCounts,omitempty" xml:"ScanRowCounts,omitempty"`
	AccountName          *string   `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	QueryId              *string   `json:"QueryId,omitempty" xml:"QueryId,omitempty"`
	SliceIds             []*string `json:"SliceIds,omitempty" xml:"SliceIds,omitempty" type:"Repeated"`
}

func (s DescribeSQLLogByQueryIdResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLLogByQueryIdResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogByQueryIdResponseBodyItems) SetOperationClass(v string) *DescribeSQLLogByQueryIdResponseBodyItems {
	s.OperationClass = &v
	return s
}

func (s *DescribeSQLLogByQueryIdResponseBodyItems) SetExecuteState(v string) *DescribeSQLLogByQueryIdResponseBodyItems {
	s.ExecuteState = &v
	return s
}

func (s *DescribeSQLLogByQueryIdResponseBodyItems) SetExecuteCost(v float32) *DescribeSQLLogByQueryIdResponseBodyItems {
	s.ExecuteCost = &v
	return s
}

func (s *DescribeSQLLogByQueryIdResponseBodyItems) SetSQLText(v string) *DescribeSQLLogByQueryIdResponseBodyItems {
	s.SQLText = &v
	return s
}

func (s *DescribeSQLLogByQueryIdResponseBodyItems) SetSourcePort(v int32) *DescribeSQLLogByQueryIdResponseBodyItems {
	s.SourcePort = &v
	return s
}

func (s *DescribeSQLLogByQueryIdResponseBodyItems) SetDBRole(v string) *DescribeSQLLogByQueryIdResponseBodyItems {
	s.DBRole = &v
	return s
}

func (s *DescribeSQLLogByQueryIdResponseBodyItems) SetOperationType(v string) *DescribeSQLLogByQueryIdResponseBodyItems {
	s.OperationType = &v
	return s
}

func (s *DescribeSQLLogByQueryIdResponseBodyItems) SetSourceIP(v string) *DescribeSQLLogByQueryIdResponseBodyItems {
	s.SourceIP = &v
	return s
}

func (s *DescribeSQLLogByQueryIdResponseBodyItems) SetSQLPlan(v string) *DescribeSQLLogByQueryIdResponseBodyItems {
	s.SQLPlan = &v
	return s
}

func (s *DescribeSQLLogByQueryIdResponseBodyItems) SetReturnRowCounts(v int64) *DescribeSQLLogByQueryIdResponseBodyItems {
	s.ReturnRowCounts = &v
	return s
}

func (s *DescribeSQLLogByQueryIdResponseBodyItems) SetDBName(v string) *DescribeSQLLogByQueryIdResponseBodyItems {
	s.DBName = &v
	return s
}

func (s *DescribeSQLLogByQueryIdResponseBodyItems) SetOperationExecuteTime(v string) *DescribeSQLLogByQueryIdResponseBodyItems {
	s.OperationExecuteTime = &v
	return s
}

func (s *DescribeSQLLogByQueryIdResponseBodyItems) SetScanRowCounts(v int64) *DescribeSQLLogByQueryIdResponseBodyItems {
	s.ScanRowCounts = &v
	return s
}

func (s *DescribeSQLLogByQueryIdResponseBodyItems) SetAccountName(v string) *DescribeSQLLogByQueryIdResponseBodyItems {
	s.AccountName = &v
	return s
}

func (s *DescribeSQLLogByQueryIdResponseBodyItems) SetQueryId(v string) *DescribeSQLLogByQueryIdResponseBodyItems {
	s.QueryId = &v
	return s
}

func (s *DescribeSQLLogByQueryIdResponseBodyItems) SetSliceIds(v []*string) *DescribeSQLLogByQueryIdResponseBodyItems {
	s.SliceIds = v
	return s
}

type DescribeSQLLogByQueryIdResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSQLLogByQueryIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSQLLogByQueryIdResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLLogByQueryIdResponse) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogByQueryIdResponse) SetHeaders(v map[string]*string) *DescribeSQLLogByQueryIdResponse {
	s.Headers = v
	return s
}

func (s *DescribeSQLLogByQueryIdResponse) SetBody(v *DescribeSQLLogByQueryIdResponseBody) *DescribeSQLLogByQueryIdResponse {
	s.Body = v
	return s
}

type DescribeSQLLogCountRequest struct {
	DBInstanceId   *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	QueryKeywords  *string `json:"QueryKeywords,omitempty" xml:"QueryKeywords,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Database       *string `json:"Database,omitempty" xml:"Database,omitempty"`
	User           *string `json:"User,omitempty" xml:"User,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ExecuteCost    *string `json:"ExecuteCost,omitempty" xml:"ExecuteCost,omitempty"`
	SourceIP       *string `json:"SourceIP,omitempty" xml:"SourceIP,omitempty"`
	ExecuteState   *string `json:"ExecuteState,omitempty" xml:"ExecuteState,omitempty"`
	OperationClass *string `json:"OperationClass,omitempty" xml:"OperationClass,omitempty"`
	OperationType  *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	MaxExecuteCost *string `json:"MaxExecuteCost,omitempty" xml:"MaxExecuteCost,omitempty"`
	MinExecuteCost *string `json:"MinExecuteCost,omitempty" xml:"MinExecuteCost,omitempty"`
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

func (s *DescribeSQLLogCountRequest) SetQueryKeywords(v string) *DescribeSQLLogCountRequest {
	s.QueryKeywords = &v
	return s
}

func (s *DescribeSQLLogCountRequest) SetStartTime(v string) *DescribeSQLLogCountRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSQLLogCountRequest) SetDatabase(v string) *DescribeSQLLogCountRequest {
	s.Database = &v
	return s
}

func (s *DescribeSQLLogCountRequest) SetUser(v string) *DescribeSQLLogCountRequest {
	s.User = &v
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

func (s *DescribeSQLLogCountRequest) SetSourceIP(v string) *DescribeSQLLogCountRequest {
	s.SourceIP = &v
	return s
}

func (s *DescribeSQLLogCountRequest) SetExecuteState(v string) *DescribeSQLLogCountRequest {
	s.ExecuteState = &v
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

func (s *DescribeSQLLogCountRequest) SetMaxExecuteCost(v string) *DescribeSQLLogCountRequest {
	s.MaxExecuteCost = &v
	return s
}

func (s *DescribeSQLLogCountRequest) SetMinExecuteCost(v string) *DescribeSQLLogCountRequest {
	s.MinExecuteCost = &v
	return s
}

type DescribeSQLLogCountResponseBody struct {
	EndTime     *string                                 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RequestId   *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime   *string                                 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	DBClusterId *string                                 `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	Items       []*DescribeSQLLogCountResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
}

func (s DescribeSQLLogCountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLLogCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogCountResponseBody) SetEndTime(v string) *DescribeSQLLogCountResponseBody {
	s.EndTime = &v
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

func (s *DescribeSQLLogCountResponseBody) SetDBClusterId(v string) *DescribeSQLLogCountResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeSQLLogCountResponseBody) SetItems(v []*DescribeSQLLogCountResponseBodyItems) *DescribeSQLLogCountResponseBody {
	s.Items = v
	return s
}

type DescribeSQLLogCountResponseBodyItems struct {
	Series []*DescribeSQLLogCountResponseBodyItemsSeries `json:"Series,omitempty" xml:"Series,omitempty" type:"Repeated"`
	Name   *string                                       `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeSQLLogCountResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLLogCountResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogCountResponseBodyItems) SetSeries(v []*DescribeSQLLogCountResponseBodyItemsSeries) *DescribeSQLLogCountResponseBodyItems {
	s.Series = v
	return s
}

func (s *DescribeSQLLogCountResponseBodyItems) SetName(v string) *DescribeSQLLogCountResponseBodyItems {
	s.Name = &v
	return s
}

type DescribeSQLLogCountResponseBodyItemsSeries struct {
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
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSQLLogCountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeSQLLogCountResponse) SetBody(v *DescribeSQLLogCountResponseBody) *DescribeSQLLogCountResponse {
	s.Body = v
	return s
}

type DescribeSQLLogFilesRequest struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	FileName     *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber   *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s DescribeSQLLogFilesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLLogFilesRequest) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogFilesRequest) SetDBInstanceId(v string) *DescribeSQLLogFilesRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeSQLLogFilesRequest) SetFileName(v string) *DescribeSQLLogFilesRequest {
	s.FileName = &v
	return s
}

func (s *DescribeSQLLogFilesRequest) SetPageSize(v int32) *DescribeSQLLogFilesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSQLLogFilesRequest) SetPageNumber(v int32) *DescribeSQLLogFilesRequest {
	s.PageNumber = &v
	return s
}

type DescribeSQLLogFilesResponseBody struct {
	RequestId        *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber       *int32                                `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageRecordCount  *int32                                `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	TotalRecordCount *int32                                `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
	Items            *DescribeSQLLogFilesResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
}

func (s DescribeSQLLogFilesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLLogFilesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogFilesResponseBody) SetRequestId(v string) *DescribeSQLLogFilesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSQLLogFilesResponseBody) SetPageNumber(v int32) *DescribeSQLLogFilesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSQLLogFilesResponseBody) SetPageRecordCount(v int32) *DescribeSQLLogFilesResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeSQLLogFilesResponseBody) SetTotalRecordCount(v int32) *DescribeSQLLogFilesResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeSQLLogFilesResponseBody) SetItems(v *DescribeSQLLogFilesResponseBodyItems) *DescribeSQLLogFilesResponseBody {
	s.Items = v
	return s
}

type DescribeSQLLogFilesResponseBodyItems struct {
	LogFile []*DescribeSQLLogFilesResponseBodyItemsLogFile `json:"LogFile,omitempty" xml:"LogFile,omitempty" type:"Repeated"`
}

func (s DescribeSQLLogFilesResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLLogFilesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogFilesResponseBodyItems) SetLogFile(v []*DescribeSQLLogFilesResponseBodyItemsLogFile) *DescribeSQLLogFilesResponseBodyItems {
	s.LogFile = v
	return s
}

type DescribeSQLLogFilesResponseBodyItemsLogFile struct {
	FileID         *string `json:"FileID,omitempty" xml:"FileID,omitempty"`
	LogStartTime   *string `json:"LogStartTime,omitempty" xml:"LogStartTime,omitempty"`
	LogSize        *string `json:"LogSize,omitempty" xml:"LogSize,omitempty"`
	LogDownloadURL *string `json:"LogDownloadURL,omitempty" xml:"LogDownloadURL,omitempty"`
	LogEndTime     *string `json:"LogEndTime,omitempty" xml:"LogEndTime,omitempty"`
	LogStatus      *string `json:"LogStatus,omitempty" xml:"LogStatus,omitempty"`
}

func (s DescribeSQLLogFilesResponseBodyItemsLogFile) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLLogFilesResponseBodyItemsLogFile) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogFilesResponseBodyItemsLogFile) SetFileID(v string) *DescribeSQLLogFilesResponseBodyItemsLogFile {
	s.FileID = &v
	return s
}

func (s *DescribeSQLLogFilesResponseBodyItemsLogFile) SetLogStartTime(v string) *DescribeSQLLogFilesResponseBodyItemsLogFile {
	s.LogStartTime = &v
	return s
}

func (s *DescribeSQLLogFilesResponseBodyItemsLogFile) SetLogSize(v string) *DescribeSQLLogFilesResponseBodyItemsLogFile {
	s.LogSize = &v
	return s
}

func (s *DescribeSQLLogFilesResponseBodyItemsLogFile) SetLogDownloadURL(v string) *DescribeSQLLogFilesResponseBodyItemsLogFile {
	s.LogDownloadURL = &v
	return s
}

func (s *DescribeSQLLogFilesResponseBodyItemsLogFile) SetLogEndTime(v string) *DescribeSQLLogFilesResponseBodyItemsLogFile {
	s.LogEndTime = &v
	return s
}

func (s *DescribeSQLLogFilesResponseBodyItemsLogFile) SetLogStatus(v string) *DescribeSQLLogFilesResponseBodyItemsLogFile {
	s.LogStatus = &v
	return s
}

type DescribeSQLLogFilesResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSQLLogFilesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSQLLogFilesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLLogFilesResponse) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogFilesResponse) SetHeaders(v map[string]*string) *DescribeSQLLogFilesResponse {
	s.Headers = v
	return s
}

func (s *DescribeSQLLogFilesResponse) SetBody(v *DescribeSQLLogFilesResponseBody) *DescribeSQLLogFilesResponse {
	s.Body = v
	return s
}

type DescribeSQLLogRecordsRequest struct {
	DBInstanceId  *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	QueryKeywords *string `json:"QueryKeywords,omitempty" xml:"QueryKeywords,omitempty"`
	StartTime     *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Database      *string `json:"Database,omitempty" xml:"Database,omitempty"`
	User          *string `json:"User,omitempty" xml:"User,omitempty"`
	Form          *string `json:"Form,omitempty" xml:"Form,omitempty"`
	EndTime       *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s DescribeSQLLogRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLLogRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogRecordsRequest) SetDBInstanceId(v string) *DescribeSQLLogRecordsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeSQLLogRecordsRequest) SetQueryKeywords(v string) *DescribeSQLLogRecordsRequest {
	s.QueryKeywords = &v
	return s
}

func (s *DescribeSQLLogRecordsRequest) SetStartTime(v string) *DescribeSQLLogRecordsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSQLLogRecordsRequest) SetDatabase(v string) *DescribeSQLLogRecordsRequest {
	s.Database = &v
	return s
}

func (s *DescribeSQLLogRecordsRequest) SetUser(v string) *DescribeSQLLogRecordsRequest {
	s.User = &v
	return s
}

func (s *DescribeSQLLogRecordsRequest) SetForm(v string) *DescribeSQLLogRecordsRequest {
	s.Form = &v
	return s
}

func (s *DescribeSQLLogRecordsRequest) SetEndTime(v string) *DescribeSQLLogRecordsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSQLLogRecordsRequest) SetPageSize(v int32) *DescribeSQLLogRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSQLLogRecordsRequest) SetPageNumber(v int32) *DescribeSQLLogRecordsRequest {
	s.PageNumber = &v
	return s
}

type DescribeSQLLogRecordsResponseBody struct {
	RequestId        *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber       *int32                                  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageRecordCount  *int32                                  `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	TotalRecordCount *int32                                  `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
	Items            *DescribeSQLLogRecordsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
}

func (s DescribeSQLLogRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLLogRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogRecordsResponseBody) SetRequestId(v string) *DescribeSQLLogRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSQLLogRecordsResponseBody) SetPageNumber(v int32) *DescribeSQLLogRecordsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSQLLogRecordsResponseBody) SetPageRecordCount(v int32) *DescribeSQLLogRecordsResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeSQLLogRecordsResponseBody) SetTotalRecordCount(v int32) *DescribeSQLLogRecordsResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeSQLLogRecordsResponseBody) SetItems(v *DescribeSQLLogRecordsResponseBodyItems) *DescribeSQLLogRecordsResponseBody {
	s.Items = v
	return s
}

type DescribeSQLLogRecordsResponseBodyItems struct {
	SQLRecord []*DescribeSQLLogRecordsResponseBodyItemsSQLRecord `json:"SQLRecord,omitempty" xml:"SQLRecord,omitempty" type:"Repeated"`
}

func (s DescribeSQLLogRecordsResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLLogRecordsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogRecordsResponseBodyItems) SetSQLRecord(v []*DescribeSQLLogRecordsResponseBodyItemsSQLRecord) *DescribeSQLLogRecordsResponseBodyItems {
	s.SQLRecord = v
	return s
}

type DescribeSQLLogRecordsResponseBodyItemsSQLRecord struct {
	HostAddress         *string `json:"HostAddress,omitempty" xml:"HostAddress,omitempty"`
	SQLText             *string `json:"SQLText,omitempty" xml:"SQLText,omitempty"`
	ReturnRowCounts     *int64  `json:"ReturnRowCounts,omitempty" xml:"ReturnRowCounts,omitempty"`
	DBName              *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	ExecuteTime         *string `json:"ExecuteTime,omitempty" xml:"ExecuteTime,omitempty"`
	ThreadID            *string `json:"ThreadID,omitempty" xml:"ThreadID,omitempty"`
	TotalExecutionTimes *int64  `json:"TotalExecutionTimes,omitempty" xml:"TotalExecutionTimes,omitempty"`
	AccountName         *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
}

func (s DescribeSQLLogRecordsResponseBodyItemsSQLRecord) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLLogRecordsResponseBodyItemsSQLRecord) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogRecordsResponseBodyItemsSQLRecord) SetHostAddress(v string) *DescribeSQLLogRecordsResponseBodyItemsSQLRecord {
	s.HostAddress = &v
	return s
}

func (s *DescribeSQLLogRecordsResponseBodyItemsSQLRecord) SetSQLText(v string) *DescribeSQLLogRecordsResponseBodyItemsSQLRecord {
	s.SQLText = &v
	return s
}

func (s *DescribeSQLLogRecordsResponseBodyItemsSQLRecord) SetReturnRowCounts(v int64) *DescribeSQLLogRecordsResponseBodyItemsSQLRecord {
	s.ReturnRowCounts = &v
	return s
}

func (s *DescribeSQLLogRecordsResponseBodyItemsSQLRecord) SetDBName(v string) *DescribeSQLLogRecordsResponseBodyItemsSQLRecord {
	s.DBName = &v
	return s
}

func (s *DescribeSQLLogRecordsResponseBodyItemsSQLRecord) SetExecuteTime(v string) *DescribeSQLLogRecordsResponseBodyItemsSQLRecord {
	s.ExecuteTime = &v
	return s
}

func (s *DescribeSQLLogRecordsResponseBodyItemsSQLRecord) SetThreadID(v string) *DescribeSQLLogRecordsResponseBodyItemsSQLRecord {
	s.ThreadID = &v
	return s
}

func (s *DescribeSQLLogRecordsResponseBodyItemsSQLRecord) SetTotalExecutionTimes(v int64) *DescribeSQLLogRecordsResponseBodyItemsSQLRecord {
	s.TotalExecutionTimes = &v
	return s
}

func (s *DescribeSQLLogRecordsResponseBodyItemsSQLRecord) SetAccountName(v string) *DescribeSQLLogRecordsResponseBodyItemsSQLRecord {
	s.AccountName = &v
	return s
}

type DescribeSQLLogRecordsResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSQLLogRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSQLLogRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLLogRecordsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogRecordsResponse) SetHeaders(v map[string]*string) *DescribeSQLLogRecordsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSQLLogRecordsResponse) SetBody(v *DescribeSQLLogRecordsResponseBody) *DescribeSQLLogRecordsResponse {
	s.Body = v
	return s
}

type DescribeSQLLogsRequest struct {
	DBInstanceId   *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	QueryKeywords  *string `json:"QueryKeywords,omitempty" xml:"QueryKeywords,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Database       *string `json:"Database,omitempty" xml:"Database,omitempty"`
	User           *string `json:"User,omitempty" xml:"User,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber     *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	ExecuteCost    *string `json:"ExecuteCost,omitempty" xml:"ExecuteCost,omitempty"`
	SourceIP       *string `json:"SourceIP,omitempty" xml:"SourceIP,omitempty"`
	ExecuteState   *string `json:"ExecuteState,omitempty" xml:"ExecuteState,omitempty"`
	OperationClass *string `json:"OperationClass,omitempty" xml:"OperationClass,omitempty"`
	OperationType  *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	MaxExecuteCost *string `json:"MaxExecuteCost,omitempty" xml:"MaxExecuteCost,omitempty"`
	MinExecuteCost *string `json:"MinExecuteCost,omitempty" xml:"MinExecuteCost,omitempty"`
}

func (s DescribeSQLLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLLogsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogsRequest) SetDBInstanceId(v string) *DescribeSQLLogsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeSQLLogsRequest) SetQueryKeywords(v string) *DescribeSQLLogsRequest {
	s.QueryKeywords = &v
	return s
}

func (s *DescribeSQLLogsRequest) SetStartTime(v string) *DescribeSQLLogsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSQLLogsRequest) SetDatabase(v string) *DescribeSQLLogsRequest {
	s.Database = &v
	return s
}

func (s *DescribeSQLLogsRequest) SetUser(v string) *DescribeSQLLogsRequest {
	s.User = &v
	return s
}

func (s *DescribeSQLLogsRequest) SetEndTime(v string) *DescribeSQLLogsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSQLLogsRequest) SetPageSize(v int32) *DescribeSQLLogsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSQLLogsRequest) SetPageNumber(v int32) *DescribeSQLLogsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSQLLogsRequest) SetExecuteCost(v string) *DescribeSQLLogsRequest {
	s.ExecuteCost = &v
	return s
}

func (s *DescribeSQLLogsRequest) SetSourceIP(v string) *DescribeSQLLogsRequest {
	s.SourceIP = &v
	return s
}

func (s *DescribeSQLLogsRequest) SetExecuteState(v string) *DescribeSQLLogsRequest {
	s.ExecuteState = &v
	return s
}

func (s *DescribeSQLLogsRequest) SetOperationClass(v string) *DescribeSQLLogsRequest {
	s.OperationClass = &v
	return s
}

func (s *DescribeSQLLogsRequest) SetOperationType(v string) *DescribeSQLLogsRequest {
	s.OperationType = &v
	return s
}

func (s *DescribeSQLLogsRequest) SetMaxExecuteCost(v string) *DescribeSQLLogsRequest {
	s.MaxExecuteCost = &v
	return s
}

func (s *DescribeSQLLogsRequest) SetMinExecuteCost(v string) *DescribeSQLLogsRequest {
	s.MinExecuteCost = &v
	return s
}

type DescribeSQLLogsResponseBody struct {
	PageRecordCount *int32                              `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	RequestId       *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber      *int32                              `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Items           []*DescribeSQLLogsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
}

func (s DescribeSQLLogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLLogsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogsResponseBody) SetPageRecordCount(v int32) *DescribeSQLLogsResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeSQLLogsResponseBody) SetRequestId(v string) *DescribeSQLLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSQLLogsResponseBody) SetPageNumber(v int32) *DescribeSQLLogsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSQLLogsResponseBody) SetItems(v []*DescribeSQLLogsResponseBodyItems) *DescribeSQLLogsResponseBody {
	s.Items = v
	return s
}

type DescribeSQLLogsResponseBodyItems struct {
	OperationClass       *string  `json:"OperationClass,omitempty" xml:"OperationClass,omitempty"`
	ExecuteState         *string  `json:"ExecuteState,omitempty" xml:"ExecuteState,omitempty"`
	ExecuteCost          *float32 `json:"ExecuteCost,omitempty" xml:"ExecuteCost,omitempty"`
	SQLText              *string  `json:"SQLText,omitempty" xml:"SQLText,omitempty"`
	SourcePort           *int32   `json:"SourcePort,omitempty" xml:"SourcePort,omitempty"`
	DBRole               *string  `json:"DBRole,omitempty" xml:"DBRole,omitempty"`
	OperationType        *string  `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	SourceIP             *string  `json:"SourceIP,omitempty" xml:"SourceIP,omitempty"`
	SQLPlan              *string  `json:"SQLPlan,omitempty" xml:"SQLPlan,omitempty"`
	ReturnRowCounts      *int64   `json:"ReturnRowCounts,omitempty" xml:"ReturnRowCounts,omitempty"`
	DBName               *string  `json:"DBName,omitempty" xml:"DBName,omitempty"`
	OperationExecuteTime *string  `json:"OperationExecuteTime,omitempty" xml:"OperationExecuteTime,omitempty"`
	ScanRowCounts        *int64   `json:"ScanRowCounts,omitempty" xml:"ScanRowCounts,omitempty"`
	AccountName          *string  `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
}

func (s DescribeSQLLogsResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLLogsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogsResponseBodyItems) SetOperationClass(v string) *DescribeSQLLogsResponseBodyItems {
	s.OperationClass = &v
	return s
}

func (s *DescribeSQLLogsResponseBodyItems) SetExecuteState(v string) *DescribeSQLLogsResponseBodyItems {
	s.ExecuteState = &v
	return s
}

func (s *DescribeSQLLogsResponseBodyItems) SetExecuteCost(v float32) *DescribeSQLLogsResponseBodyItems {
	s.ExecuteCost = &v
	return s
}

func (s *DescribeSQLLogsResponseBodyItems) SetSQLText(v string) *DescribeSQLLogsResponseBodyItems {
	s.SQLText = &v
	return s
}

func (s *DescribeSQLLogsResponseBodyItems) SetSourcePort(v int32) *DescribeSQLLogsResponseBodyItems {
	s.SourcePort = &v
	return s
}

func (s *DescribeSQLLogsResponseBodyItems) SetDBRole(v string) *DescribeSQLLogsResponseBodyItems {
	s.DBRole = &v
	return s
}

func (s *DescribeSQLLogsResponseBodyItems) SetOperationType(v string) *DescribeSQLLogsResponseBodyItems {
	s.OperationType = &v
	return s
}

func (s *DescribeSQLLogsResponseBodyItems) SetSourceIP(v string) *DescribeSQLLogsResponseBodyItems {
	s.SourceIP = &v
	return s
}

func (s *DescribeSQLLogsResponseBodyItems) SetSQLPlan(v string) *DescribeSQLLogsResponseBodyItems {
	s.SQLPlan = &v
	return s
}

func (s *DescribeSQLLogsResponseBodyItems) SetReturnRowCounts(v int64) *DescribeSQLLogsResponseBodyItems {
	s.ReturnRowCounts = &v
	return s
}

func (s *DescribeSQLLogsResponseBodyItems) SetDBName(v string) *DescribeSQLLogsResponseBodyItems {
	s.DBName = &v
	return s
}

func (s *DescribeSQLLogsResponseBodyItems) SetOperationExecuteTime(v string) *DescribeSQLLogsResponseBodyItems {
	s.OperationExecuteTime = &v
	return s
}

func (s *DescribeSQLLogsResponseBodyItems) SetScanRowCounts(v int64) *DescribeSQLLogsResponseBodyItems {
	s.ScanRowCounts = &v
	return s
}

func (s *DescribeSQLLogsResponseBodyItems) SetAccountName(v string) *DescribeSQLLogsResponseBodyItems {
	s.AccountName = &v
	return s
}

type DescribeSQLLogsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSQLLogsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSQLLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLLogsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogsResponse) SetHeaders(v map[string]*string) *DescribeSQLLogsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSQLLogsResponse) SetBody(v *DescribeSQLLogsResponseBody) *DescribeSQLLogsResponse {
	s.Body = v
	return s
}

type DescribeSQLLogsOnSliceRequest struct {
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber     *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	MaxExecuteCost *string `json:"MaxExecuteCost,omitempty" xml:"MaxExecuteCost,omitempty"`
	MinExecuteCost *string `json:"MinExecuteCost,omitempty" xml:"MinExecuteCost,omitempty"`
	ExecuteState   *string `json:"ExecuteState,omitempty" xml:"ExecuteState,omitempty"`
	DBInstanceId   *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	QueryId        *string `json:"QueryId,omitempty" xml:"QueryId,omitempty"`
	SliceId        *string `json:"SliceId,omitempty" xml:"SliceId,omitempty"`
}

func (s DescribeSQLLogsOnSliceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLLogsOnSliceRequest) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogsOnSliceRequest) SetPageSize(v int32) *DescribeSQLLogsOnSliceRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSQLLogsOnSliceRequest) SetPageNumber(v int32) *DescribeSQLLogsOnSliceRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSQLLogsOnSliceRequest) SetMaxExecuteCost(v string) *DescribeSQLLogsOnSliceRequest {
	s.MaxExecuteCost = &v
	return s
}

func (s *DescribeSQLLogsOnSliceRequest) SetMinExecuteCost(v string) *DescribeSQLLogsOnSliceRequest {
	s.MinExecuteCost = &v
	return s
}

func (s *DescribeSQLLogsOnSliceRequest) SetExecuteState(v string) *DescribeSQLLogsOnSliceRequest {
	s.ExecuteState = &v
	return s
}

func (s *DescribeSQLLogsOnSliceRequest) SetDBInstanceId(v string) *DescribeSQLLogsOnSliceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeSQLLogsOnSliceRequest) SetQueryId(v string) *DescribeSQLLogsOnSliceRequest {
	s.QueryId = &v
	return s
}

func (s *DescribeSQLLogsOnSliceRequest) SetSliceId(v string) *DescribeSQLLogsOnSliceRequest {
	s.SliceId = &v
	return s
}

type DescribeSQLLogsOnSliceResponseBody struct {
	RequestId       *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageRecordCount *int32                                             `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	PageNumber      *int32                                             `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	SliceLogItems   []*DescribeSQLLogsOnSliceResponseBodySliceLogItems `json:"SliceLogItems,omitempty" xml:"SliceLogItems,omitempty" type:"Repeated"`
}

func (s DescribeSQLLogsOnSliceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLLogsOnSliceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogsOnSliceResponseBody) SetRequestId(v string) *DescribeSQLLogsOnSliceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSQLLogsOnSliceResponseBody) SetPageRecordCount(v int32) *DescribeSQLLogsOnSliceResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeSQLLogsOnSliceResponseBody) SetPageNumber(v int32) *DescribeSQLLogsOnSliceResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSQLLogsOnSliceResponseBody) SetSliceLogItems(v []*DescribeSQLLogsOnSliceResponseBodySliceLogItems) *DescribeSQLLogsOnSliceResponseBody {
	s.SliceLogItems = v
	return s
}

type DescribeSQLLogsOnSliceResponseBodySliceLogItems struct {
	ExecuteStatus           *string  `json:"ExecuteStatus,omitempty" xml:"ExecuteStatus,omitempty"`
	ExecuteCost             *float32 `json:"ExecuteCost,omitempty" xml:"ExecuteCost,omitempty"`
	ReturnRowCounts         *int64   `json:"ReturnRowCounts,omitempty" xml:"ReturnRowCounts,omitempty"`
	OperationExecuteTime    *string  `json:"OperationExecuteTime,omitempty" xml:"OperationExecuteTime,omitempty"`
	SegmentId               *string  `json:"SegmentId,omitempty" xml:"SegmentId,omitempty"`
	PeakMemory              *float32 `json:"PeakMemory,omitempty" xml:"PeakMemory,omitempty"`
	OperationExecuteEndTime *string  `json:"OperationExecuteEndTime,omitempty" xml:"OperationExecuteEndTime,omitempty"`
	SegmentName             *string  `json:"SegmentName,omitempty" xml:"SegmentName,omitempty"`
}

func (s DescribeSQLLogsOnSliceResponseBodySliceLogItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLLogsOnSliceResponseBodySliceLogItems) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogsOnSliceResponseBodySliceLogItems) SetExecuteStatus(v string) *DescribeSQLLogsOnSliceResponseBodySliceLogItems {
	s.ExecuteStatus = &v
	return s
}

func (s *DescribeSQLLogsOnSliceResponseBodySliceLogItems) SetExecuteCost(v float32) *DescribeSQLLogsOnSliceResponseBodySliceLogItems {
	s.ExecuteCost = &v
	return s
}

func (s *DescribeSQLLogsOnSliceResponseBodySliceLogItems) SetReturnRowCounts(v int64) *DescribeSQLLogsOnSliceResponseBodySliceLogItems {
	s.ReturnRowCounts = &v
	return s
}

func (s *DescribeSQLLogsOnSliceResponseBodySliceLogItems) SetOperationExecuteTime(v string) *DescribeSQLLogsOnSliceResponseBodySliceLogItems {
	s.OperationExecuteTime = &v
	return s
}

func (s *DescribeSQLLogsOnSliceResponseBodySliceLogItems) SetSegmentId(v string) *DescribeSQLLogsOnSliceResponseBodySliceLogItems {
	s.SegmentId = &v
	return s
}

func (s *DescribeSQLLogsOnSliceResponseBodySliceLogItems) SetPeakMemory(v float32) *DescribeSQLLogsOnSliceResponseBodySliceLogItems {
	s.PeakMemory = &v
	return s
}

func (s *DescribeSQLLogsOnSliceResponseBodySliceLogItems) SetOperationExecuteEndTime(v string) *DescribeSQLLogsOnSliceResponseBodySliceLogItems {
	s.OperationExecuteEndTime = &v
	return s
}

func (s *DescribeSQLLogsOnSliceResponseBodySliceLogItems) SetSegmentName(v string) *DescribeSQLLogsOnSliceResponseBodySliceLogItems {
	s.SegmentName = &v
	return s
}

type DescribeSQLLogsOnSliceResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSQLLogsOnSliceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSQLLogsOnSliceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLLogsOnSliceResponse) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogsOnSliceResponse) SetHeaders(v map[string]*string) *DescribeSQLLogsOnSliceResponse {
	s.Headers = v
	return s
}

func (s *DescribeSQLLogsOnSliceResponse) SetBody(v *DescribeSQLLogsOnSliceResponseBody) *DescribeSQLLogsOnSliceResponse {
	s.Body = v
	return s
}

type DescribeTagsRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceType         *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s DescribeTagsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTagsRequest) GoString() string {
	return s.String()
}

func (s *DescribeTagsRequest) SetOwnerId(v int64) *DescribeTagsRequest {
	s.OwnerId = &v
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

func (s *DescribeTagsRequest) SetOwnerAccount(v string) *DescribeTagsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeTagsRequest) SetRegionId(v string) *DescribeTagsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeTagsRequest) SetResourceType(v string) *DescribeTagsRequest {
	s.ResourceType = &v
	return s
}

type DescribeTagsResponseBody struct {
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Tags      []*DescribeTagsResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
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
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
}

func (s DescribeTagsResponseBodyTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeTagsResponseBodyTags) GoString() string {
	return s.String()
}

func (s *DescribeTagsResponseBodyTags) SetTagValue(v string) *DescribeTagsResponseBodyTags {
	s.TagValue = &v
	return s
}

func (s *DescribeTagsResponseBodyTags) SetTagKey(v string) *DescribeTagsResponseBodyTags {
	s.TagKey = &v
	return s
}

type DescribeTagsResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeTagsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeTagsResponse) SetBody(v *DescribeTagsResponseBody) *DescribeTagsResponse {
	s.Body = v
	return s
}

type DescribeUserEncryptionKeyListRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeUserEncryptionKeyListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserEncryptionKeyListRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserEncryptionKeyListRequest) SetRegionId(v string) *DescribeUserEncryptionKeyListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeUserEncryptionKeyListRequest) SetPageNumber(v string) *DescribeUserEncryptionKeyListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeUserEncryptionKeyListRequest) SetPageSize(v string) *DescribeUserEncryptionKeyListRequest {
	s.PageSize = &v
	return s
}

type DescribeUserEncryptionKeyListResponseBody struct {
	RequestId *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	KmsKeys   []*DescribeUserEncryptionKeyListResponseBodyKmsKeys `json:"KmsKeys,omitempty" xml:"KmsKeys,omitempty" type:"Repeated"`
}

func (s DescribeUserEncryptionKeyListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserEncryptionKeyListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserEncryptionKeyListResponseBody) SetRequestId(v string) *DescribeUserEncryptionKeyListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserEncryptionKeyListResponseBody) SetKmsKeys(v []*DescribeUserEncryptionKeyListResponseBodyKmsKeys) *DescribeUserEncryptionKeyListResponseBody {
	s.KmsKeys = v
	return s
}

type DescribeUserEncryptionKeyListResponseBodyKmsKeys struct {
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
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeUserEncryptionKeyListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeUserEncryptionKeyListResponse) SetBody(v *DescribeUserEncryptionKeyListResponseBody) *DescribeUserEncryptionKeyListResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	OwnerId              *int64                        `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string                       `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                        `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string                       `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	RegionId             *string                       `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceType         *string                       `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	NextToken            *string                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	ResourceId           []*string                     `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	Tag                  []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) SetOwnerId(v int64) *ListTagResourcesRequest {
	s.OwnerId = &v
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

func (s *ListTagResourcesRequest) SetOwnerAccount(v string) *ListTagResourcesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListTagResourcesRequest) SetRegionId(v string) *ListTagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceType(v string) *ListTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesRequest) SetNextToken(v string) *ListTagResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceId(v []*string) *ListTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *ListTagResourcesRequest) SetTag(v []*ListTagResourcesRequestTag) *ListTagResourcesRequest {
	s.Tag = v
	return s
}

type ListTagResourcesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
	NextToken    *string                                   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId    *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagValue     *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	TagKey       *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
}

func (s ListTagResourcesResponseBodyTagResourcesTagResource) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResourcesTagResource) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetResourceType(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetTagValue(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.TagValue = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetResourceId(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetTagKey(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.TagKey = &v
	return s
}

type ListTagResourcesResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListTagResourcesResponse) SetBody(v *ListTagResourcesResponseBody) *ListTagResourcesResponse {
	s.Body = v
	return s
}

type ModifyAccountDescriptionRequest struct {
	DBInstanceId       *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	AccountName        *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	AccountDescription *string `json:"AccountDescription,omitempty" xml:"AccountDescription,omitempty"`
}

func (s ModifyAccountDescriptionRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccountDescriptionRequest) GoString() string {
	return s.String()
}

func (s *ModifyAccountDescriptionRequest) SetDBInstanceId(v string) *ModifyAccountDescriptionRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyAccountDescriptionRequest) SetAccountName(v string) *ModifyAccountDescriptionRequest {
	s.AccountName = &v
	return s
}

func (s *ModifyAccountDescriptionRequest) SetAccountDescription(v string) *ModifyAccountDescriptionRequest {
	s.AccountDescription = &v
	return s
}

type ModifyAccountDescriptionResponseBody struct {
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
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyAccountDescriptionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ModifyAccountDescriptionResponse) SetBody(v *ModifyAccountDescriptionResponseBody) *ModifyAccountDescriptionResponse {
	s.Body = v
	return s
}

type ModifyBackupPolicyRequest struct {
	DBInstanceId          *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	PreferredBackupTime   *string `json:"PreferredBackupTime,omitempty" xml:"PreferredBackupTime,omitempty"`
	PreferredBackupPeriod *string `json:"PreferredBackupPeriod,omitempty" xml:"PreferredBackupPeriod,omitempty"`
	BackupRetentionPeriod *int32  `json:"BackupRetentionPeriod,omitempty" xml:"BackupRetentionPeriod,omitempty"`
	EnableRecoveryPoint   *bool   `json:"EnableRecoveryPoint,omitempty" xml:"EnableRecoveryPoint,omitempty"`
	RecoveryPointPeriod   *string `json:"RecoveryPointPeriod,omitempty" xml:"RecoveryPointPeriod,omitempty"`
}

func (s ModifyBackupPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyBackupPolicyRequest) GoString() string {
	return s.String()
}

func (s *ModifyBackupPolicyRequest) SetDBInstanceId(v string) *ModifyBackupPolicyRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetPreferredBackupTime(v string) *ModifyBackupPolicyRequest {
	s.PreferredBackupTime = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetPreferredBackupPeriod(v string) *ModifyBackupPolicyRequest {
	s.PreferredBackupPeriod = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetBackupRetentionPeriod(v int32) *ModifyBackupPolicyRequest {
	s.BackupRetentionPeriod = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetEnableRecoveryPoint(v bool) *ModifyBackupPolicyRequest {
	s.EnableRecoveryPoint = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetRecoveryPointPeriod(v string) *ModifyBackupPolicyRequest {
	s.RecoveryPointPeriod = &v
	return s
}

type ModifyBackupPolicyResponseBody struct {
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
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyBackupPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ModifyBackupPolicyResponse) SetBody(v *ModifyBackupPolicyResponseBody) *ModifyBackupPolicyResponse {
	s.Body = v
	return s
}

type ModifyDBInstanceConnectionModeRequest struct {
	DBInstanceId   *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	ConnectionMode *string `json:"ConnectionMode,omitempty" xml:"ConnectionMode,omitempty"`
}

func (s ModifyDBInstanceConnectionModeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceConnectionModeRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceConnectionModeRequest) SetDBInstanceId(v string) *ModifyDBInstanceConnectionModeRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceConnectionModeRequest) SetConnectionMode(v string) *ModifyDBInstanceConnectionModeRequest {
	s.ConnectionMode = &v
	return s
}

type ModifyDBInstanceConnectionModeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBInstanceConnectionModeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceConnectionModeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceConnectionModeResponseBody) SetRequestId(v string) *ModifyDBInstanceConnectionModeResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDBInstanceConnectionModeResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDBInstanceConnectionModeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDBInstanceConnectionModeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceConnectionModeResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceConnectionModeResponse) SetHeaders(v map[string]*string) *ModifyDBInstanceConnectionModeResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBInstanceConnectionModeResponse) SetBody(v *ModifyDBInstanceConnectionModeResponseBody) *ModifyDBInstanceConnectionModeResponse {
	s.Body = v
	return s
}

type ModifyDBInstanceConnectionStringRequest struct {
	DBInstanceId            *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	ConnectionStringPrefix  *string `json:"ConnectionStringPrefix,omitempty" xml:"ConnectionStringPrefix,omitempty"`
	Port                    *string `json:"Port,omitempty" xml:"Port,omitempty"`
	CurrentConnectionString *string `json:"CurrentConnectionString,omitempty" xml:"CurrentConnectionString,omitempty"`
}

func (s ModifyDBInstanceConnectionStringRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceConnectionStringRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceConnectionStringRequest) SetDBInstanceId(v string) *ModifyDBInstanceConnectionStringRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringRequest) SetConnectionStringPrefix(v string) *ModifyDBInstanceConnectionStringRequest {
	s.ConnectionStringPrefix = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringRequest) SetPort(v string) *ModifyDBInstanceConnectionStringRequest {
	s.Port = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringRequest) SetCurrentConnectionString(v string) *ModifyDBInstanceConnectionStringRequest {
	s.CurrentConnectionString = &v
	return s
}

type ModifyDBInstanceConnectionStringResponseBody struct {
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
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDBInstanceConnectionStringResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ModifyDBInstanceConnectionStringResponse) SetBody(v *ModifyDBInstanceConnectionStringResponseBody) *ModifyDBInstanceConnectionStringResponse {
	s.Body = v
	return s
}

type ModifyDBInstanceDescriptionRequest struct {
	DBInstanceId          *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	DBInstanceDescription *string `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
}

func (s ModifyDBInstanceDescriptionRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceDescriptionRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceDescriptionRequest) SetDBInstanceId(v string) *ModifyDBInstanceDescriptionRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceDescriptionRequest) SetDBInstanceDescription(v string) *ModifyDBInstanceDescriptionRequest {
	s.DBInstanceDescription = &v
	return s
}

type ModifyDBInstanceDescriptionResponseBody struct {
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
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDBInstanceDescriptionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ModifyDBInstanceDescriptionResponse) SetBody(v *ModifyDBInstanceDescriptionResponseBody) *ModifyDBInstanceDescriptionResponse {
	s.Body = v
	return s
}

type ModifyDBInstanceMaintainTimeRequest struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
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

func (s *ModifyDBInstanceMaintainTimeRequest) SetStartTime(v string) *ModifyDBInstanceMaintainTimeRequest {
	s.StartTime = &v
	return s
}

func (s *ModifyDBInstanceMaintainTimeRequest) SetEndTime(v string) *ModifyDBInstanceMaintainTimeRequest {
	s.EndTime = &v
	return s
}

type ModifyDBInstanceMaintainTimeResponseBody struct {
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
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDBInstanceMaintainTimeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ModifyDBInstanceMaintainTimeResponse) SetBody(v *ModifyDBInstanceMaintainTimeResponseBody) *ModifyDBInstanceMaintainTimeResponse {
	s.Body = v
	return s
}

type ModifyDBInstanceNetworkTypeRequest struct {
	DBInstanceId        *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	InstanceNetworkType *string `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
	VPCId               *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	VSwitchId           *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	PrivateIpAddress    *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
}

func (s ModifyDBInstanceNetworkTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceNetworkTypeRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceNetworkTypeRequest) SetDBInstanceId(v string) *ModifyDBInstanceNetworkTypeRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceNetworkTypeRequest) SetInstanceNetworkType(v string) *ModifyDBInstanceNetworkTypeRequest {
	s.InstanceNetworkType = &v
	return s
}

func (s *ModifyDBInstanceNetworkTypeRequest) SetVPCId(v string) *ModifyDBInstanceNetworkTypeRequest {
	s.VPCId = &v
	return s
}

func (s *ModifyDBInstanceNetworkTypeRequest) SetVSwitchId(v string) *ModifyDBInstanceNetworkTypeRequest {
	s.VSwitchId = &v
	return s
}

func (s *ModifyDBInstanceNetworkTypeRequest) SetPrivateIpAddress(v string) *ModifyDBInstanceNetworkTypeRequest {
	s.PrivateIpAddress = &v
	return s
}

type ModifyDBInstanceNetworkTypeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBInstanceNetworkTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceNetworkTypeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceNetworkTypeResponseBody) SetRequestId(v string) *ModifyDBInstanceNetworkTypeResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDBInstanceNetworkTypeResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDBInstanceNetworkTypeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDBInstanceNetworkTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceNetworkTypeResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceNetworkTypeResponse) SetHeaders(v map[string]*string) *ModifyDBInstanceNetworkTypeResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBInstanceNetworkTypeResponse) SetBody(v *ModifyDBInstanceNetworkTypeResponseBody) *ModifyDBInstanceNetworkTypeResponse {
	s.Body = v
	return s
}

type ModifyDBInstanceSSLRequest struct {
	DBInstanceId     *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	SSLEnabled       *int32  `json:"SSLEnabled,omitempty" xml:"SSLEnabled,omitempty"`
}

func (s ModifyDBInstanceSSLRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceSSLRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceSSLRequest) SetDBInstanceId(v string) *ModifyDBInstanceSSLRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceSSLRequest) SetConnectionString(v string) *ModifyDBInstanceSSLRequest {
	s.ConnectionString = &v
	return s
}

func (s *ModifyDBInstanceSSLRequest) SetSSLEnabled(v int32) *ModifyDBInstanceSSLRequest {
	s.SSLEnabled = &v
	return s
}

type ModifyDBInstanceSSLResponseBody struct {
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
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDBInstanceSSLResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ModifyDBInstanceSSLResponse) SetBody(v *ModifyDBInstanceSSLResponseBody) *ModifyDBInstanceSSLResponse {
	s.Body = v
	return s
}

type ModifyParametersRequest struct {
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	Parameters           *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	ForceRestartInstance *bool   `json:"ForceRestartInstance,omitempty" xml:"ForceRestartInstance,omitempty"`
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

func (s *ModifyParametersRequest) SetParameters(v string) *ModifyParametersRequest {
	s.Parameters = &v
	return s
}

func (s *ModifyParametersRequest) SetForceRestartInstance(v bool) *ModifyParametersRequest {
	s.ForceRestartInstance = &v
	return s
}

type ModifyParametersResponseBody struct {
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
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyParametersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ModifyParametersResponse) SetBody(v *ModifyParametersResponseBody) *ModifyParametersResponse {
	s.Body = v
	return s
}

type ModifySecurityIpsRequest struct {
	DBInstanceId               *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	SecurityIPList             *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	DBInstanceIPArrayName      *string `json:"DBInstanceIPArrayName,omitempty" xml:"DBInstanceIPArrayName,omitempty"`
	DBInstanceIPArrayAttribute *string `json:"DBInstanceIPArrayAttribute,omitempty" xml:"DBInstanceIPArrayAttribute,omitempty"`
}

func (s ModifySecurityIpsRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifySecurityIpsRequest) GoString() string {
	return s.String()
}

func (s *ModifySecurityIpsRequest) SetDBInstanceId(v string) *ModifySecurityIpsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifySecurityIpsRequest) SetSecurityIPList(v string) *ModifySecurityIpsRequest {
	s.SecurityIPList = &v
	return s
}

func (s *ModifySecurityIpsRequest) SetDBInstanceIPArrayName(v string) *ModifySecurityIpsRequest {
	s.DBInstanceIPArrayName = &v
	return s
}

func (s *ModifySecurityIpsRequest) SetDBInstanceIPArrayAttribute(v string) *ModifySecurityIpsRequest {
	s.DBInstanceIPArrayAttribute = &v
	return s
}

type ModifySecurityIpsResponseBody struct {
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
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifySecurityIpsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ModifySecurityIpsResponse) SetBody(v *ModifySecurityIpsResponseBody) *ModifySecurityIpsResponse {
	s.Body = v
	return s
}

type ModifySQLCollectorPolicyRequest struct {
	DBInstanceId       *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
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
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifySQLCollectorPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ModifySQLCollectorPolicyResponse) SetBody(v *ModifySQLCollectorPolicyResponseBody) *ModifySQLCollectorPolicyResponse {
	s.Body = v
	return s
}

type ReleaseInstancePublicConnectionRequest struct {
	DBInstanceId            *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	CurrentConnectionString *string `json:"CurrentConnectionString,omitempty" xml:"CurrentConnectionString,omitempty"`
	AddressType             *string `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
}

func (s ReleaseInstancePublicConnectionRequest) String() string {
	return tea.Prettify(s)
}

func (s ReleaseInstancePublicConnectionRequest) GoString() string {
	return s.String()
}

func (s *ReleaseInstancePublicConnectionRequest) SetDBInstanceId(v string) *ReleaseInstancePublicConnectionRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ReleaseInstancePublicConnectionRequest) SetCurrentConnectionString(v string) *ReleaseInstancePublicConnectionRequest {
	s.CurrentConnectionString = &v
	return s
}

func (s *ReleaseInstancePublicConnectionRequest) SetAddressType(v string) *ReleaseInstancePublicConnectionRequest {
	s.AddressType = &v
	return s
}

type ReleaseInstancePublicConnectionResponseBody struct {
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
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ReleaseInstancePublicConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ReleaseInstancePublicConnectionResponse) SetBody(v *ReleaseInstancePublicConnectionResponseBody) *ReleaseInstancePublicConnectionResponse {
	s.Body = v
	return s
}

type ResetAccountPasswordRequest struct {
	DBInstanceId    *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	AccountName     *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	AccountPassword *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
}

func (s ResetAccountPasswordRequest) String() string {
	return tea.Prettify(s)
}

func (s ResetAccountPasswordRequest) GoString() string {
	return s.String()
}

func (s *ResetAccountPasswordRequest) SetDBInstanceId(v string) *ResetAccountPasswordRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ResetAccountPasswordRequest) SetAccountName(v string) *ResetAccountPasswordRequest {
	s.AccountName = &v
	return s
}

func (s *ResetAccountPasswordRequest) SetAccountPassword(v string) *ResetAccountPasswordRequest {
	s.AccountPassword = &v
	return s
}

type ResetAccountPasswordResponseBody struct {
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
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ResetAccountPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ResetAccountPasswordResponse) SetBody(v *ResetAccountPasswordResponseBody) *ResetAccountPasswordResponse {
	s.Body = v
	return s
}

type RestartDBInstanceRequest struct {
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
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
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RestartDBInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RestartDBInstanceResponse) SetBody(v *RestartDBInstanceResponseBody) *RestartDBInstanceResponse {
	s.Body = v
	return s
}

type SwitchDBInstanceNetTypeRequest struct {
	DBInstanceId           *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	ConnectionStringPrefix *string `json:"ConnectionStringPrefix,omitempty" xml:"ConnectionStringPrefix,omitempty"`
	Port                   *string `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s SwitchDBInstanceNetTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s SwitchDBInstanceNetTypeRequest) GoString() string {
	return s.String()
}

func (s *SwitchDBInstanceNetTypeRequest) SetDBInstanceId(v string) *SwitchDBInstanceNetTypeRequest {
	s.DBInstanceId = &v
	return s
}

func (s *SwitchDBInstanceNetTypeRequest) SetConnectionStringPrefix(v string) *SwitchDBInstanceNetTypeRequest {
	s.ConnectionStringPrefix = &v
	return s
}

func (s *SwitchDBInstanceNetTypeRequest) SetPort(v string) *SwitchDBInstanceNetTypeRequest {
	s.Port = &v
	return s
}

type SwitchDBInstanceNetTypeResponseBody struct {
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
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SwitchDBInstanceNetTypeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *SwitchDBInstanceNetTypeResponse) SetBody(v *SwitchDBInstanceNetTypeResponseBody) *SwitchDBInstanceNetTypeResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	OwnerId              *int64                    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string                   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string                   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	RegionId             *string                   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceType         *string                   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	ResourceId           []*string                 `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	Tag                  []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s TagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) SetOwnerId(v int64) *TagResourcesRequest {
	s.OwnerId = &v
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

func (s *TagResourcesRequest) SetOwnerAccount(v string) *TagResourcesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *TagResourcesRequest) SetRegionId(v string) *TagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceType(v string) *TagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesRequest) SetResourceId(v []*string) *TagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *TagResourcesRequest) SetTag(v []*TagResourcesRequestTag) *TagResourcesRequest {
	s.Tag = v
	return s
}

type TagResourcesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *TagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *TagResourcesResponse) SetBody(v *TagResourcesResponseBody) *TagResourcesResponse {
	s.Body = v
	return s
}

type UntagResourcesRequest struct {
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	RegionId             *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceType         *string   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	All                  *bool     `json:"All,omitempty" xml:"All,omitempty"`
	ResourceId           []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	TagKey               []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s UntagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesRequest) SetOwnerId(v int64) *UntagResourcesRequest {
	s.OwnerId = &v
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

func (s *UntagResourcesRequest) SetOwnerAccount(v string) *UntagResourcesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UntagResourcesRequest) SetRegionId(v string) *UntagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceType(v string) *UntagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesRequest) SetAll(v bool) *UntagResourcesRequest {
	s.All = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceId(v []*string) *UntagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *UntagResourcesRequest) SetTagKey(v []*string) *UntagResourcesRequest {
	s.TagKey = v
	return s
}

type UntagResourcesResponseBody struct {
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
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UntagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UntagResourcesResponse) SetBody(v *UntagResourcesResponseBody) *UntagResourcesResponse {
	s.Body = v
	return s
}

type UpgradeDBInstanceRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DBInstanceClass      *string `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty"`
	DBInstanceGroupCount *string `json:"DBInstanceGroupCount,omitempty" xml:"DBInstanceGroupCount,omitempty"`
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	PayType              *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
}

func (s UpgradeDBInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpgradeDBInstanceRequest) GoString() string {
	return s.String()
}

func (s *UpgradeDBInstanceRequest) SetOwnerId(v int64) *UpgradeDBInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *UpgradeDBInstanceRequest) SetRegionId(v string) *UpgradeDBInstanceRequest {
	s.RegionId = &v
	return s
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

func (s *UpgradeDBInstanceRequest) SetPayType(v string) *UpgradeDBInstanceRequest {
	s.PayType = &v
	return s
}

type UpgradeDBInstanceResponseBody struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	OrderId      *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
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

func (s *UpgradeDBInstanceResponseBody) SetRequestId(v string) *UpgradeDBInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeDBInstanceResponseBody) SetOrderId(v string) *UpgradeDBInstanceResponseBody {
	s.OrderId = &v
	return s
}

type UpgradeDBInstanceResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpgradeDBInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpgradeDBInstanceResponse) SetBody(v *UpgradeDBInstanceResponseBody) *UpgradeDBInstanceResponse {
	s.Body = v
	return s
}

type UpgradeDBVersionRequest struct {
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DBInstanceId   *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	MinorVersion   *string `json:"MinorVersion,omitempty" xml:"MinorVersion,omitempty"`
	MajorVersion   *string `json:"MajorVersion,omitempty" xml:"MajorVersion,omitempty"`
	SwitchTimeMode *string `json:"SwitchTimeMode,omitempty" xml:"SwitchTimeMode,omitempty"`
	SwitchTime     *string `json:"SwitchTime,omitempty" xml:"SwitchTime,omitempty"`
}

func (s UpgradeDBVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpgradeDBVersionRequest) GoString() string {
	return s.String()
}

func (s *UpgradeDBVersionRequest) SetOwnerId(v int64) *UpgradeDBVersionRequest {
	s.OwnerId = &v
	return s
}

func (s *UpgradeDBVersionRequest) SetRegionId(v string) *UpgradeDBVersionRequest {
	s.RegionId = &v
	return s
}

func (s *UpgradeDBVersionRequest) SetDBInstanceId(v string) *UpgradeDBVersionRequest {
	s.DBInstanceId = &v
	return s
}

func (s *UpgradeDBVersionRequest) SetMinorVersion(v string) *UpgradeDBVersionRequest {
	s.MinorVersion = &v
	return s
}

func (s *UpgradeDBVersionRequest) SetMajorVersion(v string) *UpgradeDBVersionRequest {
	s.MajorVersion = &v
	return s
}

func (s *UpgradeDBVersionRequest) SetSwitchTimeMode(v string) *UpgradeDBVersionRequest {
	s.SwitchTimeMode = &v
	return s
}

func (s *UpgradeDBVersionRequest) SetSwitchTime(v string) *UpgradeDBVersionRequest {
	s.SwitchTime = &v
	return s
}

type UpgradeDBVersionResponseBody struct {
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	DBInstanceId   *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId         *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s UpgradeDBVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpgradeDBVersionResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeDBVersionResponseBody) SetDBInstanceName(v string) *UpgradeDBVersionResponseBody {
	s.DBInstanceName = &v
	return s
}

func (s *UpgradeDBVersionResponseBody) SetDBInstanceId(v string) *UpgradeDBVersionResponseBody {
	s.DBInstanceId = &v
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
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpgradeDBVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpgradeDBVersionResponse) SetBody(v *UpgradeDBVersionResponseBody) *UpgradeDBVersionResponse {
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

func (client *Client) AddBuDBInstanceRelationWithOptions(request *AddBuDBInstanceRelationRequest, runtime *util.RuntimeOptions) (_result *AddBuDBInstanceRelationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddBuDBInstanceRelationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddBuDBInstanceRelation"), tea.String("2016-05-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddBuDBInstanceRelation(request *AddBuDBInstanceRelationRequest) (_result *AddBuDBInstanceRelationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddBuDBInstanceRelationResponse{}
	_body, _err := client.AddBuDBInstanceRelationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AllocateInstancePublicConnectionWithOptions(request *AllocateInstancePublicConnectionRequest, runtime *util.RuntimeOptions) (_result *AllocateInstancePublicConnectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AllocateInstancePublicConnectionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AllocateInstancePublicConnection"), tea.String("2016-05-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CheckServiceLinkedRoleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CheckServiceLinkedRole"), tea.String("2016-05-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) CreateAccountWithOptions(request *CreateAccountRequest, runtime *util.RuntimeOptions) (_result *CreateAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateAccountResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateAccount"), tea.String("2016-05-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) CreateDBInstanceWithOptions(request *CreateDBInstanceRequest, runtime *util.RuntimeOptions) (_result *CreateDBInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateDBInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateDBInstance"), tea.String("2016-05-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) CreateECSDBInstanceWithOptions(request *CreateECSDBInstanceRequest, runtime *util.RuntimeOptions) (_result *CreateECSDBInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateECSDBInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateECSDBInstance"), tea.String("2016-05-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateECSDBInstance(request *CreateECSDBInstanceRequest) (_result *CreateECSDBInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateECSDBInstanceResponse{}
	_body, _err := client.CreateECSDBInstanceWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateServiceLinkedRoleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateServiceLinkedRole"), tea.String("2016-05-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DeleteDatabaseWithOptions(request *DeleteDatabaseRequest, runtime *util.RuntimeOptions) (_result *DeleteDatabaseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteDatabaseResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteDatabase"), tea.String("2016-05-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDatabase(request *DeleteDatabaseRequest) (_result *DeleteDatabaseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDatabaseResponse{}
	_body, _err := client.DeleteDatabaseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDBInstanceWithOptions(request *DeleteDBInstanceRequest, runtime *util.RuntimeOptions) (_result *DeleteDBInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteDBInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteDBInstance"), tea.String("2016-05-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) DescribeAccountsWithOptions(request *DescribeAccountsRequest, runtime *util.RuntimeOptions) (_result *DescribeAccountsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAccountsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAccounts"), tea.String("2016-05-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) DescribeAvailableResourcesWithOptions(request *DescribeAvailableResourcesRequest, runtime *util.RuntimeOptions) (_result *DescribeAvailableResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAvailableResourcesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAvailableResources"), tea.String("2016-05-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) DescribeBackupPolicyWithOptions(request *DescribeBackupPolicyRequest, runtime *util.RuntimeOptions) (_result *DescribeBackupPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeBackupPolicyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeBackupPolicy"), tea.String("2016-05-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) DescribeDataBackupsWithOptions(request *DescribeDataBackupsRequest, runtime *util.RuntimeOptions) (_result *DescribeDataBackupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDataBackupsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDataBackups"), tea.String("2016-05-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeDBClusterPerformanceWithOptions(request *DescribeDBClusterPerformanceRequest, runtime *util.RuntimeOptions) (_result *DescribeDBClusterPerformanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDBClusterPerformanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDBClusterPerformance"), tea.String("2016-05-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) DescribeDBInstanceAttributeWithOptions(request *DescribeDBInstanceAttributeRequest, runtime *util.RuntimeOptions) (_result *DescribeDBInstanceAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDBInstanceAttributeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDBInstanceAttribute"), tea.String("2016-05-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) DescribeDBInstanceIPArrayListWithOptions(request *DescribeDBInstanceIPArrayListRequest, runtime *util.RuntimeOptions) (_result *DescribeDBInstanceIPArrayListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDBInstanceIPArrayListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDBInstanceIPArrayList"), tea.String("2016-05-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) DescribeDBInstanceNetInfoWithOptions(request *DescribeDBInstanceNetInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeDBInstanceNetInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDBInstanceNetInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDBInstanceNetInfo"), tea.String("2016-05-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeDBInstanceOnECSAttributeWithOptions(request *DescribeDBInstanceOnECSAttributeRequest, runtime *util.RuntimeOptions) (_result *DescribeDBInstanceOnECSAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDBInstanceOnECSAttributeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDBInstanceOnECSAttribute"), tea.String("2016-05-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBInstanceOnECSAttribute(request *DescribeDBInstanceOnECSAttributeRequest) (_result *DescribeDBInstanceOnECSAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBInstanceOnECSAttributeResponse{}
	_body, _err := client.DescribeDBInstanceOnECSAttributeWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDBInstancePerformanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDBInstancePerformance"), tea.String("2016-05-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeDBInstancesWithOptions(tmpReq *DescribeDBInstancesRequest, runtime *util.RuntimeOptions) (_result *DescribeDBInstancesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DescribeDBInstancesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.InstanceDeployTypes)) {
		request.InstanceDeployTypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InstanceDeployTypes, tea.String("InstanceDeployTypes"), tea.String("simple"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.DBInstanceStatuses)) {
		request.DBInstanceStatusesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DBInstanceStatuses, tea.String("DBInstanceStatuses"), tea.String("simple"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.DBInstanceCategories)) {
		request.DBInstanceCategoriesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DBInstanceCategories, tea.String("DBInstanceCategories"), tea.String("simple"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDBInstancesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDBInstances"), tea.String("2016-05-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) DescribeDBInstanceSQLPatternsWithOptions(request *DescribeDBInstanceSQLPatternsRequest, runtime *util.RuntimeOptions) (_result *DescribeDBInstanceSQLPatternsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDBInstanceSQLPatternsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDBInstanceSQLPatterns"), tea.String("2016-05-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBInstanceSQLPatterns(request *DescribeDBInstanceSQLPatternsRequest) (_result *DescribeDBInstanceSQLPatternsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBInstanceSQLPatternsResponse{}
	_body, _err := client.DescribeDBInstanceSQLPatternsWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDBInstanceSSLResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDBInstanceSSL"), tea.String("2016-05-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeLogBackupsWithOptions(request *DescribeLogBackupsRequest, runtime *util.RuntimeOptions) (_result *DescribeLogBackupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeLogBackupsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeLogBackups"), tea.String("2016-05-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeModifyParameterLogResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeModifyParameterLog"), tea.String("2016-05-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeParametersWithOptions(request *DescribeParametersRequest, runtime *util.RuntimeOptions) (_result *DescribeParametersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeParametersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeParameters"), tea.String("2016-05-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) DescribeRdsVpcsWithOptions(request *DescribeRdsVpcsRequest, runtime *util.RuntimeOptions) (_result *DescribeRdsVpcsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeRdsVpcsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRdsVpcs"), tea.String("2016-05-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) DescribeRdsVSwitchsWithOptions(request *DescribeRdsVSwitchsRequest, runtime *util.RuntimeOptions) (_result *DescribeRdsVSwitchsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeRdsVSwitchsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRdsVSwitchs"), tea.String("2016-05-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) DescribeRegionsWithOptions(request *DescribeRegionsRequest, runtime *util.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRegions"), tea.String("2016-05-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) DescribeResourceUsageWithOptions(request *DescribeResourceUsageRequest, runtime *util.RuntimeOptions) (_result *DescribeResourceUsageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeResourceUsageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeResourceUsage"), tea.String("2016-05-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeResourceUsage(request *DescribeResourceUsageRequest) (_result *DescribeResourceUsageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeResourceUsageResponse{}
	_body, _err := client.DescribeResourceUsageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSlowLogRecordsWithOptions(request *DescribeSlowLogRecordsRequest, runtime *util.RuntimeOptions) (_result *DescribeSlowLogRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSlowLogRecordsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSlowLogRecords"), tea.String("2016-05-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSlowLogRecords(request *DescribeSlowLogRecordsRequest) (_result *DescribeSlowLogRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSlowLogRecordsResponse{}
	_body, _err := client.DescribeSlowLogRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSlowSQLLogsWithOptions(request *DescribeSlowSQLLogsRequest, runtime *util.RuntimeOptions) (_result *DescribeSlowSQLLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSlowSQLLogsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSlowSQLLogs"), tea.String("2016-05-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSlowSQLLogs(request *DescribeSlowSQLLogsRequest) (_result *DescribeSlowSQLLogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSlowSQLLogsResponse{}
	_body, _err := client.DescribeSlowSQLLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSpecificationWithOptions(request *DescribeSpecificationRequest, runtime *util.RuntimeOptions) (_result *DescribeSpecificationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSpecificationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSpecification"), tea.String("2016-05-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSpecification(request *DescribeSpecificationRequest) (_result *DescribeSpecificationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSpecificationResponse{}
	_body, _err := client.DescribeSpecificationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSQLCollectorPolicyWithOptions(request *DescribeSQLCollectorPolicyRequest, runtime *util.RuntimeOptions) (_result *DescribeSQLCollectorPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSQLCollectorPolicyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSQLCollectorPolicy"), tea.String("2016-05-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSQLCollectorPolicy(request *DescribeSQLCollectorPolicyRequest) (_result *DescribeSQLCollectorPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSQLCollectorPolicyResponse{}
	_body, _err := client.DescribeSQLCollectorPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSQLLogByQueryIdWithOptions(request *DescribeSQLLogByQueryIdRequest, runtime *util.RuntimeOptions) (_result *DescribeSQLLogByQueryIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSQLLogByQueryIdResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSQLLogByQueryId"), tea.String("2016-05-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSQLLogByQueryId(request *DescribeSQLLogByQueryIdRequest) (_result *DescribeSQLLogByQueryIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSQLLogByQueryIdResponse{}
	_body, _err := client.DescribeSQLLogByQueryIdWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSQLLogCountResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSQLLogCount"), tea.String("2016-05-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeSQLLogFilesWithOptions(request *DescribeSQLLogFilesRequest, runtime *util.RuntimeOptions) (_result *DescribeSQLLogFilesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSQLLogFilesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSQLLogFiles"), tea.String("2016-05-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSQLLogFiles(request *DescribeSQLLogFilesRequest) (_result *DescribeSQLLogFilesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSQLLogFilesResponse{}
	_body, _err := client.DescribeSQLLogFilesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSQLLogRecordsWithOptions(request *DescribeSQLLogRecordsRequest, runtime *util.RuntimeOptions) (_result *DescribeSQLLogRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSQLLogRecordsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSQLLogRecords"), tea.String("2016-05-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSQLLogRecords(request *DescribeSQLLogRecordsRequest) (_result *DescribeSQLLogRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSQLLogRecordsResponse{}
	_body, _err := client.DescribeSQLLogRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSQLLogsWithOptions(request *DescribeSQLLogsRequest, runtime *util.RuntimeOptions) (_result *DescribeSQLLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSQLLogsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSQLLogs"), tea.String("2016-05-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSQLLogs(request *DescribeSQLLogsRequest) (_result *DescribeSQLLogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSQLLogsResponse{}
	_body, _err := client.DescribeSQLLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSQLLogsOnSliceWithOptions(request *DescribeSQLLogsOnSliceRequest, runtime *util.RuntimeOptions) (_result *DescribeSQLLogsOnSliceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSQLLogsOnSliceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSQLLogsOnSlice"), tea.String("2016-05-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSQLLogsOnSlice(request *DescribeSQLLogsOnSliceRequest) (_result *DescribeSQLLogsOnSliceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSQLLogsOnSliceResponse{}
	_body, _err := client.DescribeSQLLogsOnSliceWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeTagsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeTags"), tea.String("2016-05-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeUserEncryptionKeyListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeUserEncryptionKeyList"), tea.String("2016-05-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, runtime *util.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListTagResources"), tea.String("2016-05-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyAccountDescriptionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyAccountDescription"), tea.String("2016-05-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyBackupPolicyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyBackupPolicy"), tea.String("2016-05-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ModifyDBInstanceConnectionModeWithOptions(request *ModifyDBInstanceConnectionModeRequest, runtime *util.RuntimeOptions) (_result *ModifyDBInstanceConnectionModeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyDBInstanceConnectionModeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyDBInstanceConnectionMode"), tea.String("2016-05-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDBInstanceConnectionMode(request *ModifyDBInstanceConnectionModeRequest) (_result *ModifyDBInstanceConnectionModeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDBInstanceConnectionModeResponse{}
	_body, _err := client.ModifyDBInstanceConnectionModeWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyDBInstanceConnectionStringResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyDBInstanceConnectionString"), tea.String("2016-05-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ModifyDBInstanceDescriptionWithOptions(request *ModifyDBInstanceDescriptionRequest, runtime *util.RuntimeOptions) (_result *ModifyDBInstanceDescriptionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyDBInstanceDescriptionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyDBInstanceDescription"), tea.String("2016-05-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) ModifyDBInstanceMaintainTimeWithOptions(request *ModifyDBInstanceMaintainTimeRequest, runtime *util.RuntimeOptions) (_result *ModifyDBInstanceMaintainTimeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyDBInstanceMaintainTimeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyDBInstanceMaintainTime"), tea.String("2016-05-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) ModifyDBInstanceNetworkTypeWithOptions(request *ModifyDBInstanceNetworkTypeRequest, runtime *util.RuntimeOptions) (_result *ModifyDBInstanceNetworkTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyDBInstanceNetworkTypeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyDBInstanceNetworkType"), tea.String("2016-05-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDBInstanceNetworkType(request *ModifyDBInstanceNetworkTypeRequest) (_result *ModifyDBInstanceNetworkTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDBInstanceNetworkTypeResponse{}
	_body, _err := client.ModifyDBInstanceNetworkTypeWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyDBInstanceSSLResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyDBInstanceSSL"), tea.String("2016-05-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ModifyParametersWithOptions(request *ModifyParametersRequest, runtime *util.RuntimeOptions) (_result *ModifyParametersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyParametersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyParameters"), tea.String("2016-05-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) ModifySecurityIpsWithOptions(request *ModifySecurityIpsRequest, runtime *util.RuntimeOptions) (_result *ModifySecurityIpsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifySecurityIpsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifySecurityIps"), tea.String("2016-05-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) ModifySQLCollectorPolicyWithOptions(request *ModifySQLCollectorPolicyRequest, runtime *util.RuntimeOptions) (_result *ModifySQLCollectorPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifySQLCollectorPolicyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifySQLCollectorPolicy"), tea.String("2016-05-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ReleaseInstancePublicConnectionWithOptions(request *ReleaseInstancePublicConnectionRequest, runtime *util.RuntimeOptions) (_result *ReleaseInstancePublicConnectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ReleaseInstancePublicConnectionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ReleaseInstancePublicConnection"), tea.String("2016-05-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ResetAccountPasswordResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ResetAccountPassword"), tea.String("2016-05-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) RestartDBInstanceWithOptions(request *RestartDBInstanceRequest, runtime *util.RuntimeOptions) (_result *RestartDBInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RestartDBInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RestartDBInstance"), tea.String("2016-05-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) SwitchDBInstanceNetTypeWithOptions(request *SwitchDBInstanceNetTypeRequest, runtime *util.RuntimeOptions) (_result *SwitchDBInstanceNetTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SwitchDBInstanceNetTypeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SwitchDBInstanceNetType"), tea.String("2016-05-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("TagResources"), tea.String("2016-05-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) UntagResourcesWithOptions(request *UntagResourcesRequest, runtime *util.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UntagResources"), tea.String("2016-05-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) UpgradeDBInstanceWithOptions(request *UpgradeDBInstanceRequest, runtime *util.RuntimeOptions) (_result *UpgradeDBInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpgradeDBInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpgradeDBInstance"), tea.String("2016-05-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpgradeDBVersionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpgradeDBVersion"), tea.String("2016-05-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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