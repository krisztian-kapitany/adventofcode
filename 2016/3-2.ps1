[cmdletbinding()]
Param()

$cluestrs = gc "clue3.txt"
$cnt = 0
$good_order = @()

for ($i = 0; $i -lt 3; $i++) {

    $cluestrs | foreach {
        $numbers = @()
        $_ -split "\s+" | foreach { $numbers += [int]$_ }

        $good_order += $numbers[$i+1]
    }
}

for ($i = 0; $i -lt $good_order.Count; $i+=3) {
    
    $sides = @()
    $sides += $good_order[$i]
    $sides += $good_order[$i+1]
    $sides += $good_order[$i+2]

    $sides = $sides | sort

    $longest = $sides | select -Last 1

    $isTriangle = $(0 -lt ((($sides[0]) + $($sides[1])) - $longest))

    if($isTriangle) {
        $cnt += 1 
    }

    Write-Verbose "------------------------------"
    Write-Verbose "triangle: $($sides[0]),$($sides[1]),$($sides[2])"
    Write-Verbose "triangle: $longest"
    Write-Verbose "triange?: $isTriangle"
    Write-Verbose "cnt     : $cnt"

}

Write-Host "Valid triangles: $cnt"