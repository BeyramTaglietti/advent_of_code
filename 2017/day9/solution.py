
def solution():
  def part1():
    with open("day9/input.txt", "r") as f:
      lines = f.readlines()

      input = lines[0]

      def cleanup(s: str) -> str:
        res = ""

        insideGarbage = False
        i = 0
        while i < len(s):

          match s[i]:
            case ">":
              insideGarbage = False
            case "<": 
              insideGarbage = True
            case "!":
              i += 1
            case _:
              if not insideGarbage:
                res += s[i]
          
          i += 1
          
        return res
      
      cleanedUpInput = cleanup(input)

      print("cleaned up input", cleanedUpInput)

      level = 0
      score = 0
      for i in range(len(cleanedUpInput)):
        match cleanedUpInput[i]:
          case "{":
            level += 1
          case "}":
            level -= 1
            score += level + 1


      print("score", score)

  def part2():
    with open("day9/input.txt", "r") as f:
      lines = f.readlines()

      input = lines[0]

      def cleanup(s: str) -> int:
        insideGarbage = False
        cleaned = 0
        i = 0
        while i < len(s):

          if insideGarbage:
            match s[i]:
              case "!":
                i += 1
              case ">":
                insideGarbage = False
              case _:
                cleaned += 1
          else:
            match s[i]:
              case "<":
                insideGarbage = True

          i += 1
          
        return cleaned
      
      cleanedUpInput = cleanup(input)

      print("cleaned up input removed", cleanedUpInput, "characters")

  part2()


solution()
