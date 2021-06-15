## XML grammar
Modification of `XML` grammar from [Egg](https://github.com/bruceiv/egg/blob/deriv/grammars/XML-u.egg) to test XML gramar.

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

EncodingDecl   : S "encoding" Eq quoEncNam ;
        quoEncNam    : '\'' EncName '\''  
                     | '\"' EncName '\"' ;
        optEncodDecl : [ EncodingDecl ] ;

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
VersionNum     : < NAME_ CHAR > ;

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

S              : < any " \t\r\n" > ;
        optS              : [ S ] ;
CHAR_DATA      :  < any "^<&" > ; 

```
