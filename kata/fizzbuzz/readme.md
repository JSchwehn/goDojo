# Fizz Buzz
The task is to count from 0 to 100 and while doing it to replace any number which is divisible by three the number gets replaced by the word fizz and any number divisible by five is replaced with buzz. Numbers divisible by five and three become fuzzbuzz.
To use this in a test, the function must return a slice of strings

```pascal
function Fuzzbuzz(): array of string
begin
	returnValues : array of string;
	for i := 0 to 100 do 
		if(i modulo 3 = 0) && (i modulo 5 = 0)) 
			returnValues[i] := "fizzbuzz";
		else if(i modulo 5 = 0)
			returnValues[i] := "buzz";
		else if(i modulo 3 = 0)
			returnValues[i] := "fuzz";
		else
			returnValues[i] := (string)i;
	endFor;
	return returnValues;
end;
```

## Source
http://en.wikipedia.org/wiki/Fizz_buzz