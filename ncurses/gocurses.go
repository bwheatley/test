package gocurses

/*
#define NCURSES_ENABLE_STDBOOL_H 0
#include <ncurses.h>
#include <stdlib.h>

int Printw (const char* str) { return printw (str); }

*/
import "C"

import (
    "fmt";
    "unsafe";
)

func Hello (s string) {
    C.initscr ();

    p := C.CString (fmt.Sprintf ("Hello, %s", s));
    C.Printw (p);
    C.free (unsafe.Pointer (p));

    C.refresh ();
    C.getch ();
    C.endwin ()


}
