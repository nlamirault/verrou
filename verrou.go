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
	"image"
	"os"
	"os/exec"

	"github.com/disintegration/imaging"
	"github.com/vova616/screenshot"
)

const (
	// BANNER is what is printed for help/info output.
	BANNER = "Verrou"
	// VERSION is the binary version.
	VERSION = "0.1.0"

	blurSigma  = 5.0
	background = "/tmp/verrou.png"
)

var (
	// Flags
	dryrun bool
	vrs    bool
)

func init() {
	// parse flags
	flag.BoolVar(&vrs, "version", false, "print version and exit")
	flag.BoolVar(&dryrun, "dry-run", false, "Create background image and not lock screen")

	flag.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprintf("%s - v%s\n", BANNER, VERSION))
		flag.PrintDefaults()
	}

	flag.Parse()

	if vrs {
		fmt.Printf("v%s\n", VERSION)
		os.Exit(0)
	}
}

func lockScreen(path string) {
	if err := exec.Command("i3lock", "-n", "-i"+path).Run(); err != nil {
		panic(err)
	}
}

func makeScreenshot() (*image.RGBA, error) {
	return screenshot.CaptureScreen()
}

func main() {
	screenshot, err := makeScreenshot()
	if err != nil {
		panic(err)
	}

	blurredImage := imaging.Blur(screenshot, blurSigma)
	err = imaging.Save(blurredImage, background)
	if err != nil {
		panic(err)
	}

	if !dryrun {
		lockScreen(background)
	}
}
