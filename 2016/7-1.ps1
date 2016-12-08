$cluestr = gc "clue7.txt"
$cnt = 0;

function isABBA() {
  param(
    [Parameter(Mandatory=$true)]
    [ValidateScript({$_.Length -eq 4})]
    [string] $string
  )

  [char[]]$char_array = $string.ToCharArray()

  if( ($char_array[1] -eq $char_array[2]) -and ($char_array[0] -ne $char_array[1]) -and ($char_array[0] -eq $char_array[3])) {
    return $true
  }
  
  return $false
}

foreach ($clue in $cluestr) {
  $abba = $false

  for ($i = 0; $i -lt $clue.Length-3; $i++) {
    if(isABBA -string $clue.substring($i,4)) {
      if(([regex]::Matches($clue.substring(0,$i), "\[" )).Count -gt ([regex]::Matches($clue.substring(0,$i), ']' )).Count) {
        $abba = $false
        break
      }

      $abba = $true
    }

  }

  if($abba) {
    $cnt++
  }
  
}

Write-Host $cnt

