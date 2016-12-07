$clues = gc "clue6.txt"

$result = @()

for ($i = 0; $i -lt 8; $i++) {
  $result += $clues | Group-Object -Property { $_.substring($i,1) } | Sort-Object Count -Descending | Select-Object Name -First 1
}

Write-Host "Decoded: $($result.Name -join '')"
