# **`Java` GRAMMAR**
### **AUTHORSHIP INFORMATION**
#### *Authors :* Brynn Harrington and Emily Hoppe Copyright (C) 2021
#### *Adapted from :* Aaron Moss"s [`Java` Egg Grammar](https:github.com/bruceiv/egg/blob/deriv/grammars/Java-u.egg) and Roman Reziejowski's [`Java` Mouse Parser-Generator](http://home.swipnet.se/redz/roman)
#### *Creation Date :* June 11, 2021 
#### *Last Modified :* June 22, 2021
#### *Copyright and Licensing Information :* See end of file.

###  **GENERAL DESCRIPTION**
A modification of `Java` [Egg](https:github.com/bruceiv/egg/blob/deriv/grammars/Java-u.egg) parsing grammar ported into GoGLL to test `Java` input files.

### **STATUS ON GRAMMAR**
#### *Markdown File Creation:* Not working 
#### *Parser Generated :* Incomplete
#### *Test File Creation:* Incomplete
#### *Testing Results:* Unknown

### **`Java` GRAMMAR GUIDE**
The following grammar tests input files from the `Java` language.
```
package "Java"
```
#### ***Compilation Unit***
```
CompUnit          : WS OptPackDecl RepImpDecl0x RepSemiModDecl0x ;
      OptPackDecl : PackDecl                      
                  / empty                             ;
     RepImpDecl0x : ImportDecl RepImpDecl0x
                  / empty                             ; 
 RepSemiModDecl0x : SemiModDecl RepSemiModDecl0x
                  / empty                             ;
      SemiModDecl :  SEMI 
                  | RepModif0 DeclAlts                ;
      DeclAlts    : ClsDecl  
                  | IntfDecl                          ;

PackDecl          : PACKAGE QualifiedID SEMI          ;

ImportDecl        : IMPORT OptStatic QualifiedID OptDotStar SEMI ;
      OptDotStar  : DOT STAR 
                  / empty                     ;
```
#### ***Class Declarations***
- Note: The following are the representations of the 
    MemDecl (Member Decl): 
    - Type ID FormalParams RepDim0x OptThrowClsTypLst MemAlts = Method
    - VOID ID FormalParams OptThrowClsTypLst MemAlts = Void Method
    - ID FormalParams OptThrowClsTypLst Block = Constructor
    - IntfDecl = Interface
    - ClsDecl = Class
    - Type VarDecl RepComVDecl0x = Field
- Note: The following are the representations of the 
     ClsBdyDecl (ClassBodyDeclaration): 
    - SEMI = Semicolon
    - OptStatic Block  = Static or Instance Initializer
    - RepModif0 MemDecl = Class Member Declaration  
```
ClsDecl           : Cls ID OptExtClsType OptImpClsLst ClsBdy ;
    OptExtClsType : EXTENDS ClsType              
                  / empty                             ;
     OptImpClsLst : IMPLEMENTS ClsTypeList 
                  / empty       ;
      
ClsBdy            : LWING RepClsBDecl0x RWING         ;
    RepClsBDecl0x : ClsBdyDecl RepClsBDecl0x
                  / empty                             ;            

ClsBdyDecl        : SEMI
                  | OptStatic Block 
                  | RepModif0 MemDecl                 ;
      OptStatic : STATIC
                  / empty                             ; 

MemDecl           : Type ID FormalParams RepDim0x OptThrowClsTypLst MemAlts
                  | VOID ID FormalParams OptThrowClsTypLst MemAlts
                  | ID FormalParams OptThrowClsTypLst Block
                  | IntfDecl 
                  | ClsDecl
                  | Type VarDecl RepComVDecl0x        ; 
      MemAlts     : SEMI 
                  | Block                             ;
```
#### ***Interface Declarations***
```
IntfDecl   : Intf ID OptExtendsClsLis IntfBdy         ;
      OptExtendsClsLis : EXTENDS ClsTypeList 
                        / empty      ;

IntfBdy           : LWING RepInBodDecl0x RWING        ;
   RepInBodDecl0x : IntfBdyDecl RepInBodDecl0x
                  / empty                             ;

IntfBdyDecl       : RepModif0 IntfMemDecl 
                  | SEMI                              ;

IntfMemDecl       : IntfMethFieldDecl
                  | VOID ID VoidIntfMethDeclRst
                  | IntfDecl
                  | ClsDecl                           ;

IntfMethFieldDecl: Type ID IntfMethFieldRest          ;

IntfMethFieldRest : ConstDeclsRest SEMI 
                  | IntfMethDeclRest                  ;

IntfMethDeclRest  : FormalParams RepDim0x OptThrowClsTypLst SEMI ;

VoidIntfMethDeclRst: FormalParams OptThrowClsTypLst SEMI;
 OptThrowClsTypLst: THROWS ClsTypeList
                  / empty                             ;

ConstDeclsRest    : ConstDeclRest RepComCnstDecl0x    ;
 RepComCnstDecl0x : COMMA ConstDecl RepComCnstDecl0x
                  / empty                             ;

ConstDecl         : ID ConstDeclRest                  ;

ConstDeclRest     : RepDim0x EQU VarInitial           ;

```
#### ***Variable Declarations***
```
LocalVarDeclStmt  : OptFinType VarDecl RepComVDecl0x SEMI ;

VarDecl           :  ID RepDim0x OptEqVarInit         ;
     OptEqVarInit : EQU  VarInitial
                  / empty                             ;
```
#### ***Formal Parameters***
```
FormalParams      : LPAR OptFormPDecl RPAR            ;
     OptFormPDecl : FormalParamDecls 
                  / empty                             ;

FormalParam       : OptFinType VarDelID               ;
 
FormalParamDecls  : OptFinType FormalParamDeclsRest   ;

FormalParamDeclsRest :  VarDelID OptComFormPDecl      ;
  OptComFormPDecl : COMMA FormalParamDecls 
                  / empty                             ;

VarDelID          : ID RepDim0x                       ;

```
#### ***Statements***
```
Block             : LWING RepBlkSt0x RWING            ;
     RepBlkSt0x   : BlockStmt RepBlkSt0x
                  / empty                             ;

BlockStmt         : LocalVarDeclStmt
                  | RepModif0 ClsDecl
                  | Stmt                              ;
   RepModif0      : Modifier RepModif0
                  / empty                             ;

Stmt              : Block
                  | ASSERT Expr OptColExpr SEMI
                  | IF ParExpr Stmt OptElse
                  | FOR LPAR OptForInit SEMI OptExpr SEMI OptForUpd RPAR Stmt
                  | WHILE ParExpr Stmt
                  | DO Stmt WHILE ParExpr SEMI
                  | TRY Block CatchBlk
                  | SWITCH ParExpr LWING RepSwBlkStmt0x RWING
                  | SYNCHRONIZED ParExpr Block
                  | RETURN OptExpr SEMI
                  | THROW Expr SEMI
                  | BREAK OptID SEMI
                  | CONTINUE OptID SEMI
                  | SEMI
                  | StmtExpr SEMI
                  | ID COLON Stmt                     ; 
      OptColExpr  : COLON Expr 
                  / empty                             ;
      OptElse     : ELSE Stmt  
                  / empty                             ;
      OptForInit  : ForInit  
                  / empty                             ;
      OptForUpd   : ForUpdate 
                  / empty                             ;
      OptExpr     : Expr 
                  / empty                             ;
      CatchBlk    : Catch RepCatch0x OptFinly  
                  | Finally                           ;
      RepCatch0x  : Catch RepCatch0x 
                  / empty                             ;
      OptFinly    : Finally 
                  / empty                             ;
   RepSwBlkStmt0x : SwitchBlockStmtGrp RepSwBlkStmt0x
                  / empty                             ;
      OptID       : ID 
                  / empty                             ;       

Catch             : CATCH LPAR FormalParam RPAR Block ;
   
Finally           : FINALLY Block                     ;

SwitchBlockStmtGrp: SwitchLabel RepBlkSt0x            ;


SwitchLabel       : CASE ConstExpr COLON 
                  | DEFAULT COLON                     ;

ForInit           : OptFinType Type VarDeclInit       
                  | StmtExpr RepComSExpr0x            ; 
   OptFinType     : FINAL Type                        ;
   OptFin         : FINAL 
                  / empty                             ;
   VarDeclInit    : VarDecl RepComVDecl0x             ; 
   RepComVDecl0x  : COMMA VarDecl RepComVDecl0x
                  / empty                             ;

ForUpdate         : StmtExpr RepComSExpr0x            ;
    RepComSExpr0x : COMMA StmtExpr RepComSExpr0x
                  / empty                             ;
```
#### ***Expressions***
- Note: Some of the shorthand names are:
    - Cond = Conditional
    - IOR = Inclusive-Or 
    - XOR = Exclusive-Or
- Note: The definition of Expr is part of the modification in 
JLS Ch. 18 to minimize look ahead. In JLS Ch. 15.27, Expr
is defined as AssignmentExpr, which is effectively defined as
(LeftHandSide AssignOp) * CondExpr. The above is obtained by 
allowing ANY CondExpr as LeftHandSide, which results in 
accepting Stmts like 5 : a.
```
StmtExpr          : Expr                              ;
   
ConstExpr         : Expr                              ;

Expr              : CondExpr RepAsscExpr0x            ;
    RepAsscExpr0x : AssignOp CondExpr RepAsscExpr0x
                  / empty                             ;

AssignOp          : EQU
                  | PLUS_EQU
                  | MINUS_EQU
                  | STAR_EQU
                  | DIV_EQU
                  | AND_EQU
                  | OR_EQU
                  | HAT_EQU
                  | MOD_EQU
                  | SL_EQU
                  | SR_EQU
                  | BSR_EQU                           ;

CondExpr          : CondORExpr RepCondition0x;
   RepCondition0x : QUERY Expr COLON CondORExpr RepCondition0x
                  / empty                             ;

CondORExpr        : CondANDExpr RepORcAND0x           ;
      RepORcAND0x : OR_OR CondANDExpr RepORcAND0x
                  / empty                             ;

CondANDExpr       : IORExpr RepANDIOR0x               ;
      RepANDIOR0x : AND_AND IORExpr RepANDIOR0x
                  / empty                             ;

IORExpr           : XORExpr RepORXOR0x                ; 
      RepORXOR0x  : ORXOR RepORXOR0x
                  / empty                             ;
      ORXOR       : OR XORExpr                        ;

XORExpr           : ANDExpr RepHatAND0x               ;
      RepHatAND0x : HAT ANDExpr RepHatAND0x
                  / empty                             ;

ANDExpr           : EqualExpr RepANDEq0x              ;
      RepANDEq0x  : AND EqualExpr RepANDEq0x
                  / empty                             ;

EqualExpr         : RelateExpr RepEqExpr0x            ;
      RepEqExpr0x : EqAlts RelateExpr RepEqExpr0x
                  / empty                             ;
      EqAlts      : EQUAL 
                  | NOT_EQUAL                         ;

RelateExpr        : ShiftExpr RepESInst0x             ;
      RepESInst0x : ESInst RepESInst0x
                  / empty                             ;
      ESInst      : EqShift 
                  | INSTANCEOF ReferenceType          ;
      EqShift     : EqCheck ShiftExpr                 ;
      EqCheck     : LE | GE | LT | GT                 ;

ShiftExpr   : AddExpr ShiftAlts                       ;
      CarrotAlts  : SL 
                  | SR 
                  | BSR                               ;
      ShiftAlts   : CarrotAlts AddExpr ShiftAlts
                  / empty                             ;


AddExpr           : MultExpr RepAddAltsMult0x         ;
 RepAddAltsMult0x : AddAlts MultExpr 
                  / empty                             ;
      AddAlts     : PLUS | MINUS                      ;

MultExpr          : UnaryExpr RepSDMUExpr0x           ;
    RepSDMUExpr0x : SDM UnaryExpr RepSDMUExpr0x
                  / empty                             ;
      SDM         : STAR | DIV | MOD                  ;

UnaryExpr         : PrefixOp UnaryExpr
                  | LPAR Type RPAR UnaryExpr
                  | Primary RepSel0x RepPfOp0x        ;
   RepSel0x       : Selector RepSel0x
                  / empty                             ;
   RepPfOp0x      : PostfixOp RepPfOp0x
                  / empty                             ;

Primary           : ParExpr
                  | THIS OptArgs
                  | SUPER SuperSuffix
                  | Literal
                  | NEW Creator
                  | QualifiedID OptIDSuff
                  | BasicType RepDim0x DOT Cls
                  | VOID DOT Cls                      ;            
      OptIDSuff   : IDSuffix 
                  / empty                             ;

IDSuffix          : LBRK RBRKAlts 
                  | Arguments 
                  | DOT OtherAlts                     ;
      RBRKAlts    : RBRK RepDim0x DOT Cls 
                  | Expr RBRK                         ; 
      OtherAlts   : Cls 
                  | THIS  
                  | SUPER Arguments 
                  | NEW InnerCreator                  ;

PrefixOp          : INC 
                  | DEC 
                  | BANG 
                  | TILDA 
                  | PLUS 
                  | MINUS                             ;     

PostfixOp         : INC 
                  | DEC                               ;

Selector          : DOT ID OptArgs
                  | DOT THIS
                  | DOT SUPER SuperSuffix
                  | DOT NEW InnerCreator
                  | DimExpr                           ;

SuperSuffix       : Arguments 
                  | DOT ID OptArgs;
      OptArgs     : Arguments
                  | empty                             ;

BasicType         : BasicTypeLit notLorD              ;
     BasicTypeLit : "byte"
                  | "short"
                  | "char"
                  | "int"
                  | "long"
                  | "float"
                  | "double"
                  | "boolean"                         ;

Arguments         : LPAR OptExprs RPAR                ;
      OptExprs    : Expr RepComExp0x
                  / empty                             ;
      RepComExp0x : COMMA Expr RepComExp0x
                  / empty                             ;

Creator           : CreatedName ClsCreatorRest 
                  | TypeAlts ArrayCreatorRest         ;

CreatedName       : ID RepDotID0x                     ;

InnerCreator      : ID ClsCreatorRest                 ;

ArrayCreatorRest  : LBRK ArrayRest                    ;
      ArrayRest   : RBRK RepDim0x ArrayInitializer 
                  | Expr RBRK RepDimExpr0x RepDim0x   ;
     RepDimExpr0x : DimExpr RepDimExpr0x
                  / empty                             ;
   
ClsCreatorRest    : Arguments OptClsBdy               ;
      OptClsBdy   : ClsBdy 
                  / empty                             ;

ArrayInitializer  : LWING OptVarInit RWING            ;
      OptVarInit  : VarInitial RepComInit0x OptCom
                  | empty                             ;
      RepComInit0x : COMMA VarInitial RepComInit0x 
                  / empty                             ;
      OptCom      : COMMA
                  | empty                             ;

VarInitial        : ArrayInitializer                
                  | Expr                              ;

ParExpr           : LPAR Expr RPAR                    ;

QualifiedID       : ID RepDotID0x                     ;

Dim               : LBRK RBRK                         ;
DimExpr           : LBRK Expr RBRK                    ; 
```
#### ***TYPES AND MODIFIERS***
```
Type              : TypeAlts RepDim0x                 ; 
      TypeAlts    : BasicType 
                  | ClsType                           ;

ReferenceType     : BasicType Dim RepDim0x 
                  | ClsType RepDim0x                  ;
      RepDim0x    : Dim RepDim0x
                  / empty                             ;

ClsType           : ID RepDotID0x                      ;
      RepDotID0x  : DOT ID RepDotID0x
                  / empty                             ;
ClsTypeList       : ClsType RepComCls0x               ;
   RepComCls0x    : COMMA ClsType RepComCls0x
                  / empty                             ;

Modifier          : Modifs notLorD                    ;
      Modifs      : "public"
                  | "protected"
                  | "private"
                  | "static"
                  | "abstract"
                  | "final"
                  | "native"
                  | "synchronized"
                  | "transient"
                  | "volatile"
                  | "strictfp"                        ;
```
#### ***Identifiers***
            NOTKEYWORD NEEDS FIXING
```

ID                : notKeyword LetterLorD             ;   

      LetterLorD  : Letter RepLorD0x WS               ;
      RepLorD0x   : LorD  RepLorD0x 
                  / empty                             ; 
      notKeyword  : not "abstract" 
                  not "assert" 
                  not "boolean"
                  not "break" 
                  not "byte" 
                  not "case" 
                  not "catch" 
                  not "char" 
                  not "Cls" 
                  not "continue" 
                  not "default" 
                  not "double" 
                  not "do" 
                  not "else" 
                  not "enum" 
                  not "extends" 
                  not "false"    
                  not "finally"  
                  not "final"    
                  not "float"    
                  not "for"      
                  not "if"       
                  not "implements"
                  not "import"   
                  not "Intf"
                  not "int"      
                  not "instanceof"
                  not "long"     
                  not "native"   
                  not "new"      
                  not "null"     
                  not "package"  
                  not "private"  
                  not "protected"
                  not "public"   
                  not "return"   
                  not "short"    
                  not "static"   
                  not "strictfp" 
                  not "super"    
                  not "switch"   
                  not "synchronized"
                  not "this"
                  not "throws"   
                  not "throw"    
                  not "transient"
                  not "true"     
                  not "try"      
                  not "void"     
                  not "volatile" 
                  not "while"                         ;

```
#### ***Keywords***
```
Keyword           : Words notLorD                     ;
      Words       : "abstract" 
                  | "assert"   
                  | "boolean"  
                  | "break"    
                  | "byte"     
                  | "case"     
                  | "catch"    
                  | "char"     
                  | "Cls"    
                  | "continue" 
                  | "default"  
                  | "double"   
                  | "do"       
                  | "else"     
                  | "enum"     
                  | "extends"  
                  | "false"    
                  | "finally"  
                  | "final"    
                  | "float"    
                  | "for"      
                  | "if"       
                  | "implements"
                  | "import"   
                  | "Intf"
                  | "int"      
                  | "instanceof"
                  | "long"     
                  | "native"   
                  | "new"      
                  | "null"     
                  | "package"  
                  | "private"  
                  | "protected"
                  | "public"   
                  | "return"   
                  | "short"    
                  | "static"   
                  | "strictfp" 
                  | "super"    
                  | "switch"   
                  | "synchronized"
                  | "this"
                  | "throws"   
                  | "throw"    
                  | "transient"
                  | "true"     
                  | "try"      
                  | "void"     
                  | "volatile" 
                  | "while"                           ;

ASSERT            : "assert"       notLorD            ; 
BREAK             : "break"        notLorD            ;
CASE              : "case"         notLorD            ;
CATCH             : "catch"        notLorD            ;
Cls               : "Cls"          notLorD            ;
CONTINUE          : "continue"     notLorD            ;
DEFAULT           : "default"      notLorD            ;
DO                : "do"           notLorD            ;
ELSE              : "else"         notLorD            ;
ENUM              : "enum"         notLorD            ;
EXTENDS           : "extends"      notLorD            ;
FINALLY           : "finally"      notLorD            ; 
FINAL             : "final"        notLorD            ;
FOR               : "for"          notLorD            ;
IF                : "if"           notLorD            ; 
IMPLEMENTS        : "implements"   notLorD            ; 
IMPORT            : "import"       notLorD            ;
Intf              : "Intf"         notLorD            ;
INSTANCEOF        : "instanceof"   notLorD            ;
NEW               : "new"          notLorD            ;
PACKAGE           : "package"      notLorD            ; 
RETURN            : "return"       notLorD            ;
STATIC            : "static"       notLorD            ; 
SUPER             : "super"        notLorD            ;
SWITCH            : "switch"       notLorD            ;
SYNCHRONIZED      : "synchronized" notLorD            ; 
THIS              : "this"         notLorD            ;
THROWS            : "throws"       notLorD            ;
THROW             : "throw"        notLorD            ;
TRY               : "try"          notLorD            ;
VOID              : "void"         notLorD            ;
WHILE             : "while"        notLorD            ; 
```
### ***General Literal Definition***
```
Literal           : LitAlts WS                         ;
      LitAlts     : FloatLiteral
                  | IntegerLiteral 
                  | CharLiteral
                  | StringLiteral
                  | "true"  notLorD
                  | "false" notLorD
                  | "null"  notLorD                   ;
```
### ***Basic Identifiers*** 
                  FIX notLorD VERY IMPORTANT
- Note: These are traditional definitions of letters and
digits. JLS defines letters and digits as Unicode characters recognized as such by special Java procedures, which is difficult to express in terms of Parsing Expressions.
```

notLorD           : not "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890" ;        
LorD              : Letter 
                  | num                               ; 
Letter            : let 
                  | WS                                ;
let               : letter                            ;
                  
```
### ***Character and String Literals***
- Note: Unicode escape is not defined in JLS syntax because 
unicode characters are processed very early.
```
CharLiteral       : sinQuo EscSlash sinQuo                   ;
      EscSlash    : Escape 
                  | EscUp                              ;
      EscUp       : "^" 
                  | sinQuo 
                  | dubSlash                           ;
      carrot      : '^'                                ;
                  
                  
StringLiteral     : dubQuo StrClose                    ;
      StrClose    : dubQuo 
                  / OptEsc StrClose                    ;
      dubQuo      : any "\""                           ;
      sinQuo      : '\''                               ;
      OptEsc      : Escape 
                  | anyCarSl                           ;
            anyCarSl : any "^\\" ;
Escape            : dubSlash Escs                      ;
      Escs        : esc
                  | OctalEscape 
                  | UnicodeEscape                      ;
      esc         : any "bfnrt\"'\\"                   ;
      dubSlash    : '\\' ;
   
UnicodeEscape : "u" hexDigit hexDigit hexDigit hexDigit ;
```
#### ***General Numeric Literals***
- Note: In IntegerLiteral, OctalNumeral may prefix 
HexNumeral and DecimalNumeral may prefix OctalNumeral
```

FloatLiteral      : HexFloat 
                  | DecimalFloat                      ;

IntegerLiteral    : NumeralAlts optOneL               ;
      NumeralAlts : HexNumeral 
                  | octalNumeral  
                  | DecimalNumeral                    ;
      optOneL     : [ any "1L" ]                      ;

DecimalFloat      :  repDig1x dot repDig0x optExpo fF_dD 
                  | dot RepDig1xExp 
                  | RepDig1xExp fF_dD
                  | RepDig1xOptExp fF_dD                   ;
   RepDig1xOptExp : repDig1x optExpo                       ;
      RepDig1xExp : repDig1x exponent                      ;
      optExpo     : [ any "eE" { number } [ any "+\\-" ] ] ;
      repDig1x    : < number >                             ;
      repDig0x    : { number }                             ;
      dot         : '.'                                    ;
      fF_dD       : [ any "fFdD" ]                         ;

```
#### ***BASE-SIXTEEN AND BASE-EIGHT LITERALS***
Incomplete decimalnumeral!!!
```

DecimalNumeral   : ze
                  | onenine repNumx0                ;  
    repNumx0      : { number }                      ;
    onenine       : any "123456789"                 ;
    ze            : '0'                             ;

HexFloat          : HexSignificand  Beoptfd         ;
      Beoptfd     : binaryExponent optfFdD          ;
      optfFdD     : [ any "fFdD" ]                  ;
      
HexSignificand    : HexNumeral OptDot 
                  | RepHex0xDot hexDigit repHex0x   ;
      RepHex0xDot : Any0xX repHex0x "."             ;
      OptDot      : "." 
                  / empty                           ; 
      
HexNumeral      : Any0xX repHex1x ; 
      Any0xX    : "0"
                | "x"
                | "X"                               ;
      repHex1x  : < < number any "abcdefABCDEF" > > ;
      repHex0x  : < number any "abcdefABCDEF" >     ;

hexDigit        : < number any "abcdefABCDEF" >     ;

octalNumeral    : '0' < any "01234567" >            ; 
                
OctalEscape     : int03Two07
                / two07
                / int07                             ;
    int03Two07  : any "0123" any "01234567" any "01234567" ;
      two07     : any "01234567" any "01234567"     ;
      int07     : any "01234567"                    ;
      any4567   : any "4567"                        ;
      int03     : any "0123"                        ;
```
#### ***Exponent and Digital Literals***
Original Egg grammar had a NT "Digit", which is replaced here in GoGll by the reserved word "number" and occasionally the NT "num" which has been assigned to number.
```
exponent        : any "eE" [ any "+\\-" ] { number }    ;
binaryExponent  : any "pP" any "+\\-" < number >        ; 

num             : number                                ;     

```   
#### ***Separators and Operators***
```
AT                  :  "@"            WS                 ;
AND                 :  "&" notEqAnd   WS                 ;
AND_AND             :  "&&"           WS                 ;
AND_EQU             :  "&="           WS                 ;
BANG                :  "!" nEq        WS                 ;
BSR                 :  ">>>" nEq      WS                 ;
BSR_EQU             :  ">>>="         WS                 ;
COLON               :  ":"            WS                 ;
COMMA               :  ","            WS                 ;
DEC                 :  "--"           WS                 ;
DIV                 :  "/" nEq        WS                 ;
DIV_EQU             :  "/="           WS                 ;
DOT                 :  "."            WS                 ;
EQU                 :  "=" nEq        WS                 ;
EQUAL               :  "=="           WS                 ;
GE                  :  ">="           WS                 ;
GT                  :  ">" notEqCar   WS                 ;
HAT                 :  "^" nEq        WS                 ;
HAT_EQU             :  "^="           WS                 ;

notEqAnd            :  not "=" not "&"                   ;

INC                 :  "++"           WS                 ;
LBRK                :  "["            WS                 ;
LE                  :  "<="           WS                 ;
LPAR                :  "("            WS                 ;
LPOINT              :  "<"            WS                 ;
LT                  :  "<" notEqCar2  WS                 ;
LWING               :  "{"            WS                 ;
MINUS               : "-" notEqSlDash WS                 ;
MINUS_EQU           :  "-="           WS                 ;
MOD                 :  "%" nEq        WS                 ;
MOD_EQU             :  "%="           WS                 ;
NOT_EQUAL           :  "!="           WS                 ;   
OR                  :  "|" notEqPipe  WS                 ;
OR_EQU              :  "|="           WS                 ;
OR_OR               :  "||"           WS                 ;
PLUS                :  "+" notEqPlus  WS                 ;
PLUS_EQU            :  "+="           WS                 ;

notEqPipe           :  not "=" not "|"                   ; 
notEqPlus           :  not "=" not "+"                   ;
notEqCar2           :  not "=" not "<"                   ;
notEqSlDash         :  not "=" not "\\" not "-"          ;

QUERY               :  "?"            WS                 ;
RBRK                :  "]"            WS                 ;
RPAR                :  ")"            WS                 ;
RPOINT              :  ">"            WS                 ;
RWING               :  "}"            WS                 ;
SEMI                :  ";"            WS                 ;
SL                  :  "<<" nEq       WS                 ;
SL_EQU              :  "<<="          WS                 ;
SR                  :  ">>" notEqCar  WS                 ;
SR_EQU              :  ">>="          WS                 ;
STAR                :  "*" nEq        WS                 ;
STAR_EQU            :  "*="           WS                 ;
TILDA               :  "~"            WS                 ;
    

notEqCar            :  not "=" not ">"                   ; 
nEq                 : not "="                            ;


```
### ***Escape Characters/Sequences, Comments, and Spacing***
- Note: To match the -> operator in GoGLL, the following syntax is used:
    
    (Egg): XtoY : X -> Y
    
    (GoGLL): XtoY : Y / X XtoY;
    
```
WS                : EscOrLineOrBlock     
                  | empty                             ;
EscOrLineOrBlock  : line_comment 
                  | block_comment                     
                  | escCharSp                         ;
      escCharSp   : < any " \t\r\n" >                 ;
      
!line_comment : '/' '/' {not "\n"} ;
!block_comment : '/''*' {not "*" | '*' not "/"} '*''/' ;
      newline       : any "\r\n"                       ;

```
### **COPYRIGHT AND LICENSING INFORMATION**
**Copyright 2021 Brynn Harrington and Emily Hoppe**

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License.

You may obtain a copy of the License [here](http://www.apache.org/licenses/LICENSE-2.0) or at:

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.