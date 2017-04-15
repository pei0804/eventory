package main

import (
	"../client"
	"../tool/cli"
	"fmt"
	goaclient "github.com/goadesign/goa/client"
	"github.com/spf13/cobra"
	"net/http"
	"os"
	"time"
)

func main() {
	// Create command line parser
	app := &cobra.Command{
		Use:   "eventory-cli",
		Short: `CLI client for the eventory service (https://github.com/tikasan/eventory/wiki)`,
	}

	// Create client struct
	httpClient := newHTTPClient()
	c := client.New(goaclient.HTTPClientDoer(httpClient))

	// Register global flags
	app.PersistentFlags().StringVarP(&c.Scheme, "scheme", "s", "", "Set the requests scheme")
	app.PersistentFlags().StringVarP(&c.Host, "host", "H", "localhost:8080", "API hostname")
	app.PersistentFlags().DurationVarP(&httpClient.Timeout, "timeout", "t", time.Duration(20)*time.Second, "Set the request timeout")
	app.PersistentFlags().BoolVar(&c.Dump, "dump", false, "Dump HTTP request and response.")

	// Register signer flags
	var key, format string
	app.PersistentFlags().StringVar(&key, "key", "", "API key used for authentication")
	app.PersistentFlags().StringVar(&format, "format", "Bearer %s", "Format used to create auth header or query from key")

	// Parse flags and setup signers
	app.ParseFlags(os.Args)
	userTokenSigner := newUserTokenSigner(key, format)
	cronTokenSigner := newCronTokenSigner(key, format)

	// Initialize API client
	c.SetUserTokenSigner(userTokenSigner)
	c.SetCronTokenSigner(cronTokenSigner)
	c.UserAgent = "eventory-cli/0"

	// Register API commands
	cli.RegisterCommands(app, c)

	// Execute!
	if err := app.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(-1)
	}
}

// newHTTPClient returns the HTTP client used by the API client to make requests to the service.
func newHTTPClient() *http.Client {
	// TBD: Change as needed (e.g. to use a different transport to control redirection policy or
	// disable cert validation or...)
	return http.DefaultClient
}

// newUserTokenSigner returns the request signer used for authenticating
// against the userToken security scheme.
func newUserTokenSigner(key, format string) goaclient.Signer {
	return &goaclient.APIKeySigner{
		SignQuery: false,
		KeyName:   "X-Authorization",
		KeyValue:  key,
		Format:    format,
	}

}

// newCronTokenSigner returns the request signer used for authenticating
// against the cronToken security scheme.
func newCronTokenSigner(key, format string) goaclient.Signer {
	return &goaclient.APIKeySigner{
		SignQuery: false,
		KeyName:   "X-Appengine-Cron",
		KeyValue:  key,
		Format:    format,
	}

}
