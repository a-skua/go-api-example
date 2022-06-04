package env

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func FuzzEnv(f *testing.F) {
	f.Fuzz(func(t *testing.T, name, value string) {
		env := &env{
			name:  name,
			value: value,
		}

		if name != env.Name() {
			t.Fatalf("want=%v, got=%v.", name, env.Name())
		}

		if value != env.Value() {
			t.Fatalf("want=%v, got=%v.", name, env.Value())
		}

		want := fmt.Sprintf("%s=%s", name, value)
		if want != env.String() {
			t.Fatalf("want=%v, got=%v.", want, env.String())
		}
	})
}

func FuzzSecureEnv(f *testing.F) {
	f.Fuzz(func(t *testing.T, name, value string) {
		env := &secureEnv{
			name:  name,
			value: value,
		}

		if name != env.Name() {
			t.Fatalf("want=%v, got=%v.", name, env.Name())
		}

		if value != env.Value() {
			t.Fatalf("want=%v, got=%v.", name, env.Value())
		}

		want := fmt.Sprintf("%s=******", name)
		if want != env.String() {
			t.Fatalf("want=%v, got=%v.", want, env.String())
		}
	})
}

func TestGet(t *testing.T) {
	type test struct {
		name string
		want Env
	}

	do := func(tt test) {
		got := Get(tt.name)

		if !reflect.DeepEqual(tt.want, got) {
			t.Fatalf("want=%v, got=%v.", tt.want, got)
		}
	}

	tests := []test{
		func() test {
			err := os.Setenv("FOO", "foo")
			if err != nil {
				panic(err)
			}
			return test{
				name: "FOO",
				want: &env{name: "FOO", value: "foo"},
			}
		}(),
		func() test {
			err := os.Setenv("BAR", "bar")
			if err != nil {
				panic(err)
			}
			return test{
				name: "BAR",
				want: &env{name: "BAR", value: "bar"},
			}
		}(),
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestEnvName(t *testing.T) {
	type test struct {
		env  Env
		want string
	}

	do := func(tt test) {
		got := tt.env.Name()
		if tt.want != got {
			t.Fatalf("want=%v, got=%v.", tt.want, got)
		}
	}

	tests := []test{
		{
			env:  &env{name: ""},
			want: "",
		},
		{
			env:  &env{name: "FOO"},
			want: "FOO",
		},
		{
			env:  &env{name: "BAR"},
			want: "BAR",
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestEnvValue(t *testing.T) {
	type test struct {
		env  Env
		want string
	}

	do := func(tt test) {
		got := tt.env.Value()
		if tt.want != got {
			t.Fatalf("want=%v, got=%v.", tt.want, got)
		}
	}

	tests := []test{
		{
			env:  &env{value: ""},
			want: "",
		},
		{
			env:  &env{value: "foo"},
			want: "foo",
		},
		{
			env:  &env{value: "bar"},
			want: "bar",
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestEnvString(t *testing.T) {
	type test struct {
		env  Env
		want string
	}

	do := func(tt test) {
		got := tt.env.String()
		if tt.want != got {
			t.Fatalf("want=%v, got=%v.", tt.want, got)
		}
	}

	tests := []test{
		{
			env:  &env{name: "", value: ""},
			want: "=",
		},
		{
			env:  &env{name: "", value: "foo"},
			want: "=foo",
		},
		{
			env:  &env{name: "BAR", value: ""},
			want: "BAR=",
		},
		{
			env:  &env{name: "BAZ", value: "baz"},
			want: "BAZ=baz",
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestGetSecure(t *testing.T) {
	type test struct {
		name string
		want Env
	}

	do := func(tt test) {
		got := GetSecure(tt.name)

		if !reflect.DeepEqual(tt.want, got) {
			t.Fatalf("want=%v, got=%v.", tt.want, got)
		}
	}

	tests := []test{
		func() test {
			name, value := "FOO", "foo"
			err := os.Setenv(name, value)
			if err != nil {
				panic(err)
			}
			return test{
				name: name,
				want: &secureEnv{
					name:  name,
					value: value,
				},
			}
		}(),
		func() test {
			name, value := "BAR", "bar"
			err := os.Setenv(name, value)
			if err != nil {
				panic(err)
			}
			return test{
				name: name,
				want: &secureEnv{
					name:  name,
					value: value,
				},
			}
		}(),
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestSecureEnvName(t *testing.T) {
	type test struct {
		env  Env
		want string
	}

	do := func(tt test) {
		got := tt.env.Name()
		if tt.want != got {
			t.Fatalf("want=%v, got=%v.", tt.want, got)
		}
	}

	tests := []test{
		{
			env:  &secureEnv{name: ""},
			want: "",
		},
		{
			env:  &secureEnv{name: "FOO"},
			want: "FOO",
		},
		{
			env:  &secureEnv{name: "BAR"},
			want: "BAR",
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestSecureEnvValue(t *testing.T) {
	type test struct {
		env  Env
		want string
	}

	do := func(tt test) {
		got := tt.env.Value()
		if tt.want != got {
			t.Fatalf("want=%v, got=%v.", tt.want, got)
		}
	}

	tests := []test{
		{
			env:  &secureEnv{value: ""},
			want: "",
		},
		{
			env:  &secureEnv{value: "foo"},
			want: "foo",
		},
		{
			env:  &secureEnv{value: "bar"},
			want: "bar",
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestSecureEnvString(t *testing.T) {
	type test struct {
		env  Env
		want string
	}

	do := func(tt test) {
		got := tt.env.String()
		if tt.want != got {
			t.Fatalf("want=%v, got=%v.", tt.want, got)
		}
	}

	tests := []test{
		{
			env:  &secureEnv{name: "", value: ""},
			want: "=******",
		},
		{
			env:  &secureEnv{name: "", value: "foo"},
			want: "=******",
		},
		{
			env:  &secureEnv{name: "BAR", value: ""},
			want: "BAR=******",
		},
		{
			env:  &secureEnv{name: "BAZ", value: "baz"},
			want: "BAZ=******",
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}
