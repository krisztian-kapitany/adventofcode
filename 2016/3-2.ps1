[cmdletbinding()]
Param()

$cluestrs = gc "clue3.txt"
$cnt = 0
$good_order = @()

for ($i = 0; $i -lt $cluestrs.Count; $i+=3) {
    $sidelengths0 = @()
     $cluestrs[$i] -split "\s+" | foreach { $sidelengths0 += [int]$_ }

    $sidelengths1 = @()
     $cluestrs[$i+1] -split "\s+" | foreach { $sidelengths1 += [int]$_ }

    $sidelengths2 = @()
     $cluestrs[$i+2] -split "\s+" | foreach { $sidelengths2 += [int]$_ }

    $good_order += $sidelengths0[1]
    $good_order += $sidelengths1[1]
    $good_order += $sidelengths2[1]

    $good_order += $sidelengths0[2]
    $good_order += $sidelengths1[2]
    $good_order += $sidelengths2[2]

    $good_order += $sidelengths0[3]
    $good_order += $sidelengths1[3]
    $good_order += $sidelengths2[3]
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