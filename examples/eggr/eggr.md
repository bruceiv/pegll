# **`eggr` Grammar**

### **AUTHORSHIP INFORMATION**
#### *Author :* Brynn Harrington and Emily Hoppe Copyright (C) 2021
#### *Adapted from :* Aaron Moss's [`eggr` Egg Grammar](https://github.com/bruceiv/egg/blob/deriv/grammars/eggr.egg)
#### *Creation Date :* June 11, 2021 
#### *Last Modified :* June 18, 2021
#### *Copyright and Licensing Information :* See end of file.

###  **GENERAL DESCRIPTION**
An originally Egg Parsing grammar created by Aaron Moss ported into the GoGLL grammar to test a given structure.Modification of `eggr` grammar from [Egg](https://github.com/bruceiv/egg/blob/deriv/grammars/eggr.egg) to test an example structure.
### **STATUS ON GRAMMAR**
#### *Markdown File Creation:* Not working 
#### *Parser Generated :* Incomplete
#### *Test File Creation:* Incomplete
#### *Testing Results:* Unknown
### **`eggr` Grammar Guide**
NEED TO FIX:
- end of line MAY NOT WORK PROPERLY
- space 

```
package "eggr"

```
`LineOrBlock` represents the semantic rule for either a line or a block comment. 
`!line_comment` is a lexical rule representing a C-style line comment. Everything from the first slash to the end of line is a comment. 
`!block_comment` is a lexical rule representing a C-style block comment. Everything between and including `/*` and `*/` is a comment. 
The `!` in front of `!line_comment` and `!block_comment` instructs the lexer to suppress those tokens. See the [grammar for details.](../../gogll.md) 
*Note:* `!line_comment` and `!block_comment` were taken from [comments.md.](https://github.com/bruceiv/pegll/tree/main/examples/comments) 
```
LineOrBlock     : line_comment 
                | block_comment ;
!line_comment   : '/' '/' {not "\n"} ;
!block_comment  : '/''*' 
                { not "*" 
                | '*' not "/" 
                } '*''/' ;
end_of_line     : any "\r\n" ;

```
#### ORIGINAL GRAMMAR
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
                        | LineOrBlock }

        space           : ' ' 
                        | '\t' 
                        | end_of_line ;



        end_of_line     : "\r\n" 
                        / '\n' 
                        / '\r' ;

#### PARTIALLY WORKING GRAMMAR
!line_comment : '/' '/' {not "\n"} ;
!block_comment : '/''*' {not "*" | '*' not "/"} '*''/' ;
#
### **COPYRIGHT AND LICENSING INFORMATION**
**Copyright 2021 Brynn Harrington and Emily Hoppe**

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License.

You may obtain a copy of the License [here](http://www.apache.org/licenses/LICENSE-2.0) or at:

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.