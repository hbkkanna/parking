# parking lot

parkinglot.go is the main module in the system, it creates parking lot system 
based on input config , parking lot type and slot count of each vehicle type are configurable with tariff models listed below . 

### Requirement:
go 1.16 and above.

### How to test ?
parkinglot_test.go has the test cases for all type of parking lot , will be able to run individual test using "go test -run ^TestStadiumParkingLot$ github.com/hbkkanna/parking  -v" command  
1. TestMallParkingLot
2. TestStadiumParkingLot 
3. TestAirportParkingLot

### Tariff Models :
Tariff model are the basic units of tariff calculator, these models are listed to calculate based on the parking lot requirements . 
* EveryHour - Hourly price, calculates for every hour.
* EveryDay  - Daily price  
* HourInterval - Fixed price in the hour range.
* PreviousHourInterval - sums up all the previous hour tariff. 
* EveryHourInInterval - Hourly price in an interval. 

### Tariff Matcher : 
Parkinglot system uses tariff matcher to calculate the cost , matchers will have list of tariff models . 
* SingleTariffMatcher - matches with single model in the collection of models 
* MultipleTariffMatcher - sums up all the matches 


### Test run commands :
* go test ./... -v   
* go test  github.com/hbkkanna/parking  -v 
* go test -run ^TestStadiumParkingLot$ github.com/hbkkanna/parking  -v   


