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

// ListOfficeConversionTask invokes the imm.ListOfficeConversionTask API synchronously
// api document: https://help.aliyun.com/api/imm/listofficeconversiontask.html
func (client *Client) ListOfficeConversionTask(request *ListOfficeConversionTaskRequest) (response *ListOfficeConversionTaskResponse, err error) {
	response = CreateListOfficeConversionTaskResponse()
	err = client.DoAction(request, response)
	return
}

// ListOfficeConversionTaskWithChan invokes the imm.ListOfficeConversionTask API asynchronously
// api document: https://help.aliyun.com/api/imm/listofficeconversiontask.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListOfficeConversionTaskWithChan(request *ListOfficeConversionTaskRequest) (<-chan *ListOfficeConversionTaskResponse, <-chan error) {
	responseChan := make(chan *ListOfficeConversionTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListOfficeConversionTask(request)
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

// ListOfficeConversionTaskWithCallback invokes the imm.ListOfficeConversionTask API asynchronously
// api document: https://help.aliyun.com/api/imm/listofficeconversiontask.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListOfficeConversionTaskWithCallback(request *ListOfficeConversionTaskRequest, callback func(response *ListOfficeConversionTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListOfficeConversionTaskResponse
		var err error
		defer close(result)
		response, err = client.ListOfficeConversionTask(request)
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

// ListOfficeConversionTaskRequest is the request struct for api ListOfficeConversionTask
type ListOfficeConversionTaskRequest struct {
	*requests.RpcRequest
	MaxKeys requests.Integer `position:"Query" name:"MaxKeys"`
	Marker  string           `position:"Query" name:"Marker"`
	Project string           `position:"Query" name:"Project"`
}

// ListOfficeConversionTaskResponse is the response struct for api ListOfficeConversionTask
type ListOfficeConversionTaskResponse struct {
	*responses.BaseResponse
	RequestId  string      `json:"RequestId" xml:"RequestId"`
	NextMarker string      `json:"NextMarker" xml:"NextMarker"`
	Tasks      []TasksItem `json:"Tasks" xml:"Tasks"`
}

// CreateListOfficeConversionTaskRequest creates a request to invoke ListOfficeConversionTask API
func CreateListOfficeConversionTaskRequest() (request *ListOfficeConversionTaskRequest) {
	request = &ListOfficeConversionTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imm", "2017-09-06", "ListOfficeConversionTask", "2017-09-06", "openAPI")
	return
}

// CreateListOfficeConversionTaskResponse creates a response to parse from ListOfficeConversionTask response
func CreateListOfficeConversionTaskResponse() (response *ListOfficeConversionTaskResponse) {
	response = &ListOfficeConversionTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
