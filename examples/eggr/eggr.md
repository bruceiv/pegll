## Eggr grammar
Modification of `eggr` grammar from [Egg](https://github.com/bruceiv/egg/blob/deriv/grammars/eggr.egg) to test an example structure.

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
