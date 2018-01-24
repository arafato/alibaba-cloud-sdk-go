package ddospro

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

func (client *Client) CreateDomain(request *CreateDomainRequest) (response *CreateDomainResponse, err error) {
	response = CreateCreateDomainResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) CreateDomainWithChan(request *CreateDomainRequest) (<-chan *CreateDomainResponse, <-chan error) {
	responseChan := make(chan *CreateDomainResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateDomain(request)
		responseChan <- response
		errChan <- err
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

func (client *Client) CreateDomainWithCallback(request *CreateDomainRequest, callback func(response *CreateDomainResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateDomainResponse
		var err error
		defer close(result)
		response, err = client.CreateDomain(request)
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

type CreateDomainRequest struct {
	*requests.RpcRequest
	CcEnable        requests.Boolean `position:"Query" name:"CcEnable"`
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Domain          string           `position:"Query" name:"Domain"`
	Ip              string           `position:"Query" name:"Ip"`
	ProxyType       *[]string        `position:"Query" name:"ProxyType"  type:"Repeated"`
	RealServer      *[]string        `position:"Query" name:"RealServer"  type:"Repeated"`
	Type            string           `position:"Query" name:"Type"`
}

type CreateDomainResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateCreateDomainRequest() (request *CreateDomainRequest) {
	request = &CreateDomainRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("DDoSPro", "2017-07-25", "CreateDomain", "", "")
	return
}

func CreateCreateDomainResponse() (response *CreateDomainResponse) {
	response = &CreateDomainResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}