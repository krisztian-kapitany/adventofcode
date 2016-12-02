 [cmdletbinding()]
 Param()

function step() {
  param(
    $point,
    [string]$dir
  )
  $ret = @{'x' = $point.x; 'y' = $point.y}

  switch($dir) {
    "U" { if($point.x -ne 1) { $ret.x = $point.x+1 }}
    "R" { if($point.y -ne 1) { $ret.y = $point.y+1 }}
    "D" { if($point.x -ne -1) { $ret.x = $point.x-1 }}
    "L" { if($point.y -ne -1) { $ret.y = $point.y-1 }}
    default {
      Write-Error "No such dir"
    }
  }
  return $ret
}

$cluestr = gc "clue2.txt"
$position = @{'x' = 0; 'y' = 0}

foreach ($clue in $cluestr) {

  $clue -split '(?<=.)(?=.)' | foreach {
    Write-Verbose "Char: $_"

    $position = step -point $position -dir $_

  }

  Write-Host "Step: $($position.x) , $($position.y)"
}



