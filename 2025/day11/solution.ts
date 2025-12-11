class Node {
  children: Node[];
  value: string;

  constructor(value: string, children: Node[]) {
    this.value = value;
    this.children = children;
  }
}

export function solveP1(lines: Array<string>): number {
  const treeMap = buildTree(lines);

  const startingPoint = treeMap.get("you")!;

  /**
   *
   * @param node
   * @returns amount of roads that reach the "out" point
   */
  const dfs = (node: Node): number => {
    if (node.value === "out") return 1;
    if (node.children.length < 1) return 0;

    return node.children.reduce((acc, item) => acc + dfs(item), 0);
  };

  const nodesThatReachTheEnd = dfs(startingPoint);

  return nodesThatReachTheEnd;
}

export function solveP2(lines: Array<string>): number {
  const treeMap = buildTree(lines);

  const startingPoint = treeMap.get("svr")!;

  const stringifyKey = (
    node: Node,
    visitedNodes: Record<string, boolean>
  ): string => {
    return `${node.value}#${Object.entries(visitedNodes)
      .map(([key, value]) => `${key}:${value}`)
      .join(",")}`;
  };

  const CACHE: Record<string, number> = {};
  const dfs = (node: Node, nodesToVisit: Record<string, boolean>): number => {
    const stringifiedKey = stringifyKey(node, nodesToVisit);
    if (stringifiedKey in CACHE) {
      return CACHE[stringifiedKey]!;
    }
    if (node.value === "out") {
      if (Object.values(nodesToVisit).every((x) => x === true)) return 1;
      else return 0;
    }

    if (node.children.length < 1) return 0;

    if (nodesToVisit[node.value] !== undefined) {
      nodesToVisit[node.value] = true;
    }

    let total = 0;

    for (let i = 0; i < node.children.length; i++) {
      total += dfs(node.children[i]!, { ...nodesToVisit });
    }

    CACHE[stringifiedKey] = total;

    return total;
  };

  const nodesThatReachTheEnd = dfs(startingPoint, {
    dac: false,
    fft: false,
  });

  return nodesThatReachTheEnd;
}

const buildTree = (lines: string[]): Map<string, Node> => {
  const nodes = new Map<string, Node>();

  for (let i = 0; i < lines.length; i++) {
    const [parentString, childrenStrings] = parseLine(lines[i]!);

    const children: Node[] = childrenStrings.map((x) => {
      if (nodes.has(x)) {
        return nodes.get(x)!;
      }

      const node = new Node(x, []);
      nodes.set(x, node);
      return node;
    });

    const parent: Node = nodes.has(parentString)
      ? nodes.get(parentString)!
      : new Node(parentString, []);
    nodes.set(parentString, parent);

    parent.children.push(...children);
  }

  return nodes;
};

const parseLine = (line: string): [string, string[]] => {
  const [parent, rest] = line.split(":");

  const children = rest?.split(" ").filter((x) => x !== "")!;

  return [parent!, children];
};

const input = await Bun.file("./input.txt").text();
console.log("solution 1", solveP1(input.split("\n")));
console.log("solution 2", solveP2(input.split("\n")));
