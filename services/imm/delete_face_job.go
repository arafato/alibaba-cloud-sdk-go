package imm

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

// DeleteFaceJob invokes the imm.DeleteFaceJob API synchronously
// api document: https://help.aliyun.com/api/imm/deletefacejob.html
func (client *Client) DeleteFaceJob(request *DeleteFaceJobRequest) (response *DeleteFaceJobResponse, err error) {
	response = CreateDeleteFaceJobResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteFaceJobWithChan invokes the imm.DeleteFaceJob API asynchronously
// api document: https://help.aliyun.com/api/imm/deletefacejob.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteFaceJobWithChan(request *DeleteFaceJobRequest) (<-chan *DeleteFaceJobResponse, <-chan error) {
	responseChan := make(chan *DeleteFaceJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteFaceJob(request)
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

// DeleteFaceJobWithCallback invokes the imm.DeleteFaceJob API asynchronously
// api document: https://help.aliyun.com/api/imm/deletefacejob.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteFaceJobWithCallback(request *DeleteFaceJobRequest, callback func(response *DeleteFaceJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteFaceJobResponse
		var err error
		defer close(result)
		response, err = client.DeleteFaceJob(request)
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

// DeleteFaceJobRequest is the request struct for api DeleteFaceJob
type DeleteFaceJobRequest struct {
	*requests.RpcRequest
	JobId          string `position:"Query" name:"JobId"`
	Project        string `position:"Query" name:"Project"`
	ClearIndexData string `position:"Query" name:"ClearIndexData"`
}

// DeleteFaceJobResponse is the response struct for api DeleteFaceJob
type DeleteFaceJobResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteFaceJobRequest creates a request to invoke DeleteFaceJob API
func CreateDeleteFaceJobRequest() (request *DeleteFaceJobRequest) {
	request = &DeleteFaceJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imm", "2017-09-06", "DeleteFaceJob", "2017-09-06", "openAPI")
	return
}

// CreateDeleteFaceJobResponse creates a response to parse from DeleteFaceJob response
func CreateDeleteFaceJobResponse() (response *DeleteFaceJobResponse) {
	response = &DeleteFaceJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
