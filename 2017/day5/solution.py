
def solution():
  def part1():
    with open("day5/input.txt", "r") as f:
      lines = f.readlines()

      instructions = [int(line) for line in lines]

      currentIndex = 0
      steps = 0

      while currentIndex < len(lines):
        idx = currentIndex
        currentIndex += instructions[currentIndex]
        instructions[idx] += 1
        steps += 1

      print("exited after", steps, "steps")

  def part2():
    with open("day5/input.txt", "r") as f:
      lines = f.readlines()

      instructions = [int(line) for line in lines]

      currentIndex = 0
      steps = 0

      while currentIndex < len(lines):
        idx = currentIndex
        currentIndex += instructions[currentIndex]
        if instructions[idx] >= 3:
          instructions[idx] -= 1
        else:
          instructions[idx] += 1
        steps += 1

      print("exited after", steps, "steps")
      
  part2()


solution()
