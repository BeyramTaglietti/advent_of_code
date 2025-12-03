export function solveP1(lines: Array<string>): number {
  return 0;
}

export function solveP2(lines: Array<string>): number {
  return 0;
}

const input = await Bun.file("./input.txt").text();
console.log("solution 1", solveP1(input.split("\n")));
console.log("solution 2", solveP2(input.split("\n")));
