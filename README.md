# Ascii-art

**Ascii-art** is a Go program that takes a string as an argument and outputs a stylized ASCII representation of that string. With an optional color flag, users can also colorize specific parts of the ASCII text output.

## Features

- **Text to ASCII Art**: Converts a given string into a visual representation using ASCII characters.
- **Terminal output Alignment**: Changes the alignment of the output using the `--align=<type>` flag.
  - Specify a banner name to be used.
  - 4 different font to use for free.
  - 4 different alignment methods.

## Installation

**Clone the repository:**
```bash
git clone https://github.com/your-username/ascii-art.git
```
**Navigate to the project directory:**
```bash
cd ascii-art
```
**Run the program:**
```bash
go run . [STRING]
```

## Usage

### Basic Usage
```bash
go run . [STRING]
```
### Change Ascii Font
```bash
go run . [STRING] [BANNER]
```
### Change the alignment of the output
```bash
go run . [OPTION] [STRING] [BANNER]
```
### Examples

**Display text as ASCII art**
```bash
go run . "Ascii Art"
```
**Change Ascii Font**
```bash
go run . "Ascii Art" shadow
```
**Create .txt File with selected Font**
```bash
EX: go run . --align=right something standard
```
### Supported Fonts
#### Standard
```bash
       _                         _                      _  
      | |                       | |                    | | 
 ___  | |_    __ _   _ __     __| |   __ _   _ __    __| | 
/ __| | __|  / _` | | '_ \   / _` |  / _` | | '__|  / _` | 
\__ \ \ |_  | (_| | | | | | | (_| | | (_| | | |    | (_| | 
|___/  \__|  \__,_| |_| |_|  \__,_|  \__,_| |_|     \__,_| 
```
#### Shadow
```bash
         _|                      _|
  _|_|_| _|_|_|     _|_|_|   _|_|_|   _|_|   _|      _|      _| 
_|_|     _|    _| _|    _| _|    _| _|    _| _|      _|      _| 
    _|_| _|    _| _|    _| _|    _| _|    _|   _|  _|  _|  _|   
_|_|_|   _|    _|   _|_|_|   _|_|_|   _|_|       _|      _|     
```
#### Thinkertoy
```bash
 o  o           o             o
 |  |    o      | /           |
-o- O--o   o-o  OO   o-o o-o -o- o-o o  o 
 |  |  | | |  | | \  |-' |    |  | | |  | 
 o  o  o | o  o o  o o-o o    o  o-o o--O 
                                        | 
                                     o--o 
```
#### Other
```bash                                                    
           **     **                         
  ****   ******** ******     ****   **  **** 
**    **   **     **    ** ******** ****     
**    **   **     **    ** **       **       
  ****       **** **    **   ****** **
  ```
## Authors

Thanos Ziagakis ❄️
<!-- 🍉              -->
Sofia Busho ☃️
<!-- 🌸    -->
Maria Tzemanaki 🦌
<!-- 🍓           -->