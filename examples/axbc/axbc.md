## A~BC grammar
Modification of `abc` grammar from [Egg](https://github.com/bruceiv/egg/blob/deriv/grammars/abc.egg) to meaningfully exercise ordered choice by eliminating `"ab"` branch.

```
package "axbc"

AxBC : AorB "c" ;

AorB : As / "a" "b" ;

As : "a" As / empty ;
```

### Previous versions
Version of the grammar using lexer (lexer's insistence on making everything a valid sequence of tokens defeats PEG-level ordered choice)

`AxBC : AorB "c" ;`
`AorB : as / "ab" ;`
`as : { 'a' } ;`


