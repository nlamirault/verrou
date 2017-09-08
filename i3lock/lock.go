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

package i3lock

import (
	"os/exec"

	"github.com/nlamirault/verrou/screen"
)

const (
	ScreenLocker = "i3lock"
)

func LockScreen(path string) error {
	if err := screen.GenerateImage(path); err != nil {
		return err
	}
	if err := exec.Command("i3lock", "-n", "-i"+path).Run(); err != nil {
		return err
	}
	return nil
}
