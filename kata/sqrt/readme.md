# Square Root of a numbers

Everyone who has taken the go tour http://tour.golang.org should know this one :)

The task is to implement the calculation of the square root of a number, using Newton's method.

### Pseudocode Newton
```pascal
function sqrt(n1:float): float
begin
	for abs(old - z)>0.001 do
		old := z;
		z := z - ((z^2)-x) / (2*z);
	endFor;
	return z
end;	
```