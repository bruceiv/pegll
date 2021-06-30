Simple identifier grammar to test negative lookahead

Author: Aaron Moss

To ensure tokenization does not interfere with the PEG semantics in this 
version of the grammar, all multi-character rules are semantic, rather than 
lexical.
```
Ident : !Keyword idChar RepidChar0x ;
Keyword : "i" "f" | "f" "o" "r" ;

idChar : letter | number ;
RepidChar0x : idChar RepidChar0x / empty ;
```