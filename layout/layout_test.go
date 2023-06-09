package layout

import (
	"reflect"
	"testing"
)

func TestLayout(t *testing.T) {
	type args struct {
		blocks []Block
	}
	tests := []struct {
		name string
		args args
		want []LayoutResult
	}{
		{
			name: "test 1",
			args: args{
				blocks: []Block{
					{
						Id: 738,
						Form: [][]int64{
							{1, 0},
							{1, 1},
						},
					},
					{
						Id: 841,
						Form: [][]int64{
							{1, 1},
							{0, 1},
						},
					},
				},
			},
			want: []LayoutResult{
				{
					BlockId:   738,
					Position:  1,
					IsRotated: false,
				},
				{
					BlockId:   841,
					Position:  2,
					IsRotated: false,
				},
			},
		},
		{
			name: "test 2",
			args: args{
				blocks: []Block{
					{
						Id: 443,
						Form: [][]int64{
							{1, 0, 1},
							{1, 1, 1},
						},
					},
					{
						Id: 327,
						Form: [][]int64{
							{0, 1, 0},
							{1, 1, 1},
							{1, 1, 1},
							{1, 1, 0},
							{0, 1, 0},
						},
					},
					{
						Id: 891,
						Form: [][]int64{
							{0, 0, 1},
							{1, 0, 1},
							{1, 1, 1},
						},
					},
				},
			},
			want: []LayoutResult{
				{
					BlockId:   443,
					Position:  1,
					IsRotated: false,
				},
				{
					BlockId:   327,
					Position:  2,
					IsRotated: true,
				},
				{
					BlockId:   891,
					Position:  3,
					IsRotated: true,
				},
			},
		},
		{
			name: "test 3",
			args: args{
				blocks: []Block{
					{
						Id: 4892,
						Form: [][]int64{
							{0, 0, 1},
							{1, 0, 1},
							{1, 1, 1},
							{1, 1, 1},
							{1, 1, 1},
							{1, 1, 1},
							{1, 1, 1},
							{1, 1, 1},
						},
					},
					{
						Id: 1839,
						Form: [][]int64{
							{1, 1, 1},
							{1, 1, 1},
							{1, 1, 1},
							{1, 1, 1},
							{1, 0, 0},
						},
					},
					{
						Id: 8183,
						Form: [][]int64{
							{0, 1, 1},
							{1, 1, 1},
							{1, 1, 1},
							{1, 1, 0},
							{0, 1, 0},
						},
					},
				},
			},
			want: []LayoutResult{
				{
					BlockId:   4892,
					Position:  1,
					IsRotated: false,
				},
				{
					BlockId:   8183,
					Position:  2,
					IsRotated: false,
				},
				{
					BlockId:   1839,
					Position:  3,
					IsRotated: false,
				},
			},
		},
		{
			name: "test 4",
			args: args{
				blocks: []Block{
					{
						Id: 1,
						Form: [][]int64{
							{1, 0, 1},
							{1, 1, 1},
							{1, 1, 1},
						},
					},
					{
						Id: 2,
						Form: [][]int64{
							{0, 0, 1},
							{1, 1, 1},
							{1, 1, 1},
							{1, 1, 1},
							{1, 1, 1},
						},
					},
					{
						Id: 3,
						Form: [][]int64{
							{0, 1, 1},
							{1, 1, 1},
							{0, 1, 0},
						},
					},
				},
			},
			want: []LayoutResult{
				{
					BlockId:   1,
					Position:  1,
					IsRotated: false,
				},
				{
					BlockId:   3,
					Position:  2,
					IsRotated: false,
				},
				{
					BlockId:   2,
					Position:  3,
					IsRotated: true,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Layout(tt.args.blocks); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Layout() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fit(t *testing.T) {
	type args struct {
		low  [][]int64
		high [][]int64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "should fit",
			args: args{
				high: [][]int64{
					{0, 1, 1},
				},
				low: [][]int64{
					{1, 0, 0},
				},
			},
			want: true,
		},
		{
			name: "should fit to fully closed part",
			args: args{
				high: [][]int64{},
				low:  [][]int64{},
			},
			want: true,
		},
		{
			name: "should not fit",
			args: args{
				high: [][]int64{
					{0, 1, 1},
					{0, 1, 0},
				},
				low: [][]int64{
					{1, 0, 1},
					{1, 0, 1},
				},
			},
			want: false,
		},
		{
			name: "should not fit - have different height of open part",
			args: args{
				high: [][]int64{
					{0, 1, 0},
				},
				low: [][]int64{
					{1, 0, 0},
					{1, 0, 1},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fit(tt.args.low, tt.args.high); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_openBlock(t *testing.T) {
	type args struct {
		form [][]int64
		dir  Direction
	}
	tests := []struct {
		name string
		args args
		want [][]int64
	}{
		{
			name: "should return empty slice - top closed",
			args: args{
				form: [][]int64{
					{1, 1, 1},
					{1, 1, 1},
				},
				dir: Top,
			},
			want: [][]int64{},
		},
		{
			name: "should return empty slice - bottom closed",
			args: args{
				form: [][]int64{
					{1, 1, 1},
					{1, 1, 1},
				},
				dir: Bottom,
			},
			want: [][]int64{},
		},
		{
			name: "should return open bottom part",
			args: args{
				form: [][]int64{
					{1, 1, 1},
					{1, 0, 1},
					{1, 0, 1},
				},
				dir: Bottom,
			},
			want: [][]int64{
				{1, 0, 1},
				{1, 0, 1},
			},
		},
		{
			name: "should return open top part",
			args: args{
				form: [][]int64{
					{1, 0, 1},
					{1, 0, 1},
					{1, 1, 1},
				},
				dir: Top,
			},
			want: [][]int64{
				{1, 0, 1},
				{1, 0, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := openBlock(tt.args.form, tt.args.dir); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("openBlock() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rotate(t *testing.T) {
	type args struct {
		form [][]int64
	}
	tests := []struct {
		name string
		args args
		want [][]int64
	}{
		{
			name: "test 1",
			args: args{
				form: [][]int64{
					{1, 1},
					{1, 0},
				},
			},
			want: [][]int64{
				{0, 1},
				{1, 1},
			},
		},
		{
			name: "test 2",
			args: args{
				form: [][]int64{
					{1, 0, 1},
					{1, 1, 0},
				},
			},
			want: [][]int64{
				{0, 1, 1},
				{1, 0, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotate(tt.args.form); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rotate() = %v, want %v", got, tt.want)
			}
		})
	}
}
