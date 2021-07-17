# **`Optional` GRAMMAR**

### **AUTHORSHIP INFORMATION**
#### *Authors :* Brynn Harrington and Emily Hoppe Copyright (C) 2021
#### *Designed to* test syntax operator `?`
#### *Creation Date :* July 15, 2021 
#### *Last Modified :* July 15, 2021
#### *Copyright and Licensing Information :* See end of file.

### **GENERAL DESCRIPTION**
A grammar to test the added `?` operator in `Pegll`.

### **STATUS ON GRAMMAR**
#### *Markdown File Creation:* Complete
#### *Parser Generated :* Complete
#### *Test File Creation:* Incomplete
#### *Testing Results:* Unknown

### **`Optional` GRAMMAR GUIDE**

```
package "Optional"

S1       : Required             ;

Required : "Required"           ;

Optional : Base?               ;

Base     : "Base"               ;

```
#
### **COPYRIGHT AND LICENSING INFORMATION**
**Copyright 2021 Brynn Harrington and Emily Hoppe**

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License.

You may obtain a copy of the License [here](http://www.apache.org/licenses/LICENSE-2.0) or at:

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.
