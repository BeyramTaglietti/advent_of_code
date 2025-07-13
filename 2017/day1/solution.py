
def solution():
  def part1():
    with open("day1/input.txt", "r") as f:
      lines = f.readlines()

      inputLine = lines[0]

      inputLine += inputLine[0]

      sm = 0
      for i in range(len(inputLine) - 1):
        if inputLine[i] == inputLine[i + 1]:
          sm += int(inputLine[i])
      
      print("sum", sm)
    
  def part2():
    with open("day1/input.txt", "r") as f:
      lines = f.readlines()

      inputLine = lines[0]
      n = len(inputLine) // 2

      inputLine += inputLine[:n]

      print("new input line", inputLine)

      sm = 0
      for i in range(len(inputLine) - n):
        if inputLine[i] == inputLine[i + n]:
          sm += int(inputLine[i])

      print("sum", sm)

      """
      sm = 0
      for i in range(len(inputLine)):
        if inputLine[i] == inputLine[i + 1]:
          sm += int(inputLine[i])
      
      print("sum", sm) 
      """
  
  part2()



solution()