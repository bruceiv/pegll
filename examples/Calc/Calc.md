## Calc grammar
Modification of `Calc` grammar from [Egg](https://github.com/bruceiv/egg/blob/deriv/grammars/Calc.egg) to test an example calculator grammar.

```
package "Calc"

expr :  _ sum ;

sum  : prod { PLUS prod | MINUS } ;

prod : elem { TIMES elem | DIVIDE elem } ;

elem : OPEN sum CLOSE | num ;

num  :  < number > _ ;

PLUS   : '+' _ ;
MINUS  : '-' _ ;
TIMES  : '*' _ ; 
DIVIDE : '/' _ ;
OPEN   : '(' _ ; 
CLOSE  : ')' _ ;

_ : { ' ' | '\t' } ;

```
