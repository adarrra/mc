/*
 * Minio Client (C) 2014, 2015 Minio, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"regexp"
	"runtime"
	"strings"

	"github.com/hashicorp/go-version"
	"github.com/minio/mc/pkg/console"
	"github.com/minio/minio/pkg/iodine"
)

var minWindowsGolangVersion = "1.5"
var minGolangVersion = "1.3"

// following code handles the current Golang release styles, we might have to update them in future
// if golang community divulges from the below formatting style.
const (
	betaRegexp = "beta[0-9]"
	rcRegexp   = "rc[0-9]"
)

func getNormalizedGolangVersion() string {
	version := strings.TrimPrefix(runtime.Version(), "go")
	br := regexp.MustCompile(betaRegexp)
	rr := regexp.MustCompile(rcRegexp)
	betaStr := br.FindString(version)
	version = strings.TrimRight(version, betaStr)
	rcStr := rr.FindString(version)
	version = strings.TrimRight(version, rcStr)
	return version
}

func checkGolangVersion() {
	v1, err := version.NewVersion(getNormalizedGolangVersion())
	if err != nil {
		console.Fatalf("Unable to parse runtime version, %s\n", NewIodine(iodine.New(err, nil)))
	}
	switch runtime.GOOS {
	case "windows":
		v2, err := version.NewVersion(minWindowsGolangVersion)
		if err != nil {
			console.Fatalf("Unable to parse minimum version, %s\n", NewIodine(iodine.New(err, nil)))
		}
		if v1.LessThan(v2) {
			console.Errorln("Minimum Golang runtime expected on windows is go1.5, please compile ‘mc’ with go1.5")
		}
	default:
		v2, err := version.NewVersion(minGolangVersion)
		if err != nil {
			console.Fatalf("Unable to parse minimum version, %s\n", NewIodine(iodine.New(err, nil)))
		}
		if v1.LessThan(v2) {
			console.Errorln("Minimum Golang runtime expected on windows is go1.3, please compile ‘mc’ with go1.3")
		}
	}
}

func firstTimeRun() {
	if !isMcConfigExists() {
		if err := createMcConfigDir(); err != nil {
			console.Fatalf("Unable to create ‘mc’ config folder. %s\n", err)
		}
		config, err := newConfig()
		if err != nil {
			console.Fatalln(NewIodine(iodine.New(err, nil)))
		}
		err = writeConfig(config)
		if err != nil {
			console.Fatalln(NewIodine(iodine.New(err, nil)))
		}
		console.Infoln("Configuration written to [" + mustGetMcConfigPath() + "]. Please update your access credentials.")
	}
	if !isSessionDirExists() {
		if err := createSessionDir(); err != nil {
			console.Fatalf("Unable to create ‘mc’ session folder. %s\n", err)
		}
	}
	checkGolangVersion()
}
