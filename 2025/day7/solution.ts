export function solveP1(lines: Array<string>): number {
  const M = lines.length;
  const splitters = new Set<string>();
  let beams = new Set<string>();

  for (let y = 0; y < M; y++) {
    for (let x = 0; x < lines[y]!.length; x++) {
      const cell = lines[y]![x]!;
      if (cell === "^") splitters.add(`${y}#${x}`);
      if (cell === "S") beams.add(`${y}#${x}`);
    }
  }

  let splits = 0;

  for (let y = 1; y < M; y++) {
    const nextBeams = new Set<string>();

    for (const beam of beams) {
      const [, x] = beam.split("#").map(Number);
      const nextPos = `${y}#${x}`;

      if (splitters.has(nextPos)) {
        nextBeams.add(`${y}#${x! - 1}`);
        nextBeams.add(`${y}#${x! + 1}`);
        splits++;
      } else {
        nextBeams.add(nextPos);
      }
    }

    beams = nextBeams;
  }

  return splits;
}

export function solveP2(lines: Array<string>): number {
  const M = lines.length;
  const splitters = new Set<string>();
  let startY = 0,
    startX = 0;

  for (let y = 0; y < M; y++) {
    for (let x = 0; x < lines[y]!.length; x++) {
      const cell = lines[y]![x]!;
      if (cell === "^") splitters.add(`${y}#${x}`);
      if (cell === "S") {
        startY = y;
        startX = x;
      }
    }
  }

  const CACHE = new Map<string, number>();
  const dfs = (y: number, x: number): number => {
    const newY = y + 1;
    const key = `${newY}#${x}`;

    if (newY === M) return 1;
    if (CACHE.has(key)) return CACHE.get(key)!;

    if (splitters.has(key)) {
      const result = dfs(newY, x - 1) + dfs(newY, x + 1);
      CACHE.set(key, result);
      return result;
    }

    return dfs(newY, x);
  };

  return dfs(startY, startX);
}

const input = await Bun.file("./input.txt").text();
console.log("solution 1", solveP1(input.split("\n")));
console.log("solution 2", solveP2(input.split("\n")));
