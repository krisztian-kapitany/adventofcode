function moveit() {
  param(
    $point,
    [int]$dir,
    [int]$dist
  )
  $ret = @{'x' = $point.x; 'y' = $point.y}

  switch($dir){
    0 { $ret.x = $point.x+$dist }
    1 { $ret.y = $point.y+$dist }
    2 { $ret.x = $point.x-$dist }
    3 { $ret.y = $point.y-$dist }
    default {
      Write-Error "No such dir"
    }
  }
  return $ret
}

function turn([int]$direction, $side) {
  [int]$ret = 0

  switch($side) {
    "L" {
      if($direction -eq 0) {
        $ret = 3
      }
      else {
        $ret = $direction - 1
      }    
    }
    "R" { $ret = ($direction+1)%4 }
    default { Write-Error "No such side" }
  }
  return $ret
}

$cluestr = "R1, L3, R5, R5, R5, L4, R5, R1, R2, L1, L1, R5, R1, L3, L5, L2, R4, L1, R4, R5, L3, R5, L1, R3, L5, R1, L2, R1, L5, L1, R1, R4, R1, L1, L3, R3, R5, L3, R4, L4, R5, L5, L1, L2, R4, R3, R3, L185, R3, R4, L5, L4, R48, R1, R2, L1, R1, L4, L4, R77, R5, L2, R192, R2, R5, L4, L5, L3, R2, L4, R1, L5, R5, R4, R1, R2, L3, R4, R4, L2, L4, L3, R5, R4, L2, L1, L3, R1, R5, R5, R2, L5, L2, L3, L4, R2, R1, L4, L1, R1, R5, R3, R3, R4, L1, L4, R1, L2, R3, L3, L2, L1, L2, L2, L1, L2, R3, R1, L4, R1, L1, L4, R1, L2, L5, R3, L5, L2, L2, L3, R1, L4, R1, R1, R2, L1, L4, L4, R2, R2, R2, R2, R5, R1, L1, L4, L5, R2, R4, L3, L5, R2, R3, L4, L1, R2, R3, R5, L2, L3, R3, R1, R3"

$point = @{'x' = 0; 'y' = 0}
$locations = @()
$locations += New-Object -TypeName PSObject -Property @{'x' = 0; 'y' = 0}

$direction = 0

$clues = $cluestr.split(',')

$clues | foreach {
  $clue = $_.Trim()
  Write-Verbose "$clue"

  $turn_to = $clue.Substring(0,1)
  $move_dist = [Convert]::ToInt32( $clue.Substring(1, $clue.Length-1))

  $direction = turn -direction $direction -side $turn_to

  for($i=1; $i -le $move_dist; $i++) {
    $step = moveit -point $step -dir $direction -dist 1

	# Comment out for 1-1
    foreach($location in $locations) {
      if ( ($location.x -eq $step.x) -and ($location.y -eq $step.y) ) {
        Write-Host "==========Match============="
        Write-Host "x= $($step.x)"
        Write-Host "y= $($step.y)"
        Write-Host "----------------------------"
        Write-Host "Distance: $([math]::abs($step.x) + [math]::abs($step.y))"
        exit 0
		}
	}

  Write-Verbose "Step: $($step.x) , $($step.y)"
  $locations += New-Object -TypeName PSObject -Property @{'x' = $step.x; 'y' = $step.y}

  }

  $point = moveit -point $point -dir $direction -dist $move_dist

}

Write-Host "$($point.x) , $($point.y)"
Write-Host "direction= $direction"
Write-Host "----------------------------"
Write-Host "Distance: $([math]::abs($point.x) + [math]::abs($point.y))"



