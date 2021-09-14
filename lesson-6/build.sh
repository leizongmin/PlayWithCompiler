#!/usr/bin/env bash

antlr -Dlanguage=Go -o hello_parser Hello.g4
antlr -Dlanguage=Go -o calc_parser Calc.g4
