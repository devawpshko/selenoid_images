package cmd

import (
	"github.com/aerokube/selenoid-images/build"
	"github.com/spf13/cobra"
)

var (
	yandexCmd = &cobra.Command{
		Use:   "yandex",
		Short: "build Yandex.Browser image",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := build.Requirements{
				BrowserSource:  build.BrowserSource(browserSource),
				BrowserChannel: browserChannel,
				DriverVersion:  driverVersion,
				NoCache:        noCache,
				TestsDir:       testsDir,
				RunTests:       test,
				Tags:           tags,
				PushImage:      push,
			}
			chrome := &build.YandexBrowser{req}
			return chrome.Build()
		},
	}
)
