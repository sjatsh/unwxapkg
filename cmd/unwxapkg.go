// Copyright 2020 SunJun <i@sjis.me>
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
package main

import (
	"flag"

	"github.com/sjatsh/unwxapkg/pkg"
)

var f = flag.String("f", "", "wechat wxapkg file path")
var out = flag.String("o", ".", "output file path")

func main() {
	flag.Parse()
	extract := new(pkg.UnWxapkg)
	extract.InPath = *f
	extract.OutPath = *out
	if err := extract.Unwxapkg(); err != nil {
		panic(err)
	}
}
