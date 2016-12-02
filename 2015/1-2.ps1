$cluestr = gc "clue1.txt"

$result = 0
$i = 0

$cluestr -split '(?<=.)(?=.)' | foreach {
  if($_ -eq '(') { $result = $result+1 }
  else { $result = $result-1 }

  $i = $i + 1
  if($result -eq -1) {
    Write-Host $i
    exit 0
  }
}

Write-Host $result
