package main

/*
@Time : 2020/1/10
@Author : hejun
*/
func main() {
	s := make([]int, 0, 1)
	s = append(s, 5)
}

/**
go tool compile -S -l -N main.go

0x0000 00000 (main.go:7)        TEXT    "".main(SB), NOSPLIT|ABIInternal, $40-0
0x0000 00000 (main.go:7)        SUBQ    $40, SP
0x0004 00004 (main.go:7)        MOVQ    BP, 32(SP)
0x0009 00009 (main.go:7)        LEAQ    32(SP), BP
0x000e 00014 (main.go:7)        FUNCDATA        $0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
0x000e 00014 (main.go:7)        FUNCDATA        $1, gclocals·54241e171da8af6ae173d69da0236748(SB)
0x000e 00014 (main.go:7)        FUNCDATA        $3, gclocals·9fb7f0986f647f17cb53dda1484e0f7a(SB)
0x000e 00014 (main.go:8)        PCDATA  $2, $0
0x000e 00014 (main.go:8)        PCDATA  $0, $0
0x000e 00014 (main.go:8)        MOVQ    $0, ""..autotmp_1(SP)
0x0016 00022 (main.go:8)        PCDATA  $2, $1
0x0016 00022 (main.go:8)        LEAQ    ""..autotmp_1(SP), AX
0x001a 00026 (main.go:8)        TESTB   AL, (AX)
0x001c 00028 (main.go:8)        JMP     30
0x001e 00030 (main.go:8)        PCDATA  $2, $-2
0x001e 00030 (main.go:8)        PCDATA  $0, $-2
0x001e 00030 (main.go:8)        JMP     32
0x0020 00032 (main.go:8)        PCDATA  $2, $1
0x0020 00032 (main.go:8)        PCDATA  $0, $0
0x0020 00032 (main.go:8)        MOVQ    AX, "".s+8(SP)
0x0025 00037 (main.go:8)        MOVQ    $0, "".s+16(SP)
0x002e 00046 (main.go:8)        MOVQ    $1, "".s+24(SP)
0x0037 00055 (main.go:9)        JMP     57
0x0039 00057 (main.go:9)        MOVQ    $5, ""..autotmp_1(SP)
0x0041 00065 (main.go:9)        PCDATA  $2, $0
0x0041 00065 (main.go:9)        MOVQ    AX, "".s+8(SP)
0x0046 00070 (main.go:9)        MOVQ    $1, "".s+16(SP)
0x004f 00079 (main.go:9)        MOVQ    $1, "".s+24(SP)
0x0058 00088 (main.go:10)       MOVQ    32(SP), BP
0x005d 00093 (main.go:10)       ADDQ    $40, SP
0x0061 00097 (main.go:10)       RET
*/
