//(C) Copyright [2020] Hewlett Packard Enterprise Development LP
//
//Licensed under the Apache License, Version 2.0 (the "License"); you may
//not use this file except in compliance with the License. You may obtain
//a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
//WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
//License for the specific language governing permissions and limitations
// under the License.

// Package account ...
package account

// ---------------------------------------------------------------------------------------
// IMPORT Section
// ---------------------------------------------------------------------------------------
import (
	"log"
	"net/http"

	"github.com/ODIM-Project/ODIM/lib-utilities/common"
	"github.com/ODIM-Project/ODIM/lib-utilities/errors"
	"github.com/ODIM-Project/ODIM/lib-utilities/response"
	"github.com/ODIM-Project/ODIM/svc-account-session/asmodel"
)

// Delete defines deletion of an existing account.
//
// Two parameters need to be passed to the function which are
// the Session, which contains all the session related data, espically the ConfigureUsers privilege
// and the accountID which is used for identifing the account to be deleted.
//
// As return parameters RPC response, which contains status code, message, headers and data,
// error will be passed back.
func Delete(session *asmodel.Session, accountID string) response.RPC {
	var resp response.RPC

	// Default admin user account should not be deleted
	if accountID == defaultAdminAccount {
		errorMessage := "default user account can not be deleted"
		resp.StatusCode = http.StatusBadRequest
		resp.StatusMessage = response.ResourceCannotBeDeleted
		args := response.Args{
			Code:    response.GeneralError,
			Message: "",
			ErrorArgs: []response.ErrArgs{
				response.ErrArgs{
					StatusMessage: resp.StatusMessage,
					ErrorMessage:  errorMessage,
				},
			},
		}
		resp.Body = args.CreateGenericErrorResponse()
		resp.Header = map[string]string{
			"Content-type": "application/json; charset=utf-8", // TODO: add all error headers
		}
		log.Printf(errorMessage)
		return resp
	}

	if !(session.Privileges[common.PrivilegeConfigureUsers]) {
		errorMessage := "error: " + session.UserName + " does not have the privilege to delete user"
		resp.StatusCode = http.StatusForbidden
		resp.StatusMessage = response.InsufficientPrivilege
		args := response.Args{
			Code:    response.GeneralError,
			Message: "",
			ErrorArgs: []response.ErrArgs{
				response.ErrArgs{
					StatusMessage: resp.StatusMessage,
					ErrorMessage:  errorMessage,
					MessageArgs:   []interface{}{},
				},
			},
		}
		resp.Body = args.CreateGenericErrorResponse()
		resp.Header = map[string]string{
			"Content-type": "application/json; charset=utf-8", // TODO: add all error headers
		}
		log.Printf(errorMessage)
		return resp
	}

	if derr := asmodel.DeleteUser(accountID); derr != nil {
		errorMessage := "error while deleting user: " + derr.Error()
		if errors.DBKeyNotFound == derr.ErrNo() {
			resp.StatusCode = http.StatusNotFound
			resp.StatusMessage = response.ResourceNotFound
			args := response.Args{
				Code:    response.GeneralError,
				Message: "",
				ErrorArgs: []response.ErrArgs{
					response.ErrArgs{
						StatusMessage: resp.StatusMessage,
						ErrorMessage:  errorMessage,
						MessageArgs:   []interface{}{"Account", accountID},
					},
				},
			}
			resp.Body = args.CreateGenericErrorResponse()
		} else {
			resp.CreateInternalErrorResponse(errorMessage)
		}
		resp.Header = map[string]string{
			"Content-type": "application/json; charset=utf-8", // TODO: add all error headers
		}
		log.Printf(errorMessage)
		return resp
	}

	resp.StatusCode = http.StatusNoContent
	resp.StatusMessage = response.AccountRemoved

	resp.Header = map[string]string{
		"Cache-Control":     "no-cache",
		"Connection":        "keep-alive",
		"Transfer-Encoding": "chunked",
		"Content-type":      "application/json; charset=utf-8",
		"OData-Version":     "4.0",
		"X-Frame-Options":   "sameorigin",
	}

	return resp
}
