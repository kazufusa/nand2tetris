// function Sys.init 0
(Sys.init)

// call Sys.main 0
@retAddr1
D=A
@SP
A=M
M=D
@SP
M=M+1

@LCL
D=M
@SP
A=M
M=D
@SP
M=M+1

@ARG
D=M
@SP
A=M
M=D
@SP
M=M+1

@THIS
D=M
@SP
A=M
M=D
@SP
M=M+1

@THAT
D=M
@SP
A=M
M=D
@SP
M=M+1


@SP
D=M
D=D-1
D=D-1
D=D-1
D=D-1
D=D-1

@ARG
M=D

@SP
D=M
@LCL
M=D

@Sys.main
0;JMP
(retAddr1)

// pop temp 1
@SP
M=M-1
A=M
D=M
@R6
M=D

// label LOOP
(LOOP)

// goto LOOP
@LOOP
0;JMP

// function Sys.main 0
(Sys.main)

// push constant 123
@123
D=A
@SP
A=M
M=D
@SP
M=M+1

// call Sys.add12 1
@retAddr2
D=A
@SP
A=M
M=D
@SP
M=M+1

@LCL
D=M
@SP
A=M
M=D
@SP
M=M+1

@ARG
D=M
@SP
A=M
M=D
@SP
M=M+1

@THIS
D=M
@SP
A=M
M=D
@SP
M=M+1

@THAT
D=M
@SP
A=M
M=D
@SP
M=M+1


@SP
D=M
D=D-1
D=D-1
D=D-1
D=D-1
D=D-1
D=D-1

@ARG
M=D

@SP
D=M
@LCL
M=D

@Sys.add12
0;JMP
(retAddr2)

// pop temp 0
@SP
M=M-1
A=M
D=M
@R5
M=D

// push constant 246
@246
D=A
@SP
A=M
M=D
@SP
M=M+1

// return
@LCL
D=M
@R13
M=D

@SP
M=M-1
A=M
D=M
@ARG
A=M
M=D

@ARG
D=M+1
@SP
M=D

@R13
M=M-1
D=M
@THAT
M=D

@R13
M=M-1
D=M
@THIS
M=D

@R13
M=M-1
D=M
@ARG
M=D

@R13
M=M-1
D=M
@LCL
M=D

@R13
M=M-1
A=M
0;JMP

// function Sys.add12 3
(Sys.add12)
@SP
A=M
M=0
@SP
M=M+1
@SP
A=M
M=0
@SP
M=M+1
@SP
A=M
M=0
@SP
M=M+1

// push argument 0
@0
D=A
@ARG
A=D+M
D=M
@SP
A=M
M=D
@SP
M=M+1

// push constant 12
@12
D=A
@SP
A=M
M=D
@SP
M=M+1

// add
@SP
M=M-1
A=M
D=M
@SP
A=M-1
D=D+M
@SP
A=M-1
M=D

// return
@LCL
D=M
@R13
M=D

@SP
M=M-1
A=M
D=M
@ARG
A=M
M=D

@ARG
D=M+1
@SP
M=D

@R13
M=M-1
D=M
@THAT
M=D

@R13
M=M-1
D=M
@THIS
M=D

@R13
M=M-1
D=M
@ARG
M=D

@R13
M=M-1
D=M
@LCL
M=D

@R13
M=M-1
A=M
0;JMP
