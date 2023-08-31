package t_logic

import "go_tetris/figure_struct"

/*
На вход мы получаем матрицу-поле и координаты фигуры(матрица не содержит падающую фигуру)
Если падающая фигура не соприкасается с фигурами поля, изменяются только координаты фигуры,
иначе текущая фигура прикрепляется к фигурам поля и запускается удаление заполненных слоёв
*/
func next_gen(field [20][10]int, figure figure_struct.Figure) [20][10]int {

}

/*
Функция проходит по слоям марицы (по горизонтали) и если встречает заполненный слой уничтожает его

На вход подаётс матрица с учётом, что координата (0,0) совпадает с field[0][0], возможно потребуется
перевернуть матрицу
*/
func find_ful_row(field [20][10]int) {
	nul_count := 0
	null_row := true
	for j := 0; j < 20; j++ {
		for i := 0; i < 10; i++ {
			if field[i][j] != 0 {
				null_row = false
				break
			}
		}
		if null_row {
			nul_count++
		} else {
			if nul_count != 0 {
				field = del_ful_row(field, j, nul_count) //перемещает ряды начиная с j и до конца
				j -= nul_count
				nul_count = 0
			}
		}
	}
}

func del_ful_row(field [20][10]int, row int, count int) [20][10]int {
	for j := row; j < 20; j++ {
		for i := 0; i < 10; i++ {

		}
	}
}
