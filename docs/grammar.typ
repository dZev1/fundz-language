= Grammar

$
  E ::= & I \
      | & E + E \
      | & E - E \
      | & E * E \
      | & E \/ E \
      | & E < E \
      | & E "<=" E \
      | & E > E \
      | & E ">=" E \
      | & E "==" E \
      | & E "!=" E \
      | & E "&&" E \
      | & E "||" E \
      | & E "^|" E \
      | & E <- E \
      | & E => E \
      | & E "|&" E \
      | & I : T \
      | & E : T \
      | & E :: E \
      | & E "^" E \
      | & "if" (E) "then" {E} "else" {E} \
      | & (E) \
      | & {E} \
      | & [E] \
      | & "let" I : T = E \
      | & <E,E> \
      | & [E (,E)^*] \
      | & (I) "=>" {E} \
$

$
  T ::= & "int" \
      | & "bool" \
      | & "char" \
      | & "float" \
      | & T -> T \
      | & [T] \
      | & <T, T> \
$
