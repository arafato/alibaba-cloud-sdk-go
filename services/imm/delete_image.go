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

// DeleteImage invokes the imm.DeleteImage API synchronously
// api document: https://help.aliyun.com/api/imm/deleteimage.html
func (client *Client) DeleteImage(request *DeleteImageRequest) (response *DeleteImageResponse, err error) {
	response = CreateDeleteImageResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteImageWithChan invokes the imm.DeleteImage API asynchronously
// api document: https://help.aliyun.com/api/imm/deleteimage.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteImageWithChan(request *DeleteImageRequest) (<-chan *DeleteImageResponse, <-chan error) {
	responseChan := make(chan *DeleteImageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteImage(request)
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

// DeleteImageWithCallback invokes the imm.DeleteImage API asynchronously
// api document: https://help.aliyun.com/api/imm/deleteimage.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteImageWithCallback(request *DeleteImageRequest, callback func(response *DeleteImageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteImageResponse
		var err error
		defer close(result)
		response, err = client.DeleteImage(request)
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

// DeleteImageRequest is the request struct for api DeleteImage
type DeleteImageRequest struct {
	*requests.RpcRequest
	ImageUri string `position:"Query" name:"ImageUri"`
	Project  string `position:"Query" name:"Project"`
	SetId    string `position:"Query" name:"SetId"`
}

// DeleteImageResponse is the response struct for api DeleteImage
type DeleteImageResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	SetId     string `json:"SetId" xml:"SetId"`
	ImageUri  string `json:"ImageUri" xml:"ImageUri"`
}

// CreateDeleteImageRequest creates a request to invoke DeleteImage API
func CreateDeleteImageRequest() (request *DeleteImageRequest) {
	request = &DeleteImageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imm", "2017-09-06", "DeleteImage", "2017-09-06", "openAPI")
	return
}

// CreateDeleteImageResponse creates a response to parse from DeleteImage response
func CreateDeleteImageResponse() (response *DeleteImageResponse) {
	response = &DeleteImageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
