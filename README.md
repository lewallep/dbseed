# DBSeed generates random data for a database

This program designed to create randomized data inside of a database for testing purposes.  The application which was the catalyst for this project was DBBeagle.  Currently the only database supported is MSSSQL.  The intention is to use the same concept with 

### Command line arguments
- **append** Accepts a boolean.  If append is set to false this will erase the existing database data
- **rowsToAdd** Accepts an integer.  Appends the specified amount of rows to the database.
- **numTables** Accepts an integer.  Specifies how many random tables to add. 
- **minCols** Accepts an integer.  Specifies the minimum amount of columns a table will contain.
- **maxCols** Accepts an integer.  Specifies the maximum amount of columns a table will contain.
