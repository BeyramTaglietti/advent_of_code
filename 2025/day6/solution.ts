export function solveP1(lines: Array<string>): number {
  let total = 0;

  const operations = lines[0]!.split(" ").length;

  const grid: Array<Array<number>> = Array.from(
    { length: operations - 1 },
    () => []
  );

  for (let i = 0; i < lines.length; i++) {
    let operationMembers = lines[i]!.split(" ").filter((x) => x !== "");

    if (i === lines.length - 1) {
      for (let j = 0; j < operationMembers.length; j++) {
        const multiplication = operationMembers[j]! === "*";
        const operationResult = grid[j]!.reduce(
          (acc, item) => {
            if (multiplication) {
              acc *= item;
            } else {
              acc += item;
            }

            return acc;
          },
          multiplication ? 1 : 0
        );

        total += operationResult;
      }
    } else {
      for (let j = 0; j < operationMembers.length; j++) {
        grid[j]!.push(Number(operationMembers[j]!));
      }
    }
  }

  return total;
}

export function solveP2(lines: Array<string>): number {
  const operations: Array<number> = [];
  const operationsLine = lines[lines.length - 1]!;

  for (let x = 0; x < operationsLine.length; x++) {
    if (operationsLine[x] === "+" || operationsLine[x] === "*") {
      operations.push(x);
    }
  }

  const members: Array<Array<string>> = Array.from(
    { length: lines.length - 1 },
    () => []
  );

  for (let y = 0; y < lines.length - 1; y++) {
    for (let o = 0; o < operations.length; o++) {
      if (o === operations.length - 1) {
        members[y]?.push(lines[y]!.slice(operations[o]));
      } else {
        members[y]?.push(
          lines[y]!.slice(operations[o], operations[o + 1]! - 1)
        );
      }
    }
  }

  const verticalNumbers: Array<Array<string>> = Array.from({
    length: operations.length,
  });

  for (let x = 0; x < members[0]!.length; x++) {
    const maxMembersLength = members[0]![x]!.length;
    const verticalMembers = Array.from({ length: maxMembersLength }, () =>
      Array.from({ length: members.length })
    );

    for (let y = 0; y < members.length; y++) {
      for (let d = 0; d < maxMembersLength; d++) {
        verticalMembers[d]![y] = members[y]![x]![d]!;
      }
    }

    verticalNumbers[x] = verticalMembers.map((x) =>
      x.filter((x) => x !== " ").join("")
    );
  }

  let total = 0;

  for (let i = 0; i < operations.length; i++) {
    const isMultiplication = lines[lines.length - 1]![operations[i]!] === "*";

    const operationResult = verticalNumbers[i]!.reduce(
      (acc, item) => {
        if (isMultiplication) {
          acc *= Number(item);
        } else {
          acc += Number(item);
        }

        return acc;
      },
      isMultiplication ? 1 : 0
    );

    total += operationResult;
  }

  return total;
}

const input = await Bun.file("./input.txt").text();
console.log("solution 1", solveP1(input.split("\n")));
console.log("solution 2", solveP2(input.split("\n")));
