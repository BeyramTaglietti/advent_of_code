class Point {
  y = 0;
  x = 0;

  constructor(y: number, x: number) {
    this.y = y;
    this.x = x;
  }

  stringify(): string {
    return `${this.y}#${this.x}`;
  }

  static parse(s: string): Point {
    const [y, x] = s.split("#");
    return new Point(Number(y!), Number(x!));
  }
}

export function solveP1(lines: Array<string>): number {
  const grid = new Map<string, boolean>();

  const m = lines.length,
    n = lines[0]!.length;

  for (let y = 0; y < m; y++) {
    for (let x = 0; x < n; x++) {
      const p = lines[y]![x]!;
      if (p === "@") {
        const point = new Point(y, x);
        grid.set(point.stringify(), true);
      }
    }
  }

  let canBeAccessed = 0;
  for (const [key] of grid) {
    if (countNeighbours(Point.parse(key), grid) < 4) canBeAccessed++;
  }

  return canBeAccessed;
}

export function solveP2(lines: Array<string>): number {
  const grid = new Map<string, boolean>();

  const m = lines.length,
    n = lines[0]!.length;

  for (let y = 0; y < m; y++) {
    for (let x = 0; x < n; x++) {
      const p = lines[y]![x]!;
      if (p === "@") {
        const point = new Point(y, x);
        grid.set(point.stringify(), true);
      }
    }
  }

  const candidates = new Set<string>();
  for (const [key] of grid) {
    if (countNeighbours(Point.parse(key), grid) < 4) {
      candidates.add(key);
    }
  }

  let canBeAccessed = 0;

  while (candidates.size > 0) {
    const key = candidates.values().next().value!;
    candidates.delete(key);

    if (!grid.has(key)) continue;

    const point = Point.parse(key);

    if (countNeighbours(point, grid) >= 4) continue;

    grid.delete(key);
    canBeAccessed++;

    for (const [dy, dx] of directions) {
      const neighbor = new Point(point.y + dy!, point.x + dx!);
      const neighborKey = neighbor.stringify();
      if (grid.has(neighborKey)) {
        candidates.add(neighborKey);
      }
    }
  }

  return canBeAccessed;
}

const directions = [
  [-1, -1],
  [-1, 0],
  [-1, 1],
  [0, -1],
  [0, 1],
  [1, -1],
  [1, 0],
  [1, 1],
];

const countNeighbours = (point: Point, grid: Map<string, boolean>): number => {
  let counter = 0;

  for (const direction of directions) {
    const [dy, dx] = direction;
    const [ny, nx] = [point.y + dy!, point.x + dx!];
    const nextPoint = new Point(ny, nx);

    if (grid.has(nextPoint.stringify())) counter++;
  }

  return counter;
};

const input = await Bun.file("./input.txt").text();
console.log("solution 1", solveP1(input.split("\n")));
console.log("solution 2", solveP2(input.split("\n")));
