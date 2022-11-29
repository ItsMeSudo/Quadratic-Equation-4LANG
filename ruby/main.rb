puts("Add meg az A értékét: ")
a_val = gets.to_i
puts("Add meg az B értékét: ")
b_val = gets.to_i
puts("Add meg az C értékét: ")
c_val = gets.to_i

puts("A értéke: " + a_val.to_s  + ", B értéke: "+ b_val.to_s  + ", C értéke: " + c_val.to_s )
a_val.to_i
b_val.to_i
c_val.to_i

diszkriminans = ((b_val*b_val)-(4*a_val*c_val))
puts("A diszkrimináns értéke: "+ diszkriminans.to_s)

if diszkriminans < 0 
    puts("Az egyenletnek nincs megoldása a valós számok halmazán!")
    exit(false)
elsif diszkriminans == 0
    puts("Az egyenletnek két egyenlő (kettős gyöke) van!")
elsif diszkriminans > 0
    puts("Az egyenletnek két különböző valós gyöke van!")
else 
    puts("HIBA")
    exit(false)
end

res_1 = (-b_val) + Math.sqrt(diszkriminans)
res_2 = (-b_val) - Math.sqrt(diszkriminans)
result_1 = res_1 / (2*a_val)
result_2 = res_2 / (2*a_val)

puts("X1 értéke: " + result_1.to_s)
puts("X2 értéke: " + result_2.to_s)
puts("SUDO#0814 2022")