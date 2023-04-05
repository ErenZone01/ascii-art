go run . "" | cat -e
go run . "\n" | cat -e
go run . "hello" | cat -e
go run . "HELLO" | cat -e
go run . "HeLlo HuMaN" | cat -e
go run . "1Hello 2There" | cat -e
go run . "Hello\nThere" | cat -e
go run . "Hello\n\nThere" | cat -e
go run . "{Hello & There #}" | cat -e
go run . 'hello There 1 to 2!' | cat -e
go run . "MaD3IrA&LiSboN" | cat -e
go run . "1a\"#FdwHywR&/()=" | cat -e
go run . "{|}~" | cat -e
go run . "[\]^_ 'a" | cat -e
go run . "RGB" | cat -e
go run . ":;<=>?@" | cat -e
go run . '\!" #$%&'"'"'()*+,-./' | cat -e
go run . "ABCDEFGHIJKLMNOPQRSTUVWXYZ" | cat -e
go run . "abcdefghijklmnopqrstuvwxyz" | cat -e
go run . "abcdeWXC" | cat -e
go run . "abcde 12" | cat -e
go run . "A@&!" | cat -e
go run . "ab  2!%AZE" | cat -e