// Code generated by goagen v1.1.0, command line:
// $ goagen
// --design=github.com/tikasan/eventory/design
// --out=$(GOPATH)
// --version=v1.1.0-dirty
//
// API "eventory": users TestHelpers
//
// The content of this file is auto-generated, DO NOT MODIFY

package test

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"

	"github.com/goadesign/goa"
	"github.com/goadesign/goa/goatest"
	"github.com/tikasan/eventory/app"
	"golang.org/x/net/context"
)

// AccountCreateUsersBadRequest runs the method AccountCreate of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func AccountCreateUsersBadRequest(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.UsersController, email string, identifier string) (http.ResponseWriter, error) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	query := url.Values{}
	{
		sliceVal := []string{email}
		query["email"] = sliceVal
	}
	{
		sliceVal := []string{identifier}
		query["identifier"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/api/v2/users/new"),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	{
		sliceVal := []string{email}
		prms["email"] = sliceVal
	}
	{
		sliceVal := []string{identifier}
		prms["identifier"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "UsersTest"), rw, req, prms)
	accountCreateCtx, err := app.NewAccountCreateUsersContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}

	// Perform action
	err = ctrl.AccountCreate(accountCreateCtx)

	// Validate response
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}
	if rw.Code != 400 {
		t.Errorf("invalid response status code: got %+v, expected 400", rw.Code)
	}
	var mt error
	if resp != nil {
		var ok bool
		mt, ok = resp.(error)
		if !ok {
			t.Fatalf("invalid response media: got %+v, expected instance of error", resp)
		}
	}

	// Return results
	return rw, mt
}

// AccountCreateUsersOK runs the method AccountCreate of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func AccountCreateUsersOK(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.UsersController, email string, identifier string) (http.ResponseWriter, *app.Message) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	query := url.Values{}
	{
		sliceVal := []string{email}
		query["email"] = sliceVal
	}
	{
		sliceVal := []string{identifier}
		query["identifier"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/api/v2/users/new"),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	{
		sliceVal := []string{email}
		prms["email"] = sliceVal
	}
	{
		sliceVal := []string{identifier}
		prms["identifier"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "UsersTest"), rw, req, prms)
	accountCreateCtx, err := app.NewAccountCreateUsersContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}

	// Perform action
	err = ctrl.AccountCreate(accountCreateCtx)

	// Validate response
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt *app.Message
	if resp != nil {
		var ok bool
		mt, ok = resp.(*app.Message)
		if !ok {
			t.Fatalf("invalid response media: got %+v, expected instance of app.Message", resp)
		}
	}

	// Return results
	return rw, mt
}

// AccountCreateUsersUnauthorized runs the method AccountCreate of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func AccountCreateUsersUnauthorized(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.UsersController, email string, identifier string) http.ResponseWriter {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	query := url.Values{}
	{
		sliceVal := []string{email}
		query["email"] = sliceVal
	}
	{
		sliceVal := []string{identifier}
		query["identifier"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/api/v2/users/new"),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	{
		sliceVal := []string{email}
		prms["email"] = sliceVal
	}
	{
		sliceVal := []string{identifier}
		prms["identifier"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "UsersTest"), rw, req, prms)
	accountCreateCtx, err := app.NewAccountCreateUsersContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}

	// Perform action
	err = ctrl.AccountCreate(accountCreateCtx)

	// Validate response
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}
	if rw.Code != 401 {
		t.Errorf("invalid response status code: got %+v, expected 401", rw.Code)
	}

	// Return results
	return rw
}

// AccountTerminalStatusUpdateUsersBadRequest runs the method AccountTerminalStatusUpdate of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func AccountTerminalStatusUpdateUsersBadRequest(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.UsersController, clientVersion string, platform string) (http.ResponseWriter, error) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	query := url.Values{}
	{
		sliceVal := []string{clientVersion}
		query["client_version"] = sliceVal
	}
	{
		sliceVal := []string{platform}
		query["platform"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/api/v2/users/status"),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("PUT", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	{
		sliceVal := []string{clientVersion}
		prms["client_version"] = sliceVal
	}
	{
		sliceVal := []string{platform}
		prms["platform"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "UsersTest"), rw, req, prms)
	accountTerminalStatusUpdateCtx, err := app.NewAccountTerminalStatusUpdateUsersContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}

	// Perform action
	err = ctrl.AccountTerminalStatusUpdate(accountTerminalStatusUpdateCtx)

	// Validate response
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}
	if rw.Code != 400 {
		t.Errorf("invalid response status code: got %+v, expected 400", rw.Code)
	}
	var mt error
	if resp != nil {
		var ok bool
		mt, ok = resp.(error)
		if !ok {
			t.Fatalf("invalid response media: got %+v, expected instance of error", resp)
		}
	}

	// Return results
	return rw, mt
}

// AccountTerminalStatusUpdateUsersOK runs the method AccountTerminalStatusUpdate of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func AccountTerminalStatusUpdateUsersOK(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.UsersController, clientVersion string, platform string) http.ResponseWriter {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	query := url.Values{}
	{
		sliceVal := []string{clientVersion}
		query["client_version"] = sliceVal
	}
	{
		sliceVal := []string{platform}
		query["platform"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/api/v2/users/status"),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("PUT", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	{
		sliceVal := []string{clientVersion}
		prms["client_version"] = sliceVal
	}
	{
		sliceVal := []string{platform}
		prms["platform"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "UsersTest"), rw, req, prms)
	accountTerminalStatusUpdateCtx, err := app.NewAccountTerminalStatusUpdateUsersContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}

	// Perform action
	err = ctrl.AccountTerminalStatusUpdate(accountTerminalStatusUpdateCtx)

	// Validate response
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}

	// Return results
	return rw
}

// TmpAccountCreateUsersBadRequest runs the method TmpAccountCreate of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func TmpAccountCreateUsersBadRequest(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.UsersController, clientVersion string, identifier string, platform string) (http.ResponseWriter, error) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	query := url.Values{}
	{
		sliceVal := []string{clientVersion}
		query["client_version"] = sliceVal
	}
	{
		sliceVal := []string{identifier}
		query["identifier"] = sliceVal
	}
	{
		sliceVal := []string{platform}
		query["platform"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/api/v2/users/tmp"),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	{
		sliceVal := []string{clientVersion}
		prms["client_version"] = sliceVal
	}
	{
		sliceVal := []string{identifier}
		prms["identifier"] = sliceVal
	}
	{
		sliceVal := []string{platform}
		prms["platform"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "UsersTest"), rw, req, prms)
	tmpAccountCreateCtx, err := app.NewTmpAccountCreateUsersContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}

	// Perform action
	err = ctrl.TmpAccountCreate(tmpAccountCreateCtx)

	// Validate response
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}
	if rw.Code != 400 {
		t.Errorf("invalid response status code: got %+v, expected 400", rw.Code)
	}
	var mt error
	if resp != nil {
		var ok bool
		mt, ok = resp.(error)
		if !ok {
			t.Fatalf("invalid response media: got %+v, expected instance of error", resp)
		}
	}

	// Return results
	return rw, mt
}

// TmpAccountCreateUsersOK runs the method TmpAccountCreate of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func TmpAccountCreateUsersOK(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.UsersController, clientVersion string, identifier string, platform string) (http.ResponseWriter, *app.Token) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	query := url.Values{}
	{
		sliceVal := []string{clientVersion}
		query["client_version"] = sliceVal
	}
	{
		sliceVal := []string{identifier}
		query["identifier"] = sliceVal
	}
	{
		sliceVal := []string{platform}
		query["platform"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/api/v2/users/tmp"),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	{
		sliceVal := []string{clientVersion}
		prms["client_version"] = sliceVal
	}
	{
		sliceVal := []string{identifier}
		prms["identifier"] = sliceVal
	}
	{
		sliceVal := []string{platform}
		prms["platform"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "UsersTest"), rw, req, prms)
	tmpAccountCreateCtx, err := app.NewTmpAccountCreateUsersContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}

	// Perform action
	err = ctrl.TmpAccountCreate(tmpAccountCreateCtx)

	// Validate response
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt *app.Token
	if resp != nil {
		var ok bool
		mt, ok = resp.(*app.Token)
		if !ok {
			t.Fatalf("invalid response media: got %+v, expected instance of app.Token", resp)
		}
	}

	// Return results
	return rw, mt
}
