// Copyright 2019 Tamás Gulácsi. All rights reserved.

package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"

	apiclient "github.com/UNO-SOFT/mantis-rest/client"
	config "github.com/UNO-SOFT/mantis-rest/client/config"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

func main() {
	if err := Main(); err != nil {
		log.Fatalf("%+v", err)
	}
}

func Main() error {
	flagMantis := flag.String("url", "", "Mantis URL to target (for example https://www.example.com/sub/path/mantis/api/rest/index.php)")
	flagTokenEnv := flag.String("token-env", "API_ACCESS_TOKEN", "API token environment name")
	flag.Parse()

	u, err := url.Parse(*flagMantis)
	if err != nil {
		return fmt.Errorf("%s: %w", *flagMantis, err)
	}

	client := apiclient.NewHTTPClientWithConfig(strfmt.Default, &apiclient.TransportConfig{
		Host:     u.Host,
		BasePath: u.Path,
		Schemes:  []string{u.Scheme},
	})
	log.Println(client.Transport)

	tokenAuth := httptransport.APIKeyAuth("Authorization", "header", os.Getenv(*flagTokenEnv))
	log.Println(tokenAuth)

	configGetParams := config.NewConfigGetParams()
	configGetParams.Option = []string{"csv_separator"}
	resp, err := client.Config.ConfigGet(configGetParams, tokenAuth)
	if err != nil {
		return err
	}
	log.Println(resp.Payload)
	return nil
}
