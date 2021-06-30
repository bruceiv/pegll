Matches the a^n b^n c^n (n >= 1) grammar which can't be matched by a CFG

Author: Aaron Moss 
  (adapted from Bryan Ford's "Packrat Parsing: a Practical Linear-Time Algorithm with Backtracking")

To ensure tokenization does not interfere with the PEG semantics in this 
version of the grammar, all rules are semantic, rather than lexical.
```
G : &Ac "a" Repa0x B ;
A : "a" A "b" / "a" "b" ;
B : "b" B "c" / "b" "c" ;

Ac : A "c" ;
Repa0x : "a" Repa0x / empty ;
```
