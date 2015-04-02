TEXT    ·Add+0(SB),$0-24
MOVQ    a+0(FP),BX
MOVQ    b+8(FP),BP
ADDQ    BP,BX
MOVQ    BX,res+16(FP)
RET     ,
