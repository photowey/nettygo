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

package logger

type Config struct {
	Level           string `toml:"level" json:"level" yaml:"level"`
	Path            string `toml:"path" json:"path" yaml:"path"`
	FileName        string `toml:"file_name" json:"fileName" yaml:"fileName"`
	MaxSize         int    `toml:"max_size" json:"maxSize" yaml:"maxSize"`
	MaxBackups      int    `default:"0" toml:"max_backups" json:"maxBackups" yaml:"maxBackups"`
	MaxAgeDay       int    `toml:"max_age_day" json:"maxAgeDay" yaml:"maxAgeDay"`
	CompressEnabled bool   `toml:"compress_enabled" json:"compressEnabled" yaml:"compressEnabled"`
	StdoutEnabled   bool   `toml:"stdout_enabled" json:"stdoutEnabled" yaml:"stdoutEnabled"`
	JsonEnabled     bool   `toml:"json_enabled" json:"jsonEnabled" yaml:"jsonEnabled"`
}
