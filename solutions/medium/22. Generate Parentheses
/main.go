
// 3 2 1 3 2 1
// 3 2 3 1 2 1
// 3 2 2 1 1 1
// 3 1 2 1 1 1
// 3 1 2 1 1 1
/*
   array for opener[n]
   array for closer[0]
   open() {
       append closer
       remove opener
   }
*/

type element struct {
	opener, closer int
	text           string
}

func generateParenthesis(n int) []string {
	e := element{
		opener: n,
		closer: 0,
		text:   "",
	}

	return appending(e)
}

func appending(e element) []string {
	fmt.Println(e)
	if e.closer == 0 && e.opener == 0 {
		return []string{e.text}
	} else {
		var res []string
		if e.opener > 0 {
			fmt.Println("opener: ")
			a := open(e)
			res = append(res, appending(a)...)
		}
		if e.closer > 0 {
			fmt.Println("closer: ")
			b := close(e)
			res = append(res, appending(b)...)
		}
		return res
	}
}

func open(e element) element {
	e.opener -= 1
	e.closer += 1
	e.text += string('(')
	return e
}

func close(e element) element {
	e.closer -= 1
	e.text += string(')')
	return e
}
