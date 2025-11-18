// Fuel required to launch a given module is based on its mass. Specifically, to find the fuel required for a module, take its mass, divide by three, round down, and subtract 2.

export async function solveP1() {
  const text = await Bun.file("./day1/input.txt").text();

  let total = 0;

  for (let line of text.split("\n")) {
    total += fuelPerModule(Number(line));
  }

  return total;
}

export async function solveP2() {
  const text = await Bun.file("./day1/input.txt").text();

  let total = 0;

  for (let line of text.split("\n")) {
    total += fuelPerModule(Number(line), true);
  }

  return total;
}

const CACHE = new Map<number, number>();

const fuelPerModule = (module: number, recurrent: boolean = false): number => {
  if (CACHE.has(module)) {
    return CACHE.get(module)!;
  }

  let result = Math.floor(module / 3) - 2;

  if (recurrent) {
    if (result > 0) {
      result += fuelPerModule(result);
    } else {
      result -= result;
    }
  }

  CACHE.set(module, result);

  return result;
};
