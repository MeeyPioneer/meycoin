/*
 * @file
 * @copyright defined in meycoin/LICENSE.txt
 */

package p2pcommon

import (
	"errors"
	"github.com/coreos/go-semver/semver"
	"regexp"
)

// MeyCoinVersion follows sementic versioning https://semver.org/
type MeyCoinVersion = semver.Version

// const verPattern = `v([0-9]+)\.([0-9]+)\..([0-9]+)(-(.+))?`
const verPattern = `v[0-9].+`
var checker, _ = regexp.Compile(verPattern)

// ParseMeyCoinVersion parse version string to sementic version with slightly different manner. This function allows not only standard sementic version but also version strings containing v in prefixes.
func ParseMeyCoinVersion(verStr string) (MeyCoinVersion, error) {
	if checker.MatchString(verStr) {
		verStr = verStr[1:]
	}
	ver, err := semver.NewVersion(verStr)
	if err != nil {
		return MeyCoinVersion{}, errors.New("invalid version "+verStr)
	}
	return MeyCoinVersion(*ver), nil
}

// Supported MeyCoin version. wezen will register meycoinsvr within the version range. This version range should be modified when new release is born.
const (
	MinimumMeyCoinVersion = "v2.0.0"
	MaximumMeyCoinVersion = "v3.0.0"
)
var (
	minMeyCoinVersion MeyCoinVersion // inclusive
	maxMeyCoinVersion MeyCoinVersion // exclusive
)

func init() {
	var err error
	minMeyCoinVersion, err = ParseMeyCoinVersion(MinimumMeyCoinVersion)
	if err != nil {
		panic("Invalid minimum version "+MinimumMeyCoinVersion)
	}
	maxMeyCoinVersion, err = ParseMeyCoinVersion(MaximumMeyCoinVersion)
	if err != nil {
		panic("Invalid maximum version "+MaximumMeyCoinVersion)
	}
}


func CheckVersion(version string) bool {
	ver, err := ParseMeyCoinVersion(version)
	if err != nil {
		return false
	}

	return ver.Compare(minMeyCoinVersion) >= 0 && ver.Compare(maxMeyCoinVersion) < 0
}
