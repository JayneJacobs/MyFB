# Space Complexity


Space complexity is a measure of how efficient your code is in terms of memory used.
Space complexity analysis happens almost in the same way time complexity analysis happens.

For example, consider the following code :

```vector<int> V;

for (int i = 0; i < N; i++) V.push_back(i);
```


The code snippet ends up creating a vector of size N. So, space complexity of the code is O(N).

Additional space / memory is measured in terms of:
       - the largest memory use by the program when it runs. 

That is to say, if you allocated O(N) memory, and later free it, that does not make the space complexity of your program O(1).