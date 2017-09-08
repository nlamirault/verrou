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

package screen

import (
	"image"

	"github.com/disintegration/imaging"
	"github.com/vova616/screenshot"
)

const (
	blurSigma = 5.0
)

func GenerateImage(filename string) error {
	screenshot, err := takeScreenshot()
	if err != nil {
		return err
	}
	if err := blurring(screenshot, blurSigma, filename); err != nil {
		return err
	}
	return nil
}

func takeScreenshot() (*image.RGBA, error) {
	return screenshot.CaptureScreen()
}

func blurring(img image.Image, sigma float64, filename string) error {
	blurredImage := imaging.Blur(img, sigma)
	return imaging.Save(blurredImage, filename)
}
