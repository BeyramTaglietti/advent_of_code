
def solution():
  def part1():
    with open("day2/input.txt", "r") as f:
      lines = f.readlines()
      sm = 0
      for line in lines:
        row = [int(x.strip()) for x in line.split("\t")]
        sm += max(row) - min(row)
      
      print("sum", sm)

  def part2():
    with open("day2/input.txt", "r") as f:
      lines = f.readlines()
      sm = 0
      for line in lines:
        row = [int(x.strip()) for x in line.split("\t")]

        for j in range(len(row) - 1):
          for y in range(j + 1, len(row)):
            M, m = max(row[j], row[y]), min(row[j], row[y])
            if M % m == 0:
              sm += M // m
      
      print("sum", sm)

  part2()


solution()