Matches the a^n b^n c^n (n >= 1) grammar which can't be matched by a CFG

Author: Aaron Moss 
  (adapted from Bryan Ford's "Packrat Parsing: a Practical Linear-Time Algorithm with Backtracking")

```
package "anbncn"
```

To ensure tokenization does not interfere with the PEG semantics in this 
version of the grammar, all rules are semantic, rather than lexical.
```
G1 : &Ac "a" Repa0x B1 ;
A1 : "a" A1 "b" / "a" "b" ;
B1 : "b" B1 "c" / "b" "c" ;

Ac : A1 "c" ;
Repa0x : "a" Repa0x / empty ;
```
