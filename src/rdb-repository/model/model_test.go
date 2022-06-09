package model

import (
	"testing"
	"time"
)

func TestCurrentTime(t *testing.T) {
	type test struct {
		testcase string
	}

	do := func(tt *test) {
		t.Run(tt.testcase, func(t *testing.T) {
			want := time.Now().Round(time.Second)
			got := currentTime()
			if want != got {
				// NOTE 稀に失敗する可能性あり
				t.Fatalf("want=%v, got=%v.", want, got)
			}
		})
	}

	tests := []*test{
		{},
		{},
		{},
	}

	for _, tt := range tests {
		do(tt)
	}
}
