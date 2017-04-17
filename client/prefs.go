package client

import (
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
)

// FollowPrefsPath computes a request path to the follow action of prefs.
func FollowPrefsPath(prefID int) string {
	return fmt.Sprintf("/api/v2/prefs/%v/follow", prefID)
}

// FollowPrefsPath2 computes a request path to the follow action of prefs.
func FollowPrefsPath2(prefID int) string {
	return fmt.Sprintf("/api/v2/prefs/%v/follow", prefID)
}

// <b>都道府県フォロー操作</b><br>
// PUTでフォロー、DELETEでアンフォローをする。<br>
// HTTPメソッド意外は同じパラメーターで動作する。<br>
// 存在しない都道府県へのリクエストは404エラーを返す。
func (c *Client) FollowPrefs(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewFollowPrefsRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewFollowPrefsRequest create the request corresponding to the follow action endpoint of the prefs resource.
func (c *Client) NewFollowPrefsRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("PUT", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.UserTokenSigner != nil {
		c.UserTokenSigner.Sign(req)
	}
	return req, nil
}
