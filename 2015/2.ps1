 [cmdletbinding()]
 Param()

$clues = gc "clue2.txt"
$sum = 0

foreach($clue in $clues) {
  Write-Verbose "Clue: $clue"

  $dimensions = @()
  $clue.split("x") | foreach { $dimensions += [int]$_ }
  
  $dimensions = $dimensions | Sort-Object

  $smallest = $dimensions | select -first 2
  Write-Verbose "Smallest 2: $($smallest[0]), $($smallest[1])"
  Write-Verbose "------------------"

  $surface = 2 * (($dimensions[0] * $dimensions[1]) + ($dimensions[0] * $dimensions[2]) + ($dimensions[1] * $dimensions[2]))
  $slack = $smallest[0] * $smallest[1]

  $volume = $dimensions[0] * $dimensions[1] * $dimensions[2]
  $wrap = 2 * ($smallest[0] + $smallest[1])

  $sum = $sum + $surface + $slack
  $wrapsum = $wrapsum + $wrap + $volume

  
}

Write-Host "Sum: $sum"
Write-Host "Wrap Sum: $wrapsum"