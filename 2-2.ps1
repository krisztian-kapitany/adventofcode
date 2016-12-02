 [cmdletbinding()]
 Param()

function vec_abs() {
  param(
    $vector
  )
  return $([math]::abs($vector.x) + [math]::abs($vector.y))
}

function step() {
  param(
    $point,
    [string]$dir
  )

  $ret = @{'x' = $point.x; 'y' = $point.y}

  switch($dir) {
    "U" { $ret.x = $point.x+1 }
    "R" { $ret.y = $point.y+1 }
    "D" { $ret.x = $point.x-1 }
    "L" { $ret.y = $point.y-1 }
    default {
      Write-Error "No such dir"
    }
  }

  if((vec_abs $ret) -le 2) { return $ret }
  else { return $point }
}

$cluestr = gc "clue2.txt"

$position = @{'x' = 0; 'y' = -2}



foreach ($clue in $cluestr) {

  $clue -split '(?<=.)(?=.)' | foreach {
    Write-Verbose "Char: $_"

    $position = step -point $position -dir $_

  }

  Write-Host "Step: $($position.x) , $($position.y)"
}



Write-Host "$($point.x) , $($point.y)"
Write-Host "direction= $direction"
Write-Host "----------------------------"
Write-Host "Distance: $([math]::abs($point.x) + [math]::abs($point.y))"



