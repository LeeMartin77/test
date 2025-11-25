# Test 003: Business Logic Endpoint

## Requirement
We need to add a handler under `/businesslogic` - it should take four float parameters (a-d). It should then:

- sum a and b
- then subtract c from result
- then multiply result by d
- finally divide result by a

## Expected Result
```bash
curl "http://localhost:8080/businesslogic?a=10&b=5&c=3&d=2"
# Calculation: ((10 + 5) - 3) * 2 / 10 = 2.40
# Returns: 2.40
```
