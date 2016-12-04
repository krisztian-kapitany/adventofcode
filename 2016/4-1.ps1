[cmdletbinding()]
Param()

$cluestrs = gc "clue4.txt"
$cnt = 0

foreach($line in $cluestrs) {
    Write-Verbose "--------------------"

    $sections = $line -split "-"
    [string]$last = $sections | select -last 1

    $encoded = $sections | select -SkipLast 1
 
    $sectorid = [int]($last.Substring($last.IndexOf("[")-3, 3))
    $checksum = $last.Substring($last.IndexOf("[")+1, 5)

    $most_common = $encoded -split '(?<=.)(?=.)' | Sort-Object | Group-Object `
        | Sort-Object @{Expression = "Count"; Descending = $True}, @{Expression = "Name"; Descending = $False} `
        | Select-Object -Property Name -ExpandProperty Name -First 5
    
    $calc_checksum = ([string]$most_common) -replace " ",""

    if($calc_checksum -eq $checksum) { $cnt += $sectorid }

    Write-Verbose "encoded: $encoded"
    Write-Verbose "sectorid: $sectorid"
    Write-Verbose "checksum: $checksum"
    Write-Verbose "checksum: $calc_checksum"
    Write-Verbose "summ: $cnt"
    
}

Write-Host "Sum of real ids: $cnt"