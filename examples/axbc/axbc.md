# **`A~BC` GRAMMAR**
### **AUTHORSHIP INFORMATION**
#### *Author :* Aaron Moss Copyright (C) 2021
#### *Adapted from :* Aaron Moss's [`abc` Egg Grammar](https://github.com/bruceiv/egg/blob/deriv/grammars/abc.egg)
#### *Creation Date :* July 30, 2021 
#### *Last Modified :* July 30, 2021
#### *Copyright and Licensing Information :* See end of file.

###  **GENERAL DESCRIPTION**
A modification of `abc` [Egg](https://github.com/bruceiv/egg/blob/deriv/grammars/abc.egg) parsing grammar to test and meaningfully exercise ordered choice through the elimination of the `"ab"` branch.

### **STATUS ON GRAMMAR**
#### *Markdown File Creation:* Complete
#### *Parser Generated :* Complete
#### *Test File Creation:* Complete
#### *Testing Results:* Passed
#### *Errors:* None

### **`A~BC` Grammar Guide**
The following grammar tests ordered choice in GoGLL. This grammar attempts to exercise ordered choice through the elimination of the `"ab"` branch.
```
package "axbc"
```
`AxBC` is a semantic starting rule that is composed of the semantic rule `AorB` followed by the character 'c'.
```
AxBC    : AorB "c"  ;
```
`AorB` is a semantic rule that represents the ordered choice of a string beginning with `'a'` zero or more times, with the alternate choice being the string `"ab"`. Here the syntactic rule `/` is utilized for ordered choice. See the [grammar for details.](../../gogll.md). A key aspect of this grammar is that each character must be an individual token, otherwise the lexer's longest-match tokenization process defeats the PEG semantics.
```
AorB    : AStar / "a" "b" ;
```
`AStar` consists of the repeated token of `a` zero or more times. Similarly to `AorB`, this must be done in a semantic rule to preserve the proper PEG behavior.
```
AStar  : "a"*;
```
#
### **COPYRIGHT AND LICENSING INFORMATION**
**Copyright 2021 Aaron Moss**

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License.

You may obtain a copy of the License [here](http://www.apache.org/licenses/LICENSE-2.0) or at:

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.

