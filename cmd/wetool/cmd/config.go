/*
 * @file
 * @copyright defined in meycoin/LICENSE.txt
 */

package cmd

import (
	"github.com/meeypioneer/mey-library/config"
	"github.com/meeypioneer/meycoin/wezen/common"
)

const (
	EnvironmentPrefix             = "AG"
	defaultWetoolHomePath          = ".wezen"
	defaultWetoolConfigFileName = "wetool.toml"
)

type CliContext struct {
	config.BaseContext
}

func NewCliContext(homePath string, configFilePath string) *CliContext {
	cliCtx := &CliContext{}
	cliCtx.BaseContext = config.NewBaseContext(cliCtx, homePath, configFilePath, EnvironmentPrefix)

	return cliCtx
}

// CliConfig is configs for meycoin cli.
type CliConfig struct {
	Host string `mapstructure:"host" description:"Target server host. default is localhost"`
	Port int    `mapstructure:"port" description:"Target server port. default is 8915"`
}

// GetDefaultConfig return cliconfig with default value. It ALWAYS returns NEW object.
func (ctx *CliContext) GetDefaultConfig() interface{} {
	return CliConfig{
		Host: "localhost",
		Port: common.DefaultRPCPort,
	}
}

func (ctx *CliContext) GetHomePath() string {
	return defaultWetoolHomePath
}

func (ctx *CliContext) GetConfigFileName() string {
	return defaultWetoolConfigFileName
}

func (ctx *CliContext) GetTemplate() string {
	return configTemplate
}

const configTemplate = `# meycoin cli TOML Configuration File (https://github.com/toml-lang/toml)
host = "{{.Host}}"
port = "{{.Port}}"
`
