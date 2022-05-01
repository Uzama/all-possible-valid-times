# all-possible-valid-times
Given a list of 4 integers, find all possible valid 24 hour times (eg: 12:34) that the given integers can be assembled into and return the total number of valid times.

### rules
- You can not use the same number twice.
- Times such as 34:12 and 12:60 are not valid.
- Provided integers can not be negative and all integers within 0-9.

### example
```
  Input: [1, 2, 3, 4]
  Valid times: ["12:34", "12:43", "13:24", "13:42", "14:23", "14:32", "23:14", "23:41", "21:34", "21:43"]
  Return: 10
```
