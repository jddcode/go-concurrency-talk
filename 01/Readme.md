# One

In our first application we are passing incoming requests to our two downstream systems, *SlowSyS* and 
*MediumSyS*. It is our scneario that it is necessary to pass requests on to these systems. You can make a GET
request to `/` on port `8080` to test the application - your request will eventually complete after 90 seconds.