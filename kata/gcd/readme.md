# Greatest Common Divisor

Task is to calculate the greates common divisor or also known as greates common factor of two numbers. There are three methods to get to the right answer prepared all are based on the Euclid's algorithm. 

## Method 1 - the "old" one
The first step is to check if the first number n1 is zero, in this case we will return the first number as the answer. After we have done that, we loop until the the second number n2 will be zero and we return n1 as the gcd.
In the loop we are substract the lesser value from the higher value of both numbers. If the numbers are equal, this means that both numbers are equal and so we return n1.

### Pseudocode
```pascal
function euclid_old(n1:integer, n2:integer): integer
begin
	if n1 = 0 then return n2;
	while n2 <> 0 do 
		if n1 > n2 then
			n1 := n1 - n2;
		else
			n2 := n2 - n1;
	endWhile
	return n1
end;	
```

## Method 2 - the "modern" one
The trick in this method is to use the mod operater ( _you have to import the math lib to use that method_ )
You can do this on two ways, iterative or recursive.

### Pseudocode iterative
```pascal
function euclid_modernIterative(n1:integer, n2:integer): integer
begin
	while n2 <> 0 do 
		rest := n1 modulo n2;
		n1 := n2;
		n2 := rest;
	endWhile;
	return n1
end;	
```

### Pseudocode recursive
```pascal
function euclid_modernRecursive(n1:integer, n2:integer): integer
begin
	if n2 = 0 then return n1
	else 
		return euclid_modernRecursive(n2, n1 modulo n2)
end;	
```

## Sources

http://en.wikipedia.org/wiki/Greatest_common_divisor

https://www.khanacademy.org/math/cc-sixth-grade-math/cc-6th-factors-and-multiples/cc-6th-gcf/v/greatest-common-divisor