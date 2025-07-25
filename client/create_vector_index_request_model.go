// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVectorIndexRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollection(v string) *CreateVectorIndexRequest
	GetCollection() *string
	SetDBInstanceId(v string) *CreateVectorIndexRequest
	GetDBInstanceId() *string
	SetDimension(v int32) *CreateVectorIndexRequest
	GetDimension() *int32
	SetExternalStorage(v int32) *CreateVectorIndexRequest
	GetExternalStorage() *int32
	SetHnswEfConstruction(v int32) *CreateVectorIndexRequest
	GetHnswEfConstruction() *int32
	SetHnswM(v int32) *CreateVectorIndexRequest
	GetHnswM() *int32
	SetManagerAccount(v string) *CreateVectorIndexRequest
	GetManagerAccount() *string
	SetManagerAccountPassword(v string) *CreateVectorIndexRequest
	GetManagerAccountPassword() *string
	SetMetrics(v string) *CreateVectorIndexRequest
	GetMetrics() *string
	SetNamespace(v string) *CreateVectorIndexRequest
	GetNamespace() *string
	SetOwnerId(v int64) *CreateVectorIndexRequest
	GetOwnerId() *int64
	SetPqEnable(v int32) *CreateVectorIndexRequest
	GetPqEnable() *int32
	SetRegionId(v string) *CreateVectorIndexRequest
	GetRegionId() *string
	SetType(v string) *CreateVectorIndexRequest
	GetType() *string
}

type CreateVectorIndexRequest struct {
	// Collection name.
	//
	// > You can use the [ListCollections](https://help.aliyun.com/document_detail/2401503.html) API to view the list.
	//
	// This parameter is required.
	//
	// example:
	//
	// document
	Collection *string `json:"Collection,omitempty" xml:"Collection,omitempty"`
	// Instance ID.
	//
	// > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) API to view details of all AnalyticDB PostgreSQL instances in the target region, including the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// Vector dimension.
	//
	// > This value must be consistent with the length of the vector data (Rows. Vector) uploaded via the [UpsertCollectionData](https://help.aliyun.com/document_detail/2401493.html) API.
	//
	// example:
	//
	// 1024
	Dimension *int32 `json:"Dimension,omitempty" xml:"Dimension,omitempty"`
	// Whether to use mmap to build the HNSW index, default is 0. If the data does not need to be deleted and there are performance requirements for uploading data, it is recommended to set this to 1.
	//
	// >
	//
	// > - When set to 0, the segment-page storage mode is used to build the index, which can use the shared_buffer in PostgreSQL for caching and supports deletion and update operations.
	//
	// > - When set to 1, the index is built using mmap, which does not support deletion and update operations.
	//
	// example:
	//
	// 0
	ExternalStorage    *int32 `json:"ExternalStorage,omitempty" xml:"ExternalStorage,omitempty"`
	HnswEfConstruction *int32 `json:"HnswEfConstruction,omitempty" xml:"HnswEfConstruction,omitempty"`
	// The maximum number of neighbors in the HNSW algorithm, ranging from 1 to 1000. The API will automatically set this value based on the vector dimension, and it generally does not need to be manually set.
	//
	// > It is suggested to set this based on the vector dimension as follows:
	//
	// > - Less than or equal to 384: 16
	//
	// > - Greater than 384 and less than or equal to 768: 32
	//
	// > - Greater than 768 and less than or equal to 1024: 64
	//
	// > - Greater than 1024: 128
	//
	// example:
	//
	// 64
	HnswM *int32 `json:"HnswM,omitempty" xml:"HnswM,omitempty"`
	// Name of the management account with rds_superuser permissions.
	//
	// > You can create an account through the console -> Account Management, or by using the [CreateAccount](https://help.aliyun.com/document_detail/2361789.html) API.
	//
	// This parameter is required.
	//
	// example:
	//
	// testaccount
	ManagerAccount *string `json:"ManagerAccount,omitempty" xml:"ManagerAccount,omitempty"`
	// Management account password.
	//
	// This parameter is required.
	//
	// example:
	//
	// testpassword
	ManagerAccountPassword *string `json:"ManagerAccountPassword,omitempty" xml:"ManagerAccountPassword,omitempty"`
	// Method used for building the vector index. Value description:
	//
	// - l2: Euclidean distance.
	//
	// - ip: Inner product (dot product) distance.
	//
	// - cosine: Cosine similarity.
	//
	// example:
	//
	// cosine
	Metrics *string `json:"Metrics,omitempty" xml:"Metrics,omitempty"`
	// Namespace, default is public.
	//
	// > You can use the [ListNamespaces](https://help.aliyun.com/document_detail/2401502.html) API to view the list.
	//
	// example:
	//
	// mynamespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Whether to enable PQ (Product Quantization) algorithm acceleration for the index. It is recommended to enable this when the data volume exceeds 500,000. Value description:
	//
	// - 0: Disabled.
	//
	// - 1: Enabled (default).
	//
	// example:
	//
	// 1
	PqEnable *int32 `json:"PqEnable,omitempty" xml:"PqEnable,omitempty"`
	// Region ID where the instance is located.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateVectorIndexRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVectorIndexRequest) GoString() string {
	return s.String()
}

func (s *CreateVectorIndexRequest) GetCollection() *string {
	return s.Collection
}

func (s *CreateVectorIndexRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CreateVectorIndexRequest) GetDimension() *int32 {
	return s.Dimension
}

func (s *CreateVectorIndexRequest) GetExternalStorage() *int32 {
	return s.ExternalStorage
}

func (s *CreateVectorIndexRequest) GetHnswEfConstruction() *int32 {
	return s.HnswEfConstruction
}

func (s *CreateVectorIndexRequest) GetHnswM() *int32 {
	return s.HnswM
}

func (s *CreateVectorIndexRequest) GetManagerAccount() *string {
	return s.ManagerAccount
}

func (s *CreateVectorIndexRequest) GetManagerAccountPassword() *string {
	return s.ManagerAccountPassword
}

func (s *CreateVectorIndexRequest) GetMetrics() *string {
	return s.Metrics
}

func (s *CreateVectorIndexRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *CreateVectorIndexRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateVectorIndexRequest) GetPqEnable() *int32 {
	return s.PqEnable
}

func (s *CreateVectorIndexRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateVectorIndexRequest) GetType() *string {
	return s.Type
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

func (s *CreateVectorIndexRequest) SetExternalStorage(v int32) *CreateVectorIndexRequest {
	s.ExternalStorage = &v
	return s
}

func (s *CreateVectorIndexRequest) SetHnswEfConstruction(v int32) *CreateVectorIndexRequest {
	s.HnswEfConstruction = &v
	return s
}

func (s *CreateVectorIndexRequest) SetHnswM(v int32) *CreateVectorIndexRequest {
	s.HnswM = &v
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

func (s *CreateVectorIndexRequest) SetPqEnable(v int32) *CreateVectorIndexRequest {
	s.PqEnable = &v
	return s
}

func (s *CreateVectorIndexRequest) SetRegionId(v string) *CreateVectorIndexRequest {
	s.RegionId = &v
	return s
}

func (s *CreateVectorIndexRequest) SetType(v string) *CreateVectorIndexRequest {
	s.Type = &v
	return s
}

func (s *CreateVectorIndexRequest) Validate() error {
	return dara.Validate(s)
}
