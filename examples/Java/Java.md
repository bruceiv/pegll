// // Java grammar
Modification of `Java` grammar from [Egg](https://github.com/bruceiv/egg/blob/deriv/grammars/Java-u.egg) to test Java grammar.

```
package "Java"

//-------------------------------------------------------------
//  @ Author : Roman R Redziejowski Copyright (C) 2006
//  (http://home.swipnet.se/redz/roman).
//
//  This is free software; you can redistribute and/or modify
//  it under the terms of the GNU Library General Public License
//  as published by the Free Software Foundation; either version 2
//  of the License or (at your option) any later version.
//
//  This file is distributed in the hope that it will be useful,
//  but WITHOUT ANY WARRANTY; without even the implied warranty of
//  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
//
//  For more details, see the GNU Library General Public License
//  at http://www.fsf.org/copyleft/gpl.html.
//
//  @ Modified to GoGLL Grammar : Brynn Harrington and Emily Hoppe
//  Date : June 10, 2021 
//  Last Updated: June 12, 2021

//-------------------------------------------------------------
//  COMPILATION Unit
//-------------------------------------------------------------

CompUnit          : _ OptPackDecl RepImpDecl0x RepSemiModDecl0x ;
      OptPackDecl : [ PackDecl ]                      ;
     RepImpDecl0x : { ImportDecl }                    ; 
 RepSemiModDecl0x : { SEMI 
                  | RepModif0 DeclAlts }              ;
      DeclAlts    : ClsDecl  
                  | IntfDecl                          ;

PackDecl          : PACKAGE QualifiedID SEMI          ;

ImportDecl        : IMPORT OptStatic QualifiedID OptDotStar SEMI ;
      OptDotStar  : [ DOT STAR ]                      ;

//-------------------------------------------------------------
//  CLASS DECLARATIONS
//    - Note: The following are the representations of the 
//    MemDecl (Member Decl): 
//          Type ID FormalParams RepDim0x OptThrowClsTypLst MemAlts = Method
//          VOID ID FormalParams OptThrowClsTypLst MemAlts = Void Method
//          ID FormalParams OptThrowClsTypLst Block = Constructor
//          IntfDecl = Interface
//          ClsDecl = Class
//          Type VarDecl RepComVDecl0x = Field
//    - Note: The following are the representations of the 
//     ClsBdyDecl(ClassBodyDeclaration): 
//          SEMI = Semicolon
//          OptStatic Block  = Static or Instance Initializer
//          RepModif0 MemDecl = Class Member Declaration  
//-------------------------------------------------------------

ClsDecl           : Cls ID OptExtClsType OptImpClsLst ClsBdy ;
    OptExtClsType : [ EXTENDS ClsType ]               ;
     OptImpClsLst : [ IMPLEMENTS ClsTypeList ]        ;
      
ClsBdy            : LWING RepClsBDecl0x RWING         ;
    RepClsBDecl0x : { ClsBdyDecl }                    ;            

ClsBdyDecl        : SEMI
                  | OptStatic Block 
                  | RepModif0 MemDecl                 ;
      OptStatic : [ STATIC ]                          ; 

MemDecl           : Type ID FormalParams RepDim0x OptThrowClsTypLst MemAlts
                  | VOID ID FormalParams OptThrowClsTypLst MemAlts
                  | ID FormalParams OptThrowClsTypLst Block
                  | IntfDecl 
                  | ClsDecl
                  | Type VarDecl RepComVDecl0x        ; 
      MemAlts     : SEMI 
                  | Block                             ;

//-------------------------------------------------------------
//  INTERFACE DECLARATION
//-------------------------------------------------------------
IntfDecl   : Intf ID OptExtendsClsLis IntfBdy         ;
      OptExtendsClsLis : [ EXTENDS ClsTypeList ]      ;

IntfBdy           : LWING RepInBodDecl0x RWING        ;
   RepInBodDecl0x : { IntfBdyDecl }                   ;

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
 OptThrowClsTypLst: [ THROWS ClsTypeList ]            ;

ConstDeclsRest    : ConstDeclRest RepComCnstDecl0x    ;
 RepComCnstDecl0x : { COMMA ConstDecl }               ;

ConstDecl         : ID ConstDeclRest                  ;

ConstDeclRest     : RepDim0x EQU VarInitial           ;
    
//-------------------------------------------------------------
//  VARIABLE DECLARATIONS
//-------------------------------------------------------------

LocalVarDeclStmt  : OptFinType VarDecl RepComVDecl0x SEMI ;

VarDecl           :  ID RepDim0x OptEqVarInit         ;
     OptEqVarInit : [ EQU  VarInitial ]               ;

//-------------------------------------------------------------
//  FORMAL PARAMETERS
//-------------------------------------------------------------

FormalParams      : LPAR OptFormPDecl RPAR            ;
     OptFormPDecl : [ FormalParamDecls ]              ;

FormalParam       : OptFinType VarDelID               ;
 
FormalParamDecls  : OptFinType FormalParamDeclsRest   ;

FormalParamDeclsRest :  VarDelID OptComFormPDecl      ;
  OptComFormPDecl : [ COMMA FormalParamDecls ]        ;

VarDelID          : ID RepDim0x                       ;

//-------------------------------------------------------------
//  STATEMENTS
//-------------------------------------------------------------

Block             : LWING RepBlkSt0x RWING            ;
     RepBlkSt0x   : { BlockStmt }                     ;

BlockStmt         : LocalVarDeclStmt
                  | RepModif0 ClsDecl
                  | Stmt                              ;
   RepModif0      : { Modifier }                      ;

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
      OptColExpr  : [ COLON Expr ]                    ;
      OptElse     : [ ELSE Stmt ]                     ;
      OptForInit  : [ ForInit ]                       ;
      OptForUpd   : [ ForUpdate ]                     ;
      OptExpr     : [ Expr ]                          ;
      CatchBlk    : RepCatch1x OptFinly 
                  | Finally                           ;
      RepCatch1x  : < Catch >                         ;
      OptFinly    : [ Finally ]                       ;
   RepSwBlkStmt0x : { SwitchBlockStmtGrp }            ;
      OptID       : [ ID ]                            ;       

Catch             : CATCH LPAR FormalParam RPAR Block ;
   
Finally           : FINALLY Block                     ;

SwitchBlockStmtGrp: SwitchLabel RepBlkSt0x            ;
      RepBlkSt0x  : { BlockStmt }                     ;


SwitchLabel       : CASE ConstExpr COLON 
                  | DEFAULT COLON                     ;

ForInit           : OptFinType Type VarDeclInit       ;
                  | StmtExpr RepComSExpr0x            ; 
   OptFinType     : [ FINAL ] Type                    ;
   VarDeclInit    : VarDecl RepComVDecl0x             ; 
   RepComVDecl0x  : { COMMA VarDecl }                 ;

ForUpdate         : StmtExpr RepComSExpr0x            ;
    RepComSExpr0x : { COMMA StmtExpr }                ;

//-------------------------------------------------------------
//  EXPRESSIONS
//    - Note: Some of the shorthand names are:
//          Cond = Conditional
//          IOR = Inclusive-Or 
//          XOR = Exclusive-Or
//    - Note: StmtExpr is more generous than definition in 14.8,
//    which allows only specific forms of Expr.
//    - Note: The definition of Expr is part of the modification
//    in JLS Ch. 18 to minimize look ahead. In JLS Ch. 15.27, Expr
//    is defined as AssignmentExpr, which is effectively defined as
//    (LeftHandSide AssignOp)* CondExpr. The above is obtained by 
//    allowing ANY CondExpr as LeftHandSide, which results in 
//    accepting Stmts like 5 : a.
//    - Note: ArrayCreator is more generous than JLS 15.10.
//    According to that definition, BasicType must be followed 
//    by at least one DimExpr or by ArrayInitializer.
//-------------------------------------------------------------

StmtExpr          : Expr                              ;
   
ConstExpr         : Expr                              ;

Expr              : CondExpr RepAsscExpr0x            ;
    RepAsscExpr0x : { AssignOp CondExpr }             ;

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
   RepCondition0x : { QUERY Expr COLON CondORExpr }   ;

CondORExpr        : CondANDExpr RepORcAND0x           ;
      RepORcAND0x : { OR_OR CondANDExpr }             ;

CondANDExpr       : IORExpr RepANDIOR0x               ;
      RepANDIOR0x : { AND_AND IORExpr }               ;

IORExpr           : XORExpr RepORXOR0x                ; 
      RepORXOR0x  : { ORXOR }                         ;
      ORXOR       : OR XORExpr                        ;

XORExpr           : ANDExpr RepHatAND0x               ;
      RepHatAND0x : { HAT ANDExpr }                   ;

ANDExpr           : EqualExpr RepANDEq0x              ;
      RepANDEq0x  : { AND EqualExpr }                 ;

EqualExpr         : RelateExpr RepEqExpr0x            ;
      RepEqExpr0x : { EqAlts RelateExpr }             ;
      EqAlts      : EQUAL 
                  | NOT_EQUAL                         ;

RelateExpr        : ShiftExpr RepESInst0x             ;
      RepESInst0x : { ESInst }                        ;
      ESInst      : EqShift 
                  | INSTANCEOF ReferenceType          ;
      EqShift     : EqCheck ShiftExpr                 ;
      EqCheck     : LE | GE | LT | GT                 ;

ShiftExpr   : AddExpr ShiftAlts                       ;
      CarrotAlts  : SL 
                  | SR 
                  | BSR                               ;
      ShiftAlts   : { CarrotAlts AddExpr }            ;


AddExpr           : MultExpr RepAddAltsMult0x         ;
 RepAddAltsMult0x : { AddAlts MultExpr }              ;
      AddAlts     : PLUS | MINUS                      ;

MultExpr          : UnaryExpr RepSDMUExpr0x           ;
    RepSDMUExpr0x : { SDM UnaryExpr }                 ;
      SDM         : STAR | DIV | MOD                  ;

UnaryExpr         : PrefixOp UnaryExpr
                  | LPAR Type RPAR UnaryExpr
                  | Primary RepSel0x RepPfOp0x        ;
   RepSel0x       : { Selector }                      ;
   RepPfOp0x      : { PostfixOp }                     ;

Primary           : ParExpr
                  | THIS OptArgs
                  | SUPER SuperSuffix
                  | Literal
                  | NEW Creator
                  | QualifiedID OptIDSuff
                  | BasicType RepDim0x DOT Cls
                  | VOID DOT Cls                    ;            
      OptIDSuff   : [ IDSuffix ]                      ;

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
      OptArgs     : [ Arguments ]                     ;

BasicType         : BasicTypeLit NotLorD              ;
     BasicTypeLit : "byte"
                  | "short"
                  | "char"
                  | "int"
                  | "long"
                  | "float"
                  | "double"
                  | "boolean"                         ;

Arguments         : LPAR OptExprs RPAR                ;
      OptExprs    : [ Expr RepComExp0x ]              ;
      RepComExp0x : { COMMA Expr }                    ;

Creator           : CreatedName ClsCreatorRest 
                  | TypeAlts ArrayCreatorRest         ;

CreatedName       : ID RepDotID0x                     ;

InnerCreator      : ID ClsCreatorRest               ;

ArrayCreatorRest  : LBRK ArrayRest                    ;
      ArrayRest   : RBRK RepDim0x ArrayInitializer 
                  | Expr RBRK RepDimExpr0x RepDim0x   ;
      RepDimExpr0x: { DimExpr }                       ;
   
ClsCreatorRest  :  Arguments OptClsBdy           ;
      OptClsBdy: [ ClsBdy ]                     ;

ArrayInitializer  : LWING OptVarInit RWING            ;
      OptVarInit  : [ VarInitial RepComInit0x OptCom ];
     RepComInit0x : { COMMA VarInitial }              ;
      OptCom      : [ COMMA ]                         ;

VarInitial        : ArrayInitializer                
                  | Expr                              ;

ParExpr           : LPAR Expr RPAR                    ;

QualifiedID       : ID RepDotID0x                     ;

Dim               : LBRK RBRK                         ;
DimExpr           : LBRK Expr RBRK                    ; 
////////////////////////////////////////Dim is used MANY times before this??? NEED TO MOVE FURTHER DOWN!!
///// but also uses expression in the definition so like????

//-------------------------------------------------------------
//  TYPES AND MODIFIERS
//    - Note: This common definition of Modifier is part of the 
//    modification in JLS Chapter 18 to minimize look ahead. 
//    The main Bdy of JLS has different lists of modifiers 
//    for different language elements.
//-------------------------------------------------------------

Type              : TypeAlts RepDim0x                 ; 
      TypeAlts    : BasicType 
                  | ClsType                         ;

ReferenceType     : BasicType RepDim1x 
                  | ClsType RepDim0x                ;
      RepDim0x    : { Dim }                           ;
      RepDim1x    : < Dim >                           ;

ClsType         : ID RepDotID0x                     ;
      RepDotID0x  : { DOT Indentifier}                ;

ClsTypeList     : ClsType RepComCls0x           ;
   RepComCls0x  : { COMMA ClsType }               ;

Modifier          : Modifs NotLorD                    ;
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
   
//-------------------------------------------------------------
//  IDENFITIFERS
//    - Note: ID is used to represent identifiers.
//-------------------------------------------------------------

ID                : not Keyword LetterLorD            ;   
      LetterLorD   : Letter RepLorD0x _               ;
      RepLorD0x   :  { LorD }                         ; 

//-------------------------------------------------------------
//  KEYWORDS
//-------------------------------------------------------------
Keyword           : Words NotLorD                     ;
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

ASSERT            : "assert"       NotLorD            ; 
BREAK             : "break"        NotLorD            ;
CASE              : "case"         NotLorD            ;
CATCH             : "catch"        NotLorD            ;
Cls             : "Cls"        NotLorD            ;
CONTINUE          : "continue"     NotLorD            ;
DEFAULT           : "default"      NotLorD            ;
DO                : "do"           NotLorD            ;
ELSE              : "else"         NotLorD            ;
ENUM              : "enum"         NotLorD            ;
EXTENDS           : "extends"      NotLorD            ;
FINALLY           : "finally"      NotLorD            ; 
FINAL             : "final"        NotLorD            ;
FOR               : "for"          NotLorD            ;
IF                : "if"           NotLorD            ; 
IMPLEMENTS        : "implements"   NotLorD            ; 
IMPORT            : "import"       NotLorD            ;
Intf              : "Intf"         NotLorD            ;
INSTANCEOF        : "instanceof"   NotLorD            ;
NEW               : "new"          NotLorD            ;
PACKAGE           : "package"      NotLorD            ; 
RETURN            : "return"       NotLorD            ;
STATIC            : "static"       NotLorD            ; 
SUPER             : "super"        NotLorD            ;
SWITCH            : "switch"       NotLorD            ;
SYNCHRONIZED      : "synchronized" NotLorD            ; 
THIS              : "this"         NotLorD            ;
THROWS            : "throws"       NotLorD            ;
THROW             : "throw"        NotLorD            ;
TRY               : "try"          NotLorD            ;
VOID              : "void"         NotLorD            ;
WHILE             : "while"        NotLorD            ; 

//-------------------------------------------------------------
//  GENERAL LITERAL DEFINITION 
//-------------------------------------------------------------

Literal           : LitAlts _                         ;
      LitAlts     : FloatLiteral
                  | IntegerLiteral 
                  | CharLiteral
                  | StringLiteral
                  | "true"  NotLorD
                  | "false" NotLorD
                  | "null"  NotLorD                   ;

//-------------------------------------------------------------
//  BASIC IDENTIFIER DEFINITIONS
//    - Note: These are traditional definitions of letters and
//    digits. JLS defines letters and digits as Unicode  
//    characters recognized as such by special Java procedures,
//    which is difficult to express in terms of Parsing 
//    Expressions.
//-------------------------------------------------------------

NotLorD           : not LorD                          ;        
LorD              : Letter 
                  | Digit 
                  | _                                 ; ///////////////////// do we need this since the '_' is already inside the letter definition/rule??
Letter            : letter 
                  | _                                 ;

//-------------------------------------------------------------
//  CHARACTER AND STRING LITERALS
//    - Note: Unicode escape is not defined in JLS syntax because 
//    unicode characters are processed very early.
//-------------------------------------------------------------

CharLiteral       : '\'' EscSlash '\''                ;
      EscSlash    : ( Escape 
                  | EscUp )                           ;
      EscUp       : '^' 
                  | '\'' 
                  | '\\'                              ;

StringLiteral     : '\"' StrClose                     ;
      StrClose    : '\"' 
                  / OptEsc StrClose                   ;
      OptEsc      : ( Escape 
                  | [^\\] )                           ;

Escape            : "\\" Escs                         ;
      Escs        : ( EsChars 
                  | OctalEscape 
                  | UnicodeEscape )                   ;
      EscChars    : '\\' 
                  | '\"' 
                  | '\'' 
                  | 'b' 
                  | 'f' 
                  | 'n' 
                  | 'r' 
                  | 't'                               ;
   
UnicodeEscape : "u" HexDigit HexDigit HexDigit HexDigit ;
 
//-------------------------------------------------------------
//  GENERAL NUMBERIC LITERALS
//    - Note: In IntegerLiteral, OctalNumeral may prefix 
//    HexNumeral and DecimalNumeral may prefix OctalNumeral
//-------------------------------------------------------------

FloatLiteral      : HexFloat 
                  | DecimalFloat                      ;

IntegerLiteral    : NumeralAlts OptOneL               ;
      NumeralAlts : HexNumeral 
                  | OctalNumeral  
                  | DecimalNumeral                    ;
      OptOneL     : [ any "1L" ]                      ;

DecimalFloat      :  RepDig1x "." RepDig0x OptExpo fF_dD 
                  | "." RepDig1xExp 
                  | RepDig1xExp fF_dD
                  | RepDig1xOptExp fF_dD              ;
   RepDig1xOptExp : RepDig1x OptExpo                  ;
      RepDig1xExp : RepDig1x Exponent                 ;
      OptExpo     : [ Exponent ]                      ;
            
//-------------------------------------------------------------
//  BASE-SIXTEEN AND BASE-EIGHT LITERALS
//-------------------------------------------------------------

HexFloat          : HexSignificand BinaryExponent fF_dD ;
      fF_dD       : [ any "fFdD" ]                    ; 

HexSignificand    : HexNumeral OptDot 
                  | RepHex0xDot RepHex1x              ;
      RepHex0xDot : zeroxX RepHex0x "."               ;
      OptDot      :    [ '.' ]                        ; 

HexNumeral        : zeroxX RepHex1x                   ; 
      zeroxX      : '0' xX                            ;
      xX          : any "xX"                          ;          
      RepHex0x    : { HexDigit }                      ;  
      RepHex1x    : < HexDigit >                      ;  

HexDigit          : < Digit aA-fF >                   ;
      aA-fF       : any "abcdefABCDEF"                ;  
 
OctalNumeral      : "0" Rep0-7_1x                     ; 
      Rep0-7_1x   : < Int0-7 >                        ;

OctalEscape       : Int0-3 Two0-7
                  / Two0-7
                  / Int0-7                            ;
      Two0-7      : Int0-7 Int0-7                     ;
      Int0-7      : any Int0-3
                  | any "4567"                        ;
      Int0-3      : any "0123"                        ;

//-------------------------------------------------------------
//  EXPONENT AND DIGIT LITERALS
//-------------------------------------------------------------

Exponent          : eE OptPSM RepDig0x                ;
      eE          : any "eE"                          ;
      RepDig0x    : { Digit }                         ; 

BinaryExponent    : pP PSM RepDig1x                   ;
      pP          : 'p' | 'P'                         ;     
      PSM         : any "+\-"                         ;
      RepDig1x :  < Digit >                           ;        

Digit             : number                            ;   

//-------------------------------------------------------------
//  SEPERATORS AND OPERATORS
//-------------------------------------------------------------

AT                :  '@'            _                 ;
AND               :  '&'![=&]       _                 ;
AND_AND           :  "&&"           _                 ;
AND_EQU           :  "&="           _                 ;
BANG              :  '!' !'='       _                 ;
BSR               :  ">>>"!'='      _                 ;
BSR_EQU           :  ">>>="         _                 ;
COLON             :  ':'            _                 ;
COMMA             :  ','            _                 ;
DEC               :  "--"           _                 ;
DIV               :  '/' !'='       _                 ;
DIV_EQU           :  "/="           _                 ;
DOT               :  '.'            _                 ;
EQU               :  '=' !'='       _                 ;
EQUAL             :  "=="           _                 ;
GE                :  ">="           _                 ;
GT                :  '>'![=>]       _                 ;
HAT               :  '^'!'='        _                 ;
HAT_EQU           :  "^="           _                 ;
INC               :  "++"           _                 ;
LBRK              :  '['            _                 ;
LE                :  "<="           _                 ;
LPAR              :  '('            _                 ;
LPOINT            :  '<'            _                 ;
LT                :  '<'![=<]       _                 ;
LWING             :  '{'            _                 ;
MINUS             :  '-'![=\-]      _                 ;
MINUS_EQU         :  "-="           _                 ;
MOD               :  '%'!'='        _                 ;
MOD_EQU           :  "%="           _                 ;
NOT_EQUAL         :  "!="           _                 ;   
OR                :  '|'![=|]       _                 ;
OR_EQU            :  "|="           _                 ;
OR_OR             :  "||"           _                 ;
PLUS              :  '+'![=+]       _                 ;
PLUS_EQU          :  "+="           _                 ;
QUERY             :  '?'            _                 ;
RBRK              :  ']'            _                 ;
RPAR              :  ')'            _                 ;
RPOINT            :  '>'            _                 ;
RWING             :  '}'            _                 ;
SEMI              :  ';'            _                 ;
SL                :  "<<"!'='       _                 ;
SL_EQU            :  "<<="          _                 ;
SR                :  ">>"![=>]      _                 ;
SR_EQU            :  ">>="          _                 ;
STAR              :  '*'!'='        _                 ;
STAR_EQU          :  "*="           _                 ;
TILDA             :  '~'            _                 ;

//-------------------------------------------------------------
//  ESCAPES, COMMENTING, AND SPACING
//    Note: To match the -> operator in GoGLL, the following
//    syntax is used:
//    (Egg): XtoY : X -> Y
//    (GoGLL): XtoY : Y / X XtoY;
//-------------------------------------------------------------
_                 : { EscCharSp     
                  | BlockComment 
                  | Comment }                         ;
      EscCharSp   : < ' ' 
                  | EscChar >                         ;
      EscChar     : '\t' 
                  | NewLine                           ;      
     BlockComment : "*/" 
                  / "/*" BlockComment                 ;
      Comment     : NewLine 
                  / "//" Comment                      ; 
      NewLine     : '\r' 
                  | '\n'                              ;


```