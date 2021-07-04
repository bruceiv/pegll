## **`string` GRAMMAR**
### **AUTHORSHIP INFORMATION**
#### *Authors :* Brynn Harrington Copyright (C) 2021
#### *Creation Date :* July 4, 2021 
#### *Last Modified :* July 4, 2021
#### *Copyright and Licensing Information :* See end of file.

### **GENERAL DESCRIPTION**
A GoGLL parsing grammar to test reading a JSON-syntax string.

### **STATUS ON GRAMMAR**
#### *Markdown File Creation:* Complete
#### *Parser Generated :* Complete
#### *Test File Creation:* Incomplete
#### *Testing Results:* Unknown

### **`string` GRAMMAR GUIDE**
The following grammar tests input strings from the JSON language syntax.
```
package "string"

String          : string_ns                  ;
!string_ns      : '"' { not "\"" letter } '"' ;  
```
WS              : EscOrComment WS
                / empty                                 ;
EscOrComment    : escCharSpace 
                / line_comment
                / block_comment                         
                / empty                                 ;
escCharSpace    : any " \t\r\n"                     ;

!line_comment   : '/' '/' { not "\r\n" }  ;               
!block_comment  : '/''*' { not "*" | '*' not "/" } '*''/' ;
```

```

### **COPYRIGHT AND LICENSING INFORMATION**
**Copyright 2021 Brynn Harrington**

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License.

You may obtain a copy of the License [here](http://www.apache.org/licenses/LICENSE-2.0) or at:

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.