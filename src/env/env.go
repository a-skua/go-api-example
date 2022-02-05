// 環境変数を扱うための package
package env

import (
	"fmt"
	"os"
)

type Env interface {
	// 表示に利用
	fmt.Stringer
	// 値の取得に利用
	Value() string
}

type env string

// os.Getenv の Wrap
func Get(name string) Env {
	return env(os.Getenv(name))
}

func (e env) String() string {
	return string(e)
}

func (e env) Value() string {
	return string(e)
}

// ログに情報を残さないようにする
type secureEnv string

// os.Getenv の Wrap
func GetSecure(name string) Env {
	return secureEnv(os.Getenv(name))
}

func (e secureEnv) String() string {
	return "*****"
}

func (e secureEnv) Value() string {
	return string(e)
}
