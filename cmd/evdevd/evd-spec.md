# The Evd Programming Language Specification

## 1. Introduction

**Evd** is a domain-specific language (**DSL**) designed to map Linux
`evdev` devices to a `uinput` device managed by each instance of the
`evdevd` (**evdev daemon**) program. Each `evdevd` instance operates with
its own script and corresponding `uinput` device.

**Evd** combines both **imperative** and **declarative** paradigms. It
supports basic types such as `int`, `bool`, `string`, and `float`, along
with control structures like `if`, `switch`, and `for` loops.

The syntax is inspired by **Go**, aiming to provide a declarative feel
while allowing imperative functions for powerful scripting capabilities.
Global variables are **mutex-protected** by default. The language supports
**closures** and treats **functions as first-class citizens**, enabling
expressive functional abstractions.

**Evd** is designed to be **interpreted**, similar to shell scripting
languages. This allows for **rapid development**, **easier debugging**,
and **dynamic execution** without a separate compilation step.

## Notation

The syntax of Evd is described using **Backus-Naur Form (BNF)**, a formal
notation for defining the grammar of programming languages. Each rule
consists of a non-terminal symbol, followed by the symbol `::=`, and then
a sequence of terminals and/or non-terminals that define its structure.

- **Non-terminals** are enclosed in angle brackets (e.g., `<identifier>`)
  and represent abstract syntactic categories.
- **Terminals** are literal characters or tokens that appear in the source
  code.
- **Alternatives** are separated by the `|` symbol.
- **Repetition and optionality** are expressed through recursive rules or
  empty productions (`""`).
- **Escaping**: In terminal tokens, only the double quote character `"` is
  escaped using a backslash (`\"`).

This notation provides a precise and modular way to specify Evd’s syntax,
making it easier to reason about parsing behavior and implement a lexer or
parser that conforms to the language’s rules.

### Source code representation

Evd source files are encoded in UTF-8 and consist of Unicode text. The
language does not normalize characters, so a precomposed accented
character is treated differently from a decomposed sequence of accent and
base letter.

Each Unicode code point is considered distinct, including case
differences between uppercase and lowercase letters.

**Implementation note:** To maintain compatibility with external tools,
interpreters may reject the NUL character (`U+0000`) in source files.

**Implementation note:** Interpreters may ignore a UTF-8 byte order mark
(`U+FEFF`) if it appears as the first code point in the file. However, a
BOM may be disallowed elsewhere in the source.

### Characters

The following terms define specific Unicode character categories used in
Evd. These classifications follow Unicode Standard 16.0. When certain
constraints cannot be captured directly in BNF, such as excluding specific
code points, additional prose is used to clarify the intended behavior.

```
<newline>        ::= U+000A (newline)
<unicode_char>   ::= any Unicode code point excluding U+000A (newline)
<unicode_letter> ::= any Unicode code point in the Letter categories
                     Lu | Ll | Lt | Lm | Lo
<unicode_digit>  ::= any Unicode code point in the Number category Nd
```

### Letters and digits

The underscore character _ (U+005F) is considered a lowercase letter.

```
<letter>        ::= <unicode_letter> | "_"
<decimal_digit> ::= "0" | "1" | "2" | "3" | "4" | "5" | "6" | "7" | "8" |
                    "9"
<binary_digit>  ::= "0" | "1"
<octal_digit>   ::= "0" | "1" | "2" | "3" | "4" | "5" | "6" | "7"
<hex_digit>     ::= "0" | "1" | "2" | "3" | "4" | "5" | "6" | "7" | "8" |
                    "9" |
                    "A" | "B" | "C" | "D" | "E" | "F" |
                    "a" | "b" | "c" | "d" | "e" | "f"
```

## Lexical elements

### Comments

Comments are used to document Evd scripts and come in two styles:

- **Line comments** begin with `//` and continue until the end of the
  line.

- **General comments** begin with `/*` and end at the first occurrence of
  `*/`.

Comments cannot begin inside a string literal or another comment. If a
general comment contains no newline characters, it behaves like a space.
Otherwise, it behaves like a newline.

### Tokens

Tokens are the building blocks of Evd programs. They fall into four
categories: *identifiers*, *keywords*, *operators and punctuation*, and
*literals*.

Whitespace, consisting of spaces (`U+0020`), horizontal tabs (`U+0009`),
carriage returns (`U+000D`), and newlines (`U+000A`), is ignored except
when needed to separate adjacent tokens that would otherwise merge.

A newline or end-of-file may also trigger the insertion of a semicolon.
When scanning input, the lexer always selects the longest valid sequence
of characters as the next token.

### Semicolons

Semicolons (`;`) are used as statement terminators in Evd's formal syntax.
However, most semicolons can be omitted in scripts thanks to two automatic
insertion rules:

1. A semicolon is automatically inserted after the final token on a line
   if that token is:

   - an identifier
   - an integer, floating-point, or string literal
   - one of the keywords `break`, `continue`, `fallthrough`, or `return`
   - one of the operators or punctuation: `++`, `--`, `)`, `]`, or `}`

2. A semicolon may be omitted before a closing `)` or `}` to allow complex
   statements on a single line.

Code examples in this document omit semicolons where permitted to reflect
idiomatic usage.

### Identifiers

Identifiers are used to name program entities such as variables and types.
An identifier consists of one or more letters and digits, and must begin
with a letter.

```
<identifier>       ::= <letter> <identifier_tail>
<identifier_tail>  ::= "" | <letter> <identifier_tail> |
                       <unicode_digit> <identifier_tail>
```

```
a
_x9
ThisVariableIsExported
αβ
```

Some identifiers are predeclared and have special meaning within the
language.

### Keywords

The following keywords are reserved and may not be used as identifiers.

```
break       const        continue     else
fallthrough for          func         if
import      map          range        return
switch
```

### Operators and punctuation

The following character sequences are recognized as operators and
punctuation in Evd. This includes arithmetic, logical, comparison,
assignment, and grouping symbols.

```
+     &     +=     &=     &&     ==     !=     (     )
-     |     -=     |=     ||     <      <=     [     ]
*     ^     *=     ^=     >      >=     &^=    {     }
/     <<    /=     <<=    ++     =      ,      ;     &^
%     >>    %=     >>=    --     !      ...    .     :
```

### Integer literals

Integer literals represent constant whole numbers. They may begin with an
optional prefix to indicate a non-decimal base: `0b` or `0B` for binary,
`0`, `0o`, or `0O` for octal, and `0x` or `0X` for hexadecimal. A
standalone `0` is treated as a decimal zero. In hexadecimal literals, the
letters `a`–`f` and `A`–`F` correspond to values 10 through 15.

For improved readability, underscores (`_`) may be placed after the base
prefix or between digits. These underscores have no effect on the
literal’s value.

```
<int_lit>             ::= <decimal_lit> | <binary_lit> | <octal_lit> |
                          <hex_lit>

<decimal_lit>         ::= "0" | <decimal_lit_start> <decimal_digits>
<decimal_lit_start>   ::= "1" | "2" | "3" | "4" | "5" | "6" | "7" | "8" | "9"

<binary_lit>          ::= "0b" <binary_lit_tail> | "0B" <binary_lit_tail>
<binary_lit_tail>     ::= "_" <binary_digits> | <binary_digits>

<octal_lit>           ::= "0" <octal_lit_tail> | "0o" <octal_lit_tail> |
                          "0O" <octal_lit_tail>
<octal_lit_tail>      ::= "_" <octal_digits> | <octal_digits>

<hex_lit>             ::= "0x" <hex_lit_tail> | "0X" <hex_lit_tail>
<hex_lit_tail>        ::= "_" <hex_digits> | <hex_digits>

<decimal_digits>      ::= <decimal_digit> |
                          <decimal_digit> <decimal_digits_tail>
<decimal_digits_tail> ::= "" | "_" <decimal_digit> <decimal_digits_tail> |
                          <decimal_digit> <decimal_digits_tail>

<binary_digits>       ::= <binary_digit> |
                          <binary_digit> <binary_digits_tail>
<binary_digits_tail>  ::= "" | "_" <binary_digit> <binary_digits_tail> |
                          <binary_digit> <binary_digits_tail>

<octal_digits>        ::= <octal_digit> |
                          <octal_digit> <octal_digits_tail>
<octal_digits_tail>   ::= "" | "_" <octal_digit> <octal_digits_tail> |
                          <octal_digit> <octal_digits_tail>

<hex_digits>          ::= <hex_digit> | <hex_digit> <hex_digits_tail>
<hex_digits_tail>     ::= "" | "_" <hex_digit> <hex_digits_tail> |
                          <hex_digit> <hex_digits_tail>
```

```
42
4_2
0600
0_600
0o600
0O600       // second character is capital letter 'O'
0xBadFace
0xBad_Face
0x_67_7a_2f_cc_40_c6
170141183460469231731687303715884105727
170_141183_460469_231731_687303_715884_105727

_42         // an identifier, not an integer literal
42_         // invalid: _ must separate successive digits
4__2        // invalid: only one _ at a time
0_xBadFace  // invalid: _ must separate successive digits
```

### Floating-point literals

Floating-point literals represent constant real numbers in either decimal
or hexadecimal form.

A **decimal floating-point literal** consists of:
- an integer part (decimal digits)
- a decimal point
- a fractional part (decimal digits)
- an optional exponent part (`e` or `E`, followed by an optional sign and
  decimal digits)

Either the integer part or the fractional part may be omitted. Likewise,
the decimal point or exponent may be omitted. The exponent scales the
mantissa (the combined integer and fractional part) by `10^exp`.

A **hexadecimal floating-point literal** begins with a `0x` or `0X` prefix
and includes:
- an integer part (hex digits)
- a radix point
- a fractional part (hex digits)
- a required exponent part (`p` or `P`, followed by an optional sign and
  decimal digits)

Either the integer or fractional part may be omitted. The radix point may
also be omitted, but the exponent is mandatory. This format follows the
IEEE 754-2008 §5.12.3 specification. The exponent scales the mantissa by
`2^exp`.

For readability, underscores (`_`) may appear after the base prefix or
between digits. These do not affect the literal’s value.

```
<float_lit>                  ::= <decimal_float_lit> | <hex_float_lit>

<decimal_float_lit>          ::= <decimal_digits> "."
                                 <decimal_float_lit_fraction>
                                 <decimal_float_lit_exponent> |
                                 <decimal_digits> <decimal_exponent> |
                                 "." <decimal_digits>
                                 <decimal_float_lit_exponent>
<decimal_float_lit_fraction> ::= "" | <decimal_digits>
<decimal_float_lit_exponent> ::= "" | <decimal_exponent>

<decimal_exponent>           ::= "e" <decimal_exponent_tail> |
                                 "E" <decimal_exponent_tail>
<decimal_exponent_tail>      ::= "+" <decimal_digits> | "-" <decimal_digits> |
                                 <decimal_digits>

<hex_float_lit>              ::= "0x" <hex_mantissa> <hex_exponent> |
                                 "0X" <hex_mantissa> <hex_exponent>
<hex_mantissa>               ::= <hex_mantissa_start> "." |
                                 <hex_mantissa_start> "." <hex_digits> |
                                 <hex_mantissa_start> |
                                 "." <hex_digits>
<hex_mantissa_start>         ::= "_" <hex_digits> | <hex_digits>

<hex_exponent>               ::= "p" <decimal_exponent_tail> |
                                 "P" <decimal_exponent_tail>
<hex_exponent_tail>          ::= "+" <decimal_digits> | "-" <decimal_digits> |
                                 <decimal_digits>
```

```
0.
72.40
072.40       // == 72.40
2.71828
1.e+0
6.67428e-11
1E6
.25
.12345E+5
1_5.         // == 15.0
0.15e+0_2    // == 15.0

0x1p-2       // == 0.25
0x2.p10      // == 2048.0
0x1.Fp+0     // == 1.9375
0X.8p-0      // == 0.5
0X_1FFFP-16  // == 0.1249847412109375
0x15e-2      // == 0x15e - 2 (integer subtraction)

0x.p1        // invalid: mantissa has no digits
1p-2         // invalid: p exponent requires hexadecimal mantissa
0x1.5e-2     // invalid: hexadecimal mantissa requires p exponent
1_.5         // invalid: _ must separate successive digits
1._5         // invalid: _ must separate successive digits
1.5_e1       // invalid: _ must separate successive digits
1.5e_1       // invalid: _ must separate successive digits
1.5e1_       // invalid: _ must separate successive digits
```

### String literals

String literals represent constant byte sequences derived from a series of
characters. Evd supports two forms: *raw string literals* and *interpreted
string literals*.

**Raw string literals** are enclosed in back quotes, like `` `foo` ``.
Inside the quotes, any character is allowed except the back quote itself.
The contents are treated as-is, without escape processing, and are
implicitly UTF-8 encoded. Backslashes have no special meaning, and
newlines are preserved. Carriage return characters (`\r`) are discarded.

**Interpreted string literals** are enclosed in double quotes, like `"bar"`.
Within the quotes, any character may appear except unescaped double quotes
and newlines. Escape sequences follow standard conventions: octal (`\nnn`)
and hexadecimal (`\xnn`) escapes represent individual bytes. Other escapes
encode Unicode code points using UTF-8, which may result in multi-byte
sequences.

For example, `\377` and `\xFF` both yield a single byte `0xFF`, while
`ÿ`, `\u00FF`, `\U000000FF`, and `\xc3\xbf` all produce the two-byte UTF-8
sequence `0xc3 0xbf` for character U+00FF.

```
<string_lit>                ::= <raw_string_lit> | <interpreted_string_lit>

<raw_string_lit>            ::= "`" <raw_string_middle>  "`"
<raw_string_middle>         ::= "" | <unicode_char> <raw_string_middle> |
                              <newline> <raw_string_middle>

<interpreted_string_lit>    ::= "\"" <interpreted_string_middle> "\""
<interpreted_string_middle> ::= "" |
                                <unicode_value> <interpreted_string_middle> |
                                <byte_value> <interpreted_string_middle>

<unicode_value>             ::= <unicode_char> | <little_u_value> |
                                <big_u_value> | <escaped_char>
<byte_value>                ::= <octal_byte_value> | <hex_byte_value>
<octal_byte_value>          ::= "\" <octal_digit> <octal_digit> <octal_digit>
<hex_byte_value>            ::= "\x" <hex_digit> <hex_digit>
<little_u_value>            ::= "\u" <hex_digit> <hex_digit> <hex_digit>
                                <hex_digit>
<big_u_value>               ::= "\U" <hex_digit> <hex_digit> <hex_digit>
                                <hex_digit> <hex_digit> <hex_digit>
                                <hex_digit> <hex_digit>
<escaped_char>              ::= "\a" | "\b" | "\f" | "\n" | "\r" | "\t" |
                                "\v" | "\\" | "\'" | "\\""
```

```
`abc`                // same as "abc"
`\n
\n`                  // same as "\\n\n\\n"
"\n"
"\""                 // same as `"`
"Hello, world!\n"
"日本語"
"\u65e5本\U00008a9e"
"\xff\u00FF"
"\uD800"             // illegal: surrogate half
"\U00110000"         // illegal: invalid Unicode code point
```

These examples all represent the same string:

```
"日本語"                               // UTF-8 input text
`日本語`                               // UTF-8 input text as a raw literal
"\u65e5\u672c\u8a9e"                   // the explicit Unicode code points
"\U000065e5\U0000672c\U00008a9e"       // the explicit Unicode code points
"\xe6\x97\xa5\xe6\x9c\xac\xe8\xaa\x9e" // the explicit UTF-8 bytes
```

If the source code represents a character as two code points, such as a
combining form involving an accent and a letter, it will appear as two
code points in a string literal.
