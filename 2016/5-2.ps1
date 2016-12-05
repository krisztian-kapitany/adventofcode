$clue = "ffykfhsq"

$md5 = new-object -TypeName System.Security.Cryptography.MD5CryptoServiceProvider
$utf8 = new-object -TypeName System.Text.UTF8Encoding

[char[]]$results = "-","-","-","-","-","-","-","-"

Write-Host "Hacking..."

$errorActionPreference = 'SilentlyContinue'

for ($i = 0; $results -contains "-"; $i++) {
  $probe = $clue + $i

  $hash = [System.BitConverter]::ToString($md5.ComputeHash($utf8.GetBytes($probe)))
  $hash = $hash -replace "-",""

  if(($hash -like "00000*") -and ( [int]($hash.substring(5,1)) -le 7 ) ) {
    $index = [int]($hash.substring(5,1))

    if($results[$index] -eq "-") {
      $hash
      $results[$index] = $hash.substring(6,1)

      $results -join ""
    }
  }
}

Write-Host "======================"
Write-Host "Result: $($results -join '')"



