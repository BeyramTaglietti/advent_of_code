type range = [number, number];

export function solveP1(lines: Array<string>): number {
  let ranges: Array<range> = [];

  for (const line of lines) {
    const r = parseLine(line);
    ranges = [...ranges, ...r];
  }

  let total = 0;

  for (const [rangeStart, rangeEnd] of ranges) {
    for (let start = rangeStart; start < rangeEnd + 1; start++) {
      const isValid = isValidID(start);
      if (!isValid) {
        total += start;
      }
    }
  }

  return total;
}

export function solveP2(lines: Array<string>): number {
  let ranges: Array<range> = [];

  for (const line of lines) {
    const r = parseLine(line);
    ranges = [...ranges, ...r];
  }

  let total = 0;

  for (const [rangeStart, rangeEnd] of ranges) {
    for (let start = rangeStart; start < rangeEnd + 1; start++) {
      const isValid = isValidID2(start);
      if (!isValid) {
        total += start;
      }
    }
  }

  return total;
}

const parseLine = (line: string): Array<range> => {
  const ranges = line.split(",");
  return ranges.map((x) => {
    const s = x.split("-").map((x) => Number(x));
    return [s[0]!, s[1]!];
  });
};

export const isValidID = (value: number): boolean => {
  const sValue = value.toString();

  const half = sValue.length / 2;

  return sValue.slice(0, half) !== sValue.slice(half, sValue.length);
};

export const isValidID2 = (value: number): boolean => {
  const sValue = value.toString();
  const doubled = sValue + sValue;
  const pos = doubled.indexOf(sValue, 1);

  // Valid if the pattern is NOT found before the end (i.e., not repeating)
  return pos >= sValue.length;
};

const input = await Bun.file("./input.txt").text();
console.log("solution", solveP2(input.split("\n")));
