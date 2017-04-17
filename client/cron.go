package client

import (
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
)

// AppendGenreCronPath computes a request path to the append genre action of cron.
func AppendGenreCronPath() string {
	return fmt.Sprintf("/api/v2/cron/events/appendgenre")
}

// <b>イベントにジャンルを付加する<b>
func (c *Client) AppendGenreCron(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewAppendGenreCronRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewAppendGenreCronRequest create the request corresponding to the append genre action endpoint of the cron resource.
func (c *Client) NewAppendGenreCronRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.CronTokenSigner != nil {
		c.CronTokenSigner.Sign(req)
	}
	return req, nil
}

// FixUserFollowCronPath computes a request path to the fix user follow action of cron.
func FixUserFollowCronPath() string {
	return fmt.Sprintf("/api/v2/cron/user/events/fixfollow")
}

// <b>イベントフォロー操作の確定</b><br>
// user_follow_eventsテーブルのbatch_processedをtrueに変更する
func (c *Client) FixUserFollowCron(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewFixUserFollowCronRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewFixUserFollowCronRequest create the request corresponding to the fix user follow action endpoint of the cron resource.
func (c *Client) NewFixUserFollowCronRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.CronTokenSigner != nil {
		c.CronTokenSigner.Sign(req)
	}
	return req, nil
}

// NewEventFetchCronPath computes a request path to the new event fetch action of cron.
func NewEventFetchCronPath() string {
	return fmt.Sprintf("/api/v2/cron/events/fetch")
}

// <b>最新イベント情報の取得<b>
func (c *Client) NewEventFetchCron(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewNewEventFetchCronRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewNewEventFetchCronRequest create the request corresponding to the new event fetch action endpoint of the cron resource.
func (c *Client) NewNewEventFetchCronRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.CronTokenSigner != nil {
		c.CronTokenSigner.Sign(req)
	}
	return req, nil
}
