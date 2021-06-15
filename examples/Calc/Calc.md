## Calc grammar
Modification of `Calc` grammar from [Egg](https://github.com/bruceiv/egg/blob/deriv/grammars/Calc.egg) to test an example calculator grammar.

```
package "Calc"

expr : int = _ sum : psVal !.

sum : int = prod : psVal (
            PLUS prod : i { psVal += i; }
            | MINUS prod : i { psVal -= i; } )*
prod : int = elem : psVal (
             TIMES elem : i { psVal *= i; }
             | DIVIDE elem : i { psVal /= i; } )*
elem : int = OPEN sum : psVal CLOSE
             | num : psVal

num : int = < [0-9]+ > : s { psVal = atoi(s.c_str()); } _

PLUS : '+' _ ;
MINUS : '-' _ ;
TIMES : '*' _ ; 
DIVIDE : '/' _ ;
OPEN : '(' _ ; 
CLOSE : ')' _ ;

_ : { ' ' | '\t' } ;

```
