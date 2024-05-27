@echo build provider
go build -o .

pause
@copy to local provider

move terraform-provider-viettelidc.exe C:\Users\tacjl\AppData\Roaming\terraform.d\plugins\terraform.local\local\viettelidc\1.0.0\windows_amd64

