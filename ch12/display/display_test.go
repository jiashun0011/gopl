package display

import (
	"io"
	"net"
	"os"
	"reflect"
	"sync"
	"testing"

	"gopl/ch7/eval"
)

func Example_expr() {
	e, _ := eval.Parse("sqrt(A / pi)")
	Display("e", e)
	// Output:
	// Display e (eval.call):
	// e.fn = "sqrt"
	// e.args[0].type = eval.binary
	// e.args[0].value.op = 47
	// e.args[0].value.x.type = eval.Var
	// e.args[0].value.x.value = "A"
	// e.args[0].value.y.type = eval.Var
	// e.args[0].value.y.value = "pi"
}

func Example_slice() {
	Display("slice", []*int{new(int), nil})
	// Output:
	// Display slice ([]*int):
	// (*slice[0]) = 0
	// slice[1] = nil
}

func Example_nilInterface() {
	var w io.Writer
	Display("w", w)
	// Output:
	// Display w (<nil>):
	// w = invalid
}

func Example_ptrToInterface() {
	var w io.Writer
	Display("&w", &w)
	// Output:
	// Display &w (*io.Writer):
	// (*&w) = nil
}

func Example_sturct() {
	Display("x", struct{ x interface{} }{3})
	// Output:
	// Display x (struct { x interface {} }):
	// x.x.type = int
	// x.x.value = 3
}

func Example_interface() {
	var i interface{} = 3
	Display("i", i)
	// Output:
	// Display i (int):
	// i = 3
}

func Example_ptrToInterface2() {
	var i interface{} = 3
	Display("&i", &i)
	// Output:
	// Display &i (*interface {}):
	// (*&i).type = int
	// (*&i).value = 3
}

func Example_array() {
	Display("x", [1]interface{}{3})
	// Output:
	// Display x ([1]interface {}):
	// x[0].type = int
	// x[0].value = 3
}

func Example_movie() {
	type Movie struct {
		Title, Subtitle string
		Year            int
		Color           bool
		Actor           map[string]string
		Oscars          []string
		Sequel          *string
	}

	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
		Color:    false,
		Actor: map[string]string{
			"Dr. Strangelove":            "Peter Sellers",
			"Grp. Capt. Lionel Mandrake": "Peter Sellers",
			"Pres. Merkin Muffley":       "Peter Sellers",
			"Gen. Buck Turgidson":        "George C. Scott",
			"Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
			`Maj. T.J. "King" Kong`:      "Slim Pickens",
		},

		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
	}

	Display("strangelove", strangelove)

	/*
		//!+output
		Display strangelove (display.Movie):
		strangelove.Title = "Dr. Strangelove"
		strangelove.Subtitle = "How I Learned to Stop Worrying and Love the Bomb"
		strangelove.Year = 1964
		strangelove.Color = false
		strangelove.Actor["Gen. Buck Turgidson"] = "George C. Scott"
		strangelove.Actor["Brig. Gen. Jack D. Ripper"] = "Sterling Hayden"
		strangelove.Actor["Maj. T.J. \"King\" Kong"] = "Slim Pickens"
		strangelove.Actor["Dr. Strangelove"] = "Peter Sellers"
		strangelove.Actor["Grp. Capt. Lionel Mandrake"] = "Peter Sellers"
		strangelove.Actor["Pres. Merkin Muffley"] = "Peter Sellers"
		strangelove.Oscars[0] = "Best Actor (Nomin.)"
		strangelove.Oscars[1] = "Best Adapted Screenplay (Nomin.)"
		strangelove.Oscars[2] = "Best Director (Nomin.)"
		strangelove.Oscars[3] = "Best Picture (Nomin.)"
		strangelove.Sequel = nil
		//!-output
	*/
}

func Test(t *testing.T) {
	Display("os.Stderr", os.Stderr)

	var w io.Writer = os.Stderr
	Display("&w", &w)

	var locker sync.Locker = new(sync.Mutex)
	Display("(&locker)", &locker)

	Display("locker", locker)

	locker = nil
	Display("(&locker)", &locker)

	ips, _ := net.LookupHost("golang.org")
	Display("ips", ips)

	Display("rV", reflect.ValueOf(os.Stderr))

	type P *P
	var p P
	p = &p
	if false {
		Display("p", p)
	}

	type M map[string]M
	m := make(M)
	m[""] = m
	if false {
		Display("m", m)
	}

	type S []S
	s := make(S, 1)
	s[0] = s
	if false {
		Display("s", s)
	}

	type Cycle struct {
		Value int
		Tail  *Cycle
	}
	var c Cycle
	c = Cycle{42, &c}
	if false {
		Display("c", c)
	}
}
