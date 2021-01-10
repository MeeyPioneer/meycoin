/**
 *  @file
 *  @copyright defined in meycoin/LICENSE.txt
 */

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"text/template"

	"github.com/meeypioneer/meycoin/types"
)

var tpl = `// Code generated by go run main.go {{.Input}} {{.Output}}; DO NOT EDIT.

package {{.Package}}

import (
	"fmt"
	"reflect"
	
	"github.com/meeypioneer/meycoin/types"
)

var (
	MainNetHardforkConfig = &HardforkConfig{
{{- range .Hardforks}}
		V{{.Version}}: types.BlockNo({{.MainNetHeight}}),
{{- end}}
	}
	TestNetHardforkConfig = &HardforkConfig{
{{- range .Hardforks}}
		V{{.Version}}: types.BlockNo({{.TestNetHeight}}),
{{- end}}
	}
	AllEnabledHardforkConfig = &HardforkConfig{
{{- range .Hardforks}}
		V{{.Version}}: types.BlockNo(0),
{{- end}}
	}
)

const hardforkConfigTmpl = ` + "`[hardfork]\n" + `
{{- range .Hardforks}}
v{{.Version}} = "{{"{{.Hardfork.V"}}{{.Version}}}}"
{{- end}}
` + "`" + `

type HardforkConfig struct {
{{- range .Hardforks}}
	V{{.Version}} types.BlockNo ` + "`" + `mapstructure:"v{{.Version}}" description:"a block number of the hardfork version {{.Version}}"` + "`" + `
{{- end}}
}

type HardforkDbConfig map[string]types.BlockNo
{{range .Hardforks}}
func (c *HardforkConfig) IsV{{.Version}}Fork(h types.BlockNo) bool {
	return isFork(c.V{{.Version}}, h)
}
{{end}}
func (c *HardforkConfig) CheckCompatibility(dbCfg HardforkDbConfig, h types.BlockNo) error {
	if err := c.validate(); err != nil {
		return err
	}
{{- range .Hardforks}}
	if (isFork(c.V{{.Version}}, h) || isFork(dbCfg["V{{.Version}}"], h)) && c.V{{.Version}} != dbCfg["V{{.Version}}"] {
		return newForkError("V{{.Version}}", h, c.V{{.Version}}, dbCfg["V{{.Version}}"])
	}
{{- end}}
	return checkOlderNode({{.MaxVersion}}, h, dbCfg)
}

func (c *HardforkConfig) Version(h types.BlockNo) int32 {
	v := reflect.ValueOf(*c)
	for i := v.NumField() - 1; i >= 0; i-- {
		if v.Field(i).Uint() <= h {
			return int32(i + 2)
		}
	}
	return int32(0)
}

func (c *HardforkConfig) validate() error {
	prev := uint64(0)
	v := reflect.ValueOf(*c)
	for i := 0; i < v.NumField(); i++ {
		curr := v.Field(i).Uint()
		if prev > curr {
			return fmt.Errorf("version %d has a lower block number: %d, %d(v%d)", i+2, curr, prev, i+1)
		}
		prev = curr
	}
	return nil
}
`

const versionStartNo = uint64(2)

type hardforkElem struct {
	Version                      uint64
	MainNetHeight, TestNetHeight types.BlockNo
}

type hardforkData struct {
	Hardforks  []hardforkElem
	MaxVersion uint64
	Package    string
	Input      string
	Output     string
}

func main() {
	if len(os.Args) != 3 {
		panic("Usage: go run main.go <input-file> <output-file>")
	}
	b, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	var hardforks []hardforkElem
	err = json.Unmarshal(b, &hardforks)
	if err != nil {
		panic(err)
	}
	bindVal := &hardforkData{
		Hardforks: hardforks,
	}
	err = validate(bindVal)
	if err != nil {
		panic(err)
	}
	tt := template.Must(template.New("hadfork").Parse(tpl))
	f, err := os.OpenFile(os.Args[2], os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	bindVal.Input = os.Args[1]
	bindVal.Output = os.Args[2]
	tt.Execute(f, bindVal)
	f.Close()
}

func validate(v *hardforkData) error {
	v.Package = os.Getenv("GOPACKAGE")
	v.MaxVersion = uint64(len(v.Hardforks) + 1)
	MainNetMax, TestNetMax := types.BlockNo(0), types.BlockNo(0)
	for i := versionStartNo; i <= v.MaxVersion; i++ {
		hf := v.Hardforks[i-versionStartNo]
		if i != hf.Version {
			return fmt.Errorf("version %d expected, but got %d", i, hf.Version)
		}
		if MainNetMax > hf.MainNetHeight {
			return fmt.Errorf("version %d, mainnet block number %d is too low", i, hf.MainNetHeight)
		}
		MainNetMax = hf.MainNetHeight
		if TestNetMax > hf.TestNetHeight {
			return fmt.Errorf("version %d, testnet block number %d is too low", i, hf.TestNetHeight)
		}
		TestNetMax = hf.TestNetHeight
	}
	return nil
}