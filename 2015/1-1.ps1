$cluestr = gc "clue1.txt"

$result = 0

$cluestr -split '(?<=.)(?=.)' | foreach {
  if($_ -eq '(') { $result = $result+1 }
  else { $result = $result-1 }
}

Write-Host $result
