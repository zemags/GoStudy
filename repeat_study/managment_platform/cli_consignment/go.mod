module github.com/zemags/GoStudy/repeat_study/managment_platform/cli_consignment

go 1.15

replace github.com/zemags/GoStudy/repeat_study/managment_platform/service_consignment => ../service_consignment

require (
	github.com/zemags/GoStudy/repeat_study/managment_platform/service_consignment v0.0.0-00010101000000-000000000000
	golang.org/x/net v0.0.0-20210405180319-a5a99cb37ef4 // indirect
	golang.org/x/sys v0.0.0-20210403161142-5e06dd20ab57 // indirect
	golang.org/x/text v0.3.6 // indirect
	google.golang.org/genproto v0.0.0-20210406143921-e86de6bf7a46 // indirect
	google.golang.org/grpc v1.37.0
)
