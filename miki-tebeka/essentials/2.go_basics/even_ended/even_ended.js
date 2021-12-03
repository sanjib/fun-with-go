console.time("start")

const min = 1000
const max = 9999

let count = 0

for (let i = min; i <= max; i++) {
  for (let j = i; j <= max; j++) {
    const n = i * j
    const s = String(n)
    if (s[0] === s[s.length-1]) {
      count++
    }
  }
}

console.log("found", count, "even-ended numbers")
console.timeEnd("start")
