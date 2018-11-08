/*
 *    Copyright 2018 Insolar
 *
 *    Licensed under the Apache License, Version 2.0 (the License);
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an AS IS BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path"

	"github.com/insolar/insolar/log"
	"gopkg.in/yaml.v2"
)

type SFeature struct {
	StartVersion string `yaml:"startversion"`
	Description  string `yaml:"description"`
}

type SVersionTable struct {
	V map[string]SFeature `yaml:"versiontable"`
}

//go:generate go run template.go
func main() {
	vtContent, _ := ioutil.ReadFile(path.Join("..", "..", "..", "versiontable.yml"))
	buffer, err := parseYaml(vtContent)
	if err != nil {
		log.Error(err)
	}
	if err := ioutil.WriteFile(path.Join("..", "versiontable.go"), buffer.Bytes(), 0644); err != nil {
		log.Error(err)
	}
}

func parseYaml(vtContent []byte) (*bytes.Buffer, error) {
	vt := &SVersionTable{}
	if err := yaml.Unmarshal(vtContent, vt); err != nil {
		return nil, err
	}

	buffer := bytes.NewBuffer(nil)
	fmt.Fprintln(buffer, `// Code generated by "go run template.go"; DO NOT EDIT.`)
	fmt.Fprintln(buffer, "")
	fmt.Fprintln(buffer, "package manager")
	fmt.Fprintln(buffer, "")
	if len(vt.V) > 0 {
		fmt.Fprintln(buffer, "import (\n	\"github.com/insolar/insolar/log\"\n)")
		fmt.Fprintln(buffer, "")
	}
	fmt.Fprintln(buffer, "func (vm *VersionManager) loadVersionTable() {")
	if len(vt.V) > 0 {
		fmt.Fprintln(buffer, "	var err error")
		fmt.Fprintln(buffer, "")
	}
	for key, value := range vt.V {
		fmt.Fprintln(buffer, `	vm.VersionTable["`+key+`"], err = NewFeature("`+key+`","`+value.StartVersion+`", "`+value.Description+`")`)
		fmt.Fprintln(buffer, "	if(err!=nil){")
		fmt.Fprintln(buffer, `		log.Warn("Error loading from versiontable.yml, verify structure, key='`+key+"', startVersion='"+value.StartVersion+
			`', message: "+ err.Error())`)
		fmt.Fprintln(buffer, "	}")
		fmt.Fprintln(buffer, "")
	}

	fmt.Fprintln(buffer, "	return")
	fmt.Fprintf(buffer, "}")
	return buffer, nil
}
