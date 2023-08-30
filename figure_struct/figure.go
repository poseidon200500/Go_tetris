package figure_struct

type Coords struct {
	x, y int
}

type Figure struct {
	f_type int
	cords  [4]Coords
}

/*
координаты относительно центра
1 тип -- I

	{-2,0}{-1,0}{0,0}{1,0}

2 тип -- J

		{1,0}
		{1,1}
		{1,2}
	{0,3}{1,3}

3 тип -- L

	{0,0}
	{0,1}
	{0,2}
	{0,3}{1,3}

4 тип -- O

	{0,0}{1,0}
	{0,1}{1,1}

5 тип -- S

		{0,0}{1,0}
	{-1,1}{0,1}

6 тип -- T

	{-1,0}{0,0}{1,0}
		{0,1}

7 тип -- Z

	{-1,0}{0,0}
	{0,1}{1,1}
*/
func Make_Coods(cor [4][2]int) [4]Coords {
	result := [4]Coords{}
	for i := 0; i < 4; i++ {
		result[i] = Coords{cor[i][0], cor[0][1]}
	}
	return result
}

func Make_figure(input_type int) Figure {
	f_cord := [7][4][2]int{
		{{3, -1}, {4, -1}, {5, -1}, {6, -1}},
		{{5, -3}, {5, -2}, {5, -1}, {4, -1}},
		{{4, -3}, {4, -2}, {4, -1}, {5, -1}},
		{{4, -2}, {5, -2}, {4, -1}, {5, -1}},
		{{5, -2}, {6, -2}, {4, -1}, {5, -1}},
		{{4, -2}, {5, -2}, {6, -2}, {5, -1}},
		{{4, -2}, {5, -2}, {5, -1}, {6, -1}}}
	return Figure{f_type: input_type, cords: Make_Coods(f_cord[input_type-1])}
}
