## Miniegg grammar
Modification of `miniegg` grammar from [Egg](https://github.com/bruceiv/egg/blob/deriv/grammars/miniegg.egg) to test an example structure.

```
package "miniegg"

gram    : _ ruler ;
    ruler   : < rule > ;
rule    : id '=' _ exprr ;
id      : upcase _ ;
    exprr   : < expr > ;
expr    : id not '=' ;
_       : [ ' ' ] ;

```
