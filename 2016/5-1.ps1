$clue = "ffykfhsq"

$md5 = new-object -TypeName System.Security.Cryptography.MD5CryptoServiceProvider
$utf8 = new-object -TypeName System.Text.UTF8Encoding

$results = @()

for ($i = 0; $results.Count -lt 8; $i++) {
  $probe = $clue + $i

  $hash = [System.BitConverter]::ToString($md5.ComputeHash($utf8.GetBytes($probe)))

  $hash = $hash -replace "-",""

  if($hash -like "00000*"){
    Write-host $hash.substring(5,1)
    $results += $hash
  }
}

Write-Host "======================"
Write-Host "Result: $($results -join '')" 



