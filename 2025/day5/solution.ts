type range = [number, number];

export function solveP1(lines: Array<string>): number {
  const [ranges, ingredients] = parseInput(lines);
  const orderedRanges = orderRanges(ranges);
  const mergedRanges = mergeRanges(orderedRanges);

  const freshIngredients = [];

  for (const ingredient of ingredients) {
    for (const [start, end] of mergedRanges) {
      if (ingredient >= start && ingredient <= end) {
        freshIngredients.push(ingredient);
        break;
      }
    }
  }

  return freshIngredients.length;
}

export function solveP2(lines: Array<string>): number {
  const [ranges] = parseInput(lines);
  const orderedRanges = orderRanges(ranges);
  const mergedRanges = mergeRanges(orderedRanges);

  let total = 0;
  for (const [start, end] of mergedRanges) {
    total += end - start + 1;
  }

  return total;
}

const parseInput = (lines: Array<string>): [Array<range>, Array<number>] => {
  const ranges: Array<range> = [];
  const ingredients: Array<number> = [];

  let foundEmptyLine = false;
  for (let i = 0; i < lines.length; i++) {
    if (foundEmptyLine) {
      ingredients.push(Number(lines[i]!));
    } else {
      if (lines[i]! === "") {
        foundEmptyLine = true;
      } else {
        ranges.push(lines[i]!.split("-").map((x) => Number(x)) as range);
      }
    }
  }
  return [ranges, ingredients];
};

const orderRanges = (ranges: Array<range>): Array<range> => {
  const orderedRanges = ranges.slice().sort((a, b) => {
    if (a[0] === b[0]) {
      return a[1] - b[1];
    }

    return a[0] - b[0];
  });

  return orderedRanges;
};

const mergeRanges = (ranges: Array<range>): Array<range> => {
  const sortedInterval = ranges.slice().sort((a, b) => a[0] - b[0]);
  const output: Array<range> = [];
  output.push(sortedInterval[0]!);

  for (const interval of sortedInterval) {
    const start = interval[0]!;
    const end = interval[1]!;
    const lastEnd = output[output.length - 1]![1]!;

    if (start <= lastEnd) {
      output[output.length - 1]![1] = Math.max(lastEnd, end);
    } else {
      output.push([start, end]);
    }
  }
  return output;
};

const input = await Bun.file("./input.txt").text();
console.log("solution 1", solveP1(input.split("\n")));
console.log("solution 2", solveP2(input.split("\n")));
