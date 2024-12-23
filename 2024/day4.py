xmass_matrix =[]

with open("input-day4.txt") as input_file:
    xmass_matrix = [ [c for c in l.strip()] for l in input_file]

directions = {
    'n': (-1, 0),
    'ne': (-1, 1),
    'e': (0, 1),
    'se': (1, 1),
    's': (1, 0),
    'sw': (1, -1),
    'w': (0, -1),
    'nw': (-1, -1)
}

def check_next(x: int, y: int, d: tuple[int, int], part: str) -> bool:
    nx, ny = d
    nx += x
    ny += y

    if  nx < 0 or ny < 0:
        return False
    if len(xmass_matrix)-1 < ny or len(xmass_matrix[0])-1 < nx:
        return False
    
    part += xmass_matrix[nx][ny]

    if part == 'XMAS':
        return True
    
    if 'XMAS'.startswith(part):
        return check_next(nx, ny, d, part)

    return False
   

xmass_count = 0
for ix, l in enumerate(xmass_matrix):
    for iy, c in enumerate(l):
        if not c == 'X':
            continue

        for dir in directions.keys():
            if check_next(ix, iy, directions[dir], 'X'):
                xmass_count += 1

print(xmass_count)

# Part 2
def is_mas(a: str, b: str):
    if a == 'M' and b == 'S':
        return True
    if a == 'S' and b == 'M':
        return True
    return False

def get_cell(x: int, y: int, d: tuple[int,int]) -> str:
    nx, ny = d
    return xmass_matrix[x+nx][y+ny] 

def check_cross(x: int, y: int):
    a = get_cell(x, y, directions['nw'])
    b = get_cell(x, y, directions['se'])

    c = get_cell(x, y, directions['ne'])
    d = get_cell(x, y, directions['sw'])

    if is_mas(a, b) and is_mas(c, d):
        return True

    return False

x_mass_count = 0
for ix in range(1, len(xmass_matrix[0])-1):
    for iy in range(1, len(xmass_matrix)-1):
        if not xmass_matrix[ix][iy] == 'A':
            continue

        if check_cross(ix, iy):
            x_mass_count += 1

print(x_mass_count)