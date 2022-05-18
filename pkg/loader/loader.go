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

package loader

import (
	"fmt"

	"github.com/spf13/viper"
)

const (
	EnvKeyPrefix = "NETTYGO"
)

// Load a func load config form fs by viper[github](https://github.com/spf13/viper).
//
// @param fileName the target config file's name
//
// @param fileType the target config file's type
//
// @param dst the target struct of decode
//
// @param searchPath search paths
//
/*
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml") // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("/etc/appname/")  // path to look for the config file in
	viper.AddConfigPath("$HOME/.appname") // call multiple times to add many search paths
	viper.AddConfigPath(".")              // optionally look for config in the working directory
	err := viper.ReadInConfig() // Find and read the config file
...
*/
func Load(fileName, fileType string, dst any, searchPath ...string) error {
	viper.SetConfigName(fileName)
	viper.SetConfigType(fileType)
	viper.SetEnvPrefix(EnvKeyPrefix)
	for _, path := range searchPath {
		viper.AddConfigPath(path)
	}

	err := viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("Fatal error config file: %w \n", err)
	}

	err = viper.Unmarshal(dst)
	if err != nil {
		return fmt.Errorf("unable to decode into struct, %v", err)
	}

	return nil
}
