/*
 * Copyright © 2022 photowey (photowey@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package loader_test

import (
	"testing"

	"github.com/photowey/nettygo/pkg/loader"
)

type (
	config struct {
		Author
	}
	Author struct {
		Name  string
		Email string
	}
)

func TestLoad(t *testing.T) {
	type args struct {
		fileName string
		fileType string
		dst      any
		pathList []string
	}
	var dst config

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "test load config file",
			args: args{
				fileName: "config",
				fileType: "yml",
				dst:      &dst,
				pathList: []string{"./testdata/"},
			},
			want:    "photowey",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := loader.Load(tt.args.fileName, tt.args.fileType, tt.args.dst, tt.args.pathList...); (err != nil) != tt.wantErr {
				t.Errorf("Load() error = %v, wantErr %v", err, tt.wantErr)
			}
			if dst.Name != tt.want {
				t.Errorf("Load() error got = %v, want = %v", dst.Name, tt.want)
			}
			t.Log("ok")
		})
	}
}
