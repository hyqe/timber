package timber

import (
	"testing"

	"github.com/hyqe/assert"
)

func TestLevel_GTE(t *testing.T) {
	assert.Want(t, true, DEBUG.GTE(ALERT))
	assert.Want(t, false, SILENT.GTE(ALERT))
}

func TestLevel_GTE_DEBUG(t *testing.T) {
	lvl := DEBUG
	type args struct {
		lvl Level
	}
	tests := []struct {
		name string
		l    Level
		args args
		want bool
	}{
		{
			name: "vs DEBUG",
			l:    lvl,
			args: args{
				lvl: DEBUG,
			},
			want: true,
		},
		{
			name: "vs ERROR",
			l:    lvl,
			args: args{
				lvl: ERROR,
			},
			want: true,
		},
		{
			name: "vs ALERT",
			l:    lvl,
			args: args{
				lvl: ALERT,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.GTE(tt.args.lvl); got != tt.want {
				t.Errorf("Level.GTE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLevel_GTE_ERROR(t *testing.T) {
	lvl := ERROR
	type args struct {
		lvl Level
	}
	tests := []struct {
		name string
		l    Level
		args args
		want bool
	}{
		{
			name: "vs DEBUG",
			l:    lvl,
			args: args{
				lvl: DEBUG,
			},
			want: false,
		},
		{
			name: "vs ERROR",
			l:    lvl,
			args: args{
				lvl: ERROR,
			},
			want: true,
		},
		{
			name: "vs ALERT",
			l:    lvl,
			args: args{
				lvl: ALERT,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.GTE(tt.args.lvl); got != tt.want {
				t.Errorf("Level.GTE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLevel_GTE_ALERT(t *testing.T) {
	lvl := ALERT
	type args struct {
		lvl Level
	}
	tests := []struct {
		name string
		l    Level
		args args
		want bool
	}{
		{
			name: "vs DEBUG",
			l:    lvl,
			args: args{
				lvl: DEBUG,
			},
			want: false,
		},
		{
			name: "vs ERROR",
			l:    lvl,
			args: args{
				lvl: ERROR,
			},
			want: false,
		},
		{
			name: "vs ALERT",
			l:    lvl,
			args: args{
				lvl: ALERT,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.GTE(tt.args.lvl); got != tt.want {
				t.Errorf("Level.GTE() = %v, want %v", got, tt.want)
			}
		})
	}
}
