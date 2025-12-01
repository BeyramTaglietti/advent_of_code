function isPickValid(pick: string, limits: Record<string, number>): boolean {
  const cubes = pick.split(", ");

  return cubes.every((cube) => {
    const [count, color] = cube.split(" ");
    return Number(count) <= limits[color!]!;
  });
}

function parseLine(line: string): number {
  return Number(line.split(":")[0]?.split(" ")[1]);
}

export function solveP1(lines: string[]): number {
  const limits = { red: 12, green: 13, blue: 14 };

  return lines.reduce((sum, line) => {
    if (!line.trim()) return sum;

    const [_, rest] = line.split(": ");
    const picks = rest?.split("; ");

    const valid = picks?.every((pick) => isPickValid(pick, limits));
    return valid ? sum + parseLine(line) : sum;
  }, 0);
}

export function solveP2() {}
