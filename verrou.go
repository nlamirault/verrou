// Copyright (C) 2017 Nicolas Lamirault <nicolas.lamirault@gmail.com>

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/nlamirault/verrou/i3lock"
	"github.com/nlamirault/verrou/screen"
)

const (
	// BANNER is what is printed for help/info output.
	BANNER = "Verrou"
	// VERSION is the binary version.
	VERSION = "0.2.0"

	background = "/tmp/verrou.png"
)

var (
	// Flags
	dryrun bool
	vrs    bool
	use    string
)

func init() {
	// parse flags
	flag.BoolVar(&vrs, "version", false, "print version and exit")
	flag.BoolVar(&dryrun, "dry-run", false, "Create background image and not lock screen")
	flag.StringVar(&use, "use", "", "Which screen locker to use")

	flag.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprintf("%s - v%s\n", BANNER, VERSION))
		flag.PrintDefaults()
	}

	flag.Parse()

	if vrs {
		fmt.Printf("v%s\n", VERSION)
		os.Exit(0)
	}

	if dryrun {

	}
}

func main() {
	if dryrun {
		if err := screen.GenerateImage(background); err != nil {
			panic(err)
		}
		fmt.Printf("See screen lock image: %s\n", background)
		return
	}

	switch use {
	case i3lock.ScreenLocker:
		if err := i3lock.LockScreen(background); err != nil {
			panic(err)
		}
	default:
		fmt.Printf("%s is not a supported screen locker.\n", use)
	}
}
