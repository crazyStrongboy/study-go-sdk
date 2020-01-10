package main

/*
@Time : 2020/1/10 20:25
@Author : hejun
*/

func main() {
	s := make([]int, 0, 1)
	s = append(s, 5, 6)
}

/**
go tool compile -S -l -N main.go

0x0000 00000 (main.go:8)        TEXT    "".main(SB), ABIInternal, $104-0
0x0000 00000 (main.go:8)        MOVQ    TLS, CX
0x0009 00009 (main.go:8)        MOVQ    (CX)(TLS*2), CX
0x0010 00016 (main.go:8)        CMPQ    SP, 16(CX)
0x0014 00020 (main.go:8)        JLS     194
0x001a 00026 (main.go:8)        SUBQ    $104, SP
0x001e 00030 (main.go:8)        MOVQ    BP, 96(SP)
0x0023 00035 (main.go:8)        LEAQ    96(SP), BP
0x0028 00040 (main.go:8)        FUNCDATA        $0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
0x0028 00040 (main.go:8)        FUNCDATA        $1, gclocals·54241e171da8af6ae173d69da0236748(SB)
0x0028 00040 (main.go:8)        FUNCDATA        $3, gclocals·bfec7e55b3f043d1941c093912808913(SB)
0x0028 00040 (main.go:9)        PCDATA  $2, $0
0x0028 00040 (main.go:9)        PCDATA  $0, $0
0x0028 00040 (main.go:9)        MOVQ    $0, ""..autotmp_1+64(SP)
0x0031 00049 (main.go:9)        PCDATA  $2, $1
0x0031 00049 (main.go:9)        LEAQ    ""..autotmp_1+64(SP), AX
0x0036 00054 (main.go:9)        TESTB   AL, (AX)
0x0038 00056 (main.go:9)        PCDATA  $2, $-2
0x0038 00056 (main.go:9)        PCDATA  $0, $-2
0x0038 00056 (main.go:9)        JMP     58
0x003a 00058 (main.go:9)        JMP     60
0x003c 00060 (main.go:9)        PCDATA  $2, $1
0x003c 00060 (main.go:9)        PCDATA  $0, $0
0x003c 00060 (main.go:9)        MOVQ    AX, "".s+72(SP)
0x0041 00065 (main.go:9)        MOVQ    $0, "".s+80(SP)
0x004a 00074 (main.go:9)        MOVQ    $1, "".s+88(SP)
0x0053 00083 (main.go:10)       JMP     85
0x0055 00085 (main.go:10)       PCDATA  $2, $2
0x0055 00085 (main.go:10)       LEAQ    type.int(SB), CX
0x005c 00092 (main.go:10)       PCDATA  $2, $1
0x005c 00092 (main.go:10)       MOVQ    CX, (SP)
0x0060 00096 (main.go:10)       PCDATA  $2, $0
0x0060 00096 (main.go:10)       MOVQ    AX, 8(SP)
0x0065 00101 (main.go:10)       MOVQ    $0, 16(SP)
0x006e 00110 (main.go:10)       MOVQ    $1, 24(SP)
0x0077 00119 (main.go:10)       MOVQ    $2, 32(SP)
0x0080 00128 (main.go:10)       CALL    runtime.growslice(SB)
0x0085 00133 (main.go:10)       PCDATA  $2, $1
0x0085 00133 (main.go:10)       MOVQ    40(SP), AX
0x008a 00138 (main.go:10)       MOVQ    48(SP), CX
0x008f 00143 (main.go:10)       MOVQ    56(SP), DX
0x0094 00148 (main.go:10)       ADDQ    $2, CX
0x0098 00152 (main.go:10)       JMP     154
0x009a 00154 (main.go:10)       MOVQ    $5, (AX)
0x00a1 00161 (main.go:10)       MOVQ    $6, 8(AX)
0x00a9 00169 (main.go:10)       PCDATA  $2, $0
0x00a9 00169 (main.go:10)       MOVQ    AX, "".s+72(SP)
0x00ae 00174 (main.go:10)       MOVQ    CX, "".s+80(SP)
0x00b3 00179 (main.go:10)       MOVQ    DX, "".s+88(SP)
0x00b8 00184 (main.go:11)       MOVQ    96(SP), BP
0x00bd 00189 (main.go:11)       ADDQ    $104, SP
0x00c1 00193 (main.go:11)       RET
*/
