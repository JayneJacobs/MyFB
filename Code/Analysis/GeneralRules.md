# General Rules


Analyze Time Complexity for:
 a) Very Large input size
 b) worst case scenario


 Example: 
 t(n) = n^3 + ~~3n^2 + 4n + 2~~
              Insignificant for large input sizes

<u>~</u> n^3  (n --> ∞)



Rule Big O:  
            drop lower order terms
            drop constant multiplier

T(n) = 17n^4 ~~+ 3n^3 + 4n + 8~~  = O(n4)
T(n) = 16n + lgn = On
Rule Running time
   ∑ Running time of all fragments
   
   int a;
   a=5 
   a++

loop
```
   for (i=0;i<n;i++) {       // O(n)
       //statement;
   }
```
simple statement fragment 1


nested loop 
```
   for (i=0;i<n;i++) {       // O(n)
            for (i=0;i<n;i++) {       // O(n2)
                //statement;  C
            }
   }
```
 n+1 (n+1*C)  O(n^2)


```   for (i=0;i<n;i++) {       // O(n)
                //statement;  C
            
   }
      for (i=0;i<n;i++) {       // O(n)
            for (i=0;i<n;i++) {       // O(n2)
                //statement;  C
            }
   }
```

t(n) ..... O(n^2)


## Conditional 
```py
if(condition) {
    for (i=0;i<n;i++) {           // O(n)
                //statement;       //C        
     } 
   } else {
      for (i=0;i<n;i++) {            // O(n)  
            for (i=0;i<n;i++) {       // O(n)
                //statement;         //C
            }
       }
   }                                 //O(n2)
```
Worst Case 0(n2)