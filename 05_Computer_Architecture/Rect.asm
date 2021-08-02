@0      // a=0
D=M     // D=RAM[0] = 高さ
@23     // a = 23
D;JLE   // D<=0? jump to 23 else next, 高さ分表示したか
@16     // a = 16
M=D     // ram[16]=D=高さ
@16384  // a=screen
D=A     // D = screen
@17     // a = 17
M=D     // ram[17] = D = screen
@17     // a = 17
A=M     // a = ram[17] = screen
M=-1    // ram[screen] = -1
@17
D=M
@32
D=D+A
@17
M=D
@16
MD=M-1
@10
D;JGT
@23
0;JMP
