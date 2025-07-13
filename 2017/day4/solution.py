
def solution():
  def part1():
    with open("day4/input.txt", "r") as f:
      
      lines = f.readlines()

      valids = 0

      for line in lines:
        line = line.strip()

        phrases = line.split(" ")

        s = set()
        invalid = False
        for phrase in phrases:
          if phrase in s:
            invalid = True
            break
          else:
            s.add(phrase)

        if not invalid:
          valids += 1


      print("found", valids, "valid passwords")

  def part2():
    with open("day4/input.txt", "r") as f:
      
      lines = f.readlines()

      valids = 0

      for line in lines:
        line = line.strip()

        phrases = line.split(" ")

        s = set()
        invalid = False
        for phrase in phrases:
          sortedPhrase = ''.join(sorted(phrase))
          if sortedPhrase in s:
            invalid = True
            break
          else:
            s.add(sortedPhrase)

        if not invalid:
          valids += 1


      print("found", valids, "valid passwords")

  part2()


solution()
