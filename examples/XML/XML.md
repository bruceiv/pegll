## **`XML` Grammar**
### **AUTHORSHIP INFORMATION**
#### *Authors :* Brynn Harrington and Emily Hoppe Copyright (C) 2021
#### *Adapted from :* Aaron Moss's [`XML` Egg Grammar](https://github.com/bruceiv/egg/blob/deriv/grammars/XML-u.egg)
#### *Creation Date :* June 11, 2021 
#### *Last Modified :* June 23, 2021
#### *Copyright and Licensing Information :* See end of file.

### **GENERAL DESCRIPTION**
An originally Egg Parsing grammar created by Aaron Moss ported into the GoGLL grammar for the `XML` language tests. Modification of `XML` grammar from [Egg](https://github.com/bruceiv/egg/blob/deriv/grammars/XML-u.egg) to test `XML` input files under the parser generated.
### **`XML` GRAMMAR GUIDE**
GRAMMAR IS WORKING 


```
package "XML"

        Document       : Prolog Element RepMisc0                ;

        Prolog 	       : OptXMLDecl RepMisc0                    ;
        XMLDecl        : xmlDeclStart VersionInfo OptEncodDecl optSpaceEsc xmlDeclEnd ;
                OptXMLDecl : XMLDecl 
                           / empty                              ;
                xmlDeclStart : '<' '?' 'x' 'm' 'l'              ;
                xmlDeclEnd   :  '?' '>'                         ;

        VersionInfo    : spaceEsc version Eq QuoVerNum          ;
                QuoVerNum    : sinQu VersionNum sinQu  
                             | dubQu VersionNum dubQu           ;
        VersionNum     : NAME_CHAR NameCharRep                  ;
                NameCharRep  : NAME_CHAR NameCharRep
                             / empty                            ;
        EncodingDecl   : spaceEsc encoding Eq QuoEncNam         ;
                QuoEncNam    : sinQu EncName sinQu  
                             | dubQu EncName dubQu              ;
                OptEncodDecl : EncodingDecl / empty             ;
                encoding     : 'e' 'n' 'c' 'o' 'd' 'i' 'n' 'g'  ;
                version      : 'v' 'e' 'r' 's' 'i' 'o' 'n'      ;
```
#### ***Values and References***
```
        ATT_VALUE       : dubQu DubCondClose 
                        | sinQu SinCondClose           ;
                SinCondClose : sinQu
                        / SymRefAlts SinCondClose      ;
                DubCondClose  : dubQu 
                        / SymRefAlts DubCondClose      ;
                SymRefAlts    : andCarrs
                              | REFERENCE              ;
                andCarrs : any "^<&"    ;
                dubQu    : '"'          ;
                sinQu    : '\''         ;

        REFERENCE       : ENTITY_REF 
                        | CHAR_REF                     ;
        ENTITY_REF      : "&" NAME ";"                 ;
        CHAR_REF        : "&#x" Hex ";"  
                        | "&#" repNum1 ";"             ;
                semi  : ';'         ;  
                Hex             : HexAlts RepHexAlts    ;
                RepHexAlts      : HexAlts Hex
                                / empty                 ;
                HexAlts         : num
                                | anyafAF    ;  
                repNum1         : < number >            ;
                anyafAF         : any "abcdefABCDEF"    ;
```
#### ***Commenting, Elements, and Attributes***
```
        Content         : RepContentAltsx0              ;
                RepContentAltsx0 : ContentAlts RepContentAltsx0
                                / empty                 ;
                ContentAlts     : COMMENT 
                                | Element 
                                | REFERENCE 
                                | charData              ;
        Misc 	         : COMMENT 
                         | spaceEsc                     ; 
                RepMisc0 : Misc RepMisc0 
                        / empty                         ;

        COMMENT        : ComStart ComEnterior clCarr1   ;
                ComStart        : opCarr1 excla DubDash ;
                DubDash         : "--"            ;
                ComEnterior     : DubDash 
                                / lets ComEnterior      ;


        Element        : opCarr1 NAME RepSAttx0 optSpaceEsc  ElemCloseAlts ;
                SAtt          : spaceEsc Attribute ;
                RepSAttx0     : SAtt RepSAttx0  
                              / empty ;
                ElemCloseAlts : clCarr1 Content opCarr2 NAME optSpaceEsc  clCarr1 
                              | clCarr2 ;
                opCarr1 : '<'     ;
                opCarr2 : '<' '/' ;
                clCarr1 : '>'     ;
                clCarr2 : '/' '>' ;
                excla   : '!'     ;

        Attribute       : NAME optSpaceEsc eq optSpaceEsc ATT_VALUE ;
        Eq              : optSpaceEsc eq optSpaceEsc    ;

        optSpaceEsc     : [ < any " \t\r\n" > ]         ;
        spaceEsc        : < any " \t\r\n" >             ;
        charData        :  < any "^<&" >                ;
        eq              : '='                           ;
```
#### ***Names, Encoding, and (Whitespace/Escape) Characters***
```
        NAME           : LetColonAlts RepNameChar0 ;
                LetColonAlts      : lets | anyColUn                     ;
                RepNameChar0      :  NAME_CHAR RepNameChar0 / empty     ;
                anyColUn          : any "_:"                            ; 

        NAME_CHAR      : lets 
                       | num
                       | anyDotDashEtc2                                 ;
        anyDotDashEtc2 : any "\\-._:"                                   ;

        EncName        : lets LetDigSymAltsRepx0                        ;
                LetDigSymAltsRepx0 : LetDigSymAlts LetDigSymAltsRepx0
                                   / empty                              ;
                LetDigSymAlts      : lets 
                                   | num
                                   | anyDotDashEtc                      ;
                anyDotDashEtc      : any "._\\-"                        ;
                lets               : letter                             ;
                num                : number                                ;

```

### **COPYRIGHT AND LICENSING INFORMATION**
**Copyright 2021 Brynn Harrington and Emily Hoppe**

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License.

You may obtain a copy of the License [here](http://www.apache.org/licenses/LICENSE-2.0) or at:

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.