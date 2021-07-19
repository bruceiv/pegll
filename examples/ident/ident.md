Simple identifier grammar to test negative lookahead

Author: Aaron Moss

```
package "ident"
```

To ensure tokenization does not interfere with the PEG semantics in this 
version of the grammar, all multi-character rules are semantic, rather than 
lexical.
```
Ident : !Keyword IdChar RepidChar0x ;
Keyword : "i" "f" | "f" "o" "r" ;

IdChar : "i" | "f" | "o" | "r" | idChar ;
idChar : ( letter | number ) ;
RepidChar0x : IdChar RepidChar0x / empty ;
```