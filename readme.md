----------
PART 1
----------
To run the application, cd into the cmd directory.
$ cd ./cmd

Then run the main file.
$ go run main.go 

Update the connection string in main.go with the local PostgreSQL address/credentials as needed.

NOTE: As I'm entirely new to Go, I was only able to get the basic functionality working in the alloted time. Therefore, there is no concurrency, unit tests, and very minimal error handling. I'd never written anything in Go prior to today, so it was a wild ride.


----------
PART 2
----------
1. Poor error handling
 - There is no error handling for opening the DB connection, querying the DB, scanning rows, or parsing requests.
 - Unhandled errors could result in erratic behavior, skewed telemetry, or application downtime.
 - These can all be fixed with a similar check that exists on line 56 of the file.

2. WaitGroup as a global variable
 - The wg variable is used as a global variable.
 - Long running DB operations will negatively affect the other/faster functions, resulting in slowed performance or deadlocks.
 - Individual channels can be used for each endpoint.

3. No API response
 - No response is returned to the user when actions are complete.
 - If no feedback is provided to the user for creation of a user, this could result in multiple attempts and an influx of requests.
 - For each request, send a response indicating the status of the user's request.

----------
PART 3
----------

 See system-design.pdf