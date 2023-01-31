# RUNNABLE README
[![release](https://github.com/R3DRUN3/runnable-readme/actions/workflows/release.yaml/badge.svg)](https://github.com/R3DRUN3/runnable-readme/actions/workflows/release.yaml)
[![Go Report Card](https://goreportcard.com/badge/github.com/R3DRUN3/runnable-readme)](https://goreportcard.com/report/github.com/R3DRUN3/runnable-readme)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
<br/>

<img src="images/gopher.png" width="350" height="350">

<br/>
Make your readme executable! 📖 ⚙️

<br/>

## Abstract
*Runnable Readme* is a simple utility that allows people to execute bash commands that are found in *README* (markdown) files in automatic fashion.
<br/>
The advantage is twofold:
1. Allow those who follow a documentation to avoid tedious copy-paste.
2. Encourage developers to write documentation that is as clear and reproducible as possible.

We can call this *Documentation as Code* 😊

## Instructions
Download this utility from the release or clone this repo and launch it locally (it requires go v1.19).
<br/>
example, clone the repo:  
```console
git clone https://github.com/r3drun3/runnable-readme.git && cd runnable-readme
```

<br/>

test it against a *markdown* file:  
```console
go run main.go exec examples/example.md
```

<br/>

Output:  
```console
Executing command number  1
Output: Hello World!

Executing command number  2
Output: My external IP is: 87.0.22.162
Sleeping for 3 seconds
Goodbye!

Executing command number  3
Output: where am I?
/tmp
let's create folders in loop
Sleeping for 2 seconds
removing the folders
Operation OK: true
```

<br/>





