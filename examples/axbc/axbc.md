## A~BC grammar
Modification of `abc` grammar from [Egg](https://github.com/bruceiv/egg/blob/deriv/grammars/abc.egg) to meaningfully exercise ordered choice by eliminating `"ab"` branch.

```
package "axbc"

AxBC : AorB "c" ;

AorB : as / "ab" ;

as : { 'a' } ;
```

