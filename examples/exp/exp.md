## Exp grammar
Modification of `exp` grammar from [Egg](https://github.com/bruceiv/egg/blob/deriv/grammars/exp.egg) to test an example grammar with possible exponential runtime under recursive descent.

```
package "exp"

S : A '\n' ;

A : 'a' A 'b' 
  | 'a' A 'c' 
  | empty ;

```
