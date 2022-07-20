package resourcemanager

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

// SendVerificationCodeForEnableRD invokes the resourcemanager.SendVerificationCodeForEnableRD API synchronously
func (client *Client) SendVerificationCodeForEnableRD(request *SendVerificationCodeForEnableRDRequest) (response *SendVerificationCodeForEnableRDResponse, err error) {
	response = CreateSendVerificationCodeForEnableRDResponse()
	err = client.DoAction(request, response)
	return
}

// SendVerificationCodeForEnableRDWithChan invokes the resourcemanager.SendVerificationCodeForEnableRD API asynchronously
func (client *Client) SendVerificationCodeForEnableRDWithChan(request *SendVerificationCodeForEnableRDRequest) (<-chan *SendVerificationCodeForEnableRDResponse, <-chan error) {
	responseChan := make(chan *SendVerificationCodeForEnableRDResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SendVerificationCodeForEnableRD(request)
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

// SendVerificationCodeForEnableRDWithCallback invokes the resourcemanager.SendVerificationCodeForEnableRD API asynchronously
func (client *Client) SendVerificationCodeForEnableRDWithCallback(request *SendVerificationCodeForEnableRDRequest, callback func(response *SendVerificationCodeForEnableRDResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SendVerificationCodeForEnableRDResponse
		var err error
		defer close(result)
		response, err = client.SendVerificationCodeForEnableRD(request)
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

// SendVerificationCodeForEnableRDRequest is the request struct for api SendVerificationCodeForEnableRD
type SendVerificationCodeForEnableRDRequest struct {
	*requests.RpcRequest
	SecureMobilePhone string `position:"Query" name:"SecureMobilePhone"`
}

// SendVerificationCodeForEnableRDResponse is the response struct for api SendVerificationCodeForEnableRD
type SendVerificationCodeForEnableRDResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateSendVerificationCodeForEnableRDRequest creates a request to invoke SendVerificationCodeForEnableRD API
func CreateSendVerificationCodeForEnableRDRequest() (request *SendVerificationCodeForEnableRDRequest) {
	request = &SendVerificationCodeForEnableRDRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ResourceManager", "2020-03-31", "SendVerificationCodeForEnableRD", "", "")
	request.Method = requests.POST
	return
}

// CreateSendVerificationCodeForEnableRDResponse creates a response to parse from SendVerificationCodeForEnableRD response
func CreateSendVerificationCodeForEnableRDResponse() (response *SendVerificationCodeForEnableRDResponse) {
	response = &SendVerificationCodeForEnableRDResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
