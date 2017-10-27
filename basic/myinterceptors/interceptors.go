package myinterceptors

import (
	"strings"
	"net/http"

	"github.com/rihtim/core/log"
	"github.com/rihtim/core/utils"
	"github.com/rihtim/core/messages"
	"github.com/rihtim/core/requestscope"
	"github.com/rihtim/core/dataprovider"
)

// Checks path, returns error if path equals to '/wrong'
func CheckPath(rs requestscope.RequestScope, extras interface{}, req, res messages.Message, dp dataprovider.Provider) (editedReq, editedRes messages.Message, editedRs requestscope.RequestScope, err *utils.Error) {

	if strings.EqualFold(req.Res, "/wrong") {
		log.Info("Wrong path visited, returning error")
		err = &utils.Error{Code: http.StatusBadGateway, Message: "Wrong way!"}
	}

	return
}

// Checks body fields, returns error if one of email, firstName and lastName is missing
func ValidateUserInput(rs requestscope.RequestScope, extras interface{}, req, res messages.Message, dp dataprovider.Provider) (editedReq, editedRes messages.Message, editedRs requestscope.RequestScope, err *utils.Error) {

	if _, containsEmail := req.Body["email"]; !containsEmail {
		err = &utils.Error{Code: http.StatusBadRequest, Message: "Email must be provided in the body."}
	} else if _, containsFirstName := req.Body["firstName"]; !containsFirstName {
		err = &utils.Error{Code: http.StatusBadRequest, Message: "First name must be provided in the body."}
	} else if _, containsLastName := req.Body["lastName"]; !containsLastName {
		err = &utils.Error{Code: http.StatusBadRequest, Message: "Last name must be provided in the body."}
	}

	return
}

// Checks body fields, returns error if one of email, firstName and lastName is missing
func AddChangedFieldInfoToResponse(rs requestscope.RequestScope, extras interface{}, req, res messages.Message, dp dataprovider.Provider) (editedReq, editedRes messages.Message, editedRs requestscope.RequestScope, err *utils.Error) {

	log.Info("Adding changed fields info to response")

	changedFields := make([]string, 0)
	for key := range req.Body {
		changedFields = append(changedFields, key)
	}

	editedRes = res
	editedRes.Body["changedFields"] = strings.Join(changedFields[:], ", ")

	return
}

// Checks if headers contain "If-Modified-Since", return 304 (StatusNotModified) if header exists
func CheckIfModifiedSinceHeader(rs requestscope.RequestScope, extras interface{}, req, res messages.Message, dp dataprovider.Provider) (editedReq, editedRes messages.Message, editedRs requestscope.RequestScope, err *utils.Error) {

	if lastModifiedHeader, hasLastModifiedHeader := req.Headers["If-Modified-Since"]; hasLastModifiedHeader {

		if _, ptErr := http.ParseTime(lastModifiedHeader[0]); ptErr != nil {
			err = &utils.Error{Code: http.StatusBadRequest, Message: "Last modified header value format is not correct."}
			return
		}

		log.Info("'If-Modified-Since' header found, returning 304 - StatusNotModified")
		editedRes = messages.Message{Status: http.StatusNotModified}
		return
	}
	return
}
