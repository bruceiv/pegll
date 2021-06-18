## **`XML` Grammar**
### **AUTHORSHIP INFORMATION**
#### *Authors :* Brynn Harrington and Emily Hoppe Copyright (C) 2021
#### *Creation Date :* June 11, 2021 
#### *Last Modified :* June 18, 2021
#### *Copyright and Licensing Information :* See end of file.

### **GENERAL DESCRIPTION**
An originally Egg Parsing grammar created by Dr. Aaron Moss ported into the GoGLL grammar for the `XML` language tests. Modification of `XML` grammar from [Egg](https://github.com/bruceiv/egg/blob/deriv/grammars/XML-u.egg) to test `XML` input files under the parser generated.
### **`XML` Grammar Guide**
NEED TO FINISH ONE GRAMMAR IS WORKING 
#### ***Top-Level Parsing***
```
package "XML"

Document       : prolog Element repMisc0 ;

Content        : { COMMENT 
               | Element 
               | REFERENCE 
               | CHAR_DATA } ;

prolog 	       : optXMLDecl repMisc0 ;
XMLDecl        : "<?xml" VersionInfo optEncodDecl optS "?>" ;
        optXMLDecl : [ XMLDecl ] ;

VersionInfo    : S "version" Eq quoVerNum ;
        quoVerNum    : '\'' VersionNum '\''  
                     | '\"' VersionNum '\"' ;
VersionNum     : < NAME_ CHAR > ;
EncodingDecl   : S "encoding" Eq quoEncNam ;
        quoEncNam    : '\'' EncName '\''  
                     | '\"' EncName '\"' ;
        optEncodDecl : [ EncodingDecl ] ;
```
#### ***Values and References***
```
ATT_VALUE       : '\"' dubCondClose 
                | '\'' sinCondClose ;
        singCondClose :   '\'' 
                      / SymRefAlts singCondClose ;
        dubCondClose  :   '\"' 
                      / SymRefAlts dubCondClose ;
                SymRefAlts   : any "^<&" 
                             | REFERENCE  ;

REFERENCE       : ENTITY_REF 
                | CHAR_REF ;
ENTITY_REF      : '&' NAME ';' ;
CHAR_REF        : "&#x" hex ';' 
                | "&#" repNum1 ';' ;
        hex             : < number 
                        | any "abcdefABCDEF" > ;
        repNum1         : < number > ;
```
#### ***Commenting, Elements, and Attributes***
```
Misc 	         : COMMENT 
                 | S ;
        repMisc0 : { Misc } ;

COMMENT        : "<!--" comEnterior '>' ;
        comEnterior     : "--" 
                        / letter comEnterior ;

Element        : '<' NAME RepSAtt0 optS elemCloseAlts ;
        RepSAtt0      : { S Attribute } ;
        elemCloseAlts : '>' Content  "</" NAME optS '>' 
                      | "/>" ;

Attribute      : NAME optS '=' optS ATT_VALUE ;

Eq             : optS '=' optS ;
```
#### ***Names, Encoding, and (Whitespace/Escape) Characters***
```
NAME           : letColonAlts repNameChar0 ;
        letColonAlts      : letter | any "_:" ;
        repNameChar0      : { NAME_CHAR } ;
NAME_CHAR      : letter 
               | number 
               | any "\-._:" ;

EncName        : letter letDigSymAlts ;
        letDigSymAlts     : { letter 
                          | number 
                          | any "._\-" } ;

SP             : < any " \t\r\n" > ;
        optS              : [ S ] ;
CHAR_DATA      :  < any "^<&" > ; 
```
#
### **COPYRIGHT AND LICENSING INFORMATION**
**Copyright 2021 Brynn Harrington and Emily Hoppe**

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License.

You may obtain a copy of the License [here](http://www.apache.org/licenses/LICENSE-2.0) or at:

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.