import math

def solution():
  def part1():
    input = 361527

    # find the ring number: y x y where y is the ring size y^2

    # the bottom right corner would be of value y^2

    # ring_level = (ring_size - 1) // 2

    # min(abs(num - m) for m in midpoints) 

    s = math.ceil(math.sqrt(input))
    if s % 2 == 0:
        s += 1
    ring_level = (s - 1) // 2
    max_val = s ** 2
    side_len = s - 1
    
    midpoints = [max_val - side_len//2 - side_len*i for i in range(4)]

    offset = min(abs(input - m) for m in midpoints)

    print("distance", ring_level + offset)

  def part2():
    input_val = 361527

    directions = [(1,0), (0,1), (-1,0), (0,-1)]
    dir_idx = 0
    x = y = 0
    grid = {(0,0): 1}

    step_size = 1
    while True:
        for _ in range(2):
            dx, dy = directions[dir_idx]
            for _ in range(step_size):
                x += dx
                y += dy

                neighbors = [(x + nx, y + ny) for nx in [-1,0,1] for ny in [-1,0,1] if not (nx == 0 and ny == 0)]
                val = sum(grid.get(pos, 0) for pos in neighbors)
                grid[(x,y)] = val

                if val > input_val:
                    print("val", val)
                    return

            dir_idx = (dir_idx + 1) % 4
        step_size += 1

  part2()


solution()