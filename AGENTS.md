# Project Summary: Monkey Interpreter

## 🐒 Overview
Monkey is a custom, interpreted programming language implemented in Go. It utilizes a standard compiler/interpreter pipeline: Lexer -> Parser -> AST -> Evaluator -> Object.

## 📂 Structure & Key Components
*   **`lexer/`**: Tokenizes raw input.
*   **`token/`**: Defines language grammar (keywords, tokens).
*   **`parser/`**: Generates the Abstract Syntax Tree (AST).
*   **`ast/`**: Defines the hierarchical code structure.
*   **`evaluator/`**: Executes the AST.
*   **`object/`**: Defines runtime data types (Integer, Float, String, Error, etc.) and their operations.
*   **`main.go` / `repl/`**: Entry point and interactive console.

## 🐛 Known Issues & Future Work
*   **FIXED (Recent):** Division by zero is now handled gracefully by returning an `Error` object in `object/object.go`.
*   **Major Feature Gap:** The language lacks support for `while` and `for` loops.
*   **Major Feature Gap:** Missing dedicated variable assignment operator (`=`) for modifying existing variables.
*   **Lexer Improvement:** Needs robust handling for unterminated string literals.
*   **Runtime Improvement:** Floating-point comparisons should consider an epsilon tolerance for strict equality checks.