function getNumericValue(substring: string): number[] {
  const NUMS = new Map<string, number>([
    ["one", 1],
    ["two", 2],
    ["three", 3],
    ["four", 4],
    ["five", 5],
    ["six", 6],
    ["seven", 7],
    ["eight", 8],
    ["nine", 9],
  ]);

  const result: [number, number][] = []; // index, value

  for (let i = 0; i < substring.length; i++) {
    const char = substring[i]!;
    if (char >= "0" && char <= "9") {
      result.push([i, Number(char)]);
      continue;
    }

    for (const [word, value] of NUMS.entries()) {
      if (substring.slice(i, i + word.length) === word) {
        result.push([i, value]);
        break;
      }
    }
  }

  result.sort((a, b) => a[0] - b[0]);
  return result.map((x) => x[1]);
}

export function solveP1(values: string[]): number {
  let result = 0;

  for (const val of values) {
    let currentValue = [];
    for (const letter of val) {
      if (!isNaN(Number(letter))) currentValue.push(Number(letter));
    }

    result += Number(
      [currentValue[0]!, currentValue[currentValue.length - 1]!].join("")
    );
  }

  return result;
}
export function solveP2(values: string[]): number {
  let result = 0;

  for (const val of values) {
    const currentValue = getNumericValue(val);

    result += Number(
      [currentValue[0]!, currentValue[currentValue.length - 1]!].join("")
    );
  }

  return result;
}
