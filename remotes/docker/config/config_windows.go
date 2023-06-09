//go:build windows
// +build windows

/*
   Copyright The containerd Authors.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package config

import (
	"crypto/x509"
	"path/filepath"
	"strings"
)

func hostPaths(root, host string) []string {
	ch := hostDirectory(host)
	if ch == host {
		return []string{filepath.Join(root, host)}
	}

	return []string{
		filepath.Join(root, strings.Replace(ch, ":", "", -1)),
		filepath.Join(root, strings.Replace(host, ":", "", -1)),
	}
}

func rootSystemPool() (*x509.CertPool, error) {
	return x509.NewCertPool(), nil
}
