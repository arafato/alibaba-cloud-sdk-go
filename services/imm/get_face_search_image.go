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

// GetFaceSearchImage invokes the imm.GetFaceSearchImage API synchronously
// api document: https://help.aliyun.com/api/imm/getfacesearchimage.html
func (client *Client) GetFaceSearchImage(request *GetFaceSearchImageRequest) (response *GetFaceSearchImageResponse, err error) {
	response = CreateGetFaceSearchImageResponse()
	err = client.DoAction(request, response)
	return
}

// GetFaceSearchImageWithChan invokes the imm.GetFaceSearchImage API asynchronously
// api document: https://help.aliyun.com/api/imm/getfacesearchimage.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetFaceSearchImageWithChan(request *GetFaceSearchImageRequest) (<-chan *GetFaceSearchImageResponse, <-chan error) {
	responseChan := make(chan *GetFaceSearchImageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetFaceSearchImage(request)
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

// GetFaceSearchImageWithCallback invokes the imm.GetFaceSearchImage API asynchronously
// api document: https://help.aliyun.com/api/imm/getfacesearchimage.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetFaceSearchImageWithCallback(request *GetFaceSearchImageRequest, callback func(response *GetFaceSearchImageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetFaceSearchImageResponse
		var err error
		defer close(result)
		response, err = client.GetFaceSearchImage(request)
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

// GetFaceSearchImageRequest is the request struct for api GetFaceSearchImage
type GetFaceSearchImageRequest struct {
	*requests.RpcRequest
	ImageId   string `position:"Query" name:"ImageId"`
	Project   string `position:"Query" name:"Project"`
	SrcUri    string `position:"Query" name:"SrcUri"`
	GroupName string `position:"Query" name:"GroupName"`
	User      string `position:"Query" name:"User"`
}

// GetFaceSearchImageResponse is the response struct for api GetFaceSearchImage
type GetFaceSearchImageResponse struct {
	*responses.BaseResponse
	RequestId string   `json:"RequestId" xml:"RequestId"`
	GroupName string   `json:"GroupName" xml:"GroupName"`
	GroupId   string   `json:"GroupId" xml:"GroupId"`
	Gender    string   `json:"Gender" xml:"Gender"`
	Age       int      `json:"Age" xml:"Age"`
	Pitch     float64  `json:"Pitch" xml:"Pitch"`
	ImageUri  string   `json:"ImageUri" xml:"ImageUri"`
	ImageMd5  string   `json:"ImageMd5" xml:"ImageMd5"`
	Roll      float64  `json:"Roll" xml:"Roll"`
	FaceId    string   `json:"FaceId" xml:"FaceId"`
	Yaw       float64  `json:"Yaw" xml:"Yaw"`
	Quality   float64  `json:"Quality" xml:"Quality"`
	Hat       int      `json:"Hat" xml:"Hat"`
	Glasses   int      `json:"Glasses" xml:"Glasses"`
	ImageId   string   `json:"ImageId" xml:"ImageId"`
	User      string   `json:"User" xml:"User"`
	Axis      []string `json:"Axis" xml:"Axis"`
}

// CreateGetFaceSearchImageRequest creates a request to invoke GetFaceSearchImage API
func CreateGetFaceSearchImageRequest() (request *GetFaceSearchImageRequest) {
	request = &GetFaceSearchImageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imm", "2017-09-06", "GetFaceSearchImage", "2017-09-06", "openAPI")
	return
}

// CreateGetFaceSearchImageResponse creates a response to parse from GetFaceSearchImage response
func CreateGetFaceSearchImageResponse() (response *GetFaceSearchImageResponse) {
	response = &GetFaceSearchImageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
