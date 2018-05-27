## [0.1.1] - 2018-05-27
### Added
- New field process_at for Transactions
- RAML API documentation
- Log exception data
- Catch runtime errors
- Budget removal webservice

### Changed
- Transaction type field is now a string

## [0.1.0] - 2018-04-09
### Added
- Sector creation
- Exceptions handling
- Budget slug

### Changed
- Packages architecture
- Transactions are now contained in budget document
- Sectors are now an array field of budget containing transactions
- Budgets and sectors are now identified by their slug

## [0.0.2] - 2017-09-15
### Added
- Budget creation
- Budgets list
- Budget details
- Mongo database name configuration

### Changed
- Transactions are now bound to a budget
- Transactions amount are now float type
- Transaction "Date" field is now "CreatedAt"

## [0.0.1] - 2017-09-14
### Added
- API Docker container configuration
- MongoDB connection configuration
- Transactions list
- Transaction show
- Transaction creation
