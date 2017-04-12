package client

import (
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
)

// CreateGenresPath computes a request path to the create action of genres.
func CreateGenresPath() string {
	return fmt.Sprintf("/api/v2/genres/new")
}

// ジャンルの新規作成
func (c *Client) CreateGenres(ctx context.Context, path string, name string) (*http.Response, error) {
	req, err := c.NewCreateGenresRequest(ctx, path, name)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreateGenresRequest create the request corresponding to the create action endpoint of the genres resource.
func (c *Client) NewCreateGenresRequest(ctx context.Context, path string, name string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	values.Set("name", name)
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

// FollowGenreGenresPath computes a request path to the follow genre action of genres.
func FollowGenreGenresPath(genreID int) string {
	return fmt.Sprintf("/api/v2/genres/%v/follow", genreID)
}

// FollowGenreGenresPath2 computes a request path to the follow genre action of genres.
func FollowGenreGenresPath2(genreID int) string {
	return fmt.Sprintf("/api/v2/genres/%v/follow", genreID)
}

// ジャンルお気に入り操作
func (c *Client) FollowGenreGenres(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewFollowGenreGenresRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewFollowGenreGenresRequest create the request corresponding to the follow genre action endpoint of the genres resource.
func (c *Client) NewFollowGenreGenresRequest(ctx context.Context, path string) (*http.Request, error) {
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

// ListGenresPath computes a request path to the list action of genres.
func ListGenresPath() string {
	return fmt.Sprintf("/api/v2/genres")
}

// ジャンル取得
func (c *Client) ListGenres(ctx context.Context, path string, q *string) (*http.Response, error) {
	req, err := c.NewListGenresRequest(ctx, path, q)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListGenresRequest create the request corresponding to the list action endpoint of the genres resource.
func (c *Client) NewListGenresRequest(ctx context.Context, path string, q *string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if q != nil {
		values.Set("q", *q)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.KeySigner != nil {
		c.KeySigner.Sign(req)
	}
	return req, nil
}
