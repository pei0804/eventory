package client

import (
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
)

// AccountCreateUsersPath computes a request path to the account create action of users.
func AccountCreateUsersPath() string {
	return fmt.Sprintf("/api/v2/users/new")
}

// 正規ユーザーの作成
func (c *Client) AccountCreateUsers(ctx context.Context, path string, email string, identifier string) (*http.Response, error) {
	req, err := c.NewAccountCreateUsersRequest(ctx, path, email, identifier)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewAccountCreateUsersRequest create the request corresponding to the account create action endpoint of the users resource.
func (c *Client) NewAccountCreateUsersRequest(ctx context.Context, path string, email string, identifier string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	values.Set("email", email)
	values.Set("identifier", identifier)
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.KeySigner != nil {
		c.KeySigner.Sign(req)
	}
	return req, nil
}

// AccountTerminalStatusUpdateUsersPath computes a request path to the account terminal status update action of users.
func AccountTerminalStatusUpdateUsersPath() string {
	return fmt.Sprintf("/api/v2/users/status")
}

// 一時ユーザーの作成
func (c *Client) AccountTerminalStatusUpdateUsers(ctx context.Context, path string, clientVersion string, platform string) (*http.Response, error) {
	req, err := c.NewAccountTerminalStatusUpdateUsersRequest(ctx, path, clientVersion, platform)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewAccountTerminalStatusUpdateUsersRequest create the request corresponding to the account terminal status update action endpoint of the users resource.
func (c *Client) NewAccountTerminalStatusUpdateUsersRequest(ctx context.Context, path string, clientVersion string, platform string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	values.Set("client_version", clientVersion)
	values.Set("platform", platform)
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("PUT", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.KeySigner != nil {
		c.KeySigner.Sign(req)
	}
	return req, nil
}

// TmpAccountCreateUsersPath computes a request path to the tmp account create action of users.
func TmpAccountCreateUsersPath() string {
	return fmt.Sprintf("/api/v2/users/tmp")
}

// 一時ユーザーの作成
func (c *Client) TmpAccountCreateUsers(ctx context.Context, path string, clientVersion string, identifier string, platform string) (*http.Response, error) {
	req, err := c.NewTmpAccountCreateUsersRequest(ctx, path, clientVersion, identifier, platform)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewTmpAccountCreateUsersRequest create the request corresponding to the tmp account create action endpoint of the users resource.
func (c *Client) NewTmpAccountCreateUsersRequest(ctx context.Context, path string, clientVersion string, identifier string, platform string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	values.Set("client_version", clientVersion)
	values.Set("identifier", identifier)
	values.Set("platform", platform)
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
