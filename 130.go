package main


func look(m [][]byte, result [][]int, v [][]bool, x, y int) (bool, [][]int) {
    if x == 0 || x == len(m)-1 || y == 0 || y == len(m[0])-1 {
        if len(result) > 0 {
            return m[x][y] == 'X', result
        }
    }

    f := func(a, b int) bool {
        if a < 0 || a > len(m)-1 || b < 0 || b > len(m[0])-1 {
            return false
        }
        v[a][b] = true
        if m[a][b] == byte('O') && v[a][b] == false {
            result = append(result, []int{a, b})
            println(a, b)
            
            var ok bool
            ok, result = look(m, result, v, a, b)
            return ok
        }
        return true
    }
    return f(x+1, y) && f(x-1, y) && f(x, y+1) && f(x, y-1), result
}

func solve(board [][]byte) {
    if len(board) == 0 || len(board[0]) == 0 {
        return
    }


    v := make([][]bool, len(board))
    result := [][]int{}
    for i := range board {
        v[i] = make([]bool, len(board[0]))
    }

    for i := range board {
        for j := range board[0] {
            if board[i][j] == 'O' && v[i][j] == false {
                println(i,j, 'f')
                result = append(result, []int{i,j})
                v[i][j]=true
                
                if ok, result := look(board, result, v, i, j); ok && len(result) > 0 {
                    for i := range result {
                        board[result[i][0]][result[i][1]] = byte('X')
                    }
                }
                result = [][]int{}
            }
        }
    }
}