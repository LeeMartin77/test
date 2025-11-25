# Test 001: Fix Multiplication Bug

## Problem Description
There is a buggy implementation of multiplication in this server - fix it and get it returning the right results.

## Getting Started
The multiplication endpoint can be found at `/mul` and should work similarly to the add and subtract endpoints.

## What You Need to Do
1. **Identify the Issue**: Try running the server and testing the multiply endpoint:
   ```bash
   curl "http://localhost:8080/mul?a=5&b=3"
   ```
   You should notice that something goes wrong.

3. **Fix the Issue**: 
   - Ensure the mathematical operation is correct (should multiply a × b)
   - Test your fixes thoroughly

## Expected Behavior
After fixing the bugs:
```bash
curl "http://localhost:8080/mul?a=5&b=3"
# Should return: 15.00

curl "http://localhost:8080/mul?a=4&b=7"
# Should return: 28.00
```

## Testing Your Solution
Make sure to test both:
- Valid input cases (both parameters provided and valid numbers)
- Invalid input cases (missing parameters, invalid numbers)
