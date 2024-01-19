import sympy as s;print(sum(int(n%sum(map(int,str(n)))==0 and s.isprime(n//sum(map(int,str(n))))) for n in range(2,100000000)))
