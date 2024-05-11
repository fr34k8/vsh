package kms

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

// Encrypt invokes the kms.Encrypt API synchronously
func (client *Client) Encrypt(request *EncryptRequest) (response *EncryptResponse, err error) {
	response = CreateEncryptResponse()
	err = client.DoAction(request, response)
	return
}

// EncryptWithChan invokes the kms.Encrypt API asynchronously
func (client *Client) EncryptWithChan(request *EncryptRequest) (<-chan *EncryptResponse, <-chan error) {
	responseChan := make(chan *EncryptResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.Encrypt(request)
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

// EncryptWithCallback invokes the kms.Encrypt API asynchronously
func (client *Client) EncryptWithCallback(request *EncryptRequest, callback func(response *EncryptResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *EncryptResponse
		var err error
		defer close(result)
		response, err = client.Encrypt(request)
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

// EncryptRequest is the request struct for api Encrypt
type EncryptRequest struct {
	*requests.RpcRequest
	KeyId             string `position:"Query" name:"KeyId"`
	Plaintext         string `position:"Query" name:"Plaintext"`
	EncryptionContext string `position:"Query" name:"EncryptionContext"`
}

// EncryptResponse is the response struct for api Encrypt
type EncryptResponse struct {
	*responses.BaseResponse
	KeyVersionId   string `json:"KeyVersionId" xml:"KeyVersionId"`
	KeyId          string `json:"KeyId" xml:"KeyId"`
	CiphertextBlob string `json:"CiphertextBlob" xml:"CiphertextBlob"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
}

// CreateEncryptRequest creates a request to invoke Encrypt API
func CreateEncryptRequest() (request *EncryptRequest) {
	request = &EncryptRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Kms", "2016-01-20", "Encrypt", "kms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateEncryptResponse creates a response to parse from Encrypt response
func CreateEncryptResponse() (response *EncryptResponse) {
	response = &EncryptResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
