[cmdletbinding()]
Param()

$cluestrs = gc "clue4.txt"
$cnt = 0

foreach($line in $cluestrs) {
    Write-Verbose "--------------------"

    $sections = $line -split "-"
    [string]$last = $sections | select -last 1

    [string]$encoded = $sections | select -SkipLast 1
 
    $sectorid = [int]($last.Substring($last.IndexOf("[")-3, 3))

    #$most_common = $encoded -split '(?<=.)(?=.)'

    Write-Verbose "encoded: $encoded"
    Write-Verbose "sectorid: $sectorid"

    # 122-97+1
    $shift = $sectorid % 26
    [char[]]$decoded_array = ""

    foreach($letter in $encoded -split '(?<=.)(?=.)') {
        
        $ascii_d = [int][char]$letter

        # 32 = space
        if($ascii_d -ne 32) { 
            $ascii_d += $shift 
            # 97-122 (97==a, 122==z)
            if($ascii_d -gt 122) {$ascii_d -= 26}
            if($ascii_d -lt 97) {$ascii_d += 26}
        }

        $decoded_array += [char]$ascii_d 
    }  

    $decoded = $decoded_array -join ''
    Write-Verbose "decoded: $($decoded)"

    if($decoded -like "*north*pole*") {
        Write-Host "encoded: $encoded"
        Write-Host "sectorid: $sectorid"
        Write-Host "decoded: $($decoded)"
        exit 0
    }   
}

