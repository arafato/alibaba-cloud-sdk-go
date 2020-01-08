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

// DescribeDBInstanceHAConfig invokes the rds.DescribeDBInstanceHAConfig API synchronously
// api document: https://help.aliyun.com/api/rds/describedbinstancehaconfig.html
func (client *Client) DescribeDBInstanceHAConfig(request *DescribeDBInstanceHAConfigRequest) (response *DescribeDBInstanceHAConfigResponse, err error) {
	response = CreateDescribeDBInstanceHAConfigResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDBInstanceHAConfigWithChan invokes the rds.DescribeDBInstanceHAConfig API asynchronously
// api document: https://help.aliyun.com/api/rds/describedbinstancehaconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDBInstanceHAConfigWithChan(request *DescribeDBInstanceHAConfigRequest) (<-chan *DescribeDBInstanceHAConfigResponse, <-chan error) {
	responseChan := make(chan *DescribeDBInstanceHAConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDBInstanceHAConfig(request)
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

// DescribeDBInstanceHAConfigWithCallback invokes the rds.DescribeDBInstanceHAConfig API asynchronously
// api document: https://help.aliyun.com/api/rds/describedbinstancehaconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDBInstanceHAConfigWithCallback(request *DescribeDBInstanceHAConfigRequest, callback func(response *DescribeDBInstanceHAConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDBInstanceHAConfigResponse
		var err error
		defer close(result)
		response, err = client.DescribeDBInstanceHAConfig(request)
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

// DescribeDBInstanceHAConfigRequest is the request struct for api DescribeDBInstanceHAConfig
type DescribeDBInstanceHAConfigRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
}

// DescribeDBInstanceHAConfigResponse is the response struct for api DescribeDBInstanceHAConfig
type DescribeDBInstanceHAConfigResponse struct {
	*responses.BaseResponse
	RequestId         string            `json:"RequestId" xml:"RequestId"`
	DBInstanceId      string            `json:"DBInstanceId" xml:"DBInstanceId"`
	SyncMode          string            `json:"SyncMode" xml:"SyncMode"`
	HAMode            string            `json:"HAMode" xml:"HAMode"`
	HostInstanceInfos HostInstanceInfos `json:"HostInstanceInfos" xml:"HostInstanceInfos"`
}

// CreateDescribeDBInstanceHAConfigRequest creates a request to invoke DescribeDBInstanceHAConfig API
func CreateDescribeDBInstanceHAConfigRequest() (request *DescribeDBInstanceHAConfigRequest) {
	request = &DescribeDBInstanceHAConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeDBInstanceHAConfig", "Rds", "openAPI")
	return
}

// CreateDescribeDBInstanceHAConfigResponse creates a response to parse from DescribeDBInstanceHAConfig response
func CreateDescribeDBInstanceHAConfigResponse() (response *DescribeDBInstanceHAConfigResponse) {
	response = &DescribeDBInstanceHAConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
