# **`ax` GRAMMAR**

### **AUTHORSHIP INFORMATION**
#### *Authors :* Brynn Harrington and Emily Hoppe Copyright (C) 2021
#### *Adapted from :* Aaron Moss's [`astar` Egg Grammar](https://github.com/bruceiv/egg/blob/deriv/grammars/abc.egg)
#### *Creation Date :* June 10, 2021 
#### *Last Modified :* June 24, 2021
#### *Copyright and Licensing Information :* See end of file.

### **GENERAL DESCRIPTION**
A modification of the `astar` [Egg](https://github.com/bruceiv/egg/blob/deriv/grammars/astar.egg) Parsing Grammar ported into GoGLL to test repeatability within the grammar. 

### **STATUS ON GRAMMAR**
#### *Markdown File Creation:* Complete
#### *Parser Generated :* Complete
#### *Test File Creation:* Incomplete
#### *Testing Results:* Unknown

### **`ax` GRAMMAR GUIDE**
The following grammar tests the repeatability of a character, in this case 'a', within GoGLL. `Repa0x` represents the start rule while `repa0x` consists of the repeated token of `a` zero or more times. The use of the bracketed expression`{}` within this grammar is to indicate a set of tokens or identifiers repeated zero or more times. See the [grammar for details.](../../gogll.md)

```
package "ax"

Repa0x : repa0x ;

repa0x : { 'a' } ;
```
#
### **COPYRIGHT AND LICENSING INFORMATION**
**Copyright 2021 Brynn Harrington and Emily Hoppe**

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License.

You may obtain a copy of the License [here](http://www.apache.org/licenses/LICENSE-2.0) or at:

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.

