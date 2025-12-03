export function solveP1(banks: Array<string>): number {
  return banks.reduce((acc, item) => {
    const largest = findLargestVoltage(item, 2);
    acc += largest;
    return acc;
  }, 0);
}

export function solveP2(banks: Array<string>): number {
  return banks.reduce((acc, item) => {
    const largest = findLargestVoltage(item, 12);
    acc += largest;
    return acc;
  }, 0);
}

function findLargestVoltage(bank: string, digits: number): number {
  let drop = bank.length - digits;
  const stack: string[] = [];

  for (const digit of bank) {
    while (drop > 0 && stack.length > 0 && stack[stack.length - 1]! < digit) {
      stack.pop();
      drop--;
    }
    stack.push(digit);
  }

  return Number(stack.slice(0, digits).join(""));
}

const input = await Bun.file("./input.txt").text();
console.log("solution 1", solveP1(input.split("\n")));
console.log("solution 2", solveP2(input.split("\n")));
