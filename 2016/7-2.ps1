$cluestr = gc "clue7.txt"
$cnt = 0;

function isABA() {
  param(
    [Parameter(Mandatory=$true)]
    [ValidateScript({$_.Length -eq 3})]
    [string] $string
  )

  [char[]]$char_array = $string.ToCharArray()

  if( ($char_array[0] -eq $char_array[2]) -and ($char_array[0] -ne $char_array[1]) -and !($char_array -contains "\]") -and !($char_array -contains "\[") ) {
    return $true
  }

  return $false
}

foreach ($clue in $cluestr) {
  $ssl = $false

  $aba_map = @{}
  $bab_map = @{}

  for ($i = 0; $i -lt $clue.Length-2; $i++) {
    if(isABA -string $clue.substring($i,3)) {
      if(([regex]::Matches($clue.substring(0,$i), "\[" )).Count -gt ([regex]::Matches($clue.substring(0,$i), ']' )).Count) {
        $bab_map.Add($clue.substring($i,3),$i)
      }
      else {
        $aba_map.Add($clue.substring($i,3),$i)
      }
    }
  }

  foreach($aba in $aba_map.Keys){
    $bab_probe = "$($aba.substring(1,1))$($aba.substring(0,1))$($aba.substring(1,1))"

    if($bab_map.ContainsKey($bab_probe)) {
      $cnt++

      $aba_index = $aba_map[$aba]
      $bab_index = $bab_map[$bab_probe]

      $first_index = $aba_index
      $second_index = $bab_index

      if($aba_index -gt $bab_index) {
        $first_index = $bab_index
        $second_index = $aba_index
      }

      Write-Host "$cnt : " -NoNewline
      Write-Host $clue.substring(0,$first_index) -NoNewline
      Write-Host $clue.substring($first_index,3) -NoNewline -ForegroundColor "Yellow"
      Write-Host $clue.substring($first_index+3,$second_index-$first_index-3) -NoNewline
      Write-Host $clue.substring($second_index,3) -NoNewline -ForegroundColor "Yellow"
      Write-Host $clue.substring($second_index+3,$clue.Length-$second_index-3)
      
      break
    }
  }
}
Write-Host "======================="
Write-Host $cnt

