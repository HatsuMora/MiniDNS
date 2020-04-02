# MiniDNS
This is my first Go project; on the way to learn Go, it is a DNS server

## Installation
Build it and run it

## Features
- Reply to an A request based on a json.zone file

### Upcoming updates (Not implemented yet)
- Answer any kind of DNS request
    - Answer proper error responses instead of panic
- Query foreign resolver if answer is not preset
    - cache for this
- Management UI (WebUI)
- Support for secondary dns server 
- Zone transfer
- Maintenance queries

## Technical debt
Error handling and unit test are not being prioritize at the moment, looking forward if anyone can put some attention in 
how to do real error handling for MiniDNS. 

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License
[MIT](https://choosealicense.com/licenses/mit/)
