const generateId = (n, m) => {
  //Cantor pairing
  return ((n + m) * (n + m + 1)) / 2 + m;
};

function main() {
  console.log(generateId(0, 1));
  console.log(generateId(0, 0));
  console.log(generateId(1, 0));
  console.log(generateId(1, 1));
  console.log(generateId(1000, 1000));
}

main();
