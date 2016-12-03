[cmdletbinding()]
Param()

$cluestrs = gc "clue3.txt"
$cnt = 0

foreach($clueline in $cluestrs) {

    $sides = @()
    $clueline -split "\s+" | foreach { $sides += [int]$_ }
    $sides = $sides | sort

    $longest = $sides | select -Last 1

    $isTriangle = $(0 -lt ((($sides[1]) + $($sides[2])) - $longest))

    if($isTriangle) {
        $cnt += 1 
    }

    Write-Verbose "------------------------------"
    Write-Verbose "triangle: $($sides[1]),$($sides[2]),$($sides[3])"
    Write-Verbose "triangle: $longest"
    Write-Verbose "triange?: $isTriangle"
    Write-Verbose "cnt     : $cnt"

    #if($cnt -eq 10) {exit 0}

}

Write-Host "Valid triangles: $cnt"