// 環境変数を扱うための package
package env

import (
	"fmt"
	"os"
)

// 環境変数を扱う
type Env interface {
	// ENV_NAME=ENV_VALUE
	fmt.Stringer
	Name() string
	Value() string
}

// 環境変数の取得
func Get(name string) Env {
	return get(name)
}

type env struct {
	name, value string
}

func get(name string) *env {
	return &env{
		name:  name,
		value: os.Getenv(name),
	}
}

func (e *env) String() string {
	return fmt.Sprintf("%s=%s", e.name, e.value)
}

func (e *env) Name() string {
	return e.name
}

func (e *env) Value() string {
	return e.value
}

// 環境変数の取得
// ログに値を残したくない場合、こちらを利用する
func GetSecure(name string) Env {
	env := secureEnv(*(get(name)))
	return &env
}

type secureEnv env

func (e *secureEnv) String() string {
	return fmt.Sprintf("%s=******", e.name)
}

func (e *secureEnv) Name() string {
	return e.name
}

func (e *secureEnv) Value() string {
	return e.value
}
