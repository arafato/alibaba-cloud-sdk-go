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

// CloneParameterGroup invokes the rds.CloneParameterGroup API synchronously
// api document: https://help.aliyun.com/api/rds/cloneparametergroup.html
func (client *Client) CloneParameterGroup(request *CloneParameterGroupRequest) (response *CloneParameterGroupResponse, err error) {
	response = CreateCloneParameterGroupResponse()
	err = client.DoAction(request, response)
	return
}

// CloneParameterGroupWithChan invokes the rds.CloneParameterGroup API asynchronously
// api document: https://help.aliyun.com/api/rds/cloneparametergroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CloneParameterGroupWithChan(request *CloneParameterGroupRequest) (<-chan *CloneParameterGroupResponse, <-chan error) {
	responseChan := make(chan *CloneParameterGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CloneParameterGroup(request)
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

// CloneParameterGroupWithCallback invokes the rds.CloneParameterGroup API asynchronously
// api document: https://help.aliyun.com/api/rds/cloneparametergroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CloneParameterGroupWithCallback(request *CloneParameterGroupRequest, callback func(response *CloneParameterGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CloneParameterGroupResponse
		var err error
		defer close(result)
		response, err = client.CloneParameterGroup(request)
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

// CloneParameterGroupRequest is the request struct for api CloneParameterGroup
type CloneParameterGroupRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ParameterGroupId     string           `position:"Query" name:"ParameterGroupId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	TargetRegionId       string           `position:"Query" name:"TargetRegionId"`
	ParameterGroupName   string           `position:"Query" name:"ParameterGroupName"`
	ParameterGroupDesc   string           `position:"Query" name:"ParameterGroupDesc"`
}

// CloneParameterGroupResponse is the response struct for api CloneParameterGroup
type CloneParameterGroupResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCloneParameterGroupRequest creates a request to invoke CloneParameterGroup API
func CreateCloneParameterGroupRequest() (request *CloneParameterGroupRequest) {
	request = &CloneParameterGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "CloneParameterGroup", "Rds", "openAPI")
	return
}

// CreateCloneParameterGroupResponse creates a response to parse from CloneParameterGroup response
func CreateCloneParameterGroupResponse() (response *CloneParameterGroupResponse) {
	response = &CloneParameterGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
