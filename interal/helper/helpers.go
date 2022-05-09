package helper

import (
	`os`
)

func Match(expect bool, standard, standBy string) string {
	if expect {
		return standard
	}
	return standBy
}

func IsExist(path string) bool {
	_, err := os.Stat(path)

	return err != nil || os.IsExist(err)
}
