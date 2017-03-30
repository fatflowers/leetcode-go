package main
import "testing"

func TestSolve(t *testing.T) {
    /*
    board := [][]byte{{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'O', 'X', 'X'}}
    t.Log("board", board)
    solve(board)
    t.Log("after", board)

    board = [][]byte{{'X'}, {'O'}}
    t.Log("board", board)
    solve(board)
    t.Log("after", board)

    board := [][]byte{{'X', 'O', 'X'}, {'O', 'X', 'O'}, {'X', 'O', 'X'}}
    t.Log("board", board)
    solve(board)
    t.Log("after", board)
*/


    board := [][]byte{
        {'O', 'O', 'O', 'O', 'X', 'X'}, 
        {'O', 'O', 'O', 'O', 'O', 'O'}, 
        {'O', 'X', 'O', 'X', 'O', 'O'},
        {'O', 'X', 'O', 'O', 'X', 'O'},
        {'O', 'X', 'O', 'X', 'O', 'O'},
        {'O', 'X', 'O', 'O', 'O', 'O'},
    }
    t.Log("board", board)
    solve(board)
    t.Log("after", board)
}