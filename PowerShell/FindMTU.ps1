$dhost = "1.1.1.1"
if ($args.Length -gt 0) {
    $dhost = $args[0]
}
Write-Host "Finding MTU max:"
Write-Host "================"
$min = 68
$max = 9000

while ($min -le $max) {

    $mtu = [int]($min + $max) / 2
    ping -w 3 -f -n 1 -l $mtu $dhost > $null
    if ($?) {
        $min = $mtu + 1
        Write-Host -NoNewline "+"

    } else {
        $max = $mtu - 1
        Write-Host -NoNewline "-"

    }
}
$mtu = $max + 28
Write-Host "`r`nMaximum MTU value for $dhost : $mtu"