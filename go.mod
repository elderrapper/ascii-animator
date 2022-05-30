module github.com/davidhsingyuchen/ascii-animator

go 1.17

require (
	github.com/leaanthony/go-ansi-parser v1.4.0
	github.com/rivo/uniseg v0.2.0 // indirect
)

// TODO: Remove this after https://github.com/leaanthony/go-ansi-parser/issues/3 is fixed.
replace github.com/leaanthony/go-ansi-parser v1.4.0 => github.com/davidhsingyuchen/go-ansi-parser v1.4.1 // indirect
