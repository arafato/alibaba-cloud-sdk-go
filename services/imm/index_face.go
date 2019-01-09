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

// IndexFace invokes the imm.IndexFace API synchronously
// api document: https://help.aliyun.com/api/imm/indexface.html
func (client *Client) IndexFace(request *IndexFaceRequest) (response *IndexFaceResponse, err error) {
	response = CreateIndexFaceResponse()
	err = client.DoAction(request, response)
	return
}

// IndexFaceWithChan invokes the imm.IndexFace API asynchronously
// api document: https://help.aliyun.com/api/imm/indexface.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) IndexFaceWithChan(request *IndexFaceRequest) (<-chan *IndexFaceResponse, <-chan error) {
	responseChan := make(chan *IndexFaceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.IndexFace(request)
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

// IndexFaceWithCallback invokes the imm.IndexFace API asynchronously
// api document: https://help.aliyun.com/api/imm/indexface.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) IndexFaceWithCallback(request *IndexFaceRequest, callback func(response *IndexFaceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *IndexFaceResponse
		var err error
		defer close(result)
		response, err = client.IndexFace(request)
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

// IndexFaceRequest is the request struct for api IndexFace
type IndexFaceRequest struct {
	*requests.RpcRequest
	Chain   string `position:"Query" name:"Chain"`
	SrcUris string `position:"Query" name:"SrcUris"`
	Project string `position:"Query" name:"Project"`
	SetId   string `position:"Query" name:"SetId"`
	Force   string `position:"Query" name:"Force"`
}

// IndexFaceResponse is the response struct for api IndexFace
type IndexFaceResponse struct {
	*responses.BaseResponse
	RequestId      string               `json:"RequestId" xml:"RequestId"`
	SetId          string               `json:"SetId" xml:"SetId"`
	SrcUris        []string             `json:"SrcUris" xml:"SrcUris"`
	SuccessDetails []SuccessDetailsItem `json:"SuccessDetails" xml:"SuccessDetails"`
	FailDetails    []FailDetailsItem    `json:"FailDetails" xml:"FailDetails"`
}

// CreateIndexFaceRequest creates a request to invoke IndexFace API
func CreateIndexFaceRequest() (request *IndexFaceRequest) {
	request = &IndexFaceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imm", "2017-09-06", "IndexFace", "2017-09-06", "openAPI")
	return
}

// CreateIndexFaceResponse creates a response to parse from IndexFace response
func CreateIndexFaceResponse() (response *IndexFaceResponse) {
	response = &IndexFaceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
