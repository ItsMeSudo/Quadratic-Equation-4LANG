const prompt = require('prompt-sync')();

const a_val = prompt("Add meg az A értékét: ")
const b_val = prompt("Add meg az B értékét: ")
const c_val = prompt("Add meg az C értékét: ")
console.log("A értéke: " + a_val + ", B értéke: " + b_val + ", C értéke: " + c_val)
const diszkriminans = ((b_val*b_val)-(4*a_val*c_val))
console.log("A diszkrimináns értéke: "+ diszkriminans)

if (diszkriminans < 0){
    console.log("Az egyenletnek nincs megoldása a valós számok halmazán!")
    process.exit(0)
} else if (diszkriminans == 0){
    console.log("Az egyenletnek két egyenlő (kettős gyöke) van!")
} else if (diszkriminans > 0){
    console.log("Az egyenletnek két különböző valós gyöke van!")
} else {
    console.log("HIBA")
    process.exit(0)
}

const res_1 = (-b_val) + Math.sqrt(diszkriminans)
const res_2 = (-b_val) - Math.sqrt(diszkriminans)
const result_1 = res_1 / (2*a_val)
const result_2 = res_2 / (2*a_val)

console.log("X1 értéke:", result_1)
console.log("X2 értéke:", result_2)
console.log("SUDO#0814 2022")