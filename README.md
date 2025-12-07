# fundz-language

A project where I will try to implement a simple functional language, based on lambda calculus. The core functionalities will be variables, applications and abstractions (functions). I will also implement types and type inference.

## Implementation

### Lexer

The lexer will be implemented using regular expressions that are going to be turned into NFAs which will then be determinized.
Then we will read each character to form tokens.

### Parser

The parser will be implemented using a pushdown automaton.
