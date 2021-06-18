# **`eggr` Grammar**

### **AUTHORSHIP INFORMATION**
#### *Author :* Brynn Harrington and Emily Hoppe Copyright (C) 2021
#### *Adapted from :* Aaron Moss's [`eggr` Egg Grammar](https://github.com/bruceiv/egg/blob/deriv/grammars/eggr.egg)
#### *Creation Date :* June 11, 2021 
#### *Last Modified :* June 18, 2021
#### *Copyright and Licensing Information :* See end of file.

###  **GENERAL DESCRIPTION**
An originally Egg Parsing grammar created by Aaron Moss ported into the GoGLL grammar to test a given structure.Modification of `eggr` grammar from [Egg](https://github.com/bruceiv/egg/blob/deriv/grammars/eggr.egg) to test an example structure.

### **`eggr` Grammar Guide**
NEED TO FINISH ONE GRAMMAR IS WORKING 

See the [grammar for details.](../../gogll.md)

### **STATUS ON GRAMMAR**
#### *Markdown File Creation:* Not working 
#### *Parser Generated :* Incomplete
#### *Test File Creation:* Incomplete
#### *Testing Results:* Unknown
```
package "eggr"

grammar         : _ rule_rep ; 
        rule_rep        : < rule > ;
rule            : identifier EQUAL choice ;

choice          : sequence pipe_seq ;
        pipe_seq        : { PIPE sequence } ;

sequence        : < expression > ;

expression      : AND primary 
                | NOT primary 
                | primary optStarPlus_rep ;
        optStarPlus_rep : [ OPT 
                        | STAR 
                        | PLUS ] ;

primary         : identifier not EQUAL 
                | OPEN choice CLOSE
                | char_literal
                | str_literal
                | char_class
                | ANY
                | EMPTY ;

identifier      : let_ let_num _ ;
        let_            : letter 
                        | _ ;
        let_num         : { letter 
                        | _ 
                        | number } ;

char_literal    : '\'' character '\'' _ ;
str_literal     : '\"' str '\"' _ ;
        str             :  { character } ;
char_class      : '[' unclosed_chars ']' _ ;
        unclosed_chars  : { not ']' character } ;
character       : not "\'\"\\" 
                | '\\' any "nrt\'\"\\" ;

EQUAL : '=' _ ;
PIPE  : '|' _ ;
AND   : '&' _ ;
NOT   : '!' _ ;
OPT   : '?' _ ;
STAR  : '*' _ ;
PLUS  : '+' _ ;
OPEN  : '(' _ ;
CLOSE : ')' _ ;
ANY   : '.' _ ;
EMPTY : ';' _ ;

_               : { space 
                | comment }

space           : ' ' 
                | '\t' 
                | end_of_line ;

comment         : line_comment 
                | block_comment ;

line_comment    : '/' '/' notNLn0 ;
        notNLn0         : {not "\n"} ;

block_comment   : '/''*' notStarAlts0 '*''/' ;
        notStarAlts0    : {not "*" | '*' not "/"} ;

end_of_line     : "\r\n" 
                / '\n' 
                / '\r' ;

```
#
### **COPYRIGHT AND LICENSING INFORMATION**
**Copyright 2021 Brynn Harrington and Emily Hoppe**

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License.

You may obtain a copy of the License [here](http://www.apache.org/licenses/LICENSE-2.0) or at:

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.