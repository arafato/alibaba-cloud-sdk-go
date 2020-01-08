package rds

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DescribeResourceUsage invokes the rds.DescribeResourceUsage API synchronously
// api document: https://help.aliyun.com/api/rds/describeresourceusage.html
func (client *Client) DescribeResourceUsage(request *DescribeResourceUsageRequest) (response *DescribeResourceUsageResponse, err error) {
	response = CreateDescribeResourceUsageResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeResourceUsageWithChan invokes the rds.DescribeResourceUsage API asynchronously
// api document: https://help.aliyun.com/api/rds/describeresourceusage.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeResourceUsageWithChan(request *DescribeResourceUsageRequest) (<-chan *DescribeResourceUsageResponse, <-chan error) {
	responseChan := make(chan *DescribeResourceUsageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeResourceUsage(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// DescribeResourceUsageWithCallback invokes the rds.DescribeResourceUsage API asynchronously
// api document: https://help.aliyun.com/api/rds/describeresourceusage.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeResourceUsageWithCallback(request *DescribeResourceUsageRequest, callback func(response *DescribeResourceUsageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeResourceUsageResponse
		var err error
		defer close(result)
		response, err = client.DescribeResourceUsage(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// DescribeResourceUsageRequest is the request struct for api DescribeResourceUsage
type DescribeResourceUsageRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
}

// DescribeResourceUsageResponse is the response struct for api DescribeResourceUsage
type DescribeResourceUsageResponse struct {
	*responses.BaseResponse
	RequestId         string `json:"RequestId" xml:"RequestId"`
	DBInstanceId      string `json:"DBInstanceId" xml:"DBInstanceId"`
	Engine            string `json:"Engine" xml:"Engine"`
	DiskUsed          int64  `json:"DiskUsed" xml:"DiskUsed"`
	DataSize          int64  `json:"DataSize" xml:"DataSize"`
	LogSize           int64  `json:"LogSize" xml:"LogSize"`
	BackupSize        int64  `json:"BackupSize" xml:"BackupSize"`
	BackupOssDataSize int64  `json:"BackupOssDataSize" xml:"BackupOssDataSize"`
	BackupOssLogSize  int64  `json:"BackupOssLogSize" xml:"BackupOssLogSize"`
	SQLSize           int64  `json:"SQLSize" xml:"SQLSize"`
	ColdBackupSize    int64  `json:"ColdBackupSize" xml:"ColdBackupSize"`
}

// CreateDescribeResourceUsageRequest creates a request to invoke DescribeResourceUsage API
func CreateDescribeResourceUsageRequest() (request *DescribeResourceUsageRequest) {
	request = &DescribeResourceUsageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeResourceUsage", "Rds", "openAPI")
	return
}

// CreateDescribeResourceUsageResponse creates a response to parse from DescribeResourceUsage response
func CreateDescribeResourceUsageResponse() (response *DescribeResourceUsageResponse) {
	response = &DescribeResourceUsageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
