# Test 002: Add Division Handler

## Requirements
We need to add a division handler to this project.

## Acceptance Criteria
- Create a new endpoint `/div` that accepts `a` and `b` query parameters
- The endpoint should return the result of a / b
- Follow the same patterns as existing endpoints (add, subtract, multiply)
- Handle division by zero appropriately - return an error message
- Return results formatted to 2 decimal places
- Use the existing utility functions for parameter parsing
- Add the new method to the domain service interface and implementation

## Expected API Behavior
```bash
# Successful division
curl "http://localhost:8080/div?a=15&b=3"
# Returns: 5.00

# Division by zero
curl "http://localhost:8080/div?a=10&b=0"
# Should return appropriate error message

# Invalid parameters
curl "http://localhost:8080/div?a=10"
# Should return parameter error message
```

## Definition of Done
- Division endpoint is functional and follows project patterns
- Error handling is consistent with other endpoints
- Division by zero is handled gracefully
- Code follows the same architectural structure as existing handlers