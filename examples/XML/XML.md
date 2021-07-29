## **`XML` GRAMMAR**
### **AUTHORSHIP INFORMATION**
#### *Authors :* Brynn Harrington and Emily Hoppe Copyright (C) 2021
#### *Adapted from :* Aaron Moss's [`XML` Egg Grammar](https://github.com/bruceiv/egg/blob/deriv/grammars/XML-u.egg)
#### *Creation Date :* June 11, 2021 
#### *Last Modified :* July 29, 2021
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
Document                : Prolog Element Misc*                  ;

Prolog 	                : XMLDecl? Misc*                        ;
XMLDecl                 : "<?xml" VersionInfo EncodingDecl? 
                         optSpaceEsc "?>"                       ;

VersionInfo             : sp "version" Eq QuoVerNum       ;
        QuoVerNum       : "'" VersionNum "'"  
                        | dubQu VersionNum dubQu                ;

VersionNum              : NAME_CHAR+                            ;

EncodingDecl            : sp "encoding" Eq QuoEncNam      ;
        QuoEncNam       : "'" EncName "'"  
                        / dubQu EncName dubQu                   ;
        
```
#### ***Values and References***
sinQu isn't necessary but there is an error in the lexer if it is not there
```

ATT_VALUE               : dubQu DubConClose 
                        | "'" SinConClose                       ;
        SinConClose     : "'"
                        / SymRefAlts SinConClose                ;
        DubConClose     : dubQu 
                        / SymRefAlts DubConClose                ;
        SymRefAlts      : andCars
                        | REFERENCE                             ;
        andCars         : any "^<&"                             ;
        dubQu           : '"'                                   ;
        sinQu           : '\''                                  ;

REFERENCE               : ENTITY_REF 
                        | CHAR_REF                              ;

ENTITY_REF              : "&" NAME ";"                          ;        

CHAR_REF                : "&#x" Hex ";"  
                        | "&#" num+ ";"                     ;
        Hex             : HexAlts+                              ;     
        HexAlts         : num
                        | aA_fF                                 ;  
        aA_fF           : any "abcdefABCDEF"                    ;
```
#### ***Commenting, Elements, and Attributes***
```
Content                 : ContentAlts*                          ;
        ContentAlts     : comment 
                        | Element 
                        | REFERENCE 
                        | charData                              ;

Misc                    : comment 
                        | sp                                    ; 

!comment                : '<''!''-''-'
                        { not "-->"
                        | '-' not "->"
                        | '-''-' not ">"
                        } '-''-''>'                             ;  


Element                 : angLBrk NAME SAtt* sp? ElemCloseAlts  ;
        SAtt            : sp Attribute                          ;      
        ElemCloseAlts   : ">" Content "</" NAME sp? ">" 
                        | "</"                                  ;
        angLBrk         : '<'                                   ;

Attribute               : NAME Eq ATT_VALUE                     ;
        Eq              : sp? "=" sp?                           ;
        sp              : < any " \t\r\n" >                     ;

charData                : < any "^<&" >                         ;
```
#### ***Names, Encoding, and (Whitespace/Escape) Characters***
```
NAME                    : LetColonAlts NAME_CHAR*               ;
        LetColonAlts    : let 
                        | ":"
                        | "_"                                   ;


NAME_CHAR               : let 
                        | num
                        | ":"
                        | "_"
                        | dot_BSlashDash                        ;

EncName                 : let LetDigSymAlts*                    ;
        LetDigSymAlts   : let   
                        | num
                        | "_"
                        | dot_BSlashDash                        ;       
        dot_BSlashDash  : any ".\\-"                            ;
        let             : letter                                ;
        num             : number                                ;

```

### **COPYRIGHT AND LICENSING INFORMATION**
**Copyright 2021 Brynn Harrington and Emily Hoppe**

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License.

You may obtain a copy of the License [here](http://www.apache.org/licenses/LICENSE-2.0) or at:

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.