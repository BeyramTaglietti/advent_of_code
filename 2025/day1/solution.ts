type direction = "L" | "R";

export function solveP1(instructions: Array<string>): number {
  let currentPosition = 50;
  let result = 0;
  for (const instruction of instructions) {
    const [direction, ticks] = parseInstruction(instruction);
    const [newPosition] = rotateDial(currentPosition, ticks, direction);
    currentPosition = newPosition;
    if (currentPosition === 0) result++;
  }

  return result;
}

export function solveP2(instructions: Array<string>): number {
  let currentPosition = 50;
  let result = 0;
  for (const instruction of instructions) {
    const [direction, ticks] = parseInstruction(instruction);
    const [newPosition, timesCrossedZero] = rotateDial(
      currentPosition,
      ticks,
      direction
    );
    currentPosition = newPosition;
    result += timesCrossedZero;
  }

  return result;
}

const parseInstruction = (instruction: string): [direction, number] => {
  return [instruction[0]! as direction, Number(instruction.slice(1))];
};

const rotateDial = (
  currentValue: number,
  ticks: number,
  direction: direction
): [number, number] => {
  if (direction === "R") {
    const endPosition = currentValue + ticks;
    const timesHitZero =
      Math.floor(endPosition / 100) - Math.floor(currentValue / 100);
    return [endPosition % 100, timesHitZero];
  } else {
    const endPosition = currentValue - ticks;
    const timesHitZero =
      Math.ceil(currentValue / 100) - Math.ceil(endPosition / 100);

    let wrappedPosition = endPosition % 100;
    if (wrappedPosition < 0) {
      wrappedPosition += 100;
    }
    return [wrappedPosition, timesHitZero];
  }
};

const input = await Bun.file("./input.txt").text();
console.log("solution", solveP2(input.split("\n")));
