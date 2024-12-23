import numpy as np

patrol_matrix =[]

with open("input-day6-ex.txt") as input_file:
    patrol_matrix = [ [c for c in l.strip()] for l in input_file]

directions = {
    '^': (-1, 0),
    '>': (0, 1),
    'v': (1, 0),
    '<': (0, -1)
}

def right_turn(dir: str) -> str:
    if( dir == '^'): return '>'
    if( dir == '>'): return 'v'
    if( dir == 'v'): return '<'
    if( dir == '<'): return '^'

# 0: done
#-1: turn
# 1: step
def check_next(x: int, y: int, d: tuple[int, int]) -> tuple[int, int, int]:
    nx, ny = d
    nx += x
    ny += y

    if  nx < 0 or ny < 0:
        return 0, nx, ny
    if len(patrol_matrix)-1 < ny or len(patrol_matrix[0])-1 < nx:
        return 0, nx, ny
    
    if patrol_matrix[nx][ny] == '#':
        return -1, x, y
    
    return 1, nx, ny

x = 0
y = 0
for ix, l in enumerate(patrol_matrix):
    for iy, c in enumerate(l):
        if not c in directions.keys():
            continue
        x = ix
        y = iy
        break
    
dir = patrol_matrix[x][y]
is_patroling = True
action, nx, ny = check_next(x, y, directions[dir])
while not action == 0:
    #patrol_matrix[nx][ny] = 'X'
    if action == 1:
        if dir == '<' or dir == '>':
            patrol_matrix[nx][ny] = '-'
        if dir == '^' or dir == 'v':
            patrol_matrix[nx][ny] = '|'

    if action == -1:
        patrol_matrix[nx][ny] = '+'
        dir = right_turn(dir)

    action, nx, ny = check_next(nx, ny, directions[dir])
    
cnt = sum([line.count('X') for line in patrol_matrix])
print(cnt)

np.savetxt("save.txt", patrol_matrix, '%s')