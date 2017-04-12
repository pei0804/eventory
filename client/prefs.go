package client

import (
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
)

// PrefFollowPrefsPath computes a request path to the pref follow action of prefs.
func PrefFollowPrefsPath(prefID int) string {
	return fmt.Sprintf("/api/v2/prefs/%v/follow", prefID)
}

// PrefFollowPrefsPath2 computes a request path to the pref follow action of prefs.
func PrefFollowPrefsPath2(prefID int) string {
	return fmt.Sprintf("/api/v2/prefs/%v/follow", prefID)
}

// ジャンルお気に入り操作
func (c *Client) PrefFollowPrefs(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewPrefFollowPrefsRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewPrefFollowPrefsRequest create the request corresponding to the pref follow action endpoint of the prefs resource.
func (c *Client) NewPrefFollowPrefsRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("PUT", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.KeySigner != nil {
		c.KeySigner.Sign(req)
	}
	return req, nil
}
