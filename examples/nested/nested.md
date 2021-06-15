## Nested grammar
Modification of `nested` grammar from [Egg](https://github.com/bruceiv/egg/blob/deriv/grammars/nested.egg) to test an example structure similar to XML.

```
package "nested"

String      : Content ;
Content     : { Parens 
            / Char } ;
Parens      : '(' Content ')' ;
Char        : < letter > ;

```
