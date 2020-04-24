# Input
    - rate of growth of time
- Define model machine
-  above factors
-  1 unit of time for 
    -  arithmetical l
    -  logic 
    -  assignement
    -  return

## Calculate

```
  rate of growth 
      time/input
```

Psudo code

```
      Sum(a,b)     // 1 unit of time
      {return a+b} // 1 unit of time
```
Function

```
   Tsum = 2 (constant)


```  

<b> Function </bb>
 
```
Sumoflist ( A[], n) //array, size of array
                            Units              Executions          Constant
{
    total = 0                // 1                  1                  c1
    for i=0 to n-1           //1 each = 2          n+1                c2
        total = total + Ai   //1 each = 2          n                  c3
    return total             //1                                      c4
}
```

### Analysis
```

   Tsum = 2 (constant)
  1 + 2(n+1) +2n + 1

  1 + 2n + 2 + 2n + 1
  4 + 4n

Tn   = cn + c'
c    = c2 + c3
c'   = c1 + c2 + c4


Tsum = K                             O(1)
Tsum of list c*n + c'                O(n)



Tsum of Matrix  =  an^2 + bn + c    On^2
```

[Asymptotic Notation](AsymptoticNotation.md)