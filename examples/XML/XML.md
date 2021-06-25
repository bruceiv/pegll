## **`XML` GRAMMAR**
### **AUTHORSHIP INFORMATION**
#### *Authors :* Brynn Harrington and Emily Hoppe Copyright (C) 2021
#### *Adapted from :* Aaron Moss's [`XML` Egg Grammar](https://github.com/bruceiv/egg/blob/deriv/grammars/XML-u.egg)
#### *Creation Date :* June 11, 2021 
#### *Last Modified :* June 24, 2021
#### *Copyright and Licensing Information :* See end of file.

### **GENERAL DESCRIPTION**
A modification of `XML` [Egg](https://github.com/bruceiv/egg/blob/deriv/grammars/XML-u.egg) parsing grammar ported into GoGLL to test `XML` input files.

### **STATUS ON GRAMMAR**
#### *Markdown File Creation:* Complete
#### *Parser Generated :* Complete
#### *Test File Creation:* Incomplete
#### *Testing Results:* Unknown

### **`XML` GRAMMAR GUIDE**
The following grammar tests input files from the `XML` language.

```
package "XML"
```
#### ***Higher-Level Language Structures***
```
Document                : Prolog Element RepMisc0x              ;

Prolog 	                : OptXMLDecl RepMisc0x                  ;
XMLDecl                 : xmlDeclStart VersionInfo OptEncodDecl 
                        optSpaceEsc xmlDeclEnd                  ;
        OptXMLDecl      : XMLDecl 
                        / empty                                 ;
        xmlDeclStart    : '<' '?' 'x' 'm' 'l'                   ;
        xmlDeclEnd      :  '?' '>'                              ;

        VersionInfo     : spaceEsc version Eq QuoVerNum         ;
        QuoVerNum       : sinQu VersionNum sinQu  
                        | dubQu VersionNum dubQu                ;

VersionNum              : NAME_CHAR NameCharRep                 ;
        NameCharRep     : NAME_CHAR NameCharRep
                        / empty                                 ;

EncodingDecl            : spaceEsc encoding Eq QuoEncNam        ;
        QuoEncNam       : sinQu EncName sinQu  
                        | dubQu EncName dubQu                   ;
        OptEncodDecl    : EncodingDecl 
                        / empty                                 ;
        encoding        : 'e' 'n' 'c' 'o' 'd' 'i' 'n' 'g'       ;
        version         : 'v' 'e' 'r' 's' 'i' 'o' 'n'           ;
```
#### ***Values and References***
```
ATT_VALUE               : dubQu DubCondClose 
                        | sinQu SinCondClose                    ;
        SinCondClose    : sinQu
                        / SymRefAlts SinCondClose               ;
        DubCondClose    : dubQu 
                        / SymRefAlts DubCondClose               ;
        SymRefAlts      : andCarrs
                        | REFERENCE                             ;
        andCarrs        : any "^<&"                             ;
        dubQu           : '"'                                   ;
        sinQu           : '\''                                  ;

REFERENCE               : ENTITY_REF 
                        | CHAR_REF                              ;

ENTITY_REF              : "&" NAME ";"                          ;        

CHAR_REF                : "&#x" Hex ";"  
                        | "&#" repNum1x ";"                     ;
        Hex             : HexAlts RepHexAlts0x                  ;
        RepHexAlts0x    : HexAlts Hex   
                        / empty                                 ;       
        HexAlts         : num
                        | aA_fF                                 ;  
        repNum1x        : < number >                            ;
        aA_fF           : any "abcdefABCDEF"                    ;
```
#### ***Commenting, Elements, and Attributes***
```
Content                 : ContentAlts Content
                        / empty                                 ;
        ContentAlts     : COMMENT 
                        | Element 
                        | REFERENCE 
                        | charData                              ;
Misc                    : COMMENT 
                        | spaceEsc                              ; 
        RepMisc0x       : Misc RepMisc0x 
                        / empty                                 ;

COMMENT                 : ComStart ComEnterior angRBrk          ;
        ComStart        : angLBrk exclamation DubDash           ;
        DubDash         : "--"                                  ;
        ComEnterior     : DubDash 
                        / let ComEnterior                       ;


Element                 : angLBrk NAME RepSAttx0x optSpaceEsc 
                        ElemCloseAlts                           ;
        SAtt            : spaceEsc Attribute                    ;      
        RepSAttx0x      : SAtt RepSAttx0x  
                        / empty                                 ;
        ElemCloseAlts   : angRBrk Content slashAngLBrk NAME optSpaceEsc angRBrk 
                        | slashAngRBrk                          ;
        angLBrk         : '<'                                   ;
        slashAngLBrk    : '<' '/'                               ;
        angRBrk         : '>'                                   ;
        slashAngRBrk    : '/' '>'                               ;
        exclamation     : '!'                                   ;

Attribute               : NAME optSpaceEsc eq optSpaceEsc 
                        ATT_VALUE                               ;
        Eq              : optSpaceEsc eq optSpaceEsc            ;
        optSpaceEsc     : [ < any " \t\r\n" > ]                 ;
        spaceEsc        : < any " \t\r\n" >                     ;
        charData        :  < any "^<&" >                        ;
        eq              : '='                                   ;
```
#### ***Names, Encoding, and (Whitespace/Escape) Characters***
```
NAME                    : LetColonAlts RepNameChar0x            ;
        LetColonAlts    : let 
                        | col_                                  ;
        RepNameChar0x   :  NAME_CHAR RepNameChar0x 
                        / empty                                 ;
        col_            : any "_:"                              ; 

NAME_CHAR               : let 
                        | num
                        | dot_BSlashDashCol                     ;
      dot_BSlashDashCol : any "\\-._:"                          ;

EncName                 : let RepLDSAlts0x                      ;
        RepLDSAlts0x    : LetDigSymAlts RepLDSAlts0x
                        / empty                                 ;
        LetDigSymAlts   : let   
                        | num
                        | dot_BSlashDash                        ;       
        dot_BSlashDash  : any "._\\-"                           ;
        let             : letter                                ;
        num             : number                                ;

```

### **COPYRIGHT AND LICENSING INFORMATION**
**Copyright 2021 Brynn Harrington and Emily Hoppe**

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License.

You may obtain a copy of the License [here](http://www.apache.org/licenses/LICENSE-2.0) or at:

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.