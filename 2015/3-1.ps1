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

$position = @{ 'x' = 0; 'y'= 0 }

$visited = @()
$visited += $position

foreach($direction in $cluestr -split '(?<=.)(?=.)') {
  $next_pos = @{ 'x' = $position.x; 'y'= $position.y }

  switch($direction){
    "^" {$next_pos.y += 1}
    "v" {$next_pos.y -= 1}
    ">" {$next_pos.x += 1}
    "<" {$next_pos.x -= 1}
    default { Write-Error "No such direction"}
  }
  $position = $next_pos

  if( !(contains_pos -positions_array $visited -new_position $next_pos) ){
    $visited += $next_pos
  }
}

Write-Host "========================="
Write-Host "Different places visited: #$($visited.Count)"