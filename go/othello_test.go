package othello

import (
	"reflect"
	"testing"
	"google.golang.org/appengine/aetest"
)

func TestMinMax(t *testing.T){
	ctx, done, _ := aetest.NewContext()
	defer done()
	board := Board{
		Pieces: [8][8]Piece{
			{2, 2, 2, 2, 2, 2, 2, 2},
			{2, 2, 2, 2, 2, 2, 2, 2},
			{2, 2, 2, 2, 2, 2, 2, 2},
			{1, 2, 1, 2, 2, 2, 2, 1},
			{1, 2, 1, 2, 2, 2, 2, 1},
			{1, 2, 1, 2, 2, 2, 2, 1},
			{1, 2, 1, 2, 2, 2, 2, 1},
			{1, 0, 0, 0, 0, 0, 0, 1},
		},
		Next: Black,
	}
	var move Move 
	_, bestMove := board.Score(ctx, 3, move)
	want := Move{
		Where: Position{2, 8},
		As: Black,
	}
	if !reflect.DeepEqual(bestMove, want) {
		t.Errorf("bestMove = %v but want %v", bestMove, want)
	}
}
