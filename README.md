# windy-judge

**windy-judge** is an automated testing tool designed to compare program input and output, verifying whether the execution results match expectations. It is ideal for algorithm competitions, automatic evaluation, and CI/CD testing workflows.

## ✨ Features

- **Automatic Comparison**: Compares actual program output against expected output.
- **Detailed Reports**: Displays test case results, including execution time and comparison details.
- **Command-Line Support**: Run directly via CLI for seamless integration into development workflows.

## 🚀 Usage

### 1️⃣ Run Tests

Execute the following command to run a test (replace `test` and `add` with your actual test case and program logic):

```bash

./windy-judge test add https://example.com
```

### 2️⃣ Example Output

After execution, you will see results similar to the following:

```bash

# Test Case 0 - Result:
----------------------------------------------
[Timestamps]
- Test Time: 2025-03-07 23:13:45
- Execution Time: 0.01s

Input:
1 2

Expected Output:
3

Program Output:
3

[Comparison Result] ✅ Accepted! Your output matches the expected result.

🎉 Congratulations! All 1 test case passed successfully! ✅🎯 Used time: 0.01s Keep up the great work! 🚀🔥
```

## ⚙️ Requirements

- Supports **Linux / macOS**
- Requires **Go1.24** 
- May require **Python** or other environments depending on test cases

## 📌 Future Plans

-  Support multi-threaded parallel testing
-  Enhance test result analysis

## 📜 License

This project is licensed under the **MIT License**, allowing free use and modification.
