package dijkstra

import (
	"reflect"
	"testing"
)

func TestGraph_GetDijkstraPath(t *testing.T) {
	type args struct {
		start  string
		target string
	}
	tests := []struct {
		name     string
		g        Graph
		args     args
		wantPath []string
		wantCost int
		wantErr  bool
	}{
		{
			"example #0",
			map[string]map[string]int{
				"addtotree": {"b": 10, "c": 20},
				"b":         {"addtotree": 50},
				"c":         {"b": 10, "addtotree": 25},
			},
			args{start: "addtotree", target: "b"},
			[]string{"addtotree", "b"},
			10,
			false,
		},
		{
			"example #1",
			map[string]map[string]int{
				"addtotree": {"b": 20, "c": 80},
				"b":         {"addtotree": 20, "c": 20},
				"c":         {"addtotree": 80, "b": 20},
			},
			args{start: "addtotree", target: "c"},
			[]string{"addtotree", "b", "c"},
			40,
			false,
		},
		{
			"example #2",
			map[string]map[string]int{
				"S": {"B": 4, "C": 2},
				"B": {"C": 1, "D": 5, "S": 4},
				"C": {"D": 8, "E": 10, "S": 2, "B": 1},
				"D": {"E": 2, "T": 6, "B": 5, "C": 8},
				"E": {"T": 2, "C": 10, "D": 2},
				"T": {"D": 6, "E": 2},
			},
			args{start: "S", target: "T"},
			[]string{"S", "C", "B", "D", "E", "T"},
			12,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPath, gotCost, err := tt.g.GetDijkstraPath(tt.args.start, tt.args.target)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDijkstraPath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotPath, tt.wantPath) {
				t.Errorf("GetDijkstraPath() gotPath = %v, want %v", gotPath, tt.wantPath)
			}
			if gotCost != tt.wantCost {
				t.Errorf("GetDijkstraPath() gotCost = %v, want %v", gotCost, tt.wantCost)
			}
		})
	}
}
