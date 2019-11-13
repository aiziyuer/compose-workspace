# See: https://www.petri.com/using-nat-virtual-switch-hyper-v

If ("NAT-Switch" -in (Get-VMSwitch | Select-Object -ExpandProperty Name) -eq $FALSE) {
    'Creating Internal-only switch named "NAT-Switch" on Windows Hyper-V host...'

    New-VMSwitch -SwitchName "NAT-Switch" -SwitchType Internal

    New-NetIPAddress -IPAddress 192.168.0.1 -PrefixLength 24 -InterfaceAlias "vEthernet (NAT-Switch)"

    New-NetNAT -Name "NATNetwork" -InternalIPInterfaceAddressPrefix 192.168.0.0/24
}
else {
    '"NAT-Switch" for static IP configuration already exists; skipping'
}

If ("192.168.0.1" -in (Get-NetIPAddress | Select-Object -ExpandProperty IPAddress) -eq $FALSE) {
    'Registering new IP address 192.168.0.1 on Windows Hyper-V host...'

    New-NetIPAddress -IPAddress 192.168.0.1 -PrefixLength 24 -InterfaceAlias "vEthernet (NAT-Switch)"
}
else {
    '"192.168.0.1" for static IP configuration already registered; skipping'
}

If ("192.168.0.0/24" -in (Get-NetNAT | Select-Object -ExpandProperty InternalIPInterfaceAddressPrefix) -eq $FALSE) {
    'Registering new NAT adapter for 192.168.0.0/24 on Windows Hyper-V host...'

    New-NetNAT -Name "NATNetwork" -InternalIPInterfaceAddressPrefix 192.168.0.0/24
}
else {
    '"192.168.0.0/24" for static IP configuration already registered; skipping'
}