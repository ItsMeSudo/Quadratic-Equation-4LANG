import math

a_val = float(input("Add meg az A értékét: "))
b_val = float(input("Add meg az B értékét: "))
c_val = float(input("Add meg az C értékét: "))

print("A értéke: " + str(a_val) + ", B értéke: " + str(b_val) + ", C értéke: " + str(c_val))
diszkriminans = ((b_val*b_val)-(4*a_val*c_val))
print("A diszkrimináns értéke: "+ str(diszkriminans))

if diszkriminans < 0:
    print("Az egyenletnek nincs megoldása a valós számok halmazán!")
    quit()
elif diszkriminans == 0:
    print("Az egyenletnek két egyenlő (kettős gyöke) van!")
elif diszkriminans > 0:
    print("Az egyenletnek két különböző valós gyöke van!")
else:
    print("HIBA")
    quit()

res_1 = (-b_val) + math.sqrt(diszkriminans)
res_2 = (-b_val) - math.sqrt(diszkriminans)
result_1 = res_1 / (2*a_val)
result_2 = res_2 / (2*a_val)

print("X1 értéke:", result_1)
print("X2 értéke:", result_2)
print("SUDO#0814 2022")