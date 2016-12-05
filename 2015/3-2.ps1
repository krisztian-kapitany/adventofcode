function contains_pos(){
  param(
    $positions_array,
    $new_position
  )

  foreach($position in $positions_array) {
    if(($position.x -eq $new_position.x) -and ($position.y -eq $new_position.y)) {
      return $true
    }
  }
  return $false
}

$cluestr = gc "clue3.txt"

$position_s = @{ 'x' = 0; 'y'= 0 }
$position_r = @{ 'x' = 0; 'y'= 0 }

$visited = @()
$visited += $position_s

$directions = $cluestr -split '(?<=.)(?=.)'

for ($i = 0; $i -lt $directions.Count; $i++) {

  if(($i % 2) -eq 0) {
    $next_pos = @{ 'x' = $position_s.x; 'y'= $position_s.y }
  }
  else {
    $next_pos = @{ 'x' = $position_r.x; 'y'= $position_r.y }
  }
  
  switch($directions[$i]){
    "^" {$next_pos.y += 1}
    "v" {$next_pos.y -= 1}
    ">" {$next_pos.x += 1}
    "<" {$next_pos.x -= 1}
    default { Write-Error "No such direction"}
  }

  if(($i % 2) -eq 0) {
    $position_s = $next_pos
  }
  else {
    $position_r = $next_pos
  }

  if( !(contains_pos -positions_array $visited -new_position $next_pos) ){
    $visited += $next_pos
  }
}

Write-Host "========================="
Write-Host "Different places visited: #$($visited.Count)"